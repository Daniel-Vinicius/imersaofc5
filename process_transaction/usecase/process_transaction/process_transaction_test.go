package process_transaction

import (
	"testing"
	"time"

	mock_broker "github.com/Daniel-Vinicus/imersaofc5/adapter/broker/mock"
	"github.com/Daniel-Vinicus/imersaofc5/domain/entity"
	mock_repository "github.com/Daniel-Vinicus/imersaofc5/domain/repository/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestProcessTransaction_ExecuteInvalidCreditCard(t *testing.T) {
	nextYear := time.Now().AddDate(1, 0, 0).Year()

	input := TransactionDtoInput{
		ID:                        "1",
		AccountID:                 "1",
		CreditCardNumber:          "0000000000000000",
		CreditCardName:            "Daniel Vinícius",
		CreditCardExpirationMonth: 12,
		CreditCardExpirationYear:  nextYear,
		CreditCardExpirationCVV:   123,
		Amount:                    200,
	}

	expectedOutput := TransactionDtoOutput{
		ID:           "1",
		Status:       entity.REJECTED,
		ErrorMessage: "invalid credit card number",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock_repository.NewMockTransactionRepository(ctrl)
	repositoryMock.EXPECT().
		Insert(input.ID, input.AccountID, input.Amount, expectedOutput.Status, expectedOutput.ErrorMessage).
		Return(nil)

	producerMock := mock_broker.NewMockProducerInterface(ctrl)
	producerMock.EXPECT().
		Publish(expectedOutput, []byte(expectedOutput.ID), "transactions_result").
		Return(nil)

	usecase := NewProcessTransaction(repositoryMock, producerMock, "transactions_result")
	output, err := usecase.Execute(input)

	// Expected that variable err be nil
	assert.Nil(t, err)

	// Expected that variable output be equal to expectedOutput
	assert.Equal(t, expectedOutput, output)
}

func TestProcessTransaction_ExecuteRejectedTransaction(t *testing.T) {
	nextYear := time.Now().AddDate(1, 0, 0).Year()

	input := TransactionDtoInput{
		ID:                        "1",
		AccountID:                 "1",
		CreditCardNumber:          "5288050701449222",
		CreditCardName:            "Daniel Vinícius",
		CreditCardExpirationMonth: 12,
		CreditCardExpirationYear:  nextYear,
		CreditCardExpirationCVV:   123,
		Amount:                    2000,
	}

	expectedOutput := TransactionDtoOutput{
		ID:           "1",
		Status:       entity.REJECTED,
		ErrorMessage: "you dont have limit for this transaction",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock_repository.NewMockTransactionRepository(ctrl)
	repositoryMock.EXPECT().
		Insert(input.ID, input.AccountID, input.Amount, expectedOutput.Status, expectedOutput.ErrorMessage).
		Return(nil)

	producerMock := mock_broker.NewMockProducerInterface(ctrl)
	producerMock.EXPECT().
		Publish(expectedOutput, []byte(expectedOutput.ID), "transactions_result").
		Return(nil)

	usecase := NewProcessTransaction(repositoryMock, producerMock, "transactions_result")
	output, err := usecase.Execute(input)

	// Expected that variable err be nil
	assert.Nil(t, err)

	// Expected that variable output be equal to expectedOutput
	assert.Equal(t, expectedOutput, output)
}

func TestProcessTransaction_ExecuteApprovedTransaction(t *testing.T) {
	nextYear := time.Now().AddDate(1, 0, 0).Year()

	input := TransactionDtoInput{
		ID:                        "1",
		AccountID:                 "1",
		CreditCardNumber:          "5288050701449222",
		CreditCardName:            "Daniel Vinícius",
		CreditCardExpirationMonth: 12,
		CreditCardExpirationYear:  nextYear,
		CreditCardExpirationCVV:   123,
		Amount:                    800,
	}

	expectedOutput := TransactionDtoOutput{
		ID:           "1",
		Status:       entity.APPROVED,
		ErrorMessage: "",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock_repository.NewMockTransactionRepository(ctrl)
	repositoryMock.EXPECT().
		Insert(input.ID, input.AccountID, input.Amount, expectedOutput.Status, expectedOutput.ErrorMessage).
		Return(nil)

	producerMock := mock_broker.NewMockProducerInterface(ctrl)
	producerMock.EXPECT().
		Publish(expectedOutput, []byte(expectedOutput.ID), "transactions_result").
		Return(nil)

	usecase := NewProcessTransaction(repositoryMock, producerMock, "transactions_result")
	output, err := usecase.Execute(input)

	// Expected that variable err be nil
	assert.Nil(t, err)

	// Expected that variable output be equal to expectedOutput
	assert.Equal(t, expectedOutput, output)
}
