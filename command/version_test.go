package command

import (
	"fmt"
	"testing"

	"github.com/hashicorp/packer/version"
	"github.com/mitchellh/cli"
	"github.com/stretchr/testify/assert"
)

func TestVersionCommand_implements(t *testing.T) {
	var _ cli.Command = &VersionCommand{}
}

func Test_version(t *testing.T) {
	tc := []struct {
		command  []string
		env      []string
		expected string
	}{
		{[]string{"version"}, nil, fmt.Sprintf("Packer v%s", version.FormattedVersion()) + "\n"},
		{[]string{"version", "&"}, nil, fmt.Sprintf("Packer v%s", version.FormattedVersion()) + "\n"},
	}

	for _, tc := range tc {
		t.Run(fmt.Sprintf("packer %s", tc.command), func(t *testing.T) {
			p := helperCommand(t, tc.command...)
			bs, err := p.Output()
			fmt.Println(err)
			if err != nil {
				t.Fatalf("%v: %s", err, bs)
			}
			assert.Equal(t, tc.expected, string(bs))
		})
	}
}
