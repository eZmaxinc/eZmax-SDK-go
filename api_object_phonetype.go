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


// ObjectPhonetypeAPIService ObjectPhonetypeAPI service
type ObjectPhonetypeAPIService service

type ApiPhonetypeGetAutocompleteV2Request struct {
	ctx context.Context
	ApiService *ObjectPhonetypeAPIService
	sSelector string
	eFilterActive *string
	sQuery *string
	acceptLanguage *HeaderAcceptLanguage
}

// Specify which results we want to display.
func (r ApiPhonetypeGetAutocompleteV2Request) EFilterActive(eFilterActive string) ApiPhonetypeGetAutocompleteV2Request {
	r.eFilterActive = &eFilterActive
	return r
}

// Allow to filter the returned results
func (r ApiPhonetypeGetAutocompleteV2Request) SQuery(sQuery string) ApiPhonetypeGetAutocompleteV2Request {
	r.sQuery = &sQuery
	return r
}

func (r ApiPhonetypeGetAutocompleteV2Request) AcceptLanguage(acceptLanguage HeaderAcceptLanguage) ApiPhonetypeGetAutocompleteV2Request {
	r.acceptLanguage = &acceptLanguage
	return r
}

func (r ApiPhonetypeGetAutocompleteV2Request) Execute() (*PhonetypeGetAutocompleteV2Response, *http.Response, error) {
	return r.ApiService.PhonetypeGetAutocompleteV2Execute(r)
}

/*
PhonetypeGetAutocompleteV2 Retrieve Phonetypes and IDs

Get the list of Phonetype to be used in a dropdown or autocomplete control.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param sSelector The type of Phonetypes to return
 @return ApiPhonetypeGetAutocompleteV2Request
*/
func (a *ObjectPhonetypeAPIService) PhonetypeGetAutocompleteV2(ctx context.Context, sSelector string) ApiPhonetypeGetAutocompleteV2Request {
	return ApiPhonetypeGetAutocompleteV2Request{
		ApiService: a,
		ctx: ctx,
		sSelector: sSelector,
	}
}

// Execute executes the request
//  @return PhonetypeGetAutocompleteV2Response
func (a *ObjectPhonetypeAPIService) PhonetypeGetAutocompleteV2Execute(r ApiPhonetypeGetAutocompleteV2Request) (*PhonetypeGetAutocompleteV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PhonetypeGetAutocompleteV2Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectPhonetypeAPIService.PhonetypeGetAutocompleteV2")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/2/object/phonetype/getAutocomplete/{sSelector}"
	localVarPath = strings.Replace(localVarPath, "{"+"sSelector"+"}", url.PathEscape(parameterValueToString(r.sSelector, "sSelector")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.eFilterActive != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "eFilterActive", r.eFilterActive, "form", "")
	} else {
		var defaultValue string = "Active"
		r.eFilterActive = &defaultValue
	}
	if r.sQuery != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sQuery", r.sQuery, "form", "")
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
