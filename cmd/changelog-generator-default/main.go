package main

import (
	defaultGenerator "github.com/ted-vo/changelog-generator-default/pkg/generator"
	"github.com/ted-vo/semantic-release/v3/pkg/generator"
	"github.com/ted-vo/semantic-release/v3/pkg/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ChangelogGenerator: func() generator.ChangelogGenerator {
			return &defaultGenerator.DefaultChangelogGenerator{}
		},
	})
}
