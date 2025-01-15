# eZmaxAPI\ObjectWebhookAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WebhookCreateObjectV2**](ObjectWebhookAPI.md#WebhookCreateObjectV2) | **Post** /2/object/webhook | Create a new Webhook
[**WebhookDeleteObjectV1**](ObjectWebhookAPI.md#WebhookDeleteObjectV1) | **Delete** /1/object/webhook/{pkiWebhookID} | Delete an existing Webhook
[**WebhookEditObjectV1**](ObjectWebhookAPI.md#WebhookEditObjectV1) | **Put** /1/object/webhook/{pkiWebhookID} | Edit an existing Webhook
[**WebhookGetHistoryV1**](ObjectWebhookAPI.md#WebhookGetHistoryV1) | **Get** /1/object/webhook/{pkiWebhookID}/getHistory | Retrieve the logs for recent Webhook calls
[**WebhookGetListV1**](ObjectWebhookAPI.md#WebhookGetListV1) | **Get** /1/object/webhook/getList | Retrieve Webhook list
[**WebhookGetObjectV2**](ObjectWebhookAPI.md#WebhookGetObjectV2) | **Get** /2/object/webhook/{pkiWebhookID} | Retrieve an existing Webhook
[**WebhookRegenerateApikeyV1**](ObjectWebhookAPI.md#WebhookRegenerateApikeyV1) | **Post** /1/object/webhook/{pkiWebhookID}/regenerateApikey | Regenerate the Apikey
[**WebhookSendWebhookV1**](ObjectWebhookAPI.md#WebhookSendWebhookV1) | **Post** /1/object/webhook/sendWebhook | Emit a Webhook event
[**WebhookTestV1**](ObjectWebhookAPI.md#WebhookTestV1) | **Post** /1/object/webhook/{pkiWebhookID}/test | Test the Webhook by calling the Url



## WebhookCreateObjectV2

> WebhookCreateObjectV2Response WebhookCreateObjectV2(ctx).WebhookCreateObjectV2Request(webhookCreateObjectV2Request).Execute()

Create a new Webhook



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
	webhookCreateObjectV2Request := *openapiclient.NewWebhookCreateObjectV2Request([]openapiclient.WebhookRequestCompound{*openapiclient.NewWebhookRequestCompound("Import into our system", openapiclient.Field-eWebhookModule("Ezsign"), "https://www.example.com", "email@example.com", true, false)}) // WebhookCreateObjectV2Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectWebhookAPI.WebhookCreateObjectV2(context.Background()).WebhookCreateObjectV2Request(webhookCreateObjectV2Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectWebhookAPI.WebhookCreateObjectV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhookCreateObjectV2`: WebhookCreateObjectV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectWebhookAPI.WebhookCreateObjectV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookCreateObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookCreateObjectV2Request** | [**WebhookCreateObjectV2Request**](WebhookCreateObjectV2Request.md) |  | 

### Return type

[**WebhookCreateObjectV2Response**](WebhookCreateObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookDeleteObjectV1

> CommonResponse WebhookDeleteObjectV1(ctx, pkiWebhookID).Execute()

Delete an existing Webhook



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
	pkiWebhookID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectWebhookAPI.WebhookDeleteObjectV1(context.Background(), pkiWebhookID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectWebhookAPI.WebhookDeleteObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhookDeleteObjectV1`: CommonResponse
	fmt.Fprintf(os.Stdout, "Response from `ObjectWebhookAPI.WebhookDeleteObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiWebhookID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWebhookDeleteObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CommonResponse**](CommonResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookEditObjectV1

> CommonResponse WebhookEditObjectV1(ctx, pkiWebhookID).WebhookEditObjectV1Request(webhookEditObjectV1Request).Execute()

Edit an existing Webhook



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
	pkiWebhookID := int32(56) // int32 | 
	webhookEditObjectV1Request := *openapiclient.NewWebhookEditObjectV1Request(*openapiclient.NewWebhookRequestCompound("Import into our system", openapiclient.Field-eWebhookModule("Ezsign"), "https://www.example.com", "email@example.com", true, false)) // WebhookEditObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectWebhookAPI.WebhookEditObjectV1(context.Background(), pkiWebhookID).WebhookEditObjectV1Request(webhookEditObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectWebhookAPI.WebhookEditObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhookEditObjectV1`: CommonResponse
	fmt.Fprintf(os.Stdout, "Response from `ObjectWebhookAPI.WebhookEditObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiWebhookID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWebhookEditObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **webhookEditObjectV1Request** | [**WebhookEditObjectV1Request**](WebhookEditObjectV1Request.md) |  | 

### Return type

[**CommonResponse**](CommonResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookGetHistoryV1

> WebhookGetHistoryV1Response WebhookGetHistoryV1(ctx, pkiWebhookID).EWebhookHistoryinterval(eWebhookHistoryinterval).Execute()

Retrieve the logs for recent Webhook calls



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
	pkiWebhookID := int32(56) // int32 | 
	eWebhookHistoryinterval := "eWebhookHistoryinterval_example" // string | The number of days to return

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectWebhookAPI.WebhookGetHistoryV1(context.Background(), pkiWebhookID).EWebhookHistoryinterval(eWebhookHistoryinterval).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectWebhookAPI.WebhookGetHistoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhookGetHistoryV1`: WebhookGetHistoryV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectWebhookAPI.WebhookGetHistoryV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiWebhookID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWebhookGetHistoryV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eWebhookHistoryinterval** | **string** | The number of days to return | 

### Return type

[**WebhookGetHistoryV1Response**](WebhookGetHistoryV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookGetListV1

> WebhookGetListV1Response WebhookGetListV1(ctx).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()

Retrieve Webhook list



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
	resp, r, err := apiClient.ObjectWebhookAPI.WebhookGetListV1(context.Background()).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectWebhookAPI.WebhookGetListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhookGetListV1`: WebhookGetListV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectWebhookAPI.WebhookGetListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookGetListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eOrderBy** | **string** | Specify how you want the results to be sorted | 
 **iRowMax** | **int32** |  | 
 **iRowOffset** | **int32** |  | [default to 0]
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 
 **sFilter** | **string** |  | 

### Return type

[**WebhookGetListV1Response**](WebhookGetListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookGetObjectV2

> WebhookGetObjectV2Response WebhookGetObjectV2(ctx, pkiWebhookID).Execute()

Retrieve an existing Webhook



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
	pkiWebhookID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectWebhookAPI.WebhookGetObjectV2(context.Background(), pkiWebhookID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectWebhookAPI.WebhookGetObjectV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhookGetObjectV2`: WebhookGetObjectV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectWebhookAPI.WebhookGetObjectV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiWebhookID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWebhookGetObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WebhookGetObjectV2Response**](WebhookGetObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookRegenerateApikeyV1

> WebhookRegenerateApikeyV1Response WebhookRegenerateApikeyV1(ctx, pkiWebhookID).WebhookRegenerateApikeyV1Request(webhookRegenerateApikeyV1Request).Execute()

Regenerate the Apikey



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
	pkiWebhookID := int32(56) // int32 | 
	webhookRegenerateApikeyV1Request := *openapiclient.NewWebhookRegenerateApikeyV1Request() // WebhookRegenerateApikeyV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectWebhookAPI.WebhookRegenerateApikeyV1(context.Background(), pkiWebhookID).WebhookRegenerateApikeyV1Request(webhookRegenerateApikeyV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectWebhookAPI.WebhookRegenerateApikeyV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhookRegenerateApikeyV1`: WebhookRegenerateApikeyV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectWebhookAPI.WebhookRegenerateApikeyV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiWebhookID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWebhookRegenerateApikeyV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **webhookRegenerateApikeyV1Request** | [**WebhookRegenerateApikeyV1Request**](WebhookRegenerateApikeyV1Request.md) |  | 

### Return type

[**WebhookRegenerateApikeyV1Response**](WebhookRegenerateApikeyV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookSendWebhookV1

> CommonResponse WebhookSendWebhookV1(ctx).WebhookSendWebhookV1Request(webhookSendWebhookV1Request).Execute()

Emit a Webhook event

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
	webhookSendWebhookV1Request := *openapiclient.NewWebhookSendWebhookV1Request(openapiclient.Field-eWebhookModule("Ezsign")) // WebhookSendWebhookV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectWebhookAPI.WebhookSendWebhookV1(context.Background()).WebhookSendWebhookV1Request(webhookSendWebhookV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectWebhookAPI.WebhookSendWebhookV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhookSendWebhookV1`: CommonResponse
	fmt.Fprintf(os.Stdout, "Response from `ObjectWebhookAPI.WebhookSendWebhookV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookSendWebhookV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookSendWebhookV1Request** | [**WebhookSendWebhookV1Request**](WebhookSendWebhookV1Request.md) |  | 

### Return type

[**CommonResponse**](CommonResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookTestV1

> CommonResponse WebhookTestV1(ctx, pkiWebhookID).Body(body).Execute()

Test the Webhook by calling the Url



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
	pkiWebhookID := int32(56) // int32 | 
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectWebhookAPI.WebhookTestV1(context.Background(), pkiWebhookID).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectWebhookAPI.WebhookTestV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhookTestV1`: CommonResponse
	fmt.Fprintf(os.Stdout, "Response from `ObjectWebhookAPI.WebhookTestV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiWebhookID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWebhookTestV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

[**CommonResponse**](CommonResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

