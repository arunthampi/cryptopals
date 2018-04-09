package set1

import (
	"bytes"
	"fmt"
)

func HexToBase64(input *bytes.Buffer) *bytes.Buffer {
	inputBytes := input.Bytes()
	outputBytes := bytes.NewBuffer([]byte{})

	for _, i := range inputBytes {
		fmt.Printf("%d", i)
	}

	fmt.Println("")

	return outputBytes
}
