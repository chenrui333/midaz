package portfolio

import (
	"bytes"
	"testing"

	"github.com/LerianStudio/midaz/common/mmodel"
	"github.com/LerianStudio/midaz/components/mdz/internal/domain/repository"
	"github.com/LerianStudio/midaz/components/mdz/pkg/factory"
	"github.com/LerianStudio/midaz/components/mdz/pkg/iostreams"
	"github.com/LerianStudio/midaz/components/mdz/pkg/ptr"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_newCmdPortfolioCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockPortfolio(ctrl)

	portfolioID := "01931c99-adef-7b98-ad68-72d7e263066a"
	ledgerID := "0192e251-328d-7390-99f5-5c54980115ed"
	organizationID := "0192e250-ed9d-7e5c-a614-9b294151b572"
	entityID := "9b294151b57so"

	name := "Romaguera and Sons"
	code := "ACTIVE"
	description := "Teste Ledger"

	metadata := map[string]any{
		"chave1": "valor1",
		"chave2": 2,
		"chave3": true,
	}

	orgFactory := factoryPortfolioCreate{
		factory: &factory.Factory{IOStreams: &iostreams.IOStreams{
			Out: &bytes.Buffer{},
			Err: &bytes.Buffer{},
		}},
		repoPortfolio: mockRepo,
		tuiInput: func(message string) (string, error) {
			return name, nil
		},
		flagsCreate: flagsCreate{
			OrganizationID: organizationID,
			LedgerID:       ledgerID,
			EntityID:       entityID,
			Name:           name,
			Code:           code,
			Description:    description,
			Metadata:       "{\"chave1\": \"valor1\", \"chave2\": 2, \"chave3\": true}",
		},
	}

	cmd := newCmdPortfolioCreate(&orgFactory)
	cmd.SetArgs([]string{
		"--organization-id", organizationID,
		"--ledger-id", ledgerID,
		"--entity-id", entityID,
		"--name", name,
		"--status-code", code,
		"--status-description", description,
		"--metadata", "{\"chave1\": \"valor1\", \"chave2\": 2, \"chave3\": true}",
	})

	result := &mmodel.Portfolio{
		ID:             portfolioID,
		LedgerID:       ledgerID,
		OrganizationID: organizationID,
		Name:           name,
		Status: mmodel.Status{
			Code:        code,
			Description: ptr.StringPtr(description),
		},
		Metadata: metadata,
	}

	mockRepo.EXPECT().Create(gomock.Any(), gomock.Any(), gomock.Any()).Return(result, nil)
	err := cmd.Execute()
	assert.NoError(t, err)

	output := orgFactory.factory.IOStreams.Out.(*bytes.Buffer).String()
	assert.Contains(t, output, "The Portfolio ID 01931c99-adef-7b98-ad68-72d7e263066a has been successfully created")

}