/*
eZmax API Definition (Full)

Testing ObjectRejectedoffertopurchaseAPIService

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

func Test_eZmaxApi_ObjectRejectedoffertopurchaseAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ObjectRejectedoffertopurchaseAPIService RejectedoffertopurchaseGetCommunicationCountV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiRejectedoffertopurchaseID int32

		resp, httpRes, err := apiClient.ObjectRejectedoffertopurchaseAPI.RejectedoffertopurchaseGetCommunicationCountV1(context.Background(), pkiRejectedoffertopurchaseID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectRejectedoffertopurchaseAPIService RejectedoffertopurchaseGetCommunicationListV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiRejectedoffertopurchaseID int32

		resp, httpRes, err := apiClient.ObjectRejectedoffertopurchaseAPI.RejectedoffertopurchaseGetCommunicationListV1(context.Background(), pkiRejectedoffertopurchaseID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectRejectedoffertopurchaseAPIService RejectedoffertopurchaseGetCommunicationrecipientsV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiRejectedoffertopurchaseID int32

		resp, httpRes, err := apiClient.ObjectRejectedoffertopurchaseAPI.RejectedoffertopurchaseGetCommunicationrecipientsV1(context.Background(), pkiRejectedoffertopurchaseID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectRejectedoffertopurchaseAPIService RejectedoffertopurchaseGetCommunicationsendersV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiRejectedoffertopurchaseID int32

		resp, httpRes, err := apiClient.ObjectRejectedoffertopurchaseAPI.RejectedoffertopurchaseGetCommunicationsendersV1(context.Background(), pkiRejectedoffertopurchaseID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
