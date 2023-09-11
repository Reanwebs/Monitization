package utils

import "fmt"

func CoinReward(Minutes int32, ConferenceType string) (int32, error) {
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

	return coins, nil

}
