# VariableexpenseGetObjectV2Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjDebugPayload** | [**CommonResponseObjDebugPayload**](CommonResponseObjDebugPayload.md) |  | 
**ObjDebug** | Pointer to [**CommonResponseObjDebug**](CommonResponseObjDebug.md) |  | [optional] 
**MPayload** | [**VariableexpenseGetObjectV2ResponseMPayload**](VariableexpenseGetObjectV2ResponseMPayload.md) |  | 

## Methods

### NewVariableexpenseGetObjectV2Response

`func NewVariableexpenseGetObjectV2Response(objDebugPayload CommonResponseObjDebugPayload, mPayload VariableexpenseGetObjectV2ResponseMPayload, ) *VariableexpenseGetObjectV2Response`

NewVariableexpenseGetObjectV2Response instantiates a new VariableexpenseGetObjectV2Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariableexpenseGetObjectV2ResponseWithDefaults

`func NewVariableexpenseGetObjectV2ResponseWithDefaults() *VariableexpenseGetObjectV2Response`

NewVariableexpenseGetObjectV2ResponseWithDefaults instantiates a new VariableexpenseGetObjectV2Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjDebugPayload

`func (o *VariableexpenseGetObjectV2Response) GetObjDebugPayload() CommonResponseObjDebugPayload`

GetObjDebugPayload returns the ObjDebugPayload field if non-nil, zero value otherwise.

### GetObjDebugPayloadOk

`func (o *VariableexpenseGetObjectV2Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool)`

GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDebugPayload

`func (o *VariableexpenseGetObjectV2Response) SetObjDebugPayload(v CommonResponseObjDebugPayload)`

SetObjDebugPayload sets ObjDebugPayload field to given value.


### GetObjDebug

`func (o *VariableexpenseGetObjectV2Response) GetObjDebug() CommonResponseObjDebug`

GetObjDebug returns the ObjDebug field if non-nil, zero value otherwise.

### GetObjDebugOk

`func (o *VariableexpenseGetObjectV2Response) GetObjDebugOk() (*CommonResponseObjDebug, bool)`

GetObjDebugOk returns a tuple with the ObjDebug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDebug

`func (o *VariableexpenseGetObjectV2Response) SetObjDebug(v CommonResponseObjDebug)`

SetObjDebug sets ObjDebug field to given value.

### HasObjDebug

`func (o *VariableexpenseGetObjectV2Response) HasObjDebug() bool`

HasObjDebug returns a boolean if a field has been set.

### GetMPayload

`func (o *VariableexpenseGetObjectV2Response) GetMPayload() VariableexpenseGetObjectV2ResponseMPayload`

GetMPayload returns the MPayload field if non-nil, zero value otherwise.

### GetMPayloadOk

`func (o *VariableexpenseGetObjectV2Response) GetMPayloadOk() (*VariableexpenseGetObjectV2ResponseMPayload, bool)`

GetMPayloadOk returns a tuple with the MPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMPayload

`func (o *VariableexpenseGetObjectV2Response) SetMPayload(v VariableexpenseGetObjectV2ResponseMPayload)`

SetMPayload sets MPayload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


