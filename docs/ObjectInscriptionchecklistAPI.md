# eZmaxAPI\ObjectInscriptionchecklistAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InscriptionchecklistGetAutocompleteV2**](ObjectInscriptionchecklistAPI.md#InscriptionchecklistGetAutocompleteV2) | **Get** /2/object/inscriptionchecklist/getAutocomplete/{sSelector} | Retrieve Inscriptionchecklists and IDs
[**InscriptionchecklistGetAutocompleteV3**](ObjectInscriptionchecklistAPI.md#InscriptionchecklistGetAutocompleteV3) | **Get** /3/object/inscriptionchecklist/getAutocomplete/{sSelector} | Retrieve Inscriptionchecklists and IDs



## InscriptionchecklistGetAutocompleteV2

> InscriptionchecklistGetAutocompleteV2Response InscriptionchecklistGetAutocompleteV2(ctx, sSelector).FkiID(fkiID).EType(eType).EFilterActive(eFilterActive).SQuery(sQuery).AcceptLanguage(acceptLanguage).Execute()

Retrieve Inscriptionchecklists and IDs



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
	sSelector := "sSelector_example" // string | The type of Inscriptionchecklist to return
	fkiID := "fkiID_example" // string | Specify which fkiID we want to display. (optional)
	eType := "eType_example" // string | The type of Inscriptionchecklist (optional)
	eFilterActive := "eFilterActive_example" // string | Specify which results we want to display. (optional) (default to "Active")
	sQuery := "sQuery_example" // string | Allow to filter the returned results (optional)
	acceptLanguage := openapiclient.Header-Accept-Language("*") // HeaderAcceptLanguage |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectInscriptionchecklistAPI.InscriptionchecklistGetAutocompleteV2(context.Background(), sSelector).FkiID(fkiID).EType(eType).EFilterActive(eFilterActive).SQuery(sQuery).AcceptLanguage(acceptLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectInscriptionchecklistAPI.InscriptionchecklistGetAutocompleteV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InscriptionchecklistGetAutocompleteV2`: InscriptionchecklistGetAutocompleteV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectInscriptionchecklistAPI.InscriptionchecklistGetAutocompleteV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sSelector** | **string** | The type of Inscriptionchecklist to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiInscriptionchecklistGetAutocompleteV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fkiID** | **string** | Specify which fkiID we want to display. | 
 **eType** | **string** | The type of Inscriptionchecklist | 
 **eFilterActive** | **string** | Specify which results we want to display. | [default to &quot;Active&quot;]
 **sQuery** | **string** | Allow to filter the returned results | 
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 

### Return type

[**InscriptionchecklistGetAutocompleteV2Response**](InscriptionchecklistGetAutocompleteV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InscriptionchecklistGetAutocompleteV3

> InscriptionchecklistGetAutocompleteV3Response InscriptionchecklistGetAutocompleteV3(ctx, sSelector).FkiBuyercontractID(fkiBuyercontractID).FkiInscriptionID(fkiInscriptionID).FkiInscriptionnotauthenticatedID(fkiInscriptionnotauthenticatedID).FkiInscriptiontempID(fkiInscriptiontempID).FkiAgentID(fkiAgentID).FkiBrokerID(fkiBrokerID).FkiOtherincomeID(fkiOtherincomeID).FkiRejectedoffertopurchaseID(fkiRejectedoffertopurchaseID).EFilterActive(eFilterActive).SQuery(sQuery).AcceptLanguage(acceptLanguage).Execute()

Retrieve Inscriptionchecklists and IDs



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
	sSelector := "sSelector_example" // string | The type of Inscriptionchecklist to return
	fkiBuyercontractID := "fkiBuyercontractID_example" // string | Specify which Buyercontract we want to display. (optional)
	fkiInscriptionID := "fkiInscriptionID_example" // string | Specify which Inscription we want to display. (optional)
	fkiInscriptionnotauthenticatedID := "fkiInscriptionnotauthenticatedID_example" // string | Specify which Inscriptionnotauthenticated we want to display. (optional)
	fkiInscriptiontempID := "fkiInscriptiontempID_example" // string | Specify which Inscriptiontemp we want to display. (optional)
	fkiAgentID := "fkiAgentID_example" // string | Specify which Agent we want to display. (optional)
	fkiBrokerID := "fkiBrokerID_example" // string | Specify which Broker we want to display. (optional)
	fkiOtherincomeID := "fkiOtherincomeID_example" // string | Specify which Otherincome we want to display. (optional)
	fkiRejectedoffertopurchaseID := "fkiRejectedoffertopurchaseID_example" // string | Specify which Rejectedoffertopurchase we want to display. (optional)
	eFilterActive := "eFilterActive_example" // string | Specify which results we want to display. (optional) (default to "Active")
	sQuery := "sQuery_example" // string | Allow to filter the returned results (optional)
	acceptLanguage := openapiclient.Header-Accept-Language("*") // HeaderAcceptLanguage |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectInscriptionchecklistAPI.InscriptionchecklistGetAutocompleteV3(context.Background(), sSelector).FkiBuyercontractID(fkiBuyercontractID).FkiInscriptionID(fkiInscriptionID).FkiInscriptionnotauthenticatedID(fkiInscriptionnotauthenticatedID).FkiInscriptiontempID(fkiInscriptiontempID).FkiAgentID(fkiAgentID).FkiBrokerID(fkiBrokerID).FkiOtherincomeID(fkiOtherincomeID).FkiRejectedoffertopurchaseID(fkiRejectedoffertopurchaseID).EFilterActive(eFilterActive).SQuery(sQuery).AcceptLanguage(acceptLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectInscriptionchecklistAPI.InscriptionchecklistGetAutocompleteV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InscriptionchecklistGetAutocompleteV3`: InscriptionchecklistGetAutocompleteV3Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectInscriptionchecklistAPI.InscriptionchecklistGetAutocompleteV3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sSelector** | **string** | The type of Inscriptionchecklist to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiInscriptionchecklistGetAutocompleteV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fkiBuyercontractID** | **string** | Specify which Buyercontract we want to display. | 
 **fkiInscriptionID** | **string** | Specify which Inscription we want to display. | 
 **fkiInscriptionnotauthenticatedID** | **string** | Specify which Inscriptionnotauthenticated we want to display. | 
 **fkiInscriptiontempID** | **string** | Specify which Inscriptiontemp we want to display. | 
 **fkiAgentID** | **string** | Specify which Agent we want to display. | 
 **fkiBrokerID** | **string** | Specify which Broker we want to display. | 
 **fkiOtherincomeID** | **string** | Specify which Otherincome we want to display. | 
 **fkiRejectedoffertopurchaseID** | **string** | Specify which Rejectedoffertopurchase we want to display. | 
 **eFilterActive** | **string** | Specify which results we want to display. | [default to &quot;Active&quot;]
 **sQuery** | **string** | Allow to filter the returned results | 
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 

### Return type

[**InscriptionchecklistGetAutocompleteV3Response**](InscriptionchecklistGetAutocompleteV3Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

