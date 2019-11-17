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

type GetTaskStatusResponse struct {
	// Unique ID of the Task
	Id string `json:"id,omitempty"`
	// UTC time of creation of the Task in RFC-3339 format
	CreatedAt string `json:"created_at,omitempty"`
	// UTC last time the Task was updated in RFC-3339 format
	UpdatedAt string `json:"updated_at,omitempty"`
	Result Result `json:"result,omitempty"`
	State TaskState `json:"state,omitempty"`
	// The Blockchain-generated hash of the Transaction; populated after the Blockchain Node accepted the Transaction
	TxHash string `json:"tx_hash,omitempty"`
	// The QEDIT-generated hash of the Transaction; generated after proof generation, but prior to Broadcast by the QEDIT Node
	QeditTxHash string `json:"qedit_tx_hash,omitempty"`
	Type TaskType `json:"type,omitempty"`
	Data TaskData `json:"data,omitempty"`
	// In case of failure this field reports the reason for the failure
	Error string `json:"error,omitempty"`
}
