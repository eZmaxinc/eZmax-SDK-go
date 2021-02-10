/*
 * eZmax API Definition
 *
 * This API expose all the functionnalities for the eZmax and eZsign application.  We provide SDKs for customers. They are generated using OpenAPI codegen, we encourage customers to use them as we also provide samples for them.  You can choose to build your own implementation manually or can use any compatible OpenAPI 3.0 generator like Swagger Codegen, OpenAPI codegen or any commercial generators.  If you need helping understanding how to use this API, don't waste too much time looking for it. Contact support-api@ezmax.ca, we're here to help. We are developpers so we know programmers don't like bad documentation. If you don't find what you need in the documentation, let us know, we'll improve it and put you rapidly up on track.
 *
 * API version: 1.0.29
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

// ObjectEzsigndocumentApiService ObjectEzsigndocumentApi service
type ObjectEzsigndocumentApiService service

type ApiEzsigndocumentApplyEzsigntemplateV1Request struct {
	ctx _context.Context
	ApiService *ObjectEzsigndocumentApiService
	pkiEzsigndocumentID int32
	ezsigndocumentApplyEzsigntemplateV1Request *EzsigndocumentApplyEzsigntemplateV1Request
}

func (r ApiEzsigndocumentApplyEzsigntemplateV1Request) EzsigndocumentApplyEzsigntemplateV1Request(ezsigndocumentApplyEzsigntemplateV1Request EzsigndocumentApplyEzsigntemplateV1Request) ApiEzsigndocumentApplyEzsigntemplateV1Request {
	r.ezsigndocumentApplyEzsigntemplateV1Request = &ezsigndocumentApplyEzsigntemplateV1Request
	return r
}

func (r ApiEzsigndocumentApplyEzsigntemplateV1Request) Execute() (EzsigndocumentApplyEzsigntemplateV1Response, *_nethttp.Response, error) {
	return r.ApiService.EzsigndocumentApplyEzsigntemplateV1Execute(r)
}

/*
 * EzsigndocumentApplyEzsigntemplateV1 Apply an Ezsign Template to the Ezsigndocument.
 * This endpoint applies a predefined template to the ezsign document.
This allows to automatically apply all the form and signature fields on a document in a single step.

The document must not already have fields otherwise an error will be returned.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param pkiEzsigndocumentID The unique ID of the Ezsigndocument
 * @return ApiEzsigndocumentApplyEzsigntemplateV1Request
 */
func (a *ObjectEzsigndocumentApiService) EzsigndocumentApplyEzsigntemplateV1(ctx _context.Context, pkiEzsigndocumentID int32) ApiEzsigndocumentApplyEzsigntemplateV1Request {
	return ApiEzsigndocumentApplyEzsigntemplateV1Request{
		ApiService: a,
		ctx: ctx,
		pkiEzsigndocumentID: pkiEzsigndocumentID,
	}
}

/*
 * Execute executes the request
 * @return EzsigndocumentApplyEzsigntemplateV1Response
 */
func (a *ObjectEzsigndocumentApiService) EzsigndocumentApplyEzsigntemplateV1Execute(r ApiEzsigndocumentApplyEzsigntemplateV1Request) (EzsigndocumentApplyEzsigntemplateV1Response, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  EzsigndocumentApplyEzsigntemplateV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectEzsigndocumentApiService.EzsigndocumentApplyEzsigntemplateV1")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/ezsigndocument/{pkiEzsigndocumentID}/applyezsigntemplate"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiEzsigndocumentID"+"}", _neturl.PathEscape(parameterToString(r.pkiEzsigndocumentID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.ezsigndocumentApplyEzsigntemplateV1Request == nil {
		return localVarReturnValue, nil, reportError("ezsigndocumentApplyEzsigntemplateV1Request is required and must be specified")
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
	localVarPostBody = r.ezsigndocumentApplyEzsigntemplateV1Request
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

