package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type VaultKeys struct {
	Keys []string `json:"keys"`
}

type VaultMeta struct {
	Data VaultKeys `json:"data"`
}

func main() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	tokenFile := "/home/vault/.vault-token"
	token, err := ioutil.ReadFile(tokenFile)
	if err != nil {
		panic(err)
	}
	vaultToken := string(token)

	vaultAddr, exists := os.LookupEnv("VAULT_ADDR")
	if !exists {
		panic("VAULT_ADDR env variable not found. Please control deployment yaml and/or secrets object for the namespace")
	}

	keyPath, exist := os.LookupEnv("VAULT_KEY_PATH")
	if !exist {
		panic("KEY_PATH env variable not found. Please control deployment yaml and/or secrets object for the namespace")
	}

	keysDestinationPath, exist := os.LookupEnv("KEY_DESTINATION_PATH")
	if !exist {
		panic("KEY_DESTINATION_PATH env variable not found. Please control deployment yaml and/or secrets object for the namespace")
	}

	configFile := "/etc/consultemplate/consul-template-config.hcl"
	config, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(err)
	}
	consulConfig := string(config)

	consulConfigTemp := strings.ReplaceAll(consulConfig, "VAULT_ADDR", vaultAddr)
	consulConfigModified := strings.ReplaceAll(consulConfigTemp, "VAULT_TOKEN", vaultToken)

	client := &http.Client{}
	url := vaultAddr + "/v1/" + keyPath
	req, err := http.NewRequest("LIST", url, nil)
	req.Header.Set("X-Vault-Token", vaultToken)
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	var vaultMeta VaultMeta
	json.NewDecoder(res.Body).Decode(&vaultMeta)
	if len(vaultMeta.Data.Keys) > 0 {
		var str strings.Builder
		str.WriteString(consulConfigModified)
		for _, key := range vaultMeta.Data.Keys {
			str.WriteString(fmt.Sprintf("\ntemplate { \n  contents = \"{{- with secret \\\"keys/%s\\\" -}}{{- .Data.content | base64Decode -}}{{- end -}}\" \n  destination = \"%s%s\" \n }",
				key, keysDestinationPath, key))
		}
		configBytes := []byte(str.String())
		err := ioutil.WriteFile("/etc/consultemplatemodified/consul-template-config.hcl", configBytes, 0644)
		if err != nil {
			panic(err)
		}
	} else {
		panic("No keys has been found")
	}

}
