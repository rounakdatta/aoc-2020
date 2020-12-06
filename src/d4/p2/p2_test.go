package p2

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrint(t *testing.T) {
	output := driver()
	fmt.Println(output)
}

func TestHgtRegexValidator(t *testing.T) {
	assert.Equal(t, regexValidator("^((1[5-8][0-9]|19[0-3])cm|(59|6[0-9]|7[0-6])in)$", "193cm"), 1)
	assert.Equal(t, regexValidator("^((1[5-8][0-9]|19[0-3])cm|(59|6[0-9]|7[0-6])in)$", "79in"), 0)
}

func TestHclRegexValidator(t *testing.T) {
	assert.Equal(t, regexValidator("^#([0-9]|[a-f]){6}$", "#7fe94a"), 1)
}

func TestResult(t *testing.T) {
	output := driver()
	assert.Equal(t, output, 137)
}
