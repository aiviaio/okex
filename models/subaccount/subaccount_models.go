package subaccount

import (
	"github.com/aiviaio/okex"
)

type (
	SubAccount struct {
		SubAcct string        `json:"subAcct,omitempty"`
		Label   string        `json:"label,omitempty"`
		Mobile  string        `json:"mobile,omitempty"`
		GAuth   bool          `json:"gAuth"`
		Enable  bool          `json:"enable"`
		TS      okex.JSONTime `json:"ts"`
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
	CreateSubAccount struct {
		SubAcct string        `json:"subAcct"`
		Label   string        `json:"label"`
		AcctLv  string        `json:"acctLv"`
		UUID    string        `json:"uuid"`
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
	SetFeeRateSubAccount struct {
		SubAcct string `json:"subAcct,omitempty"`
		EffDate string `json:"effDate,omitempty"`
	}
	CreateDepositAddress struct {
		Ccy   string        `json:"ccy,omitempty"`
		Addr  string        `json:"addr,omitempty"`
		Chain string        `json:"chain,omitempty"`
		PmtID string        `json:"pmtId,omitempty"`
		Tag   string        `json:"tag,omitempty"`
		Memo  string        `json:"memo,omitempty"`
		TS    okex.JSONTime `json:"ts,omitempty"`
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
)
