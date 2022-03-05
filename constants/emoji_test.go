package constants

import (
	"testing"
)

func TestZombieMapsToEmoji(t *testing.T) {
	for _, zMap := range ZombieMapsArr {
		_, err := ZMapToEmoji(zMap)
		if err != nil {
			t.Errorf("Expected nil but received %s", err)
		}
	}

	_, err := ZMapToEmoji("invalidMap")
	if err == nil {
		t.Errorf("Expected error but received nil")
	}
}