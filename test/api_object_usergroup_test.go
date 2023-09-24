/*
eZmax API Definition (Full)

Testing ObjectUsergroupAPIService

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

func Test_eZmaxApi_ObjectUsergroupAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ObjectUsergroupAPIService UsergroupCreateObjectV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ObjectUsergroupAPI.UsergroupCreateObjectV1(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectUsergroupAPIService UsergroupEditObjectV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiUsergroupID int32

		resp, httpRes, err := apiClient.ObjectUsergroupAPI.UsergroupEditObjectV1(context.Background(), pkiUsergroupID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectUsergroupAPIService UsergroupEditPermissionsV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiUsergroupID int32

		resp, httpRes, err := apiClient.ObjectUsergroupAPI.UsergroupEditPermissionsV1(context.Background(), pkiUsergroupID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectUsergroupAPIService UsergroupEditUsergroupdelegationsV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiUsergroupID int32

		resp, httpRes, err := apiClient.ObjectUsergroupAPI.UsergroupEditUsergroupdelegationsV1(context.Background(), pkiUsergroupID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectUsergroupAPIService UsergroupEditUsergroupmembershipsV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiUsergroupID int32

		resp, httpRes, err := apiClient.ObjectUsergroupAPI.UsergroupEditUsergroupmembershipsV1(context.Background(), pkiUsergroupID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectUsergroupAPIService UsergroupGetAutocompleteV2", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sSelector string

		resp, httpRes, err := apiClient.ObjectUsergroupAPI.UsergroupGetAutocompleteV2(context.Background(), sSelector).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectUsergroupAPIService UsergroupGetListV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ObjectUsergroupAPI.UsergroupGetListV1(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectUsergroupAPIService UsergroupGetObjectV2", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiUsergroupID int32

		resp, httpRes, err := apiClient.ObjectUsergroupAPI.UsergroupGetObjectV2(context.Background(), pkiUsergroupID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectUsergroupAPIService UsergroupGetPermissionsV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiUsergroupID int32

		resp, httpRes, err := apiClient.ObjectUsergroupAPI.UsergroupGetPermissionsV1(context.Background(), pkiUsergroupID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectUsergroupAPIService UsergroupGetUsergroupdelegationsV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiUsergroupID int32

		resp, httpRes, err := apiClient.ObjectUsergroupAPI.UsergroupGetUsergroupdelegationsV1(context.Background(), pkiUsergroupID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectUsergroupAPIService UsergroupGetUsergroupmembershipsV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiUsergroupID int32

		resp, httpRes, err := apiClient.ObjectUsergroupAPI.UsergroupGetUsergroupmembershipsV1(context.Background(), pkiUsergroupID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}