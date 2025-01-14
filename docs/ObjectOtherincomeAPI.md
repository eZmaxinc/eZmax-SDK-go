# eZmaxAPI\ObjectOtherincomeAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OtherincomeGetCommunicationCountV1**](ObjectOtherincomeAPI.md#OtherincomeGetCommunicationCountV1) | **Get** /1/object/otherincome/{pkiOtherincomeID}/getCommunicationCount | Retrieve Communication count
[**OtherincomeGetCommunicationListV1**](ObjectOtherincomeAPI.md#OtherincomeGetCommunicationListV1) | **Get** /1/object/otherincome/{pkiOtherincomeID}/getCommunicationList | Retrieve Communication list
[**OtherincomeGetCommunicationrecipientsV1**](ObjectOtherincomeAPI.md#OtherincomeGetCommunicationrecipientsV1) | **Get** /1/object/otherincome/{pkiOtherincomeID}/getCommunicationrecipients | Retrieve Otherincome&#39;s Communicationrecipient
[**OtherincomeGetCommunicationsendersV1**](ObjectOtherincomeAPI.md#OtherincomeGetCommunicationsendersV1) | **Get** /1/object/otherincome/{pkiOtherincomeID}/getCommunicationsenders | Retrieve Otherincome&#39;s Communicationsender



## OtherincomeGetCommunicationCountV1

> OtherincomeGetCommunicationCountV1Response OtherincomeGetCommunicationCountV1(ctx, pkiOtherincomeID).Execute()

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
	pkiOtherincomeID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectOtherincomeAPI.OtherincomeGetCommunicationCountV1(context.Background(), pkiOtherincomeID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectOtherincomeAPI.OtherincomeGetCommunicationCountV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OtherincomeGetCommunicationCountV1`: OtherincomeGetCommunicationCountV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectOtherincomeAPI.OtherincomeGetCommunicationCountV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiOtherincomeID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOtherincomeGetCommunicationCountV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OtherincomeGetCommunicationCountV1Response**](OtherincomeGetCommunicationCountV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OtherincomeGetCommunicationListV1

> OtherincomeGetCommunicationListV1Response OtherincomeGetCommunicationListV1(ctx, pkiOtherincomeID).Execute()

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
	pkiOtherincomeID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectOtherincomeAPI.OtherincomeGetCommunicationListV1(context.Background(), pkiOtherincomeID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectOtherincomeAPI.OtherincomeGetCommunicationListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OtherincomeGetCommunicationListV1`: OtherincomeGetCommunicationListV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectOtherincomeAPI.OtherincomeGetCommunicationListV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiOtherincomeID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOtherincomeGetCommunicationListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OtherincomeGetCommunicationListV1Response**](OtherincomeGetCommunicationListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OtherincomeGetCommunicationrecipientsV1

> OtherincomeGetCommunicationrecipientsV1Response OtherincomeGetCommunicationrecipientsV1(ctx, pkiOtherincomeID).Execute()

Retrieve Otherincome's Communicationrecipient



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
	pkiOtherincomeID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectOtherincomeAPI.OtherincomeGetCommunicationrecipientsV1(context.Background(), pkiOtherincomeID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectOtherincomeAPI.OtherincomeGetCommunicationrecipientsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OtherincomeGetCommunicationrecipientsV1`: OtherincomeGetCommunicationrecipientsV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectOtherincomeAPI.OtherincomeGetCommunicationrecipientsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiOtherincomeID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOtherincomeGetCommunicationrecipientsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OtherincomeGetCommunicationrecipientsV1Response**](OtherincomeGetCommunicationrecipientsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OtherincomeGetCommunicationsendersV1

> OtherincomeGetCommunicationsendersV1Response OtherincomeGetCommunicationsendersV1(ctx, pkiOtherincomeID).Execute()

Retrieve Otherincome's Communicationsender



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
	pkiOtherincomeID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectOtherincomeAPI.OtherincomeGetCommunicationsendersV1(context.Background(), pkiOtherincomeID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectOtherincomeAPI.OtherincomeGetCommunicationsendersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OtherincomeGetCommunicationsendersV1`: OtherincomeGetCommunicationsendersV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectOtherincomeAPI.OtherincomeGetCommunicationsendersV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiOtherincomeID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOtherincomeGetCommunicationsendersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OtherincomeGetCommunicationsendersV1Response**](OtherincomeGetCommunicationsendersV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

