# eZmaxAPI\ObjectPaymentgatewayAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PaymentgatewayCreateObjectV1**](ObjectPaymentgatewayAPI.md#PaymentgatewayCreateObjectV1) | **Post** /1/object/paymentgateway | Create a new Paymentgateway
[**PaymentgatewayEditObjectV1**](ObjectPaymentgatewayAPI.md#PaymentgatewayEditObjectV1) | **Put** /1/object/paymentgateway/{pkiPaymentgatewayID} | Edit an existing Paymentgateway
[**PaymentgatewayGetAutocompleteV2**](ObjectPaymentgatewayAPI.md#PaymentgatewayGetAutocompleteV2) | **Get** /2/object/paymentgateway/getAutocomplete/{sSelector} | Retrieve Paymentgateways and IDs
[**PaymentgatewayGetListV1**](ObjectPaymentgatewayAPI.md#PaymentgatewayGetListV1) | **Get** /1/object/paymentgateway/getList | Retrieve Paymentgateway list
[**PaymentgatewayGetObjectV2**](ObjectPaymentgatewayAPI.md#PaymentgatewayGetObjectV2) | **Get** /2/object/paymentgateway/{pkiPaymentgatewayID} | Retrieve an existing Paymentgateway



## PaymentgatewayCreateObjectV1

> PaymentgatewayCreateObjectV1Response PaymentgatewayCreateObjectV1(ctx).PaymentgatewayCreateObjectV1Request(paymentgatewayCreateObjectV1Request).Execute()

Create a new Paymentgateway



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
	paymentgatewayCreateObjectV1Request := *openapiclient.NewPaymentgatewayCreateObjectV1Request([]openapiclient.PaymentgatewayRequestCompound{*openapiclient.NewPaymentgatewayRequestCompound(openapiclient.Field-ePaymentgatewayProcessor("Moneris"), *openapiclient.NewMultilingualPaymentgatewayDescription())}) // PaymentgatewayCreateObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectPaymentgatewayAPI.PaymentgatewayCreateObjectV1(context.Background()).PaymentgatewayCreateObjectV1Request(paymentgatewayCreateObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectPaymentgatewayAPI.PaymentgatewayCreateObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentgatewayCreateObjectV1`: PaymentgatewayCreateObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectPaymentgatewayAPI.PaymentgatewayCreateObjectV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentgatewayCreateObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentgatewayCreateObjectV1Request** | [**PaymentgatewayCreateObjectV1Request**](PaymentgatewayCreateObjectV1Request.md) |  | 

### Return type

[**PaymentgatewayCreateObjectV1Response**](PaymentgatewayCreateObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentgatewayEditObjectV1

> PaymentgatewayEditObjectV1Response PaymentgatewayEditObjectV1(ctx, pkiPaymentgatewayID).PaymentgatewayEditObjectV1Request(paymentgatewayEditObjectV1Request).Execute()

Edit an existing Paymentgateway



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
	pkiPaymentgatewayID := int32(56) // int32 | The unique ID of the Paymentgateway
	paymentgatewayEditObjectV1Request := *openapiclient.NewPaymentgatewayEditObjectV1Request(*openapiclient.NewPaymentgatewayRequestCompound(openapiclient.Field-ePaymentgatewayProcessor("Moneris"), *openapiclient.NewMultilingualPaymentgatewayDescription())) // PaymentgatewayEditObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectPaymentgatewayAPI.PaymentgatewayEditObjectV1(context.Background(), pkiPaymentgatewayID).PaymentgatewayEditObjectV1Request(paymentgatewayEditObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectPaymentgatewayAPI.PaymentgatewayEditObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentgatewayEditObjectV1`: PaymentgatewayEditObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectPaymentgatewayAPI.PaymentgatewayEditObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiPaymentgatewayID** | **int32** | The unique ID of the Paymentgateway | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaymentgatewayEditObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paymentgatewayEditObjectV1Request** | [**PaymentgatewayEditObjectV1Request**](PaymentgatewayEditObjectV1Request.md) |  | 

### Return type

[**PaymentgatewayEditObjectV1Response**](PaymentgatewayEditObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentgatewayGetAutocompleteV2

> PaymentgatewayGetAutocompleteV2Response PaymentgatewayGetAutocompleteV2(ctx, sSelector).EFilterActive(eFilterActive).SQuery(sQuery).AcceptLanguage(acceptLanguage).Execute()

Retrieve Paymentgateways and IDs



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
	sSelector := "sSelector_example" // string | The type of Paymentgateways to return
	eFilterActive := "eFilterActive_example" // string | Specify which results we want to display. (optional) (default to "Active")
	sQuery := "sQuery_example" // string | Allow to filter the returned results (optional)
	acceptLanguage := openapiclient.Header-Accept-Language("*") // HeaderAcceptLanguage |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectPaymentgatewayAPI.PaymentgatewayGetAutocompleteV2(context.Background(), sSelector).EFilterActive(eFilterActive).SQuery(sQuery).AcceptLanguage(acceptLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectPaymentgatewayAPI.PaymentgatewayGetAutocompleteV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentgatewayGetAutocompleteV2`: PaymentgatewayGetAutocompleteV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectPaymentgatewayAPI.PaymentgatewayGetAutocompleteV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sSelector** | **string** | The type of Paymentgateways to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaymentgatewayGetAutocompleteV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eFilterActive** | **string** | Specify which results we want to display. | [default to &quot;Active&quot;]
 **sQuery** | **string** | Allow to filter the returned results | 
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 

### Return type

[**PaymentgatewayGetAutocompleteV2Response**](PaymentgatewayGetAutocompleteV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentgatewayGetListV1

> PaymentgatewayGetListV1Response PaymentgatewayGetListV1(ctx).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()

Retrieve Paymentgateway list



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
	resp, r, err := apiClient.ObjectPaymentgatewayAPI.PaymentgatewayGetListV1(context.Background()).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectPaymentgatewayAPI.PaymentgatewayGetListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentgatewayGetListV1`: PaymentgatewayGetListV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectPaymentgatewayAPI.PaymentgatewayGetListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentgatewayGetListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eOrderBy** | **string** | Specify how you want the results to be sorted | 
 **iRowMax** | **int32** |  | 
 **iRowOffset** | **int32** |  | [default to 0]
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 
 **sFilter** | **string** |  | 

### Return type

[**PaymentgatewayGetListV1Response**](PaymentgatewayGetListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentgatewayGetObjectV2

> PaymentgatewayGetObjectV2Response PaymentgatewayGetObjectV2(ctx, pkiPaymentgatewayID).Execute()

Retrieve an existing Paymentgateway



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
	pkiPaymentgatewayID := int32(56) // int32 | The unique ID of the Paymentgateway

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectPaymentgatewayAPI.PaymentgatewayGetObjectV2(context.Background(), pkiPaymentgatewayID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectPaymentgatewayAPI.PaymentgatewayGetObjectV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentgatewayGetObjectV2`: PaymentgatewayGetObjectV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectPaymentgatewayAPI.PaymentgatewayGetObjectV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiPaymentgatewayID** | **int32** | The unique ID of the Paymentgateway | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaymentgatewayGetObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PaymentgatewayGetObjectV2Response**](PaymentgatewayGetObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

