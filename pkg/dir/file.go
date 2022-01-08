package dir

import (
	"io/fs"
	"log"
	"os"
	"path"
	"path/filepath"
)

// FindFiles with ext with os package
func FindFiles(dir, ext string, recursive bool) []os.DirEntry {
	slice := make([]os.DirEntry, 0)
	files, err := os.ReadDir(dir)
	if err != nil {
		return slice
	}

	for _, f := range files {
		if f.IsDir() {
			slice = append(slice, FindFiles(dir+"/"+f.Name(), ext, recursive)...)
		}
		if path.Ext(f.Name()) == ext {
			slice = append(slice, f)
		}
	}
	return slice
}

// Files list all files in dir and subdirs with/out ext
func Files(fsys fs.FS, ext string) []fs.DirEntry {
	slice := make([]fs.DirEntry, 0)
	err := fs.WalkDir(fsys, ".", func(p string, d fs.DirEntry, err error) error {
		if d.IsDir() || ext != "" && filepath.Ext(p) != ext {
			return nil
		}

		slice = append(slice, d)
		return nil
	})
	if err != nil {
		log.Printf("fs.WalkDir returns error %v", err)
	}
	return slice
}
