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


// GlobalEzmaxclientAPIService GlobalEzmaxclientAPI service
type GlobalEzmaxclientAPIService service

type ApiGlobalEzmaxclientVersionV1Request struct {
	ctx context.Context
	ApiService *GlobalEzmaxclientAPIService
	pksEzmaxclientOs FieldPksEzmaxclientOs
}

func (r ApiGlobalEzmaxclientVersionV1Request) Execute() (*GlobalEzmaxclientVersionV1Response, *http.Response, error) {
	return r.ApiService.GlobalEzmaxclientVersionV1Execute(r)
}

/*
GlobalEzmaxclientVersionV1 Retrieve the latest version of the Ezmaxclient

Retrieve the latest version of the Ezmaxclient that is available on the store.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pksEzmaxclientOs
 @return ApiGlobalEzmaxclientVersionV1Request
*/
func (a *GlobalEzmaxclientAPIService) GlobalEzmaxclientVersionV1(ctx context.Context, pksEzmaxclientOs FieldPksEzmaxclientOs) ApiGlobalEzmaxclientVersionV1Request {
	return ApiGlobalEzmaxclientVersionV1Request{
		ApiService: a,
		ctx: ctx,
		pksEzmaxclientOs: pksEzmaxclientOs,
	}
}

// Execute executes the request
//  @return GlobalEzmaxclientVersionV1Response
func (a *GlobalEzmaxclientAPIService) GlobalEzmaxclientVersionV1Execute(r ApiGlobalEzmaxclientVersionV1Request) (*GlobalEzmaxclientVersionV1Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GlobalEzmaxclientVersionV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GlobalEzmaxclientAPIService.GlobalEzmaxclientVersionV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/ezmaxclient/{pksEzmaxclientOs}/version"
	localVarPath = strings.Replace(localVarPath, "{"+"pksEzmaxclientOs"+"}", url.PathEscape(parameterValueToString(r.pksEzmaxclientOs, "pksEzmaxclientOs")), -1)

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
