package main

import (
	"fmt"
	t "hello/src/estructuraT"
)

func main() {

	//Ejemplo de uso de la fución reshape en 1 dimensión
	//sampleReShape1D()
	//sampleReShape2D()
	//sampleReShapeND()
	//sampleIndexSelectN()
	sampleHadamard2X2()
	sampleHadamardND()

}

// Permite verificar el funcionamiento del la funcion RashSahpe para un T de una dimension
func sampleReShape1D() {

	shape := []int64{2}
	data := []int{1, 2}
	var myStructT t.StructT
	myStructT.Shape = shape
	myStructT.SetData(data)
	fmt.Println("Tensor Creado:", myStructT)
	newShape := []int64{2, 1}
	fmt.Println("Aplicando ReShape:", newShape)
	myStructT.Reshape(newShape)
	fmt.Println("Tensor Creado:", myStructT.Reshape(newShape))

}

func sampleReShape2D() {

	shape := []int64{3, 2}
	data := [][]float32{{2, 3}, {4, 5}, {6, 7}}
	var mytensor t.StructT
	mytensor.Shape = shape
	mytensor.SetData(data)
	fmt.Println("Tensor Creado:", mytensor)
	newShape := []int64{2, 1, 3}
	fmt.Println("Aplicando RaShape:", newShape)
	mytensor.Reshape(newShape)
	fmt.Println("Tensor Creado:", mytensor.Reshape(newShape))

}

// Ejemplo de un T de 4 Dimensiones inicializada con auto generacion datos aleatorios al pasar el parametro nul en los Datos
func sampleReShapeND() {

	shape := []int64{3, 2, 1, 4}

	var myStructT t.StructT
	myStructT.Shape = shape
	myStructT.SetData(nil)
	fmt.Println("Estructura creada:", myStructT)
	newShape := []int64{4, 2, 1, 3}
	fmt.Println("Aplicando RaShape:", newShape)
	fmt.Println("Estructura Resultado:", myStructT.Reshape(newShape))

}

// Ejemplo //1. IndexSelect([[1, 2], [3, 4]], 0, [0])    			[[1, 2]]
func sampleIndexSelect1() {
	shape := []int64{2, 2}
	data := [][]int{{1, 2}, {3, 4}}
	dim := 0
	indexVector := []int64{0, 0}
	var myStructT t.StructT
	myStructT.Shape = shape
	myStructT.SetData(data)
	fmt.Println("Estructura Creada:", myStructT)
	fmt.Println("Aplicando IndexSelect:", indexVector)
	fmt.Println("Estructura Resultado:", myStructT.IndexSelect(dim, indexVector))

}

// //4. IndexSelect([[1, 2], [3, 4]], 1, [0])
func sampleIndexSelect4() {
	shape := []int64{2, 2}
	data := [][]int{{1, 2}, {3, 4}}
	dim := 1
	indexVector := []int64{0}
	var myStructT t.StructT
	myStructT.Shape = shape
	myStructT.SetData(data)
	fmt.Println("Estructura Creada:", myStructT)
	fmt.Println("Aplicando IndexSelect:", indexVector)
	fmt.Println("Estructura Resultado:", myStructT.IndexSelect(dim, indexVector))

}

// Ejemplo de una T con 4 Dimensiones generado aletoreamente
func sampleIndexSelectN() {
	shape := []int64{3, 3, 2, 2}

	dim := 3
	indexVector := []int64{0, 1, 1, 0, 1}
	var myStructT t.StructT
	myStructT.Shape = shape
	myStructT.SetData(nil)
	fmt.Println("Estructura Creado:", myStructT)
	fmt.Println("Aplicando IndexSelect:", indexVector)
	fmt.Println("Estructura Resultado:", myStructT.IndexSelect(dim, indexVector))

}

func sampleHadamard2X2() {
	shape := []int64{2, 2}
	dataA := [][]int{{1, 2}, {3, 4}}
	dataB := [][]int{{2, 2}, {4, 4}}
	var myStructA t.StructT
	myStructA.Shape = shape
	myStructA.SetData(dataA)
	fmt.Println("Estructura A Creada:", myStructA)
	var myStructB t.StructT
	myStructB.Shape = shape
	myStructB.SetData(dataB)
	fmt.Println("Estructura A Creada:", myStructA)
	fmt.Println("Estructura Resultado:", myStructA.HadamardProduct(myStructA, myStructB))

}

func sampleHadamardND() {
	shape := []int64{2, 2, 3, 1}

	var myStructA t.StructT
	myStructA.Shape = shape
	myStructA.SetData(nil)
	fmt.Println("Estructura A Creada:", myStructA)
	var myStructB t.StructT
	myStructB.Shape = shape
	myStructB.SetData(nil)
	fmt.Println("Estructura A Creada:", myStructA)
	fmt.Println("Estructura Resultado:", myStructA.HadamardProduct(myStructA, myStructB))

}
