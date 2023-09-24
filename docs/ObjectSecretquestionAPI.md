# eZmaxAPI\ObjectSecretquestionAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SecretquestionGetAutocompleteV2**](ObjectSecretquestionAPI.md#SecretquestionGetAutocompleteV2) | **Get** /2/object/secretquestion/getAutocomplete/{sSelector} | Retrieve Secretquestions and IDs



## SecretquestionGetAutocompleteV2

> SecretquestionGetAutocompleteV2Response SecretquestionGetAutocompleteV2(ctx, sSelector).EFilterActive(eFilterActive).SQuery(sQuery).AcceptLanguage(acceptLanguage).Execute()

Retrieve Secretquestions and IDs



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
    sSelector := "sSelector_example" // string | The type of Secretquestions to return
    eFilterActive := "eFilterActive_example" // string | Specify which results we want to display. (optional) (default to "Active")
    sQuery := "sQuery_example" // string | Allow to filter the returned results (optional)
    acceptLanguage := openapiclient.Header-Accept-Language("*") // HeaderAcceptLanguage |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectSecretquestionAPI.SecretquestionGetAutocompleteV2(context.Background(), sSelector).EFilterActive(eFilterActive).SQuery(sQuery).AcceptLanguage(acceptLanguage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectSecretquestionAPI.SecretquestionGetAutocompleteV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecretquestionGetAutocompleteV2`: SecretquestionGetAutocompleteV2Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectSecretquestionAPI.SecretquestionGetAutocompleteV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sSelector** | **string** | The type of Secretquestions to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecretquestionGetAutocompleteV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eFilterActive** | **string** | Specify which results we want to display. | [default to &quot;Active&quot;]
 **sQuery** | **string** | Allow to filter the returned results | 
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 

### Return type

[**SecretquestionGetAutocompleteV2Response**](SecretquestionGetAutocompleteV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

