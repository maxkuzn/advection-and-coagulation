package coagulator1d

import "github.com/maxkuzn/advection-and-coagulation/internal/field1d"

type Coagulator interface {
	Process(field, buff field1d.Field) (field1d.Field, field1d.Field)
	Start() error
	Stop() error
}
