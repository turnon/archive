package backup

import (
	"os"
)

func init() {
	switch os.Args[1] {
	case "-d":
		dumpBackup()
	case "-l":
		loadBackups()
	}
}
