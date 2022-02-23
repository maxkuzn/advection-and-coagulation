package advector1d

import "github.com/maxkuzn/advection-and-coagulation/internal/field1d"

type Advector interface {
	Process(field, buff field1d.Field) (field1d.Field, field1d.Field)
}
