# SignatureCreateObjectV1Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjDebugPayload** | [**CommonResponseObjDebugPayload**](CommonResponseObjDebugPayload.md) |  | 
**ObjDebug** | Pointer to [**CommonResponseObjDebug**](CommonResponseObjDebug.md) |  | [optional] 
**MPayload** | [**SignatureCreateObjectV1ResponseMPayload**](SignatureCreateObjectV1ResponseMPayload.md) |  | 

## Methods

### NewSignatureCreateObjectV1Response

`func NewSignatureCreateObjectV1Response(objDebugPayload CommonResponseObjDebugPayload, mPayload SignatureCreateObjectV1ResponseMPayload, ) *SignatureCreateObjectV1Response`

NewSignatureCreateObjectV1Response instantiates a new SignatureCreateObjectV1Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignatureCreateObjectV1ResponseWithDefaults

`func NewSignatureCreateObjectV1ResponseWithDefaults() *SignatureCreateObjectV1Response`

NewSignatureCreateObjectV1ResponseWithDefaults instantiates a new SignatureCreateObjectV1Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjDebugPayload

`func (o *SignatureCreateObjectV1Response) GetObjDebugPayload() CommonResponseObjDebugPayload`

GetObjDebugPayload returns the ObjDebugPayload field if non-nil, zero value otherwise.

### GetObjDebugPayloadOk

`func (o *SignatureCreateObjectV1Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool)`

GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDebugPayload

`func (o *SignatureCreateObjectV1Response) SetObjDebugPayload(v CommonResponseObjDebugPayload)`

SetObjDebugPayload sets ObjDebugPayload field to given value.


### GetObjDebug

`func (o *SignatureCreateObjectV1Response) GetObjDebug() CommonResponseObjDebug`

GetObjDebug returns the ObjDebug field if non-nil, zero value otherwise.

### GetObjDebugOk

`func (o *SignatureCreateObjectV1Response) GetObjDebugOk() (*CommonResponseObjDebug, bool)`

GetObjDebugOk returns a tuple with the ObjDebug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDebug

`func (o *SignatureCreateObjectV1Response) SetObjDebug(v CommonResponseObjDebug)`

SetObjDebug sets ObjDebug field to given value.

### HasObjDebug

`func (o *SignatureCreateObjectV1Response) HasObjDebug() bool`

HasObjDebug returns a boolean if a field has been set.

### GetMPayload

`func (o *SignatureCreateObjectV1Response) GetMPayload() SignatureCreateObjectV1ResponseMPayload`

GetMPayload returns the MPayload field if non-nil, zero value otherwise.

### GetMPayloadOk

`func (o *SignatureCreateObjectV1Response) GetMPayloadOk() (*SignatureCreateObjectV1ResponseMPayload, bool)`

GetMPayloadOk returns a tuple with the MPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMPayload

`func (o *SignatureCreateObjectV1Response) SetMPayload(v SignatureCreateObjectV1ResponseMPayload)`

SetMPayload sets MPayload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


