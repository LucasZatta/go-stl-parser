package decode

import "fmt"

type Model struct {
	solidName string
	facets    []Facet
}

func (m Model) ModelVolume() float32 {
	return 0
}

func (m Model) SurfaceArea() float32 {
	return 0
}

func (m Model) Facets() {
	for facet, i := range m.facets {
		fmt.Println("Facet number", i)
		fmt.Println(facet)
	}
}

type Facet struct {
	vertexes    [][]float32
	facetNormal []float32
}

func (f Facet) FacetArea() float32 {
	return 0
}
