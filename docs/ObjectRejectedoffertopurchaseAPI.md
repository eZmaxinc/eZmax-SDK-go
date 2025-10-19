# eZmaxAPI\ObjectRejectedoffertopurchaseAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RejectedoffertopurchaseGetCommunicationCountV1**](ObjectRejectedoffertopurchaseAPI.md#RejectedoffertopurchaseGetCommunicationCountV1) | **Get** /1/object/rejectedoffertopurchase/{pkiRejectedoffertopurchaseID}/getCommunicationCount | Retrieve Communication count
[**RejectedoffertopurchaseGetCommunicationListV1**](ObjectRejectedoffertopurchaseAPI.md#RejectedoffertopurchaseGetCommunicationListV1) | **Get** /1/object/rejectedoffertopurchase/{pkiRejectedoffertopurchaseID}/getCommunicationList | Retrieve Communication list
[**RejectedoffertopurchaseGetCommunicationrecipientsV1**](ObjectRejectedoffertopurchaseAPI.md#RejectedoffertopurchaseGetCommunicationrecipientsV1) | **Get** /1/object/rejectedoffertopurchase/{pkiRejectedoffertopurchaseID}/getCommunicationrecipients | Retrieve Rejectedoffertopurchase&#39;s Communicationrecipient
[**RejectedoffertopurchaseGetCommunicationsendersV1**](ObjectRejectedoffertopurchaseAPI.md#RejectedoffertopurchaseGetCommunicationsendersV1) | **Get** /1/object/rejectedoffertopurchase/{pkiRejectedoffertopurchaseID}/getCommunicationsenders | Retrieve Rejectedoffertopurchase&#39;s Communicationsender
[**RejectedoffertopurchaseGetListV1**](ObjectRejectedoffertopurchaseAPI.md#RejectedoffertopurchaseGetListV1) | **Get** /1/object/rejectedoffertopurchase/getList | Retrieve Rejectedoffertopurchase list
[**RejectedoffertopurchaseImportIntoEDMV1**](ObjectRejectedoffertopurchaseAPI.md#RejectedoffertopurchaseImportIntoEDMV1) | **Post** /1/object/rejectedoffertopurchase/{pkiRejectedoffertopurchaseID}/importIntoEDM | Import attachments into the Rejectedoffertopurchase



## RejectedoffertopurchaseGetCommunicationCountV1

> RejectedoffertopurchaseGetCommunicationCountV1Response RejectedoffertopurchaseGetCommunicationCountV1(ctx, pkiRejectedoffertopurchaseID).Execute()

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
	pkiRejectedoffertopurchaseID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectRejectedoffertopurchaseAPI.RejectedoffertopurchaseGetCommunicationCountV1(context.Background(), pkiRejectedoffertopurchaseID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectRejectedoffertopurchaseAPI.RejectedoffertopurchaseGetCommunicationCountV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RejectedoffertopurchaseGetCommunicationCountV1`: RejectedoffertopurchaseGetCommunicationCountV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectRejectedoffertopurchaseAPI.RejectedoffertopurchaseGetCommunicationCountV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiRejectedoffertopurchaseID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRejectedoffertopurchaseGetCommunicationCountV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RejectedoffertopurchaseGetCommunicationCountV1Response**](RejectedoffertopurchaseGetCommunicationCountV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RejectedoffertopurchaseGetCommunicationListV1

> RejectedoffertopurchaseGetCommunicationListV1Response RejectedoffertopurchaseGetCommunicationListV1(ctx, pkiRejectedoffertopurchaseID).Execute()

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
	pkiRejectedoffertopurchaseID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectRejectedoffertopurchaseAPI.RejectedoffertopurchaseGetCommunicationListV1(context.Background(), pkiRejectedoffertopurchaseID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectRejectedoffertopurchaseAPI.RejectedoffertopurchaseGetCommunicationListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RejectedoffertopurchaseGetCommunicationListV1`: RejectedoffertopurchaseGetCommunicationListV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectRejectedoffertopurchaseAPI.RejectedoffertopurchaseGetCommunicationListV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiRejectedoffertopurchaseID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRejectedoffertopurchaseGetCommunicationListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RejectedoffertopurchaseGetCommunicationListV1Response**](RejectedoffertopurchaseGetCommunicationListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RejectedoffertopurchaseGetCommunicationrecipientsV1

> RejectedoffertopurchaseGetCommunicationrecipientsV1Response RejectedoffertopurchaseGetCommunicationrecipientsV1(ctx, pkiRejectedoffertopurchaseID).Execute()

Retrieve Rejectedoffertopurchase's Communicationrecipient



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
	pkiRejectedoffertopurchaseID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectRejectedoffertopurchaseAPI.RejectedoffertopurchaseGetCommunicationrecipientsV1(context.Background(), pkiRejectedoffertopurchaseID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectRejectedoffertopurchaseAPI.RejectedoffertopurchaseGetCommunicationrecipientsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RejectedoffertopurchaseGetCommunicationrecipientsV1`: RejectedoffertopurchaseGetCommunicationrecipientsV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectRejectedoffertopurchaseAPI.RejectedoffertopurchaseGetCommunicationrecipientsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiRejectedoffertopurchaseID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRejectedoffertopurchaseGetCommunicationrecipientsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RejectedoffertopurchaseGetCommunicationrecipientsV1Response**](RejectedoffertopurchaseGetCommunicationrecipientsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RejectedoffertopurchaseGetCommunicationsendersV1

> RejectedoffertopurchaseGetCommunicationsendersV1Response RejectedoffertopurchaseGetCommunicationsendersV1(ctx, pkiRejectedoffertopurchaseID).Execute()

Retrieve Rejectedoffertopurchase's Communicationsender



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
	pkiRejectedoffertopurchaseID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectRejectedoffertopurchaseAPI.RejectedoffertopurchaseGetCommunicationsendersV1(context.Background(), pkiRejectedoffertopurchaseID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectRejectedoffertopurchaseAPI.RejectedoffertopurchaseGetCommunicationsendersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RejectedoffertopurchaseGetCommunicationsendersV1`: RejectedoffertopurchaseGetCommunicationsendersV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectRejectedoffertopurchaseAPI.RejectedoffertopurchaseGetCommunicationsendersV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiRejectedoffertopurchaseID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRejectedoffertopurchaseGetCommunicationsendersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RejectedoffertopurchaseGetCommunicationsendersV1Response**](RejectedoffertopurchaseGetCommunicationsendersV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RejectedoffertopurchaseGetListV1

> RejectedoffertopurchaseGetListV1Response RejectedoffertopurchaseGetListV1(ctx).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()

Retrieve Rejectedoffertopurchase list



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
	resp, r, err := apiClient.ObjectRejectedoffertopurchaseAPI.RejectedoffertopurchaseGetListV1(context.Background()).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectRejectedoffertopurchaseAPI.RejectedoffertopurchaseGetListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RejectedoffertopurchaseGetListV1`: RejectedoffertopurchaseGetListV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectRejectedoffertopurchaseAPI.RejectedoffertopurchaseGetListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRejectedoffertopurchaseGetListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eOrderBy** | **string** | Specify how you want the results to be sorted | 
 **iRowMax** | **int32** |  | 
 **iRowOffset** | **int32** |  | [default to 0]
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 
 **sFilter** | **string** |  | 

### Return type

[**RejectedoffertopurchaseGetListV1Response**](RejectedoffertopurchaseGetListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RejectedoffertopurchaseImportIntoEDMV1

> RejectedoffertopurchaseImportIntoEDMV1Response RejectedoffertopurchaseImportIntoEDMV1(ctx, pkiRejectedoffertopurchaseID).RejectedoffertopurchaseImportIntoEDMV1Request(rejectedoffertopurchaseImportIntoEDMV1Request).Execute()

Import attachments into the Rejectedoffertopurchase



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
	pkiRejectedoffertopurchaseID := int32(56) // int32 | 
	rejectedoffertopurchaseImportIntoEDMV1Request := *openapiclient.NewRejectedoffertopurchaseImportIntoEDMV1Request([]openapiclient.CustomAttachmentImportIntoEDMRequest{*openapiclient.NewCustomAttachmentImportIntoEDMRequest("EAttachmentSource_example", "Document.pdf", "Inscription", openapiclient.Field-eAttachmentPrivacy("All"))}) // RejectedoffertopurchaseImportIntoEDMV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectRejectedoffertopurchaseAPI.RejectedoffertopurchaseImportIntoEDMV1(context.Background(), pkiRejectedoffertopurchaseID).RejectedoffertopurchaseImportIntoEDMV1Request(rejectedoffertopurchaseImportIntoEDMV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectRejectedoffertopurchaseAPI.RejectedoffertopurchaseImportIntoEDMV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RejectedoffertopurchaseImportIntoEDMV1`: RejectedoffertopurchaseImportIntoEDMV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectRejectedoffertopurchaseAPI.RejectedoffertopurchaseImportIntoEDMV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiRejectedoffertopurchaseID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRejectedoffertopurchaseImportIntoEDMV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rejectedoffertopurchaseImportIntoEDMV1Request** | [**RejectedoffertopurchaseImportIntoEDMV1Request**](RejectedoffertopurchaseImportIntoEDMV1Request.md) |  | 

### Return type

[**RejectedoffertopurchaseImportIntoEDMV1Response**](RejectedoffertopurchaseImportIntoEDMV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

