package helpers

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const shortURLLength = 8
const alphabet = "ynAJfoSgdXHB5VasEMtcbPCr1uNZ4LG723ehWkvwYR6KpxjTm8QUFqz9D"

var alphabetLen = uint32(len(alphabet))

func Shorten(longUrl string) string {
	// Генерируем случайную соль длиной 8 символов
	rand.Seed(time.Now().UnixNano())
	salt := make([]byte, 8)
	for i := range salt {
		salt[i] = alphabet[rand.Intn(len(alphabet))]
	}
	saltStr := string(salt)

	// Добавляем соль к длинной ссылке
	longUrlWithSalt := longUrl + saltStr

	// Используем алгоритм MD5 для генерации хеша из длинной ссылки с солью
	hasher := md5.New()
	hasher.Write([]byte(longUrlWithSalt))
	hash := hex.EncodeToString(hasher.Sum(nil))

	// Переводим первые 8 символов хеша в 32-битное целое число
	hash32, err := strconv.ParseUint(hash[:8], 16, 32)
	if err != nil {
		panic(err)
	}

	// Генерируем короткую ссылку на основе целочисленного значения хеша
	var shortURLBuilder strings.Builder
	for i := 0; i < shortURLLength; i++ {
		idx := hash32 % uint64(alphabetLen)
		shortURLBuilder.WriteByte(alphabet[idx])
		hash32 /= uint64(alphabetLen)
	}
	return shortURLBuilder.String()
}
