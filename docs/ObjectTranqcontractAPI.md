# eZmaxAPI\ObjectTranqcontractAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TranqcontractGetCommunicationCountV1**](ObjectTranqcontractAPI.md#TranqcontractGetCommunicationCountV1) | **Get** /1/object/tranqcontract/{pkiTranqcontractID}/getCommunicationCount | Retrieve Communication count
[**TranqcontractGetCommunicationListV1**](ObjectTranqcontractAPI.md#TranqcontractGetCommunicationListV1) | **Get** /1/object/tranqcontract/{pkiTranqcontractID}/getCommunicationList | Retrieve Communication list
[**TranqcontractGetCommunicationrecipientsV1**](ObjectTranqcontractAPI.md#TranqcontractGetCommunicationrecipientsV1) | **Get** /1/object/tranqcontract/{pkiTranqcontractID}/getCommunicationrecipients | Retrieve Tranqcontract&#39;s Communicationrecipient
[**TranqcontractGetCommunicationsendersV1**](ObjectTranqcontractAPI.md#TranqcontractGetCommunicationsendersV1) | **Get** /1/object/tranqcontract/{pkiTranqcontractID}/getCommunicationsenders | Retrieve Tranqcontract&#39;s Communicationsender



## TranqcontractGetCommunicationCountV1

> TranqcontractGetCommunicationCountV1Response TranqcontractGetCommunicationCountV1(ctx, pkiTranqcontractID).Execute()

Retrieve Communication count



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
	pkiTranqcontractID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectTranqcontractAPI.TranqcontractGetCommunicationCountV1(context.Background(), pkiTranqcontractID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectTranqcontractAPI.TranqcontractGetCommunicationCountV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TranqcontractGetCommunicationCountV1`: TranqcontractGetCommunicationCountV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectTranqcontractAPI.TranqcontractGetCommunicationCountV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiTranqcontractID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTranqcontractGetCommunicationCountV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TranqcontractGetCommunicationCountV1Response**](TranqcontractGetCommunicationCountV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TranqcontractGetCommunicationListV1

> TranqcontractGetCommunicationListV1Response TranqcontractGetCommunicationListV1(ctx, pkiTranqcontractID).Execute()

Retrieve Communication list



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
	pkiTranqcontractID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectTranqcontractAPI.TranqcontractGetCommunicationListV1(context.Background(), pkiTranqcontractID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectTranqcontractAPI.TranqcontractGetCommunicationListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TranqcontractGetCommunicationListV1`: TranqcontractGetCommunicationListV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectTranqcontractAPI.TranqcontractGetCommunicationListV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiTranqcontractID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTranqcontractGetCommunicationListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TranqcontractGetCommunicationListV1Response**](TranqcontractGetCommunicationListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TranqcontractGetCommunicationrecipientsV1

> TranqcontractGetCommunicationrecipientsV1Response TranqcontractGetCommunicationrecipientsV1(ctx, pkiTranqcontractID).Execute()

Retrieve Tranqcontract's Communicationrecipient



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
	pkiTranqcontractID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectTranqcontractAPI.TranqcontractGetCommunicationrecipientsV1(context.Background(), pkiTranqcontractID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectTranqcontractAPI.TranqcontractGetCommunicationrecipientsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TranqcontractGetCommunicationrecipientsV1`: TranqcontractGetCommunicationrecipientsV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectTranqcontractAPI.TranqcontractGetCommunicationrecipientsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiTranqcontractID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTranqcontractGetCommunicationrecipientsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TranqcontractGetCommunicationrecipientsV1Response**](TranqcontractGetCommunicationrecipientsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TranqcontractGetCommunicationsendersV1

> TranqcontractGetCommunicationsendersV1Response TranqcontractGetCommunicationsendersV1(ctx, pkiTranqcontractID).Execute()

Retrieve Tranqcontract's Communicationsender



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
	pkiTranqcontractID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectTranqcontractAPI.TranqcontractGetCommunicationsendersV1(context.Background(), pkiTranqcontractID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectTranqcontractAPI.TranqcontractGetCommunicationsendersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TranqcontractGetCommunicationsendersV1`: TranqcontractGetCommunicationsendersV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectTranqcontractAPI.TranqcontractGetCommunicationsendersV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiTranqcontractID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTranqcontractGetCommunicationsendersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TranqcontractGetCommunicationsendersV1Response**](TranqcontractGetCommunicationsendersV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

