package field1d

import (
	"bufio"
	"fmt"
	"io"
)

type Saver struct {
	w         *bufio.Writer
	firstSave bool
}

func NewSaver(w io.Writer) *Saver {
	return &Saver{
		w:         bufio.NewWriter(w),
		firstSave: true,
	}
}

func (s *Saver) saveHeader(f Field) error {
	_, err := fmt.Fprintf(s.w, "%d\n", f.Len())
	return err
}

func (s *Saver) Save(f Field) error {
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

func (s *Saver) Flush() error {
	return s.w.Flush()
}