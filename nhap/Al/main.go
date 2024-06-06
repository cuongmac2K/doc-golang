package main

import (
	"fmt"
	"log"

	"gorgonia.org/gorgonia"
	"gorgonia.org/tensor"
)

func main() {
	// Dữ liệu huấn luyện giả định
	xs := []float64{1, 2, 3, 4}
	ys := []float64{2, 4, 6, 8}

	// Tạo graph
	g := gorgonia.NewGraph()

	// Tạo input node
	x := gorgonia.NewMatrix(g, tensor.Float64, gorgonia.WithShape(4, 1), gorgonia.WithValue(tensor.New(tensor.Of(tensor.Float64), tensor.WithBacking(xs))))

	// Tạo target node
	y := gorgonia.NewMatrix(g, tensor.Float64, gorgonia.WithShape(4, 1), gorgonia.WithValue(tensor.New(tensor.Of(tensor.Float64), tensor.WithBacking(ys))))

	// Tạo biến tham số
	w := gorgonia.NewMatrix(g, tensor.Float64, gorgonia.WithShape(1, 1), gorgonia.WithName("w"), gorgonia.WithValue(gorgonia.NewTensor(g, tensor.Float64, tensor.WithShape(1, 1), tensor.WithBacking([]float64{0}))))
	b := gorgonia.NewMatrix(g, tensor.Float64, gorgonia.WithShape(1), gorgonia.WithName("b"), gorgonia.WithValue(gorgonia.NewTensor(g, tensor.Float64, tensor.WithShape(1), tensor.WithBacking([]float64{0}))))

	// Tạo phương trình hồi quy tuyến tính
	yPred := gorgonia.Must(gorgonia.Add(gorgonia.Must(gorgonia.Mul(x, w)), b))

	// Tính toán loss (sai số)
	cost := gorgonia.Must(gorgonia.Mean(gorgonia.Must(gorgonia.Square(gorgonia.Must(gorgonia.Sub(yPred, y))))))

	// Khởi tạo gradient tape
	grad, err := gorgonia.Gradient(cost, w, b)
	if err != nil {
		log.Fatal(err)
	}

	// Tạo VM
	vm := gorgonia.NewTapeMachine(g, gorgonia.BindDualValues(w, b))

	// Huấn luyện
	for i := 0; i < 1000; i++ {
		vm.Reset()
		if err := vm.RunAll(); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Epoch: %d, Loss: %.2f\n", i+1, cost.Value().Data().(float64))
	}

	// In kết quả
	fmt.Printf("Trained parameters - w: %.2f, b: %.2f\n", w.Value().Data().(float64), b.Value().Data().(float64))
}
