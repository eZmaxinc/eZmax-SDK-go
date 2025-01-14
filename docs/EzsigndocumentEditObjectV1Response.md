# EzsigndocumentEditObjectV1Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjDebugPayload** | [**CommonResponseObjDebugPayload**](CommonResponseObjDebugPayload.md) |  | 
**ObjDebug** | Pointer to [**CommonResponseObjDebug**](CommonResponseObjDebug.md) |  | [optional] 
**AObjWarning** | Pointer to [**[]CommonResponseWarning**](CommonResponseWarning.md) |  | [optional] 

## Methods

### NewEzsigndocumentEditObjectV1Response

`func NewEzsigndocumentEditObjectV1Response(objDebugPayload CommonResponseObjDebugPayload, ) *EzsigndocumentEditObjectV1Response`

NewEzsigndocumentEditObjectV1Response instantiates a new EzsigndocumentEditObjectV1Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigndocumentEditObjectV1ResponseWithDefaults

`func NewEzsigndocumentEditObjectV1ResponseWithDefaults() *EzsigndocumentEditObjectV1Response`

NewEzsigndocumentEditObjectV1ResponseWithDefaults instantiates a new EzsigndocumentEditObjectV1Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjDebugPayload

`func (o *EzsigndocumentEditObjectV1Response) GetObjDebugPayload() CommonResponseObjDebugPayload`

GetObjDebugPayload returns the ObjDebugPayload field if non-nil, zero value otherwise.

### GetObjDebugPayloadOk

`func (o *EzsigndocumentEditObjectV1Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool)`

GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDebugPayload

`func (o *EzsigndocumentEditObjectV1Response) SetObjDebugPayload(v CommonResponseObjDebugPayload)`

SetObjDebugPayload sets ObjDebugPayload field to given value.


### GetObjDebug

`func (o *EzsigndocumentEditObjectV1Response) GetObjDebug() CommonResponseObjDebug`

GetObjDebug returns the ObjDebug field if non-nil, zero value otherwise.

### GetObjDebugOk

`func (o *EzsigndocumentEditObjectV1Response) GetObjDebugOk() (*CommonResponseObjDebug, bool)`

GetObjDebugOk returns a tuple with the ObjDebug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDebug

`func (o *EzsigndocumentEditObjectV1Response) SetObjDebug(v CommonResponseObjDebug)`

SetObjDebug sets ObjDebug field to given value.

### HasObjDebug

`func (o *EzsigndocumentEditObjectV1Response) HasObjDebug() bool`

HasObjDebug returns a boolean if a field has been set.

### GetAObjWarning

`func (o *EzsigndocumentEditObjectV1Response) GetAObjWarning() []CommonResponseWarning`

GetAObjWarning returns the AObjWarning field if non-nil, zero value otherwise.

### GetAObjWarningOk

`func (o *EzsigndocumentEditObjectV1Response) GetAObjWarningOk() (*[]CommonResponseWarning, bool)`

GetAObjWarningOk returns a tuple with the AObjWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjWarning

`func (o *EzsigndocumentEditObjectV1Response) SetAObjWarning(v []CommonResponseWarning)`

SetAObjWarning sets AObjWarning field to given value.

### HasAObjWarning

`func (o *EzsigndocumentEditObjectV1Response) HasAObjWarning() bool`

HasAObjWarning returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


