/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package client

import (
    "errors"
    "time"

    "github.com/google/uuid"
    "github.com/transchain/sdk-go/crypto/ed25519"
    "github.com/transchain/sdk-go/crypto/nacl"

    "github.com/katena-chain/sdk-go/api"
    "github.com/katena-chain/sdk-go/entity"
    "github.com/katena-chain/sdk-go/entity/account"
    entityApi "github.com/katena-chain/sdk-go/entity/api"
    "github.com/katena-chain/sdk-go/entity/certify"
    "github.com/katena-chain/sdk-go/entity/common"
)

// Transactor provides helper methods to hide the complexity of Tx creation, signature and API dialog.
type Transactor struct {
    apiHandler     *api.Handler
    chainId        string
    txSigner       *ed25519.PrivateKey
    companyChainId string
}

// Transactor constructor.
func NewTransactor(apiUrl string, chainId string, companyChainId string, txSigner *ed25519.PrivateKey) *Transactor {
    return &Transactor{
        apiHandler:     api.NewHandler(apiUrl),
        chainId:        chainId,
        txSigner:       txSigner,
        companyChainId: companyChainId,
    }
}

// SendCertificateRawV1 creates a CertificateRaw (V1), wraps it in a tx and sends it to the API.
func (t Transactor) SendCertificateRawV1(uuid string, value []byte) (*entityApi.TxStatus, error) {
    certificate := certify.CertificateRawV1{
        Id:    common.FormatTxid(t.companyChainId, uuid),
        Value: value,
    }

    return t.SendTx(certificate)
}

// SendCertificateEd25519V1 creates a CertificateEd25519 (V1), wraps it in a tx and sends it to the API.
func (t Transactor) SendCertificateEd25519V1(
    uuid string,
    signer ed25519.PublicKey,
    signature ed25519.Signature,
) (*entityApi.TxStatus, error) {
    certificate := certify.CertificateEd25519V1{
        Id:        common.FormatTxid(t.companyChainId, uuid),
        Signer:    signer,
        Signature: signature,
    }

    return t.SendTx(certificate)
}

// SendKeyCreateV1 creates a KeyCreate (V1), wraps it in a tx and sends it to the API.
func (t Transactor) SendKeyCreateV1(uuid string, publicKey ed25519.PublicKey, role string) (*entityApi.TxStatus, error) {
    keyCreate := account.NewKeyCreateV1(common.FormatTxid(t.companyChainId, uuid), publicKey, role)
    return t.SendTx(keyCreate)
}

// SendKeyRevokeV1 creates a KeyRevoke (V1), wraps it in a tx and sends it to the API.
func (t Transactor) SendKeyRevokeV1(uuid string, publicKey ed25519.PublicKey) (*entityApi.TxStatus, error) {
    keyRevoke := account.NewKeyRevokeV1(common.FormatTxid(t.companyChainId, uuid), publicKey)
    return t.SendTx(keyRevoke)
}

// RetrieveCertificate fetches the API to find the corresponding tx and returns a tx wrapper or an error.
func (t Transactor) RetrieveCertificate(companyChainId string, uuid string) (*entityApi.TxWrapper, error) {
    return t.apiHandler.RetrieveCertificate(common.FormatTxid(companyChainId, uuid))
}

// RetrieveKeysCreate fetches the API to find the corresponding tx and returns a tx wrapper or an error.
func (t Transactor) RetrieveKeysCreate(companyChainId string, uuid string, page int, txPerPage int) (*entityApi.TxWrappers, error) {
    return t.apiHandler.GetTxs(account.GetQueryKeyCreateID(), common.FormatTxid(companyChainId, uuid), page, txPerPage)
}

// RetrieveKeysRevoke fetches the API to find the corresponding tx and returns a tx wrapper or an error.
func (t Transactor) RetrieveKeysRevoke(companyChainId string, uuid string, page int, txPerPage int) (*entityApi.TxWrappers, error) {
    return t.apiHandler.GetTxs(account.GetQueryKeyRevokeID(), common.FormatTxid(companyChainId, uuid), page, txPerPage)
}

// Retrieve any Txs fetches the API to find the corresponding tx and returns a tx wrapper or an error.
func (t Transactor) RetrieveTxs(queryKey string, companyChainId string, uuid string, page int, txPerPage int) (*entityApi.TxWrappers, error) {
    return t.apiHandler.GetTxs(queryKey, common.FormatTxid(companyChainId, uuid), page, txPerPage)
}

// RetrieveCertificatesHistory fetches the API to find the corresponding txs and returns tx wrappers or an error.
func (t Transactor) RetrieveCertificatesHistory(companyChainId string, uuid string, page int, txPerPage int) (*entityApi.TxWrappers, error) {
    return t.apiHandler.RetrieveCertificatesHistory(common.FormatTxid(companyChainId, uuid), page, txPerPage)
}

// SendSecretNaclBoxV1 creates a SecretNaclBox (V1), wraps it in a tx and sends it to the API.
func (t Transactor) SendSecretNaclBoxV1(
    uuid string,
    sender nacl.PublicKey,
    nonce nacl.BoxNonce,
    content []byte,
) (*entityApi.TxStatus, error) {
    secret := certify.SecretNaclBoxV1{
        Id:      common.FormatTxid(t.companyChainId, uuid),
        Sender:  sender,
        Nonce:   nonce,
        Content: content,
    }

    return t.SendTx(secret)
}

// RetrieveSecrets fetches the API to find the corresponding txs and returns tx wrappers or an error.
func (t Transactor) RetrieveSecrets(companyChainId string, uuid string, page int, txPerPage int) (*entityApi.TxWrappers, error) {
    return t.apiHandler.RetrieveSecrets(common.FormatTxid(companyChainId, uuid), page, txPerPage)
}

// GetTx signs a tx data and returns a new tx ready to be sent.
func (t Transactor) SendTx(txData entity.TxData) (status *entityApi.TxStatus, err error) {
    if t.txSigner == nil || t.chainId == "" {
        return nil, errors.New("impossible to create txs without a private key or chain id")
    }

    tx := t.apiHandler.SignTx(t.txSigner, t.chainId, t.getCurrentNonceTime(), txData)
    txBytes, err := t.apiHandler.EncodeTx(tx)
    if err != nil {
        return nil, err
    }
    return t.apiHandler.SendTx(txBytes)
}

func (t Transactor) getCurrentNonceTime() entity.Time {
    return entity.Time{
        Time: time.Now(),
    }
}

func (t Transactor) GenerateUuidV4() string {
    return uuid.New().String()
}
