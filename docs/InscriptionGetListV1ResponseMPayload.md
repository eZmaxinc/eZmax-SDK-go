# InscriptionGetListV1ResponseMPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IRowReturned** | **int32** | The number of rows returned | 
**IRowFiltered** | **int32** | The number of rows matching your filters (if any) or the total number of rows | 
**AObjInscription** | [**[]InscriptionListElement**](InscriptionListElement.md) |  | 

## Methods

### NewInscriptionGetListV1ResponseMPayload

`func NewInscriptionGetListV1ResponseMPayload(iRowReturned int32, iRowFiltered int32, aObjInscription []InscriptionListElement, ) *InscriptionGetListV1ResponseMPayload`

NewInscriptionGetListV1ResponseMPayload instantiates a new InscriptionGetListV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInscriptionGetListV1ResponseMPayloadWithDefaults

`func NewInscriptionGetListV1ResponseMPayloadWithDefaults() *InscriptionGetListV1ResponseMPayload`

NewInscriptionGetListV1ResponseMPayloadWithDefaults instantiates a new InscriptionGetListV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIRowReturned

`func (o *InscriptionGetListV1ResponseMPayload) GetIRowReturned() int32`

GetIRowReturned returns the IRowReturned field if non-nil, zero value otherwise.

### GetIRowReturnedOk

`func (o *InscriptionGetListV1ResponseMPayload) GetIRowReturnedOk() (*int32, bool)`

GetIRowReturnedOk returns a tuple with the IRowReturned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowReturned

`func (o *InscriptionGetListV1ResponseMPayload) SetIRowReturned(v int32)`

SetIRowReturned sets IRowReturned field to given value.


### GetIRowFiltered

`func (o *InscriptionGetListV1ResponseMPayload) GetIRowFiltered() int32`

GetIRowFiltered returns the IRowFiltered field if non-nil, zero value otherwise.

### GetIRowFilteredOk

`func (o *InscriptionGetListV1ResponseMPayload) GetIRowFilteredOk() (*int32, bool)`

GetIRowFilteredOk returns a tuple with the IRowFiltered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowFiltered

`func (o *InscriptionGetListV1ResponseMPayload) SetIRowFiltered(v int32)`

SetIRowFiltered sets IRowFiltered field to given value.


### GetAObjInscription

`func (o *InscriptionGetListV1ResponseMPayload) GetAObjInscription() []InscriptionListElement`

GetAObjInscription returns the AObjInscription field if non-nil, zero value otherwise.

### GetAObjInscriptionOk

`func (o *InscriptionGetListV1ResponseMPayload) GetAObjInscriptionOk() (*[]InscriptionListElement, bool)`

GetAObjInscriptionOk returns a tuple with the AObjInscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjInscription

`func (o *InscriptionGetListV1ResponseMPayload) SetAObjInscription(v []InscriptionListElement)`

SetAObjInscription sets AObjInscription field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


