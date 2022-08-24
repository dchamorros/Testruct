package main

import (
	"fmt"
	tensor "hello/src/tensor"
)

func intScanln(n int) ([]int, error) {
	x := make([]int, n)
	y := make([]interface{}, len(x))
	for i := range x {
		y[i] = &x[i]
	}
	n, err := fmt.Scanln(y...)
	x = x[:n]
	return x, err
}

func main() {
	//var numDimen int
	fmt.Println("Ingreso el entero que represnta el numero de las dimensiones del tensor : ")
	//fmt.Scanln(&numDimen)
	//rows = 4

	//colums = 4
	fmt.Println("Ingreso los tamaños para cada dimencion: ")
	var mytensor tensor.Tensor
	// shape, err := intScanln(numDimen)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	//IndexSelect([1, 2, 3, 4], 0, [0, 0, 2])[1, 1, 3]

	//1. IndexSelect([[1, 2], [3, 4]], 0, [0])    			[[1, 2]]
	//2. IndexSelect([[1, 2], [3, 4]], 0, [0, 0])			[[1, 2], [1, 2]]
	//3. IndexSelect([[1, 2], [3, 4]], 0, [0, 0, 1, 1])	[[1, 2], [1, 2], [3, 4], [3, 4]]
	//4. IndexSelect([[1, 2], [3, 4]], 1, [0])				[[1], [3]]
	//5. IndexSelect([[1, 2], [3, 4]], 1, [0, 0])			[[1, 1], [3, 3]]
	//6. IndexSelect([[1, 2], [3, 4]], 1, [0, 0, 1, 1])	[[1, 1, 2, 2], [3, 3, 4, 4]]
	shape := []int64{3, 3, 3} //  1. {1,2} -->[[1, 2]]2 repreesnta al ultimo de cada shape y el primero representa el tamoño de index
	shapes := [][]int{{2, 2}, {3, 4}}
	fmt.Println("shapes")
	fmt.Println(shapes)
	//selecDimen := 0S
	//indexVector := []int{0, 0, 2}
	mytensor.Shape = shape
	mytensor.DataInit()

	var mytensorB tensor.Tensor
	mytensorB.Shape = shape
	mytensorB.DataInit()

	mytensor.HadamardProduct(mytensor, mytensorB)
	//dim := 0
	//indexVector := []int64{0}
	//mytensor.IndexSelect(dim, indexVector)
	//indexVector = []int64{0, 0}
	//mytensor.IndexSelect(dim, indexVector)
	//indexVector = []int64{0, 1, 0, 1}
	//mytensor.IndexSelect(dim, indexVector)
	//fmt.Scanln(&colums)

	// matrix := initMatrix(rows, colums)
	// fmt.Println("matrix1:", matrix)
	// matrix = populateRandomValues(matrix)
	// fmt.Println("matrix1:", matrix)
	// matrix = tensorReshape(matrix, 8, 2)
	// fmt.Println("matrix1:", matrix)

}

func initMatrix(rows, colums int) [][]float32 {
	matrix := make([][]float32, rows)
	for i := range matrix {
		matrix[i] = make([]float32, colums)
	}

	return matrix
}
