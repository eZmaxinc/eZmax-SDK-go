# EzsignsignatureGetObjectV4Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjDebugPayload** | [**CommonResponseObjDebugPayload**](CommonResponseObjDebugPayload.md) |  | 
**ObjDebug** | Pointer to [**CommonResponseObjDebug**](CommonResponseObjDebug.md) |  | [optional] 
**MPayload** | [**EzsignsignatureGetObjectV4ResponseMPayload**](EzsignsignatureGetObjectV4ResponseMPayload.md) |  | 

## Methods

### NewEzsignsignatureGetObjectV4Response

`func NewEzsignsignatureGetObjectV4Response(objDebugPayload CommonResponseObjDebugPayload, mPayload EzsignsignatureGetObjectV4ResponseMPayload, ) *EzsignsignatureGetObjectV4Response`

NewEzsignsignatureGetObjectV4Response instantiates a new EzsignsignatureGetObjectV4Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignsignatureGetObjectV4ResponseWithDefaults

`func NewEzsignsignatureGetObjectV4ResponseWithDefaults() *EzsignsignatureGetObjectV4Response`

NewEzsignsignatureGetObjectV4ResponseWithDefaults instantiates a new EzsignsignatureGetObjectV4Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjDebugPayload

`func (o *EzsignsignatureGetObjectV4Response) GetObjDebugPayload() CommonResponseObjDebugPayload`

GetObjDebugPayload returns the ObjDebugPayload field if non-nil, zero value otherwise.

### GetObjDebugPayloadOk

`func (o *EzsignsignatureGetObjectV4Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool)`

GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDebugPayload

`func (o *EzsignsignatureGetObjectV4Response) SetObjDebugPayload(v CommonResponseObjDebugPayload)`

SetObjDebugPayload sets ObjDebugPayload field to given value.


### GetObjDebug

`func (o *EzsignsignatureGetObjectV4Response) GetObjDebug() CommonResponseObjDebug`

GetObjDebug returns the ObjDebug field if non-nil, zero value otherwise.

### GetObjDebugOk

`func (o *EzsignsignatureGetObjectV4Response) GetObjDebugOk() (*CommonResponseObjDebug, bool)`

GetObjDebugOk returns a tuple with the ObjDebug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDebug

`func (o *EzsignsignatureGetObjectV4Response) SetObjDebug(v CommonResponseObjDebug)`

SetObjDebug sets ObjDebug field to given value.

### HasObjDebug

`func (o *EzsignsignatureGetObjectV4Response) HasObjDebug() bool`

HasObjDebug returns a boolean if a field has been set.

### GetMPayload

`func (o *EzsignsignatureGetObjectV4Response) GetMPayload() EzsignsignatureGetObjectV4ResponseMPayload`

GetMPayload returns the MPayload field if non-nil, zero value otherwise.

### GetMPayloadOk

`func (o *EzsignsignatureGetObjectV4Response) GetMPayloadOk() (*EzsignsignatureGetObjectV4ResponseMPayload, bool)`

GetMPayloadOk returns a tuple with the MPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMPayload

`func (o *EzsignsignatureGetObjectV4Response) SetMPayload(v EzsignsignatureGetObjectV4ResponseMPayload)`

SetMPayload sets MPayload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


