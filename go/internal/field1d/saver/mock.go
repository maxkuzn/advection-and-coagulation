package saver

import "github.com/maxkuzn/advection-and-coagulation/internal/field1d"

func NewMock() *mockSaver {
	return &mockSaver{}
}

type mockSaver struct{}

func (s *mockSaver) Save(f field1d.Field) error {
	return nil
}

func (s *mockSaver) Flush() error {
	return nil
}
