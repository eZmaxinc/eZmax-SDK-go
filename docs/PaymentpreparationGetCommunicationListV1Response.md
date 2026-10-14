# PaymentpreparationGetCommunicationListV1Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjDebugPayload** | [**CommonResponseObjDebugPayloadGetList**](CommonResponseObjDebugPayloadGetList.md) |  | 
**ObjDebug** | Pointer to [**CommonResponseObjDebug**](CommonResponseObjDebug.md) |  | [optional] 
**MPayload** | [**PaymentpreparationGetCommunicationListV1ResponseMPayload**](PaymentpreparationGetCommunicationListV1ResponseMPayload.md) |  | 

## Methods

### NewPaymentpreparationGetCommunicationListV1Response

`func NewPaymentpreparationGetCommunicationListV1Response(objDebugPayload CommonResponseObjDebugPayloadGetList, mPayload PaymentpreparationGetCommunicationListV1ResponseMPayload, ) *PaymentpreparationGetCommunicationListV1Response`

NewPaymentpreparationGetCommunicationListV1Response instantiates a new PaymentpreparationGetCommunicationListV1Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentpreparationGetCommunicationListV1ResponseWithDefaults

`func NewPaymentpreparationGetCommunicationListV1ResponseWithDefaults() *PaymentpreparationGetCommunicationListV1Response`

NewPaymentpreparationGetCommunicationListV1ResponseWithDefaults instantiates a new PaymentpreparationGetCommunicationListV1Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjDebugPayload

`func (o *PaymentpreparationGetCommunicationListV1Response) GetObjDebugPayload() CommonResponseObjDebugPayloadGetList`

GetObjDebugPayload returns the ObjDebugPayload field if non-nil, zero value otherwise.

### GetObjDebugPayloadOk

`func (o *PaymentpreparationGetCommunicationListV1Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayloadGetList, bool)`

GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDebugPayload

`func (o *PaymentpreparationGetCommunicationListV1Response) SetObjDebugPayload(v CommonResponseObjDebugPayloadGetList)`

SetObjDebugPayload sets ObjDebugPayload field to given value.


### GetObjDebug

`func (o *PaymentpreparationGetCommunicationListV1Response) GetObjDebug() CommonResponseObjDebug`

GetObjDebug returns the ObjDebug field if non-nil, zero value otherwise.

### GetObjDebugOk

`func (o *PaymentpreparationGetCommunicationListV1Response) GetObjDebugOk() (*CommonResponseObjDebug, bool)`

GetObjDebugOk returns a tuple with the ObjDebug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDebug

`func (o *PaymentpreparationGetCommunicationListV1Response) SetObjDebug(v CommonResponseObjDebug)`

SetObjDebug sets ObjDebug field to given value.

### HasObjDebug

`func (o *PaymentpreparationGetCommunicationListV1Response) HasObjDebug() bool`

HasObjDebug returns a boolean if a field has been set.

### GetMPayload

`func (o *PaymentpreparationGetCommunicationListV1Response) GetMPayload() PaymentpreparationGetCommunicationListV1ResponseMPayload`

GetMPayload returns the MPayload field if non-nil, zero value otherwise.

### GetMPayloadOk

`func (o *PaymentpreparationGetCommunicationListV1Response) GetMPayloadOk() (*PaymentpreparationGetCommunicationListV1ResponseMPayload, bool)`

GetMPayloadOk returns a tuple with the MPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMPayload

`func (o *PaymentpreparationGetCommunicationListV1Response) SetMPayload(v PaymentpreparationGetCommunicationListV1ResponseMPayload)`

SetMPayload sets MPayload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


