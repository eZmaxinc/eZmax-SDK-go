# eZmaxAPI\ObjectEzsignbulksenddocumentmappingAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EzsignbulksenddocumentmappingCreateObjectV1**](ObjectEzsignbulksenddocumentmappingAPI.md#EzsignbulksenddocumentmappingCreateObjectV1) | **Post** /1/object/ezsignbulksenddocumentmapping | Create a new Ezsignbulksenddocumentmapping
[**EzsignbulksenddocumentmappingDeleteObjectV1**](ObjectEzsignbulksenddocumentmappingAPI.md#EzsignbulksenddocumentmappingDeleteObjectV1) | **Delete** /1/object/ezsignbulksenddocumentmapping/{pkiEzsignbulksenddocumentmappingID} | Delete an existing Ezsignbulksenddocumentmapping
[**EzsignbulksenddocumentmappingGetObjectV2**](ObjectEzsignbulksenddocumentmappingAPI.md#EzsignbulksenddocumentmappingGetObjectV2) | **Get** /2/object/ezsignbulksenddocumentmapping/{pkiEzsignbulksenddocumentmappingID} | Retrieve an existing Ezsignbulksenddocumentmapping



## EzsignbulksenddocumentmappingCreateObjectV1

> EzsignbulksenddocumentmappingCreateObjectV1Response EzsignbulksenddocumentmappingCreateObjectV1(ctx).EzsignbulksenddocumentmappingCreateObjectV1Request(ezsignbulksenddocumentmappingCreateObjectV1Request).Execute()

Create a new Ezsignbulksenddocumentmapping



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
	ezsignbulksenddocumentmappingCreateObjectV1Request := *openapiclient.NewEzsignbulksenddocumentmappingCreateObjectV1Request([]openapiclient.EzsignbulksenddocumentmappingRequestCompound{*openapiclient.NewEzsignbulksenddocumentmappingRequestCompound(int32(8))}) // EzsignbulksenddocumentmappingCreateObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsignbulksenddocumentmappingAPI.EzsignbulksenddocumentmappingCreateObjectV1(context.Background()).EzsignbulksenddocumentmappingCreateObjectV1Request(ezsignbulksenddocumentmappingCreateObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignbulksenddocumentmappingAPI.EzsignbulksenddocumentmappingCreateObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsignbulksenddocumentmappingCreateObjectV1`: EzsignbulksenddocumentmappingCreateObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignbulksenddocumentmappingAPI.EzsignbulksenddocumentmappingCreateObjectV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEzsignbulksenddocumentmappingCreateObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ezsignbulksenddocumentmappingCreateObjectV1Request** | [**EzsignbulksenddocumentmappingCreateObjectV1Request**](EzsignbulksenddocumentmappingCreateObjectV1Request.md) |  | 

### Return type

[**EzsignbulksenddocumentmappingCreateObjectV1Response**](EzsignbulksenddocumentmappingCreateObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignbulksenddocumentmappingDeleteObjectV1

> EzsignbulksenddocumentmappingDeleteObjectV1Response EzsignbulksenddocumentmappingDeleteObjectV1(ctx, pkiEzsignbulksenddocumentmappingID).Execute()

Delete an existing Ezsignbulksenddocumentmapping



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
	pkiEzsignbulksenddocumentmappingID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsignbulksenddocumentmappingAPI.EzsignbulksenddocumentmappingDeleteObjectV1(context.Background(), pkiEzsignbulksenddocumentmappingID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignbulksenddocumentmappingAPI.EzsignbulksenddocumentmappingDeleteObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsignbulksenddocumentmappingDeleteObjectV1`: EzsignbulksenddocumentmappingDeleteObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignbulksenddocumentmappingAPI.EzsignbulksenddocumentmappingDeleteObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignbulksenddocumentmappingID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignbulksenddocumentmappingDeleteObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsignbulksenddocumentmappingDeleteObjectV1Response**](EzsignbulksenddocumentmappingDeleteObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignbulksenddocumentmappingGetObjectV2

> EzsignbulksenddocumentmappingGetObjectV2Response EzsignbulksenddocumentmappingGetObjectV2(ctx, pkiEzsignbulksenddocumentmappingID).Execute()

Retrieve an existing Ezsignbulksenddocumentmapping



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
	pkiEzsignbulksenddocumentmappingID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsignbulksenddocumentmappingAPI.EzsignbulksenddocumentmappingGetObjectV2(context.Background(), pkiEzsignbulksenddocumentmappingID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignbulksenddocumentmappingAPI.EzsignbulksenddocumentmappingGetObjectV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsignbulksenddocumentmappingGetObjectV2`: EzsignbulksenddocumentmappingGetObjectV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignbulksenddocumentmappingAPI.EzsignbulksenddocumentmappingGetObjectV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignbulksenddocumentmappingID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignbulksenddocumentmappingGetObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsignbulksenddocumentmappingGetObjectV2Response**](EzsignbulksenddocumentmappingGetObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

