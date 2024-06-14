package main

import (
	"log"
	"os"

	"github.com/lucaszatta/go-stl-parser/internal/decode"
)

func main() {
	f, err := os.Open("cube.stl")
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
