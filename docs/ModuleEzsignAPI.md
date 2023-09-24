# eZmaxAPI\ModuleEzsignAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EzsignSuggestSignersV1**](ModuleEzsignAPI.md#EzsignSuggestSignersV1) | **Get** /1/module/ezsign/suggestSigners | Suggest signers
[**EzsignSuggestTemplatesV1**](ModuleEzsignAPI.md#EzsignSuggestTemplatesV1) | **Get** /1/module/ezsign/suggestTemplates | Suggest templates



## EzsignSuggestSignersV1

> EzsignSuggestSignersV1Response EzsignSuggestSignersV1(ctx).Execute()

Suggest signers



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
    resp, r, err := apiClient.ModuleEzsignAPI.EzsignSuggestSignersV1(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModuleEzsignAPI.EzsignSuggestSignersV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsignSuggestSignersV1`: EzsignSuggestSignersV1Response
    fmt.Fprintf(os.Stdout, "Response from `ModuleEzsignAPI.EzsignSuggestSignersV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignSuggestSignersV1Request struct via the builder pattern


### Return type

[**EzsignSuggestSignersV1Response**](EzsignSuggestSignersV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignSuggestTemplatesV1

> EzsignSuggestTemplatesV1Response EzsignSuggestTemplatesV1(ctx).FkiEzsignfoldertypeID(fkiEzsignfoldertypeID).Execute()

Suggest templates



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
    fkiEzsignfoldertypeID := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ModuleEzsignAPI.EzsignSuggestTemplatesV1(context.Background()).FkiEzsignfoldertypeID(fkiEzsignfoldertypeID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModuleEzsignAPI.EzsignSuggestTemplatesV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsignSuggestTemplatesV1`: EzsignSuggestTemplatesV1Response
    fmt.Fprintf(os.Stdout, "Response from `ModuleEzsignAPI.EzsignSuggestTemplatesV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEzsignSuggestTemplatesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fkiEzsignfoldertypeID** | **int32** |  | 

### Return type

[**EzsignSuggestTemplatesV1Response**](EzsignSuggestTemplatesV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

