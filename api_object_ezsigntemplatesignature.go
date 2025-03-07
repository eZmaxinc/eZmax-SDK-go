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


// ObjectEzsigntemplatesignatureAPIService ObjectEzsigntemplatesignatureAPI service
type ObjectEzsigntemplatesignatureAPIService service

type ApiEzsigntemplatesignatureCreateObjectV2Request struct {
	ctx context.Context
	ApiService *ObjectEzsigntemplatesignatureAPIService
	ezsigntemplatesignatureCreateObjectV2Request *EzsigntemplatesignatureCreateObjectV2Request
}

func (r ApiEzsigntemplatesignatureCreateObjectV2Request) EzsigntemplatesignatureCreateObjectV2Request(ezsigntemplatesignatureCreateObjectV2Request EzsigntemplatesignatureCreateObjectV2Request) ApiEzsigntemplatesignatureCreateObjectV2Request {
	r.ezsigntemplatesignatureCreateObjectV2Request = &ezsigntemplatesignatureCreateObjectV2Request
	return r
}

func (r ApiEzsigntemplatesignatureCreateObjectV2Request) Execute() (*EzsigntemplatesignatureCreateObjectV2Response, *http.Response, error) {
	return r.ApiService.EzsigntemplatesignatureCreateObjectV2Execute(r)
}

/*
EzsigntemplatesignatureCreateObjectV2 Create a new Ezsigntemplatesignature

The endpoint allows to create one or many elements at once.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiEzsigntemplatesignatureCreateObjectV2Request
*/
func (a *ObjectEzsigntemplatesignatureAPIService) EzsigntemplatesignatureCreateObjectV2(ctx context.Context) ApiEzsigntemplatesignatureCreateObjectV2Request {
	return ApiEzsigntemplatesignatureCreateObjectV2Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return EzsigntemplatesignatureCreateObjectV2Response
func (a *ObjectEzsigntemplatesignatureAPIService) EzsigntemplatesignatureCreateObjectV2Execute(r ApiEzsigntemplatesignatureCreateObjectV2Request) (*EzsigntemplatesignatureCreateObjectV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EzsigntemplatesignatureCreateObjectV2Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectEzsigntemplatesignatureAPIService.EzsigntemplatesignatureCreateObjectV2")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/2/object/ezsigntemplatesignature"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.ezsigntemplatesignatureCreateObjectV2Request == nil {
		return localVarReturnValue, nil, reportError("ezsigntemplatesignatureCreateObjectV2Request is required and must be specified")
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
	localVarPostBody = r.ezsigntemplatesignatureCreateObjectV2Request
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

type ApiEzsigntemplatesignatureEditObjectV2Request struct {
	ctx context.Context
	ApiService *ObjectEzsigntemplatesignatureAPIService
	pkiEzsigntemplatesignatureID int32
	ezsigntemplatesignatureEditObjectV2Request *EzsigntemplatesignatureEditObjectV2Request
}

func (r ApiEzsigntemplatesignatureEditObjectV2Request) EzsigntemplatesignatureEditObjectV2Request(ezsigntemplatesignatureEditObjectV2Request EzsigntemplatesignatureEditObjectV2Request) ApiEzsigntemplatesignatureEditObjectV2Request {
	r.ezsigntemplatesignatureEditObjectV2Request = &ezsigntemplatesignatureEditObjectV2Request
	return r
}

func (r ApiEzsigntemplatesignatureEditObjectV2Request) Execute() (*EzsigntemplatesignatureEditObjectV2Response, *http.Response, error) {
	return r.ApiService.EzsigntemplatesignatureEditObjectV2Execute(r)
}

/*
EzsigntemplatesignatureEditObjectV2 Edit an existing Ezsigntemplatesignature



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pkiEzsigntemplatesignatureID
 @return ApiEzsigntemplatesignatureEditObjectV2Request
*/
func (a *ObjectEzsigntemplatesignatureAPIService) EzsigntemplatesignatureEditObjectV2(ctx context.Context, pkiEzsigntemplatesignatureID int32) ApiEzsigntemplatesignatureEditObjectV2Request {
	return ApiEzsigntemplatesignatureEditObjectV2Request{
		ApiService: a,
		ctx: ctx,
		pkiEzsigntemplatesignatureID: pkiEzsigntemplatesignatureID,
	}
}

// Execute executes the request
//  @return EzsigntemplatesignatureEditObjectV2Response
func (a *ObjectEzsigntemplatesignatureAPIService) EzsigntemplatesignatureEditObjectV2Execute(r ApiEzsigntemplatesignatureEditObjectV2Request) (*EzsigntemplatesignatureEditObjectV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EzsigntemplatesignatureEditObjectV2Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectEzsigntemplatesignatureAPIService.EzsigntemplatesignatureEditObjectV2")
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
	if r.ezsigntemplatesignatureEditObjectV2Request == nil {
		return localVarReturnValue, nil, reportError("ezsigntemplatesignatureEditObjectV2Request is required and must be specified")
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
	localVarPostBody = r.ezsigntemplatesignatureEditObjectV2Request
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

type ApiEzsigntemplatesignatureGetObjectV3Request struct {
	ctx context.Context
	ApiService *ObjectEzsigntemplatesignatureAPIService
	pkiEzsigntemplatesignatureID int32
}

func (r ApiEzsigntemplatesignatureGetObjectV3Request) Execute() (*EzsigntemplatesignatureGetObjectV3Response, *http.Response, error) {
	return r.ApiService.EzsigntemplatesignatureGetObjectV3Execute(r)
}

/*
EzsigntemplatesignatureGetObjectV3 Retrieve an existing Ezsigntemplatesignature



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pkiEzsigntemplatesignatureID
 @return ApiEzsigntemplatesignatureGetObjectV3Request
*/
func (a *ObjectEzsigntemplatesignatureAPIService) EzsigntemplatesignatureGetObjectV3(ctx context.Context, pkiEzsigntemplatesignatureID int32) ApiEzsigntemplatesignatureGetObjectV3Request {
	return ApiEzsigntemplatesignatureGetObjectV3Request{
		ApiService: a,
		ctx: ctx,
		pkiEzsigntemplatesignatureID: pkiEzsigntemplatesignatureID,
	}
}

// Execute executes the request
//  @return EzsigntemplatesignatureGetObjectV3Response
func (a *ObjectEzsigntemplatesignatureAPIService) EzsigntemplatesignatureGetObjectV3Execute(r ApiEzsigntemplatesignatureGetObjectV3Request) (*EzsigntemplatesignatureGetObjectV3Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EzsigntemplatesignatureGetObjectV3Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectEzsigntemplatesignatureAPIService.EzsigntemplatesignatureGetObjectV3")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/3/object/ezsigntemplatesignature/{pkiEzsigntemplatesignatureID}"
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
