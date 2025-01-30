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


// ObjectTranqcontractAPIService ObjectTranqcontractAPI service
type ObjectTranqcontractAPIService service

type ApiTranqcontractGetCommunicationCountV1Request struct {
	ctx context.Context
	ApiService *ObjectTranqcontractAPIService
	pkiTranqcontractID int32
}

func (r ApiTranqcontractGetCommunicationCountV1Request) Execute() (*TranqcontractGetCommunicationCountV1Response, *http.Response, error) {
	return r.ApiService.TranqcontractGetCommunicationCountV1Execute(r)
}

/*
TranqcontractGetCommunicationCountV1 Retrieve Communication count



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pkiTranqcontractID
 @return ApiTranqcontractGetCommunicationCountV1Request
*/
func (a *ObjectTranqcontractAPIService) TranqcontractGetCommunicationCountV1(ctx context.Context, pkiTranqcontractID int32) ApiTranqcontractGetCommunicationCountV1Request {
	return ApiTranqcontractGetCommunicationCountV1Request{
		ApiService: a,
		ctx: ctx,
		pkiTranqcontractID: pkiTranqcontractID,
	}
}

// Execute executes the request
//  @return TranqcontractGetCommunicationCountV1Response
func (a *ObjectTranqcontractAPIService) TranqcontractGetCommunicationCountV1Execute(r ApiTranqcontractGetCommunicationCountV1Request) (*TranqcontractGetCommunicationCountV1Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TranqcontractGetCommunicationCountV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectTranqcontractAPIService.TranqcontractGetCommunicationCountV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/tranqcontract/{pkiTranqcontractID}/getCommunicationCount"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiTranqcontractID"+"}", url.PathEscape(parameterValueToString(r.pkiTranqcontractID, "pkiTranqcontractID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pkiTranqcontractID < 0 {
		return localVarReturnValue, nil, reportError("pkiTranqcontractID must be greater than 0")
	}
	if r.pkiTranqcontractID > 16777215 {
		return localVarReturnValue, nil, reportError("pkiTranqcontractID must be less than 16777215")
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

type ApiTranqcontractGetCommunicationListV1Request struct {
	ctx context.Context
	ApiService *ObjectTranqcontractAPIService
	pkiTranqcontractID int32
}

func (r ApiTranqcontractGetCommunicationListV1Request) Execute() (*TranqcontractGetCommunicationListV1Response, *http.Response, error) {
	return r.ApiService.TranqcontractGetCommunicationListV1Execute(r)
}

/*
TranqcontractGetCommunicationListV1 Retrieve Communication list



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pkiTranqcontractID
 @return ApiTranqcontractGetCommunicationListV1Request
*/
func (a *ObjectTranqcontractAPIService) TranqcontractGetCommunicationListV1(ctx context.Context, pkiTranqcontractID int32) ApiTranqcontractGetCommunicationListV1Request {
	return ApiTranqcontractGetCommunicationListV1Request{
		ApiService: a,
		ctx: ctx,
		pkiTranqcontractID: pkiTranqcontractID,
	}
}

// Execute executes the request
//  @return TranqcontractGetCommunicationListV1Response
func (a *ObjectTranqcontractAPIService) TranqcontractGetCommunicationListV1Execute(r ApiTranqcontractGetCommunicationListV1Request) (*TranqcontractGetCommunicationListV1Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TranqcontractGetCommunicationListV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectTranqcontractAPIService.TranqcontractGetCommunicationListV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/tranqcontract/{pkiTranqcontractID}/getCommunicationList"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiTranqcontractID"+"}", url.PathEscape(parameterValueToString(r.pkiTranqcontractID, "pkiTranqcontractID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pkiTranqcontractID < 0 {
		return localVarReturnValue, nil, reportError("pkiTranqcontractID must be greater than 0")
	}
	if r.pkiTranqcontractID > 16777215 {
		return localVarReturnValue, nil, reportError("pkiTranqcontractID must be less than 16777215")
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

type ApiTranqcontractGetCommunicationrecipientsV1Request struct {
	ctx context.Context
	ApiService *ObjectTranqcontractAPIService
	pkiTranqcontractID int32
}

func (r ApiTranqcontractGetCommunicationrecipientsV1Request) Execute() (*TranqcontractGetCommunicationrecipientsV1Response, *http.Response, error) {
	return r.ApiService.TranqcontractGetCommunicationrecipientsV1Execute(r)
}

/*
TranqcontractGetCommunicationrecipientsV1 Retrieve Tranqcontract's Communicationrecipient



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pkiTranqcontractID
 @return ApiTranqcontractGetCommunicationrecipientsV1Request
*/
func (a *ObjectTranqcontractAPIService) TranqcontractGetCommunicationrecipientsV1(ctx context.Context, pkiTranqcontractID int32) ApiTranqcontractGetCommunicationrecipientsV1Request {
	return ApiTranqcontractGetCommunicationrecipientsV1Request{
		ApiService: a,
		ctx: ctx,
		pkiTranqcontractID: pkiTranqcontractID,
	}
}

// Execute executes the request
//  @return TranqcontractGetCommunicationrecipientsV1Response
func (a *ObjectTranqcontractAPIService) TranqcontractGetCommunicationrecipientsV1Execute(r ApiTranqcontractGetCommunicationrecipientsV1Request) (*TranqcontractGetCommunicationrecipientsV1Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TranqcontractGetCommunicationrecipientsV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectTranqcontractAPIService.TranqcontractGetCommunicationrecipientsV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/tranqcontract/{pkiTranqcontractID}/getCommunicationrecipients"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiTranqcontractID"+"}", url.PathEscape(parameterValueToString(r.pkiTranqcontractID, "pkiTranqcontractID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pkiTranqcontractID < 0 {
		return localVarReturnValue, nil, reportError("pkiTranqcontractID must be greater than 0")
	}
	if r.pkiTranqcontractID > 16777215 {
		return localVarReturnValue, nil, reportError("pkiTranqcontractID must be less than 16777215")
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

type ApiTranqcontractGetCommunicationsendersV1Request struct {
	ctx context.Context
	ApiService *ObjectTranqcontractAPIService
	pkiTranqcontractID int32
}

func (r ApiTranqcontractGetCommunicationsendersV1Request) Execute() (*TranqcontractGetCommunicationsendersV1Response, *http.Response, error) {
	return r.ApiService.TranqcontractGetCommunicationsendersV1Execute(r)
}

/*
TranqcontractGetCommunicationsendersV1 Retrieve Tranqcontract's Communicationsender



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pkiTranqcontractID
 @return ApiTranqcontractGetCommunicationsendersV1Request
*/
func (a *ObjectTranqcontractAPIService) TranqcontractGetCommunicationsendersV1(ctx context.Context, pkiTranqcontractID int32) ApiTranqcontractGetCommunicationsendersV1Request {
	return ApiTranqcontractGetCommunicationsendersV1Request{
		ApiService: a,
		ctx: ctx,
		pkiTranqcontractID: pkiTranqcontractID,
	}
}

// Execute executes the request
//  @return TranqcontractGetCommunicationsendersV1Response
func (a *ObjectTranqcontractAPIService) TranqcontractGetCommunicationsendersV1Execute(r ApiTranqcontractGetCommunicationsendersV1Request) (*TranqcontractGetCommunicationsendersV1Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TranqcontractGetCommunicationsendersV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectTranqcontractAPIService.TranqcontractGetCommunicationsendersV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/tranqcontract/{pkiTranqcontractID}/getCommunicationsenders"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiTranqcontractID"+"}", url.PathEscape(parameterValueToString(r.pkiTranqcontractID, "pkiTranqcontractID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pkiTranqcontractID < 0 {
		return localVarReturnValue, nil, reportError("pkiTranqcontractID must be greater than 0")
	}
	if r.pkiTranqcontractID > 16777215 {
		return localVarReturnValue, nil, reportError("pkiTranqcontractID must be less than 16777215")
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
