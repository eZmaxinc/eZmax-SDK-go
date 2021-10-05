# EzsignfoldertypeGetListV1ResponseMPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AObjEzsignfoldertype** | [**[]EzsignfoldertypeListElement**](EzsignfoldertypeListElement.md) |  | 
**IRowReturned** | **int32** | The number of rows returned | 
**IRowFiltered** | **int32** | The number of rows matching your filters (if any) or the total number of rows | 

## Methods

### NewEzsignfoldertypeGetListV1ResponseMPayload

`func NewEzsignfoldertypeGetListV1ResponseMPayload(aObjEzsignfoldertype []EzsignfoldertypeListElement, iRowReturned int32, iRowFiltered int32, ) *EzsignfoldertypeGetListV1ResponseMPayload`

NewEzsignfoldertypeGetListV1ResponseMPayload instantiates a new EzsignfoldertypeGetListV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignfoldertypeGetListV1ResponseMPayloadWithDefaults

`func NewEzsignfoldertypeGetListV1ResponseMPayloadWithDefaults() *EzsignfoldertypeGetListV1ResponseMPayload`

NewEzsignfoldertypeGetListV1ResponseMPayloadWithDefaults instantiates a new EzsignfoldertypeGetListV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAObjEzsignfoldertype

`func (o *EzsignfoldertypeGetListV1ResponseMPayload) GetAObjEzsignfoldertype() []EzsignfoldertypeListElement`

GetAObjEzsignfoldertype returns the AObjEzsignfoldertype field if non-nil, zero value otherwise.

### GetAObjEzsignfoldertypeOk

`func (o *EzsignfoldertypeGetListV1ResponseMPayload) GetAObjEzsignfoldertypeOk() (*[]EzsignfoldertypeListElement, bool)`

GetAObjEzsignfoldertypeOk returns a tuple with the AObjEzsignfoldertype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzsignfoldertype

`func (o *EzsignfoldertypeGetListV1ResponseMPayload) SetAObjEzsignfoldertype(v []EzsignfoldertypeListElement)`

SetAObjEzsignfoldertype sets AObjEzsignfoldertype field to given value.


### GetIRowReturned

`func (o *EzsignfoldertypeGetListV1ResponseMPayload) GetIRowReturned() int32`

GetIRowReturned returns the IRowReturned field if non-nil, zero value otherwise.

### GetIRowReturnedOk

`func (o *EzsignfoldertypeGetListV1ResponseMPayload) GetIRowReturnedOk() (*int32, bool)`

GetIRowReturnedOk returns a tuple with the IRowReturned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowReturned

`func (o *EzsignfoldertypeGetListV1ResponseMPayload) SetIRowReturned(v int32)`

SetIRowReturned sets IRowReturned field to given value.


### GetIRowFiltered

`func (o *EzsignfoldertypeGetListV1ResponseMPayload) GetIRowFiltered() int32`

GetIRowFiltered returns the IRowFiltered field if non-nil, zero value otherwise.

### GetIRowFilteredOk

`func (o *EzsignfoldertypeGetListV1ResponseMPayload) GetIRowFilteredOk() (*int32, bool)`

GetIRowFilteredOk returns a tuple with the IRowFiltered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowFiltered

`func (o *EzsignfoldertypeGetListV1ResponseMPayload) SetIRowFiltered(v int32)`

SetIRowFiltered sets IRowFiltered field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


