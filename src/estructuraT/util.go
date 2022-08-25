package estructuraT

import "fmt"

func testEq(a, b []int64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func addSalto(dim int, toString string) string {
	if len(toString) > 0 {
		lastCh := toString[len(toString)-1:]
		if lastCh == "]" {
			toString += "\n"
			if lastCh != ("\n") {
				for i := 0; i < dim; i++ {
					toString += " "
				}

			}
		}
	}
	return toString
}

func addSpace(toString string) string {
	if len(toString) > 0 {
		lastCh := toString[len(toString)-1:]
		if lastCh != "[" {
			toString += " "
		}
	}
	return toString
}

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
