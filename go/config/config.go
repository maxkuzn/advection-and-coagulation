package config

import (
	"encoding/json"
	"errors"
	"os"
)

type Config struct {
	FieldSize         float64 `json:"field_size"`
	FieldCellsSize    int     `json:"field_cells_size"`
	ParticlesSizesNum int     `json:"particles_sizes_num"`
	MinParticleSize   float64 `json:"min_particle_size"`
	MaxParticleSize   float64 `json:"max_particle_size"`

	TotalTime float64 `json:"total_time"`
	TimeSteps int     `json:"time_steps"`
	TimeStep  float64 `json:"-"`

	AdvectionCoef float64 `json:"advection_coef"`

	SaverName             string `json:"saver"`
	AdvectorName          string `json:"advector"`
	BaseCoagulatorName    string `json:"base_coagulator"`
	CoagulatorName        string `json:"coagulator"`
	CoagulationKernelName string `json:"coagulation_kernel"`
}

func (c *Config) validateAndFill() error {
	if c.FieldSize <= 0 {
		return errors.New("field size should be specified as positive")
	}

	if c.FieldCellsSize <= 0 {
		return errors.New("field cells size should be specified as positive")
	}

	if c.ParticlesSizesNum <= 0 {
		return errors.New("particles sizes num should be specified as positive")
	}

	if c.MinParticleSize <= 0 {
		return errors.New("min particle size should be specified as positive")
	}

	if c.MaxParticleSize <= 0 {
		return errors.New("max particle size should be specified as positive")
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

	if c.SaverName == "" {
		return errors.New("saver name should be specified")
	}

	if c.AdvectorName == "" {
		return errors.New("advection name should be specified")
	}

	if c.BaseCoagulatorName == "" {
		return errors.New("base coagulation name should be specified")
	}

	if c.CoagulatorName == "" {
		return errors.New("coagulation name should be specified")
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
