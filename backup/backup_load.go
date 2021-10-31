package backup

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type fileEntry struct {
	size int64
	path string
}

func loadBackups() {
	_loadBackups(os.Args[2:]...)
}

func _loadBackups(baks ...string) {
	for _, bak := range baks {
		_load(bak)
	}
}

func _load(bak string) {
	file, _ := os.Open(bak)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		size, path := sizeAndPath(line)
		fe := fileEntry{size: size, path: path}
		fmt.Println(fe)
	}
}

func sizeAndPath(line string) (size int64, path string) {
	sizePath := strings.SplitN(line, " ", 2)
	size, _ = strconv.ParseInt(sizePath[0], 10, 64)
	path = sizePath[1]
	return
}
