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

	// Crea un cubo con esquinas en (0, 0, 0), (1, 0, 0), (1, 1, 0), (0, 1, 0), (0, 0, 1), (1, 0, 1), (1, 1, 1) y (0, 1, 1)
	cube, err := plotter.NewCuboid(plotter.XYZ(0, 0, 0), plotter.XYZ(1, 1, 1))
	if err != nil {
		panic(err)
	}

	// Añade el cubo al gráfico
	p.Add(cube)

	// Establece el título del gráfico y los ejes
	p.Title.Text = "Cubo en 3D"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"
	p.Z.Label.Text = "Z"

	// Guarda el gráfico en formato PNG
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "cubo3D.png"); err != nil {
		panic(err)
	}
}
