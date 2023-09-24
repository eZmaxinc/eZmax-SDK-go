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


// ObjectUsergroupdelegationAPIService ObjectUsergroupdelegationAPI service
type ObjectUsergroupdelegationAPIService service

type ApiUsergroupdelegationCreateObjectV1Request struct {
	ctx context.Context
	ApiService *ObjectUsergroupdelegationAPIService
	usergroupdelegationCreateObjectV1Request *UsergroupdelegationCreateObjectV1Request
}

func (r ApiUsergroupdelegationCreateObjectV1Request) UsergroupdelegationCreateObjectV1Request(usergroupdelegationCreateObjectV1Request UsergroupdelegationCreateObjectV1Request) ApiUsergroupdelegationCreateObjectV1Request {
	r.usergroupdelegationCreateObjectV1Request = &usergroupdelegationCreateObjectV1Request
	return r
}

func (r ApiUsergroupdelegationCreateObjectV1Request) Execute() (*UsergroupdelegationCreateObjectV1Response, *http.Response, error) {
	return r.ApiService.UsergroupdelegationCreateObjectV1Execute(r)
}

/*
UsergroupdelegationCreateObjectV1 Create a new Usergroupdelegation

The endpoint allows to create one or many elements at once.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUsergroupdelegationCreateObjectV1Request
*/
func (a *ObjectUsergroupdelegationAPIService) UsergroupdelegationCreateObjectV1(ctx context.Context) ApiUsergroupdelegationCreateObjectV1Request {
	return ApiUsergroupdelegationCreateObjectV1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return UsergroupdelegationCreateObjectV1Response
func (a *ObjectUsergroupdelegationAPIService) UsergroupdelegationCreateObjectV1Execute(r ApiUsergroupdelegationCreateObjectV1Request) (*UsergroupdelegationCreateObjectV1Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UsergroupdelegationCreateObjectV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectUsergroupdelegationAPIService.UsergroupdelegationCreateObjectV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/usergroupdelegation"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.usergroupdelegationCreateObjectV1Request == nil {
		return localVarReturnValue, nil, reportError("usergroupdelegationCreateObjectV1Request is required and must be specified")
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
	localVarPostBody = r.usergroupdelegationCreateObjectV1Request
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

type ApiUsergroupdelegationDeleteObjectV1Request struct {
	ctx context.Context
	ApiService *ObjectUsergroupdelegationAPIService
	pkiUsergroupdelegationID int32
}

func (r ApiUsergroupdelegationDeleteObjectV1Request) Execute() (*UsergroupdelegationDeleteObjectV1Response, *http.Response, error) {
	return r.ApiService.UsergroupdelegationDeleteObjectV1Execute(r)
}

/*
UsergroupdelegationDeleteObjectV1 Delete an existing Usergroupdelegation



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pkiUsergroupdelegationID The unique ID of the Usergroupdelegation
 @return ApiUsergroupdelegationDeleteObjectV1Request
*/
func (a *ObjectUsergroupdelegationAPIService) UsergroupdelegationDeleteObjectV1(ctx context.Context, pkiUsergroupdelegationID int32) ApiUsergroupdelegationDeleteObjectV1Request {
	return ApiUsergroupdelegationDeleteObjectV1Request{
		ApiService: a,
		ctx: ctx,
		pkiUsergroupdelegationID: pkiUsergroupdelegationID,
	}
}

// Execute executes the request
//  @return UsergroupdelegationDeleteObjectV1Response
func (a *ObjectUsergroupdelegationAPIService) UsergroupdelegationDeleteObjectV1Execute(r ApiUsergroupdelegationDeleteObjectV1Request) (*UsergroupdelegationDeleteObjectV1Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UsergroupdelegationDeleteObjectV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectUsergroupdelegationAPIService.UsergroupdelegationDeleteObjectV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/usergroupdelegation/{pkiUsergroupdelegationID}"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiUsergroupdelegationID"+"}", url.PathEscape(parameterValueToString(r.pkiUsergroupdelegationID, "pkiUsergroupdelegationID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pkiUsergroupdelegationID < 0 {
		return localVarReturnValue, nil, reportError("pkiUsergroupdelegationID must be greater than 0")
	}
	if r.pkiUsergroupdelegationID > 65535 {
		return localVarReturnValue, nil, reportError("pkiUsergroupdelegationID must be less than 65535")
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

type ApiUsergroupdelegationEditObjectV1Request struct {
	ctx context.Context
	ApiService *ObjectUsergroupdelegationAPIService
	pkiUsergroupdelegationID int32
	usergroupdelegationEditObjectV1Request *UsergroupdelegationEditObjectV1Request
}

func (r ApiUsergroupdelegationEditObjectV1Request) UsergroupdelegationEditObjectV1Request(usergroupdelegationEditObjectV1Request UsergroupdelegationEditObjectV1Request) ApiUsergroupdelegationEditObjectV1Request {
	r.usergroupdelegationEditObjectV1Request = &usergroupdelegationEditObjectV1Request
	return r
}

func (r ApiUsergroupdelegationEditObjectV1Request) Execute() (*UsergroupdelegationEditObjectV1Response, *http.Response, error) {
	return r.ApiService.UsergroupdelegationEditObjectV1Execute(r)
}

/*
UsergroupdelegationEditObjectV1 Edit an existing Usergroupdelegation



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pkiUsergroupdelegationID The unique ID of the Usergroupdelegation
 @return ApiUsergroupdelegationEditObjectV1Request
*/
func (a *ObjectUsergroupdelegationAPIService) UsergroupdelegationEditObjectV1(ctx context.Context, pkiUsergroupdelegationID int32) ApiUsergroupdelegationEditObjectV1Request {
	return ApiUsergroupdelegationEditObjectV1Request{
		ApiService: a,
		ctx: ctx,
		pkiUsergroupdelegationID: pkiUsergroupdelegationID,
	}
}

// Execute executes the request
//  @return UsergroupdelegationEditObjectV1Response
func (a *ObjectUsergroupdelegationAPIService) UsergroupdelegationEditObjectV1Execute(r ApiUsergroupdelegationEditObjectV1Request) (*UsergroupdelegationEditObjectV1Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UsergroupdelegationEditObjectV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectUsergroupdelegationAPIService.UsergroupdelegationEditObjectV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/usergroupdelegation/{pkiUsergroupdelegationID}"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiUsergroupdelegationID"+"}", url.PathEscape(parameterValueToString(r.pkiUsergroupdelegationID, "pkiUsergroupdelegationID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pkiUsergroupdelegationID < 0 {
		return localVarReturnValue, nil, reportError("pkiUsergroupdelegationID must be greater than 0")
	}
	if r.pkiUsergroupdelegationID > 65535 {
		return localVarReturnValue, nil, reportError("pkiUsergroupdelegationID must be less than 65535")
	}
	if r.usergroupdelegationEditObjectV1Request == nil {
		return localVarReturnValue, nil, reportError("usergroupdelegationEditObjectV1Request is required and must be specified")
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
	localVarPostBody = r.usergroupdelegationEditObjectV1Request
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

type ApiUsergroupdelegationGetObjectV2Request struct {
	ctx context.Context
	ApiService *ObjectUsergroupdelegationAPIService
	pkiUsergroupdelegationID int32
}

func (r ApiUsergroupdelegationGetObjectV2Request) Execute() (*UsergroupdelegationGetObjectV2Response, *http.Response, error) {
	return r.ApiService.UsergroupdelegationGetObjectV2Execute(r)
}

/*
UsergroupdelegationGetObjectV2 Retrieve an existing Usergroupdelegation



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pkiUsergroupdelegationID The unique ID of the Usergroupdelegation
 @return ApiUsergroupdelegationGetObjectV2Request
*/
func (a *ObjectUsergroupdelegationAPIService) UsergroupdelegationGetObjectV2(ctx context.Context, pkiUsergroupdelegationID int32) ApiUsergroupdelegationGetObjectV2Request {
	return ApiUsergroupdelegationGetObjectV2Request{
		ApiService: a,
		ctx: ctx,
		pkiUsergroupdelegationID: pkiUsergroupdelegationID,
	}
}

// Execute executes the request
//  @return UsergroupdelegationGetObjectV2Response
func (a *ObjectUsergroupdelegationAPIService) UsergroupdelegationGetObjectV2Execute(r ApiUsergroupdelegationGetObjectV2Request) (*UsergroupdelegationGetObjectV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UsergroupdelegationGetObjectV2Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectUsergroupdelegationAPIService.UsergroupdelegationGetObjectV2")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/2/object/usergroupdelegation/{pkiUsergroupdelegationID}"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiUsergroupdelegationID"+"}", url.PathEscape(parameterValueToString(r.pkiUsergroupdelegationID, "pkiUsergroupdelegationID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pkiUsergroupdelegationID < 0 {
		return localVarReturnValue, nil, reportError("pkiUsergroupdelegationID must be greater than 0")
	}
	if r.pkiUsergroupdelegationID > 65535 {
		return localVarReturnValue, nil, reportError("pkiUsergroupdelegationID must be less than 65535")
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