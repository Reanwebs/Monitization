package repository

import (
	"monit/pkg/common/utils"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepoMethods {
	return userRepo{
		db: db,
	}
}

type UserRepoMethods interface {
	CreateWallet(string) error
	GetWallet(string) (utils.Wallet, error)
	UpdateWallet(string, uint) error
	UpdateWalletHistory(utils.UserRewardHistory) error
	GetRewardHistory(string) (*[]utils.UserRewardHistory, error)
	GetRewardHistoryBySort(string, string) (*[]utils.UserRewardHistory, error)
}

func (r userRepo) CreateWallet(userID string) error {
	wallet := utils.Wallet{
		UserID: userID,
		Coins:  10,
	}
	result := r.db.Create(&wallet)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r userRepo) GetWallet(userID string) (utils.Wallet, error) {
	var wallet utils.Wallet
	result := r.db.Where("user_id = ?", userID).First(&wallet)
	if result.Error != nil {
		return utils.Wallet{}, result.Error
	}
	return wallet, nil
}

func (r userRepo) UpdateWallet(userID string, coin uint) error {
	var wallet utils.Wallet
	result := r.db.Where("user_id = ?", userID).First(&wallet)
	if result.Error != nil {
		return result.Error
	}
	wallet.Coins += coin
	result = r.db.Save(&wallet)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r userRepo) UpdateWalletHistory(history utils.UserRewardHistory) error {
	result := r.db.Create(&history)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r userRepo) GetRewardHistory(userID string) (*[]utils.UserRewardHistory, error) {
	var history []utils.UserRewardHistory

	result := r.db.Where("user_id = ?", userID).Find(&history)
	if result.Error != nil {
		return nil, result.Error
	}

	return &history, nil
}

func (r userRepo) GetRewardHistoryBySort(userID string, sort string) (*[]utils.UserRewardHistory, error) {
	var history []utils.UserRewardHistory
	query := r.db.Where("user_id = ? AND reward_reason = ?", userID, sort)

	result := query.Find(&history)
	if result.Error != nil {
		return nil, result.Error
	}
	return &history, nil
}
