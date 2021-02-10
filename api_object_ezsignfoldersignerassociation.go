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

// ObjectEzsignfoldersignerassociationApiService ObjectEzsignfoldersignerassociationApi service
type ObjectEzsignfoldersignerassociationApiService service

type ApiEzsignfoldersignerassociationCreateObjectV1Request struct {
	ctx _context.Context
	ApiService *ObjectEzsignfoldersignerassociationApiService
	ezsignfoldersignerassociationCreateObjectV1Request *[]EzsignfoldersignerassociationCreateObjectV1Request
}

func (r ApiEzsignfoldersignerassociationCreateObjectV1Request) EzsignfoldersignerassociationCreateObjectV1Request(ezsignfoldersignerassociationCreateObjectV1Request []EzsignfoldersignerassociationCreateObjectV1Request) ApiEzsignfoldersignerassociationCreateObjectV1Request {
	r.ezsignfoldersignerassociationCreateObjectV1Request = &ezsignfoldersignerassociationCreateObjectV1Request
	return r
}

func (r ApiEzsignfoldersignerassociationCreateObjectV1Request) Execute() (EzsignfoldersignerassociationCreateObjectV1Response, *_nethttp.Response, error) {
	return r.ApiService.EzsignfoldersignerassociationCreateObjectV1Execute(r)
}

/*
 * EzsignfoldersignerassociationCreateObjectV1 Create a new Ezsignfoldersignerassociation
 * The endpoint allows to create one or many elements at once.

The array can contain simple (Just the object) or compound (The object and its child) objects.

Creating compound elements allows to reduce the multiple requests to create all child objects.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiEzsignfoldersignerassociationCreateObjectV1Request
 */
