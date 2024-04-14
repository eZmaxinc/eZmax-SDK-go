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


// ObjectEzsigntemplatesignatureAPIService ObjectEzsigntemplatesignatureAPI service
type ObjectEzsigntemplatesignatureAPIService service

type ApiEzsigntemplatesignatureCreateObjectV1Request struct {
	ctx context.Context
	ApiService *ObjectEzsigntemplatesignatureAPIService
	ezsigntemplatesignatureCreateObjectV1Request *EzsigntemplatesignatureCreateObjectV1Request
}

func (r ApiEzsigntemplatesignatureCreateObjectV1Request) EzsigntemplatesignatureCreateObjectV1Request(ezsigntemplatesignatureCreateObjectV1Request EzsigntemplatesignatureCreateObjectV1Request) ApiEzsigntemplatesignatureCreateObjectV1Request {
	r.ezsigntemplatesignatureCreateObjectV1Request = &ezsigntemplatesignatureCreateObjectV1Request
	return r
}

func (r ApiEzsigntemplatesignatureCreateObjectV1Request) Execute() (*EzsigntemplatesignatureCreateObjectV1Response, *http.Response, error) {
	return r.ApiService.EzsigntemplatesignatureCreateObjectV1Execute(r)
}

/*
EzsigntemplatesignatureCreateObjectV1 Create a new Ezsigntemplatesignature

The endpoint allows to create one or many elements at once.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiEzsigntemplatesignatureCreateObjectV1Request
*/
func (a *ObjectEzsigntemplatesignatureAPIService) EzsigntemplatesignatureCreateObjectV1(ctx context.Context) ApiEzsigntemplatesignatureCreateObjectV1Request {
	return ApiEzsigntemplatesignatureCreateObjectV1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return EzsigntemplatesignatureCreateObjectV1Response
func (a *ObjectEzsigntemplatesignatureAPIService) EzsigntemplatesignatureCreateObjectV1Execute(r ApiEzsigntemplatesignatureCreateObjectV1Request) (*EzsigntemplatesignatureCreateObjectV1Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EzsigntemplatesignatureCreateObjectV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectEzsigntemplatesignatureAPIService.EzsigntemplatesignatureCreateObjectV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/ezsigntemplatesignature"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.ezsigntemplatesignatureCreateObjectV1Request == nil {
		return localVarReturnValue, nil, reportError("ezsigntemplatesignatureCreateObjectV1Request is required and must be specified")
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
	localVarPostBody = r.ezsigntemplatesignatureCreateObjectV1Request
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

type ApiEzsigntemplatesignatureDeleteObjectV1Request struct {
	ctx context.Context
	ApiService *ObjectEzsigntemplatesignatureAPIService
	pkiEzsigntemplatesignatureID int32
}

func (r ApiEzsigntemplatesignatureDeleteObjectV1Request) Execute() (*EzsigntemplatesignatureDeleteObjectV1Response, *http.Response, error) {
	return r.ApiService.EzsigntemplatesignatureDeleteObjectV1Execute(r)
}

/*
EzsigntemplatesignatureDeleteObjectV1 Delete an existing Ezsigntemplatesignature



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pkiEzsigntemplatesignatureID
 @return ApiEzsigntemplatesignatureDeleteObjectV1Request
*/
func (a *ObjectEzsigntemplatesignatureAPIService) EzsigntemplatesignatureDeleteObjectV1(ctx context.Context, pkiEzsigntemplatesignatureID int32) ApiEzsigntemplatesignatureDeleteObjectV1Request {
	return ApiEzsigntemplatesignatureDeleteObjectV1Request{
		ApiService: a,
		ctx: ctx,
		pkiEzsigntemplatesignatureID: pkiEzsigntemplatesignatureID,
	}
}

// Execute executes the request
//  @return EzsigntemplatesignatureDeleteObjectV1Response
func (a *ObjectEzsigntemplatesignatureAPIService) EzsigntemplatesignatureDeleteObjectV1Execute(r ApiEzsigntemplatesignatureDeleteObjectV1Request) (*EzsigntemplatesignatureDeleteObjectV1Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EzsigntemplatesignatureDeleteObjectV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectEzsigntemplatesignatureAPIService.EzsigntemplatesignatureDeleteObjectV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/ezsigntemplatesignature/{pkiEzsigntemplatesignatureID}"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiEzsigntemplatesignatureID"+"}", url.PathEscape(parameterValueToString(r.pkiEzsigntemplatesignatureID, "pkiEzsigntemplatesignatureID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pkiEzsigntemplatesignatureID < 0 {
		return localVarReturnValue, nil, reportError("pkiEzsigntemplatesignatureID must be greater than 0")
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

type ApiEzsigntemplatesignatureEditObjectV1Request struct {
	ctx context.Context
	ApiService *ObjectEzsigntemplatesignatureAPIService
	pkiEzsigntemplatesignatureID int32
	ezsigntemplatesignatureEditObjectV1Request *EzsigntemplatesignatureEditObjectV1Request
}

func (r ApiEzsigntemplatesignatureEditObjectV1Request) EzsigntemplatesignatureEditObjectV1Request(ezsigntemplatesignatureEditObjectV1Request EzsigntemplatesignatureEditObjectV1Request) ApiEzsigntemplatesignatureEditObjectV1Request {
	r.ezsigntemplatesignatureEditObjectV1Request = &ezsigntemplatesignatureEditObjectV1Request
	return r
}

func (r ApiEzsigntemplatesignatureEditObjectV1Request) Execute() (*EzsigntemplatesignatureEditObjectV1Response, *http.Response, error) {
	return r.ApiService.EzsigntemplatesignatureEditObjectV1Execute(r)
}

/*
EzsigntemplatesignatureEditObjectV1 Edit an existing Ezsigntemplatesignature



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pkiEzsigntemplatesignatureID
 @return ApiEzsigntemplatesignatureEditObjectV1Request
*/
func (a *ObjectEzsigntemplatesignatureAPIService) EzsigntemplatesignatureEditObjectV1(ctx context.Context, pkiEzsigntemplatesignatureID int32) ApiEzsigntemplatesignatureEditObjectV1Request {
	return ApiEzsigntemplatesignatureEditObjectV1Request{
		ApiService: a,
		ctx: ctx,
		pkiEzsigntemplatesignatureID: pkiEzsigntemplatesignatureID,
	}
}

// Execute executes the request
//  @return EzsigntemplatesignatureEditObjectV1Response
func (a *ObjectEzsigntemplatesignatureAPIService) EzsigntemplatesignatureEditObjectV1Execute(r ApiEzsigntemplatesignatureEditObjectV1Request) (*EzsigntemplatesignatureEditObjectV1Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EzsigntemplatesignatureEditObjectV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectEzsigntemplatesignatureAPIService.EzsigntemplatesignatureEditObjectV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/ezsigntemplatesignature/{pkiEzsigntemplatesignatureID}"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiEzsigntemplatesignatureID"+"}", url.PathEscape(parameterValueToString(r.pkiEzsigntemplatesignatureID, "pkiEzsigntemplatesignatureID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pkiEzsigntemplatesignatureID < 0 {
		return localVarReturnValue, nil, reportError("pkiEzsigntemplatesignatureID must be greater than 0")
	}
	if r.ezsigntemplatesignatureEditObjectV1Request == nil {
		return localVarReturnValue, nil, reportError("ezsigntemplatesignatureEditObjectV1Request is required and must be specified")
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
	localVarPostBody = r.ezsigntemplatesignatureEditObjectV1Request
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

type ApiEzsigntemplatesignatureGetObjectV2Request struct {
	ctx context.Context
	ApiService *ObjectEzsigntemplatesignatureAPIService
	pkiEzsigntemplatesignatureID int32
}

func (r ApiEzsigntemplatesignatureGetObjectV2Request) Execute() (*EzsigntemplatesignatureGetObjectV2Response, *http.Response, error) {
	return r.ApiService.EzsigntemplatesignatureGetObjectV2Execute(r)
}

/*
EzsigntemplatesignatureGetObjectV2 Retrieve an existing Ezsigntemplatesignature



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pkiEzsigntemplatesignatureID
 @return ApiEzsigntemplatesignatureGetObjectV2Request
*/
func (a *ObjectEzsigntemplatesignatureAPIService) EzsigntemplatesignatureGetObjectV2(ctx context.Context, pkiEzsigntemplatesignatureID int32) ApiEzsigntemplatesignatureGetObjectV2Request {
	return ApiEzsigntemplatesignatureGetObjectV2Request{
		ApiService: a,
		ctx: ctx,
		pkiEzsigntemplatesignatureID: pkiEzsigntemplatesignatureID,
	}
}

// Execute executes the request
//  @return EzsigntemplatesignatureGetObjectV2Response
func (a *ObjectEzsigntemplatesignatureAPIService) EzsigntemplatesignatureGetObjectV2Execute(r ApiEzsigntemplatesignatureGetObjectV2Request) (*EzsigntemplatesignatureGetObjectV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EzsigntemplatesignatureGetObjectV2Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectEzsigntemplatesignatureAPIService.EzsigntemplatesignatureGetObjectV2")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/2/object/ezsigntemplatesignature/{pkiEzsigntemplatesignatureID}"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiEzsigntemplatesignatureID"+"}", url.PathEscape(parameterValueToString(r.pkiEzsigntemplatesignatureID, "pkiEzsigntemplatesignatureID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pkiEzsigntemplatesignatureID < 0 {
		return localVarReturnValue, nil, reportError("pkiEzsigntemplatesignatureID must be greater than 0")
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
