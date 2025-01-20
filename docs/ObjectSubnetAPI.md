# eZmaxAPI\ObjectSubnetAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubnetCreateObjectV1**](ObjectSubnetAPI.md#SubnetCreateObjectV1) | **Post** /1/object/subnet | Create a new Subnet
[**SubnetDeleteObjectV1**](ObjectSubnetAPI.md#SubnetDeleteObjectV1) | **Delete** /1/object/subnet/{pkiSubnetID} | Delete an existing Subnet
[**SubnetEditObjectV1**](ObjectSubnetAPI.md#SubnetEditObjectV1) | **Put** /1/object/subnet/{pkiSubnetID} | Edit an existing Subnet
[**SubnetGetObjectV2**](ObjectSubnetAPI.md#SubnetGetObjectV2) | **Get** /2/object/subnet/{pkiSubnetID} | Retrieve an existing Subnet



## SubnetCreateObjectV1

> SubnetCreateObjectV1Response SubnetCreateObjectV1(ctx).SubnetCreateObjectV1Request(subnetCreateObjectV1Request).Execute()

Create a new Subnet



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
	subnetCreateObjectV1Request := *openapiclient.NewSubnetCreateObjectV1Request([]openapiclient.SubnetRequestCompound{*openapiclient.NewSubnetRequestCompound(*openapiclient.NewMultilingualSubnetDescription(), int64(134744064), int64(4294967040))}) // SubnetCreateObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectSubnetAPI.SubnetCreateObjectV1(context.Background()).SubnetCreateObjectV1Request(subnetCreateObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectSubnetAPI.SubnetCreateObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubnetCreateObjectV1`: SubnetCreateObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectSubnetAPI.SubnetCreateObjectV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubnetCreateObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subnetCreateObjectV1Request** | [**SubnetCreateObjectV1Request**](SubnetCreateObjectV1Request.md) |  | 

### Return type

[**SubnetCreateObjectV1Response**](SubnetCreateObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubnetDeleteObjectV1

> CommonResponse SubnetDeleteObjectV1(ctx, pkiSubnetID).Execute()

Delete an existing Subnet



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
	pkiSubnetID := int32(56) // int32 | The unique ID of the Subnet

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectSubnetAPI.SubnetDeleteObjectV1(context.Background(), pkiSubnetID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectSubnetAPI.SubnetDeleteObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubnetDeleteObjectV1`: CommonResponse
	fmt.Fprintf(os.Stdout, "Response from `ObjectSubnetAPI.SubnetDeleteObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiSubnetID** | **int32** | The unique ID of the Subnet | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubnetDeleteObjectV1Request struct via the builder pattern


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


## SubnetEditObjectV1

> CommonResponse SubnetEditObjectV1(ctx, pkiSubnetID).SubnetEditObjectV1Request(subnetEditObjectV1Request).Execute()

Edit an existing Subnet



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
	pkiSubnetID := int32(56) // int32 | The unique ID of the Subnet
	subnetEditObjectV1Request := *openapiclient.NewSubnetEditObjectV1Request(*openapiclient.NewSubnetRequestCompound(*openapiclient.NewMultilingualSubnetDescription(), int64(134744064), int64(4294967040))) // SubnetEditObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectSubnetAPI.SubnetEditObjectV1(context.Background(), pkiSubnetID).SubnetEditObjectV1Request(subnetEditObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectSubnetAPI.SubnetEditObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubnetEditObjectV1`: CommonResponse
	fmt.Fprintf(os.Stdout, "Response from `ObjectSubnetAPI.SubnetEditObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiSubnetID** | **int32** | The unique ID of the Subnet | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubnetEditObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subnetEditObjectV1Request** | [**SubnetEditObjectV1Request**](SubnetEditObjectV1Request.md) |  | 

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


## SubnetGetObjectV2

> SubnetGetObjectV2Response SubnetGetObjectV2(ctx, pkiSubnetID).Execute()

Retrieve an existing Subnet



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
	pkiSubnetID := int32(56) // int32 | The unique ID of the Subnet

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectSubnetAPI.SubnetGetObjectV2(context.Background(), pkiSubnetID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectSubnetAPI.SubnetGetObjectV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubnetGetObjectV2`: SubnetGetObjectV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectSubnetAPI.SubnetGetObjectV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiSubnetID** | **int32** | The unique ID of the Subnet | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubnetGetObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SubnetGetObjectV2Response**](SubnetGetObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

