package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/pkg/profile"
	"github.com/schollz/progressbar/v3"

	"github.com/maxkuzn/advection-and-coagulation/algorithm/advector1d"
	"github.com/maxkuzn/advection-and-coagulation/algorithm/coagulator1d"
	"github.com/maxkuzn/advection-and-coagulation/algorithm/coagulator1d/kernel1d"
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

	field := field1d.New(conf.FieldCellsSize, conf.ParticlesSizesNum)
	buff := field1d.New(conf.FieldCellsSize, conf.ParticlesSizesNum)

	advector, err := newAdvector(conf)
	if err != nil {
		log.Fatal(err)
	}

	coagulator, err := newCoagulator(conf)
	if err != nil {
		log.Fatal(err)
	}

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

func newAdvector(conf *config.Config) (advector1d.Advector, error) {
	switch conf.AdvectorName {
	case "CentralDifference":
		return advector1d.NewCentralDifference(conf.AdvectionCoef), nil
	default:
		return nil, fmt.Errorf("unknown advector name %q", conf.AdvectorName)
	}
}

func newCoagulator(conf *config.Config) (coagulator1d.Coagulator, error) {
	var kernel coagulator1d.Kernel
	switch conf.CoagulationKernelName {
	case "Identity":
		kernel = kernel1d.NewIdentity()
	default:
		return nil, fmt.Errorf("unknown coagulation kernel name %q", conf.CoagulationKernelName)
	}

	switch conf.CoagulatorName {
	case "Sequential":
		return coagulator1d.NewSequential(kernel, conf.TimeStep), nil
	case "Parallel":
		return coagulator1d.NewParallel(kernel, conf.TimeStep), nil
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
