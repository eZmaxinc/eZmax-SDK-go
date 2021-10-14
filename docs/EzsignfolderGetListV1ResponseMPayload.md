# EzsignfolderGetListV1ResponseMPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AObjEzsignfolder** | [**[]EzsignfolderListElement**](EzsignfolderListElement.md) |  | 
**IRowReturned** | **int32** | The number of rows returned | 
**IRowFiltered** | **int32** | The number of rows matching your filters (if any) or the total number of rows | 

## Methods

### NewEzsignfolderGetListV1ResponseMPayload

`func NewEzsignfolderGetListV1ResponseMPayload(aObjEzsignfolder []EzsignfolderListElement, iRowReturned int32, iRowFiltered int32, ) *EzsignfolderGetListV1ResponseMPayload`

NewEzsignfolderGetListV1ResponseMPayload instantiates a new EzsignfolderGetListV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignfolderGetListV1ResponseMPayloadWithDefaults

`func NewEzsignfolderGetListV1ResponseMPayloadWithDefaults() *EzsignfolderGetListV1ResponseMPayload`

NewEzsignfolderGetListV1ResponseMPayloadWithDefaults instantiates a new EzsignfolderGetListV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAObjEzsignfolder

`func (o *EzsignfolderGetListV1ResponseMPayload) GetAObjEzsignfolder() []EzsignfolderListElement`

GetAObjEzsignfolder returns the AObjEzsignfolder field if non-nil, zero value otherwise.

### GetAObjEzsignfolderOk

`func (o *EzsignfolderGetListV1ResponseMPayload) GetAObjEzsignfolderOk() (*[]EzsignfolderListElement, bool)`

GetAObjEzsignfolderOk returns a tuple with the AObjEzsignfolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzsignfolder

`func (o *EzsignfolderGetListV1ResponseMPayload) SetAObjEzsignfolder(v []EzsignfolderListElement)`

SetAObjEzsignfolder sets AObjEzsignfolder field to given value.


### GetIRowReturned

`func (o *EzsignfolderGetListV1ResponseMPayload) GetIRowReturned() int32`

GetIRowReturned returns the IRowReturned field if non-nil, zero value otherwise.

### GetIRowReturnedOk

`func (o *EzsignfolderGetListV1ResponseMPayload) GetIRowReturnedOk() (*int32, bool)`

GetIRowReturnedOk returns a tuple with the IRowReturned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowReturned

`func (o *EzsignfolderGetListV1ResponseMPayload) SetIRowReturned(v int32)`

SetIRowReturned sets IRowReturned field to given value.


### GetIRowFiltered

`func (o *EzsignfolderGetListV1ResponseMPayload) GetIRowFiltered() int32`

GetIRowFiltered returns the IRowFiltered field if non-nil, zero value otherwise.

### GetIRowFilteredOk

`func (o *EzsignfolderGetListV1ResponseMPayload) GetIRowFilteredOk() (*int32, bool)`

GetIRowFilteredOk returns a tuple with the IRowFiltered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowFiltered

`func (o *EzsignfolderGetListV1ResponseMPayload) SetIRowFiltered(v int32)`

SetIRowFiltered sets IRowFiltered field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


