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


// ObjectInscriptionnotauthenticatedAPIService ObjectInscriptionnotauthenticatedAPI service
type ObjectInscriptionnotauthenticatedAPIService service

type ApiInscriptionnotauthenticatedGetCommunicationListV1Request struct {
	ctx context.Context
	ApiService *ObjectInscriptionnotauthenticatedAPIService
	pkiInscriptionnotauthenticatedID int32
}

func (r ApiInscriptionnotauthenticatedGetCommunicationListV1Request) Execute() (*InscriptionnotauthenticatedGetCommunicationListV1Response, *http.Response, error) {
	return r.ApiService.InscriptionnotauthenticatedGetCommunicationListV1Execute(r)
}

/*
InscriptionnotauthenticatedGetCommunicationListV1 Retrieve Communication list



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pkiInscriptionnotauthenticatedID
 @return ApiInscriptionnotauthenticatedGetCommunicationListV1Request
*/
func (a *ObjectInscriptionnotauthenticatedAPIService) InscriptionnotauthenticatedGetCommunicationListV1(ctx context.Context, pkiInscriptionnotauthenticatedID int32) ApiInscriptionnotauthenticatedGetCommunicationListV1Request {
	return ApiInscriptionnotauthenticatedGetCommunicationListV1Request{
		ApiService: a,
		ctx: ctx,
		pkiInscriptionnotauthenticatedID: pkiInscriptionnotauthenticatedID,
	}
}

// Execute executes the request
//  @return InscriptionnotauthenticatedGetCommunicationListV1Response
func (a *ObjectInscriptionnotauthenticatedAPIService) InscriptionnotauthenticatedGetCommunicationListV1Execute(r ApiInscriptionnotauthenticatedGetCommunicationListV1Request) (*InscriptionnotauthenticatedGetCommunicationListV1Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InscriptionnotauthenticatedGetCommunicationListV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectInscriptionnotauthenticatedAPIService.InscriptionnotauthenticatedGetCommunicationListV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/inscriptionnotauthenticated/{pkiInscriptionnotauthenticatedID}/getCommunicationList"
	localVarPath = strings.Replace(localVarPath, "{"+"pkiInscriptionnotauthenticatedID"+"}", url.PathEscape(parameterValueToString(r.pkiInscriptionnotauthenticatedID, "pkiInscriptionnotauthenticatedID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pkiInscriptionnotauthenticatedID < 0 {
		return localVarReturnValue, nil, reportError("pkiInscriptionnotauthenticatedID must be greater than 0")
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
