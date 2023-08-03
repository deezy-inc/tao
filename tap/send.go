package tap

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type TapSendRequest struct {
	Address []string `json:"tap_addrs" form:"tap_addrs" bson:"tap_addrs"`
	// ScriptKey   string `json:"script_key" form:"script_key" bson:"script_key"`
	// InternalKey string `json:"internal_key" form:"internal_key" bson:"internal_key"`
}

type TapSendResponse struct {
	AnchorTxHash string `json:"anchor_tx_hash"`
}

func (client *TapClient) CreateSend(address []string) (hash TapSendResponse, err error) {
	/*if groupKey != "" {
		groupKeyHex, err := hex.DecodeString(groupKey)
		if err != nil {
			return address, err
		}
		groupKey = base64.URLEncoding.EncodeToString(groupKeyHex)
	}*/

	resp, err := client.sendPostRequestJSON("v1/taproot-assets/send", &TapSendRequest{
		Address: address,
	})
	if err != nil {
		log.Println(err)
		return hash, err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return hash, err
	}

	if err := json.Unmarshal(bodyBytes, &hash); err != nil {
		log.Println(err)
		return hash, err
	}

	/*str, _ := base64.StdEncoding.DecodeString(address.AssetId)
	address.AssetId = hex.EncodeToString(str)
	str, _ = base64.StdEncoding.DecodeString(address.GroupKey)
	address.GroupKey = hex.EncodeToString(str)
	str, _ = base64.StdEncoding.DecodeString(address.ScriptKey)
	address.ScriptKey = hex.EncodeToString(str)
	str, _ = base64.StdEncoding.DecodeString(address.InternalKey)
	address.InternalKey = hex.EncodeToString(str)
	str, _ = base64.StdEncoding.DecodeString(address.TaprootOutputKey)
	address.TaprootOutputKey = hex.EncodeToString(str)*/

	return hash, err
}
