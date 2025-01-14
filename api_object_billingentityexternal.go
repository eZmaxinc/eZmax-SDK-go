/*
eZmax API Definition (Full)

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.2.1
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// ObjectBillingentityexternalAPIService ObjectBillingentityexternalAPI service
type ObjectBillingentityexternalAPIService service

type ApiBillingentityexternalGenerateFederationTokenV1Request struct {
	ctx context.Context
	ApiService *ObjectBillingentityexternalAPIService
	pkiBillingentityexternalID int32
	billingentityexternalGenerateFederationTokenV1Request *BillingentityexternalGenerateFederationTokenV1Request
}

func (r ApiBillingentityexternalGenerateFederationTokenV1Request) BillingentityexternalGenerateFederationTokenV1Request(billingentityexternalGenerateFederationTokenV1Request BillingentityexternalGenerateFederationTokenV1Request) ApiBillingentityexternalGenerateFederationTokenV1Request {
	r.billingentityexternalGenerateFederationTokenV1Request = &billingentityexternalGenerateFederationTokenV1Request
	return r
}

func (r ApiBillingentityexternalGenerateFederationTokenV1Request) Execute() (*BillingentityexternalGenerateFederationTokenV1Response, *http.Response, error) {
	return r.ApiService.BillingentityexternalGenerateFederationTokenV1Execute(r)
}

/*
BillingentityexternalGenerateFederationTokenV1 Generate a federation token



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pkiBillingentityexternalID
 @return ApiBillingentityexternalGenerateFederationTokenV1Request
*/
func (a *ObjectBillingentityexternalAPIService) BillingentityexternalGenerateFederationTokenV1(ctx context.Context, pkiBillingentityexternalID int32) ApiBillingentityexternalGenerateFederationTokenV1Request {
	return ApiBillingentityexternalGenerateFederationTokenV1Request{
		ApiService: a,
		ctx: ctx,
		pkiBillingentityexternalID: pkiBillingentityexternalID,
	}
}

// Execute executes the request
//  @return BillingentityexternalGenerateFederationTokenV1Response
func (a *ObjectBillingentityexternalAPIService) BillingentityexternalGenerateFederationTokenV1Execute(r ApiBillingentityexternalGenerateFederationTokenV1Request) (*BillingentityexternalGenerateFederationTokenV1Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BillingentityexternalGenerateFederationTokenV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectBillingentityexternalAPIService.BillingentityexternalGenerateFederationTokenV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/billingentityexternal/{pkiBillingentityexternalID}/generateFederationToken"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiBillingentityexternalID"+"}", url.PathEscape(parameterValueToString(r.pkiBillingentityexternalID, "pkiBillingentityexternalID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pkiBillingentityexternalID < 1 {
		return localVarReturnValue, nil, reportError("pkiBillingentityexternalID must be greater than 1")
	}
	if r.billingentityexternalGenerateFederationTokenV1Request == nil {
		return localVarReturnValue, nil, reportError("billingentityexternalGenerateFederationTokenV1Request is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.billingentityexternalGenerateFederationTokenV1Request
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Authorization"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v CommonResponseError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiBillingentityexternalGetAutocompleteV2Request struct {
	ctx context.Context
	ApiService *ObjectBillingentityexternalAPIService
	sSelector string
	eFilterActive *string
	sQuery *string
	acceptLanguage *HeaderAcceptLanguage
}

// Specify which results we want to display.
func (r ApiBillingentityexternalGetAutocompleteV2Request) EFilterActive(eFilterActive string) ApiBillingentityexternalGetAutocompleteV2Request {
	r.eFilterActive = &eFilterActive
	return r
}

// Allow to filter the returned results
func (r ApiBillingentityexternalGetAutocompleteV2Request) SQuery(sQuery string) ApiBillingentityexternalGetAutocompleteV2Request {
	r.sQuery = &sQuery
	return r
}

func (r ApiBillingentityexternalGetAutocompleteV2Request) AcceptLanguage(acceptLanguage HeaderAcceptLanguage) ApiBillingentityexternalGetAutocompleteV2Request {
	r.acceptLanguage = &acceptLanguage
	return r
}

func (r ApiBillingentityexternalGetAutocompleteV2Request) Execute() (*BillingentityexternalGetAutocompleteV2Response, *http.Response, error) {
	return r.ApiService.BillingentityexternalGetAutocompleteV2Execute(r)
}

/*
BillingentityexternalGetAutocompleteV2 Retrieve Billingentityexternals and IDs

Get the list of Billingentityexternal to be used in a dropdown or autocomplete control.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param sSelector The type of Billingentityexternals to return
 @return ApiBillingentityexternalGetAutocompleteV2Request
*/
func (a *ObjectBillingentityexternalAPIService) BillingentityexternalGetAutocompleteV2(ctx context.Context, sSelector string) ApiBillingentityexternalGetAutocompleteV2Request {
	return ApiBillingentityexternalGetAutocompleteV2Request{
		ApiService: a,
		ctx: ctx,
		sSelector: sSelector,
	}
}

// Execute executes the request
//  @return BillingentityexternalGetAutocompleteV2Response
func (a *ObjectBillingentityexternalAPIService) BillingentityexternalGetAutocompleteV2Execute(r ApiBillingentityexternalGetAutocompleteV2Request) (*BillingentityexternalGetAutocompleteV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BillingentityexternalGetAutocompleteV2Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectBillingentityexternalAPIService.BillingentityexternalGetAutocompleteV2")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/2/object/billingentityexternal/getAutocomplete/{sSelector}"
	localVarPath = strings.Replace(localVarPath, "{"+"sSelector"+"}", url.PathEscape(parameterValueToString(r.sSelector, "sSelector")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.eFilterActive != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "eFilterActive", r.eFilterActive, "form", "")
	} else {
		var defaultValue string = "Active"
		r.eFilterActive = &defaultValue
	}
	if r.sQuery != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sQuery", r.sQuery, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.acceptLanguage != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Accept-Language", r.acceptLanguage, "simple", "")
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Authorization"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
