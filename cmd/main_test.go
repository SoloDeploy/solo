package cmd

import "testing"

func TestNewCmdSolo(t *testing.T) {
	cmd := NewCmdSolo()

	if cmd.Use != "solo" {
		t.Errorf("Use is not correct")
	}
}
