package utils

import (
	"time"

	"gorm.io/gorm"
)

type Models struct {
	Wallet
	UserRewardHistory
}

type Wallet struct {
	gorm.Model
	UserID string
	Coins  uint
}
type UserRewardHistory struct {
	UserID          string
	RewardReason    string
	TransactionType string
	Referal         string
	CoinCount       uint
	Time            time.Time
}
