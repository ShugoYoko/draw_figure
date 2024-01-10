package main

import (
	"math"

	"github.com/fogleman/gg"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func main() {
	p := plot.New()

	p.Title.Text = "Plotutil example"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	err := plotutil.AddLinePoints(p,
		"y=x^2+3", funcPoints(15),
	)
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "points.png"); err != nil {
		panic(err)
	}

	drawPic()
}

func funcPoints(n int) plotter.XYs {
	pts := make(plotter.XYs, n)
	for i := range pts {
		x := float64(i)
		pts[i].X = x
		pts[i].Y = math.Pow(x, 2) + 3
	}
	return pts
}

func drawPic() {
	dc := gg.NewContext(1000, 1000)
	// 背景の描画
	dc.DrawRectangle(0, 0, 1000, 1000)
	dc.SetRGB(1, 1, 1)
	dc.Fill()

	//円
	dc.DrawCircle(400, 400, 300)
	dc.SetRGB(0, 0, 1)
	dc.Stroke()
	//dc.Fill()

	//プロット
	dc.DrawPoint(400, 400, 10)
	dc.SetRGB(0, 0, 1)
	dc.Fill()

	//線（三角形も作成可能）
	dc.DrawLine(100, 100, 200, 200)
	dc.SetRGB(0, 1, 0)
	dc.SetLineWidth(5.0)
	dc.Stroke()

	//円弧
	ang1 := gg.Radians(30)
	ang2 := gg.Radians(270)
	dc.DrawArc(200, 200, 100, ang2, ang1)
	dc.SetRGB(1, 1, 0)
	dc.Stroke()

	dc.SavePNG("out.png")
}
