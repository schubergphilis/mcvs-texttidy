package application

import (
	"testing"

	"github.com/schubergphilis/mcvs-texttidy/internal/app/mcvs-texttidy/data"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCheckFilesForForbiddenWords(t *testing.T) {
	t.Parallel()

	forbidden := []string{"forbiddenWord", "forbiddenWORD"}

	testCases := []TestCase{
		{
			Name:           "BasicCheckWithForbiddenWords",
			Forbidden:      forbidden,
			Dirs:           []string{},
			Files:          []string{},
			ExpectedFiles:  1,
			ExpectedLines:  2,
			ExpectedErrMsg: "",
		},
		{
			Name:           "IgnoreSpecificFile",
			Forbidden:      forbidden,
			Dirs:           []string{},
			Files:          []string{"application_test.go"},
			ExpectedFiles:  0,
			ExpectedLines:  0,
			ExpectedErrMsg: "",
		},
		{
			Name:           "IgnoreCurrentDirectory",
			Forbidden:      forbidden,
			Dirs:           []string{"."},
			Files:          []string{},
			ExpectedFiles:  0,
			ExpectedLines:  0,
			ExpectedErrMsg: "",
		},
	}

	for _, testCase := range testCases {
		test := createTestCase(t, testCase)
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			stats := &data.ForbiddenWordStats{FileCount: 0, LineCount: 0}

			err := CheckFilesForForbiddenWords(test.config, stats)
			if test.expectedErrMsg != "" {
				assert.EqualError(t, err, test.expectedErrMsg)

				return
			}

			require.NoError(t, err)
			assert.Equal(t, test.expectedFiles, stats.FileCount, "unexpected file count")
			assert.Equal(t, test.expectedLines, stats.LineCount, "unexpected line count")
		})
	}
}
