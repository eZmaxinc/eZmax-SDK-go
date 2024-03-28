# eZmaxAPI\ObjectEzmaxinvoicingAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EzmaxinvoicingGetAutocompleteV2**](ObjectEzmaxinvoicingAPI.md#EzmaxinvoicingGetAutocompleteV2) | **Get** /2/object/ezmaxinvoicing/getAutocomplete/{sSelector} | Retrieve Ezmaxinvoicings and IDs
[**EzmaxinvoicingGetObjectV2**](ObjectEzmaxinvoicingAPI.md#EzmaxinvoicingGetObjectV2) | **Get** /2/object/ezmaxinvoicing/{pkiEzmaxinvoicingID} | Retrieve an existing Ezmaxinvoicing
[**EzmaxinvoicingGetProvisionalV1**](ObjectEzmaxinvoicingAPI.md#EzmaxinvoicingGetProvisionalV1) | **Get** /1/object/ezmaxinvoicing/getProvisional | Retrieve provisional Ezmaxinvoicing



## EzmaxinvoicingGetAutocompleteV2

> EzmaxinvoicingGetAutocompleteV2Response EzmaxinvoicingGetAutocompleteV2(ctx, sSelector).EFilterActive(eFilterActive).SQuery(sQuery).AcceptLanguage(acceptLanguage).Execute()

Retrieve Ezmaxinvoicings and IDs



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
	sSelector := "sSelector_example" // string | The type of Ezmaxinvoicings to return
	eFilterActive := "eFilterActive_example" // string | Specify which results we want to display. (optional) (default to "Active")
	sQuery := "sQuery_example" // string | Allow to filter the returned results (optional)
	acceptLanguage := openapiclient.Header-Accept-Language("*") // HeaderAcceptLanguage |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzmaxinvoicingAPI.EzmaxinvoicingGetAutocompleteV2(context.Background(), sSelector).EFilterActive(eFilterActive).SQuery(sQuery).AcceptLanguage(acceptLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzmaxinvoicingAPI.EzmaxinvoicingGetAutocompleteV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzmaxinvoicingGetAutocompleteV2`: EzmaxinvoicingGetAutocompleteV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzmaxinvoicingAPI.EzmaxinvoicingGetAutocompleteV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sSelector** | **string** | The type of Ezmaxinvoicings to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzmaxinvoicingGetAutocompleteV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eFilterActive** | **string** | Specify which results we want to display. | [default to &quot;Active&quot;]
 **sQuery** | **string** | Allow to filter the returned results | 
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 

### Return type

[**EzmaxinvoicingGetAutocompleteV2Response**](EzmaxinvoicingGetAutocompleteV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzmaxinvoicingGetObjectV2

> EzmaxinvoicingGetObjectV2Response EzmaxinvoicingGetObjectV2(ctx, pkiEzmaxinvoicingID).Execute()

Retrieve an existing Ezmaxinvoicing



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
	pkiEzmaxinvoicingID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzmaxinvoicingAPI.EzmaxinvoicingGetObjectV2(context.Background(), pkiEzmaxinvoicingID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzmaxinvoicingAPI.EzmaxinvoicingGetObjectV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzmaxinvoicingGetObjectV2`: EzmaxinvoicingGetObjectV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzmaxinvoicingAPI.EzmaxinvoicingGetObjectV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzmaxinvoicingID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzmaxinvoicingGetObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzmaxinvoicingGetObjectV2Response**](EzmaxinvoicingGetObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzmaxinvoicingGetProvisionalV1

> EzmaxinvoicingGetProvisionalV1Response EzmaxinvoicingGetProvisionalV1(ctx).Execute()

Retrieve provisional Ezmaxinvoicing



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzmaxinvoicingAPI.EzmaxinvoicingGetProvisionalV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzmaxinvoicingAPI.EzmaxinvoicingGetProvisionalV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzmaxinvoicingGetProvisionalV1`: EzmaxinvoicingGetProvisionalV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzmaxinvoicingAPI.EzmaxinvoicingGetProvisionalV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiEzmaxinvoicingGetProvisionalV1Request struct via the builder pattern


### Return type

[**EzmaxinvoicingGetProvisionalV1Response**](EzmaxinvoicingGetProvisionalV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

