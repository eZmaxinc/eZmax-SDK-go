# EzmaxinvoicingGetObjectV2Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjDebugPayload** | [**CommonResponseObjDebugPayload**](CommonResponseObjDebugPayload.md) |  | 
**ObjDebug** | Pointer to [**CommonResponseObjDebug**](CommonResponseObjDebug.md) |  | [optional] 
**MPayload** | [**EzmaxinvoicingGetObjectV2ResponseMPayload**](EzmaxinvoicingGetObjectV2ResponseMPayload.md) |  | 

## Methods

### NewEzmaxinvoicingGetObjectV2Response

`func NewEzmaxinvoicingGetObjectV2Response(objDebugPayload CommonResponseObjDebugPayload, mPayload EzmaxinvoicingGetObjectV2ResponseMPayload, ) *EzmaxinvoicingGetObjectV2Response`

NewEzmaxinvoicingGetObjectV2Response instantiates a new EzmaxinvoicingGetObjectV2Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzmaxinvoicingGetObjectV2ResponseWithDefaults

`func NewEzmaxinvoicingGetObjectV2ResponseWithDefaults() *EzmaxinvoicingGetObjectV2Response`

NewEzmaxinvoicingGetObjectV2ResponseWithDefaults instantiates a new EzmaxinvoicingGetObjectV2Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjDebugPayload

`func (o *EzmaxinvoicingGetObjectV2Response) GetObjDebugPayload() CommonResponseObjDebugPayload`

GetObjDebugPayload returns the ObjDebugPayload field if non-nil, zero value otherwise.

### GetObjDebugPayloadOk

`func (o *EzmaxinvoicingGetObjectV2Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool)`

GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDebugPayload

`func (o *EzmaxinvoicingGetObjectV2Response) SetObjDebugPayload(v CommonResponseObjDebugPayload)`

SetObjDebugPayload sets ObjDebugPayload field to given value.


### GetObjDebug

`func (o *EzmaxinvoicingGetObjectV2Response) GetObjDebug() CommonResponseObjDebug`

GetObjDebug returns the ObjDebug field if non-nil, zero value otherwise.

### GetObjDebugOk

`func (o *EzmaxinvoicingGetObjectV2Response) GetObjDebugOk() (*CommonResponseObjDebug, bool)`

GetObjDebugOk returns a tuple with the ObjDebug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDebug

`func (o *EzmaxinvoicingGetObjectV2Response) SetObjDebug(v CommonResponseObjDebug)`

SetObjDebug sets ObjDebug field to given value.

### HasObjDebug

`func (o *EzmaxinvoicingGetObjectV2Response) HasObjDebug() bool`

HasObjDebug returns a boolean if a field has been set.

### GetMPayload

`func (o *EzmaxinvoicingGetObjectV2Response) GetMPayload() EzmaxinvoicingGetObjectV2ResponseMPayload`

GetMPayload returns the MPayload field if non-nil, zero value otherwise.

### GetMPayloadOk

`func (o *EzmaxinvoicingGetObjectV2Response) GetMPayloadOk() (*EzmaxinvoicingGetObjectV2ResponseMPayload, bool)`

GetMPayloadOk returns a tuple with the MPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMPayload

`func (o *EzmaxinvoicingGetObjectV2Response) SetMPayload(v EzmaxinvoicingGetObjectV2ResponseMPayload)`

SetMPayload sets MPayload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


