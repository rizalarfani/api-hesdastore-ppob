package helper

import (
	"hesdastore/api-ppob/domain/model"

	"github.com/gin-gonic/gin"
)

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
