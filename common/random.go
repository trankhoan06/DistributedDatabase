package common

import (
	"math/rand"
	"time"
)

var letter = []rune("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890")

func Randomreques(length int) string {
	b := make([]rune, length)
	s1 := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s1)
	for i := range b {
		b[i] = letter[r.Intn(9999999%len(letter))]
	}
	return string(b)
}
func GetSalt(length int) string {
	if length < 1 {
		length = 50
	}
	return Randomreques(length)
}

func GenerateRandomCode() int {
	rand.Seed(time.Now().UnixNano()) // Khởi tạo seed cho số ngẫu nhiên
	min := 100000
	max := 999999
	return rand.Intn(max-min+1) + min
}
