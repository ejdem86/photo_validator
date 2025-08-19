package validator

import (
	"fmt"
	"image/jpeg"
	"os"
)

func ValidateJpeg(fname string) error {
	f, err := os.Open(fname)
	if err != nil {
		return fmt.Errorf("failed to open the file: %w", err)
	}
	defer f.Close()

	photo, err := jpeg.Decode(f)
	if err != nil {
		return fmt.Errorf("failed to decode JPEG image: %w", err)
	}

	if photo == nil {
		return fmt.Errorf("decoded image is nil")
	}

	return nil
}
