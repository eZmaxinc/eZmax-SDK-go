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


// ObjectCreditcardclientAPIService ObjectCreditcardclientAPI service
type ObjectCreditcardclientAPIService service

type ApiCreditcardclientCreateObjectV1Request struct {
	ctx context.Context
	ApiService *ObjectCreditcardclientAPIService
	creditcardclientCreateObjectV1Request *CreditcardclientCreateObjectV1Request
}

func (r ApiCreditcardclientCreateObjectV1Request) CreditcardclientCreateObjectV1Request(creditcardclientCreateObjectV1Request CreditcardclientCreateObjectV1Request) ApiCreditcardclientCreateObjectV1Request {
	r.creditcardclientCreateObjectV1Request = &creditcardclientCreateObjectV1Request
	return r
}

func (r ApiCreditcardclientCreateObjectV1Request) Execute() (*CreditcardclientCreateObjectV1Response, *http.Response, error) {
	return r.ApiService.CreditcardclientCreateObjectV1Execute(r)
}

/*
CreditcardclientCreateObjectV1 Create a new Creditcardclient

The endpoint allows to create one or many elements at once.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreditcardclientCreateObjectV1Request
*/
func (a *ObjectCreditcardclientAPIService) CreditcardclientCreateObjectV1(ctx context.Context) ApiCreditcardclientCreateObjectV1Request {
	return ApiCreditcardclientCreateObjectV1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CreditcardclientCreateObjectV1Response
func (a *ObjectCreditcardclientAPIService) CreditcardclientCreateObjectV1Execute(r ApiCreditcardclientCreateObjectV1Request) (*CreditcardclientCreateObjectV1Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreditcardclientCreateObjectV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectCreditcardclientAPIService.CreditcardclientCreateObjectV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/creditcardclient"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.creditcardclientCreateObjectV1Request == nil {
		return localVarReturnValue, nil, reportError("creditcardclientCreateObjectV1Request is required and must be specified")
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
	localVarPostBody = r.creditcardclientCreateObjectV1Request
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
			var v CommonResponseErrorCreditcardValidation
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

type ApiCreditcardclientDeleteObjectV1Request struct {
	ctx context.Context
	ApiService *ObjectCreditcardclientAPIService
	pkiCreditcardclientID int32
}

func (r ApiCreditcardclientDeleteObjectV1Request) Execute() (*CommonResponse, *http.Response, error) {
	return r.ApiService.CreditcardclientDeleteObjectV1Execute(r)
}

/*
CreditcardclientDeleteObjectV1 Delete an existing Creditcardclient



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pkiCreditcardclientID The unique ID of the Creditcardclient
 @return ApiCreditcardclientDeleteObjectV1Request
*/
func (a *ObjectCreditcardclientAPIService) CreditcardclientDeleteObjectV1(ctx context.Context, pkiCreditcardclientID int32) ApiCreditcardclientDeleteObjectV1Request {
	return ApiCreditcardclientDeleteObjectV1Request{
		ApiService: a,
		ctx: ctx,
		pkiCreditcardclientID: pkiCreditcardclientID,
	}
}

// Execute executes the request
//  @return CommonResponse
func (a *ObjectCreditcardclientAPIService) CreditcardclientDeleteObjectV1Execute(r ApiCreditcardclientDeleteObjectV1Request) (*CommonResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CommonResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectCreditcardclientAPIService.CreditcardclientDeleteObjectV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/creditcardclient/{pkiCreditcardclientID}"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiCreditcardclientID"+"}", url.PathEscape(parameterValueToString(r.pkiCreditcardclientID, "pkiCreditcardclientID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pkiCreditcardclientID < 0 {
		return localVarReturnValue, nil, reportError("pkiCreditcardclientID must be greater than 0")
	}
	if r.pkiCreditcardclientID > 65535 {
		return localVarReturnValue, nil, reportError("pkiCreditcardclientID must be less than 65535")
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

type ApiCreditcardclientEditObjectV1Request struct {
	ctx context.Context
	ApiService *ObjectCreditcardclientAPIService
	pkiCreditcardclientID int32
	creditcardclientEditObjectV1Request *CreditcardclientEditObjectV1Request
}

func (r ApiCreditcardclientEditObjectV1Request) CreditcardclientEditObjectV1Request(creditcardclientEditObjectV1Request CreditcardclientEditObjectV1Request) ApiCreditcardclientEditObjectV1Request {
	r.creditcardclientEditObjectV1Request = &creditcardclientEditObjectV1Request
	return r
}

func (r ApiCreditcardclientEditObjectV1Request) Execute() (*CommonResponse, *http.Response, error) {
	return r.ApiService.CreditcardclientEditObjectV1Execute(r)
}

/*
CreditcardclientEditObjectV1 Edit an existing Creditcardclient



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pkiCreditcardclientID The unique ID of the Creditcardclient
 @return ApiCreditcardclientEditObjectV1Request
*/
func (a *ObjectCreditcardclientAPIService) CreditcardclientEditObjectV1(ctx context.Context, pkiCreditcardclientID int32) ApiCreditcardclientEditObjectV1Request {
	return ApiCreditcardclientEditObjectV1Request{
		ApiService: a,
		ctx: ctx,
		pkiCreditcardclientID: pkiCreditcardclientID,
	}
}

// Execute executes the request
//  @return CommonResponse
func (a *ObjectCreditcardclientAPIService) CreditcardclientEditObjectV1Execute(r ApiCreditcardclientEditObjectV1Request) (*CommonResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CommonResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectCreditcardclientAPIService.CreditcardclientEditObjectV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/creditcardclient/{pkiCreditcardclientID}"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiCreditcardclientID"+"}", url.PathEscape(parameterValueToString(r.pkiCreditcardclientID, "pkiCreditcardclientID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pkiCreditcardclientID < 0 {
		return localVarReturnValue, nil, reportError("pkiCreditcardclientID must be greater than 0")
	}
	if r.pkiCreditcardclientID > 65535 {
		return localVarReturnValue, nil, reportError("pkiCreditcardclientID must be less than 65535")
	}
	if r.creditcardclientEditObjectV1Request == nil {
		return localVarReturnValue, nil, reportError("creditcardclientEditObjectV1Request is required and must be specified")
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
	localVarPostBody = r.creditcardclientEditObjectV1Request
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
			var v CommonResponseErrorCreditcardValidation
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

type ApiCreditcardclientGetAutocompleteV2Request struct {
	ctx context.Context
	ApiService *ObjectCreditcardclientAPIService
	sSelector string
	eFilterActive *string
	sQuery *string
	acceptLanguage *HeaderAcceptLanguage
}

// Specify which results we want to display.
func (r ApiCreditcardclientGetAutocompleteV2Request) EFilterActive(eFilterActive string) ApiCreditcardclientGetAutocompleteV2Request {
	r.eFilterActive = &eFilterActive
	return r
}

// Allow to filter the returned results
func (r ApiCreditcardclientGetAutocompleteV2Request) SQuery(sQuery string) ApiCreditcardclientGetAutocompleteV2Request {
	r.sQuery = &sQuery
	return r
}

func (r ApiCreditcardclientGetAutocompleteV2Request) AcceptLanguage(acceptLanguage HeaderAcceptLanguage) ApiCreditcardclientGetAutocompleteV2Request {
	r.acceptLanguage = &acceptLanguage
	return r
}

func (r ApiCreditcardclientGetAutocompleteV2Request) Execute() (*CreditcardclientGetAutocompleteV2Response, *http.Response, error) {
	return r.ApiService.CreditcardclientGetAutocompleteV2Execute(r)
}

/*
CreditcardclientGetAutocompleteV2 Retrieve Creditcardclients and IDs

Get the list of Creditcardclient to be used in a dropdown or autocomplete control.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param sSelector The type of Creditcardclients to return
 @return ApiCreditcardclientGetAutocompleteV2Request
*/
func (a *ObjectCreditcardclientAPIService) CreditcardclientGetAutocompleteV2(ctx context.Context, sSelector string) ApiCreditcardclientGetAutocompleteV2Request {
	return ApiCreditcardclientGetAutocompleteV2Request{
		ApiService: a,
		ctx: ctx,
		sSelector: sSelector,
	}
}

// Execute executes the request
//  @return CreditcardclientGetAutocompleteV2Response
func (a *ObjectCreditcardclientAPIService) CreditcardclientGetAutocompleteV2Execute(r ApiCreditcardclientGetAutocompleteV2Request) (*CreditcardclientGetAutocompleteV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreditcardclientGetAutocompleteV2Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectCreditcardclientAPIService.CreditcardclientGetAutocompleteV2")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/2/object/creditcardclient/getAutocomplete/{sSelector}"
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

type ApiCreditcardclientGetListV1Request struct {
	ctx context.Context
	ApiService *ObjectCreditcardclientAPIService
	eOrderBy *string
	iRowMax *int32
	iRowOffset *int32
	acceptLanguage *HeaderAcceptLanguage
	sFilter *string
}

// Specify how you want the results to be sorted
func (r ApiCreditcardclientGetListV1Request) EOrderBy(eOrderBy string) ApiCreditcardclientGetListV1Request {
	r.eOrderBy = &eOrderBy
	return r
}

func (r ApiCreditcardclientGetListV1Request) IRowMax(iRowMax int32) ApiCreditcardclientGetListV1Request {
	r.iRowMax = &iRowMax
	return r
}

func (r ApiCreditcardclientGetListV1Request) IRowOffset(iRowOffset int32) ApiCreditcardclientGetListV1Request {
	r.iRowOffset = &iRowOffset
	return r
}

func (r ApiCreditcardclientGetListV1Request) AcceptLanguage(acceptLanguage HeaderAcceptLanguage) ApiCreditcardclientGetListV1Request {
	r.acceptLanguage = &acceptLanguage
	return r
}

func (r ApiCreditcardclientGetListV1Request) SFilter(sFilter string) ApiCreditcardclientGetListV1Request {
	r.sFilter = &sFilter
	return r
}

func (r ApiCreditcardclientGetListV1Request) Execute() (*CreditcardclientGetListV1Response, *http.Response, error) {
	return r.ApiService.CreditcardclientGetListV1Execute(r)
}

/*
CreditcardclientGetListV1 Retrieve Creditcardclient list



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreditcardclientGetListV1Request
*/
func (a *ObjectCreditcardclientAPIService) CreditcardclientGetListV1(ctx context.Context) ApiCreditcardclientGetListV1Request {
	return ApiCreditcardclientGetListV1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CreditcardclientGetListV1Response
func (a *ObjectCreditcardclientAPIService) CreditcardclientGetListV1Execute(r ApiCreditcardclientGetListV1Request) (*CreditcardclientGetListV1Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreditcardclientGetListV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectCreditcardclientAPIService.CreditcardclientGetListV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/creditcardclient/getList"

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

type ApiCreditcardclientGetObjectV2Request struct {
	ctx context.Context
	ApiService *ObjectCreditcardclientAPIService
	pkiCreditcardclientID int32
}

func (r ApiCreditcardclientGetObjectV2Request) Execute() (*CreditcardclientGetObjectV2Response, *http.Response, error) {
	return r.ApiService.CreditcardclientGetObjectV2Execute(r)
}

/*
CreditcardclientGetObjectV2 Retrieve an existing Creditcardclient



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pkiCreditcardclientID The unique ID of the Creditcardclient
 @return ApiCreditcardclientGetObjectV2Request
*/
func (a *ObjectCreditcardclientAPIService) CreditcardclientGetObjectV2(ctx context.Context, pkiCreditcardclientID int32) ApiCreditcardclientGetObjectV2Request {
	return ApiCreditcardclientGetObjectV2Request{
		ApiService: a,
		ctx: ctx,
		pkiCreditcardclientID: pkiCreditcardclientID,
	}
}

// Execute executes the request
//  @return CreditcardclientGetObjectV2Response
func (a *ObjectCreditcardclientAPIService) CreditcardclientGetObjectV2Execute(r ApiCreditcardclientGetObjectV2Request) (*CreditcardclientGetObjectV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreditcardclientGetObjectV2Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectCreditcardclientAPIService.CreditcardclientGetObjectV2")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/2/object/creditcardclient/{pkiCreditcardclientID}"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiCreditcardclientID"+"}", url.PathEscape(parameterValueToString(r.pkiCreditcardclientID, "pkiCreditcardclientID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pkiCreditcardclientID < 0 {
		return localVarReturnValue, nil, reportError("pkiCreditcardclientID must be greater than 0")
	}
	if r.pkiCreditcardclientID > 65535 {
		return localVarReturnValue, nil, reportError("pkiCreditcardclientID must be less than 65535")
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

type ApiCreditcardclientPatchObjectV1Request struct {
	ctx context.Context
	ApiService *ObjectCreditcardclientAPIService
	pkiCreditcardclientID int32
	creditcardclientPatchObjectV1Request *CreditcardclientPatchObjectV1Request
}

func (r ApiCreditcardclientPatchObjectV1Request) CreditcardclientPatchObjectV1Request(creditcardclientPatchObjectV1Request CreditcardclientPatchObjectV1Request) ApiCreditcardclientPatchObjectV1Request {
	r.creditcardclientPatchObjectV1Request = &creditcardclientPatchObjectV1Request
	return r
}

func (r ApiCreditcardclientPatchObjectV1Request) Execute() (*CommonResponse, *http.Response, error) {
	return r.ApiService.CreditcardclientPatchObjectV1Execute(r)
}

/*
CreditcardclientPatchObjectV1 Patch an existing Creditcardclient



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pkiCreditcardclientID The unique ID of the Creditcardclient
 @return ApiCreditcardclientPatchObjectV1Request
*/
func (a *ObjectCreditcardclientAPIService) CreditcardclientPatchObjectV1(ctx context.Context, pkiCreditcardclientID int32) ApiCreditcardclientPatchObjectV1Request {
	return ApiCreditcardclientPatchObjectV1Request{
		ApiService: a,
		ctx: ctx,
		pkiCreditcardclientID: pkiCreditcardclientID,
	}
}

// Execute executes the request
//  @return CommonResponse
func (a *ObjectCreditcardclientAPIService) CreditcardclientPatchObjectV1Execute(r ApiCreditcardclientPatchObjectV1Request) (*CommonResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CommonResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectCreditcardclientAPIService.CreditcardclientPatchObjectV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/creditcardclient/{pkiCreditcardclientID}"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiCreditcardclientID"+"}", url.PathEscape(parameterValueToString(r.pkiCreditcardclientID, "pkiCreditcardclientID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pkiCreditcardclientID < 0 {
		return localVarReturnValue, nil, reportError("pkiCreditcardclientID must be greater than 0")
	}
	if r.pkiCreditcardclientID > 65535 {
		return localVarReturnValue, nil, reportError("pkiCreditcardclientID must be less than 65535")
	}
	if r.creditcardclientPatchObjectV1Request == nil {
		return localVarReturnValue, nil, reportError("creditcardclientPatchObjectV1Request is required and must be specified")
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
	localVarPostBody = r.creditcardclientPatchObjectV1Request
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
