# CustomerGetListV1ResponseMPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IRowReturned** | **int32** | The number of rows returned | 
**IRowFiltered** | **int32** | The number of rows matching your filters (if any) or the total number of rows | 
**AObjCustomer** | [**[]CustomerListElement**](CustomerListElement.md) |  | 

## Methods

### NewCustomerGetListV1ResponseMPayload

`func NewCustomerGetListV1ResponseMPayload(iRowReturned int32, iRowFiltered int32, aObjCustomer []CustomerListElement, ) *CustomerGetListV1ResponseMPayload`

NewCustomerGetListV1ResponseMPayload instantiates a new CustomerGetListV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerGetListV1ResponseMPayloadWithDefaults

`func NewCustomerGetListV1ResponseMPayloadWithDefaults() *CustomerGetListV1ResponseMPayload`

NewCustomerGetListV1ResponseMPayloadWithDefaults instantiates a new CustomerGetListV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIRowReturned

`func (o *CustomerGetListV1ResponseMPayload) GetIRowReturned() int32`

GetIRowReturned returns the IRowReturned field if non-nil, zero value otherwise.

### GetIRowReturnedOk

`func (o *CustomerGetListV1ResponseMPayload) GetIRowReturnedOk() (*int32, bool)`

GetIRowReturnedOk returns a tuple with the IRowReturned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowReturned

`func (o *CustomerGetListV1ResponseMPayload) SetIRowReturned(v int32)`

SetIRowReturned sets IRowReturned field to given value.


### GetIRowFiltered

`func (o *CustomerGetListV1ResponseMPayload) GetIRowFiltered() int32`

GetIRowFiltered returns the IRowFiltered field if non-nil, zero value otherwise.

### GetIRowFilteredOk

`func (o *CustomerGetListV1ResponseMPayload) GetIRowFilteredOk() (*int32, bool)`

GetIRowFilteredOk returns a tuple with the IRowFiltered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowFiltered

`func (o *CustomerGetListV1ResponseMPayload) SetIRowFiltered(v int32)`

SetIRowFiltered sets IRowFiltered field to given value.


### GetAObjCustomer

`func (o *CustomerGetListV1ResponseMPayload) GetAObjCustomer() []CustomerListElement`

GetAObjCustomer returns the AObjCustomer field if non-nil, zero value otherwise.

### GetAObjCustomerOk

`func (o *CustomerGetListV1ResponseMPayload) GetAObjCustomerOk() (*[]CustomerListElement, bool)`

GetAObjCustomerOk returns a tuple with the AObjCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjCustomer

`func (o *CustomerGetListV1ResponseMPayload) SetAObjCustomer(v []CustomerListElement)`

SetAObjCustomer sets AObjCustomer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


