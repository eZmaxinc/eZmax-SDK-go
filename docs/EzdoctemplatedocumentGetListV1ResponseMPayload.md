# EzdoctemplatedocumentGetListV1ResponseMPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IRowReturned** | **int32** | The number of rows returned | 
**IRowFiltered** | **int32** | The number of rows matching your filters (if any) or the total number of rows | 
**AObjEzdoctemplatedocument** | [**[]EzdoctemplatedocumentListElement**](EzdoctemplatedocumentListElement.md) |  | 

## Methods

### NewEzdoctemplatedocumentGetListV1ResponseMPayload

`func NewEzdoctemplatedocumentGetListV1ResponseMPayload(iRowReturned int32, iRowFiltered int32, aObjEzdoctemplatedocument []EzdoctemplatedocumentListElement, ) *EzdoctemplatedocumentGetListV1ResponseMPayload`

NewEzdoctemplatedocumentGetListV1ResponseMPayload instantiates a new EzdoctemplatedocumentGetListV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzdoctemplatedocumentGetListV1ResponseMPayloadWithDefaults

`func NewEzdoctemplatedocumentGetListV1ResponseMPayloadWithDefaults() *EzdoctemplatedocumentGetListV1ResponseMPayload`

NewEzdoctemplatedocumentGetListV1ResponseMPayloadWithDefaults instantiates a new EzdoctemplatedocumentGetListV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIRowReturned

`func (o *EzdoctemplatedocumentGetListV1ResponseMPayload) GetIRowReturned() int32`

GetIRowReturned returns the IRowReturned field if non-nil, zero value otherwise.

### GetIRowReturnedOk

`func (o *EzdoctemplatedocumentGetListV1ResponseMPayload) GetIRowReturnedOk() (*int32, bool)`

GetIRowReturnedOk returns a tuple with the IRowReturned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowReturned

`func (o *EzdoctemplatedocumentGetListV1ResponseMPayload) SetIRowReturned(v int32)`

SetIRowReturned sets IRowReturned field to given value.


### GetIRowFiltered

`func (o *EzdoctemplatedocumentGetListV1ResponseMPayload) GetIRowFiltered() int32`

GetIRowFiltered returns the IRowFiltered field if non-nil, zero value otherwise.

### GetIRowFilteredOk

`func (o *EzdoctemplatedocumentGetListV1ResponseMPayload) GetIRowFilteredOk() (*int32, bool)`

GetIRowFilteredOk returns a tuple with the IRowFiltered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowFiltered

`func (o *EzdoctemplatedocumentGetListV1ResponseMPayload) SetIRowFiltered(v int32)`

SetIRowFiltered sets IRowFiltered field to given value.


### GetAObjEzdoctemplatedocument

`func (o *EzdoctemplatedocumentGetListV1ResponseMPayload) GetAObjEzdoctemplatedocument() []EzdoctemplatedocumentListElement`

GetAObjEzdoctemplatedocument returns the AObjEzdoctemplatedocument field if non-nil, zero value otherwise.

### GetAObjEzdoctemplatedocumentOk

`func (o *EzdoctemplatedocumentGetListV1ResponseMPayload) GetAObjEzdoctemplatedocumentOk() (*[]EzdoctemplatedocumentListElement, bool)`

GetAObjEzdoctemplatedocumentOk returns a tuple with the AObjEzdoctemplatedocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzdoctemplatedocument

`func (o *EzdoctemplatedocumentGetListV1ResponseMPayload) SetAObjEzdoctemplatedocument(v []EzdoctemplatedocumentListElement)`

SetAObjEzdoctemplatedocument sets AObjEzdoctemplatedocument field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


