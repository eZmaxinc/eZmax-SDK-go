# LeadListElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiLeadID** | **int32** | The unique ID of the Lead | 
**FkiLeadsourceID** | **int32** | The unique ID of the Leadsource | 
**SLeadsourceNameX** | **string** | The name of the Leadsource in the language of the requester | 
**ELeadStatus** | [**FieldELeadStatus**](FieldELeadStatus.md) |  | 
**DtLeadExpiration** | **string** | The expiration of the Lead | 
**BLeadIsactive** | **bool** | Whether the lead is active or not | 
**SLeadCode** | **string** | The code of the Lead | 
**SLeadContacts** | Pointer to **string** | The contacts&#39; name of the Lead | [optional] 

## Methods

### NewLeadListElement

`func NewLeadListElement(pkiLeadID int32, fkiLeadsourceID int32, sLeadsourceNameX string, eLeadStatus FieldELeadStatus, dtLeadExpiration string, bLeadIsactive bool, sLeadCode string, ) *LeadListElement`

NewLeadListElement instantiates a new LeadListElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeadListElementWithDefaults

`func NewLeadListElementWithDefaults() *LeadListElement`

NewLeadListElementWithDefaults instantiates a new LeadListElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiLeadID

`func (o *LeadListElement) GetPkiLeadID() int32`

GetPkiLeadID returns the PkiLeadID field if non-nil, zero value otherwise.

### GetPkiLeadIDOk

`func (o *LeadListElement) GetPkiLeadIDOk() (*int32, bool)`

GetPkiLeadIDOk returns a tuple with the PkiLeadID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiLeadID

`func (o *LeadListElement) SetPkiLeadID(v int32)`

SetPkiLeadID sets PkiLeadID field to given value.


### GetFkiLeadsourceID

`func (o *LeadListElement) GetFkiLeadsourceID() int32`

GetFkiLeadsourceID returns the FkiLeadsourceID field if non-nil, zero value otherwise.

### GetFkiLeadsourceIDOk

`func (o *LeadListElement) GetFkiLeadsourceIDOk() (*int32, bool)`

GetFkiLeadsourceIDOk returns a tuple with the FkiLeadsourceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiLeadsourceID

`func (o *LeadListElement) SetFkiLeadsourceID(v int32)`

SetFkiLeadsourceID sets FkiLeadsourceID field to given value.


### GetSLeadsourceNameX

`func (o *LeadListElement) GetSLeadsourceNameX() string`

GetSLeadsourceNameX returns the SLeadsourceNameX field if non-nil, zero value otherwise.

### GetSLeadsourceNameXOk

`func (o *LeadListElement) GetSLeadsourceNameXOk() (*string, bool)`

GetSLeadsourceNameXOk returns a tuple with the SLeadsourceNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSLeadsourceNameX

`func (o *LeadListElement) SetSLeadsourceNameX(v string)`

SetSLeadsourceNameX sets SLeadsourceNameX field to given value.


### GetELeadStatus

`func (o *LeadListElement) GetELeadStatus() FieldELeadStatus`

GetELeadStatus returns the ELeadStatus field if non-nil, zero value otherwise.

### GetELeadStatusOk

`func (o *LeadListElement) GetELeadStatusOk() (*FieldELeadStatus, bool)`

GetELeadStatusOk returns a tuple with the ELeadStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetELeadStatus

`func (o *LeadListElement) SetELeadStatus(v FieldELeadStatus)`

SetELeadStatus sets ELeadStatus field to given value.


### GetDtLeadExpiration

`func (o *LeadListElement) GetDtLeadExpiration() string`

GetDtLeadExpiration returns the DtLeadExpiration field if non-nil, zero value otherwise.

### GetDtLeadExpirationOk

`func (o *LeadListElement) GetDtLeadExpirationOk() (*string, bool)`

GetDtLeadExpirationOk returns a tuple with the DtLeadExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtLeadExpiration

`func (o *LeadListElement) SetDtLeadExpiration(v string)`

SetDtLeadExpiration sets DtLeadExpiration field to given value.


### GetBLeadIsactive

`func (o *LeadListElement) GetBLeadIsactive() bool`

GetBLeadIsactive returns the BLeadIsactive field if non-nil, zero value otherwise.

### GetBLeadIsactiveOk

`func (o *LeadListElement) GetBLeadIsactiveOk() (*bool, bool)`

GetBLeadIsactiveOk returns a tuple with the BLeadIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBLeadIsactive

`func (o *LeadListElement) SetBLeadIsactive(v bool)`

SetBLeadIsactive sets BLeadIsactive field to given value.


### GetSLeadCode

`func (o *LeadListElement) GetSLeadCode() string`

GetSLeadCode returns the SLeadCode field if non-nil, zero value otherwise.

### GetSLeadCodeOk

`func (o *LeadListElement) GetSLeadCodeOk() (*string, bool)`

GetSLeadCodeOk returns a tuple with the SLeadCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSLeadCode

`func (o *LeadListElement) SetSLeadCode(v string)`

SetSLeadCode sets SLeadCode field to given value.


### GetSLeadContacts

`func (o *LeadListElement) GetSLeadContacts() string`

GetSLeadContacts returns the SLeadContacts field if non-nil, zero value otherwise.

### GetSLeadContactsOk

`func (o *LeadListElement) GetSLeadContactsOk() (*string, bool)`

GetSLeadContactsOk returns a tuple with the SLeadContacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSLeadContacts

`func (o *LeadListElement) SetSLeadContacts(v string)`

SetSLeadContacts sets SLeadContacts field to given value.

### HasSLeadContacts

`func (o *LeadListElement) HasSLeadContacts() bool`

HasSLeadContacts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


