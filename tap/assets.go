package tap

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"
)

type TapAssetsResponse struct {
	TapAssets []TapAssetResponse `json:"assets"`
}

type TapAssetResponse struct {
	Version      int `json:"version"`
	AssetGenesis struct {
		GenesisPoint string `json:"genesis_point"`
		Name         string `json:"name"`
		Meta         string `json:"meta_hash"`
		AssetID      string `json:"asset_id"`
		OutputIndex  int    `json:"output_index"`
		Version      int    `json:"version"`
	} `json:"asset_genesis"`
	AssetType        string `json:"asset_type"`
	Amount           string `json:"amount"`
	LockTime         int    `json:"lock_time"`
	RelativeLockTime int    `json:"relative_lock_time"`
	ScriptVersion    int    `json:"script_version"`
	ScriptKey        string `json:"script_key"`
	ScriptKeyIsLocal bool   `json:"script_key_is_local"`
	AssetGroup       struct {
		RawGroupKey     string `json:"raw_group_key"`
		TweakedGroupKey string `json:"tweaked_group_key"`
		AssetIDSig      string `json:"asset_id_sig"`
	} `json:"asset_group"`
	ChainAnchor struct {
		AnchorTx         string `json:"anchor_tx"`
		AnchorTxid       string `json:"anchor_txid"`
		AnchorBlockHash  string `json:"anchor_block_hash"`
		AnchorOutpoint   string `json:"anchor_outpoint"`
		InternalKey      string `json:"internal_key"`
		MerkleRoot       string `json:"merkle_root"`
		TapscriptSibling string `json:"tapscript_sibling"`
	} `json:"chain_anchor"`
	PrevWitnesses []struct {
		PrevId struct {
			AnchorPoint string `json:"anchor_point"`
			AssetId     string `json:"asset_id"`
			ScriptKey   string `json:"script_key"`
			Amount      string `json:"amount"`
		} `json:"prev_id"`
		TxWitness       []string `json:"tx_witness"`
		SplitCommitment struct {
			RootAsset struct {
				Version      int `json:"version"`
				AssetGenesis struct {
					GenesisPoint string `json:"genesis_point"`
					Name         string `json:"name"`
					MetaHash     string `json:"meta_hash"`
					AssetID      string `json:"asset_id"`
					OutputIndex  int    `json:"output_index"`
					Version      int    `json:"version"`
				} `json:"asset_genesis"`
				AssetType        string      `json:"asset_type"`
				Amount           string      `json:"amount"`
				LockTime         int         `json:"lock_time"`
				RelativeLockTime int         `json:"relative_lock_time"`
				ScriptVersion    int         `json:"script_version"`
				ScriptKey        string      `json:"script_key"`
				ScriptKeyIsLocal bool        `json:"script_key_is_local"`
				AssetGroup       interface{} `json:"asset_group"`
				ChainAnchor      interface{} `json:"chain_anchor"`
				PrevWitnesses    []struct {
					PrevID struct {
						AnchorPoint string `json:"anchor_point"`
						AssetID     string `json:"asset_id"`
						ScriptKey   string `json:"script_key"`
						Amount      string `json:"amount"`
					} `json:"prev_id"`
					TxWitness       []string    `json:"tx_witness"`
					SplitCommitment interface{} `json:"split_commitment"`
				} `json:"prev_witnesses"`
				IsSpent bool `json:"is_spent"`
			} `json:"root_asset"`
		} `json:"split_commitment"`
	} `json:"prev_witnesses"`
	IsSpent bool `json:"is_spent"`
}

func (client *TapClient) GetAsset(assetName string) (assetResponse TapAssetResponse, err error) {
	var assets TapAssetsResponse
	if len(client.CachedAssetResponse.TapAssets) == 0 {
		assets, err = client.ListAssets()
		if err != nil {
			return assetResponse, nil
		}
	} else {
		assets = client.CachedAssetResponse
	}

	for _, asset := range assets.TapAssets {
		if asset.AssetGenesis.Name == assetName {
			assetResponse = asset
			break
		}
	}

	return assetResponse, err
}

type TapExportProofRequest struct {
	AssetId   string `json:"asset_id"`
	ScriptKey string `json:"script_key"`
}

type TapDecodeProofRequest struct {
	RawProof          string `json:"raw_proof"`
	WithPrevWitnesses bool   `json:"with_prev_witnesses"`
	WithMetaReveal    bool   `json:"with_meta_reveal"`
}

type TapProofResponse struct {
	Proof TapProof `json:"decoded_proof"`
}

type TapProof struct {
	ProofAtDepth   int              `json:"proof_at_depth"`
	NumberOfProofs int              `json:"number_of_proofs"`
	Asset          TapAssetResponse `json:"asset"`
	MetaReveal     struct {
		Data     string `json:"data"`
		Type     string `json:"type"`
		MetaHash string `json:"meta_hash"`
	} `json:"meta_reveal"`
	TxMerkleProof   string   `json:"tx_merkle_proof"`
	InclusionProof  string   `json:"inclusion_proof"`
	ExclusionProofs []string `json:"exclusion_proofs"`
}

func (client *TapClient) GetAssetProof(assetName string) (proofResponse TapProofResponse, err error) {
	var assets TapAssetsResponse
	var asset TapAssetResponse
	if len(client.CachedAssetResponse.TapAssets) == 0 {
		assets, err = client.ListAssets()
		if err != nil {
			return proofResponse, nil
		}
	} else {
		assets = client.CachedAssetResponse
	}

	for _, a := range assets.TapAssets {
		if a.AssetGenesis.Name == assetName {
			asset = a
			break
		}
	}

	log.Println("hello2")

	var encodedProof TapDecodeProofRequest
	resp, err := client.sendPostRequestJSON("v1/taproot-assets/proofs/export", &TapExportProofRequest{
		asset.AssetGenesis.AssetID,
		asset.ScriptKey,
	})
	if err != nil {
		log.Println(err)
		return proofResponse, err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return proofResponse, err
	}

	if err := json.Unmarshal(bodyBytes, &encodedProof); err != nil {
		log.Println(err)
		return proofResponse, err
	}

	resp, err = client.sendPostRequestJSON("v1/taproot-assets/proofs/decode", &TapDecodeProofRequest{
		encodedProof.RawProof, true, true,
	})
	if err != nil {
		log.Println(err)
		return proofResponse, err
	}

	bodyBytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return proofResponse, err
	}

	if strings.Contains(string(bodyBytes), "unable to fetch asset meta") {
		// try fetching proof info without meta data
		resp, err = client.sendPostRequestJSON("v1/taproot-assets/proofs/decode", &TapDecodeProofRequest{
			encodedProof.RawProof, true, false,
		})
		if err != nil {
			// if error contains "unable to fetch asset meta" then try again without it
			log.Println(err)
			return proofResponse, err
		}

		bodyBytes, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			return proofResponse, err
		}
	}

	if err := json.Unmarshal(bodyBytes, &proofResponse); err != nil {
		log.Println(err)
		return proofResponse, err
	}

	log.Println(proofResponse)

	return proofResponse, err
}

func (client *TapClient) ListAssets() (assets TapAssetsResponse, err error) {
	resp, err := client.sendGetRequest("v1/taproot-assets/assets")
	if err != nil {
		log.Println(err)
		return assets, err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("death")
		log.Println(err)
		return assets, err
	}

	log.Println(string(bodyBytes))

	if err := json.Unmarshal(bodyBytes, &assets); err != nil {
		log.Println("death2")
		log.Println(err)
		return assets, err
	}

	client.CachedAssetResponse = assets

	return assets, err
}
