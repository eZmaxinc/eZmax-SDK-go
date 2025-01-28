# eZmaxAPI\ObjectCreditcardmerchantAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreditcardmerchantCreateObjectV1**](ObjectCreditcardmerchantAPI.md#CreditcardmerchantCreateObjectV1) | **Post** /1/object/creditcardmerchant | Create a new Creditcardmerchant
[**CreditcardmerchantDeleteObjectV1**](ObjectCreditcardmerchantAPI.md#CreditcardmerchantDeleteObjectV1) | **Delete** /1/object/creditcardmerchant/{pkiCreditcardmerchantID} | Delete an existing Creditcardmerchant
[**CreditcardmerchantEditObjectV1**](ObjectCreditcardmerchantAPI.md#CreditcardmerchantEditObjectV1) | **Put** /1/object/creditcardmerchant/{pkiCreditcardmerchantID} | Edit an existing Creditcardmerchant
[**CreditcardmerchantGetAutocompleteV2**](ObjectCreditcardmerchantAPI.md#CreditcardmerchantGetAutocompleteV2) | **Get** /2/object/creditcardmerchant/getAutocomplete/{sSelector} | Retrieve Creditcardmerchants and IDs
[**CreditcardmerchantGetListV1**](ObjectCreditcardmerchantAPI.md#CreditcardmerchantGetListV1) | **Get** /1/object/creditcardmerchant/getList | Retrieve Creditcardmerchant list
[**CreditcardmerchantGetObjectV2**](ObjectCreditcardmerchantAPI.md#CreditcardmerchantGetObjectV2) | **Get** /2/object/creditcardmerchant/{pkiCreditcardmerchantID} | Retrieve an existing Creditcardmerchant



## CreditcardmerchantCreateObjectV1

> CreditcardmerchantCreateObjectV1Response CreditcardmerchantCreateObjectV1(ctx).CreditcardmerchantCreateObjectV1Request(creditcardmerchantCreateObjectV1Request).Execute()

Create a new Creditcardmerchant



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
	creditcardmerchantCreateObjectV1Request := *openapiclient.NewCreditcardmerchantCreateObjectV1Request([]openapiclient.CreditcardmerchantRequestCompound{*openapiclient.NewCreditcardmerchantRequestCompound(int32(46), true, true, true, true, "Moneris", "REPLACEME")}) // CreditcardmerchantCreateObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectCreditcardmerchantAPI.CreditcardmerchantCreateObjectV1(context.Background()).CreditcardmerchantCreateObjectV1Request(creditcardmerchantCreateObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectCreditcardmerchantAPI.CreditcardmerchantCreateObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreditcardmerchantCreateObjectV1`: CreditcardmerchantCreateObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectCreditcardmerchantAPI.CreditcardmerchantCreateObjectV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreditcardmerchantCreateObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **creditcardmerchantCreateObjectV1Request** | [**CreditcardmerchantCreateObjectV1Request**](CreditcardmerchantCreateObjectV1Request.md) |  | 

### Return type

[**CreditcardmerchantCreateObjectV1Response**](CreditcardmerchantCreateObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreditcardmerchantDeleteObjectV1

> CreditcardmerchantDeleteObjectV1Response CreditcardmerchantDeleteObjectV1(ctx, pkiCreditcardmerchantID).Execute()

Delete an existing Creditcardmerchant



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
	pkiCreditcardmerchantID := int32(56) // int32 | The unique ID of the Creditcardmerchant

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectCreditcardmerchantAPI.CreditcardmerchantDeleteObjectV1(context.Background(), pkiCreditcardmerchantID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectCreditcardmerchantAPI.CreditcardmerchantDeleteObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreditcardmerchantDeleteObjectV1`: CreditcardmerchantDeleteObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectCreditcardmerchantAPI.CreditcardmerchantDeleteObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiCreditcardmerchantID** | **int32** | The unique ID of the Creditcardmerchant | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreditcardmerchantDeleteObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreditcardmerchantDeleteObjectV1Response**](CreditcardmerchantDeleteObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreditcardmerchantEditObjectV1

> CreditcardmerchantEditObjectV1Response CreditcardmerchantEditObjectV1(ctx, pkiCreditcardmerchantID).CreditcardmerchantEditObjectV1Request(creditcardmerchantEditObjectV1Request).Execute()

Edit an existing Creditcardmerchant



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
	pkiCreditcardmerchantID := int32(56) // int32 | The unique ID of the Creditcardmerchant
	creditcardmerchantEditObjectV1Request := *openapiclient.NewCreditcardmerchantEditObjectV1Request(*openapiclient.NewCreditcardmerchantRequestCompound(int32(46), true, true, true, true, "Moneris", "REPLACEME")) // CreditcardmerchantEditObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectCreditcardmerchantAPI.CreditcardmerchantEditObjectV1(context.Background(), pkiCreditcardmerchantID).CreditcardmerchantEditObjectV1Request(creditcardmerchantEditObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectCreditcardmerchantAPI.CreditcardmerchantEditObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreditcardmerchantEditObjectV1`: CreditcardmerchantEditObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectCreditcardmerchantAPI.CreditcardmerchantEditObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiCreditcardmerchantID** | **int32** | The unique ID of the Creditcardmerchant | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreditcardmerchantEditObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **creditcardmerchantEditObjectV1Request** | [**CreditcardmerchantEditObjectV1Request**](CreditcardmerchantEditObjectV1Request.md) |  | 

### Return type

[**CreditcardmerchantEditObjectV1Response**](CreditcardmerchantEditObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreditcardmerchantGetAutocompleteV2

> CreditcardmerchantGetAutocompleteV2Response CreditcardmerchantGetAutocompleteV2(ctx, sSelector).EFilterActive(eFilterActive).SQuery(sQuery).AcceptLanguage(acceptLanguage).Execute()

Retrieve Creditcardmerchants and IDs



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
	sSelector := "sSelector_example" // string | The type of Creditcardmerchants to return
	eFilterActive := "eFilterActive_example" // string | Specify which results we want to display. (optional) (default to "Active")
	sQuery := "sQuery_example" // string | Allow to filter the returned results (optional)
	acceptLanguage := openapiclient.Header-Accept-Language("*") // HeaderAcceptLanguage |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectCreditcardmerchantAPI.CreditcardmerchantGetAutocompleteV2(context.Background(), sSelector).EFilterActive(eFilterActive).SQuery(sQuery).AcceptLanguage(acceptLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectCreditcardmerchantAPI.CreditcardmerchantGetAutocompleteV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreditcardmerchantGetAutocompleteV2`: CreditcardmerchantGetAutocompleteV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectCreditcardmerchantAPI.CreditcardmerchantGetAutocompleteV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sSelector** | **string** | The type of Creditcardmerchants to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreditcardmerchantGetAutocompleteV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eFilterActive** | **string** | Specify which results we want to display. | [default to &quot;Active&quot;]
 **sQuery** | **string** | Allow to filter the returned results | 
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 

### Return type

[**CreditcardmerchantGetAutocompleteV2Response**](CreditcardmerchantGetAutocompleteV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreditcardmerchantGetListV1

> CreditcardmerchantGetListV1Response CreditcardmerchantGetListV1(ctx).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()

Retrieve Creditcardmerchant list



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
	resp, r, err := apiClient.ObjectCreditcardmerchantAPI.CreditcardmerchantGetListV1(context.Background()).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectCreditcardmerchantAPI.CreditcardmerchantGetListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreditcardmerchantGetListV1`: CreditcardmerchantGetListV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectCreditcardmerchantAPI.CreditcardmerchantGetListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreditcardmerchantGetListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eOrderBy** | **string** | Specify how you want the results to be sorted | 
 **iRowMax** | **int32** |  | 
 **iRowOffset** | **int32** |  | [default to 0]
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 
 **sFilter** | **string** |  | 

### Return type

[**CreditcardmerchantGetListV1Response**](CreditcardmerchantGetListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreditcardmerchantGetObjectV2

> CreditcardmerchantGetObjectV2Response CreditcardmerchantGetObjectV2(ctx, pkiCreditcardmerchantID).Execute()

Retrieve an existing Creditcardmerchant



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
	pkiCreditcardmerchantID := int32(56) // int32 | The unique ID of the Creditcardmerchant

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectCreditcardmerchantAPI.CreditcardmerchantGetObjectV2(context.Background(), pkiCreditcardmerchantID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectCreditcardmerchantAPI.CreditcardmerchantGetObjectV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreditcardmerchantGetObjectV2`: CreditcardmerchantGetObjectV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectCreditcardmerchantAPI.CreditcardmerchantGetObjectV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiCreditcardmerchantID** | **int32** | The unique ID of the Creditcardmerchant | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreditcardmerchantGetObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreditcardmerchantGetObjectV2Response**](CreditcardmerchantGetObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

