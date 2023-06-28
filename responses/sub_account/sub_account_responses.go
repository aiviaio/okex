package sub_account

import (
	"github.com/aiviaio/okex/models/account"
	models "github.com/aiviaio/okex/models/subaccount"
	"github.com/aiviaio/okex/responses"
)

type (
	ViewList struct {
		responses.Basic
		SubAccounts []*models.SubAccount `json:"data,omitempty"`
	}
	APIKey struct {
		responses.Basic
		APIKeys []*models.APIKey `json:"data,omitempty"`
	}
	GetBalance struct {
		responses.Basic
		Balances []*account.Balance `json:"data,omitempty"`
	}
	HistoryTransfer struct {
		responses.Basic
		HistoryTransfers []*models.HistoryTransfer `json:"data,omitempty"`
	}
	ManageTransfer struct {
		responses.Basic
		Transfers []*models.Transfer `json:"data,omitempty"`
	}
	CreateSubAccount struct {
		responses.Basic
		CreateSubAccount []*models.CreateSubAccount `json:"data,omitempty"`
	}
	DeleteSubAccount struct {
		responses.Basic
		DeleteSubAccount []*models.DeleteSubAccount `json:"data,omitempty"`
	}
	CreatAPIKeySubAccount struct {
		responses.Basic
		CreatAPIKeySubAccount []*models.CreatAPIKeySubAccount `json:"data,omitempty"`
	}
	UpdateAPIKEySubAccount struct {
		responses.Basic
		UpdateSubAccount []*models.UpdateAPIKEySubAccount `json:"data,omitempty"`
	}
	DeleteAPIKeySubAccount struct {
		responses.Basic
		DeleteAPIKeySubAccount []*models.DeleteAPIKeySubAccount `json:"data,omitempty"`
	}
	SetLevelSubAccount struct {
		responses.Basic
		SetLevelSubAccount []*models.SetLevelSubAccount `json:"data,omitempty"`
	}
	SetFeeRateSubAccount struct {
		responses.Basic
		SetFeeRateSubAccount []*models.SetFeeRateSubAccount `json:"data,omitempty"`
	}
	CreateDepositAddress struct {
		responses.Basic
		CreateDepositAddress []*models.CreateDepositAddress `json:"data,omitempty"`
	}
	UpdateDepositAddress struct {
		responses.Basic
		UpdateDepositAddress []*models.UpdateDepositAddress `json:"data,omitempty"`
	}
	GetDepositAddress struct {
		responses.Basic
		GetDepositAddress []*models.GetDepositAddress `json:"data,omitempty"`
	}
)
