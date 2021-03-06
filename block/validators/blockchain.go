package validators

import (
	"blockchain/block"
	"bytes"
	"fmt"
)

func Validate(bc *block.Blockchain) bool {
	for _, b := range bc.GetBlocks() {
		var f *block.Block = &block.Block{}
		if bytes.Compare(f.GetPrevHash(), ([]byte{})) != 0 {
			if bytes.Compare(f.GetPrevHash(), b.GetHash()) != 0 {
				fmt.Println(f, b)
				return false
			}
		}
		f = b
	}

	return true
}
