package decode

import (
	"fmt"

	"github.com/quartercastle/vector"
)

type Model struct {
	solidName string
	facets    []Facet
}

func (m Model) ModelVolume() float64 {
	return 0
}

func (m Model) SurfaceArea() {
	var totalSurfaceArea float64
	for _, facet := range m.facets {
		totalSurfaceArea += facet.FacetArea()
		fmt.Println(facet.FacetArea())
	}

	fmt.Println(totalSurfaceArea)
}

func (m Model) Facets() {
	fmt.Println("Total facets", len(m.facets))
	for facet, i := range m.facets {
		fmt.Println(i)

		fmt.Printf("%+v\n", facet)
	}
}

type Facet struct {
	vertexes    [][]float64
	facetNormal []float64
}

func (f Facet) FacetArea() float64 {
	ab := vector.Vector(f.vertexes[0]).Sub(vector.Vector(f.vertexes[1]))
	ac := vector.Vector(f.vertexes[0]).Sub(vector.Vector(f.vertexes[2]))

	crossFactor, err := ab.Cross(ac)
	if err != nil {
		fmt.Println(err.Error())
	}

	return 0.5 * crossFactor.Magnitude()
}
