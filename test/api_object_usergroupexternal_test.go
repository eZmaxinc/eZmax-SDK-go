/*
eZmax API Definition (Full)

Testing ObjectUsergroupexternalAPIService

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

func Test_eZmaxApi_ObjectUsergroupexternalAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ObjectUsergroupexternalAPIService UsergroupexternalCreateObjectV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ObjectUsergroupexternalAPI.UsergroupexternalCreateObjectV1(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectUsergroupexternalAPIService UsergroupexternalDeleteObjectV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiUsergroupexternalID int32

		resp, httpRes, err := apiClient.ObjectUsergroupexternalAPI.UsergroupexternalDeleteObjectV1(context.Background(), pkiUsergroupexternalID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectUsergroupexternalAPIService UsergroupexternalEditObjectV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiUsergroupexternalID int32

		resp, httpRes, err := apiClient.ObjectUsergroupexternalAPI.UsergroupexternalEditObjectV1(context.Background(), pkiUsergroupexternalID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectUsergroupexternalAPIService UsergroupexternalGetAutocompleteV2", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sSelector string

		resp, httpRes, err := apiClient.ObjectUsergroupexternalAPI.UsergroupexternalGetAutocompleteV2(context.Background(), sSelector).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectUsergroupexternalAPIService UsergroupexternalGetListV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ObjectUsergroupexternalAPI.UsergroupexternalGetListV1(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectUsergroupexternalAPIService UsergroupexternalGetObjectV2", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiUsergroupexternalID int32

		resp, httpRes, err := apiClient.ObjectUsergroupexternalAPI.UsergroupexternalGetObjectV2(context.Background(), pkiUsergroupexternalID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectUsergroupexternalAPIService UsergroupexternalGetUsergroupexternalmembershipsV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiUsergroupexternalID int32

		resp, httpRes, err := apiClient.ObjectUsergroupexternalAPI.UsergroupexternalGetUsergroupexternalmembershipsV1(context.Background(), pkiUsergroupexternalID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectUsergroupexternalAPIService UsergroupexternalGetUsergroupsV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiUsergroupexternalID int32

		resp, httpRes, err := apiClient.ObjectUsergroupexternalAPI.UsergroupexternalGetUsergroupsV1(context.Background(), pkiUsergroupexternalID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
