package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/pkg/profile"
	"github.com/schollz/progressbar/v3"

	"github.com/maxkuzn/advection-and-coagulation/algorithm/advector1d"
	"github.com/maxkuzn/advection-and-coagulation/algorithm/coagulator"
	"github.com/maxkuzn/advection-and-coagulation/algorithm/coagulator/kernel"
	"github.com/maxkuzn/advection-and-coagulation/algorithm/coagulator1d"
	"github.com/maxkuzn/advection-and-coagulation/algorithm/coagulator1d/naiveparallel"
	"github.com/maxkuzn/advection-and-coagulation/algorithm/coagulator1d/parallelpool"
	"github.com/maxkuzn/advection-and-coagulation/algorithm/coagulator1d/sequential"
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

	saver := field1d.NewSaver(saveFile)
	defer func() {
		err := saver.Flush()
		if err != nil {
			log.Fatal(err)
		}
	}()

	field, buff := initFields(conf)

	advector, err := newAdvector(conf)
	if err != nil {
		log.Fatal(err)
	}

	coagulator, err := newCoagulator(conf)
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

	run(conf, field, buff, saver, advector, coagulator)
}

func run(
	conf *config.Config,
	field, buff field1d.Field,
	saver *field1d.Saver,
	advector advector1d.Advector,
	coagulator coagulator1d.Coagulator,
) {
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

func newAdvector(conf *config.Config) (advector1d.Advector, error) {
	switch conf.AdvectorName {
	case "CentralDifference":
		return advector1d.NewCentralDifference(conf.AdvectionCoef), nil
	default:
		return nil, fmt.Errorf("unknown advector name %q", conf.AdvectorName)
	}
}

func newCoagulator(conf *config.Config) (coagulator1d.Coagulator, error) {
	var kern coagulator.Kernel
	switch conf.CoagulationKernelName {
	case "Identity":
		kern = kernel.NewIdentity()
	default:
		return nil, fmt.Errorf("unknown coagulation kernel name %q", conf.CoagulationKernelName)
	}

	base := coagulator.New(kern, conf.TimeStep)

	switch conf.CoagulatorName {
	case "Sequential":
		return sequential.New(base), nil
	case "NaiveParallel":
		return naiveparallel.New(base), nil
	case "ParallelPool":
		return parallelpool.New(base), nil
	default:
		return nil, fmt.Errorf("unknown coagulator name %q", conf.CoagulatorName)
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
