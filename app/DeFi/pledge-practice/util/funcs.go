package util

import (
	"fmt"
	"math"
	mrand "math/rand"
	"regexp"
	"time"
)

func GenerateCode(figures int) (randNum string) {
	startNum := math.Pow(10, float64(figures))
	number := mrand.New(mrand.NewSource(time.Now().UnixNano())).Int31n(int32(startNum))
	return fmt.Sprintf("%06d", number)
}

var (
	EMAIL_REG = "^[0-9A-Za-z]+@[0-9A-Za-z]+\\.[a-z]{2,4}$]"
)

func CheckEmailRegular(email string) bool {
	result, err := regexp.MatchString(EMAIL_REG, email)
	if err != nil {
		return false
	}
	return result
}
