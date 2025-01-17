# AuthenticationexternalGetObjectV2Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjDebugPayload** | [**CommonResponseObjDebugPayload**](CommonResponseObjDebugPayload.md) |  | 
**ObjDebug** | Pointer to [**CommonResponseObjDebug**](CommonResponseObjDebug.md) |  | [optional] 
**MPayload** | [**AuthenticationexternalGetObjectV2ResponseMPayload**](AuthenticationexternalGetObjectV2ResponseMPayload.md) |  | 

## Methods

### NewAuthenticationexternalGetObjectV2Response

`func NewAuthenticationexternalGetObjectV2Response(objDebugPayload CommonResponseObjDebugPayload, mPayload AuthenticationexternalGetObjectV2ResponseMPayload, ) *AuthenticationexternalGetObjectV2Response`

NewAuthenticationexternalGetObjectV2Response instantiates a new AuthenticationexternalGetObjectV2Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationexternalGetObjectV2ResponseWithDefaults

`func NewAuthenticationexternalGetObjectV2ResponseWithDefaults() *AuthenticationexternalGetObjectV2Response`

NewAuthenticationexternalGetObjectV2ResponseWithDefaults instantiates a new AuthenticationexternalGetObjectV2Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjDebugPayload

`func (o *AuthenticationexternalGetObjectV2Response) GetObjDebugPayload() CommonResponseObjDebugPayload`

GetObjDebugPayload returns the ObjDebugPayload field if non-nil, zero value otherwise.

### GetObjDebugPayloadOk

`func (o *AuthenticationexternalGetObjectV2Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool)`

GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDebugPayload

`func (o *AuthenticationexternalGetObjectV2Response) SetObjDebugPayload(v CommonResponseObjDebugPayload)`

SetObjDebugPayload sets ObjDebugPayload field to given value.


### GetObjDebug

`func (o *AuthenticationexternalGetObjectV2Response) GetObjDebug() CommonResponseObjDebug`

GetObjDebug returns the ObjDebug field if non-nil, zero value otherwise.

### GetObjDebugOk

`func (o *AuthenticationexternalGetObjectV2Response) GetObjDebugOk() (*CommonResponseObjDebug, bool)`

GetObjDebugOk returns a tuple with the ObjDebug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDebug

`func (o *AuthenticationexternalGetObjectV2Response) SetObjDebug(v CommonResponseObjDebug)`

SetObjDebug sets ObjDebug field to given value.

### HasObjDebug

`func (o *AuthenticationexternalGetObjectV2Response) HasObjDebug() bool`

HasObjDebug returns a boolean if a field has been set.

### GetMPayload

`func (o *AuthenticationexternalGetObjectV2Response) GetMPayload() AuthenticationexternalGetObjectV2ResponseMPayload`

GetMPayload returns the MPayload field if non-nil, zero value otherwise.

### GetMPayloadOk

`func (o *AuthenticationexternalGetObjectV2Response) GetMPayloadOk() (*AuthenticationexternalGetObjectV2ResponseMPayload, bool)`

GetMPayloadOk returns a tuple with the MPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMPayload

`func (o *AuthenticationexternalGetObjectV2Response) SetMPayload(v AuthenticationexternalGetObjectV2ResponseMPayload)`

SetMPayload sets MPayload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


