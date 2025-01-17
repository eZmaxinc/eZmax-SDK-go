# eZmaxAPI\ObjectEzsigndocumentAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EzsigndocumentApplyEzsigntemplateV1**](ObjectEzsigndocumentAPI.md#EzsigndocumentApplyEzsigntemplateV1) | **Post** /1/object/ezsigndocument/{pkiEzsigndocumentID}/applyezsigntemplate | Apply an Ezsigntemplate to the Ezsigndocument.
[**EzsigndocumentApplyEzsigntemplateV2**](ObjectEzsigndocumentAPI.md#EzsigndocumentApplyEzsigntemplateV2) | **Post** /2/object/ezsigndocument/{pkiEzsigndocumentID}/applyEzsigntemplate | Apply an Ezsigntemplate to the Ezsigndocument.
[**EzsigndocumentApplyEzsigntemplateglobalV1**](ObjectEzsigndocumentAPI.md#EzsigndocumentApplyEzsigntemplateglobalV1) | **Post** /1/object/ezsigndocument/{pkiEzsigndocumentID}/applyEzsigntemplateglobal | Apply an Ezsigntemplateglobal to the Ezsigndocument.
[**EzsigndocumentCreateEzsignelementsPositionedByWordV1**](ObjectEzsigndocumentAPI.md#EzsigndocumentCreateEzsignelementsPositionedByWordV1) | **Post** /1/object/ezsigndocument/{pkiEzsigndocumentID}/createEzsignelementsPositionedByWord | Create multiple Ezsignsignatures/Ezsignformfieldgroups
[**EzsigndocumentCreateObjectV1**](ObjectEzsigndocumentAPI.md#EzsigndocumentCreateObjectV1) | **Post** /1/object/ezsigndocument | Create a new Ezsigndocument
[**EzsigndocumentCreateObjectV2**](ObjectEzsigndocumentAPI.md#EzsigndocumentCreateObjectV2) | **Post** /2/object/ezsigndocument | Create a new Ezsigndocument
[**EzsigndocumentCreateObjectV3**](ObjectEzsigndocumentAPI.md#EzsigndocumentCreateObjectV3) | **Post** /3/object/ezsigndocument | Create a new Ezsigndocument
[**EzsigndocumentDeclineToSignV1**](ObjectEzsigndocumentAPI.md#EzsigndocumentDeclineToSignV1) | **Post** /1/object/ezsigndocument/{pkiEzsigndocumentID}/declineToSign | Decline to sign
[**EzsigndocumentDeleteObjectV1**](ObjectEzsigndocumentAPI.md#EzsigndocumentDeleteObjectV1) | **Delete** /1/object/ezsigndocument/{pkiEzsigndocumentID} | Delete an existing Ezsigndocument
[**EzsigndocumentEditEzsignannotationsV1**](ObjectEzsigndocumentAPI.md#EzsigndocumentEditEzsignannotationsV1) | **Put** /1/object/ezsigndocument/{pkiEzsigndocumentID}/editEzsignannotations | Edit multiple Ezsignannotations
[**EzsigndocumentEditEzsignformfieldgroupsV1**](ObjectEzsigndocumentAPI.md#EzsigndocumentEditEzsignformfieldgroupsV1) | **Put** /1/object/ezsigndocument/{pkiEzsigndocumentID}/editEzsignformfieldgroups | Edit multiple Ezsignformfieldgroups
[**EzsigndocumentEditEzsignsignaturesV1**](ObjectEzsigndocumentAPI.md#EzsigndocumentEditEzsignsignaturesV1) | **Put** /1/object/ezsigndocument/{pkiEzsigndocumentID}/editEzsignsignatures | Edit multiple Ezsignsignatures
[**EzsigndocumentEditObjectV1**](ObjectEzsigndocumentAPI.md#EzsigndocumentEditObjectV1) | **Put** /1/object/ezsigndocument/{pkiEzsigndocumentID} | Edit an existing Ezsigndocument
[**EzsigndocumentEndPrematurelyV1**](ObjectEzsigndocumentAPI.md#EzsigndocumentEndPrematurelyV1) | **Post** /1/object/ezsigndocument/{pkiEzsigndocumentID}/endPrematurely | End prematurely
[**EzsigndocumentExtractTextV1**](ObjectEzsigndocumentAPI.md#EzsigndocumentExtractTextV1) | **Post** /1/object/ezsigndocument/{pkiEzsigndocumentID}/extractText | Extract text from Ezsigndocument area
[**EzsigndocumentFlattenV1**](ObjectEzsigndocumentAPI.md#EzsigndocumentFlattenV1) | **Post** /1/object/ezsigndocument/{pkiEzsigndocumentID}/flatten | Flatten
[**EzsigndocumentGetActionableElementsV1**](ObjectEzsigndocumentAPI.md#EzsigndocumentGetActionableElementsV1) | **Get** /1/object/ezsigndocument/{pkiEzsigndocumentID}/getActionableElements | Retrieve actionable elements for the Ezsigndocument
[**EzsigndocumentGetAttachmentsV1**](ObjectEzsigndocumentAPI.md#EzsigndocumentGetAttachmentsV1) | **Get** /1/object/ezsigndocument/{pkiEzsigndocumentID}/getAttachments | Retrieve Ezsigndocument&#39;s Attachments
[**EzsigndocumentGetCompletedElementsV1**](ObjectEzsigndocumentAPI.md#EzsigndocumentGetCompletedElementsV1) | **Get** /1/object/ezsigndocument/{pkiEzsigndocumentID}/getCompletedElements | Retrieve completed elements for the Ezsigndocument
[**EzsigndocumentGetDownloadUrlV1**](ObjectEzsigndocumentAPI.md#EzsigndocumentGetDownloadUrlV1) | **Get** /1/object/ezsigndocument/{pkiEzsigndocumentID}/getDownloadUrl/{eDocumentType} | Retrieve a URL to download documents.
[**EzsigndocumentGetEzsignannotationsV1**](ObjectEzsigndocumentAPI.md#EzsigndocumentGetEzsignannotationsV1) | **Get** /1/object/ezsigndocument/{pkiEzsigndocumentID}/getEzsignannotations | Retrieve an existing Ezsigndocument&#39;s Ezsignannotations
[**EzsigndocumentGetEzsigndiscussionsV1**](ObjectEzsigndocumentAPI.md#EzsigndocumentGetEzsigndiscussionsV1) | **Get** /1/object/ezsigndocument/{pkiEzsigndocumentID}/getEzsigndiscussions | Retrieve an existing Ezsigndocument&#39;s Ezsigndiscussions
[**EzsigndocumentGetEzsignformfieldgroupsV1**](ObjectEzsigndocumentAPI.md#EzsigndocumentGetEzsignformfieldgroupsV1) | **Get** /1/object/ezsigndocument/{pkiEzsigndocumentID}/getEzsignformfieldgroups | Retrieve an existing Ezsigndocument&#39;s Ezsignformfieldgroups
[**EzsigndocumentGetEzsignpagesV1**](ObjectEzsigndocumentAPI.md#EzsigndocumentGetEzsignpagesV1) | **Get** /1/object/ezsigndocument/{pkiEzsigndocumentID}/getEzsignpages | Retrieve an existing Ezsigndocument&#39;s Ezsignpages
[**EzsigndocumentGetEzsignsignaturesAutomaticV1**](ObjectEzsigndocumentAPI.md#EzsigndocumentGetEzsignsignaturesAutomaticV1) | **Get** /1/object/ezsigndocument/{pkiEzsigndocumentID}/getEzsignsignaturesAutomatic | Retrieve an existing Ezsigndocument&#39;s automatic Ezsignsignatures
[**EzsigndocumentGetEzsignsignaturesV1**](ObjectEzsigndocumentAPI.md#EzsigndocumentGetEzsignsignaturesV1) | **Get** /1/object/ezsigndocument/{pkiEzsigndocumentID}/getEzsignsignatures | Retrieve an existing Ezsigndocument&#39;s Ezsignsignatures
[**EzsigndocumentGetFormDataV1**](ObjectEzsigndocumentAPI.md#EzsigndocumentGetFormDataV1) | **Get** /1/object/ezsigndocument/{pkiEzsigndocumentID}/getFormData | Retrieve an existing Ezsigndocument&#39;s Form Data
[**EzsigndocumentGetObjectV1**](ObjectEzsigndocumentAPI.md#EzsigndocumentGetObjectV1) | **Get** /1/object/ezsigndocument/{pkiEzsigndocumentID} | Retrieve an existing Ezsigndocument
[**EzsigndocumentGetObjectV2**](ObjectEzsigndocumentAPI.md#EzsigndocumentGetObjectV2) | **Get** /2/object/ezsigndocument/{pkiEzsigndocumentID} | Retrieve an existing Ezsigndocument
[**EzsigndocumentGetTemporaryProofV1**](ObjectEzsigndocumentAPI.md#EzsigndocumentGetTemporaryProofV1) | **Get** /1/object/ezsigndocument/{pkiEzsigndocumentID}/getTemporaryProof | Retrieve the temporary proof
[**EzsigndocumentGetWordsPositionsV1**](ObjectEzsigndocumentAPI.md#EzsigndocumentGetWordsPositionsV1) | **Post** /1/object/ezsigndocument/{pkiEzsigndocumentID}/getWordsPositions | Retrieve positions X,Y of given words from a Ezsigndocument
[**EzsigndocumentPatchObjectV1**](ObjectEzsigndocumentAPI.md#EzsigndocumentPatchObjectV1) | **Patch** /1/object/ezsigndocument/{pkiEzsigndocumentID} | Patch an existing Ezsigndocument
[**EzsigndocumentSubmitEzsignformV1**](ObjectEzsigndocumentAPI.md#EzsigndocumentSubmitEzsignformV1) | **Post** /1/object/ezsigndocument/{pkiEzsigndocumentID}/submitEzsignform | Submit the Ezsignform
[**EzsigndocumentUnsendV1**](ObjectEzsigndocumentAPI.md#EzsigndocumentUnsendV1) | **Post** /1/object/ezsigndocument/{pkiEzsigndocumentID}/unsend | Unsend the Ezsigndocument



## EzsigndocumentApplyEzsigntemplateV1

> CommonResponse EzsigndocumentApplyEzsigntemplateV1(ctx, pkiEzsigndocumentID).EzsigndocumentApplyEzsigntemplateV1Request(ezsigndocumentApplyEzsigntemplateV1Request).Execute()

Apply an Ezsigntemplate to the Ezsigndocument.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ezmaxinc/ezmax-sdk-go"
)

func main() {
	pkiEzsigndocumentID := int32(56) // int32 | 
	ezsigndocumentApplyEzsigntemplateV1Request := *openapiclient.NewEzsigndocumentApplyEzsigntemplateV1Request(int32(36), []string{"John"}, []int32{int32(20)}) // EzsigndocumentApplyEzsigntemplateV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigndocumentAPI.EzsigndocumentApplyEzsigntemplateV1(context.Background(), pkiEzsigndocumentID).EzsigndocumentApplyEzsigntemplateV1Request(ezsigndocumentApplyEzsigntemplateV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigndocumentAPI.EzsigndocumentApplyEzsigntemplateV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigndocumentApplyEzsigntemplateV1`: CommonResponse
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigndocumentAPI.EzsigndocumentApplyEzsigntemplateV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigndocumentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigndocumentApplyEzsigntemplateV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezsigndocumentApplyEzsigntemplateV1Request** | [**EzsigndocumentApplyEzsigntemplateV1Request**](EzsigndocumentApplyEzsigntemplateV1Request.md) |  | 

### Return type

[**CommonResponse**](CommonResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigndocumentApplyEzsigntemplateV2

> EzsigndocumentApplyEzsigntemplateV2Response EzsigndocumentApplyEzsigntemplateV2(ctx, pkiEzsigndocumentID).EzsigndocumentApplyEzsigntemplateV2Request(ezsigndocumentApplyEzsigntemplateV2Request).Execute()

Apply an Ezsigntemplate to the Ezsigndocument.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ezmaxinc/ezmax-sdk-go"
)

func main() {
	pkiEzsigndocumentID := int32(56) // int32 | 
	ezsigndocumentApplyEzsigntemplateV2Request := *openapiclient.NewEzsigndocumentApplyEzsigntemplateV2Request(int32(36), []string{"John"}, []int32{int32(20)}) // EzsigndocumentApplyEzsigntemplateV2Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigndocumentAPI.EzsigndocumentApplyEzsigntemplateV2(context.Background(), pkiEzsigndocumentID).EzsigndocumentApplyEzsigntemplateV2Request(ezsigndocumentApplyEzsigntemplateV2Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigndocumentAPI.EzsigndocumentApplyEzsigntemplateV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigndocumentApplyEzsigntemplateV2`: EzsigndocumentApplyEzsigntemplateV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigndocumentAPI.EzsigndocumentApplyEzsigntemplateV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigndocumentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigndocumentApplyEzsigntemplateV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezsigndocumentApplyEzsigntemplateV2Request** | [**EzsigndocumentApplyEzsigntemplateV2Request**](EzsigndocumentApplyEzsigntemplateV2Request.md) |  | 

### Return type

[**EzsigndocumentApplyEzsigntemplateV2Response**](EzsigndocumentApplyEzsigntemplateV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigndocumentApplyEzsigntemplateglobalV1

> EzsigndocumentApplyEzsigntemplateglobalV1Response EzsigndocumentApplyEzsigntemplateglobalV1(ctx, pkiEzsigndocumentID).EzsigndocumentApplyEzsigntemplateglobalV1Request(ezsigndocumentApplyEzsigntemplateglobalV1Request).Execute()

Apply an Ezsigntemplateglobal to the Ezsigndocument.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ezmaxinc/ezmax-sdk-go"
)

func main() {
	pkiEzsigndocumentID := int32(56) // int32 | 
	ezsigndocumentApplyEzsigntemplateglobalV1Request := *openapiclient.NewEzsigndocumentApplyEzsigntemplateglobalV1Request(int32(36), []string{"John"}, []int32{int32(20)}) // EzsigndocumentApplyEzsigntemplateglobalV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigndocumentAPI.EzsigndocumentApplyEzsigntemplateglobalV1(context.Background(), pkiEzsigndocumentID).EzsigndocumentApplyEzsigntemplateglobalV1Request(ezsigndocumentApplyEzsigntemplateglobalV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigndocumentAPI.EzsigndocumentApplyEzsigntemplateglobalV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigndocumentApplyEzsigntemplateglobalV1`: EzsigndocumentApplyEzsigntemplateglobalV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigndocumentAPI.EzsigndocumentApplyEzsigntemplateglobalV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigndocumentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigndocumentApplyEzsigntemplateglobalV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezsigndocumentApplyEzsigntemplateglobalV1Request** | [**EzsigndocumentApplyEzsigntemplateglobalV1Request**](EzsigndocumentApplyEzsigntemplateglobalV1Request.md) |  | 

### Return type

[**EzsigndocumentApplyEzsigntemplateglobalV1Response**](EzsigndocumentApplyEzsigntemplateglobalV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigndocumentCreateEzsignelementsPositionedByWordV1

> EzsigndocumentCreateEzsignelementsPositionedByWordV1Response EzsigndocumentCreateEzsignelementsPositionedByWordV1(ctx, pkiEzsigndocumentID).EzsigndocumentCreateEzsignelementsPositionedByWordV1Request(ezsigndocumentCreateEzsignelementsPositionedByWordV1Request).Execute()

Create multiple Ezsignsignatures/Ezsignformfieldgroups



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ezmaxinc/ezmax-sdk-go"
)

func main() {
	pkiEzsigndocumentID := int32(56) // int32 | 
	ezsigndocumentCreateEzsignelementsPositionedByWordV1Request := *openapiclient.NewEzsigndocumentCreateEzsignelementsPositionedByWordV1Request([]openapiclient.CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest{*openapiclient.NewCustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest(*openapiclient.NewCustomCreateEzsignelementsPositionedByWordRequest("SCreateezsignelementspositionedbywordPattern_example", int32(123), int32(123), "ECreateezsignelementspositionedbywordOccurance_example"), int32(97), openapiclient.Field-eEzsignformfieldgroupType("Text"), "Allergies", int32(1), int32(1), int32(2), false, []openapiclient.EzsignformfieldgroupsignerRequestCompound{*openapiclient.NewEzsignformfieldgroupsignerRequest(int32(20))}, []openapiclient.EzsignformfieldRequestCompound{*openapiclient.NewEzsignformfieldRequestCompound(int32(1), "Peanuts", int32(200), int32(300), int32(102), int32(22))})}, []openapiclient.CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest{*openapiclient.NewCustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest(*openapiclient.NewCustomCreateEzsignelementsPositionedByWordRequest("SCreateezsignelementspositionedbywordPattern_example", int32(123), int32(123), "ECreateezsignelementspositionedbywordOccurance_example"), int32(20), int32(1), int32(200), int32(300), int32(1), openapiclient.Field-eEzsignsignatureType("Acknowledgement"), int32(97))}) // EzsigndocumentCreateEzsignelementsPositionedByWordV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigndocumentAPI.EzsigndocumentCreateEzsignelementsPositionedByWordV1(context.Background(), pkiEzsigndocumentID).EzsigndocumentCreateEzsignelementsPositionedByWordV1Request(ezsigndocumentCreateEzsignelementsPositionedByWordV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigndocumentAPI.EzsigndocumentCreateEzsignelementsPositionedByWordV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigndocumentCreateEzsignelementsPositionedByWordV1`: EzsigndocumentCreateEzsignelementsPositionedByWordV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigndocumentAPI.EzsigndocumentCreateEzsignelementsPositionedByWordV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigndocumentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigndocumentCreateEzsignelementsPositionedByWordV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezsigndocumentCreateEzsignelementsPositionedByWordV1Request** | [**EzsigndocumentCreateEzsignelementsPositionedByWordV1Request**](EzsigndocumentCreateEzsignelementsPositionedByWordV1Request.md) |  | 

### Return type

[**EzsigndocumentCreateEzsignelementsPositionedByWordV1Response**](EzsigndocumentCreateEzsignelementsPositionedByWordV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigndocumentCreateObjectV1

> EzsigndocumentCreateObjectV1Response EzsigndocumentCreateObjectV1(ctx).EzsigndocumentCreateObjectV1Request(ezsigndocumentCreateObjectV1Request).Execute()

Create a new Ezsigndocument



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ezmaxinc/ezmax-sdk-go"
)

func main() {
	ezsigndocumentCreateObjectV1Request := []openapiclient.EzsigndocumentCreateObjectV1Request{*openapiclient.NewEzsigndocumentCreateObjectV1Request()} // []EzsigndocumentCreateObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigndocumentAPI.EzsigndocumentCreateObjectV1(context.Background()).EzsigndocumentCreateObjectV1Request(ezsigndocumentCreateObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigndocumentAPI.EzsigndocumentCreateObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigndocumentCreateObjectV1`: EzsigndocumentCreateObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigndocumentAPI.EzsigndocumentCreateObjectV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEzsigndocumentCreateObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ezsigndocumentCreateObjectV1Request** | [**[]EzsigndocumentCreateObjectV1Request**](EzsigndocumentCreateObjectV1Request.md) |  | 

### Return type

[**EzsigndocumentCreateObjectV1Response**](EzsigndocumentCreateObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigndocumentCreateObjectV2

> EzsigndocumentCreateObjectV2Response EzsigndocumentCreateObjectV2(ctx).EzsigndocumentCreateObjectV2Request(ezsigndocumentCreateObjectV2Request).Execute()

Create a new Ezsigndocument



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ezmaxinc/ezmax-sdk-go"
)

func main() {
	ezsigndocumentCreateObjectV2Request := *openapiclient.NewEzsigndocumentCreateObjectV2Request([]openapiclient.EzsigndocumentRequestCompound{*openapiclient.NewEzsigndocumentRequest(int32(33), int32(2), "EEzsigndocumentSource_example", "2020-12-31 23:59:59", "Contract #123")}) // EzsigndocumentCreateObjectV2Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigndocumentAPI.EzsigndocumentCreateObjectV2(context.Background()).EzsigndocumentCreateObjectV2Request(ezsigndocumentCreateObjectV2Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigndocumentAPI.EzsigndocumentCreateObjectV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigndocumentCreateObjectV2`: EzsigndocumentCreateObjectV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigndocumentAPI.EzsigndocumentCreateObjectV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEzsigndocumentCreateObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ezsigndocumentCreateObjectV2Request** | [**EzsigndocumentCreateObjectV2Request**](EzsigndocumentCreateObjectV2Request.md) |  | 

### Return type

[**EzsigndocumentCreateObjectV2Response**](EzsigndocumentCreateObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigndocumentCreateObjectV3

> EzsigndocumentCreateObjectV3Response EzsigndocumentCreateObjectV3(ctx).EzsigndocumentCreateObjectV3Request(ezsigndocumentCreateObjectV3Request).Execute()

Create a new Ezsigndocument



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ezmaxinc/ezmax-sdk-go"
)

func main() {
	ezsigndocumentCreateObjectV3Request := *openapiclient.NewEzsigndocumentCreateObjectV3Request([]openapiclient.EzsigndocumentRequestCompound{*openapiclient.NewEzsigndocumentRequest(int32(33), int32(2), "EEzsigndocumentSource_example", "2020-12-31 23:59:59", "Contract #123")}) // EzsigndocumentCreateObjectV3Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigndocumentAPI.EzsigndocumentCreateObjectV3(context.Background()).EzsigndocumentCreateObjectV3Request(ezsigndocumentCreateObjectV3Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigndocumentAPI.EzsigndocumentCreateObjectV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigndocumentCreateObjectV3`: EzsigndocumentCreateObjectV3Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigndocumentAPI.EzsigndocumentCreateObjectV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEzsigndocumentCreateObjectV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ezsigndocumentCreateObjectV3Request** | [**EzsigndocumentCreateObjectV3Request**](EzsigndocumentCreateObjectV3Request.md) |  | 

### Return type

[**EzsigndocumentCreateObjectV3Response**](EzsigndocumentCreateObjectV3Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigndocumentDeclineToSignV1

> CommonResponse EzsigndocumentDeclineToSignV1(ctx, pkiEzsigndocumentID).EzsigndocumentDeclineToSignV1Request(ezsigndocumentDeclineToSignV1Request).Execute()

Decline to sign



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ezmaxinc/ezmax-sdk-go"
)

func main() {
	pkiEzsigndocumentID := int32(56) // int32 | 
	ezsigndocumentDeclineToSignV1Request := *openapiclient.NewEzsigndocumentDeclineToSignV1Request("Wrong document") // EzsigndocumentDeclineToSignV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigndocumentAPI.EzsigndocumentDeclineToSignV1(context.Background(), pkiEzsigndocumentID).EzsigndocumentDeclineToSignV1Request(ezsigndocumentDeclineToSignV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigndocumentAPI.EzsigndocumentDeclineToSignV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigndocumentDeclineToSignV1`: CommonResponse
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigndocumentAPI.EzsigndocumentDeclineToSignV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigndocumentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigndocumentDeclineToSignV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezsigndocumentDeclineToSignV1Request** | [**EzsigndocumentDeclineToSignV1Request**](EzsigndocumentDeclineToSignV1Request.md) |  | 

### Return type

[**CommonResponse**](CommonResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigndocumentDeleteObjectV1

> CommonResponse EzsigndocumentDeleteObjectV1(ctx, pkiEzsigndocumentID).Execute()

Delete an existing Ezsigndocument



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ezmaxinc/ezmax-sdk-go"
)

func main() {
	pkiEzsigndocumentID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigndocumentAPI.EzsigndocumentDeleteObjectV1(context.Background(), pkiEzsigndocumentID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigndocumentAPI.EzsigndocumentDeleteObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigndocumentDeleteObjectV1`: CommonResponse
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigndocumentAPI.EzsigndocumentDeleteObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigndocumentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigndocumentDeleteObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CommonResponse**](CommonResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigndocumentEditEzsignannotationsV1

> EzsigndocumentEditEzsignannotationsV1Response EzsigndocumentEditEzsignannotationsV1(ctx, pkiEzsigndocumentID).EzsigndocumentEditEzsignannotationsV1Request(ezsigndocumentEditEzsignannotationsV1Request).Execute()

Edit multiple Ezsignannotations



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ezmaxinc/ezmax-sdk-go"
)

func main() {
	pkiEzsigndocumentID := int32(56) // int32 | 
	ezsigndocumentEditEzsignannotationsV1Request := *openapiclient.NewEzsigndocumentEditEzsignannotationsV1Request([]openapiclient.EzsignannotationRequestCompound{*openapiclient.NewEzsignannotationRequestCompound(int32(97), openapiclient.Field-eEzsignannotationType("StrikethroughBlock"), int32(50), int32(50), int32(1))}) // EzsigndocumentEditEzsignannotationsV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigndocumentAPI.EzsigndocumentEditEzsignannotationsV1(context.Background(), pkiEzsigndocumentID).EzsigndocumentEditEzsignannotationsV1Request(ezsigndocumentEditEzsignannotationsV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigndocumentAPI.EzsigndocumentEditEzsignannotationsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigndocumentEditEzsignannotationsV1`: EzsigndocumentEditEzsignannotationsV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigndocumentAPI.EzsigndocumentEditEzsignannotationsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigndocumentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigndocumentEditEzsignannotationsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezsigndocumentEditEzsignannotationsV1Request** | [**EzsigndocumentEditEzsignannotationsV1Request**](EzsigndocumentEditEzsignannotationsV1Request.md) |  | 

### Return type

[**EzsigndocumentEditEzsignannotationsV1Response**](EzsigndocumentEditEzsignannotationsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigndocumentEditEzsignformfieldgroupsV1

> EzsigndocumentEditEzsignformfieldgroupsV1Response EzsigndocumentEditEzsignformfieldgroupsV1(ctx, pkiEzsigndocumentID).EzsigndocumentEditEzsignformfieldgroupsV1Request(ezsigndocumentEditEzsignformfieldgroupsV1Request).Execute()

Edit multiple Ezsignformfieldgroups



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ezmaxinc/ezmax-sdk-go"
)

func main() {
	pkiEzsigndocumentID := int32(56) // int32 | 
	ezsigndocumentEditEzsignformfieldgroupsV1Request := *openapiclient.NewEzsigndocumentEditEzsignformfieldgroupsV1Request([]openapiclient.EzsignformfieldgroupRequestCompound{*openapiclient.NewEzsignformfieldgroupRequestCompound([]openapiclient.EzsignformfieldgroupsignerRequestCompound{*openapiclient.NewEzsignformfieldgroupsignerRequest(int32(20))}, []openapiclient.EzsignformfieldRequestCompound{*openapiclient.NewEzsignformfieldRequestCompound(int32(1), "Peanuts", int32(200), int32(300), int32(102), int32(22))}, int32(97), openapiclient.Field-eEzsignformfieldgroupType("Text"), "Allergies", int32(1), int32(1), int32(2), false)}) // EzsigndocumentEditEzsignformfieldgroupsV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigndocumentAPI.EzsigndocumentEditEzsignformfieldgroupsV1(context.Background(), pkiEzsigndocumentID).EzsigndocumentEditEzsignformfieldgroupsV1Request(ezsigndocumentEditEzsignformfieldgroupsV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigndocumentAPI.EzsigndocumentEditEzsignformfieldgroupsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigndocumentEditEzsignformfieldgroupsV1`: EzsigndocumentEditEzsignformfieldgroupsV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigndocumentAPI.EzsigndocumentEditEzsignformfieldgroupsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigndocumentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigndocumentEditEzsignformfieldgroupsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezsigndocumentEditEzsignformfieldgroupsV1Request** | [**EzsigndocumentEditEzsignformfieldgroupsV1Request**](EzsigndocumentEditEzsignformfieldgroupsV1Request.md) |  | 

### Return type

[**EzsigndocumentEditEzsignformfieldgroupsV1Response**](EzsigndocumentEditEzsignformfieldgroupsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigndocumentEditEzsignsignaturesV1

> EzsigndocumentEditEzsignsignaturesV1Response EzsigndocumentEditEzsignsignaturesV1(ctx, pkiEzsigndocumentID).EzsigndocumentEditEzsignsignaturesV1Request(ezsigndocumentEditEzsignsignaturesV1Request).Execute()

Edit multiple Ezsignsignatures



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ezmaxinc/ezmax-sdk-go"
)

func main() {
	pkiEzsigndocumentID := int32(56) // int32 | 
	ezsigndocumentEditEzsignsignaturesV1Request := *openapiclient.NewEzsigndocumentEditEzsignsignaturesV1Request([]openapiclient.EzsignsignatureRequestCompound{*openapiclient.NewEzsignsignatureRequestCompound(int32(20), int32(1), int32(200), int32(300), int32(1), openapiclient.Field-eEzsignsignatureType("Acknowledgement"), int32(97))}) // EzsigndocumentEditEzsignsignaturesV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigndocumentAPI.EzsigndocumentEditEzsignsignaturesV1(context.Background(), pkiEzsigndocumentID).EzsigndocumentEditEzsignsignaturesV1Request(ezsigndocumentEditEzsignsignaturesV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigndocumentAPI.EzsigndocumentEditEzsignsignaturesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigndocumentEditEzsignsignaturesV1`: EzsigndocumentEditEzsignsignaturesV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigndocumentAPI.EzsigndocumentEditEzsignsignaturesV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigndocumentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigndocumentEditEzsignsignaturesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezsigndocumentEditEzsignsignaturesV1Request** | [**EzsigndocumentEditEzsignsignaturesV1Request**](EzsigndocumentEditEzsignsignaturesV1Request.md) |  | 

### Return type

[**EzsigndocumentEditEzsignsignaturesV1Response**](EzsigndocumentEditEzsignsignaturesV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigndocumentEditObjectV1

> EzsigndocumentEditObjectV1Response EzsigndocumentEditObjectV1(ctx, pkiEzsigndocumentID).EzsigndocumentEditObjectV1Request(ezsigndocumentEditObjectV1Request).Execute()

Edit an existing Ezsigndocument



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ezmaxinc/ezmax-sdk-go"
)

func main() {
	pkiEzsigndocumentID := int32(56) // int32 | 
	ezsigndocumentEditObjectV1Request := *openapiclient.NewEzsigndocumentEditObjectV1Request(*openapiclient.NewEzsigndocumentRequest(int32(33), int32(2), "EEzsigndocumentSource_example", "2020-12-31 23:59:59", "Contract #123")) // EzsigndocumentEditObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigndocumentAPI.EzsigndocumentEditObjectV1(context.Background(), pkiEzsigndocumentID).EzsigndocumentEditObjectV1Request(ezsigndocumentEditObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigndocumentAPI.EzsigndocumentEditObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigndocumentEditObjectV1`: EzsigndocumentEditObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigndocumentAPI.EzsigndocumentEditObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigndocumentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigndocumentEditObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezsigndocumentEditObjectV1Request** | [**EzsigndocumentEditObjectV1Request**](EzsigndocumentEditObjectV1Request.md) |  | 

### Return type

[**EzsigndocumentEditObjectV1Response**](EzsigndocumentEditObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigndocumentEndPrematurelyV1

> CommonResponse EzsigndocumentEndPrematurelyV1(ctx, pkiEzsigndocumentID).Body(body).Execute()

End prematurely



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ezmaxinc/ezmax-sdk-go"
)

func main() {
	pkiEzsigndocumentID := int32(56) // int32 | 
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigndocumentAPI.EzsigndocumentEndPrematurelyV1(context.Background(), pkiEzsigndocumentID).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigndocumentAPI.EzsigndocumentEndPrematurelyV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigndocumentEndPrematurelyV1`: CommonResponse
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigndocumentAPI.EzsigndocumentEndPrematurelyV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigndocumentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigndocumentEndPrematurelyV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

[**CommonResponse**](CommonResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigndocumentExtractTextV1

> EzsigndocumentExtractTextV1Response EzsigndocumentExtractTextV1(ctx, pkiEzsigndocumentID).EzsigndocumentExtractTextV1Request(ezsigndocumentExtractTextV1Request).Execute()

Extract text from Ezsigndocument area



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ezmaxinc/ezmax-sdk-go"
)

func main() {
	pkiEzsigndocumentID := int32(56) // int32 | 
	ezsigndocumentExtractTextV1Request := *openapiclient.NewEzsigndocumentExtractTextV1Request(int32(1)) // EzsigndocumentExtractTextV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigndocumentAPI.EzsigndocumentExtractTextV1(context.Background(), pkiEzsigndocumentID).EzsigndocumentExtractTextV1Request(ezsigndocumentExtractTextV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigndocumentAPI.EzsigndocumentExtractTextV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigndocumentExtractTextV1`: EzsigndocumentExtractTextV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigndocumentAPI.EzsigndocumentExtractTextV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigndocumentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigndocumentExtractTextV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezsigndocumentExtractTextV1Request** | [**EzsigndocumentExtractTextV1Request**](EzsigndocumentExtractTextV1Request.md) |  | 

### Return type

[**EzsigndocumentExtractTextV1Response**](EzsigndocumentExtractTextV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigndocumentFlattenV1

> CommonResponse EzsigndocumentFlattenV1(ctx, pkiEzsigndocumentID).Body(body).Execute()

Flatten



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ezmaxinc/ezmax-sdk-go"
)

func main() {
	pkiEzsigndocumentID := int32(56) // int32 | 
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigndocumentAPI.EzsigndocumentFlattenV1(context.Background(), pkiEzsigndocumentID).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigndocumentAPI.EzsigndocumentFlattenV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigndocumentFlattenV1`: CommonResponse
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigndocumentAPI.EzsigndocumentFlattenV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigndocumentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigndocumentFlattenV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

[**CommonResponse**](CommonResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigndocumentGetActionableElementsV1

> EzsigndocumentGetActionableElementsV1Response EzsigndocumentGetActionableElementsV1(ctx, pkiEzsigndocumentID).Execute()

Retrieve actionable elements for the Ezsigndocument



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ezmaxinc/ezmax-sdk-go"
)

func main() {
	pkiEzsigndocumentID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigndocumentAPI.EzsigndocumentGetActionableElementsV1(context.Background(), pkiEzsigndocumentID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigndocumentAPI.EzsigndocumentGetActionableElementsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigndocumentGetActionableElementsV1`: EzsigndocumentGetActionableElementsV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigndocumentAPI.EzsigndocumentGetActionableElementsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigndocumentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigndocumentGetActionableElementsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsigndocumentGetActionableElementsV1Response**](EzsigndocumentGetActionableElementsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigndocumentGetAttachmentsV1

> EzsigndocumentGetAttachmentsV1Response EzsigndocumentGetAttachmentsV1(ctx, pkiEzsigndocumentID).Execute()

Retrieve Ezsigndocument's Attachments



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ezmaxinc/ezmax-sdk-go"
)

func main() {
	pkiEzsigndocumentID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigndocumentAPI.EzsigndocumentGetAttachmentsV1(context.Background(), pkiEzsigndocumentID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigndocumentAPI.EzsigndocumentGetAttachmentsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigndocumentGetAttachmentsV1`: EzsigndocumentGetAttachmentsV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigndocumentAPI.EzsigndocumentGetAttachmentsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigndocumentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigndocumentGetAttachmentsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsigndocumentGetAttachmentsV1Response**](EzsigndocumentGetAttachmentsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigndocumentGetCompletedElementsV1

> EzsigndocumentGetCompletedElementsV1Response EzsigndocumentGetCompletedElementsV1(ctx, pkiEzsigndocumentID).Execute()

Retrieve completed elements for the Ezsigndocument



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ezmaxinc/ezmax-sdk-go"
)

func main() {
	pkiEzsigndocumentID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigndocumentAPI.EzsigndocumentGetCompletedElementsV1(context.Background(), pkiEzsigndocumentID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigndocumentAPI.EzsigndocumentGetCompletedElementsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigndocumentGetCompletedElementsV1`: EzsigndocumentGetCompletedElementsV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigndocumentAPI.EzsigndocumentGetCompletedElementsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigndocumentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigndocumentGetCompletedElementsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsigndocumentGetCompletedElementsV1Response**](EzsigndocumentGetCompletedElementsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigndocumentGetDownloadUrlV1

> EzsigndocumentGetDownloadUrlV1Response EzsigndocumentGetDownloadUrlV1(ctx, pkiEzsigndocumentID, eDocumentType).Execute()

Retrieve a URL to download documents.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ezmaxinc/ezmax-sdk-go"
)

func main() {
	pkiEzsigndocumentID := int32(56) // int32 | 
	eDocumentType := "eDocumentType_example" // string | The type of document to retrieve.  1. **Initial** Is the initial document before any signature were applied. 2. **SignatureReady** Is the version containing the annotations/form to show the signer. 3. **Signed** Is the final document once all signatures were applied in current document if eEzsignfolderCompletion is PerEzsigndocument.<br>     Is the final document once all signatures were applied in all documents if eEzsignfolderCompletion is PerEzsignfolder. 4. **Proofdocument** Is the evidence report. 5. **Proof** Is the complete evidence archive including all of the above and more. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigndocumentAPI.EzsigndocumentGetDownloadUrlV1(context.Background(), pkiEzsigndocumentID, eDocumentType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigndocumentAPI.EzsigndocumentGetDownloadUrlV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigndocumentGetDownloadUrlV1`: EzsigndocumentGetDownloadUrlV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigndocumentAPI.EzsigndocumentGetDownloadUrlV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigndocumentID** | **int32** |  | 
**eDocumentType** | **string** | The type of document to retrieve.  1. **Initial** Is the initial document before any signature were applied. 2. **SignatureReady** Is the version containing the annotations/form to show the signer. 3. **Signed** Is the final document once all signatures were applied in current document if eEzsignfolderCompletion is PerEzsigndocument.&lt;br&gt;     Is the final document once all signatures were applied in all documents if eEzsignfolderCompletion is PerEzsignfolder. 4. **Proofdocument** Is the evidence report. 5. **Proof** Is the complete evidence archive including all of the above and more.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigndocumentGetDownloadUrlV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EzsigndocumentGetDownloadUrlV1Response**](EzsigndocumentGetDownloadUrlV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigndocumentGetEzsignannotationsV1

> EzsigndocumentGetEzsignannotationsV1Response EzsigndocumentGetEzsignannotationsV1(ctx, pkiEzsigndocumentID).Execute()

Retrieve an existing Ezsigndocument's Ezsignannotations



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ezmaxinc/ezmax-sdk-go"
)

func main() {
	pkiEzsigndocumentID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigndocumentAPI.EzsigndocumentGetEzsignannotationsV1(context.Background(), pkiEzsigndocumentID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigndocumentAPI.EzsigndocumentGetEzsignannotationsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigndocumentGetEzsignannotationsV1`: EzsigndocumentGetEzsignannotationsV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigndocumentAPI.EzsigndocumentGetEzsignannotationsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigndocumentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigndocumentGetEzsignannotationsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsigndocumentGetEzsignannotationsV1Response**](EzsigndocumentGetEzsignannotationsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigndocumentGetEzsigndiscussionsV1

> EzsigndocumentGetEzsigndiscussionsV1Response EzsigndocumentGetEzsigndiscussionsV1(ctx, pkiEzsigndocumentID).Execute()

Retrieve an existing Ezsigndocument's Ezsigndiscussions



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ezmaxinc/ezmax-sdk-go"
)

func main() {
	pkiEzsigndocumentID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigndocumentAPI.EzsigndocumentGetEzsigndiscussionsV1(context.Background(), pkiEzsigndocumentID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigndocumentAPI.EzsigndocumentGetEzsigndiscussionsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigndocumentGetEzsigndiscussionsV1`: EzsigndocumentGetEzsigndiscussionsV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigndocumentAPI.EzsigndocumentGetEzsigndiscussionsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigndocumentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigndocumentGetEzsigndiscussionsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsigndocumentGetEzsigndiscussionsV1Response**](EzsigndocumentGetEzsigndiscussionsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigndocumentGetEzsignformfieldgroupsV1

> EzsigndocumentGetEzsignformfieldgroupsV1Response EzsigndocumentGetEzsignformfieldgroupsV1(ctx, pkiEzsigndocumentID).Execute()

Retrieve an existing Ezsigndocument's Ezsignformfieldgroups



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ezmaxinc/ezmax-sdk-go"
)

func main() {
	pkiEzsigndocumentID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigndocumentAPI.EzsigndocumentGetEzsignformfieldgroupsV1(context.Background(), pkiEzsigndocumentID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigndocumentAPI.EzsigndocumentGetEzsignformfieldgroupsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigndocumentGetEzsignformfieldgroupsV1`: EzsigndocumentGetEzsignformfieldgroupsV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigndocumentAPI.EzsigndocumentGetEzsignformfieldgroupsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigndocumentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigndocumentGetEzsignformfieldgroupsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsigndocumentGetEzsignformfieldgroupsV1Response**](EzsigndocumentGetEzsignformfieldgroupsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigndocumentGetEzsignpagesV1

> EzsigndocumentGetEzsignpagesV1Response EzsigndocumentGetEzsignpagesV1(ctx, pkiEzsigndocumentID).Execute()

Retrieve an existing Ezsigndocument's Ezsignpages



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ezmaxinc/ezmax-sdk-go"
)

func main() {
	pkiEzsigndocumentID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigndocumentAPI.EzsigndocumentGetEzsignpagesV1(context.Background(), pkiEzsigndocumentID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigndocumentAPI.EzsigndocumentGetEzsignpagesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigndocumentGetEzsignpagesV1`: EzsigndocumentGetEzsignpagesV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigndocumentAPI.EzsigndocumentGetEzsignpagesV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigndocumentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigndocumentGetEzsignpagesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsigndocumentGetEzsignpagesV1Response**](EzsigndocumentGetEzsignpagesV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigndocumentGetEzsignsignaturesAutomaticV1

> EzsigndocumentGetEzsignsignaturesAutomaticV1Response EzsigndocumentGetEzsignsignaturesAutomaticV1(ctx, pkiEzsigndocumentID).Execute()

Retrieve an existing Ezsigndocument's automatic Ezsignsignatures



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ezmaxinc/ezmax-sdk-go"
)

func main() {
	pkiEzsigndocumentID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigndocumentAPI.EzsigndocumentGetEzsignsignaturesAutomaticV1(context.Background(), pkiEzsigndocumentID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigndocumentAPI.EzsigndocumentGetEzsignsignaturesAutomaticV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigndocumentGetEzsignsignaturesAutomaticV1`: EzsigndocumentGetEzsignsignaturesAutomaticV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigndocumentAPI.EzsigndocumentGetEzsignsignaturesAutomaticV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigndocumentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigndocumentGetEzsignsignaturesAutomaticV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsigndocumentGetEzsignsignaturesAutomaticV1Response**](EzsigndocumentGetEzsignsignaturesAutomaticV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigndocumentGetEzsignsignaturesV1

> EzsigndocumentGetEzsignsignaturesV1Response EzsigndocumentGetEzsignsignaturesV1(ctx, pkiEzsigndocumentID).Execute()

Retrieve an existing Ezsigndocument's Ezsignsignatures



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ezmaxinc/ezmax-sdk-go"
)

func main() {
	pkiEzsigndocumentID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigndocumentAPI.EzsigndocumentGetEzsignsignaturesV1(context.Background(), pkiEzsigndocumentID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigndocumentAPI.EzsigndocumentGetEzsignsignaturesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigndocumentGetEzsignsignaturesV1`: EzsigndocumentGetEzsignsignaturesV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigndocumentAPI.EzsigndocumentGetEzsignsignaturesV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigndocumentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigndocumentGetEzsignsignaturesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsigndocumentGetEzsignsignaturesV1Response**](EzsigndocumentGetEzsignsignaturesV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigndocumentGetFormDataV1

> EzsigndocumentGetFormDataV1Response EzsigndocumentGetFormDataV1(ctx, pkiEzsigndocumentID).Execute()

Retrieve an existing Ezsigndocument's Form Data



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ezmaxinc/ezmax-sdk-go"
)

func main() {
	pkiEzsigndocumentID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigndocumentAPI.EzsigndocumentGetFormDataV1(context.Background(), pkiEzsigndocumentID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigndocumentAPI.EzsigndocumentGetFormDataV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigndocumentGetFormDataV1`: EzsigndocumentGetFormDataV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigndocumentAPI.EzsigndocumentGetFormDataV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigndocumentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigndocumentGetFormDataV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsigndocumentGetFormDataV1Response**](EzsigndocumentGetFormDataV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/zip, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigndocumentGetObjectV1

> EzsigndocumentGetObjectV1Response EzsigndocumentGetObjectV1(ctx, pkiEzsigndocumentID).Execute()

Retrieve an existing Ezsigndocument

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ezmaxinc/ezmax-sdk-go"
)

func main() {
	pkiEzsigndocumentID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigndocumentAPI.EzsigndocumentGetObjectV1(context.Background(), pkiEzsigndocumentID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigndocumentAPI.EzsigndocumentGetObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigndocumentGetObjectV1`: EzsigndocumentGetObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigndocumentAPI.EzsigndocumentGetObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigndocumentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigndocumentGetObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsigndocumentGetObjectV1Response**](EzsigndocumentGetObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigndocumentGetObjectV2

> EzsigndocumentGetObjectV2Response EzsigndocumentGetObjectV2(ctx, pkiEzsigndocumentID).Execute()

Retrieve an existing Ezsigndocument



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ezmaxinc/ezmax-sdk-go"
)

func main() {
	pkiEzsigndocumentID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigndocumentAPI.EzsigndocumentGetObjectV2(context.Background(), pkiEzsigndocumentID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigndocumentAPI.EzsigndocumentGetObjectV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigndocumentGetObjectV2`: EzsigndocumentGetObjectV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigndocumentAPI.EzsigndocumentGetObjectV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigndocumentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigndocumentGetObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsigndocumentGetObjectV2Response**](EzsigndocumentGetObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigndocumentGetTemporaryProofV1

> EzsigndocumentGetTemporaryProofV1Response EzsigndocumentGetTemporaryProofV1(ctx, pkiEzsigndocumentID).Execute()

Retrieve the temporary proof



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ezmaxinc/ezmax-sdk-go"
)

func main() {
	pkiEzsigndocumentID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigndocumentAPI.EzsigndocumentGetTemporaryProofV1(context.Background(), pkiEzsigndocumentID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigndocumentAPI.EzsigndocumentGetTemporaryProofV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigndocumentGetTemporaryProofV1`: EzsigndocumentGetTemporaryProofV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigndocumentAPI.EzsigndocumentGetTemporaryProofV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigndocumentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigndocumentGetTemporaryProofV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsigndocumentGetTemporaryProofV1Response**](EzsigndocumentGetTemporaryProofV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigndocumentGetWordsPositionsV1

> EzsigndocumentGetWordsPositionsV1Response EzsigndocumentGetWordsPositionsV1(ctx, pkiEzsigndocumentID).EzsigndocumentGetWordsPositionsV1Request(ezsigndocumentGetWordsPositionsV1Request).Execute()

Retrieve positions X,Y of given words from a Ezsigndocument



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ezmaxinc/ezmax-sdk-go"
)

func main() {
	pkiEzsigndocumentID := int32(56) // int32 | 
	ezsigndocumentGetWordsPositionsV1Request := *openapiclient.NewEzsigndocumentGetWordsPositionsV1Request("EGet_example", false) // EzsigndocumentGetWordsPositionsV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigndocumentAPI.EzsigndocumentGetWordsPositionsV1(context.Background(), pkiEzsigndocumentID).EzsigndocumentGetWordsPositionsV1Request(ezsigndocumentGetWordsPositionsV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigndocumentAPI.EzsigndocumentGetWordsPositionsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigndocumentGetWordsPositionsV1`: EzsigndocumentGetWordsPositionsV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigndocumentAPI.EzsigndocumentGetWordsPositionsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigndocumentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigndocumentGetWordsPositionsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezsigndocumentGetWordsPositionsV1Request** | [**EzsigndocumentGetWordsPositionsV1Request**](EzsigndocumentGetWordsPositionsV1Request.md) |  | 

### Return type

[**EzsigndocumentGetWordsPositionsV1Response**](EzsigndocumentGetWordsPositionsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigndocumentPatchObjectV1

> CommonResponse EzsigndocumentPatchObjectV1(ctx, pkiEzsigndocumentID).EzsigndocumentPatchObjectV1Request(ezsigndocumentPatchObjectV1Request).Execute()

Patch an existing Ezsigndocument



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ezmaxinc/ezmax-sdk-go"
)

func main() {
	pkiEzsigndocumentID := int32(56) // int32 | 
	ezsigndocumentPatchObjectV1Request := *openapiclient.NewEzsigndocumentPatchObjectV1Request(*openapiclient.NewEzsigndocumentRequestPatch()) // EzsigndocumentPatchObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigndocumentAPI.EzsigndocumentPatchObjectV1(context.Background(), pkiEzsigndocumentID).EzsigndocumentPatchObjectV1Request(ezsigndocumentPatchObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigndocumentAPI.EzsigndocumentPatchObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigndocumentPatchObjectV1`: CommonResponse
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigndocumentAPI.EzsigndocumentPatchObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigndocumentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigndocumentPatchObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezsigndocumentPatchObjectV1Request** | [**EzsigndocumentPatchObjectV1Request**](EzsigndocumentPatchObjectV1Request.md) |  | 

### Return type

[**CommonResponse**](CommonResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigndocumentSubmitEzsignformV1

> CommonResponse EzsigndocumentSubmitEzsignformV1(ctx, pkiEzsigndocumentID).EzsigndocumentSubmitEzsignformV1Request(ezsigndocumentSubmitEzsignformV1Request).Execute()

Submit the Ezsignform



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ezmaxinc/ezmax-sdk-go"
)

func main() {
	pkiEzsigndocumentID := int32(56) // int32 | 
	ezsigndocumentSubmitEzsignformV1Request := *openapiclient.NewEzsigndocumentSubmitEzsignformV1Request(false, []openapiclient.CustomEzsignformfieldgroupRequest{map[string]interface{}(123)}) // EzsigndocumentSubmitEzsignformV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigndocumentAPI.EzsigndocumentSubmitEzsignformV1(context.Background(), pkiEzsigndocumentID).EzsigndocumentSubmitEzsignformV1Request(ezsigndocumentSubmitEzsignformV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigndocumentAPI.EzsigndocumentSubmitEzsignformV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigndocumentSubmitEzsignformV1`: CommonResponse
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigndocumentAPI.EzsigndocumentSubmitEzsignformV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigndocumentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigndocumentSubmitEzsignformV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezsigndocumentSubmitEzsignformV1Request** | [**EzsigndocumentSubmitEzsignformV1Request**](EzsigndocumentSubmitEzsignformV1Request.md) |  | 

### Return type

[**CommonResponse**](CommonResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigndocumentUnsendV1

> CommonResponse EzsigndocumentUnsendV1(ctx, pkiEzsigndocumentID).Body(body).Execute()

Unsend the Ezsigndocument



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ezmaxinc/ezmax-sdk-go"
)

func main() {
	pkiEzsigndocumentID := int32(56) // int32 | 
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigndocumentAPI.EzsigndocumentUnsendV1(context.Background(), pkiEzsigndocumentID).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigndocumentAPI.EzsigndocumentUnsendV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigndocumentUnsendV1`: CommonResponse
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigndocumentAPI.EzsigndocumentUnsendV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigndocumentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigndocumentUnsendV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

[**CommonResponse**](CommonResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

