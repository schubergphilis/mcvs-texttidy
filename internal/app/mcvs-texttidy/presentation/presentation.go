package presentation

import (
	"flag"
	"fmt"
	"os"

	"github.com/schubergphilis/mcvs-texttidy/internal/app/mcvs-texttidy/application"
	"github.com/schubergphilis/mcvs-texttidy/internal/app/mcvs-texttidy/data"
	"github.com/schubergphilis/mcvs-texttidy/internal/pkg/constants"
	log "github.com/sirupsen/logrus"
)

func CLI() (*data.ForbiddenWordStats, error) {
	showVersion := flag.Bool("version", false, "display the version of the application")
	flag.Parse()

	if *showVersion {
		log.Infof("version: '%s'", constants.AppVersion)
		os.Exit(0)
	}

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
