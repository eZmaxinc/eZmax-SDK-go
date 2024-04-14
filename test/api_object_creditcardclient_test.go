/*
eZmax API Definition (Full)

Testing ObjectCreditcardclientAPIService

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

func Test_eZmaxApi_ObjectCreditcardclientAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ObjectCreditcardclientAPIService CreditcardclientCreateObjectV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ObjectCreditcardclientAPI.CreditcardclientCreateObjectV1(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectCreditcardclientAPIService CreditcardclientDeleteObjectV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiCreditcardclientID int32

		resp, httpRes, err := apiClient.ObjectCreditcardclientAPI.CreditcardclientDeleteObjectV1(context.Background(), pkiCreditcardclientID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectCreditcardclientAPIService CreditcardclientEditObjectV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiCreditcardclientID int32

		resp, httpRes, err := apiClient.ObjectCreditcardclientAPI.CreditcardclientEditObjectV1(context.Background(), pkiCreditcardclientID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectCreditcardclientAPIService CreditcardclientGetAutocompleteV2", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sSelector string

		resp, httpRes, err := apiClient.ObjectCreditcardclientAPI.CreditcardclientGetAutocompleteV2(context.Background(), sSelector).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectCreditcardclientAPIService CreditcardclientGetListV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ObjectCreditcardclientAPI.CreditcardclientGetListV1(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectCreditcardclientAPIService CreditcardclientGetObjectV2", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiCreditcardclientID int32

		resp, httpRes, err := apiClient.ObjectCreditcardclientAPI.CreditcardclientGetObjectV2(context.Background(), pkiCreditcardclientID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}