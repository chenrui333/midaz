package query

import (
	"context"
	"errors"
	"reflect"

	"github.com/LerianStudio/midaz/common"
	"github.com/LerianStudio/midaz/common/mlog"
	commonHTTP "github.com/LerianStudio/midaz/common/net/http"
	"github.com/LerianStudio/midaz/components/ledger/internal/app"
	r "github.com/LerianStudio/midaz/components/ledger/internal/domain/portfolio/product"
	"github.com/google/uuid"
)

// GetAllMetadataProducts fetch all Products from the repository
func (uc *UseCase) GetAllMetadataProducts(ctx context.Context, organizationID, ledgerID string, filter commonHTTP.QueryHeader) ([]*r.Product, error) {
	logger := mlog.NewLoggerFromContext(ctx)
	logger.Infof("Retrieving products")

	metadata, err := uc.MetadataRepo.FindList(ctx, reflect.TypeOf(r.Product{}).Name(), filter)
	if err != nil || metadata == nil {
		return nil, common.EntityNotFoundError{
			EntityType: reflect.TypeOf(r.Product{}).Name(),
			Code:       "0057",
			Title:      "No Products Found",
			Message:    "No products were found in the search. Please review the search criteria and try again.",
			Err:        err,
		}
	}

	uuids := make([]uuid.UUID, len(metadata))
	metadataMap := make(map[string]map[string]any, len(metadata))

	for i, meta := range metadata {
		uuids[i] = uuid.MustParse(meta.EntityID)
		metadataMap[meta.EntityID] = meta.Data
	}

	products, err := uc.ProductRepo.FindByIDs(ctx, uuid.MustParse(organizationID), uuid.MustParse(ledgerID), uuids)
	if err != nil {
		logger.Errorf("Error getting products on repo by query params: %v", err)

		if errors.Is(err, app.ErrDatabaseItemNotFound) {
			return nil, common.EntityNotFoundError{
				EntityType: reflect.TypeOf(r.Product{}).Name(),
				Code:       "0057",
				Title:      "No Products Found",
				Message:    "No products were found in the search. Please review the search criteria and try again.",
				Err:        err,
			}
		}

		return nil, err
	}

	for i := range products {
		if data, ok := metadataMap[products[i].ID]; ok {
			products[i].Metadata = data
		}
	}

	return products, nil
}
