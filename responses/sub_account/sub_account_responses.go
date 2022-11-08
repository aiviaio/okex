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
)
