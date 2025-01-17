# SignatureGetObjectV3Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjDebugPayload** | [**CommonResponseObjDebugPayload**](CommonResponseObjDebugPayload.md) |  | 
**ObjDebug** | Pointer to [**CommonResponseObjDebug**](CommonResponseObjDebug.md) |  | [optional] 
**MPayload** | [**SignatureGetObjectV3ResponseMPayload**](SignatureGetObjectV3ResponseMPayload.md) |  | 

## Methods

### NewSignatureGetObjectV3Response

`func NewSignatureGetObjectV3Response(objDebugPayload CommonResponseObjDebugPayload, mPayload SignatureGetObjectV3ResponseMPayload, ) *SignatureGetObjectV3Response`

NewSignatureGetObjectV3Response instantiates a new SignatureGetObjectV3Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignatureGetObjectV3ResponseWithDefaults

`func NewSignatureGetObjectV3ResponseWithDefaults() *SignatureGetObjectV3Response`

NewSignatureGetObjectV3ResponseWithDefaults instantiates a new SignatureGetObjectV3Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjDebugPayload

`func (o *SignatureGetObjectV3Response) GetObjDebugPayload() CommonResponseObjDebugPayload`

GetObjDebugPayload returns the ObjDebugPayload field if non-nil, zero value otherwise.

### GetObjDebugPayloadOk

`func (o *SignatureGetObjectV3Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool)`

GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDebugPayload

`func (o *SignatureGetObjectV3Response) SetObjDebugPayload(v CommonResponseObjDebugPayload)`

SetObjDebugPayload sets ObjDebugPayload field to given value.


### GetObjDebug

`func (o *SignatureGetObjectV3Response) GetObjDebug() CommonResponseObjDebug`

GetObjDebug returns the ObjDebug field if non-nil, zero value otherwise.

### GetObjDebugOk

`func (o *SignatureGetObjectV3Response) GetObjDebugOk() (*CommonResponseObjDebug, bool)`

GetObjDebugOk returns a tuple with the ObjDebug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDebug

`func (o *SignatureGetObjectV3Response) SetObjDebug(v CommonResponseObjDebug)`

SetObjDebug sets ObjDebug field to given value.

### HasObjDebug

`func (o *SignatureGetObjectV3Response) HasObjDebug() bool`

HasObjDebug returns a boolean if a field has been set.

### GetMPayload

`func (o *SignatureGetObjectV3Response) GetMPayload() SignatureGetObjectV3ResponseMPayload`

GetMPayload returns the MPayload field if non-nil, zero value otherwise.

### GetMPayloadOk

`func (o *SignatureGetObjectV3Response) GetMPayloadOk() (*SignatureGetObjectV3ResponseMPayload, bool)`

GetMPayloadOk returns a tuple with the MPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMPayload

`func (o *SignatureGetObjectV3Response) SetMPayload(v SignatureGetObjectV3ResponseMPayload)`

SetMPayload sets MPayload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


