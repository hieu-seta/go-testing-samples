package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestMockAdd demonstrates how to use a mock for dependency isolation.
func TestMockAdd(t *testing.T) {
	// Create a mock repository instance
	mockRepo := new(MockRepository)

	// Set expectation: when Add is called with [2, 3], return 5 and nil error
	mockRepo.On("Add", []int{2, 3}).Return(5, nil)

	// Use the mock as a MathService implementation
	testService := MathService(mockRepo)

	// Execute the test
	result, err := testService.Add([]int{2, 3})

	// Verify that all expectations were met
	mockRepo.AssertExpectations(t)

	// Assert the results
	assert.Equal(t, 5, result, "Expected sum to be 5")
	assert.Nil(t, err, "Expected no error")
}
