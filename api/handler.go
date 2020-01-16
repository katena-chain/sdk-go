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
    "strconv"

    "github.com/transchain/sdk-go/api"
    "github.com/transchain/sdk-go/crypto/ed25519"
    "github.com/valyala/fasthttp"

    "github.com/katena-chain/sdk-go/entity"
    entityApi "github.com/katena-chain/sdk-go/entity/api"
)

const (
    routeCertificates = "certificates"
    routeSecrets      = "secrets"
    pathHistory       = "history"
    routeTxs          = "txs"

    PageParam           = "page"
    PerPageParam        = "per_page"
    DefaultPerPageParam = 10
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

// RetrieveCertificate fetches the API and returns a tx wrapper or an error.
func (h *Handler) RetrieveCertificate(id string) (*entityApi.TxWrapper, error) {
    apiResponse, err := h.query(fmt.Sprintf("%s/%s", routeCertificates, id), map[string]string{})
    if err != nil {
        return nil, err
    }

    var txWrapper entityApi.TxWrapper
    if err := unmarshalApiResponse(apiResponse, &txWrapper); err != nil {
        return nil, err
    }
    return &txWrapper, nil
}

// RetrieveCertificatesHistory fetches the API and returns tx wrappers or an error.
func (h *Handler) RetrieveCertificatesHistory(id string, page int, txPerPage int) (*entityApi.TxWrappers, error) {
    queryParams := map[string]string{
        PageParam:    strconv.Itoa(page),
        PerPageParam: strconv.Itoa(txPerPage),
    }
    apiResponse, err := h.query(fmt.Sprintf("%s/%s/%s", routeCertificates, id, pathHistory), queryParams)
    if err != nil {
        return nil, err
    }

    var txWrappers entityApi.TxWrappers
    if err := unmarshalApiResponse(apiResponse, &txWrappers); err != nil {
        return nil, err
    }
    return &txWrappers, nil
}

// RetrieveSecrets fetches the API and returns a tx wrapper list or an error.
func (h *Handler) RetrieveSecrets(id string, page int, txPerPage int) (*entityApi.TxWrappers, error) {
    queryParams := map[string]string{
        PageParam:    strconv.Itoa(page),
        PerPageParam: strconv.Itoa(txPerPage),
    }
    apiResponse, err := h.query(fmt.Sprintf("%s/%s", routeSecrets, id), queryParams)
    if err != nil {
        return nil, err
    }
    var txWrappers entityApi.TxWrappers
    if err := unmarshalApiResponse(apiResponse, &txWrappers); err != nil {
        return nil, err
    }
    return &txWrappers, nil
}

func (h *Handler) GetTx(queryKey string, id string) (*entityApi.TxWrapper, error) {
    txWrappers, err := h.queryTxs(queryKey, id, 1, DefaultPerPageParam)
    if err != nil {
        return nil, err
    }

    // Calculate the last page
    if txWrappers.Total > DefaultPerPageParam {
        pageToFetch := (txWrappers.Total / DefaultPerPageParam) + 1
        txWrappers, err = h.queryTxs(queryKey, id, int(pageToFetch), DefaultPerPageParam)
        if err != nil {
            return nil, err
        }
    }

    // Decode ResultTxSearch
    lastTx := txWrappers.Txs[len(txWrappers.Txs)-1]
    return lastTx, nil
}

func (h *Handler) GetTxs(queryKey string, id string, page int, txPerPage int) (*entityApi.TxWrappers, error) {
    return h.queryTxs(queryKey, id, page, txPerPage)
}

func (h *Handler) queryTxs(queryKey string, queryValue string, page int, txPerPage int) (_ *entityApi.TxWrappers, katenaError error) {
    queryParams := map[string]string{
        PageParam:    strconv.Itoa(page),
        PerPageParam: strconv.Itoa(txPerPage),
    }
    apiResponse, err := h.query(fmt.Sprintf("%s/%s/%s", routeTxs, queryKey, queryValue), queryParams)
    if err != nil {
        return nil, err
    }

    var txWrappers entityApi.TxWrappers
    if err := unmarshalApiResponse(apiResponse, &txWrappers); err != nil {
        return nil, err
    }

    return &txWrappers, nil
}

func (h *Handler) query(route string, queryParams map[string]string) (_ *api.RawResponse, katenaError error) {
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

func (h *Handler) SignTx(privateKey *ed25519.PrivateKey, chainId string, nonceTime entity.Time, txData entity.TxData) *entity.Tx {
    txDataState := entity.GetTxDataStateBytes(chainId, nonceTime, txData)
    signature := privateKey.Sign(txDataState)

    return &entity.Tx{
        NonceTime: nonceTime,
        Data:      txData,
        Signer:    privateKey.GetPublicKey(),
        Signature: signature,
    }
}

func (h *Handler) EncodeTx(tx *entity.Tx) ([]byte, error) {
    return json.Marshal(tx)
}

func (h *Handler) SendTx(txBytes []byte) (_ *entityApi.TxStatus, katenaError error) {
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

    apiResponse, err := h.apiClient.Post(routeTxs, txBytes, nil, nil)
    if err != nil {
        return nil, err
    }
    var txStatus entityApi.TxStatus
    if err := unmarshalApiResponse(apiResponse, &txStatus); err != nil {
        return nil, err
    }
    return &txStatus, nil
}

// unmarshalApiResponse tries to parse the api response body if the API returns a 2xx HTTP code.
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
