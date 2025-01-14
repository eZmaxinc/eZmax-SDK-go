/*
eZmax API Definition (Full)

Testing ObjectWebhookAPIService

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

func Test_eZmaxApi_ObjectWebhookAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ObjectWebhookAPIService WebhookCreateObjectV2", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ObjectWebhookAPI.WebhookCreateObjectV2(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectWebhookAPIService WebhookDeleteObjectV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiWebhookID int32

		resp, httpRes, err := apiClient.ObjectWebhookAPI.WebhookDeleteObjectV1(context.Background(), pkiWebhookID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectWebhookAPIService WebhookEditObjectV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiWebhookID int32

		resp, httpRes, err := apiClient.ObjectWebhookAPI.WebhookEditObjectV1(context.Background(), pkiWebhookID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectWebhookAPIService WebhookGetHistoryV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiWebhookID int32

		resp, httpRes, err := apiClient.ObjectWebhookAPI.WebhookGetHistoryV1(context.Background(), pkiWebhookID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectWebhookAPIService WebhookGetListV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ObjectWebhookAPI.WebhookGetListV1(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectWebhookAPIService WebhookGetObjectV2", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiWebhookID int32

		resp, httpRes, err := apiClient.ObjectWebhookAPI.WebhookGetObjectV2(context.Background(), pkiWebhookID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectWebhookAPIService WebhookRegenerateApikeyV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiWebhookID int32

		resp, httpRes, err := apiClient.ObjectWebhookAPI.WebhookRegenerateApikeyV1(context.Background(), pkiWebhookID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectWebhookAPIService WebhookSendWebhookV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ObjectWebhookAPI.WebhookSendWebhookV1(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectWebhookAPIService WebhookTestV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiWebhookID int32

		resp, httpRes, err := apiClient.ObjectWebhookAPI.WebhookTestV1(context.Background(), pkiWebhookID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
