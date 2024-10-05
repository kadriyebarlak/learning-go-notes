package calc

import (
	"fmt"
	"testing"

	asrt "github.com/stretchr/testify/assert"
)

type TestCalculator struct {
	add      map[string]int
	subtract map[string]int
	divide   map[string]float64
}

func (c TestCalculator) On(f string, x, y int, rt float64) {
	switch f {
	case "Add":
		c.add[fmt.Sprintf("%d%d", x, y)] = int(rt)
	case "Subtract":
		c.subtract[fmt.Sprintf("%d%d", x, y)] = int(rt)
	case "Divide":
		c.divide[fmt.Sprintf("%d%d", x, y)] = rt
	}
}

func (c TestCalculator) Add(x, y int) int {
	return c.add[fmt.Sprintf("%d%d", x, y)]
}

func (c TestCalculator) Subtract(x, y int) int {
	return c.subtract[fmt.Sprintf("%d%d", x, y)]
}

func (c TestCalculator) Divide(x, y int) float64 {
	return c.divide[fmt.Sprintf("%d%d", x, y)]
}

func TestAmount(t *testing.T) {
	assert := asrt.New(t)

	testCalculator := TestCalculator{
		add:      map[string]int{},
		subtract: map[string]int{},
		divide:   map[string]float64{},
	}

	d := NewMinPriceDiscount(11, testCalculator)

	tables := []struct {
		amount   float64
		discount float64
		expected float64
	}{
		{100, 20, 80},
		{10, 20, 10},
		{100, 120, 100},
	}

	testCalculator.On("Subtract", 100, 20, 80)

	for _, v := range tables {
		actual := d.Amount(v.amount, v.discount)

		assert.Equal(v.expected, actual)
	}

}
