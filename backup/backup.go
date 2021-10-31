package backup

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	archivePaths []string
	archiveRoots []*Entry
	entries      map[string]*Entry
)

const (
	slash     = "/"
	backSlash = `\`
)

type Entry struct {
	dirname  string
	basename string
	entries  []*Entry
}

func init() {
	archivePaths = make([]string, len(os.Args)-1)
	copy(archivePaths, os.Args[1:])
	entries = make(map[string]*Entry)

	for _, path := range archivePaths {
		readArchive(path)
	}

	fmt.Println(entries, len(entries))
	fmt.Println(archiveRoots, len(archiveRoots))
}

func readArchive(path string) {
	file, _ := os.Open(path)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var sep string

	for scanner.Scan() {
		line := scanner.Text()
		if sep == "" {
			if !strings.Contains(line, slash) {
				sep = backSlash
			} else {
				sep = slash
			}
		}

		segs := strings.Split(line, sep)
		basename, dirname := segs[len(segs)-1], strings.Join(segs[0:len(segs)-1], sep)
		entry := &Entry{dirname: path + sep + dirname, basename: basename}
		entries[entry.dirname+sep+entry.basename] = entry

		parent, ok := entries[entry.dirname]
		if !ok {
			parent = &Entry{}
			entries[entry.dirname] = parent
			archiveRoots = append(archiveRoots, parent)
		}
		parent.entries = append(parent.entries, entry)
	}
}
