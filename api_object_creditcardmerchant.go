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


// ObjectCreditcardmerchantAPIService ObjectCreditcardmerchantAPI service
type ObjectCreditcardmerchantAPIService service

type ApiCreditcardmerchantCreateObjectV1Request struct {
	ctx context.Context
	ApiService *ObjectCreditcardmerchantAPIService
	creditcardmerchantCreateObjectV1Request *CreditcardmerchantCreateObjectV1Request
}

func (r ApiCreditcardmerchantCreateObjectV1Request) CreditcardmerchantCreateObjectV1Request(creditcardmerchantCreateObjectV1Request CreditcardmerchantCreateObjectV1Request) ApiCreditcardmerchantCreateObjectV1Request {
	r.creditcardmerchantCreateObjectV1Request = &creditcardmerchantCreateObjectV1Request
	return r
}

func (r ApiCreditcardmerchantCreateObjectV1Request) Execute() (*CreditcardmerchantCreateObjectV1Response, *http.Response, error) {
	return r.ApiService.CreditcardmerchantCreateObjectV1Execute(r)
}

/*
CreditcardmerchantCreateObjectV1 Create a new Creditcardmerchant

The endpoint allows to create one or many elements at once.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreditcardmerchantCreateObjectV1Request
*/
func (a *ObjectCreditcardmerchantAPIService) CreditcardmerchantCreateObjectV1(ctx context.Context) ApiCreditcardmerchantCreateObjectV1Request {
	return ApiCreditcardmerchantCreateObjectV1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CreditcardmerchantCreateObjectV1Response
func (a *ObjectCreditcardmerchantAPIService) CreditcardmerchantCreateObjectV1Execute(r ApiCreditcardmerchantCreateObjectV1Request) (*CreditcardmerchantCreateObjectV1Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreditcardmerchantCreateObjectV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectCreditcardmerchantAPIService.CreditcardmerchantCreateObjectV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/creditcardmerchant"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.creditcardmerchantCreateObjectV1Request == nil {
		return localVarReturnValue, nil, reportError("creditcardmerchantCreateObjectV1Request is required and must be specified")
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
	localVarPostBody = r.creditcardmerchantCreateObjectV1Request
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

type ApiCreditcardmerchantDeleteObjectV1Request struct {
	ctx context.Context
	ApiService *ObjectCreditcardmerchantAPIService
	pkiCreditcardmerchantID int32
}

func (r ApiCreditcardmerchantDeleteObjectV1Request) Execute() (*CreditcardmerchantDeleteObjectV1Response, *http.Response, error) {
	return r.ApiService.CreditcardmerchantDeleteObjectV1Execute(r)
}

/*
CreditcardmerchantDeleteObjectV1 Delete an existing Creditcardmerchant



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pkiCreditcardmerchantID The unique ID of the Creditcardmerchant
 @return ApiCreditcardmerchantDeleteObjectV1Request
*/
func (a *ObjectCreditcardmerchantAPIService) CreditcardmerchantDeleteObjectV1(ctx context.Context, pkiCreditcardmerchantID int32) ApiCreditcardmerchantDeleteObjectV1Request {
	return ApiCreditcardmerchantDeleteObjectV1Request{
		ApiService: a,
		ctx: ctx,
		pkiCreditcardmerchantID: pkiCreditcardmerchantID,
	}
}

// Execute executes the request
//  @return CreditcardmerchantDeleteObjectV1Response
func (a *ObjectCreditcardmerchantAPIService) CreditcardmerchantDeleteObjectV1Execute(r ApiCreditcardmerchantDeleteObjectV1Request) (*CreditcardmerchantDeleteObjectV1Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreditcardmerchantDeleteObjectV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectCreditcardmerchantAPIService.CreditcardmerchantDeleteObjectV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/creditcardmerchant/{pkiCreditcardmerchantID}"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiCreditcardmerchantID"+"}", url.PathEscape(parameterValueToString(r.pkiCreditcardmerchantID, "pkiCreditcardmerchantID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pkiCreditcardmerchantID < 0 {
		return localVarReturnValue, nil, reportError("pkiCreditcardmerchantID must be greater than 0")
	}
	if r.pkiCreditcardmerchantID > 255 {
		return localVarReturnValue, nil, reportError("pkiCreditcardmerchantID must be less than 255")
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

type ApiCreditcardmerchantEditObjectV1Request struct {
	ctx context.Context
	ApiService *ObjectCreditcardmerchantAPIService
	pkiCreditcardmerchantID int32
	creditcardmerchantEditObjectV1Request *CreditcardmerchantEditObjectV1Request
}

func (r ApiCreditcardmerchantEditObjectV1Request) CreditcardmerchantEditObjectV1Request(creditcardmerchantEditObjectV1Request CreditcardmerchantEditObjectV1Request) ApiCreditcardmerchantEditObjectV1Request {
	r.creditcardmerchantEditObjectV1Request = &creditcardmerchantEditObjectV1Request
	return r
}

func (r ApiCreditcardmerchantEditObjectV1Request) Execute() (*CreditcardmerchantEditObjectV1Response, *http.Response, error) {
	return r.ApiService.CreditcardmerchantEditObjectV1Execute(r)
}

/*
CreditcardmerchantEditObjectV1 Edit an existing Creditcardmerchant



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pkiCreditcardmerchantID The unique ID of the Creditcardmerchant
 @return ApiCreditcardmerchantEditObjectV1Request
*/
func (a *ObjectCreditcardmerchantAPIService) CreditcardmerchantEditObjectV1(ctx context.Context, pkiCreditcardmerchantID int32) ApiCreditcardmerchantEditObjectV1Request {
	return ApiCreditcardmerchantEditObjectV1Request{
		ApiService: a,
		ctx: ctx,
		pkiCreditcardmerchantID: pkiCreditcardmerchantID,
	}
}

// Execute executes the request
//  @return CreditcardmerchantEditObjectV1Response
func (a *ObjectCreditcardmerchantAPIService) CreditcardmerchantEditObjectV1Execute(r ApiCreditcardmerchantEditObjectV1Request) (*CreditcardmerchantEditObjectV1Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreditcardmerchantEditObjectV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectCreditcardmerchantAPIService.CreditcardmerchantEditObjectV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/creditcardmerchant/{pkiCreditcardmerchantID}"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiCreditcardmerchantID"+"}", url.PathEscape(parameterValueToString(r.pkiCreditcardmerchantID, "pkiCreditcardmerchantID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pkiCreditcardmerchantID < 0 {
		return localVarReturnValue, nil, reportError("pkiCreditcardmerchantID must be greater than 0")
	}
	if r.pkiCreditcardmerchantID > 255 {
		return localVarReturnValue, nil, reportError("pkiCreditcardmerchantID must be less than 255")
	}
	if r.creditcardmerchantEditObjectV1Request == nil {
		return localVarReturnValue, nil, reportError("creditcardmerchantEditObjectV1Request is required and must be specified")
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
	localVarPostBody = r.creditcardmerchantEditObjectV1Request
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

type ApiCreditcardmerchantGetAutocompleteV2Request struct {
	ctx context.Context
	ApiService *ObjectCreditcardmerchantAPIService
	sSelector string
	eFilterActive *string
	sQuery *string
	acceptLanguage *HeaderAcceptLanguage
}

// Specify which results we want to display.
func (r ApiCreditcardmerchantGetAutocompleteV2Request) EFilterActive(eFilterActive string) ApiCreditcardmerchantGetAutocompleteV2Request {
	r.eFilterActive = &eFilterActive
	return r
}

// Allow to filter the returned results
func (r ApiCreditcardmerchantGetAutocompleteV2Request) SQuery(sQuery string) ApiCreditcardmerchantGetAutocompleteV2Request {
	r.sQuery = &sQuery
	return r
}

func (r ApiCreditcardmerchantGetAutocompleteV2Request) AcceptLanguage(acceptLanguage HeaderAcceptLanguage) ApiCreditcardmerchantGetAutocompleteV2Request {
	r.acceptLanguage = &acceptLanguage
	return r
}

func (r ApiCreditcardmerchantGetAutocompleteV2Request) Execute() (*CreditcardmerchantGetAutocompleteV2Response, *http.Response, error) {
	return r.ApiService.CreditcardmerchantGetAutocompleteV2Execute(r)
}

/*
CreditcardmerchantGetAutocompleteV2 Retrieve Creditcardmerchants and IDs

Get the list of Creditcardmerchant to be used in a dropdown or autocomplete control.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param sSelector The type of Creditcardmerchants to return
 @return ApiCreditcardmerchantGetAutocompleteV2Request
*/
func (a *ObjectCreditcardmerchantAPIService) CreditcardmerchantGetAutocompleteV2(ctx context.Context, sSelector string) ApiCreditcardmerchantGetAutocompleteV2Request {
	return ApiCreditcardmerchantGetAutocompleteV2Request{
		ApiService: a,
		ctx: ctx,
		sSelector: sSelector,
	}
}

// Execute executes the request
//  @return CreditcardmerchantGetAutocompleteV2Response
func (a *ObjectCreditcardmerchantAPIService) CreditcardmerchantGetAutocompleteV2Execute(r ApiCreditcardmerchantGetAutocompleteV2Request) (*CreditcardmerchantGetAutocompleteV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreditcardmerchantGetAutocompleteV2Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectCreditcardmerchantAPIService.CreditcardmerchantGetAutocompleteV2")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/2/object/creditcardmerchant/getAutocomplete/{sSelector}"
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

type ApiCreditcardmerchantGetListV1Request struct {
	ctx context.Context
	ApiService *ObjectCreditcardmerchantAPIService
	eOrderBy *string
	iRowMax *int32
	iRowOffset *int32
	acceptLanguage *HeaderAcceptLanguage
	sFilter *string
}

// Specify how you want the results to be sorted
func (r ApiCreditcardmerchantGetListV1Request) EOrderBy(eOrderBy string) ApiCreditcardmerchantGetListV1Request {
	r.eOrderBy = &eOrderBy
	return r
}

func (r ApiCreditcardmerchantGetListV1Request) IRowMax(iRowMax int32) ApiCreditcardmerchantGetListV1Request {
	r.iRowMax = &iRowMax
	return r
}

func (r ApiCreditcardmerchantGetListV1Request) IRowOffset(iRowOffset int32) ApiCreditcardmerchantGetListV1Request {
	r.iRowOffset = &iRowOffset
	return r
}

func (r ApiCreditcardmerchantGetListV1Request) AcceptLanguage(acceptLanguage HeaderAcceptLanguage) ApiCreditcardmerchantGetListV1Request {
	r.acceptLanguage = &acceptLanguage
	return r
}

func (r ApiCreditcardmerchantGetListV1Request) SFilter(sFilter string) ApiCreditcardmerchantGetListV1Request {
	r.sFilter = &sFilter
	return r
}

func (r ApiCreditcardmerchantGetListV1Request) Execute() (*CreditcardmerchantGetListV1Response, *http.Response, error) {
	return r.ApiService.CreditcardmerchantGetListV1Execute(r)
}

/*
CreditcardmerchantGetListV1 Retrieve Creditcardmerchant list



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreditcardmerchantGetListV1Request
*/
func (a *ObjectCreditcardmerchantAPIService) CreditcardmerchantGetListV1(ctx context.Context) ApiCreditcardmerchantGetListV1Request {
	return ApiCreditcardmerchantGetListV1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CreditcardmerchantGetListV1Response
func (a *ObjectCreditcardmerchantAPIService) CreditcardmerchantGetListV1Execute(r ApiCreditcardmerchantGetListV1Request) (*CreditcardmerchantGetListV1Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreditcardmerchantGetListV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectCreditcardmerchantAPIService.CreditcardmerchantGetListV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/creditcardmerchant/getList"

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

type ApiCreditcardmerchantGetObjectV2Request struct {
	ctx context.Context
	ApiService *ObjectCreditcardmerchantAPIService
	pkiCreditcardmerchantID int32
}

func (r ApiCreditcardmerchantGetObjectV2Request) Execute() (*CreditcardmerchantGetObjectV2Response, *http.Response, error) {
	return r.ApiService.CreditcardmerchantGetObjectV2Execute(r)
}

/*
CreditcardmerchantGetObjectV2 Retrieve an existing Creditcardmerchant



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pkiCreditcardmerchantID The unique ID of the Creditcardmerchant
 @return ApiCreditcardmerchantGetObjectV2Request
*/
func (a *ObjectCreditcardmerchantAPIService) CreditcardmerchantGetObjectV2(ctx context.Context, pkiCreditcardmerchantID int32) ApiCreditcardmerchantGetObjectV2Request {
	return ApiCreditcardmerchantGetObjectV2Request{
		ApiService: a,
		ctx: ctx,
		pkiCreditcardmerchantID: pkiCreditcardmerchantID,
	}
}

// Execute executes the request
//  @return CreditcardmerchantGetObjectV2Response
func (a *ObjectCreditcardmerchantAPIService) CreditcardmerchantGetObjectV2Execute(r ApiCreditcardmerchantGetObjectV2Request) (*CreditcardmerchantGetObjectV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreditcardmerchantGetObjectV2Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectCreditcardmerchantAPIService.CreditcardmerchantGetObjectV2")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/2/object/creditcardmerchant/{pkiCreditcardmerchantID}"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiCreditcardmerchantID"+"}", url.PathEscape(parameterValueToString(r.pkiCreditcardmerchantID, "pkiCreditcardmerchantID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pkiCreditcardmerchantID < 0 {
		return localVarReturnValue, nil, reportError("pkiCreditcardmerchantID must be greater than 0")
	}
	if r.pkiCreditcardmerchantID > 255 {
		return localVarReturnValue, nil, reportError("pkiCreditcardmerchantID must be less than 255")
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
