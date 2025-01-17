# CompanyGetAutocompleteV2Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjDebugPayload** | [**CommonResponseObjDebugPayload**](CommonResponseObjDebugPayload.md) |  | 
**ObjDebug** | Pointer to [**CommonResponseObjDebug**](CommonResponseObjDebug.md) |  | [optional] 
**MPayload** | [**CompanyGetAutocompleteV2ResponseMPayload**](CompanyGetAutocompleteV2ResponseMPayload.md) |  | 

## Methods

### NewCompanyGetAutocompleteV2Response

`func NewCompanyGetAutocompleteV2Response(objDebugPayload CommonResponseObjDebugPayload, mPayload CompanyGetAutocompleteV2ResponseMPayload, ) *CompanyGetAutocompleteV2Response`

NewCompanyGetAutocompleteV2Response instantiates a new CompanyGetAutocompleteV2Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyGetAutocompleteV2ResponseWithDefaults

`func NewCompanyGetAutocompleteV2ResponseWithDefaults() *CompanyGetAutocompleteV2Response`

NewCompanyGetAutocompleteV2ResponseWithDefaults instantiates a new CompanyGetAutocompleteV2Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjDebugPayload

`func (o *CompanyGetAutocompleteV2Response) GetObjDebugPayload() CommonResponseObjDebugPayload`

GetObjDebugPayload returns the ObjDebugPayload field if non-nil, zero value otherwise.

### GetObjDebugPayloadOk

`func (o *CompanyGetAutocompleteV2Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool)`

GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDebugPayload

`func (o *CompanyGetAutocompleteV2Response) SetObjDebugPayload(v CommonResponseObjDebugPayload)`

SetObjDebugPayload sets ObjDebugPayload field to given value.


### GetObjDebug

`func (o *CompanyGetAutocompleteV2Response) GetObjDebug() CommonResponseObjDebug`

GetObjDebug returns the ObjDebug field if non-nil, zero value otherwise.

### GetObjDebugOk

`func (o *CompanyGetAutocompleteV2Response) GetObjDebugOk() (*CommonResponseObjDebug, bool)`

GetObjDebugOk returns a tuple with the ObjDebug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDebug

`func (o *CompanyGetAutocompleteV2Response) SetObjDebug(v CommonResponseObjDebug)`

SetObjDebug sets ObjDebug field to given value.

### HasObjDebug

`func (o *CompanyGetAutocompleteV2Response) HasObjDebug() bool`

HasObjDebug returns a boolean if a field has been set.

### GetMPayload

`func (o *CompanyGetAutocompleteV2Response) GetMPayload() CompanyGetAutocompleteV2ResponseMPayload`

GetMPayload returns the MPayload field if non-nil, zero value otherwise.

### GetMPayloadOk

`func (o *CompanyGetAutocompleteV2Response) GetMPayloadOk() (*CompanyGetAutocompleteV2ResponseMPayload, bool)`

GetMPayloadOk returns a tuple with the MPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMPayload

`func (o *CompanyGetAutocompleteV2Response) SetMPayload(v CompanyGetAutocompleteV2ResponseMPayload)`

SetMPayload sets MPayload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


