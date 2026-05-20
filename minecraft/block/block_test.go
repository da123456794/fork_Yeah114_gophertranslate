package block

import (
	"testing"

	"github.com/Yeah114/gophertranslate/minecraft/block/v1v26v20"
)

func TestBlockRuntimeIDTable(t *testing.T) {
	// v1.26.20
	table := v1v26v20.NewBlockRuntimeIDTable(false)
	airRuntimeID, found := table.StateToRuntimeID("minecraft:air", nil)
	if !found {
		t.Fatal("Failed to find runtime ID for air block")
	}
	if airRuntimeID != table.AirRuntimeID() {
		t.Fatalf("Expected air block runtime ID to be %d, got %d", table.AirRuntimeID(), airRuntimeID)
	}
	t.Logf("Air block runtime ID: %d", airRuntimeID)
}
