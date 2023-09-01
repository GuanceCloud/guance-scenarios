package killerkoda

import (
	"fmt"
	"testing"
)

func TestSnapshot(t *testing.T) {
	root, err := ParseCollection("./testdata/input")
	if err != nil {
		t.Fatal(err)
	}

	for _, scenario := range root.Scenarios() {
		t.Run(scenario.Title, func(t *testing.T) {
			s, err := (&Generator{}).GenerateScenarioDocs(scenario)
			if err != nil {
				t.Fatal(err)
			}

			fmt.Println(s)
		})
	}
}
