/*
eZmax API Definition (Full)

Testing ObjectPermissionAPIService

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

func Test_eZmaxApi_ObjectPermissionAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ObjectPermissionAPIService PermissionCreateObjectV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ObjectPermissionAPI.PermissionCreateObjectV1(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectPermissionAPIService PermissionDeleteObjectV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiPermissionID int32

		resp, httpRes, err := apiClient.ObjectPermissionAPI.PermissionDeleteObjectV1(context.Background(), pkiPermissionID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectPermissionAPIService PermissionEditObjectV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiPermissionID int32

		resp, httpRes, err := apiClient.ObjectPermissionAPI.PermissionEditObjectV1(context.Background(), pkiPermissionID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectPermissionAPIService PermissionGetObjectV2", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiPermissionID int32

		resp, httpRes, err := apiClient.ObjectPermissionAPI.PermissionGetObjectV2(context.Background(), pkiPermissionID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
