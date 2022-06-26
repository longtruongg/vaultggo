package main

import (
	"encoding/json"
	"log"
	"main/api"
	"main/models"
)

func main() {

	vault, err := api.NewVaultToken(api.Address, api.Token)
	if err != nil {
		log.Printf("Vault init %s", err)
	}

	// err = vault.WriteDataToSecret(api.PathSecrets, api.ValKv)
	// if err != nil {
	// 	log.Printf("WriteData %s", err)
	// }

	data, err := vault.ReadDataFromSecret(api.PathKvData, "")
	if err != nil {
		log.Printf("ReadData %s", err)
	}
	jsonData, err := json.Marshal(data)

	if err != nil {
		log.Printf("Data error %s", err)
	}
	kv, err := models.UnmarshalKv(jsonData)
	if err != nil {
		log.Printf("Unmarshal %s", err)
	}
	log.Printf("Kv %v", kv.Key)
}
