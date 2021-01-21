package tools

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcalateMaxFloat64(t *testing.T) {
	_assert := assert.New(t)
	var data = []float64{1, 2, 3, 4, 5, 6, 7, 6, 5, 9, 3, 3}
	var max = CalcalateMaxFloat64(data)
	_assert.Equal(float64(9), max)
}

func TestCalculateMinFloat64(t *testing.T) {
	_assert := assert.New(t)
	var data = []float64{1, 2, 3, 4, 5, 6, 7, 6, 5, 9, 3, 3}
	var max = CalcalateMinFloat64(data)
	_assert.Equal(float64(1), max)
}
