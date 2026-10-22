# DocumenttypechecklistGetObjectV2Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjDebugPayload** | [**CommonResponseObjDebugPayload**](CommonResponseObjDebugPayload.md) |  | 
**ObjDebug** | Pointer to [**CommonResponseObjDebug**](CommonResponseObjDebug.md) |  | [optional] 
**MPayload** | [**DocumenttypechecklistGetObjectV2ResponseMPayload**](DocumenttypechecklistGetObjectV2ResponseMPayload.md) |  | 

## Methods

### NewDocumenttypechecklistGetObjectV2Response

`func NewDocumenttypechecklistGetObjectV2Response(objDebugPayload CommonResponseObjDebugPayload, mPayload DocumenttypechecklistGetObjectV2ResponseMPayload, ) *DocumenttypechecklistGetObjectV2Response`

NewDocumenttypechecklistGetObjectV2Response instantiates a new DocumenttypechecklistGetObjectV2Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumenttypechecklistGetObjectV2ResponseWithDefaults

`func NewDocumenttypechecklistGetObjectV2ResponseWithDefaults() *DocumenttypechecklistGetObjectV2Response`

NewDocumenttypechecklistGetObjectV2ResponseWithDefaults instantiates a new DocumenttypechecklistGetObjectV2Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjDebugPayload

`func (o *DocumenttypechecklistGetObjectV2Response) GetObjDebugPayload() CommonResponseObjDebugPayload`

GetObjDebugPayload returns the ObjDebugPayload field if non-nil, zero value otherwise.

### GetObjDebugPayloadOk

`func (o *DocumenttypechecklistGetObjectV2Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool)`

GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDebugPayload

`func (o *DocumenttypechecklistGetObjectV2Response) SetObjDebugPayload(v CommonResponseObjDebugPayload)`

SetObjDebugPayload sets ObjDebugPayload field to given value.


### GetObjDebug

`func (o *DocumenttypechecklistGetObjectV2Response) GetObjDebug() CommonResponseObjDebug`

GetObjDebug returns the ObjDebug field if non-nil, zero value otherwise.

### GetObjDebugOk

`func (o *DocumenttypechecklistGetObjectV2Response) GetObjDebugOk() (*CommonResponseObjDebug, bool)`

GetObjDebugOk returns a tuple with the ObjDebug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDebug

`func (o *DocumenttypechecklistGetObjectV2Response) SetObjDebug(v CommonResponseObjDebug)`

SetObjDebug sets ObjDebug field to given value.

### HasObjDebug

`func (o *DocumenttypechecklistGetObjectV2Response) HasObjDebug() bool`

HasObjDebug returns a boolean if a field has been set.

### GetMPayload

`func (o *DocumenttypechecklistGetObjectV2Response) GetMPayload() DocumenttypechecklistGetObjectV2ResponseMPayload`

GetMPayload returns the MPayload field if non-nil, zero value otherwise.

### GetMPayloadOk

`func (o *DocumenttypechecklistGetObjectV2Response) GetMPayloadOk() (*DocumenttypechecklistGetObjectV2ResponseMPayload, bool)`

GetMPayloadOk returns a tuple with the MPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMPayload

`func (o *DocumenttypechecklistGetObjectV2Response) SetMPayload(v DocumenttypechecklistGetObjectV2ResponseMPayload)`

SetMPayload sets MPayload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


