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


// ObjectEzsigntemplateformfieldgroupAPIService ObjectEzsigntemplateformfieldgroupAPI service
type ObjectEzsigntemplateformfieldgroupAPIService service

type ApiEzsigntemplateformfieldgroupCreateObjectV1Request struct {
	ctx context.Context
	ApiService *ObjectEzsigntemplateformfieldgroupAPIService
	ezsigntemplateformfieldgroupCreateObjectV1Request *EzsigntemplateformfieldgroupCreateObjectV1Request
}

func (r ApiEzsigntemplateformfieldgroupCreateObjectV1Request) EzsigntemplateformfieldgroupCreateObjectV1Request(ezsigntemplateformfieldgroupCreateObjectV1Request EzsigntemplateformfieldgroupCreateObjectV1Request) ApiEzsigntemplateformfieldgroupCreateObjectV1Request {
	r.ezsigntemplateformfieldgroupCreateObjectV1Request = &ezsigntemplateformfieldgroupCreateObjectV1Request
	return r
}

func (r ApiEzsigntemplateformfieldgroupCreateObjectV1Request) Execute() (*EzsigntemplateformfieldgroupCreateObjectV1Response, *http.Response, error) {
	return r.ApiService.EzsigntemplateformfieldgroupCreateObjectV1Execute(r)
}

