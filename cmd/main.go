package main

import "fmt"

func main() {
	params := [][]byte{}
	param1 := []byte{'.', '.', '4', '.', '.', '.', '6', '3', '.'}
	param2 := []byte{'.', '.', '.', '.', '.', '.', '.', '.', '.'}
	param3 := []byte{'5', '.', '.', '.', '.', '.', '.', '9', '.'}
	param4 := []byte{'.', '.', '.', '5', '6', '.', '.', '.', '.'}
	param5 := []byte{'4', '.', '3', '.', '.', '.', '.', '.', '1'}
	param6 := []byte{'.', '.', '.', '7', '.', '.', '.', '.', '.'}
	param7 := []byte{'.', '.', '.', '5', '.', '.', '.', '.', '.'}
	param8 := []byte{'.', '.', '.', '.', '.', '.', '.', '.', '.'}
	param9 := []byte{'.', '.', '.', '.', '.', '.', '.', '.', '.'}

	params = append(params, param1)
	params = append(params, param2)
	params = append(params, param3)
	params = append(params, param4)
	params = append(params, param5)
	params = append(params, param6)
	params = append(params, param7)
	params = append(params, param8)
	params = append(params, param9)

	fmt.Println(isValidSudoku(params))
}
