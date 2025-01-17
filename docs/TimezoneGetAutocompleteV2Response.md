# TimezoneGetAutocompleteV2Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjDebugPayload** | [**CommonResponseObjDebugPayload**](CommonResponseObjDebugPayload.md) |  | 
**ObjDebug** | Pointer to [**CommonResponseObjDebug**](CommonResponseObjDebug.md) |  | [optional] 
**MPayload** | [**TimezoneGetAutocompleteV2ResponseMPayload**](TimezoneGetAutocompleteV2ResponseMPayload.md) |  | 

## Methods

### NewTimezoneGetAutocompleteV2Response

`func NewTimezoneGetAutocompleteV2Response(objDebugPayload CommonResponseObjDebugPayload, mPayload TimezoneGetAutocompleteV2ResponseMPayload, ) *TimezoneGetAutocompleteV2Response`

NewTimezoneGetAutocompleteV2Response instantiates a new TimezoneGetAutocompleteV2Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimezoneGetAutocompleteV2ResponseWithDefaults

`func NewTimezoneGetAutocompleteV2ResponseWithDefaults() *TimezoneGetAutocompleteV2Response`

NewTimezoneGetAutocompleteV2ResponseWithDefaults instantiates a new TimezoneGetAutocompleteV2Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjDebugPayload

`func (o *TimezoneGetAutocompleteV2Response) GetObjDebugPayload() CommonResponseObjDebugPayload`

GetObjDebugPayload returns the ObjDebugPayload field if non-nil, zero value otherwise.

### GetObjDebugPayloadOk

`func (o *TimezoneGetAutocompleteV2Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool)`

GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDebugPayload

`func (o *TimezoneGetAutocompleteV2Response) SetObjDebugPayload(v CommonResponseObjDebugPayload)`

SetObjDebugPayload sets ObjDebugPayload field to given value.


### GetObjDebug

`func (o *TimezoneGetAutocompleteV2Response) GetObjDebug() CommonResponseObjDebug`

GetObjDebug returns the ObjDebug field if non-nil, zero value otherwise.

### GetObjDebugOk

`func (o *TimezoneGetAutocompleteV2Response) GetObjDebugOk() (*CommonResponseObjDebug, bool)`

GetObjDebugOk returns a tuple with the ObjDebug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDebug

`func (o *TimezoneGetAutocompleteV2Response) SetObjDebug(v CommonResponseObjDebug)`

SetObjDebug sets ObjDebug field to given value.

### HasObjDebug

`func (o *TimezoneGetAutocompleteV2Response) HasObjDebug() bool`

HasObjDebug returns a boolean if a field has been set.

### GetMPayload

`func (o *TimezoneGetAutocompleteV2Response) GetMPayload() TimezoneGetAutocompleteV2ResponseMPayload`

GetMPayload returns the MPayload field if non-nil, zero value otherwise.

### GetMPayloadOk

`func (o *TimezoneGetAutocompleteV2Response) GetMPayloadOk() (*TimezoneGetAutocompleteV2ResponseMPayload, bool)`

GetMPayloadOk returns a tuple with the MPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMPayload

`func (o *TimezoneGetAutocompleteV2Response) SetMPayload(v TimezoneGetAutocompleteV2ResponseMPayload)`

SetMPayload sets MPayload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


