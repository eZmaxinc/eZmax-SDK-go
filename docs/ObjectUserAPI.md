# eZmaxAPI\ObjectUserAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserCreateObjectV1**](ObjectUserAPI.md#UserCreateObjectV1) | **Post** /1/object/user | Create a new User
[**UserEditObjectV1**](ObjectUserAPI.md#UserEditObjectV1) | **Put** /1/object/user/{pkiUserID} | Edit an existing User
[**UserEditPermissionsV1**](ObjectUserAPI.md#UserEditPermissionsV1) | **Put** /1/object/user/{pkiUserID}/editPermissions | Edit multiple Permissions
[**UserGetApikeysV1**](ObjectUserAPI.md#UserGetApikeysV1) | **Get** /1/object/user/{pkiUserID}/getApikeys | Retrieve an existing User&#39;s Apikeys
[**UserGetAutocompleteV2**](ObjectUserAPI.md#UserGetAutocompleteV2) | **Get** /2/object/user/getAutocomplete/{sSelector} | Retrieve Users and IDs
[**UserGetEffectivePermissionsV1**](ObjectUserAPI.md#UserGetEffectivePermissionsV1) | **Get** /1/object/user/{pkiUserID}/getEffectivePermissions | Retrieve an existing User&#39;s Effective Permissions
[**UserGetListV1**](ObjectUserAPI.md#UserGetListV1) | **Get** /1/object/user/getList | Retrieve User list
[**UserGetObjectV2**](ObjectUserAPI.md#UserGetObjectV2) | **Get** /2/object/user/{pkiUserID} | Retrieve an existing User
[**UserGetPermissionsV1**](ObjectUserAPI.md#UserGetPermissionsV1) | **Get** /1/object/user/{pkiUserID}/getPermissions | Retrieve an existing User&#39;s Permissions
[**UserGetSubnetsV1**](ObjectUserAPI.md#UserGetSubnetsV1) | **Get** /1/object/user/{pkiUserID}/getSubnets | Retrieve an existing User&#39;s Subnets
[**UserSendPasswordResetV1**](ObjectUserAPI.md#UserSendPasswordResetV1) | **Post** /1/object/user/{pkiUserID}/sendPasswordReset | Send password reset



## UserCreateObjectV1

> UserCreateObjectV1Response UserCreateObjectV1(ctx).UserCreateObjectV1Request(userCreateObjectV1Request).Execute()

Create a new User



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
    userCreateObjectV1Request := *openapiclient.NewUserCreateObjectV1Request([]openapiclient.UserRequestCompound{*openapiclient.NewUserRequestCompound(int32(1), int32(21), int32(247), int32(2), *openapiclient.NewEmailRequestCompound(int32(1), "email@example.com"), int32(1), openapiclient.Field-eUserType("AgentBroker"), openapiclient.Field-eUserLogintype("Password"), "John", "Doe", "JohnDoe", openapiclient.Field-eUserEzsignaccess("No"), true)}) // UserCreateObjectV1Request | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectUserAPI.UserCreateObjectV1(context.Background()).UserCreateObjectV1Request(userCreateObjectV1Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectUserAPI.UserCreateObjectV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserCreateObjectV1`: UserCreateObjectV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectUserAPI.UserCreateObjectV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserCreateObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userCreateObjectV1Request** | [**UserCreateObjectV1Request**](UserCreateObjectV1Request.md) |  | 

### Return type

[**UserCreateObjectV1Response**](UserCreateObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserEditObjectV1

> UserEditObjectV1Response UserEditObjectV1(ctx, pkiUserID).UserEditObjectV1Request(userEditObjectV1Request).Execute()

Edit an existing User



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
    pkiUserID := int32(56) // int32 | The unique ID of the User
    userEditObjectV1Request := *openapiclient.NewUserEditObjectV1Request(*openapiclient.NewUserRequestCompound(int32(1), int32(21), int32(247), int32(2), *openapiclient.NewEmailRequestCompound(int32(1), "email@example.com"), int32(1), openapiclient.Field-eUserType("AgentBroker"), openapiclient.Field-eUserLogintype("Password"), "John", "Doe", "JohnDoe", openapiclient.Field-eUserEzsignaccess("No"), true)) // UserEditObjectV1Request | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectUserAPI.UserEditObjectV1(context.Background(), pkiUserID).UserEditObjectV1Request(userEditObjectV1Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectUserAPI.UserEditObjectV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserEditObjectV1`: UserEditObjectV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectUserAPI.UserEditObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiUserID** | **int32** | The unique ID of the User | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserEditObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userEditObjectV1Request** | [**UserEditObjectV1Request**](UserEditObjectV1Request.md) |  | 

### Return type

[**UserEditObjectV1Response**](UserEditObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserEditPermissionsV1

> UserEditPermissionsV1Response UserEditPermissionsV1(ctx, pkiUserID).UserEditPermissionsV1Request(userEditPermissionsV1Request).Execute()

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
    pkiUserID := int32(56) // int32 | 
    userEditPermissionsV1Request := *openapiclient.NewUserEditPermissionsV1Request([]openapiclient.PermissionRequestCompound{*openapiclient.NewPermissionRequestCompound(int32(53))}) // UserEditPermissionsV1Request | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectUserAPI.UserEditPermissionsV1(context.Background(), pkiUserID).UserEditPermissionsV1Request(userEditPermissionsV1Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectUserAPI.UserEditPermissionsV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserEditPermissionsV1`: UserEditPermissionsV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectUserAPI.UserEditPermissionsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiUserID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserEditPermissionsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userEditPermissionsV1Request** | [**UserEditPermissionsV1Request**](UserEditPermissionsV1Request.md) |  | 

### Return type

[**UserEditPermissionsV1Response**](UserEditPermissionsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserGetApikeysV1

> UserGetApikeysV1Response UserGetApikeysV1(ctx, pkiUserID).Execute()

Retrieve an existing User's Apikeys

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
    pkiUserID := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectUserAPI.UserGetApikeysV1(context.Background(), pkiUserID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectUserAPI.UserGetApikeysV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserGetApikeysV1`: UserGetApikeysV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectUserAPI.UserGetApikeysV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiUserID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserGetApikeysV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserGetApikeysV1Response**](UserGetApikeysV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserGetAutocompleteV2

> UserGetAutocompleteV2Response UserGetAutocompleteV2(ctx, sSelector).EFilterActive(eFilterActive).SQuery(sQuery).AcceptLanguage(acceptLanguage).Execute()

Retrieve Users and IDs



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
    sSelector := "sSelector_example" // string | The type of Users to return
    eFilterActive := "eFilterActive_example" // string | Specify which results we want to display. (optional) (default to "Active")
    sQuery := "sQuery_example" // string | Allow to filter the returned results (optional)
    acceptLanguage := openapiclient.Header-Accept-Language("*") // HeaderAcceptLanguage |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectUserAPI.UserGetAutocompleteV2(context.Background(), sSelector).EFilterActive(eFilterActive).SQuery(sQuery).AcceptLanguage(acceptLanguage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectUserAPI.UserGetAutocompleteV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserGetAutocompleteV2`: UserGetAutocompleteV2Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectUserAPI.UserGetAutocompleteV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sSelector** | **string** | The type of Users to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserGetAutocompleteV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eFilterActive** | **string** | Specify which results we want to display. | [default to &quot;Active&quot;]
 **sQuery** | **string** | Allow to filter the returned results | 
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 

### Return type

[**UserGetAutocompleteV2Response**](UserGetAutocompleteV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserGetEffectivePermissionsV1

> UserGetEffectivePermissionsV1Response UserGetEffectivePermissionsV1(ctx, pkiUserID).Execute()

Retrieve an existing User's Effective Permissions



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
    pkiUserID := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectUserAPI.UserGetEffectivePermissionsV1(context.Background(), pkiUserID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectUserAPI.UserGetEffectivePermissionsV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserGetEffectivePermissionsV1`: UserGetEffectivePermissionsV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectUserAPI.UserGetEffectivePermissionsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiUserID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserGetEffectivePermissionsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserGetEffectivePermissionsV1Response**](UserGetEffectivePermissionsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserGetListV1

> UserGetListV1Response UserGetListV1(ctx).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()

Retrieve User list



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
    resp, r, err := apiClient.ObjectUserAPI.UserGetListV1(context.Background()).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectUserAPI.UserGetListV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserGetListV1`: UserGetListV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectUserAPI.UserGetListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserGetListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eOrderBy** | **string** | Specify how you want the results to be sorted | 
 **iRowMax** | **int32** |  | 
 **iRowOffset** | **int32** |  | [default to 0]
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 
 **sFilter** | **string** |  | 

### Return type

[**UserGetListV1Response**](UserGetListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserGetObjectV2

> UserGetObjectV2Response UserGetObjectV2(ctx, pkiUserID).Execute()

Retrieve an existing User



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
    pkiUserID := int32(56) // int32 | The unique ID of the User

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectUserAPI.UserGetObjectV2(context.Background(), pkiUserID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectUserAPI.UserGetObjectV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserGetObjectV2`: UserGetObjectV2Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectUserAPI.UserGetObjectV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiUserID** | **int32** | The unique ID of the User | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserGetObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserGetObjectV2Response**](UserGetObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserGetPermissionsV1

> UserGetPermissionsV1Response UserGetPermissionsV1(ctx, pkiUserID).Execute()

Retrieve an existing User's Permissions

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
    pkiUserID := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectUserAPI.UserGetPermissionsV1(context.Background(), pkiUserID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectUserAPI.UserGetPermissionsV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserGetPermissionsV1`: UserGetPermissionsV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectUserAPI.UserGetPermissionsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiUserID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserGetPermissionsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserGetPermissionsV1Response**](UserGetPermissionsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserGetSubnetsV1

> UserGetSubnetsV1Response UserGetSubnetsV1(ctx, pkiUserID).Execute()

Retrieve an existing User's Subnets

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
    pkiUserID := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectUserAPI.UserGetSubnetsV1(context.Background(), pkiUserID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectUserAPI.UserGetSubnetsV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserGetSubnetsV1`: UserGetSubnetsV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectUserAPI.UserGetSubnetsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiUserID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserGetSubnetsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserGetSubnetsV1Response**](UserGetSubnetsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserSendPasswordResetV1

> UserSendPasswordResetV1Response UserSendPasswordResetV1(ctx, pkiUserID).Body(body).Execute()

Send password reset



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
    pkiUserID := int32(56) // int32 | 
    body := map[string]interface{}{ ... } // map[string]interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectUserAPI.UserSendPasswordResetV1(context.Background(), pkiUserID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectUserAPI.UserSendPasswordResetV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserSendPasswordResetV1`: UserSendPasswordResetV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectUserAPI.UserSendPasswordResetV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiUserID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserSendPasswordResetV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

[**UserSendPasswordResetV1Response**](UserSendPasswordResetV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

