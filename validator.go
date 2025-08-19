package main

import (
	"errors"
	"flag"
	"log"
	"os"
	"strings"

	"github.com/ejdem86/photo_validator/pkg/download"
	"github.com/ejdem86/photo_validator/pkg/validator"
	"github.com/ejdem86/photo_validator/pkg/walker"
)

var (
	sourceDirectory = flag.String("s", "", "Where to download the backup image from?")
	fname           = flag.String("f", "", "Path to the image")
	prefix          = flag.String("p", "", "Prefix to ignore in the source path")
	confirm         = flag.Bool("a", true, "Ask to confirm the replacement")
)

func main() {
	flag.Parse()

	if *fname == "" {
		log.Fatal("Usage: photo_validator -f <image_file>")
	}

	filesToCopy := walker.SingleWalker(*fname)
	if i, err := os.Stat(*fname); err == nil && i.IsDir() {
		filesToCopy = walker.Dir(*fname)
		log.Printf("target is a directory, num of files: %d\n", len(filesToCopy))
	}

	recreatedFiles := 0
	for _, fileToCheck := range filesToCopy {
		// Validate the image data
		var err error
		if strings.HasSuffix(strings.ToLower(fileToCheck), ".raf") {
			err = validator.ValidateRaw(fileToCheck)
		} else {
			err = validator.ValidateJpeg(fileToCheck)
		}
		if err != nil {
			if *sourceDirectory == "" {
				log.Fatalf("Invalid image: %v", err)
			}

			backupFile := *sourceDirectory + strings.ReplaceAll(fileToCheck, *prefix, "")
			log.Printf("Attempt to download the file from backup: %s", backupFile)
			if err := download.Replace(fileToCheck, backupFile, *confirm); err != nil {
				if errors.Is(err, download.ErrSkip) {
					log.Println(err)
					continue
				}
				log.Fatalf("failed to replace file: %v", err)
			}
			recreatedFiles++
		}
	}

	log.Printf("Re-created %d out of %d files.\n", recreatedFiles, len(filesToCopy))
}
