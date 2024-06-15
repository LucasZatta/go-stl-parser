package main

import (
	"flag"
	"log"
	"os"

	"github.com/lucaszatta/go-stl-parser/internal/decode"
)

func main() {
	filePath := flag.String("fp", "../../stl/solid.stl", "STL file path")
	flag.Parse()

	f, err := os.Open(*filePath)
	if err != nil {
		log.Fatal(err)
	}

	mesh, err := decode.DecodeSTL(f)
	if err != nil {
		log.Fatal(err)
	}

	mesh.Facets()
	mesh.SurfaceArea()
}
