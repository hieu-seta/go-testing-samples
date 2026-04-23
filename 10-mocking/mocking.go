package tests

import (
	"github.com/stretchr/testify/mock"
)

// Real implementation

type Calculator struct{}

// Add sums a slice of integers and returns the total.
func (*Calculator) Add(numbers []int) (int, error) {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total, nil
}

// Mock implementation for testing

type MockRepository struct {
	mock.Mock
}

// Add is the mocked method matching the MathService interface.
func (m *MockRepository) Add(numbers []int) (int, error) {
	args := m.Called(numbers)
	return args.Get(0).(int), args.Error(1)
}

// MathService defines the behavior required for the tests.
type MathService interface {
	Add([]int) (int, error)
}
