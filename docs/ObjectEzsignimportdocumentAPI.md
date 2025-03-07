# eZmaxAPI\ObjectEzsignimportdocumentAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EzsignimportdocumentDownloadV1**](ObjectEzsignimportdocumentAPI.md#EzsignimportdocumentDownloadV1) | **Get** /1/object/ezsignimportdocument/{pkiEzsignimportdocumentID}/download | Retrieve the content



## EzsignimportdocumentDownloadV1

> EzsignimportdocumentDownloadV1Response EzsignimportdocumentDownloadV1(ctx, pkiEzsignimportdocumentID).Execute()

Retrieve the content

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
	pkiEzsignimportdocumentID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsignimportdocumentAPI.EzsignimportdocumentDownloadV1(context.Background(), pkiEzsignimportdocumentID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignimportdocumentAPI.EzsignimportdocumentDownloadV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsignimportdocumentDownloadV1`: EzsignimportdocumentDownloadV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignimportdocumentAPI.EzsignimportdocumentDownloadV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignimportdocumentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignimportdocumentDownloadV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsignimportdocumentDownloadV1Response**](EzsignimportdocumentDownloadV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

