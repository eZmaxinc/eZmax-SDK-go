# eZmaxAPI\ObjectEzsigntemplatepackagesignerAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EzsigntemplatepackagesignerCreateObjectV1**](ObjectEzsigntemplatepackagesignerAPI.md#EzsigntemplatepackagesignerCreateObjectV1) | **Post** /1/object/ezsigntemplatepackagesigner | Create a new Ezsigntemplatepackagesigner
[**EzsigntemplatepackagesignerDeleteObjectV1**](ObjectEzsigntemplatepackagesignerAPI.md#EzsigntemplatepackagesignerDeleteObjectV1) | **Delete** /1/object/ezsigntemplatepackagesigner/{pkiEzsigntemplatepackagesignerID} | Delete an existing Ezsigntemplatepackagesigner
[**EzsigntemplatepackagesignerEditObjectV1**](ObjectEzsigntemplatepackagesignerAPI.md#EzsigntemplatepackagesignerEditObjectV1) | **Put** /1/object/ezsigntemplatepackagesigner/{pkiEzsigntemplatepackagesignerID} | Edit an existing Ezsigntemplatepackagesigner
[**EzsigntemplatepackagesignerGetObjectV2**](ObjectEzsigntemplatepackagesignerAPI.md#EzsigntemplatepackagesignerGetObjectV2) | **Get** /2/object/ezsigntemplatepackagesigner/{pkiEzsigntemplatepackagesignerID} | Retrieve an existing Ezsigntemplatepackagesigner



## EzsigntemplatepackagesignerCreateObjectV1

> EzsigntemplatepackagesignerCreateObjectV1Response EzsigntemplatepackagesignerCreateObjectV1(ctx).EzsigntemplatepackagesignerCreateObjectV1Request(ezsigntemplatepackagesignerCreateObjectV1Request).Execute()

Create a new Ezsigntemplatepackagesigner



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
	ezsigntemplatepackagesignerCreateObjectV1Request := *openapiclient.NewEzsigntemplatepackagesignerCreateObjectV1Request([]openapiclient.EzsigntemplatepackagesignerRequestCompound{*openapiclient.NewEzsigntemplatepackagesignerRequestCompound(int32(99), "Customer")}) // EzsigntemplatepackagesignerCreateObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigntemplatepackagesignerAPI.EzsigntemplatepackagesignerCreateObjectV1(context.Background()).EzsigntemplatepackagesignerCreateObjectV1Request(ezsigntemplatepackagesignerCreateObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplatepackagesignerAPI.EzsigntemplatepackagesignerCreateObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigntemplatepackagesignerCreateObjectV1`: EzsigntemplatepackagesignerCreateObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplatepackagesignerAPI.EzsigntemplatepackagesignerCreateObjectV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplatepackagesignerCreateObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ezsigntemplatepackagesignerCreateObjectV1Request** | [**EzsigntemplatepackagesignerCreateObjectV1Request**](EzsigntemplatepackagesignerCreateObjectV1Request.md) |  | 

### Return type

[**EzsigntemplatepackagesignerCreateObjectV1Response**](EzsigntemplatepackagesignerCreateObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigntemplatepackagesignerDeleteObjectV1

> EzsigntemplatepackagesignerDeleteObjectV1Response EzsigntemplatepackagesignerDeleteObjectV1(ctx, pkiEzsigntemplatepackagesignerID).Execute()

Delete an existing Ezsigntemplatepackagesigner



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
	pkiEzsigntemplatepackagesignerID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigntemplatepackagesignerAPI.EzsigntemplatepackagesignerDeleteObjectV1(context.Background(), pkiEzsigntemplatepackagesignerID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplatepackagesignerAPI.EzsigntemplatepackagesignerDeleteObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigntemplatepackagesignerDeleteObjectV1`: EzsigntemplatepackagesignerDeleteObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplatepackagesignerAPI.EzsigntemplatepackagesignerDeleteObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigntemplatepackagesignerID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplatepackagesignerDeleteObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsigntemplatepackagesignerDeleteObjectV1Response**](EzsigntemplatepackagesignerDeleteObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigntemplatepackagesignerEditObjectV1

> CommonResponse EzsigntemplatepackagesignerEditObjectV1(ctx, pkiEzsigntemplatepackagesignerID).EzsigntemplatepackagesignerEditObjectV1Request(ezsigntemplatepackagesignerEditObjectV1Request).Execute()

Edit an existing Ezsigntemplatepackagesigner



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
	pkiEzsigntemplatepackagesignerID := int32(56) // int32 | 
	ezsigntemplatepackagesignerEditObjectV1Request := *openapiclient.NewEzsigntemplatepackagesignerEditObjectV1Request(*openapiclient.NewEzsigntemplatepackagesignerRequestCompound(int32(99), "Customer")) // EzsigntemplatepackagesignerEditObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigntemplatepackagesignerAPI.EzsigntemplatepackagesignerEditObjectV1(context.Background(), pkiEzsigntemplatepackagesignerID).EzsigntemplatepackagesignerEditObjectV1Request(ezsigntemplatepackagesignerEditObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplatepackagesignerAPI.EzsigntemplatepackagesignerEditObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigntemplatepackagesignerEditObjectV1`: CommonResponse
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplatepackagesignerAPI.EzsigntemplatepackagesignerEditObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigntemplatepackagesignerID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplatepackagesignerEditObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezsigntemplatepackagesignerEditObjectV1Request** | [**EzsigntemplatepackagesignerEditObjectV1Request**](EzsigntemplatepackagesignerEditObjectV1Request.md) |  | 

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


## EzsigntemplatepackagesignerGetObjectV2

> EzsigntemplatepackagesignerGetObjectV2Response EzsigntemplatepackagesignerGetObjectV2(ctx, pkiEzsigntemplatepackagesignerID).Execute()

Retrieve an existing Ezsigntemplatepackagesigner



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
	pkiEzsigntemplatepackagesignerID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigntemplatepackagesignerAPI.EzsigntemplatepackagesignerGetObjectV2(context.Background(), pkiEzsigntemplatepackagesignerID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplatepackagesignerAPI.EzsigntemplatepackagesignerGetObjectV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigntemplatepackagesignerGetObjectV2`: EzsigntemplatepackagesignerGetObjectV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplatepackagesignerAPI.EzsigntemplatepackagesignerGetObjectV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigntemplatepackagesignerID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplatepackagesignerGetObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsigntemplatepackagesignerGetObjectV2Response**](EzsigntemplatepackagesignerGetObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

