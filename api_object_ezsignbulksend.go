/*
eZmax API Definition

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.1.3
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
)

// Linger please
var (
	_ _context.Context
)

// ObjectEzsignbulksendApiService ObjectEzsignbulksendApi service
type ObjectEzsignbulksendApiService service

type ApiEzsignbulksendGetListV1Request struct {
	ctx _context.Context
	ApiService *ObjectEzsignbulksendApiService
	eOrderBy *string
	iRowMax *int32
	iRowOffset *int32
	acceptLanguage *HeaderAcceptLanguage
	sFilter *string
}

// Specify how you want the results to be sorted
func (r ApiEzsignbulksendGetListV1Request) EOrderBy(eOrderBy string) ApiEzsignbulksendGetListV1Request {
	r.eOrderBy = &eOrderBy
	return r
}
func (r ApiEzsignbulksendGetListV1Request) IRowMax(iRowMax int32) ApiEzsignbulksendGetListV1Request {
	r.iRowMax = &iRowMax
	return r
}
func (r ApiEzsignbulksendGetListV1Request) IRowOffset(iRowOffset int32) ApiEzsignbulksendGetListV1Request {
	r.iRowOffset = &iRowOffset
	return r
}
func (r ApiEzsignbulksendGetListV1Request) AcceptLanguage(acceptLanguage HeaderAcceptLanguage) ApiEzsignbulksendGetListV1Request {
	r.acceptLanguage = &acceptLanguage
	return r
}
func (r ApiEzsignbulksendGetListV1Request) SFilter(sFilter string) ApiEzsignbulksendGetListV1Request {
	r.sFilter = &sFilter
	return r
}

func (r ApiEzsignbulksendGetListV1Request) Execute() (EzsignbulksendGetListV1Response, *_nethttp.Response, error) {
	return r.ApiService.EzsignbulksendGetListV1Execute(r)
}

/*
EzsignbulksendGetListV1 Retrieve Ezsignbulksend list

Enum values that can be filtered in query parameter *sFilter*:

| Variable | Valid values |
|---|---|
| eEzsignfoldertypePrivacylevel | User<br>Usergroup |

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiEzsignbulksendGetListV1Request
*/
func (a *ObjectEzsignbulksendApiService) EzsignbulksendGetListV1(ctx _context.Context) ApiEzsignbulksendGetListV1Request {
	return ApiEzsignbulksendGetListV1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return EzsignbulksendGetListV1Response
func (a *ObjectEzsignbulksendApiService) EzsignbulksendGetListV1Execute(r ApiEzsignbulksendGetListV1Request) (EzsignbulksendGetListV1Response, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  EzsignbulksendGetListV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectEzsignbulksendApiService.EzsignbulksendGetListV1")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/object/ezsignbulksend/getList"

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
