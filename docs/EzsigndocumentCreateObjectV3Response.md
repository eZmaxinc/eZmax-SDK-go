# EzsigndocumentCreateObjectV3Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjDebugPayload** | [**CommonResponseObjDebugPayload**](CommonResponseObjDebugPayload.md) |  | 
**ObjDebug** | Pointer to [**CommonResponseObjDebug**](CommonResponseObjDebug.md) |  | [optional] 
**MPayload** | [**EzsigndocumentCreateObjectV3ResponseMPayload**](EzsigndocumentCreateObjectV3ResponseMPayload.md) |  | 

## Methods

### NewEzsigndocumentCreateObjectV3Response

`func NewEzsigndocumentCreateObjectV3Response(objDebugPayload CommonResponseObjDebugPayload, mPayload EzsigndocumentCreateObjectV3ResponseMPayload, ) *EzsigndocumentCreateObjectV3Response`

NewEzsigndocumentCreateObjectV3Response instantiates a new EzsigndocumentCreateObjectV3Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigndocumentCreateObjectV3ResponseWithDefaults

`func NewEzsigndocumentCreateObjectV3ResponseWithDefaults() *EzsigndocumentCreateObjectV3Response`

NewEzsigndocumentCreateObjectV3ResponseWithDefaults instantiates a new EzsigndocumentCreateObjectV3Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjDebugPayload

`func (o *EzsigndocumentCreateObjectV3Response) GetObjDebugPayload() CommonResponseObjDebugPayload`

GetObjDebugPayload returns the ObjDebugPayload field if non-nil, zero value otherwise.

### GetObjDebugPayloadOk

`func (o *EzsigndocumentCreateObjectV3Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool)`

GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDebugPayload

`func (o *EzsigndocumentCreateObjectV3Response) SetObjDebugPayload(v CommonResponseObjDebugPayload)`

SetObjDebugPayload sets ObjDebugPayload field to given value.


### GetObjDebug

`func (o *EzsigndocumentCreateObjectV3Response) GetObjDebug() CommonResponseObjDebug`

GetObjDebug returns the ObjDebug field if non-nil, zero value otherwise.

### GetObjDebugOk

`func (o *EzsigndocumentCreateObjectV3Response) GetObjDebugOk() (*CommonResponseObjDebug, bool)`

GetObjDebugOk returns a tuple with the ObjDebug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDebug

`func (o *EzsigndocumentCreateObjectV3Response) SetObjDebug(v CommonResponseObjDebug)`

SetObjDebug sets ObjDebug field to given value.

### HasObjDebug

`func (o *EzsigndocumentCreateObjectV3Response) HasObjDebug() bool`

HasObjDebug returns a boolean if a field has been set.

### GetMPayload

`func (o *EzsigndocumentCreateObjectV3Response) GetMPayload() EzsigndocumentCreateObjectV3ResponseMPayload`

GetMPayload returns the MPayload field if non-nil, zero value otherwise.

### GetMPayloadOk

`func (o *EzsigndocumentCreateObjectV3Response) GetMPayloadOk() (*EzsigndocumentCreateObjectV3ResponseMPayload, bool)`

GetMPayloadOk returns a tuple with the MPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMPayload

`func (o *EzsigndocumentCreateObjectV3Response) SetMPayload(v EzsigndocumentCreateObjectV3ResponseMPayload)`

SetMPayload sets MPayload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


