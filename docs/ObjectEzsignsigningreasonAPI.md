# eZmaxAPI\ObjectEzsignsigningreasonAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EzsignsigningreasonCreateObjectV1**](ObjectEzsignsigningreasonAPI.md#EzsignsigningreasonCreateObjectV1) | **Post** /1/object/ezsignsigningreason | Create a new Ezsignsigningreason
[**EzsignsigningreasonEditObjectV1**](ObjectEzsignsigningreasonAPI.md#EzsignsigningreasonEditObjectV1) | **Put** /1/object/ezsignsigningreason/{pkiEzsignsigningreasonID} | Edit an existing Ezsignsigningreason
[**EzsignsigningreasonGetAutocompleteV2**](ObjectEzsignsigningreasonAPI.md#EzsignsigningreasonGetAutocompleteV2) | **Get** /2/object/ezsignsigningreason/getAutocomplete/{sSelector} | Retrieve Ezsignsigningreasons and IDs
[**EzsignsigningreasonGetListV1**](ObjectEzsignsigningreasonAPI.md#EzsignsigningreasonGetListV1) | **Get** /1/object/ezsignsigningreason/getList | Retrieve Ezsignsigningreason list
[**EzsignsigningreasonGetObjectV2**](ObjectEzsignsigningreasonAPI.md#EzsignsigningreasonGetObjectV2) | **Get** /2/object/ezsignsigningreason/{pkiEzsignsigningreasonID} | Retrieve an existing Ezsignsigningreason



## EzsignsigningreasonCreateObjectV1

> EzsignsigningreasonCreateObjectV1Response EzsignsigningreasonCreateObjectV1(ctx).EzsignsigningreasonCreateObjectV1Request(ezsignsigningreasonCreateObjectV1Request).Execute()

Create a new Ezsignsigningreason



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
	ezsignsigningreasonCreateObjectV1Request := *openapiclient.NewEzsignsigningreasonCreateObjectV1Request([]openapiclient.EzsignsigningreasonRequestCompound{*openapiclient.NewEzsignsigningreasonRequestCompound(*openapiclient.NewMultilingualEzsignsigningreasonDescription(), true)}) // EzsignsigningreasonCreateObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsignsigningreasonAPI.EzsignsigningreasonCreateObjectV1(context.Background()).EzsignsigningreasonCreateObjectV1Request(ezsignsigningreasonCreateObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignsigningreasonAPI.EzsignsigningreasonCreateObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsignsigningreasonCreateObjectV1`: EzsignsigningreasonCreateObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignsigningreasonAPI.EzsignsigningreasonCreateObjectV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEzsignsigningreasonCreateObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ezsignsigningreasonCreateObjectV1Request** | [**EzsignsigningreasonCreateObjectV1Request**](EzsignsigningreasonCreateObjectV1Request.md) |  | 

### Return type

[**EzsignsigningreasonCreateObjectV1Response**](EzsignsigningreasonCreateObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignsigningreasonEditObjectV1

> CommonResponse EzsignsigningreasonEditObjectV1(ctx, pkiEzsignsigningreasonID).EzsignsigningreasonEditObjectV1Request(ezsignsigningreasonEditObjectV1Request).Execute()

Edit an existing Ezsignsigningreason



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
	pkiEzsignsigningreasonID := int32(56) // int32 | The unique ID of the Ezsignsigningreason
	ezsignsigningreasonEditObjectV1Request := *openapiclient.NewEzsignsigningreasonEditObjectV1Request(*openapiclient.NewEzsignsigningreasonRequestCompound(*openapiclient.NewMultilingualEzsignsigningreasonDescription(), true)) // EzsignsigningreasonEditObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsignsigningreasonAPI.EzsignsigningreasonEditObjectV1(context.Background(), pkiEzsignsigningreasonID).EzsignsigningreasonEditObjectV1Request(ezsignsigningreasonEditObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignsigningreasonAPI.EzsignsigningreasonEditObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsignsigningreasonEditObjectV1`: CommonResponse
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignsigningreasonAPI.EzsignsigningreasonEditObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignsigningreasonID** | **int32** | The unique ID of the Ezsignsigningreason | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignsigningreasonEditObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezsignsigningreasonEditObjectV1Request** | [**EzsignsigningreasonEditObjectV1Request**](EzsignsigningreasonEditObjectV1Request.md) |  | 

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


## EzsignsigningreasonGetAutocompleteV2

> EzsignsigningreasonGetAutocompleteV2Response EzsignsigningreasonGetAutocompleteV2(ctx, sSelector).EFilterActive(eFilterActive).SQuery(sQuery).AcceptLanguage(acceptLanguage).Execute()

Retrieve Ezsignsigningreasons and IDs



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
	sSelector := "sSelector_example" // string | The type of Ezsignsigningreasons to return
	eFilterActive := "eFilterActive_example" // string | Specify which results we want to display. (optional) (default to "Active")
	sQuery := "sQuery_example" // string | Allow to filter the returned results (optional)
	acceptLanguage := openapiclient.Header-Accept-Language("*") // HeaderAcceptLanguage |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsignsigningreasonAPI.EzsignsigningreasonGetAutocompleteV2(context.Background(), sSelector).EFilterActive(eFilterActive).SQuery(sQuery).AcceptLanguage(acceptLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignsigningreasonAPI.EzsignsigningreasonGetAutocompleteV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsignsigningreasonGetAutocompleteV2`: EzsignsigningreasonGetAutocompleteV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignsigningreasonAPI.EzsignsigningreasonGetAutocompleteV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sSelector** | **string** | The type of Ezsignsigningreasons to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignsigningreasonGetAutocompleteV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eFilterActive** | **string** | Specify which results we want to display. | [default to &quot;Active&quot;]
 **sQuery** | **string** | Allow to filter the returned results | 
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 

### Return type

[**EzsignsigningreasonGetAutocompleteV2Response**](EzsignsigningreasonGetAutocompleteV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignsigningreasonGetListV1

> EzsignsigningreasonGetListV1Response EzsignsigningreasonGetListV1(ctx).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()

Retrieve Ezsignsigningreason list



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
	resp, r, err := apiClient.ObjectEzsignsigningreasonAPI.EzsignsigningreasonGetListV1(context.Background()).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignsigningreasonAPI.EzsignsigningreasonGetListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsignsigningreasonGetListV1`: EzsignsigningreasonGetListV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignsigningreasonAPI.EzsignsigningreasonGetListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEzsignsigningreasonGetListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eOrderBy** | **string** | Specify how you want the results to be sorted | 
 **iRowMax** | **int32** |  | 
 **iRowOffset** | **int32** |  | [default to 0]
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 
 **sFilter** | **string** |  | 

### Return type

[**EzsignsigningreasonGetListV1Response**](EzsignsigningreasonGetListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignsigningreasonGetObjectV2

> EzsignsigningreasonGetObjectV2Response EzsignsigningreasonGetObjectV2(ctx, pkiEzsignsigningreasonID).Execute()

Retrieve an existing Ezsignsigningreason



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
	pkiEzsignsigningreasonID := int32(56) // int32 | The unique ID of the Ezsignsigningreason

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsignsigningreasonAPI.EzsignsigningreasonGetObjectV2(context.Background(), pkiEzsignsigningreasonID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignsigningreasonAPI.EzsignsigningreasonGetObjectV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsignsigningreasonGetObjectV2`: EzsignsigningreasonGetObjectV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignsigningreasonAPI.EzsignsigningreasonGetObjectV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignsigningreasonID** | **int32** | The unique ID of the Ezsignsigningreason | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignsigningreasonGetObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsignsigningreasonGetObjectV2Response**](EzsignsigningreasonGetObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

