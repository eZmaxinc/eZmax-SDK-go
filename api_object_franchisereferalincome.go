/*
 * eZmax API Definition
 *
 * This API expose all the functionnalities for the eZmax and eZsign application.
 *
 * API version: 1.0.30
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
)

// Linger please
var (
	_ _context.Context
)

// ObjectFranchisereferalincomeApiService ObjectFranchisereferalincomeApi service
type ObjectFranchisereferalincomeApiService service

type ApiFranchisereferalincomeCreateObjectV1Request struct {
	ctx _context.Context
	ApiService *ObjectFranchisereferalincomeApiService
	franchisereferalincomeCreateObjectV1Request *[]FranchisereferalincomeCreateObjectV1Request
}

func (r ApiFranchisereferalincomeCreateObjectV1Request) FranchisereferalincomeCreateObjectV1Request(franchisereferalincomeCreateObjectV1Request []FranchisereferalincomeCreateObjectV1Request) ApiFranchisereferalincomeCreateObjectV1Request {
	r.franchisereferalincomeCreateObjectV1Request = &franchisereferalincomeCreateObjectV1Request
	return r
}

func (r ApiFranchisereferalincomeCreateObjectV1Request) Execute() (FranchisereferalincomeCreateObjectV1Response, *_nethttp.Response, error) {
	return r.ApiService.FranchisereferalincomeCreateObjectV1Execute(r)
}

/*
 * FranchisereferalincomeCreateObjectV1 Create a new Franchisereferalincome
 * The endpoint allows to create one or many elements at once.

The array can contain simple (Just the object) or compound (The object and its child) objects.

Creating compound elements allows to reduce the multiple requests to create all child objects.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiFranchisereferalincomeCreateObjectV1Request
 */
func (a *ObjectFranchisereferalincomeApiService) FranchisereferalincomeCreateObjectV1(ctx _context.Context) ApiFranchisereferalincomeCreateObjectV1Request {
	return ApiFranchisereferalincomeCreateObjectV1Request{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return FranchisereferalincomeCreateObjectV1Response
 */
func (a *ObjectFranchisereferalincomeApiService) FranchisereferalincomeCreateObjectV1Execute(r ApiFranchisereferalincomeCreateObjectV1Request) (FranchisereferalincomeCreateObjectV1Response, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  FranchisereferalincomeCreateObjectV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectFranchisereferalincomeApiService.FranchisereferalincomeCreateObjectV1")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/franchisereferalincome"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.franchisereferalincomeCreateObjectV1Request == nil {
		return localVarReturnValue, nil, reportError("franchisereferalincomeCreateObjectV1Request is required and must be specified")
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
	localVarPostBody = r.franchisereferalincomeCreateObjectV1Request
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
