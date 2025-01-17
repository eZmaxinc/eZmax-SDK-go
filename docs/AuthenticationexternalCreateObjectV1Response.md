# AuthenticationexternalCreateObjectV1Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjDebugPayload** | [**CommonResponseObjDebugPayload**](CommonResponseObjDebugPayload.md) |  | 
**ObjDebug** | Pointer to [**CommonResponseObjDebug**](CommonResponseObjDebug.md) |  | [optional] 
**MPayload** | [**AuthenticationexternalCreateObjectV1ResponseMPayload**](AuthenticationexternalCreateObjectV1ResponseMPayload.md) |  | 

## Methods

### NewAuthenticationexternalCreateObjectV1Response

`func NewAuthenticationexternalCreateObjectV1Response(objDebugPayload CommonResponseObjDebugPayload, mPayload AuthenticationexternalCreateObjectV1ResponseMPayload, ) *AuthenticationexternalCreateObjectV1Response`

NewAuthenticationexternalCreateObjectV1Response instantiates a new AuthenticationexternalCreateObjectV1Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationexternalCreateObjectV1ResponseWithDefaults

`func NewAuthenticationexternalCreateObjectV1ResponseWithDefaults() *AuthenticationexternalCreateObjectV1Response`

NewAuthenticationexternalCreateObjectV1ResponseWithDefaults instantiates a new AuthenticationexternalCreateObjectV1Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjDebugPayload

`func (o *AuthenticationexternalCreateObjectV1Response) GetObjDebugPayload() CommonResponseObjDebugPayload`

GetObjDebugPayload returns the ObjDebugPayload field if non-nil, zero value otherwise.

### GetObjDebugPayloadOk

`func (o *AuthenticationexternalCreateObjectV1Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool)`

GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDebugPayload

`func (o *AuthenticationexternalCreateObjectV1Response) SetObjDebugPayload(v CommonResponseObjDebugPayload)`

SetObjDebugPayload sets ObjDebugPayload field to given value.


### GetObjDebug

`func (o *AuthenticationexternalCreateObjectV1Response) GetObjDebug() CommonResponseObjDebug`

GetObjDebug returns the ObjDebug field if non-nil, zero value otherwise.

### GetObjDebugOk

`func (o *AuthenticationexternalCreateObjectV1Response) GetObjDebugOk() (*CommonResponseObjDebug, bool)`

GetObjDebugOk returns a tuple with the ObjDebug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDebug

`func (o *AuthenticationexternalCreateObjectV1Response) SetObjDebug(v CommonResponseObjDebug)`

SetObjDebug sets ObjDebug field to given value.

### HasObjDebug

`func (o *AuthenticationexternalCreateObjectV1Response) HasObjDebug() bool`

HasObjDebug returns a boolean if a field has been set.

### GetMPayload

`func (o *AuthenticationexternalCreateObjectV1Response) GetMPayload() AuthenticationexternalCreateObjectV1ResponseMPayload`

GetMPayload returns the MPayload field if non-nil, zero value otherwise.

### GetMPayloadOk

`func (o *AuthenticationexternalCreateObjectV1Response) GetMPayloadOk() (*AuthenticationexternalCreateObjectV1ResponseMPayload, bool)`

GetMPayloadOk returns a tuple with the MPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMPayload

`func (o *AuthenticationexternalCreateObjectV1Response) SetMPayload(v AuthenticationexternalCreateObjectV1ResponseMPayload)`

SetMPayload sets MPayload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


