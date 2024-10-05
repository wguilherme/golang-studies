package tax

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type TaxRepositoryMock struct {
	mock.Mock
}

// mock
func (m *TaxRepositoryMock) Save(tax float64) error {
	args := m.Called(tax)
	return args.Error(0)
}

func TestCalculateTax(t *testing.T) {
	tax, err := CalculateTax(1000.0)
	assert.Nil(t, err)
	assert.Equal(t, 10.0, tax)

	tax, err = CalculateTax(0)
	assert.Error(t, err, "amount must be greater than 0")
	assert.Equal(t, 0.0, tax)
	assert.Contains(t, err.Error(), "greater than 0")
}

func TestCalculateTaxAndSave(t *testing.T) {
	repository := &TaxRepositoryMock{}
	repository.On("Save", 10.0).Return(nil)
	// repository.On("SaveTax", 0).Return(errors.New("amount must be greater than 0"))

	err := CalculateTaxAndSave(1000.0, repository)

	assert.Nil(t, err)

	repository.AssertExpectations(t)
	repository.AssertNumberOfCalls(t, "Save", 1)
}
