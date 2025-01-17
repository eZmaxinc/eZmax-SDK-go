# DomainGetListV1Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjDebugPayload** | [**CommonResponseObjDebugPayloadGetList**](CommonResponseObjDebugPayloadGetList.md) |  | 
**ObjDebug** | Pointer to [**CommonResponseObjDebug**](CommonResponseObjDebug.md) |  | [optional] 
**MPayload** | [**DomainGetListV1ResponseMPayload**](DomainGetListV1ResponseMPayload.md) |  | 

## Methods

### NewDomainGetListV1Response

`func NewDomainGetListV1Response(objDebugPayload CommonResponseObjDebugPayloadGetList, mPayload DomainGetListV1ResponseMPayload, ) *DomainGetListV1Response`

NewDomainGetListV1Response instantiates a new DomainGetListV1Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainGetListV1ResponseWithDefaults

`func NewDomainGetListV1ResponseWithDefaults() *DomainGetListV1Response`

NewDomainGetListV1ResponseWithDefaults instantiates a new DomainGetListV1Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjDebugPayload

`func (o *DomainGetListV1Response) GetObjDebugPayload() CommonResponseObjDebugPayloadGetList`

GetObjDebugPayload returns the ObjDebugPayload field if non-nil, zero value otherwise.

### GetObjDebugPayloadOk

`func (o *DomainGetListV1Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayloadGetList, bool)`

GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDebugPayload

`func (o *DomainGetListV1Response) SetObjDebugPayload(v CommonResponseObjDebugPayloadGetList)`

SetObjDebugPayload sets ObjDebugPayload field to given value.


### GetObjDebug

`func (o *DomainGetListV1Response) GetObjDebug() CommonResponseObjDebug`

GetObjDebug returns the ObjDebug field if non-nil, zero value otherwise.

### GetObjDebugOk

`func (o *DomainGetListV1Response) GetObjDebugOk() (*CommonResponseObjDebug, bool)`

GetObjDebugOk returns a tuple with the ObjDebug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDebug

`func (o *DomainGetListV1Response) SetObjDebug(v CommonResponseObjDebug)`

SetObjDebug sets ObjDebug field to given value.

### HasObjDebug

`func (o *DomainGetListV1Response) HasObjDebug() bool`

HasObjDebug returns a boolean if a field has been set.

### GetMPayload

`func (o *DomainGetListV1Response) GetMPayload() DomainGetListV1ResponseMPayload`

GetMPayload returns the MPayload field if non-nil, zero value otherwise.

### GetMPayloadOk

`func (o *DomainGetListV1Response) GetMPayloadOk() (*DomainGetListV1ResponseMPayload, bool)`

GetMPayloadOk returns a tuple with the MPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMPayload

`func (o *DomainGetListV1Response) SetMPayload(v DomainGetListV1ResponseMPayload)`

SetMPayload sets MPayload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


