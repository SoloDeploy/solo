package cmd

import (
	"testing"

	"github.com/SoloDeploy/solo/core/configuration"
)

func TestNewCmdSolo(t *testing.T) {
	configuration := configuration.Configuration{}
	cmd := NewCmdSolo(&configuration)

	if cmd.Use != "solo" {
		t.Errorf("Use is not correct")
	}
}
