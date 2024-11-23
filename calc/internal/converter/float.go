package converter

import (
	"math/big"
)

const (
	SetPrec = 256
	MaxPrec = 50
)

func ToBigFloatFromString(numStr string) *big.Float {
	num := new(big.Float)
	num.SetPrec(SetPrec).SetString(numStr)
	//fmt.Println("log: converted str ", numStr, num.Text('f', MaxPrec))
	return num
}
