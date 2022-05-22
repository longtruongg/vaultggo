package api

import (
	vault "github.com/hashicorp/vault/api"
)

type Vault struct {
	token   string
	address string
}
type SpecialError struct {
	mess string
}

func (s SpecialError) Error() string {
	return s.mess
}
func NewVaultToken(address, token string) (*Vault, error) {
	v := &Vault{token: token, address: address}
	_, err := v.setup()
	if err != nil {
		return nil, &SpecialError{mess: err.Error()}
	}
	return v, nil
}
func (recierver Vault) setup() (*vault.Client, error) {
	vaultConfig := vault.DefaultConfig()
	vaultConfig.Address = recierver.address
	client, err := vault.NewClient(vaultConfig)
	if err != nil {
		return nil, &SpecialError{mess: err.Error()}
	}
	client.SetToken(recierver.token)
	return client, nil
}
func (recierver Vault) WriteDataToSecret(pathSecret string, data map[string]interface{}) (string, error) {
	client, err := recierver.setup()
	if err != nil {
		return "", &SpecialError{mess: err.Error()}
	}
	writeData, err := client.Logical().Write(pathSecret, data)
	if err != nil {
		return "", &SpecialError{mess: err.Error()}
	}
	return writeData.RequestID, nil
}
func (recierver Vault) ReadDataFromSecret(pathToSecret, field string) (map[string]interface{}, error) {
	client, err := recierver.setup()
	if err != nil {
		return nil, &SpecialError{mess: err.Error()}
	}
	secret, err := client.Logical().Read(pathToSecret)
	if err != nil {
		return nil, &SpecialError{mess: err.Error()}
	}
	data, ok := secret.Data["data"].(map[string]interface{})
	if !ok {
		return nil, &SpecialError{mess: err.Error()}
	}
	return data, nil
}
func (recierver Vault) CreatePolicy() (bool, error) {
	_, err := recierver.setup()
	if err != nil {
		return false, &SpecialError{mess: err.Error()}
	}
	return false, nil
}
