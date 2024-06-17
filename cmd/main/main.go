package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/lucaszatta/go-stl-parser/internal/decode"
)

func main() {
	filePath := flag.String("fp", "././stl/solid.stl", "STL file path")
	flag.Parse()

	f, err := os.Open(*filePath)
	if err != nil {
		log.Fatal(err)
	}

	mesh, err := decode.DecodeSTL(f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Number of Triangles: %d \n", mesh.Facets())
	fmt.Printf("Surface Area: %f \n", mesh.SurfaceArea())
}
