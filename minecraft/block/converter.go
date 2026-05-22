package block

import (
	"github.com/Yeah114/bedrock-world-operator/block"
	"github.com/Yeah114/gophertunnel/minecraft/protocol"
	"github.com/Yeah114/worlddowngrader/blockdowngrader"
	"github.com/Yeah114/worldupgrader/blockupgrader"
)

// BlockConverter is a struct that can convert block states and runtime IDs between two protocol versions.
type BlockConverter struct {
	srcInfo  protocol.Info
	dstInfo  protocol.Info
	srcTable *block.BlockRuntimeIDTable
	dstTable *block.BlockRuntimeIDTable
}

// NewBlockConverter creates a new BlockConverter that can convert block states and runtime IDs between two protocol versions.
func NewBlockConverter(srcProtocol int32, srcTable *block.BlockRuntimeIDTable, dstProtocol int32, dstTable *block.BlockRuntimeIDTable) *BlockConverter {
	srcInfo := protocol.NewInfoByProtocol(srcProtocol)
	dstInfo := protocol.NewInfoByProtocol(dstProtocol)

	return &BlockConverter{
		srcInfo:  srcInfo,
		dstInfo:  dstInfo,
		srcTable: srcTable,
		dstTable: dstTable,
	}
}

// DstInfo returns the protocol info of the destination protocol version.
func (c *BlockConverter) SrcInfo() protocol.Info {
	return c.srcInfo
}

// DstInfo returns the protocol info of the destination protocol version.
func (c *BlockConverter) DstInfo() protocol.Info {
	return c.dstInfo
}

// SrcTable returns the block runtime ID table of the source protocol version.
func (c *BlockConverter) SrcTable() *block.BlockRuntimeIDTable {
	return c.srcTable
}

// DstTable returns the block runtime ID table of the destination protocol version.
func (c *BlockConverter) DstTable() *block.BlockRuntimeIDTable {
	return c.dstTable
}

// ConvertBlockState converts a block state from one protocol version to another.
// It returns the converted block state and a boolean indicating whether the conversion was successful.
func (c *BlockConverter) ConvertBlockState(name string, properties map[string]interface{}) (string, map[string]interface{}, bool) {
	if c.srcInfo.Version() < c.dstInfo.Version() {
		blockState := blockupgrader.BlockState{
			Name:       name,
			Properties: properties,
			Version:    c.srcInfo.Version(),
		}
		dstBlockState := blockupgrader.UpgradeToVersion(blockState, c.dstInfo.Ver())
		return dstBlockState.Name, dstBlockState.Properties, true
	} else if c.srcInfo.Version() > c.dstInfo.Version() {
		blockState := blockdowngrader.BlockState{
			Name:       name,
			Properties: properties,
			Version:    c.srcInfo.Version(),
		}
		dstBlockState := blockdowngrader.DowngradeToVersion(blockState, c.dstInfo.Ver())
		return dstBlockState.Name, dstBlockState.Properties, true
	}
	return name, properties, true
}

// ConvertBlockRuntimeID converts a block runtime ID from one protocol version to another.
// It returns the converted block runtime ID and a boolean indicating whether the conversion was successful.
func (c *BlockConverter) ConvertBlockRuntimeID(runtimeID uint32) (uint32, bool) {
	name, properties, found := c.srcTable.RuntimeIDToState(runtimeID)
	if !found {
		return 0, false
	}

	dstName, dstProperties, found := c.ConvertBlockState(name, properties)
	if !found {
		return 0, false
	}

	return c.dstTable.StateToRuntimeID(dstName, dstProperties)
}
