/*
eZmax API Definition (Full)

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.2.0
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


// ObjectInvoiceAPIService ObjectInvoiceAPI service
type ObjectInvoiceAPIService service

type ApiInvoiceGetAttachmentsV1Request struct {
	ctx context.Context
	ApiService *ObjectInvoiceAPIService
	pkiInvoiceID int32
}

func (r ApiInvoiceGetAttachmentsV1Request) Execute() (*InvoiceGetAttachmentsV1Response, *http.Response, error) {
	return r.ApiService.InvoiceGetAttachmentsV1Execute(r)
}

/*
InvoiceGetAttachmentsV1 Retrieve Invoice's Attachments



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pkiInvoiceID
 @return ApiInvoiceGetAttachmentsV1Request
*/
func (a *ObjectInvoiceAPIService) InvoiceGetAttachmentsV1(ctx context.Context, pkiInvoiceID int32) ApiInvoiceGetAttachmentsV1Request {
	return ApiInvoiceGetAttachmentsV1Request{
		ApiService: a,
		ctx: ctx,
		pkiInvoiceID: pkiInvoiceID,
	}
}

// Execute executes the request
//  @return InvoiceGetAttachmentsV1Response
func (a *ObjectInvoiceAPIService) InvoiceGetAttachmentsV1Execute(r ApiInvoiceGetAttachmentsV1Request) (*InvoiceGetAttachmentsV1Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InvoiceGetAttachmentsV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectInvoiceAPIService.InvoiceGetAttachmentsV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/invoice/{pkiInvoiceID}/getAttachments"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiInvoiceID"+"}", url.PathEscape(parameterValueToString(r.pkiInvoiceID, "pkiInvoiceID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pkiInvoiceID < 0 {
		return localVarReturnValue, nil, reportError("pkiInvoiceID must be greater than 0")
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

type ApiInvoiceGetCommunicationListV1Request struct {
	ctx context.Context
	ApiService *ObjectInvoiceAPIService
	pkiInvoiceID int32
}

func (r ApiInvoiceGetCommunicationListV1Request) Execute() (*InvoiceGetCommunicationListV1Response, *http.Response, error) {
	return r.ApiService.InvoiceGetCommunicationListV1Execute(r)
}

/*
InvoiceGetCommunicationListV1 Retrieve Communication list



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pkiInvoiceID
 @return ApiInvoiceGetCommunicationListV1Request
*/
func (a *ObjectInvoiceAPIService) InvoiceGetCommunicationListV1(ctx context.Context, pkiInvoiceID int32) ApiInvoiceGetCommunicationListV1Request {
	return ApiInvoiceGetCommunicationListV1Request{
		ApiService: a,
		ctx: ctx,
		pkiInvoiceID: pkiInvoiceID,
	}
}

// Execute executes the request
//  @return InvoiceGetCommunicationListV1Response
func (a *ObjectInvoiceAPIService) InvoiceGetCommunicationListV1Execute(r ApiInvoiceGetCommunicationListV1Request) (*InvoiceGetCommunicationListV1Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InvoiceGetCommunicationListV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectInvoiceAPIService.InvoiceGetCommunicationListV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/invoice/{pkiInvoiceID}/getCommunicationList"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiInvoiceID"+"}", url.PathEscape(parameterValueToString(r.pkiInvoiceID, "pkiInvoiceID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pkiInvoiceID < 0 {
		return localVarReturnValue, nil, reportError("pkiInvoiceID must be greater than 0")
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
