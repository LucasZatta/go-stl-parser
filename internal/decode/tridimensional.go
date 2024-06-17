package decode

import (
	"fmt"

	"github.com/quartercastle/vector"
)

type Model struct {
	solidName string
	facets    []Facet
}

// func (m Model) ModelVolume() float64 {
// 	return 0
// }

func (m Model) SurfaceArea() float64 {
	var totalSurfaceArea float64
	for _, facet := range m.facets {
		totalSurfaceArea += facet.FacetArea()
	}
	return totalSurfaceArea
}

func (m Model) Facets() int {
	return len(m.facets)
}

type Facet struct {
	vertices    [][]float64
	facetNormal []float64
}

func (f Facet) FacetArea() float64 {
	vectorAB := vector.Vector(f.vertices[0]).Sub(vector.Vector(f.vertices[1]))
	vectorAC := vector.Vector(f.vertices[0]).Sub(vector.Vector(f.vertices[2]))

	crossFactor, err := vectorAB.Cross(vectorAC)
	if err != nil {
		fmt.Println(err.Error())
	}

	return 0.5 * crossFactor.Magnitude()
}
