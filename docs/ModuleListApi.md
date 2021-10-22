# eZmaxAPI\ModuleListApi

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListListpresentationV1**](ModuleListApi.md#ListListpresentationV1) | **Post** /1/module/list/listpresentation/{sListName} | Save all Listpresentation for a specific list



## ListListpresentationV1

> ListSaveListpresentationV1Response ListListpresentationV1(ctx, sListName).ListSaveListpresentationV1Request(listSaveListpresentationV1Request).Execute()

Save all Listpresentation for a specific list



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sListName := "sListName_example" // string | The list Name
    listSaveListpresentationV1Request := *openapiclient.NewListSaveListpresentationV1Request([]openapiclient.ListpresentationRequest{*openapiclient.NewListpresentationRequest("SListpresentationDescription_example", "bField1 eq true and iField2 gte 0 and iField2 lte 1000 and sField3 eq 'Other' and eField4 eq 'Paid' and sField5 like '%needle%'", "SListpresentationOrderby_example", []string{"ASColumnName_example"}, int32(100), int32(0))}) // ListSaveListpresentationV1Request | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ModuleListApi.ListListpresentationV1(context.Background(), sListName).ListSaveListpresentationV1Request(listSaveListpresentationV1Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModuleListApi.ListListpresentationV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListListpresentationV1`: ListSaveListpresentationV1Response
    fmt.Fprintf(os.Stdout, "Response from `ModuleListApi.ListListpresentationV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sListName** | **string** | The list Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListListpresentationV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **listSaveListpresentationV1Request** | [**ListSaveListpresentationV1Request**](ListSaveListpresentationV1Request.md) |  | 

### Return type

[**ListSaveListpresentationV1Response**](ListSaveListpresentationV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

