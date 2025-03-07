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


// ObjectOtherincomeAPIService ObjectOtherincomeAPI service
type ObjectOtherincomeAPIService service

type ApiOtherincomeGetCommunicationCountV1Request struct {
	ctx context.Context
	ApiService *ObjectOtherincomeAPIService
	pkiOtherincomeID int32
}

func (r ApiOtherincomeGetCommunicationCountV1Request) Execute() (*OtherincomeGetCommunicationCountV1Response, *http.Response, error) {
	return r.ApiService.OtherincomeGetCommunicationCountV1Execute(r)
}

/*
OtherincomeGetCommunicationCountV1 Retrieve Communication count



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pkiOtherincomeID
 @return ApiOtherincomeGetCommunicationCountV1Request
*/
func (a *ObjectOtherincomeAPIService) OtherincomeGetCommunicationCountV1(ctx context.Context, pkiOtherincomeID int32) ApiOtherincomeGetCommunicationCountV1Request {
	return ApiOtherincomeGetCommunicationCountV1Request{
		ApiService: a,
		ctx: ctx,
		pkiOtherincomeID: pkiOtherincomeID,
	}
}

// Execute executes the request
//  @return OtherincomeGetCommunicationCountV1Response
func (a *ObjectOtherincomeAPIService) OtherincomeGetCommunicationCountV1Execute(r ApiOtherincomeGetCommunicationCountV1Request) (*OtherincomeGetCommunicationCountV1Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OtherincomeGetCommunicationCountV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectOtherincomeAPIService.OtherincomeGetCommunicationCountV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/otherincome/{pkiOtherincomeID}/getCommunicationCount"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiOtherincomeID"+"}", url.PathEscape(parameterValueToString(r.pkiOtherincomeID, "pkiOtherincomeID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pkiOtherincomeID < 1 {
		return localVarReturnValue, nil, reportError("pkiOtherincomeID must be greater than 1")
	}
	if r.pkiOtherincomeID > 65535 {
		return localVarReturnValue, nil, reportError("pkiOtherincomeID must be less than 65535")
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

type ApiOtherincomeGetCommunicationListV1Request struct {
	ctx context.Context
	ApiService *ObjectOtherincomeAPIService
	pkiOtherincomeID int32
}

func (r ApiOtherincomeGetCommunicationListV1Request) Execute() (*OtherincomeGetCommunicationListV1Response, *http.Response, error) {
	return r.ApiService.OtherincomeGetCommunicationListV1Execute(r)
}

/*
OtherincomeGetCommunicationListV1 Retrieve Communication list



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pkiOtherincomeID
 @return ApiOtherincomeGetCommunicationListV1Request
*/
func (a *ObjectOtherincomeAPIService) OtherincomeGetCommunicationListV1(ctx context.Context, pkiOtherincomeID int32) ApiOtherincomeGetCommunicationListV1Request {
	return ApiOtherincomeGetCommunicationListV1Request{
		ApiService: a,
		ctx: ctx,
		pkiOtherincomeID: pkiOtherincomeID,
	}
}

// Execute executes the request
//  @return OtherincomeGetCommunicationListV1Response
func (a *ObjectOtherincomeAPIService) OtherincomeGetCommunicationListV1Execute(r ApiOtherincomeGetCommunicationListV1Request) (*OtherincomeGetCommunicationListV1Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OtherincomeGetCommunicationListV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectOtherincomeAPIService.OtherincomeGetCommunicationListV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/otherincome/{pkiOtherincomeID}/getCommunicationList"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiOtherincomeID"+"}", url.PathEscape(parameterValueToString(r.pkiOtherincomeID, "pkiOtherincomeID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pkiOtherincomeID < 1 {
		return localVarReturnValue, nil, reportError("pkiOtherincomeID must be greater than 1")
	}
	if r.pkiOtherincomeID > 65535 {
		return localVarReturnValue, nil, reportError("pkiOtherincomeID must be less than 65535")
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

type ApiOtherincomeGetCommunicationrecipientsV1Request struct {
	ctx context.Context
	ApiService *ObjectOtherincomeAPIService
	pkiOtherincomeID int32
}

func (r ApiOtherincomeGetCommunicationrecipientsV1Request) Execute() (*OtherincomeGetCommunicationrecipientsV1Response, *http.Response, error) {
	return r.ApiService.OtherincomeGetCommunicationrecipientsV1Execute(r)
}

/*
OtherincomeGetCommunicationrecipientsV1 Retrieve Otherincome's Communicationrecipient



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pkiOtherincomeID
 @return ApiOtherincomeGetCommunicationrecipientsV1Request
*/
func (a *ObjectOtherincomeAPIService) OtherincomeGetCommunicationrecipientsV1(ctx context.Context, pkiOtherincomeID int32) ApiOtherincomeGetCommunicationrecipientsV1Request {
	return ApiOtherincomeGetCommunicationrecipientsV1Request{
		ApiService: a,
		ctx: ctx,
		pkiOtherincomeID: pkiOtherincomeID,
	}
}

// Execute executes the request
//  @return OtherincomeGetCommunicationrecipientsV1Response
func (a *ObjectOtherincomeAPIService) OtherincomeGetCommunicationrecipientsV1Execute(r ApiOtherincomeGetCommunicationrecipientsV1Request) (*OtherincomeGetCommunicationrecipientsV1Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OtherincomeGetCommunicationrecipientsV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectOtherincomeAPIService.OtherincomeGetCommunicationrecipientsV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/otherincome/{pkiOtherincomeID}/getCommunicationrecipients"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiOtherincomeID"+"}", url.PathEscape(parameterValueToString(r.pkiOtherincomeID, "pkiOtherincomeID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pkiOtherincomeID < 1 {
		return localVarReturnValue, nil, reportError("pkiOtherincomeID must be greater than 1")
	}
	if r.pkiOtherincomeID > 65535 {
		return localVarReturnValue, nil, reportError("pkiOtherincomeID must be less than 65535")
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

type ApiOtherincomeGetCommunicationsendersV1Request struct {
	ctx context.Context
	ApiService *ObjectOtherincomeAPIService
	pkiOtherincomeID int32
}

func (r ApiOtherincomeGetCommunicationsendersV1Request) Execute() (*OtherincomeGetCommunicationsendersV1Response, *http.Response, error) {
	return r.ApiService.OtherincomeGetCommunicationsendersV1Execute(r)
}

/*
OtherincomeGetCommunicationsendersV1 Retrieve Otherincome's Communicationsender



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pkiOtherincomeID
 @return ApiOtherincomeGetCommunicationsendersV1Request
*/
func (a *ObjectOtherincomeAPIService) OtherincomeGetCommunicationsendersV1(ctx context.Context, pkiOtherincomeID int32) ApiOtherincomeGetCommunicationsendersV1Request {
	return ApiOtherincomeGetCommunicationsendersV1Request{
		ApiService: a,
		ctx: ctx,
		pkiOtherincomeID: pkiOtherincomeID,
	}
}

// Execute executes the request
//  @return OtherincomeGetCommunicationsendersV1Response
func (a *ObjectOtherincomeAPIService) OtherincomeGetCommunicationsendersV1Execute(r ApiOtherincomeGetCommunicationsendersV1Request) (*OtherincomeGetCommunicationsendersV1Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OtherincomeGetCommunicationsendersV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectOtherincomeAPIService.OtherincomeGetCommunicationsendersV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/otherincome/{pkiOtherincomeID}/getCommunicationsenders"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiOtherincomeID"+"}", url.PathEscape(parameterValueToString(r.pkiOtherincomeID, "pkiOtherincomeID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pkiOtherincomeID < 1 {
		return localVarReturnValue, nil, reportError("pkiOtherincomeID must be greater than 1")
	}
	if r.pkiOtherincomeID > 65535 {
		return localVarReturnValue, nil, reportError("pkiOtherincomeID must be less than 65535")
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
