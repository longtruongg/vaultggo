package main

import (
	"log"
	"main/api"
)

func main() {
	address := "http://127.0.0.1:8200"
	token := "hvs.u5YQ0tFIwPfekiR1WwBBlspc"
	path := "secret/data/mydata"
	vault, err := api.NewVaultToken(address, token)
	if err != nil {
		log.Printf("Vault init %s", err)
	}
	secretData := map[string]interface{}{
		"data": map[string]interface{}{
			"password": "1234",
		},
	}

	requestID, err := vault.WriteDataToSecret(path, secretData)
	if err != nil {
		log.Printf("WriteData %s", err)
	}
	log.Printf("RequestID %s", requestID)
	// read data

	data, err := vault.ReadDataFromSecret(path, "")
	if err != nil {
		log.Printf("ReadData %s", err)
	}
	log.Printf("Data from %s", data)
}
