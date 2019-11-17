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

type AnalyticsOutput struct {
	// Boolean signifying whether the Issuance was done confidentially; false if the Issuance was done publicly
	IsConfidential bool `json:"is_confidential,omitempty"`
	PublicIssuanceDescription AnalyticsPublicIssuanceDescription `json:"public_issuance_description,omitempty"`
	ConfidentialIssuanceDescription AnalyticsConfidentialIssuanceDescription `json:"confidential_issuance_description,omitempty"`
	OutputDescription AnalyticsOutputDescription `json:"output_description,omitempty"`
}
