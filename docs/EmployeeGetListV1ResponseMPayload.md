# EmployeeGetListV1ResponseMPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IRowReturned** | **int32** | The number of rows returned | 
**IRowFiltered** | **int32** | The number of rows matching your filters (if any) or the total number of rows | 
**AObjEmployee** | [**[]EmployeeListElement**](EmployeeListElement.md) |  | 

## Methods

### NewEmployeeGetListV1ResponseMPayload

`func NewEmployeeGetListV1ResponseMPayload(iRowReturned int32, iRowFiltered int32, aObjEmployee []EmployeeListElement, ) *EmployeeGetListV1ResponseMPayload`

NewEmployeeGetListV1ResponseMPayload instantiates a new EmployeeGetListV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmployeeGetListV1ResponseMPayloadWithDefaults

`func NewEmployeeGetListV1ResponseMPayloadWithDefaults() *EmployeeGetListV1ResponseMPayload`

NewEmployeeGetListV1ResponseMPayloadWithDefaults instantiates a new EmployeeGetListV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIRowReturned

`func (o *EmployeeGetListV1ResponseMPayload) GetIRowReturned() int32`

GetIRowReturned returns the IRowReturned field if non-nil, zero value otherwise.

### GetIRowReturnedOk

`func (o *EmployeeGetListV1ResponseMPayload) GetIRowReturnedOk() (*int32, bool)`

GetIRowReturnedOk returns a tuple with the IRowReturned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowReturned

`func (o *EmployeeGetListV1ResponseMPayload) SetIRowReturned(v int32)`

SetIRowReturned sets IRowReturned field to given value.


### GetIRowFiltered

`func (o *EmployeeGetListV1ResponseMPayload) GetIRowFiltered() int32`

GetIRowFiltered returns the IRowFiltered field if non-nil, zero value otherwise.

### GetIRowFilteredOk

`func (o *EmployeeGetListV1ResponseMPayload) GetIRowFilteredOk() (*int32, bool)`

GetIRowFilteredOk returns a tuple with the IRowFiltered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowFiltered

`func (o *EmployeeGetListV1ResponseMPayload) SetIRowFiltered(v int32)`

SetIRowFiltered sets IRowFiltered field to given value.


### GetAObjEmployee

`func (o *EmployeeGetListV1ResponseMPayload) GetAObjEmployee() []EmployeeListElement`

GetAObjEmployee returns the AObjEmployee field if non-nil, zero value otherwise.

### GetAObjEmployeeOk

`func (o *EmployeeGetListV1ResponseMPayload) GetAObjEmployeeOk() (*[]EmployeeListElement, bool)`

GetAObjEmployeeOk returns a tuple with the AObjEmployee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEmployee

`func (o *EmployeeGetListV1ResponseMPayload) SetAObjEmployee(v []EmployeeListElement)`

SetAObjEmployee sets AObjEmployee field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


