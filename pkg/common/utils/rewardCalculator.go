package utils

import "fmt"

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

func CoinRewardCreator(Type string) (uint, error) {
	if Type == "Referal" {
		return 10, nil
	}
	return 0, nil
}
