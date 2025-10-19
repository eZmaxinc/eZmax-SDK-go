# LeadGetListV1ResponseMPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IRowReturned** | **int32** | The number of rows returned | 
**IRowFiltered** | **int32** | The number of rows matching your filters (if any) or the total number of rows | 
**AObjLead** | [**[]LeadListElement**](LeadListElement.md) |  | 

## Methods

### NewLeadGetListV1ResponseMPayload

`func NewLeadGetListV1ResponseMPayload(iRowReturned int32, iRowFiltered int32, aObjLead []LeadListElement, ) *LeadGetListV1ResponseMPayload`

NewLeadGetListV1ResponseMPayload instantiates a new LeadGetListV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeadGetListV1ResponseMPayloadWithDefaults

`func NewLeadGetListV1ResponseMPayloadWithDefaults() *LeadGetListV1ResponseMPayload`

NewLeadGetListV1ResponseMPayloadWithDefaults instantiates a new LeadGetListV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIRowReturned

`func (o *LeadGetListV1ResponseMPayload) GetIRowReturned() int32`

GetIRowReturned returns the IRowReturned field if non-nil, zero value otherwise.

### GetIRowReturnedOk

`func (o *LeadGetListV1ResponseMPayload) GetIRowReturnedOk() (*int32, bool)`

GetIRowReturnedOk returns a tuple with the IRowReturned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowReturned

`func (o *LeadGetListV1ResponseMPayload) SetIRowReturned(v int32)`

SetIRowReturned sets IRowReturned field to given value.


### GetIRowFiltered

`func (o *LeadGetListV1ResponseMPayload) GetIRowFiltered() int32`

GetIRowFiltered returns the IRowFiltered field if non-nil, zero value otherwise.

### GetIRowFilteredOk

`func (o *LeadGetListV1ResponseMPayload) GetIRowFilteredOk() (*int32, bool)`

GetIRowFilteredOk returns a tuple with the IRowFiltered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowFiltered

`func (o *LeadGetListV1ResponseMPayload) SetIRowFiltered(v int32)`

SetIRowFiltered sets IRowFiltered field to given value.


### GetAObjLead

`func (o *LeadGetListV1ResponseMPayload) GetAObjLead() []LeadListElement`

GetAObjLead returns the AObjLead field if non-nil, zero value otherwise.

### GetAObjLeadOk

`func (o *LeadGetListV1ResponseMPayload) GetAObjLeadOk() (*[]LeadListElement, bool)`

GetAObjLeadOk returns a tuple with the AObjLead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjLead

`func (o *LeadGetListV1ResponseMPayload) SetAObjLead(v []LeadListElement)`

SetAObjLead sets AObjLead field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


