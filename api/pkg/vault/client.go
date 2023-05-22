package vault

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type VaultClient struct {
	Client  *http.Client
	ApiKey  string
	BaseUrl string
}

func NewVaultClient(apikey string, baseurl string) *VaultClient {
	return &VaultClient{
		Client:  &http.Client{},
		ApiKey:  apikey,
		BaseUrl: baseurl,
	}
}

func (v *VaultClient) AddToCollection(ledger string, collection string, formid string, form map[string]interface{}) (bool, error) {
	form["__formid"] = formid
	toBytes, err := json.Marshal(form)
	if err != nil {
		log.Println("INVALID JSON FORM: " + err.Error())
		return false, err
	}
	requestBody := bytes.NewBuffer(toBytes)
	req, err := http.NewRequest("PUT", v.BaseUrl+fmt.Sprintf("/ledger/%s/collection/%s/document", ledger, collection), requestBody)
	req.Header.Set("X-API-KEY", v.ApiKey)
	req.Header.Set("Content-type", "application/json")
	if err != nil {
		log.Println("INVALID REQUEST: " + err.Error())
		return false, err
	}
	resp, err := v.Client.Do(req)

	if err != nil {
		log.Printf("An Error Occured %v", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return false, err
	}
	sb := string(body)
	log.Println("Vault response: ", sb)
	return resp.StatusCode == 200, nil
}
