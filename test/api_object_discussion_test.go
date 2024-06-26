/*
eZmax API Definition (Full)

Testing ObjectDiscussionAPIService

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

func Test_eZmaxApi_ObjectDiscussionAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ObjectDiscussionAPIService DiscussionCreateObjectV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ObjectDiscussionAPI.DiscussionCreateObjectV1(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectDiscussionAPIService DiscussionDeleteObjectV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiDiscussionID int32

		resp, httpRes, err := apiClient.ObjectDiscussionAPI.DiscussionDeleteObjectV1(context.Background(), pkiDiscussionID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectDiscussionAPIService DiscussionGetObjectV2", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiDiscussionID int32

		resp, httpRes, err := apiClient.ObjectDiscussionAPI.DiscussionGetObjectV2(context.Background(), pkiDiscussionID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectDiscussionAPIService DiscussionPatchObjectV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiDiscussionID int32

		resp, httpRes, err := apiClient.ObjectDiscussionAPI.DiscussionPatchObjectV1(context.Background(), pkiDiscussionID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectDiscussionAPIService DiscussionUpdateDiscussionreadstatusV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiDiscussionID int32

		resp, httpRes, err := apiClient.ObjectDiscussionAPI.DiscussionUpdateDiscussionreadstatusV1(context.Background(), pkiDiscussionID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
