# eZmaxAPI\ObjectCorsAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CorsCreateObjectV1**](ObjectCorsAPI.md#CorsCreateObjectV1) | **Post** /1/object/cors | Create a new Cors
[**CorsDeleteObjectV1**](ObjectCorsAPI.md#CorsDeleteObjectV1) | **Delete** /1/object/cors/{pkiCorsID} | Delete an existing Cors
[**CorsEditObjectV1**](ObjectCorsAPI.md#CorsEditObjectV1) | **Put** /1/object/cors/{pkiCorsID} | Edit an existing Cors
[**CorsGetObjectV2**](ObjectCorsAPI.md#CorsGetObjectV2) | **Get** /2/object/cors/{pkiCorsID} | Retrieve an existing Cors



## CorsCreateObjectV1

> CorsCreateObjectV1Response CorsCreateObjectV1(ctx).CorsCreateObjectV1Request(corsCreateObjectV1Request).Execute()

Create a new Cors



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
	corsCreateObjectV1Request := *openapiclient.NewCorsCreateObjectV1Request([]openapiclient.CorsRequestCompound{*openapiclient.NewCorsRequestCompound(int32(99), "Https://www.example.com")}) // CorsCreateObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectCorsAPI.CorsCreateObjectV1(context.Background()).CorsCreateObjectV1Request(corsCreateObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectCorsAPI.CorsCreateObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CorsCreateObjectV1`: CorsCreateObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectCorsAPI.CorsCreateObjectV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCorsCreateObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **corsCreateObjectV1Request** | [**CorsCreateObjectV1Request**](CorsCreateObjectV1Request.md) |  | 

### Return type

[**CorsCreateObjectV1Response**](CorsCreateObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CorsDeleteObjectV1

> CorsDeleteObjectV1Response CorsDeleteObjectV1(ctx, pkiCorsID).Execute()

Delete an existing Cors



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
	pkiCorsID := int32(56) // int32 | The unique ID of the Cors

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectCorsAPI.CorsDeleteObjectV1(context.Background(), pkiCorsID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectCorsAPI.CorsDeleteObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CorsDeleteObjectV1`: CorsDeleteObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectCorsAPI.CorsDeleteObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiCorsID** | **int32** | The unique ID of the Cors | 

### Other Parameters

Other parameters are passed through a pointer to a apiCorsDeleteObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CorsDeleteObjectV1Response**](CorsDeleteObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CorsEditObjectV1

> CorsEditObjectV1Response CorsEditObjectV1(ctx, pkiCorsID).CorsEditObjectV1Request(corsEditObjectV1Request).Execute()

Edit an existing Cors



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
	pkiCorsID := int32(56) // int32 | The unique ID of the Cors
	corsEditObjectV1Request := *openapiclient.NewCorsEditObjectV1Request(*openapiclient.NewCorsRequestCompound(int32(99), "Https://www.example.com")) // CorsEditObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectCorsAPI.CorsEditObjectV1(context.Background(), pkiCorsID).CorsEditObjectV1Request(corsEditObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectCorsAPI.CorsEditObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CorsEditObjectV1`: CorsEditObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectCorsAPI.CorsEditObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiCorsID** | **int32** | The unique ID of the Cors | 

### Other Parameters

Other parameters are passed through a pointer to a apiCorsEditObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **corsEditObjectV1Request** | [**CorsEditObjectV1Request**](CorsEditObjectV1Request.md) |  | 

### Return type

[**CorsEditObjectV1Response**](CorsEditObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CorsGetObjectV2

> CorsGetObjectV2Response CorsGetObjectV2(ctx, pkiCorsID).Execute()

Retrieve an existing Cors



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
	pkiCorsID := int32(56) // int32 | The unique ID of the Cors

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectCorsAPI.CorsGetObjectV2(context.Background(), pkiCorsID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectCorsAPI.CorsGetObjectV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CorsGetObjectV2`: CorsGetObjectV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectCorsAPI.CorsGetObjectV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiCorsID** | **int32** | The unique ID of the Cors | 

### Other Parameters

Other parameters are passed through a pointer to a apiCorsGetObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CorsGetObjectV2Response**](CorsGetObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

