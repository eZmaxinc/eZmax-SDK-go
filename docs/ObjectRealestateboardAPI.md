# eZmaxAPI\ObjectRealestateboardAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RealestateboardGetAutocompleteV2**](ObjectRealestateboardAPI.md#RealestateboardGetAutocompleteV2) | **Get** /2/object/realestateboard/getAutocomplete/{sSelector} | Retrieve Realestateboards and IDs



## RealestateboardGetAutocompleteV2

> RealestateboardGetAutocompleteV2Response RealestateboardGetAutocompleteV2(ctx, sSelector).EFilterActive(eFilterActive).SQuery(sQuery).AcceptLanguage(acceptLanguage).FkiProvinceID(fkiProvinceID).Execute()

Retrieve Realestateboards and IDs



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
	sSelector := "sSelector_example" // string | The type of Realestateboards to return
	eFilterActive := "eFilterActive_example" // string | Specify which results we want to display. (optional) (default to "Active")
	sQuery := "sQuery_example" // string | Allow to filter the returned results (optional)
	acceptLanguage := openapiclient.Header-Accept-Language("*") // HeaderAcceptLanguage |  (optional)
	fkiProvinceID := "fkiProvinceID_example" // string | The province ID to filter the results expected (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectRealestateboardAPI.RealestateboardGetAutocompleteV2(context.Background(), sSelector).EFilterActive(eFilterActive).SQuery(sQuery).AcceptLanguage(acceptLanguage).FkiProvinceID(fkiProvinceID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectRealestateboardAPI.RealestateboardGetAutocompleteV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RealestateboardGetAutocompleteV2`: RealestateboardGetAutocompleteV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectRealestateboardAPI.RealestateboardGetAutocompleteV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sSelector** | **string** | The type of Realestateboards to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiRealestateboardGetAutocompleteV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eFilterActive** | **string** | Specify which results we want to display. | [default to &quot;Active&quot;]
 **sQuery** | **string** | Allow to filter the returned results | 
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 
 **fkiProvinceID** | **string** | The province ID to filter the results expected | 

### Return type

[**RealestateboardGetAutocompleteV2Response**](RealestateboardGetAutocompleteV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

