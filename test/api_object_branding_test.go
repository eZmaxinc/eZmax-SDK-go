/*
eZmax API Definition (Full)

Testing ObjectBrandingAPIService

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

func Test_eZmaxApi_ObjectBrandingAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ObjectBrandingAPIService BrandingCreateObjectV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ObjectBrandingAPI.BrandingCreateObjectV1(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectBrandingAPIService BrandingEditObjectV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiBrandingID int32

		resp, httpRes, err := apiClient.ObjectBrandingAPI.BrandingEditObjectV1(context.Background(), pkiBrandingID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectBrandingAPIService BrandingGetAutocompleteV2", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sSelector string

		resp, httpRes, err := apiClient.ObjectBrandingAPI.BrandingGetAutocompleteV2(context.Background(), sSelector).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectBrandingAPIService BrandingGetListV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ObjectBrandingAPI.BrandingGetListV1(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectBrandingAPIService BrandingGetObjectV2", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiBrandingID int32

		resp, httpRes, err := apiClient.ObjectBrandingAPI.BrandingGetObjectV2(context.Background(), pkiBrandingID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
