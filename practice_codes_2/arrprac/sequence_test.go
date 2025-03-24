package arrprac

import (
	"fmt"
	"testing"
)

func TestBracketSequence(t *testing.T) {
	//valid := make([]byte, 5)
	byteVal := []byte("()()(")
	//valid = append(valid, byteVal)
	//_ = byteVal

	count := BracketSequence(byteVal)
	fmt.Println("valid parantheseis count: ", count)

	// if count != "valid sequence" {
	// 	t.Error("Invalid sequence passed")
	// }

}
