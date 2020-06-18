package main

import "fmt"

func main() {
	params := [][]byte{}
	param1 := []byte{'5', '3', '.', '.', '7', '.', '.', '.', '.'}
	param2 := []byte{'6', '.', '.', '1', '9', '5', '.', '.', '.'}
	param3 := []byte{'.', '9', '8', '.', '.', '.', '.', '6', '.'}
	param4 := []byte{'8', '.', '.', '.', '6', '.', '.', '.', '3'}
	param5 := []byte{'4', '.', '.', '8', '.', '3', '.', '.', '1'}
	param6 := []byte{'7', '.', '.', '.', '2', '.', '.', '.', '6'}
	param7 := []byte{'.', '6', '.', '.', '.', '.', '2', '8', '.'}
	param8 := []byte{'.', '.', '.', '4', '1', '9', '.', '.', '5'}
	param9 := []byte{'.', '.', '.', '.', '8', '.', '.', '7', '9'}

	params = append(params, param1)
	params = append(params, param2)
	params = append(params, param3)
	params = append(params, param4)
	params = append(params, param5)
	params = append(params, param6)
	params = append(params, param7)
	params = append(params, param8)
	params = append(params, param9)

	solveSudoku(params)
	fmt.Println(params)
}
