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
	"github.com/transchain/sdk-go/crypto/ed25519"
	"github.com/valyala/fasthttp"

	"github.com/katena-chain/sdk-go/entity"
	"github.com/katena-chain/sdk-go/entity/account"
	entityApi "github.com/katena-chain/sdk-go/entity/api"
)

const (
	CertificatesPath = "/certificates"
	LastPath         = "/last"
	SecretsPath      = "/secrets"
	TxsPath          = "/txs"
	CompaniesPath    = "/companies"
	KeysPath         = "/keys"
)

// Handler provides helper methods to send and retrieve txs without directly interacting with the HTTP Client.
type Handler struct {
	apiClient api.Client
}

// Handler constructor.
func NewHandler(apiUrl string) *Handler {
	return &Handler{
		apiClient: api.NewFastHttpClient(apiUrl),
	}
}

// SendTx accepts an encoded tx and sends it to the Api to return a tx status or an error.
func (h *Handler) SendTx(txBytes []byte) (*entityApi.TxStatus, error) {
	apiResponse, err := h.SafePost(TxsPath, txBytes)
	if err != nil {
		return nil, err
	}
	var txStatus entityApi.TxStatus
	if err := unmarshalApiResponse(apiResponse, &txStatus); err != nil {
		return nil, err
	}
	return &txStatus, nil
}

// RetrieveCertificates fetches the API and returns a tx wrapper list or an error.
func (h *Handler) RetrieveCertificates(id string, page int, txPerPage int) (*entityApi.TxWrappers, error) {
	queryParams := api.GetPaginationQueryParams(page, txPerPage)
	apiResponse, err := h.SafeGet(fmt.Sprintf("%s/%s", CertificatesPath, id), queryParams)
	if err != nil {
		return nil, err
	}

	var txWrappers entityApi.TxWrappers
	if err := unmarshalApiResponse(apiResponse, &txWrappers); err != nil {
		return nil, err
	}
	return &txWrappers, nil
}

// RetrieveLastCertificate fetches the API and returns a tx wrapper or an error.
func (h *Handler) RetrieveLastCertificate(id string) (*entityApi.TxWrapper, error) {
	apiResponse, err := h.SafeGet(fmt.Sprintf("%s/%s%s", CertificatesPath, id, LastPath), nil)
	if err != nil {
		return nil, err
	}

	var txWrapper entityApi.TxWrapper
	if err := unmarshalApiResponse(apiResponse, &txWrapper); err != nil {
		return nil, err
	}
	return &txWrapper, nil
}

// RetrieveSecrets fetches the API and returns a tx wrapper list or an error.
func (h *Handler) RetrieveSecrets(id string, page int, txPerPage int) (*entityApi.TxWrappers, error) {
	queryParams := api.GetPaginationQueryParams(page, txPerPage)
	apiResponse, err := h.SafeGet(fmt.Sprintf("%s/%s", SecretsPath, id), queryParams)
	if err != nil {
		return nil, err
	}
	var txWrappers entityApi.TxWrappers
	if err := unmarshalApiResponse(apiResponse, &txWrappers); err != nil {
		return nil, err
	}
	return &txWrappers, nil
}

// RetrieveLastTx fetches the API and returns a tx wrapper or an error.
func (h *Handler) RetrieveLastTx(txCategory string, id string) (*entityApi.TxWrapper, error) {
	txWrappers, err := h.RetrieveTxs(txCategory, id, 1, api.DefaultPerPageParam)
	if err != nil {
		return nil, err
	}

	// Calculate the last page
	if txWrappers.Total > api.DefaultPerPageParam {
		pageToFetch := api.GetNumPage(api.DefaultPerPageParam, int(txWrappers.Total))
		txWrappers, err = h.RetrieveTxs(txCategory, id, pageToFetch, api.DefaultPerPageParam)
		if err != nil {
			return nil, err
		}
	}

	lastTx := txWrappers.Txs[len(txWrappers.Txs)-1]
	return lastTx, nil
}

// RetrieveTxs fetches the API and returns a tx wrapper list or an error.
func (h *Handler) RetrieveTxs(txCategory string, id string, page int, txPerPage int) (*entityApi.TxWrappers, error) {
	queryParams := api.GetPaginationQueryParams(page, txPerPage)
	apiResponse, err := h.SafeGet(fmt.Sprintf("%s/%s/%s", TxsPath, txCategory, id), queryParams)
	if err != nil {
		return nil, err
	}

	var txWrappers entityApi.TxWrappers
	if err := unmarshalApiResponse(apiResponse, &txWrappers); err != nil {
		return nil, err
	}

	return &txWrappers, nil
}

// RetrieveCompanyKeys fetches the API and returns the list of keyV1 for a company.
func (h *Handler) RetrieveCompanyKeys(companyBcid string, page int, txPerPage int) ([]*account.KeyV1, error) {
	queryParams := api.GetPaginationQueryParams(page, txPerPage)
	apiResponse, err := h.SafeGet(fmt.Sprintf("%s/%s%s", CompaniesPath, companyBcid, KeysPath), queryParams)
	if err != nil {
		return nil, err
	}

	var keys []*account.KeyV1
	if err := unmarshalApiResponse(apiResponse, &keys); err != nil {
		return nil, err
	}
	return keys, nil
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

// SignTx creates a tx data state, signs it and returns a tx ready to encode and send.
func (h *Handler) SignTx(privateKey *ed25519.PrivateKey, chainID string, nonceTime entity.Time, txData entity.TxData) *entity.Tx {
	txDataState := entity.GetTxDataStateBytes(chainID, nonceTime, txData)
	signature := privateKey.Sign(txDataState)

	return &entity.Tx{
		NonceTime: nonceTime,
		Data:      txData,
		Signer:    privateKey.GetPublicKey(),
		Signature: signature,
	}
}

// EncodeTx defines the way the tx is encoded (here with the json marshaller).
func (h *Handler) EncodeTx(tx *entity.Tx) ([]byte, error) {
	return json.Marshal(tx)
}

// unmarshalApiResponse tries to parse the api response body into the provided interface if the API returns a 200 or a
// 202 HTTP code. If not, it tries to parse it in a PublicError.
func unmarshalApiResponse(apiResponse *api.RawResponse, dest interface{}) error {
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
