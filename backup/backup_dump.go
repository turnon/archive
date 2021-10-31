package backup

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func dumpBackup() {
	dir, dumpFile := os.Args[2], os.Args[3]

	var (
		content strings.Builder
		space   = " "
		newline = "\n"
	)

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		var size int64
		if !info.IsDir() {
			size = info.Size()
		}

		sizeStr := strconv.FormatInt(size, 10)
		content.WriteString(sizeStr)
		content.WriteString(space)
		content.WriteString(path)
		content.WriteString(newline)

		return nil
	})

	ioutil.WriteFile(dumpFile, []byte(content.String()), 0777)
}
