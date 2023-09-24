/*
eZmax API Definition (Full)

Testing ObjectPaymenttermAPIService

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

func Test_eZmaxApi_ObjectPaymenttermAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ObjectPaymenttermAPIService PaymenttermCreateObjectV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ObjectPaymenttermAPI.PaymenttermCreateObjectV1(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectPaymenttermAPIService PaymenttermEditObjectV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiPaymenttermID int32

		resp, httpRes, err := apiClient.ObjectPaymenttermAPI.PaymenttermEditObjectV1(context.Background(), pkiPaymenttermID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectPaymenttermAPIService PaymenttermGetAutocompleteV2", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sSelector string

		resp, httpRes, err := apiClient.ObjectPaymenttermAPI.PaymenttermGetAutocompleteV2(context.Background(), sSelector).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectPaymenttermAPIService PaymenttermGetListV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ObjectPaymenttermAPI.PaymenttermGetListV1(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectPaymenttermAPIService PaymenttermGetObjectV2", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiPaymenttermID int32

		resp, httpRes, err := apiClient.ObjectPaymenttermAPI.PaymenttermGetObjectV2(context.Background(), pkiPaymenttermID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}