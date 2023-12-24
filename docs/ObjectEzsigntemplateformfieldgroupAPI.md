# eZmaxAPI\ObjectEzsigntemplateformfieldgroupAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EzsigntemplateformfieldgroupCreateObjectV1**](ObjectEzsigntemplateformfieldgroupAPI.md#EzsigntemplateformfieldgroupCreateObjectV1) | **Post** /1/object/ezsigntemplateformfieldgroup | Create a new Ezsigntemplateformfieldgroup
[**EzsigntemplateformfieldgroupDeleteObjectV1**](ObjectEzsigntemplateformfieldgroupAPI.md#EzsigntemplateformfieldgroupDeleteObjectV1) | **Delete** /1/object/ezsigntemplateformfieldgroup/{pkiEzsigntemplateformfieldgroupID} | Delete an existing Ezsigntemplateformfieldgroup
[**EzsigntemplateformfieldgroupEditObjectV1**](ObjectEzsigntemplateformfieldgroupAPI.md#EzsigntemplateformfieldgroupEditObjectV1) | **Put** /1/object/ezsigntemplateformfieldgroup/{pkiEzsigntemplateformfieldgroupID} | Edit an existing Ezsigntemplateformfieldgroup
[**EzsigntemplateformfieldgroupGetObjectV2**](ObjectEzsigntemplateformfieldgroupAPI.md#EzsigntemplateformfieldgroupGetObjectV2) | **Get** /2/object/ezsigntemplateformfieldgroup/{pkiEzsigntemplateformfieldgroupID} | Retrieve an existing Ezsigntemplateformfieldgroup



## EzsigntemplateformfieldgroupCreateObjectV1

> EzsigntemplateformfieldgroupCreateObjectV1Response EzsigntemplateformfieldgroupCreateObjectV1(ctx).EzsigntemplateformfieldgroupCreateObjectV1Request(ezsigntemplateformfieldgroupCreateObjectV1Request).Execute()

Create a new Ezsigntemplateformfieldgroup



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
	ezsigntemplateformfieldgroupCreateObjectV1Request := *openapiclient.NewEzsigntemplateformfieldgroupCreateObjectV1Request([]openapiclient.EzsigntemplateformfieldgroupRequestCompound{*openapiclient.NewEzsigntemplateformfieldgroupRequestCompound(int32(133), openapiclient.Field-eEzsigntemplateformfieldgroupType("Text"), openapiclient.Field-eEzsigntemplateformfieldgroupSignerrequirement("All"), "Allergies", int32(1), "Foo", int32(1), int32(2), false, []openapiclient.EzsigntemplateformfieldgroupsignerRequestCompound{*openapiclient.NewEzsigntemplateformfieldgroupsignerRequestCompound(int32(9))}, []openapiclient.EzsigntemplateformfieldRequestCompound{*openapiclient.NewEzsigntemplateformfieldRequestCompound(int32(1), "Peanuts", int32(200), int32(300), int32(102), int32(22))})}) // EzsigntemplateformfieldgroupCreateObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigntemplateformfieldgroupAPI.EzsigntemplateformfieldgroupCreateObjectV1(context.Background()).EzsigntemplateformfieldgroupCreateObjectV1Request(ezsigntemplateformfieldgroupCreateObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplateformfieldgroupAPI.EzsigntemplateformfieldgroupCreateObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigntemplateformfieldgroupCreateObjectV1`: EzsigntemplateformfieldgroupCreateObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplateformfieldgroupAPI.EzsigntemplateformfieldgroupCreateObjectV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplateformfieldgroupCreateObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ezsigntemplateformfieldgroupCreateObjectV1Request** | [**EzsigntemplateformfieldgroupCreateObjectV1Request**](EzsigntemplateformfieldgroupCreateObjectV1Request.md) |  | 

### Return type

[**EzsigntemplateformfieldgroupCreateObjectV1Response**](EzsigntemplateformfieldgroupCreateObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigntemplateformfieldgroupDeleteObjectV1

> EzsigntemplateformfieldgroupDeleteObjectV1Response EzsigntemplateformfieldgroupDeleteObjectV1(ctx, pkiEzsigntemplateformfieldgroupID).Execute()

Delete an existing Ezsigntemplateformfieldgroup



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
	pkiEzsigntemplateformfieldgroupID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigntemplateformfieldgroupAPI.EzsigntemplateformfieldgroupDeleteObjectV1(context.Background(), pkiEzsigntemplateformfieldgroupID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplateformfieldgroupAPI.EzsigntemplateformfieldgroupDeleteObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigntemplateformfieldgroupDeleteObjectV1`: EzsigntemplateformfieldgroupDeleteObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplateformfieldgroupAPI.EzsigntemplateformfieldgroupDeleteObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigntemplateformfieldgroupID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplateformfieldgroupDeleteObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsigntemplateformfieldgroupDeleteObjectV1Response**](EzsigntemplateformfieldgroupDeleteObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigntemplateformfieldgroupEditObjectV1

> EzsigntemplateformfieldgroupEditObjectV1Response EzsigntemplateformfieldgroupEditObjectV1(ctx, pkiEzsigntemplateformfieldgroupID).EzsigntemplateformfieldgroupEditObjectV1Request(ezsigntemplateformfieldgroupEditObjectV1Request).Execute()

Edit an existing Ezsigntemplateformfieldgroup



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
	pkiEzsigntemplateformfieldgroupID := int32(56) // int32 | 
	ezsigntemplateformfieldgroupEditObjectV1Request := *openapiclient.NewEzsigntemplateformfieldgroupEditObjectV1Request(*openapiclient.NewEzsigntemplateformfieldgroupRequestCompound(int32(133), openapiclient.Field-eEzsigntemplateformfieldgroupType("Text"), openapiclient.Field-eEzsigntemplateformfieldgroupSignerrequirement("All"), "Allergies", int32(1), "Foo", int32(1), int32(2), false, []openapiclient.EzsigntemplateformfieldgroupsignerRequestCompound{*openapiclient.NewEzsigntemplateformfieldgroupsignerRequestCompound(int32(9))}, []openapiclient.EzsigntemplateformfieldRequestCompound{*openapiclient.NewEzsigntemplateformfieldRequestCompound(int32(1), "Peanuts", int32(200), int32(300), int32(102), int32(22))})) // EzsigntemplateformfieldgroupEditObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigntemplateformfieldgroupAPI.EzsigntemplateformfieldgroupEditObjectV1(context.Background(), pkiEzsigntemplateformfieldgroupID).EzsigntemplateformfieldgroupEditObjectV1Request(ezsigntemplateformfieldgroupEditObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplateformfieldgroupAPI.EzsigntemplateformfieldgroupEditObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigntemplateformfieldgroupEditObjectV1`: EzsigntemplateformfieldgroupEditObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplateformfieldgroupAPI.EzsigntemplateformfieldgroupEditObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigntemplateformfieldgroupID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplateformfieldgroupEditObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezsigntemplateformfieldgroupEditObjectV1Request** | [**EzsigntemplateformfieldgroupEditObjectV1Request**](EzsigntemplateformfieldgroupEditObjectV1Request.md) |  | 

### Return type

[**EzsigntemplateformfieldgroupEditObjectV1Response**](EzsigntemplateformfieldgroupEditObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigntemplateformfieldgroupGetObjectV2

> EzsigntemplateformfieldgroupGetObjectV2Response EzsigntemplateformfieldgroupGetObjectV2(ctx, pkiEzsigntemplateformfieldgroupID).Execute()

Retrieve an existing Ezsigntemplateformfieldgroup



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
	pkiEzsigntemplateformfieldgroupID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigntemplateformfieldgroupAPI.EzsigntemplateformfieldgroupGetObjectV2(context.Background(), pkiEzsigntemplateformfieldgroupID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplateformfieldgroupAPI.EzsigntemplateformfieldgroupGetObjectV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigntemplateformfieldgroupGetObjectV2`: EzsigntemplateformfieldgroupGetObjectV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplateformfieldgroupAPI.EzsigntemplateformfieldgroupGetObjectV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigntemplateformfieldgroupID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplateformfieldgroupGetObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsigntemplateformfieldgroupGetObjectV2Response**](EzsigntemplateformfieldgroupGetObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

