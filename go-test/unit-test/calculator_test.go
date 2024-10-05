package calc

import (
	"testing"

	asrt "github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	x, y := 2, 3

	c := Calculate{}

	actual := c.Add(x, y)

	expected := 5

	if actual != expected {
		t.Errorf("Calculate.Add(%d,%d) failed. Expected: %d Actual:%d", x, y, expected, actual)
	}
}

func TestDivide(t *testing.T) {
	assert := asrt.New(t)
	c := Calculate{}

	tables := []struct {
		x        float64
		y        float64
		expected float64
	}{
		{2, 2, 1},
		{6, 2, 3},
	}

	for _, v := range tables {
		actual := c.Divide(v.x, v.y)

		assert.Equal(v.expected, actual)
	}
}
