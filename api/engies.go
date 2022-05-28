package api

import (
	"encoding/json"
)

func UnmarshalConfigVault(data []byte) (ConfigVault, error) {
	var r ConfigVault
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConfigVault) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ConfigVault struct {
	Type string `json:"type"`
}

func UnmarshalConfigDB(data []byte) (ConfigDB, error) {
	var r ConfigDB
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConfigDB) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ConfigDB struct {
	PluginName    string `json:"plugin_name"`
	AllowedRoles  string `json:"allowed_roles"`
	ConnectionURL string `json:"connection_url"`
	Username      string `json:"username"`
	Password      string `json:"password"`
}
