package main

import (
	"math/rand"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/hashicorp/packer/command"
	"github.com/mitchellh/cli"
)

func TestExcludeHelpFunc(t *testing.T) {
	commands := map[string]cli.CommandFactory{
		"build": func() (cli.Command, error) {
			return &command.BuildCommand{
				Meta: command.Meta{},
			}, nil
		},

		"fix": func() (cli.Command, error) {
			return &command.FixCommand{
				Meta: command.Meta{},
			}, nil
		},
	}

	helpFunc := excludeHelpFunc(commands, []string{"fix"})
	helpText := helpFunc(commands)

	if strings.Contains(helpText, "fix") {
		t.Fatalf("Found fix in help text even though we excluded it: \n\n%s\n\n", helpText)
	}
}

func TestExtractMachineReadable(t *testing.T) {
	var args, expected, result []string
	var mr bool

	// Not
	args = []string{"foo", "bar", "baz"}
	result, mr = extractMachineReadable(args)
	expected = []string{"foo", "bar", "baz"}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("bad: %#v", result)
	}

	if mr {
		t.Fatal("should not be mr")
	}

	// Yes
	args = []string{"foo", "-machine-readable", "baz"}
	result, mr = extractMachineReadable(args)
	expected = []string{"foo", "baz"}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("bad: %#v", result)
	}

	if !mr {
		t.Fatal("should be mr")
	}
}

func TestRandom(t *testing.T) {
	if rand.Intn(9999999) == 8498210 {
		t.Fatal("math.rand is not seeded properly")
	}
}

func ExampleWrappedMain() {
	os.Setenv("PACKER_WRAP_COOKIE", "49C22B1A-3A93-4C98-97FA-E07D18C787B5")
	backgroundCheckFn = func(_ int) (bool, error) { return true, nil }
	os.Args = []string{"packer", "version"}
	wrappedMain()
	//Output: Packer v1.7.2-dev
}
