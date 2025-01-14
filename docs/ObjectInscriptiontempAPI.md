# eZmaxAPI\ObjectInscriptiontempAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InscriptiontempGetCommunicationCountV1**](ObjectInscriptiontempAPI.md#InscriptiontempGetCommunicationCountV1) | **Get** /1/object/inscriptiontemp/{pkiInscriptiontempID}/getCommunicationCount | Retrieve Communication count
[**InscriptiontempGetCommunicationListV1**](ObjectInscriptiontempAPI.md#InscriptiontempGetCommunicationListV1) | **Get** /1/object/inscriptiontemp/{pkiInscriptiontempID}/getCommunicationList | Retrieve Communication list
[**InscriptiontempGetCommunicationrecipientsV1**](ObjectInscriptiontempAPI.md#InscriptiontempGetCommunicationrecipientsV1) | **Get** /1/object/inscriptiontemp/{pkiInscriptiontempID}/getCommunicationrecipients | Retrieve Inscriptiontemp&#39;s Communicationrecipient
[**InscriptiontempGetCommunicationsendersV1**](ObjectInscriptiontempAPI.md#InscriptiontempGetCommunicationsendersV1) | **Get** /1/object/inscriptiontemp/{pkiInscriptiontempID}/getCommunicationsenders | Retrieve Inscriptiontemp&#39;s Communicationsender



## InscriptiontempGetCommunicationCountV1

> InscriptiontempGetCommunicationCountV1Response InscriptiontempGetCommunicationCountV1(ctx, pkiInscriptiontempID).Execute()

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
	pkiInscriptiontempID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectInscriptiontempAPI.InscriptiontempGetCommunicationCountV1(context.Background(), pkiInscriptiontempID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectInscriptiontempAPI.InscriptiontempGetCommunicationCountV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InscriptiontempGetCommunicationCountV1`: InscriptiontempGetCommunicationCountV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectInscriptiontempAPI.InscriptiontempGetCommunicationCountV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiInscriptiontempID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInscriptiontempGetCommunicationCountV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InscriptiontempGetCommunicationCountV1Response**](InscriptiontempGetCommunicationCountV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InscriptiontempGetCommunicationListV1

> InscriptiontempGetCommunicationListV1Response InscriptiontempGetCommunicationListV1(ctx, pkiInscriptiontempID).Execute()

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
	pkiInscriptiontempID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectInscriptiontempAPI.InscriptiontempGetCommunicationListV1(context.Background(), pkiInscriptiontempID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectInscriptiontempAPI.InscriptiontempGetCommunicationListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InscriptiontempGetCommunicationListV1`: InscriptiontempGetCommunicationListV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectInscriptiontempAPI.InscriptiontempGetCommunicationListV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiInscriptiontempID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInscriptiontempGetCommunicationListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InscriptiontempGetCommunicationListV1Response**](InscriptiontempGetCommunicationListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InscriptiontempGetCommunicationrecipientsV1

> InscriptiontempGetCommunicationrecipientsV1Response InscriptiontempGetCommunicationrecipientsV1(ctx, pkiInscriptiontempID).Execute()

Retrieve Inscriptiontemp's Communicationrecipient



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
	pkiInscriptiontempID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectInscriptiontempAPI.InscriptiontempGetCommunicationrecipientsV1(context.Background(), pkiInscriptiontempID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectInscriptiontempAPI.InscriptiontempGetCommunicationrecipientsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InscriptiontempGetCommunicationrecipientsV1`: InscriptiontempGetCommunicationrecipientsV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectInscriptiontempAPI.InscriptiontempGetCommunicationrecipientsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiInscriptiontempID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInscriptiontempGetCommunicationrecipientsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InscriptiontempGetCommunicationrecipientsV1Response**](InscriptiontempGetCommunicationrecipientsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InscriptiontempGetCommunicationsendersV1

> InscriptiontempGetCommunicationsendersV1Response InscriptiontempGetCommunicationsendersV1(ctx, pkiInscriptiontempID).Execute()

Retrieve Inscriptiontemp's Communicationsender



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
	pkiInscriptiontempID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectInscriptiontempAPI.InscriptiontempGetCommunicationsendersV1(context.Background(), pkiInscriptiontempID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectInscriptiontempAPI.InscriptiontempGetCommunicationsendersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InscriptiontempGetCommunicationsendersV1`: InscriptiontempGetCommunicationsendersV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectInscriptiontempAPI.InscriptiontempGetCommunicationsendersV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiInscriptiontempID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInscriptiontempGetCommunicationsendersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InscriptiontempGetCommunicationsendersV1Response**](InscriptiontempGetCommunicationsendersV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

