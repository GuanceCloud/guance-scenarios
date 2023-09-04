package killerkoda

import (
	"fmt"
	"strings"

	"github.com/shurcooL/markdownfmt/markdown"
)

type Generator struct {
}

func (g *Generator) GenerateScenarioDocs(scenario Scenario) (string, error) {
	blocks := []string{
		fmt.Sprintf("# %s\n", scenario.Title),
		scenario.Description,
		scenario.Details.Intro.Markdown,
	}
	for i, step := range scenario.Details.Steps {
		blocks = append(blocks, fmt.Sprintf("## Step %d: %s\n", i, strings.TrimLeft(step.Markdown, "# ")))
	}

	result, err := markdown.Process("", []byte(strings.Join(blocks, "\n\n")), nil)
	if err != nil {
		return "", fmt.Errorf("formatting markdown failed: %w", err)
	}
	return string(result), nil
}
