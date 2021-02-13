/*
 * eZmax API Definition
 *
 * This API expose all the functionnalities for the eZmax and eZsign application.
 *
 * API version: 1.0.30
 * Contact: support-api@ezmax.ca
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxinc/eZmax-SDK-go

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

// ObjectEzsignsignatureApiService ObjectEzsignsignatureApi service
type ObjectEzsignsignatureApiService service

type ApiEzsignsignatureCreateObjectV1Request struct {
	ctx _context.Context
	ApiService *ObjectEzsignsignatureApiService
	ezsignsignatureCreateObjectV1Request *[]EzsignsignatureCreateObjectV1Request
}

func (r ApiEzsignsignatureCreateObjectV1Request) EzsignsignatureCreateObjectV1Request(ezsignsignatureCreateObjectV1Request []EzsignsignatureCreateObjectV1Request) ApiEzsignsignatureCreateObjectV1Request {
	r.ezsignsignatureCreateObjectV1Request = &ezsignsignatureCreateObjectV1Request
	return r
}

func (r ApiEzsignsignatureCreateObjectV1Request) Execute() (EzsignsignatureCreateObjectV1Response, *_nethttp.Response, error) {
	return r.ApiService.EzsignsignatureCreateObjectV1Execute(r)
}

/*
 * EzsignsignatureCreateObjectV1 Create a new Ezsignsignature
 * The endpoint allows to create one or many elements at once.

The array can contain simple (Just the object) or compound (The object and its child) objects.

Creating compound elements allows to reduce the multiple requests to create all child objects.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiEzsignsignatureCreateObjectV1Request
 */
