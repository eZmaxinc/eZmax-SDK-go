# eZmaxAPI\ScimGroupsAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupsCreateObjectScimV2**](ScimGroupsAPI.md#GroupsCreateObjectScimV2) | **Post** /2/scim/Groups | Create a new Usergroup
[**GroupsDeleteObjectScimV2**](ScimGroupsAPI.md#GroupsDeleteObjectScimV2) | **Delete** /2/scim/Groups/{groupId} | Delete an existing Usergroup
[**GroupsEditObjectScimV2**](ScimGroupsAPI.md#GroupsEditObjectScimV2) | **Put** /2/scim/Groups/{groupId} | Edit an existing Usergroup
[**GroupsGetListScimV2**](ScimGroupsAPI.md#GroupsGetListScimV2) | **Get** /2/scim/Groups | Retrieve Usergroup list
[**GroupsGetObjectScimV2**](ScimGroupsAPI.md#GroupsGetObjectScimV2) | **Get** /2/scim/Groups/{groupId} | Retrieve an existing Usergroup



## GroupsCreateObjectScimV2

> ScimGroup GroupsCreateObjectScimV2(ctx).ScimGroup(scimGroup).Execute()

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
    scimGroup := *openapiclient.NewScimGroup("Administration") // ScimGroup | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScimGroupsAPI.GroupsCreateObjectScimV2(context.Background()).ScimGroup(scimGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScimGroupsAPI.GroupsCreateObjectScimV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsCreateObjectScimV2`: ScimGroup
    fmt.Fprintf(os.Stdout, "Response from `ScimGroupsAPI.GroupsCreateObjectScimV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGroupsCreateObjectScimV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scimGroup** | [**ScimGroup**](ScimGroup.md) |  | 

### Return type

[**ScimGroup**](ScimGroup.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsDeleteObjectScimV2

> GroupsDeleteObjectScimV2(ctx, groupId).Execute()

Delete an existing Usergroup

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
    groupId := "groupId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ScimGroupsAPI.GroupsDeleteObjectScimV2(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScimGroupsAPI.GroupsDeleteObjectScimV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsDeleteObjectScimV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsEditObjectScimV2

> ScimGroup GroupsEditObjectScimV2(ctx, groupId).ScimGroup(scimGroup).Execute()

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
    groupId := "groupId_example" // string | 
    scimGroup := *openapiclient.NewScimGroup("Administration") // ScimGroup | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScimGroupsAPI.GroupsEditObjectScimV2(context.Background(), groupId).ScimGroup(scimGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScimGroupsAPI.GroupsEditObjectScimV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsEditObjectScimV2`: ScimGroup
    fmt.Fprintf(os.Stdout, "Response from `ScimGroupsAPI.GroupsEditObjectScimV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsEditObjectScimV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scimGroup** | [**ScimGroup**](ScimGroup.md) |  | 

### Return type

[**ScimGroup**](ScimGroup.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsGetListScimV2

> ScimGroup GroupsGetListScimV2(ctx).Filter(filter).Execute()

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
    filter := "filter_example" // string | Filter expression for searching groups (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScimGroupsAPI.GroupsGetListScimV2(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScimGroupsAPI.GroupsGetListScimV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsGetListScimV2`: ScimGroup
    fmt.Fprintf(os.Stdout, "Response from `ScimGroupsAPI.GroupsGetListScimV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGroupsGetListScimV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | Filter expression for searching groups | 

### Return type

[**ScimGroup**](ScimGroup.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsGetObjectScimV2

> ScimGroup GroupsGetObjectScimV2(ctx, groupId).Execute()

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
    groupId := "groupId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScimGroupsAPI.GroupsGetObjectScimV2(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScimGroupsAPI.GroupsGetObjectScimV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsGetObjectScimV2`: ScimGroup
    fmt.Fprintf(os.Stdout, "Response from `ScimGroupsAPI.GroupsGetObjectScimV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsGetObjectScimV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ScimGroup**](ScimGroup.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

