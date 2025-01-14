# eZmaxAPI\ObjectInscriptionnotauthenticatedAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InscriptionnotauthenticatedGetCommunicationCountV1**](ObjectInscriptionnotauthenticatedAPI.md#InscriptionnotauthenticatedGetCommunicationCountV1) | **Get** /1/object/inscriptionnotauthenticated/{pkiInscriptionnotauthenticatedID}/getCommunicationCount | Retrieve Communication count
[**InscriptionnotauthenticatedGetCommunicationListV1**](ObjectInscriptionnotauthenticatedAPI.md#InscriptionnotauthenticatedGetCommunicationListV1) | **Get** /1/object/inscriptionnotauthenticated/{pkiInscriptionnotauthenticatedID}/getCommunicationList | Retrieve Communication list
[**InscriptionnotauthenticatedGetCommunicationrecipientsV1**](ObjectInscriptionnotauthenticatedAPI.md#InscriptionnotauthenticatedGetCommunicationrecipientsV1) | **Get** /1/object/inscriptionnotauthenticated/{pkiInscriptionnotauthenticatedID}/getCommunicationrecipients | Retrieve Inscriptionnotauthenticated&#39;s Communicationrecipient
[**InscriptionnotauthenticatedGetCommunicationsendersV1**](ObjectInscriptionnotauthenticatedAPI.md#InscriptionnotauthenticatedGetCommunicationsendersV1) | **Get** /1/object/inscriptionnotauthenticated/{pkiInscriptionnotauthenticatedID}/getCommunicationsenders | Retrieve Inscriptionnotauthenticated&#39;s Communicationsender



## InscriptionnotauthenticatedGetCommunicationCountV1

> InscriptionnotauthenticatedGetCommunicationCountV1Response InscriptionnotauthenticatedGetCommunicationCountV1(ctx, pkiInscriptionnotauthenticatedID).Execute()

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
	pkiInscriptionnotauthenticatedID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectInscriptionnotauthenticatedAPI.InscriptionnotauthenticatedGetCommunicationCountV1(context.Background(), pkiInscriptionnotauthenticatedID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectInscriptionnotauthenticatedAPI.InscriptionnotauthenticatedGetCommunicationCountV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InscriptionnotauthenticatedGetCommunicationCountV1`: InscriptionnotauthenticatedGetCommunicationCountV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectInscriptionnotauthenticatedAPI.InscriptionnotauthenticatedGetCommunicationCountV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiInscriptionnotauthenticatedID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInscriptionnotauthenticatedGetCommunicationCountV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InscriptionnotauthenticatedGetCommunicationCountV1Response**](InscriptionnotauthenticatedGetCommunicationCountV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InscriptionnotauthenticatedGetCommunicationListV1

> InscriptionnotauthenticatedGetCommunicationListV1Response InscriptionnotauthenticatedGetCommunicationListV1(ctx, pkiInscriptionnotauthenticatedID).Execute()

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
	pkiInscriptionnotauthenticatedID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectInscriptionnotauthenticatedAPI.InscriptionnotauthenticatedGetCommunicationListV1(context.Background(), pkiInscriptionnotauthenticatedID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectInscriptionnotauthenticatedAPI.InscriptionnotauthenticatedGetCommunicationListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InscriptionnotauthenticatedGetCommunicationListV1`: InscriptionnotauthenticatedGetCommunicationListV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectInscriptionnotauthenticatedAPI.InscriptionnotauthenticatedGetCommunicationListV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiInscriptionnotauthenticatedID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInscriptionnotauthenticatedGetCommunicationListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InscriptionnotauthenticatedGetCommunicationListV1Response**](InscriptionnotauthenticatedGetCommunicationListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InscriptionnotauthenticatedGetCommunicationrecipientsV1

> InscriptionnotauthenticatedGetCommunicationrecipientsV1Response InscriptionnotauthenticatedGetCommunicationrecipientsV1(ctx, pkiInscriptionnotauthenticatedID).Execute()

Retrieve Inscriptionnotauthenticated's Communicationrecipient



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
	pkiInscriptionnotauthenticatedID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectInscriptionnotauthenticatedAPI.InscriptionnotauthenticatedGetCommunicationrecipientsV1(context.Background(), pkiInscriptionnotauthenticatedID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectInscriptionnotauthenticatedAPI.InscriptionnotauthenticatedGetCommunicationrecipientsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InscriptionnotauthenticatedGetCommunicationrecipientsV1`: InscriptionnotauthenticatedGetCommunicationrecipientsV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectInscriptionnotauthenticatedAPI.InscriptionnotauthenticatedGetCommunicationrecipientsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiInscriptionnotauthenticatedID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInscriptionnotauthenticatedGetCommunicationrecipientsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InscriptionnotauthenticatedGetCommunicationrecipientsV1Response**](InscriptionnotauthenticatedGetCommunicationrecipientsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InscriptionnotauthenticatedGetCommunicationsendersV1

> InscriptionnotauthenticatedGetCommunicationsendersV1Response InscriptionnotauthenticatedGetCommunicationsendersV1(ctx, pkiInscriptionnotauthenticatedID).Execute()

Retrieve Inscriptionnotauthenticated's Communicationsender



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
	pkiInscriptionnotauthenticatedID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectInscriptionnotauthenticatedAPI.InscriptionnotauthenticatedGetCommunicationsendersV1(context.Background(), pkiInscriptionnotauthenticatedID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectInscriptionnotauthenticatedAPI.InscriptionnotauthenticatedGetCommunicationsendersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InscriptionnotauthenticatedGetCommunicationsendersV1`: InscriptionnotauthenticatedGetCommunicationsendersV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectInscriptionnotauthenticatedAPI.InscriptionnotauthenticatedGetCommunicationsendersV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiInscriptionnotauthenticatedID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInscriptionnotauthenticatedGetCommunicationsendersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InscriptionnotauthenticatedGetCommunicationsendersV1Response**](InscriptionnotauthenticatedGetCommunicationsendersV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

