# DisclosureListElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiDisclosureID** | **int32** | The unique ID of the Disclosure | 
**FkiAgentID** | Pointer to **int32** | The unique ID of the Agent. | [optional] 
**FkiBrokerID** | Pointer to **int32** | The unique ID of the Broker. | [optional] 
**EDisclosureInteresttype** | [**FieldEDisclosureInteresttype**](FieldEDisclosureInteresttype.md) |  | 
**SDisclosureRelation** | **string** | The relation of the Disclosure | 
**DtDisclosureRedactiondate** | **string** | The redactiondate of the Disclosure | 
**DtDisclosureAcceptationdate** | **string** | The acceptationdate of the Disclosure | 
**DtDisclosureReceptiondate** | **string** | The receptiondate of the Disclosure | 
**BDisclosureIsactive** | **bool** | Whether the disclosure is active or not | 
**SDisclosureNumber** | Pointer to **string** | The number of a Disclosure | [optional] 
**SContactFirstname** | Pointer to **string** | The First name of the contact | [optional] 
**SContactLastname** | Pointer to **string** | The Last name of the contact | [optional] 
**SBrokerName** | Pointer to **string** | The name of the Broker | [optional] 

## Methods

### NewDisclosureListElement

`func NewDisclosureListElement(pkiDisclosureID int32, eDisclosureInteresttype FieldEDisclosureInteresttype, sDisclosureRelation string, dtDisclosureRedactiondate string, dtDisclosureAcceptationdate string, dtDisclosureReceptiondate string, bDisclosureIsactive bool, ) *DisclosureListElement`

NewDisclosureListElement instantiates a new DisclosureListElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisclosureListElementWithDefaults

`func NewDisclosureListElementWithDefaults() *DisclosureListElement`

NewDisclosureListElementWithDefaults instantiates a new DisclosureListElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiDisclosureID

`func (o *DisclosureListElement) GetPkiDisclosureID() int32`

GetPkiDisclosureID returns the PkiDisclosureID field if non-nil, zero value otherwise.

### GetPkiDisclosureIDOk

`func (o *DisclosureListElement) GetPkiDisclosureIDOk() (*int32, bool)`

GetPkiDisclosureIDOk returns a tuple with the PkiDisclosureID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiDisclosureID

`func (o *DisclosureListElement) SetPkiDisclosureID(v int32)`

SetPkiDisclosureID sets PkiDisclosureID field to given value.


### GetFkiAgentID

`func (o *DisclosureListElement) GetFkiAgentID() int32`

GetFkiAgentID returns the FkiAgentID field if non-nil, zero value otherwise.

### GetFkiAgentIDOk

`func (o *DisclosureListElement) GetFkiAgentIDOk() (*int32, bool)`

GetFkiAgentIDOk returns a tuple with the FkiAgentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiAgentID

`func (o *DisclosureListElement) SetFkiAgentID(v int32)`

SetFkiAgentID sets FkiAgentID field to given value.

### HasFkiAgentID

`func (o *DisclosureListElement) HasFkiAgentID() bool`

HasFkiAgentID returns a boolean if a field has been set.

### GetFkiBrokerID

`func (o *DisclosureListElement) GetFkiBrokerID() int32`

GetFkiBrokerID returns the FkiBrokerID field if non-nil, zero value otherwise.

### GetFkiBrokerIDOk

`func (o *DisclosureListElement) GetFkiBrokerIDOk() (*int32, bool)`

GetFkiBrokerIDOk returns a tuple with the FkiBrokerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiBrokerID

`func (o *DisclosureListElement) SetFkiBrokerID(v int32)`

SetFkiBrokerID sets FkiBrokerID field to given value.

### HasFkiBrokerID

`func (o *DisclosureListElement) HasFkiBrokerID() bool`

HasFkiBrokerID returns a boolean if a field has been set.

### GetEDisclosureInteresttype

`func (o *DisclosureListElement) GetEDisclosureInteresttype() FieldEDisclosureInteresttype`

GetEDisclosureInteresttype returns the EDisclosureInteresttype field if non-nil, zero value otherwise.

### GetEDisclosureInteresttypeOk

`func (o *DisclosureListElement) GetEDisclosureInteresttypeOk() (*FieldEDisclosureInteresttype, bool)`

GetEDisclosureInteresttypeOk returns a tuple with the EDisclosureInteresttype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEDisclosureInteresttype

`func (o *DisclosureListElement) SetEDisclosureInteresttype(v FieldEDisclosureInteresttype)`

SetEDisclosureInteresttype sets EDisclosureInteresttype field to given value.


### GetSDisclosureRelation

`func (o *DisclosureListElement) GetSDisclosureRelation() string`

GetSDisclosureRelation returns the SDisclosureRelation field if non-nil, zero value otherwise.

### GetSDisclosureRelationOk

`func (o *DisclosureListElement) GetSDisclosureRelationOk() (*string, bool)`

GetSDisclosureRelationOk returns a tuple with the SDisclosureRelation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSDisclosureRelation

`func (o *DisclosureListElement) SetSDisclosureRelation(v string)`

SetSDisclosureRelation sets SDisclosureRelation field to given value.


### GetDtDisclosureRedactiondate

`func (o *DisclosureListElement) GetDtDisclosureRedactiondate() string`

