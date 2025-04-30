package service

import (
	"github.com/ethereum/go-ethereum/accounts"
	"miniapp/model"
	"miniapp/utils"
)

func CreateAccount(pass string) (accounts.Account, string) {
	wallet, secret := utils.CreateWallet(pass)
	utils.GetDb().Create(model.NewAccountB(wallet.URL.Path, pass, wallet.Address.Hex()))
	return wallet, secret
}

func SelectAll() *model.ResultData {
	var accounts []model.AccountB
	utils.GetDb().Model(&model.AccountB{}).Scan(&accounts)
	if len(accounts) > 0 {
		return &model.ResultData{
			Data: accounts,
			Code: 200,
		}
	}

	return nil
}
