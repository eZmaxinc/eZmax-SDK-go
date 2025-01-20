# eZmaxAPI\ObjectDiscussionAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DiscussionCreateObjectV1**](ObjectDiscussionAPI.md#DiscussionCreateObjectV1) | **Post** /1/object/discussion | Create a new Discussion
[**DiscussionDeleteObjectV1**](ObjectDiscussionAPI.md#DiscussionDeleteObjectV1) | **Delete** /1/object/discussion/{pkiDiscussionID} | Delete an existing Discussion
[**DiscussionGetObjectV2**](ObjectDiscussionAPI.md#DiscussionGetObjectV2) | **Get** /2/object/discussion/{pkiDiscussionID} | Retrieve an existing Discussion
[**DiscussionPatchObjectV1**](ObjectDiscussionAPI.md#DiscussionPatchObjectV1) | **Patch** /1/object/discussion/{pkiDiscussionID} | Patch an existing Discussion
[**DiscussionUpdateDiscussionreadstatusV1**](ObjectDiscussionAPI.md#DiscussionUpdateDiscussionreadstatusV1) | **Post** /1/object/discussion/{pkiDiscussionID}/updateDiscussionreadstatus | Update the read status of the discussion



## DiscussionCreateObjectV1

> DiscussionCreateObjectV1Response DiscussionCreateObjectV1(ctx).DiscussionCreateObjectV1Request(discussionCreateObjectV1Request).Execute()

Create a new Discussion



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
	discussionCreateObjectV1Request := *openapiclient.NewDiscussionCreateObjectV1Request([]openapiclient.DiscussionRequestCompound{*openapiclient.NewDiscussionRequestCompound("John Doe")}) // DiscussionCreateObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectDiscussionAPI.DiscussionCreateObjectV1(context.Background()).DiscussionCreateObjectV1Request(discussionCreateObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectDiscussionAPI.DiscussionCreateObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscussionCreateObjectV1`: DiscussionCreateObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectDiscussionAPI.DiscussionCreateObjectV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDiscussionCreateObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **discussionCreateObjectV1Request** | [**DiscussionCreateObjectV1Request**](DiscussionCreateObjectV1Request.md) |  | 

### Return type

[**DiscussionCreateObjectV1Response**](DiscussionCreateObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiscussionDeleteObjectV1

> CommonResponse DiscussionDeleteObjectV1(ctx, pkiDiscussionID).Execute()

Delete an existing Discussion



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
	pkiDiscussionID := int32(56) // int32 | The unique ID of the Discussion

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectDiscussionAPI.DiscussionDeleteObjectV1(context.Background(), pkiDiscussionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectDiscussionAPI.DiscussionDeleteObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscussionDeleteObjectV1`: CommonResponse
	fmt.Fprintf(os.Stdout, "Response from `ObjectDiscussionAPI.DiscussionDeleteObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiDiscussionID** | **int32** | The unique ID of the Discussion | 

### Other Parameters

Other parameters are passed through a pointer to a apiDiscussionDeleteObjectV1Request struct via the builder pattern


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


## DiscussionGetObjectV2

> DiscussionGetObjectV2Response DiscussionGetObjectV2(ctx, pkiDiscussionID).Execute()

Retrieve an existing Discussion



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
	pkiDiscussionID := int32(56) // int32 | The unique ID of the Discussion

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectDiscussionAPI.DiscussionGetObjectV2(context.Background(), pkiDiscussionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectDiscussionAPI.DiscussionGetObjectV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscussionGetObjectV2`: DiscussionGetObjectV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectDiscussionAPI.DiscussionGetObjectV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiDiscussionID** | **int32** | The unique ID of the Discussion | 

### Other Parameters

Other parameters are passed through a pointer to a apiDiscussionGetObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DiscussionGetObjectV2Response**](DiscussionGetObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiscussionPatchObjectV1

> CommonResponse DiscussionPatchObjectV1(ctx, pkiDiscussionID).DiscussionPatchObjectV1Request(discussionPatchObjectV1Request).Execute()

Patch an existing Discussion



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
	pkiDiscussionID := int32(56) // int32 | The unique ID of the Discussion
	discussionPatchObjectV1Request := *openapiclient.NewDiscussionPatchObjectV1Request(*openapiclient.NewDiscussionRequestPatch()) // DiscussionPatchObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectDiscussionAPI.DiscussionPatchObjectV1(context.Background(), pkiDiscussionID).DiscussionPatchObjectV1Request(discussionPatchObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectDiscussionAPI.DiscussionPatchObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscussionPatchObjectV1`: CommonResponse
	fmt.Fprintf(os.Stdout, "Response from `ObjectDiscussionAPI.DiscussionPatchObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiDiscussionID** | **int32** | The unique ID of the Discussion | 

### Other Parameters

Other parameters are passed through a pointer to a apiDiscussionPatchObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **discussionPatchObjectV1Request** | [**DiscussionPatchObjectV1Request**](DiscussionPatchObjectV1Request.md) |  | 

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


## DiscussionUpdateDiscussionreadstatusV1

> CommonResponse DiscussionUpdateDiscussionreadstatusV1(ctx, pkiDiscussionID).DiscussionUpdateDiscussionreadstatusV1Request(discussionUpdateDiscussionreadstatusV1Request).Execute()

Update the read status of the discussion

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
	pkiDiscussionID := int32(56) // int32 | 
	discussionUpdateDiscussionreadstatusV1Request := *openapiclient.NewDiscussionUpdateDiscussionreadstatusV1Request() // DiscussionUpdateDiscussionreadstatusV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectDiscussionAPI.DiscussionUpdateDiscussionreadstatusV1(context.Background(), pkiDiscussionID).DiscussionUpdateDiscussionreadstatusV1Request(discussionUpdateDiscussionreadstatusV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectDiscussionAPI.DiscussionUpdateDiscussionreadstatusV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscussionUpdateDiscussionreadstatusV1`: CommonResponse
	fmt.Fprintf(os.Stdout, "Response from `ObjectDiscussionAPI.DiscussionUpdateDiscussionreadstatusV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiDiscussionID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDiscussionUpdateDiscussionreadstatusV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **discussionUpdateDiscussionreadstatusV1Request** | [**DiscussionUpdateDiscussionreadstatusV1Request**](DiscussionUpdateDiscussionreadstatusV1Request.md) |  | 

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

