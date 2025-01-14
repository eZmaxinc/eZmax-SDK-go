/*
eZmax API Definition (Full)

Testing ObjectUserAPIService

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

func Test_eZmaxApi_ObjectUserAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ObjectUserAPIService UserCreateObjectV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ObjectUserAPI.UserCreateObjectV1(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectUserAPIService UserCreateObjectV2", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ObjectUserAPI.UserCreateObjectV2(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectUserAPIService UserEditColleaguesV2", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiUserID int32

		resp, httpRes, err := apiClient.ObjectUserAPI.UserEditColleaguesV2(context.Background(), pkiUserID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectUserAPIService UserEditObjectV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiUserID int32

		resp, httpRes, err := apiClient.ObjectUserAPI.UserEditObjectV1(context.Background(), pkiUserID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectUserAPIService UserEditPermissionsV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiUserID int32

		resp, httpRes, err := apiClient.ObjectUserAPI.UserEditPermissionsV1(context.Background(), pkiUserID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectUserAPIService UserGetApikeysV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiUserID int32

		resp, httpRes, err := apiClient.ObjectUserAPI.UserGetApikeysV1(context.Background(), pkiUserID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectUserAPIService UserGetAutocompleteV2", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sSelector string

		resp, httpRes, err := apiClient.ObjectUserAPI.UserGetAutocompleteV2(context.Background(), sSelector).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectUserAPIService UserGetColleaguesV2", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiUserID int32

		resp, httpRes, err := apiClient.ObjectUserAPI.UserGetColleaguesV2(context.Background(), pkiUserID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectUserAPIService UserGetEffectivePermissionsV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiUserID int32

		resp, httpRes, err := apiClient.ObjectUserAPI.UserGetEffectivePermissionsV1(context.Background(), pkiUserID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectUserAPIService UserGetListV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ObjectUserAPI.UserGetListV1(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectUserAPIService UserGetObjectV2", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiUserID int32

		resp, httpRes, err := apiClient.ObjectUserAPI.UserGetObjectV2(context.Background(), pkiUserID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectUserAPIService UserGetPermissionsV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiUserID int32

		resp, httpRes, err := apiClient.ObjectUserAPI.UserGetPermissionsV1(context.Background(), pkiUserID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectUserAPIService UserGetSubnetsV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiUserID int32

		resp, httpRes, err := apiClient.ObjectUserAPI.UserGetSubnetsV1(context.Background(), pkiUserID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectUserAPIService UserGetUsergroupexternalsV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiUserID int32

		resp, httpRes, err := apiClient.ObjectUserAPI.UserGetUsergroupexternalsV1(context.Background(), pkiUserID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectUserAPIService UserGetUsergroupsV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiUserID int32

		resp, httpRes, err := apiClient.ObjectUserAPI.UserGetUsergroupsV1(context.Background(), pkiUserID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectUserAPIService UserSendPasswordResetV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiUserID int32

		resp, httpRes, err := apiClient.ObjectUserAPI.UserSendPasswordResetV1(context.Background(), pkiUserID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
