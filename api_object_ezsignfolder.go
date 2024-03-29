/*
eZmax API Definition

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.1.3
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

// ObjectEzsignfolderApiService ObjectEzsignfolderApi service
type ObjectEzsignfolderApiService service

type ApiEzsignfolderCreateObjectV1Request struct {
	ctx _context.Context
	ApiService *ObjectEzsignfolderApiService
	ezsignfolderCreateObjectV1Request *[]EzsignfolderCreateObjectV1Request
}

func (r ApiEzsignfolderCreateObjectV1Request) EzsignfolderCreateObjectV1Request(ezsignfolderCreateObjectV1Request []EzsignfolderCreateObjectV1Request) ApiEzsignfolderCreateObjectV1Request {
	r.ezsignfolderCreateObjectV1Request = &ezsignfolderCreateObjectV1Request
	return r
}

func (r ApiEzsignfolderCreateObjectV1Request) Execute() (EzsignfolderCreateObjectV1Response, *_nethttp.Response, error) {
	return r.ApiService.EzsignfolderCreateObjectV1Execute(r)
}

/*
EzsignfolderCreateObjectV1 Create a new Ezsignfolder

The endpoint allows to create one or many elements at once.

The array can contain simple (Just the object) or compound (The object and its child) objects.

Creating compound elements allows to reduce the multiple requests to create all child objects.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiEzsignfolderCreateObjectV1Request
*/
func (a *ObjectEzsignfolderApiService) EzsignfolderCreateObjectV1(ctx _context.Context) ApiEzsignfolderCreateObjectV1Request {
	return ApiEzsignfolderCreateObjectV1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return EzsignfolderCreateObjectV1Response
func (a *ObjectEzsignfolderApiService) EzsignfolderCreateObjectV1Execute(r ApiEzsignfolderCreateObjectV1Request) (EzsignfolderCreateObjectV1Response, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  EzsignfolderCreateObjectV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectEzsignfolderApiService.EzsignfolderCreateObjectV1")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/ezsignfolder"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.ezsignfolderCreateObjectV1Request == nil {
		return localVarReturnValue, nil, reportError("ezsignfolderCreateObjectV1Request is required and must be specified")
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
	localVarPostBody = r.ezsignfolderCreateObjectV1Request
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

type ApiEzsignfolderDeleteObjectV1Request struct {
	ctx _context.Context
	ApiService *ObjectEzsignfolderApiService
	pkiEzsignfolderID int32
}


func (r ApiEzsignfolderDeleteObjectV1Request) Execute() (EzsignfolderDeleteObjectV1Response, *_nethttp.Response, error) {
	return r.ApiService.EzsignfolderDeleteObjectV1Execute(r)
}

/*
EzsignfolderDeleteObjectV1 Delete an existing Ezsignfolder

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pkiEzsignfolderID
 @return ApiEzsignfolderDeleteObjectV1Request
*/
func (a *ObjectEzsignfolderApiService) EzsignfolderDeleteObjectV1(ctx _context.Context, pkiEzsignfolderID int32) ApiEzsignfolderDeleteObjectV1Request {
	return ApiEzsignfolderDeleteObjectV1Request{
		ApiService: a,
		ctx: ctx,
		pkiEzsignfolderID: pkiEzsignfolderID,
	}
}

// Execute executes the request
//  @return EzsignfolderDeleteObjectV1Response
func (a *ObjectEzsignfolderApiService) EzsignfolderDeleteObjectV1Execute(r ApiEzsignfolderDeleteObjectV1Request) (EzsignfolderDeleteObjectV1Response, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  EzsignfolderDeleteObjectV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectEzsignfolderApiService.EzsignfolderDeleteObjectV1")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/ezsignfolder/{pkiEzsignfolderID}"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiEzsignfolderID"+"}", _neturl.PathEscape(parameterToString(r.pkiEzsignfolderID, "")), -1)

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

type ApiEzsignfolderGetChildrenV1Request struct {
	ctx _context.Context
	ApiService *ObjectEzsignfolderApiService
	pkiEzsignfolderID int32
}


func (r ApiEzsignfolderGetChildrenV1Request) Execute() (*_nethttp.Response, error) {
	return r.ApiService.EzsignfolderGetChildrenV1Execute(r)
}

/*
EzsignfolderGetChildrenV1 Retrieve an existing Ezsignfolder's children IDs

## ⚠️EARLY ADOPTERS WARNING

### This endpoint is not officially released. Its definition might still change and it might not be available in every environment and region.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pkiEzsignfolderID
 @return ApiEzsignfolderGetChildrenV1Request
*/
func (a *ObjectEzsignfolderApiService) EzsignfolderGetChildrenV1(ctx _context.Context, pkiEzsignfolderID int32) ApiEzsignfolderGetChildrenV1Request {
	return ApiEzsignfolderGetChildrenV1Request{
		ApiService: a,
		ctx: ctx,
		pkiEzsignfolderID: pkiEzsignfolderID,
	}
}

// Execute executes the request
func (a *ObjectEzsignfolderApiService) EzsignfolderGetChildrenV1Execute(r ApiEzsignfolderGetChildrenV1Request) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectEzsignfolderApiService.EzsignfolderGetChildrenV1")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/ezsignfolder/{pkiEzsignfolderID}/getChildren"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiEzsignfolderID"+"}", _neturl.PathEscape(parameterToString(r.pkiEzsignfolderID, "")), -1)

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
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

type ApiEzsignfolderGetFormsDataV1Request struct {
	ctx _context.Context
	ApiService *ObjectEzsignfolderApiService
	pkiEzsignfolderID int32
}


func (r ApiEzsignfolderGetFormsDataV1Request) Execute() (EzsignfolderGetFormsDataV1Response, *_nethttp.Response, error) {
	return r.ApiService.EzsignfolderGetFormsDataV1Execute(r)
}

/*
EzsignfolderGetFormsDataV1 Retrieve an existing Ezsignfolder's forms data

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pkiEzsignfolderID
 @return ApiEzsignfolderGetFormsDataV1Request
*/
func (a *ObjectEzsignfolderApiService) EzsignfolderGetFormsDataV1(ctx _context.Context, pkiEzsignfolderID int32) ApiEzsignfolderGetFormsDataV1Request {
	return ApiEzsignfolderGetFormsDataV1Request{
		ApiService: a,
		ctx: ctx,
		pkiEzsignfolderID: pkiEzsignfolderID,
	}
}

// Execute executes the request
//  @return EzsignfolderGetFormsDataV1Response
func (a *ObjectEzsignfolderApiService) EzsignfolderGetFormsDataV1Execute(r ApiEzsignfolderGetFormsDataV1Request) (EzsignfolderGetFormsDataV1Response, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  EzsignfolderGetFormsDataV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectEzsignfolderApiService.EzsignfolderGetFormsDataV1")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/ezsignfolder/{pkiEzsignfolderID}/getFormsData"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiEzsignfolderID"+"}", _neturl.PathEscape(parameterToString(r.pkiEzsignfolderID, "")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json", "application/zip"}

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
		if localVarHTTPResponse.StatusCode == 406 {
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

type ApiEzsignfolderGetListV1Request struct {
	ctx _context.Context
	ApiService *ObjectEzsignfolderApiService
	eOrderBy *string
	iRowMax *int32
	iRowOffset *int32
	acceptLanguage *HeaderAcceptLanguage
	sFilter *string
}

// Specify how you want the results to be sorted
func (r ApiEzsignfolderGetListV1Request) EOrderBy(eOrderBy string) ApiEzsignfolderGetListV1Request {
	r.eOrderBy = &eOrderBy
	return r
}
func (r ApiEzsignfolderGetListV1Request) IRowMax(iRowMax int32) ApiEzsignfolderGetListV1Request {
	r.iRowMax = &iRowMax
	return r
}
func (r ApiEzsignfolderGetListV1Request) IRowOffset(iRowOffset int32) ApiEzsignfolderGetListV1Request {
	r.iRowOffset = &iRowOffset
	return r
}
func (r ApiEzsignfolderGetListV1Request) AcceptLanguage(acceptLanguage HeaderAcceptLanguage) ApiEzsignfolderGetListV1Request {
	r.acceptLanguage = &acceptLanguage
	return r
}
func (r ApiEzsignfolderGetListV1Request) SFilter(sFilter string) ApiEzsignfolderGetListV1Request {
	r.sFilter = &sFilter
	return r
}

func (r ApiEzsignfolderGetListV1Request) Execute() (EzsignfolderGetListV1Response, *_nethttp.Response, error) {
	return r.ApiService.EzsignfolderGetListV1Execute(r)
}

/*
EzsignfolderGetListV1 Retrieve Ezsignfolder list

Enum values that can be filtered in query parameter *sFilter*:

| Variable | Valid values |
|---|---|
| eEzsignfolderStep | Unsent<br>Sent<br>PartiallySigned<br>Expired<br>Completed<br>Archived |
| eEzsignfoldertypePrivacylevel | User<br>Usergroup |

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiEzsignfolderGetListV1Request
*/
func (a *ObjectEzsignfolderApiService) EzsignfolderGetListV1(ctx _context.Context) ApiEzsignfolderGetListV1Request {
	return ApiEzsignfolderGetListV1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return EzsignfolderGetListV1Response
func (a *ObjectEzsignfolderApiService) EzsignfolderGetListV1Execute(r ApiEzsignfolderGetListV1Request) (EzsignfolderGetListV1Response, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  EzsignfolderGetListV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectEzsignfolderApiService.EzsignfolderGetListV1")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/ezsignfolder/getList"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.eOrderBy != nil {
		localVarQueryParams.Add("eOrderBy", parameterToString(*r.eOrderBy, ""))
	}
	if r.iRowMax != nil {
		localVarQueryParams.Add("iRowMax", parameterToString(*r.iRowMax, ""))
	}
	if r.iRowOffset != nil {
		localVarQueryParams.Add("iRowOffset", parameterToString(*r.iRowOffset, ""))
	}
	if r.sFilter != nil {
		localVarQueryParams.Add("sFilter", parameterToString(*r.sFilter, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.acceptLanguage != nil {
		localVarHeaderParams["Accept-Language"] = parameterToString(*r.acceptLanguage, "")
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
		if localVarHTTPResponse.StatusCode == 406 {
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

type ApiEzsignfolderGetObjectV1Request struct {
	ctx _context.Context
	ApiService *ObjectEzsignfolderApiService
	pkiEzsignfolderID int32
}


func (r ApiEzsignfolderGetObjectV1Request) Execute() (EzsignfolderGetObjectV1Response, *_nethttp.Response, error) {
	return r.ApiService.EzsignfolderGetObjectV1Execute(r)
}

/*
EzsignfolderGetObjectV1 Retrieve an existing Ezsignfolder

## ⚠️EARLY ADOPTERS WARNING

### This endpoint is not officially released. Its definition might still change and it might not be available in every environment and region.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pkiEzsignfolderID
 @return ApiEzsignfolderGetObjectV1Request
*/
func (a *ObjectEzsignfolderApiService) EzsignfolderGetObjectV1(ctx _context.Context, pkiEzsignfolderID int32) ApiEzsignfolderGetObjectV1Request {
	return ApiEzsignfolderGetObjectV1Request{
		ApiService: a,
		ctx: ctx,
		pkiEzsignfolderID: pkiEzsignfolderID,
	}
}

// Execute executes the request
//  @return EzsignfolderGetObjectV1Response
func (a *ObjectEzsignfolderApiService) EzsignfolderGetObjectV1Execute(r ApiEzsignfolderGetObjectV1Request) (EzsignfolderGetObjectV1Response, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  EzsignfolderGetObjectV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectEzsignfolderApiService.EzsignfolderGetObjectV1")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/ezsignfolder/{pkiEzsignfolderID}"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiEzsignfolderID"+"}", _neturl.PathEscape(parameterToString(r.pkiEzsignfolderID, "")), -1)

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

type ApiEzsignfolderSendV1Request struct {
	ctx _context.Context
	ApiService *ObjectEzsignfolderApiService
	pkiEzsignfolderID int32
	ezsignfolderSendV1Request *EzsignfolderSendV1Request
}

func (r ApiEzsignfolderSendV1Request) EzsignfolderSendV1Request(ezsignfolderSendV1Request EzsignfolderSendV1Request) ApiEzsignfolderSendV1Request {
	r.ezsignfolderSendV1Request = &ezsignfolderSendV1Request
	return r
}

func (r ApiEzsignfolderSendV1Request) Execute() (EzsignfolderSendV1Response, *_nethttp.Response, error) {
	return r.ApiService.EzsignfolderSendV1Execute(r)
}

/*
EzsignfolderSendV1 Send the Ezsignfolder to the signatories for signature

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pkiEzsignfolderID
 @return ApiEzsignfolderSendV1Request
*/
func (a *ObjectEzsignfolderApiService) EzsignfolderSendV1(ctx _context.Context, pkiEzsignfolderID int32) ApiEzsignfolderSendV1Request {
	return ApiEzsignfolderSendV1Request{
		ApiService: a,
		ctx: ctx,
		pkiEzsignfolderID: pkiEzsignfolderID,
	}
}

// Execute executes the request
//  @return EzsignfolderSendV1Response
func (a *ObjectEzsignfolderApiService) EzsignfolderSendV1Execute(r ApiEzsignfolderSendV1Request) (EzsignfolderSendV1Response, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  EzsignfolderSendV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectEzsignfolderApiService.EzsignfolderSendV1")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/ezsignfolder/{pkiEzsignfolderID}/send"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiEzsignfolderID"+"}", _neturl.PathEscape(parameterToString(r.pkiEzsignfolderID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.ezsignfolderSendV1Request == nil {
		return localVarReturnValue, nil, reportError("ezsignfolderSendV1Request is required and must be specified")
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
	localVarPostBody = r.ezsignfolderSendV1Request
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

type ApiEzsignfolderUnsendV1Request struct {
	ctx _context.Context
	ApiService *ObjectEzsignfolderApiService
	pkiEzsignfolderID int32
	body *string
}

func (r ApiEzsignfolderUnsendV1Request) Body(body string) ApiEzsignfolderUnsendV1Request {
	r.body = &body
	return r
}

func (r ApiEzsignfolderUnsendV1Request) Execute() (EzsignfolderUnsendV1Response, *_nethttp.Response, error) {
	return r.ApiService.EzsignfolderUnsendV1Execute(r)
}

/*
EzsignfolderUnsendV1 Unsend the Ezsignfolder

Once an Ezsignfolder has been sent to signatories, it cannot be modified.

Using this endpoint, you can unsend the Ezsignfolder and make it modifiable again.

Signatories will receive an email informing them the signature process was aborted and they might receive a new invitation to sign.

⚠️ Warning: Any signature previously made by signatories on "Non-completed" Ezsigndocuments will be lost.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pkiEzsignfolderID
 @return ApiEzsignfolderUnsendV1Request
*/
func (a *ObjectEzsignfolderApiService) EzsignfolderUnsendV1(ctx _context.Context, pkiEzsignfolderID int32) ApiEzsignfolderUnsendV1Request {
	return ApiEzsignfolderUnsendV1Request{
		ApiService: a,
		ctx: ctx,
		pkiEzsignfolderID: pkiEzsignfolderID,
	}
}

// Execute executes the request
//  @return EzsignfolderUnsendV1Response
func (a *ObjectEzsignfolderApiService) EzsignfolderUnsendV1Execute(r ApiEzsignfolderUnsendV1Request) (EzsignfolderUnsendV1Response, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  EzsignfolderUnsendV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectEzsignfolderApiService.EzsignfolderUnsendV1")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/ezsignfolder/{pkiEzsignfolderID}/unsend"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiEzsignfolderID"+"}", _neturl.PathEscape(parameterToString(r.pkiEzsignfolderID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
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
	localVarPostBody = r.body
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
