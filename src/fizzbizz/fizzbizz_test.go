package fizzbizz

import (
	"testing"

	"github.com/go-playground/assert"

	"example_module/src/fizzbizz"
)

func TestMain(t *testing.T) {

	got := fizzbizz.Sample(5)
	assert.Equal(t, got, []string{"1", "2", "Fizz", "4", "Buzz"})
}
