# InscriptiontempListElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiInscriptiontempID** | **int32** | The unique ID of the Inscriptiontemp | 
**EInscriptiontempStatus** | [**FieldEInscriptiontempStatus**](FieldEInscriptiontempStatus.md) |  | 
**SInscriptiontempMLS** | Pointer to **string** | The mls of the Inscriptiontemp | [optional] 
**SInscriptiontempDescription** | **string** | The description of the Inscriptiontemp | 
**BInscriptiontempIsactive** | **bool** | Whether the inscriptiontemp is active or not | 
**DtCreatedDate** | **string** | The date and time at which the object was created | 
**DtModifiedDate** | **string** | The date and time at which the object was last modified | 

## Methods

### NewInscriptiontempListElement

`func NewInscriptiontempListElement(pkiInscriptiontempID int32, eInscriptiontempStatus FieldEInscriptiontempStatus, sInscriptiontempDescription string, bInscriptiontempIsactive bool, dtCreatedDate string, dtModifiedDate string, ) *InscriptiontempListElement`

NewInscriptiontempListElement instantiates a new InscriptiontempListElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInscriptiontempListElementWithDefaults

`func NewInscriptiontempListElementWithDefaults() *InscriptiontempListElement`

NewInscriptiontempListElementWithDefaults instantiates a new InscriptiontempListElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiInscriptiontempID

`func (o *InscriptiontempListElement) GetPkiInscriptiontempID() int32`

GetPkiInscriptiontempID returns the PkiInscriptiontempID field if non-nil, zero value otherwise.

### GetPkiInscriptiontempIDOk

`func (o *InscriptiontempListElement) GetPkiInscriptiontempIDOk() (*int32, bool)`

GetPkiInscriptiontempIDOk returns a tuple with the PkiInscriptiontempID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiInscriptiontempID

`func (o *InscriptiontempListElement) SetPkiInscriptiontempID(v int32)`

SetPkiInscriptiontempID sets PkiInscriptiontempID field to given value.


### GetEInscriptiontempStatus

`func (o *InscriptiontempListElement) GetEInscriptiontempStatus() FieldEInscriptiontempStatus`

GetEInscriptiontempStatus returns the EInscriptiontempStatus field if non-nil, zero value otherwise.

### GetEInscriptiontempStatusOk

`func (o *InscriptiontempListElement) GetEInscriptiontempStatusOk() (*FieldEInscriptiontempStatus, bool)`

GetEInscriptiontempStatusOk returns a tuple with the EInscriptiontempStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEInscriptiontempStatus

`func (o *InscriptiontempListElement) SetEInscriptiontempStatus(v FieldEInscriptiontempStatus)`

SetEInscriptiontempStatus sets EInscriptiontempStatus field to given value.


### GetSInscriptiontempMLS

`func (o *InscriptiontempListElement) GetSInscriptiontempMLS() string`

GetSInscriptiontempMLS returns the SInscriptiontempMLS field if non-nil, zero value otherwise.

### GetSInscriptiontempMLSOk

`func (o *InscriptiontempListElement) GetSInscriptiontempMLSOk() (*string, bool)`

GetSInscriptiontempMLSOk returns a tuple with the SInscriptiontempMLS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSInscriptiontempMLS

`func (o *InscriptiontempListElement) SetSInscriptiontempMLS(v string)`

SetSInscriptiontempMLS sets SInscriptiontempMLS field to given value.

### HasSInscriptiontempMLS

`func (o *InscriptiontempListElement) HasSInscriptiontempMLS() bool`

HasSInscriptiontempMLS returns a boolean if a field has been set.

### GetSInscriptiontempDescription

`func (o *InscriptiontempListElement) GetSInscriptiontempDescription() string`

GetSInscriptiontempDescription returns the SInscriptiontempDescription field if non-nil, zero value otherwise.

### GetSInscriptiontempDescriptionOk

`func (o *InscriptiontempListElement) GetSInscriptiontempDescriptionOk() (*string, bool)`

GetSInscriptiontempDescriptionOk returns a tuple with the SInscriptiontempDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSInscriptiontempDescription

`func (o *InscriptiontempListElement) SetSInscriptiontempDescription(v string)`

SetSInscriptiontempDescription sets SInscriptiontempDescription field to given value.


### GetBInscriptiontempIsactive

`func (o *InscriptiontempListElement) GetBInscriptiontempIsactive() bool`

GetBInscriptiontempIsactive returns the BInscriptiontempIsactive field if non-nil, zero value otherwise.

### GetBInscriptiontempIsactiveOk

`func (o *InscriptiontempListElement) GetBInscriptiontempIsactiveOk() (*bool, bool)`

GetBInscriptiontempIsactiveOk returns a tuple with the BInscriptiontempIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBInscriptiontempIsactive

`func (o *InscriptiontempListElement) SetBInscriptiontempIsactive(v bool)`

SetBInscriptiontempIsactive sets BInscriptiontempIsactive field to given value.


### GetDtCreatedDate

`func (o *InscriptiontempListElement) GetDtCreatedDate() string`

GetDtCreatedDate returns the DtCreatedDate field if non-nil, zero value otherwise.

### GetDtCreatedDateOk

`func (o *InscriptiontempListElement) GetDtCreatedDateOk() (*string, bool)`

GetDtCreatedDateOk returns a tuple with the DtCreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtCreatedDate

`func (o *InscriptiontempListElement) SetDtCreatedDate(v string)`

SetDtCreatedDate sets DtCreatedDate field to given value.


### GetDtModifiedDate

`func (o *InscriptiontempListElement) GetDtModifiedDate() string`

GetDtModifiedDate returns the DtModifiedDate field if non-nil, zero value otherwise.

### GetDtModifiedDateOk

`func (o *InscriptiontempListElement) GetDtModifiedDateOk() (*string, bool)`

GetDtModifiedDateOk returns a tuple with the DtModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtModifiedDate

`func (o *InscriptiontempListElement) SetDtModifiedDate(v string)`

SetDtModifiedDate sets DtModifiedDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