/*
EzsigntemplateformfieldgroupCreateObjectV1 Create a new Ezsigntemplateformfieldgroup

The endpoint allows to create one or many elements at once.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiEzsigntemplateformfieldgroupCreateObjectV1Request
*/
func (a *ObjectEzsigntemplateformfieldgroupAPIService) EzsigntemplateformfieldgroupCreateObjectV1(ctx context.Context) ApiEzsigntemplateformfieldgroupCreateObjectV1Request {
	return ApiEzsigntemplateformfieldgroupCreateObjectV1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return EzsigntemplateformfieldgroupCreateObjectV1Response
func (a *ObjectEzsigntemplateformfieldgroupAPIService) EzsigntemplateformfieldgroupCreateObjectV1Execute(r ApiEzsigntemplateformfieldgroupCreateObjectV1Request) (*EzsigntemplateformfieldgroupCreateObjectV1Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EzsigntemplateformfieldgroupCreateObjectV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectEzsigntemplateformfieldgroupAPIService.EzsigntemplateformfieldgroupCreateObjectV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/ezsigntemplateformfieldgroup"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.ezsigntemplateformfieldgroupCreateObjectV1Request == nil {
		return localVarReturnValue, nil, reportError("ezsigntemplateformfieldgroupCreateObjectV1Request is required and must be specified")
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
	localVarPostBody = r.ezsigntemplateformfieldgroupCreateObjectV1Request
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

type ApiEzsigntemplateformfieldgroupDeleteObjectV1Request struct {
	ctx context.Context
	ApiService *ObjectEzsigntemplateformfieldgroupAPIService
	pkiEzsigntemplateformfieldgroupID int32
}

func (r ApiEzsigntemplateformfieldgroupDeleteObjectV1Request) Execute() (*CommonResponse, *http.Response, error) {
	return r.ApiService.EzsigntemplateformfieldgroupDeleteObjectV1Execute(r)
}

/*
EzsigntemplateformfieldgroupDeleteObjectV1 Delete an existing Ezsigntemplateformfieldgroup



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pkiEzsigntemplateformfieldgroupID
 @return ApiEzsigntemplateformfieldgroupDeleteObjectV1Request
*/
func (a *ObjectEzsigntemplateformfieldgroupAPIService) EzsigntemplateformfieldgroupDeleteObjectV1(ctx context.Context, pkiEzsigntemplateformfieldgroupID int32) ApiEzsigntemplateformfieldgroupDeleteObjectV1Request {
	return ApiEzsigntemplateformfieldgroupDeleteObjectV1Request{
		ApiService: a,
		ctx: ctx,
		pkiEzsigntemplateformfieldgroupID: pkiEzsigntemplateformfieldgroupID,
	}
}

// Execute executes the request
//  @return CommonResponse
func (a *ObjectEzsigntemplateformfieldgroupAPIService) EzsigntemplateformfieldgroupDeleteObjectV1Execute(r ApiEzsigntemplateformfieldgroupDeleteObjectV1Request) (*CommonResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CommonResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectEzsigntemplateformfieldgroupAPIService.EzsigntemplateformfieldgroupDeleteObjectV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/ezsigntemplateformfieldgroup/{pkiEzsigntemplateformfieldgroupID}"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiEzsigntemplateformfieldgroupID"+"}", url.PathEscape(parameterValueToString(r.pkiEzsigntemplateformfieldgroupID, "pkiEzsigntemplateformfieldgroupID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pkiEzsigntemplateformfieldgroupID < 0 {
		return localVarReturnValue, nil, reportError("pkiEzsigntemplateformfieldgroupID must be greater than 0")
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

type ApiEzsigntemplateformfieldgroupEditObjectV1Request struct {
	ctx context.Context
	ApiService *ObjectEzsigntemplateformfieldgroupAPIService
	pkiEzsigntemplateformfieldgroupID int32
	ezsigntemplateformfieldgroupEditObjectV1Request *EzsigntemplateformfieldgroupEditObjectV1Request
}

func (r ApiEzsigntemplateformfieldgroupEditObjectV1Request) EzsigntemplateformfieldgroupEditObjectV1Request(ezsigntemplateformfieldgroupEditObjectV1Request EzsigntemplateformfieldgroupEditObjectV1Request) ApiEzsigntemplateformfieldgroupEditObjectV1Request {
	r.ezsigntemplateformfieldgroupEditObjectV1Request = &ezsigntemplateformfieldgroupEditObjectV1Request
	return r
}

func (r ApiEzsigntemplateformfieldgroupEditObjectV1Request) Execute() (*CommonResponse, *http.Response, error) {
	return r.ApiService.EzsigntemplateformfieldgroupEditObjectV1Execute(r)
}

/*
EzsigntemplateformfieldgroupEditObjectV1 Edit an existing Ezsigntemplateformfieldgroup



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pkiEzsigntemplateformfieldgroupID
 @return ApiEzsigntemplateformfieldgroupEditObjectV1Request
*/
func (a *ObjectEzsigntemplateformfieldgroupAPIService) EzsigntemplateformfieldgroupEditObjectV1(ctx context.Context, pkiEzsigntemplateformfieldgroupID int32) ApiEzsigntemplateformfieldgroupEditObjectV1Request {
	return ApiEzsigntemplateformfieldgroupEditObjectV1Request{
		ApiService: a,
		ctx: ctx,
		pkiEzsigntemplateformfieldgroupID: pkiEzsigntemplateformfieldgroupID,
	}
}

// Execute executes the request
//  @return CommonResponse
func (a *ObjectEzsigntemplateformfieldgroupAPIService) EzsigntemplateformfieldgroupEditObjectV1Execute(r ApiEzsigntemplateformfieldgroupEditObjectV1Request) (*CommonResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CommonResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectEzsigntemplateformfieldgroupAPIService.EzsigntemplateformfieldgroupEditObjectV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/ezsigntemplateformfieldgroup/{pkiEzsigntemplateformfieldgroupID}"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiEzsigntemplateformfieldgroupID"+"}", url.PathEscape(parameterValueToString(r.pkiEzsigntemplateformfieldgroupID, "pkiEzsigntemplateformfieldgroupID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pkiEzsigntemplateformfieldgroupID < 0 {
		return localVarReturnValue, nil, reportError("pkiEzsigntemplateformfieldgroupID must be greater than 0")
	}
	if r.ezsigntemplateformfieldgroupEditObjectV1Request == nil {
		return localVarReturnValue, nil, reportError("ezsigntemplateformfieldgroupEditObjectV1Request is required and must be specified")
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
	localVarPostBody = r.ezsigntemplateformfieldgroupEditObjectV1Request
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

type ApiEzsigntemplateformfieldgroupGetObjectV2Request struct {
	ctx context.Context
	ApiService *ObjectEzsigntemplateformfieldgroupAPIService
	pkiEzsigntemplateformfieldgroupID int32
}

func (r ApiEzsigntemplateformfieldgroupGetObjectV2Request) Execute() (*EzsigntemplateformfieldgroupGetObjectV2Response, *http.Response, error) {
	return r.ApiService.EzsigntemplateformfieldgroupGetObjectV2Execute(r)
}

/*
EzsigntemplateformfieldgroupGetObjectV2 Retrieve an existing Ezsigntemplateformfieldgroup



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pkiEzsigntemplateformfieldgroupID
 @return ApiEzsigntemplateformfieldgroupGetObjectV2Request
*/
func (a *ObjectEzsigntemplateformfieldgroupAPIService) EzsigntemplateformfieldgroupGetObjectV2(ctx context.Context, pkiEzsigntemplateformfieldgroupID int32) ApiEzsigntemplateformfieldgroupGetObjectV2Request {
	return ApiEzsigntemplateformfieldgroupGetObjectV2Request{
		ApiService: a,
		ctx: ctx,
		pkiEzsigntemplateformfieldgroupID: pkiEzsigntemplateformfieldgroupID,
	}
}

// Execute executes the request
//  @return EzsigntemplateformfieldgroupGetObjectV2Response
func (a *ObjectEzsigntemplateformfieldgroupAPIService) EzsigntemplateformfieldgroupGetObjectV2Execute(r ApiEzsigntemplateformfieldgroupGetObjectV2Request) (*EzsigntemplateformfieldgroupGetObjectV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EzsigntemplateformfieldgroupGetObjectV2Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectEzsigntemplateformfieldgroupAPIService.EzsigntemplateformfieldgroupGetObjectV2")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/2/object/ezsigntemplateformfieldgroup/{pkiEzsigntemplateformfieldgroupID}"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiEzsigntemplateformfieldgroupID"+"}", url.PathEscape(parameterValueToString(r.pkiEzsigntemplateformfieldgroupID, "pkiEzsigntemplateformfieldgroupID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pkiEzsigntemplateformfieldgroupID < 0 {
		return localVarReturnValue, nil, reportError("pkiEzsigntemplateformfieldgroupID must be greater than 0")
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