GetDtDisclosureRedactiondate returns the DtDisclosureRedactiondate field if non-nil, zero value otherwise.

### GetDtDisclosureRedactiondateOk

`func (o *DisclosureListElement) GetDtDisclosureRedactiondateOk() (*string, bool)`

GetDtDisclosureRedactiondateOk returns a tuple with the DtDisclosureRedactiondate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtDisclosureRedactiondate

`func (o *DisclosureListElement) SetDtDisclosureRedactiondate(v string)`

SetDtDisclosureRedactiondate sets DtDisclosureRedactiondate field to given value.


### GetDtDisclosureAcceptationdate

`func (o *DisclosureListElement) GetDtDisclosureAcceptationdate() string`

GetDtDisclosureAcceptationdate returns the DtDisclosureAcceptationdate field if non-nil, zero value otherwise.

### GetDtDisclosureAcceptationdateOk

`func (o *DisclosureListElement) GetDtDisclosureAcceptationdateOk() (*string, bool)`

GetDtDisclosureAcceptationdateOk returns a tuple with the DtDisclosureAcceptationdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtDisclosureAcceptationdate

`func (o *DisclosureListElement) SetDtDisclosureAcceptationdate(v string)`

SetDtDisclosureAcceptationdate sets DtDisclosureAcceptationdate field to given value.


### GetDtDisclosureReceptiondate

`func (o *DisclosureListElement) GetDtDisclosureReceptiondate() string`

GetDtDisclosureReceptiondate returns the DtDisclosureReceptiondate field if non-nil, zero value otherwise.

### GetDtDisclosureReceptiondateOk

`func (o *DisclosureListElement) GetDtDisclosureReceptiondateOk() (*string, bool)`

GetDtDisclosureReceptiondateOk returns a tuple with the DtDisclosureReceptiondate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtDisclosureReceptiondate

`func (o *DisclosureListElement) SetDtDisclosureReceptiondate(v string)`

SetDtDisclosureReceptiondate sets DtDisclosureReceptiondate field to given value.


### GetBDisclosureIsactive

`func (o *DisclosureListElement) GetBDisclosureIsactive() bool`

GetBDisclosureIsactive returns the BDisclosureIsactive field if non-nil, zero value otherwise.

### GetBDisclosureIsactiveOk

`func (o *DisclosureListElement) GetBDisclosureIsactiveOk() (*bool, bool)`

GetBDisclosureIsactiveOk returns a tuple with the BDisclosureIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDisclosureIsactive

`func (o *DisclosureListElement) SetBDisclosureIsactive(v bool)`

SetBDisclosureIsactive sets BDisclosureIsactive field to given value.


### GetSDisclosureNumber

`func (o *DisclosureListElement) GetSDisclosureNumber() string`

GetSDisclosureNumber returns the SDisclosureNumber field if non-nil, zero value otherwise.

### GetSDisclosureNumberOk

`func (o *DisclosureListElement) GetSDisclosureNumberOk() (*string, bool)`

GetSDisclosureNumberOk returns a tuple with the SDisclosureNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSDisclosureNumber

`func (o *DisclosureListElement) SetSDisclosureNumber(v string)`

SetSDisclosureNumber sets SDisclosureNumber field to given value.

### HasSDisclosureNumber

`func (o *DisclosureListElement) HasSDisclosureNumber() bool`

HasSDisclosureNumber returns a boolean if a field has been set.

### GetSContactFirstname

`func (o *DisclosureListElement) GetSContactFirstname() string`

GetSContactFirstname returns the SContactFirstname field if non-nil, zero value otherwise.

### GetSContactFirstnameOk

`func (o *DisclosureListElement) GetSContactFirstnameOk() (*string, bool)`

GetSContactFirstnameOk returns a tuple with the SContactFirstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSContactFirstname

`func (o *DisclosureListElement) SetSContactFirstname(v string)`

SetSContactFirstname sets SContactFirstname field to given value.

### HasSContactFirstname

`func (o *DisclosureListElement) HasSContactFirstname() bool`

HasSContactFirstname returns a boolean if a field has been set.

### GetSContactLastname

`func (o *DisclosureListElement) GetSContactLastname() string`

GetSContactLastname returns the SContactLastname field if non-nil, zero value otherwise.

### GetSContactLastnameOk

`func (o *DisclosureListElement) GetSContactLastnameOk() (*string, bool)`

GetSContactLastnameOk returns a tuple with the SContactLastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSContactLastname

`func (o *DisclosureListElement) SetSContactLastname(v string)`

SetSContactLastname sets SContactLastname field to given value.

### HasSContactLastname

`func (o *DisclosureListElement) HasSContactLastname() bool`

HasSContactLastname returns a boolean if a field has been set.

### GetSBrokerName

`func (o *DisclosureListElement) GetSBrokerName() string`

GetSBrokerName returns the SBrokerName field if non-nil, zero value otherwise.

### GetSBrokerNameOk

`func (o *DisclosureListElement) GetSBrokerNameOk() (*string, bool)`

GetSBrokerNameOk returns a tuple with the SBrokerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSBrokerName

`func (o *DisclosureListElement) SetSBrokerName(v string)`

SetSBrokerName sets SBrokerName field to given value.

### HasSBrokerName

`func (o *DisclosureListElement) HasSBrokerName() bool`

HasSBrokerName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


