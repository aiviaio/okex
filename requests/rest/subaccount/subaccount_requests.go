package subaccount

import "github.com/aiviaio/okex"

type (
	ViewList struct {
		SubAcct string `json:"subAcct,omitempty"`
		Enable  bool   `json:"enable,omitempty"`
		After   int64  `json:"after,omitempty,string"`
		Before  int64  `json:"before,omitempty,string"`
		Limit   int64  `json:"limit,omitempty,string"`
	}
	CreateAPIKey struct {
		Pwd        string            `json:"pwd"`
		SubAcct    string            `json:"subAcct"`
		Label      string            `json:"label"`
		Passphrase string            `json:"Passphrase"`
		IP         []string          `json:"ip,omitempty"`
		Perm       okex.APIKeyAccess `json:"perm,omitempty"`
	}
	ResetAPIKey struct {
		SubAcct string            `json:"subAcct"`
		ApiKey  string            `json:"apiKey"`
		Label   string            `json:"label"`
		IP      []string          `json:"ip,omitempty"`
		Perm    okex.APIKeyAccess `json:"perm,omitempty"`
	}
	QueryAPIKey struct {
		APIKey  string `json:"apiKey"`
		SubAcct string `json:"subAcct"`
	}
	DeleteAPIKey struct {
		Pwd     string `json:"pwd"`
		APIKey  string `json:"apiKey"`
		SubAcct string `json:"subAcct"`
	}
	GetBalance struct {
		SubAcct string `json:"subAcct"`
	}
	GetBalancesFunding struct {
		SubAcct string `json:"subAcct"`
		Ccy     string `json:"ccy,omitempty"`
	}
	HistoryTransfer struct {
		Ccy     string            `json:"ccy,omitempty"`
		SubAcct string            `json:"subAcct,omitempty"`
		After   int64             `json:"after,omitempty,string"`
		Before  int64             `json:"before,omitempty,string"`
		Limit   int64             `json:"limit,omitempty,string"`
		Type    okex.TransferType `json:"type,omitempty,string"`
	}
	ManageTransfers struct {
		Ccy            string           `json:"ccy"`
		Amt            float64          `json:"amt,string"`
		From           okex.AccountType `json:"from,string"`
		To             okex.AccountType `json:"to,string"`
		FromSubAccount string           `json:"fromSubAccount"`
		ToSubAccount   string           `json:"toSubAccount"`
		LoanTrans      bool             `json:"loanTrans"`
		OmitPosRisk    string           `json:"omitPosRisk"`
	}
	ListSubAccount struct {
		SubAcct string `json:"subAcct,omitempty"`
		UUID    string `json:"uid,omitempty"`
		Page    string `json:"page,omitempty"`
		Limit   string `json:"limit,omitempty"` // Number of results per request. The maximum is 100; the default is 100
	}
	CreateSubAccount struct {
		SubAcct string `json:"subAcct"`
		Label   string `json:"label,omitempty"`
	}
	DeleteSubAccount struct {
		SubAcct string `json:"subAcct"`
	}
	CreatAPIKeySubAccount struct {
		SubAcct    string            `json:"subAcct"`
		Label      string            `json:"label"`
		Passphrase string            `json:"passphrase"`
		IP         []string          `json:"ip,omitempty"`
		Perm       okex.APIKeyAccess `json:"perm,omitempty"`
	}
	UpdateAPIKEySubAccount struct {
		SubAcct string `json:"subAcct"`
		Label   string `json:"label,omitempty"`
		APIKey  string `json:"apiKey"`
		Perm    string `json:"perm,omitempty"`
		IP      string `json:"ip,omitempty"`
	}
	DeleteAPIKeySubAccount struct {
		SubAcct string `json:"subAcct"`
		APIKey  string `json:"apiKey"`
	}
	SetLevelSubAccount struct {
		SubAcct string `json:"subAcct"`
		AcctLv  string `json:"acctLv"` //Account level 1: Simple 2: Single-currency margin 3: Multi-currency margin 4：Portfolio margin
	}
	GetFeeRatesSubAccount struct {
		InstType   string `json:"instType"`             // SPOT MARGIN SWAP FUTURES OPTION
		InstID     string `json:"instId,omitempty"`     // Instrument ID, e.g. BTC-USDT Applicable to SPOT/MARGIN
		Uly        string `json:"uly,omitempty"`        // Underlying, e.g. BTC-USD Applicable to FUTURES/SWAP/OPTION
		InstFamily string `json:"instFamily,omitempty"` // Instrument family, e.g. BTC-USDApplicable to FUTURES/SWAP/OPTION
	}
	SetFeeRateSubAccount struct {
		SubAcct  string `json:"subAcct,omitempty"`
		InstType string `json:"instType,omitempty"`
		MgnType  string `json:"mgnType,omitempty"`
		ChgType  string `json:"chgType"`
		ChgTaker string `json:"chgTaker"`
		ChgMaker string `json:"chgMaker"`
		EffDate  string `json:"effDate,omitempty"`
	}
	CreateDepositAddress struct {
		SubAcct  string           `json:"subAcct"`
		Ccy      string           `json:"ccy"`
		Chain    string           `json:"chain,omitempty"`
		AddrType string           `json:"addrType,omitempty"` // 1: Regular address, 2:SegWit address (Only applicable to BTC/LTC), Default is 1
		TO       okex.AccountType `json:"to,string"`          // 6:Funding, 18:Trading account, Default is 6
	}
	UpdateDepositAddress struct {
		SubAcct  string `json:"subAcct"`
		Ccy      string `json:"ccy"`
		Chain    string `json:"chain,omitempty"`
		AddrType string `json:"addrType"` // 1: Regular address, 2:SegWit address (Only applicable to BTC/LTC), Default is 1
		TO       string `json:"to"`       // 6:Funding, 18:Trading account, Default is 6
	}
	GetDepositAddress struct {
		SubAcct string `json:"subAcct"`
		Ccy     string `json:"ccy"`
	}
	GetDepositHistory struct {
		SubAcct string `json:"subAcct"`
		Ccy     string `json:"ccy"`
		TxId    string `json:"txId"`
		Type    string `json:"type"`
		State   string `json:"state"`
		After   string `json:"after"`
		Before  string `json:"before"`
		Limit   string `json:"limit"`
	}
	GetWithdrawHistory struct {
		SubAcct  string `json:"subAcct"`
		Ccy      string `json:"ccy"`
		WdID     string `json:"wdId"`
		ClientId string `json:"clientId"`
		TxID     string `json:"txId"`
		Type     string `json:"type"`
		State    string `json:"state"`
		After    string `json:"after"`
		Before   string `json:"before"`
		Limit    string `json:"limit"`
	}
)
