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


// ObjectEzsignfoldertypeAPIService ObjectEzsignfoldertypeAPI service
type ObjectEzsignfoldertypeAPIService service

type ApiEzsignfoldertypeCreateObjectV3Request struct {
	ctx context.Context
	ApiService *ObjectEzsignfoldertypeAPIService
	ezsignfoldertypeCreateObjectV3Request *EzsignfoldertypeCreateObjectV3Request
}

func (r ApiEzsignfoldertypeCreateObjectV3Request) EzsignfoldertypeCreateObjectV3Request(ezsignfoldertypeCreateObjectV3Request EzsignfoldertypeCreateObjectV3Request) ApiEzsignfoldertypeCreateObjectV3Request {
	r.ezsignfoldertypeCreateObjectV3Request = &ezsignfoldertypeCreateObjectV3Request
	return r
}

func (r ApiEzsignfoldertypeCreateObjectV3Request) Execute() (*EzsignfoldertypeCreateObjectV3Response, *http.Response, error) {
	return r.ApiService.EzsignfoldertypeCreateObjectV3Execute(r)
}

/*
EzsignfoldertypeCreateObjectV3 Create a new Ezsignfoldertype

The endpoint allows to create one or many elements at once.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiEzsignfoldertypeCreateObjectV3Request
*/
func (a *ObjectEzsignfoldertypeAPIService) EzsignfoldertypeCreateObjectV3(ctx context.Context) ApiEzsignfoldertypeCreateObjectV3Request {
	return ApiEzsignfoldertypeCreateObjectV3Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return EzsignfoldertypeCreateObjectV3Response
func (a *ObjectEzsignfoldertypeAPIService) EzsignfoldertypeCreateObjectV3Execute(r ApiEzsignfoldertypeCreateObjectV3Request) (*EzsignfoldertypeCreateObjectV3Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EzsignfoldertypeCreateObjectV3Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectEzsignfoldertypeAPIService.EzsignfoldertypeCreateObjectV3")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/3/object/ezsignfoldertype"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.ezsignfoldertypeCreateObjectV3Request == nil {
		return localVarReturnValue, nil, reportError("ezsignfoldertypeCreateObjectV3Request is required and must be specified")
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
	localVarPostBody = r.ezsignfoldertypeCreateObjectV3Request
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

type ApiEzsignfoldertypeEditObjectV3Request struct {
	ctx context.Context
	ApiService *ObjectEzsignfoldertypeAPIService
	pkiEzsignfoldertypeID int32
	ezsignfoldertypeEditObjectV3Request *EzsignfoldertypeEditObjectV3Request
}

func (r ApiEzsignfoldertypeEditObjectV3Request) EzsignfoldertypeEditObjectV3Request(ezsignfoldertypeEditObjectV3Request EzsignfoldertypeEditObjectV3Request) ApiEzsignfoldertypeEditObjectV3Request {
	r.ezsignfoldertypeEditObjectV3Request = &ezsignfoldertypeEditObjectV3Request
	return r
}

func (r ApiEzsignfoldertypeEditObjectV3Request) Execute() (*EzsignfoldertypeEditObjectV3Response, *http.Response, error) {
	return r.ApiService.EzsignfoldertypeEditObjectV3Execute(r)
}

/*
EzsignfoldertypeEditObjectV3 Edit an existing Ezsignfoldertype



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pkiEzsignfoldertypeID
 @return ApiEzsignfoldertypeEditObjectV3Request
*/
func (a *ObjectEzsignfoldertypeAPIService) EzsignfoldertypeEditObjectV3(ctx context.Context, pkiEzsignfoldertypeID int32) ApiEzsignfoldertypeEditObjectV3Request {
	return ApiEzsignfoldertypeEditObjectV3Request{
		ApiService: a,
		ctx: ctx,
		pkiEzsignfoldertypeID: pkiEzsignfoldertypeID,
	}
}

// Execute executes the request
//  @return EzsignfoldertypeEditObjectV3Response
func (a *ObjectEzsignfoldertypeAPIService) EzsignfoldertypeEditObjectV3Execute(r ApiEzsignfoldertypeEditObjectV3Request) (*EzsignfoldertypeEditObjectV3Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EzsignfoldertypeEditObjectV3Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectEzsignfoldertypeAPIService.EzsignfoldertypeEditObjectV3")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/3/object/ezsignfoldertype/{pkiEzsignfoldertypeID}"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiEzsignfoldertypeID"+"}", url.PathEscape(parameterValueToString(r.pkiEzsignfoldertypeID, "pkiEzsignfoldertypeID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pkiEzsignfoldertypeID < 0 {
		return localVarReturnValue, nil, reportError("pkiEzsignfoldertypeID must be greater than 0")
	}
	if r.pkiEzsignfoldertypeID > 65535 {
		return localVarReturnValue, nil, reportError("pkiEzsignfoldertypeID must be less than 65535")
	}
	if r.ezsignfoldertypeEditObjectV3Request == nil {
		return localVarReturnValue, nil, reportError("ezsignfoldertypeEditObjectV3Request is required and must be specified")
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
	localVarPostBody = r.ezsignfoldertypeEditObjectV3Request
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

type ApiEzsignfoldertypeGetAutocompleteV2Request struct {
	ctx context.Context
	ApiService *ObjectEzsignfoldertypeAPIService
	sSelector string
	eFilterActive *string
	sQuery *string
	acceptLanguage *HeaderAcceptLanguage
}

// Specify which results we want to display.
func (r ApiEzsignfoldertypeGetAutocompleteV2Request) EFilterActive(eFilterActive string) ApiEzsignfoldertypeGetAutocompleteV2Request {
	r.eFilterActive = &eFilterActive
	return r
}

// Allow to filter the returned results
func (r ApiEzsignfoldertypeGetAutocompleteV2Request) SQuery(sQuery string) ApiEzsignfoldertypeGetAutocompleteV2Request {
	r.sQuery = &sQuery
	return r
}

func (r ApiEzsignfoldertypeGetAutocompleteV2Request) AcceptLanguage(acceptLanguage HeaderAcceptLanguage) ApiEzsignfoldertypeGetAutocompleteV2Request {
	r.acceptLanguage = &acceptLanguage
	return r
}

func (r ApiEzsignfoldertypeGetAutocompleteV2Request) Execute() (*EzsignfoldertypeGetAutocompleteV2Response, *http.Response, error) {
	return r.ApiService.EzsignfoldertypeGetAutocompleteV2Execute(r)
}

/*
EzsignfoldertypeGetAutocompleteV2 Retrieve Ezsignfoldertypes and IDs

Get the list of Ezsignfoldertype to be used in a dropdown or autocomplete control.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param sSelector The type of Ezsignfoldertypes to return
 @return ApiEzsignfoldertypeGetAutocompleteV2Request
*/
func (a *ObjectEzsignfoldertypeAPIService) EzsignfoldertypeGetAutocompleteV2(ctx context.Context, sSelector string) ApiEzsignfoldertypeGetAutocompleteV2Request {
	return ApiEzsignfoldertypeGetAutocompleteV2Request{
		ApiService: a,
		ctx: ctx,
		sSelector: sSelector,
	}
}

// Execute executes the request
//  @return EzsignfoldertypeGetAutocompleteV2Response
func (a *ObjectEzsignfoldertypeAPIService) EzsignfoldertypeGetAutocompleteV2Execute(r ApiEzsignfoldertypeGetAutocompleteV2Request) (*EzsignfoldertypeGetAutocompleteV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EzsignfoldertypeGetAutocompleteV2Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectEzsignfoldertypeAPIService.EzsignfoldertypeGetAutocompleteV2")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/2/object/ezsignfoldertype/getAutocomplete/{sSelector}"
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

type ApiEzsignfoldertypeGetListV1Request struct {
	ctx context.Context
	ApiService *ObjectEzsignfoldertypeAPIService
	eOrderBy *string
	iRowMax *int32
	iRowOffset *int32
	acceptLanguage *HeaderAcceptLanguage
	sFilter *string
}

// Specify how you want the results to be sorted
func (r ApiEzsignfoldertypeGetListV1Request) EOrderBy(eOrderBy string) ApiEzsignfoldertypeGetListV1Request {
	r.eOrderBy = &eOrderBy
	return r
}

func (r ApiEzsignfoldertypeGetListV1Request) IRowMax(iRowMax int32) ApiEzsignfoldertypeGetListV1Request {
	r.iRowMax = &iRowMax
	return r
}

func (r ApiEzsignfoldertypeGetListV1Request) IRowOffset(iRowOffset int32) ApiEzsignfoldertypeGetListV1Request {
	r.iRowOffset = &iRowOffset
	return r
}

func (r ApiEzsignfoldertypeGetListV1Request) AcceptLanguage(acceptLanguage HeaderAcceptLanguage) ApiEzsignfoldertypeGetListV1Request {
	r.acceptLanguage = &acceptLanguage
	return r
}

func (r ApiEzsignfoldertypeGetListV1Request) SFilter(sFilter string) ApiEzsignfoldertypeGetListV1Request {
	r.sFilter = &sFilter
	return r
}

func (r ApiEzsignfoldertypeGetListV1Request) Execute() (*EzsignfoldertypeGetListV1Response, *http.Response, error) {
	return r.ApiService.EzsignfoldertypeGetListV1Execute(r)
}

/*
EzsignfoldertypeGetListV1 Retrieve Ezsignfoldertype list

Enum values that can be filtered in query parameter *sFilter*:

| Variable | Valid values |
|---|---|
| eEzsignfoldertypePrivacylevel | User<br>Usergroup |

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiEzsignfoldertypeGetListV1Request
*/
func (a *ObjectEzsignfoldertypeAPIService) EzsignfoldertypeGetListV1(ctx context.Context) ApiEzsignfoldertypeGetListV1Request {
	return ApiEzsignfoldertypeGetListV1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return EzsignfoldertypeGetListV1Response
func (a *ObjectEzsignfoldertypeAPIService) EzsignfoldertypeGetListV1Execute(r ApiEzsignfoldertypeGetListV1Request) (*EzsignfoldertypeGetListV1Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EzsignfoldertypeGetListV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectEzsignfoldertypeAPIService.EzsignfoldertypeGetListV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/ezsignfoldertype/getList"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.eOrderBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "eOrderBy", r.eOrderBy, "form", "")
	}
	if r.iRowMax != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "iRowMax", r.iRowMax, "form", "")
	}
	if r.iRowOffset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "iRowOffset", r.iRowOffset, "form", "")
	} else {
		var defaultValue int32 = 0
		r.iRowOffset = &defaultValue
	}
	if r.sFilter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sFilter", r.sFilter, "form", "")
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

type ApiEzsignfoldertypeGetObjectV2Request struct {
	ctx context.Context
	ApiService *ObjectEzsignfoldertypeAPIService
	pkiEzsignfoldertypeID int32
}

func (r ApiEzsignfoldertypeGetObjectV2Request) Execute() (*EzsignfoldertypeGetObjectV2Response, *http.Response, error) {
	return r.ApiService.EzsignfoldertypeGetObjectV2Execute(r)
}

/*
EzsignfoldertypeGetObjectV2 Retrieve an existing Ezsignfoldertype



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pkiEzsignfoldertypeID
 @return ApiEzsignfoldertypeGetObjectV2Request

Deprecated
*/
func (a *ObjectEzsignfoldertypeAPIService) EzsignfoldertypeGetObjectV2(ctx context.Context, pkiEzsignfoldertypeID int32) ApiEzsignfoldertypeGetObjectV2Request {
	return ApiEzsignfoldertypeGetObjectV2Request{
		ApiService: a,
		ctx: ctx,
		pkiEzsignfoldertypeID: pkiEzsignfoldertypeID,
	}
}

// Execute executes the request
//  @return EzsignfoldertypeGetObjectV2Response
// Deprecated
func (a *ObjectEzsignfoldertypeAPIService) EzsignfoldertypeGetObjectV2Execute(r ApiEzsignfoldertypeGetObjectV2Request) (*EzsignfoldertypeGetObjectV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EzsignfoldertypeGetObjectV2Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectEzsignfoldertypeAPIService.EzsignfoldertypeGetObjectV2")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/2/object/ezsignfoldertype/{pkiEzsignfoldertypeID}"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiEzsignfoldertypeID"+"}", url.PathEscape(parameterValueToString(r.pkiEzsignfoldertypeID, "pkiEzsignfoldertypeID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pkiEzsignfoldertypeID < 0 {
		return localVarReturnValue, nil, reportError("pkiEzsignfoldertypeID must be greater than 0")
	}
	if r.pkiEzsignfoldertypeID > 65535 {
		return localVarReturnValue, nil, reportError("pkiEzsignfoldertypeID must be less than 65535")
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

type ApiEzsignfoldertypeGetObjectV4Request struct {
	ctx context.Context
	ApiService *ObjectEzsignfoldertypeAPIService
	pkiEzsignfoldertypeID int32
}

func (r ApiEzsignfoldertypeGetObjectV4Request) Execute() (*EzsignfoldertypeGetObjectV4Response, *http.Response, error) {
	return r.ApiService.EzsignfoldertypeGetObjectV4Execute(r)
}

/*
EzsignfoldertypeGetObjectV4 Retrieve an existing Ezsignfoldertype



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pkiEzsignfoldertypeID
 @return ApiEzsignfoldertypeGetObjectV4Request
*/
func (a *ObjectEzsignfoldertypeAPIService) EzsignfoldertypeGetObjectV4(ctx context.Context, pkiEzsignfoldertypeID int32) ApiEzsignfoldertypeGetObjectV4Request {
	return ApiEzsignfoldertypeGetObjectV4Request{
		ApiService: a,
		ctx: ctx,
		pkiEzsignfoldertypeID: pkiEzsignfoldertypeID,
	}
}

// Execute executes the request
//  @return EzsignfoldertypeGetObjectV4Response
func (a *ObjectEzsignfoldertypeAPIService) EzsignfoldertypeGetObjectV4Execute(r ApiEzsignfoldertypeGetObjectV4Request) (*EzsignfoldertypeGetObjectV4Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EzsignfoldertypeGetObjectV4Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectEzsignfoldertypeAPIService.EzsignfoldertypeGetObjectV4")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/4/object/ezsignfoldertype/{pkiEzsignfoldertypeID}"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiEzsignfoldertypeID"+"}", url.PathEscape(parameterValueToString(r.pkiEzsignfoldertypeID, "pkiEzsignfoldertypeID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pkiEzsignfoldertypeID < 0 {
		return localVarReturnValue, nil, reportError("pkiEzsignfoldertypeID must be greater than 0")
	}
	if r.pkiEzsignfoldertypeID > 65535 {
		return localVarReturnValue, nil, reportError("pkiEzsignfoldertypeID must be less than 65535")
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
