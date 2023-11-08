# eZmaxAPI\ObjectUsergroupAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsergroupCreateObjectV1**](ObjectUsergroupAPI.md#UsergroupCreateObjectV1) | **Post** /1/object/usergroup | Create a new Usergroup
[**UsergroupEditObjectV1**](ObjectUsergroupAPI.md#UsergroupEditObjectV1) | **Put** /1/object/usergroup/{pkiUsergroupID} | Edit an existing Usergroup
[**UsergroupEditPermissionsV1**](ObjectUsergroupAPI.md#UsergroupEditPermissionsV1) | **Put** /1/object/usergroup/{pkiUsergroupID}/editPermissions | Edit multiple Permissions
[**UsergroupEditUsergroupdelegationsV1**](ObjectUsergroupAPI.md#UsergroupEditUsergroupdelegationsV1) | **Put** /1/object/usergroup/{pkiUsergroupID}/editUsergroupdelegations | Edit multiple Usergroupdelegations
[**UsergroupEditUsergroupmembershipsV1**](ObjectUsergroupAPI.md#UsergroupEditUsergroupmembershipsV1) | **Put** /1/object/usergroup/{pkiUsergroupID}/editUsergroupmemberships | Edit multiple Usergroupmemberships
[**UsergroupGetAutocompleteV2**](ObjectUsergroupAPI.md#UsergroupGetAutocompleteV2) | **Get** /2/object/usergroup/getAutocomplete/{sSelector} | Retrieve Usergroups and IDs
[**UsergroupGetListV1**](ObjectUsergroupAPI.md#UsergroupGetListV1) | **Get** /1/object/usergroup/getList | Retrieve Usergroup list
[**UsergroupGetObjectV2**](ObjectUsergroupAPI.md#UsergroupGetObjectV2) | **Get** /2/object/usergroup/{pkiUsergroupID} | Retrieve an existing Usergroup
[**UsergroupGetPermissionsV1**](ObjectUsergroupAPI.md#UsergroupGetPermissionsV1) | **Get** /1/object/usergroup/{pkiUsergroupID}/getPermissions | Retrieve an existing Usergroup&#39;s Permissions
[**UsergroupGetUsergroupdelegationsV1**](ObjectUsergroupAPI.md#UsergroupGetUsergroupdelegationsV1) | **Get** /1/object/usergroup/{pkiUsergroupID}/getUsergroupdelegations | Retrieve an existing Usergroup&#39;s Usergroupdelegations
[**UsergroupGetUsergroupmembershipsV1**](ObjectUsergroupAPI.md#UsergroupGetUsergroupmembershipsV1) | **Get** /1/object/usergroup/{pkiUsergroupID}/getUsergroupmemberships | Retrieve an existing Usergroup&#39;s Usergroupmemberships



## UsergroupCreateObjectV1

> UsergroupCreateObjectV1Response UsergroupCreateObjectV1(ctx).UsergroupCreateObjectV1Request(usergroupCreateObjectV1Request).Execute()

Create a new Usergroup



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
    usergroupCreateObjectV1Request := *openapiclient.NewUsergroupCreateObjectV1Request([]openapiclient.UsergroupRequestCompound{*openapiclient.NewUsergroupRequestCompound(*openapiclient.NewMultilingualUsergroupName())}) // UsergroupCreateObjectV1Request | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectUsergroupAPI.UsergroupCreateObjectV1(context.Background()).UsergroupCreateObjectV1Request(usergroupCreateObjectV1Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectUsergroupAPI.UsergroupCreateObjectV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsergroupCreateObjectV1`: UsergroupCreateObjectV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectUsergroupAPI.UsergroupCreateObjectV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsergroupCreateObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **usergroupCreateObjectV1Request** | [**UsergroupCreateObjectV1Request**](UsergroupCreateObjectV1Request.md) |  | 

### Return type

[**UsergroupCreateObjectV1Response**](UsergroupCreateObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsergroupEditObjectV1

> UsergroupEditObjectV1Response UsergroupEditObjectV1(ctx, pkiUsergroupID).UsergroupEditObjectV1Request(usergroupEditObjectV1Request).Execute()

Edit an existing Usergroup



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
    pkiUsergroupID := int32(56) // int32 | 
    usergroupEditObjectV1Request := *openapiclient.NewUsergroupEditObjectV1Request(*openapiclient.NewUsergroupRequestCompound(*openapiclient.NewMultilingualUsergroupName())) // UsergroupEditObjectV1Request | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectUsergroupAPI.UsergroupEditObjectV1(context.Background(), pkiUsergroupID).UsergroupEditObjectV1Request(usergroupEditObjectV1Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectUsergroupAPI.UsergroupEditObjectV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsergroupEditObjectV1`: UsergroupEditObjectV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectUsergroupAPI.UsergroupEditObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiUsergroupID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsergroupEditObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **usergroupEditObjectV1Request** | [**UsergroupEditObjectV1Request**](UsergroupEditObjectV1Request.md) |  | 

### Return type

[**UsergroupEditObjectV1Response**](UsergroupEditObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsergroupEditPermissionsV1

> UsergroupEditPermissionsV1Response UsergroupEditPermissionsV1(ctx, pkiUsergroupID).UsergroupEditPermissionsV1Request(usergroupEditPermissionsV1Request).Execute()

Edit multiple Permissions



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
    pkiUsergroupID := int32(56) // int32 | 
    usergroupEditPermissionsV1Request := *openapiclient.NewUsergroupEditPermissionsV1Request([]openapiclient.PermissionRequestCompound{*openapiclient.NewPermissionRequestCompound(int32(53))}) // UsergroupEditPermissionsV1Request | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectUsergroupAPI.UsergroupEditPermissionsV1(context.Background(), pkiUsergroupID).UsergroupEditPermissionsV1Request(usergroupEditPermissionsV1Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectUsergroupAPI.UsergroupEditPermissionsV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsergroupEditPermissionsV1`: UsergroupEditPermissionsV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectUsergroupAPI.UsergroupEditPermissionsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiUsergroupID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsergroupEditPermissionsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **usergroupEditPermissionsV1Request** | [**UsergroupEditPermissionsV1Request**](UsergroupEditPermissionsV1Request.md) |  | 

### Return type

[**UsergroupEditPermissionsV1Response**](UsergroupEditPermissionsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsergroupEditUsergroupdelegationsV1

> UsergroupEditUsergroupdelegationsV1Response UsergroupEditUsergroupdelegationsV1(ctx, pkiUsergroupID).UsergroupEditUsergroupdelegationsV1Request(usergroupEditUsergroupdelegationsV1Request).Execute()

Edit multiple Usergroupdelegations



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
    pkiUsergroupID := int32(56) // int32 | 
    usergroupEditUsergroupdelegationsV1Request := *openapiclient.NewUsergroupEditUsergroupdelegationsV1Request([]openapiclient.UsergroupdelegationRequestCompound{*openapiclient.NewUsergroupdelegationRequestCompound(int32(2), int32(70))}) // UsergroupEditUsergroupdelegationsV1Request | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectUsergroupAPI.UsergroupEditUsergroupdelegationsV1(context.Background(), pkiUsergroupID).UsergroupEditUsergroupdelegationsV1Request(usergroupEditUsergroupdelegationsV1Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectUsergroupAPI.UsergroupEditUsergroupdelegationsV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsergroupEditUsergroupdelegationsV1`: UsergroupEditUsergroupdelegationsV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectUsergroupAPI.UsergroupEditUsergroupdelegationsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiUsergroupID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsergroupEditUsergroupdelegationsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **usergroupEditUsergroupdelegationsV1Request** | [**UsergroupEditUsergroupdelegationsV1Request**](UsergroupEditUsergroupdelegationsV1Request.md) |  | 

### Return type

[**UsergroupEditUsergroupdelegationsV1Response**](UsergroupEditUsergroupdelegationsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsergroupEditUsergroupmembershipsV1

> UsergroupEditUsergroupmembershipsV1Response UsergroupEditUsergroupmembershipsV1(ctx, pkiUsergroupID).UsergroupEditUsergroupmembershipsV1Request(usergroupEditUsergroupmembershipsV1Request).Execute()

Edit multiple Usergroupmemberships



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
    pkiUsergroupID := int32(56) // int32 | 
    usergroupEditUsergroupmembershipsV1Request := *openapiclient.NewUsergroupEditUsergroupmembershipsV1Request([]openapiclient.UsergroupmembershipRequestCompound{*openapiclient.NewUsergroupmembershipRequestCompound(int32(2), int32(70))}) // UsergroupEditUsergroupmembershipsV1Request | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectUsergroupAPI.UsergroupEditUsergroupmembershipsV1(context.Background(), pkiUsergroupID).UsergroupEditUsergroupmembershipsV1Request(usergroupEditUsergroupmembershipsV1Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectUsergroupAPI.UsergroupEditUsergroupmembershipsV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsergroupEditUsergroupmembershipsV1`: UsergroupEditUsergroupmembershipsV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectUsergroupAPI.UsergroupEditUsergroupmembershipsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiUsergroupID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsergroupEditUsergroupmembershipsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **usergroupEditUsergroupmembershipsV1Request** | [**UsergroupEditUsergroupmembershipsV1Request**](UsergroupEditUsergroupmembershipsV1Request.md) |  | 

### Return type

[**UsergroupEditUsergroupmembershipsV1Response**](UsergroupEditUsergroupmembershipsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsergroupGetAutocompleteV2

> UsergroupGetAutocompleteV2Response UsergroupGetAutocompleteV2(ctx, sSelector).EFilterActive(eFilterActive).SQuery(sQuery).AcceptLanguage(acceptLanguage).Execute()

Retrieve Usergroups and IDs



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
    sSelector := "sSelector_example" // string | The type of Usergroups to return
    eFilterActive := "eFilterActive_example" // string | Specify which results we want to display. (optional) (default to "Active")
    sQuery := "sQuery_example" // string | Allow to filter the returned results (optional)
    acceptLanguage := openapiclient.Header-Accept-Language("*") // HeaderAcceptLanguage |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectUsergroupAPI.UsergroupGetAutocompleteV2(context.Background(), sSelector).EFilterActive(eFilterActive).SQuery(sQuery).AcceptLanguage(acceptLanguage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectUsergroupAPI.UsergroupGetAutocompleteV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsergroupGetAutocompleteV2`: UsergroupGetAutocompleteV2Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectUsergroupAPI.UsergroupGetAutocompleteV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sSelector** | **string** | The type of Usergroups to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsergroupGetAutocompleteV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eFilterActive** | **string** | Specify which results we want to display. | [default to &quot;Active&quot;]
 **sQuery** | **string** | Allow to filter the returned results | 
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 

### Return type

[**UsergroupGetAutocompleteV2Response**](UsergroupGetAutocompleteV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsergroupGetListV1

> UsergroupGetListV1Response UsergroupGetListV1(ctx).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()

Retrieve Usergroup list



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
    resp, r, err := apiClient.ObjectUsergroupAPI.UsergroupGetListV1(context.Background()).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectUsergroupAPI.UsergroupGetListV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsergroupGetListV1`: UsergroupGetListV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectUsergroupAPI.UsergroupGetListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsergroupGetListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eOrderBy** | **string** | Specify how you want the results to be sorted | 
 **iRowMax** | **int32** |  | 
 **iRowOffset** | **int32** |  | [default to 0]
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 
 **sFilter** | **string** |  | 

### Return type

[**UsergroupGetListV1Response**](UsergroupGetListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsergroupGetObjectV2

> UsergroupGetObjectV2Response UsergroupGetObjectV2(ctx, pkiUsergroupID).Execute()

Retrieve an existing Usergroup



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
    pkiUsergroupID := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectUsergroupAPI.UsergroupGetObjectV2(context.Background(), pkiUsergroupID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectUsergroupAPI.UsergroupGetObjectV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsergroupGetObjectV2`: UsergroupGetObjectV2Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectUsergroupAPI.UsergroupGetObjectV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiUsergroupID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsergroupGetObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UsergroupGetObjectV2Response**](UsergroupGetObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsergroupGetPermissionsV1

> UsergroupGetPermissionsV1Response UsergroupGetPermissionsV1(ctx, pkiUsergroupID).Execute()

Retrieve an existing Usergroup's Permissions

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
    pkiUsergroupID := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectUsergroupAPI.UsergroupGetPermissionsV1(context.Background(), pkiUsergroupID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectUsergroupAPI.UsergroupGetPermissionsV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsergroupGetPermissionsV1`: UsergroupGetPermissionsV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectUsergroupAPI.UsergroupGetPermissionsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiUsergroupID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsergroupGetPermissionsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UsergroupGetPermissionsV1Response**](UsergroupGetPermissionsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsergroupGetUsergroupdelegationsV1

> UsergroupGetUsergroupdelegationsV1Response UsergroupGetUsergroupdelegationsV1(ctx, pkiUsergroupID).Execute()

Retrieve an existing Usergroup's Usergroupdelegations

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
    pkiUsergroupID := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectUsergroupAPI.UsergroupGetUsergroupdelegationsV1(context.Background(), pkiUsergroupID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectUsergroupAPI.UsergroupGetUsergroupdelegationsV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsergroupGetUsergroupdelegationsV1`: UsergroupGetUsergroupdelegationsV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectUsergroupAPI.UsergroupGetUsergroupdelegationsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiUsergroupID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsergroupGetUsergroupdelegationsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UsergroupGetUsergroupdelegationsV1Response**](UsergroupGetUsergroupdelegationsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsergroupGetUsergroupmembershipsV1

> UsergroupGetUsergroupmembershipsV1Response UsergroupGetUsergroupmembershipsV1(ctx, pkiUsergroupID).Execute()

Retrieve an existing Usergroup's Usergroupmemberships

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
    pkiUsergroupID := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectUsergroupAPI.UsergroupGetUsergroupmembershipsV1(context.Background(), pkiUsergroupID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectUsergroupAPI.UsergroupGetUsergroupmembershipsV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsergroupGetUsergroupmembershipsV1`: UsergroupGetUsergroupmembershipsV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectUsergroupAPI.UsergroupGetUsergroupmembershipsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiUsergroupID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsergroupGetUsergroupmembershipsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UsergroupGetUsergroupmembershipsV1Response**](UsergroupGetUsergroupmembershipsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

