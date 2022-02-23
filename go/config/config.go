package config

import (
	"encoding/json"
	"errors"
	"os"
)

type Config struct {
	FieldCellsSize    int `json:"field_cells_size"`
	ParticlesSizesNum int `json:"particles_sizes_num"`

	TotalTime float64 `json:"total_time"`
	TimeSteps int     `json:"time_steps"`
	TimeStep  float64 `json:"-"`

	AdvectionCoef float64 `json:"advection_coef"`

	AdvectorName          string `json:"advector"`
	CoagulationKernelName string `json:"coagulation_kernel"`
}

func (c *Config) validateAndFill() error {
	if c.FieldCellsSize <= 0 {
		return errors.New("field cells size should be specified as positive")
	}

	if c.ParticlesSizesNum <= 0 {
		return errors.New("particles sizes num should be specified as positive")
	}

	if c.TotalTime <= 0 {
		return errors.New("total time should be specified as positive")
	}

	if c.TimeSteps <= 0 {
		return errors.New("time steps should be specified as positive")
	}

	if c.AdvectionCoef <= 0 {
		return errors.New("advection coefficient should be specified as positive")
	}

	if c.AdvectorName == "" {
		return errors.New("advection name should be specified")
	}

	if c.CoagulationKernelName == "" {
		return errors.New("coagulation kernel name should be specified")
	}

	c.TimeStep = c.TotalTime / float64(c.TimeSteps)
	return nil
}

func Read(filename string) (*Config, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	d := json.NewDecoder(f)
	d.DisallowUnknownFields()

	err = d.Decode(config)
	if err != nil {
		return nil, err
	}

	err = config.validateAndFill()
	if err != nil {
		return nil, err
	}

	return config, nil
}
