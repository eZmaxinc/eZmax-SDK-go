# eZmaxAPI\ObjectExternalbrokerAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExternalbrokerImportIntoEDMV1**](ObjectExternalbrokerAPI.md#ExternalbrokerImportIntoEDMV1) | **Post** /1/object/externalbroker/{pkiExternalbrokerID}/importIntoEDM | Import attachments into the Externalbroker



## ExternalbrokerImportIntoEDMV1

> ExternalbrokerImportIntoEDMV1Response ExternalbrokerImportIntoEDMV1(ctx, pkiExternalbrokerID).ExternalbrokerImportIntoEDMV1Request(externalbrokerImportIntoEDMV1Request).Execute()

Import attachments into the Externalbroker



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
	pkiExternalbrokerID := int32(56) // int32 | 
	externalbrokerImportIntoEDMV1Request := *openapiclient.NewExternalbrokerImportIntoEDMV1Request([]openapiclient.CustomAttachmentImportIntoEDMRequest{*openapiclient.NewCustomAttachmentImportIntoEDMRequest("EAttachmentSource_example", "Document.pdf", "Inscription", openapiclient.Field-eAttachmentPrivacy("All"))}) // ExternalbrokerImportIntoEDMV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectExternalbrokerAPI.ExternalbrokerImportIntoEDMV1(context.Background(), pkiExternalbrokerID).ExternalbrokerImportIntoEDMV1Request(externalbrokerImportIntoEDMV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectExternalbrokerAPI.ExternalbrokerImportIntoEDMV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExternalbrokerImportIntoEDMV1`: ExternalbrokerImportIntoEDMV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectExternalbrokerAPI.ExternalbrokerImportIntoEDMV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiExternalbrokerID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExternalbrokerImportIntoEDMV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **externalbrokerImportIntoEDMV1Request** | [**ExternalbrokerImportIntoEDMV1Request**](ExternalbrokerImportIntoEDMV1Request.md) |  | 

### Return type

[**ExternalbrokerImportIntoEDMV1Response**](ExternalbrokerImportIntoEDMV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

