package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func dirTree(output io.Writer, path string, printFiles bool) error {
	err := dTree(output, path, printFiles, "")
	if err != nil {
		panic(err.Error())
	}
	return nil
}

func dTree(output io.Writer, sDir string, printFiles bool, sIndent string) error {
	files, err := ioutil.ReadDir(sDir)
	if err != nil {
		return nil
	}

	if printFiles != true {
		n := 0
		for _, file := range files {
			if file.IsDir() {
				files[n] = file
				n++
			}
		}
		files = files[:n]
	}

	numberOfFiles := len(files)
	counterOfFiles := 0
	prefix := "├───"
	extraIndent := "│\t"

	for _, file := range files {
		counterOfFiles++
		if counterOfFiles == numberOfFiles {
			prefix = "└───"
			extraIndent = "\t"
		}
		if file.IsDir() {
			fmt.Fprintf(output, sIndent+prefix+file.Name()+"\n")
			dTree(output, sDir+string(os.PathSeparator)+file.Name(), printFiles, sIndent+extraIndent)
		} else {
			if file.Size() == 0 {
				fmt.Fprintf(output, sIndent+prefix+file.Name()+" (empty)\n")
			} else {
				fmt.Fprintf(output, sIndent+prefix+file.Name()+" (%vb)\n", file.Size())
			}
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
