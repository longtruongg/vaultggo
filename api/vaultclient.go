package api

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net/http"
)

func InitVaultClient() (*http.Client, error) {
	cacert, err := ioutil.ReadFile("/home/ningen/workspace/vaultggo/api/root.crt")
	if err != nil {
		return nil, err
	}
	cacertPool := x509.NewCertPool()
	cacertPool.AppendCertsFromPEM(cacert)
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: cacertPool,
			},
		},
	}
	return client, nil

}
