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

    "github.com/transchain/sdk-go/crypto/ed25519"
    "github.com/transchain/sdk-go/crypto/nacl"

    "github.com/katena-chain/sdk-go/api"
    "github.com/katena-chain/sdk-go/entity"
    entityApi "github.com/katena-chain/sdk-go/entity/api"
    "github.com/katena-chain/sdk-go/entity/certify"
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
        Id:    certify.FormatTxid(t.companyChainId, uuid),
        Value: value,
    }
    tx, err := t.GetTx(certificate)
    if err != nil {
        return nil, err
    }
    return t.apiHandler.SendCertificate(tx)
}

// SendCertificateEd25519V1 creates a CertificateEd25519 (V1), wraps it in a tx and sends it to the API.
func (t Transactor) SendCertificateEd25519V1(
    uuid string,
    signer ed25519.PublicKey,
    signature ed25519.Signature,
) (*entityApi.TxStatus, error) {
    certificate := certify.CertificateEd25519V1{
        Id:        certify.FormatTxid(t.companyChainId, uuid),
        Signer:    signer,
        Signature: signature,
    }
    tx, err := t.GetTx(certificate)
    if err != nil {
        return nil, err
    }
    return t.apiHandler.SendCertificate(tx)
}

// RetrieveCertificate fetches the API to find the corresponding tx and returns a tx wrapper or an error.
func (t Transactor) RetrieveCertificate(companyChainId string, uuid string) (*entityApi.TxWrapper, error) {
    return t.apiHandler.RetrieveCertificate(certify.FormatTxid(companyChainId, uuid))
}

// RetrieveCertificatesHistory fetches the API to find the corresponding txs and returns tx wrappers or an error.
func (t Transactor) RetrieveCertificatesHistory(companyChainId string, uuid string) (*entityApi.TxWrappers, error) {
    return t.apiHandler.RetrieveCertificatesHistory(certify.FormatTxid(companyChainId, uuid))
}

// SendSecretNaclBoxV1 creates a SecretNaclBox (V1), wraps it in a tx and sends it to the API.
func (t Transactor) SendSecretNaclBoxV1(
    uuid string,
    sender nacl.PublicKey,
    nonce nacl.BoxNonce,
    content []byte,
) (*entityApi.TxStatus, error) {
    secret := certify.SecretNaclBoxV1{
        Id:      certify.FormatTxid(t.companyChainId, uuid),
        Sender:  sender,
        Nonce:   nonce,
        Content: content,
    }
    tx, err := t.GetTx(secret)
    if err != nil {
        return nil, err
    }
    return t.apiHandler.SendSecret(tx)
}

// RetrieveSecrets fetches the API to find the corresponding txs and returns tx wrappers or an error.
func (t Transactor) RetrieveSecrets(companyChainId string, uuid string) (*entityApi.TxWrappers, error) {
    return t.apiHandler.RetrieveSecrets(certify.FormatTxid(companyChainId, uuid))
}

// GetTx signs a tx data and returns a new tx ready to be sent.
func (t Transactor) GetTx(txData entity.TxData) (*entity.Tx, error) {
    if t.txSigner == nil || t.companyChainId == "" {
        return nil, errors.New("impossible to create txs without a private key or company chain id")
    }

    nonceTime := entity.Time{
        Time: time.Now(),
    }

    txDataState := entity.GetTxDataStateBytes(t.chainId, nonceTime, txData)
    signature := t.txSigner.Sign(txDataState)

    tx := &entity.Tx{
        NonceTime: nonceTime,
        Data:      txData,
        Signer:    t.txSigner.GetPublicKey(),
        Signature: signature,
    }

    return tx, nil
}
