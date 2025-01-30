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


// ObjectRejectedoffertopurchaseAPIService ObjectRejectedoffertopurchaseAPI service
type ObjectRejectedoffertopurchaseAPIService service

type ApiRejectedoffertopurchaseGetCommunicationCountV1Request struct {
	ctx context.Context
	ApiService *ObjectRejectedoffertopurchaseAPIService
	pkiRejectedoffertopurchaseID int32
}

func (r ApiRejectedoffertopurchaseGetCommunicationCountV1Request) Execute() (*RejectedoffertopurchaseGetCommunicationCountV1Response, *http.Response, error) {
	return r.ApiService.RejectedoffertopurchaseGetCommunicationCountV1Execute(r)
}

/*
RejectedoffertopurchaseGetCommunicationCountV1 Retrieve Communication count



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pkiRejectedoffertopurchaseID
 @return ApiRejectedoffertopurchaseGetCommunicationCountV1Request
*/
func (a *ObjectRejectedoffertopurchaseAPIService) RejectedoffertopurchaseGetCommunicationCountV1(ctx context.Context, pkiRejectedoffertopurchaseID int32) ApiRejectedoffertopurchaseGetCommunicationCountV1Request {
	return ApiRejectedoffertopurchaseGetCommunicationCountV1Request{
		ApiService: a,
		ctx: ctx,
		pkiRejectedoffertopurchaseID: pkiRejectedoffertopurchaseID,
	}
}

