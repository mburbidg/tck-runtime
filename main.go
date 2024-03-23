package main

import (
	"github.com/cucumber/godog"
	"github.com/opengql/tck/runtime"
	"log"
	"testing"
)

func main() {
	suite := godog.TestSuite{
		ScenarioInitializer: runtime.InitializeScenario,
		Options: &godog.Options{
			Format: "pretty",
			FeatureContents: []godog.Feature{
				{
					Name: "Test",
					Contents: []byte(
						`
							Feature: test step regular expressions
								Scenario:
									Given the catalog/catalog1 catalog
									Then an exception condition should be raised: G2000
						`,
					),
				},
			},
		},
	}

	if suite.Run() != 0 {
		log.Fatalf("non-zero status returned, failed to run feature tests\n")
	}
}

func TestGQLFeatures(t *testing.T) {
}
