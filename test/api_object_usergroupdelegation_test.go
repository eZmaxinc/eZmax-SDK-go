/*
eZmax API Definition (Full)

Testing ObjectUsergroupdelegationAPIService

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

func Test_eZmaxApi_ObjectUsergroupdelegationAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ObjectUsergroupdelegationAPIService UsergroupdelegationCreateObjectV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ObjectUsergroupdelegationAPI.UsergroupdelegationCreateObjectV1(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectUsergroupdelegationAPIService UsergroupdelegationDeleteObjectV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiUsergroupdelegationID int32

		resp, httpRes, err := apiClient.ObjectUsergroupdelegationAPI.UsergroupdelegationDeleteObjectV1(context.Background(), pkiUsergroupdelegationID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectUsergroupdelegationAPIService UsergroupdelegationEditObjectV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiUsergroupdelegationID int32

		resp, httpRes, err := apiClient.ObjectUsergroupdelegationAPI.UsergroupdelegationEditObjectV1(context.Background(), pkiUsergroupdelegationID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectUsergroupdelegationAPIService UsergroupdelegationGetObjectV2", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiUsergroupdelegationID int32

		resp, httpRes, err := apiClient.ObjectUsergroupdelegationAPI.UsergroupdelegationGetObjectV2(context.Background(), pkiUsergroupdelegationID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
