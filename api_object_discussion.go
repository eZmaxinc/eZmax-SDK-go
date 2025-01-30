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


// ObjectDiscussionAPIService ObjectDiscussionAPI service
type ObjectDiscussionAPIService service

type ApiDiscussionCreateObjectV1Request struct {
	ctx context.Context
	ApiService *ObjectDiscussionAPIService
	discussionCreateObjectV1Request *DiscussionCreateObjectV1Request
}

func (r ApiDiscussionCreateObjectV1Request) DiscussionCreateObjectV1Request(discussionCreateObjectV1Request DiscussionCreateObjectV1Request) ApiDiscussionCreateObjectV1Request {
	r.discussionCreateObjectV1Request = &discussionCreateObjectV1Request
	return r
}

func (r ApiDiscussionCreateObjectV1Request) Execute() (*DiscussionCreateObjectV1Response, *http.Response, error) {
	return r.ApiService.DiscussionCreateObjectV1Execute(r)
}

/*
DiscussionCreateObjectV1 Create a new Discussion

The endpoint allows to create one or many elements at once.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDiscussionCreateObjectV1Request
*/
func (a *ObjectDiscussionAPIService) DiscussionCreateObjectV1(ctx context.Context) ApiDiscussionCreateObjectV1Request {
	return ApiDiscussionCreateObjectV1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return DiscussionCreateObjectV1Response
func (a *ObjectDiscussionAPIService) DiscussionCreateObjectV1Execute(r ApiDiscussionCreateObjectV1Request) (*DiscussionCreateObjectV1Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DiscussionCreateObjectV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectDiscussionAPIService.DiscussionCreateObjectV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/discussion"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.discussionCreateObjectV1Request == nil {
		return localVarReturnValue, nil, reportError("discussionCreateObjectV1Request is required and must be specified")
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
	localVarPostBody = r.discussionCreateObjectV1Request
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

type ApiDiscussionDeleteObjectV1Request struct {
	ctx context.Context
	ApiService *ObjectDiscussionAPIService
	pkiDiscussionID int32
}

func (r ApiDiscussionDeleteObjectV1Request) Execute() (*DiscussionDeleteObjectV1Response, *http.Response, error) {
	return r.ApiService.DiscussionDeleteObjectV1Execute(r)
}

/*
DiscussionDeleteObjectV1 Delete an existing Discussion



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pkiDiscussionID The unique ID of the Discussion
 @return ApiDiscussionDeleteObjectV1Request
*/
func (a *ObjectDiscussionAPIService) DiscussionDeleteObjectV1(ctx context.Context, pkiDiscussionID int32) ApiDiscussionDeleteObjectV1Request {
	return ApiDiscussionDeleteObjectV1Request{
		ApiService: a,
		ctx: ctx,
		pkiDiscussionID: pkiDiscussionID,
	}
}

// Execute executes the request
//  @return DiscussionDeleteObjectV1Response
func (a *ObjectDiscussionAPIService) DiscussionDeleteObjectV1Execute(r ApiDiscussionDeleteObjectV1Request) (*DiscussionDeleteObjectV1Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DiscussionDeleteObjectV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectDiscussionAPIService.DiscussionDeleteObjectV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/discussion/{pkiDiscussionID}"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiDiscussionID"+"}", url.PathEscape(parameterValueToString(r.pkiDiscussionID, "pkiDiscussionID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pkiDiscussionID < 0 {
		return localVarReturnValue, nil, reportError("pkiDiscussionID must be greater than 0")
	}
	if r.pkiDiscussionID > 16777215 {
		return localVarReturnValue, nil, reportError("pkiDiscussionID must be less than 16777215")
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

type ApiDiscussionGetObjectV2Request struct {
	ctx context.Context
	ApiService *ObjectDiscussionAPIService
	pkiDiscussionID int32
}

func (r ApiDiscussionGetObjectV2Request) Execute() (*DiscussionGetObjectV2Response, *http.Response, error) {
	return r.ApiService.DiscussionGetObjectV2Execute(r)
}

/*
DiscussionGetObjectV2 Retrieve an existing Discussion



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pkiDiscussionID The unique ID of the Discussion
 @return ApiDiscussionGetObjectV2Request
*/
func (a *ObjectDiscussionAPIService) DiscussionGetObjectV2(ctx context.Context, pkiDiscussionID int32) ApiDiscussionGetObjectV2Request {
	return ApiDiscussionGetObjectV2Request{
		ApiService: a,
		ctx: ctx,
		pkiDiscussionID: pkiDiscussionID,
	}
}

// Execute executes the request
//  @return DiscussionGetObjectV2Response
func (a *ObjectDiscussionAPIService) DiscussionGetObjectV2Execute(r ApiDiscussionGetObjectV2Request) (*DiscussionGetObjectV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DiscussionGetObjectV2Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectDiscussionAPIService.DiscussionGetObjectV2")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/2/object/discussion/{pkiDiscussionID}"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiDiscussionID"+"}", url.PathEscape(parameterValueToString(r.pkiDiscussionID, "pkiDiscussionID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pkiDiscussionID < 0 {
		return localVarReturnValue, nil, reportError("pkiDiscussionID must be greater than 0")
	}
	if r.pkiDiscussionID > 16777215 {
		return localVarReturnValue, nil, reportError("pkiDiscussionID must be less than 16777215")
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

type ApiDiscussionPatchObjectV1Request struct {
	ctx context.Context
	ApiService *ObjectDiscussionAPIService
	pkiDiscussionID int32
	discussionPatchObjectV1Request *DiscussionPatchObjectV1Request
}

func (r ApiDiscussionPatchObjectV1Request) DiscussionPatchObjectV1Request(discussionPatchObjectV1Request DiscussionPatchObjectV1Request) ApiDiscussionPatchObjectV1Request {
	r.discussionPatchObjectV1Request = &discussionPatchObjectV1Request
	return r
}

func (r ApiDiscussionPatchObjectV1Request) Execute() (*DiscussionPatchObjectV1Response, *http.Response, error) {
	return r.ApiService.DiscussionPatchObjectV1Execute(r)
}

/*
DiscussionPatchObjectV1 Patch an existing Discussion



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pkiDiscussionID The unique ID of the Discussion
 @return ApiDiscussionPatchObjectV1Request
*/
func (a *ObjectDiscussionAPIService) DiscussionPatchObjectV1(ctx context.Context, pkiDiscussionID int32) ApiDiscussionPatchObjectV1Request {
	return ApiDiscussionPatchObjectV1Request{
		ApiService: a,
		ctx: ctx,
		pkiDiscussionID: pkiDiscussionID,
	}
}

// Execute executes the request
//  @return DiscussionPatchObjectV1Response
func (a *ObjectDiscussionAPIService) DiscussionPatchObjectV1Execute(r ApiDiscussionPatchObjectV1Request) (*DiscussionPatchObjectV1Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DiscussionPatchObjectV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectDiscussionAPIService.DiscussionPatchObjectV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/discussion/{pkiDiscussionID}"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiDiscussionID"+"}", url.PathEscape(parameterValueToString(r.pkiDiscussionID, "pkiDiscussionID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pkiDiscussionID < 0 {
		return localVarReturnValue, nil, reportError("pkiDiscussionID must be greater than 0")
	}
	if r.pkiDiscussionID > 16777215 {
		return localVarReturnValue, nil, reportError("pkiDiscussionID must be less than 16777215")
	}
	if r.discussionPatchObjectV1Request == nil {
		return localVarReturnValue, nil, reportError("discussionPatchObjectV1Request is required and must be specified")
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
	localVarPostBody = r.discussionPatchObjectV1Request
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

type ApiDiscussionUpdateDiscussionreadstatusV1Request struct {
	ctx context.Context
	ApiService *ObjectDiscussionAPIService
	pkiDiscussionID int32
	discussionUpdateDiscussionreadstatusV1Request *DiscussionUpdateDiscussionreadstatusV1Request
}

func (r ApiDiscussionUpdateDiscussionreadstatusV1Request) DiscussionUpdateDiscussionreadstatusV1Request(discussionUpdateDiscussionreadstatusV1Request DiscussionUpdateDiscussionreadstatusV1Request) ApiDiscussionUpdateDiscussionreadstatusV1Request {
	r.discussionUpdateDiscussionreadstatusV1Request = &discussionUpdateDiscussionreadstatusV1Request
	return r
}

func (r ApiDiscussionUpdateDiscussionreadstatusV1Request) Execute() (*DiscussionUpdateDiscussionreadstatusV1Response, *http.Response, error) {
	return r.ApiService.DiscussionUpdateDiscussionreadstatusV1Execute(r)
}

/*
DiscussionUpdateDiscussionreadstatusV1 Update the read status of the discussion

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pkiDiscussionID
 @return ApiDiscussionUpdateDiscussionreadstatusV1Request
*/
func (a *ObjectDiscussionAPIService) DiscussionUpdateDiscussionreadstatusV1(ctx context.Context, pkiDiscussionID int32) ApiDiscussionUpdateDiscussionreadstatusV1Request {
	return ApiDiscussionUpdateDiscussionreadstatusV1Request{
		ApiService: a,
		ctx: ctx,
		pkiDiscussionID: pkiDiscussionID,
	}
}

// Execute executes the request
//  @return DiscussionUpdateDiscussionreadstatusV1Response
func (a *ObjectDiscussionAPIService) DiscussionUpdateDiscussionreadstatusV1Execute(r ApiDiscussionUpdateDiscussionreadstatusV1Request) (*DiscussionUpdateDiscussionreadstatusV1Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DiscussionUpdateDiscussionreadstatusV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectDiscussionAPIService.DiscussionUpdateDiscussionreadstatusV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/discussion/{pkiDiscussionID}/updateDiscussionreadstatus"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiDiscussionID"+"}", url.PathEscape(parameterValueToString(r.pkiDiscussionID, "pkiDiscussionID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pkiDiscussionID < 0 {
		return localVarReturnValue, nil, reportError("pkiDiscussionID must be greater than 0")
	}
	if r.pkiDiscussionID > 16777215 {
		return localVarReturnValue, nil, reportError("pkiDiscussionID must be less than 16777215")
	}
	if r.discussionUpdateDiscussionreadstatusV1Request == nil {
		return localVarReturnValue, nil, reportError("discussionUpdateDiscussionreadstatusV1Request is required and must be specified")
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
	localVarPostBody = r.discussionUpdateDiscussionreadstatusV1Request
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
