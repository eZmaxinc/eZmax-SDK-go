/*
eZmax API Definition (Full)

Testing ObjectEzsignfolderAPIService

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

func Test_eZmaxApi_ObjectEzsignfolderAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ObjectEzsignfolderAPIService EzsignfolderArchiveV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiEzsignfolderID int32

		resp, httpRes, err := apiClient.ObjectEzsignfolderAPI.EzsignfolderArchiveV1(context.Background(), pkiEzsignfolderID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectEzsignfolderAPIService EzsignfolderBatchDownloadV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiEzsignfolderID int32

		resp, httpRes, err := apiClient.ObjectEzsignfolderAPI.EzsignfolderBatchDownloadV1(context.Background(), pkiEzsignfolderID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectEzsignfolderAPIService EzsignfolderCreateObjectV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ObjectEzsignfolderAPI.EzsignfolderCreateObjectV1(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectEzsignfolderAPIService EzsignfolderCreateObjectV2", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ObjectEzsignfolderAPI.EzsignfolderCreateObjectV2(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectEzsignfolderAPIService EzsignfolderDeleteObjectV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiEzsignfolderID int32

		resp, httpRes, err := apiClient.ObjectEzsignfolderAPI.EzsignfolderDeleteObjectV1(context.Background(), pkiEzsignfolderID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectEzsignfolderAPIService EzsignfolderDisposeEzsignfoldersV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ObjectEzsignfolderAPI.EzsignfolderDisposeEzsignfoldersV1(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectEzsignfolderAPIService EzsignfolderDisposeV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiEzsignfolderID int32

		resp, httpRes, err := apiClient.ObjectEzsignfolderAPI.EzsignfolderDisposeV1(context.Background(), pkiEzsignfolderID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectEzsignfolderAPIService EzsignfolderEditObjectV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiEzsignfolderID int32

		resp, httpRes, err := apiClient.ObjectEzsignfolderAPI.EzsignfolderEditObjectV1(context.Background(), pkiEzsignfolderID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectEzsignfolderAPIService EzsignfolderGetActionableElementsV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiEzsignfolderID int32

		resp, httpRes, err := apiClient.ObjectEzsignfolderAPI.EzsignfolderGetActionableElementsV1(context.Background(), pkiEzsignfolderID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectEzsignfolderAPIService EzsignfolderGetCommunicationCountV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiEzsignfolderID int32

		resp, httpRes, err := apiClient.ObjectEzsignfolderAPI.EzsignfolderGetCommunicationCountV1(context.Background(), pkiEzsignfolderID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectEzsignfolderAPIService EzsignfolderGetCommunicationListV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiEzsignfolderID int32

		resp, httpRes, err := apiClient.ObjectEzsignfolderAPI.EzsignfolderGetCommunicationListV1(context.Background(), pkiEzsignfolderID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectEzsignfolderAPIService EzsignfolderGetEzsigndocumentsV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiEzsignfolderID int32

		resp, httpRes, err := apiClient.ObjectEzsignfolderAPI.EzsignfolderGetEzsigndocumentsV1(context.Background(), pkiEzsignfolderID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectEzsignfolderAPIService EzsignfolderGetEzsignfoldersignerassociationsV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiEzsignfolderID int32

		resp, httpRes, err := apiClient.ObjectEzsignfolderAPI.EzsignfolderGetEzsignfoldersignerassociationsV1(context.Background(), pkiEzsignfolderID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectEzsignfolderAPIService EzsignfolderGetEzsignfoldersignerassociationsmineV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiEzsignfolderID int32

		resp, httpRes, err := apiClient.ObjectEzsignfolderAPI.EzsignfolderGetEzsignfoldersignerassociationsmineV1(context.Background(), pkiEzsignfolderID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectEzsignfolderAPIService EzsignfolderGetEzsignsignaturesAutomaticV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiEzsignfolderID int32

		resp, httpRes, err := apiClient.ObjectEzsignfolderAPI.EzsignfolderGetEzsignsignaturesAutomaticV1(context.Background(), pkiEzsignfolderID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectEzsignfolderAPIService EzsignfolderGetFormsDataV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiEzsignfolderID int32

		resp, httpRes, err := apiClient.ObjectEzsignfolderAPI.EzsignfolderGetFormsDataV1(context.Background(), pkiEzsignfolderID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectEzsignfolderAPIService EzsignfolderGetListV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ObjectEzsignfolderAPI.EzsignfolderGetListV1(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectEzsignfolderAPIService EzsignfolderGetObjectV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiEzsignfolderID int32

		resp, httpRes, err := apiClient.ObjectEzsignfolderAPI.EzsignfolderGetObjectV1(context.Background(), pkiEzsignfolderID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectEzsignfolderAPIService EzsignfolderGetObjectV2", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiEzsignfolderID int32

		resp, httpRes, err := apiClient.ObjectEzsignfolderAPI.EzsignfolderGetObjectV2(context.Background(), pkiEzsignfolderID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectEzsignfolderAPIService EzsignfolderImportEzsignfoldersignerassociationsV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiEzsignfolderID int32

		resp, httpRes, err := apiClient.ObjectEzsignfolderAPI.EzsignfolderImportEzsignfoldersignerassociationsV1(context.Background(), pkiEzsignfolderID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectEzsignfolderAPIService EzsignfolderImportEzsigntemplatepackageV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiEzsignfolderID int32

		resp, httpRes, err := apiClient.ObjectEzsignfolderAPI.EzsignfolderImportEzsigntemplatepackageV1(context.Background(), pkiEzsignfolderID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectEzsignfolderAPIService EzsignfolderReorderV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiEzsignfolderID int32

		resp, httpRes, err := apiClient.ObjectEzsignfolderAPI.EzsignfolderReorderV1(context.Background(), pkiEzsignfolderID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectEzsignfolderAPIService EzsignfolderSendV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiEzsignfolderID int32

		resp, httpRes, err := apiClient.ObjectEzsignfolderAPI.EzsignfolderSendV1(context.Background(), pkiEzsignfolderID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectEzsignfolderAPIService EzsignfolderSendV2", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiEzsignfolderID int32

		resp, httpRes, err := apiClient.ObjectEzsignfolderAPI.EzsignfolderSendV2(context.Background(), pkiEzsignfolderID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectEzsignfolderAPIService EzsignfolderSendV3", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiEzsignfolderID int32

		resp, httpRes, err := apiClient.ObjectEzsignfolderAPI.EzsignfolderSendV3(context.Background(), pkiEzsignfolderID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectEzsignfolderAPIService EzsignfolderUnsendV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiEzsignfolderID int32

		resp, httpRes, err := apiClient.ObjectEzsignfolderAPI.EzsignfolderUnsendV1(context.Background(), pkiEzsignfolderID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
