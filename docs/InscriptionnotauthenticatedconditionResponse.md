# InscriptionnotauthenticatedconditionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiInscriptionnotauthenticatedconditionID** | **int32** | The unique ID of the Inscriptionnotauthenticatedcondition | 
**FkiInscriptionnotauthenticatedconditiontypeID** | **int32** | The unique ID of the Inscriptionnotauthenticatedconditiontype | 
**SInscriptionnotauthenticatedconditiontypeNameX** | **string** | The name of the Inscriptionnotauthenticatedconditiontype in the language of the requester | 
**FkiInscriptionnotauthenticatedID** | **int32** | The unique ID of the Inscriptionnotauthenticated. | 
**BInscriptionnotauthenticatedconditionFilled** | **bool** | Can access attachment when we clone a user | 
**DtInscriptionnotauthenticatedconditionCompleted** | Pointer to **string** | The date the Inscriptionnotauthenticatedcondition was completed | [optional] 
**DtInscriptionnotauthenticatedconditionDue** | Pointer to **string** | The date the Inscriptionnotauthenticatedcondition is due | [optional] 
**TInscriptionnotauthenticatedconditionComment** | **string** | The comment of the Inscriptionnotauthenticatedcondition | 

## Methods

### NewInscriptionnotauthenticatedconditionResponse

`func NewInscriptionnotauthenticatedconditionResponse(pkiInscriptionnotauthenticatedconditionID int32, fkiInscriptionnotauthenticatedconditiontypeID int32, sInscriptionnotauthenticatedconditiontypeNameX string, fkiInscriptionnotauthenticatedID int32, bInscriptionnotauthenticatedconditionFilled bool, tInscriptionnotauthenticatedconditionComment string, ) *InscriptionnotauthenticatedconditionResponse`

NewInscriptionnotauthenticatedconditionResponse instantiates a new InscriptionnotauthenticatedconditionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInscriptionnotauthenticatedconditionResponseWithDefaults

`func NewInscriptionnotauthenticatedconditionResponseWithDefaults() *InscriptionnotauthenticatedconditionResponse`

NewInscriptionnotauthenticatedconditionResponseWithDefaults instantiates a new InscriptionnotauthenticatedconditionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiInscriptionnotauthenticatedconditionID

`func (o *InscriptionnotauthenticatedconditionResponse) GetPkiInscriptionnotauthenticatedconditionID() int32`

GetPkiInscriptionnotauthenticatedconditionID returns the PkiInscriptionnotauthenticatedconditionID field if non-nil, zero value otherwise.

### GetPkiInscriptionnotauthenticatedconditionIDOk

`func (o *InscriptionnotauthenticatedconditionResponse) GetPkiInscriptionnotauthenticatedconditionIDOk() (*int32, bool)`

GetPkiInscriptionnotauthenticatedconditionIDOk returns a tuple with the PkiInscriptionnotauthenticatedconditionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiInscriptionnotauthenticatedconditionID

`func (o *InscriptionnotauthenticatedconditionResponse) SetPkiInscriptionnotauthenticatedconditionID(v int32)`

SetPkiInscriptionnotauthenticatedconditionID sets PkiInscriptionnotauthenticatedconditionID field to given value.


### GetFkiInscriptionnotauthenticatedconditiontypeID

`func (o *InscriptionnotauthenticatedconditionResponse) GetFkiInscriptionnotauthenticatedconditiontypeID() int32`

GetFkiInscriptionnotauthenticatedconditiontypeID returns the FkiInscriptionnotauthenticatedconditiontypeID field if non-nil, zero value otherwise.

### GetFkiInscriptionnotauthenticatedconditiontypeIDOk

`func (o *InscriptionnotauthenticatedconditionResponse) GetFkiInscriptionnotauthenticatedconditiontypeIDOk() (*int32, bool)`

GetFkiInscriptionnotauthenticatedconditiontypeIDOk returns a tuple with the FkiInscriptionnotauthenticatedconditiontypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiInscriptionnotauthenticatedconditiontypeID

`func (o *InscriptionnotauthenticatedconditionResponse) SetFkiInscriptionnotauthenticatedconditiontypeID(v int32)`

SetFkiInscriptionnotauthenticatedconditiontypeID sets FkiInscriptionnotauthenticatedconditiontypeID field to given value.


### GetSInscriptionnotauthenticatedconditiontypeNameX

`func (o *InscriptionnotauthenticatedconditionResponse) GetSInscriptionnotauthenticatedconditiontypeNameX() string`

GetSInscriptionnotauthenticatedconditiontypeNameX returns the SInscriptionnotauthenticatedconditiontypeNameX field if non-nil, zero value otherwise.

### GetSInscriptionnotauthenticatedconditiontypeNameXOk

`func (o *InscriptionnotauthenticatedconditionResponse) GetSInscriptionnotauthenticatedconditiontypeNameXOk() (*string, bool)`

GetSInscriptionnotauthenticatedconditiontypeNameXOk returns a tuple with the SInscriptionnotauthenticatedconditiontypeNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSInscriptionnotauthenticatedconditiontypeNameX

`func (o *InscriptionnotauthenticatedconditionResponse) SetSInscriptionnotauthenticatedconditiontypeNameX(v string)`

SetSInscriptionnotauthenticatedconditiontypeNameX sets SInscriptionnotauthenticatedconditiontypeNameX field to given value.


### GetFkiInscriptionnotauthenticatedID

