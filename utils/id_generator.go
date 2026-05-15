package utils

import (
	"crypto/rand"

	"github.com/zahidmahfudz/collabforge-platform/config"
)

var Logger = config.Logger

const characterSet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

/*
fungsi untuk generate id dengan prefix dan panjang sebanyak 16 karakter setelah prefix

/ contoh:
// GenerateID("usr", 16)
// hasil:
// usr_a8Ks91LmPq2XzT7Q

*/

func GenerateID(prefix string) (string, error) {
	Logger.Debugf("memasuki fungsi GenerateID dengan Parameter prefix: %s", prefix)
	//panjang id setelah prefix
	length := 16

	//buat slice untuk menyimpan karakter-karakter id
	bytes := make([]byte, length)

	//generate karakter acak untuk id
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	//buat string builder untuk membangun id
	for i := 0; i < length; i++ {
		bytes[i] = characterSet[bytes[i]%byte(len(characterSet))]
	}

	//gabungkan prefix dengan karakter-karakter id
	id := prefix + "_" + string(bytes)

	Logger.Debugf("berhasil generate id: %s", id)
	return id, nil


}