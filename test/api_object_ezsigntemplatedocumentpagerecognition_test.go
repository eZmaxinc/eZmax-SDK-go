/*
eZmax API Definition (Full)

Testing ObjectEzsigntemplatedocumentpagerecognitionAPIService

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

func Test_eZmaxApi_ObjectEzsigntemplatedocumentpagerecognitionAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ObjectEzsigntemplatedocumentpagerecognitionAPIService EzsigntemplatedocumentpagerecognitionCreateObjectV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ObjectEzsigntemplatedocumentpagerecognitionAPI.EzsigntemplatedocumentpagerecognitionCreateObjectV1(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectEzsigntemplatedocumentpagerecognitionAPIService EzsigntemplatedocumentpagerecognitionDeleteObjectV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiEzsigntemplatedocumentpagerecognitionID int32

		resp, httpRes, err := apiClient.ObjectEzsigntemplatedocumentpagerecognitionAPI.EzsigntemplatedocumentpagerecognitionDeleteObjectV1(context.Background(), pkiEzsigntemplatedocumentpagerecognitionID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectEzsigntemplatedocumentpagerecognitionAPIService EzsigntemplatedocumentpagerecognitionEditObjectV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiEzsigntemplatedocumentpagerecognitionID int32

		resp, httpRes, err := apiClient.ObjectEzsigntemplatedocumentpagerecognitionAPI.EzsigntemplatedocumentpagerecognitionEditObjectV1(context.Background(), pkiEzsigntemplatedocumentpagerecognitionID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectEzsigntemplatedocumentpagerecognitionAPIService EzsigntemplatedocumentpagerecognitionGetObjectV2", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiEzsigntemplatedocumentpagerecognitionID int32

		resp, httpRes, err := apiClient.ObjectEzsigntemplatedocumentpagerecognitionAPI.EzsigntemplatedocumentpagerecognitionGetObjectV2(context.Background(), pkiEzsigntemplatedocumentpagerecognitionID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
