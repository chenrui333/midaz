package model

import (
	"math"
	"math/big"
	"strings"

	"github.com/LerianStudio/midaz/pkg"
	"github.com/LerianStudio/midaz/pkg/constant"
	a "github.com/LerianStudio/midaz/pkg/mgrpc/account"
)

// GetListAccounts this func collects all the account identifiers from the Source.From and Distribute.To fields of the transaction.
func GetListAccounts(transaction Transaction) *Responses {
	flatTransaction := &Responses{
		Total:         transaction.Send.Value,
		Asset:         transaction.Send.Asset,
		Rates:         make(map[string]Rate),
		From:          make(map[string]Amount),
		To:            make(map[string]Amount),
		Sources:       make([]string, 0),
		Destinations:  make([]string, 0),
		AliasesAssets: make(map[string]string),
	}

	for i := range transaction.Send.Source.From {
		flatTransaction.Sources = append(flatTransaction.Sources, transaction.Send.Source.From[i].Account)
	}

	for i := range transaction.Distribute.To {
		flatTransaction.Destinations = append(flatTransaction.Destinations, transaction.Distribute.To[i].Account)
	}

	return flatTransaction
}

// ValidateAccounts function with some validates in accounts and DSL operations
func ValidateAccounts(transaction Transaction, accounts []*a.Account) (map[string]string, error) {
	aliasesWithAssets := map[string]string{}
	assets := map[string]string{}

	if len(accounts) != (len(transaction.Send.Source.From) + len(transaction.Distribute.To)) {
		return nil, pkg.ValidateBusinessError(constant.ErrAccountIneligibility, "ValidateAccounts")
	}

	for _, acc := range accounts {
		for i := range transaction.Send.Source.From {
			if acc.Id == transaction.Send.Source.From[i].Account || acc.Alias == transaction.Send.Source.From[i].Account {
				if !acc.AllowSending {
					return nil, pkg.ValidateBusinessError(constant.ErrAccountStatusTransactionRestriction, "ValidateAccounts")
				}

				if acc.AssetCode != transaction.Send.Asset {
					return nil, pkg.ValidateBusinessError(constant.ErrAccountStatusTransactionRestriction, "ValidateAccounts")
				}

				assets[acc.AssetCode] = acc.AssetCode

				aliasesWithAssets[acc.Alias] = acc.AssetCode
			}
		}

		for i := range transaction.Distribute.To {
			if acc.Id == transaction.Distribute.To[i].Account || acc.Alias == transaction.Distribute.To[i].Account && !acc.AllowReceiving {
				return nil, pkg.ValidateBusinessError(constant.ErrAccountStatusTransactionRestriction, "ValidateAccounts")
			}

			if acc.AssetCode != transaction.Send.Asset && acc.Type == constant.DefaultExternalAccountType {
				return nil, pkg.ValidateBusinessError(constant.ErrAccountStatusTransactionRestriction, "ValidateAccounts")
			}

			assets[acc.AssetCode] = acc.AssetCode

			aliasesWithAssets[acc.Alias] = acc.AssetCode
		}
	}

	if len(assets) > 2 {
		return nil, pkg.ValidateBusinessError(constant.ErrAccountIneligibility, "ValidateAccounts")
	}

	return aliasesWithAssets, nil
}

// ValidateFromToOperation func that validate operate balance amount
func ValidateFromToOperation(ft FromTo, validate Responses, acc *a.Account) (Amount, Balance, error) {
	amount := Amount{}

	balanceAfter := Balance{}

	if ft.IsFrom {
		ba, err := OperateAmounts(validate.From[ft.Account], acc.Balance, constant.DEBIT)
		if err != nil {
			return amount, balanceAfter, err
		}

		if ba.Available < 0 && acc.Type != constant.DefaultExternalAccountType {
			return amount, balanceAfter, pkg.ValidateBusinessError(constant.ErrInsufficientFunds, "ValidateFromToOperation", acc.Alias)
		}

		amount = Amount{
			Value: validate.From[ft.Account].Value,
			Scale: validate.From[ft.Account].Scale,
		}

		balanceAfter = ba
	} else {
		ba, err := OperateAmounts(validate.To[ft.Account], acc.Balance, constant.CREDIT)
		if err != nil {
			return amount, balanceAfter, err
		}

		amount = Amount{
			Value: validate.To[ft.Account].Value,
			Scale: validate.To[ft.Account].Scale,
		}

		balanceAfter = ba
	}

	return amount, balanceAfter, nil
}

