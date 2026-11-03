# eZmaxAPI\ObjectDocumenttypechecklistAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DocumenttypechecklistDeleteObjectV1**](ObjectDocumenttypechecklistAPI.md#DocumenttypechecklistDeleteObjectV1) | **Delete** /1/object/documenttypechecklist/{pkiDocumenttypechecklistID} | Delete an existing Documenttypechecklist
[**DocumenttypechecklistPatchObjectV1**](ObjectDocumenttypechecklistAPI.md#DocumenttypechecklistPatchObjectV1) | **Patch** /1/object/documenttypechecklist/{pkiDocumenttypechecklistID} | Patch an existing Documenttypechecklist



## DocumenttypechecklistDeleteObjectV1

> DocumenttypechecklistDeleteObjectV1Response DocumenttypechecklistDeleteObjectV1(ctx, pkiDocumenttypechecklistID).Execute()

Delete an existing Documenttypechecklist



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
	pkiDocumenttypechecklistID := int32(56) // int32 | The unique ID of the Documenttypechecklist

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectDocumenttypechecklistAPI.DocumenttypechecklistDeleteObjectV1(context.Background(), pkiDocumenttypechecklistID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectDocumenttypechecklistAPI.DocumenttypechecklistDeleteObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DocumenttypechecklistDeleteObjectV1`: DocumenttypechecklistDeleteObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectDocumenttypechecklistAPI.DocumenttypechecklistDeleteObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiDocumenttypechecklistID** | **int32** | The unique ID of the Documenttypechecklist | 

### Other Parameters

Other parameters are passed through a pointer to a apiDocumenttypechecklistDeleteObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DocumenttypechecklistDeleteObjectV1Response**](DocumenttypechecklistDeleteObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DocumenttypechecklistPatchObjectV1

> DocumenttypechecklistPatchObjectV1Response DocumenttypechecklistPatchObjectV1(ctx, pkiDocumenttypechecklistID).DocumenttypechecklistPatchObjectV1Request(documenttypechecklistPatchObjectV1Request).Execute()

Patch an existing Documenttypechecklist



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
	pkiDocumenttypechecklistID := int32(56) // int32 | The unique ID of the Documenttypechecklist
	documenttypechecklistPatchObjectV1Request := *openapiclient.NewDocumenttypechecklistPatchObjectV1Request(*openapiclient.NewDocumenttypechecklistRequestPatch()) // DocumenttypechecklistPatchObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectDocumenttypechecklistAPI.DocumenttypechecklistPatchObjectV1(context.Background(), pkiDocumenttypechecklistID).DocumenttypechecklistPatchObjectV1Request(documenttypechecklistPatchObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectDocumenttypechecklistAPI.DocumenttypechecklistPatchObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DocumenttypechecklistPatchObjectV1`: DocumenttypechecklistPatchObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectDocumenttypechecklistAPI.DocumenttypechecklistPatchObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiDocumenttypechecklistID** | **int32** | The unique ID of the Documenttypechecklist | 

### Other Parameters

Other parameters are passed through a pointer to a apiDocumenttypechecklistPatchObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **documenttypechecklistPatchObjectV1Request** | [**DocumenttypechecklistPatchObjectV1Request**](DocumenttypechecklistPatchObjectV1Request.md) |  | 

### Return type

[**DocumenttypechecklistPatchObjectV1Response**](DocumenttypechecklistPatchObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

