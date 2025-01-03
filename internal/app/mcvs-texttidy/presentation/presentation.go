package presentation

import (
	"fmt"

	"github.com/schubergphilis/mcvs-texttidy/internal/app/mcvs-texttidy/application"
	"github.com/schubergphilis/mcvs-texttidy/internal/app/mcvs-texttidy/data"
)

func CLI() (*data.ForbiddenWordStats, error) {
	config, err := data.ParseYAMLConfig()
	if err != nil {
		return nil, fmt.Errorf("error parsing YAML configuration: %w", err)
	}

	var stats data.ForbiddenWordStats

	if err := application.CheckFilesForForbiddenWords(
		config,
		&stats,
	); err != nil {
		return nil, fmt.Errorf("error during forbidden word check: %w", err)
	}

	return &stats, nil
}
