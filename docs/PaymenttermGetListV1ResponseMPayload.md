# PaymenttermGetListV1ResponseMPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IRowReturned** | **int32** | The number of rows returned | 
**IRowFiltered** | **int32** | The number of rows matching your filters (if any) or the total number of rows | 
**AObjPaymentterm** | [**[]PaymenttermListElement**](PaymenttermListElement.md) |  | 

## Methods

### NewPaymenttermGetListV1ResponseMPayload

`func NewPaymenttermGetListV1ResponseMPayload(iRowReturned int32, iRowFiltered int32, aObjPaymentterm []PaymenttermListElement, ) *PaymenttermGetListV1ResponseMPayload`

NewPaymenttermGetListV1ResponseMPayload instantiates a new PaymenttermGetListV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymenttermGetListV1ResponseMPayloadWithDefaults

`func NewPaymenttermGetListV1ResponseMPayloadWithDefaults() *PaymenttermGetListV1ResponseMPayload`

NewPaymenttermGetListV1ResponseMPayloadWithDefaults instantiates a new PaymenttermGetListV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIRowReturned

`func (o *PaymenttermGetListV1ResponseMPayload) GetIRowReturned() int32`

GetIRowReturned returns the IRowReturned field if non-nil, zero value otherwise.

### GetIRowReturnedOk

`func (o *PaymenttermGetListV1ResponseMPayload) GetIRowReturnedOk() (*int32, bool)`

GetIRowReturnedOk returns a tuple with the IRowReturned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowReturned

`func (o *PaymenttermGetListV1ResponseMPayload) SetIRowReturned(v int32)`

SetIRowReturned sets IRowReturned field to given value.


### GetIRowFiltered

`func (o *PaymenttermGetListV1ResponseMPayload) GetIRowFiltered() int32`

GetIRowFiltered returns the IRowFiltered field if non-nil, zero value otherwise.

### GetIRowFilteredOk

`func (o *PaymenttermGetListV1ResponseMPayload) GetIRowFilteredOk() (*int32, bool)`

GetIRowFilteredOk returns a tuple with the IRowFiltered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowFiltered

`func (o *PaymenttermGetListV1ResponseMPayload) SetIRowFiltered(v int32)`

SetIRowFiltered sets IRowFiltered field to given value.


### GetAObjPaymentterm

`func (o *PaymenttermGetListV1ResponseMPayload) GetAObjPaymentterm() []PaymenttermListElement`

GetAObjPaymentterm returns the AObjPaymentterm field if non-nil, zero value otherwise.

### GetAObjPaymenttermOk

`func (o *PaymenttermGetListV1ResponseMPayload) GetAObjPaymenttermOk() (*[]PaymenttermListElement, bool)`

GetAObjPaymenttermOk returns a tuple with the AObjPaymentterm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjPaymentterm

`func (o *PaymenttermGetListV1ResponseMPayload) SetAObjPaymentterm(v []PaymenttermListElement)`

SetAObjPaymentterm sets AObjPaymentterm field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


