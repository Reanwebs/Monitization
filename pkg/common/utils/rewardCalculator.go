package utils

import (
	"fmt"
)

func CoinReward(Minutes int32, ConferenceType string) (uint, error) {
	var coins int32

	switch ConferenceType {
	case "public":

		coins = (Minutes / 60) * 12
	case "group":

		coins = (Minutes / 60) * 8
	case "private":

		coins = (Minutes / 60) * 4
	default:
		return 0, fmt.Errorf("unknown ConferenceType: %s", ConferenceType)
	}

	return uint(coins), nil

}

func CoinRewardCreator(input RewardCreator) (float32, error) {
	if input.Reason == "Referal" {
		return 10, nil
	} else if input.Reason == "Views" {
		return (float32(input.Count) * 0.1), nil
	} else if input.Reason == "Paid" {
		return float32(input.Coins) - (float32(input.Coins) / 100), nil
	}
	return 0, nil
}
