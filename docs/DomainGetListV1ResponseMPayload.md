# DomainGetListV1ResponseMPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IRowReturned** | **int32** | The number of rows returned | 
**IRowFiltered** | **int32** | The number of rows matching your filters (if any) or the total number of rows | 
**AObjDomain** | [**[]DomainListElement**](DomainListElement.md) |  | 

## Methods

### NewDomainGetListV1ResponseMPayload

`func NewDomainGetListV1ResponseMPayload(iRowReturned int32, iRowFiltered int32, aObjDomain []DomainListElement, ) *DomainGetListV1ResponseMPayload`

NewDomainGetListV1ResponseMPayload instantiates a new DomainGetListV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainGetListV1ResponseMPayloadWithDefaults

`func NewDomainGetListV1ResponseMPayloadWithDefaults() *DomainGetListV1ResponseMPayload`

NewDomainGetListV1ResponseMPayloadWithDefaults instantiates a new DomainGetListV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIRowReturned

`func (o *DomainGetListV1ResponseMPayload) GetIRowReturned() int32`

GetIRowReturned returns the IRowReturned field if non-nil, zero value otherwise.

### GetIRowReturnedOk

`func (o *DomainGetListV1ResponseMPayload) GetIRowReturnedOk() (*int32, bool)`

GetIRowReturnedOk returns a tuple with the IRowReturned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowReturned

`func (o *DomainGetListV1ResponseMPayload) SetIRowReturned(v int32)`

SetIRowReturned sets IRowReturned field to given value.


### GetIRowFiltered

`func (o *DomainGetListV1ResponseMPayload) GetIRowFiltered() int32`

GetIRowFiltered returns the IRowFiltered field if non-nil, zero value otherwise.

### GetIRowFilteredOk

`func (o *DomainGetListV1ResponseMPayload) GetIRowFilteredOk() (*int32, bool)`

GetIRowFilteredOk returns a tuple with the IRowFiltered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowFiltered

`func (o *DomainGetListV1ResponseMPayload) SetIRowFiltered(v int32)`

SetIRowFiltered sets IRowFiltered field to given value.


### GetAObjDomain

`func (o *DomainGetListV1ResponseMPayload) GetAObjDomain() []DomainListElement`

GetAObjDomain returns the AObjDomain field if non-nil, zero value otherwise.

### GetAObjDomainOk

`func (o *DomainGetListV1ResponseMPayload) GetAObjDomainOk() (*[]DomainListElement, bool)`

GetAObjDomainOk returns a tuple with the AObjDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjDomain

`func (o *DomainGetListV1ResponseMPayload) SetAObjDomain(v []DomainListElement)`

SetAObjDomain sets AObjDomain field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


