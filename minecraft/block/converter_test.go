package block

import (
	"testing"

	"github.com/Yeah114/gophertunnel/minecraft/protocol"
)

func TestBlockConverterConvertBlockStateSameVersion(t *testing.T) {
	converter := newTestBlockConverter(t, protocol.Protocol1v26v20, protocol.Protocol1v26v20)
	properties := map[string]any{"foo": "bar"}
	name, convertedProperties, ok := converter.ConvertBlockState("minecraft:test", properties)
	if !ok {
		t.Fatal("expected conversion to succeed")
	}
	if name != "minecraft:test" {
		t.Fatalf("expected name to stay unchanged, got %q", name)
	}
	if convertedProperties["foo"] != "bar" {
		t.Fatalf("expected properties to stay unchanged, got %#v", convertedProperties)
	}
}

func TestBlockConverterConvertBlockRuntimeID(t *testing.T) {
	converter := newTestBlockConverter(t, protocol.Protocol1v26v10, protocol.Protocol1v26v20)

	srcRuntimeID, found := converter.srcTable.StateToRuntimeID("minecraft:air", nil)
	if !found {
		t.Fatal("failed to find source air runtime ID")
	}
	dstRuntimeID, ok := converter.ConvertBlockRuntimeID(srcRuntimeID)
	if !ok {
		t.Fatal("expected runtime ID conversion to succeed")
	}
	if dstRuntimeID != converter.dstTable.AirRuntimeID() {
		t.Fatalf("expected destination air runtime ID %d, got %d", converter.dstTable.AirRuntimeID(), dstRuntimeID)
	}
	if dstRuntimeID == srcRuntimeID {
		t.Fatalf("expected air runtime ID to change between versions, both were %d", dstRuntimeID)
	}
}

func TestBlockConverterConvertBlockRuntimeIDMissing(t *testing.T) {
	converter := newTestBlockConverter(t, protocol.Protocol1v26v10, protocol.Protocol1v26v20)
	if runtimeID, ok := converter.ConvertBlockRuntimeID(^uint32(0)); ok {
		t.Fatalf("expected conversion to fail for unknown runtime ID, got %d", runtimeID)
	}
}

func newTestBlockConverter(t *testing.T, srcProtocol, dstProtocol int32) *BlockConverter {
	t.Helper()
	srcConstructor, ok := Pool[srcProtocol]
	if !ok {
		t.Fatalf("missing source table constructor for protocol %d", srcProtocol)
	}
	dstConstructor, ok := Pool[dstProtocol]
	if !ok {
		t.Fatalf("missing destination table constructor for protocol %d", dstProtocol)
	}
	return NewBlockConverter(srcProtocol, srcConstructor(false), dstProtocol, dstConstructor(false))
}
