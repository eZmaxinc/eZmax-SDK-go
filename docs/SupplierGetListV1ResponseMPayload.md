# SupplierGetListV1ResponseMPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IRowReturned** | **int32** | The number of rows returned | 
**IRowFiltered** | **int32** | The number of rows matching your filters (if any) or the total number of rows | 
**AObjSupplier** | [**[]SupplierListElement**](SupplierListElement.md) |  | 

## Methods

### NewSupplierGetListV1ResponseMPayload

`func NewSupplierGetListV1ResponseMPayload(iRowReturned int32, iRowFiltered int32, aObjSupplier []SupplierListElement, ) *SupplierGetListV1ResponseMPayload`

NewSupplierGetListV1ResponseMPayload instantiates a new SupplierGetListV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupplierGetListV1ResponseMPayloadWithDefaults

`func NewSupplierGetListV1ResponseMPayloadWithDefaults() *SupplierGetListV1ResponseMPayload`

NewSupplierGetListV1ResponseMPayloadWithDefaults instantiates a new SupplierGetListV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIRowReturned

`func (o *SupplierGetListV1ResponseMPayload) GetIRowReturned() int32`

GetIRowReturned returns the IRowReturned field if non-nil, zero value otherwise.

### GetIRowReturnedOk

`func (o *SupplierGetListV1ResponseMPayload) GetIRowReturnedOk() (*int32, bool)`

GetIRowReturnedOk returns a tuple with the IRowReturned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowReturned

`func (o *SupplierGetListV1ResponseMPayload) SetIRowReturned(v int32)`

SetIRowReturned sets IRowReturned field to given value.


### GetIRowFiltered

`func (o *SupplierGetListV1ResponseMPayload) GetIRowFiltered() int32`

GetIRowFiltered returns the IRowFiltered field if non-nil, zero value otherwise.

### GetIRowFilteredOk

`func (o *SupplierGetListV1ResponseMPayload) GetIRowFilteredOk() (*int32, bool)`

GetIRowFilteredOk returns a tuple with the IRowFiltered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowFiltered

`func (o *SupplierGetListV1ResponseMPayload) SetIRowFiltered(v int32)`

SetIRowFiltered sets IRowFiltered field to given value.


### GetAObjSupplier

`func (o *SupplierGetListV1ResponseMPayload) GetAObjSupplier() []SupplierListElement`

GetAObjSupplier returns the AObjSupplier field if non-nil, zero value otherwise.

### GetAObjSupplierOk

`func (o *SupplierGetListV1ResponseMPayload) GetAObjSupplierOk() (*[]SupplierListElement, bool)`

GetAObjSupplierOk returns a tuple with the AObjSupplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjSupplier

`func (o *SupplierGetListV1ResponseMPayload) SetAObjSupplier(v []SupplierListElement)`

SetAObjSupplier sets AObjSupplier field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


