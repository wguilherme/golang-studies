package tax

import "errors"

func CalculateTax(amount float64) (float64, error) {
	if amount <= 0 {
		return 0, errors.New("amount must be greater than 0")
	}

	if amount >= 1000 && amount < 20000 {
		return 10.0, nil
	}
	if amount >= 20000 {
		return 20.0, nil
	}
	return 5.0, nil
}

type SaveTaxRepository interface {
	Save(amount float64) error
}

func CalculateTaxAndSave(amount float64, repository SaveTaxRepository) error {

	if amount <= 0 {
		return errors.New("amount must be greater than 0")
	}

	tax, err := CalculateTax(amount)

	if err != nil {
		return err
	}

	return repository.Save(tax)
}