func (a *ObjectEzsignfoldersignerassociationApiService) EzsignfoldersignerassociationCreateObjectV1(ctx _context.Context) ApiEzsignfoldersignerassociationCreateObjectV1Request {
	return ApiEzsignfoldersignerassociationCreateObjectV1Request{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return EzsignfoldersignerassociationCreateObjectV1Response
 */
func (a *ObjectEzsignfoldersignerassociationApiService) EzsignfoldersignerassociationCreateObjectV1Execute(r ApiEzsignfoldersignerassociationCreateObjectV1Request) (EzsignfoldersignerassociationCreateObjectV1Response, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  EzsignfoldersignerassociationCreateObjectV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectEzsignfoldersignerassociationApiService.EzsignfoldersignerassociationCreateObjectV1")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/ezsignfoldersignerassociation"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.ezsignfoldersignerassociationCreateObjectV1Request == nil {
		return localVarReturnValue, nil, reportError("ezsignfoldersignerassociationCreateObjectV1Request is required and must be specified")
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
	localVarPostBody = r.ezsignfoldersignerassociationCreateObjectV1Request
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

type ApiEzsignfoldersignerassociationDeleteObjectV1Request struct {
	ctx _context.Context
	ApiService *ObjectEzsignfoldersignerassociationApiService
	pkiEzsignfoldersignerassociationID int32
}


func (r ApiEzsignfoldersignerassociationDeleteObjectV1Request) Execute() (EzsignfoldersignerassociationDeleteObjectV1Response, *_nethttp.Response, error) {
	return r.ApiService.EzsignfoldersignerassociationDeleteObjectV1Execute(r)
}

/*
 * EzsignfoldersignerassociationDeleteObjectV1 Delete an existing Ezsignfoldersignerassociation
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param pkiEzsignfoldersignerassociationID The unique ID of the Ezsignfoldersignerassociation
 * @return ApiEzsignfoldersignerassociationDeleteObjectV1Request
 */
func (a *ObjectEzsignfoldersignerassociationApiService) EzsignfoldersignerassociationDeleteObjectV1(ctx _context.Context, pkiEzsignfoldersignerassociationID int32) ApiEzsignfoldersignerassociationDeleteObjectV1Request {
	return ApiEzsignfoldersignerassociationDeleteObjectV1Request{
		ApiService: a,
		ctx: ctx,
		pkiEzsignfoldersignerassociationID: pkiEzsignfoldersignerassociationID,
	}
}

/*
 * Execute executes the request
 * @return EzsignfoldersignerassociationDeleteObjectV1Response
 */
func (a *ObjectEzsignfoldersignerassociationApiService) EzsignfoldersignerassociationDeleteObjectV1Execute(r ApiEzsignfoldersignerassociationDeleteObjectV1Request) (EzsignfoldersignerassociationDeleteObjectV1Response, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  EzsignfoldersignerassociationDeleteObjectV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectEzsignfoldersignerassociationApiService.EzsignfoldersignerassociationDeleteObjectV1")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/ezsignfoldersignerassociation/{pkiEzsignfoldersignerassociationID}"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiEzsignfoldersignerassociationID"+"}", _neturl.PathEscape(parameterToString(r.pkiEzsignfoldersignerassociationID, "")), -1)

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

type ApiEzsignfoldersignerassociationEditObjectV1Request struct {
	ctx _context.Context
	ApiService *ObjectEzsignfoldersignerassociationApiService
	pkiEzsignfoldersignerassociationID int32
	ezsignfoldersignerassociationEditObjectV1Request *EzsignfoldersignerassociationEditObjectV1Request
}

func (r ApiEzsignfoldersignerassociationEditObjectV1Request) EzsignfoldersignerassociationEditObjectV1Request(ezsignfoldersignerassociationEditObjectV1Request EzsignfoldersignerassociationEditObjectV1Request) ApiEzsignfoldersignerassociationEditObjectV1Request {
	r.ezsignfoldersignerassociationEditObjectV1Request = &ezsignfoldersignerassociationEditObjectV1Request
	return r
}

func (r ApiEzsignfoldersignerassociationEditObjectV1Request) Execute() (EzsignfoldersignerassociationEditObjectV1Response, *_nethttp.Response, error) {
	return r.ApiService.EzsignfoldersignerassociationEditObjectV1Execute(r)
}

/*
 * EzsignfoldersignerassociationEditObjectV1 Modify an existing Ezsignfoldersignerassociation
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param pkiEzsignfoldersignerassociationID The unique ID of the Ezsignfoldersignerassociation
 * @return ApiEzsignfoldersignerassociationEditObjectV1Request
 */
func (a *ObjectEzsignfoldersignerassociationApiService) EzsignfoldersignerassociationEditObjectV1(ctx _context.Context, pkiEzsignfoldersignerassociationID int32) ApiEzsignfoldersignerassociationEditObjectV1Request {
	return ApiEzsignfoldersignerassociationEditObjectV1Request{
		ApiService: a,
		ctx: ctx,
		pkiEzsignfoldersignerassociationID: pkiEzsignfoldersignerassociationID,
	}
}

/*
 * Execute executes the request
 * @return EzsignfoldersignerassociationEditObjectV1Response
 */
func (a *ObjectEzsignfoldersignerassociationApiService) EzsignfoldersignerassociationEditObjectV1Execute(r ApiEzsignfoldersignerassociationEditObjectV1Request) (EzsignfoldersignerassociationEditObjectV1Response, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  EzsignfoldersignerassociationEditObjectV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectEzsignfoldersignerassociationApiService.EzsignfoldersignerassociationEditObjectV1")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/ezsignfoldersignerassociation/{pkiEzsignfoldersignerassociationID}"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiEzsignfoldersignerassociationID"+"}", _neturl.PathEscape(parameterToString(r.pkiEzsignfoldersignerassociationID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.ezsignfoldersignerassociationEditObjectV1Request == nil {
		return localVarReturnValue, nil, reportError("ezsignfoldersignerassociationEditObjectV1Request is required and must be specified")
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
	localVarPostBody = r.ezsignfoldersignerassociationEditObjectV1Request
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

type ApiEzsignfoldersignerassociationGetChildrenV1Request struct {
	ctx _context.Context
	ApiService *ObjectEzsignfoldersignerassociationApiService
	pkiEzsignfoldersignerassociationID int32
}


func (r ApiEzsignfoldersignerassociationGetChildrenV1Request) Execute() (*_nethttp.Response, error) {
	return r.ApiService.EzsignfoldersignerassociationGetChildrenV1Execute(r)
}

/*
 * EzsignfoldersignerassociationGetChildrenV1 Retrieve an existing Ezsignfoldersignerassociation's children IDs
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param pkiEzsignfoldersignerassociationID The unique ID of the Ezsignfoldersignerassociation
 * @return ApiEzsignfoldersignerassociationGetChildrenV1Request
 */
func (a *ObjectEzsignfoldersignerassociationApiService) EzsignfoldersignerassociationGetChildrenV1(ctx _context.Context, pkiEzsignfoldersignerassociationID int32) ApiEzsignfoldersignerassociationGetChildrenV1Request {
	return ApiEzsignfoldersignerassociationGetChildrenV1Request{
		ApiService: a,
		ctx: ctx,
		pkiEzsignfoldersignerassociationID: pkiEzsignfoldersignerassociationID,
	}
}

/*
 * Execute executes the request
 */
func (a *ObjectEzsignfoldersignerassociationApiService) EzsignfoldersignerassociationGetChildrenV1Execute(r ApiEzsignfoldersignerassociationGetChildrenV1Request) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectEzsignfoldersignerassociationApiService.EzsignfoldersignerassociationGetChildrenV1")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/ezsignfoldersignerassociation/{pkiEzsignfoldersignerassociationID}/getChildren"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiEzsignfoldersignerassociationID"+"}", _neturl.PathEscape(parameterToString(r.pkiEzsignfoldersignerassociationID, "")), -1)

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

type ApiEzsignfoldersignerassociationGetInPersonLoginUrlV1Request struct {
	ctx _context.Context
	ApiService *ObjectEzsignfoldersignerassociationApiService
	pkiEzsignfoldersignerassociationID int32
}


func (r ApiEzsignfoldersignerassociationGetInPersonLoginUrlV1Request) Execute() (EzsignfoldersignerassociationGetInPersonLoginUrlV1Response, *_nethttp.Response, error) {
	return r.ApiService.EzsignfoldersignerassociationGetInPersonLoginUrlV1Execute(r)
}

/*
 * EzsignfoldersignerassociationGetInPersonLoginUrlV1 Retrieve a Login Url to allow In-Person signing
 * This endpoint returns a Login Url that can be used in a browser or embedded in an I-Frame to allow in person signing.

The signer Login type must be configured as In-Person.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param pkiEzsignfoldersignerassociationID The unique ID of the Ezsignfoldersignerassociation
 * @return ApiEzsignfoldersignerassociationGetInPersonLoginUrlV1Request
 */
func (a *ObjectEzsignfoldersignerassociationApiService) EzsignfoldersignerassociationGetInPersonLoginUrlV1(ctx _context.Context, pkiEzsignfoldersignerassociationID int32) ApiEzsignfoldersignerassociationGetInPersonLoginUrlV1Request {
	return ApiEzsignfoldersignerassociationGetInPersonLoginUrlV1Request{
		ApiService: a,
		ctx: ctx,
		pkiEzsignfoldersignerassociationID: pkiEzsignfoldersignerassociationID,
	}
}

/*
 * Execute executes the request
 * @return EzsignfoldersignerassociationGetInPersonLoginUrlV1Response
 */
func (a *ObjectEzsignfoldersignerassociationApiService) EzsignfoldersignerassociationGetInPersonLoginUrlV1Execute(r ApiEzsignfoldersignerassociationGetInPersonLoginUrlV1Request) (EzsignfoldersignerassociationGetInPersonLoginUrlV1Response, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  EzsignfoldersignerassociationGetInPersonLoginUrlV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectEzsignfoldersignerassociationApiService.EzsignfoldersignerassociationGetInPersonLoginUrlV1")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/ezsignfoldersignerassociation/{pkiEzsignfoldersignerassociationID}/getInPersonLoginUrl"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiEzsignfoldersignerassociationID"+"}", _neturl.PathEscape(parameterToString(r.pkiEzsignfoldersignerassociationID, "")), -1)

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

type ApiEzsignfoldersignerassociationGetObjectV1Request struct {
	ctx _context.Context
	ApiService *ObjectEzsignfoldersignerassociationApiService
	pkiEzsignfoldersignerassociationID int32
}


func (r ApiEzsignfoldersignerassociationGetObjectV1Request) Execute() (EzsignfoldersignerassociationGetObjectV1Response, *_nethttp.Response, error) {
	return r.ApiService.EzsignfoldersignerassociationGetObjectV1Execute(r)
}

/*
 * EzsignfoldersignerassociationGetObjectV1 Retrieve an existing Ezsignfoldersignerassociation
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param pkiEzsignfoldersignerassociationID The unique ID of the Ezsignfoldersignerassociation
 * @return ApiEzsignfoldersignerassociationGetObjectV1Request
 */
func (a *ObjectEzsignfoldersignerassociationApiService) EzsignfoldersignerassociationGetObjectV1(ctx _context.Context, pkiEzsignfoldersignerassociationID int32) ApiEzsignfoldersignerassociationGetObjectV1Request {
	return ApiEzsignfoldersignerassociationGetObjectV1Request{
		ApiService: a,
		ctx: ctx,
		pkiEzsignfoldersignerassociationID: pkiEzsignfoldersignerassociationID,
	}
}

/*
 * Execute executes the request
 * @return EzsignfoldersignerassociationGetObjectV1Response
 */
func (a *ObjectEzsignfoldersignerassociationApiService) EzsignfoldersignerassociationGetObjectV1Execute(r ApiEzsignfoldersignerassociationGetObjectV1Request) (EzsignfoldersignerassociationGetObjectV1Response, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  EzsignfoldersignerassociationGetObjectV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectEzsignfoldersignerassociationApiService.EzsignfoldersignerassociationGetObjectV1")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/ezsignfoldersignerassociation/{pkiEzsignfoldersignerassociationID}"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiEzsignfoldersignerassociationID"+"}", _neturl.PathEscape(parameterToString(r.pkiEzsignfoldersignerassociationID, "")), -1)

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
