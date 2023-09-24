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
)


// ScimServiceProviderConfigAPIService ScimServiceProviderConfigAPI service
type ScimServiceProviderConfigAPIService service

type ApiServiceProviderConfigGetObjectScimV2Request struct {
	ctx context.Context
	ApiService *ScimServiceProviderConfigAPIService
}

func (r ApiServiceProviderConfigGetObjectScimV2Request) Execute() (*ScimServiceProviderConfig, *http.Response, error) {
	return r.ApiService.ServiceProviderConfigGetObjectScimV2Execute(r)
}

/*
ServiceProviderConfigGetObjectScimV2 Get Service Provider Configuration

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiServiceProviderConfigGetObjectScimV2Request
*/
func (a *ScimServiceProviderConfigAPIService) ServiceProviderConfigGetObjectScimV2(ctx context.Context) ApiServiceProviderConfigGetObjectScimV2Request {
	return ApiServiceProviderConfigGetObjectScimV2Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ScimServiceProviderConfig
func (a *ScimServiceProviderConfigAPIService) ServiceProviderConfigGetObjectScimV2Execute(r ApiServiceProviderConfigGetObjectScimV2Request) (*ScimServiceProviderConfig, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ScimServiceProviderConfig
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ScimServiceProviderConfigAPIService.ServiceProviderConfigGetObjectScimV2")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/2/scim/ServiceProviderConfig"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
