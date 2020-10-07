package main

import (
	"crypto/rand"
	"math/big"
)

func getRandomNumber(total int64) (int64, error) {
	randomNumber, err := rand.Int(rand.Reader, big.NewInt(total))
	if err != nil {
		return -1, err
	}
	return randomNumber.Int64(), nil
}
