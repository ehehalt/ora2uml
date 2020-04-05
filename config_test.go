package ora2uml

import (
	"testing"
)

func TestConfigRead(t *testing.T) {
	filename := "test/sample.json"

	config, err := Read(filename)

	if err != nil {
		t.Errorf("File %v should readed without errors", filename)
	}

	if config.User.UserId != "system" {
		t.Errorf("User.UserId should be 'system' but was '%v'", config.User.UserId)
	}

	if config.User.Password != "sysadm" {
		t.Errorf("User.Password should be 'sysadm' but was '%v'", config.User.Password)
	}
}
