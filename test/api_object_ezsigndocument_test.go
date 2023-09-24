/*
eZmax API Definition (Full)

Testing ObjectEzsigndocumentAPIService

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

func Test_eZmaxApi_ObjectEzsigndocumentAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ObjectEzsigndocumentAPIService EzsigndocumentApplyEzsigntemplateV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiEzsigndocumentID int32

		resp, httpRes, err := apiClient.ObjectEzsigndocumentAPI.EzsigndocumentApplyEzsigntemplateV1(context.Background(), pkiEzsigndocumentID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectEzsigndocumentAPIService EzsigndocumentApplyEzsigntemplateV2", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiEzsigndocumentID int32

		resp, httpRes, err := apiClient.ObjectEzsigndocumentAPI.EzsigndocumentApplyEzsigntemplateV2(context.Background(), pkiEzsigndocumentID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectEzsigndocumentAPIService EzsigndocumentCreateObjectV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ObjectEzsigndocumentAPI.EzsigndocumentCreateObjectV1(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectEzsigndocumentAPIService EzsigndocumentCreateObjectV2", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ObjectEzsigndocumentAPI.EzsigndocumentCreateObjectV2(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectEzsigndocumentAPIService EzsigndocumentDeclineToSignV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiEzsigndocumentID int32

		resp, httpRes, err := apiClient.ObjectEzsigndocumentAPI.EzsigndocumentDeclineToSignV1(context.Background(), pkiEzsigndocumentID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectEzsigndocumentAPIService EzsigndocumentDeleteObjectV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiEzsigndocumentID int32

		resp, httpRes, err := apiClient.ObjectEzsigndocumentAPI.EzsigndocumentDeleteObjectV1(context.Background(), pkiEzsigndocumentID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectEzsigndocumentAPIService EzsigndocumentEditEzsignformfieldgroupsV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiEzsigndocumentID int32

		resp, httpRes, err := apiClient.ObjectEzsigndocumentAPI.EzsigndocumentEditEzsignformfieldgroupsV1(context.Background(), pkiEzsigndocumentID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectEzsigndocumentAPIService EzsigndocumentEditEzsignsignaturesV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiEzsigndocumentID int32

		resp, httpRes, err := apiClient.ObjectEzsigndocumentAPI.EzsigndocumentEditEzsignsignaturesV1(context.Background(), pkiEzsigndocumentID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectEzsigndocumentAPIService EzsigndocumentEndPrematurelyV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiEzsigndocumentID int32

		resp, httpRes, err := apiClient.ObjectEzsigndocumentAPI.EzsigndocumentEndPrematurelyV1(context.Background(), pkiEzsigndocumentID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectEzsigndocumentAPIService EzsigndocumentFlattenV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiEzsigndocumentID int32

		resp, httpRes, err := apiClient.ObjectEzsigndocumentAPI.EzsigndocumentFlattenV1(context.Background(), pkiEzsigndocumentID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectEzsigndocumentAPIService EzsigndocumentGetActionableElementsV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiEzsigndocumentID int32

		resp, httpRes, err := apiClient.ObjectEzsigndocumentAPI.EzsigndocumentGetActionableElementsV1(context.Background(), pkiEzsigndocumentID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectEzsigndocumentAPIService EzsigndocumentGetCompletedElementsV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiEzsigndocumentID int32

		resp, httpRes, err := apiClient.ObjectEzsigndocumentAPI.EzsigndocumentGetCompletedElementsV1(context.Background(), pkiEzsigndocumentID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectEzsigndocumentAPIService EzsigndocumentGetDownloadUrlV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiEzsigndocumentID int32
		var eDocumentType string

		resp, httpRes, err := apiClient.ObjectEzsigndocumentAPI.EzsigndocumentGetDownloadUrlV1(context.Background(), pkiEzsigndocumentID, eDocumentType).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectEzsigndocumentAPIService EzsigndocumentGetEzsignannotationsV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiEzsigndocumentID int32

		resp, httpRes, err := apiClient.ObjectEzsigndocumentAPI.EzsigndocumentGetEzsignannotationsV1(context.Background(), pkiEzsigndocumentID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectEzsigndocumentAPIService EzsigndocumentGetEzsignformfieldgroupsV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiEzsigndocumentID int32

		resp, httpRes, err := apiClient.ObjectEzsigndocumentAPI.EzsigndocumentGetEzsignformfieldgroupsV1(context.Background(), pkiEzsigndocumentID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectEzsigndocumentAPIService EzsigndocumentGetEzsignpagesV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiEzsigndocumentID int32

		resp, httpRes, err := apiClient.ObjectEzsigndocumentAPI.EzsigndocumentGetEzsignpagesV1(context.Background(), pkiEzsigndocumentID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectEzsigndocumentAPIService EzsigndocumentGetEzsignsignaturesAutomaticV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiEzsigndocumentID int32

		resp, httpRes, err := apiClient.ObjectEzsigndocumentAPI.EzsigndocumentGetEzsignsignaturesAutomaticV1(context.Background(), pkiEzsigndocumentID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectEzsigndocumentAPIService EzsigndocumentGetEzsignsignaturesV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiEzsigndocumentID int32

		resp, httpRes, err := apiClient.ObjectEzsigndocumentAPI.EzsigndocumentGetEzsignsignaturesV1(context.Background(), pkiEzsigndocumentID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectEzsigndocumentAPIService EzsigndocumentGetFormDataV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiEzsigndocumentID int32

		resp, httpRes, err := apiClient.ObjectEzsigndocumentAPI.EzsigndocumentGetFormDataV1(context.Background(), pkiEzsigndocumentID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectEzsigndocumentAPIService EzsigndocumentGetObjectV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiEzsigndocumentID int32

		resp, httpRes, err := apiClient.ObjectEzsigndocumentAPI.EzsigndocumentGetObjectV1(context.Background(), pkiEzsigndocumentID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectEzsigndocumentAPIService EzsigndocumentGetObjectV2", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiEzsigndocumentID int32

		resp, httpRes, err := apiClient.ObjectEzsigndocumentAPI.EzsigndocumentGetObjectV2(context.Background(), pkiEzsigndocumentID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectEzsigndocumentAPIService EzsigndocumentGetTemporaryProofV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiEzsigndocumentID int32

		resp, httpRes, err := apiClient.ObjectEzsigndocumentAPI.EzsigndocumentGetTemporaryProofV1(context.Background(), pkiEzsigndocumentID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectEzsigndocumentAPIService EzsigndocumentGetWordsPositionsV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiEzsigndocumentID int32

		resp, httpRes, err := apiClient.ObjectEzsigndocumentAPI.EzsigndocumentGetWordsPositionsV1(context.Background(), pkiEzsigndocumentID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectEzsigndocumentAPIService EzsigndocumentPatchObjectV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiEzsigndocumentID int32

		resp, httpRes, err := apiClient.ObjectEzsigndocumentAPI.EzsigndocumentPatchObjectV1(context.Background(), pkiEzsigndocumentID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectEzsigndocumentAPIService EzsigndocumentSubmitEzsignformV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiEzsigndocumentID int32

		resp, httpRes, err := apiClient.ObjectEzsigndocumentAPI.EzsigndocumentSubmitEzsignformV1(context.Background(), pkiEzsigndocumentID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectEzsigndocumentAPIService EzsigndocumentUnsendV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pkiEzsigndocumentID int32

		resp, httpRes, err := apiClient.ObjectEzsigndocumentAPI.EzsigndocumentUnsendV1(context.Background(), pkiEzsigndocumentID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
