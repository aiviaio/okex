package subaccount

import (
	"github.com/aiviaio/okex"
)

type (
	SubAccount struct {
		SubAcct     string        `json:"subAcct,omitempty"`
		Type        string        `json:"type"`
		UUID        string        `json:"uid"`
		Label       string        `json:"label,omitempty"`
		Mobile      string        `json:"mobile,omitempty"`
		GAuth       bool          `json:"gAuth"`
		Enable      bool          `json:"enable"`
		TS          okex.JSONTime `json:"ts"`
		CanTransOut string        `json:"canTransOut"`
	}
	APIKey struct {
		SubAcct    string        `json:"subAcct,omitempty"`
		Label      string        `json:"label,omitempty"`
		APIKey     string        `json:"apiKey,omitempty"`
		SecretKey  string        `json:"secretKey,omitempty"`
		Passphrase string        `json:"Passphrase,omitempty"`
		Perm       string        `json:"perm,omitempty"`
		IP         string        `json:"ip,omitempty"`
		TS         okex.JSONTime `json:"ts,omitempty"`
	}
	ResetAPIKey struct {
		SubAcct string        `json:"subAcct,omitempty"`
		Label   string        `json:"label,omitempty"`
		ApiKey  string        `json:"apiKey,omitempty"`
		Perm    string        `json:"perm,omitempty"`
		IP      string        `json:"ip,omitempty"`
		TS      okex.JSONTime `json:"ts,omitempty"`
	}
	HistoryTransfer struct {
		SubAcct string         `json:"subAcct,omitempty"`
		Ccy     string         `json:"ccy,omitempty"`
		BillID  okex.JSONInt64 `json:"billId,omitempty"`
		Type    okex.BillType  `json:"type,omitempty,string"`
		TS      okex.JSONTime  `json:"ts,omitempty"`
	}
	Transfer struct {
		TransID okex.JSONInt64 `json:"transId"`
	}
	ListSubAccount struct {
		TotalPage string                  `json:"totalPage"`
		Page      string                  `json:"page"`
		Details   []*ListSubAccountDetail `json:"details"`
	}
	ListSubAccountDetail struct {
		SubAcct string        `json:"subAcct"`
		UUID    string        `json:"uid"`
		Label   string        `json:"label"`
		AcctLv  string        `json:"acctLv"`
		TS      okex.JSONTime `json:"ts"`
	}
	CreateSubAccount struct {
		SubAcct string        `json:"subAcct"`
		Label   string        `json:"label"`
		AcctLv  string        `json:"acctLv"`
		UID     string        `json:"uid"`
		TS      okex.JSONTime `json:"ts"`
	}
	DeleteSubAccount struct {
		SubAcct string `json:"subAcct"`
	}
	CreatAPIKeySubAccount struct {
		SubAcct    string        `json:"subAcct,omitempty"`
		Label      string        `json:"label,omitempty"`
		APIKey     string        `json:"apiKey,omitempty"`
		SecretKey  string        `json:"secretKey,omitempty"`
		Passphrase string        `json:"passphrase,omitempty"`
		Perm       string        `json:"perm,omitempty"`
		IP         string        `json:"ip,omitempty"`
		TS         okex.JSONTime `json:"ts,omitempty"`
	}
	UpdateAPIKEySubAccount struct {
		SubAcct string        `json:"subAcct,omitempty"`
		Label   string        `json:"label,omitempty"`
		APIKey  string        `json:"apiKey,omitempty"`
		Perm    string        `json:"perm,omitempty"`
		IP      string        `json:"ip,omitempty"`
		TS      okex.JSONTime `json:"ts,omitempty"`
	}
	DeleteAPIKeySubAccount struct {
		SubAcct string `json:"subAcct,omitempty"`
	}
	SetLevelSubAccount struct {
		AcctLv string `json:"acctLv"` // Account level 1: Simple 2: Single-currency margin 3: Multi-currency margin 4：Portfolio margin
	}
	GetFeeRatesSubAccount struct {
		Level     string        `json:"level"`
		Taker     string        `json:"taker"`
		Maker     string        `json:"maker"`
		TakerU    string        `json:"takerU"`
		MakerU    string        `json:"makerU"`
		Delivery  string        `json:"delivery"`
		Exercise  string        `json:"exercise"`
		InstType  string        `json:"instType"`
		TakerUSDC string        `json:"takerUSDC"`
		MakerUSDC string        `json:"makerUSDC"`
		TS        okex.JSONTime `json:"ts,omitempty"`
		Category  string        `json:"category"`
	}
	SetFeeRateSubAccount struct {
		SubAcct string `json:"subAcct,omitempty"`
		EffDate string `json:"effDate,omitempty"`
	}
	CreateDepositAddress struct {
		Ccy     string        `json:"ccy,omitempty"`
		Addr    string        `json:"addr,omitempty"`
		Chain   string        `json:"chain,omitempty"`
		PmtID   string        `json:"pmtId,omitempty"`
		Tag     string        `json:"tag,omitempty"`
		Memo    string        `json:"memo,omitempty"`
		Comment string        `json:"comment,omitempty"`
		TS      okex.JSONTime `json:"ts,omitempty"`
	}
	UpdateDepositAddress struct {
		Ccy   string        `json:"ccy,omitempty"`
		Addr  string        `json:"addr,omitempty"`
		Chain string        `json:"chain,omitempty"`
		PmtID string        `json:"pmtId,omitempty"`
		Tag   string        `json:"tag,omitempty"`
		Memo  string        `json:"memo,omitempty"`
		TO    string        `json:"to,omitempty"` // 6:Funding, 18:Trading account, Default is 6
		TS    okex.JSONTime `json:"ts,omitempty"`
	}
	GetDepositAddress struct {
		Addr     string `json:"addr,omitempty"`
		Tag      string `json:"tag,omitempty"`
		Memo     string `json:"memo,omitempty"`
		PmtID    string `json:"pmtId,omitempty"`
		Ccy      string `json:"ccy,omitempty"`
		Chain    string `json:"chain,omitempty"`
		TO       string `json:"to,omitempty"` // 6:Funding, 18:Trading account, Default is 6
		Selected bool   `json:"selected"`     // Return true if the current deposit address is selected on the website page
		CtAddr   string `json:"ctAddr,omitempty"`
	}
	GetDepositHistory struct {
		SubAcct             string `json:"subAcct,omitempty"`
		Ccy                 string `json:"ccy,omitempty"`
		Chain               string `json:"chain,omitempty"`
		Amt                 string `json:"amt,omitempty"`
		From                string `json:"from,omitempty"`
		AreaCodeFrom        string `json:"areaCodeFrom,omitempty"`
		TO                  string `json:"to,omitempty"` // 6:Funding, 18:Trading account, Default is 6
		TxId                string `json:"txId,omitempty"`
		TS                  string `json:"ts,omitempty"`
		State               string `json:"state,omitempty"`
		DepId               string `json:"depId,omitempty"`
		ActualDepBlkConfirm string `json:"actualDepBlkConfirm,omitempty"`
	}
	WithdrawHistory struct {
		SubAcct          string               `json:"subAcct"`
		Ccy              string               `json:"ccy"`
		Chain            string               `json:"chain"`
		NonTradableAsset string               `json:"nonTradableAsset"`
		Amt              okex.JSONFloat64     `json:"amt"`
		TS               okex.JSONTime        `json:"ts"`
		To               string               `json:"to"`
		AreaCodeTo       string               `json:"areaCodeTo,omitempty"`
		Tag              string               `json:"tag,omitempty"`
		PmtID            string               `json:"pmtId,omitempty"`
		Memo             string               `json:"memo,omitempty"`
		AddrEx           string               `json:"addrEx,omitempty"`
		TxID             string               `json:"txId"`
		Fee              okex.JSONFloat64     `json:"fee"`
		FeeCcy           string               `json:"feeCcy"`
		State            okex.WithdrawalState `json:"state,string"`
		WdID             okex.JSONInt64       `json:"wdId"`
		ClientID         string               `json:"clientId"`
	}
)
