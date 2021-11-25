# EzsigntemplatepackageGetListV1ResponseMPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AObjEzsigntemplatepackage** | [**[]EzsigntemplatepackageListElement**](EzsigntemplatepackageListElement.md) |  | 
**IRowReturned** | **int32** | The number of rows returned | 
**IRowFiltered** | **int32** | The number of rows matching your filters (if any) or the total number of rows | 

## Methods

### NewEzsigntemplatepackageGetListV1ResponseMPayload

`func NewEzsigntemplatepackageGetListV1ResponseMPayload(aObjEzsigntemplatepackage []EzsigntemplatepackageListElement, iRowReturned int32, iRowFiltered int32, ) *EzsigntemplatepackageGetListV1ResponseMPayload`

NewEzsigntemplatepackageGetListV1ResponseMPayload instantiates a new EzsigntemplatepackageGetListV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigntemplatepackageGetListV1ResponseMPayloadWithDefaults

`func NewEzsigntemplatepackageGetListV1ResponseMPayloadWithDefaults() *EzsigntemplatepackageGetListV1ResponseMPayload`

NewEzsigntemplatepackageGetListV1ResponseMPayloadWithDefaults instantiates a new EzsigntemplatepackageGetListV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAObjEzsigntemplatepackage

`func (o *EzsigntemplatepackageGetListV1ResponseMPayload) GetAObjEzsigntemplatepackage() []EzsigntemplatepackageListElement`

GetAObjEzsigntemplatepackage returns the AObjEzsigntemplatepackage field if non-nil, zero value otherwise.

### GetAObjEzsigntemplatepackageOk

`func (o *EzsigntemplatepackageGetListV1ResponseMPayload) GetAObjEzsigntemplatepackageOk() (*[]EzsigntemplatepackageListElement, bool)`

GetAObjEzsigntemplatepackageOk returns a tuple with the AObjEzsigntemplatepackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzsigntemplatepackage

`func (o *EzsigntemplatepackageGetListV1ResponseMPayload) SetAObjEzsigntemplatepackage(v []EzsigntemplatepackageListElement)`

SetAObjEzsigntemplatepackage sets AObjEzsigntemplatepackage field to given value.


### GetIRowReturned

`func (o *EzsigntemplatepackageGetListV1ResponseMPayload) GetIRowReturned() int32`

GetIRowReturned returns the IRowReturned field if non-nil, zero value otherwise.

### GetIRowReturnedOk

`func (o *EzsigntemplatepackageGetListV1ResponseMPayload) GetIRowReturnedOk() (*int32, bool)`

GetIRowReturnedOk returns a tuple with the IRowReturned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowReturned

`func (o *EzsigntemplatepackageGetListV1ResponseMPayload) SetIRowReturned(v int32)`

SetIRowReturned sets IRowReturned field to given value.


### GetIRowFiltered

`func (o *EzsigntemplatepackageGetListV1ResponseMPayload) GetIRowFiltered() int32`

GetIRowFiltered returns the IRowFiltered field if non-nil, zero value otherwise.

### GetIRowFilteredOk

`func (o *EzsigntemplatepackageGetListV1ResponseMPayload) GetIRowFilteredOk() (*int32, bool)`

GetIRowFilteredOk returns a tuple with the IRowFiltered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowFiltered

`func (o *EzsigntemplatepackageGetListV1ResponseMPayload) SetIRowFiltered(v int32)`

SetIRowFiltered sets IRowFiltered field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


