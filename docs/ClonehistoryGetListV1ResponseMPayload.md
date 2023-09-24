# ClonehistoryGetListV1ResponseMPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IRowReturned** | **int32** | The number of rows returned | 
**IRowFiltered** | **int32** | The number of rows matching your filters (if any) or the total number of rows | 
**AObjClonehistory** | [**[]ClonehistoryListElement**](ClonehistoryListElement.md) |  | 

## Methods

### NewClonehistoryGetListV1ResponseMPayload

`func NewClonehistoryGetListV1ResponseMPayload(iRowReturned int32, iRowFiltered int32, aObjClonehistory []ClonehistoryListElement, ) *ClonehistoryGetListV1ResponseMPayload`

NewClonehistoryGetListV1ResponseMPayload instantiates a new ClonehistoryGetListV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClonehistoryGetListV1ResponseMPayloadWithDefaults

`func NewClonehistoryGetListV1ResponseMPayloadWithDefaults() *ClonehistoryGetListV1ResponseMPayload`

NewClonehistoryGetListV1ResponseMPayloadWithDefaults instantiates a new ClonehistoryGetListV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIRowReturned

`func (o *ClonehistoryGetListV1ResponseMPayload) GetIRowReturned() int32`

GetIRowReturned returns the IRowReturned field if non-nil, zero value otherwise.

### GetIRowReturnedOk

`func (o *ClonehistoryGetListV1ResponseMPayload) GetIRowReturnedOk() (*int32, bool)`

GetIRowReturnedOk returns a tuple with the IRowReturned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowReturned

`func (o *ClonehistoryGetListV1ResponseMPayload) SetIRowReturned(v int32)`

SetIRowReturned sets IRowReturned field to given value.


### GetIRowFiltered

`func (o *ClonehistoryGetListV1ResponseMPayload) GetIRowFiltered() int32`

GetIRowFiltered returns the IRowFiltered field if non-nil, zero value otherwise.

### GetIRowFilteredOk

`func (o *ClonehistoryGetListV1ResponseMPayload) GetIRowFilteredOk() (*int32, bool)`

GetIRowFilteredOk returns a tuple with the IRowFiltered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowFiltered

`func (o *ClonehistoryGetListV1ResponseMPayload) SetIRowFiltered(v int32)`

SetIRowFiltered sets IRowFiltered field to given value.


### GetAObjClonehistory

`func (o *ClonehistoryGetListV1ResponseMPayload) GetAObjClonehistory() []ClonehistoryListElement`

GetAObjClonehistory returns the AObjClonehistory field if non-nil, zero value otherwise.

### GetAObjClonehistoryOk

`func (o *ClonehistoryGetListV1ResponseMPayload) GetAObjClonehistoryOk() (*[]ClonehistoryListElement, bool)`

GetAObjClonehistoryOk returns a tuple with the AObjClonehistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjClonehistory

`func (o *ClonehistoryGetListV1ResponseMPayload) SetAObjClonehistory(v []ClonehistoryListElement)`

SetAObjClonehistory sets AObjClonehistory field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


