package p1

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrint(t *testing.T) {
	output := driver()
	fmt.Println(output)
}

func TestResult(t *testing.T) {
	output := driver()
	assert.Equal(t, output, 202)
}
