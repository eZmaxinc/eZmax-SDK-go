/*
eZmax API Definition (Full)

Testing ObjectApikeyAPIService

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

func Test_eZmaxApi_ObjectApikeyAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ObjectApikeyAPIService ApikeyCreateObjectV2", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ObjectApikeyAPI.ApikeyCreateObjectV2(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectApikeyAPIService ApikeyEditObjectV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiApikeyID int32

		resp, httpRes, err := apiClient.ObjectApikeyAPI.ApikeyEditObjectV1(context.Background(), pkiApikeyID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectApikeyAPIService ApikeyEditPermissionsV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiApikeyID int32

		resp, httpRes, err := apiClient.ObjectApikeyAPI.ApikeyEditPermissionsV1(context.Background(), pkiApikeyID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectApikeyAPIService ApikeyGetCorsV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiApikeyID int32

		resp, httpRes, err := apiClient.ObjectApikeyAPI.ApikeyGetCorsV1(context.Background(), pkiApikeyID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectApikeyAPIService ApikeyGetListV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ObjectApikeyAPI.ApikeyGetListV1(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectApikeyAPIService ApikeyGetObjectV2", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiApikeyID int32

		resp, httpRes, err := apiClient.ObjectApikeyAPI.ApikeyGetObjectV2(context.Background(), pkiApikeyID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectApikeyAPIService ApikeyGetPermissionsV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiApikeyID int32

		resp, httpRes, err := apiClient.ObjectApikeyAPI.ApikeyGetPermissionsV1(context.Background(), pkiApikeyID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectApikeyAPIService ApikeyGetSubnetsV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiApikeyID int32

		resp, httpRes, err := apiClient.ObjectApikeyAPI.ApikeyGetSubnetsV1(context.Background(), pkiApikeyID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectApikeyAPIService ApikeyRegenerateV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiApikeyID int32

		resp, httpRes, err := apiClient.ObjectApikeyAPI.ApikeyRegenerateV1(context.Background(), pkiApikeyID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}