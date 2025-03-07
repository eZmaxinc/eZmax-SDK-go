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


// ObjectEzsignuserAPIService ObjectEzsignuserAPI service
type ObjectEzsignuserAPIService service

type ApiEzsignuserEditObjectV1Request struct {
	ctx context.Context
	ApiService *ObjectEzsignuserAPIService
	pkiEzsignuserID int32
	ezsignuserEditObjectV1Request *EzsignuserEditObjectV1Request
}

func (r ApiEzsignuserEditObjectV1Request) EzsignuserEditObjectV1Request(ezsignuserEditObjectV1Request EzsignuserEditObjectV1Request) ApiEzsignuserEditObjectV1Request {
	r.ezsignuserEditObjectV1Request = &ezsignuserEditObjectV1Request
	return r
}

func (r ApiEzsignuserEditObjectV1Request) Execute() (*EzsignuserEditObjectV1Response, *http.Response, error) {
	return r.ApiService.EzsignuserEditObjectV1Execute(r)
}

/*
EzsignuserEditObjectV1 Edit an existing Ezsignuser



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pkiEzsignuserID The unique ID of the Ezsignuser
 @return ApiEzsignuserEditObjectV1Request
*/
func (a *ObjectEzsignuserAPIService) EzsignuserEditObjectV1(ctx context.Context, pkiEzsignuserID int32) ApiEzsignuserEditObjectV1Request {
	return ApiEzsignuserEditObjectV1Request{
		ApiService: a,
		ctx: ctx,
		pkiEzsignuserID: pkiEzsignuserID,
	}
}

// Execute executes the request
//  @return EzsignuserEditObjectV1Response
func (a *ObjectEzsignuserAPIService) EzsignuserEditObjectV1Execute(r ApiEzsignuserEditObjectV1Request) (*EzsignuserEditObjectV1Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EzsignuserEditObjectV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectEzsignuserAPIService.EzsignuserEditObjectV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/ezsignuser/{pkiEzsignuserID}"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiEzsignuserID"+"}", url.PathEscape(parameterValueToString(r.pkiEzsignuserID, "pkiEzsignuserID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pkiEzsignuserID < 0 {
		return localVarReturnValue, nil, reportError("pkiEzsignuserID must be greater than 0")
	}
	if r.pkiEzsignuserID > 65535 {
		return localVarReturnValue, nil, reportError("pkiEzsignuserID must be less than 65535")
	}
	if r.ezsignuserEditObjectV1Request == nil {
		return localVarReturnValue, nil, reportError("ezsignuserEditObjectV1Request is required and must be specified")
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
	localVarPostBody = r.ezsignuserEditObjectV1Request
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

type ApiEzsignuserGetObjectV2Request struct {
	ctx context.Context
	ApiService *ObjectEzsignuserAPIService
	pkiEzsignuserID int32
}

func (r ApiEzsignuserGetObjectV2Request) Execute() (*EzsignuserGetObjectV2Response, *http.Response, error) {
	return r.ApiService.EzsignuserGetObjectV2Execute(r)
}

/*
EzsignuserGetObjectV2 Retrieve an existing Ezsignuser



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pkiEzsignuserID The unique ID of the Ezsignuser
 @return ApiEzsignuserGetObjectV2Request
*/
func (a *ObjectEzsignuserAPIService) EzsignuserGetObjectV2(ctx context.Context, pkiEzsignuserID int32) ApiEzsignuserGetObjectV2Request {
	return ApiEzsignuserGetObjectV2Request{
		ApiService: a,
		ctx: ctx,
		pkiEzsignuserID: pkiEzsignuserID,
	}
}

// Execute executes the request
//  @return EzsignuserGetObjectV2Response
func (a *ObjectEzsignuserAPIService) EzsignuserGetObjectV2Execute(r ApiEzsignuserGetObjectV2Request) (*EzsignuserGetObjectV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EzsignuserGetObjectV2Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectEzsignuserAPIService.EzsignuserGetObjectV2")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/2/object/ezsignuser/{pkiEzsignuserID}"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiEzsignuserID"+"}", url.PathEscape(parameterValueToString(r.pkiEzsignuserID, "pkiEzsignuserID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pkiEzsignuserID < 0 {
		return localVarReturnValue, nil, reportError("pkiEzsignuserID must be greater than 0")
	}
	if r.pkiEzsignuserID > 65535 {
		return localVarReturnValue, nil, reportError("pkiEzsignuserID must be less than 65535")
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
