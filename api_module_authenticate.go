/*
eZmax API Definition

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.1.4
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// ModuleAuthenticateApiService ModuleAuthenticateApi service
type ModuleAuthenticateApiService service

type ApiAuthenticateAuthenticateV2Request struct {
	ctx _context.Context
	ApiService *ModuleAuthenticateApiService
	eSessionType string
	authenticateAuthenticateV2Request *AuthenticateAuthenticateV2Request
}

func (r ApiAuthenticateAuthenticateV2Request) AuthenticateAuthenticateV2Request(authenticateAuthenticateV2Request AuthenticateAuthenticateV2Request) ApiAuthenticateAuthenticateV2Request {
	r.authenticateAuthenticateV2Request = &authenticateAuthenticateV2Request
	return r
}

func (r ApiAuthenticateAuthenticateV2Request) Execute() (AuthenticateAuthenticateV2Response, *_nethttp.Response, error) {
	return r.ApiService.AuthenticateAuthenticateV2Execute(r)
}

/*
AuthenticateAuthenticateV2 Authenticate a user

This endpoint authenticates a user.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param eSessionType
 @return ApiAuthenticateAuthenticateV2Request
*/
func (a *ModuleAuthenticateApiService) AuthenticateAuthenticateV2(ctx _context.Context, eSessionType string) ApiAuthenticateAuthenticateV2Request {
	return ApiAuthenticateAuthenticateV2Request{
		ApiService: a,
		ctx: ctx,
		eSessionType: eSessionType,
	}
}

// Execute executes the request
//  @return AuthenticateAuthenticateV2Response
func (a *ModuleAuthenticateApiService) AuthenticateAuthenticateV2Execute(r ApiAuthenticateAuthenticateV2Request) (AuthenticateAuthenticateV2Response, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  AuthenticateAuthenticateV2Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ModuleAuthenticateApiService.AuthenticateAuthenticateV2")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/2/module/authenticate/authenticate/{eSessionType}"
	localVarPath = strings.Replace(localVarPath, "{"+"eSessionType"+"}", _neturl.PathEscape(parameterToString(r.eSessionType, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.authenticateAuthenticateV2Request == nil {
		return localVarReturnValue, nil, reportError("authenticateAuthenticateV2Request is required and must be specified")
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
	localVarPostBody = r.authenticateAuthenticateV2Request
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

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v CommonResponseError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
