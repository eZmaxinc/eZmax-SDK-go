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


// ObjectEzsigndiscussionAPIService ObjectEzsigndiscussionAPI service
type ObjectEzsigndiscussionAPIService service

type ApiEzsigndiscussionCreateObjectV1Request struct {
	ctx context.Context
	ApiService *ObjectEzsigndiscussionAPIService
	ezsigndiscussionCreateObjectV1Request *EzsigndiscussionCreateObjectV1Request
}

func (r ApiEzsigndiscussionCreateObjectV1Request) EzsigndiscussionCreateObjectV1Request(ezsigndiscussionCreateObjectV1Request EzsigndiscussionCreateObjectV1Request) ApiEzsigndiscussionCreateObjectV1Request {
	r.ezsigndiscussionCreateObjectV1Request = &ezsigndiscussionCreateObjectV1Request
	return r
}

func (r ApiEzsigndiscussionCreateObjectV1Request) Execute() (*EzsigndiscussionCreateObjectV1Response, *http.Response, error) {
	return r.ApiService.EzsigndiscussionCreateObjectV1Execute(r)
}

/*
EzsigndiscussionCreateObjectV1 Create a new Ezsigndiscussion

The endpoint allows to create one or many elements at once.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiEzsigndiscussionCreateObjectV1Request
*/
func (a *ObjectEzsigndiscussionAPIService) EzsigndiscussionCreateObjectV1(ctx context.Context) ApiEzsigndiscussionCreateObjectV1Request {
	return ApiEzsigndiscussionCreateObjectV1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return EzsigndiscussionCreateObjectV1Response
func (a *ObjectEzsigndiscussionAPIService) EzsigndiscussionCreateObjectV1Execute(r ApiEzsigndiscussionCreateObjectV1Request) (*EzsigndiscussionCreateObjectV1Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EzsigndiscussionCreateObjectV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectEzsigndiscussionAPIService.EzsigndiscussionCreateObjectV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/ezsigndiscussion"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.ezsigndiscussionCreateObjectV1Request == nil {
		return localVarReturnValue, nil, reportError("ezsigndiscussionCreateObjectV1Request is required and must be specified")
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
	localVarPostBody = r.ezsigndiscussionCreateObjectV1Request
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

type ApiEzsigndiscussionDeleteObjectV1Request struct {
	ctx context.Context
	ApiService *ObjectEzsigndiscussionAPIService
	pkiEzsigndiscussionID int32
}

func (r ApiEzsigndiscussionDeleteObjectV1Request) Execute() (*EzsigndiscussionDeleteObjectV1Response, *http.Response, error) {
	return r.ApiService.EzsigndiscussionDeleteObjectV1Execute(r)
}

/*
EzsigndiscussionDeleteObjectV1 Delete an existing Ezsigndiscussion



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pkiEzsigndiscussionID The unique ID of the Ezsigndiscussion
 @return ApiEzsigndiscussionDeleteObjectV1Request
*/
func (a *ObjectEzsigndiscussionAPIService) EzsigndiscussionDeleteObjectV1(ctx context.Context, pkiEzsigndiscussionID int32) ApiEzsigndiscussionDeleteObjectV1Request {
	return ApiEzsigndiscussionDeleteObjectV1Request{
		ApiService: a,
		ctx: ctx,
		pkiEzsigndiscussionID: pkiEzsigndiscussionID,
	}
}

// Execute executes the request
//  @return EzsigndiscussionDeleteObjectV1Response
func (a *ObjectEzsigndiscussionAPIService) EzsigndiscussionDeleteObjectV1Execute(r ApiEzsigndiscussionDeleteObjectV1Request) (*EzsigndiscussionDeleteObjectV1Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EzsigndiscussionDeleteObjectV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectEzsigndiscussionAPIService.EzsigndiscussionDeleteObjectV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/ezsigndiscussion/{pkiEzsigndiscussionID}"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiEzsigndiscussionID"+"}", url.PathEscape(parameterValueToString(r.pkiEzsigndiscussionID, "pkiEzsigndiscussionID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pkiEzsigndiscussionID < 0 {
		return localVarReturnValue, nil, reportError("pkiEzsigndiscussionID must be greater than 0")
	}
	if r.pkiEzsigndiscussionID > 16777215 {
		return localVarReturnValue, nil, reportError("pkiEzsigndiscussionID must be less than 16777215")
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

type ApiEzsigndiscussionGetObjectV2Request struct {
	ctx context.Context
	ApiService *ObjectEzsigndiscussionAPIService
	pkiEzsigndiscussionID int32
}

func (r ApiEzsigndiscussionGetObjectV2Request) Execute() (*EzsigndiscussionGetObjectV2Response, *http.Response, error) {
	return r.ApiService.EzsigndiscussionGetObjectV2Execute(r)
}

/*
EzsigndiscussionGetObjectV2 Retrieve an existing Ezsigndiscussion



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pkiEzsigndiscussionID The unique ID of the Ezsigndiscussion
 @return ApiEzsigndiscussionGetObjectV2Request
*/
func (a *ObjectEzsigndiscussionAPIService) EzsigndiscussionGetObjectV2(ctx context.Context, pkiEzsigndiscussionID int32) ApiEzsigndiscussionGetObjectV2Request {
	return ApiEzsigndiscussionGetObjectV2Request{
		ApiService: a,
		ctx: ctx,
		pkiEzsigndiscussionID: pkiEzsigndiscussionID,
	}
}

// Execute executes the request
//  @return EzsigndiscussionGetObjectV2Response
func (a *ObjectEzsigndiscussionAPIService) EzsigndiscussionGetObjectV2Execute(r ApiEzsigndiscussionGetObjectV2Request) (*EzsigndiscussionGetObjectV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EzsigndiscussionGetObjectV2Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectEzsigndiscussionAPIService.EzsigndiscussionGetObjectV2")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/2/object/ezsigndiscussion/{pkiEzsigndiscussionID}"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiEzsigndiscussionID"+"}", url.PathEscape(parameterValueToString(r.pkiEzsigndiscussionID, "pkiEzsigndiscussionID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pkiEzsigndiscussionID < 0 {
		return localVarReturnValue, nil, reportError("pkiEzsigndiscussionID must be greater than 0")
	}
	if r.pkiEzsigndiscussionID > 16777215 {
		return localVarReturnValue, nil, reportError("pkiEzsigndiscussionID must be less than 16777215")
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