package command

import (
	m "github.com/LerianStudio/midaz/components/ledger/internal/domain/metadata"
	l "github.com/LerianStudio/midaz/components/ledger/internal/domain/onboarding/ledger"
	o "github.com/LerianStudio/midaz/components/ledger/internal/domain/onboarding/organization"
	a "github.com/LerianStudio/midaz/components/ledger/internal/domain/portfolio/account"
	i "github.com/LerianStudio/midaz/components/ledger/internal/domain/portfolio/instrument"
	p "github.com/LerianStudio/midaz/components/ledger/internal/domain/portfolio/portfolio"
	r "github.com/LerianStudio/midaz/components/ledger/internal/domain/portfolio/product"
)

// UseCase is a struct that aggregates various repositories for simplified access in use case implementations.
type UseCase struct {
	// OrganizationRepository provides an abstraction on top of the organization data source.
	OrganizationRepository o.Repository

	// LedgerRepository provides an abstraction on top of the ledger data source.
	LedgerRepository l.Repository

	// ProductRepository provides an abstraction on top of the product data source.
	ProductRepository r.Repository

	// PortfolioRepository provides an abstraction on top of the portfolio data source.
	PortfolioRepository p.Repository

	// AccountRepository provides an abstraction on top of the account data source.
	AccountRepository a.Repository

	// InstrumentRepository provides an abstraction on top of the instrument data source.
	InstrumentRepository i.Repository

	// MetadataRepository provides an abstraction on top of the metadata data source.
	MetadataRepository m.Repository
}
