/*
eZmax API Definition (Full)

Testing ObjectCreditcardmerchantAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package eZmaxApi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/ezmaxinc/ezmax-sdk-go"
)

func Test_eZmaxApi_ObjectCreditcardmerchantAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ObjectCreditcardmerchantAPIService CreditcardmerchantCreateObjectV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ObjectCreditcardmerchantAPI.CreditcardmerchantCreateObjectV1(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectCreditcardmerchantAPIService CreditcardmerchantDeleteObjectV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiCreditcardmerchantID int32

		resp, httpRes, err := apiClient.ObjectCreditcardmerchantAPI.CreditcardmerchantDeleteObjectV1(context.Background(), pkiCreditcardmerchantID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectCreditcardmerchantAPIService CreditcardmerchantEditObjectV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiCreditcardmerchantID int32

		resp, httpRes, err := apiClient.ObjectCreditcardmerchantAPI.CreditcardmerchantEditObjectV1(context.Background(), pkiCreditcardmerchantID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectCreditcardmerchantAPIService CreditcardmerchantGetAutocompleteV2", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sSelector string

		resp, httpRes, err := apiClient.ObjectCreditcardmerchantAPI.CreditcardmerchantGetAutocompleteV2(context.Background(), sSelector).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectCreditcardmerchantAPIService CreditcardmerchantGetListV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ObjectCreditcardmerchantAPI.CreditcardmerchantGetListV1(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectCreditcardmerchantAPIService CreditcardmerchantGetObjectV2", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiCreditcardmerchantID int32

		resp, httpRes, err := apiClient.ObjectCreditcardmerchantAPI.CreditcardmerchantGetObjectV2(context.Background(), pkiCreditcardmerchantID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
