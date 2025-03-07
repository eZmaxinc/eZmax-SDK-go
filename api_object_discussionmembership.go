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


// ObjectDiscussionmembershipAPIService ObjectDiscussionmembershipAPI service
type ObjectDiscussionmembershipAPIService service

type ApiDiscussionmembershipCreateObjectV1Request struct {
	ctx context.Context
	ApiService *ObjectDiscussionmembershipAPIService
	discussionmembershipCreateObjectV1Request *DiscussionmembershipCreateObjectV1Request
}

func (r ApiDiscussionmembershipCreateObjectV1Request) DiscussionmembershipCreateObjectV1Request(discussionmembershipCreateObjectV1Request DiscussionmembershipCreateObjectV1Request) ApiDiscussionmembershipCreateObjectV1Request {
	r.discussionmembershipCreateObjectV1Request = &discussionmembershipCreateObjectV1Request
	return r
}

func (r ApiDiscussionmembershipCreateObjectV1Request) Execute() (*DiscussionmembershipCreateObjectV1Response, *http.Response, error) {
	return r.ApiService.DiscussionmembershipCreateObjectV1Execute(r)
}

/*
DiscussionmembershipCreateObjectV1 Create a new Discussionmembership

The endpoint allows to create one or many elements at once.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDiscussionmembershipCreateObjectV1Request
*/
func (a *ObjectDiscussionmembershipAPIService) DiscussionmembershipCreateObjectV1(ctx context.Context) ApiDiscussionmembershipCreateObjectV1Request {
	return ApiDiscussionmembershipCreateObjectV1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return DiscussionmembershipCreateObjectV1Response
func (a *ObjectDiscussionmembershipAPIService) DiscussionmembershipCreateObjectV1Execute(r ApiDiscussionmembershipCreateObjectV1Request) (*DiscussionmembershipCreateObjectV1Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DiscussionmembershipCreateObjectV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectDiscussionmembershipAPIService.DiscussionmembershipCreateObjectV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/discussionmembership"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.discussionmembershipCreateObjectV1Request == nil {
		return localVarReturnValue, nil, reportError("discussionmembershipCreateObjectV1Request is required and must be specified")
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
	localVarPostBody = r.discussionmembershipCreateObjectV1Request
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

type ApiDiscussionmembershipDeleteObjectV1Request struct {
	ctx context.Context
	ApiService *ObjectDiscussionmembershipAPIService
	pkiDiscussionmembershipID int32
}

func (r ApiDiscussionmembershipDeleteObjectV1Request) Execute() (*DiscussionmembershipDeleteObjectV1Response, *http.Response, error) {
	return r.ApiService.DiscussionmembershipDeleteObjectV1Execute(r)
}

/*
DiscussionmembershipDeleteObjectV1 Delete an existing Discussionmembership



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pkiDiscussionmembershipID The unique ID of the Discussionmembership
 @return ApiDiscussionmembershipDeleteObjectV1Request
*/
func (a *ObjectDiscussionmembershipAPIService) DiscussionmembershipDeleteObjectV1(ctx context.Context, pkiDiscussionmembershipID int32) ApiDiscussionmembershipDeleteObjectV1Request {
	return ApiDiscussionmembershipDeleteObjectV1Request{
		ApiService: a,
		ctx: ctx,
		pkiDiscussionmembershipID: pkiDiscussionmembershipID,
	}
}

// Execute executes the request
//  @return DiscussionmembershipDeleteObjectV1Response
func (a *ObjectDiscussionmembershipAPIService) DiscussionmembershipDeleteObjectV1Execute(r ApiDiscussionmembershipDeleteObjectV1Request) (*DiscussionmembershipDeleteObjectV1Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DiscussionmembershipDeleteObjectV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectDiscussionmembershipAPIService.DiscussionmembershipDeleteObjectV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/discussionmembership/{pkiDiscussionmembershipID}"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiDiscussionmembershipID"+"}", url.PathEscape(parameterValueToString(r.pkiDiscussionmembershipID, "pkiDiscussionmembershipID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pkiDiscussionmembershipID < 0 {
		return localVarReturnValue, nil, reportError("pkiDiscussionmembershipID must be greater than 0")
	}
	if r.pkiDiscussionmembershipID > 16777215 {
		return localVarReturnValue, nil, reportError("pkiDiscussionmembershipID must be less than 16777215")
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
