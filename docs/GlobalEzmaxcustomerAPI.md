# eZmaxAPI\GlobalEzmaxcustomerAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GlobalEzmaxcustomerGetConfigurationV1**](GlobalEzmaxcustomerAPI.md#GlobalEzmaxcustomerGetConfigurationV1) | **Get** /1/ezmaxcustomer/{pksEzmaxcustomerCode}/getConfiguration | Get ezmaxcustomer configuration



## GlobalEzmaxcustomerGetConfigurationV1

> GlobalEzmaxcustomerGetConfigurationV1Response GlobalEzmaxcustomerGetConfigurationV1(ctx, pksEzmaxcustomerCode).Execute()

Get ezmaxcustomer configuration



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
	pksEzmaxcustomerCode := "pksEzmaxcustomerCode_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlobalEzmaxcustomerAPI.GlobalEzmaxcustomerGetConfigurationV1(context.Background(), pksEzmaxcustomerCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalEzmaxcustomerAPI.GlobalEzmaxcustomerGetConfigurationV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GlobalEzmaxcustomerGetConfigurationV1`: GlobalEzmaxcustomerGetConfigurationV1Response
	fmt.Fprintf(os.Stdout, "Response from `GlobalEzmaxcustomerAPI.GlobalEzmaxcustomerGetConfigurationV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pksEzmaxcustomerCode** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGlobalEzmaxcustomerGetConfigurationV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GlobalEzmaxcustomerGetConfigurationV1Response**](GlobalEzmaxcustomerGetConfigurationV1Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