type ApiEzsigndocumentCreateObjectV1Request struct {
	ctx _context.Context
	ApiService *ObjectEzsigndocumentApiService
	ezsigndocumentCreateObjectV1Request *[]EzsigndocumentCreateObjectV1Request
}

func (r ApiEzsigndocumentCreateObjectV1Request) EzsigndocumentCreateObjectV1Request(ezsigndocumentCreateObjectV1Request []EzsigndocumentCreateObjectV1Request) ApiEzsigndocumentCreateObjectV1Request {
	r.ezsigndocumentCreateObjectV1Request = &ezsigndocumentCreateObjectV1Request
	return r
}

func (r ApiEzsigndocumentCreateObjectV1Request) Execute() (EzsigndocumentCreateObjectV1Response, *_nethttp.Response, error) {
	return r.ApiService.EzsigndocumentCreateObjectV1Execute(r)
}

/*
 * EzsigndocumentCreateObjectV1 Create a new Ezsigndocument
 * The endpoint allows to create one or many elements at once.

The array can contain simple (Just the object) or compound (The object and its child) objects.

Creating compound elements allows to reduce the multiple requests to create all child objects.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiEzsigndocumentCreateObjectV1Request
 */
func (a *ObjectEzsigndocumentApiService) EzsigndocumentCreateObjectV1(ctx _context.Context) ApiEzsigndocumentCreateObjectV1Request {
	return ApiEzsigndocumentCreateObjectV1Request{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return EzsigndocumentCreateObjectV1Response
 */
func (a *ObjectEzsigndocumentApiService) EzsigndocumentCreateObjectV1Execute(r ApiEzsigndocumentCreateObjectV1Request) (EzsigndocumentCreateObjectV1Response, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  EzsigndocumentCreateObjectV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectEzsigndocumentApiService.EzsigndocumentCreateObjectV1")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/ezsigndocument"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.ezsigndocumentCreateObjectV1Request == nil {
		return localVarReturnValue, nil, reportError("ezsigndocumentCreateObjectV1Request is required and must be specified")
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
	localVarPostBody = r.ezsigndocumentCreateObjectV1Request
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

type ApiEzsigndocumentDeleteObjectV1Request struct {
	ctx _context.Context
	ApiService *ObjectEzsigndocumentApiService
	pkiEzsigndocumentID int32
}


func (r ApiEzsigndocumentDeleteObjectV1Request) Execute() (EzsigndocumentDeleteObjectV1Response, *_nethttp.Response, error) {
	return r.ApiService.EzsigndocumentDeleteObjectV1Execute(r)
}

/*
 * EzsigndocumentDeleteObjectV1 Delete an existing Ezsigndocument
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param pkiEzsigndocumentID The unique ID of the Ezsigndocument
 * @return ApiEzsigndocumentDeleteObjectV1Request
 */
func (a *ObjectEzsigndocumentApiService) EzsigndocumentDeleteObjectV1(ctx _context.Context, pkiEzsigndocumentID int32) ApiEzsigndocumentDeleteObjectV1Request {
	return ApiEzsigndocumentDeleteObjectV1Request{
		ApiService: a,
		ctx: ctx,
		pkiEzsigndocumentID: pkiEzsigndocumentID,
	}
}

/*
 * Execute executes the request
 * @return EzsigndocumentDeleteObjectV1Response
 */
func (a *ObjectEzsigndocumentApiService) EzsigndocumentDeleteObjectV1Execute(r ApiEzsigndocumentDeleteObjectV1Request) (EzsigndocumentDeleteObjectV1Response, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  EzsigndocumentDeleteObjectV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectEzsigndocumentApiService.EzsigndocumentDeleteObjectV1")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/ezsigndocument/{pkiEzsigndocumentID}"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiEzsigndocumentID"+"}", _neturl.PathEscape(parameterToString(r.pkiEzsigndocumentID, "")), -1)

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

type ApiEzsigndocumentEditObjectV1Request struct {
	ctx _context.Context
	ApiService *ObjectEzsigndocumentApiService
	pkiEzsigndocumentID int32
	ezsigndocumentEditObjectV1Request *EzsigndocumentEditObjectV1Request
}

func (r ApiEzsigndocumentEditObjectV1Request) EzsigndocumentEditObjectV1Request(ezsigndocumentEditObjectV1Request EzsigndocumentEditObjectV1Request) ApiEzsigndocumentEditObjectV1Request {
	r.ezsigndocumentEditObjectV1Request = &ezsigndocumentEditObjectV1Request
	return r
}

func (r ApiEzsigndocumentEditObjectV1Request) Execute() (EzsigndocumentEditObjectV1Response, *_nethttp.Response, error) {
	return r.ApiService.EzsigndocumentEditObjectV1Execute(r)
}

/*
 * EzsigndocumentEditObjectV1 Modify an existing Ezsigndocument
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param pkiEzsigndocumentID The unique ID of the Ezsigndocument
 * @return ApiEzsigndocumentEditObjectV1Request
 */
func (a *ObjectEzsigndocumentApiService) EzsigndocumentEditObjectV1(ctx _context.Context, pkiEzsigndocumentID int32) ApiEzsigndocumentEditObjectV1Request {
	return ApiEzsigndocumentEditObjectV1Request{
		ApiService: a,
		ctx: ctx,
		pkiEzsigndocumentID: pkiEzsigndocumentID,
	}
}

/*
 * Execute executes the request
 * @return EzsigndocumentEditObjectV1Response
 */
func (a *ObjectEzsigndocumentApiService) EzsigndocumentEditObjectV1Execute(r ApiEzsigndocumentEditObjectV1Request) (EzsigndocumentEditObjectV1Response, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  EzsigndocumentEditObjectV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectEzsigndocumentApiService.EzsigndocumentEditObjectV1")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/ezsigndocument/{pkiEzsigndocumentID}"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiEzsigndocumentID"+"}", _neturl.PathEscape(parameterToString(r.pkiEzsigndocumentID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.ezsigndocumentEditObjectV1Request == nil {
		return localVarReturnValue, nil, reportError("ezsigndocumentEditObjectV1Request is required and must be specified")
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
	localVarPostBody = r.ezsigndocumentEditObjectV1Request
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

type ApiEzsigndocumentGetChildrenV1Request struct {
	ctx _context.Context
	ApiService *ObjectEzsigndocumentApiService
	pkiEzsigndocumentID int32
}


func (r ApiEzsigndocumentGetChildrenV1Request) Execute() (*_nethttp.Response, error) {
	return r.ApiService.EzsigndocumentGetChildrenV1Execute(r)
}

/*
 * EzsigndocumentGetChildrenV1 Retrieve an existing Ezsigndocument's children IDs
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param pkiEzsigndocumentID The unique ID of the Ezsigndocument
 * @return ApiEzsigndocumentGetChildrenV1Request
 */
func (a *ObjectEzsigndocumentApiService) EzsigndocumentGetChildrenV1(ctx _context.Context, pkiEzsigndocumentID int32) ApiEzsigndocumentGetChildrenV1Request {
	return ApiEzsigndocumentGetChildrenV1Request{
		ApiService: a,
		ctx: ctx,
		pkiEzsigndocumentID: pkiEzsigndocumentID,
	}
}

/*
 * Execute executes the request
 */
func (a *ObjectEzsigndocumentApiService) EzsigndocumentGetChildrenV1Execute(r ApiEzsigndocumentGetChildrenV1Request) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectEzsigndocumentApiService.EzsigndocumentGetChildrenV1")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/ezsigndocument/{pkiEzsigndocumentID}/getChildren"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiEzsigndocumentID"+"}", _neturl.PathEscape(parameterToString(r.pkiEzsigndocumentID, "")), -1)

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

type ApiEzsigndocumentGetDownloadUrlV1Request struct {
	ctx _context.Context
	ApiService *ObjectEzsigndocumentApiService
	pkiEzsigndocumentID int32
	eDocumentType string
}


func (r ApiEzsigndocumentGetDownloadUrlV1Request) Execute() (EzsigndocumentGetDownloadUrlV1Response, *_nethttp.Response, error) {
	return r.ApiService.EzsigndocumentGetDownloadUrlV1Execute(r)
}

/*
 * EzsigndocumentGetDownloadUrlV1 Retrieve a URL to download documents.
 * This endpoint returns URLs to different files that can be downloaded during the signing process.

These links will expire after 5 minutes so the download of the file should be made soon after retrieving the link.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param pkiEzsigndocumentID The unique ID of the Ezsigndocument
 * @param eDocumentType The type of document to retrieve.  1. **Initial** Is the initial document before any signature were applied. 2. **Signed** Is the final document once all signatures were applied. 3. **Proofdocument** Is the evidence report. 4. **Proof** Is the complete evidence archive including all of the above and more. 
 * @return ApiEzsigndocumentGetDownloadUrlV1Request
 */
func (a *ObjectEzsigndocumentApiService) EzsigndocumentGetDownloadUrlV1(ctx _context.Context, pkiEzsigndocumentID int32, eDocumentType string) ApiEzsigndocumentGetDownloadUrlV1Request {
	return ApiEzsigndocumentGetDownloadUrlV1Request{
		ApiService: a,
		ctx: ctx,
		pkiEzsigndocumentID: pkiEzsigndocumentID,
		eDocumentType: eDocumentType,
	}
}

/*
 * Execute executes the request
 * @return EzsigndocumentGetDownloadUrlV1Response
 */
func (a *ObjectEzsigndocumentApiService) EzsigndocumentGetDownloadUrlV1Execute(r ApiEzsigndocumentGetDownloadUrlV1Request) (EzsigndocumentGetDownloadUrlV1Response, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  EzsigndocumentGetDownloadUrlV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectEzsigndocumentApiService.EzsigndocumentGetDownloadUrlV1")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/ezsigndocument/{pkiEzsigndocumentID}/getDownloadUrl/{eDocumentType}"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiEzsigndocumentID"+"}", _neturl.PathEscape(parameterToString(r.pkiEzsigndocumentID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"eDocumentType"+"}", _neturl.PathEscape(parameterToString(r.eDocumentType, "")), -1)

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

type ApiEzsigndocumentGetObjectV1Request struct {
	ctx _context.Context
	ApiService *ObjectEzsigndocumentApiService
	pkiEzsigndocumentID int32
}


func (r ApiEzsigndocumentGetObjectV1Request) Execute() (EzsigndocumentGetObjectV1Response, *_nethttp.Response, error) {
	return r.ApiService.EzsigndocumentGetObjectV1Execute(r)
}

/*
 * EzsigndocumentGetObjectV1 Retrieve an existing Ezsigndocument
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param pkiEzsigndocumentID The unique ID of the Ezsigndocument
 * @return ApiEzsigndocumentGetObjectV1Request
 */
func (a *ObjectEzsigndocumentApiService) EzsigndocumentGetObjectV1(ctx _context.Context, pkiEzsigndocumentID int32) ApiEzsigndocumentGetObjectV1Request {
	return ApiEzsigndocumentGetObjectV1Request{
		ApiService: a,
		ctx: ctx,
		pkiEzsigndocumentID: pkiEzsigndocumentID,
	}
}

/*
 * Execute executes the request
 * @return EzsigndocumentGetObjectV1Response
 */
func (a *ObjectEzsigndocumentApiService) EzsigndocumentGetObjectV1Execute(r ApiEzsigndocumentGetObjectV1Request) (EzsigndocumentGetObjectV1Response, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  EzsigndocumentGetObjectV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectEzsigndocumentApiService.EzsigndocumentGetObjectV1")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/ezsigndocument/{pkiEzsigndocumentID}"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiEzsigndocumentID"+"}", _neturl.PathEscape(parameterToString(r.pkiEzsigndocumentID, "")), -1)

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
