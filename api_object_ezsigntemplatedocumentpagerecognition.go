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


// ObjectEzsigntemplatedocumentpagerecognitionAPIService ObjectEzsigntemplatedocumentpagerecognitionAPI service
type ObjectEzsigntemplatedocumentpagerecognitionAPIService service

type ApiEzsigntemplatedocumentpagerecognitionCreateObjectV1Request struct {
	ctx context.Context
	ApiService *ObjectEzsigntemplatedocumentpagerecognitionAPIService
	ezsigntemplatedocumentpagerecognitionCreateObjectV1Request *EzsigntemplatedocumentpagerecognitionCreateObjectV1Request
}

func (r ApiEzsigntemplatedocumentpagerecognitionCreateObjectV1Request) EzsigntemplatedocumentpagerecognitionCreateObjectV1Request(ezsigntemplatedocumentpagerecognitionCreateObjectV1Request EzsigntemplatedocumentpagerecognitionCreateObjectV1Request) ApiEzsigntemplatedocumentpagerecognitionCreateObjectV1Request {
	r.ezsigntemplatedocumentpagerecognitionCreateObjectV1Request = &ezsigntemplatedocumentpagerecognitionCreateObjectV1Request
	return r
}

func (r ApiEzsigntemplatedocumentpagerecognitionCreateObjectV1Request) Execute() (*EzsigntemplatedocumentpagerecognitionCreateObjectV1Response, *http.Response, error) {
	return r.ApiService.EzsigntemplatedocumentpagerecognitionCreateObjectV1Execute(r)
}

