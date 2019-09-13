/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package api

import (
    "encoding/json"
    "fmt"

    "github.com/transchain/sdk-go/api"
    "github.com/valyala/fasthttp"

    "github.com/katena-chain/sdk-go/entity"
    entityApi "github.com/katena-chain/sdk-go/entity/api"
)

const routeCertificates = "certificates"
const routeSecrets = "secrets"
const pathCertify = "certify"
const pathHistory = "history"

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

// SendCertificate accepts a tx and sends it to the appropriate certificate API route.
func (h *Handler) SendCertificate(tx *entity.Tx) (*entityApi.TxStatus, error) {
    data, err := json.Marshal(tx)
    if err != nil {
        return nil, err
    }
    return h.SendTx(fmt.Sprintf("%s/%s", routeCertificates, pathCertify), data)
}

// SendSecret accepts a tx and sends it to the appropriate secret API route.
func (h *Handler) SendSecret(tx *entity.Tx) (*entityApi.TxStatus, error) {
    data, err := json.Marshal(tx)
    if err != nil {
        return nil, err
    }
    return h.SendTx(fmt.Sprintf("%s/%s", routeSecrets, pathCertify), data)
}

// RetrieveCertificate fetches the API and returns a tx wrapper or an error.
func (h *Handler) RetrieveCertificate(id string) (*entityApi.TxWrapper, error) {
    apiResponse, err := h.apiClient.Get(fmt.Sprintf("%s/%s", routeCertificates, id), nil, nil)
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
func (h *Handler) RetrieveCertificatesHistory(id string) (*entityApi.TxWrappers, error) {
    apiResponse, err := h.apiClient.Get(fmt.Sprintf("%s/%s/%s", routeCertificates, id, pathHistory), nil, nil)
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
func (h *Handler) RetrieveSecrets(id string) (*entityApi.TxWrappers, error) {
    apiResponse, err := h.apiClient.Get(fmt.Sprintf("%s/%s", routeSecrets, id), nil, nil)
    if err != nil {
        return nil, err
    }
    var txWrappers entityApi.TxWrappers
    if err := unmarshalApiResponse(apiResponse, &txWrappers); err != nil {
        return nil, err
    }
    return &txWrappers, nil
}

// SendTx tries to send a tx to the API and returns a tx status or an error.
func (h *Handler) SendTx(route string, tx []byte) (*entityApi.TxStatus, error) {
    apiResponse, err := h.apiClient.Post(route, tx, nil, nil)
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
        var apiError entityApi.Error
        if err := json.Unmarshal(apiResponse.Body, &apiError); err != nil {
            return err
        }
        return apiError
    }
}
