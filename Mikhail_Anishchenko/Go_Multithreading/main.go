package main

import (
	"fmt"
	"log"
	"os"
	"sync"
)

func run(dirname string) {
	files, err := os.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}
	var wg sync.WaitGroup
	for _, f := range files {
		wg.Add(1)
		go func(filepath string) {
			defer wg.Done()
			matrix := ReadMatrixFromFile(filepath)
			fmt.Println("Determinant of matrix in ", filepath, " = ", det(matrix))
		}(dirname + "/" + f.Name())
	}
	wg.Wait()
}

func main() {
	dirname := "matrices"
	if len(os.Args) > 1 {
		dirname = os.Args[1]
	}
	run(dirname)
}
