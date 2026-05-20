package utils

import (
	"fmt"
	"bytes"

	"github.com/Yeah114/bedrock-world-operator/define"
	"github.com/Yeah114/gophertunnel/minecraft/nbt"
)

// DecodeBlockStates decodes block states from the given byte slice.
func DecodeBlockStates(blockStatesBytes []byte) (blockStates []define.BlockState) {
	dec := nbt.NewDecoder(bytes.NewBuffer(blockStatesBytes))
	for {
		var s define.BlockState
		if err := dec.Decode(&s); err != nil {
			if _, ok := err.(nbt.BufferOverrunError); ok {
				break
			}
			panic(fmt.Errorf("DecodeBlockStates: Failed to decode block state from NBT: %v", err))
		}
		blockStates = append(blockStates, s)
	}
	return blockStates
}
