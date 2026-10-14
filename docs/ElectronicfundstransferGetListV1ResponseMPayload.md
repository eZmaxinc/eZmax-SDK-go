# ElectronicfundstransferGetListV1ResponseMPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IRowReturned** | **int32** | The number of rows returned | 
**IRowFiltered** | **int32** | The number of rows matching your filters (if any) or the total number of rows | 
**AObjElectronicfundstransfer** | [**[]ElectronicfundstransferListElement**](ElectronicfundstransferListElement.md) |  | 

## Methods

### NewElectronicfundstransferGetListV1ResponseMPayload

`func NewElectronicfundstransferGetListV1ResponseMPayload(iRowReturned int32, iRowFiltered int32, aObjElectronicfundstransfer []ElectronicfundstransferListElement, ) *ElectronicfundstransferGetListV1ResponseMPayload`

NewElectronicfundstransferGetListV1ResponseMPayload instantiates a new ElectronicfundstransferGetListV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewElectronicfundstransferGetListV1ResponseMPayloadWithDefaults

`func NewElectronicfundstransferGetListV1ResponseMPayloadWithDefaults() *ElectronicfundstransferGetListV1ResponseMPayload`

NewElectronicfundstransferGetListV1ResponseMPayloadWithDefaults instantiates a new ElectronicfundstransferGetListV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIRowReturned

`func (o *ElectronicfundstransferGetListV1ResponseMPayload) GetIRowReturned() int32`

GetIRowReturned returns the IRowReturned field if non-nil, zero value otherwise.

### GetIRowReturnedOk

`func (o *ElectronicfundstransferGetListV1ResponseMPayload) GetIRowReturnedOk() (*int32, bool)`

GetIRowReturnedOk returns a tuple with the IRowReturned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowReturned

`func (o *ElectronicfundstransferGetListV1ResponseMPayload) SetIRowReturned(v int32)`

SetIRowReturned sets IRowReturned field to given value.


### GetIRowFiltered

`func (o *ElectronicfundstransferGetListV1ResponseMPayload) GetIRowFiltered() int32`

GetIRowFiltered returns the IRowFiltered field if non-nil, zero value otherwise.

### GetIRowFilteredOk

`func (o *ElectronicfundstransferGetListV1ResponseMPayload) GetIRowFilteredOk() (*int32, bool)`

GetIRowFilteredOk returns a tuple with the IRowFiltered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowFiltered

`func (o *ElectronicfundstransferGetListV1ResponseMPayload) SetIRowFiltered(v int32)`

SetIRowFiltered sets IRowFiltered field to given value.


### GetAObjElectronicfundstransfer

`func (o *ElectronicfundstransferGetListV1ResponseMPayload) GetAObjElectronicfundstransfer() []ElectronicfundstransferListElement`

GetAObjElectronicfundstransfer returns the AObjElectronicfundstransfer field if non-nil, zero value otherwise.

### GetAObjElectronicfundstransferOk

`func (o *ElectronicfundstransferGetListV1ResponseMPayload) GetAObjElectronicfundstransferOk() (*[]ElectronicfundstransferListElement, bool)`

GetAObjElectronicfundstransferOk returns a tuple with the AObjElectronicfundstransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjElectronicfundstransfer

`func (o *ElectronicfundstransferGetListV1ResponseMPayload) SetAObjElectronicfundstransfer(v []ElectronicfundstransferListElement)`

SetAObjElectronicfundstransfer sets AObjElectronicfundstransfer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


