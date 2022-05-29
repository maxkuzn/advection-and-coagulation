package saver

import (
	"bufio"
	"fmt"
	"io"

	"github.com/maxkuzn/advection-and-coagulation/internal/field1d"
)

type saver struct {
	w         *bufio.Writer
	firstSave bool
}

func NewSaver(w io.Writer) *saver {
	return &saver{
		w:         bufio.NewWriter(w),
		firstSave: true,
	}
}

func (s *saver) saveHeader(f field1d.Field) error {
	_, err := fmt.Fprintf(s.w, "%d\n", f.Len())
	return err
}

func (s *saver) Save(f field1d.Field) error {
	if s.firstSave {
		err := s.saveHeader(f)
		if err != nil {
			return err
		}

		s.firstSave = false
	}

	for i := 0; i < f.Len(); i++ {
		c := f.Cell(i)
		for j, currentSize := range c {
			format := "%v "
			if j+1 == len(c) {
				format = "%v\n"
			}

			_, err := fmt.Fprintf(s.w, format, currentSize)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *saver) Flush() error {
	return s.w.Flush()
}
