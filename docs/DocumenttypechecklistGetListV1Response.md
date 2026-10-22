# DocumenttypechecklistGetListV1Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjDebugPayload** | [**CommonResponseObjDebugPayloadGetList**](CommonResponseObjDebugPayloadGetList.md) |  | 
**ObjDebug** | Pointer to [**CommonResponseObjDebug**](CommonResponseObjDebug.md) |  | [optional] 
**MPayload** | [**DocumenttypechecklistGetListV1ResponseMPayload**](DocumenttypechecklistGetListV1ResponseMPayload.md) |  | 

## Methods

### NewDocumenttypechecklistGetListV1Response

`func NewDocumenttypechecklistGetListV1Response(objDebugPayload CommonResponseObjDebugPayloadGetList, mPayload DocumenttypechecklistGetListV1ResponseMPayload, ) *DocumenttypechecklistGetListV1Response`

NewDocumenttypechecklistGetListV1Response instantiates a new DocumenttypechecklistGetListV1Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumenttypechecklistGetListV1ResponseWithDefaults

`func NewDocumenttypechecklistGetListV1ResponseWithDefaults() *DocumenttypechecklistGetListV1Response`

NewDocumenttypechecklistGetListV1ResponseWithDefaults instantiates a new DocumenttypechecklistGetListV1Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjDebugPayload

`func (o *DocumenttypechecklistGetListV1Response) GetObjDebugPayload() CommonResponseObjDebugPayloadGetList`

GetObjDebugPayload returns the ObjDebugPayload field if non-nil, zero value otherwise.

### GetObjDebugPayloadOk

`func (o *DocumenttypechecklistGetListV1Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayloadGetList, bool)`

GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDebugPayload

`func (o *DocumenttypechecklistGetListV1Response) SetObjDebugPayload(v CommonResponseObjDebugPayloadGetList)`

SetObjDebugPayload sets ObjDebugPayload field to given value.


### GetObjDebug

`func (o *DocumenttypechecklistGetListV1Response) GetObjDebug() CommonResponseObjDebug`

GetObjDebug returns the ObjDebug field if non-nil, zero value otherwise.

### GetObjDebugOk

`func (o *DocumenttypechecklistGetListV1Response) GetObjDebugOk() (*CommonResponseObjDebug, bool)`

GetObjDebugOk returns a tuple with the ObjDebug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDebug

`func (o *DocumenttypechecklistGetListV1Response) SetObjDebug(v CommonResponseObjDebug)`

SetObjDebug sets ObjDebug field to given value.

### HasObjDebug

`func (o *DocumenttypechecklistGetListV1Response) HasObjDebug() bool`

HasObjDebug returns a boolean if a field has been set.

### GetMPayload

`func (o *DocumenttypechecklistGetListV1Response) GetMPayload() DocumenttypechecklistGetListV1ResponseMPayload`

GetMPayload returns the MPayload field if non-nil, zero value otherwise.

### GetMPayloadOk

`func (o *DocumenttypechecklistGetListV1Response) GetMPayloadOk() (*DocumenttypechecklistGetListV1ResponseMPayload, bool)`

GetMPayloadOk returns a tuple with the MPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMPayload

`func (o *DocumenttypechecklistGetListV1Response) SetMPayload(v DocumenttypechecklistGetListV1ResponseMPayload)`

SetMPayload sets MPayload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


