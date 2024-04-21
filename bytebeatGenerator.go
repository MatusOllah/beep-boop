package main

import (
	crand "crypto/rand"
	"encoding/binary"
	"log/slog"
	"math"
	"math/bits"
	"math/rand"

	"github.com/dop251/goja"
)

type BytebeatGenerator struct {
	T int

	prog *goja.Program
	vm   *goja.Runtime
}

func NewBytebeatGenerator(expr string) *BytebeatGenerator {
	var seed int64
	if err := binary.Read(crand.Reader, binary.LittleEndian, &seed); err != nil {
		panic(err)
	}

	r := rand.New(rand.NewSource(seed))

	vm := goja.New()

	// math stuff
	vm.Set("E", math.E)
	vm.Set("LN10", math.Ln10)
	vm.Set("LN2", math.Ln2)
	vm.Set("LOG10E", math.Log10E)
	vm.Set("LOG2E", math.Log2E)
	vm.Set("PI", math.Pi)
	vm.Set("SQRT1_2", math.Sqrt(0.5))
	vm.Set("SQRT2", math.Sqrt2)

	vm.Set("abs", math.Abs)
	vm.Set("acos", math.Acos)
	vm.Set("acosh", math.Acosh)
	vm.Set("asin", math.Asin)
	vm.Set("asinh", math.Asinh)
	vm.Set("atan", math.Atan)
	vm.Set("atan2", math.Atan2)
	vm.Set("atanh", math.Atanh)
	vm.Set("cbrt", math.Cbrt)
	vm.Set("ceil", math.Ceil)
	vm.Set("clz32", bits.LeadingZeros32)
	vm.Set("cos", math.Cos)
	vm.Set("cosh", math.Cosh)
	vm.Set("exp", math.Exp)
	vm.Set("expm1", math.Expm1)
	vm.Set("floor", math.Floor)
	//TODO: fround
	vm.Set("hypot", math.Hypot)
	//TODO: imul
	vm.Set("log", math.Log)
	vm.Set("log10", math.Log10)
	vm.Set("log1p", math.Log1p)
	vm.Set("log2", math.Log2)
	vm.Set("max", math.Max)
	vm.Set("min", math.Min)
	vm.Set("pow", math.Pow)
	vm.Set("random", r.Float64)
	vm.Set("round", math.Round)
	vm.Set("sign", func(x float64) float64 {
		switch {
		case x == 0:
			return 0
		case x > 0:
			return 1
		case x < 0:
			return -1
		}
		panic("unreachable")
	})
	vm.Set("sin", math.Sin)
	vm.Set("sinh", math.Sinh)
	vm.Set("sqrt", math.Sqrt)
	vm.Set("tan", math.Tan)
	vm.Set("tanh", math.Tanh)
	vm.Set("trunc", math.Trunc)

	var prog *goja.Program
	prog, err := goja.Compile("", expr, false)
	if err != nil {
		slog.Error("failed to compile expression", "err", err)
		prog = goja.MustCompile("", "t", false)
	}

	return &BytebeatGenerator{
		T:    0,
		prog: prog,
		vm:   vm,
	}
}

func (g *BytebeatGenerator) Stream(samples [][2]float64) (n int, ok bool) {
	for i := range samples {
		g.vm.Set("t", g.T)
		_sample, err := g.vm.RunProgram(g.prog)
		if err != nil {
			slog.Error(err.Error())
			return len(samples), false
		}
		switch mode {
		case "Bytebeat":
			sample := _sample.ToInteger()
			//samples[i][0] = float64(sample&255)/127 - 1
			samples[i][0] = (float64(sample % 256)) / 256
		case "Floatbeat":
			sample := _sample.ToFloat()
			samples[i][0] = sample
		}
		samples[i][1] = samples[i][0]
		g.T++
	}

	return len(samples), true
}

func (g *BytebeatGenerator) Err() error {
	return nil
}

func (g *BytebeatGenerator) SetExpr(expr string) error {
	prog, err := goja.Compile("", expr, false)
	if err != nil {
		return err
	}

	g.prog = prog

	return nil
}
