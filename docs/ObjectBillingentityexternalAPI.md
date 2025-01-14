# eZmaxAPI\ObjectBillingentityexternalAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BillingentityexternalGenerateFederationTokenV1**](ObjectBillingentityexternalAPI.md#BillingentityexternalGenerateFederationTokenV1) | **Post** /1/object/billingentityexternal/{pkiBillingentityexternalID}/generateFederationToken | Generate a federation token
[**BillingentityexternalGetAutocompleteV2**](ObjectBillingentityexternalAPI.md#BillingentityexternalGetAutocompleteV2) | **Get** /2/object/billingentityexternal/getAutocomplete/{sSelector} | Retrieve Billingentityexternals and IDs



## BillingentityexternalGenerateFederationTokenV1

> BillingentityexternalGenerateFederationTokenV1Response BillingentityexternalGenerateFederationTokenV1(ctx, pkiBillingentityexternalID).BillingentityexternalGenerateFederationTokenV1Request(billingentityexternalGenerateFederationTokenV1Request).Execute()

Generate a federation token



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
	pkiBillingentityexternalID := int32(56) // int32 | 
	billingentityexternalGenerateFederationTokenV1Request := *openapiclient.NewBillingentityexternalGenerateFederationTokenV1Request("demo") // BillingentityexternalGenerateFederationTokenV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectBillingentityexternalAPI.BillingentityexternalGenerateFederationTokenV1(context.Background(), pkiBillingentityexternalID).BillingentityexternalGenerateFederationTokenV1Request(billingentityexternalGenerateFederationTokenV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectBillingentityexternalAPI.BillingentityexternalGenerateFederationTokenV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingentityexternalGenerateFederationTokenV1`: BillingentityexternalGenerateFederationTokenV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectBillingentityexternalAPI.BillingentityexternalGenerateFederationTokenV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiBillingentityexternalID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBillingentityexternalGenerateFederationTokenV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **billingentityexternalGenerateFederationTokenV1Request** | [**BillingentityexternalGenerateFederationTokenV1Request**](BillingentityexternalGenerateFederationTokenV1Request.md) |  | 

### Return type

[**BillingentityexternalGenerateFederationTokenV1Response**](BillingentityexternalGenerateFederationTokenV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingentityexternalGetAutocompleteV2

> BillingentityexternalGetAutocompleteV2Response BillingentityexternalGetAutocompleteV2(ctx, sSelector).EFilterActive(eFilterActive).SQuery(sQuery).AcceptLanguage(acceptLanguage).Execute()

Retrieve Billingentityexternals and IDs



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
	sSelector := "sSelector_example" // string | The type of Billingentityexternals to return
	eFilterActive := "eFilterActive_example" // string | Specify which results we want to display. (optional) (default to "Active")
	sQuery := "sQuery_example" // string | Allow to filter the returned results (optional)
	acceptLanguage := openapiclient.Header-Accept-Language("*") // HeaderAcceptLanguage |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectBillingentityexternalAPI.BillingentityexternalGetAutocompleteV2(context.Background(), sSelector).EFilterActive(eFilterActive).SQuery(sQuery).AcceptLanguage(acceptLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectBillingentityexternalAPI.BillingentityexternalGetAutocompleteV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingentityexternalGetAutocompleteV2`: BillingentityexternalGetAutocompleteV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectBillingentityexternalAPI.BillingentityexternalGetAutocompleteV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sSelector** | **string** | The type of Billingentityexternals to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiBillingentityexternalGetAutocompleteV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eFilterActive** | **string** | Specify which results we want to display. | [default to &quot;Active&quot;]
 **sQuery** | **string** | Allow to filter the returned results | 
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 

### Return type

[**BillingentityexternalGetAutocompleteV2Response**](BillingentityexternalGetAutocompleteV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

