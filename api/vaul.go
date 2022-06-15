package api

import (
	vault "github.com/hashicorp/vault/api"
	"log"
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

func (recierver Vault) CreateSecretsEngie(path string, data map[string]interface{}) (bool, error) {
	client, err := recierver.setup()
	if err != nil {
		return false, &SpecialError{mess: err.Error()}
	}
	_, err = client.Logical().Write(path, data)
	if err != nil {
		return false, &SpecialError{mess: err.Error()}
	}
	return true, nil
}
func (recierver Vault) WriteDatabaseConfig(path string, data map[string]interface{}) (bool, error) {
	client, err := recierver.setup()
	if err != nil {
		return false, &SpecialError{mess: err.Error()}
	}
	_, err = client.Logical().Write(path, data)
	if err != nil {
		return false, &SpecialError{mess: err.Error()}
	}
	return true, nil
}
func (recierver Vault) ListConfig(path string) (map[string]interface{}, error) {
	client, err := recierver.setup()
	if err != nil {
		return nil, &SpecialError{mess: err.Error()}
	}
	list, err := client.Logical().List(path)
	if err != nil {
		return nil, err
	}

	return list.Data, nil
}
func (recierver Vault) RotateRoot(path string) (bool, error) {
	client, err := recierver.setup()
	if err != nil {
		return false, &SpecialError{mess: err.Error()}
	}
	_, err = client.Logical().Write(path, nil)
	if err != nil {
		return false, &SpecialError{mess: err.Error()}
	}
	return true, nil
}

func (recierver Vault) CreateStaticRole(path string, data map[string]interface{}) (bool, error) {
	client, err := recierver.setup()
	if err != nil {
		return false, &SpecialError{mess: err.Error()}
	}
	_, err = client.Logical().Write(path, data)
	if err != nil {
		return false, &SpecialError{mess: err.Error()}
	}
	return true, nil
}
func (recierver Vault) ReadStaticRole(path string) (interface{}, error) {
	client, err := recierver.setup()
	if err != nil {
		return nil, &SpecialError{mess: err.Error()}
	}
	res, err := client.Logical().Read(path)
	if err != nil {
		return nil, &SpecialError{mess: err.Error()}
	}
	return res, nil
}
func (recierver Vault) CreatePolicy(path string, data map[string]interface{}) (bool, error) {
	client, err := recierver.setup()
	if err != nil {
		return false, &SpecialError{mess: err.Error()}
	}
	res, err := client.Logical().Write(path, data)
	if err != nil {
		return false, &SpecialError{mess: err.Error()}
	}
	log.Printf("Data from request %s", res)
	return true, nil
}
func (recierver Vault) CreateRotateToken(path string, data map[string]interface{}) (string, error) {
	client, err := recierver.setup()
	if err != nil {
		return "", &SpecialError{mess: err.Error()}
	}
	result, err := client.Logical().Write(path, data)
	if err != nil {
		return "", &SpecialError{mess: err.Error()}
	}
	return result.Auth.ClientToken, nil
}
func (recierver Vault) EndpointRotateRoleByAdmin(path string) (bool, error) {
	client, err := recierver.setup()
	if err != nil {
		return false, &SpecialError{mess: err.Error()}
	}
	_, err = client.Logical().Write(path, nil)
	if err != nil {
		return false, err
	}
	return true, nil
}
func (recierver Vault) RotateTokenByAdmin(path, token string) (map[string]interface{}, error) {
	client, err := recierver.setup()
	if err != nil {
		return nil, &SpecialError{mess: err.Error()}
	}
	client.SetToken(token)
	res, err := client.Logical().Read(path)
	if err != nil {
		return nil, &SpecialError{mess: err.Error()}
	}
	return res.Data, nil
}
