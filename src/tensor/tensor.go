package tensor

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

type Dimension struct {
	Dimen            int
	Data             float32
	MultiDimensional []Dimension
}

type Tensor struct {
	Shape      []int64
	DataTensor []Dimension
}

func (t *Tensor) DataInit(dataIni interface{}) {
	shapeLen := len(t.Shape)
	var data []Dimension
	//for i := 0; i < shapeLen; i++ {
	var index int = 0
	var elements []float32
	if dataIni != nil {
		elements = getInitialElements(dataIni)
		if numElements(t.Shape) != int64(len(elements)) {
			fmt.Println("Datos Iniciales no concuerdan con las dimensiones", t)
			return
		}
	}
	data = createDataTensor(0, t.Shape, shapeLen, &elements, &index)
	t.DataTensor = data
	fmt.Println("Datos del Tensor", t)

}

func createDataTensor(dimen int, shapeValue []int64, shapeLen int, elements *[]float32, index *int) []Dimension {

	//    if(data==nil){
	data := make([]Dimension, shapeValue[dimen])
	//}
	//vector
	for j := range data {
		var dataDimen Dimension
		if dimen == (shapeLen - 1) {

			if elements == nil {
				dataDimen.Data = rand.Float32() * 5
			} else {
				dataDimen.Data = (*elements)[*index]
				*index++
			}
		} else {
			dataDimen.MultiDimensional = createDataTensor(dimen+1, shapeValue, shapeLen, elements, index)
		}
		dataDimen.Dimen = dimen + 1
		data[j] = dataDimen

	}

	return data
}

func (t *Tensor) Reshape(newShape []int64) Tensor {
	numElem := numElements(t.Shape)
	elements := make([]float32, numElem)
	var tensorResult Tensor
	if numElem != numElements(newShape) {
		fmt.Println("Nuevo Tamaños incompatibles", t.Shape, newShape)
		return tensorResult
	} else {
		shapeLen := len(t.Shape)
		var index int = 0
		setElementsFromOriginalTensor(0, &index, t.DataTensor, shapeLen, elements)
		//fmt.Println("Nuevo Tamaños elements", elements)
		index = 0
		data := createDataTensor(0, t.Shape, shapeLen, &elements, &index)
		tensorResult.Shape = t.Shape
		tensorResult.DataTensor = data
		fmt.Println("Nuevo Tensor elements", tensorResult)
	}
	return tensorResult
}

func (t *Tensor) IndexSelect(dimen int, indexVector []int64) Tensor {
	shapeLen := len(t.Shape)
	var index int = 0
	outPutShape := t.Shape
	outPutShape[dimen] = int64(len(indexVector))
	numElem := numElements(outPutShape)
	elements := make([]float32, numElem)
	setElementsFromOTandIndexV(0, &index, t.DataTensor, shapeLen, elements, indexVector, dimen)
	index = 0
	data := createDataTensor(0, t.Shape, shapeLen, &elements, &index)
	var out Tensor
	out.DataTensor = data
	out.Shape = outPutShape
	fmt.Println("Selected index REsult", out)
	return out

}

func (t *Tensor) HadamardProduct(tensorA, tensorB Tensor) Tensor {
	var out Tensor
	if testEq(tensorA.Shape, tensorB.Shape) {
		shapeLen := len(t.Shape)
		numElem := numElements(t.Shape)
		elements := make([]float32, numElem)
		var index int = 0

		setElementsFromHandmardO(0, &index, tensorA.DataTensor, tensorB.DataTensor, shapeLen, elements)
		index = 0
		data := createDataTensor(0, t.Shape, shapeLen, &elements, &index)

		out.DataTensor = data
		out.Shape = tensorA.Shape
		fmt.Println("Hadmard Result", out)
	}
	return out

}

func (t Tensor) String() string {

	//shapeLen := len(t.Shape)
	//toString := fmt.Sprintln("Dim:", shapeLen)
	toString := fmt.Sprintln("Shape:", t.Shape)
	toString += "["
	for dim := range t.DataTensor {

		toString = getString(1, t.DataTensor[dim], toString)

	}
	toString += "]"
	return toString
}

func getString(dim int, data Dimension, toString string) string {
	//+=toString = addSalto(toString)
	if data.Data == 0 {
		toString = addSalto(dim, toString)
		toString += "["

		for _, value := range data.MultiDimensional {

			toString = getString(dim+1, value, toString)

		}
		toString += "]"
	} else {
		toString = addSpace(toString)
		toString += fmt.Sprintf("%1.2f", data.Data)
		// 	for _, d := range shape {

	}
	return toString
}

func setElementsFromOriginalTensor(dimen int, index *int, data []Dimension, shapeLen int, elements []float32) {
	rows := len(data)

	for i := 0; i < rows; i++ {

		if dimen == (shapeLen - 1) {
			value := data[i].Data
			elements[(*index)] = value
			*index++

		} else {
			setElementsFromOriginalTensor(dimen+1, index, data[i].MultiDimensional, shapeLen, elements)
		}
	}

}

func setElementsFromHandmardO(dimen int, index *int, dataA, dataB []Dimension, shapeLen int, elements []float32) {
	rows := len(dataA)

	for i := 0; i < rows; i++ {

		if dimen == (shapeLen - 1) {
			value := dataA[i].Data * dataB[i].Data
			elements[(*index)] = value
			*index++

		} else {
			setElementsFromHandmardO(dimen+1, index, dataA[i].MultiDimensional, dataB[i].MultiDimensional, shapeLen, elements)
		}
	}

}

func setElementsFromOTandIndexV(dimen int, index *int, data []Dimension, shapeLen int, elements []float32, indexVector []int64, dimSelect int) {

	rows := len(data)

	indexSelect := dimen == dimSelect
	if indexSelect {
		rows = len(indexVector)
	}
	for i := 0; i < rows; i++ {

		if dimen == (shapeLen - 1) {
			//fmt.Println(*index, ":", i, "-->", data[i].Data)
			var value float32
			if indexSelect {
				value = data[indexVector[i]].Data
			} else {
				value = data[i].Data
			}
			elements[*index] = value
			*index++

		} else {
			var value []Dimension
			if indexSelect {
				value = data[indexVector[i]].MultiDimensional
			} else {
				value = data[i].MultiDimensional
			}
			setElementsFromOTandIndexV(dimen+1, index, value, shapeLen, elements, indexVector, dimSelect)

		}
	}

}

func numElements(shape []int64) int64 {
	n := int64(1)
	for _, d := range shape {
		n *= d
	}
	return n
}

func getInitialElements(elementsIni interface{}) []float32 {
	elemtString := fmt.Sprint(elementsIni)
	var elementos []float32
	element := strings.Split(elemtString, " ")
	for i := range element {
		t := strings.Replace(element[i], "[", "", -1)
		t = strings.Replace(t, "]", "", -1)
		f, err := (strconv.ParseFloat(t, 2))
		if err != nil {
			return nil
		} else {
			elementos = append(elementos, float32(f))
		}
	}
	return elementos

}
