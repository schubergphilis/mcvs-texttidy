package data

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

const textTidyIgnoreFile = ".mcvs-texttidy.yml"

type Config struct {
	Forbidden []string `yaml:"forbidden"`
	Ignore    struct {
		Dirs  []string               `yaml:"dirs"`
		Files []string               `yaml:"files"`
		Words map[string][]FileLines `yaml:"words"`
	} `yaml:"ignore"`
}

type FileLines struct {
	File  string `yaml:"file"`
	Lines []int  `yaml:"lines"`
}

type ForbiddenWordStats struct {
	FileCount int
	LineCount int
}

type FileProcessor struct {
	ForbiddenWords []string
	IgnoreWords    map[string][]FileLines
}

func ParseYAMLConfig() (Config, error) {
	file, err := os.Open(filepath.Clean(textTidyIgnoreFile))
	if err != nil {
		return Config{}, fmt.Errorf("failed to open config file: %s. Error: %w", textTidyIgnoreFile, err)
	}
	defer file.Close()

	fileContent, err := io.ReadAll(file)
	if err != nil {
		return Config{}, fmt.Errorf("failed to read config file: %s. Error: %w", textTidyIgnoreFile, err)
	}

	var config Config
	if err := yaml.Unmarshal(fileContent, &config); err != nil {
		return Config{}, fmt.Errorf("failed to unmarshal config file: %s. Error: %w", textTidyIgnoreFile, err)
	}

	return config, nil
}

func (fp *FileProcessor) ProcessFile(filePath string, stats *ForbiddenWordStats) (err error) {
	file, err := os.Open(filepath.Clean(filePath))
	if err != nil {
		return fmt.Errorf("unable to open file %s: %w", filePath, err)
	}

	defer func() {
		if cerr := file.Close(); cerr != nil {
			err = fmt.Errorf("unable to close file: %s. Error: %w", filePath, cerr)
		}
	}()

	lineCount, err := countForbiddenWords(file, filePath, fp.ForbiddenWords, fp.IgnoreWords)
	if err != nil {
		return fmt.Errorf("error counting forbidden words in file %s: %w", filePath, err)
	}

	stats.LineCount += lineCount
	if lineCount > 0 {
		stats.FileCount++
	}

	return nil
}

func countForbiddenWords(
	file *os.File,
	filePath string,
	forbiddenWords []string,
	ignoreWords map[string][]FileLines,
) (int, error) {
	scanner := bufio.NewScanner(file)
	lineNumber := 1
	numberOfLinesWithForbiddenWord := 0

	for scanner.Scan() {
		line := scanner.Text()

		for _, word := range forbiddenWords {
			if strings.Contains(line, word) {
				if shouldIgnoreWord(filePath, lineNumber, ignoreWords[word]) {
					continue
				}

				log.Errorf("found forbidden word: '%s' in file: %s:%d", word, filePath, lineNumber)

				numberOfLinesWithForbiddenWord++
			}
		}

		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("non-EOF error during scan: %w", err)
	}

	return numberOfLinesWithForbiddenWord, nil
}

func shouldIgnoreWord(filePath string, lineNumber int, ignoreList []FileLines) bool {
	for _, fl := range ignoreList {
		if fl.File == filePath {
			for _, line := range fl.Lines {
				if line == lineNumber {
					return true
				}
			}
		}
	}

	return false
}
