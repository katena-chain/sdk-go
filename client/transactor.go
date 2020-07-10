/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package client

import (
	"github.com/katena-chain/sdk-go/api"
	"github.com/katena-chain/sdk-go/crypto/ed25519"
	"github.com/katena-chain/sdk-go/crypto/nacl"
	"github.com/katena-chain/sdk-go/entity"
	"github.com/katena-chain/sdk-go/entity/account"
	entityApi "github.com/katena-chain/sdk-go/entity/api"
	"github.com/katena-chain/sdk-go/entity/certify"
	"github.com/katena-chain/sdk-go/entity/common"
)

// Transactor provides helper methods to hide the complexity of Tx creation, signature and API dialog.
type Transactor struct {
	apiHandler *api.Handler
	chainId    string
	txSigner   *entity.TxSigner
}

// Transactor constructor.
func NewTransactor(apiUrl string, chainId string, txSigner *entity.TxSigner) *Transactor {
	return &Transactor{
		apiHandler: api.NewHandler(apiUrl),
		chainId:    chainId,
		txSigner:   txSigner,
	}
}

// SendCertificateRawV1Tx creates a CertificateRawV1 TxData and sends it to the API.
func (t Transactor) SendCertificateRawV1Tx(id string, value []byte) (*entityApi.SendTxResult, error) {
	certificate := certify.NewCertificateRawV1(id, value)
	return t.SendTx(certificate)
}

// SendCertificateEd25519V1Tx creates a CertificateEd25519V1 TxData and sends it to the API.
func (t Transactor) SendCertificateEd25519V1Tx(id string, signer ed25519.PublicKey, signature ed25519.Signature) (*entityApi.SendTxResult, error) {
	certificate := certify.NewCertificateEd25519V1(id, signer, signature)
	return t.SendTx(certificate)
}

// SendSecretNaclBoxV1Tx creates a SecretNaclBoxV1 TxData and sends it to the API.
func (t Transactor) SendSecretNaclBoxV1Tx(id string, sender nacl.PublicKey, nonce nacl.BoxNonce, content []byte) (*entityApi.SendTxResult, error) {
	secret := certify.NewSecretNaclBoxV1(id, sender, nonce, content)
	return t.SendTx(secret)
}

// SendKeyCreateV1Tx creates a KeyCreateV1 TxData and sends it to the API.
func (t Transactor) SendKeyCreateV1Tx(id string, publicKey ed25519.PublicKey, role string) (*entityApi.SendTxResult, error) {
	keyCreate := account.NewKeyCreateV1(id, publicKey, role)
	return t.SendTx(keyCreate)
}

// SendKeyRotateV1Tx creates a KeyRotateV1 TxData and sends it to the API.
func (t Transactor) SendKeyRotateV1Tx(id string, publicKey ed25519.PublicKey) (*entityApi.SendTxResult, error) {
	keyRotate := account.NewKeyRotateV1(id, publicKey)
	return t.SendTx(keyRotate)
}

// SendKeyRevokeV1Tx creates a KeyRevokeV1 TxData and sends it to the API.
func (t Transactor) SendKeyRevokeV1Tx(id string) (*entityApi.SendTxResult, error) {
	keyRevoke := account.NewKeyRevokeV1(id)
	return t.SendTx(keyRevoke)
}

// SendTx creates a tx from a tx data and the provided tx signer info and chain id, signs it, encodes it and sends it
// to the API.
func (t Transactor) SendTx(txData entity.TxData) (status *entityApi.SendTxResult, err error) {
	return t.apiHandler.SendTx(txData, t.txSigner, t.chainId)
}

// RetrieveCertificateTxs fetches the API and returns all txs related to a certificate fqid.
func (t Transactor) RetrieveCertificateTxs(companyBcId string, id string, page int, txPerPage int) (*entityApi.TxResults, error) {
	return t.apiHandler.RetrieveCertificateTxs(common.ConcatFqId(companyBcId, id), page, txPerPage)
}

// RetrieveLastCertificateTx fetches the API and returns the last tx related to a certificate fqid.
func (t Transactor) RetrieveLastCertificateTx(companyBcId string, id string) (*entityApi.TxResult, error) {
	return t.apiHandler.RetrieveLastCertificateTx(common.ConcatFqId(companyBcId, id))
}

// RetrieveSecretTxs fetches the API and returns all txs related to a secret fqid.
func (t Transactor) RetrieveSecretTxs(companyBcId string, id string, page int, txPerPage int) (*entityApi.TxResults, error) {
	return t.apiHandler.RetrieveSecretTxs(common.ConcatFqId(companyBcId, id), page, txPerPage)
}

// RetrieveLastSecretTx fetches the API and returns the last tx related to a secret fqid.
func (t Transactor) RetrieveLastSecretTx(companyBcId string, id string) (*entityApi.TxResult, error) {
	return t.apiHandler.RetrieveLastSecretTx(common.ConcatFqId(companyBcId, id))
}

// RetrieveKeyTxs fetches the API and returns all txs related to a key fqid.
func (t Transactor) RetrieveKeyTxs(companyBcId string, id string, page int, txPerPage int) (*entityApi.TxResults, error) {
	return t.apiHandler.RetrieveKeyTxs(common.ConcatFqId(companyBcId, id), page, txPerPage)
}

// RetrieveKey fetches the API and returns the last tx related to a key fqid.
func (t Transactor) RetrieveLastKeyTx(companyBcId string, id string) (*entityApi.TxResult, error) {
	return t.apiHandler.RetrieveLastKeyTx(common.ConcatFqId(companyBcId, id))
}

// RetrieveKey fetches the API and return any tx by its hash.
func (t Transactor) RetrieveTx(hash string) (*entityApi.TxResult, error) {
	return t.apiHandler.RetrieveTx(hash)
}

// RetrieveCertificate fetches the API and returns a certificate from the state.
func (t Transactor) RetrieveCertificate(companyBcId string, id string) (entity.TxData, error) {
	return t.apiHandler.RetrieveCertificate(common.ConcatFqId(companyBcId, id))
}

// RetrieveSecret fetches the API and returns a secret from the state.
func (t Transactor) RetrieveSecret(companyBcId string, id string) (entity.TxData, error) {
	return t.apiHandler.RetrieveSecret(common.ConcatFqId(companyBcId, id))
}

// RetrieveKey fetches the API and returns a key from the state.
func (t Transactor) RetrieveKey(companyBcId string, id string) (*account.KeyV1, error) {
	return t.apiHandler.RetrieveKey(common.ConcatFqId(companyBcId, id))
}

// RetrieveCompanyKeys fetches the API and returns a list of keys for a company from the state.
func (t Transactor) RetrieveCompanyKeys(companyBcId string, page int, txPerPage int) ([]*account.KeyV1, error) {
	return t.apiHandler.RetrieveCompanyKeys(companyBcId, page, txPerPage)
}
