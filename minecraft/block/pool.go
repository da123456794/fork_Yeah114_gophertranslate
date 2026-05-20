package block

import (
	"github.com/Yeah114/bedrock-world-operator/block"
	"github.com/Yeah114/gophertunnel/minecraft/protocol"
	"github.com/Yeah114/gophertranslate/minecraft/block/v1v26v20"
)

// Pool holds functions that create BlockRuntimeIDTables for different Minecraft versions. The key is the protocol version.
var Pool = map[int32]func(bool) *block.BlockRuntimeIDTable{
	protocol.Protocol1v26v20: v1v26v20.NewBlockRuntimeIDTable,
}
