package main

import (
	"fmt"
	"os"

	"github.com/Abhi13027/go-arrow/arrow"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()
	userID := os.Getenv("USER_ID")
	password := os.Getenv("PASSWORD")
	totp_key := os.Getenv("TOTP_KEY")
	appID := os.Getenv("APP_ID")
	appSecret := os.Getenv("APP_SECRET")

	fmt.Println(userID, password, totp_key, appID, appSecret)

	client := arrow.NewClient(appID, appSecret)

	err := client.AutoLogin(userID, password, totp_key)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Login successful!")

	// Get user details
	user, err := client.GetUserDetails()

	if err != nil {
		fmt.Println("Error:", err)
		return

	}

	fmt.Printf("User Details: %+v\n", user)

	orders, err := client.GetOrderBook()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Order Book: %+v\n", orders)

	holdings, err := client.GetHoldings()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Holdings: %+v\n", holdings)

}
