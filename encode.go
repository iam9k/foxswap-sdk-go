package foxswap

import (
	"encoding/base64"
	"github.com/vmihailenco/msgpack/v5"
)

const (
	TransactionTypeAdd    = "Add"
	TransactionTypeRemove = "Remove"
	TransactionTypeSwap   = "Swap"
)

type TransactionAction struct {
	Type    string `json:"t,omitempty" msgpack:"t,omitempty"`
	AssetID string `json:"a,omitempty" msgpack:"a,omitempty"`
	Minimum string `json:"m,omitempty" msgpack:"m,omitempty"`
}

func EncodeAction(action TransactionAction) string {
	b, err := msgpack.Marshal(action)
	if err != nil {
		panic(err)
	}

	return base64.StdEncoding.EncodeToString(b)
}
