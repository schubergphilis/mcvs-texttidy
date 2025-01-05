package main

import (
	"github.com/schubergphilis/mcvs-texttidy/internal/app/mcvs-texttidy/presentation"
	log "github.com/sirupsen/logrus"
)

func main() {
	forbiddenWordStats, err := presentation.CLI()
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Fatal("unable to run CLI")
	}

	if forbiddenWordStats.FileCount > 0 {
		log.WithFields(log.Fields{
			"numberOfFiles":                  forbiddenWordStats.FileCount,
			"numberOfLinesContainingTheWord": forbiddenWordStats.LineCount,
		}).Fatal("found forbidden words in files")
	}
}
