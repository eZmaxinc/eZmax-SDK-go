/*
eZmax API Definition (Full)

Testing ScimGroupsAPIService

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

func Test_eZmaxApi_ScimGroupsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ScimGroupsAPIService GroupsCreateObjectScimV2", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ScimGroupsAPI.GroupsCreateObjectScimV2(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ScimGroupsAPIService GroupsDeleteObjectScimV2", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var groupId string

		httpRes, err := apiClient.ScimGroupsAPI.GroupsDeleteObjectScimV2(context.Background(), groupId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ScimGroupsAPIService GroupsEditObjectScimV2", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var groupId string

		resp, httpRes, err := apiClient.ScimGroupsAPI.GroupsEditObjectScimV2(context.Background(), groupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ScimGroupsAPIService GroupsGetListScimV2", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ScimGroupsAPI.GroupsGetListScimV2(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ScimGroupsAPIService GroupsGetObjectScimV2", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var groupId string

		resp, httpRes, err := apiClient.ScimGroupsAPI.GroupsGetObjectScimV2(context.Background(), groupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}