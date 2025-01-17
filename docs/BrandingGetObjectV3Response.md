# BrandingGetObjectV3Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjDebugPayload** | [**CommonResponseObjDebugPayload**](CommonResponseObjDebugPayload.md) |  | 
**ObjDebug** | Pointer to [**CommonResponseObjDebug**](CommonResponseObjDebug.md) |  | [optional] 
**MPayload** | [**BrandingGetObjectV3ResponseMPayload**](BrandingGetObjectV3ResponseMPayload.md) |  | 

## Methods

### NewBrandingGetObjectV3Response

`func NewBrandingGetObjectV3Response(objDebugPayload CommonResponseObjDebugPayload, mPayload BrandingGetObjectV3ResponseMPayload, ) *BrandingGetObjectV3Response`

NewBrandingGetObjectV3Response instantiates a new BrandingGetObjectV3Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrandingGetObjectV3ResponseWithDefaults

`func NewBrandingGetObjectV3ResponseWithDefaults() *BrandingGetObjectV3Response`

NewBrandingGetObjectV3ResponseWithDefaults instantiates a new BrandingGetObjectV3Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjDebugPayload

`func (o *BrandingGetObjectV3Response) GetObjDebugPayload() CommonResponseObjDebugPayload`

GetObjDebugPayload returns the ObjDebugPayload field if non-nil, zero value otherwise.

### GetObjDebugPayloadOk

`func (o *BrandingGetObjectV3Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool)`

GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDebugPayload

`func (o *BrandingGetObjectV3Response) SetObjDebugPayload(v CommonResponseObjDebugPayload)`

SetObjDebugPayload sets ObjDebugPayload field to given value.


### GetObjDebug

`func (o *BrandingGetObjectV3Response) GetObjDebug() CommonResponseObjDebug`

GetObjDebug returns the ObjDebug field if non-nil, zero value otherwise.

### GetObjDebugOk

`func (o *BrandingGetObjectV3Response) GetObjDebugOk() (*CommonResponseObjDebug, bool)`

GetObjDebugOk returns a tuple with the ObjDebug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDebug

`func (o *BrandingGetObjectV3Response) SetObjDebug(v CommonResponseObjDebug)`

SetObjDebug sets ObjDebug field to given value.

### HasObjDebug

`func (o *BrandingGetObjectV3Response) HasObjDebug() bool`

HasObjDebug returns a boolean if a field has been set.

### GetMPayload

`func (o *BrandingGetObjectV3Response) GetMPayload() BrandingGetObjectV3ResponseMPayload`

GetMPayload returns the MPayload field if non-nil, zero value otherwise.

### GetMPayloadOk

`func (o *BrandingGetObjectV3Response) GetMPayloadOk() (*BrandingGetObjectV3ResponseMPayload, bool)`

GetMPayloadOk returns a tuple with the MPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMPayload

`func (o *BrandingGetObjectV3Response) SetMPayload(v BrandingGetObjectV3ResponseMPayload)`

SetMPayload sets MPayload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


