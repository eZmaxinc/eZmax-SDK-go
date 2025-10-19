# eZmaxAPI\ObjectOtherincomeAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OtherincomeGetCommunicationCountV1**](ObjectOtherincomeAPI.md#OtherincomeGetCommunicationCountV1) | **Get** /1/object/otherincome/{pkiOtherincomeID}/getCommunicationCount | Retrieve Communication count
[**OtherincomeGetCommunicationListV1**](ObjectOtherincomeAPI.md#OtherincomeGetCommunicationListV1) | **Get** /1/object/otherincome/{pkiOtherincomeID}/getCommunicationList | Retrieve Communication list
[**OtherincomeGetCommunicationrecipientsV1**](ObjectOtherincomeAPI.md#OtherincomeGetCommunicationrecipientsV1) | **Get** /1/object/otherincome/{pkiOtherincomeID}/getCommunicationrecipients | Retrieve Otherincome&#39;s Communicationrecipient
[**OtherincomeGetCommunicationsendersV1**](ObjectOtherincomeAPI.md#OtherincomeGetCommunicationsendersV1) | **Get** /1/object/otherincome/{pkiOtherincomeID}/getCommunicationsenders | Retrieve Otherincome&#39;s Communicationsender
[**OtherincomeGetListV1**](ObjectOtherincomeAPI.md#OtherincomeGetListV1) | **Get** /1/object/otherincome/getList | Retrieve Otherincome list
[**OtherincomeImportIntoEDMV1**](ObjectOtherincomeAPI.md#OtherincomeImportIntoEDMV1) | **Post** /1/object/otherincome/{pkiOtherincomeID}/importIntoEDM | Import attachments into the Otherincome



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


## OtherincomeGetListV1

> OtherincomeGetListV1Response OtherincomeGetListV1(ctx).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()

Retrieve Otherincome list



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
	resp, r, err := apiClient.ObjectOtherincomeAPI.OtherincomeGetListV1(context.Background()).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectOtherincomeAPI.OtherincomeGetListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OtherincomeGetListV1`: OtherincomeGetListV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectOtherincomeAPI.OtherincomeGetListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOtherincomeGetListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eOrderBy** | **string** | Specify how you want the results to be sorted | 
 **iRowMax** | **int32** |  | 
 **iRowOffset** | **int32** |  | [default to 0]
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 
 **sFilter** | **string** |  | 

### Return type

[**OtherincomeGetListV1Response**](OtherincomeGetListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OtherincomeImportIntoEDMV1

> OtherincomeImportIntoEDMV1Response OtherincomeImportIntoEDMV1(ctx, pkiOtherincomeID).OtherincomeImportIntoEDMV1Request(otherincomeImportIntoEDMV1Request).Execute()

Import attachments into the Otherincome



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
	otherincomeImportIntoEDMV1Request := *openapiclient.NewOtherincomeImportIntoEDMV1Request([]openapiclient.CustomAttachmentImportIntoEDMRequest{*openapiclient.NewCustomAttachmentImportIntoEDMRequest("EAttachmentSource_example", "Document.pdf", "Inscription", openapiclient.Field-eAttachmentPrivacy("All"))}) // OtherincomeImportIntoEDMV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectOtherincomeAPI.OtherincomeImportIntoEDMV1(context.Background(), pkiOtherincomeID).OtherincomeImportIntoEDMV1Request(otherincomeImportIntoEDMV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectOtherincomeAPI.OtherincomeImportIntoEDMV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OtherincomeImportIntoEDMV1`: OtherincomeImportIntoEDMV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectOtherincomeAPI.OtherincomeImportIntoEDMV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiOtherincomeID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOtherincomeImportIntoEDMV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **otherincomeImportIntoEDMV1Request** | [**OtherincomeImportIntoEDMV1Request**](OtherincomeImportIntoEDMV1Request.md) |  | 

### Return type

[**OtherincomeImportIntoEDMV1Response**](OtherincomeImportIntoEDMV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

