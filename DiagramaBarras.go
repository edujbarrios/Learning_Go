package main

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Crea un nuevo gráfico
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	// Establece el título del gráfico y los ejes
	p.Title.Text = "Diagrama de barras"
	p.Y.Label.Text = "Valores"
	p.X.Label.Text = "Valor_Categoria"

	// Crea un conjunto de datos con los valores a representar en el diagrama
	values := []float64{3, 4, 5, 6, 7}
	labels := []string{"A1", "B1", "C1", "D1", "E1"}

	// Crea un diagrama de barras a partir de los datos
	bar, err := plotter.NewBarChart(values, labels)
	if err != nil {
		panic(err)
	}

	// Añade el diagrama de barras al gráfico
	p.Add(bar)

	// Guarda el gráfico en formato PNG
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "NuevoGrafico.png"); err != nil {
		panic(err)
	}
}
