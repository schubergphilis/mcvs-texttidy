package application

import (
	"testing"

	"github.com/schubergphilis/mcvs-texttidy/internal/app/mcvs-texttidy/data"
)

type TestCase struct {
	Name           string
	Forbidden      []string
	Dirs           []string
	Files          []string
	ExpectedFiles  int
	ExpectedLines  int
	ExpectedErrMsg string
}

func createTestConfig(forbidden []string, dirsToIgnore []string, filesToIgnore []string) *data.Config {
	return &data.Config{
		Forbidden: forbidden,
		Ignore: data.IgnoreConfig{
			Words: make(map[string][]data.FileLines),
			Dirs:  dirsToIgnore,
			Files: filesToIgnore,
		},
	}
}

func createTestCase(t *testing.T, testCase TestCase) struct {
	name           string
	config         *data.Config
	expectedFiles  int
	expectedLines  int
	expectedErrMsg string
} {
	t.Helper()

	return struct {
		name           string
		config         *data.Config
		expectedFiles  int
		expectedLines  int
		expectedErrMsg string
	}{
		name:           testCase.Name,
		config:         createTestConfig(testCase.Forbidden, testCase.Dirs, testCase.Files),
		expectedFiles:  testCase.ExpectedFiles,
		expectedLines:  testCase.ExpectedLines,
		expectedErrMsg: testCase.ExpectedErrMsg,
	}
}
