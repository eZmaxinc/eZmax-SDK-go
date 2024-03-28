/*
eZmax API Definition (Full)

Testing ObjectDiscussionmessageAPIService

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

func Test_eZmaxApi_ObjectDiscussionmessageAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ObjectDiscussionmessageAPIService DiscussionmessageCreateObjectV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ObjectDiscussionmessageAPI.DiscussionmessageCreateObjectV1(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectDiscussionmessageAPIService DiscussionmessageDeleteObjectV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiDiscussionmessageID int32

		resp, httpRes, err := apiClient.ObjectDiscussionmessageAPI.DiscussionmessageDeleteObjectV1(context.Background(), pkiDiscussionmessageID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectDiscussionmessageAPIService DiscussionmessagePatchObjectV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiDiscussionmessageID int32

		resp, httpRes, err := apiClient.ObjectDiscussionmessageAPI.DiscussionmessagePatchObjectV1(context.Background(), pkiDiscussionmessageID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}