package utils

import "time"

type UserRewardHistory struct {
	UserID          string
	RewardType      string
	TransactionType string
	CoinCount       uint
	Time            time.Time
}
