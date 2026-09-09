# NotaryImportIntoEDMV1Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjDebugPayload** | [**CommonResponseObjDebugPayload**](CommonResponseObjDebugPayload.md) |  | 
**ObjDebug** | Pointer to [**CommonResponseObjDebug**](CommonResponseObjDebug.md) |  | [optional] 
**MPayload** | [**NotaryImportIntoEDMV1ResponseMPayload**](NotaryImportIntoEDMV1ResponseMPayload.md) |  | 

## Methods

### NewNotaryImportIntoEDMV1Response

`func NewNotaryImportIntoEDMV1Response(objDebugPayload CommonResponseObjDebugPayload, mPayload NotaryImportIntoEDMV1ResponseMPayload, ) *NotaryImportIntoEDMV1Response`

NewNotaryImportIntoEDMV1Response instantiates a new NotaryImportIntoEDMV1Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotaryImportIntoEDMV1ResponseWithDefaults

`func NewNotaryImportIntoEDMV1ResponseWithDefaults() *NotaryImportIntoEDMV1Response`

NewNotaryImportIntoEDMV1ResponseWithDefaults instantiates a new NotaryImportIntoEDMV1Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjDebugPayload

`func (o *NotaryImportIntoEDMV1Response) GetObjDebugPayload() CommonResponseObjDebugPayload`

GetObjDebugPayload returns the ObjDebugPayload field if non-nil, zero value otherwise.

### GetObjDebugPayloadOk

`func (o *NotaryImportIntoEDMV1Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool)`

GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDebugPayload

`func (o *NotaryImportIntoEDMV1Response) SetObjDebugPayload(v CommonResponseObjDebugPayload)`

SetObjDebugPayload sets ObjDebugPayload field to given value.


### GetObjDebug

`func (o *NotaryImportIntoEDMV1Response) GetObjDebug() CommonResponseObjDebug`

GetObjDebug returns the ObjDebug field if non-nil, zero value otherwise.

### GetObjDebugOk

`func (o *NotaryImportIntoEDMV1Response) GetObjDebugOk() (*CommonResponseObjDebug, bool)`

GetObjDebugOk returns a tuple with the ObjDebug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDebug

`func (o *NotaryImportIntoEDMV1Response) SetObjDebug(v CommonResponseObjDebug)`

SetObjDebug sets ObjDebug field to given value.

### HasObjDebug

`func (o *NotaryImportIntoEDMV1Response) HasObjDebug() bool`

HasObjDebug returns a boolean if a field has been set.

### GetMPayload

`func (o *NotaryImportIntoEDMV1Response) GetMPayload() NotaryImportIntoEDMV1ResponseMPayload`

GetMPayload returns the MPayload field if non-nil, zero value otherwise.

### GetMPayloadOk

`func (o *NotaryImportIntoEDMV1Response) GetMPayloadOk() (*NotaryImportIntoEDMV1ResponseMPayload, bool)`

GetMPayloadOk returns a tuple with the MPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMPayload

`func (o *NotaryImportIntoEDMV1Response) SetMPayload(v NotaryImportIntoEDMV1ResponseMPayload)`

SetMPayload sets MPayload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


