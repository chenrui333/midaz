package query

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/LerianStudio/midaz/common"
	"github.com/LerianStudio/midaz/common/mmodel"
	"github.com/LerianStudio/midaz/common/mpointers"
	"github.com/LerianStudio/midaz/components/ledger/internal/adapters/postgres/account"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

// TestGetAccountByIDWithDeletedSuccess is responsible to test GetAccountByIDWithDeleted with success
func TestGetAccountByIDWithDeletedSuccess(t *testing.T) {
	organizationID := common.GenerateUUIDv7()
	ledgerID := common.GenerateUUIDv7()
	portfolioID := common.GenerateUUIDv7()
	id := common.GenerateUUIDv7()

	a := &mmodel.Account{
		ID:             id.String(),
		OrganizationID: organizationID.String(),
		LedgerID:       ledgerID.String(),
		PortfolioID:    mpointers.String(portfolioID.String()),
		DeletedAt:      mpointers.Time(time.Now()),
	}

	uc := UseCase{
		AccountRepo: account.NewMockRepository(gomock.NewController(t)),
	}

	uc.AccountRepo.(*account.MockRepository).
		EXPECT().
		Find(gomock.Any(), organizationID, ledgerID, &portfolioID, id).
		Return(a, nil).
		Times(1)
	res, err := uc.AccountRepo.Find(context.TODO(), organizationID, ledgerID, &portfolioID, id)

	assert.Equal(t, res, a)
	assert.Nil(t, err)
}

// TestGetAccountByIDWithDeletedWithoutPortfolioSuccess is responsible to test GetAccountByIDWithDeleted without portfolio with success
func TestGetAccountByIDWithDeletedWithoutPortfolioSuccess(t *testing.T) {
	organizationID := common.GenerateUUIDv7()
	ledgerID := common.GenerateUUIDv7()
	id := common.GenerateUUIDv7()

	a := &mmodel.Account{
		ID:             id.String(),
		OrganizationID: organizationID.String(),
		LedgerID:       ledgerID.String(),
		PortfolioID:    nil,
		DeletedAt:      mpointers.Time(time.Now()),
	}

	uc := UseCase{
		AccountRepo: account.NewMockRepository(gomock.NewController(t)),
	}

	uc.AccountRepo.(*account.MockRepository).
		EXPECT().
		Find(gomock.Any(), organizationID, ledgerID, nil, id).
		Return(a, nil).
		Times(1)
	res, err := uc.AccountRepo.Find(context.TODO(), organizationID, ledgerID, nil, id)

	assert.Equal(t, res, a)
	assert.Nil(t, err)
}

// TestGetAccountByIDWithDeletedError is responsible to test GetAccountByIDWithDeleted with error
func TestGetAccountByIDWithDeletedError(t *testing.T) {
	organizationID := common.GenerateUUIDv7()
	ledgerID := common.GenerateUUIDv7()
	portfolioID := common.GenerateUUIDv7()
	id := common.GenerateUUIDv7()

	errMSG := "errDatabaseItemNotFound"

	uc := UseCase{
		AccountRepo: account.NewMockRepository(gomock.NewController(t)),
	}

	uc.AccountRepo.(*account.MockRepository).
		EXPECT().
		Find(gomock.Any(), organizationID, ledgerID, &portfolioID, id).
		Return(nil, errors.New(errMSG)).
		Times(1)
	res, err := uc.AccountRepo.Find(context.TODO(), organizationID, ledgerID, &portfolioID, id)

	assert.NotEmpty(t, err)
	assert.Equal(t, err.Error(), errMSG)
	assert.Nil(t, res)
}
