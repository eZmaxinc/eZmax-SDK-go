# eZmaxAPI\ObjectEzsignfolderAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EzsignfolderArchiveV1**](ObjectEzsignfolderAPI.md#EzsignfolderArchiveV1) | **Post** /1/object/ezsignfolder/{pkiEzsignfolderID}/archive | Archive the Ezsignfolder
[**EzsignfolderBatchDownloadV1**](ObjectEzsignfolderAPI.md#EzsignfolderBatchDownloadV1) | **Post** /1/object/ezsignfolder/{pkiEzsignfolderID}/batchDownload | Download multiples files from an Ezsignfolder
[**EzsignfolderCreateObjectV1**](ObjectEzsignfolderAPI.md#EzsignfolderCreateObjectV1) | **Post** /1/object/ezsignfolder | Create a new Ezsignfolder
[**EzsignfolderCreateObjectV2**](ObjectEzsignfolderAPI.md#EzsignfolderCreateObjectV2) | **Post** /2/object/ezsignfolder | Create a new Ezsignfolder
[**EzsignfolderDeleteObjectV1**](ObjectEzsignfolderAPI.md#EzsignfolderDeleteObjectV1) | **Delete** /1/object/ezsignfolder/{pkiEzsignfolderID} | Delete an existing Ezsignfolder
[**EzsignfolderDisposeEzsignfoldersV1**](ObjectEzsignfolderAPI.md#EzsignfolderDisposeEzsignfoldersV1) | **Post** /1/object/ezsignfolder/disposeEzsignfolders | Dispose Ezsignfolders
[**EzsignfolderDisposeV1**](ObjectEzsignfolderAPI.md#EzsignfolderDisposeV1) | **Post** /1/object/ezsignfolder/{pkiEzsignfolderID}/dispose | Dispose the Ezsignfolder
[**EzsignfolderEditObjectV1**](ObjectEzsignfolderAPI.md#EzsignfolderEditObjectV1) | **Put** /1/object/ezsignfolder/{pkiEzsignfolderID} | Edit an existing Ezsignfolder
[**EzsignfolderGetActionableElementsV1**](ObjectEzsignfolderAPI.md#EzsignfolderGetActionableElementsV1) | **Get** /1/object/ezsignfolder/{pkiEzsignfolderID}/getActionableElements | Retrieve actionable elements for the Ezsignfolder
[**EzsignfolderGetCommunicationCountV1**](ObjectEzsignfolderAPI.md#EzsignfolderGetCommunicationCountV1) | **Get** /1/object/ezsignfolder/{pkiEzsignfolderID}/getCommunicationCount | Retrieve Communication count
[**EzsignfolderGetCommunicationListV1**](ObjectEzsignfolderAPI.md#EzsignfolderGetCommunicationListV1) | **Get** /1/object/ezsignfolder/{pkiEzsignfolderID}/getCommunicationList | Retrieve Communication list
[**EzsignfolderGetEzsigndocumentsV1**](ObjectEzsignfolderAPI.md#EzsignfolderGetEzsigndocumentsV1) | **Get** /1/object/ezsignfolder/{pkiEzsignfolderID}/getEzsigndocuments | Retrieve an existing Ezsignfolder&#39;s Ezsigndocuments
[**EzsignfolderGetEzsignfoldersignerassociationsV1**](ObjectEzsignfolderAPI.md#EzsignfolderGetEzsignfoldersignerassociationsV1) | **Get** /1/object/ezsignfolder/{pkiEzsignfolderID}/getEzsignfoldersignerassociations | Retrieve an existing Ezsignfolder&#39;s Ezsignfoldersignerassociations
[**EzsignfolderGetEzsignsignaturesAutomaticV1**](ObjectEzsignfolderAPI.md#EzsignfolderGetEzsignsignaturesAutomaticV1) | **Get** /1/object/ezsignfolder/{pkiEzsignfolderID}/getEzsignsignaturesAutomatic | Retrieve an existing Ezsignfolder&#39;s automatic Ezsignsignatures
[**EzsignfolderGetFormsDataV1**](ObjectEzsignfolderAPI.md#EzsignfolderGetFormsDataV1) | **Get** /1/object/ezsignfolder/{pkiEzsignfolderID}/getFormsData | Retrieve an existing Ezsignfolder&#39;s forms data
[**EzsignfolderGetListV1**](ObjectEzsignfolderAPI.md#EzsignfolderGetListV1) | **Get** /1/object/ezsignfolder/getList | Retrieve Ezsignfolder list
[**EzsignfolderGetObjectV1**](ObjectEzsignfolderAPI.md#EzsignfolderGetObjectV1) | **Get** /1/object/ezsignfolder/{pkiEzsignfolderID} | Retrieve an existing Ezsignfolder
[**EzsignfolderGetObjectV2**](ObjectEzsignfolderAPI.md#EzsignfolderGetObjectV2) | **Get** /2/object/ezsignfolder/{pkiEzsignfolderID} | Retrieve an existing Ezsignfolder
[**EzsignfolderImportEzsignfoldersignerassociationsV1**](ObjectEzsignfolderAPI.md#EzsignfolderImportEzsignfoldersignerassociationsV1) | **Post** /1/object/ezsignfolder/{pkiEzsignfolderID}/importEzsignfoldersignerassociations | Import an existing Ezsignfoldersignerassociation into this Ezsignfolder
[**EzsignfolderImportEzsigntemplatepackageV1**](ObjectEzsignfolderAPI.md#EzsignfolderImportEzsigntemplatepackageV1) | **Post** /1/object/ezsignfolder/{pkiEzsignfolderID}/importEzsigntemplatepackage | Import an Ezsigntemplatepackage in the Ezsignfolder.
[**EzsignfolderReorderV1**](ObjectEzsignfolderAPI.md#EzsignfolderReorderV1) | **Post** /1/object/ezsignfolder/{pkiEzsignfolderID}/reorder | Reorder Ezsigndocuments in the Ezsignfolder
[**EzsignfolderSendV1**](ObjectEzsignfolderAPI.md#EzsignfolderSendV1) | **Post** /1/object/ezsignfolder/{pkiEzsignfolderID}/send | Send the Ezsignfolder to the signatories for signature
[**EzsignfolderSendV2**](ObjectEzsignfolderAPI.md#EzsignfolderSendV2) | **Post** /2/object/ezsignfolder/{pkiEzsignfolderID}/send | Send the Ezsignfolder to the signatories for signature
[**EzsignfolderSendV3**](ObjectEzsignfolderAPI.md#EzsignfolderSendV3) | **Post** /3/object/ezsignfolder/{pkiEzsignfolderID}/send | Send the Ezsignfolder to the signatories for signature
[**EzsignfolderUnsendV1**](ObjectEzsignfolderAPI.md#EzsignfolderUnsendV1) | **Post** /1/object/ezsignfolder/{pkiEzsignfolderID}/unsend | Unsend the Ezsignfolder



## EzsignfolderArchiveV1

> EzsignfolderArchiveV1Response EzsignfolderArchiveV1(ctx, pkiEzsignfolderID).Body(body).Execute()

Archive the Ezsignfolder



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
    pkiEzsignfolderID := int32(56) // int32 | 
    body := map[string]interface{}{ ... } // map[string]interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsignfolderAPI.EzsignfolderArchiveV1(context.Background(), pkiEzsignfolderID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignfolderAPI.EzsignfolderArchiveV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsignfolderArchiveV1`: EzsignfolderArchiveV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignfolderAPI.EzsignfolderArchiveV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignfolderID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignfolderArchiveV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

[**EzsignfolderArchiveV1Response**](EzsignfolderArchiveV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignfolderBatchDownloadV1

> *os.File EzsignfolderBatchDownloadV1(ctx, pkiEzsignfolderID).EzsignfolderBatchDownloadV1Request(ezsignfolderBatchDownloadV1Request).Execute()

Download multiples files from an Ezsignfolder

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
    pkiEzsignfolderID := int32(56) // int32 | 
    ezsignfolderBatchDownloadV1Request := *openapiclient.NewEzsignfolderBatchDownloadV1Request([]int32{int32(97)}, []string{"AEDocumentType_example"}) // EzsignfolderBatchDownloadV1Request | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsignfolderAPI.EzsignfolderBatchDownloadV1(context.Background(), pkiEzsignfolderID).EzsignfolderBatchDownloadV1Request(ezsignfolderBatchDownloadV1Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignfolderAPI.EzsignfolderBatchDownloadV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsignfolderBatchDownloadV1`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignfolderAPI.EzsignfolderBatchDownloadV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignfolderID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignfolderBatchDownloadV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezsignfolderBatchDownloadV1Request** | [**EzsignfolderBatchDownloadV1Request**](EzsignfolderBatchDownloadV1Request.md) |  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/zip, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignfolderCreateObjectV1

> EzsignfolderCreateObjectV1Response EzsignfolderCreateObjectV1(ctx).EzsignfolderCreateObjectV1Request(ezsignfolderCreateObjectV1Request).Execute()

Create a new Ezsignfolder



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
    ezsignfolderCreateObjectV1Request := []openapiclient.EzsignfolderCreateObjectV1Request{*openapiclient.NewEzsignfolderCreateObjectV1Request()} // []EzsignfolderCreateObjectV1Request | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsignfolderAPI.EzsignfolderCreateObjectV1(context.Background()).EzsignfolderCreateObjectV1Request(ezsignfolderCreateObjectV1Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignfolderAPI.EzsignfolderCreateObjectV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsignfolderCreateObjectV1`: EzsignfolderCreateObjectV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignfolderAPI.EzsignfolderCreateObjectV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEzsignfolderCreateObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ezsignfolderCreateObjectV1Request** | [**[]EzsignfolderCreateObjectV1Request**](EzsignfolderCreateObjectV1Request.md) |  | 

### Return type

[**EzsignfolderCreateObjectV1Response**](EzsignfolderCreateObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignfolderCreateObjectV2

> EzsignfolderCreateObjectV2Response EzsignfolderCreateObjectV2(ctx).EzsignfolderCreateObjectV2Request(ezsignfolderCreateObjectV2Request).Execute()

Create a new Ezsignfolder



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
    ezsignfolderCreateObjectV2Request := *openapiclient.NewEzsignfolderCreateObjectV2Request([]openapiclient.EzsignfolderRequestCompound{*openapiclient.NewEzsignfolderRequestCompound(int32(5), "Test eZsign Folder", openapiclient.Field-eEzsignfolderSendreminderfrequency("None"))}) // EzsignfolderCreateObjectV2Request | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsignfolderAPI.EzsignfolderCreateObjectV2(context.Background()).EzsignfolderCreateObjectV2Request(ezsignfolderCreateObjectV2Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignfolderAPI.EzsignfolderCreateObjectV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsignfolderCreateObjectV2`: EzsignfolderCreateObjectV2Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignfolderAPI.EzsignfolderCreateObjectV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEzsignfolderCreateObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ezsignfolderCreateObjectV2Request** | [**EzsignfolderCreateObjectV2Request**](EzsignfolderCreateObjectV2Request.md) |  | 

### Return type

[**EzsignfolderCreateObjectV2Response**](EzsignfolderCreateObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignfolderDeleteObjectV1

> EzsignfolderDeleteObjectV1Response EzsignfolderDeleteObjectV1(ctx, pkiEzsignfolderID).Execute()

Delete an existing Ezsignfolder

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
    pkiEzsignfolderID := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsignfolderAPI.EzsignfolderDeleteObjectV1(context.Background(), pkiEzsignfolderID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignfolderAPI.EzsignfolderDeleteObjectV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsignfolderDeleteObjectV1`: EzsignfolderDeleteObjectV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignfolderAPI.EzsignfolderDeleteObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignfolderID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignfolderDeleteObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsignfolderDeleteObjectV1Response**](EzsignfolderDeleteObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignfolderDisposeEzsignfoldersV1

> EzsignfolderDisposeEzsignfoldersV1Response EzsignfolderDisposeEzsignfoldersV1(ctx).EzsignfolderDisposeEzsignfoldersV1Request(ezsignfolderDisposeEzsignfoldersV1Request).Execute()

Dispose Ezsignfolders



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
    ezsignfolderDisposeEzsignfoldersV1Request := *openapiclient.NewEzsignfolderDisposeEzsignfoldersV1Request([]int32{int32(33)}) // EzsignfolderDisposeEzsignfoldersV1Request | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsignfolderAPI.EzsignfolderDisposeEzsignfoldersV1(context.Background()).EzsignfolderDisposeEzsignfoldersV1Request(ezsignfolderDisposeEzsignfoldersV1Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignfolderAPI.EzsignfolderDisposeEzsignfoldersV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsignfolderDisposeEzsignfoldersV1`: EzsignfolderDisposeEzsignfoldersV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignfolderAPI.EzsignfolderDisposeEzsignfoldersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEzsignfolderDisposeEzsignfoldersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ezsignfolderDisposeEzsignfoldersV1Request** | [**EzsignfolderDisposeEzsignfoldersV1Request**](EzsignfolderDisposeEzsignfoldersV1Request.md) |  | 

### Return type

[**EzsignfolderDisposeEzsignfoldersV1Response**](EzsignfolderDisposeEzsignfoldersV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignfolderDisposeV1

> EzsignfolderDisposeV1Response EzsignfolderDisposeV1(ctx, pkiEzsignfolderID).Body(body).Execute()

Dispose the Ezsignfolder



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
    pkiEzsignfolderID := int32(56) // int32 | 
    body := map[string]interface{}{ ... } // map[string]interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsignfolderAPI.EzsignfolderDisposeV1(context.Background(), pkiEzsignfolderID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignfolderAPI.EzsignfolderDisposeV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsignfolderDisposeV1`: EzsignfolderDisposeV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignfolderAPI.EzsignfolderDisposeV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignfolderID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignfolderDisposeV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

[**EzsignfolderDisposeV1Response**](EzsignfolderDisposeV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignfolderEditObjectV1

> EzsignfolderEditObjectV1Response EzsignfolderEditObjectV1(ctx, pkiEzsignfolderID).EzsignfolderEditObjectV1Request(ezsignfolderEditObjectV1Request).Execute()

Edit an existing Ezsignfolder



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
    pkiEzsignfolderID := int32(56) // int32 | 
    ezsignfolderEditObjectV1Request := *openapiclient.NewEzsignfolderEditObjectV1Request(*openapiclient.NewEzsignfolderRequestCompound(int32(5), "Test eZsign Folder", openapiclient.Field-eEzsignfolderSendreminderfrequency("None"))) // EzsignfolderEditObjectV1Request | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsignfolderAPI.EzsignfolderEditObjectV1(context.Background(), pkiEzsignfolderID).EzsignfolderEditObjectV1Request(ezsignfolderEditObjectV1Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignfolderAPI.EzsignfolderEditObjectV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsignfolderEditObjectV1`: EzsignfolderEditObjectV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignfolderAPI.EzsignfolderEditObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignfolderID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignfolderEditObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezsignfolderEditObjectV1Request** | [**EzsignfolderEditObjectV1Request**](EzsignfolderEditObjectV1Request.md) |  | 

### Return type

[**EzsignfolderEditObjectV1Response**](EzsignfolderEditObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignfolderGetActionableElementsV1

> EzsignfolderGetActionableElementsV1Response EzsignfolderGetActionableElementsV1(ctx, pkiEzsignfolderID).Execute()

Retrieve actionable elements for the Ezsignfolder



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
    pkiEzsignfolderID := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsignfolderAPI.EzsignfolderGetActionableElementsV1(context.Background(), pkiEzsignfolderID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignfolderAPI.EzsignfolderGetActionableElementsV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsignfolderGetActionableElementsV1`: EzsignfolderGetActionableElementsV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignfolderAPI.EzsignfolderGetActionableElementsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignfolderID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignfolderGetActionableElementsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsignfolderGetActionableElementsV1Response**](EzsignfolderGetActionableElementsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignfolderGetCommunicationCountV1

> EzsignfolderGetCommunicationCountV1Response EzsignfolderGetCommunicationCountV1(ctx, pkiEzsignfolderID).Execute()

Retrieve Communication count



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
    pkiEzsignfolderID := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsignfolderAPI.EzsignfolderGetCommunicationCountV1(context.Background(), pkiEzsignfolderID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignfolderAPI.EzsignfolderGetCommunicationCountV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsignfolderGetCommunicationCountV1`: EzsignfolderGetCommunicationCountV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignfolderAPI.EzsignfolderGetCommunicationCountV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignfolderID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignfolderGetCommunicationCountV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsignfolderGetCommunicationCountV1Response**](EzsignfolderGetCommunicationCountV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignfolderGetCommunicationListV1

> EzsignfolderGetCommunicationListV1Response EzsignfolderGetCommunicationListV1(ctx, pkiEzsignfolderID).Execute()

Retrieve Communication list



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
    pkiEzsignfolderID := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsignfolderAPI.EzsignfolderGetCommunicationListV1(context.Background(), pkiEzsignfolderID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignfolderAPI.EzsignfolderGetCommunicationListV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsignfolderGetCommunicationListV1`: EzsignfolderGetCommunicationListV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignfolderAPI.EzsignfolderGetCommunicationListV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignfolderID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignfolderGetCommunicationListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsignfolderGetCommunicationListV1Response**](EzsignfolderGetCommunicationListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignfolderGetEzsigndocumentsV1

> EzsignfolderGetEzsigndocumentsV1Response EzsignfolderGetEzsigndocumentsV1(ctx, pkiEzsignfolderID).Execute()

Retrieve an existing Ezsignfolder's Ezsigndocuments



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
    pkiEzsignfolderID := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsignfolderAPI.EzsignfolderGetEzsigndocumentsV1(context.Background(), pkiEzsignfolderID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignfolderAPI.EzsignfolderGetEzsigndocumentsV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsignfolderGetEzsigndocumentsV1`: EzsignfolderGetEzsigndocumentsV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignfolderAPI.EzsignfolderGetEzsigndocumentsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignfolderID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignfolderGetEzsigndocumentsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsignfolderGetEzsigndocumentsV1Response**](EzsignfolderGetEzsigndocumentsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignfolderGetEzsignfoldersignerassociationsV1

> EzsignfolderGetEzsignfoldersignerassociationsV1Response EzsignfolderGetEzsignfoldersignerassociationsV1(ctx, pkiEzsignfolderID).Execute()

Retrieve an existing Ezsignfolder's Ezsignfoldersignerassociations



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
    pkiEzsignfolderID := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsignfolderAPI.EzsignfolderGetEzsignfoldersignerassociationsV1(context.Background(), pkiEzsignfolderID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignfolderAPI.EzsignfolderGetEzsignfoldersignerassociationsV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsignfolderGetEzsignfoldersignerassociationsV1`: EzsignfolderGetEzsignfoldersignerassociationsV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignfolderAPI.EzsignfolderGetEzsignfoldersignerassociationsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignfolderID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignfolderGetEzsignfoldersignerassociationsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsignfolderGetEzsignfoldersignerassociationsV1Response**](EzsignfolderGetEzsignfoldersignerassociationsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignfolderGetEzsignsignaturesAutomaticV1

> EzsignfolderGetEzsignsignaturesAutomaticV1Response EzsignfolderGetEzsignsignaturesAutomaticV1(ctx, pkiEzsignfolderID).Execute()

Retrieve an existing Ezsignfolder's automatic Ezsignsignatures



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
    pkiEzsignfolderID := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsignfolderAPI.EzsignfolderGetEzsignsignaturesAutomaticV1(context.Background(), pkiEzsignfolderID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignfolderAPI.EzsignfolderGetEzsignsignaturesAutomaticV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsignfolderGetEzsignsignaturesAutomaticV1`: EzsignfolderGetEzsignsignaturesAutomaticV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignfolderAPI.EzsignfolderGetEzsignsignaturesAutomaticV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignfolderID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignfolderGetEzsignsignaturesAutomaticV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsignfolderGetEzsignsignaturesAutomaticV1Response**](EzsignfolderGetEzsignsignaturesAutomaticV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignfolderGetFormsDataV1

> EzsignfolderGetFormsDataV1Response EzsignfolderGetFormsDataV1(ctx, pkiEzsignfolderID).Execute()

Retrieve an existing Ezsignfolder's forms data



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
    pkiEzsignfolderID := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsignfolderAPI.EzsignfolderGetFormsDataV1(context.Background(), pkiEzsignfolderID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignfolderAPI.EzsignfolderGetFormsDataV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsignfolderGetFormsDataV1`: EzsignfolderGetFormsDataV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignfolderAPI.EzsignfolderGetFormsDataV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignfolderID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignfolderGetFormsDataV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsignfolderGetFormsDataV1Response**](EzsignfolderGetFormsDataV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/zip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignfolderGetListV1

> EzsignfolderGetListV1Response EzsignfolderGetListV1(ctx).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()

Retrieve Ezsignfolder list



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
    iRowMax := int32(56) // int32 |  (optional) (default to 10000)
    iRowOffset := int32(56) // int32 |  (optional) (default to 0)
    acceptLanguage := openapiclient.Header-Accept-Language("*") // HeaderAcceptLanguage |  (optional)
    sFilter := "sFilter_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsignfolderAPI.EzsignfolderGetListV1(context.Background()).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignfolderAPI.EzsignfolderGetListV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsignfolderGetListV1`: EzsignfolderGetListV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignfolderAPI.EzsignfolderGetListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEzsignfolderGetListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eOrderBy** | **string** | Specify how you want the results to be sorted | 
 **iRowMax** | **int32** |  | [default to 10000]
 **iRowOffset** | **int32** |  | [default to 0]
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 
 **sFilter** | **string** |  | 

### Return type

[**EzsignfolderGetListV1Response**](EzsignfolderGetListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignfolderGetObjectV1

> EzsignfolderGetObjectV1Response EzsignfolderGetObjectV1(ctx, pkiEzsignfolderID).Execute()

Retrieve an existing Ezsignfolder

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
    pkiEzsignfolderID := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsignfolderAPI.EzsignfolderGetObjectV1(context.Background(), pkiEzsignfolderID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignfolderAPI.EzsignfolderGetObjectV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsignfolderGetObjectV1`: EzsignfolderGetObjectV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignfolderAPI.EzsignfolderGetObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignfolderID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignfolderGetObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsignfolderGetObjectV1Response**](EzsignfolderGetObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignfolderGetObjectV2

> EzsignfolderGetObjectV2Response EzsignfolderGetObjectV2(ctx, pkiEzsignfolderID).Execute()

Retrieve an existing Ezsignfolder



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
    pkiEzsignfolderID := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsignfolderAPI.EzsignfolderGetObjectV2(context.Background(), pkiEzsignfolderID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignfolderAPI.EzsignfolderGetObjectV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsignfolderGetObjectV2`: EzsignfolderGetObjectV2Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignfolderAPI.EzsignfolderGetObjectV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignfolderID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignfolderGetObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsignfolderGetObjectV2Response**](EzsignfolderGetObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignfolderImportEzsignfoldersignerassociationsV1

> EzsignfolderImportEzsignfoldersignerassociationsV1Response EzsignfolderImportEzsignfoldersignerassociationsV1(ctx, pkiEzsignfolderID).EzsignfolderImportEzsignfoldersignerassociationsV1Request(ezsignfolderImportEzsignfoldersignerassociationsV1Request).Execute()

Import an existing Ezsignfoldersignerassociation into this Ezsignfolder



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
    pkiEzsignfolderID := int32(56) // int32 | 
    ezsignfolderImportEzsignfoldersignerassociationsV1Request := *openapiclient.NewEzsignfolderImportEzsignfoldersignerassociationsV1Request([]int32{int32(20)}) // EzsignfolderImportEzsignfoldersignerassociationsV1Request | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsignfolderAPI.EzsignfolderImportEzsignfoldersignerassociationsV1(context.Background(), pkiEzsignfolderID).EzsignfolderImportEzsignfoldersignerassociationsV1Request(ezsignfolderImportEzsignfoldersignerassociationsV1Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignfolderAPI.EzsignfolderImportEzsignfoldersignerassociationsV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsignfolderImportEzsignfoldersignerassociationsV1`: EzsignfolderImportEzsignfoldersignerassociationsV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignfolderAPI.EzsignfolderImportEzsignfoldersignerassociationsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignfolderID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignfolderImportEzsignfoldersignerassociationsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezsignfolderImportEzsignfoldersignerassociationsV1Request** | [**EzsignfolderImportEzsignfoldersignerassociationsV1Request**](EzsignfolderImportEzsignfoldersignerassociationsV1Request.md) |  | 

### Return type

[**EzsignfolderImportEzsignfoldersignerassociationsV1Response**](EzsignfolderImportEzsignfoldersignerassociationsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignfolderImportEzsigntemplatepackageV1

> EzsignfolderImportEzsigntemplatepackageV1Response EzsignfolderImportEzsigntemplatepackageV1(ctx, pkiEzsignfolderID).EzsignfolderImportEzsigntemplatepackageV1Request(ezsignfolderImportEzsigntemplatepackageV1Request).Execute()

Import an Ezsigntemplatepackage in the Ezsignfolder.



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
    pkiEzsignfolderID := int32(56) // int32 | 
    ezsignfolderImportEzsigntemplatepackageV1Request := *openapiclient.NewEzsignfolderImportEzsigntemplatepackageV1Request(int32(99), "2020-12-31 23:59:59", []openapiclient.CustomImportEzsigntemplatepackageRelationRequest{*openapiclient.NewCustomImportEzsigntemplatepackageRelationRequest(int32(20))}) // EzsignfolderImportEzsigntemplatepackageV1Request | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsignfolderAPI.EzsignfolderImportEzsigntemplatepackageV1(context.Background(), pkiEzsignfolderID).EzsignfolderImportEzsigntemplatepackageV1Request(ezsignfolderImportEzsigntemplatepackageV1Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignfolderAPI.EzsignfolderImportEzsigntemplatepackageV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsignfolderImportEzsigntemplatepackageV1`: EzsignfolderImportEzsigntemplatepackageV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignfolderAPI.EzsignfolderImportEzsigntemplatepackageV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignfolderID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignfolderImportEzsigntemplatepackageV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezsignfolderImportEzsigntemplatepackageV1Request** | [**EzsignfolderImportEzsigntemplatepackageV1Request**](EzsignfolderImportEzsigntemplatepackageV1Request.md) |  | 

### Return type

[**EzsignfolderImportEzsigntemplatepackageV1Response**](EzsignfolderImportEzsigntemplatepackageV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignfolderReorderV1

> EzsignfolderReorderV1Response EzsignfolderReorderV1(ctx, pkiEzsignfolderID).EzsignfolderReorderV1Request(ezsignfolderReorderV1Request).Execute()

Reorder Ezsigndocuments in the Ezsignfolder

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
    pkiEzsignfolderID := int32(56) // int32 | 
    ezsignfolderReorderV1Request := *openapiclient.NewEzsignfolderReorderV1Request([]int32{int32(97)}) // EzsignfolderReorderV1Request | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsignfolderAPI.EzsignfolderReorderV1(context.Background(), pkiEzsignfolderID).EzsignfolderReorderV1Request(ezsignfolderReorderV1Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignfolderAPI.EzsignfolderReorderV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsignfolderReorderV1`: EzsignfolderReorderV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignfolderAPI.EzsignfolderReorderV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignfolderID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignfolderReorderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezsignfolderReorderV1Request** | [**EzsignfolderReorderV1Request**](EzsignfolderReorderV1Request.md) |  | 

### Return type

[**EzsignfolderReorderV1Response**](EzsignfolderReorderV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignfolderSendV1

> EzsignfolderSendV1Response EzsignfolderSendV1(ctx, pkiEzsignfolderID).EzsignfolderSendV1Request(ezsignfolderSendV1Request).Execute()

Send the Ezsignfolder to the signatories for signature



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
    pkiEzsignfolderID := int32(56) // int32 | 
    ezsignfolderSendV1Request := *openapiclient.NewEzsignfolderSendV1Request("Hi John,

This is the document I need you to review.

Could you sign it before Monday please.

Best Regards.

Mary") // EzsignfolderSendV1Request | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsignfolderAPI.EzsignfolderSendV1(context.Background(), pkiEzsignfolderID).EzsignfolderSendV1Request(ezsignfolderSendV1Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignfolderAPI.EzsignfolderSendV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsignfolderSendV1`: EzsignfolderSendV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignfolderAPI.EzsignfolderSendV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignfolderID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignfolderSendV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezsignfolderSendV1Request** | [**EzsignfolderSendV1Request**](EzsignfolderSendV1Request.md) |  | 

### Return type

[**EzsignfolderSendV1Response**](EzsignfolderSendV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignfolderSendV2

> EzsignfolderSendV2Response EzsignfolderSendV2(ctx, pkiEzsignfolderID).EzsignfolderSendV2Request(ezsignfolderSendV2Request).Execute()

Send the Ezsignfolder to the signatories for signature



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
    pkiEzsignfolderID := int32(56) // int32 | 
    ezsignfolderSendV2Request := *openapiclient.NewEzsignfolderSendV2Request("Hi everyone,

This is the document I need you to review.

Could you sign it before Monday please.

Best Regards.

Mary", []int32{int32(20)}, []openapiclient.CustomEzsignfoldersignerassociationmessageRequest{*openapiclient.NewCustomEzsignfoldersignerassociationmessageRequest(int32(20))}) // EzsignfolderSendV2Request | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsignfolderAPI.EzsignfolderSendV2(context.Background(), pkiEzsignfolderID).EzsignfolderSendV2Request(ezsignfolderSendV2Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignfolderAPI.EzsignfolderSendV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsignfolderSendV2`: EzsignfolderSendV2Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignfolderAPI.EzsignfolderSendV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignfolderID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignfolderSendV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezsignfolderSendV2Request** | [**EzsignfolderSendV2Request**](EzsignfolderSendV2Request.md) |  | 

### Return type

[**EzsignfolderSendV2Response**](EzsignfolderSendV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignfolderSendV3

> EzsignfolderSendV3Response EzsignfolderSendV3(ctx, pkiEzsignfolderID).EzsignfolderSendV3Request(ezsignfolderSendV3Request).Execute()

Send the Ezsignfolder to the signatories for signature



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
    pkiEzsignfolderID := int32(56) // int32 | 
    ezsignfolderSendV3Request := *openapiclient.NewEzsignfolderSendV3Request([]int32{int32(20)}) // EzsignfolderSendV3Request | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsignfolderAPI.EzsignfolderSendV3(context.Background(), pkiEzsignfolderID).EzsignfolderSendV3Request(ezsignfolderSendV3Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignfolderAPI.EzsignfolderSendV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsignfolderSendV3`: EzsignfolderSendV3Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignfolderAPI.EzsignfolderSendV3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignfolderID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignfolderSendV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezsignfolderSendV3Request** | [**EzsignfolderSendV3Request**](EzsignfolderSendV3Request.md) |  | 

### Return type

[**EzsignfolderSendV3Response**](EzsignfolderSendV3Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignfolderUnsendV1

> EzsignfolderUnsendV1Response EzsignfolderUnsendV1(ctx, pkiEzsignfolderID).Body(body).Execute()

Unsend the Ezsignfolder



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
    pkiEzsignfolderID := int32(56) // int32 | 
    body := map[string]interface{}{ ... } // map[string]interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsignfolderAPI.EzsignfolderUnsendV1(context.Background(), pkiEzsignfolderID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignfolderAPI.EzsignfolderUnsendV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsignfolderUnsendV1`: EzsignfolderUnsendV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignfolderAPI.EzsignfolderUnsendV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignfolderID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignfolderUnsendV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

[**EzsignfolderUnsendV1Response**](EzsignfolderUnsendV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

