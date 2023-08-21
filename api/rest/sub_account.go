package rest

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/aiviaio/okex"
	requests "github.com/aiviaio/okex/requests/rest/subaccount"
	responses "github.com/aiviaio/okex/responses/sub_account"
)

// SubAccount
//
// https://www.okex.com/docs-v5/en/#rest-api-subaccount
type SubAccount struct {
	client *ClientRest
}

// NewSubAccount returns a pointer to a fresh SubAccount
func NewSubAccount(c *ClientRest) *SubAccount {
	return &SubAccount{c}
}

// ViewList
// applies to master accounts only
//
// https://www.okx.com/docs-v5/en/#sub-account-rest-api-get-sub-account-list
func (c *SubAccount) ViewList(req requests.ViewList) (response responses.ViewList, err error) {
	p := "/api/v5/users/subaccount/list"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}

// CreateAPIKey
// applies to master accounts only
//
// https://www.okex.com/docs-v5/en/#rest-api-subaccount-create-an-apikey-for-a-sub-account
func (c *SubAccount) CreateAPIKey(req requests.CreateAPIKey) (response responses.APIKey, err error) {
	p := "/api/v5/users/subaccount/apikey"
	m := okex.S2M(req)
	if len(req.IP) > 0 {
		m["ip"] = strings.Join(req.IP, ",")
	}
	res, err := c.client.Do(http.MethodPost, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}

// QueryAPIKey
// applies to master accounts only
//
// https://www.okex.com/docs-v5/en/#rest-api-subaccount-query-the-apikey-of-a-sub-account
func (c *SubAccount) QueryAPIKey(req requests.QueryAPIKey) (response responses.APIKey, err error) {
	p := "/api/v5/users/subaccount/apikey"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}

// ResetAPIKey
// applies to master accounts only
//
// https://www.okex.com/docs-v5/en/#rest-api-subaccount-reset-the-apikey-of-a-sub-account
func (c *SubAccount) ResetAPIKey(req requests.CreateAPIKey) (response responses.APIKey, err error) {
	p := "/api/v5/users/subaccount/modify-apikey"
	m := okex.S2M(req)
	if len(req.IP) > 0 {
		m["ip"] = strings.Join(req.IP, ",")
	}
	res, err := c.client.Do(http.MethodPost, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}

// DeleteAPIKey
// applies to master accounts only
//
// https://www.okex.com/docs-v5/en/#rest-api-subaccount-delete-the-apikey-of-sub-accounts
func (c *SubAccount) DeleteAPIKey(req requests.DeleteAPIKey) (response responses.APIKey, err error) {
	p := "/api/v5/users/subaccount/delete-apikey"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodPost, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}

// GetBalance
// Query detailed balance info of Trading Account of a sub-account via the master account
// (applies to master accounts only)
//
// https://www.okex.com/docs-v5/en/#rest-api-subaccount-get-sub-account-balance
func (c *SubAccount) GetBalance(req requests.GetBalance) (response responses.GetBalance, err error) {
	p := "/api/v5/account/subaccount/balances"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}

// GetBalancesFunding
// Query detailed balance info of Funding Account of a sub-account via the master account (applies to master accounts only)
//
// https://www.okx.com/docs-v5/en/#sub-account-rest-api-get-sub-account-funding-balance
func (c *SubAccount) GetBalancesFunding(req requests.GetBalancesFunding) (response responses.GetBalancesFunding, err error) {
	p := "/api/v5/asset/subaccount/balances"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}

// HistoryTransfer
// applies to master accounts only
//
// https://www.okex.com/docs-v5/en/#rest-api-subaccount-history-of-sub-account-transfer
func (c *SubAccount) HistoryTransfer(req requests.HistoryTransfer) (response responses.HistoryTransfer, err error) {
	p := "/api/v5/account/subaccount/bills"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}

// ManageTransfers
// applies to master accounts only
//
// https://www.okex.com/docs-v5/en/#rest-api-subaccount-master-accounts-manage-the-transfers-between-sub-accounts
func (c *SubAccount) ManageTransfers(req requests.ManageTransfers) (response responses.ManageTransfer, err error) {
	p := "/api/v5/account/subaccount/transfer"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodPost, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}

// ListSubAccount
// applies to master accounts only
//
// https://www.okx.com/docs-v5/broker_en/#non-disclosed-broker-api-get-sub-account-list
func (c *SubAccount) ListSubAccount(req requests.ListSubAccount) (response responses.ListSubAccount, err error) {
	p := "/api/v5/broker/nd/subaccount-info"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}

// CreateSubAccount
// applies to master accounts only
//
// https://www.okx.com/docs-v5/broker_en/#non-disclosed-broker-api-create-sub-account
func (c *SubAccount) CreateSubAccount(req requests.CreateSubAccount) (response responses.CreateSubAccount, err error) {
	p := "/api/v5/broker/nd/create-subaccount"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodPost, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}

// DeleteSubAccount
// applies to master accounts only
//
// https://www.okx.com/docs-v5/broker_en/#non-disclosed-broker-api-delete-sub-account
func (c *SubAccount) DeleteSubAccount(req requests.DeleteSubAccount) (response responses.DeleteSubAccount, err error) {
	p := "/api/v5/broker/nd/delete-subaccount"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodPost, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}

// CreateAPIKeySubAccount
// applies to master accounts only
//
// https://www.okx.com/docs-v5/broker_en/#non-disclosed-broker-api-create-an-api-key-for-a-sub-account
func (c *SubAccount) CreateAPIKeySubAccount(req requests.CreatAPIKeySubAccount) (
	response responses.CreatAPIKeySubAccount, err error) {
	p := "/api/v5/broker/nd/subaccount/apikey"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodPost, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}

// UpdateAPIKeySubAccount
// applies to master accounts only
//
// https://www.okx.com/docs-v5/broker_en/#non-disclosed-broker-api-reset-the-api-key-of-a-sub-account
func (c *SubAccount) UpdateAPIKeySubAccount(req requests.UpdateAPIKEySubAccount) (
	response responses.UpdateAPIKEySubAccount, err error) {
	p := "/api/v5/broker/nd/subaccount/modify-apikey"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodPost, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}

// DeleteAPIKEySubAccount
// applies to master accounts only
//
// https://www.okx.com/docs-v5/broker_en/#non-disclosed-broker-api-delete-the-api-key-of-sub-accounts
func (c *SubAccount) DeleteAPIKEySubAccount(req requests.DeleteAPIKeySubAccount) (
	response responses.DeleteAPIKeySubAccount, err error) {
	p := "/api/v5/broker/nd/subaccount/delete-apikey"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodPost, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}

// SetLevelSubAccount
// applies to master accounts only
//
// https://www.okx.com/docs-v5/broker_en/#non-disclosed-broker-api-set-the-account-level-of-the-sub-account
func (c *SubAccount) SetLevelSubAccount(req requests.SetLevelSubAccount) (response responses.SetLevelSubAccount,
	err error) {
	p := "/api/v5/broker/nd/set-subaccount-level"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodPost, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}

// GetFeeRatesSubAccount
//
// https://www.okx.com/docs-v5/broker_en/#non-disclosed-broker-api-get-sub-account-deposit-address
func (c *SubAccount) GetFeeRatesSubAccount(req requests.GetFeeRatesSubAccount) (response responses.GetFeeRatesSubAccount,
	err error) {
	p := "/api/v5/account/trade-fee"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}

// SetFeeRateSubAccount
// applies to master accounts only
//
// https://www.okx.com/docs-v5/broker_en/#non-disclosed-broker-api-set-trading-fee-rate-for-the-sub-account
func (c *SubAccount) SetFeeRateSubAccount(req requests.SetFeeRateSubAccount) (response responses.SetFeeRateSubAccount,
	err error) {
	p := "/api/v5/broker/nd/set-subaccount-fee-rate"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodPost, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}

// CreateDepositAddressSubAccount
// applies to master accounts only
//
// https://www.okx.com/docs-v5/broker_en/#non-disclosed-broker-api-create-deposit-address-for-sub-account
func (c *SubAccount) CreateDepositAddressSubAccount(req requests.CreateDepositAddress) (
	response responses.CreateDepositAddress, err error) {
	p := "/api/v5/asset/broker/nd/subaccount-deposit-address"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodPost, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}

// UpdateDepositAddressSubAccount
// applies to master accounts only
//
// https://www.okx.com/docs-v5/broker_en/#non-disclosed-broker-api-modify-sub-account-deposit-address
func (c *SubAccount) UpdateDepositAddressSubAccount(req requests.UpdateDepositAddress) (
	response responses.UpdateDepositAddress, err error) {
	p := "/api/v5/asset/broker/nd/modify-subaccount-deposit-address"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodPost, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}

// GetDepositAddressSubAccount
// applies to master accounts only
//
// https://www.okx.com/docs-v5/broker_en/#non-disclosed-broker-api-get-sub-account-deposit-address
func (c *SubAccount) GetDepositAddressSubAccount(req requests.GetDepositAddress) (response responses.GetDepositAddress,
	err error) {
	p := "/api/v5/asset/broker/nd/subaccount-deposit-address"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}

// GetDepositHistorySubAccount
// applies to master accounts only
//
// https://www.okx.com/docs-v5/broker_en/#non-disclosed-broker-api-get-sub-account-deposit-history
func (c *SubAccount) GetDepositHistorySubAccount(req requests.GetDepositHistory) (response responses.GetDepositHistory,
	err error) {
	p := "/api/v5/asset/broker/nd/subaccount-deposit-history"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}
