/*
eZmax API Definition (Full)

Testing ObjectEzsigndiscussionAPIService

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

func Test_eZmaxApi_ObjectEzsigndiscussionAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ObjectEzsigndiscussionAPIService EzsigndiscussionCreateObjectV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ObjectEzsigndiscussionAPI.EzsigndiscussionCreateObjectV1(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectEzsigndiscussionAPIService EzsigndiscussionDeleteObjectV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiEzsigndiscussionID int32

		resp, httpRes, err := apiClient.ObjectEzsigndiscussionAPI.EzsigndiscussionDeleteObjectV1(context.Background(), pkiEzsigndiscussionID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectEzsigndiscussionAPIService EzsigndiscussionGetObjectV2", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiEzsigndiscussionID int32

		resp, httpRes, err := apiClient.ObjectEzsigndiscussionAPI.EzsigndiscussionGetObjectV2(context.Background(), pkiEzsigndiscussionID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
