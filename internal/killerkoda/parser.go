package killerkoda

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/hashicorp/go-multierror"
)

// Collection is the struct for the collection of scenarios
type Collection struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Items       []Item `json:"items"`
	Path        string `json:"-"`
}

// Scenarios returns all scenarios of a collection
func (c Collection) Scenarios() []Scenario {
	var scenarios []Scenario
	for _, item := range c.Items {
		if item.Scenario != nil {
			scenarios = append(scenarios, *item.Scenario)
		}
		if item.Collection != nil {
			scenarios = append(scenarios, item.Collection.Scenarios()...)
		}
	}
	return scenarios
}

// Item is the struct for the items of a collection
type Item struct {
	Path       string      `json:"path"`
	Scenario   *Scenario   `json:"-"`
	Collection *Collection `json:"-"`
}

// Scenario is the main struct for a scenario
type Scenario struct {
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Details     Details            `json:"details"`
	Backend     Backend            `json:"backend"`
	Interface   Interface          `json:"interface"`
	Assets      map[string][]Asset `json:"assets"`
	Path        string             `json:"-"`
}

// Details is the struct for the details of a scenario
type Details struct {
	Intro Intro  `json:"intro"`
	Steps []Step `json:"steps"`
}

// Intro is the struct for the intro of a scenario
type Intro struct {
	Text       string `json:"text"`
	Background string `json:"background"`
	Markdown   string `json:"-"`
}

// Step is the struct for the steps of a scenario
type Step struct {
	Text     string `json:"text"`
	Markdown string `json:"-"`
}

// Backend is the struct for the backend of a scenario
type Backend struct {
	ImageID string `json:"imageid"`
}

// Interface is the struct for the interface of a scenario
type Interface struct {
	Layout string `json:"layout"`
}

// Asset is the struct for the assets of a scenario
type Asset struct {
	File   string `json:"file"`
	Target string `json:"target"`
	Chmod  string `json:"chmod"`
}

// ParseCollection parses a collection
func ParseCollection(root string) (Collection, error) {
	var collection Collection
	err := unmarshalFile(getStructurePath(root), &collection)
	if err != nil {
		return collection, fmt.Errorf("reading structure.json failed: %w", err)
	}
	var mErr error
	for i := 0; i < len(collection.Items); i++ {
		item := &collection.Items[i]
		itemPath := filepath.Join(root, item.Path)

		if hasFile(getStructurePath(itemPath)) {
			v, err := ParseCollection(itemPath)
			if err != nil {
				mErr = multierror.Append(mErr, fmt.Errorf("parsing scenario failed: %w", err))
				continue
			}
			item.Collection = &v
		} else if hasFile(getIndexPath(itemPath)) {
			v, err := ParseScenario(itemPath)
			if err != nil {
				mErr = multierror.Append(mErr, fmt.Errorf("parsing collection failed: %w", err))
				continue
			}
			item.Scenario = &v
		} else {
			continue
		}
	}
	collection.Path = root
	return collection, mErr
}

// ParseScenario parses a scenario
func ParseScenario(root string) (Scenario, error) {
	var scenario Scenario
	indexPath := getIndexPath(root)
	err := unmarshalFile(indexPath, &scenario)
	if err != nil {
		return scenario, fmt.Errorf("reading index.json failed: %w", err)
	}

	introPath := filepath.Join(root, scenario.Details.Intro.Text)
	if hasFile(introPath) {
		content, err := os.ReadFile(introPath)
		if err != nil {
			return scenario, fmt.Errorf("reading %q failed: %w", introPath, err)
		}
		scenario.Details.Intro.Markdown = string(content)
	}

	var mErr error
	for i := 0; i < len(scenario.Details.Steps); i++ {
		step := &scenario.Details.Steps[i]
		stepPath := filepath.Join(root, step.Text)
		if hasFile(stepPath) {
			content, err := os.ReadFile(stepPath)
			if err != nil {
				mErr = multierror.Append(mErr, fmt.Errorf("reading %q failed: %w", stepPath, err))
				continue
			}
			step.Markdown = string(content)
		}
	}
	scenario.Path = root
	return scenario, mErr
}

func getStructurePath(root string) string {
	return filepath.Join(root, "structure.json")
}

func getIndexPath(root string) string {
	return filepath.Join(root, "index.json")
}

func hasFile(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil
}

func unmarshalFile(fileName string, v interface{}) error {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("reading %q failed: %w", fileName, err)
	}
	err = json.Unmarshal(content, v)
	if err != nil {
		return fmt.Errorf("unmarshalling %q failed: %w", fileName, err)
	}
	return nil
}
