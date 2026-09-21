# InvoiceGetListV1ResponseMPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IRowReturned** | **int32** | The number of rows returned | 
**IRowFiltered** | **int32** | The number of rows matching your filters (if any) or the total number of rows | 
**AObjInvoice** | [**[]InvoiceListElement**](InvoiceListElement.md) |  | 

## Methods

### NewInvoiceGetListV1ResponseMPayload

`func NewInvoiceGetListV1ResponseMPayload(iRowReturned int32, iRowFiltered int32, aObjInvoice []InvoiceListElement, ) *InvoiceGetListV1ResponseMPayload`

NewInvoiceGetListV1ResponseMPayload instantiates a new InvoiceGetListV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceGetListV1ResponseMPayloadWithDefaults

`func NewInvoiceGetListV1ResponseMPayloadWithDefaults() *InvoiceGetListV1ResponseMPayload`

NewInvoiceGetListV1ResponseMPayloadWithDefaults instantiates a new InvoiceGetListV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIRowReturned

`func (o *InvoiceGetListV1ResponseMPayload) GetIRowReturned() int32`

GetIRowReturned returns the IRowReturned field if non-nil, zero value otherwise.

### GetIRowReturnedOk

`func (o *InvoiceGetListV1ResponseMPayload) GetIRowReturnedOk() (*int32, bool)`

GetIRowReturnedOk returns a tuple with the IRowReturned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowReturned

`func (o *InvoiceGetListV1ResponseMPayload) SetIRowReturned(v int32)`

SetIRowReturned sets IRowReturned field to given value.


### GetIRowFiltered

`func (o *InvoiceGetListV1ResponseMPayload) GetIRowFiltered() int32`

GetIRowFiltered returns the IRowFiltered field if non-nil, zero value otherwise.

### GetIRowFilteredOk

`func (o *InvoiceGetListV1ResponseMPayload) GetIRowFilteredOk() (*int32, bool)`

GetIRowFilteredOk returns a tuple with the IRowFiltered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowFiltered

`func (o *InvoiceGetListV1ResponseMPayload) SetIRowFiltered(v int32)`

SetIRowFiltered sets IRowFiltered field to given value.


### GetAObjInvoice

`func (o *InvoiceGetListV1ResponseMPayload) GetAObjInvoice() []InvoiceListElement`

GetAObjInvoice returns the AObjInvoice field if non-nil, zero value otherwise.

### GetAObjInvoiceOk

`func (o *InvoiceGetListV1ResponseMPayload) GetAObjInvoiceOk() (*[]InvoiceListElement, bool)`

GetAObjInvoiceOk returns a tuple with the AObjInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjInvoice

`func (o *InvoiceGetListV1ResponseMPayload) SetAObjInvoice(v []InvoiceListElement)`

SetAObjInvoice sets AObjInvoice field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


