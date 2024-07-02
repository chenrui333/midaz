package query

import (
	"context"
	"errors"
	"testing"

	p "github.com/LerianStudio/midaz/components/ledger/internal/domain/portfolio/portfolio"
	mock "github.com/LerianStudio/midaz/components/ledger/internal/gen/mock/portfolio"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

// TestGetAllPortfoliosError is responsible to test GetAllPortfolios with success and error
func TestGetAllPortfolios(t *testing.T) {
	organizationID := uuid.New()
	ledgerID := uuid.New()
	limit := 10
	page := 1

	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockPortfolioRepository := mock.NewMockRepository(ctrl)

	uc := UseCase{
		PortfolioRepository: mockPortfolioRepository,
	}

	t.Run("Success", func(t *testing.T) {
		portfolios := []*p.Portfolio{{}}
		mockPortfolioRepository.
			EXPECT().
			FindAll(gomock.Any(), organizationID, ledgerID, limit, page).
			Return(portfolios, nil).
			Times(1)
		res, err := uc.PortfolioRepository.FindAll(context.TODO(), organizationID, ledgerID, limit, page)

		assert.NoError(t, err)
		assert.Len(t, res, 1)
	})

	t.Run("Error", func(t *testing.T) {
		errMsg := "errDatabaseItemNotFound"
		mockPortfolioRepository.
			EXPECT().
			FindAll(gomock.Any(), organizationID, ledgerID, limit, page).
			Return(nil, errors.New(errMsg)).
			Times(1)
		res, err := uc.PortfolioRepository.FindAll(context.TODO(), organizationID, ledgerID, limit, page)

		assert.EqualError(t, err, errMsg)
		assert.Nil(t, res)
	})
}