func (a *ObjectEzsignsignatureApiService) EzsignsignatureCreateObjectV1(ctx _context.Context) ApiEzsignsignatureCreateObjectV1Request {
	return ApiEzsignsignatureCreateObjectV1Request{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return EzsignsignatureCreateObjectV1Response
 */
func (a *ObjectEzsignsignatureApiService) EzsignsignatureCreateObjectV1Execute(r ApiEzsignsignatureCreateObjectV1Request) (EzsignsignatureCreateObjectV1Response, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  EzsignsignatureCreateObjectV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectEzsignsignatureApiService.EzsignsignatureCreateObjectV1")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/ezsignsignature"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.ezsignsignatureCreateObjectV1Request == nil {
		return localVarReturnValue, nil, reportError("ezsignsignatureCreateObjectV1Request is required and must be specified")
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
	localVarPostBody = r.ezsignsignatureCreateObjectV1Request
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
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

type ApiEzsignsignatureDeleteObjectV1Request struct {
	ctx _context.Context
	ApiService *ObjectEzsignsignatureApiService
	pkiEzsignsignatureID int32
}


func (r ApiEzsignsignatureDeleteObjectV1Request) Execute() (EzsignsignatureDeleteObjectV1Response, *_nethttp.Response, error) {
	return r.ApiService.EzsignsignatureDeleteObjectV1Execute(r)
}

/*
 * EzsignsignatureDeleteObjectV1 Delete an existing Ezsignsignature
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param pkiEzsignsignatureID The unique ID of the Ezsignsignature
 * @return ApiEzsignsignatureDeleteObjectV1Request
 */
func (a *ObjectEzsignsignatureApiService) EzsignsignatureDeleteObjectV1(ctx _context.Context, pkiEzsignsignatureID int32) ApiEzsignsignatureDeleteObjectV1Request {
	return ApiEzsignsignatureDeleteObjectV1Request{
		ApiService: a,
		ctx: ctx,
		pkiEzsignsignatureID: pkiEzsignsignatureID,
	}
}

/*
 * Execute executes the request
 * @return EzsignsignatureDeleteObjectV1Response
 */
func (a *ObjectEzsignsignatureApiService) EzsignsignatureDeleteObjectV1Execute(r ApiEzsignsignatureDeleteObjectV1Request) (EzsignsignatureDeleteObjectV1Response, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  EzsignsignatureDeleteObjectV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectEzsignsignatureApiService.EzsignsignatureDeleteObjectV1")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/ezsignsignature/{pkiEzsignsignatureID}"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiEzsignsignatureID"+"}", _neturl.PathEscape(parameterToString(r.pkiEzsignsignatureID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
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
		if localVarHTTPResponse.StatusCode == 404 {
			var v CommonResponseError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
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

type ApiEzsignsignatureEditObjectV1Request struct {
	ctx _context.Context
	ApiService *ObjectEzsignsignatureApiService
	pkiEzsignsignatureID int32
	ezsignsignatureEditObjectV1Request *EzsignsignatureEditObjectV1Request
}

func (r ApiEzsignsignatureEditObjectV1Request) EzsignsignatureEditObjectV1Request(ezsignsignatureEditObjectV1Request EzsignsignatureEditObjectV1Request) ApiEzsignsignatureEditObjectV1Request {
	r.ezsignsignatureEditObjectV1Request = &ezsignsignatureEditObjectV1Request
	return r
}

func (r ApiEzsignsignatureEditObjectV1Request) Execute() (EzsignsignatureEditObjectV1Response, *_nethttp.Response, error) {
	return r.ApiService.EzsignsignatureEditObjectV1Execute(r)
}

/*
 * EzsignsignatureEditObjectV1 Modify an existing Ezsignsignature
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param pkiEzsignsignatureID The unique ID of the Ezsignsignature
 * @return ApiEzsignsignatureEditObjectV1Request
 */
func (a *ObjectEzsignsignatureApiService) EzsignsignatureEditObjectV1(ctx _context.Context, pkiEzsignsignatureID int32) ApiEzsignsignatureEditObjectV1Request {
	return ApiEzsignsignatureEditObjectV1Request{
		ApiService: a,
		ctx: ctx,
		pkiEzsignsignatureID: pkiEzsignsignatureID,
	}
}

/*
 * Execute executes the request
 * @return EzsignsignatureEditObjectV1Response
 */
func (a *ObjectEzsignsignatureApiService) EzsignsignatureEditObjectV1Execute(r ApiEzsignsignatureEditObjectV1Request) (EzsignsignatureEditObjectV1Response, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  EzsignsignatureEditObjectV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectEzsignsignatureApiService.EzsignsignatureEditObjectV1")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/ezsignsignature/{pkiEzsignsignatureID}"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiEzsignsignatureID"+"}", _neturl.PathEscape(parameterToString(r.pkiEzsignsignatureID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.ezsignsignatureEditObjectV1Request == nil {
		return localVarReturnValue, nil, reportError("ezsignsignatureEditObjectV1Request is required and must be specified")
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
	localVarPostBody = r.ezsignsignatureEditObjectV1Request
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
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
		if localVarHTTPResponse.StatusCode == 404 {
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

type ApiEzsignsignatureGetChildrenV1Request struct {
	ctx _context.Context
	ApiService *ObjectEzsignsignatureApiService
	pkiEzsignsignatureID int32
}


func (r ApiEzsignsignatureGetChildrenV1Request) Execute() (*_nethttp.Response, error) {
	return r.ApiService.EzsignsignatureGetChildrenV1Execute(r)
}

/*
 * EzsignsignatureGetChildrenV1 Retrieve an existing Ezsignsignature's children IDs
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param pkiEzsignsignatureID The unique ID of the Ezsignsignature
 * @return ApiEzsignsignatureGetChildrenV1Request
 */
func (a *ObjectEzsignsignatureApiService) EzsignsignatureGetChildrenV1(ctx _context.Context, pkiEzsignsignatureID int32) ApiEzsignsignatureGetChildrenV1Request {
	return ApiEzsignsignatureGetChildrenV1Request{
		ApiService: a,
		ctx: ctx,
		pkiEzsignsignatureID: pkiEzsignsignatureID,
	}
}

/*
 * Execute executes the request
 */
func (a *ObjectEzsignsignatureApiService) EzsignsignatureGetChildrenV1Execute(r ApiEzsignsignatureGetChildrenV1Request) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectEzsignsignatureApiService.EzsignsignatureGetChildrenV1")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/ezsignsignature/{pkiEzsignsignatureID}/getChildren"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiEzsignsignatureID"+"}", _neturl.PathEscape(parameterToString(r.pkiEzsignsignatureID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v CommonResponseError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiEzsignsignatureGetObjectV1Request struct {
	ctx _context.Context
	ApiService *ObjectEzsignsignatureApiService
	pkiEzsignsignatureID int32
}


func (r ApiEzsignsignatureGetObjectV1Request) Execute() (EzsignsignatureGetObjectV1Response, *_nethttp.Response, error) {
	return r.ApiService.EzsignsignatureGetObjectV1Execute(r)
}

/*
 * EzsignsignatureGetObjectV1 Retrieve an existing Ezsignsignature
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param pkiEzsignsignatureID The unique ID of the Ezsignsignature
 * @return ApiEzsignsignatureGetObjectV1Request
 */
func (a *ObjectEzsignsignatureApiService) EzsignsignatureGetObjectV1(ctx _context.Context, pkiEzsignsignatureID int32) ApiEzsignsignatureGetObjectV1Request {
	return ApiEzsignsignatureGetObjectV1Request{
		ApiService: a,
		ctx: ctx,
		pkiEzsignsignatureID: pkiEzsignsignatureID,
	}
}

/*
 * Execute executes the request
 * @return EzsignsignatureGetObjectV1Response
 */
func (a *ObjectEzsignsignatureApiService) EzsignsignatureGetObjectV1Execute(r ApiEzsignsignatureGetObjectV1Request) (EzsignsignatureGetObjectV1Response, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  EzsignsignatureGetObjectV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectEzsignsignatureApiService.EzsignsignatureGetObjectV1")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/ezsignsignature/{pkiEzsignsignatureID}"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiEzsignsignatureID"+"}", _neturl.PathEscape(parameterToString(r.pkiEzsignsignatureID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
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
		if localVarHTTPResponse.StatusCode == 404 {
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
