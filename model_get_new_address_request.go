/*
 * QEDIT - Asset Transfers
 *
 * This SDK provides a programmatic way for interacting with QEDIT's _Asset Transfer_ API. The specification definition file is publicly available [in this repository](https://github.com/QED-it/asset_transfers_dev_guide/).
 *
 * API version: 1.7.2
 * Contact: dev@qed-it.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package goqedit

type GetNewAddressRequest struct {
	// The ID of the Wallet for which to generate an Address
	WalletId string `json:"wallet_id"`
	// An optional 11-byte (22 hexadecimal characters) input which is used to generate different Addresses. A unique Address will be generated for each different diversifier used. If omitted, the Node selects a random diversifier.
	Diversifier string `json:"diversifier,omitempty"`
}
