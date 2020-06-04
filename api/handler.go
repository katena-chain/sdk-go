/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package api

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/transchain/sdk-go/api"
	"github.com/valyala/fasthttp"

	"github.com/katena-chain/sdk-go/entity"
	"github.com/katena-chain/sdk-go/entity/account"
	entityApi "github.com/katena-chain/sdk-go/entity/api"
)

const (
	LastPath         = "/last"
	StatePath        = "/state"
	TxsPath          = "/txs"
	CertificatesPath = "/certificates"
	SecretsPath      = "/secrets"
	CompaniesPath    = "/companies"
	KeysPath         = "/keys"
)

// Handler provides helper methods to send and retrieve txs without directly interacting with the HTTP Client.
type Handler struct {
	apiClient api.Client
}

// Handler constructor.
func NewHandler(apiUrl string) *Handler {
	client := api.NewFastHttpClient(apiUrl)
	client.AddHeader(fasthttp.HeaderContentType, "application/json;charset=UTF-8")
	return &Handler{
		apiClient: client,
	}
}

// RetrieveCertificateTxs fetches the API to return all txs related to a certificate fqid.
func (h *Handler) RetrieveCertificateTxs(fqId string, page int, txPerPage int) (*entityApi.TxResults, error) {
	var txResults entityApi.TxResults
	err := h.GetAndFormat(fmt.Sprintf("%s/%s", CertificatesPath, fqId), api.GetPaginationQueryParams(page, txPerPage), &txResults)
	if err != nil {
		return nil, err
	}
	return &txResults, nil
}

// RetrieveLastCertificateTx fetches the API to return the last tx related to a certificate fqid.
func (h *Handler) RetrieveLastCertificateTx(fqId string) (*entityApi.TxResult, error) {
	var txResult entityApi.TxResult
	err := h.GetAndFormat(fmt.Sprintf("%s/%s%s", CertificatesPath, fqId, LastPath), nil, &txResult)
	if err != nil {
		return nil, err
	}
	return &txResult, nil
}

// RetrieveSecretTxs fetches the API to return all txs related to a secret fqid.
func (h *Handler) RetrieveSecretTxs(fqId string, page int, txPerPage int) (*entityApi.TxResults, error) {
	var txResults entityApi.TxResults
	err := h.GetAndFormat(fmt.Sprintf("%s/%s", SecretsPath, fqId), api.GetPaginationQueryParams(page, txPerPage), &txResults)
	if err != nil {
		return nil, err
	}
	return &txResults, nil
}

// RetrieveLastSecretTxs fetches the API to return the last tx related to a secret fqid.
func (h *Handler) RetrieveLastSecretTx(fqId string) (*entityApi.TxResult, error) {
	var txResult entityApi.TxResult
	err := h.GetAndFormat(fmt.Sprintf("%s/%s%s", SecretsPath, fqId, LastPath), nil, &txResult)
	if err != nil {
		return nil, err
	}
	return &txResult, nil
}

// RetrieveKeyTxs fetches the API to return all txs related to a key fqid.
func (h *Handler) RetrieveKeyTxs(fqId string, page int, txPerPage int) (*entityApi.TxResults, error) {
	var txResults entityApi.TxResults
	err := h.GetAndFormat(fmt.Sprintf("%s/%s", KeysPath, fqId), api.GetPaginationQueryParams(page, txPerPage), &txResults)
	if err != nil {
		return nil, err
	}
	return &txResults, nil
}

// RetrieveLastKeyTxs fetches the API to return the last txs related to a key fqid.
func (h *Handler) RetrieveLastKeyTx(fqId string) (*entityApi.TxResult, error) {
	var txResult entityApi.TxResult
	err := h.GetAndFormat(fmt.Sprintf("%s/%s%s", KeysPath, fqId, LastPath), nil, &txResult)
	if err != nil {
		return nil, err
	}
	return &txResult, nil
}

// RetrieveTx fetches the API to return any tx by its hash.
func (h *Handler) RetrieveTx(hash string) (*entityApi.TxResult, error) {
	var txResult entityApi.TxResult
	err := h.GetAndFormat(fmt.Sprintf("%s/%s", TxsPath, hash), nil, &txResult)
	if err != nil {
		return nil, err
	}
	return &txResult, nil
}

// RetrieveKey fetches the API and returns a key from the state.
func (h *Handler) RetrieveKey(fqId string) (*account.KeyV1, error) {
	var key account.KeyV1
	err := h.GetAndFormat(fmt.Sprintf("%s%s/%s", StatePath, KeysPath, fqId), nil, &key)
	if err != nil {
		return nil, err
	}
	return &key, nil
}

