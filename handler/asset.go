package handler

import (
	"context"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"html/template"
	"image"
	"log"
	"net/http"
	"strings"

	"github.com/deezy-inc/tao/configs"
	"github.com/deezy-inc/tao/tap"
	"github.com/labstack/echo"
)

func (h *Handler) FetchAsset(c echo.Context) (err error) {
	var (
		ctx = context.Background()
	)

	ctx, _ = configs.LoadConfig(ctx)
	tapClient := tap.NewClient(ctx)

	resp, err := tapClient.GetAsset(c.Param("name"))
	if err != nil {
		log.Println("Error creating address")
		log.Println(err)
		return err
	}

	resp = decodeAssetFields(resp)

	str, _ := json.MarshalIndent(resp, "", "\t")

	str1 := strings.Replace(string(str), "\n", "<br>", -1)
	str1 = strings.Replace(string(str1), "\t", "&emsp;", -1)

	return c.Render(http.StatusOK, "asset.html", map[string]interface{}{
		"asset":     template.HTML(str1),
		"assetName": resp.AssetGenesis.Name,
	})
}

func (h *Handler) FetchAssetProof(c echo.Context) (err error) {
	var (
		ctx = context.Background()
	)

	ctx, _ = configs.LoadConfig(ctx)
	tapClient := tap.NewClient(ctx)

	resp, err := tapClient.GetAssetProof(c.Param("name"))
	if err != nil {
		log.Println("Error exporting proof")
		log.Println(err)
		return err
	}

	resp = decodeProofFields(resp)

	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(resp.Proof.MetaReveal.Data))
	_, format, err := image.Decode(reader)
	if err != nil {
		fmt.Println("Not an image or unable to determine type:", err)
	}

	fmt.Println("Image format is:", format)

	str, _ := json.MarshalIndent(resp, "", "\t")

	str1 := strings.Replace(string(str), "\n", "<br>", -1)
	str1 = strings.Replace(string(str1), "\t", "&emsp;", -1)

	return c.Render(http.StatusOK, "proof.html", map[string]interface{}{
		"proof": template.HTML(str1),
		"image": resp.Proof.MetaReveal.Data,
	})
}

func decodeProofFields(p tap.TapProofResponse) (finalProof tap.TapProofResponse) {
	str, _ := hex.DecodeString(p.Proof.MetaReveal.Data)
	p.Proof.MetaReveal.Data = base64.StdEncoding.EncodeToString(str)
	/*str, _ = base64.StdEncoding.DecodeString(p.Proof.TxMerkleProof)
	p.Proof.TxMerkleProof = hex.EncodeToString(str)
	str, _ = base64.StdEncoding.DecodeString(p.Proof.InclusionProof)
	p.Proof.InclusionProof = hex.EncodeToString(str)
	for i, e := range p.Proof.ExclusionProofs {
		str, _ = base64.StdEncoding.DecodeString(e)
		p.Proof.ExclusionProofs[i] = hex.EncodeToString(str)
	}
	p.Proof.Asset = decodeAssetFields(p.Proof.Asset)*/

	return p
}

func decodeAssetFields(a tap.TapAssetResponse) (finalAsset tap.TapAssetResponse) {
	/*str, _ := base64.StdEncoding.DecodeString(a.AssetGenesis.GenesisPoint)
	a.AssetGenesis.GenesisPoint = hex.EncodeToString(str)
	str, _ = base64.StdEncoding.DecodeString(a.AssetGenesis.Meta)
	a.AssetGenesis.Meta = hex.EncodeToString(str)
	str, _ = base64.StdEncoding.DecodeString(a.AssetGenesis.AssetID)
	a.AssetGenesis.AssetID = hex.EncodeToString(str)
	str, _ = base64.StdEncoding.DecodeString(a.ScriptKey)
	a.ScriptKey = hex.EncodeToString(str)
	str, _ = base64.StdEncoding.DecodeString(a.ChainAnchor.AnchorTx)
	a.ChainAnchor.AnchorTx = hex.EncodeToString(str)
	str, _ = base64.StdEncoding.DecodeString(a.ChainAnchor.InternalKey)
	a.ChainAnchor.InternalKey = hex.EncodeToString(str)*/

	return a
}
