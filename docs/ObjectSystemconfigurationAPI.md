# eZmaxAPI\ObjectSystemconfigurationAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SystemconfigurationEditObjectV2**](ObjectSystemconfigurationAPI.md#SystemconfigurationEditObjectV2) | **Put** /2/object/systemconfiguration/{pkiSystemconfigurationID} | Edit an existing Systemconfiguration
[**SystemconfigurationGetObjectV2**](ObjectSystemconfigurationAPI.md#SystemconfigurationGetObjectV2) | **Get** /2/object/systemconfiguration/{pkiSystemconfigurationID} | Retrieve an existing Systemconfiguration



## SystemconfigurationEditObjectV2

> SystemconfigurationEditObjectV2Response SystemconfigurationEditObjectV2(ctx, pkiSystemconfigurationID).SystemconfigurationEditObjectV2Request(systemconfigurationEditObjectV2Request).Execute()

Edit an existing Systemconfiguration



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
	pkiSystemconfigurationID := int32(56) // int32 | The unique ID of the Systemconfiguration
	systemconfigurationEditObjectV2Request := *openapiclient.NewSystemconfigurationEditObjectV2Request(*openapiclient.NewSystemconfigurationRequestCompoundV2(openapiclient.Field-eSystemconfigurationNewexternaluseraction("Stage"), openapiclient.Field-eSystemconfigurationLanguage1("fr_QC"), openapiclient.Field-eSystemconfigurationLanguage2("en_CA"), true, true, int32(1))) // SystemconfigurationEditObjectV2Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectSystemconfigurationAPI.SystemconfigurationEditObjectV2(context.Background(), pkiSystemconfigurationID).SystemconfigurationEditObjectV2Request(systemconfigurationEditObjectV2Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectSystemconfigurationAPI.SystemconfigurationEditObjectV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SystemconfigurationEditObjectV2`: SystemconfigurationEditObjectV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectSystemconfigurationAPI.SystemconfigurationEditObjectV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiSystemconfigurationID** | **int32** | The unique ID of the Systemconfiguration | 

### Other Parameters

Other parameters are passed through a pointer to a apiSystemconfigurationEditObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **systemconfigurationEditObjectV2Request** | [**SystemconfigurationEditObjectV2Request**](SystemconfigurationEditObjectV2Request.md) |  | 

### Return type

[**SystemconfigurationEditObjectV2Response**](SystemconfigurationEditObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SystemconfigurationGetObjectV2

> SystemconfigurationGetObjectV2Response SystemconfigurationGetObjectV2(ctx, pkiSystemconfigurationID).Execute()

Retrieve an existing Systemconfiguration



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
	pkiSystemconfigurationID := int32(56) // int32 | The unique ID of the Systemconfiguration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectSystemconfigurationAPI.SystemconfigurationGetObjectV2(context.Background(), pkiSystemconfigurationID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectSystemconfigurationAPI.SystemconfigurationGetObjectV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SystemconfigurationGetObjectV2`: SystemconfigurationGetObjectV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectSystemconfigurationAPI.SystemconfigurationGetObjectV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiSystemconfigurationID** | **int32** | The unique ID of the Systemconfiguration | 

### Other Parameters

Other parameters are passed through a pointer to a apiSystemconfigurationGetObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SystemconfigurationGetObjectV2Response**](SystemconfigurationGetObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

