package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/maxkuzn/advection-and-coagulation/internal/field1d/saver"

	"github.com/maxkuzn/advection-and-coagulation/algorithm/coagulation/fast"

	"github.com/maxkuzn/advection-and-coagulation/algorithm/coagulation/predcorr"

	"github.com/maxkuzn/advection-and-coagulation/internal/coagulator1d"
	"github.com/maxkuzn/advection-and-coagulation/internal/coagulator1d/naiveparallel"
	"github.com/maxkuzn/advection-and-coagulation/internal/coagulator1d/parallelpool"
	"github.com/maxkuzn/advection-and-coagulation/internal/coagulator1d/sequential"

	"github.com/pkg/profile"
	"github.com/schollz/progressbar/v3"

	"github.com/maxkuzn/advection-and-coagulation/algorithm/advector1d"
	"github.com/maxkuzn/advection-and-coagulation/algorithm/coagulation"
	"github.com/maxkuzn/advection-and-coagulation/algorithm/coagulation/kernel"
	"github.com/maxkuzn/advection-and-coagulation/config"
	"github.com/maxkuzn/advection-and-coagulation/internal/field1d"
)

const (
	configFilename  = "config.json"
	historyFilename = "data/history.txt"
)

var (
	profilers             = []string{"cpu", "mem", "goroutine"}
	profileNameToProfiler = map[string]func(*profile.Profile){
		"cpu":       profile.CPUProfile,
		"mem":       profile.MemProfile,
		"goroutine": profile.MemProfile,
	}

	profileFlag = flag.String("profile", "", fmt.Sprintf("which profiler to use\npossible values: %v", profilers))
)

func main() {
	flag.Parse()

	conf, err := config.Read(configFilename)
	if err != nil {
		log.Fatal(err)
	}

	saveFile, err := os.Create(historyFilename)
	if err != nil {
		log.Fatal(err)
	}

	s := saver.NewSaver(saveFile)
	defer func() {
		err := s.Flush()
		if err != nil {
			log.Fatal(err)
		}
	}()

	field, buff := initFields(conf)

	advector, err := newAdvector(conf)
	if err != nil {
		log.Fatal(err)
	}

	coagulator, err := newCoagulator(conf, field.Volumes())
	if err != nil {
		log.Fatal(err)
	}

	if err := coagulator.Start(); err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := coagulator.Stop(); err != nil {
			log.Fatal(err)
		}
	}()

	stop, err := profiler()
	if err != nil {
		log.Fatal(err)
	}
	defer stop()

	run(conf, field, buff, s, advector, coagulator)
}

func run(
	conf *config.Config,
	field, buff field1d.Field,
	saver saver.Saver,
	advector advector1d.Advector,
	coagulator coagulator1d.Coagulator,
) {
	start := time.Now()
	defer func() {
		diff := time.Since(start)
		fmt.Printf("\nTime: %.3f\n", diff.Seconds())
	}()

	err := saver.Save(field)
	if err != nil {
		log.Fatal(err)
	}

	bar := newBar(conf.TimeSteps)
	defer func() {
		err := bar.Finish()
		if err != nil {
			log.Fatal(err)
		}
	}()

	for t := 0; t < conf.TimeSteps; t++ {
		field, buff = advector.Process(field, buff)
		field, buff = coagulator.Process(field, buff)

		err = saver.Save(field)
		if err != nil {
			log.Fatal(err)
		}

		err = bar.Add(1)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func initFields(conf *config.Config) (field1d.Field, field1d.Field) {
	field := field1d.Init(conf.FieldCellsSize, conf.ParticlesSizesNum, conf.MinParticleSize, conf.MaxParticleSize)
	buff := field1d.New(conf.FieldCellsSize, conf.ParticlesSizesNum, conf.MinParticleSize, conf.MaxParticleSize)

	return field, buff
}

func newSaver(conf *config.Config, writeTo io.Writer) (saver.Saver, error) {
	switch conf.SaverName {
	case "Sync":
		return saver.NewSaver(writeTo), nil
	case "Async":
		// return saver.NewSaver(writeTo), nil
		return nil, fmt.Errorf("async saver is not implemented")
	case "-":
		return saver.NewMock(), nil
	default:
		return nil, fmt.Errorf("unknown saver name %q", conf.SaverName)
	}
}

func newAdvector(conf *config.Config) (advector1d.Advector, error) {
	switch conf.AdvectorName {
	case "CentralDifference":
		return advector1d.NewCentralDifference(conf.AdvectionCoef), nil
	default:
		return nil, fmt.Errorf("unknown advector name %q", conf.AdvectorName)
	}
}

func newCoagulator(conf *config.Config, volumes []float64) (coagulator1d.Coagulator, error) {
	var kern coagulation.Kernel
	switch conf.CoagulationKernelName {
	case "Identity":
		kern = kernel.NewIdentity()
	case "Addition":
		kern = kernel.NewAddition()
	case "Multiplication":
		kern = kernel.NewMultiplication()
	default:
		return nil, fmt.Errorf("unknown coagulation kernel name %q", conf.CoagulationKernelName)
	}

	var base coagulation.Coagulator
	switch conf.BaseCoagulatorName {
	case "PredictorCorrector":
		base = predcorr.New(kern, conf.TimeStep)
	case "Fast":
		var err error

		base, err = fast.New(kern, conf.TimeStep, volumes)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unknown base coagulation name %q", conf.BaseCoagulatorName)
	}

	switch conf.CoagulatorName {
	case "Sequential":
		return sequential.New(base), nil
	case "NaiveParallel":
		return naiveparallel.New(base, conf.CoagulatorBatchSize), nil
	case "ParallelPool":
		return parallelpool.New(base, conf.CoagulatorBatchSize), nil
	default:
		return nil, fmt.Errorf("unknown coagulation name %q", conf.CoagulatorName)
	}
}

func profiler() (stop func(), err error) {
	if profileFlag == nil || *profileFlag == "" {
		return func() {}, nil
	}

	f, ok := profileNameToProfiler[*profileFlag]
	if !ok {
		return func() {}, fmt.Errorf("unknown profiler name: %v", *profileFlag)
	}

	s := profile.Start(f, profile.ProfilePath("."))
	return func() {
		s.Stop()
	}, nil
}

func newBar(total int) *progressbar.ProgressBar {
	bar := progressbar.NewOptions(total,
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionSetWidth(40),
		progressbar.OptionSetDescription("Computing advection coagulation"),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]=[reset]",
			SaucerHead:    "[green]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}))

	return bar
}
