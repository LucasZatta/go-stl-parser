package decode

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strconv"
)

const (
	start = iota
	solid
	facetStep
	normal
	facetLoop
	outer
	loop
	vertices
	endloop
	endfacet
	endsolid
)

func DecodeSTL(r io.Reader) (*Model, error) {
	//create buffer and open file
	buffer := make([]byte, 5)
	_, err := io.ReadFull(r, buffer)
	if err != nil {
		return nil, err
	}

	//returning file reader to the start
	if s, ok := r.(io.Seeker); ok {
		_, err = s.Seek(0, io.SeekStart)
		if err != nil {
			return nil, err
		}
	} else {
		r = io.MultiReader(bytes.NewReader(buffer), r)
	}

	//calling actual parser! simple state machine that goes through the stl file
	return ParseSTL(r)
}

func ParseSTL(r io.Reader) (*Model, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	state := start

	var mesh *Model = new(Model)
	var facet *Facet = new(Facet)

	var verts [][]float64

	for scanner.Scan() {
		currentWord := scanner.Text()
		switch state {
		case start:
			state = solid

		case solid:
			if currentWord == "facet" {
				state = facetStep
			} else {
				mesh.solidName = currentWord
				state = endfacet
			}

		case endfacet:
			switch currentWord {
			case "facet":
				state = facetStep
			case "endsolid":
				return mesh, nil
			}

		case facetStep:
			state = normal

		case normal:
			normalV, err := scanCoords(scanner)
			facet.facetNormal = normalV
			if err != nil {
				return nil, fmt.Errorf("error scaning normal vector coordinates")
			}

			state = facetLoop

		case facetLoop:
			state = outer

		case outer:
			state = loop

		case loop:
			switch currentWord {
			case "vertex":
				state = vertices
			case "endloop":
				state = endloop
			}

		case vertices:
			vert, err := scanCoords(scanner)
			if err != nil {
				return nil, fmt.Errorf("error scaning verts coordinates")
			}
			verts = append(verts, vert)
			state = loop

		case endloop:
			if len(verts) != 3 {
				return nil, fmt.Errorf("expected 3 vertices but got %d", len(verts))
			}

			facet.vertices = verts
			mesh.facets = append(mesh.facets, *facet)
			verts = [][]float64{}
			facet = &Facet{}
			state = endfacet
		}

	}
	return mesh, nil
}

func scanCoords(scanner *bufio.Scanner) ([]float64, error) {
	x, err := strconv.ParseFloat(scanner.Text(), 32)
	if err != nil {
		return nil, err
	}
	y, err := scanFloat(scanner)
	if err != nil {
		return nil, err
	}
	z, err := scanFloat(scanner)
	if err != nil {
		return nil, err
	}

	return []float64{x, y, z}, nil
}

func scanFloat(scanner *bufio.Scanner) (float64, error) {
	if !scanner.Scan() {
		return 0, io.ErrUnexpectedEOF
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	n, err := strconv.ParseFloat(scanner.Text(), 64)
	return n, err
}
