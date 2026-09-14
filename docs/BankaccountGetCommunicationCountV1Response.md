# BankaccountGetCommunicationCountV1Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjDebugPayload** | [**CommonResponseObjDebugPayload**](CommonResponseObjDebugPayload.md) |  | 
**ObjDebug** | Pointer to [**CommonResponseObjDebug**](CommonResponseObjDebug.md) |  | [optional] 
**MPayload** | [**BankaccountGetCommunicationCountV1ResponseMPayload**](BankaccountGetCommunicationCountV1ResponseMPayload.md) |  | 

## Methods

### NewBankaccountGetCommunicationCountV1Response

`func NewBankaccountGetCommunicationCountV1Response(objDebugPayload CommonResponseObjDebugPayload, mPayload BankaccountGetCommunicationCountV1ResponseMPayload, ) *BankaccountGetCommunicationCountV1Response`

NewBankaccountGetCommunicationCountV1Response instantiates a new BankaccountGetCommunicationCountV1Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankaccountGetCommunicationCountV1ResponseWithDefaults

`func NewBankaccountGetCommunicationCountV1ResponseWithDefaults() *BankaccountGetCommunicationCountV1Response`

NewBankaccountGetCommunicationCountV1ResponseWithDefaults instantiates a new BankaccountGetCommunicationCountV1Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjDebugPayload

`func (o *BankaccountGetCommunicationCountV1Response) GetObjDebugPayload() CommonResponseObjDebugPayload`

GetObjDebugPayload returns the ObjDebugPayload field if non-nil, zero value otherwise.

### GetObjDebugPayloadOk

`func (o *BankaccountGetCommunicationCountV1Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool)`

GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDebugPayload

`func (o *BankaccountGetCommunicationCountV1Response) SetObjDebugPayload(v CommonResponseObjDebugPayload)`

SetObjDebugPayload sets ObjDebugPayload field to given value.


### GetObjDebug

`func (o *BankaccountGetCommunicationCountV1Response) GetObjDebug() CommonResponseObjDebug`

GetObjDebug returns the ObjDebug field if non-nil, zero value otherwise.

### GetObjDebugOk

`func (o *BankaccountGetCommunicationCountV1Response) GetObjDebugOk() (*CommonResponseObjDebug, bool)`

GetObjDebugOk returns a tuple with the ObjDebug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDebug

`func (o *BankaccountGetCommunicationCountV1Response) SetObjDebug(v CommonResponseObjDebug)`

SetObjDebug sets ObjDebug field to given value.

### HasObjDebug

`func (o *BankaccountGetCommunicationCountV1Response) HasObjDebug() bool`

HasObjDebug returns a boolean if a field has been set.

### GetMPayload

`func (o *BankaccountGetCommunicationCountV1Response) GetMPayload() BankaccountGetCommunicationCountV1ResponseMPayload`

GetMPayload returns the MPayload field if non-nil, zero value otherwise.

### GetMPayloadOk

`func (o *BankaccountGetCommunicationCountV1Response) GetMPayloadOk() (*BankaccountGetCommunicationCountV1ResponseMPayload, bool)`

GetMPayloadOk returns a tuple with the MPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMPayload

`func (o *BankaccountGetCommunicationCountV1Response) SetMPayload(v BankaccountGetCommunicationCountV1ResponseMPayload)`

SetMPayload sets MPayload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


