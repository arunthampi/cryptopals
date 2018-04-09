package set1

import (
	"bytes"
	"testing"
)

func TestHexToBase64(t *testing.T) {
	input := bytes.NewBufferString("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	output := HexToBase64(input)
	realOutput := output.String()
	expectedOutput := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	if realOutput != expectedOutput {
		t.Fatalf("Expected HexToBase64 to return %s, got %s\n", realOutput, output.String())
	}
}
