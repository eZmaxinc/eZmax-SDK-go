# BuyercontractGetListV1ResponseMPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IRowReturned** | **int32** | The number of rows returned | 
**IRowFiltered** | **int32** | The number of rows matching your filters (if any) or the total number of rows | 
**AObjBuyercontract** | [**[]BuyercontractListElement**](BuyercontractListElement.md) |  | 

## Methods

### NewBuyercontractGetListV1ResponseMPayload

`func NewBuyercontractGetListV1ResponseMPayload(iRowReturned int32, iRowFiltered int32, aObjBuyercontract []BuyercontractListElement, ) *BuyercontractGetListV1ResponseMPayload`

NewBuyercontractGetListV1ResponseMPayload instantiates a new BuyercontractGetListV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuyercontractGetListV1ResponseMPayloadWithDefaults

`func NewBuyercontractGetListV1ResponseMPayloadWithDefaults() *BuyercontractGetListV1ResponseMPayload`

NewBuyercontractGetListV1ResponseMPayloadWithDefaults instantiates a new BuyercontractGetListV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIRowReturned

`func (o *BuyercontractGetListV1ResponseMPayload) GetIRowReturned() int32`

GetIRowReturned returns the IRowReturned field if non-nil, zero value otherwise.

### GetIRowReturnedOk

`func (o *BuyercontractGetListV1ResponseMPayload) GetIRowReturnedOk() (*int32, bool)`

GetIRowReturnedOk returns a tuple with the IRowReturned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowReturned

`func (o *BuyercontractGetListV1ResponseMPayload) SetIRowReturned(v int32)`

SetIRowReturned sets IRowReturned field to given value.


### GetIRowFiltered

`func (o *BuyercontractGetListV1ResponseMPayload) GetIRowFiltered() int32`

GetIRowFiltered returns the IRowFiltered field if non-nil, zero value otherwise.

### GetIRowFilteredOk

`func (o *BuyercontractGetListV1ResponseMPayload) GetIRowFilteredOk() (*int32, bool)`

GetIRowFilteredOk returns a tuple with the IRowFiltered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowFiltered

`func (o *BuyercontractGetListV1ResponseMPayload) SetIRowFiltered(v int32)`

SetIRowFiltered sets IRowFiltered field to given value.


### GetAObjBuyercontract

`func (o *BuyercontractGetListV1ResponseMPayload) GetAObjBuyercontract() []BuyercontractListElement`

GetAObjBuyercontract returns the AObjBuyercontract field if non-nil, zero value otherwise.

### GetAObjBuyercontractOk

`func (o *BuyercontractGetListV1ResponseMPayload) GetAObjBuyercontractOk() (*[]BuyercontractListElement, bool)`

GetAObjBuyercontractOk returns a tuple with the AObjBuyercontract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjBuyercontract

`func (o *BuyercontractGetListV1ResponseMPayload) SetAObjBuyercontract(v []BuyercontractListElement)`

SetAObjBuyercontract sets AObjBuyercontract field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


