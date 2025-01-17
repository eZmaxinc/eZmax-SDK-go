# eZmaxAPI\ObjectEzsignformfieldgroupAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EzsignformfieldgroupCreateObjectV1**](ObjectEzsignformfieldgroupAPI.md#EzsignformfieldgroupCreateObjectV1) | **Post** /1/object/ezsignformfieldgroup | Create a new Ezsignformfieldgroup
[**EzsignformfieldgroupDeleteObjectV1**](ObjectEzsignformfieldgroupAPI.md#EzsignformfieldgroupDeleteObjectV1) | **Delete** /1/object/ezsignformfieldgroup/{pkiEzsignformfieldgroupID} | Delete an existing Ezsignformfieldgroup
[**EzsignformfieldgroupEditObjectV1**](ObjectEzsignformfieldgroupAPI.md#EzsignformfieldgroupEditObjectV1) | **Put** /1/object/ezsignformfieldgroup/{pkiEzsignformfieldgroupID} | Edit an existing Ezsignformfieldgroup
[**EzsignformfieldgroupGetObjectV2**](ObjectEzsignformfieldgroupAPI.md#EzsignformfieldgroupGetObjectV2) | **Get** /2/object/ezsignformfieldgroup/{pkiEzsignformfieldgroupID} | Retrieve an existing Ezsignformfieldgroup



## EzsignformfieldgroupCreateObjectV1

> EzsignformfieldgroupCreateObjectV1Response EzsignformfieldgroupCreateObjectV1(ctx).EzsignformfieldgroupCreateObjectV1Request(ezsignformfieldgroupCreateObjectV1Request).Execute()

Create a new Ezsignformfieldgroup



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
	ezsignformfieldgroupCreateObjectV1Request := *openapiclient.NewEzsignformfieldgroupCreateObjectV1Request([]openapiclient.EzsignformfieldgroupRequestCompound{*openapiclient.NewEzsignformfieldgroupRequestCompound([]openapiclient.EzsignformfieldgroupsignerRequestCompound{*openapiclient.NewEzsignformfieldgroupsignerRequest(int32(20))}, []openapiclient.EzsignformfieldRequestCompound{*openapiclient.NewEzsignformfieldRequestCompound(int32(1), "Peanuts", int32(200), int32(300), int32(102), int32(22))}, int32(97), openapiclient.Field-eEzsignformfieldgroupType("Text"), "Allergies", int32(1), int32(1), int32(2), false)}) // EzsignformfieldgroupCreateObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsignformfieldgroupAPI.EzsignformfieldgroupCreateObjectV1(context.Background()).EzsignformfieldgroupCreateObjectV1Request(ezsignformfieldgroupCreateObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignformfieldgroupAPI.EzsignformfieldgroupCreateObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsignformfieldgroupCreateObjectV1`: EzsignformfieldgroupCreateObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignformfieldgroupAPI.EzsignformfieldgroupCreateObjectV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEzsignformfieldgroupCreateObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ezsignformfieldgroupCreateObjectV1Request** | [**EzsignformfieldgroupCreateObjectV1Request**](EzsignformfieldgroupCreateObjectV1Request.md) |  | 

### Return type

[**EzsignformfieldgroupCreateObjectV1Response**](EzsignformfieldgroupCreateObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignformfieldgroupDeleteObjectV1

> CommonResponse EzsignformfieldgroupDeleteObjectV1(ctx, pkiEzsignformfieldgroupID).Execute()

Delete an existing Ezsignformfieldgroup



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
	pkiEzsignformfieldgroupID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsignformfieldgroupAPI.EzsignformfieldgroupDeleteObjectV1(context.Background(), pkiEzsignformfieldgroupID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignformfieldgroupAPI.EzsignformfieldgroupDeleteObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsignformfieldgroupDeleteObjectV1`: CommonResponse
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignformfieldgroupAPI.EzsignformfieldgroupDeleteObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignformfieldgroupID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignformfieldgroupDeleteObjectV1Request struct via the builder pattern


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


## EzsignformfieldgroupEditObjectV1

> CommonResponse EzsignformfieldgroupEditObjectV1(ctx, pkiEzsignformfieldgroupID).EzsignformfieldgroupEditObjectV1Request(ezsignformfieldgroupEditObjectV1Request).Execute()

Edit an existing Ezsignformfieldgroup



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
	pkiEzsignformfieldgroupID := int32(56) // int32 | 
	ezsignformfieldgroupEditObjectV1Request := *openapiclient.NewEzsignformfieldgroupEditObjectV1Request(*openapiclient.NewEzsignformfieldgroupRequestCompound([]openapiclient.EzsignformfieldgroupsignerRequestCompound{*openapiclient.NewEzsignformfieldgroupsignerRequest(int32(20))}, []openapiclient.EzsignformfieldRequestCompound{*openapiclient.NewEzsignformfieldRequestCompound(int32(1), "Peanuts", int32(200), int32(300), int32(102), int32(22))}, int32(97), openapiclient.Field-eEzsignformfieldgroupType("Text"), "Allergies", int32(1), int32(1), int32(2), false)) // EzsignformfieldgroupEditObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsignformfieldgroupAPI.EzsignformfieldgroupEditObjectV1(context.Background(), pkiEzsignformfieldgroupID).EzsignformfieldgroupEditObjectV1Request(ezsignformfieldgroupEditObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignformfieldgroupAPI.EzsignformfieldgroupEditObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsignformfieldgroupEditObjectV1`: CommonResponse
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignformfieldgroupAPI.EzsignformfieldgroupEditObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignformfieldgroupID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignformfieldgroupEditObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezsignformfieldgroupEditObjectV1Request** | [**EzsignformfieldgroupEditObjectV1Request**](EzsignformfieldgroupEditObjectV1Request.md) |  | 

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


## EzsignformfieldgroupGetObjectV2

> EzsignformfieldgroupGetObjectV2Response EzsignformfieldgroupGetObjectV2(ctx, pkiEzsignformfieldgroupID).Execute()

Retrieve an existing Ezsignformfieldgroup



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
	pkiEzsignformfieldgroupID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsignformfieldgroupAPI.EzsignformfieldgroupGetObjectV2(context.Background(), pkiEzsignformfieldgroupID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignformfieldgroupAPI.EzsignformfieldgroupGetObjectV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsignformfieldgroupGetObjectV2`: EzsignformfieldgroupGetObjectV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignformfieldgroupAPI.EzsignformfieldgroupGetObjectV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignformfieldgroupID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignformfieldgroupGetObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsignformfieldgroupGetObjectV2Response**](EzsignformfieldgroupGetObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

