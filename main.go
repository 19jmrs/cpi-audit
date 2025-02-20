package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2/clientcredentials"
)

type IntegrationPackages struct {
	Id string `json:Id`
	Description string `json:Description`
	Version string `json:Version`
	Vendor string `json:Vendor`
	ShortText string `json:ShortText`
	
}

func main() {
    // Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // Example: Get environment variables
    TokenUrl  := os.Getenv("tokenurl")
    ClientID := os.Getenv("clientid")
    ClientSecret := os.Getenv("clientsecret")
	TenantUrl := os.Getenv("tenanturl")
	

	config := clientcredentials.Config{
		ClientID: ClientID,
		ClientSecret: ClientSecret,
		TokenURL: TokenUrl,
	}

	ctx := context.Background()

	_,err = config.Token(ctx)
	if err != nil{
		log.Fatal("Error getting token: ", err)
	}

	client := config.Client(ctx)

	resp, err := client.Get(TenantUrl+"IntegrationPackages")

	if err != nil {
		log.Fatal("Error getting Packages: ", err)
	}

	fmt.Printf("Packages: ", resp)
	
}
