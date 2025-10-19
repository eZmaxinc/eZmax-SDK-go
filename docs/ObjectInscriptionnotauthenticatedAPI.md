# eZmaxAPI\ObjectInscriptionnotauthenticatedAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InscriptionnotauthenticatedGetCommunicationCountV1**](ObjectInscriptionnotauthenticatedAPI.md#InscriptionnotauthenticatedGetCommunicationCountV1) | **Get** /1/object/inscriptionnotauthenticated/{pkiInscriptionnotauthenticatedID}/getCommunicationCount | Retrieve Communication count
[**InscriptionnotauthenticatedGetCommunicationListV1**](ObjectInscriptionnotauthenticatedAPI.md#InscriptionnotauthenticatedGetCommunicationListV1) | **Get** /1/object/inscriptionnotauthenticated/{pkiInscriptionnotauthenticatedID}/getCommunicationList | Retrieve Communication list
[**InscriptionnotauthenticatedGetCommunicationrecipientsV1**](ObjectInscriptionnotauthenticatedAPI.md#InscriptionnotauthenticatedGetCommunicationrecipientsV1) | **Get** /1/object/inscriptionnotauthenticated/{pkiInscriptionnotauthenticatedID}/getCommunicationrecipients | Retrieve Inscriptionnotauthenticated&#39;s Communicationrecipient
[**InscriptionnotauthenticatedGetCommunicationsendersV1**](ObjectInscriptionnotauthenticatedAPI.md#InscriptionnotauthenticatedGetCommunicationsendersV1) | **Get** /1/object/inscriptionnotauthenticated/{pkiInscriptionnotauthenticatedID}/getCommunicationsenders | Retrieve Inscriptionnotauthenticated&#39;s Communicationsender
[**InscriptionnotauthenticatedGetListV1**](ObjectInscriptionnotauthenticatedAPI.md#InscriptionnotauthenticatedGetListV1) | **Get** /1/object/inscriptionnotauthenticated/getList | Retrieve Inscriptionnotauthenticated list
[**InscriptionnotauthenticatedImportIntoEDMV1**](ObjectInscriptionnotauthenticatedAPI.md#InscriptionnotauthenticatedImportIntoEDMV1) | **Post** /1/object/inscriptionnotauthenticated/{pkiInscriptionnotauthenticatedID}/importIntoEDM | Import attachments into the Inscriptionnotauthenticated



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


## InscriptionnotauthenticatedGetListV1

> InscriptionnotauthenticatedGetListV1Response InscriptionnotauthenticatedGetListV1(ctx).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()

Retrieve Inscriptionnotauthenticated list



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
	eOrderBy := "eOrderBy_example" // string | Specify how you want the results to be sorted (optional)
	iRowMax := int32(56) // int32 |  (optional)
	iRowOffset := int32(56) // int32 |  (optional) (default to 0)
	acceptLanguage := openapiclient.Header-Accept-Language("*") // HeaderAcceptLanguage |  (optional)
	sFilter := "sFilter_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectInscriptionnotauthenticatedAPI.InscriptionnotauthenticatedGetListV1(context.Background()).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectInscriptionnotauthenticatedAPI.InscriptionnotauthenticatedGetListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InscriptionnotauthenticatedGetListV1`: InscriptionnotauthenticatedGetListV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectInscriptionnotauthenticatedAPI.InscriptionnotauthenticatedGetListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInscriptionnotauthenticatedGetListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eOrderBy** | **string** | Specify how you want the results to be sorted | 
 **iRowMax** | **int32** |  | 
 **iRowOffset** | **int32** |  | [default to 0]
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 
 **sFilter** | **string** |  | 

### Return type

[**InscriptionnotauthenticatedGetListV1Response**](InscriptionnotauthenticatedGetListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InscriptionnotauthenticatedImportIntoEDMV1

> InscriptionnotauthenticatedImportIntoEDMV1Response InscriptionnotauthenticatedImportIntoEDMV1(ctx, pkiInscriptionnotauthenticatedID).InscriptionnotauthenticatedImportIntoEDMV1Request(inscriptionnotauthenticatedImportIntoEDMV1Request).Execute()

Import attachments into the Inscriptionnotauthenticated



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
	inscriptionnotauthenticatedImportIntoEDMV1Request := *openapiclient.NewInscriptionnotauthenticatedImportIntoEDMV1Request([]openapiclient.CustomAttachmentImportIntoEDMRequest{*openapiclient.NewCustomAttachmentImportIntoEDMRequest("EAttachmentSource_example", "Document.pdf", "Inscription", openapiclient.Field-eAttachmentPrivacy("All"))}) // InscriptionnotauthenticatedImportIntoEDMV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectInscriptionnotauthenticatedAPI.InscriptionnotauthenticatedImportIntoEDMV1(context.Background(), pkiInscriptionnotauthenticatedID).InscriptionnotauthenticatedImportIntoEDMV1Request(inscriptionnotauthenticatedImportIntoEDMV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectInscriptionnotauthenticatedAPI.InscriptionnotauthenticatedImportIntoEDMV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InscriptionnotauthenticatedImportIntoEDMV1`: InscriptionnotauthenticatedImportIntoEDMV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectInscriptionnotauthenticatedAPI.InscriptionnotauthenticatedImportIntoEDMV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiInscriptionnotauthenticatedID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInscriptionnotauthenticatedImportIntoEDMV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inscriptionnotauthenticatedImportIntoEDMV1Request** | [**InscriptionnotauthenticatedImportIntoEDMV1Request**](InscriptionnotauthenticatedImportIntoEDMV1Request.md) |  | 

### Return type

[**InscriptionnotauthenticatedImportIntoEDMV1Response**](InscriptionnotauthenticatedImportIntoEDMV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

