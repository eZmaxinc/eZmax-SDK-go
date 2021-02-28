# ApikeyCreateObjectV1Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MPayload** | [**ApikeyCreateObjectV1ResponseMPayload**](ApikeyCreateObjectV1ResponseMPayload.md) |  | 
**ObjDebugPayload** | Pointer to [**CommonResponseObjDebugPayload**](CommonResponseObjDebugPayload.md) |  | [optional] 
**ObjDebug** | Pointer to [**CommonResponseObjDebug**](CommonResponseObjDebug.md) |  | [optional] 

## Methods

### NewApikeyCreateObjectV1Response

`func NewApikeyCreateObjectV1Response(mPayload ApikeyCreateObjectV1ResponseMPayload, ) *ApikeyCreateObjectV1Response`

NewApikeyCreateObjectV1Response instantiates a new ApikeyCreateObjectV1Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApikeyCreateObjectV1ResponseWithDefaults

`func NewApikeyCreateObjectV1ResponseWithDefaults() *ApikeyCreateObjectV1Response`

NewApikeyCreateObjectV1ResponseWithDefaults instantiates a new ApikeyCreateObjectV1Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMPayload

`func (o *ApikeyCreateObjectV1Response) GetMPayload() ApikeyCreateObjectV1ResponseMPayload`

GetMPayload returns the MPayload field if non-nil, zero value otherwise.

### GetMPayloadOk

`func (o *ApikeyCreateObjectV1Response) GetMPayloadOk() (*ApikeyCreateObjectV1ResponseMPayload, bool)`

GetMPayloadOk returns a tuple with the MPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMPayload

`func (o *ApikeyCreateObjectV1Response) SetMPayload(v ApikeyCreateObjectV1ResponseMPayload)`

SetMPayload sets MPayload field to given value.


### GetObjDebugPayload

`func (o *ApikeyCreateObjectV1Response) GetObjDebugPayload() CommonResponseObjDebugPayload`

GetObjDebugPayload returns the ObjDebugPayload field if non-nil, zero value otherwise.

### GetObjDebugPayloadOk

`func (o *ApikeyCreateObjectV1Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool)`

GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDebugPayload

`func (o *ApikeyCreateObjectV1Response) SetObjDebugPayload(v CommonResponseObjDebugPayload)`

SetObjDebugPayload sets ObjDebugPayload field to given value.

### HasObjDebugPayload

`func (o *ApikeyCreateObjectV1Response) HasObjDebugPayload() bool`

HasObjDebugPayload returns a boolean if a field has been set.

### GetObjDebug

`func (o *ApikeyCreateObjectV1Response) GetObjDebug() CommonResponseObjDebug`

GetObjDebug returns the ObjDebug field if non-nil, zero value otherwise.

### GetObjDebugOk

`func (o *ApikeyCreateObjectV1Response) GetObjDebugOk() (*CommonResponseObjDebug, bool)`

GetObjDebugOk returns a tuple with the ObjDebug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDebug

`func (o *ApikeyCreateObjectV1Response) SetObjDebug(v CommonResponseObjDebug)`

SetObjDebug sets ObjDebug field to given value.

### HasObjDebug

`func (o *ApikeyCreateObjectV1Response) HasObjDebug() bool`

HasObjDebug returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


