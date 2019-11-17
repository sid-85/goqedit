/*
 * QEDIT - Asset Transfers
 *
 * This SDK provides a programmatic way for interacting with QEDIT's _Asset Transfer_ API. The specification definition file is publicly available [in this repository](https://github.com/QED-it/asset_transfers_dev_guide/). 
 *
 * API version: 1.6.1
 * Contact: dev@qed-it.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package goqedit

type GetWalletActivityResponse struct {
	WalletId string `json:"wallet_id,omitempty"`
	// The list of retrieved transactions
	Transactions []AnalyticWalletTx `json:"transactions,omitempty"`
}
