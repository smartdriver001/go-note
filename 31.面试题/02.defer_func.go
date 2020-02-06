package main

import "fmt"

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}
/**
output:
10 1 2 3
20 0 2 2
2 0 2 2
1 1 3 4
 */
func main() {
	a, b := 1, 2
	defer calc("1", a, calc("10", a, b)) //1. 执行10, 4. 执行"1"
	a = 0
	defer calc("2", a, calc("20", a, b)) //2. 执行20, 3. 执行"2"
	b = 1
}