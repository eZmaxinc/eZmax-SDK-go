/*
eZmax API Definition (Full)

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.2.2
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


// ObjectVariableexpenseAPIService ObjectVariableexpenseAPI service
type ObjectVariableexpenseAPIService service

type ApiVariableexpenseCreateObjectV1Request struct {
	ctx context.Context
	ApiService *ObjectVariableexpenseAPIService
	variableexpenseCreateObjectV1Request *VariableexpenseCreateObjectV1Request
}

func (r ApiVariableexpenseCreateObjectV1Request) VariableexpenseCreateObjectV1Request(variableexpenseCreateObjectV1Request VariableexpenseCreateObjectV1Request) ApiVariableexpenseCreateObjectV1Request {
	r.variableexpenseCreateObjectV1Request = &variableexpenseCreateObjectV1Request
	return r
}

func (r ApiVariableexpenseCreateObjectV1Request) Execute() (*VariableexpenseCreateObjectV1Response, *http.Response, error) {
	return r.ApiService.VariableexpenseCreateObjectV1Execute(r)
}

/*
VariableexpenseCreateObjectV1 Create a new Variableexpense

The endpoint allows to create one or many elements at once.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiVariableexpenseCreateObjectV1Request
*/
func (a *ObjectVariableexpenseAPIService) VariableexpenseCreateObjectV1(ctx context.Context) ApiVariableexpenseCreateObjectV1Request {
	return ApiVariableexpenseCreateObjectV1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return VariableexpenseCreateObjectV1Response
func (a *ObjectVariableexpenseAPIService) VariableexpenseCreateObjectV1Execute(r ApiVariableexpenseCreateObjectV1Request) (*VariableexpenseCreateObjectV1Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *VariableexpenseCreateObjectV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectVariableexpenseAPIService.VariableexpenseCreateObjectV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/variableexpense"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.variableexpenseCreateObjectV1Request == nil {
		return localVarReturnValue, nil, reportError("variableexpenseCreateObjectV1Request is required and must be specified")
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
	localVarPostBody = r.variableexpenseCreateObjectV1Request
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

type ApiVariableexpenseEditObjectV1Request struct {
	ctx context.Context
	ApiService *ObjectVariableexpenseAPIService
	pkiVariableexpenseID int32
	variableexpenseEditObjectV1Request *VariableexpenseEditObjectV1Request
}

func (r ApiVariableexpenseEditObjectV1Request) VariableexpenseEditObjectV1Request(variableexpenseEditObjectV1Request VariableexpenseEditObjectV1Request) ApiVariableexpenseEditObjectV1Request {
	r.variableexpenseEditObjectV1Request = &variableexpenseEditObjectV1Request
	return r
}

func (r ApiVariableexpenseEditObjectV1Request) Execute() (*VariableexpenseEditObjectV1Response, *http.Response, error) {
	return r.ApiService.VariableexpenseEditObjectV1Execute(r)
}

/*
VariableexpenseEditObjectV1 Edit an existing Variableexpense



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pkiVariableexpenseID
 @return ApiVariableexpenseEditObjectV1Request
*/
func (a *ObjectVariableexpenseAPIService) VariableexpenseEditObjectV1(ctx context.Context, pkiVariableexpenseID int32) ApiVariableexpenseEditObjectV1Request {
	return ApiVariableexpenseEditObjectV1Request{
		ApiService: a,
		ctx: ctx,
		pkiVariableexpenseID: pkiVariableexpenseID,
	}
}

// Execute executes the request
//  @return VariableexpenseEditObjectV1Response
func (a *ObjectVariableexpenseAPIService) VariableexpenseEditObjectV1Execute(r ApiVariableexpenseEditObjectV1Request) (*VariableexpenseEditObjectV1Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *VariableexpenseEditObjectV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectVariableexpenseAPIService.VariableexpenseEditObjectV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/variableexpense/{pkiVariableexpenseID}"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiVariableexpenseID"+"}", url.PathEscape(parameterValueToString(r.pkiVariableexpenseID, "pkiVariableexpenseID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pkiVariableexpenseID < 1 {
		return localVarReturnValue, nil, reportError("pkiVariableexpenseID must be greater than 1")
	}
	if r.pkiVariableexpenseID > 255 {
		return localVarReturnValue, nil, reportError("pkiVariableexpenseID must be less than 255")
	}
	if r.variableexpenseEditObjectV1Request == nil {
		return localVarReturnValue, nil, reportError("variableexpenseEditObjectV1Request is required and must be specified")
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
	localVarPostBody = r.variableexpenseEditObjectV1Request
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

type ApiVariableexpenseGetAutocompleteV2Request struct {
	ctx context.Context
	ApiService *ObjectVariableexpenseAPIService
	sSelector string
	eFilterActive *string
	sQuery *string
	acceptLanguage *HeaderAcceptLanguage
}

// Specify which results we want to display.
func (r ApiVariableexpenseGetAutocompleteV2Request) EFilterActive(eFilterActive string) ApiVariableexpenseGetAutocompleteV2Request {
	r.eFilterActive = &eFilterActive
	return r
}

// Allow to filter the returned results
func (r ApiVariableexpenseGetAutocompleteV2Request) SQuery(sQuery string) ApiVariableexpenseGetAutocompleteV2Request {
	r.sQuery = &sQuery
	return r
}

func (r ApiVariableexpenseGetAutocompleteV2Request) AcceptLanguage(acceptLanguage HeaderAcceptLanguage) ApiVariableexpenseGetAutocompleteV2Request {
	r.acceptLanguage = &acceptLanguage
	return r
}

func (r ApiVariableexpenseGetAutocompleteV2Request) Execute() (*VariableexpenseGetAutocompleteV2Response, *http.Response, error) {
	return r.ApiService.VariableexpenseGetAutocompleteV2Execute(r)
}

/*
VariableexpenseGetAutocompleteV2 Retrieve Variableexpenses and IDs

Get the list of Variableexpense to be used in a dropdown or autocomplete control.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param sSelector The type of Variableexpenses to return
 @return ApiVariableexpenseGetAutocompleteV2Request
*/
func (a *ObjectVariableexpenseAPIService) VariableexpenseGetAutocompleteV2(ctx context.Context, sSelector string) ApiVariableexpenseGetAutocompleteV2Request {
	return ApiVariableexpenseGetAutocompleteV2Request{
		ApiService: a,
		ctx: ctx,
		sSelector: sSelector,
	}
}

// Execute executes the request
//  @return VariableexpenseGetAutocompleteV2Response
func (a *ObjectVariableexpenseAPIService) VariableexpenseGetAutocompleteV2Execute(r ApiVariableexpenseGetAutocompleteV2Request) (*VariableexpenseGetAutocompleteV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *VariableexpenseGetAutocompleteV2Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectVariableexpenseAPIService.VariableexpenseGetAutocompleteV2")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/2/object/variableexpense/getAutocomplete/{sSelector}"
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

type ApiVariableexpenseGetListV1Request struct {
	ctx context.Context
	ApiService *ObjectVariableexpenseAPIService
	eOrderBy *string
	iRowMax *int32
	iRowOffset *int32
	acceptLanguage *HeaderAcceptLanguage
	sFilter *string
}

// Specify how you want the results to be sorted
func (r ApiVariableexpenseGetListV1Request) EOrderBy(eOrderBy string) ApiVariableexpenseGetListV1Request {
	r.eOrderBy = &eOrderBy
	return r
}

func (r ApiVariableexpenseGetListV1Request) IRowMax(iRowMax int32) ApiVariableexpenseGetListV1Request {
	r.iRowMax = &iRowMax
	return r
}

func (r ApiVariableexpenseGetListV1Request) IRowOffset(iRowOffset int32) ApiVariableexpenseGetListV1Request {
	r.iRowOffset = &iRowOffset
	return r
}

func (r ApiVariableexpenseGetListV1Request) AcceptLanguage(acceptLanguage HeaderAcceptLanguage) ApiVariableexpenseGetListV1Request {
	r.acceptLanguage = &acceptLanguage
	return r
}

func (r ApiVariableexpenseGetListV1Request) SFilter(sFilter string) ApiVariableexpenseGetListV1Request {
	r.sFilter = &sFilter
	return r
}

func (r ApiVariableexpenseGetListV1Request) Execute() (*VariableexpenseGetListV1Response, *http.Response, error) {
	return r.ApiService.VariableexpenseGetListV1Execute(r)
}

/*
VariableexpenseGetListV1 Retrieve Variableexpense list

Enum values that can be filtered in query parameter *sFilter*:

| Variable | Valid values |
|---|---|
| eVariableexpenseTaxable | Yes<br>No<br>Included |

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiVariableexpenseGetListV1Request
*/
func (a *ObjectVariableexpenseAPIService) VariableexpenseGetListV1(ctx context.Context) ApiVariableexpenseGetListV1Request {
	return ApiVariableexpenseGetListV1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return VariableexpenseGetListV1Response
func (a *ObjectVariableexpenseAPIService) VariableexpenseGetListV1Execute(r ApiVariableexpenseGetListV1Request) (*VariableexpenseGetListV1Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *VariableexpenseGetListV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectVariableexpenseAPIService.VariableexpenseGetListV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/variableexpense/getList"

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

type ApiVariableexpenseGetObjectV2Request struct {
	ctx context.Context
	ApiService *ObjectVariableexpenseAPIService
	pkiVariableexpenseID int32
}

func (r ApiVariableexpenseGetObjectV2Request) Execute() (*VariableexpenseGetObjectV2Response, *http.Response, error) {
	return r.ApiService.VariableexpenseGetObjectV2Execute(r)
}

/*
VariableexpenseGetObjectV2 Retrieve an existing Variableexpense



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pkiVariableexpenseID
 @return ApiVariableexpenseGetObjectV2Request
*/
func (a *ObjectVariableexpenseAPIService) VariableexpenseGetObjectV2(ctx context.Context, pkiVariableexpenseID int32) ApiVariableexpenseGetObjectV2Request {
	return ApiVariableexpenseGetObjectV2Request{
		ApiService: a,
		ctx: ctx,
		pkiVariableexpenseID: pkiVariableexpenseID,
	}
}

// Execute executes the request
//  @return VariableexpenseGetObjectV2Response
func (a *ObjectVariableexpenseAPIService) VariableexpenseGetObjectV2Execute(r ApiVariableexpenseGetObjectV2Request) (*VariableexpenseGetObjectV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *VariableexpenseGetObjectV2Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectVariableexpenseAPIService.VariableexpenseGetObjectV2")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/2/object/variableexpense/{pkiVariableexpenseID}"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiVariableexpenseID"+"}", url.PathEscape(parameterValueToString(r.pkiVariableexpenseID, "pkiVariableexpenseID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pkiVariableexpenseID < 1 {
		return localVarReturnValue, nil, reportError("pkiVariableexpenseID must be greater than 1")
	}
	if r.pkiVariableexpenseID > 255 {
		return localVarReturnValue, nil, reportError("pkiVariableexpenseID must be less than 255")
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
