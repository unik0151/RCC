package file

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
)

const (
	secretDirPath string = "E:\\remote-test\\test_ethereum\\data\\keystore"
)

func ReadFile() []byte {

	dir, err := os.Open(secretDirPath)
	if err != nil {
		log.Fatal(err)
	}
	defer dir.Close()

	files, err := dir.ReadDir(0)
	if err != nil {
		log.Fatal(err)
	}

	sort.Slice(files, func(i, j int) bool {
		fileInfoI, err := files[i].Info()
		if err != nil {
			log.Fatal(err)
		}
		fileInfoJ, err := files[j].Info()
		if err != nil {
			log.Fatal(err)
		}
		return fileInfoI.ModTime().Before(fileInfoJ.ModTime())
	})
	fileFull := filepath.Join(secretDirPath, files[0].Name())
	fileBytes, err := os.ReadFile(fileFull)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("fileFull: %v\n", fileFull)
	return fileBytes
}