`func (o *InscriptionnotauthenticatedconditionResponse) GetFkiInscriptionnotauthenticatedID() int32`

GetFkiInscriptionnotauthenticatedID returns the FkiInscriptionnotauthenticatedID field if non-nil, zero value otherwise.

### GetFkiInscriptionnotauthenticatedIDOk

`func (o *InscriptionnotauthenticatedconditionResponse) GetFkiInscriptionnotauthenticatedIDOk() (*int32, bool)`

GetFkiInscriptionnotauthenticatedIDOk returns a tuple with the FkiInscriptionnotauthenticatedID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiInscriptionnotauthenticatedID

`func (o *InscriptionnotauthenticatedconditionResponse) SetFkiInscriptionnotauthenticatedID(v int32)`

SetFkiInscriptionnotauthenticatedID sets FkiInscriptionnotauthenticatedID field to given value.


### GetBInscriptionnotauthenticatedconditionFilled

`func (o *InscriptionnotauthenticatedconditionResponse) GetBInscriptionnotauthenticatedconditionFilled() bool`

GetBInscriptionnotauthenticatedconditionFilled returns the BInscriptionnotauthenticatedconditionFilled field if non-nil, zero value otherwise.

### GetBInscriptionnotauthenticatedconditionFilledOk

`func (o *InscriptionnotauthenticatedconditionResponse) GetBInscriptionnotauthenticatedconditionFilledOk() (*bool, bool)`

GetBInscriptionnotauthenticatedconditionFilledOk returns a tuple with the BInscriptionnotauthenticatedconditionFilled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBInscriptionnotauthenticatedconditionFilled

`func (o *InscriptionnotauthenticatedconditionResponse) SetBInscriptionnotauthenticatedconditionFilled(v bool)`

SetBInscriptionnotauthenticatedconditionFilled sets BInscriptionnotauthenticatedconditionFilled field to given value.


### GetDtInscriptionnotauthenticatedconditionCompleted

`func (o *InscriptionnotauthenticatedconditionResponse) GetDtInscriptionnotauthenticatedconditionCompleted() string`

GetDtInscriptionnotauthenticatedconditionCompleted returns the DtInscriptionnotauthenticatedconditionCompleted field if non-nil, zero value otherwise.

### GetDtInscriptionnotauthenticatedconditionCompletedOk

`func (o *InscriptionnotauthenticatedconditionResponse) GetDtInscriptionnotauthenticatedconditionCompletedOk() (*string, bool)`

GetDtInscriptionnotauthenticatedconditionCompletedOk returns a tuple with the DtInscriptionnotauthenticatedconditionCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtInscriptionnotauthenticatedconditionCompleted

`func (o *InscriptionnotauthenticatedconditionResponse) SetDtInscriptionnotauthenticatedconditionCompleted(v string)`

SetDtInscriptionnotauthenticatedconditionCompleted sets DtInscriptionnotauthenticatedconditionCompleted field to given value.

### HasDtInscriptionnotauthenticatedconditionCompleted

`func (o *InscriptionnotauthenticatedconditionResponse) HasDtInscriptionnotauthenticatedconditionCompleted() bool`

HasDtInscriptionnotauthenticatedconditionCompleted returns a boolean if a field has been set.

### GetDtInscriptionnotauthenticatedconditionDue

`func (o *InscriptionnotauthenticatedconditionResponse) GetDtInscriptionnotauthenticatedconditionDue() string`

GetDtInscriptionnotauthenticatedconditionDue returns the DtInscriptionnotauthenticatedconditionDue field if non-nil, zero value otherwise.

### GetDtInscriptionnotauthenticatedconditionDueOk

`func (o *InscriptionnotauthenticatedconditionResponse) GetDtInscriptionnotauthenticatedconditionDueOk() (*string, bool)`

GetDtInscriptionnotauthenticatedconditionDueOk returns a tuple with the DtInscriptionnotauthenticatedconditionDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtInscriptionnotauthenticatedconditionDue

`func (o *InscriptionnotauthenticatedconditionResponse) SetDtInscriptionnotauthenticatedconditionDue(v string)`

SetDtInscriptionnotauthenticatedconditionDue sets DtInscriptionnotauthenticatedconditionDue field to given value.

### HasDtInscriptionnotauthenticatedconditionDue

`func (o *InscriptionnotauthenticatedconditionResponse) HasDtInscriptionnotauthenticatedconditionDue() bool`

HasDtInscriptionnotauthenticatedconditionDue returns a boolean if a field has been set.

### GetTInscriptionnotauthenticatedconditionComment

`func (o *InscriptionnotauthenticatedconditionResponse) GetTInscriptionnotauthenticatedconditionComment() string`

GetTInscriptionnotauthenticatedconditionComment returns the TInscriptionnotauthenticatedconditionComment field if non-nil, zero value otherwise.

### GetTInscriptionnotauthenticatedconditionCommentOk

`func (o *InscriptionnotauthenticatedconditionResponse) GetTInscriptionnotauthenticatedconditionCommentOk() (*string, bool)`

GetTInscriptionnotauthenticatedconditionCommentOk returns a tuple with the TInscriptionnotauthenticatedconditionComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTInscriptionnotauthenticatedconditionComment

`func (o *InscriptionnotauthenticatedconditionResponse) SetTInscriptionnotauthenticatedconditionComment(v string)`

SetTInscriptionnotauthenticatedconditionComment sets TInscriptionnotauthenticatedconditionComment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


