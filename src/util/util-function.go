package util

import (
	"log"

	"github.com/northbright/keygen"
)

func GetScretKey() []byte {
	var key []byte
	var err error

	key, err = keygen.GenSymmetricKey(256)
	if err != nil {

		log.Printf("logInfo :\n key is %v\n err is %v", key, err)
		return key
	}

	return key
}
