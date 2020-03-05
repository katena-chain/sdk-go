/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package client

import (
	"errors"

	"github.com/transchain/sdk-go/crypto/ed25519"
	"github.com/transchain/sdk-go/crypto/nacl"

	"github.com/katena-chain/sdk-go/api"
	"github.com/katena-chain/sdk-go/entity"
	"github.com/katena-chain/sdk-go/entity/account"
	entityApi "github.com/katena-chain/sdk-go/entity/api"
	"github.com/katena-chain/sdk-go/entity/certify"
)

// Transactor provides helper methods to hide the complexity of Tx creation, signature and API dialog.
type Transactor struct {
	apiHandler  *api.Handler
	chainID     string
	txSigner    *ed25519.PrivateKey
	companyBcid string
}

// Transactor constructor.
func NewTransactor(apiUrl string, chainID string, companyBcid string, txSigner *ed25519.PrivateKey) *Transactor {
	return &Transactor{
		apiHandler:  api.NewHandler(apiUrl),
		chainID:     chainID,
		txSigner:    txSigner,
		companyBcid: companyBcid,
	}
}

// SendCertificateRawV1 creates a CertificateRaw (V1) and sends it to the API.
func (t Transactor) SendCertificateRawV1(uuid string, value []byte) (*entityApi.TxStatus, error) {
	certificate := certify.NewCertificateRawV1(entity.FormatTxid(t.companyBcid, uuid), value)
	return t.SendTx(certificate)
}

// SendCertificateEd25519V1 creates a CertificateEd25519 (V1) and sends it to the API.
func (t Transactor) SendCertificateEd25519V1(uuid string, signer ed25519.PublicKey, signature ed25519.Signature) (*entityApi.TxStatus, error) {
	certificate := certify.NewCertificateEd25519V1(entity.FormatTxid(t.companyBcid, uuid), signer, signature)
	return t.SendTx(certificate)
}

// SendKeyCreateV1 creates a KeyCreate (V1) and sends it to the API.
func (t Transactor) SendKeyCreateV1(uuid string, publicKey ed25519.PublicKey, role string) (*entityApi.TxStatus, error) {
	keyCreate := account.NewKeyCreateV1(entity.FormatTxid(t.companyBcid, uuid), publicKey, role)
	return t.SendTx(keyCreate)
}

// SendKeyRevokeV1 creates a KeyRevoke (V1) and sends it to the API.
func (t Transactor) SendKeyRevokeV1(uuid string, publicKey ed25519.PublicKey) (*entityApi.TxStatus, error) {
	keyRevoke := account.NewKeyRevokeV1(entity.FormatTxid(t.companyBcid, uuid), publicKey)
	return t.SendTx(keyRevoke)
}

// SendSecretNaclBoxV1 creates a SecretNaclBox (V1) and sends it to the API.
func (t Transactor) SendSecretNaclBoxV1(uuid string, sender nacl.PublicKey, nonce nacl.BoxNonce, content []byte) (*entityApi.TxStatus, error) {
	secret := certify.NewSecretNaclBoxV1(entity.FormatTxid(t.companyBcid, uuid), sender, nonce, content)
	return t.SendTx(secret)
}

// RetrieveCertificate fetches the API and returns a tx wrapper or an error.
func (t Transactor) RetrieveCertificate(companyBcid string, uuid string) (*entityApi.TxWrapper, error) {
	return t.apiHandler.RetrieveCertificate(entity.FormatTxid(companyBcid, uuid))
}

// RetrieveCertificatesHistory fetches the API and returns a tx wrapper list or an error.
func (t Transactor) RetrieveCertificatesHistory(companyBcid string, uuid string, page int, txPerPage int) (*entityApi.TxWrappers, error) {
	return t.apiHandler.RetrieveCertificatesHistory(entity.FormatTxid(companyBcid, uuid), page, txPerPage)
}

// RetrieveKeysCreate fetches the API and returns a tx wrapper list or an error.
func (t Transactor) RetrieveKeysCreate(companyBcid string, uuid string, page int, txPerPage int) (*entityApi.TxWrappers, error) {
	return t.apiHandler.RetrieveTxs(account.GetCategoryKeyCreate(), entity.FormatTxid(companyBcid, uuid), page, txPerPage)
}

// RetrieveKeysRevoke fetches the API and returns a tx wrapper list or an error.
func (t Transactor) RetrieveKeysRevoke(companyBcid string, uuid string, page int, txPerPage int) (*entityApi.TxWrappers, error) {
	return t.apiHandler.RetrieveTxs(account.GetCategoryKeyRevoke(), entity.FormatTxid(companyBcid, uuid), page, txPerPage)
}

// RetrieveCompanyKeys fetches the API and returns the list of keyV1 for a company or an error.
func (t Transactor) RetrieveCompanyKeys(companyBcid string, page int, txPerPage int) ([]*account.KeyV1, error) {
	return t.apiHandler.RetrieveCompanyKeys(companyBcid, page, txPerPage)
}

// RetrieveSecrets fetches the API and returns a tx wrapper list or an error.
func (t Transactor) RetrieveSecrets(companyBcid string, uuid string, page int, txPerPage int) (*entityApi.TxWrappers, error) {
	return t.apiHandler.RetrieveSecrets(entity.FormatTxid(companyBcid, uuid), page, txPerPage)
}

// RetrieveTxs fetches the API and returns a tx wrapper list or an error.
func (t Transactor) RetrieveTxs(txCategory string, companyBcid string, uuid string, page int, txPerPage int) (*entityApi.TxWrappers, error) {
	return t.apiHandler.RetrieveTxs(txCategory, entity.FormatTxid(companyBcid, uuid), page, txPerPage)
}

// SendTx signs, encodes and send a tx to the Api.
func (t Transactor) SendTx(txData entity.TxData) (status *entityApi.TxStatus, err error) {
	if t.txSigner == nil || t.chainID == "" {
		return nil, errors.New("impossible to create txs without a private key or chain id")
	}

	tx := t.apiHandler.SignTx(t.txSigner, t.chainID, entity.GetCurrentTime(), txData)
	txBytes, err := t.apiHandler.EncodeTx(tx)
	if err != nil {
		return nil, err
	}
	return t.apiHandler.SendTx(txBytes)
}
