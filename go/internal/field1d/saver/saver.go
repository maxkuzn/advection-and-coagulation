package saver

import "github.com/maxkuzn/advection-and-coagulation/internal/field1d"

type Saver interface {
	Save(f field1d.Field) error
	Flush() error
}