// UpdateAccounts function with some updates values in accounts and
func UpdateAccounts(operation string, fromTo map[string]Amount, accounts []*a.Account, result chan []*a.Account, e chan error) {
	accs := make([]*a.Account, 0)

	for _, acc := range accounts {
		for key := range fromTo {
			if acc.Id == key || acc.Alias == key {
				b, err := OperateAmounts(fromTo[key], acc.Balance, operation)
				if err != nil {
					e <- err
				}

				balance := a.Balance{
					Available: float64(b.Available),
					Scale:     float64(b.Scale),
					OnHold:    float64(b.OnHold),
				}

				status := a.Status{
					Code:        acc.Status.Code,
					Description: acc.Status.Description,
				}

				ac := a.Account{
					Id:              acc.Id,
					Alias:           acc.Alias,
					Name:            acc.Name,
					ParentAccountId: acc.ParentAccountId,
					EntityId:        acc.EntityId,
					OrganizationId:  acc.OrganizationId,
					LedgerId:        acc.LedgerId,
					PortfolioId:     acc.PortfolioId,
					ProductId:       acc.ProductId,
					AssetCode:       acc.AssetCode,
					Balance:         &balance,
					Status:          &status,
					AllowSending:    acc.AllowSending,
					AllowReceiving:  acc.AllowReceiving,
					Type:            acc.Type,
					CreatedAt:       acc.CreatedAt,
					UpdatedAt:       acc.UpdatedAt,
				}

				accs = append(accs, &ac)

				break
			}
		}
	}

	result <- accs
}

// Scale func scale: (V * 10^ (S0-S1))
func Scale(v, s0, s1 int) float64 {
	return float64(v) * math.Pow10(s1-s0)
}

// UndoScale Function to undo the scale calculation
func UndoScale(v float64, s int) int {
	return int(v * math.Pow10(s))
}

// FindScale Function to find the scale for any value of a value
func FindScale(asset string, v float64, s int) Amount {
	valueString := big.NewFloat(v).String()
	parts := strings.Split(valueString, ".")

	scale := s
	value := int(v)

	if len(parts) > 1 {
		scale = len(parts[1])
		value = UndoScale(v, scale)

		if parts[1] != "0" {
			scale += s
		}
	}

	amount := Amount{
		Asset: asset,
		Value: value,
		Scale: scale,
	}

	return amount
}

// normalize func that normalize scale from all values
func normalize(total, amount, remaining *Amount) {
	if total.Scale < amount.Scale {
		if total.Value != 0 {
			v0 := Scale(total.Value, total.Scale, amount.Scale)

			total.Value = int(v0) + amount.Value
		} else {
			total.Value += amount.Value
		}

		total.Scale = amount.Scale
	} else {
		if total.Value != 0 {
			v0 := Scale(amount.Value, amount.Scale, total.Scale)

			total.Value += int(v0)

			amount.Value = int(v0)
			amount.Scale = total.Scale
		} else {
			total.Value += amount.Value
			total.Scale = amount.Scale
		}
	}

	if remaining.Scale < amount.Scale {
		v0 := Scale(remaining.Value, remaining.Scale, amount.Scale)

		remaining.Value = int(v0) - amount.Value
		remaining.Scale = amount.Scale
	} else {
		v0 := Scale(amount.Value, amount.Scale, remaining.Scale)

		remaining.Value -= int(v0)
	}
}

// OperateAmounts Function to sum or sub two amounts and normalize the scale
func OperateAmounts(amount Amount, balance *a.Balance, operation string) (Balance, error) {
	var scale float64

	var total float64

	switch operation {
	case constant.DEBIT:
		if int(balance.Scale) < amount.Scale {
			v0 := Scale(int(balance.Available), int(balance.Scale), amount.Scale)
			total = v0 - float64(amount.Value)
			scale = float64(amount.Scale)
		} else {
			v0 := Scale(amount.Value, amount.Scale, int(balance.Scale))
			total = balance.Available - v0
			scale = balance.Scale
		}
	default:
		if int(balance.Scale) < amount.Scale {
			v0 := Scale(int(balance.Available), int(balance.Scale), amount.Scale)
			total = v0 + float64(amount.Value)
			scale = float64(amount.Scale)
		} else {
			v0 := Scale(amount.Value, amount.Scale, int(balance.Scale))
			total = balance.Available + v0
			scale = balance.Scale
		}
	}

	blc := Balance{
		Available: int(total),
		OnHold:    int(balance.OnHold),
		Scale:     int(scale),
	}

	return blc, nil
}

