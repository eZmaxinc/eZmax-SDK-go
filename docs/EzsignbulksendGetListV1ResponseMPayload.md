# EzsignbulksendGetListV1ResponseMPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AObjEzsignfolder** | Pointer to [**[]EzsignbulksendListElement**](EzsignbulksendListElement.md) |  | [optional] 
**IRowReturned** | **int32** | The number of rows returned | 
**IRowFiltered** | **int32** | The number of rows matching your filters (if any) or the total number of rows | 

## Methods

### NewEzsignbulksendGetListV1ResponseMPayload

`func NewEzsignbulksendGetListV1ResponseMPayload(iRowReturned int32, iRowFiltered int32, ) *EzsignbulksendGetListV1ResponseMPayload`

NewEzsignbulksendGetListV1ResponseMPayload instantiates a new EzsignbulksendGetListV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignbulksendGetListV1ResponseMPayloadWithDefaults

`func NewEzsignbulksendGetListV1ResponseMPayloadWithDefaults() *EzsignbulksendGetListV1ResponseMPayload`

NewEzsignbulksendGetListV1ResponseMPayloadWithDefaults instantiates a new EzsignbulksendGetListV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAObjEzsignfolder

`func (o *EzsignbulksendGetListV1ResponseMPayload) GetAObjEzsignfolder() []EzsignbulksendListElement`

GetAObjEzsignfolder returns the AObjEzsignfolder field if non-nil, zero value otherwise.

### GetAObjEzsignfolderOk

`func (o *EzsignbulksendGetListV1ResponseMPayload) GetAObjEzsignfolderOk() (*[]EzsignbulksendListElement, bool)`

GetAObjEzsignfolderOk returns a tuple with the AObjEzsignfolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzsignfolder

`func (o *EzsignbulksendGetListV1ResponseMPayload) SetAObjEzsignfolder(v []EzsignbulksendListElement)`

SetAObjEzsignfolder sets AObjEzsignfolder field to given value.

### HasAObjEzsignfolder

`func (o *EzsignbulksendGetListV1ResponseMPayload) HasAObjEzsignfolder() bool`

HasAObjEzsignfolder returns a boolean if a field has been set.

### GetIRowReturned

`func (o *EzsignbulksendGetListV1ResponseMPayload) GetIRowReturned() int32`

GetIRowReturned returns the IRowReturned field if non-nil, zero value otherwise.

### GetIRowReturnedOk

`func (o *EzsignbulksendGetListV1ResponseMPayload) GetIRowReturnedOk() (*int32, bool)`

GetIRowReturnedOk returns a tuple with the IRowReturned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowReturned

`func (o *EzsignbulksendGetListV1ResponseMPayload) SetIRowReturned(v int32)`

SetIRowReturned sets IRowReturned field to given value.


### GetIRowFiltered

`func (o *EzsignbulksendGetListV1ResponseMPayload) GetIRowFiltered() int32`

GetIRowFiltered returns the IRowFiltered field if non-nil, zero value otherwise.

### GetIRowFilteredOk

`func (o *EzsignbulksendGetListV1ResponseMPayload) GetIRowFilteredOk() (*int32, bool)`

GetIRowFilteredOk returns a tuple with the IRowFiltered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowFiltered

`func (o *EzsignbulksendGetListV1ResponseMPayload) SetIRowFiltered(v int32)`

SetIRowFiltered sets IRowFiltered field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


