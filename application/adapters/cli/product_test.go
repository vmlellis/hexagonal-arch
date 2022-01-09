package cli_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/vmlellis/go-hexagonal/application/adapters/cli"
	mock_application "github.com/vmlellis/go-hexagonal/application/mocks"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productName := "Product Test"
	productPrice := 25.99
	productStatus := "enabled"
	productId := "abc"

	productMock := mock_application.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetId().Return(productId).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()

	service := mock_application.NewMockProductServiceInterface(ctrl)
	service.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	service.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	t.Run("create", func(t *testing.T) {
		resultExpected := fmt.Sprintf(
			"Product ID %s with the name %s has been created with the the price %f and status %s.",
			productId, productName, productPrice, productStatus,
		)

		result, err := cli.Run(service, "create", productId, productName, productPrice)
		require.Nil(t, err)
		require.Equal(t, resultExpected, result)
	})

	t.Run("enable", func(t *testing.T) {
		resultExpected := fmt.Sprintf("Product %s has been enabled.", productName)

		result, err := cli.Run(service, "enable", productId, productName, productPrice)
		require.Nil(t, err)
		require.Equal(t, resultExpected, result)
	})

	t.Run("disable", func(t *testing.T) {
		resultExpected := fmt.Sprintf("Product %s has been disabled.", productName)

		result, err := cli.Run(service, "disable", productId, productName, productPrice)
		require.Nil(t, err)
		require.Equal(t, resultExpected, result)
	})

	t.Run("get", func(t *testing.T) {
		resultExpected := fmt.Sprintf(
			"Product ID: %s\nName: %s\nPrice: %f\nStatus: %s",
			productId, productName, productPrice, productStatus,
		)

		result, err := cli.Run(service, "get", productId, productName, productPrice)
		require.Nil(t, err)
		require.Equal(t, resultExpected, result)
	})
}
