package main

import (
	"log"
	"main/api"
)

func main() {

	vault, err := api.NewVaultToken(api.Address, api.Token)
	if err != nil {
		log.Printf("Vault init %s", err)
	}
	secretData := map[string]interface{}{
		"data": map[string]interface{}{
			"password": "1234",
		},
	}

	requestID, err := vault.WriteDataToSecret(api.PathSecrets, secretData)
	if err != nil {
		log.Printf("WriteData %s", err)
	}
	log.Printf("RequestID %s", requestID)
	// read data

	data, err := vault.ReadDataFromSecret(api.PathSecrets, "")
	if err != nil {
		log.Printf("ReadData %s", err)
	}
	log.Printf("Data from Secrets %s", data)
	result, err := vault.ReadStaticRole(api.PathStatic)
	log.Printf("Data Role %s", result)
	log.Println("--------------------------------------------------")
	tok, err := vault.CreateRotateToken(api.PathCreateToken, api.ValToken)
	if err != nil {
		log.Printf("Error %s", err)
	}
	log.Printf("Token is :: %s", tok)
	result, err = vault.EndpointRotateRoleByAdmin(api.PathReCreateToken)
	if err != nil {
		log.Printf("Error from Endpoint %s", err)
	}
	log.Printf("Isok? %s", result)
	log.Println("--------------------------------------------------")
	reTok, err := vault.RotateTokenByAdmin(api.PathRotateTokenByRoot, api.Token)
	if err != nil {
		log.Printf("RotateToken %s", err)
	}
	log.Printf("RecreateToken :: %s", reTok)
}
