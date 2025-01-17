# eZmaxAPI\ObjectBillingentityinternalAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BillingentityinternalCreateObjectV1**](ObjectBillingentityinternalAPI.md#BillingentityinternalCreateObjectV1) | **Post** /1/object/billingentityinternal | Create a new Billingentityinternal
[**BillingentityinternalEditObjectV1**](ObjectBillingentityinternalAPI.md#BillingentityinternalEditObjectV1) | **Put** /1/object/billingentityinternal/{pkiBillingentityinternalID} | Edit an existing Billingentityinternal
[**BillingentityinternalGetAutocompleteV2**](ObjectBillingentityinternalAPI.md#BillingentityinternalGetAutocompleteV2) | **Get** /2/object/billingentityinternal/getAutocomplete/{sSelector} | Retrieve Billingentityinternals and IDs
[**BillingentityinternalGetListV1**](ObjectBillingentityinternalAPI.md#BillingentityinternalGetListV1) | **Get** /1/object/billingentityinternal/getList | Retrieve Billingentityinternal list
[**BillingentityinternalGetObjectV2**](ObjectBillingentityinternalAPI.md#BillingentityinternalGetObjectV2) | **Get** /2/object/billingentityinternal/{pkiBillingentityinternalID} | Retrieve an existing Billingentityinternal



## BillingentityinternalCreateObjectV1

> BillingentityinternalCreateObjectV1Response BillingentityinternalCreateObjectV1(ctx).BillingentityinternalCreateObjectV1Request(billingentityinternalCreateObjectV1Request).Execute()

Create a new Billingentityinternal



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
	billingentityinternalCreateObjectV1Request := *openapiclient.NewBillingentityinternalCreateObjectV1Request([]openapiclient.BillingentityinternalRequestCompound{*openapiclient.NewBillingentityinternalRequestCompound([]openapiclient.BillingentityinternalproductRequestCompound{*openapiclient.NewBillingentityinternalproductRequestCompound(int32(172), int32(83))}, *openapiclient.NewMultilingualBillingentityinternalDescription())}) // BillingentityinternalCreateObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectBillingentityinternalAPI.BillingentityinternalCreateObjectV1(context.Background()).BillingentityinternalCreateObjectV1Request(billingentityinternalCreateObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectBillingentityinternalAPI.BillingentityinternalCreateObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingentityinternalCreateObjectV1`: BillingentityinternalCreateObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectBillingentityinternalAPI.BillingentityinternalCreateObjectV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBillingentityinternalCreateObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **billingentityinternalCreateObjectV1Request** | [**BillingentityinternalCreateObjectV1Request**](BillingentityinternalCreateObjectV1Request.md) |  | 

### Return type

[**BillingentityinternalCreateObjectV1Response**](BillingentityinternalCreateObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingentityinternalEditObjectV1

> CommonResponse BillingentityinternalEditObjectV1(ctx, pkiBillingentityinternalID).BillingentityinternalEditObjectV1Request(billingentityinternalEditObjectV1Request).Execute()

Edit an existing Billingentityinternal



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
	pkiBillingentityinternalID := int32(56) // int32 | 
	billingentityinternalEditObjectV1Request := *openapiclient.NewBillingentityinternalEditObjectV1Request(*openapiclient.NewBillingentityinternalRequestCompound([]openapiclient.BillingentityinternalproductRequestCompound{*openapiclient.NewBillingentityinternalproductRequestCompound(int32(172), int32(83))}, *openapiclient.NewMultilingualBillingentityinternalDescription())) // BillingentityinternalEditObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectBillingentityinternalAPI.BillingentityinternalEditObjectV1(context.Background(), pkiBillingentityinternalID).BillingentityinternalEditObjectV1Request(billingentityinternalEditObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectBillingentityinternalAPI.BillingentityinternalEditObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingentityinternalEditObjectV1`: CommonResponse
	fmt.Fprintf(os.Stdout, "Response from `ObjectBillingentityinternalAPI.BillingentityinternalEditObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiBillingentityinternalID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBillingentityinternalEditObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **billingentityinternalEditObjectV1Request** | [**BillingentityinternalEditObjectV1Request**](BillingentityinternalEditObjectV1Request.md) |  | 

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


## BillingentityinternalGetAutocompleteV2

> BillingentityinternalGetAutocompleteV2Response BillingentityinternalGetAutocompleteV2(ctx, sSelector).EFilterActive(eFilterActive).SQuery(sQuery).AcceptLanguage(acceptLanguage).Execute()

Retrieve Billingentityinternals and IDs



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
	sSelector := "sSelector_example" // string | The type of Billingentityinternals to return
	eFilterActive := "eFilterActive_example" // string | Specify which results we want to display. (optional) (default to "Active")
	sQuery := "sQuery_example" // string | Allow to filter the returned results (optional)
	acceptLanguage := openapiclient.Header-Accept-Language("*") // HeaderAcceptLanguage |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectBillingentityinternalAPI.BillingentityinternalGetAutocompleteV2(context.Background(), sSelector).EFilterActive(eFilterActive).SQuery(sQuery).AcceptLanguage(acceptLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectBillingentityinternalAPI.BillingentityinternalGetAutocompleteV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingentityinternalGetAutocompleteV2`: BillingentityinternalGetAutocompleteV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectBillingentityinternalAPI.BillingentityinternalGetAutocompleteV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sSelector** | **string** | The type of Billingentityinternals to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiBillingentityinternalGetAutocompleteV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eFilterActive** | **string** | Specify which results we want to display. | [default to &quot;Active&quot;]
 **sQuery** | **string** | Allow to filter the returned results | 
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 

### Return type

[**BillingentityinternalGetAutocompleteV2Response**](BillingentityinternalGetAutocompleteV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingentityinternalGetListV1

> BillingentityinternalGetListV1Response BillingentityinternalGetListV1(ctx).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()

Retrieve Billingentityinternal list



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
	resp, r, err := apiClient.ObjectBillingentityinternalAPI.BillingentityinternalGetListV1(context.Background()).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectBillingentityinternalAPI.BillingentityinternalGetListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingentityinternalGetListV1`: BillingentityinternalGetListV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectBillingentityinternalAPI.BillingentityinternalGetListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBillingentityinternalGetListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eOrderBy** | **string** | Specify how you want the results to be sorted | 
 **iRowMax** | **int32** |  | 
 **iRowOffset** | **int32** |  | [default to 0]
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 
 **sFilter** | **string** |  | 

### Return type

[**BillingentityinternalGetListV1Response**](BillingentityinternalGetListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingentityinternalGetObjectV2

> BillingentityinternalGetObjectV2Response BillingentityinternalGetObjectV2(ctx, pkiBillingentityinternalID).Execute()

Retrieve an existing Billingentityinternal



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
	pkiBillingentityinternalID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectBillingentityinternalAPI.BillingentityinternalGetObjectV2(context.Background(), pkiBillingentityinternalID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectBillingentityinternalAPI.BillingentityinternalGetObjectV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingentityinternalGetObjectV2`: BillingentityinternalGetObjectV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectBillingentityinternalAPI.BillingentityinternalGetObjectV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiBillingentityinternalID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBillingentityinternalGetObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BillingentityinternalGetObjectV2Response**](BillingentityinternalGetObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

