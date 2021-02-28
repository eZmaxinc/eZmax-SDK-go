# EzsignsignatureGetObjectV1Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MPayload** | **map[string]interface{}** | Payload for the /1/object/ezsignsignature/getObject API Request | 
**ObjDebugPayload** | Pointer to [**CommonResponseObjDebugPayload**](CommonResponseObjDebugPayload.md) |  | [optional] 
**ObjDebug** | Pointer to [**CommonResponseObjDebug**](CommonResponseObjDebug.md) |  | [optional] 

## Methods

### NewEzsignsignatureGetObjectV1Response

`func NewEzsignsignatureGetObjectV1Response(mPayload map[string]interface{}, ) *EzsignsignatureGetObjectV1Response`

NewEzsignsignatureGetObjectV1Response instantiates a new EzsignsignatureGetObjectV1Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignsignatureGetObjectV1ResponseWithDefaults

`func NewEzsignsignatureGetObjectV1ResponseWithDefaults() *EzsignsignatureGetObjectV1Response`

NewEzsignsignatureGetObjectV1ResponseWithDefaults instantiates a new EzsignsignatureGetObjectV1Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMPayload

`func (o *EzsignsignatureGetObjectV1Response) GetMPayload() map[string]interface{}`

GetMPayload returns the MPayload field if non-nil, zero value otherwise.

### GetMPayloadOk

`func (o *EzsignsignatureGetObjectV1Response) GetMPayloadOk() (*map[string]interface{}, bool)`

GetMPayloadOk returns a tuple with the MPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMPayload

`func (o *EzsignsignatureGetObjectV1Response) SetMPayload(v map[string]interface{})`

SetMPayload sets MPayload field to given value.


### GetObjDebugPayload

`func (o *EzsignsignatureGetObjectV1Response) GetObjDebugPayload() CommonResponseObjDebugPayload`

GetObjDebugPayload returns the ObjDebugPayload field if non-nil, zero value otherwise.

### GetObjDebugPayloadOk

`func (o *EzsignsignatureGetObjectV1Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool)`

GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDebugPayload

`func (o *EzsignsignatureGetObjectV1Response) SetObjDebugPayload(v CommonResponseObjDebugPayload)`

SetObjDebugPayload sets ObjDebugPayload field to given value.

### HasObjDebugPayload

`func (o *EzsignsignatureGetObjectV1Response) HasObjDebugPayload() bool`

HasObjDebugPayload returns a boolean if a field has been set.

### GetObjDebug

`func (o *EzsignsignatureGetObjectV1Response) GetObjDebug() CommonResponseObjDebug`

GetObjDebug returns the ObjDebug field if non-nil, zero value otherwise.

### GetObjDebugOk

`func (o *EzsignsignatureGetObjectV1Response) GetObjDebugOk() (*CommonResponseObjDebug, bool)`

GetObjDebugOk returns a tuple with the ObjDebug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDebug

`func (o *EzsignsignatureGetObjectV1Response) SetObjDebug(v CommonResponseObjDebug)`

SetObjDebug sets ObjDebug field to given value.

### HasObjDebug

`func (o *EzsignsignatureGetObjectV1Response) HasObjDebug() bool`

HasObjDebug returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


