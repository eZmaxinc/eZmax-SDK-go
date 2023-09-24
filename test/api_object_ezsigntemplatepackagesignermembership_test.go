/*
eZmax API Definition (Full)

Testing ObjectEzsigntemplatepackagesignermembershipAPIService

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

func Test_eZmaxApi_ObjectEzsigntemplatepackagesignermembershipAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ObjectEzsigntemplatepackagesignermembershipAPIService EzsigntemplatepackagesignermembershipCreateObjectV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ObjectEzsigntemplatepackagesignermembershipAPI.EzsigntemplatepackagesignermembershipCreateObjectV1(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectEzsigntemplatepackagesignermembershipAPIService EzsigntemplatepackagesignermembershipDeleteObjectV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiEzsigntemplatepackagesignermembershipID int32

		resp, httpRes, err := apiClient.ObjectEzsigntemplatepackagesignermembershipAPI.EzsigntemplatepackagesignermembershipDeleteObjectV1(context.Background(), pkiEzsigntemplatepackagesignermembershipID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectEzsigntemplatepackagesignermembershipAPIService EzsigntemplatepackagesignermembershipGetObjectV2", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiEzsigntemplatepackagesignermembershipID int32

		resp, httpRes, err := apiClient.ObjectEzsigntemplatepackagesignermembershipAPI.EzsigntemplatepackagesignermembershipGetObjectV2(context.Background(), pkiEzsigntemplatepackagesignermembershipID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}