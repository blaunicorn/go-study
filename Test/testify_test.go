// 在工程里新建一个Test目录，创建一个文件testtestify_test.go，必须以_test结尾，代码如下

package Test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func main() {
	fmt.Println("hello world!")
}

// 计算并返回 x + 2.
func Calculate(x int) (result int) {
	result = x + 2
	return
}
func TestCalculate1(t *testing.T) {
	assert.Equal(t, Calculate(2), 4)
}

func TestCalculate2(t *testing.T) {
	assert := assert.New(t)
	var tests = []struct {
		input    int
		expected int
	}{
		{2, 4},
		{-1, 1},
		{0, 2},
		{-5, -3},
		{99999, 100001},
	}
	for _, test := range tests {
		assert.Equal(Calculate(test.input), test.expected)
	}
}
