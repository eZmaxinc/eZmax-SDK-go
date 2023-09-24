# EzsigntemplateGetListV1ResponseMPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IRowReturned** | **int32** | The number of rows returned | 
**IRowFiltered** | **int32** | The number of rows matching your filters (if any) or the total number of rows | 
**AObjEzsigntemplate** | [**[]EzsigntemplateListElement**](EzsigntemplateListElement.md) |  | 

## Methods

### NewEzsigntemplateGetListV1ResponseMPayload

`func NewEzsigntemplateGetListV1ResponseMPayload(iRowReturned int32, iRowFiltered int32, aObjEzsigntemplate []EzsigntemplateListElement, ) *EzsigntemplateGetListV1ResponseMPayload`

NewEzsigntemplateGetListV1ResponseMPayload instantiates a new EzsigntemplateGetListV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigntemplateGetListV1ResponseMPayloadWithDefaults

`func NewEzsigntemplateGetListV1ResponseMPayloadWithDefaults() *EzsigntemplateGetListV1ResponseMPayload`

NewEzsigntemplateGetListV1ResponseMPayloadWithDefaults instantiates a new EzsigntemplateGetListV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIRowReturned

`func (o *EzsigntemplateGetListV1ResponseMPayload) GetIRowReturned() int32`

GetIRowReturned returns the IRowReturned field if non-nil, zero value otherwise.

### GetIRowReturnedOk

`func (o *EzsigntemplateGetListV1ResponseMPayload) GetIRowReturnedOk() (*int32, bool)`

GetIRowReturnedOk returns a tuple with the IRowReturned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowReturned

`func (o *EzsigntemplateGetListV1ResponseMPayload) SetIRowReturned(v int32)`

SetIRowReturned sets IRowReturned field to given value.


### GetIRowFiltered

`func (o *EzsigntemplateGetListV1ResponseMPayload) GetIRowFiltered() int32`

GetIRowFiltered returns the IRowFiltered field if non-nil, zero value otherwise.

### GetIRowFilteredOk

`func (o *EzsigntemplateGetListV1ResponseMPayload) GetIRowFilteredOk() (*int32, bool)`

GetIRowFilteredOk returns a tuple with the IRowFiltered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowFiltered

`func (o *EzsigntemplateGetListV1ResponseMPayload) SetIRowFiltered(v int32)`

SetIRowFiltered sets IRowFiltered field to given value.


### GetAObjEzsigntemplate

`func (o *EzsigntemplateGetListV1ResponseMPayload) GetAObjEzsigntemplate() []EzsigntemplateListElement`

GetAObjEzsigntemplate returns the AObjEzsigntemplate field if non-nil, zero value otherwise.

### GetAObjEzsigntemplateOk

`func (o *EzsigntemplateGetListV1ResponseMPayload) GetAObjEzsigntemplateOk() (*[]EzsigntemplateListElement, bool)`

GetAObjEzsigntemplateOk returns a tuple with the AObjEzsigntemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzsigntemplate

`func (o *EzsigntemplateGetListV1ResponseMPayload) SetAObjEzsigntemplate(v []EzsigntemplateListElement)`

SetAObjEzsigntemplate sets AObjEzsigntemplate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


