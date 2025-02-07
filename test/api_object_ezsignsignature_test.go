/*
eZmax API Definition (Full)

Testing ObjectEzsignsignatureAPIService

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

func Test_eZmaxApi_ObjectEzsignsignatureAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ObjectEzsignsignatureAPIService EzsignsignatureCreateObjectV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ObjectEzsignsignatureAPI.EzsignsignatureCreateObjectV1(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectEzsignsignatureAPIService EzsignsignatureCreateObjectV2", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ObjectEzsignsignatureAPI.EzsignsignatureCreateObjectV2(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectEzsignsignatureAPIService EzsignsignatureCreateObjectV3", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ObjectEzsignsignatureAPI.EzsignsignatureCreateObjectV3(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectEzsignsignatureAPIService EzsignsignatureDeleteObjectV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiEzsignsignatureID int32

		resp, httpRes, err := apiClient.ObjectEzsignsignatureAPI.EzsignsignatureDeleteObjectV1(context.Background(), pkiEzsignsignatureID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectEzsignsignatureAPIService EzsignsignatureEditObjectV2", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiEzsignsignatureID int32

		resp, httpRes, err := apiClient.ObjectEzsignsignatureAPI.EzsignsignatureEditObjectV2(context.Background(), pkiEzsignsignatureID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectEzsignsignatureAPIService EzsignsignatureGetEzsignsignatureattachmentV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiEzsignsignatureID int32

		resp, httpRes, err := apiClient.ObjectEzsignsignatureAPI.EzsignsignatureGetEzsignsignatureattachmentV1(context.Background(), pkiEzsignsignatureID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectEzsignsignatureAPIService EzsignsignatureGetEzsignsignaturesAutomaticV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ObjectEzsignsignatureAPI.EzsignsignatureGetEzsignsignaturesAutomaticV1(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectEzsignsignatureAPIService EzsignsignatureGetObjectV3", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiEzsignsignatureID int32

		resp, httpRes, err := apiClient.ObjectEzsignsignatureAPI.EzsignsignatureGetObjectV3(context.Background(), pkiEzsignsignatureID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectEzsignsignatureAPIService EzsignsignatureSignV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiEzsignsignatureID int32

		resp, httpRes, err := apiClient.ObjectEzsignsignatureAPI.EzsignsignatureSignV1(context.Background(), pkiEzsignsignatureID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
