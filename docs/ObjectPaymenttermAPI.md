# eZmaxAPI\ObjectPaymenttermAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PaymenttermCreateObjectV1**](ObjectPaymenttermAPI.md#PaymenttermCreateObjectV1) | **Post** /1/object/paymentterm | Create a new Paymentterm
[**PaymenttermEditObjectV1**](ObjectPaymenttermAPI.md#PaymenttermEditObjectV1) | **Put** /1/object/paymentterm/{pkiPaymenttermID} | Edit an existing Paymentterm
[**PaymenttermGetAutocompleteV2**](ObjectPaymenttermAPI.md#PaymenttermGetAutocompleteV2) | **Get** /2/object/paymentterm/getAutocomplete/{sSelector} | Retrieve Paymentterms and IDs
[**PaymenttermGetListV1**](ObjectPaymenttermAPI.md#PaymenttermGetListV1) | **Get** /1/object/paymentterm/getList | Retrieve Paymentterm list
[**PaymenttermGetObjectV2**](ObjectPaymenttermAPI.md#PaymenttermGetObjectV2) | **Get** /2/object/paymentterm/{pkiPaymenttermID} | Retrieve an existing Paymentterm



## PaymenttermCreateObjectV1

> PaymenttermCreateObjectV1Response PaymenttermCreateObjectV1(ctx).PaymenttermCreateObjectV1Request(paymenttermCreateObjectV1Request).Execute()

Create a new Paymentterm



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
	paymenttermCreateObjectV1Request := *openapiclient.NewPaymenttermCreateObjectV1Request([]openapiclient.PaymenttermRequestCompound{*openapiclient.NewPaymenttermRequestCompound("0030", openapiclient.Field-ePaymenttermType("Days"), int32(30), *openapiclient.NewMultilingualPaymenttermDescription(), true)}) // PaymenttermCreateObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectPaymenttermAPI.PaymenttermCreateObjectV1(context.Background()).PaymenttermCreateObjectV1Request(paymenttermCreateObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectPaymenttermAPI.PaymenttermCreateObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymenttermCreateObjectV1`: PaymenttermCreateObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectPaymenttermAPI.PaymenttermCreateObjectV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymenttermCreateObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymenttermCreateObjectV1Request** | [**PaymenttermCreateObjectV1Request**](PaymenttermCreateObjectV1Request.md) |  | 

### Return type

[**PaymenttermCreateObjectV1Response**](PaymenttermCreateObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymenttermEditObjectV1

> PaymenttermEditObjectV1Response PaymenttermEditObjectV1(ctx, pkiPaymenttermID).PaymenttermEditObjectV1Request(paymenttermEditObjectV1Request).Execute()

Edit an existing Paymentterm



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
	pkiPaymenttermID := int32(56) // int32 | 
	paymenttermEditObjectV1Request := *openapiclient.NewPaymenttermEditObjectV1Request(*openapiclient.NewPaymenttermRequestCompound("0030", openapiclient.Field-ePaymenttermType("Days"), int32(30), *openapiclient.NewMultilingualPaymenttermDescription(), true)) // PaymenttermEditObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectPaymenttermAPI.PaymenttermEditObjectV1(context.Background(), pkiPaymenttermID).PaymenttermEditObjectV1Request(paymenttermEditObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectPaymenttermAPI.PaymenttermEditObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymenttermEditObjectV1`: PaymenttermEditObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectPaymenttermAPI.PaymenttermEditObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiPaymenttermID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaymenttermEditObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paymenttermEditObjectV1Request** | [**PaymenttermEditObjectV1Request**](PaymenttermEditObjectV1Request.md) |  | 

### Return type

[**PaymenttermEditObjectV1Response**](PaymenttermEditObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymenttermGetAutocompleteV2

> PaymenttermGetAutocompleteV2Response PaymenttermGetAutocompleteV2(ctx, sSelector).EFilterActive(eFilterActive).SQuery(sQuery).AcceptLanguage(acceptLanguage).Execute()

Retrieve Paymentterms and IDs



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
	sSelector := "sSelector_example" // string | The type of Paymentterms to return
	eFilterActive := "eFilterActive_example" // string | Specify which results we want to display. (optional) (default to "Active")
	sQuery := "sQuery_example" // string | Allow to filter the returned results (optional)
	acceptLanguage := openapiclient.Header-Accept-Language("*") // HeaderAcceptLanguage |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectPaymenttermAPI.PaymenttermGetAutocompleteV2(context.Background(), sSelector).EFilterActive(eFilterActive).SQuery(sQuery).AcceptLanguage(acceptLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectPaymenttermAPI.PaymenttermGetAutocompleteV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymenttermGetAutocompleteV2`: PaymenttermGetAutocompleteV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectPaymenttermAPI.PaymenttermGetAutocompleteV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sSelector** | **string** | The type of Paymentterms to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaymenttermGetAutocompleteV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eFilterActive** | **string** | Specify which results we want to display. | [default to &quot;Active&quot;]
 **sQuery** | **string** | Allow to filter the returned results | 
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 

### Return type

[**PaymenttermGetAutocompleteV2Response**](PaymenttermGetAutocompleteV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymenttermGetListV1

> PaymenttermGetListV1Response PaymenttermGetListV1(ctx).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()

Retrieve Paymentterm list

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
	resp, r, err := apiClient.ObjectPaymenttermAPI.PaymenttermGetListV1(context.Background()).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectPaymenttermAPI.PaymenttermGetListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymenttermGetListV1`: PaymenttermGetListV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectPaymenttermAPI.PaymenttermGetListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymenttermGetListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eOrderBy** | **string** | Specify how you want the results to be sorted | 
 **iRowMax** | **int32** |  | 
 **iRowOffset** | **int32** |  | [default to 0]
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 
 **sFilter** | **string** |  | 

### Return type

[**PaymenttermGetListV1Response**](PaymenttermGetListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymenttermGetObjectV2

> PaymenttermGetObjectV2Response PaymenttermGetObjectV2(ctx, pkiPaymenttermID).Execute()

Retrieve an existing Paymentterm



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
	pkiPaymenttermID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectPaymenttermAPI.PaymenttermGetObjectV2(context.Background(), pkiPaymenttermID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectPaymenttermAPI.PaymenttermGetObjectV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymenttermGetObjectV2`: PaymenttermGetObjectV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectPaymenttermAPI.PaymenttermGetObjectV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiPaymenttermID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaymenttermGetObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PaymenttermGetObjectV2Response**](PaymenttermGetObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

