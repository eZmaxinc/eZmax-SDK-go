# eZmaxAPI\ObjectAuthenticationexternalAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthenticationexternalCreateObjectV1**](ObjectAuthenticationexternalAPI.md#AuthenticationexternalCreateObjectV1) | **Post** /1/object/authenticationexternal | Create a new Authenticationexternal
[**AuthenticationexternalDeleteObjectV1**](ObjectAuthenticationexternalAPI.md#AuthenticationexternalDeleteObjectV1) | **Delete** /1/object/authenticationexternal/{pkiAuthenticationexternalID} | Delete an existing Authenticationexternal
[**AuthenticationexternalEditObjectV1**](ObjectAuthenticationexternalAPI.md#AuthenticationexternalEditObjectV1) | **Put** /1/object/authenticationexternal/{pkiAuthenticationexternalID} | Edit an existing Authenticationexternal
[**AuthenticationexternalGetAutocompleteV2**](ObjectAuthenticationexternalAPI.md#AuthenticationexternalGetAutocompleteV2) | **Get** /2/object/authenticationexternal/getAutocomplete/{sSelector} | Retrieve Authenticationexternals and IDs
[**AuthenticationexternalGetListV1**](ObjectAuthenticationexternalAPI.md#AuthenticationexternalGetListV1) | **Get** /1/object/authenticationexternal/getList | Retrieve Authenticationexternal list
[**AuthenticationexternalGetObjectV2**](ObjectAuthenticationexternalAPI.md#AuthenticationexternalGetObjectV2) | **Get** /2/object/authenticationexternal/{pkiAuthenticationexternalID} | Retrieve an existing Authenticationexternal
[**AuthenticationexternalResetAuthorizationV1**](ObjectAuthenticationexternalAPI.md#AuthenticationexternalResetAuthorizationV1) | **Post** /1/object/authenticationexternal/{pkiAuthenticationexternalID}/resetAuthorization | Reset the Authenticationexternal authorization



## AuthenticationexternalCreateObjectV1

> AuthenticationexternalCreateObjectV1Response AuthenticationexternalCreateObjectV1(ctx).AuthenticationexternalCreateObjectV1Request(authenticationexternalCreateObjectV1Request).Execute()

Create a new Authenticationexternal



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
	authenticationexternalCreateObjectV1Request := *openapiclient.NewAuthenticationexternalCreateObjectV1Request([]openapiclient.AuthenticationexternalRequestCompound{*openapiclient.NewAuthenticationexternalRequestCompound("Authentification", openapiclient.Field-eAuthenticationexternalType("Salesforce"))}) // AuthenticationexternalCreateObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectAuthenticationexternalAPI.AuthenticationexternalCreateObjectV1(context.Background()).AuthenticationexternalCreateObjectV1Request(authenticationexternalCreateObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectAuthenticationexternalAPI.AuthenticationexternalCreateObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthenticationexternalCreateObjectV1`: AuthenticationexternalCreateObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectAuthenticationexternalAPI.AuthenticationexternalCreateObjectV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticationexternalCreateObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authenticationexternalCreateObjectV1Request** | [**AuthenticationexternalCreateObjectV1Request**](AuthenticationexternalCreateObjectV1Request.md) |  | 

### Return type

[**AuthenticationexternalCreateObjectV1Response**](AuthenticationexternalCreateObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthenticationexternalDeleteObjectV1

> CommonResponse AuthenticationexternalDeleteObjectV1(ctx, pkiAuthenticationexternalID).Execute()

Delete an existing Authenticationexternal



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
	pkiAuthenticationexternalID := int32(56) // int32 | The unique ID of the Authenticationexternal

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectAuthenticationexternalAPI.AuthenticationexternalDeleteObjectV1(context.Background(), pkiAuthenticationexternalID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectAuthenticationexternalAPI.AuthenticationexternalDeleteObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthenticationexternalDeleteObjectV1`: CommonResponse
	fmt.Fprintf(os.Stdout, "Response from `ObjectAuthenticationexternalAPI.AuthenticationexternalDeleteObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiAuthenticationexternalID** | **int32** | The unique ID of the Authenticationexternal | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticationexternalDeleteObjectV1Request struct via the builder pattern


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


## AuthenticationexternalEditObjectV1

> CommonResponse AuthenticationexternalEditObjectV1(ctx, pkiAuthenticationexternalID).AuthenticationexternalEditObjectV1Request(authenticationexternalEditObjectV1Request).Execute()

Edit an existing Authenticationexternal



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
	pkiAuthenticationexternalID := int32(56) // int32 | The unique ID of the Authenticationexternal
	authenticationexternalEditObjectV1Request := *openapiclient.NewAuthenticationexternalEditObjectV1Request(*openapiclient.NewAuthenticationexternalRequestCompound("Authentification", openapiclient.Field-eAuthenticationexternalType("Salesforce"))) // AuthenticationexternalEditObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectAuthenticationexternalAPI.AuthenticationexternalEditObjectV1(context.Background(), pkiAuthenticationexternalID).AuthenticationexternalEditObjectV1Request(authenticationexternalEditObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectAuthenticationexternalAPI.AuthenticationexternalEditObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthenticationexternalEditObjectV1`: CommonResponse
	fmt.Fprintf(os.Stdout, "Response from `ObjectAuthenticationexternalAPI.AuthenticationexternalEditObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiAuthenticationexternalID** | **int32** | The unique ID of the Authenticationexternal | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticationexternalEditObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authenticationexternalEditObjectV1Request** | [**AuthenticationexternalEditObjectV1Request**](AuthenticationexternalEditObjectV1Request.md) |  | 

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


## AuthenticationexternalGetAutocompleteV2

> AuthenticationexternalGetAutocompleteV2Response AuthenticationexternalGetAutocompleteV2(ctx, sSelector).EFilterActive(eFilterActive).SQuery(sQuery).AcceptLanguage(acceptLanguage).Execute()

Retrieve Authenticationexternals and IDs



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
	sSelector := "sSelector_example" // string | The type of Authenticationexternals to return
	eFilterActive := "eFilterActive_example" // string | Specify which results we want to display. (optional) (default to "Active")
	sQuery := "sQuery_example" // string | Allow to filter the returned results (optional)
	acceptLanguage := openapiclient.Header-Accept-Language("*") // HeaderAcceptLanguage |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectAuthenticationexternalAPI.AuthenticationexternalGetAutocompleteV2(context.Background(), sSelector).EFilterActive(eFilterActive).SQuery(sQuery).AcceptLanguage(acceptLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectAuthenticationexternalAPI.AuthenticationexternalGetAutocompleteV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthenticationexternalGetAutocompleteV2`: AuthenticationexternalGetAutocompleteV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectAuthenticationexternalAPI.AuthenticationexternalGetAutocompleteV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sSelector** | **string** | The type of Authenticationexternals to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticationexternalGetAutocompleteV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eFilterActive** | **string** | Specify which results we want to display. | [default to &quot;Active&quot;]
 **sQuery** | **string** | Allow to filter the returned results | 
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 

### Return type

[**AuthenticationexternalGetAutocompleteV2Response**](AuthenticationexternalGetAutocompleteV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthenticationexternalGetListV1

> AuthenticationexternalGetListV1Response AuthenticationexternalGetListV1(ctx).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()

Retrieve Authenticationexternal list



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
	resp, r, err := apiClient.ObjectAuthenticationexternalAPI.AuthenticationexternalGetListV1(context.Background()).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectAuthenticationexternalAPI.AuthenticationexternalGetListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthenticationexternalGetListV1`: AuthenticationexternalGetListV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectAuthenticationexternalAPI.AuthenticationexternalGetListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticationexternalGetListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eOrderBy** | **string** | Specify how you want the results to be sorted | 
 **iRowMax** | **int32** |  | 
 **iRowOffset** | **int32** |  | [default to 0]
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 
 **sFilter** | **string** |  | 

### Return type

[**AuthenticationexternalGetListV1Response**](AuthenticationexternalGetListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthenticationexternalGetObjectV2

> AuthenticationexternalGetObjectV2Response AuthenticationexternalGetObjectV2(ctx, pkiAuthenticationexternalID).Execute()

Retrieve an existing Authenticationexternal



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
	pkiAuthenticationexternalID := int32(56) // int32 | The unique ID of the Authenticationexternal

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectAuthenticationexternalAPI.AuthenticationexternalGetObjectV2(context.Background(), pkiAuthenticationexternalID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectAuthenticationexternalAPI.AuthenticationexternalGetObjectV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthenticationexternalGetObjectV2`: AuthenticationexternalGetObjectV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectAuthenticationexternalAPI.AuthenticationexternalGetObjectV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiAuthenticationexternalID** | **int32** | The unique ID of the Authenticationexternal | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticationexternalGetObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuthenticationexternalGetObjectV2Response**](AuthenticationexternalGetObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthenticationexternalResetAuthorizationV1

> CommonResponse AuthenticationexternalResetAuthorizationV1(ctx, pkiAuthenticationexternalID).Body(body).Execute()

Reset the Authenticationexternal authorization



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
	pkiAuthenticationexternalID := int32(56) // int32 | 
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectAuthenticationexternalAPI.AuthenticationexternalResetAuthorizationV1(context.Background(), pkiAuthenticationexternalID).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectAuthenticationexternalAPI.AuthenticationexternalResetAuthorizationV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthenticationexternalResetAuthorizationV1`: CommonResponse
	fmt.Fprintf(os.Stdout, "Response from `ObjectAuthenticationexternalAPI.AuthenticationexternalResetAuthorizationV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiAuthenticationexternalID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticationexternalResetAuthorizationV1Request struct via the builder pattern


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

