package tap

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/deezy-inc/tao/configs"
)

type TapClient struct {
	Client              *http.Client
	Host                string
	Macaroon            string
	Context             context.Context
	CachedAssetResponse TapAssetsResponse
}

// func NewTapClient
func NewClient(ctx context.Context) (client TapClient) {
	config := configs.GetConfig(ctx)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpClient := &http.Client{
		Transport: tr,
	}
	client = TapClient{
		Client:   httpClient,
		Host:     config.TapHost,
		Macaroon: loadMacaroon(ctx),
	}

	return client
}

func loadMacaroon(ctx context.Context) (macaroon string) {
	macaroonBytes, err := ioutil.ReadFile(configs.GetConfig(ctx).TapMacaroonLocation)
	if err != nil {
		log.Println("couldnt find or open macaroon")
		log.Println(err)
		return configs.GetConfig(ctx).TapMacaroon
	}

	macaroon = hex.EncodeToString(macaroonBytes)

	log.Println(macaroon)
	return macaroon
}

func (client *TapClient) sendGetRequest(endpoint string) (*http.Response, error) {
	log.Println(client.Host + endpoint)
	req, err := http.NewRequest("GET", client.Host+endpoint, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Grpc-Metadata-macaroon", client.Macaroon)
	resp, err := client.Client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (client *TapClient) sendPostRequestJSON(endpoint string, payload interface{}) (*http.Response, error) {
	jsonStr, err := json.Marshal(payload)
	req, err := http.NewRequest("POST", client.Host+endpoint, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Grpc-Metadata-macaroon", client.Macaroon)
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return resp, nil
}

func (client *TapClient) sendPostRequest(endpoint string, payload string) (*http.Response, error) {
	jsonStr := []byte(payload)

	req, err := http.NewRequest("POST", client.Host+endpoint, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Grpc-Metadata-macaroon", client.Macaroon)
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
