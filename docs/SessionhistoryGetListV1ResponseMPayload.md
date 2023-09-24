# SessionhistoryGetListV1ResponseMPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IRowReturned** | **int32** | The number of rows returned | 
**IRowFiltered** | **int32** | The number of rows matching your filters (if any) or the total number of rows | 
**AObjSessionhistory** | [**[]SessionhistoryListElement**](SessionhistoryListElement.md) |  | 

## Methods

### NewSessionhistoryGetListV1ResponseMPayload

`func NewSessionhistoryGetListV1ResponseMPayload(iRowReturned int32, iRowFiltered int32, aObjSessionhistory []SessionhistoryListElement, ) *SessionhistoryGetListV1ResponseMPayload`

NewSessionhistoryGetListV1ResponseMPayload instantiates a new SessionhistoryGetListV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionhistoryGetListV1ResponseMPayloadWithDefaults

`func NewSessionhistoryGetListV1ResponseMPayloadWithDefaults() *SessionhistoryGetListV1ResponseMPayload`

NewSessionhistoryGetListV1ResponseMPayloadWithDefaults instantiates a new SessionhistoryGetListV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIRowReturned

`func (o *SessionhistoryGetListV1ResponseMPayload) GetIRowReturned() int32`

GetIRowReturned returns the IRowReturned field if non-nil, zero value otherwise.

### GetIRowReturnedOk

`func (o *SessionhistoryGetListV1ResponseMPayload) GetIRowReturnedOk() (*int32, bool)`

GetIRowReturnedOk returns a tuple with the IRowReturned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowReturned

`func (o *SessionhistoryGetListV1ResponseMPayload) SetIRowReturned(v int32)`

SetIRowReturned sets IRowReturned field to given value.


### GetIRowFiltered

`func (o *SessionhistoryGetListV1ResponseMPayload) GetIRowFiltered() int32`

GetIRowFiltered returns the IRowFiltered field if non-nil, zero value otherwise.

### GetIRowFilteredOk

`func (o *SessionhistoryGetListV1ResponseMPayload) GetIRowFilteredOk() (*int32, bool)`

GetIRowFilteredOk returns a tuple with the IRowFiltered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowFiltered

`func (o *SessionhistoryGetListV1ResponseMPayload) SetIRowFiltered(v int32)`

SetIRowFiltered sets IRowFiltered field to given value.


### GetAObjSessionhistory

`func (o *SessionhistoryGetListV1ResponseMPayload) GetAObjSessionhistory() []SessionhistoryListElement`

GetAObjSessionhistory returns the AObjSessionhistory field if non-nil, zero value otherwise.

### GetAObjSessionhistoryOk

`func (o *SessionhistoryGetListV1ResponseMPayload) GetAObjSessionhistoryOk() (*[]SessionhistoryListElement, bool)`

GetAObjSessionhistoryOk returns a tuple with the AObjSessionhistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjSessionhistory

`func (o *SessionhistoryGetListV1ResponseMPayload) SetAObjSessionhistory(v []SessionhistoryListElement)`

SetAObjSessionhistory sets AObjSessionhistory field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


