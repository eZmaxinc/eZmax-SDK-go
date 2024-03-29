# EzsigndocumentGetWordsPositionsV1Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MPayload** | [**[]CustomWordPositionWordResponse**](CustomWordPositionWordResponse.md) | Payload for the /1/object/ezsigndocument/{pkiEzsigndocumentID}/getWordsPositions API Request | 
**ObjDebugPayload** | Pointer to [**CommonResponseObjDebugPayload**](CommonResponseObjDebugPayload.md) |  | [optional] 
**ObjDebug** | Pointer to [**CommonResponseObjDebug**](CommonResponseObjDebug.md) |  | [optional] 

## Methods

### NewEzsigndocumentGetWordsPositionsV1Response

`func NewEzsigndocumentGetWordsPositionsV1Response(mPayload []CustomWordPositionWordResponse, ) *EzsigndocumentGetWordsPositionsV1Response`

NewEzsigndocumentGetWordsPositionsV1Response instantiates a new EzsigndocumentGetWordsPositionsV1Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigndocumentGetWordsPositionsV1ResponseWithDefaults

`func NewEzsigndocumentGetWordsPositionsV1ResponseWithDefaults() *EzsigndocumentGetWordsPositionsV1Response`

NewEzsigndocumentGetWordsPositionsV1ResponseWithDefaults instantiates a new EzsigndocumentGetWordsPositionsV1Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMPayload

`func (o *EzsigndocumentGetWordsPositionsV1Response) GetMPayload() []CustomWordPositionWordResponse`

GetMPayload returns the MPayload field if non-nil, zero value otherwise.

### GetMPayloadOk

`func (o *EzsigndocumentGetWordsPositionsV1Response) GetMPayloadOk() (*[]CustomWordPositionWordResponse, bool)`

GetMPayloadOk returns a tuple with the MPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMPayload

`func (o *EzsigndocumentGetWordsPositionsV1Response) SetMPayload(v []CustomWordPositionWordResponse)`

SetMPayload sets MPayload field to given value.


### GetObjDebugPayload

`func (o *EzsigndocumentGetWordsPositionsV1Response) GetObjDebugPayload() CommonResponseObjDebugPayload`

GetObjDebugPayload returns the ObjDebugPayload field if non-nil, zero value otherwise.

### GetObjDebugPayloadOk

`func (o *EzsigndocumentGetWordsPositionsV1Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool)`

GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDebugPayload

`func (o *EzsigndocumentGetWordsPositionsV1Response) SetObjDebugPayload(v CommonResponseObjDebugPayload)`

SetObjDebugPayload sets ObjDebugPayload field to given value.

### HasObjDebugPayload

`func (o *EzsigndocumentGetWordsPositionsV1Response) HasObjDebugPayload() bool`

HasObjDebugPayload returns a boolean if a field has been set.

### GetObjDebug

`func (o *EzsigndocumentGetWordsPositionsV1Response) GetObjDebug() CommonResponseObjDebug`

GetObjDebug returns the ObjDebug field if non-nil, zero value otherwise.

### GetObjDebugOk

`func (o *EzsigndocumentGetWordsPositionsV1Response) GetObjDebugOk() (*CommonResponseObjDebug, bool)`

GetObjDebugOk returns a tuple with the ObjDebug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDebug

`func (o *EzsigndocumentGetWordsPositionsV1Response) SetObjDebug(v CommonResponseObjDebug)`

SetObjDebug sets ObjDebug field to given value.

### HasObjDebug

`func (o *EzsigndocumentGetWordsPositionsV1Response) HasObjDebug() bool`

HasObjDebug returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


