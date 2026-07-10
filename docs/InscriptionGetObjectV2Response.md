# InscriptionGetObjectV2Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjDebugPayload** | [**CommonResponseObjDebugPayload**](CommonResponseObjDebugPayload.md) |  | 
**ObjDebug** | Pointer to [**CommonResponseObjDebug**](CommonResponseObjDebug.md) |  | [optional] 
**MPayload** | [**InscriptionGetObjectV2ResponseMPayload**](InscriptionGetObjectV2ResponseMPayload.md) |  | 

## Methods

### NewInscriptionGetObjectV2Response

`func NewInscriptionGetObjectV2Response(objDebugPayload CommonResponseObjDebugPayload, mPayload InscriptionGetObjectV2ResponseMPayload, ) *InscriptionGetObjectV2Response`

NewInscriptionGetObjectV2Response instantiates a new InscriptionGetObjectV2Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInscriptionGetObjectV2ResponseWithDefaults

`func NewInscriptionGetObjectV2ResponseWithDefaults() *InscriptionGetObjectV2Response`

NewInscriptionGetObjectV2ResponseWithDefaults instantiates a new InscriptionGetObjectV2Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjDebugPayload

`func (o *InscriptionGetObjectV2Response) GetObjDebugPayload() CommonResponseObjDebugPayload`

GetObjDebugPayload returns the ObjDebugPayload field if non-nil, zero value otherwise.

### GetObjDebugPayloadOk

`func (o *InscriptionGetObjectV2Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool)`

GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDebugPayload

`func (o *InscriptionGetObjectV2Response) SetObjDebugPayload(v CommonResponseObjDebugPayload)`

SetObjDebugPayload sets ObjDebugPayload field to given value.


### GetObjDebug

`func (o *InscriptionGetObjectV2Response) GetObjDebug() CommonResponseObjDebug`

GetObjDebug returns the ObjDebug field if non-nil, zero value otherwise.

### GetObjDebugOk

`func (o *InscriptionGetObjectV2Response) GetObjDebugOk() (*CommonResponseObjDebug, bool)`

GetObjDebugOk returns a tuple with the ObjDebug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDebug

`func (o *InscriptionGetObjectV2Response) SetObjDebug(v CommonResponseObjDebug)`

SetObjDebug sets ObjDebug field to given value.

### HasObjDebug

`func (o *InscriptionGetObjectV2Response) HasObjDebug() bool`

HasObjDebug returns a boolean if a field has been set.

### GetMPayload

`func (o *InscriptionGetObjectV2Response) GetMPayload() InscriptionGetObjectV2ResponseMPayload`

GetMPayload returns the MPayload field if non-nil, zero value otherwise.

### GetMPayloadOk

`func (o *InscriptionGetObjectV2Response) GetMPayloadOk() (*InscriptionGetObjectV2ResponseMPayload, bool)`

GetMPayloadOk returns a tuple with the MPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMPayload

`func (o *InscriptionGetObjectV2Response) SetMPayload(v InscriptionGetObjectV2ResponseMPayload)`

SetMPayload sets MPayload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


