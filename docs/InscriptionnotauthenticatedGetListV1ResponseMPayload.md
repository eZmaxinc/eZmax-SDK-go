# InscriptionnotauthenticatedGetListV1ResponseMPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IRowReturned** | **int32** | The number of rows returned | 
**IRowFiltered** | **int32** | The number of rows matching your filters (if any) or the total number of rows | 
**AObjInscriptionnotauthenticated** | [**[]InscriptionnotauthenticatedListElement**](InscriptionnotauthenticatedListElement.md) |  | 

## Methods

### NewInscriptionnotauthenticatedGetListV1ResponseMPayload

`func NewInscriptionnotauthenticatedGetListV1ResponseMPayload(iRowReturned int32, iRowFiltered int32, aObjInscriptionnotauthenticated []InscriptionnotauthenticatedListElement, ) *InscriptionnotauthenticatedGetListV1ResponseMPayload`

NewInscriptionnotauthenticatedGetListV1ResponseMPayload instantiates a new InscriptionnotauthenticatedGetListV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInscriptionnotauthenticatedGetListV1ResponseMPayloadWithDefaults

`func NewInscriptionnotauthenticatedGetListV1ResponseMPayloadWithDefaults() *InscriptionnotauthenticatedGetListV1ResponseMPayload`

NewInscriptionnotauthenticatedGetListV1ResponseMPayloadWithDefaults instantiates a new InscriptionnotauthenticatedGetListV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIRowReturned

`func (o *InscriptionnotauthenticatedGetListV1ResponseMPayload) GetIRowReturned() int32`

GetIRowReturned returns the IRowReturned field if non-nil, zero value otherwise.

### GetIRowReturnedOk

`func (o *InscriptionnotauthenticatedGetListV1ResponseMPayload) GetIRowReturnedOk() (*int32, bool)`

GetIRowReturnedOk returns a tuple with the IRowReturned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowReturned

`func (o *InscriptionnotauthenticatedGetListV1ResponseMPayload) SetIRowReturned(v int32)`

SetIRowReturned sets IRowReturned field to given value.


### GetIRowFiltered

`func (o *InscriptionnotauthenticatedGetListV1ResponseMPayload) GetIRowFiltered() int32`

GetIRowFiltered returns the IRowFiltered field if non-nil, zero value otherwise.

### GetIRowFilteredOk

`func (o *InscriptionnotauthenticatedGetListV1ResponseMPayload) GetIRowFilteredOk() (*int32, bool)`

GetIRowFilteredOk returns a tuple with the IRowFiltered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowFiltered

`func (o *InscriptionnotauthenticatedGetListV1ResponseMPayload) SetIRowFiltered(v int32)`

SetIRowFiltered sets IRowFiltered field to given value.


### GetAObjInscriptionnotauthenticated

`func (o *InscriptionnotauthenticatedGetListV1ResponseMPayload) GetAObjInscriptionnotauthenticated() []InscriptionnotauthenticatedListElement`

GetAObjInscriptionnotauthenticated returns the AObjInscriptionnotauthenticated field if non-nil, zero value otherwise.

### GetAObjInscriptionnotauthenticatedOk

`func (o *InscriptionnotauthenticatedGetListV1ResponseMPayload) GetAObjInscriptionnotauthenticatedOk() (*[]InscriptionnotauthenticatedListElement, bool)`

GetAObjInscriptionnotauthenticatedOk returns a tuple with the AObjInscriptionnotauthenticated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjInscriptionnotauthenticated

`func (o *InscriptionnotauthenticatedGetListV1ResponseMPayload) SetAObjInscriptionnotauthenticated(v []InscriptionnotauthenticatedListElement)`

SetAObjInscriptionnotauthenticated sets AObjInscriptionnotauthenticated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


