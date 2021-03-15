/*
 * eZmax API Definition
 *
 * This API expose all the functionnalities for the eZmax and eZsign applications.
 *
 * API version: 1.0.35
 * Contact: support-api@ezmax.ca
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxinc/eZmax-SDK-go

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// ObjectFranchiseofficeApiService ObjectFranchiseofficeApi service
type ObjectFranchiseofficeApiService service

type ApiFranchiseofficeGetAutocompleteV1Request struct {
	ctx _context.Context
	ApiService *ObjectFranchiseofficeApiService
	sSelector string
	sQuery *string
}

func (r ApiFranchiseofficeGetAutocompleteV1Request) SQuery(sQuery string) ApiFranchiseofficeGetAutocompleteV1Request {
	r.sQuery = &sQuery
	return r
}

func (r ApiFranchiseofficeGetAutocompleteV1Request) Execute() (CommonGetAutocompleteV1Response, *_nethttp.Response, error) {
	return r.ApiService.FranchiseofficeGetAutocompleteV1Execute(r)
}

/*
 * FranchiseofficeGetAutocompleteV1 Retrieve Franchiseoffices and IDs
 * Get the list of Franchiseoffices to be used in a dropdown or autocomplete control.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param sSelector The type of Franchiseoffices to return
 * @return ApiFranchiseofficeGetAutocompleteV1Request
 */
func (a *ObjectFranchiseofficeApiService) FranchiseofficeGetAutocompleteV1(ctx _context.Context, sSelector string) ApiFranchiseofficeGetAutocompleteV1Request {
	return ApiFranchiseofficeGetAutocompleteV1Request{
		ApiService: a,
		ctx: ctx,
		sSelector: sSelector,
	}
}

/*
 * Execute executes the request
 * @return CommonGetAutocompleteV1Response
 */
func (a *ObjectFranchiseofficeApiService) FranchiseofficeGetAutocompleteV1Execute(r ApiFranchiseofficeGetAutocompleteV1Request) (CommonGetAutocompleteV1Response, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CommonGetAutocompleteV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectFranchiseofficeApiService.FranchiseofficeGetAutocompleteV1")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/franchiseoffice/getAutocomplete/{sSelector}"
	localVarPath = strings.Replace(localVarPath, "{"+"sSelector"+"}", _neturl.PathEscape(parameterToString(r.sSelector, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.sQuery != nil {
		localVarQueryParams.Add("sQuery", parameterToString(*r.sQuery, ""))
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
