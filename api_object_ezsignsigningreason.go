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


// ObjectEzsignsigningreasonAPIService ObjectEzsignsigningreasonAPI service
type ObjectEzsignsigningreasonAPIService service

type ApiEzsignsigningreasonCreateObjectV1Request struct {
	ctx context.Context
	ApiService *ObjectEzsignsigningreasonAPIService
	ezsignsigningreasonCreateObjectV1Request *EzsignsigningreasonCreateObjectV1Request
}

func (r ApiEzsignsigningreasonCreateObjectV1Request) EzsignsigningreasonCreateObjectV1Request(ezsignsigningreasonCreateObjectV1Request EzsignsigningreasonCreateObjectV1Request) ApiEzsignsigningreasonCreateObjectV1Request {
	r.ezsignsigningreasonCreateObjectV1Request = &ezsignsigningreasonCreateObjectV1Request
	return r
}

func (r ApiEzsignsigningreasonCreateObjectV1Request) Execute() (*EzsignsigningreasonCreateObjectV1Response, *http.Response, error) {
	return r.ApiService.EzsignsigningreasonCreateObjectV1Execute(r)
}

/*
EzsignsigningreasonCreateObjectV1 Create a new Ezsignsigningreason

The endpoint allows to create one or many elements at once.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiEzsignsigningreasonCreateObjectV1Request
*/
func (a *ObjectEzsignsigningreasonAPIService) EzsignsigningreasonCreateObjectV1(ctx context.Context) ApiEzsignsigningreasonCreateObjectV1Request {
	return ApiEzsignsigningreasonCreateObjectV1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return EzsignsigningreasonCreateObjectV1Response
func (a *ObjectEzsignsigningreasonAPIService) EzsignsigningreasonCreateObjectV1Execute(r ApiEzsignsigningreasonCreateObjectV1Request) (*EzsignsigningreasonCreateObjectV1Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EzsignsigningreasonCreateObjectV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectEzsignsigningreasonAPIService.EzsignsigningreasonCreateObjectV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/ezsignsigningreason"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.ezsignsigningreasonCreateObjectV1Request == nil {
		return localVarReturnValue, nil, reportError("ezsignsigningreasonCreateObjectV1Request is required and must be specified")
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
	localVarPostBody = r.ezsignsigningreasonCreateObjectV1Request
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

type ApiEzsignsigningreasonEditObjectV1Request struct {
	ctx context.Context
	ApiService *ObjectEzsignsigningreasonAPIService
	pkiEzsignsigningreasonID int32
	ezsignsigningreasonEditObjectV1Request *EzsignsigningreasonEditObjectV1Request
}

func (r ApiEzsignsigningreasonEditObjectV1Request) EzsignsigningreasonEditObjectV1Request(ezsignsigningreasonEditObjectV1Request EzsignsigningreasonEditObjectV1Request) ApiEzsignsigningreasonEditObjectV1Request {
	r.ezsignsigningreasonEditObjectV1Request = &ezsignsigningreasonEditObjectV1Request
	return r
}

func (r ApiEzsignsigningreasonEditObjectV1Request) Execute() (*EzsignsigningreasonEditObjectV1Response, *http.Response, error) {
	return r.ApiService.EzsignsigningreasonEditObjectV1Execute(r)
}

/*
EzsignsigningreasonEditObjectV1 Edit an existing Ezsignsigningreason



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pkiEzsignsigningreasonID The unique ID of the Ezsignsigningreason
 @return ApiEzsignsigningreasonEditObjectV1Request
*/
func (a *ObjectEzsignsigningreasonAPIService) EzsignsigningreasonEditObjectV1(ctx context.Context, pkiEzsignsigningreasonID int32) ApiEzsignsigningreasonEditObjectV1Request {
	return ApiEzsignsigningreasonEditObjectV1Request{
		ApiService: a,
		ctx: ctx,
		pkiEzsignsigningreasonID: pkiEzsignsigningreasonID,
	}
}

// Execute executes the request
//  @return EzsignsigningreasonEditObjectV1Response
func (a *ObjectEzsignsigningreasonAPIService) EzsignsigningreasonEditObjectV1Execute(r ApiEzsignsigningreasonEditObjectV1Request) (*EzsignsigningreasonEditObjectV1Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EzsignsigningreasonEditObjectV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectEzsignsigningreasonAPIService.EzsignsigningreasonEditObjectV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/ezsignsigningreason/{pkiEzsignsigningreasonID}"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiEzsignsigningreasonID"+"}", url.PathEscape(parameterValueToString(r.pkiEzsignsigningreasonID, "pkiEzsignsigningreasonID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pkiEzsignsigningreasonID < 0 {
		return localVarReturnValue, nil, reportError("pkiEzsignsigningreasonID must be greater than 0")
	}
	if r.pkiEzsignsigningreasonID > 255 {
		return localVarReturnValue, nil, reportError("pkiEzsignsigningreasonID must be less than 255")
	}
	if r.ezsignsigningreasonEditObjectV1Request == nil {
		return localVarReturnValue, nil, reportError("ezsignsigningreasonEditObjectV1Request is required and must be specified")
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
	localVarPostBody = r.ezsignsigningreasonEditObjectV1Request
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

type ApiEzsignsigningreasonGetAutocompleteV2Request struct {
	ctx context.Context
	ApiService *ObjectEzsignsigningreasonAPIService
	sSelector string
	eFilterActive *string
	sQuery *string
	acceptLanguage *HeaderAcceptLanguage
}

// Specify which results we want to display.
func (r ApiEzsignsigningreasonGetAutocompleteV2Request) EFilterActive(eFilterActive string) ApiEzsignsigningreasonGetAutocompleteV2Request {
	r.eFilterActive = &eFilterActive
	return r
}

// Allow to filter the returned results
func (r ApiEzsignsigningreasonGetAutocompleteV2Request) SQuery(sQuery string) ApiEzsignsigningreasonGetAutocompleteV2Request {
	r.sQuery = &sQuery
	return r
}

func (r ApiEzsignsigningreasonGetAutocompleteV2Request) AcceptLanguage(acceptLanguage HeaderAcceptLanguage) ApiEzsignsigningreasonGetAutocompleteV2Request {
	r.acceptLanguage = &acceptLanguage
	return r
}

func (r ApiEzsignsigningreasonGetAutocompleteV2Request) Execute() (*EzsignsigningreasonGetAutocompleteV2Response, *http.Response, error) {
	return r.ApiService.EzsignsigningreasonGetAutocompleteV2Execute(r)
}

/*
EzsignsigningreasonGetAutocompleteV2 Retrieve Ezsignsigningreasons and IDs

Get the list of Ezsignsigningreason to be used in a dropdown or autocomplete control.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param sSelector The type of Ezsignsigningreasons to return
 @return ApiEzsignsigningreasonGetAutocompleteV2Request
*/
func (a *ObjectEzsignsigningreasonAPIService) EzsignsigningreasonGetAutocompleteV2(ctx context.Context, sSelector string) ApiEzsignsigningreasonGetAutocompleteV2Request {
	return ApiEzsignsigningreasonGetAutocompleteV2Request{
		ApiService: a,
		ctx: ctx,
		sSelector: sSelector,
	}
}

// Execute executes the request
//  @return EzsignsigningreasonGetAutocompleteV2Response
func (a *ObjectEzsignsigningreasonAPIService) EzsignsigningreasonGetAutocompleteV2Execute(r ApiEzsignsigningreasonGetAutocompleteV2Request) (*EzsignsigningreasonGetAutocompleteV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EzsignsigningreasonGetAutocompleteV2Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectEzsignsigningreasonAPIService.EzsignsigningreasonGetAutocompleteV2")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/2/object/ezsignsigningreason/getAutocomplete/{sSelector}"
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

type ApiEzsignsigningreasonGetListV1Request struct {
	ctx context.Context
	ApiService *ObjectEzsignsigningreasonAPIService
	eOrderBy *string
	iRowMax *int32
	iRowOffset *int32
	acceptLanguage *HeaderAcceptLanguage
	sFilter *string
}

// Specify how you want the results to be sorted
func (r ApiEzsignsigningreasonGetListV1Request) EOrderBy(eOrderBy string) ApiEzsignsigningreasonGetListV1Request {
	r.eOrderBy = &eOrderBy
	return r
}

func (r ApiEzsignsigningreasonGetListV1Request) IRowMax(iRowMax int32) ApiEzsignsigningreasonGetListV1Request {
	r.iRowMax = &iRowMax
	return r
}

func (r ApiEzsignsigningreasonGetListV1Request) IRowOffset(iRowOffset int32) ApiEzsignsigningreasonGetListV1Request {
	r.iRowOffset = &iRowOffset
	return r
}

func (r ApiEzsignsigningreasonGetListV1Request) AcceptLanguage(acceptLanguage HeaderAcceptLanguage) ApiEzsignsigningreasonGetListV1Request {
	r.acceptLanguage = &acceptLanguage
	return r
}

func (r ApiEzsignsigningreasonGetListV1Request) SFilter(sFilter string) ApiEzsignsigningreasonGetListV1Request {
	r.sFilter = &sFilter
	return r
}

func (r ApiEzsignsigningreasonGetListV1Request) Execute() (*EzsignsigningreasonGetListV1Response, *http.Response, error) {
	return r.ApiService.EzsignsigningreasonGetListV1Execute(r)
}

/*
EzsignsigningreasonGetListV1 Retrieve Ezsignsigningreason list



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiEzsignsigningreasonGetListV1Request
*/
func (a *ObjectEzsignsigningreasonAPIService) EzsignsigningreasonGetListV1(ctx context.Context) ApiEzsignsigningreasonGetListV1Request {
	return ApiEzsignsigningreasonGetListV1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return EzsignsigningreasonGetListV1Response
func (a *ObjectEzsignsigningreasonAPIService) EzsignsigningreasonGetListV1Execute(r ApiEzsignsigningreasonGetListV1Request) (*EzsignsigningreasonGetListV1Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EzsignsigningreasonGetListV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectEzsignsigningreasonAPIService.EzsignsigningreasonGetListV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/ezsignsigningreason/getList"

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

type ApiEzsignsigningreasonGetObjectV2Request struct {
	ctx context.Context
	ApiService *ObjectEzsignsigningreasonAPIService
	pkiEzsignsigningreasonID int32
}

func (r ApiEzsignsigningreasonGetObjectV2Request) Execute() (*EzsignsigningreasonGetObjectV2Response, *http.Response, error) {
	return r.ApiService.EzsignsigningreasonGetObjectV2Execute(r)
}

/*
EzsignsigningreasonGetObjectV2 Retrieve an existing Ezsignsigningreason



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pkiEzsignsigningreasonID The unique ID of the Ezsignsigningreason
 @return ApiEzsignsigningreasonGetObjectV2Request
*/
func (a *ObjectEzsignsigningreasonAPIService) EzsignsigningreasonGetObjectV2(ctx context.Context, pkiEzsignsigningreasonID int32) ApiEzsignsigningreasonGetObjectV2Request {
	return ApiEzsignsigningreasonGetObjectV2Request{
		ApiService: a,
		ctx: ctx,
		pkiEzsignsigningreasonID: pkiEzsignsigningreasonID,
	}
}

// Execute executes the request
//  @return EzsignsigningreasonGetObjectV2Response
func (a *ObjectEzsignsigningreasonAPIService) EzsignsigningreasonGetObjectV2Execute(r ApiEzsignsigningreasonGetObjectV2Request) (*EzsignsigningreasonGetObjectV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EzsignsigningreasonGetObjectV2Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectEzsignsigningreasonAPIService.EzsignsigningreasonGetObjectV2")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/2/object/ezsignsigningreason/{pkiEzsignsigningreasonID}"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiEzsignsigningreasonID"+"}", url.PathEscape(parameterValueToString(r.pkiEzsignsigningreasonID, "pkiEzsignsigningreasonID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pkiEzsignsigningreasonID < 0 {
		return localVarReturnValue, nil, reportError("pkiEzsignsigningreasonID must be greater than 0")
	}
	if r.pkiEzsignsigningreasonID > 255 {
		return localVarReturnValue, nil, reportError("pkiEzsignsigningreasonID must be less than 255")
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
