package v1v26v20

import (
	"fmt"
	_ "embed"

	"github.com/Yeah114/bedrock-world-operator/block"
	"github.com/Yeah114/bedrock-world-operator/define"
	"github.com/Yeah114/gophertranslate/minecraft/block/utils"
)

var (
	//go:embed block_states.nbt
	blockStatesBytes []byte
	blockStates      []define.BlockState
)

func init() {
	blockStates = utils.DecodeBlockStates(blockStatesBytes)
}

// NewBlockRuntimeIDTable returns a new BlockRuntimeIDTable for Minecraft version 1.26.20.
func NewBlockRuntimeIDTable(useNetworkIDHashes bool) *block.BlockRuntimeIDTable {
	table := block.NewEmptyBlockRuntimeIDTable(useNetworkIDHashes)
	for _, block := range blockStates {
		fmt.Printf("Registering block state: %s with properties %+v\n", block.Name, block.Properties)
		err := table.RegisterCustomBlock(block)
		if err != nil {
			panic(fmt.Sprintf("v1.26.20.NewBlockRuntimeIDTable: Failed to register custom block: %v", err))
		}
	}
	table.FinaliseRegister()
	return table
}
