# eZmaxAPI\ObjectCreditcardclientAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreditcardclientCreateObjectV1**](ObjectCreditcardclientAPI.md#CreditcardclientCreateObjectV1) | **Post** /1/object/creditcardclient | Create a new Creditcardclient
[**CreditcardclientDeleteObjectV1**](ObjectCreditcardclientAPI.md#CreditcardclientDeleteObjectV1) | **Delete** /1/object/creditcardclient/{pkiCreditcardclientID} | Delete an existing Creditcardclient
[**CreditcardclientEditObjectV1**](ObjectCreditcardclientAPI.md#CreditcardclientEditObjectV1) | **Put** /1/object/creditcardclient/{pkiCreditcardclientID} | Edit an existing Creditcardclient
[**CreditcardclientGetAutocompleteV2**](ObjectCreditcardclientAPI.md#CreditcardclientGetAutocompleteV2) | **Get** /2/object/creditcardclient/getAutocomplete/{sSelector} | Retrieve Creditcardclients and IDs
[**CreditcardclientGetListV1**](ObjectCreditcardclientAPI.md#CreditcardclientGetListV1) | **Get** /1/object/creditcardclient/getList | Retrieve Creditcardclient list
[**CreditcardclientGetObjectV2**](ObjectCreditcardclientAPI.md#CreditcardclientGetObjectV2) | **Get** /2/object/creditcardclient/{pkiCreditcardclientID} | Retrieve an existing Creditcardclient
[**CreditcardclientPatchObjectV1**](ObjectCreditcardclientAPI.md#CreditcardclientPatchObjectV1) | **Patch** /1/object/creditcardclient/{pkiCreditcardclientID} | Patch an existing Creditcardclient



## CreditcardclientCreateObjectV1

> CreditcardclientCreateObjectV1Response CreditcardclientCreateObjectV1(ctx).CreditcardclientCreateObjectV1Request(creditcardclientCreateObjectV1Request).Execute()

Create a new Creditcardclient



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
	creditcardclientCreateObjectV1Request := *openapiclient.NewCreditcardclientCreateObjectV1Request([]openapiclient.CreditcardclientRequestCompound{*openapiclient.NewCreditcardclientRequestCompound(true, "Visa", true, true, true, *openapiclient.NewCreditcarddetailRequest(int32(10), int32(2024), "2500", "Daniel-Johnson Blvd.", "H7T 2P6"), "SCreditcardclientCVV_example")}) // CreditcardclientCreateObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectCreditcardclientAPI.CreditcardclientCreateObjectV1(context.Background()).CreditcardclientCreateObjectV1Request(creditcardclientCreateObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectCreditcardclientAPI.CreditcardclientCreateObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreditcardclientCreateObjectV1`: CreditcardclientCreateObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectCreditcardclientAPI.CreditcardclientCreateObjectV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreditcardclientCreateObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **creditcardclientCreateObjectV1Request** | [**CreditcardclientCreateObjectV1Request**](CreditcardclientCreateObjectV1Request.md) |  | 

### Return type

[**CreditcardclientCreateObjectV1Response**](CreditcardclientCreateObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreditcardclientDeleteObjectV1

> CommonResponse CreditcardclientDeleteObjectV1(ctx, pkiCreditcardclientID).Execute()

Delete an existing Creditcardclient



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
	pkiCreditcardclientID := int32(56) // int32 | The unique ID of the Creditcardclient

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectCreditcardclientAPI.CreditcardclientDeleteObjectV1(context.Background(), pkiCreditcardclientID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectCreditcardclientAPI.CreditcardclientDeleteObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreditcardclientDeleteObjectV1`: CommonResponse
	fmt.Fprintf(os.Stdout, "Response from `ObjectCreditcardclientAPI.CreditcardclientDeleteObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiCreditcardclientID** | **int32** | The unique ID of the Creditcardclient | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreditcardclientDeleteObjectV1Request struct via the builder pattern


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


## CreditcardclientEditObjectV1

> CommonResponse CreditcardclientEditObjectV1(ctx, pkiCreditcardclientID).CreditcardclientEditObjectV1Request(creditcardclientEditObjectV1Request).Execute()

Edit an existing Creditcardclient



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
	pkiCreditcardclientID := int32(56) // int32 | The unique ID of the Creditcardclient
	creditcardclientEditObjectV1Request := *openapiclient.NewCreditcardclientEditObjectV1Request(*openapiclient.NewCreditcardclientRequestCompound(true, "Visa", true, true, true, *openapiclient.NewCreditcarddetailRequest(int32(10), int32(2024), "2500", "Daniel-Johnson Blvd.", "H7T 2P6"), "SCreditcardclientCVV_example")) // CreditcardclientEditObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectCreditcardclientAPI.CreditcardclientEditObjectV1(context.Background(), pkiCreditcardclientID).CreditcardclientEditObjectV1Request(creditcardclientEditObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectCreditcardclientAPI.CreditcardclientEditObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreditcardclientEditObjectV1`: CommonResponse
	fmt.Fprintf(os.Stdout, "Response from `ObjectCreditcardclientAPI.CreditcardclientEditObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiCreditcardclientID** | **int32** | The unique ID of the Creditcardclient | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreditcardclientEditObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **creditcardclientEditObjectV1Request** | [**CreditcardclientEditObjectV1Request**](CreditcardclientEditObjectV1Request.md) |  | 

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


## CreditcardclientGetAutocompleteV2

> CreditcardclientGetAutocompleteV2Response CreditcardclientGetAutocompleteV2(ctx, sSelector).EFilterActive(eFilterActive).SQuery(sQuery).AcceptLanguage(acceptLanguage).Execute()

Retrieve Creditcardclients and IDs



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
	sSelector := "sSelector_example" // string | The type of Creditcardclients to return
	eFilterActive := "eFilterActive_example" // string | Specify which results we want to display. (optional) (default to "Active")
	sQuery := "sQuery_example" // string | Allow to filter the returned results (optional)
	acceptLanguage := openapiclient.Header-Accept-Language("*") // HeaderAcceptLanguage |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectCreditcardclientAPI.CreditcardclientGetAutocompleteV2(context.Background(), sSelector).EFilterActive(eFilterActive).SQuery(sQuery).AcceptLanguage(acceptLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectCreditcardclientAPI.CreditcardclientGetAutocompleteV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreditcardclientGetAutocompleteV2`: CreditcardclientGetAutocompleteV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectCreditcardclientAPI.CreditcardclientGetAutocompleteV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sSelector** | **string** | The type of Creditcardclients to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreditcardclientGetAutocompleteV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eFilterActive** | **string** | Specify which results we want to display. | [default to &quot;Active&quot;]
 **sQuery** | **string** | Allow to filter the returned results | 
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 

### Return type

[**CreditcardclientGetAutocompleteV2Response**](CreditcardclientGetAutocompleteV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreditcardclientGetListV1

> CreditcardclientGetListV1Response CreditcardclientGetListV1(ctx).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()

Retrieve Creditcardclient list



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
	resp, r, err := apiClient.ObjectCreditcardclientAPI.CreditcardclientGetListV1(context.Background()).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectCreditcardclientAPI.CreditcardclientGetListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreditcardclientGetListV1`: CreditcardclientGetListV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectCreditcardclientAPI.CreditcardclientGetListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreditcardclientGetListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eOrderBy** | **string** | Specify how you want the results to be sorted | 
 **iRowMax** | **int32** |  | 
 **iRowOffset** | **int32** |  | [default to 0]
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 
 **sFilter** | **string** |  | 

### Return type

[**CreditcardclientGetListV1Response**](CreditcardclientGetListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreditcardclientGetObjectV2

> CreditcardclientGetObjectV2Response CreditcardclientGetObjectV2(ctx, pkiCreditcardclientID).Execute()

Retrieve an existing Creditcardclient



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
	pkiCreditcardclientID := int32(56) // int32 | The unique ID of the Creditcardclient

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectCreditcardclientAPI.CreditcardclientGetObjectV2(context.Background(), pkiCreditcardclientID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectCreditcardclientAPI.CreditcardclientGetObjectV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreditcardclientGetObjectV2`: CreditcardclientGetObjectV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectCreditcardclientAPI.CreditcardclientGetObjectV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiCreditcardclientID** | **int32** | The unique ID of the Creditcardclient | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreditcardclientGetObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreditcardclientGetObjectV2Response**](CreditcardclientGetObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreditcardclientPatchObjectV1

> CommonResponse CreditcardclientPatchObjectV1(ctx, pkiCreditcardclientID).CreditcardclientPatchObjectV1Request(creditcardclientPatchObjectV1Request).Execute()

Patch an existing Creditcardclient



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
	pkiCreditcardclientID := int32(56) // int32 | The unique ID of the Creditcardclient
	creditcardclientPatchObjectV1Request := *openapiclient.NewCreditcardclientPatchObjectV1Request(*openapiclient.NewCreditcardclientRequestPatch()) // CreditcardclientPatchObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectCreditcardclientAPI.CreditcardclientPatchObjectV1(context.Background(), pkiCreditcardclientID).CreditcardclientPatchObjectV1Request(creditcardclientPatchObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectCreditcardclientAPI.CreditcardclientPatchObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreditcardclientPatchObjectV1`: CommonResponse
	fmt.Fprintf(os.Stdout, "Response from `ObjectCreditcardclientAPI.CreditcardclientPatchObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiCreditcardclientID** | **int32** | The unique ID of the Creditcardclient | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreditcardclientPatchObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **creditcardclientPatchObjectV1Request** | [**CreditcardclientPatchObjectV1Request**](CreditcardclientPatchObjectV1Request.md) |  | 

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

