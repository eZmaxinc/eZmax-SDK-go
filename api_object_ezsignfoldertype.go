/*
eZmax API Definition

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.1.2
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

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

// ObjectEzsignfoldertypeApiService ObjectEzsignfoldertypeApi service
type ObjectEzsignfoldertypeApiService service

type ApiEzsignfoldertypeGetAutocompleteV1Request struct {
	ctx _context.Context
	ApiService *ObjectEzsignfoldertypeApiService
	sSelector string
	sQuery *string
	acceptLanguage *HeaderAcceptLanguage
}

// Allow to filter the returned results
func (r ApiEzsignfoldertypeGetAutocompleteV1Request) SQuery(sQuery string) ApiEzsignfoldertypeGetAutocompleteV1Request {
	r.sQuery = &sQuery
	return r
}
func (r ApiEzsignfoldertypeGetAutocompleteV1Request) AcceptLanguage(acceptLanguage HeaderAcceptLanguage) ApiEzsignfoldertypeGetAutocompleteV1Request {
	r.acceptLanguage = &acceptLanguage
	return r
}

func (r ApiEzsignfoldertypeGetAutocompleteV1Request) Execute() (CommonGetAutocompleteV1Response, *_nethttp.Response, error) {
	return r.ApiService.EzsignfoldertypeGetAutocompleteV1Execute(r)
}

/*
EzsignfoldertypeGetAutocompleteV1 Retrieve Ezsignfoldertypes and IDs

Get the list of Ezsignfoldertypes to be used in a dropdown or autocomplete control.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param sSelector The type of Ezsignfoldertypes to return
 @return ApiEzsignfoldertypeGetAutocompleteV1Request
*/
func (a *ObjectEzsignfoldertypeApiService) EzsignfoldertypeGetAutocompleteV1(ctx _context.Context, sSelector string) ApiEzsignfoldertypeGetAutocompleteV1Request {
	return ApiEzsignfoldertypeGetAutocompleteV1Request{
		ApiService: a,
		ctx: ctx,
		sSelector: sSelector,
	}
}

// Execute executes the request
//  @return CommonGetAutocompleteV1Response
func (a *ObjectEzsignfoldertypeApiService) EzsignfoldertypeGetAutocompleteV1Execute(r ApiEzsignfoldertypeGetAutocompleteV1Request) (CommonGetAutocompleteV1Response, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CommonGetAutocompleteV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectEzsignfoldertypeApiService.EzsignfoldertypeGetAutocompleteV1")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/ezsignfoldertype/getAutocomplete/{sSelector}"
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
	if r.acceptLanguage != nil {
		localVarHeaderParams["Accept-Language"] = parameterToString(*r.acceptLanguage, "")
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

type ApiEzsignfoldertypeGetListV1Request struct {
	ctx _context.Context
	ApiService *ObjectEzsignfoldertypeApiService
	eOrderBy *string
	iRowMax *int32
	iRowOffset *int32
	acceptLanguage *HeaderAcceptLanguage
	sFilter *string
}

// Specify how you want the results to be sorted
func (r ApiEzsignfoldertypeGetListV1Request) EOrderBy(eOrderBy string) ApiEzsignfoldertypeGetListV1Request {
	r.eOrderBy = &eOrderBy
	return r
}
func (r ApiEzsignfoldertypeGetListV1Request) IRowMax(iRowMax int32) ApiEzsignfoldertypeGetListV1Request {
	r.iRowMax = &iRowMax
	return r
}
func (r ApiEzsignfoldertypeGetListV1Request) IRowOffset(iRowOffset int32) ApiEzsignfoldertypeGetListV1Request {
	r.iRowOffset = &iRowOffset
	return r
}
func (r ApiEzsignfoldertypeGetListV1Request) AcceptLanguage(acceptLanguage HeaderAcceptLanguage) ApiEzsignfoldertypeGetListV1Request {
	r.acceptLanguage = &acceptLanguage
	return r
}
func (r ApiEzsignfoldertypeGetListV1Request) SFilter(sFilter string) ApiEzsignfoldertypeGetListV1Request {
	r.sFilter = &sFilter
	return r
}

func (r ApiEzsignfoldertypeGetListV1Request) Execute() (EzsignfoldertypeGetListV1Response, *_nethttp.Response, error) {
	return r.ApiService.EzsignfoldertypeGetListV1Execute(r)
}

/*
EzsignfoldertypeGetListV1 Retrieve Ezsignfoldertype list

Enum values that can be filtered in query parameter *sFilter*:

| Variable | Valid values |
|---|---|
| eEzsignfoldertypePrivacylevel | User<br>Usergroup |

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiEzsignfoldertypeGetListV1Request
*/
func (a *ObjectEzsignfoldertypeApiService) EzsignfoldertypeGetListV1(ctx _context.Context) ApiEzsignfoldertypeGetListV1Request {
	return ApiEzsignfoldertypeGetListV1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return EzsignfoldertypeGetListV1Response
func (a *ObjectEzsignfoldertypeApiService) EzsignfoldertypeGetListV1Execute(r ApiEzsignfoldertypeGetListV1Request) (EzsignfoldertypeGetListV1Response, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  EzsignfoldertypeGetListV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectEzsignfoldertypeApiService.EzsignfoldertypeGetListV1")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/ezsignfoldertype/getList"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.eOrderBy != nil {
		localVarQueryParams.Add("eOrderBy", parameterToString(*r.eOrderBy, ""))
	}
	if r.iRowMax != nil {
		localVarQueryParams.Add("iRowMax", parameterToString(*r.iRowMax, ""))
	}
	if r.iRowOffset != nil {
		localVarQueryParams.Add("iRowOffset", parameterToString(*r.iRowOffset, ""))
	}
	if r.sFilter != nil {
		localVarQueryParams.Add("sFilter", parameterToString(*r.sFilter, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.acceptLanguage != nil {
		localVarHeaderParams["Accept-Language"] = parameterToString(*r.acceptLanguage, "")
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
		if localVarHTTPResponse.StatusCode == 406 {
			var v CommonResponseError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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
