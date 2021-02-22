package p2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrint(t *testing.T) {
	output := driver()
	fmt.Println(output)
}

func TestResult(t *testing.T) {
	output := driver()
	assert.Equal(t, output, 3299)
}
