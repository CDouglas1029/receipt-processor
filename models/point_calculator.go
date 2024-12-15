package models

import (
	"math"
	"strconv"
	"strings"
	"time"
)

func CalculatePoints(receipt Receipt) int {
	points := 0

	//1 point per alphanumeric character in retailer name
	points += len(strings.ReplaceAll(receipt.Retailer, " ", ""))

	//50 points if the taotal is a round dollar amount
	if total, err := strconv.ParseFloat(receipt.Total, 64); err == nil && total == float64(int(total)) {
		points += 50
	}

	//25points if the total is a multiple of 0.25
	if total, err := strconv.ParseFloat(receipt.Total, 64); err == nil && math.Mod(total, 0.25) == 0 {
		points += 25
	}

	//5 points for every two items
	points += (len(receipt.Items) / 2) * 5

	//Points for items with descriptoin length multiple of 3
	for _, item := range receipt.Items {
		descLength := len(strings.TrimSpace(item.ShortDescription))
		if descLength%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			points += int(math.Ceil(price * 0.2))
		}
	}

	//6 points if the purchase day is odd
	purchaseDate, _ := time.Parse("2006-01-02", receipt.PurchaseData)
	if purchaseDate.Day()%2 != 0 {
		points += 6
	}

	//10 points if the purchase time is between 2:00pm and 4:00pm
	purchaseTime, _ := time.Parse("15:04", receipt.PurchaseTime)
	if purchaseTime.Hour() >= 14 && purchaseTime.Hour() < 16 {
		points += 10
	}

	return points
}
