# BrandingGetListV1ResponseMPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IRowReturned** | **int32** | The number of rows returned | 
**IRowFiltered** | **int32** | The number of rows matching your filters (if any) or the total number of rows | 
**AObjBranding** | [**[]BrandingListElement**](BrandingListElement.md) |  | 

## Methods

### NewBrandingGetListV1ResponseMPayload

`func NewBrandingGetListV1ResponseMPayload(iRowReturned int32, iRowFiltered int32, aObjBranding []BrandingListElement, ) *BrandingGetListV1ResponseMPayload`

NewBrandingGetListV1ResponseMPayload instantiates a new BrandingGetListV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrandingGetListV1ResponseMPayloadWithDefaults

`func NewBrandingGetListV1ResponseMPayloadWithDefaults() *BrandingGetListV1ResponseMPayload`

NewBrandingGetListV1ResponseMPayloadWithDefaults instantiates a new BrandingGetListV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIRowReturned

`func (o *BrandingGetListV1ResponseMPayload) GetIRowReturned() int32`

GetIRowReturned returns the IRowReturned field if non-nil, zero value otherwise.

### GetIRowReturnedOk

`func (o *BrandingGetListV1ResponseMPayload) GetIRowReturnedOk() (*int32, bool)`

GetIRowReturnedOk returns a tuple with the IRowReturned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowReturned

`func (o *BrandingGetListV1ResponseMPayload) SetIRowReturned(v int32)`

SetIRowReturned sets IRowReturned field to given value.


### GetIRowFiltered

`func (o *BrandingGetListV1ResponseMPayload) GetIRowFiltered() int32`

GetIRowFiltered returns the IRowFiltered field if non-nil, zero value otherwise.

### GetIRowFilteredOk

`func (o *BrandingGetListV1ResponseMPayload) GetIRowFilteredOk() (*int32, bool)`

GetIRowFilteredOk returns a tuple with the IRowFiltered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowFiltered

`func (o *BrandingGetListV1ResponseMPayload) SetIRowFiltered(v int32)`

SetIRowFiltered sets IRowFiltered field to given value.


### GetAObjBranding

`func (o *BrandingGetListV1ResponseMPayload) GetAObjBranding() []BrandingListElement`

GetAObjBranding returns the AObjBranding field if non-nil, zero value otherwise.

### GetAObjBrandingOk

`func (o *BrandingGetListV1ResponseMPayload) GetAObjBrandingOk() (*[]BrandingListElement, bool)`

GetAObjBrandingOk returns a tuple with the AObjBranding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjBranding

`func (o *BrandingGetListV1ResponseMPayload) SetAObjBranding(v []BrandingListElement)`

SetAObjBranding sets AObjBranding field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


