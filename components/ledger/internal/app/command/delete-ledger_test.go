package command

import (
	"context"
	"errors"
	"testing"

	mock "github.com/LerianStudio/midaz/components/ledger/internal/gen/mock/ledger"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

// TestDeleteLedgerByIDSuccess is responsible to test DeleteLedgerByID with success
func TestDeleteLedgerByIDSuccess(t *testing.T) {
	id := uuid.New()
	organizationID := uuid.New()

	uc := UseCase{
		LedgerRepository: mock.NewMockRepository(gomock.NewController(t)),
	}

	uc.LedgerRepository.(*mock.MockRepository).
		EXPECT().
		Delete(gomock.Any(), organizationID, id).
		Return(nil).
		Times(1)
	err := uc.LedgerRepository.Delete(context.TODO(), organizationID, id)

	assert.Nil(t, err)
}

// TestDeleteLedgerByIDError is responsible to test DeleteLedgerByID with error
func TestDeleteLedgerByIDError(t *testing.T) {
	id := uuid.New()
	organizationID := uuid.New()
	errMSG := "errDatabaseItemNotFound"

	uc := UseCase{
		LedgerRepository: mock.NewMockRepository(gomock.NewController(t)),
	}

	uc.LedgerRepository.(*mock.MockRepository).
		EXPECT().
		Delete(gomock.Any(), organizationID, id).
		Return(errors.New(errMSG)).
		Times(1)
	err := uc.LedgerRepository.Delete(context.TODO(), organizationID, id)

	assert.NotEmpty(t, err)
	assert.Equal(t, err.Error(), errMSG)
}
