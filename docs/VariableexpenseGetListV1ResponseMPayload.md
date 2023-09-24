# VariableexpenseGetListV1ResponseMPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IRowReturned** | **int32** | The number of rows returned | 
**IRowFiltered** | **int32** | The number of rows matching your filters (if any) or the total number of rows | 
**AObjVariableexpense** | [**[]VariableexpenseListElement**](VariableexpenseListElement.md) |  | 

## Methods

### NewVariableexpenseGetListV1ResponseMPayload

`func NewVariableexpenseGetListV1ResponseMPayload(iRowReturned int32, iRowFiltered int32, aObjVariableexpense []VariableexpenseListElement, ) *VariableexpenseGetListV1ResponseMPayload`

NewVariableexpenseGetListV1ResponseMPayload instantiates a new VariableexpenseGetListV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariableexpenseGetListV1ResponseMPayloadWithDefaults

`func NewVariableexpenseGetListV1ResponseMPayloadWithDefaults() *VariableexpenseGetListV1ResponseMPayload`

NewVariableexpenseGetListV1ResponseMPayloadWithDefaults instantiates a new VariableexpenseGetListV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIRowReturned

`func (o *VariableexpenseGetListV1ResponseMPayload) GetIRowReturned() int32`

GetIRowReturned returns the IRowReturned field if non-nil, zero value otherwise.

### GetIRowReturnedOk

`func (o *VariableexpenseGetListV1ResponseMPayload) GetIRowReturnedOk() (*int32, bool)`

GetIRowReturnedOk returns a tuple with the IRowReturned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowReturned

`func (o *VariableexpenseGetListV1ResponseMPayload) SetIRowReturned(v int32)`

SetIRowReturned sets IRowReturned field to given value.


### GetIRowFiltered

`func (o *VariableexpenseGetListV1ResponseMPayload) GetIRowFiltered() int32`

GetIRowFiltered returns the IRowFiltered field if non-nil, zero value otherwise.

### GetIRowFilteredOk

`func (o *VariableexpenseGetListV1ResponseMPayload) GetIRowFilteredOk() (*int32, bool)`

GetIRowFilteredOk returns a tuple with the IRowFiltered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowFiltered

`func (o *VariableexpenseGetListV1ResponseMPayload) SetIRowFiltered(v int32)`

SetIRowFiltered sets IRowFiltered field to given value.


### GetAObjVariableexpense

`func (o *VariableexpenseGetListV1ResponseMPayload) GetAObjVariableexpense() []VariableexpenseListElement`

GetAObjVariableexpense returns the AObjVariableexpense field if non-nil, zero value otherwise.

### GetAObjVariableexpenseOk

`func (o *VariableexpenseGetListV1ResponseMPayload) GetAObjVariableexpenseOk() (*[]VariableexpenseListElement, bool)`

GetAObjVariableexpenseOk returns a tuple with the AObjVariableexpense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjVariableexpense

`func (o *VariableexpenseGetListV1ResponseMPayload) SetAObjVariableexpense(v []VariableexpenseListElement)`

SetAObjVariableexpense sets AObjVariableexpense field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


