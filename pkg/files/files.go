package files

import (
	"os"
	"path/filepath"
)

type DeepLevel int64

const (
	Endless DeepLevel = iota
	OneShot
)

func FindByExtension(searchPath string, formats []string, level DeepLevel) ([]string, error) {
	var files []string

	if level == Endless {
		err := filepath.Walk(searchPath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() {
				fileExt := filepath.Ext(path)
				for _, format := range formats {
					if fileExt == format {
						files = append(files, path)
						break
					}
				}

			}

			return nil
		})

		if err != nil {
			return nil, err
		}

	} else if level == OneShot {
		dirEntries, err := os.ReadDir(searchPath)
		if err != nil {
			return nil, err
		}

		for _, dirEntry := range dirEntries {

			fileExt := filepath.Ext(dirEntry.Name())
			for _, format := range formats {
				if fileExt == format {
					files = append(files, filepath.Join(searchPath, dirEntry.Name()))
					break
				}
			}
		}
	}

	return files, nil
}