// calculateTotal Calculate total for sources/destinations based on shares, amounts and remains
func calculateTotal(fromTos []FromTo, send Send, t chan int, ft chan map[string]Amount) {
	fmto := make(map[string]Amount)

	total := Amount{
		Asset: send.Asset,
		Scale: 0,
		Value: 0,
	}

	remaining := Amount{
		Asset: send.Asset,
		Scale: send.Scale,
		Value: send.Value,
	}

	for i := range fromTos {
		if fromTos[i].Share != nil && fromTos[i].Share.Percentage != 0 {
			shareValue := float64(send.Value) * (float64(fromTos[i].Share.Percentage) / float64(fromTos[i].Share.PercentageOfPercentage))
			amount := FindScale(send.Asset, shareValue, send.Scale)

			normalize(&total, &amount, &remaining)
			fmto[fromTos[i].Account] = amount
		}

		if fromTos[i].Amount != nil && fromTos[i].Amount.Value > 0 && fromTos[i].Amount.Scale > -1 {
			amount := Amount{
				Asset: fromTos[i].Amount.Asset,
				Scale: fromTos[i].Amount.Scale,
				Value: fromTos[i].Amount.Value,
			}

			normalize(&total, &amount, &remaining)
			fmto[fromTos[i].Account] = amount
		}

		if !pkg.IsNilOrEmpty(fromTos[i].Remaining) {
			total.Value += remaining.Value

			fmto[fromTos[i].Account] = remaining
			fromTos[i].Amount = &remaining
		}
	}

	ttl := total.Value
	if total.Scale > send.Scale {
		ttl = int(Scale(total.Value, total.Scale, send.Scale))
	}

	t <- ttl
	ft <- fmto
}

// ValidateSendSourceAndDistribute Validate send and distribute totals
func ValidateSendSourceAndDistribute(transaction Transaction, flatTransaction *Responses) (*Responses, error) {
	var sourcesTotal int
	total := make(chan int)
	fromTo := make(chan map[string]Amount)
	go calculateTotal(transaction.Send.Source.From, transaction.Send, total, fromTo)
	sourcesTotal = <-total
	flatTransaction.From = <-fromTo

	var destinationsTotal int
	go calculateTotal(transaction.Distribute.To, transaction.Send, total, fromTo)
	destinationsTotal = <-total
	flatTransaction.To = <-fromTo

	if math.Abs(float64(flatTransaction.Total)-float64(sourcesTotal)) != 0 {
		return nil, pkg.ValidateBusinessError(constant.ErrTransactionValueMismatch, "ValidateSendSourceAndDistribute")
	}

	if flatTransaction.Rates != nil {
		destinationsTotal = 0
		externalValue := 0
		for key, value := range flatTransaction.AliasesAssets {
			if !flatTransaction.Rates[value].IsEmpty() && !flatTransaction.To[key].IsEmpty() && flatTransaction.To[key].Asset != flatTransaction.Rates[value].To {
				to := flatTransaction.To[key]
				to.Asset = flatTransaction.Rates[value].To
				to.Value = (to.Value * flatTransaction.Rates[value].Value) / 100
				flatTransaction.To[key] = to
				externalValue = externalValue + to.Value
				destinationsTotal = destinationsTotal + (to.Value / (flatTransaction.Rates[value].Value / 100))
				flatTransaction.To[constant.DefaultExternalAccountAliasPrefix+to.Asset] = Amount{
					Asset: to.Asset,
					Scale: to.Scale,
					Value: externalValue,
				}
			} else if !flatTransaction.To[key].IsEmpty() {
				to := flatTransaction.To[key]
				if to.Asset == flatTransaction.Asset {
					sourcesTotal = sourcesTotal - to.Value
					flatTransaction.From[constant.DefaultExternalAccountAliasPrefix+flatTransaction.Asset] = Amount{
						Asset: flatTransaction.Asset,
						Scale: to.Scale,
						Value: sourcesTotal,
					}
				} else {
					destinationsTotal = destinationsTotal + (to.Value / (flatTransaction.Rates[value].Value / 100))
					externalValue = externalValue + to.Value
					flatTransaction.To[constant.DefaultExternalAccountAliasPrefix+to.Asset] = Amount{
						Asset: to.Asset,
						Scale: to.Scale,
						Value: externalValue,
					}
				}
			}

		}
	}

	if math.Abs(float64(sourcesTotal)-float64(destinationsTotal)) != 0 {
		return nil, pkg.ValidateBusinessError(constant.ErrTransactionValueMismatch, "ValidateSendSourceAndDistribute")
	}

	return flatTransaction, nil
}
