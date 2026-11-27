# eZmaxAPI\ObjectCustomertemplateAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomertemplateGetDocumenttypechecklistV1**](ObjectCustomertemplateAPI.md#CustomertemplateGetDocumenttypechecklistV1) | **Get** /1/object/customertemplate/{pkiCustomertemplateID}/getDocumenttypechecklist | Retrieve Document type checklist
[**CustomertemplateSyncDocumenttypechecklistV1**](ObjectCustomertemplateAPI.md#CustomertemplateSyncDocumenttypechecklistV1) | **Post** /1/object/customertemplate/{pkiCustomertemplateID}/syncDocumenttypechecklist | Synchronize Document type checklist



## CustomertemplateGetDocumenttypechecklistV1

> CustomertemplateGetDocumenttypechecklistV1Response CustomertemplateGetDocumenttypechecklistV1(ctx, pkiCustomertemplateID).Execute()

Retrieve Document type checklist

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
	pkiCustomertemplateID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectCustomertemplateAPI.CustomertemplateGetDocumenttypechecklistV1(context.Background(), pkiCustomertemplateID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectCustomertemplateAPI.CustomertemplateGetDocumenttypechecklistV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomertemplateGetDocumenttypechecklistV1`: CustomertemplateGetDocumenttypechecklistV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectCustomertemplateAPI.CustomertemplateGetDocumenttypechecklistV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiCustomertemplateID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomertemplateGetDocumenttypechecklistV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomertemplateGetDocumenttypechecklistV1Response**](CustomertemplateGetDocumenttypechecklistV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomertemplateSyncDocumenttypechecklistV1

> CustomertemplateSyncDocumenttypechecklistV1Response CustomertemplateSyncDocumenttypechecklistV1(ctx, pkiCustomertemplateID).Body(body).Execute()

Synchronize Document type checklist

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
	pkiCustomertemplateID := int32(56) // int32 | 
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectCustomertemplateAPI.CustomertemplateSyncDocumenttypechecklistV1(context.Background(), pkiCustomertemplateID).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectCustomertemplateAPI.CustomertemplateSyncDocumenttypechecklistV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomertemplateSyncDocumenttypechecklistV1`: CustomertemplateSyncDocumenttypechecklistV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectCustomertemplateAPI.CustomertemplateSyncDocumenttypechecklistV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiCustomertemplateID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomertemplateSyncDocumenttypechecklistV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

[**CustomertemplateSyncDocumenttypechecklistV1Response**](CustomertemplateSyncDocumenttypechecklistV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

