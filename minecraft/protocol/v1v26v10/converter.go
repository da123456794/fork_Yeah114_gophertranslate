package v1v26v10

import (
	"github.com/Yeah114/gophertranslate/minecraft/block"
	"github.com/Yeah114/gophertranslate/minecraft/chunk"
	//"github.com/Yeah114/gophertunnel/minecraft/protocol"
)

type ProtocolConverter struct {
	bc *block.BlockConverter
	cc *chunk.ChunkConverter
}

func NewProtocolConverter(bc *block.BlockConverter) *ProtocolConverter {
	return &ProtocolConverter{
		bc: bc,
		cc: chunk.NewChunkConverter(bc),
	}
}
