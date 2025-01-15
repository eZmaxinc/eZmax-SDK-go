# eZmaxAPI\ObjectVariableexpenseAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VariableexpenseCreateObjectV1**](ObjectVariableexpenseAPI.md#VariableexpenseCreateObjectV1) | **Post** /1/object/variableexpense | Create a new Variableexpense
[**VariableexpenseEditObjectV1**](ObjectVariableexpenseAPI.md#VariableexpenseEditObjectV1) | **Put** /1/object/variableexpense/{pkiVariableexpenseID} | Edit an existing Variableexpense
[**VariableexpenseGetAutocompleteV2**](ObjectVariableexpenseAPI.md#VariableexpenseGetAutocompleteV2) | **Get** /2/object/variableexpense/getAutocomplete/{sSelector} | Retrieve Variableexpenses and IDs
[**VariableexpenseGetListV1**](ObjectVariableexpenseAPI.md#VariableexpenseGetListV1) | **Get** /1/object/variableexpense/getList | Retrieve Variableexpense list
[**VariableexpenseGetObjectV2**](ObjectVariableexpenseAPI.md#VariableexpenseGetObjectV2) | **Get** /2/object/variableexpense/{pkiVariableexpenseID} | Retrieve an existing Variableexpense



## VariableexpenseCreateObjectV1

> VariableexpenseCreateObjectV1Response VariableexpenseCreateObjectV1(ctx).VariableexpenseCreateObjectV1Request(variableexpenseCreateObjectV1Request).Execute()

Create a new Variableexpense



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
	variableexpenseCreateObjectV1Request := *openapiclient.NewVariableexpenseCreateObjectV1Request([]openapiclient.VariableexpenseRequestCompound{*openapiclient.NewVariableexpenseRequestCompound("EQBUR", *openapiclient.NewMultilingualVariableexpenseDescription(), openapiclient.Field-eVariableexpenseTaxable("Yes"), true)}) // VariableexpenseCreateObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectVariableexpenseAPI.VariableexpenseCreateObjectV1(context.Background()).VariableexpenseCreateObjectV1Request(variableexpenseCreateObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectVariableexpenseAPI.VariableexpenseCreateObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VariableexpenseCreateObjectV1`: VariableexpenseCreateObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectVariableexpenseAPI.VariableexpenseCreateObjectV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVariableexpenseCreateObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **variableexpenseCreateObjectV1Request** | [**VariableexpenseCreateObjectV1Request**](VariableexpenseCreateObjectV1Request.md) |  | 

### Return type

[**VariableexpenseCreateObjectV1Response**](VariableexpenseCreateObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VariableexpenseEditObjectV1

> CommonResponse VariableexpenseEditObjectV1(ctx, pkiVariableexpenseID).VariableexpenseEditObjectV1Request(variableexpenseEditObjectV1Request).Execute()

Edit an existing Variableexpense



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
	pkiVariableexpenseID := int32(56) // int32 | 
	variableexpenseEditObjectV1Request := *openapiclient.NewVariableexpenseEditObjectV1Request(*openapiclient.NewVariableexpenseRequestCompound("EQBUR", *openapiclient.NewMultilingualVariableexpenseDescription(), openapiclient.Field-eVariableexpenseTaxable("Yes"), true)) // VariableexpenseEditObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectVariableexpenseAPI.VariableexpenseEditObjectV1(context.Background(), pkiVariableexpenseID).VariableexpenseEditObjectV1Request(variableexpenseEditObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectVariableexpenseAPI.VariableexpenseEditObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VariableexpenseEditObjectV1`: CommonResponse
	fmt.Fprintf(os.Stdout, "Response from `ObjectVariableexpenseAPI.VariableexpenseEditObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiVariableexpenseID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVariableexpenseEditObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **variableexpenseEditObjectV1Request** | [**VariableexpenseEditObjectV1Request**](VariableexpenseEditObjectV1Request.md) |  | 

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


## VariableexpenseGetAutocompleteV2

> VariableexpenseGetAutocompleteV2Response VariableexpenseGetAutocompleteV2(ctx, sSelector).EFilterActive(eFilterActive).SQuery(sQuery).AcceptLanguage(acceptLanguage).Execute()

Retrieve Variableexpenses and IDs



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
	sSelector := "sSelector_example" // string | The type of Variableexpenses to return
	eFilterActive := "eFilterActive_example" // string | Specify which results we want to display. (optional) (default to "Active")
	sQuery := "sQuery_example" // string | Allow to filter the returned results (optional)
	acceptLanguage := openapiclient.Header-Accept-Language("*") // HeaderAcceptLanguage |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectVariableexpenseAPI.VariableexpenseGetAutocompleteV2(context.Background(), sSelector).EFilterActive(eFilterActive).SQuery(sQuery).AcceptLanguage(acceptLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectVariableexpenseAPI.VariableexpenseGetAutocompleteV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VariableexpenseGetAutocompleteV2`: VariableexpenseGetAutocompleteV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectVariableexpenseAPI.VariableexpenseGetAutocompleteV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sSelector** | **string** | The type of Variableexpenses to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiVariableexpenseGetAutocompleteV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eFilterActive** | **string** | Specify which results we want to display. | [default to &quot;Active&quot;]
 **sQuery** | **string** | Allow to filter the returned results | 
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 

### Return type

[**VariableexpenseGetAutocompleteV2Response**](VariableexpenseGetAutocompleteV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VariableexpenseGetListV1

> VariableexpenseGetListV1Response VariableexpenseGetListV1(ctx).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()

Retrieve Variableexpense list



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
	resp, r, err := apiClient.ObjectVariableexpenseAPI.VariableexpenseGetListV1(context.Background()).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectVariableexpenseAPI.VariableexpenseGetListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VariableexpenseGetListV1`: VariableexpenseGetListV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectVariableexpenseAPI.VariableexpenseGetListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVariableexpenseGetListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eOrderBy** | **string** | Specify how you want the results to be sorted | 
 **iRowMax** | **int32** |  | 
 **iRowOffset** | **int32** |  | [default to 0]
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 
 **sFilter** | **string** |  | 

### Return type

[**VariableexpenseGetListV1Response**](VariableexpenseGetListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VariableexpenseGetObjectV2

> VariableexpenseGetObjectV2Response VariableexpenseGetObjectV2(ctx, pkiVariableexpenseID).Execute()

Retrieve an existing Variableexpense



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
	pkiVariableexpenseID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectVariableexpenseAPI.VariableexpenseGetObjectV2(context.Background(), pkiVariableexpenseID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectVariableexpenseAPI.VariableexpenseGetObjectV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VariableexpenseGetObjectV2`: VariableexpenseGetObjectV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectVariableexpenseAPI.VariableexpenseGetObjectV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiVariableexpenseID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVariableexpenseGetObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VariableexpenseGetObjectV2Response**](VariableexpenseGetObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

