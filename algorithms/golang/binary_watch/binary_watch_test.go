package binary_watch

import (
	"fmt"
	"testing"
)

func TestReadBinaryWatch(t *testing.T) {
	fmt.Println(readBinaryWatch(1))
	fmt.Println(readBinaryWatch(9))
}
