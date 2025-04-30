package utils

import "math/big"

func EthTOWei(amount float64) *big.Int {
	return big.NewInt(int64(amount * 1e18))
}
