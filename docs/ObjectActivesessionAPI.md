# eZmaxAPI\ObjectActivesessionAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivesessionGenerateFederationTokenV1**](ObjectActivesessionAPI.md#ActivesessionGenerateFederationTokenV1) | **Post** /1/object/activesession/generateFederationToken | Generate a federation token
[**ActivesessionGetCurrentV1**](ObjectActivesessionAPI.md#ActivesessionGetCurrentV1) | **Get** /1/object/activesession/getCurrent | Get Current Activesession
[**ActivesessionGetCurrentV2**](ObjectActivesessionAPI.md#ActivesessionGetCurrentV2) | **Get** /2/object/activesession/getCurrent | Get Current Activesession
[**ActivesessionGetListV1**](ObjectActivesessionAPI.md#ActivesessionGetListV1) | **Get** /1/object/activesession/getList | Retrieve Activesession list



## ActivesessionGenerateFederationTokenV1

> ActivesessionGenerateFederationTokenV1Response ActivesessionGenerateFederationTokenV1(ctx).ActivesessionGenerateFederationTokenV1Request(activesessionGenerateFederationTokenV1Request).Execute()

Generate a federation token



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
	activesessionGenerateFederationTokenV1Request := *openapiclient.NewActivesessionGenerateFederationTokenV1Request("demo") // ActivesessionGenerateFederationTokenV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectActivesessionAPI.ActivesessionGenerateFederationTokenV1(context.Background()).ActivesessionGenerateFederationTokenV1Request(activesessionGenerateFederationTokenV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectActivesessionAPI.ActivesessionGenerateFederationTokenV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActivesessionGenerateFederationTokenV1`: ActivesessionGenerateFederationTokenV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectActivesessionAPI.ActivesessionGenerateFederationTokenV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiActivesessionGenerateFederationTokenV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **activesessionGenerateFederationTokenV1Request** | [**ActivesessionGenerateFederationTokenV1Request**](ActivesessionGenerateFederationTokenV1Request.md) |  | 

### Return type

[**ActivesessionGenerateFederationTokenV1Response**](ActivesessionGenerateFederationTokenV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActivesessionGetCurrentV1

> ActivesessionGetCurrentV1Response ActivesessionGetCurrentV1(ctx).Execute()

Get Current Activesession



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectActivesessionAPI.ActivesessionGetCurrentV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectActivesessionAPI.ActivesessionGetCurrentV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActivesessionGetCurrentV1`: ActivesessionGetCurrentV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectActivesessionAPI.ActivesessionGetCurrentV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiActivesessionGetCurrentV1Request struct via the builder pattern


### Return type

[**ActivesessionGetCurrentV1Response**](ActivesessionGetCurrentV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActivesessionGetCurrentV2

> ActivesessionGetCurrentV2Response ActivesessionGetCurrentV2(ctx).Execute()

Get Current Activesession



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectActivesessionAPI.ActivesessionGetCurrentV2(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectActivesessionAPI.ActivesessionGetCurrentV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActivesessionGetCurrentV2`: ActivesessionGetCurrentV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectActivesessionAPI.ActivesessionGetCurrentV2`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiActivesessionGetCurrentV2Request struct via the builder pattern


### Return type

[**ActivesessionGetCurrentV2Response**](ActivesessionGetCurrentV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActivesessionGetListV1

> ActivesessionGetListV1Response ActivesessionGetListV1(ctx).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()

Retrieve Activesession list

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
	resp, r, err := apiClient.ObjectActivesessionAPI.ActivesessionGetListV1(context.Background()).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectActivesessionAPI.ActivesessionGetListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActivesessionGetListV1`: ActivesessionGetListV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectActivesessionAPI.ActivesessionGetListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiActivesessionGetListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eOrderBy** | **string** | Specify how you want the results to be sorted | 
 **iRowMax** | **int32** |  | 
 **iRowOffset** | **int32** |  | [default to 0]
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 
 **sFilter** | **string** |  | 

### Return type

[**ActivesessionGetListV1Response**](ActivesessionGetListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

