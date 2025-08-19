package validator

import (
	"bytes"
	"fmt"
	"image/jpeg"

	"github.com/kladd/raw"
)

func ValidateRaw(fname string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("invalid file: %s", fname)
		}
	}()
	raf := raw.ReadRAF(fname)

	if raf == nil {
		return fmt.Errorf("failed to read RAF format")
	}

	photo, err := jpeg.Decode(bytes.NewReader(raf.Jpeg))
	if err != nil {
		return fmt.Errorf("failed to decode JPEG image: %w", err)
	}

	if photo == nil {
		return fmt.Errorf("decoded image is nil")
	}

	return nil
}
