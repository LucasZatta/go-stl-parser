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
	buffer := make([]byte, 5)
	_, err := io.ReadFull(r, buffer)

	if err != nil {
		return nil, err
	}

	if s, ok := r.(io.Seeker); ok {
		_, err = s.Seek(0, io.SeekStart)
		if err != nil {
			return nil, err
		}
	} else {
		r = io.MultiReader(bytes.NewReader(buffer), r)
	}

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
		// //fmt.Println(currentWord)
		switch state {
		case start:
			if currentWord != "solid" {
				return nil, fmt.Errorf("ASCII STL file must start with `solid`")
			}
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
				//i can initialize the facet here. but this step wouldnt do much beyond this. Maybe i cant just cut it
			case "endsolid":

				return mesh, nil
			}

		case facetStep:
			if currentWord != "normal" {
				//error
				return nil, fmt.Errorf("error 1")
			}
			state = normal

		case normal:
			normalV, err := scanTriple(scanner)
			////fmt.Println(normalV)
			facet.facetNormal = normalV
			if err != nil {
				return nil, fmt.Errorf("error 2")
			}

			state = facetLoop

		case facetLoop:
			if currentWord != "outer" {
				return nil, fmt.Errorf("expected keywords `outer loop`")
			}
			state = outer

		case outer:
			if currentWord != "loop" {
				return nil, fmt.Errorf("expected keywords `outer loop`")
			}
			state = loop

		case loop:
			switch currentWord {
			case "vertex":
				state = vertices
			case "endloop":
				state = endloop
			default:
				return nil, fmt.Errorf("expected `vertex` or `endloop`")
			}

		case vertices:
			vert, err := scanTriple(scanner)
			if err != nil {
				return nil, fmt.Errorf("error 4")
			}
			verts = append(verts, vert)
			state = loop

		case endloop:
			if currentWord != "endfacet" {
				return nil, fmt.Errorf("expected keyword `endfacet`")
			}
			if len(verts) != 3 {
				//fmt.Println("len verts %i", len(verts))
				//fmt.Println(verts)

				return nil, fmt.Errorf("expected 3 vertices")

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

func scanTriple(scanner *bufio.Scanner) ([]float64, error) {
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
	n, err := strconv.ParseFloat(scanner.Text(), 32)
	return n, err
}
