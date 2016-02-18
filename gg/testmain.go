// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"math"
	"math/rand"
	"os"

	"github.com/aclements/go-gg/gg"
	"github.com/aclements/go-gg/ggstat"
	"github.com/aclements/go-gg/table"
	"github.com/aclements/go-moremath/vec"
)

func main() {
	xs1 := vec.Linspace(-10, 10, 100)
	for i := range xs1 {
		xs1[i] = rand.Float64()*20 - 10
	}
	ys1 := vec.Map(math.Sin, xs1)

	xs2 := vec.Linspace(-10, 10, 100)
	ys2 := vec.Map(math.Cos, xs2)

	which := []string{}
	for range xs1 {
		which = append(which, "sin")
	}
	for range xs2 {
		which = append(which, "cos")
	}

	xs := vec.Concat(xs1, xs2)
	ys := vec.Concat(ys1, ys2)

	tab := new(table.Table).Add("x", xs).Add("y", ys).Add("which", which)

	plot := gg.NewPlot(tab)
	plot.Bind("x", "x").Bind("y", "y")
	plot.GroupAuto()
	plot.Add(gg.LayerLines())

	plot.SetData(ggstat.ECDF(plot.Data(), "x", ""))
	// XXX Something is wrong here because the existing y binding
	// no longer matches up with the columns. Hence, we have to
	// re-bind it, *but* that creates a new Y scale.
	plot.Bind("x", "x").Bind("y", "cumulative density")
	plot.Add(gg.LayerSteps(gg.StepHV))
	//plot.Add(gg.LayerSteps(gg.StepHMid))

	plot.WriteSVG(os.Stdout, 400, 300)
}
