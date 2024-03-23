package runtime

import (
	"context"
	"github.com/cucumber/godog"
)

func beforeScenario(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
	return ctx, nil
}

func afterScenario(ctx context.Context, sc *godog.Scenario, err error) (context.Context, error) {
	return ctx, nil
}

func anEmptyCatalog(ctx context.Context) (context.Context, error) {
	return ctx, nil
}

func anyCatalog(ctx context.Context) (context.Context, error) {
	return ctx, nil
}

func namedCatalog(ctx context.Context, path string) (context.Context, error) {
	return ctx, nil
}

func theOmittedGraph(ctx context.Context) (context.Context, error) {
	return ctx, nil
}

func anEmptyGraph(ctx context.Context) (context.Context, error) {
	return ctx, nil
}

func anyGraph(ctx context.Context) (context.Context, error) {
	return ctx, nil
}

func namedGraph(ctx context.Context, path string) (context.Context, error) {
	return ctx, nil
}

func theGraphCreatedByExecutingTheProgram(ctx context.Context, program *godog.DocString) (context.Context, error) {
	return ctx, nil
}

func havingExecutedTheProgram(ctx context.Context, program *godog.DocString) (context.Context, error) {
	return ctx, nil
}

func executingTheProgram(ctx context.Context, program *godog.DocString) (context.Context, error) {
	return ctx, nil
}

func theResultShouldBeEmpty(ctx context.Context) (context.Context, error) {
	return ctx, nil
}

func theResultShouldBeInAnyOrder(ctx context.Context, table *godog.Table) (context.Context, error) {
	return ctx, nil
}

func theSideEffectsShouldBe(ctx context.Context, table *godog.Table) (context.Context, error) {
	return ctx, nil
}

func anExceptionConditionShouldBeRaised(ctx context.Context, gqlStatus string) (context.Context, error) {
	return ctx, nil
}

func InitializeScenario(sc *godog.ScenarioContext) {
	sc.Before(beforeScenario)
	sc.After(afterScenario)
	sc.Given(`^an empty catalog$`, anEmptyCatalog)
	sc.Given(`^any catalog$`, anyCatalog)
	sc.Given(`^the ([a-zA-Z0-9\-\_\/]+) catalog$`, namedCatalog)
	sc.Given(`^the omitted graph$`, theOmittedGraph)
	sc.Given(`^an empty graph$`, anEmptyGraph)
	sc.Given(`^any graph$`, anyGraph)
	sc.Given(`^the ([a-zA-Z0-9\-\_\/]+) graph$`, namedGraph)
	sc.Given(`^the graph created by executing program:$`, theGraphCreatedByExecutingTheProgram)
	sc.Step(`^having executed the program:$`, havingExecutedTheProgram)
	sc.When(`^executing the program:$`, executingTheProgram)
	sc.Then(`^the result should be empty$`, theResultShouldBeEmpty)
	sc.Step(`^the result should be, in any order:$`, theResultShouldBeInAnyOrder)
	sc.Step(`^the side effects should be:$`, theSideEffectsShouldBe)
	sc.Then(`^an exception condition should be raised: ([A-Z0-9]{5})`, anExceptionConditionShouldBeRaised)
}
