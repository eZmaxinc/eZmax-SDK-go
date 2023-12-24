# eZmaxAPI\ModuleReportAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReportGetReportFromCacheV1**](ModuleReportAPI.md#ReportGetReportFromCacheV1) | **Get** /1/module/report/getReportFromCache/{sReportgroupCacheID} | Retrieve report from cache



## ReportGetReportFromCacheV1

> CommonGetReportV1Response ReportGetReportFromCacheV1(ctx, sReportgroupCacheID).Execute()

Retrieve report from cache



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
	sReportgroupCacheID := "sReportgroupCacheID_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModuleReportAPI.ReportGetReportFromCacheV1(context.Background(), sReportgroupCacheID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModuleReportAPI.ReportGetReportFromCacheV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReportGetReportFromCacheV1`: CommonGetReportV1Response
	fmt.Fprintf(os.Stdout, "Response from `ModuleReportAPI.ReportGetReportFromCacheV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sReportgroupCacheID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportGetReportFromCacheV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CommonGetReportV1Response**](CommonGetReportV1Response.md)

### Authorization

[Authorization](../README.md#Authorization), [Presigned](../README.md#Presigned)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/pdf, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet, application/zip, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

