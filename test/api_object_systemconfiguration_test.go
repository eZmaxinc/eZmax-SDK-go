/*
eZmax API Definition (Full)

Testing ObjectSystemconfigurationAPIService

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

func Test_eZmaxApi_ObjectSystemconfigurationAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ObjectSystemconfigurationAPIService SystemconfigurationEditObjectV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiSystemconfigurationID int32

		resp, httpRes, err := apiClient.ObjectSystemconfigurationAPI.SystemconfigurationEditObjectV1(context.Background(), pkiSystemconfigurationID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectSystemconfigurationAPIService SystemconfigurationGetObjectV2", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiSystemconfigurationID int32

		resp, httpRes, err := apiClient.ObjectSystemconfigurationAPI.SystemconfigurationGetObjectV2(context.Background(), pkiSystemconfigurationID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}