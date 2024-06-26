/*
eZmax API Definition (Full)

Testing ObjectEzsigntemplateformfieldgroupAPIService

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

func Test_eZmaxApi_ObjectEzsigntemplateformfieldgroupAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ObjectEzsigntemplateformfieldgroupAPIService EzsigntemplateformfieldgroupCreateObjectV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ObjectEzsigntemplateformfieldgroupAPI.EzsigntemplateformfieldgroupCreateObjectV1(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectEzsigntemplateformfieldgroupAPIService EzsigntemplateformfieldgroupDeleteObjectV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiEzsigntemplateformfieldgroupID int32

		resp, httpRes, err := apiClient.ObjectEzsigntemplateformfieldgroupAPI.EzsigntemplateformfieldgroupDeleteObjectV1(context.Background(), pkiEzsigntemplateformfieldgroupID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectEzsigntemplateformfieldgroupAPIService EzsigntemplateformfieldgroupEditObjectV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiEzsigntemplateformfieldgroupID int32

		resp, httpRes, err := apiClient.ObjectEzsigntemplateformfieldgroupAPI.EzsigntemplateformfieldgroupEditObjectV1(context.Background(), pkiEzsigntemplateformfieldgroupID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectEzsigntemplateformfieldgroupAPIService EzsigntemplateformfieldgroupGetObjectV2", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiEzsigntemplateformfieldgroupID int32

		resp, httpRes, err := apiClient.ObjectEzsigntemplateformfieldgroupAPI.EzsigntemplateformfieldgroupGetObjectV2(context.Background(), pkiEzsigntemplateformfieldgroupID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