/*
EzsigntemplatedocumentpagerecognitionCreateObjectV1 Create a new Ezsigntemplatedocumentpagerecognition

The endpoint allows to create one or many elements at once.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiEzsigntemplatedocumentpagerecognitionCreateObjectV1Request
*/
func (a *ObjectEzsigntemplatedocumentpagerecognitionAPIService) EzsigntemplatedocumentpagerecognitionCreateObjectV1(ctx context.Context) ApiEzsigntemplatedocumentpagerecognitionCreateObjectV1Request {
	return ApiEzsigntemplatedocumentpagerecognitionCreateObjectV1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return EzsigntemplatedocumentpagerecognitionCreateObjectV1Response
func (a *ObjectEzsigntemplatedocumentpagerecognitionAPIService) EzsigntemplatedocumentpagerecognitionCreateObjectV1Execute(r ApiEzsigntemplatedocumentpagerecognitionCreateObjectV1Request) (*EzsigntemplatedocumentpagerecognitionCreateObjectV1Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EzsigntemplatedocumentpagerecognitionCreateObjectV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectEzsigntemplatedocumentpagerecognitionAPIService.EzsigntemplatedocumentpagerecognitionCreateObjectV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/ezsigntemplatedocumentpagerecognition"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.ezsigntemplatedocumentpagerecognitionCreateObjectV1Request == nil {
		return localVarReturnValue, nil, reportError("ezsigntemplatedocumentpagerecognitionCreateObjectV1Request is required and must be specified")
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
	localVarPostBody = r.ezsigntemplatedocumentpagerecognitionCreateObjectV1Request
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

type ApiEzsigntemplatedocumentpagerecognitionDeleteObjectV1Request struct {
	ctx context.Context
	ApiService *ObjectEzsigntemplatedocumentpagerecognitionAPIService
	pkiEzsigntemplatedocumentpagerecognitionID int32
}

func (r ApiEzsigntemplatedocumentpagerecognitionDeleteObjectV1Request) Execute() (*EzsigntemplatedocumentpagerecognitionDeleteObjectV1Response, *http.Response, error) {
	return r.ApiService.EzsigntemplatedocumentpagerecognitionDeleteObjectV1Execute(r)
}

/*
EzsigntemplatedocumentpagerecognitionDeleteObjectV1 Delete an existing Ezsigntemplatedocumentpagerecognition



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pkiEzsigntemplatedocumentpagerecognitionID The unique ID of the Ezsigntemplatedocumentpagerecognition
 @return ApiEzsigntemplatedocumentpagerecognitionDeleteObjectV1Request
*/
func (a *ObjectEzsigntemplatedocumentpagerecognitionAPIService) EzsigntemplatedocumentpagerecognitionDeleteObjectV1(ctx context.Context, pkiEzsigntemplatedocumentpagerecognitionID int32) ApiEzsigntemplatedocumentpagerecognitionDeleteObjectV1Request {
	return ApiEzsigntemplatedocumentpagerecognitionDeleteObjectV1Request{
		ApiService: a,
		ctx: ctx,
		pkiEzsigntemplatedocumentpagerecognitionID: pkiEzsigntemplatedocumentpagerecognitionID,
	}
}

// Execute executes the request
//  @return EzsigntemplatedocumentpagerecognitionDeleteObjectV1Response
func (a *ObjectEzsigntemplatedocumentpagerecognitionAPIService) EzsigntemplatedocumentpagerecognitionDeleteObjectV1Execute(r ApiEzsigntemplatedocumentpagerecognitionDeleteObjectV1Request) (*EzsigntemplatedocumentpagerecognitionDeleteObjectV1Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EzsigntemplatedocumentpagerecognitionDeleteObjectV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectEzsigntemplatedocumentpagerecognitionAPIService.EzsigntemplatedocumentpagerecognitionDeleteObjectV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/ezsigntemplatedocumentpagerecognition/{pkiEzsigntemplatedocumentpagerecognitionID}"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiEzsigntemplatedocumentpagerecognitionID"+"}", url.PathEscape(parameterValueToString(r.pkiEzsigntemplatedocumentpagerecognitionID, "pkiEzsigntemplatedocumentpagerecognitionID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pkiEzsigntemplatedocumentpagerecognitionID < 0 {
		return localVarReturnValue, nil, reportError("pkiEzsigntemplatedocumentpagerecognitionID must be greater than 0")
	}
	if r.pkiEzsigntemplatedocumentpagerecognitionID > 65535 {
		return localVarReturnValue, nil, reportError("pkiEzsigntemplatedocumentpagerecognitionID must be less than 65535")
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

type ApiEzsigntemplatedocumentpagerecognitionEditObjectV1Request struct {
	ctx context.Context
	ApiService *ObjectEzsigntemplatedocumentpagerecognitionAPIService
	pkiEzsigntemplatedocumentpagerecognitionID int32
	ezsigntemplatedocumentpagerecognitionEditObjectV1Request *EzsigntemplatedocumentpagerecognitionEditObjectV1Request
}

func (r ApiEzsigntemplatedocumentpagerecognitionEditObjectV1Request) EzsigntemplatedocumentpagerecognitionEditObjectV1Request(ezsigntemplatedocumentpagerecognitionEditObjectV1Request EzsigntemplatedocumentpagerecognitionEditObjectV1Request) ApiEzsigntemplatedocumentpagerecognitionEditObjectV1Request {
	r.ezsigntemplatedocumentpagerecognitionEditObjectV1Request = &ezsigntemplatedocumentpagerecognitionEditObjectV1Request
	return r
}

func (r ApiEzsigntemplatedocumentpagerecognitionEditObjectV1Request) Execute() (*EzsigntemplatedocumentpagerecognitionEditObjectV1Response, *http.Response, error) {
	return r.ApiService.EzsigntemplatedocumentpagerecognitionEditObjectV1Execute(r)
}

/*
EzsigntemplatedocumentpagerecognitionEditObjectV1 Edit an existing Ezsigntemplatedocumentpagerecognition



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pkiEzsigntemplatedocumentpagerecognitionID The unique ID of the Ezsigntemplatedocumentpagerecognition
 @return ApiEzsigntemplatedocumentpagerecognitionEditObjectV1Request
*/
func (a *ObjectEzsigntemplatedocumentpagerecognitionAPIService) EzsigntemplatedocumentpagerecognitionEditObjectV1(ctx context.Context, pkiEzsigntemplatedocumentpagerecognitionID int32) ApiEzsigntemplatedocumentpagerecognitionEditObjectV1Request {
	return ApiEzsigntemplatedocumentpagerecognitionEditObjectV1Request{
		ApiService: a,
		ctx: ctx,
		pkiEzsigntemplatedocumentpagerecognitionID: pkiEzsigntemplatedocumentpagerecognitionID,
	}
}

// Execute executes the request
//  @return EzsigntemplatedocumentpagerecognitionEditObjectV1Response
func (a *ObjectEzsigntemplatedocumentpagerecognitionAPIService) EzsigntemplatedocumentpagerecognitionEditObjectV1Execute(r ApiEzsigntemplatedocumentpagerecognitionEditObjectV1Request) (*EzsigntemplatedocumentpagerecognitionEditObjectV1Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EzsigntemplatedocumentpagerecognitionEditObjectV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectEzsigntemplatedocumentpagerecognitionAPIService.EzsigntemplatedocumentpagerecognitionEditObjectV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/ezsigntemplatedocumentpagerecognition/{pkiEzsigntemplatedocumentpagerecognitionID}"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiEzsigntemplatedocumentpagerecognitionID"+"}", url.PathEscape(parameterValueToString(r.pkiEzsigntemplatedocumentpagerecognitionID, "pkiEzsigntemplatedocumentpagerecognitionID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pkiEzsigntemplatedocumentpagerecognitionID < 0 {
		return localVarReturnValue, nil, reportError("pkiEzsigntemplatedocumentpagerecognitionID must be greater than 0")
	}
	if r.pkiEzsigntemplatedocumentpagerecognitionID > 65535 {
		return localVarReturnValue, nil, reportError("pkiEzsigntemplatedocumentpagerecognitionID must be less than 65535")
	}
	if r.ezsigntemplatedocumentpagerecognitionEditObjectV1Request == nil {
		return localVarReturnValue, nil, reportError("ezsigntemplatedocumentpagerecognitionEditObjectV1Request is required and must be specified")
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
	localVarPostBody = r.ezsigntemplatedocumentpagerecognitionEditObjectV1Request
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

type ApiEzsigntemplatedocumentpagerecognitionGetObjectV2Request struct {
	ctx context.Context
	ApiService *ObjectEzsigntemplatedocumentpagerecognitionAPIService
	pkiEzsigntemplatedocumentpagerecognitionID int32
}

func (r ApiEzsigntemplatedocumentpagerecognitionGetObjectV2Request) Execute() (*EzsigntemplatedocumentpagerecognitionGetObjectV2Response, *http.Response, error) {
	return r.ApiService.EzsigntemplatedocumentpagerecognitionGetObjectV2Execute(r)
}

/*
EzsigntemplatedocumentpagerecognitionGetObjectV2 Retrieve an existing Ezsigntemplatedocumentpagerecognition



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pkiEzsigntemplatedocumentpagerecognitionID The unique ID of the Ezsigntemplatedocumentpagerecognition
 @return ApiEzsigntemplatedocumentpagerecognitionGetObjectV2Request
*/
func (a *ObjectEzsigntemplatedocumentpagerecognitionAPIService) EzsigntemplatedocumentpagerecognitionGetObjectV2(ctx context.Context, pkiEzsigntemplatedocumentpagerecognitionID int32) ApiEzsigntemplatedocumentpagerecognitionGetObjectV2Request {
	return ApiEzsigntemplatedocumentpagerecognitionGetObjectV2Request{
		ApiService: a,
		ctx: ctx,
		pkiEzsigntemplatedocumentpagerecognitionID: pkiEzsigntemplatedocumentpagerecognitionID,
	}
}

// Execute executes the request
//  @return EzsigntemplatedocumentpagerecognitionGetObjectV2Response
func (a *ObjectEzsigntemplatedocumentpagerecognitionAPIService) EzsigntemplatedocumentpagerecognitionGetObjectV2Execute(r ApiEzsigntemplatedocumentpagerecognitionGetObjectV2Request) (*EzsigntemplatedocumentpagerecognitionGetObjectV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EzsigntemplatedocumentpagerecognitionGetObjectV2Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectEzsigntemplatedocumentpagerecognitionAPIService.EzsigntemplatedocumentpagerecognitionGetObjectV2")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/2/object/ezsigntemplatedocumentpagerecognition/{pkiEzsigntemplatedocumentpagerecognitionID}"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiEzsigntemplatedocumentpagerecognitionID"+"}", url.PathEscape(parameterValueToString(r.pkiEzsigntemplatedocumentpagerecognitionID, "pkiEzsigntemplatedocumentpagerecognitionID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pkiEzsigntemplatedocumentpagerecognitionID < 0 {
		return localVarReturnValue, nil, reportError("pkiEzsigntemplatedocumentpagerecognitionID must be greater than 0")
	}
	if r.pkiEzsigntemplatedocumentpagerecognitionID > 65535 {
		return localVarReturnValue, nil, reportError("pkiEzsigntemplatedocumentpagerecognitionID must be less than 65535")
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
