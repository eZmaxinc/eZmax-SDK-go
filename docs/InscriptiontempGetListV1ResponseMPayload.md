# InscriptiontempGetListV1ResponseMPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IRowReturned** | **int32** | The number of rows returned | 
**IRowFiltered** | **int32** | The number of rows matching your filters (if any) or the total number of rows | 
**AObjInscriptiontemp** | [**[]InscriptiontempListElement**](InscriptiontempListElement.md) |  | 

## Methods

### NewInscriptiontempGetListV1ResponseMPayload

`func NewInscriptiontempGetListV1ResponseMPayload(iRowReturned int32, iRowFiltered int32, aObjInscriptiontemp []InscriptiontempListElement, ) *InscriptiontempGetListV1ResponseMPayload`

NewInscriptiontempGetListV1ResponseMPayload instantiates a new InscriptiontempGetListV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInscriptiontempGetListV1ResponseMPayloadWithDefaults

`func NewInscriptiontempGetListV1ResponseMPayloadWithDefaults() *InscriptiontempGetListV1ResponseMPayload`

NewInscriptiontempGetListV1ResponseMPayloadWithDefaults instantiates a new InscriptiontempGetListV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIRowReturned

`func (o *InscriptiontempGetListV1ResponseMPayload) GetIRowReturned() int32`

GetIRowReturned returns the IRowReturned field if non-nil, zero value otherwise.

### GetIRowReturnedOk

`func (o *InscriptiontempGetListV1ResponseMPayload) GetIRowReturnedOk() (*int32, bool)`

GetIRowReturnedOk returns a tuple with the IRowReturned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowReturned

`func (o *InscriptiontempGetListV1ResponseMPayload) SetIRowReturned(v int32)`

SetIRowReturned sets IRowReturned field to given value.


### GetIRowFiltered

`func (o *InscriptiontempGetListV1ResponseMPayload) GetIRowFiltered() int32`

GetIRowFiltered returns the IRowFiltered field if non-nil, zero value otherwise.

### GetIRowFilteredOk

`func (o *InscriptiontempGetListV1ResponseMPayload) GetIRowFilteredOk() (*int32, bool)`

GetIRowFilteredOk returns a tuple with the IRowFiltered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowFiltered

`func (o *InscriptiontempGetListV1ResponseMPayload) SetIRowFiltered(v int32)`

SetIRowFiltered sets IRowFiltered field to given value.


### GetAObjInscriptiontemp

`func (o *InscriptiontempGetListV1ResponseMPayload) GetAObjInscriptiontemp() []InscriptiontempListElement`

GetAObjInscriptiontemp returns the AObjInscriptiontemp field if non-nil, zero value otherwise.

### GetAObjInscriptiontempOk

`func (o *InscriptiontempGetListV1ResponseMPayload) GetAObjInscriptiontempOk() (*[]InscriptiontempListElement, bool)`

GetAObjInscriptiontempOk returns a tuple with the AObjInscriptiontemp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjInscriptiontemp

`func (o *InscriptiontempGetListV1ResponseMPayload) SetAObjInscriptiontemp(v []InscriptiontempListElement)`

SetAObjInscriptiontemp sets AObjInscriptiontemp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


