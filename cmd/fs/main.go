package main

import (
	"log"
	"os"

	"github.com/enverbisevac/golang-examples/pkg/dir"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fsys := os.DirFS(wd)
	log.Printf("Total go files %d", len(dir.Files(fsys, ".go")))
}
