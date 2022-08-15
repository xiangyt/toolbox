package file

import (
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"regexp"
)

func BatchRename(regexpStr, pathStr, fileNameSuffix, season string, recursion bool) {
	//fmt.Println(os.Getwd())
	reg := regexp.MustCompile(regexpStr)
	ext := "." + fileNameSuffix
	repl := fmt.Sprintf("[${1}]${2} S%sE${3}", season)
	filepath.Walk(pathStr, func(p string, info fs.FileInfo, err error) error {
		if pathStr == p {
			return nil
		}

		if info.IsDir() && !recursion {
			return filepath.SkipDir
		}

		fileNameAll := info.Name()
		fileExt := path.Ext(fileNameAll)
		fileName := fileNameAll[:len(fileNameAll)-len(fileExt)]
		fmt.Println(p, info.Name())
		if fileExt == ext {
			newFileName := reg.ReplaceAllString(fileName, repl)
			if newFileName != fileName {
				newPath := fmt.Sprintf("%s%s.%s", p[:len(p)-len(fileNameAll)], newFileName, fileNameSuffix)
				return os.Rename(p, newPath)
			}
		}

		return nil
	})
}
