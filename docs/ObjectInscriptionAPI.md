# eZmaxAPI\ObjectInscriptionAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InscriptionGetAttachmentsV1**](ObjectInscriptionAPI.md#InscriptionGetAttachmentsV1) | **Get** /1/object/inscription/{pkiInscriptionID}/getAttachments | Retrieve Inscription&#39;s Attachments
[**InscriptionGetCommunicationCountV1**](ObjectInscriptionAPI.md#InscriptionGetCommunicationCountV1) | **Get** /1/object/inscription/{pkiInscriptionID}/getCommunicationCount | Retrieve Communication count
[**InscriptionGetCommunicationListV1**](ObjectInscriptionAPI.md#InscriptionGetCommunicationListV1) | **Get** /1/object/inscription/{pkiInscriptionID}/getCommunicationList | Retrieve Communication list
[**InscriptionGetCommunicationrecipientsV1**](ObjectInscriptionAPI.md#InscriptionGetCommunicationrecipientsV1) | **Get** /1/object/inscription/{pkiInscriptionID}/getCommunicationrecipients | Retrieve Inscription&#39;s Communicationrecipient
[**InscriptionGetCommunicationsendersV1**](ObjectInscriptionAPI.md#InscriptionGetCommunicationsendersV1) | **Get** /1/object/inscription/{pkiInscriptionID}/getCommunicationsenders | Retrieve Inscription&#39;s Communicationsender



## InscriptionGetAttachmentsV1

> InscriptionGetAttachmentsV1Response InscriptionGetAttachmentsV1(ctx, pkiInscriptionID).Execute()

Retrieve Inscription's Attachments



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
	pkiInscriptionID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectInscriptionAPI.InscriptionGetAttachmentsV1(context.Background(), pkiInscriptionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectInscriptionAPI.InscriptionGetAttachmentsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InscriptionGetAttachmentsV1`: InscriptionGetAttachmentsV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectInscriptionAPI.InscriptionGetAttachmentsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiInscriptionID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInscriptionGetAttachmentsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InscriptionGetAttachmentsV1Response**](InscriptionGetAttachmentsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InscriptionGetCommunicationCountV1

> InscriptionGetCommunicationCountV1Response InscriptionGetCommunicationCountV1(ctx, pkiInscriptionID).Execute()

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
	pkiInscriptionID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectInscriptionAPI.InscriptionGetCommunicationCountV1(context.Background(), pkiInscriptionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectInscriptionAPI.InscriptionGetCommunicationCountV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InscriptionGetCommunicationCountV1`: InscriptionGetCommunicationCountV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectInscriptionAPI.InscriptionGetCommunicationCountV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiInscriptionID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInscriptionGetCommunicationCountV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InscriptionGetCommunicationCountV1Response**](InscriptionGetCommunicationCountV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InscriptionGetCommunicationListV1

> InscriptionGetCommunicationListV1Response InscriptionGetCommunicationListV1(ctx, pkiInscriptionID).Execute()

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
	pkiInscriptionID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectInscriptionAPI.InscriptionGetCommunicationListV1(context.Background(), pkiInscriptionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectInscriptionAPI.InscriptionGetCommunicationListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InscriptionGetCommunicationListV1`: InscriptionGetCommunicationListV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectInscriptionAPI.InscriptionGetCommunicationListV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiInscriptionID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInscriptionGetCommunicationListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InscriptionGetCommunicationListV1Response**](InscriptionGetCommunicationListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InscriptionGetCommunicationrecipientsV1

> InscriptionGetCommunicationrecipientsV1Response InscriptionGetCommunicationrecipientsV1(ctx, pkiInscriptionID).Execute()

Retrieve Inscription's Communicationrecipient



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
	pkiInscriptionID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectInscriptionAPI.InscriptionGetCommunicationrecipientsV1(context.Background(), pkiInscriptionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectInscriptionAPI.InscriptionGetCommunicationrecipientsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InscriptionGetCommunicationrecipientsV1`: InscriptionGetCommunicationrecipientsV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectInscriptionAPI.InscriptionGetCommunicationrecipientsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiInscriptionID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInscriptionGetCommunicationrecipientsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InscriptionGetCommunicationrecipientsV1Response**](InscriptionGetCommunicationrecipientsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InscriptionGetCommunicationsendersV1

> InscriptionGetCommunicationsendersV1Response InscriptionGetCommunicationsendersV1(ctx, pkiInscriptionID).Execute()

Retrieve Inscription's Communicationsender



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
	pkiInscriptionID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectInscriptionAPI.InscriptionGetCommunicationsendersV1(context.Background(), pkiInscriptionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectInscriptionAPI.InscriptionGetCommunicationsendersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InscriptionGetCommunicationsendersV1`: InscriptionGetCommunicationsendersV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectInscriptionAPI.InscriptionGetCommunicationsendersV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiInscriptionID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInscriptionGetCommunicationsendersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InscriptionGetCommunicationsendersV1Response**](InscriptionGetCommunicationsendersV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

