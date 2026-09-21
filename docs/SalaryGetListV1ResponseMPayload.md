# SalaryGetListV1ResponseMPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IRowReturned** | **int32** | The number of rows returned | 
**IRowFiltered** | **int32** | The number of rows matching your filters (if any) or the total number of rows | 
**AObjSalary** | [**[]SalaryListElement**](SalaryListElement.md) |  | 

## Methods

### NewSalaryGetListV1ResponseMPayload

`func NewSalaryGetListV1ResponseMPayload(iRowReturned int32, iRowFiltered int32, aObjSalary []SalaryListElement, ) *SalaryGetListV1ResponseMPayload`

NewSalaryGetListV1ResponseMPayload instantiates a new SalaryGetListV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSalaryGetListV1ResponseMPayloadWithDefaults

`func NewSalaryGetListV1ResponseMPayloadWithDefaults() *SalaryGetListV1ResponseMPayload`

NewSalaryGetListV1ResponseMPayloadWithDefaults instantiates a new SalaryGetListV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIRowReturned

`func (o *SalaryGetListV1ResponseMPayload) GetIRowReturned() int32`

GetIRowReturned returns the IRowReturned field if non-nil, zero value otherwise.

### GetIRowReturnedOk

`func (o *SalaryGetListV1ResponseMPayload) GetIRowReturnedOk() (*int32, bool)`

GetIRowReturnedOk returns a tuple with the IRowReturned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowReturned

`func (o *SalaryGetListV1ResponseMPayload) SetIRowReturned(v int32)`

SetIRowReturned sets IRowReturned field to given value.


### GetIRowFiltered

`func (o *SalaryGetListV1ResponseMPayload) GetIRowFiltered() int32`

GetIRowFiltered returns the IRowFiltered field if non-nil, zero value otherwise.

### GetIRowFilteredOk

`func (o *SalaryGetListV1ResponseMPayload) GetIRowFilteredOk() (*int32, bool)`

GetIRowFilteredOk returns a tuple with the IRowFiltered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowFiltered

`func (o *SalaryGetListV1ResponseMPayload) SetIRowFiltered(v int32)`

SetIRowFiltered sets IRowFiltered field to given value.


### GetAObjSalary

`func (o *SalaryGetListV1ResponseMPayload) GetAObjSalary() []SalaryListElement`

GetAObjSalary returns the AObjSalary field if non-nil, zero value otherwise.

### GetAObjSalaryOk

`func (o *SalaryGetListV1ResponseMPayload) GetAObjSalaryOk() (*[]SalaryListElement, bool)`

GetAObjSalaryOk returns a tuple with the AObjSalary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjSalary

`func (o *SalaryGetListV1ResponseMPayload) SetAObjSalary(v []SalaryListElement)`

SetAObjSalary sets AObjSalary field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


