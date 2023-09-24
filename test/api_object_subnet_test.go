/*
eZmax API Definition (Full)

Testing ObjectSubnetAPIService

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

func Test_eZmaxApi_ObjectSubnetAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ObjectSubnetAPIService SubnetCreateObjectV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ObjectSubnetAPI.SubnetCreateObjectV1(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectSubnetAPIService SubnetDeleteObjectV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiSubnetID int32

		resp, httpRes, err := apiClient.ObjectSubnetAPI.SubnetDeleteObjectV1(context.Background(), pkiSubnetID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectSubnetAPIService SubnetEditObjectV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiSubnetID int32

		resp, httpRes, err := apiClient.ObjectSubnetAPI.SubnetEditObjectV1(context.Background(), pkiSubnetID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectSubnetAPIService SubnetGetObjectV2", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiSubnetID int32

		resp, httpRes, err := apiClient.ObjectSubnetAPI.SubnetGetObjectV2(context.Background(), pkiSubnetID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}