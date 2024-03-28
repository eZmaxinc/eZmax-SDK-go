# eZmaxAPI\ObjectEzsignfoldersignerassociationAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EzsignfoldersignerassociationCreateEmbeddedUrlV1**](ObjectEzsignfoldersignerassociationAPI.md#EzsignfoldersignerassociationCreateEmbeddedUrlV1) | **Post** /1/object/ezsignfoldersignerassociation/{pkiEzsignfoldersignerassociationID}/createEmbeddedUrl | Creates an Url to allow embedded signing
[**EzsignfoldersignerassociationCreateObjectV1**](ObjectEzsignfoldersignerassociationAPI.md#EzsignfoldersignerassociationCreateObjectV1) | **Post** /1/object/ezsignfoldersignerassociation | Create a new Ezsignfoldersignerassociation
[**EzsignfoldersignerassociationCreateObjectV2**](ObjectEzsignfoldersignerassociationAPI.md#EzsignfoldersignerassociationCreateObjectV2) | **Post** /2/object/ezsignfoldersignerassociation | Create a new Ezsignfoldersignerassociation
[**EzsignfoldersignerassociationDeleteObjectV1**](ObjectEzsignfoldersignerassociationAPI.md#EzsignfoldersignerassociationDeleteObjectV1) | **Delete** /1/object/ezsignfoldersignerassociation/{pkiEzsignfoldersignerassociationID} | Delete an existing Ezsignfoldersignerassociation
[**EzsignfoldersignerassociationEditObjectV1**](ObjectEzsignfoldersignerassociationAPI.md#EzsignfoldersignerassociationEditObjectV1) | **Put** /1/object/ezsignfoldersignerassociation/{pkiEzsignfoldersignerassociationID} | Edit an existing Ezsignfoldersignerassociation
[**EzsignfoldersignerassociationForceDisconnectV1**](ObjectEzsignfoldersignerassociationAPI.md#EzsignfoldersignerassociationForceDisconnectV1) | **Post** /1/object/ezsignfoldersignerassociation/{pkiEzsignfoldersignerassociationID}/forceDisconnect | Disconnects the Ezsignfoldersignerassociation
[**EzsignfoldersignerassociationGetInPersonLoginUrlV1**](ObjectEzsignfoldersignerassociationAPI.md#EzsignfoldersignerassociationGetInPersonLoginUrlV1) | **Get** /1/object/ezsignfoldersignerassociation/{pkiEzsignfoldersignerassociationID}/getInPersonLoginUrl | Retrieve a Login Url to allow In-Person signing
[**EzsignfoldersignerassociationGetObjectV1**](ObjectEzsignfoldersignerassociationAPI.md#EzsignfoldersignerassociationGetObjectV1) | **Get** /1/object/ezsignfoldersignerassociation/{pkiEzsignfoldersignerassociationID} | Retrieve an existing Ezsignfoldersignerassociation
[**EzsignfoldersignerassociationGetObjectV2**](ObjectEzsignfoldersignerassociationAPI.md#EzsignfoldersignerassociationGetObjectV2) | **Get** /2/object/ezsignfoldersignerassociation/{pkiEzsignfoldersignerassociationID} | Retrieve an existing Ezsignfoldersignerassociation
[**EzsignfoldersignerassociationPatchObjectV1**](ObjectEzsignfoldersignerassociationAPI.md#EzsignfoldersignerassociationPatchObjectV1) | **Patch** /1/object/ezsignfoldersignerassociation/{pkiEzsignfoldersignerassociationID} | Patch an existing Ezsignfoldersignerassociation



## EzsignfoldersignerassociationCreateEmbeddedUrlV1

> EzsignfoldersignerassociationCreateEmbeddedUrlV1Response EzsignfoldersignerassociationCreateEmbeddedUrlV1(ctx, pkiEzsignfoldersignerassociationID).EzsignfoldersignerassociationCreateEmbeddedUrlV1Request(ezsignfoldersignerassociationCreateEmbeddedUrlV1Request).Execute()

Creates an Url to allow embedded signing



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
	pkiEzsignfoldersignerassociationID := int32(56) // int32 | 
	ezsignfoldersignerassociationCreateEmbeddedUrlV1Request := *openapiclient.NewEzsignfoldersignerassociationCreateEmbeddedUrlV1Request() // EzsignfoldersignerassociationCreateEmbeddedUrlV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsignfoldersignerassociationAPI.EzsignfoldersignerassociationCreateEmbeddedUrlV1(context.Background(), pkiEzsignfoldersignerassociationID).EzsignfoldersignerassociationCreateEmbeddedUrlV1Request(ezsignfoldersignerassociationCreateEmbeddedUrlV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignfoldersignerassociationAPI.EzsignfoldersignerassociationCreateEmbeddedUrlV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsignfoldersignerassociationCreateEmbeddedUrlV1`: EzsignfoldersignerassociationCreateEmbeddedUrlV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignfoldersignerassociationAPI.EzsignfoldersignerassociationCreateEmbeddedUrlV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignfoldersignerassociationID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignfoldersignerassociationCreateEmbeddedUrlV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezsignfoldersignerassociationCreateEmbeddedUrlV1Request** | [**EzsignfoldersignerassociationCreateEmbeddedUrlV1Request**](EzsignfoldersignerassociationCreateEmbeddedUrlV1Request.md) |  | 

### Return type

[**EzsignfoldersignerassociationCreateEmbeddedUrlV1Response**](EzsignfoldersignerassociationCreateEmbeddedUrlV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignfoldersignerassociationCreateObjectV1

> EzsignfoldersignerassociationCreateObjectV1Response EzsignfoldersignerassociationCreateObjectV1(ctx).EzsignfoldersignerassociationCreateObjectV1Request(ezsignfoldersignerassociationCreateObjectV1Request).Execute()

Create a new Ezsignfoldersignerassociation



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
	ezsignfoldersignerassociationCreateObjectV1Request := []openapiclient.EzsignfoldersignerassociationCreateObjectV1Request{*openapiclient.NewEzsignfoldersignerassociationCreateObjectV1Request()} // []EzsignfoldersignerassociationCreateObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsignfoldersignerassociationAPI.EzsignfoldersignerassociationCreateObjectV1(context.Background()).EzsignfoldersignerassociationCreateObjectV1Request(ezsignfoldersignerassociationCreateObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignfoldersignerassociationAPI.EzsignfoldersignerassociationCreateObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsignfoldersignerassociationCreateObjectV1`: EzsignfoldersignerassociationCreateObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignfoldersignerassociationAPI.EzsignfoldersignerassociationCreateObjectV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEzsignfoldersignerassociationCreateObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ezsignfoldersignerassociationCreateObjectV1Request** | [**[]EzsignfoldersignerassociationCreateObjectV1Request**](EzsignfoldersignerassociationCreateObjectV1Request.md) |  | 

### Return type

[**EzsignfoldersignerassociationCreateObjectV1Response**](EzsignfoldersignerassociationCreateObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignfoldersignerassociationCreateObjectV2

> EzsignfoldersignerassociationCreateObjectV2Response EzsignfoldersignerassociationCreateObjectV2(ctx).EzsignfoldersignerassociationCreateObjectV2Request(ezsignfoldersignerassociationCreateObjectV2Request).Execute()

Create a new Ezsignfoldersignerassociation



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
	ezsignfoldersignerassociationCreateObjectV2Request := *openapiclient.NewEzsignfoldersignerassociationCreateObjectV2Request([]openapiclient.EzsignfoldersignerassociationRequestCompound{*openapiclient.NewEzsignfoldersignerassociationRequestCompound(int32(33))}) // EzsignfoldersignerassociationCreateObjectV2Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsignfoldersignerassociationAPI.EzsignfoldersignerassociationCreateObjectV2(context.Background()).EzsignfoldersignerassociationCreateObjectV2Request(ezsignfoldersignerassociationCreateObjectV2Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignfoldersignerassociationAPI.EzsignfoldersignerassociationCreateObjectV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsignfoldersignerassociationCreateObjectV2`: EzsignfoldersignerassociationCreateObjectV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignfoldersignerassociationAPI.EzsignfoldersignerassociationCreateObjectV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEzsignfoldersignerassociationCreateObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ezsignfoldersignerassociationCreateObjectV2Request** | [**EzsignfoldersignerassociationCreateObjectV2Request**](EzsignfoldersignerassociationCreateObjectV2Request.md) |  | 

### Return type

[**EzsignfoldersignerassociationCreateObjectV2Response**](EzsignfoldersignerassociationCreateObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignfoldersignerassociationDeleteObjectV1

> EzsignfoldersignerassociationDeleteObjectV1Response EzsignfoldersignerassociationDeleteObjectV1(ctx, pkiEzsignfoldersignerassociationID).Execute()

Delete an existing Ezsignfoldersignerassociation



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
	pkiEzsignfoldersignerassociationID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsignfoldersignerassociationAPI.EzsignfoldersignerassociationDeleteObjectV1(context.Background(), pkiEzsignfoldersignerassociationID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignfoldersignerassociationAPI.EzsignfoldersignerassociationDeleteObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsignfoldersignerassociationDeleteObjectV1`: EzsignfoldersignerassociationDeleteObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignfoldersignerassociationAPI.EzsignfoldersignerassociationDeleteObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignfoldersignerassociationID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignfoldersignerassociationDeleteObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsignfoldersignerassociationDeleteObjectV1Response**](EzsignfoldersignerassociationDeleteObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignfoldersignerassociationEditObjectV1

> EzsignfoldersignerassociationEditObjectV1Response EzsignfoldersignerassociationEditObjectV1(ctx, pkiEzsignfoldersignerassociationID).EzsignfoldersignerassociationEditObjectV1Request(ezsignfoldersignerassociationEditObjectV1Request).Execute()

Edit an existing Ezsignfoldersignerassociation



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
	pkiEzsignfoldersignerassociationID := int32(56) // int32 | 
	ezsignfoldersignerassociationEditObjectV1Request := *openapiclient.NewEzsignfoldersignerassociationEditObjectV1Request(*openapiclient.NewEzsignfoldersignerassociationRequestCompound(int32(33))) // EzsignfoldersignerassociationEditObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsignfoldersignerassociationAPI.EzsignfoldersignerassociationEditObjectV1(context.Background(), pkiEzsignfoldersignerassociationID).EzsignfoldersignerassociationEditObjectV1Request(ezsignfoldersignerassociationEditObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignfoldersignerassociationAPI.EzsignfoldersignerassociationEditObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsignfoldersignerassociationEditObjectV1`: EzsignfoldersignerassociationEditObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignfoldersignerassociationAPI.EzsignfoldersignerassociationEditObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignfoldersignerassociationID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignfoldersignerassociationEditObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezsignfoldersignerassociationEditObjectV1Request** | [**EzsignfoldersignerassociationEditObjectV1Request**](EzsignfoldersignerassociationEditObjectV1Request.md) |  | 

### Return type

[**EzsignfoldersignerassociationEditObjectV1Response**](EzsignfoldersignerassociationEditObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignfoldersignerassociationForceDisconnectV1

> EzsignfoldersignerassociationForceDisconnectV1Response EzsignfoldersignerassociationForceDisconnectV1(ctx, pkiEzsignfoldersignerassociationID).Body(body).Execute()

Disconnects the Ezsignfoldersignerassociation



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
	pkiEzsignfoldersignerassociationID := int32(56) // int32 | 
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsignfoldersignerassociationAPI.EzsignfoldersignerassociationForceDisconnectV1(context.Background(), pkiEzsignfoldersignerassociationID).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignfoldersignerassociationAPI.EzsignfoldersignerassociationForceDisconnectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsignfoldersignerassociationForceDisconnectV1`: EzsignfoldersignerassociationForceDisconnectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignfoldersignerassociationAPI.EzsignfoldersignerassociationForceDisconnectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignfoldersignerassociationID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignfoldersignerassociationForceDisconnectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

[**EzsignfoldersignerassociationForceDisconnectV1Response**](EzsignfoldersignerassociationForceDisconnectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignfoldersignerassociationGetInPersonLoginUrlV1

> EzsignfoldersignerassociationGetInPersonLoginUrlV1Response EzsignfoldersignerassociationGetInPersonLoginUrlV1(ctx, pkiEzsignfoldersignerassociationID).Execute()

Retrieve a Login Url to allow In-Person signing



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
	pkiEzsignfoldersignerassociationID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsignfoldersignerassociationAPI.EzsignfoldersignerassociationGetInPersonLoginUrlV1(context.Background(), pkiEzsignfoldersignerassociationID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignfoldersignerassociationAPI.EzsignfoldersignerassociationGetInPersonLoginUrlV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsignfoldersignerassociationGetInPersonLoginUrlV1`: EzsignfoldersignerassociationGetInPersonLoginUrlV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignfoldersignerassociationAPI.EzsignfoldersignerassociationGetInPersonLoginUrlV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignfoldersignerassociationID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignfoldersignerassociationGetInPersonLoginUrlV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsignfoldersignerassociationGetInPersonLoginUrlV1Response**](EzsignfoldersignerassociationGetInPersonLoginUrlV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignfoldersignerassociationGetObjectV1

> EzsignfoldersignerassociationGetObjectV1Response EzsignfoldersignerassociationGetObjectV1(ctx, pkiEzsignfoldersignerassociationID).Execute()

Retrieve an existing Ezsignfoldersignerassociation



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
	pkiEzsignfoldersignerassociationID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsignfoldersignerassociationAPI.EzsignfoldersignerassociationGetObjectV1(context.Background(), pkiEzsignfoldersignerassociationID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignfoldersignerassociationAPI.EzsignfoldersignerassociationGetObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsignfoldersignerassociationGetObjectV1`: EzsignfoldersignerassociationGetObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignfoldersignerassociationAPI.EzsignfoldersignerassociationGetObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignfoldersignerassociationID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignfoldersignerassociationGetObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsignfoldersignerassociationGetObjectV1Response**](EzsignfoldersignerassociationGetObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignfoldersignerassociationGetObjectV2

> EzsignfoldersignerassociationGetObjectV2Response EzsignfoldersignerassociationGetObjectV2(ctx, pkiEzsignfoldersignerassociationID).Execute()

Retrieve an existing Ezsignfoldersignerassociation



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
	pkiEzsignfoldersignerassociationID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsignfoldersignerassociationAPI.EzsignfoldersignerassociationGetObjectV2(context.Background(), pkiEzsignfoldersignerassociationID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignfoldersignerassociationAPI.EzsignfoldersignerassociationGetObjectV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsignfoldersignerassociationGetObjectV2`: EzsignfoldersignerassociationGetObjectV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignfoldersignerassociationAPI.EzsignfoldersignerassociationGetObjectV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignfoldersignerassociationID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignfoldersignerassociationGetObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsignfoldersignerassociationGetObjectV2Response**](EzsignfoldersignerassociationGetObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignfoldersignerassociationPatchObjectV1

> EzsignfoldersignerassociationPatchObjectV1Response EzsignfoldersignerassociationPatchObjectV1(ctx, pkiEzsignfoldersignerassociationID).EzsignfoldersignerassociationPatchObjectV1Request(ezsignfoldersignerassociationPatchObjectV1Request).Execute()

Patch an existing Ezsignfoldersignerassociation

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
	pkiEzsignfoldersignerassociationID := int32(56) // int32 | 
	ezsignfoldersignerassociationPatchObjectV1Request := *openapiclient.NewEzsignfoldersignerassociationPatchObjectV1Request(*openapiclient.NewEzsignfoldersignerassociationRequestPatch()) // EzsignfoldersignerassociationPatchObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsignfoldersignerassociationAPI.EzsignfoldersignerassociationPatchObjectV1(context.Background(), pkiEzsignfoldersignerassociationID).EzsignfoldersignerassociationPatchObjectV1Request(ezsignfoldersignerassociationPatchObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignfoldersignerassociationAPI.EzsignfoldersignerassociationPatchObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsignfoldersignerassociationPatchObjectV1`: EzsignfoldersignerassociationPatchObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignfoldersignerassociationAPI.EzsignfoldersignerassociationPatchObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignfoldersignerassociationID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignfoldersignerassociationPatchObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezsignfoldersignerassociationPatchObjectV1Request** | [**EzsignfoldersignerassociationPatchObjectV1Request**](EzsignfoldersignerassociationPatchObjectV1Request.md) |  | 

### Return type

[**EzsignfoldersignerassociationPatchObjectV1Response**](EzsignfoldersignerassociationPatchObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

