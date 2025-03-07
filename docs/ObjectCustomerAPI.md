# eZmaxAPI\ObjectCustomerAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomerCreateObjectV1**](ObjectCustomerAPI.md#CustomerCreateObjectV1) | **Post** /1/object/customer | Create a new Customer
[**CustomerGetAutocompleteV2**](ObjectCustomerAPI.md#CustomerGetAutocompleteV2) | **Get** /2/object/customer/getAutocomplete/{sSelector} | Retrieve Customers and IDs
[**CustomerGetObjectV2**](ObjectCustomerAPI.md#CustomerGetObjectV2) | **Get** /2/object/customer/{pkiCustomerID} | Retrieve an existing Customer



## CustomerCreateObjectV1

> CustomerCreateObjectV1Response CustomerCreateObjectV1(ctx).CustomerCreateObjectV1Request(customerCreateObjectV1Request).Execute()

Create a new Customer



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
	customerCreateObjectV1Request := *openapiclient.NewCustomerCreateObjectV1Request([]openapiclient.CustomerRequestCompound{*openapiclient.NewCustomerRequestCompound(int32(1), int32(229), "eZmax Solutions", int32(55), int32(150), int32(164), int32(66), int32(2), int32(21), int32(166), int32(36), int32(36), int32(223), int32(1), int32(170), int32(1), int32(26), int32(18), int32(66), int32(1), int32(26), int32(18), int32(66), int32(242), int32(107), "EZMA1", "4.00", int32(7085237), int32(12316524), int32(172), int32(193), "4.00", "2020-12-31", "2020-12-31 23:59:59", "2020-12-31 23:59:59", "2020-12-31 23:59:59", true, true, true, true, true, true, openapiclient.Field-eCustomerType("Normal"), openapiclient.Field-eCustomerMarketingcorrespondence("No"), true, true, "This is a comment")}) // CustomerCreateObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectCustomerAPI.CustomerCreateObjectV1(context.Background()).CustomerCreateObjectV1Request(customerCreateObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectCustomerAPI.CustomerCreateObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerCreateObjectV1`: CustomerCreateObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectCustomerAPI.CustomerCreateObjectV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomerCreateObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerCreateObjectV1Request** | [**CustomerCreateObjectV1Request**](CustomerCreateObjectV1Request.md) |  | 

### Return type

[**CustomerCreateObjectV1Response**](CustomerCreateObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerGetAutocompleteV2

> CustomerGetAutocompleteV2Response CustomerGetAutocompleteV2(ctx, sSelector).EFilterActive(eFilterActive).SQuery(sQuery).AcceptLanguage(acceptLanguage).Execute()

Retrieve Customers and IDs



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
	sSelector := "sSelector_example" // string | The type of Customers to return
	eFilterActive := "eFilterActive_example" // string | Specify which results we want to display. (optional) (default to "Active")
	sQuery := "sQuery_example" // string | Allow to filter the returned results (optional)
	acceptLanguage := openapiclient.Header-Accept-Language("*") // HeaderAcceptLanguage |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectCustomerAPI.CustomerGetAutocompleteV2(context.Background(), sSelector).EFilterActive(eFilterActive).SQuery(sQuery).AcceptLanguage(acceptLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectCustomerAPI.CustomerGetAutocompleteV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerGetAutocompleteV2`: CustomerGetAutocompleteV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectCustomerAPI.CustomerGetAutocompleteV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sSelector** | **string** | The type of Customers to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerGetAutocompleteV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eFilterActive** | **string** | Specify which results we want to display. | [default to &quot;Active&quot;]
 **sQuery** | **string** | Allow to filter the returned results | 
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 

### Return type

[**CustomerGetAutocompleteV2Response**](CustomerGetAutocompleteV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerGetObjectV2

> CustomerGetObjectV2Response CustomerGetObjectV2(ctx, pkiCustomerID).Execute()

Retrieve an existing Customer



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
	pkiCustomerID := int32(56) // int32 | The unique ID of the Customer

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectCustomerAPI.CustomerGetObjectV2(context.Background(), pkiCustomerID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectCustomerAPI.CustomerGetObjectV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerGetObjectV2`: CustomerGetObjectV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectCustomerAPI.CustomerGetObjectV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiCustomerID** | **int32** | The unique ID of the Customer | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerGetObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomerGetObjectV2Response**](CustomerGetObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

