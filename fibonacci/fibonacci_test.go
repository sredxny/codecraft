package fibonacci

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Fibonacci(n int) int {

	var n3, n1, n2 int = 0, 0, 1

	for i := 1; i <= n; i++ {


		n3 = n1 + n2

		n1 = n2

		n2 = n3

	}
	return n1
}

func Test_Fibonacci(t *testing.T){

	for i:=0;i <= 49;i++{
		expected := 0
		if i == 0 {
			expected = 0
		}else if i == 1{
			expected = 1
		}else{
			expected= Fibonacci(i-2) + Fibonacci(i-1)
		}

		assert.Equal(t,expected,Fibonacci(i))
	}

}
