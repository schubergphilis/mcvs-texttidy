package application

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/schubergphilis/mcvs-texttidy/internal/app/mcvs-texttidy/data"
)

func CheckFilesForForbiddenWords(
	config data.Config,
	stats *data.ForbiddenWordStats,
) error {
	dirsToSkip := makeSetFromSlice(config.Ignore.Dirs)
	filesToSkip := makeSetFromSlice(config.Ignore.Files)
	processor := &data.FileProcessor{
		ForbiddenWords: config.Forbidden,
		IgnoreWords:    config.Ignore.Words,
	}

	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if shouldSkipDir(info, path, dirsToSkip) {
			return filepath.SkipDir
		}

		if shouldSkipFile(info, path, filesToSkip) {
			return nil
		}

		if !info.IsDir() {
			return processor.ProcessFile(path, stats)
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("unable to check files for forbidden words: %w", err)
	}

	return nil
}

func makeSetFromSlice(items []string) map[string]struct{} {
	set := make(map[string]struct{})
	for _, item := range items {
		set[filepath.Clean(item)] = struct{}{}
	}

	return set
}

func shouldSkipDir(info os.FileInfo, path string, dirsToSkip map[string]struct{}) bool {
	if !info.IsDir() {
		return false
	}

	if _, exists := dirsToSkip[filepath.Clean(path)]; exists {
		return true
	}

	return false
}

func shouldSkipFile(info os.FileInfo, path string, filesToSkip map[string]struct{}) bool {
	if info.IsDir() {
		return false
	}

	if _, exists := filesToSkip[filepath.Clean(path)]; exists {
		return true
	}

	return false
}
