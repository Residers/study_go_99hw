package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

func dirTree(out io.Writer, path string, printFiles bool) error {
	pathPieces := strings.Split(path, string(os.PathSeparator))
	var spaces string

	for ind := 1; ind < len(pathPieces); ind++ {
		if ind%2 == 0 {
			spaces += "│   "
		} else {
			spaces += "    "
		}
	}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})
	for ind, file := range files {
		if file.Name() == ".git" {
			continue
		}
		var separator string
		if ind == len(files)-1 {
			separator = "└───"
		} else {
			separator = "├───"
		}

		var fileSize string
		if file.Size() != 0 {
			fileSize = "(" + strconv.Itoa(int(file.Size())) + "b" + ")"
		} else {
			fileSize = "(empty)"
		}
		if !file.IsDir() && printFiles {
			fmt.Printf("%s%s%s %s\n", spaces, separator, file.Name(), fileSize)
		} else if file.IsDir() {

			fmt.Printf("%s%s%s\n", spaces, separator, file.Name())
			dirTree(out, path+"/"+file.Name(), printFiles)
		}
	}

	return nil
}
func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}
