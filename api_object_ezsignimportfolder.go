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


// ObjectEzsignimportfolderAPIService ObjectEzsignimportfolderAPI service
type ObjectEzsignimportfolderAPIService service

type ApiEzsignimportfolderDeleteObjectV1Request struct {
	ctx context.Context
	ApiService *ObjectEzsignimportfolderAPIService
	pkiEzsignimportfolderID int32
}

func (r ApiEzsignimportfolderDeleteObjectV1Request) Execute() (*EzsignimportfolderDeleteObjectV1Response, *http.Response, error) {
	return r.ApiService.EzsignimportfolderDeleteObjectV1Execute(r)
}

/*
EzsignimportfolderDeleteObjectV1 Delete an existing Ezsignimportfolder



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pkiEzsignimportfolderID The unique ID of the Ezsignimportfolder
 @return ApiEzsignimportfolderDeleteObjectV1Request
*/
func (a *ObjectEzsignimportfolderAPIService) EzsignimportfolderDeleteObjectV1(ctx context.Context, pkiEzsignimportfolderID int32) ApiEzsignimportfolderDeleteObjectV1Request {
	return ApiEzsignimportfolderDeleteObjectV1Request{
		ApiService: a,
		ctx: ctx,
		pkiEzsignimportfolderID: pkiEzsignimportfolderID,
	}
}

// Execute executes the request
//  @return EzsignimportfolderDeleteObjectV1Response
func (a *ObjectEzsignimportfolderAPIService) EzsignimportfolderDeleteObjectV1Execute(r ApiEzsignimportfolderDeleteObjectV1Request) (*EzsignimportfolderDeleteObjectV1Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EzsignimportfolderDeleteObjectV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectEzsignimportfolderAPIService.EzsignimportfolderDeleteObjectV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/ezsignimportfolder/{pkiEzsignimportfolderID}"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiEzsignimportfolderID"+"}", url.PathEscape(parameterValueToString(r.pkiEzsignimportfolderID, "pkiEzsignimportfolderID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pkiEzsignimportfolderID < 0 {
		return localVarReturnValue, nil, reportError("pkiEzsignimportfolderID must be greater than 0")
	}
	if r.pkiEzsignimportfolderID > 16777215 {
		return localVarReturnValue, nil, reportError("pkiEzsignimportfolderID must be less than 16777215")
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

type ApiEzsignimportfolderGetListV1Request struct {
	ctx context.Context
	ApiService *ObjectEzsignimportfolderAPIService
	eOrderBy *string
	iRowMax *int32
	iRowOffset *int32
	acceptLanguage *HeaderAcceptLanguage
	sFilter *string
}

// Specify how you want the results to be sorted
func (r ApiEzsignimportfolderGetListV1Request) EOrderBy(eOrderBy string) ApiEzsignimportfolderGetListV1Request {
	r.eOrderBy = &eOrderBy
	return r
}

func (r ApiEzsignimportfolderGetListV1Request) IRowMax(iRowMax int32) ApiEzsignimportfolderGetListV1Request {
	r.iRowMax = &iRowMax
	return r
}

func (r ApiEzsignimportfolderGetListV1Request) IRowOffset(iRowOffset int32) ApiEzsignimportfolderGetListV1Request {
	r.iRowOffset = &iRowOffset
	return r
}

func (r ApiEzsignimportfolderGetListV1Request) AcceptLanguage(acceptLanguage HeaderAcceptLanguage) ApiEzsignimportfolderGetListV1Request {
	r.acceptLanguage = &acceptLanguage
	return r
}

func (r ApiEzsignimportfolderGetListV1Request) SFilter(sFilter string) ApiEzsignimportfolderGetListV1Request {
	r.sFilter = &sFilter
	return r
}

func (r ApiEzsignimportfolderGetListV1Request) Execute() (*EzsignimportfolderGetListV1Response, *http.Response, error) {
	return r.ApiService.EzsignimportfolderGetListV1Execute(r)
}

/*
EzsignimportfolderGetListV1 Retrieve Ezsignimportfolder list



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiEzsignimportfolderGetListV1Request
*/
func (a *ObjectEzsignimportfolderAPIService) EzsignimportfolderGetListV1(ctx context.Context) ApiEzsignimportfolderGetListV1Request {
	return ApiEzsignimportfolderGetListV1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return EzsignimportfolderGetListV1Response
func (a *ObjectEzsignimportfolderAPIService) EzsignimportfolderGetListV1Execute(r ApiEzsignimportfolderGetListV1Request) (*EzsignimportfolderGetListV1Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EzsignimportfolderGetListV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectEzsignimportfolderAPIService.EzsignimportfolderGetListV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/ezsignimportfolder/getList"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.eOrderBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "eOrderBy", r.eOrderBy, "form", "")
	}
	if r.iRowMax != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "iRowMax", r.iRowMax, "form", "")
	}
	if r.iRowOffset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "iRowOffset", r.iRowOffset, "form", "")
	} else {
		var defaultValue int32 = 0
		r.iRowOffset = &defaultValue
	}
	if r.sFilter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sFilter", r.sFilter, "form", "")
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
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Accept-Language", r.acceptLanguage, "simple", "")
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
		if localVarHTTPResponse.StatusCode == 406 {
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

type ApiEzsignimportfolderGetObjectV2Request struct {
	ctx context.Context
	ApiService *ObjectEzsignimportfolderAPIService
	pkiEzsignimportfolderID int32
}

func (r ApiEzsignimportfolderGetObjectV2Request) Execute() (*EzsignimportfolderGetObjectV2Response, *http.Response, error) {
	return r.ApiService.EzsignimportfolderGetObjectV2Execute(r)
}

/*
EzsignimportfolderGetObjectV2 Retrieve an existing Ezsignimportfolder



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pkiEzsignimportfolderID The unique ID of the Ezsignimportfolder
 @return ApiEzsignimportfolderGetObjectV2Request
*/
func (a *ObjectEzsignimportfolderAPIService) EzsignimportfolderGetObjectV2(ctx context.Context, pkiEzsignimportfolderID int32) ApiEzsignimportfolderGetObjectV2Request {
	return ApiEzsignimportfolderGetObjectV2Request{
		ApiService: a,
		ctx: ctx,
		pkiEzsignimportfolderID: pkiEzsignimportfolderID,
	}
}

// Execute executes the request
//  @return EzsignimportfolderGetObjectV2Response
func (a *ObjectEzsignimportfolderAPIService) EzsignimportfolderGetObjectV2Execute(r ApiEzsignimportfolderGetObjectV2Request) (*EzsignimportfolderGetObjectV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EzsignimportfolderGetObjectV2Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectEzsignimportfolderAPIService.EzsignimportfolderGetObjectV2")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/2/object/ezsignimportfolder/{pkiEzsignimportfolderID}"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiEzsignimportfolderID"+"}", url.PathEscape(parameterValueToString(r.pkiEzsignimportfolderID, "pkiEzsignimportfolderID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pkiEzsignimportfolderID < 0 {
		return localVarReturnValue, nil, reportError("pkiEzsignimportfolderID must be greater than 0")
	}
	if r.pkiEzsignimportfolderID > 16777215 {
		return localVarReturnValue, nil, reportError("pkiEzsignimportfolderID must be less than 16777215")
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
