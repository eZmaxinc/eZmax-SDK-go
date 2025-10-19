# OtherincomeGetListV1ResponseMPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IRowReturned** | **int32** | The number of rows returned | 
**IRowFiltered** | **int32** | The number of rows matching your filters (if any) or the total number of rows | 
**AObjOtherincome** | [**[]OtherincomeListElement**](OtherincomeListElement.md) |  | 

## Methods

### NewOtherincomeGetListV1ResponseMPayload

`func NewOtherincomeGetListV1ResponseMPayload(iRowReturned int32, iRowFiltered int32, aObjOtherincome []OtherincomeListElement, ) *OtherincomeGetListV1ResponseMPayload`

NewOtherincomeGetListV1ResponseMPayload instantiates a new OtherincomeGetListV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOtherincomeGetListV1ResponseMPayloadWithDefaults

`func NewOtherincomeGetListV1ResponseMPayloadWithDefaults() *OtherincomeGetListV1ResponseMPayload`

NewOtherincomeGetListV1ResponseMPayloadWithDefaults instantiates a new OtherincomeGetListV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIRowReturned

`func (o *OtherincomeGetListV1ResponseMPayload) GetIRowReturned() int32`

GetIRowReturned returns the IRowReturned field if non-nil, zero value otherwise.

### GetIRowReturnedOk

`func (o *OtherincomeGetListV1ResponseMPayload) GetIRowReturnedOk() (*int32, bool)`

GetIRowReturnedOk returns a tuple with the IRowReturned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowReturned

`func (o *OtherincomeGetListV1ResponseMPayload) SetIRowReturned(v int32)`

SetIRowReturned sets IRowReturned field to given value.


### GetIRowFiltered

`func (o *OtherincomeGetListV1ResponseMPayload) GetIRowFiltered() int32`

GetIRowFiltered returns the IRowFiltered field if non-nil, zero value otherwise.

### GetIRowFilteredOk

`func (o *OtherincomeGetListV1ResponseMPayload) GetIRowFilteredOk() (*int32, bool)`

GetIRowFilteredOk returns a tuple with the IRowFiltered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowFiltered

`func (o *OtherincomeGetListV1ResponseMPayload) SetIRowFiltered(v int32)`

SetIRowFiltered sets IRowFiltered field to given value.


### GetAObjOtherincome

`func (o *OtherincomeGetListV1ResponseMPayload) GetAObjOtherincome() []OtherincomeListElement`

GetAObjOtherincome returns the AObjOtherincome field if non-nil, zero value otherwise.

### GetAObjOtherincomeOk

`func (o *OtherincomeGetListV1ResponseMPayload) GetAObjOtherincomeOk() (*[]OtherincomeListElement, bool)`

GetAObjOtherincomeOk returns a tuple with the AObjOtherincome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjOtherincome

`func (o *OtherincomeGetListV1ResponseMPayload) SetAObjOtherincome(v []OtherincomeListElement)`

SetAObjOtherincome sets AObjOtherincome field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


