/*
 * QEDIT - Asset Transfers
 *
 * This SDK provides a programmatic way for interacting with QEDIT's _Asset Transfer_ API. The specification definition file is publicly available [in this repository](https://github.com/QED-it/asset_transfers_dev_guide/). 
 *
 * API version: 1.6.1
 * Contact: dev@qed-it.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package sdk

type ImportAuditorAccessWalletRequest struct {
	// The ID under which to import the Wallet; can be different from the ID the Wallet was stored under in the exporting Node
	WalletId string `json:"wallet_id"`
	// The public key of the imported Wallet
	PublicKey string `json:"public_key"`
	// the encrypted viewing key of the imported Wallet
	AccessKey string `json:"access_key"`
}
