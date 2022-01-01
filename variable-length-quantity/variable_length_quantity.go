package variablelengthquantity

import "fmt"

func EncodeVarint(input []uint32) []byte {
	result := make([]byte, 0)

	fmt.Println(input)

	return result
}

func DecodeVarint(input []byte) ([]uint32, error) {
	result := make([]uint32, 0)

	fmt.Println(input)

	return result, nil
}
