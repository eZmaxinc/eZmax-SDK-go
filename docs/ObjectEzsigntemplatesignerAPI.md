# eZmaxAPI\ObjectEzsigntemplatesignerAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EzsigntemplatesignerCreateObjectV1**](ObjectEzsigntemplatesignerAPI.md#EzsigntemplatesignerCreateObjectV1) | **Post** /1/object/ezsigntemplatesigner | Create a new Ezsigntemplatesigner
[**EzsigntemplatesignerCreateObjectV2**](ObjectEzsigntemplatesignerAPI.md#EzsigntemplatesignerCreateObjectV2) | **Post** /2/object/ezsigntemplatesigner | Create a new Ezsigntemplatesigner
[**EzsigntemplatesignerDeleteObjectV1**](ObjectEzsigntemplatesignerAPI.md#EzsigntemplatesignerDeleteObjectV1) | **Delete** /1/object/ezsigntemplatesigner/{pkiEzsigntemplatesignerID} | Delete an existing Ezsigntemplatesigner
[**EzsigntemplatesignerEditObjectV1**](ObjectEzsigntemplatesignerAPI.md#EzsigntemplatesignerEditObjectV1) | **Put** /1/object/ezsigntemplatesigner/{pkiEzsigntemplatesignerID} | Edit an existing Ezsigntemplatesigner
[**EzsigntemplatesignerEditObjectV2**](ObjectEzsigntemplatesignerAPI.md#EzsigntemplatesignerEditObjectV2) | **Put** /2/object/ezsigntemplatesigner/{pkiEzsigntemplatesignerID} | Edit an existing Ezsigntemplatesigner
[**EzsigntemplatesignerGetObjectV2**](ObjectEzsigntemplatesignerAPI.md#EzsigntemplatesignerGetObjectV2) | **Get** /2/object/ezsigntemplatesigner/{pkiEzsigntemplatesignerID} | Retrieve an existing Ezsigntemplatesigner
[**EzsigntemplatesignerGetObjectV3**](ObjectEzsigntemplatesignerAPI.md#EzsigntemplatesignerGetObjectV3) | **Get** /3/object/ezsigntemplatesigner/{pkiEzsigntemplatesignerID} | Retrieve an existing Ezsigntemplatesigner



## EzsigntemplatesignerCreateObjectV1

> EzsigntemplatesignerCreateObjectV1Response EzsigntemplatesignerCreateObjectV1(ctx).EzsigntemplatesignerCreateObjectV1Request(ezsigntemplatesignerCreateObjectV1Request).Execute()

Create a new Ezsigntemplatesigner



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
	ezsigntemplatesignerCreateObjectV1Request := *openapiclient.NewEzsigntemplatesignerCreateObjectV1Request([]openapiclient.EzsigntemplatesignerRequestCompound{*openapiclient.NewEzsigntemplatesignerRequestCompound(int32(36), "Customer")}) // EzsigntemplatesignerCreateObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigntemplatesignerAPI.EzsigntemplatesignerCreateObjectV1(context.Background()).EzsigntemplatesignerCreateObjectV1Request(ezsigntemplatesignerCreateObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplatesignerAPI.EzsigntemplatesignerCreateObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigntemplatesignerCreateObjectV1`: EzsigntemplatesignerCreateObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplatesignerAPI.EzsigntemplatesignerCreateObjectV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplatesignerCreateObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ezsigntemplatesignerCreateObjectV1Request** | [**EzsigntemplatesignerCreateObjectV1Request**](EzsigntemplatesignerCreateObjectV1Request.md) |  | 

### Return type

[**EzsigntemplatesignerCreateObjectV1Response**](EzsigntemplatesignerCreateObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigntemplatesignerCreateObjectV2

> EzsigntemplatesignerCreateObjectV2Response EzsigntemplatesignerCreateObjectV2(ctx).EzsigntemplatesignerCreateObjectV2Request(ezsigntemplatesignerCreateObjectV2Request).Execute()

Create a new Ezsigntemplatesigner



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
	ezsigntemplatesignerCreateObjectV2Request := *openapiclient.NewEzsigntemplatesignerCreateObjectV2Request([]openapiclient.EzsigntemplatesignerRequestCompoundV2{*openapiclient.NewEzsigntemplatesignerRequestCompoundV2(int32(36), "Customer")}) // EzsigntemplatesignerCreateObjectV2Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigntemplatesignerAPI.EzsigntemplatesignerCreateObjectV2(context.Background()).EzsigntemplatesignerCreateObjectV2Request(ezsigntemplatesignerCreateObjectV2Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplatesignerAPI.EzsigntemplatesignerCreateObjectV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigntemplatesignerCreateObjectV2`: EzsigntemplatesignerCreateObjectV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplatesignerAPI.EzsigntemplatesignerCreateObjectV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplatesignerCreateObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ezsigntemplatesignerCreateObjectV2Request** | [**EzsigntemplatesignerCreateObjectV2Request**](EzsigntemplatesignerCreateObjectV2Request.md) |  | 

### Return type

[**EzsigntemplatesignerCreateObjectV2Response**](EzsigntemplatesignerCreateObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigntemplatesignerDeleteObjectV1

> EzsigntemplatesignerDeleteObjectV1Response EzsigntemplatesignerDeleteObjectV1(ctx, pkiEzsigntemplatesignerID).Execute()

Delete an existing Ezsigntemplatesigner



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
	pkiEzsigntemplatesignerID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigntemplatesignerAPI.EzsigntemplatesignerDeleteObjectV1(context.Background(), pkiEzsigntemplatesignerID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplatesignerAPI.EzsigntemplatesignerDeleteObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigntemplatesignerDeleteObjectV1`: EzsigntemplatesignerDeleteObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplatesignerAPI.EzsigntemplatesignerDeleteObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigntemplatesignerID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplatesignerDeleteObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsigntemplatesignerDeleteObjectV1Response**](EzsigntemplatesignerDeleteObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigntemplatesignerEditObjectV1

> EzsigntemplatesignerEditObjectV1Response EzsigntemplatesignerEditObjectV1(ctx, pkiEzsigntemplatesignerID).EzsigntemplatesignerEditObjectV1Request(ezsigntemplatesignerEditObjectV1Request).Execute()

Edit an existing Ezsigntemplatesigner



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
	pkiEzsigntemplatesignerID := int32(56) // int32 | 
	ezsigntemplatesignerEditObjectV1Request := *openapiclient.NewEzsigntemplatesignerEditObjectV1Request(*openapiclient.NewEzsigntemplatesignerRequestCompound(int32(36), "Customer")) // EzsigntemplatesignerEditObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigntemplatesignerAPI.EzsigntemplatesignerEditObjectV1(context.Background(), pkiEzsigntemplatesignerID).EzsigntemplatesignerEditObjectV1Request(ezsigntemplatesignerEditObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplatesignerAPI.EzsigntemplatesignerEditObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigntemplatesignerEditObjectV1`: EzsigntemplatesignerEditObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplatesignerAPI.EzsigntemplatesignerEditObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigntemplatesignerID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplatesignerEditObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezsigntemplatesignerEditObjectV1Request** | [**EzsigntemplatesignerEditObjectV1Request**](EzsigntemplatesignerEditObjectV1Request.md) |  | 

### Return type

[**EzsigntemplatesignerEditObjectV1Response**](EzsigntemplatesignerEditObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigntemplatesignerEditObjectV2

> EzsigntemplatesignerEditObjectV2Response EzsigntemplatesignerEditObjectV2(ctx, pkiEzsigntemplatesignerID).EzsigntemplatesignerEditObjectV2Request(ezsigntemplatesignerEditObjectV2Request).Execute()

Edit an existing Ezsigntemplatesigner



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
	pkiEzsigntemplatesignerID := int32(56) // int32 | 
	ezsigntemplatesignerEditObjectV2Request := *openapiclient.NewEzsigntemplatesignerEditObjectV2Request(*openapiclient.NewEzsigntemplatesignerRequestCompoundV2(int32(36), "Customer")) // EzsigntemplatesignerEditObjectV2Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigntemplatesignerAPI.EzsigntemplatesignerEditObjectV2(context.Background(), pkiEzsigntemplatesignerID).EzsigntemplatesignerEditObjectV2Request(ezsigntemplatesignerEditObjectV2Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplatesignerAPI.EzsigntemplatesignerEditObjectV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigntemplatesignerEditObjectV2`: EzsigntemplatesignerEditObjectV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplatesignerAPI.EzsigntemplatesignerEditObjectV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigntemplatesignerID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplatesignerEditObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezsigntemplatesignerEditObjectV2Request** | [**EzsigntemplatesignerEditObjectV2Request**](EzsigntemplatesignerEditObjectV2Request.md) |  | 

### Return type

[**EzsigntemplatesignerEditObjectV2Response**](EzsigntemplatesignerEditObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigntemplatesignerGetObjectV2

> EzsigntemplatesignerGetObjectV2Response EzsigntemplatesignerGetObjectV2(ctx, pkiEzsigntemplatesignerID).Execute()

Retrieve an existing Ezsigntemplatesigner



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
	pkiEzsigntemplatesignerID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigntemplatesignerAPI.EzsigntemplatesignerGetObjectV2(context.Background(), pkiEzsigntemplatesignerID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplatesignerAPI.EzsigntemplatesignerGetObjectV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigntemplatesignerGetObjectV2`: EzsigntemplatesignerGetObjectV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplatesignerAPI.EzsigntemplatesignerGetObjectV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigntemplatesignerID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplatesignerGetObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsigntemplatesignerGetObjectV2Response**](EzsigntemplatesignerGetObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigntemplatesignerGetObjectV3

> EzsigntemplatesignerGetObjectV3Response EzsigntemplatesignerGetObjectV3(ctx, pkiEzsigntemplatesignerID).Execute()

Retrieve an existing Ezsigntemplatesigner



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
	pkiEzsigntemplatesignerID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigntemplatesignerAPI.EzsigntemplatesignerGetObjectV3(context.Background(), pkiEzsigntemplatesignerID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplatesignerAPI.EzsigntemplatesignerGetObjectV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigntemplatesignerGetObjectV3`: EzsigntemplatesignerGetObjectV3Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplatesignerAPI.EzsigntemplatesignerGetObjectV3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigntemplatesignerID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplatesignerGetObjectV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsigntemplatesignerGetObjectV3Response**](EzsigntemplatesignerGetObjectV3Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

