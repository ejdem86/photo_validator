package download

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/ejdem86/photo_validator/pkg/validator"
)

var (
	ErrCancel = errors.New("user cancelled")
	ErrSkip   = errors.New("file skipped")
)

func Replace(to, from string, ask bool) error {
	if ask {
		fmt.Printf("About to replace %s with %s. Ok? [Y]es/[N]o/[S]kip: ", to, from)
		var decision string
		fmt.Scanln(&decision)
		switch strings.ToLower(decision) {
		case "y", "yes":
		case "s", "skip":
			return nil
		default:
			return ErrCancel
		}
	}

	if strings.HasSuffix(strings.ToLower(from), ".raf") {
		if err := validator.ValidateRaw(from); err != nil {
			return fmt.Errorf("source is invalid: %w", err)
		}
	} else {
		if err := validator.ValidateJpeg(from); err != nil {
			return fmt.Errorf("source is invalid: %w", err)
		}
	}

	destination, err := os.OpenFile(to, os.O_WRONLY, 0660)
	if err != nil {
		return fmt.Errorf("failed to open destination for write: %w", err)
	}
	defer destination.Close()

	source, err := os.OpenFile(from, os.O_RDONLY, 0600)
	if err != nil {
		return fmt.Errorf("failed to open destination for read: %w", err)
	}

	sourceData, err := io.ReadAll(source)
	if err != nil {
		return fmt.Errorf("failed to read source: %w", err)
	}

	wroteN, err := destination.Write(sourceData)
	if err != nil {
		return fmt.Errorf("failed to write to the destination file: %w", err)
	}
	if wroteN != len(sourceData) {
		return fmt.Errorf("wrote %d, but expected %d", wroteN, len(sourceData))
	}

	fmt.Printf("Wrote new file: %d\n", wroteN)
	return nil
}
