package solo

import "testing"

func TestNewCmdSolo(t *testing.T) {
	cmd := NewCmdSolo("0.0.0", "commitid", "date")

	if cmd.Use != "solo" {
		t.Errorf("Use is not correct")
	}
}
