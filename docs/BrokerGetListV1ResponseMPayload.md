# BrokerGetListV1ResponseMPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IRowReturned** | **int32** | The number of rows returned | 
**IRowFiltered** | **int32** | The number of rows matching your filters (if any) or the total number of rows | 
**AObjBroker** | [**[]BrokerListElement**](BrokerListElement.md) |  | 

## Methods

### NewBrokerGetListV1ResponseMPayload

`func NewBrokerGetListV1ResponseMPayload(iRowReturned int32, iRowFiltered int32, aObjBroker []BrokerListElement, ) *BrokerGetListV1ResponseMPayload`

NewBrokerGetListV1ResponseMPayload instantiates a new BrokerGetListV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrokerGetListV1ResponseMPayloadWithDefaults

`func NewBrokerGetListV1ResponseMPayloadWithDefaults() *BrokerGetListV1ResponseMPayload`

NewBrokerGetListV1ResponseMPayloadWithDefaults instantiates a new BrokerGetListV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIRowReturned

`func (o *BrokerGetListV1ResponseMPayload) GetIRowReturned() int32`

GetIRowReturned returns the IRowReturned field if non-nil, zero value otherwise.

### GetIRowReturnedOk

`func (o *BrokerGetListV1ResponseMPayload) GetIRowReturnedOk() (*int32, bool)`

GetIRowReturnedOk returns a tuple with the IRowReturned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowReturned

`func (o *BrokerGetListV1ResponseMPayload) SetIRowReturned(v int32)`

SetIRowReturned sets IRowReturned field to given value.


### GetIRowFiltered

`func (o *BrokerGetListV1ResponseMPayload) GetIRowFiltered() int32`

GetIRowFiltered returns the IRowFiltered field if non-nil, zero value otherwise.

### GetIRowFilteredOk

`func (o *BrokerGetListV1ResponseMPayload) GetIRowFilteredOk() (*int32, bool)`

GetIRowFilteredOk returns a tuple with the IRowFiltered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowFiltered

`func (o *BrokerGetListV1ResponseMPayload) SetIRowFiltered(v int32)`

SetIRowFiltered sets IRowFiltered field to given value.


### GetAObjBroker

`func (o *BrokerGetListV1ResponseMPayload) GetAObjBroker() []BrokerListElement`

GetAObjBroker returns the AObjBroker field if non-nil, zero value otherwise.

### GetAObjBrokerOk

`func (o *BrokerGetListV1ResponseMPayload) GetAObjBrokerOk() (*[]BrokerListElement, bool)`

GetAObjBrokerOk returns a tuple with the AObjBroker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjBroker

`func (o *BrokerGetListV1ResponseMPayload) SetAObjBroker(v []BrokerListElement)`

SetAObjBroker sets AObjBroker field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


