package walker

import (
	"fmt"
	"io/fs"
	"os"
	"path"
	"strings"
)

func Dir(p string) (resp []string) {
	root, err := os.OpenRoot(p)
	if err != nil {
		panic(err)
	}

	if err := fs.WalkDir(root.FS(), ".", func(cp string, d fs.DirEntry, err error) error {
		fmt.Println("walking...", cp)
		if d != nil {
			if d.IsDir() {
				return nil
			}
		}
		lowPath := strings.ToLower(cp)
		if strings.HasSuffix(lowPath, ".jpg") || strings.HasSuffix(lowPath, ".raf") {
			resp = append(resp, path.Join(p, cp))
		}
		return nil
	}); err != nil {
		panic(err)
	}

	return
}
