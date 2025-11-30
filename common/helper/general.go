package helper

import (
	"crypto/rand"
	"hesdastore/api-ppob/domain/model"
	"math/big"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	reNonDigit = regexp.MustCompile(`\D+`)
	reLocal08  = regexp.MustCompile(`^08\d+$`)
	rePlus62_8 = regexp.MustCompile(`^(?:\+?62)8(\d+)$`)
	re6288     = regexp.MustCompile(`^628?8(\d+)$`)
	re62       = regexp.MustCompile(`^62(\d+)$`)
)

const refChars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GetPriceProductByRole(role int, product model.Product) int {
	var price int
	if role == 2 {
		price = *product.PriceReseller
	} else {
		price = *product.PriceSeller
	}
	return price
}

func GetStatusProductByRole(role int, product model.Product) string {
	var status string
	if role == 2 {
		status = product.ResellerStatus
	} else {
		status = product.SellerStatus
	}
	return status
}

func GetAuthUser(ctx *gin.Context) (*model.ApiUser, bool) {
	v, ok := ctx.Get("authUser")
	if !ok {
		return nil, false
	}

	user, ok := v.(*model.ApiUser)
	if !ok {
		return nil, false
	}

	return user, true
}

func ToLocal08(msisdn string) string {
	trimmed := strings.TrimSpace(msisdn)

	digits := reNonDigit.ReplaceAllString(trimmed, "")
	if digits == "" {
		return ""
	}

	if reLocal08.MatchString(digits) {
		return digits
	}

	if m := rePlus62_8.FindStringSubmatch(trimmed); m != nil {
		return "08" + m[1]
	}

	if m := re6288.FindStringSubmatch(digits); m != nil {
		return "08" + m[1]
	}

	if m := re62.FindStringSubmatch(digits); m != nil {
		return "0" + m[1]
	}

	return digits
}

func InIntSlice(target int, list []int) bool {
	for _, v := range list {
		if v == target {
			return true
		}
	}
	return false
}

func GenerateRandomString(length int) (string, error) {
	if length <= 0 {
		return "", nil
	}

	result := make([]byte, length)
	max := big.NewInt(int64(len(refChars)))

	for i := 0; i < length; i++ {
		n, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", err
		}
		result[i] = refChars[n.Int64()]
	}

	return string(result), nil
}
