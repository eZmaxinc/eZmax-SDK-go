/*
eZmax API Definition (Full)

Testing ObjectAttachmentAPIService

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

func Test_eZmaxApi_ObjectAttachmentAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ObjectAttachmentAPIService AttachmentDownloadV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiAttachmentID int32

		httpRes, err := apiClient.ObjectAttachmentAPI.AttachmentDownloadV1(context.Background(), pkiAttachmentID).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectAttachmentAPIService AttachmentGetAttachmentlogsV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiAttachmentID int32

		resp, httpRes, err := apiClient.ObjectAttachmentAPI.AttachmentGetAttachmentlogsV1(context.Background(), pkiAttachmentID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