// RetrieveCompanyKeys fetches the API and returns a list of keys for a company from the state.
func (h *Handler) RetrieveCompanyKeys(companyBcId string, page int, txPerPage int) ([]*account.KeyV1, error) {
	var keys []*account.KeyV1
	err := h.GetAndFormat(fmt.Sprintf("%s%s/%s%s", StatePath, CompaniesPath, companyBcId, KeysPath), api.GetPaginationQueryParams(page, txPerPage), &keys)
	if err != nil {
		return nil, err
	}
	return keys, nil
}

// SendTx accepts an encoded tx and sends it to the Api to return its status and its hash.
func (h *Handler) SendRawTx(txBytes []byte) (*entityApi.SendTxResult, error) {
	apiResponse, err := h.SafePost(TxsPath, txBytes)
	if err != nil {
		return nil, err
	}
	var txResult entityApi.SendTxResult
	if err := UnmarshalApiResponse(apiResponse, &txResult); err != nil {
		return nil, err
	}
	return &txResult, nil
}

// SendTx creates a tx from a tx data and the provided tx signer info and chain id, signs it, encodes it and sends it
// to the api.
func (h *Handler) SendTx(txData entity.TxData, txSigner *entity.TxSigner, chainId string) (status *entityApi.SendTxResult, err error) {
	if txSigner == nil || txSigner.FqId == "" || txSigner.PrivateKey == nil || chainId == "" {
		return nil, errors.New("impossible to create txs without a tx signer info or chain id")
	}
	// Sign the tx with the current client time.
	tx := SignTx(txSigner, chainId, entity.GetCurrentTime(), txData)
	txBytes, err := EncodeTx(tx)
	if err != nil {
		return nil, err
	}
	return h.SendRawTx(txBytes)
}

// GetAndFormat fetches the API route and try to unmarshal the response in the provided instance.
func (h *Handler) GetAndFormat(route string, queryParams map[string]string, instance interface{}) error {
	apiResponse, err := h.SafeGet(route, queryParams)
	if err != nil {
		return err
	}
	if err := UnmarshalApiResponse(apiResponse, instance); err != nil {
		return err
	}
	return nil
}

// SafePost calls the api handler post method and recover if it panics.
func (h *Handler) SafePost(route string, body []byte) (_ *api.RawResponse, katenaError error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok := r.(error)
			if !ok {
				katenaError = errors.New(fmt.Sprintf("%v", r))
			} else {
				katenaError = err
			}
		}
	}()
	return h.apiClient.Post(route, body, nil, nil)
}

// SafeGet calls the api handler get method and recover if it panics.
func (h *Handler) SafeGet(route string, queryParams map[string]string) (_ *api.RawResponse, katenaError error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok := r.(error)
			if !ok {
				katenaError = errors.New(fmt.Sprintf("%v", r))
			} else {
				katenaError = err
			}
		}
	}()
	return h.apiClient.Get(route, nil, queryParams)
}

// SignTx creates a tx data state, signs it and returns a tx ready to be encoded and sent.
func SignTx(txSigner *entity.TxSigner, chainId string, nonceTime entity.Time, txData entity.TxData) *entity.Tx {
	txDataState := entity.GetTxDataStateBytes(chainId, nonceTime, txData)
	signature := txSigner.PrivateKey.Sign(txDataState)

	return &entity.Tx{
		NonceTime:  nonceTime,
		Data:       txData,
		SignerFqId: txSigner.FqId,
		Signature:  signature,
	}
}

// EncodeTx defines the way the tx is encoded (here with the json marshaller).
func EncodeTx(tx *entity.Tx) ([]byte, error) {
	return json.Marshal(tx)
}

// UnmarshalApiResponse tries to parse the api response body into the provided interface if the API returns a 200 or a
// 202 HTTP code. If not, it tries to parse it in a PublicError.
func UnmarshalApiResponse(apiResponse *api.RawResponse, dest interface{}) error {
	if apiResponse.StatusCode == fasthttp.StatusOK || apiResponse.StatusCode == fasthttp.StatusAccepted {
		if err := json.Unmarshal(apiResponse.Body, dest); err != nil {
			return err
		}
		return nil
	} else {
		var apiError entityApi.PublicError
		if err := json.Unmarshal(apiResponse.Body, &apiError); err != nil {
			return err
		}
		return apiError
	}
}
