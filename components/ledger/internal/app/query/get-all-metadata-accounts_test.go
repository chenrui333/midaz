package query

import (
	"context"
	"errors"
	"reflect"
	"testing"

	commonHTTP "github.com/LerianStudio/midaz/common/net/http"
	meta "github.com/LerianStudio/midaz/components/ledger/internal/domain/metadata"
	a "github.com/LerianStudio/midaz/components/ledger/internal/domain/portfolio/account"
	mock "github.com/LerianStudio/midaz/components/ledger/internal/gen/mock/metadata"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
)

// TestGetAllMetadataAccounts is responsible to test TestGetAllMetadataAccounts with success and error
func TestGetAllMetadataAccounts(t *testing.T) {
	collection := reflect.TypeOf(a.Account{}).Name()
	filter := commonHTTP.QueryHeader{
		Metadata: &bson.M{"metadata": 1},
		Limit:    10,
		Page:     1,
	}

	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockMetadataRepository := mock.NewMockRepository(gomock.NewController(t))
	uc := UseCase{
		MetadataRepository: mockMetadataRepository,
	}

	t.Run("Success", func(t *testing.T) {
		mockMetadataRepository.
			EXPECT().
			FindList(gomock.Any(), collection, filter).
			Return([]*meta.Metadata{{ID: primitive.NewObjectID()}}, nil).
			Times(1)
		res, err := uc.MetadataRepository.FindList(context.TODO(), collection, filter)

		assert.NoError(t, err)
		assert.Len(t, res, 1)
	})

	t.Run("Error", func(t *testing.T) {
		errMSG := "errDatabaseItemNotFound"
		mockMetadataRepository.
			EXPECT().
			FindList(gomock.Any(), collection, filter).
			Return(nil, errors.New(errMSG)).
			Times(1)
		res, err := uc.MetadataRepository.FindList(context.TODO(), collection, filter)

		assert.EqualError(t, err, errMSG)
		assert.Nil(t, res)
	})
}
