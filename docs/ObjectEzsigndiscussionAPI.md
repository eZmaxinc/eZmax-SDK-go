# eZmaxAPI\ObjectEzsigndiscussionAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EzsigndiscussionCreateObjectV1**](ObjectEzsigndiscussionAPI.md#EzsigndiscussionCreateObjectV1) | **Post** /1/object/ezsigndiscussion | Create a new Ezsigndiscussion
[**EzsigndiscussionDeleteObjectV1**](ObjectEzsigndiscussionAPI.md#EzsigndiscussionDeleteObjectV1) | **Delete** /1/object/ezsigndiscussion/{pkiEzsigndiscussionID} | Delete an existing Ezsigndiscussion
[**EzsigndiscussionGetObjectV2**](ObjectEzsigndiscussionAPI.md#EzsigndiscussionGetObjectV2) | **Get** /2/object/ezsigndiscussion/{pkiEzsigndiscussionID} | Retrieve an existing Ezsigndiscussion



## EzsigndiscussionCreateObjectV1

> EzsigndiscussionCreateObjectV1Response EzsigndiscussionCreateObjectV1(ctx).EzsigndiscussionCreateObjectV1Request(ezsigndiscussionCreateObjectV1Request).Execute()

Create a new Ezsigndiscussion



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
	ezsigndiscussionCreateObjectV1Request := *openapiclient.NewEzsigndiscussionCreateObjectV1Request([]openapiclient.EzsigndiscussionRequestCompound{*openapiclient.NewEzsigndiscussionRequestCompound(int32(97), int32(4), int32(57208), int32(57652), *openapiclient.NewDiscussionRequest("John Doe"))}) // EzsigndiscussionCreateObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigndiscussionAPI.EzsigndiscussionCreateObjectV1(context.Background()).EzsigndiscussionCreateObjectV1Request(ezsigndiscussionCreateObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigndiscussionAPI.EzsigndiscussionCreateObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigndiscussionCreateObjectV1`: EzsigndiscussionCreateObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigndiscussionAPI.EzsigndiscussionCreateObjectV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEzsigndiscussionCreateObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ezsigndiscussionCreateObjectV1Request** | [**EzsigndiscussionCreateObjectV1Request**](EzsigndiscussionCreateObjectV1Request.md) |  | 

### Return type

[**EzsigndiscussionCreateObjectV1Response**](EzsigndiscussionCreateObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigndiscussionDeleteObjectV1

> EzsigndiscussionDeleteObjectV1Response EzsigndiscussionDeleteObjectV1(ctx, pkiEzsigndiscussionID).Execute()

Delete an existing Ezsigndiscussion



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
	pkiEzsigndiscussionID := int32(56) // int32 | The unique ID of the Ezsigndiscussion

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigndiscussionAPI.EzsigndiscussionDeleteObjectV1(context.Background(), pkiEzsigndiscussionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigndiscussionAPI.EzsigndiscussionDeleteObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigndiscussionDeleteObjectV1`: EzsigndiscussionDeleteObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigndiscussionAPI.EzsigndiscussionDeleteObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigndiscussionID** | **int32** | The unique ID of the Ezsigndiscussion | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigndiscussionDeleteObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsigndiscussionDeleteObjectV1Response**](EzsigndiscussionDeleteObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigndiscussionGetObjectV2

> EzsigndiscussionGetObjectV2Response EzsigndiscussionGetObjectV2(ctx, pkiEzsigndiscussionID).Execute()

Retrieve an existing Ezsigndiscussion



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
	pkiEzsigndiscussionID := int32(56) // int32 | The unique ID of the Ezsigndiscussion

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigndiscussionAPI.EzsigndiscussionGetObjectV2(context.Background(), pkiEzsigndiscussionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigndiscussionAPI.EzsigndiscussionGetObjectV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigndiscussionGetObjectV2`: EzsigndiscussionGetObjectV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigndiscussionAPI.EzsigndiscussionGetObjectV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigndiscussionID** | **int32** | The unique ID of the Ezsigndiscussion | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigndiscussionGetObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsigndiscussionGetObjectV2Response**](EzsigndiscussionGetObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

