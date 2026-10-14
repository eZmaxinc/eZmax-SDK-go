# SalaryGetCommunicationListV1Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjDebugPayload** | [**CommonResponseObjDebugPayloadGetList**](CommonResponseObjDebugPayloadGetList.md) |  | 
**ObjDebug** | Pointer to [**CommonResponseObjDebug**](CommonResponseObjDebug.md) |  | [optional] 
**MPayload** | [**SalaryGetCommunicationListV1ResponseMPayload**](SalaryGetCommunicationListV1ResponseMPayload.md) |  | 

## Methods

### NewSalaryGetCommunicationListV1Response

`func NewSalaryGetCommunicationListV1Response(objDebugPayload CommonResponseObjDebugPayloadGetList, mPayload SalaryGetCommunicationListV1ResponseMPayload, ) *SalaryGetCommunicationListV1Response`

NewSalaryGetCommunicationListV1Response instantiates a new SalaryGetCommunicationListV1Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSalaryGetCommunicationListV1ResponseWithDefaults

`func NewSalaryGetCommunicationListV1ResponseWithDefaults() *SalaryGetCommunicationListV1Response`

NewSalaryGetCommunicationListV1ResponseWithDefaults instantiates a new SalaryGetCommunicationListV1Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjDebugPayload

`func (o *SalaryGetCommunicationListV1Response) GetObjDebugPayload() CommonResponseObjDebugPayloadGetList`

GetObjDebugPayload returns the ObjDebugPayload field if non-nil, zero value otherwise.

### GetObjDebugPayloadOk

`func (o *SalaryGetCommunicationListV1Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayloadGetList, bool)`

GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDebugPayload

`func (o *SalaryGetCommunicationListV1Response) SetObjDebugPayload(v CommonResponseObjDebugPayloadGetList)`

SetObjDebugPayload sets ObjDebugPayload field to given value.


### GetObjDebug

`func (o *SalaryGetCommunicationListV1Response) GetObjDebug() CommonResponseObjDebug`

GetObjDebug returns the ObjDebug field if non-nil, zero value otherwise.

### GetObjDebugOk

`func (o *SalaryGetCommunicationListV1Response) GetObjDebugOk() (*CommonResponseObjDebug, bool)`

GetObjDebugOk returns a tuple with the ObjDebug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDebug

`func (o *SalaryGetCommunicationListV1Response) SetObjDebug(v CommonResponseObjDebug)`

SetObjDebug sets ObjDebug field to given value.

### HasObjDebug

`func (o *SalaryGetCommunicationListV1Response) HasObjDebug() bool`

HasObjDebug returns a boolean if a field has been set.

### GetMPayload

`func (o *SalaryGetCommunicationListV1Response) GetMPayload() SalaryGetCommunicationListV1ResponseMPayload`

GetMPayload returns the MPayload field if non-nil, zero value otherwise.

### GetMPayloadOk

`func (o *SalaryGetCommunicationListV1Response) GetMPayloadOk() (*SalaryGetCommunicationListV1ResponseMPayload, bool)`

GetMPayloadOk returns a tuple with the MPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMPayload

`func (o *SalaryGetCommunicationListV1Response) SetMPayload(v SalaryGetCommunicationListV1ResponseMPayload)`

SetMPayload sets MPayload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


