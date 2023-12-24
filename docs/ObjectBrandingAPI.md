# eZmaxAPI\ObjectBrandingAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BrandingCreateObjectV1**](ObjectBrandingAPI.md#BrandingCreateObjectV1) | **Post** /1/object/branding | Create a new Branding
[**BrandingEditObjectV1**](ObjectBrandingAPI.md#BrandingEditObjectV1) | **Put** /1/object/branding/{pkiBrandingID} | Edit an existing Branding
[**BrandingGetAutocompleteV2**](ObjectBrandingAPI.md#BrandingGetAutocompleteV2) | **Get** /2/object/branding/getAutocomplete/{sSelector} | Retrieve Brandings and IDs
[**BrandingGetListV1**](ObjectBrandingAPI.md#BrandingGetListV1) | **Get** /1/object/branding/getList | Retrieve Branding list
[**BrandingGetObjectV2**](ObjectBrandingAPI.md#BrandingGetObjectV2) | **Get** /2/object/branding/{pkiBrandingID} | Retrieve an existing Branding



## BrandingCreateObjectV1

> BrandingCreateObjectV1Response BrandingCreateObjectV1(ctx).BrandingCreateObjectV1Request(brandingCreateObjectV1Request).Execute()

Create a new Branding



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
	brandingCreateObjectV1Request := *openapiclient.NewBrandingCreateObjectV1Request([]openapiclient.BrandingRequestCompound{*openapiclient.NewBrandingRequestCompound(*openapiclient.NewMultilingualBrandingDescription(), openapiclient.Field-eBrandingLogo("Default"), int32(3752795), int32(0), int32(16777215), int32(15658734), int32(13577007), int32(16777215), true)}) // BrandingCreateObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectBrandingAPI.BrandingCreateObjectV1(context.Background()).BrandingCreateObjectV1Request(brandingCreateObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectBrandingAPI.BrandingCreateObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BrandingCreateObjectV1`: BrandingCreateObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectBrandingAPI.BrandingCreateObjectV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBrandingCreateObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **brandingCreateObjectV1Request** | [**BrandingCreateObjectV1Request**](BrandingCreateObjectV1Request.md) |  | 

### Return type

[**BrandingCreateObjectV1Response**](BrandingCreateObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BrandingEditObjectV1

> BrandingEditObjectV1Response BrandingEditObjectV1(ctx, pkiBrandingID).BrandingEditObjectV1Request(brandingEditObjectV1Request).Execute()

Edit an existing Branding



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
	pkiBrandingID := int32(56) // int32 | 
	brandingEditObjectV1Request := *openapiclient.NewBrandingEditObjectV1Request(*openapiclient.NewBrandingRequestCompound(*openapiclient.NewMultilingualBrandingDescription(), openapiclient.Field-eBrandingLogo("Default"), int32(3752795), int32(0), int32(16777215), int32(15658734), int32(13577007), int32(16777215), true)) // BrandingEditObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectBrandingAPI.BrandingEditObjectV1(context.Background(), pkiBrandingID).BrandingEditObjectV1Request(brandingEditObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectBrandingAPI.BrandingEditObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BrandingEditObjectV1`: BrandingEditObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectBrandingAPI.BrandingEditObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiBrandingID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBrandingEditObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **brandingEditObjectV1Request** | [**BrandingEditObjectV1Request**](BrandingEditObjectV1Request.md) |  | 

### Return type

[**BrandingEditObjectV1Response**](BrandingEditObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BrandingGetAutocompleteV2

> BrandingGetAutocompleteV2Response BrandingGetAutocompleteV2(ctx, sSelector).EFilterActive(eFilterActive).SQuery(sQuery).AcceptLanguage(acceptLanguage).Execute()

Retrieve Brandings and IDs



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
	sSelector := "sSelector_example" // string | The type of Brandings to return
	eFilterActive := "eFilterActive_example" // string | Specify which results we want to display. (optional) (default to "Active")
	sQuery := "sQuery_example" // string | Allow to filter the returned results (optional)
	acceptLanguage := openapiclient.Header-Accept-Language("*") // HeaderAcceptLanguage |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectBrandingAPI.BrandingGetAutocompleteV2(context.Background(), sSelector).EFilterActive(eFilterActive).SQuery(sQuery).AcceptLanguage(acceptLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectBrandingAPI.BrandingGetAutocompleteV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BrandingGetAutocompleteV2`: BrandingGetAutocompleteV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectBrandingAPI.BrandingGetAutocompleteV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sSelector** | **string** | The type of Brandings to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiBrandingGetAutocompleteV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eFilterActive** | **string** | Specify which results we want to display. | [default to &quot;Active&quot;]
 **sQuery** | **string** | Allow to filter the returned results | 
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 

### Return type

[**BrandingGetAutocompleteV2Response**](BrandingGetAutocompleteV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BrandingGetListV1

> BrandingGetListV1Response BrandingGetListV1(ctx).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()

Retrieve Branding list



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
	eOrderBy := "eOrderBy_example" // string | Specify how you want the results to be sorted (optional)
	iRowMax := int32(56) // int32 |  (optional)
	iRowOffset := int32(56) // int32 |  (optional) (default to 0)
	acceptLanguage := openapiclient.Header-Accept-Language("*") // HeaderAcceptLanguage |  (optional)
	sFilter := "sFilter_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectBrandingAPI.BrandingGetListV1(context.Background()).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectBrandingAPI.BrandingGetListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BrandingGetListV1`: BrandingGetListV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectBrandingAPI.BrandingGetListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBrandingGetListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eOrderBy** | **string** | Specify how you want the results to be sorted | 
 **iRowMax** | **int32** |  | 
 **iRowOffset** | **int32** |  | [default to 0]
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 
 **sFilter** | **string** |  | 

### Return type

[**BrandingGetListV1Response**](BrandingGetListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BrandingGetObjectV2

> BrandingGetObjectV2Response BrandingGetObjectV2(ctx, pkiBrandingID).Execute()

Retrieve an existing Branding



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
	pkiBrandingID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectBrandingAPI.BrandingGetObjectV2(context.Background(), pkiBrandingID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectBrandingAPI.BrandingGetObjectV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BrandingGetObjectV2`: BrandingGetObjectV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectBrandingAPI.BrandingGetObjectV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiBrandingID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBrandingGetObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BrandingGetObjectV2Response**](BrandingGetObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

