package cmd

import (
	"testing"

	"github.com/SoloDeploy/solo/core/configuration"
)

func TestNewCmdSolo(t *testing.T) {
	configuration := configuration.Configuration{}
	cmd := NewCmdSolo(&configuration, "version", "commit", "date")

	if cmd.Use != "solo" {
		t.Errorf("Use is not correct")
	}
}
