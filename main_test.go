package main

import (
	"flag"
	"os"
	"testing"

	"github.com/ONSdigital/dp-charts-api/features/steps"
	componenttest "github.com/ONSdigital/dp-component-test"
	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
)

var componentFlag = flag.Bool("component", false, "perform component tests")

type ComponentTest struct {
	MongoFeature *componenttest.MongoFeature
}

func (f *ComponentTest) InitializeScenario(ctx *godog.ScenarioContext) {
	component, err := steps.NewComponent()
	if err != nil {
		panic(err)
	}

	ctx.BeforeScenario(func(*godog.Scenario) {
		component.Reset()
	})

	ctx.AfterScenario(func(*godog.Scenario, error) {
		_ = component.Close()
	})

	component.RegisterSteps(ctx)
}

func (f *ComponentTest) InitializeTestSuite(ctx *godog.TestSuiteContext) {

}

func TestComponent(t *testing.T) {
	if *componentFlag {
		status := 0

		var opts = godog.Options{
			Output: colors.Colored(os.Stdout),
			Format: "pretty",
			Paths:  flag.Args(),
		}

		f := &ComponentTest{}

		status = godog.TestSuite{
			Name:                 "feature_tests",
			ScenarioInitializer:  f.InitializeScenario,
			TestSuiteInitializer: f.InitializeTestSuite,
			Options:              &opts,
		}.Run()

		if status > 0 {
			t.Fail()
		}
	} else {
		t.Skip("component flag required to run component tests")
	}
}
