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


// ObjectEzsignformfieldgroupAPIService ObjectEzsignformfieldgroupAPI service
type ObjectEzsignformfieldgroupAPIService service

type ApiEzsignformfieldgroupCreateObjectV1Request struct {
	ctx context.Context
	ApiService *ObjectEzsignformfieldgroupAPIService
	ezsignformfieldgroupCreateObjectV1Request *EzsignformfieldgroupCreateObjectV1Request
}

func (r ApiEzsignformfieldgroupCreateObjectV1Request) EzsignformfieldgroupCreateObjectV1Request(ezsignformfieldgroupCreateObjectV1Request EzsignformfieldgroupCreateObjectV1Request) ApiEzsignformfieldgroupCreateObjectV1Request {
	r.ezsignformfieldgroupCreateObjectV1Request = &ezsignformfieldgroupCreateObjectV1Request
	return r
}

func (r ApiEzsignformfieldgroupCreateObjectV1Request) Execute() (*EzsignformfieldgroupCreateObjectV1Response, *http.Response, error) {
	return r.ApiService.EzsignformfieldgroupCreateObjectV1Execute(r)
}

/*
EzsignformfieldgroupCreateObjectV1 Create a new Ezsignformfieldgroup

The endpoint allows to create one or many elements at once.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiEzsignformfieldgroupCreateObjectV1Request
*/
func (a *ObjectEzsignformfieldgroupAPIService) EzsignformfieldgroupCreateObjectV1(ctx context.Context) ApiEzsignformfieldgroupCreateObjectV1Request {
	return ApiEzsignformfieldgroupCreateObjectV1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return EzsignformfieldgroupCreateObjectV1Response
func (a *ObjectEzsignformfieldgroupAPIService) EzsignformfieldgroupCreateObjectV1Execute(r ApiEzsignformfieldgroupCreateObjectV1Request) (*EzsignformfieldgroupCreateObjectV1Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EzsignformfieldgroupCreateObjectV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectEzsignformfieldgroupAPIService.EzsignformfieldgroupCreateObjectV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/ezsignformfieldgroup"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.ezsignformfieldgroupCreateObjectV1Request == nil {
		return localVarReturnValue, nil, reportError("ezsignformfieldgroupCreateObjectV1Request is required and must be specified")
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
	localVarPostBody = r.ezsignformfieldgroupCreateObjectV1Request
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

type ApiEzsignformfieldgroupDeleteObjectV1Request struct {
	ctx context.Context
	ApiService *ObjectEzsignformfieldgroupAPIService
	pkiEzsignformfieldgroupID int32
}

func (r ApiEzsignformfieldgroupDeleteObjectV1Request) Execute() (*EzsignformfieldgroupDeleteObjectV1Response, *http.Response, error) {
	return r.ApiService.EzsignformfieldgroupDeleteObjectV1Execute(r)
}

/*
EzsignformfieldgroupDeleteObjectV1 Delete an existing Ezsignformfieldgroup



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pkiEzsignformfieldgroupID
 @return ApiEzsignformfieldgroupDeleteObjectV1Request
*/
func (a *ObjectEzsignformfieldgroupAPIService) EzsignformfieldgroupDeleteObjectV1(ctx context.Context, pkiEzsignformfieldgroupID int32) ApiEzsignformfieldgroupDeleteObjectV1Request {
	return ApiEzsignformfieldgroupDeleteObjectV1Request{
		ApiService: a,
		ctx: ctx,
		pkiEzsignformfieldgroupID: pkiEzsignformfieldgroupID,
	}
}

// Execute executes the request
//  @return EzsignformfieldgroupDeleteObjectV1Response
func (a *ObjectEzsignformfieldgroupAPIService) EzsignformfieldgroupDeleteObjectV1Execute(r ApiEzsignformfieldgroupDeleteObjectV1Request) (*EzsignformfieldgroupDeleteObjectV1Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EzsignformfieldgroupDeleteObjectV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectEzsignformfieldgroupAPIService.EzsignformfieldgroupDeleteObjectV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/ezsignformfieldgroup/{pkiEzsignformfieldgroupID}"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiEzsignformfieldgroupID"+"}", url.PathEscape(parameterValueToString(r.pkiEzsignformfieldgroupID, "pkiEzsignformfieldgroupID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pkiEzsignformfieldgroupID < 0 {
		return localVarReturnValue, nil, reportError("pkiEzsignformfieldgroupID must be greater than 0")
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

type ApiEzsignformfieldgroupEditObjectV1Request struct {
	ctx context.Context
	ApiService *ObjectEzsignformfieldgroupAPIService
	pkiEzsignformfieldgroupID int32
	ezsignformfieldgroupEditObjectV1Request *EzsignformfieldgroupEditObjectV1Request
}

func (r ApiEzsignformfieldgroupEditObjectV1Request) EzsignformfieldgroupEditObjectV1Request(ezsignformfieldgroupEditObjectV1Request EzsignformfieldgroupEditObjectV1Request) ApiEzsignformfieldgroupEditObjectV1Request {
	r.ezsignformfieldgroupEditObjectV1Request = &ezsignformfieldgroupEditObjectV1Request
	return r
}

func (r ApiEzsignformfieldgroupEditObjectV1Request) Execute() (*EzsignformfieldgroupEditObjectV1Response, *http.Response, error) {
	return r.ApiService.EzsignformfieldgroupEditObjectV1Execute(r)
}

/*
EzsignformfieldgroupEditObjectV1 Edit an existing Ezsignformfieldgroup



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pkiEzsignformfieldgroupID
 @return ApiEzsignformfieldgroupEditObjectV1Request
*/
func (a *ObjectEzsignformfieldgroupAPIService) EzsignformfieldgroupEditObjectV1(ctx context.Context, pkiEzsignformfieldgroupID int32) ApiEzsignformfieldgroupEditObjectV1Request {
	return ApiEzsignformfieldgroupEditObjectV1Request{
		ApiService: a,
		ctx: ctx,
		pkiEzsignformfieldgroupID: pkiEzsignformfieldgroupID,
	}
}

// Execute executes the request
//  @return EzsignformfieldgroupEditObjectV1Response
func (a *ObjectEzsignformfieldgroupAPIService) EzsignformfieldgroupEditObjectV1Execute(r ApiEzsignformfieldgroupEditObjectV1Request) (*EzsignformfieldgroupEditObjectV1Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EzsignformfieldgroupEditObjectV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectEzsignformfieldgroupAPIService.EzsignformfieldgroupEditObjectV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/ezsignformfieldgroup/{pkiEzsignformfieldgroupID}"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiEzsignformfieldgroupID"+"}", url.PathEscape(parameterValueToString(r.pkiEzsignformfieldgroupID, "pkiEzsignformfieldgroupID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pkiEzsignformfieldgroupID < 0 {
		return localVarReturnValue, nil, reportError("pkiEzsignformfieldgroupID must be greater than 0")
	}
	if r.ezsignformfieldgroupEditObjectV1Request == nil {
		return localVarReturnValue, nil, reportError("ezsignformfieldgroupEditObjectV1Request is required and must be specified")
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
	localVarPostBody = r.ezsignformfieldgroupEditObjectV1Request
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

type ApiEzsignformfieldgroupGetObjectV2Request struct {
	ctx context.Context
	ApiService *ObjectEzsignformfieldgroupAPIService
	pkiEzsignformfieldgroupID int32
}

func (r ApiEzsignformfieldgroupGetObjectV2Request) Execute() (*EzsignformfieldgroupGetObjectV2Response, *http.Response, error) {
	return r.ApiService.EzsignformfieldgroupGetObjectV2Execute(r)
}

/*
EzsignformfieldgroupGetObjectV2 Retrieve an existing Ezsignformfieldgroup



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pkiEzsignformfieldgroupID
 @return ApiEzsignformfieldgroupGetObjectV2Request
*/
func (a *ObjectEzsignformfieldgroupAPIService) EzsignformfieldgroupGetObjectV2(ctx context.Context, pkiEzsignformfieldgroupID int32) ApiEzsignformfieldgroupGetObjectV2Request {
	return ApiEzsignformfieldgroupGetObjectV2Request{
		ApiService: a,
		ctx: ctx,
		pkiEzsignformfieldgroupID: pkiEzsignformfieldgroupID,
	}
}

// Execute executes the request
//  @return EzsignformfieldgroupGetObjectV2Response
func (a *ObjectEzsignformfieldgroupAPIService) EzsignformfieldgroupGetObjectV2Execute(r ApiEzsignformfieldgroupGetObjectV2Request) (*EzsignformfieldgroupGetObjectV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EzsignformfieldgroupGetObjectV2Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectEzsignformfieldgroupAPIService.EzsignformfieldgroupGetObjectV2")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/2/object/ezsignformfieldgroup/{pkiEzsignformfieldgroupID}"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiEzsignformfieldgroupID"+"}", url.PathEscape(parameterValueToString(r.pkiEzsignformfieldgroupID, "pkiEzsignformfieldgroupID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pkiEzsignformfieldgroupID < 0 {
		return localVarReturnValue, nil, reportError("pkiEzsignformfieldgroupID must be greater than 0")
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
