package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"strings"

	"github.com/MaineK00n/vuls-data-update/pkg/fetch/suse/oval"
)

type root struct {
	Generator struct {
		ProductName   string `xml:"product_name"`
		SchemaVersion string `xml:"schema_version"`
		Timestamp     string `xml:"timestamp"`
	} `xml:"generator" json:"generator,omitempty"`
	Definitions struct {
		Definition []oval.Definition `xml:"definition" json:"definition,omitempty"`
	} `xml:"definitions" json:"definitions,omitempty"`
	Tests   oval.Tests   `xml:"tests" json:"tests,omitempty"`
	Objects oval.Objects `xml:"objects" json:"objects,omitempty"`
	States  oval.States  `xml:"states" json:"states,omitempty"`
}

type m struct {
	dm map[string]oval.Definition
	tm map[string]oval.RpminfoTest
	om map[string]oval.RpminfoObject
	sm map[string]oval.RpminfoState
}

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Please provide a file path as an argument.")
		return
	}
	filePath := os.Args[1]
	ref := os.Args[2]

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var root root
	if err := xml.NewDecoder(file).Decode(&root); err != nil {
		fmt.Println("Error decoding XML:", err)
		return
	}

	m := m{
		dm: make(map[string]oval.Definition),
		tm: make(map[string]oval.RpminfoTest),
		om: make(map[string]oval.RpminfoObject),
		sm: make(map[string]oval.RpminfoState),
	}

	for _, def := range root.Definitions.Definition {
		m.dm[def.ID] = def
	}
	for _, test := range root.Tests.RpminfoTest {
		m.tm[test.ID] = test
	}
	for _, object := range root.Objects.RpminfoObject {
		m.om[object.ID] = object
	}
	for _, state := range root.States.RpminfoState {
		m.sm[state.ID] = state
	}

	ss := strings.Split(ref, ":")
	if len(ss) != 4 {
		fmt.Println("id invalid format")
		return
	}
	kind := ss[2]
	switch kind {
	case "def":
		def, found := m.dm[ref]
		if !found {
			fmt.Println("Definition not found:", ref)
			return
		}
		// fmt.Printf("======= def: %+v\n", def)
		printCriteria(m, def.Criteria)
	}
}

func printCriteria(m m, criteria oval.Criteria) {
	for _, c := range criteria.Criterias {
		printCriteria(m, c)
	}

	for _, c := range criteria.Criterions {
		fmt.Printf("Criterion: %s\n", c.Comment)
		t, found := m.tm[c.TestRef]
		if !found {
			fmt.Println("Test not found:", c.TestRef)
			continue
		}
		o, found := m.om[t.Object.ObjectRef]
		if !found {
			fmt.Println("Object not found:", t.Object.ObjectRef)
			continue
		}
		s, found := m.sm[t.State.StateRef]
		if !found {
			fmt.Println("State not found:", t.State.StateRef)
			continue
		}

		fmt.Printf("  Object Name:    %q\n", o.Name)
		fmt.Printf("  Object Version: %q\n", o.Version)
		fmt.Printf("  State ID:       %q\n", s.ID)
		fmt.Printf("  State Version:  %q (op: %q)\n", s.Version.Text, s.Version.Operation)
		fmt.Printf("  State EVR:      %q (op: %q)\n", s.Evr.Text, s.Evr.Operation)
		fmt.Printf("  State Arch:     %q (op: %q)\n", s.Arch.Text, s.Arch.Operation)
	}
}
