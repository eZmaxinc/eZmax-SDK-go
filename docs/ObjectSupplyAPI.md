# eZmaxAPI\ObjectSupplyAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SupplyCreateObjectV1**](ObjectSupplyAPI.md#SupplyCreateObjectV1) | **Post** /1/object/supply | Create a new Supply
[**SupplyDeleteObjectV1**](ObjectSupplyAPI.md#SupplyDeleteObjectV1) | **Delete** /1/object/supply/{pkiSupplyID} | Delete an existing Supply
[**SupplyEditObjectV1**](ObjectSupplyAPI.md#SupplyEditObjectV1) | **Put** /1/object/supply/{pkiSupplyID} | Edit an existing Supply
[**SupplyGetAutocompleteV2**](ObjectSupplyAPI.md#SupplyGetAutocompleteV2) | **Get** /2/object/supply/getAutocomplete/{sSelector} | Retrieve Supplys and IDs
[**SupplyGetListV1**](ObjectSupplyAPI.md#SupplyGetListV1) | **Get** /1/object/supply/getList | Retrieve Supply list
[**SupplyGetObjectV2**](ObjectSupplyAPI.md#SupplyGetObjectV2) | **Get** /2/object/supply/{pkiSupplyID} | Retrieve an existing Supply



## SupplyCreateObjectV1

> SupplyCreateObjectV1Response SupplyCreateObjectV1(ctx).SupplyCreateObjectV1Request(supplyCreateObjectV1Request).Execute()

Create a new Supply



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
	supplyCreateObjectV1Request := *openapiclient.NewSupplyCreateObjectV1Request([]openapiclient.SupplyRequestCompound{*openapiclient.NewSupplyRequestCompound(int32(2), "PPLET", *openapiclient.NewMultilingualSupplyDescription(), "8", true, true)}) // SupplyCreateObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectSupplyAPI.SupplyCreateObjectV1(context.Background()).SupplyCreateObjectV1Request(supplyCreateObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectSupplyAPI.SupplyCreateObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SupplyCreateObjectV1`: SupplyCreateObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectSupplyAPI.SupplyCreateObjectV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSupplyCreateObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **supplyCreateObjectV1Request** | [**SupplyCreateObjectV1Request**](SupplyCreateObjectV1Request.md) |  | 

### Return type

[**SupplyCreateObjectV1Response**](SupplyCreateObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SupplyDeleteObjectV1

> SupplyDeleteObjectV1Response SupplyDeleteObjectV1(ctx, pkiSupplyID).Execute()

Delete an existing Supply



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
	pkiSupplyID := int32(56) // int32 | The unique ID of the Supply

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectSupplyAPI.SupplyDeleteObjectV1(context.Background(), pkiSupplyID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectSupplyAPI.SupplyDeleteObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SupplyDeleteObjectV1`: SupplyDeleteObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectSupplyAPI.SupplyDeleteObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiSupplyID** | **int32** | The unique ID of the Supply | 

### Other Parameters

Other parameters are passed through a pointer to a apiSupplyDeleteObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SupplyDeleteObjectV1Response**](SupplyDeleteObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SupplyEditObjectV1

> SupplyEditObjectV1Response SupplyEditObjectV1(ctx, pkiSupplyID).SupplyEditObjectV1Request(supplyEditObjectV1Request).Execute()

Edit an existing Supply



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
	pkiSupplyID := int32(56) // int32 | The unique ID of the Supply
	supplyEditObjectV1Request := *openapiclient.NewSupplyEditObjectV1Request(*openapiclient.NewSupplyRequestCompound(int32(2), "PPLET", *openapiclient.NewMultilingualSupplyDescription(), "8", true, true)) // SupplyEditObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectSupplyAPI.SupplyEditObjectV1(context.Background(), pkiSupplyID).SupplyEditObjectV1Request(supplyEditObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectSupplyAPI.SupplyEditObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SupplyEditObjectV1`: SupplyEditObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectSupplyAPI.SupplyEditObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiSupplyID** | **int32** | The unique ID of the Supply | 

### Other Parameters

Other parameters are passed through a pointer to a apiSupplyEditObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **supplyEditObjectV1Request** | [**SupplyEditObjectV1Request**](SupplyEditObjectV1Request.md) |  | 

### Return type

[**SupplyEditObjectV1Response**](SupplyEditObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SupplyGetAutocompleteV2

> SupplyGetAutocompleteV2Response SupplyGetAutocompleteV2(ctx, sSelector).EFilterActive(eFilterActive).SQuery(sQuery).AcceptLanguage(acceptLanguage).Execute()

Retrieve Supplys and IDs



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
	sSelector := "sSelector_example" // string | The type of Supplys to return
	eFilterActive := "eFilterActive_example" // string | Specify which results we want to display. (optional) (default to "Active")
	sQuery := "sQuery_example" // string | Allow to filter the returned results (optional)
	acceptLanguage := openapiclient.Header-Accept-Language("*") // HeaderAcceptLanguage |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectSupplyAPI.SupplyGetAutocompleteV2(context.Background(), sSelector).EFilterActive(eFilterActive).SQuery(sQuery).AcceptLanguage(acceptLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectSupplyAPI.SupplyGetAutocompleteV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SupplyGetAutocompleteV2`: SupplyGetAutocompleteV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectSupplyAPI.SupplyGetAutocompleteV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sSelector** | **string** | The type of Supplys to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiSupplyGetAutocompleteV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eFilterActive** | **string** | Specify which results we want to display. | [default to &quot;Active&quot;]
 **sQuery** | **string** | Allow to filter the returned results | 
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 

### Return type

[**SupplyGetAutocompleteV2Response**](SupplyGetAutocompleteV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SupplyGetListV1

> SupplyGetListV1Response SupplyGetListV1(ctx).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()

Retrieve Supply list



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
	resp, r, err := apiClient.ObjectSupplyAPI.SupplyGetListV1(context.Background()).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectSupplyAPI.SupplyGetListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SupplyGetListV1`: SupplyGetListV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectSupplyAPI.SupplyGetListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSupplyGetListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eOrderBy** | **string** | Specify how you want the results to be sorted | 
 **iRowMax** | **int32** |  | 
 **iRowOffset** | **int32** |  | [default to 0]
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 
 **sFilter** | **string** |  | 

### Return type

[**SupplyGetListV1Response**](SupplyGetListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SupplyGetObjectV2

> SupplyGetObjectV2Response SupplyGetObjectV2(ctx, pkiSupplyID).Execute()

Retrieve an existing Supply



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
	pkiSupplyID := int32(56) // int32 | The unique ID of the Supply

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectSupplyAPI.SupplyGetObjectV2(context.Background(), pkiSupplyID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectSupplyAPI.SupplyGetObjectV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SupplyGetObjectV2`: SupplyGetObjectV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectSupplyAPI.SupplyGetObjectV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiSupplyID** | **int32** | The unique ID of the Supply | 

### Other Parameters

Other parameters are passed through a pointer to a apiSupplyGetObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SupplyGetObjectV2Response**](SupplyGetObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