// Execute executes the request
//  @return RejectedoffertopurchaseGetCommunicationCountV1Response
func (a *ObjectRejectedoffertopurchaseAPIService) RejectedoffertopurchaseGetCommunicationCountV1Execute(r ApiRejectedoffertopurchaseGetCommunicationCountV1Request) (*RejectedoffertopurchaseGetCommunicationCountV1Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RejectedoffertopurchaseGetCommunicationCountV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectRejectedoffertopurchaseAPIService.RejectedoffertopurchaseGetCommunicationCountV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/rejectedoffertopurchase/{pkiRejectedoffertopurchaseID}/getCommunicationCount"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiRejectedoffertopurchaseID"+"}", url.PathEscape(parameterValueToString(r.pkiRejectedoffertopurchaseID, "pkiRejectedoffertopurchaseID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pkiRejectedoffertopurchaseID < 1 {
		return localVarReturnValue, nil, reportError("pkiRejectedoffertopurchaseID must be greater than 1")
	}
	if r.pkiRejectedoffertopurchaseID > 65535 {
		return localVarReturnValue, nil, reportError("pkiRejectedoffertopurchaseID must be less than 65535")
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

type ApiRejectedoffertopurchaseGetCommunicationListV1Request struct {
	ctx context.Context
	ApiService *ObjectRejectedoffertopurchaseAPIService
	pkiRejectedoffertopurchaseID int32
}

func (r ApiRejectedoffertopurchaseGetCommunicationListV1Request) Execute() (*RejectedoffertopurchaseGetCommunicationListV1Response, *http.Response, error) {
	return r.ApiService.RejectedoffertopurchaseGetCommunicationListV1Execute(r)
}

/*
RejectedoffertopurchaseGetCommunicationListV1 Retrieve Communication list



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pkiRejectedoffertopurchaseID
 @return ApiRejectedoffertopurchaseGetCommunicationListV1Request
*/
func (a *ObjectRejectedoffertopurchaseAPIService) RejectedoffertopurchaseGetCommunicationListV1(ctx context.Context, pkiRejectedoffertopurchaseID int32) ApiRejectedoffertopurchaseGetCommunicationListV1Request {
	return ApiRejectedoffertopurchaseGetCommunicationListV1Request{
		ApiService: a,
		ctx: ctx,
		pkiRejectedoffertopurchaseID: pkiRejectedoffertopurchaseID,
	}
}

// Execute executes the request
//  @return RejectedoffertopurchaseGetCommunicationListV1Response
func (a *ObjectRejectedoffertopurchaseAPIService) RejectedoffertopurchaseGetCommunicationListV1Execute(r ApiRejectedoffertopurchaseGetCommunicationListV1Request) (*RejectedoffertopurchaseGetCommunicationListV1Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RejectedoffertopurchaseGetCommunicationListV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectRejectedoffertopurchaseAPIService.RejectedoffertopurchaseGetCommunicationListV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/rejectedoffertopurchase/{pkiRejectedoffertopurchaseID}/getCommunicationList"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiRejectedoffertopurchaseID"+"}", url.PathEscape(parameterValueToString(r.pkiRejectedoffertopurchaseID, "pkiRejectedoffertopurchaseID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pkiRejectedoffertopurchaseID < 1 {
		return localVarReturnValue, nil, reportError("pkiRejectedoffertopurchaseID must be greater than 1")
	}
	if r.pkiRejectedoffertopurchaseID > 65535 {
		return localVarReturnValue, nil, reportError("pkiRejectedoffertopurchaseID must be less than 65535")
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

type ApiRejectedoffertopurchaseGetCommunicationrecipientsV1Request struct {
	ctx context.Context
	ApiService *ObjectRejectedoffertopurchaseAPIService
	pkiRejectedoffertopurchaseID int32
}

func (r ApiRejectedoffertopurchaseGetCommunicationrecipientsV1Request) Execute() (*RejectedoffertopurchaseGetCommunicationrecipientsV1Response, *http.Response, error) {
	return r.ApiService.RejectedoffertopurchaseGetCommunicationrecipientsV1Execute(r)
}

/*
RejectedoffertopurchaseGetCommunicationrecipientsV1 Retrieve Rejectedoffertopurchase's Communicationrecipient



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pkiRejectedoffertopurchaseID
 @return ApiRejectedoffertopurchaseGetCommunicationrecipientsV1Request
*/
func (a *ObjectRejectedoffertopurchaseAPIService) RejectedoffertopurchaseGetCommunicationrecipientsV1(ctx context.Context, pkiRejectedoffertopurchaseID int32) ApiRejectedoffertopurchaseGetCommunicationrecipientsV1Request {
	return ApiRejectedoffertopurchaseGetCommunicationrecipientsV1Request{
		ApiService: a,
		ctx: ctx,
		pkiRejectedoffertopurchaseID: pkiRejectedoffertopurchaseID,
	}
}

// Execute executes the request
//  @return RejectedoffertopurchaseGetCommunicationrecipientsV1Response
func (a *ObjectRejectedoffertopurchaseAPIService) RejectedoffertopurchaseGetCommunicationrecipientsV1Execute(r ApiRejectedoffertopurchaseGetCommunicationrecipientsV1Request) (*RejectedoffertopurchaseGetCommunicationrecipientsV1Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RejectedoffertopurchaseGetCommunicationrecipientsV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectRejectedoffertopurchaseAPIService.RejectedoffertopurchaseGetCommunicationrecipientsV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/rejectedoffertopurchase/{pkiRejectedoffertopurchaseID}/getCommunicationrecipients"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiRejectedoffertopurchaseID"+"}", url.PathEscape(parameterValueToString(r.pkiRejectedoffertopurchaseID, "pkiRejectedoffertopurchaseID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pkiRejectedoffertopurchaseID < 1 {
		return localVarReturnValue, nil, reportError("pkiRejectedoffertopurchaseID must be greater than 1")
	}
	if r.pkiRejectedoffertopurchaseID > 65535 {
		return localVarReturnValue, nil, reportError("pkiRejectedoffertopurchaseID must be less than 65535")
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

type ApiRejectedoffertopurchaseGetCommunicationsendersV1Request struct {
	ctx context.Context
	ApiService *ObjectRejectedoffertopurchaseAPIService
	pkiRejectedoffertopurchaseID int32
}

func (r ApiRejectedoffertopurchaseGetCommunicationsendersV1Request) Execute() (*RejectedoffertopurchaseGetCommunicationsendersV1Response, *http.Response, error) {
	return r.ApiService.RejectedoffertopurchaseGetCommunicationsendersV1Execute(r)
}

/*
RejectedoffertopurchaseGetCommunicationsendersV1 Retrieve Rejectedoffertopurchase's Communicationsender



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pkiRejectedoffertopurchaseID
 @return ApiRejectedoffertopurchaseGetCommunicationsendersV1Request
*/
func (a *ObjectRejectedoffertopurchaseAPIService) RejectedoffertopurchaseGetCommunicationsendersV1(ctx context.Context, pkiRejectedoffertopurchaseID int32) ApiRejectedoffertopurchaseGetCommunicationsendersV1Request {
	return ApiRejectedoffertopurchaseGetCommunicationsendersV1Request{
		ApiService: a,
		ctx: ctx,
		pkiRejectedoffertopurchaseID: pkiRejectedoffertopurchaseID,
	}
}

// Execute executes the request
//  @return RejectedoffertopurchaseGetCommunicationsendersV1Response
func (a *ObjectRejectedoffertopurchaseAPIService) RejectedoffertopurchaseGetCommunicationsendersV1Execute(r ApiRejectedoffertopurchaseGetCommunicationsendersV1Request) (*RejectedoffertopurchaseGetCommunicationsendersV1Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RejectedoffertopurchaseGetCommunicationsendersV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectRejectedoffertopurchaseAPIService.RejectedoffertopurchaseGetCommunicationsendersV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/rejectedoffertopurchase/{pkiRejectedoffertopurchaseID}/getCommunicationsenders"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiRejectedoffertopurchaseID"+"}", url.PathEscape(parameterValueToString(r.pkiRejectedoffertopurchaseID, "pkiRejectedoffertopurchaseID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pkiRejectedoffertopurchaseID < 1 {
		return localVarReturnValue, nil, reportError("pkiRejectedoffertopurchaseID must be greater than 1")
	}
	if r.pkiRejectedoffertopurchaseID > 65535 {
		return localVarReturnValue, nil, reportError("pkiRejectedoffertopurchaseID must be less than 65535")
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
