/*
eZmax API Definition (Full)

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.2.0
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


// ObjectBillingentityinternalAPIService ObjectBillingentityinternalAPI service
type ObjectBillingentityinternalAPIService service

type ApiBillingentityinternalCreateObjectV1Request struct {
	ctx context.Context
	ApiService *ObjectBillingentityinternalAPIService
	billingentityinternalCreateObjectV1Request *BillingentityinternalCreateObjectV1Request
}

func (r ApiBillingentityinternalCreateObjectV1Request) BillingentityinternalCreateObjectV1Request(billingentityinternalCreateObjectV1Request BillingentityinternalCreateObjectV1Request) ApiBillingentityinternalCreateObjectV1Request {
	r.billingentityinternalCreateObjectV1Request = &billingentityinternalCreateObjectV1Request
	return r
}

func (r ApiBillingentityinternalCreateObjectV1Request) Execute() (*BillingentityinternalCreateObjectV1Response, *http.Response, error) {
	return r.ApiService.BillingentityinternalCreateObjectV1Execute(r)
}

/*
BillingentityinternalCreateObjectV1 Create a new Billingentityinternal

The endpoint allows to create one or many elements at once.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiBillingentityinternalCreateObjectV1Request
*/
func (a *ObjectBillingentityinternalAPIService) BillingentityinternalCreateObjectV1(ctx context.Context) ApiBillingentityinternalCreateObjectV1Request {
	return ApiBillingentityinternalCreateObjectV1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BillingentityinternalCreateObjectV1Response
func (a *ObjectBillingentityinternalAPIService) BillingentityinternalCreateObjectV1Execute(r ApiBillingentityinternalCreateObjectV1Request) (*BillingentityinternalCreateObjectV1Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BillingentityinternalCreateObjectV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectBillingentityinternalAPIService.BillingentityinternalCreateObjectV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/billingentityinternal"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.billingentityinternalCreateObjectV1Request == nil {
		return localVarReturnValue, nil, reportError("billingentityinternalCreateObjectV1Request is required and must be specified")
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
	localVarPostBody = r.billingentityinternalCreateObjectV1Request
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

type ApiBillingentityinternalEditObjectV1Request struct {
	ctx context.Context
	ApiService *ObjectBillingentityinternalAPIService
	pkiBillingentityinternalID int32
	billingentityinternalEditObjectV1Request *BillingentityinternalEditObjectV1Request
}

func (r ApiBillingentityinternalEditObjectV1Request) BillingentityinternalEditObjectV1Request(billingentityinternalEditObjectV1Request BillingentityinternalEditObjectV1Request) ApiBillingentityinternalEditObjectV1Request {
	r.billingentityinternalEditObjectV1Request = &billingentityinternalEditObjectV1Request
	return r
}

func (r ApiBillingentityinternalEditObjectV1Request) Execute() (*BillingentityinternalEditObjectV1Response, *http.Response, error) {
	return r.ApiService.BillingentityinternalEditObjectV1Execute(r)
}

/*
BillingentityinternalEditObjectV1 Edit an existing Billingentityinternal



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pkiBillingentityinternalID
 @return ApiBillingentityinternalEditObjectV1Request
*/
func (a *ObjectBillingentityinternalAPIService) BillingentityinternalEditObjectV1(ctx context.Context, pkiBillingentityinternalID int32) ApiBillingentityinternalEditObjectV1Request {
	return ApiBillingentityinternalEditObjectV1Request{
		ApiService: a,
		ctx: ctx,
		pkiBillingentityinternalID: pkiBillingentityinternalID,
	}
}

// Execute executes the request
//  @return BillingentityinternalEditObjectV1Response
func (a *ObjectBillingentityinternalAPIService) BillingentityinternalEditObjectV1Execute(r ApiBillingentityinternalEditObjectV1Request) (*BillingentityinternalEditObjectV1Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BillingentityinternalEditObjectV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectBillingentityinternalAPIService.BillingentityinternalEditObjectV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/billingentityinternal/{pkiBillingentityinternalID}"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiBillingentityinternalID"+"}", url.PathEscape(parameterValueToString(r.pkiBillingentityinternalID, "pkiBillingentityinternalID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pkiBillingentityinternalID < 0 {
		return localVarReturnValue, nil, reportError("pkiBillingentityinternalID must be greater than 0")
	}
	if r.billingentityinternalEditObjectV1Request == nil {
		return localVarReturnValue, nil, reportError("billingentityinternalEditObjectV1Request is required and must be specified")
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
	localVarPostBody = r.billingentityinternalEditObjectV1Request
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
		if localVarHTTPResponse.StatusCode == 404 {
			var v CommonResponseError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
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

type ApiBillingentityinternalGetAutocompleteV2Request struct {
	ctx context.Context
	ApiService *ObjectBillingentityinternalAPIService
	sSelector string
	eFilterActive *string
	sQuery *string
	acceptLanguage *HeaderAcceptLanguage
}

// Specify which results we want to display.
func (r ApiBillingentityinternalGetAutocompleteV2Request) EFilterActive(eFilterActive string) ApiBillingentityinternalGetAutocompleteV2Request {
	r.eFilterActive = &eFilterActive
	return r
}

// Allow to filter the returned results
func (r ApiBillingentityinternalGetAutocompleteV2Request) SQuery(sQuery string) ApiBillingentityinternalGetAutocompleteV2Request {
	r.sQuery = &sQuery
	return r
}

func (r ApiBillingentityinternalGetAutocompleteV2Request) AcceptLanguage(acceptLanguage HeaderAcceptLanguage) ApiBillingentityinternalGetAutocompleteV2Request {
	r.acceptLanguage = &acceptLanguage
	return r
}

func (r ApiBillingentityinternalGetAutocompleteV2Request) Execute() (*BillingentityinternalGetAutocompleteV2Response, *http.Response, error) {
	return r.ApiService.BillingentityinternalGetAutocompleteV2Execute(r)
}

/*
BillingentityinternalGetAutocompleteV2 Retrieve Billingentityinternals and IDs

Get the list of Billingentityinternal to be used in a dropdown or autocomplete control.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param sSelector The type of Billingentityinternals to return
 @return ApiBillingentityinternalGetAutocompleteV2Request
*/
func (a *ObjectBillingentityinternalAPIService) BillingentityinternalGetAutocompleteV2(ctx context.Context, sSelector string) ApiBillingentityinternalGetAutocompleteV2Request {
	return ApiBillingentityinternalGetAutocompleteV2Request{
		ApiService: a,
		ctx: ctx,
		sSelector: sSelector,
	}
}

// Execute executes the request
//  @return BillingentityinternalGetAutocompleteV2Response
func (a *ObjectBillingentityinternalAPIService) BillingentityinternalGetAutocompleteV2Execute(r ApiBillingentityinternalGetAutocompleteV2Request) (*BillingentityinternalGetAutocompleteV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BillingentityinternalGetAutocompleteV2Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectBillingentityinternalAPIService.BillingentityinternalGetAutocompleteV2")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/2/object/billingentityinternal/getAutocomplete/{sSelector}"
	localVarPath = strings.Replace(localVarPath, "{"+"sSelector"+"}", url.PathEscape(parameterValueToString(r.sSelector, "sSelector")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.eFilterActive != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "eFilterActive", r.eFilterActive, "")
	} else {
		var defaultValue string = "Active"
		r.eFilterActive = &defaultValue
	}
	if r.sQuery != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sQuery", r.sQuery, "")
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
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Accept-Language", r.acceptLanguage, "")
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

type ApiBillingentityinternalGetListV1Request struct {
	ctx context.Context
	ApiService *ObjectBillingentityinternalAPIService
	eOrderBy *string
	iRowMax *int32
	iRowOffset *int32
	acceptLanguage *HeaderAcceptLanguage
	sFilter *string
}

// Specify how you want the results to be sorted
func (r ApiBillingentityinternalGetListV1Request) EOrderBy(eOrderBy string) ApiBillingentityinternalGetListV1Request {
	r.eOrderBy = &eOrderBy
	return r
}

func (r ApiBillingentityinternalGetListV1Request) IRowMax(iRowMax int32) ApiBillingentityinternalGetListV1Request {
	r.iRowMax = &iRowMax
	return r
}

func (r ApiBillingentityinternalGetListV1Request) IRowOffset(iRowOffset int32) ApiBillingentityinternalGetListV1Request {
	r.iRowOffset = &iRowOffset
	return r
}

func (r ApiBillingentityinternalGetListV1Request) AcceptLanguage(acceptLanguage HeaderAcceptLanguage) ApiBillingentityinternalGetListV1Request {
	r.acceptLanguage = &acceptLanguage
	return r
}

func (r ApiBillingentityinternalGetListV1Request) SFilter(sFilter string) ApiBillingentityinternalGetListV1Request {
	r.sFilter = &sFilter
	return r
}

func (r ApiBillingentityinternalGetListV1Request) Execute() (*BillingentityinternalGetListV1Response, *http.Response, error) {
	return r.ApiService.BillingentityinternalGetListV1Execute(r)
}

/*
BillingentityinternalGetListV1 Retrieve Billingentityinternal list



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiBillingentityinternalGetListV1Request
*/
func (a *ObjectBillingentityinternalAPIService) BillingentityinternalGetListV1(ctx context.Context) ApiBillingentityinternalGetListV1Request {
	return ApiBillingentityinternalGetListV1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BillingentityinternalGetListV1Response
func (a *ObjectBillingentityinternalAPIService) BillingentityinternalGetListV1Execute(r ApiBillingentityinternalGetListV1Request) (*BillingentityinternalGetListV1Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BillingentityinternalGetListV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectBillingentityinternalAPIService.BillingentityinternalGetListV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/billingentityinternal/getList"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.eOrderBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "eOrderBy", r.eOrderBy, "")
	}
	if r.iRowMax != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "iRowMax", r.iRowMax, "")
	}
	if r.iRowOffset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "iRowOffset", r.iRowOffset, "")
	} else {
		var defaultValue int32 = 0
		r.iRowOffset = &defaultValue
	}
	if r.sFilter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sFilter", r.sFilter, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.acceptLanguage != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Accept-Language", r.acceptLanguage, "")
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
		if localVarHTTPResponse.StatusCode == 406 {
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

type ApiBillingentityinternalGetObjectV2Request struct {
	ctx context.Context
	ApiService *ObjectBillingentityinternalAPIService
	pkiBillingentityinternalID int32
}

func (r ApiBillingentityinternalGetObjectV2Request) Execute() (*BillingentityinternalGetObjectV2Response, *http.Response, error) {
	return r.ApiService.BillingentityinternalGetObjectV2Execute(r)
}

/*
BillingentityinternalGetObjectV2 Retrieve an existing Billingentityinternal



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pkiBillingentityinternalID
 @return ApiBillingentityinternalGetObjectV2Request
*/
func (a *ObjectBillingentityinternalAPIService) BillingentityinternalGetObjectV2(ctx context.Context, pkiBillingentityinternalID int32) ApiBillingentityinternalGetObjectV2Request {
	return ApiBillingentityinternalGetObjectV2Request{
		ApiService: a,
		ctx: ctx,
		pkiBillingentityinternalID: pkiBillingentityinternalID,
	}
}

// Execute executes the request
//  @return BillingentityinternalGetObjectV2Response
func (a *ObjectBillingentityinternalAPIService) BillingentityinternalGetObjectV2Execute(r ApiBillingentityinternalGetObjectV2Request) (*BillingentityinternalGetObjectV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BillingentityinternalGetObjectV2Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectBillingentityinternalAPIService.BillingentityinternalGetObjectV2")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/2/object/billingentityinternal/{pkiBillingentityinternalID}"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiBillingentityinternalID"+"}", url.PathEscape(parameterValueToString(r.pkiBillingentityinternalID, "pkiBillingentityinternalID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pkiBillingentityinternalID < 0 {
		return localVarReturnValue, nil, reportError("pkiBillingentityinternalID must be greater than 0")
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
		if localVarHTTPResponse.StatusCode == 404 {
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
