# eZmaxAPI\ObjectEzdoctemplatetypeAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EzdoctemplatetypeGetAutocompleteV2**](ObjectEzdoctemplatetypeAPI.md#EzdoctemplatetypeGetAutocompleteV2) | **Get** /2/object/ezdoctemplatetype/getAutocomplete/{sSelector} | Retrieve Ezdoctemplatetypes and IDs



## EzdoctemplatetypeGetAutocompleteV2

> EzdoctemplatetypeGetAutocompleteV2Response EzdoctemplatetypeGetAutocompleteV2(ctx, sSelector).EFilterActive(eFilterActive).SQuery(sQuery).AcceptLanguage(acceptLanguage).Execute()

Retrieve Ezdoctemplatetypes and IDs



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
	sSelector := "sSelector_example" // string | The type of Ezdoctemplatetypes to return
	eFilterActive := "eFilterActive_example" // string | Specify which results we want to display. (optional) (default to "Active")
	sQuery := "sQuery_example" // string | Allow to filter the returned results (optional)
	acceptLanguage := openapiclient.Header-Accept-Language("*") // HeaderAcceptLanguage |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzdoctemplatetypeAPI.EzdoctemplatetypeGetAutocompleteV2(context.Background(), sSelector).EFilterActive(eFilterActive).SQuery(sQuery).AcceptLanguage(acceptLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzdoctemplatetypeAPI.EzdoctemplatetypeGetAutocompleteV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzdoctemplatetypeGetAutocompleteV2`: EzdoctemplatetypeGetAutocompleteV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzdoctemplatetypeAPI.EzdoctemplatetypeGetAutocompleteV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sSelector** | **string** | The type of Ezdoctemplatetypes to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzdoctemplatetypeGetAutocompleteV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eFilterActive** | **string** | Specify which results we want to display. | [default to &quot;Active&quot;]
 **sQuery** | **string** | Allow to filter the returned results | 
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 

### Return type

[**EzdoctemplatetypeGetAutocompleteV2Response**](EzdoctemplatetypeGetAutocompleteV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

