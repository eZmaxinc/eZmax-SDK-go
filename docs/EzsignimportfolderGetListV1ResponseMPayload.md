# EzsignimportfolderGetListV1ResponseMPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IRowReturned** | **int32** | The number of rows returned | 
**IRowFiltered** | **int32** | The number of rows matching your filters (if any) or the total number of rows | 
**AObjEzsignimportfolder** | [**[]EzsignimportfolderListElement**](EzsignimportfolderListElement.md) |  | 

## Methods

### NewEzsignimportfolderGetListV1ResponseMPayload

`func NewEzsignimportfolderGetListV1ResponseMPayload(iRowReturned int32, iRowFiltered int32, aObjEzsignimportfolder []EzsignimportfolderListElement, ) *EzsignimportfolderGetListV1ResponseMPayload`

NewEzsignimportfolderGetListV1ResponseMPayload instantiates a new EzsignimportfolderGetListV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignimportfolderGetListV1ResponseMPayloadWithDefaults

`func NewEzsignimportfolderGetListV1ResponseMPayloadWithDefaults() *EzsignimportfolderGetListV1ResponseMPayload`

NewEzsignimportfolderGetListV1ResponseMPayloadWithDefaults instantiates a new EzsignimportfolderGetListV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIRowReturned

`func (o *EzsignimportfolderGetListV1ResponseMPayload) GetIRowReturned() int32`

GetIRowReturned returns the IRowReturned field if non-nil, zero value otherwise.

### GetIRowReturnedOk

`func (o *EzsignimportfolderGetListV1ResponseMPayload) GetIRowReturnedOk() (*int32, bool)`

GetIRowReturnedOk returns a tuple with the IRowReturned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowReturned

`func (o *EzsignimportfolderGetListV1ResponseMPayload) SetIRowReturned(v int32)`

SetIRowReturned sets IRowReturned field to given value.


### GetIRowFiltered

`func (o *EzsignimportfolderGetListV1ResponseMPayload) GetIRowFiltered() int32`

GetIRowFiltered returns the IRowFiltered field if non-nil, zero value otherwise.

### GetIRowFilteredOk

`func (o *EzsignimportfolderGetListV1ResponseMPayload) GetIRowFilteredOk() (*int32, bool)`

GetIRowFilteredOk returns a tuple with the IRowFiltered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowFiltered

`func (o *EzsignimportfolderGetListV1ResponseMPayload) SetIRowFiltered(v int32)`

SetIRowFiltered sets IRowFiltered field to given value.


### GetAObjEzsignimportfolder

`func (o *EzsignimportfolderGetListV1ResponseMPayload) GetAObjEzsignimportfolder() []EzsignimportfolderListElement`

GetAObjEzsignimportfolder returns the AObjEzsignimportfolder field if non-nil, zero value otherwise.

### GetAObjEzsignimportfolderOk

`func (o *EzsignimportfolderGetListV1ResponseMPayload) GetAObjEzsignimportfolderOk() (*[]EzsignimportfolderListElement, bool)`

GetAObjEzsignimportfolderOk returns a tuple with the AObjEzsignimportfolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzsignimportfolder

`func (o *EzsignimportfolderGetListV1ResponseMPayload) SetAObjEzsignimportfolder(v []EzsignimportfolderListElement)`

SetAObjEzsignimportfolder sets AObjEzsignimportfolder field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


