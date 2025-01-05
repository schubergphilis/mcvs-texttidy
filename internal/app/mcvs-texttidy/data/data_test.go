package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldIgnoreWord(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name         string
		filePath     string
		lineNumber   int
		ignoreList   []FileLines
		expectIgnore bool
	}{
		{
			name:         "Ignore existing line",
			filePath:     "example.go",
			lineNumber:   10,
			ignoreList:   []FileLines{{File: "example.go", Lines: []int{5, 10, 15}}},
			expectIgnore: true,
		},
		{
			name:         "Do not ignore nonexistent line",
			filePath:     "example.go",
			lineNumber:   20,
			ignoreList:   []FileLines{{File: "example.go", Lines: []int{5, 10, 15}}},
			expectIgnore: false,
		},
		{
			name:         "Do not ignore different file",
			filePath:     "other.go",
			lineNumber:   10,
			ignoreList:   []FileLines{{File: "example.go", Lines: []int{5, 10, 15}}},
			expectIgnore: false,
		},
		{
			name:         "Ignore with multiple entries",
			filePath:     "example.go",
			lineNumber:   15,
			ignoreList:   []FileLines{{File: "example.go", Lines: []int{5}}, {File: "example.go", Lines: []int{15}}},
			expectIgnore: true,
		},
		{
			name:         "Empty ignore list",
			filePath:     "example.go",
			lineNumber:   10,
			ignoreList:   []FileLines{},
			expectIgnore: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			actualIgnore := shouldIgnoreWord(testCase.filePath, testCase.lineNumber, testCase.ignoreList)
			assert.Equal(t, testCase.expectIgnore, actualIgnore)
		})
	}
}
