# EzsignfoldertypeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsignfoldertypeID** | Pointer to **int32** | The unique ID of the Ezsignfoldertype. | [optional] 
**ObjEzsignfoldertypeName** | [**MultilingualEzsignfoldertypeName**](MultilingualEzsignfoldertypeName.md) |  | 
**FkiBrandingID** | **int32** | The unique ID of the Branding | 
**FkiBillingentityinternalID** | Pointer to **int32** | The unique ID of the Billingentityinternal. | [optional] 
**FkiUsergroupID** | Pointer to **int32** | The unique ID of the Usergroup | [optional] 
**FkiUsergroupIDRestricted** | Pointer to **int32** | The unique ID of the Usergroup | [optional] 
**FkiEzsigntsarequirementID** | Pointer to **int32** | The unique ID of the Ezsigntsarequirement.  Determine if a Time Stamping Authority should add a timestamp on each of the signature. Valid values:  |Value|Description| |-|-| |1|No. TSA Timestamping will requested. This will make all signatures a lot faster since no round-trip to the TSA server will be required. Timestamping will be made using eZsign server&#39;s time.| |2|Best effort. Timestamping from a Time Stamping Authority will be requested but is not mandatory. In the very improbable case it cannot be completed, the timestamping will be made using eZsign server&#39;s time. **Additional fee applies**| |3|Mandatory. Timestamping from a Time Stamping Authority will be requested and is mandatory. In the very improbable case it cannot be completed, the signature will fail and the user will be asked to retry. **Additional fee applies**| | [optional] 
**SEmailAddressSigned** | Pointer to **string** | The email address. | [optional] 
**SEmailAddressSummary** | Pointer to **string** | The email address. | [optional] 
**EEzsignfoldertypePrivacylevel** | [**FieldEEzsignfoldertypePrivacylevel**](FieldEEzsignfoldertypePrivacylevel.md) |  | 
**EEzsignfoldertypeSendreminderfrequency** | Pointer to [**FieldEEzsignfoldertypeSendreminderfrequency**](FieldEEzsignfoldertypeSendreminderfrequency.md) |  | [optional] 
**IEzsignfoldertypeArchivaldays** | **int32** | The number of days before the archival of Ezsignfolders created using this Ezsignfoldertype | 
**EEzsignfoldertypeDisposal** | [**FieldEEzsignfoldertypeDisposal**](FieldEEzsignfoldertypeDisposal.md) |  | 
**IEzsignfoldertypeDisposaldays** | Pointer to **int32** | The number of days after the archival before the disposal of the Ezsignfolder | [optional] 
**IEzsignfoldertypeDeadlinedays** | **int32** | The number of days to get all Ezsignsignatures | 
**BEzsignfoldertypeDelegate** | Pointer to **bool** | Wheter if delegation of signature is allowed to another user or not | [optional] 
**BEzsignfoldertypeReassign** | Pointer to **bool** | Wheter if Reassignment of signature is allowed to another signatory or not | [optional] 
**BEzsignfoldertypeSendattatchmentsigner** | **bool** | Whether we send the Ezsigndocument and the proof as attachment in the email | 
**BEzsignfoldertypeSendsignedtodocumentowner** | **bool** | Whether we send the signed Ezsigndocument to the Ezsigndocument&#39;s owner | 
**BEzsignfoldertypeSendsignedtofolderowner** | **bool** | Whether we send the signed Ezsigndocument to the Ezsignfolder&#39;s owner | 
**BEzsignfoldertypeSendsignedtofullgroup** | Pointer to **bool** | Whether we send the signed Ezsigndocument to the Usergroup that has acces to all Ezsignfolders | [optional] 
**BEzsignfoldertypeSendsignedtolimitedgroup** | Pointer to **bool** | Whether we send the signed Ezsigndocument to the Usergroup that has acces to only their own Ezsignfolders | [optional] 
**BEzsignfoldertypeSendsignedtocolleague** | **bool** | Whether we send the signed Ezsigndocument to the colleagues | 
**BEzsignfoldertypeSendsummarytodocumentowner** | **bool** | Whether we send the summary to the Ezsigndocument&#39;s owner | 
**BEzsignfoldertypeSendsummarytofolderowner** | **bool** | Whether we send the summary to the Ezsignfolder&#39;s owner | 
**BEzsignfoldertypeSendsummarytofullgroup** | Pointer to **bool** | Whether we send the summary to the Usergroup that has acces to all Ezsignfolders | [optional] 
**BEzsignfoldertypeSendsummarytolimitedgroup** | Pointer to **bool** | Whether we send the summary to the Usergroup that has acces to only their own Ezsignfolders | [optional] 
**BEzsignfoldertypeSendsummarytocolleague** | **bool** | Whether we send the summary to the colleagues | 
**BEzsignfoldertypeIncludeproofsigner** | **bool** | Whether we include the proof with the signed Ezsigndocument for Ezsignsigners | 
**BEzsignfoldertypeIncludeproofuser** | **bool** | Whether we include the proof with the signed Ezsigndocument for users | 
**BEzsignfoldertypeIsactive** | **bool** | Whether the Ezsignfoldertype is active or not | 

## Methods

### NewEzsignfoldertypeRequest

`func NewEzsignfoldertypeRequest(objEzsignfoldertypeName MultilingualEzsignfoldertypeName, fkiBrandingID int32, eEzsignfoldertypePrivacylevel FieldEEzsignfoldertypePrivacylevel, iEzsignfoldertypeArchivaldays int32, eEzsignfoldertypeDisposal FieldEEzsignfoldertypeDisposal, iEzsignfoldertypeDeadlinedays int32, bEzsignfoldertypeSendattatchmentsigner bool, bEzsignfoldertypeSendsignedtodocumentowner bool, bEzsignfoldertypeSendsignedtofolderowner bool, bEzsignfoldertypeSendsignedtocolleague bool, bEzsignfoldertypeSendsummarytodocumentowner bool, bEzsignfoldertypeSendsummarytofolderowner bool, bEzsignfoldertypeSendsummarytocolleague bool, bEzsignfoldertypeIncludeproofsigner bool, bEzsignfoldertypeIncludeproofuser bool, bEzsignfoldertypeIsactive bool, ) *EzsignfoldertypeRequest`

NewEzsignfoldertypeRequest instantiates a new EzsignfoldertypeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignfoldertypeRequestWithDefaults

`func NewEzsignfoldertypeRequestWithDefaults() *EzsignfoldertypeRequest`

NewEzsignfoldertypeRequestWithDefaults instantiates a new EzsignfoldertypeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsignfoldertypeID

`func (o *EzsignfoldertypeRequest) GetPkiEzsignfoldertypeID() int32`

GetPkiEzsignfoldertypeID returns the PkiEzsignfoldertypeID field if non-nil, zero value otherwise.

### GetPkiEzsignfoldertypeIDOk

`func (o *EzsignfoldertypeRequest) GetPkiEzsignfoldertypeIDOk() (*int32, bool)`

GetPkiEzsignfoldertypeIDOk returns a tuple with the PkiEzsignfoldertypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsignfoldertypeID

`func (o *EzsignfoldertypeRequest) SetPkiEzsignfoldertypeID(v int32)`

SetPkiEzsignfoldertypeID sets PkiEzsignfoldertypeID field to given value.

### HasPkiEzsignfoldertypeID

`func (o *EzsignfoldertypeRequest) HasPkiEzsignfoldertypeID() bool`

HasPkiEzsignfoldertypeID returns a boolean if a field has been set.

### GetObjEzsignfoldertypeName

`func (o *EzsignfoldertypeRequest) GetObjEzsignfoldertypeName() MultilingualEzsignfoldertypeName`

GetObjEzsignfoldertypeName returns the ObjEzsignfoldertypeName field if non-nil, zero value otherwise.

### GetObjEzsignfoldertypeNameOk

`func (o *EzsignfoldertypeRequest) GetObjEzsignfoldertypeNameOk() (*MultilingualEzsignfoldertypeName, bool)`

GetObjEzsignfoldertypeNameOk returns a tuple with the ObjEzsignfoldertypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjEzsignfoldertypeName

`func (o *EzsignfoldertypeRequest) SetObjEzsignfoldertypeName(v MultilingualEzsignfoldertypeName)`

SetObjEzsignfoldertypeName sets ObjEzsignfoldertypeName field to given value.


### GetFkiBrandingID

`func (o *EzsignfoldertypeRequest) GetFkiBrandingID() int32`

GetFkiBrandingID returns the FkiBrandingID field if non-nil, zero value otherwise.

### GetFkiBrandingIDOk

`func (o *EzsignfoldertypeRequest) GetFkiBrandingIDOk() (*int32, bool)`

GetFkiBrandingIDOk returns a tuple with the FkiBrandingID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiBrandingID

`func (o *EzsignfoldertypeRequest) SetFkiBrandingID(v int32)`

SetFkiBrandingID sets FkiBrandingID field to given value.


### GetFkiBillingentityinternalID

`func (o *EzsignfoldertypeRequest) GetFkiBillingentityinternalID() int32`

GetFkiBillingentityinternalID returns the FkiBillingentityinternalID field if non-nil, zero value otherwise.

### GetFkiBillingentityinternalIDOk

`func (o *EzsignfoldertypeRequest) GetFkiBillingentityinternalIDOk() (*int32, bool)`

GetFkiBillingentityinternalIDOk returns a tuple with the FkiBillingentityinternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiBillingentityinternalID

`func (o *EzsignfoldertypeRequest) SetFkiBillingentityinternalID(v int32)`

SetFkiBillingentityinternalID sets FkiBillingentityinternalID field to given value.

### HasFkiBillingentityinternalID

`func (o *EzsignfoldertypeRequest) HasFkiBillingentityinternalID() bool`

HasFkiBillingentityinternalID returns a boolean if a field has been set.

### GetFkiUsergroupID

`func (o *EzsignfoldertypeRequest) GetFkiUsergroupID() int32`

GetFkiUsergroupID returns the FkiUsergroupID field if non-nil, zero value otherwise.

### GetFkiUsergroupIDOk

`func (o *EzsignfoldertypeRequest) GetFkiUsergroupIDOk() (*int32, bool)`

GetFkiUsergroupIDOk returns a tuple with the FkiUsergroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUsergroupID

`func (o *EzsignfoldertypeRequest) SetFkiUsergroupID(v int32)`

SetFkiUsergroupID sets FkiUsergroupID field to given value.

### HasFkiUsergroupID

`func (o *EzsignfoldertypeRequest) HasFkiUsergroupID() bool`

HasFkiUsergroupID returns a boolean if a field has been set.

### GetFkiUsergroupIDRestricted

`func (o *EzsignfoldertypeRequest) GetFkiUsergroupIDRestricted() int32`

GetFkiUsergroupIDRestricted returns the FkiUsergroupIDRestricted field if non-nil, zero value otherwise.

### GetFkiUsergroupIDRestrictedOk

`func (o *EzsignfoldertypeRequest) GetFkiUsergroupIDRestrictedOk() (*int32, bool)`

GetFkiUsergroupIDRestrictedOk returns a tuple with the FkiUsergroupIDRestricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUsergroupIDRestricted

`func (o *EzsignfoldertypeRequest) SetFkiUsergroupIDRestricted(v int32)`

SetFkiUsergroupIDRestricted sets FkiUsergroupIDRestricted field to given value.

### HasFkiUsergroupIDRestricted

`func (o *EzsignfoldertypeRequest) HasFkiUsergroupIDRestricted() bool`

HasFkiUsergroupIDRestricted returns a boolean if a field has been set.

### GetFkiEzsigntsarequirementID

`func (o *EzsignfoldertypeRequest) GetFkiEzsigntsarequirementID() int32`

GetFkiEzsigntsarequirementID returns the FkiEzsigntsarequirementID field if non-nil, zero value otherwise.

### GetFkiEzsigntsarequirementIDOk

`func (o *EzsignfoldertypeRequest) GetFkiEzsigntsarequirementIDOk() (*int32, bool)`

GetFkiEzsigntsarequirementIDOk returns a tuple with the FkiEzsigntsarequirementID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntsarequirementID

`func (o *EzsignfoldertypeRequest) SetFkiEzsigntsarequirementID(v int32)`

SetFkiEzsigntsarequirementID sets FkiEzsigntsarequirementID field to given value.

### HasFkiEzsigntsarequirementID

`func (o *EzsignfoldertypeRequest) HasFkiEzsigntsarequirementID() bool`

HasFkiEzsigntsarequirementID returns a boolean if a field has been set.

### GetSEmailAddressSigned

`func (o *EzsignfoldertypeRequest) GetSEmailAddressSigned() string`

GetSEmailAddressSigned returns the SEmailAddressSigned field if non-nil, zero value otherwise.

### GetSEmailAddressSignedOk

`func (o *EzsignfoldertypeRequest) GetSEmailAddressSignedOk() (*string, bool)`

GetSEmailAddressSignedOk returns a tuple with the SEmailAddressSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEmailAddressSigned

`func (o *EzsignfoldertypeRequest) SetSEmailAddressSigned(v string)`

SetSEmailAddressSigned sets SEmailAddressSigned field to given value.

### HasSEmailAddressSigned

`func (o *EzsignfoldertypeRequest) HasSEmailAddressSigned() bool`

HasSEmailAddressSigned returns a boolean if a field has been set.

### GetSEmailAddressSummary

`func (o *EzsignfoldertypeRequest) GetSEmailAddressSummary() string`

GetSEmailAddressSummary returns the SEmailAddressSummary field if non-nil, zero value otherwise.

### GetSEmailAddressSummaryOk

`func (o *EzsignfoldertypeRequest) GetSEmailAddressSummaryOk() (*string, bool)`

GetSEmailAddressSummaryOk returns a tuple with the SEmailAddressSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEmailAddressSummary

`func (o *EzsignfoldertypeRequest) SetSEmailAddressSummary(v string)`

SetSEmailAddressSummary sets SEmailAddressSummary field to given value.

### HasSEmailAddressSummary

`func (o *EzsignfoldertypeRequest) HasSEmailAddressSummary() bool`

HasSEmailAddressSummary returns a boolean if a field has been set.

### GetEEzsignfoldertypePrivacylevel

`func (o *EzsignfoldertypeRequest) GetEEzsignfoldertypePrivacylevel() FieldEEzsignfoldertypePrivacylevel`

GetEEzsignfoldertypePrivacylevel returns the EEzsignfoldertypePrivacylevel field if non-nil, zero value otherwise.

### GetEEzsignfoldertypePrivacylevelOk

`func (o *EzsignfoldertypeRequest) GetEEzsignfoldertypePrivacylevelOk() (*FieldEEzsignfoldertypePrivacylevel, bool)`

GetEEzsignfoldertypePrivacylevelOk returns a tuple with the EEzsignfoldertypePrivacylevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignfoldertypePrivacylevel

`func (o *EzsignfoldertypeRequest) SetEEzsignfoldertypePrivacylevel(v FieldEEzsignfoldertypePrivacylevel)`

SetEEzsignfoldertypePrivacylevel sets EEzsignfoldertypePrivacylevel field to given value.


### GetEEzsignfoldertypeSendreminderfrequency

`func (o *EzsignfoldertypeRequest) GetEEzsignfoldertypeSendreminderfrequency() FieldEEzsignfoldertypeSendreminderfrequency`

GetEEzsignfoldertypeSendreminderfrequency returns the EEzsignfoldertypeSendreminderfrequency field if non-nil, zero value otherwise.

### GetEEzsignfoldertypeSendreminderfrequencyOk

`func (o *EzsignfoldertypeRequest) GetEEzsignfoldertypeSendreminderfrequencyOk() (*FieldEEzsignfoldertypeSendreminderfrequency, bool)`

GetEEzsignfoldertypeSendreminderfrequencyOk returns a tuple with the EEzsignfoldertypeSendreminderfrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignfoldertypeSendreminderfrequency

`func (o *EzsignfoldertypeRequest) SetEEzsignfoldertypeSendreminderfrequency(v FieldEEzsignfoldertypeSendreminderfrequency)`

SetEEzsignfoldertypeSendreminderfrequency sets EEzsignfoldertypeSendreminderfrequency field to given value.

### HasEEzsignfoldertypeSendreminderfrequency

`func (o *EzsignfoldertypeRequest) HasEEzsignfoldertypeSendreminderfrequency() bool`

HasEEzsignfoldertypeSendreminderfrequency returns a boolean if a field has been set.

### GetIEzsignfoldertypeArchivaldays

`func (o *EzsignfoldertypeRequest) GetIEzsignfoldertypeArchivaldays() int32`

GetIEzsignfoldertypeArchivaldays returns the IEzsignfoldertypeArchivaldays field if non-nil, zero value otherwise.

### GetIEzsignfoldertypeArchivaldaysOk

`func (o *EzsignfoldertypeRequest) GetIEzsignfoldertypeArchivaldaysOk() (*int32, bool)`

GetIEzsignfoldertypeArchivaldaysOk returns a tuple with the IEzsignfoldertypeArchivaldays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignfoldertypeArchivaldays

`func (o *EzsignfoldertypeRequest) SetIEzsignfoldertypeArchivaldays(v int32)`

SetIEzsignfoldertypeArchivaldays sets IEzsignfoldertypeArchivaldays field to given value.


### GetEEzsignfoldertypeDisposal

`func (o *EzsignfoldertypeRequest) GetEEzsignfoldertypeDisposal() FieldEEzsignfoldertypeDisposal`

GetEEzsignfoldertypeDisposal returns the EEzsignfoldertypeDisposal field if non-nil, zero value otherwise.

### GetEEzsignfoldertypeDisposalOk

`func (o *EzsignfoldertypeRequest) GetEEzsignfoldertypeDisposalOk() (*FieldEEzsignfoldertypeDisposal, bool)`

GetEEzsignfoldertypeDisposalOk returns a tuple with the EEzsignfoldertypeDisposal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignfoldertypeDisposal

`func (o *EzsignfoldertypeRequest) SetEEzsignfoldertypeDisposal(v FieldEEzsignfoldertypeDisposal)`

SetEEzsignfoldertypeDisposal sets EEzsignfoldertypeDisposal field to given value.


### GetIEzsignfoldertypeDisposaldays

`func (o *EzsignfoldertypeRequest) GetIEzsignfoldertypeDisposaldays() int32`

GetIEzsignfoldertypeDisposaldays returns the IEzsignfoldertypeDisposaldays field if non-nil, zero value otherwise.

### GetIEzsignfoldertypeDisposaldaysOk

`func (o *EzsignfoldertypeRequest) GetIEzsignfoldertypeDisposaldaysOk() (*int32, bool)`

GetIEzsignfoldertypeDisposaldaysOk returns a tuple with the IEzsignfoldertypeDisposaldays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignfoldertypeDisposaldays

`func (o *EzsignfoldertypeRequest) SetIEzsignfoldertypeDisposaldays(v int32)`

SetIEzsignfoldertypeDisposaldays sets IEzsignfoldertypeDisposaldays field to given value.

### HasIEzsignfoldertypeDisposaldays

`func (o *EzsignfoldertypeRequest) HasIEzsignfoldertypeDisposaldays() bool`

HasIEzsignfoldertypeDisposaldays returns a boolean if a field has been set.

### GetIEzsignfoldertypeDeadlinedays

`func (o *EzsignfoldertypeRequest) GetIEzsignfoldertypeDeadlinedays() int32`

GetIEzsignfoldertypeDeadlinedays returns the IEzsignfoldertypeDeadlinedays field if non-nil, zero value otherwise.

### GetIEzsignfoldertypeDeadlinedaysOk

`func (o *EzsignfoldertypeRequest) GetIEzsignfoldertypeDeadlinedaysOk() (*int32, bool)`

GetIEzsignfoldertypeDeadlinedaysOk returns a tuple with the IEzsignfoldertypeDeadlinedays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignfoldertypeDeadlinedays

`func (o *EzsignfoldertypeRequest) SetIEzsignfoldertypeDeadlinedays(v int32)`

SetIEzsignfoldertypeDeadlinedays sets IEzsignfoldertypeDeadlinedays field to given value.


### GetBEzsignfoldertypeDelegate

`func (o *EzsignfoldertypeRequest) GetBEzsignfoldertypeDelegate() bool`

GetBEzsignfoldertypeDelegate returns the BEzsignfoldertypeDelegate field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeDelegateOk

`func (o *EzsignfoldertypeRequest) GetBEzsignfoldertypeDelegateOk() (*bool, bool)`

GetBEzsignfoldertypeDelegateOk returns a tuple with the BEzsignfoldertypeDelegate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeDelegate

`func (o *EzsignfoldertypeRequest) SetBEzsignfoldertypeDelegate(v bool)`

SetBEzsignfoldertypeDelegate sets BEzsignfoldertypeDelegate field to given value.

### HasBEzsignfoldertypeDelegate

`func (o *EzsignfoldertypeRequest) HasBEzsignfoldertypeDelegate() bool`

HasBEzsignfoldertypeDelegate returns a boolean if a field has been set.

### GetBEzsignfoldertypeReassign

`func (o *EzsignfoldertypeRequest) GetBEzsignfoldertypeReassign() bool`

GetBEzsignfoldertypeReassign returns the BEzsignfoldertypeReassign field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeReassignOk

`func (o *EzsignfoldertypeRequest) GetBEzsignfoldertypeReassignOk() (*bool, bool)`

GetBEzsignfoldertypeReassignOk returns a tuple with the BEzsignfoldertypeReassign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeReassign

`func (o *EzsignfoldertypeRequest) SetBEzsignfoldertypeReassign(v bool)`

SetBEzsignfoldertypeReassign sets BEzsignfoldertypeReassign field to given value.

### HasBEzsignfoldertypeReassign

`func (o *EzsignfoldertypeRequest) HasBEzsignfoldertypeReassign() bool`

HasBEzsignfoldertypeReassign returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendattatchmentsigner

`func (o *EzsignfoldertypeRequest) GetBEzsignfoldertypeSendattatchmentsigner() bool`

GetBEzsignfoldertypeSendattatchmentsigner returns the BEzsignfoldertypeSendattatchmentsigner field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendattatchmentsignerOk

`func (o *EzsignfoldertypeRequest) GetBEzsignfoldertypeSendattatchmentsignerOk() (*bool, bool)`

GetBEzsignfoldertypeSendattatchmentsignerOk returns a tuple with the BEzsignfoldertypeSendattatchmentsigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendattatchmentsigner

`func (o *EzsignfoldertypeRequest) SetBEzsignfoldertypeSendattatchmentsigner(v bool)`

SetBEzsignfoldertypeSendattatchmentsigner sets BEzsignfoldertypeSendattatchmentsigner field to given value.


### GetBEzsignfoldertypeSendsignedtodocumentowner

`func (o *EzsignfoldertypeRequest) GetBEzsignfoldertypeSendsignedtodocumentowner() bool`

GetBEzsignfoldertypeSendsignedtodocumentowner returns the BEzsignfoldertypeSendsignedtodocumentowner field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsignedtodocumentownerOk

`func (o *EzsignfoldertypeRequest) GetBEzsignfoldertypeSendsignedtodocumentownerOk() (*bool, bool)`

GetBEzsignfoldertypeSendsignedtodocumentownerOk returns a tuple with the BEzsignfoldertypeSendsignedtodocumentowner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsignedtodocumentowner

`func (o *EzsignfoldertypeRequest) SetBEzsignfoldertypeSendsignedtodocumentowner(v bool)`

SetBEzsignfoldertypeSendsignedtodocumentowner sets BEzsignfoldertypeSendsignedtodocumentowner field to given value.


### GetBEzsignfoldertypeSendsignedtofolderowner

`func (o *EzsignfoldertypeRequest) GetBEzsignfoldertypeSendsignedtofolderowner() bool`

GetBEzsignfoldertypeSendsignedtofolderowner returns the BEzsignfoldertypeSendsignedtofolderowner field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsignedtofolderownerOk

`func (o *EzsignfoldertypeRequest) GetBEzsignfoldertypeSendsignedtofolderownerOk() (*bool, bool)`

GetBEzsignfoldertypeSendsignedtofolderownerOk returns a tuple with the BEzsignfoldertypeSendsignedtofolderowner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsignedtofolderowner

`func (o *EzsignfoldertypeRequest) SetBEzsignfoldertypeSendsignedtofolderowner(v bool)`

SetBEzsignfoldertypeSendsignedtofolderowner sets BEzsignfoldertypeSendsignedtofolderowner field to given value.


### GetBEzsignfoldertypeSendsignedtofullgroup

`func (o *EzsignfoldertypeRequest) GetBEzsignfoldertypeSendsignedtofullgroup() bool`

GetBEzsignfoldertypeSendsignedtofullgroup returns the BEzsignfoldertypeSendsignedtofullgroup field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsignedtofullgroupOk

`func (o *EzsignfoldertypeRequest) GetBEzsignfoldertypeSendsignedtofullgroupOk() (*bool, bool)`

GetBEzsignfoldertypeSendsignedtofullgroupOk returns a tuple with the BEzsignfoldertypeSendsignedtofullgroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsignedtofullgroup

`func (o *EzsignfoldertypeRequest) SetBEzsignfoldertypeSendsignedtofullgroup(v bool)`

SetBEzsignfoldertypeSendsignedtofullgroup sets BEzsignfoldertypeSendsignedtofullgroup field to given value.

### HasBEzsignfoldertypeSendsignedtofullgroup

`func (o *EzsignfoldertypeRequest) HasBEzsignfoldertypeSendsignedtofullgroup() bool`

HasBEzsignfoldertypeSendsignedtofullgroup returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendsignedtolimitedgroup

`func (o *EzsignfoldertypeRequest) GetBEzsignfoldertypeSendsignedtolimitedgroup() bool`

GetBEzsignfoldertypeSendsignedtolimitedgroup returns the BEzsignfoldertypeSendsignedtolimitedgroup field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsignedtolimitedgroupOk

`func (o *EzsignfoldertypeRequest) GetBEzsignfoldertypeSendsignedtolimitedgroupOk() (*bool, bool)`

GetBEzsignfoldertypeSendsignedtolimitedgroupOk returns a tuple with the BEzsignfoldertypeSendsignedtolimitedgroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsignedtolimitedgroup

`func (o *EzsignfoldertypeRequest) SetBEzsignfoldertypeSendsignedtolimitedgroup(v bool)`

SetBEzsignfoldertypeSendsignedtolimitedgroup sets BEzsignfoldertypeSendsignedtolimitedgroup field to given value.

### HasBEzsignfoldertypeSendsignedtolimitedgroup

`func (o *EzsignfoldertypeRequest) HasBEzsignfoldertypeSendsignedtolimitedgroup() bool`

HasBEzsignfoldertypeSendsignedtolimitedgroup returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendsignedtocolleague

`func (o *EzsignfoldertypeRequest) GetBEzsignfoldertypeSendsignedtocolleague() bool`

GetBEzsignfoldertypeSendsignedtocolleague returns the BEzsignfoldertypeSendsignedtocolleague field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsignedtocolleagueOk

`func (o *EzsignfoldertypeRequest) GetBEzsignfoldertypeSendsignedtocolleagueOk() (*bool, bool)`

GetBEzsignfoldertypeSendsignedtocolleagueOk returns a tuple with the BEzsignfoldertypeSendsignedtocolleague field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsignedtocolleague

`func (o *EzsignfoldertypeRequest) SetBEzsignfoldertypeSendsignedtocolleague(v bool)`

SetBEzsignfoldertypeSendsignedtocolleague sets BEzsignfoldertypeSendsignedtocolleague field to given value.


### GetBEzsignfoldertypeSendsummarytodocumentowner

`func (o *EzsignfoldertypeRequest) GetBEzsignfoldertypeSendsummarytodocumentowner() bool`

GetBEzsignfoldertypeSendsummarytodocumentowner returns the BEzsignfoldertypeSendsummarytodocumentowner field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsummarytodocumentownerOk

`func (o *EzsignfoldertypeRequest) GetBEzsignfoldertypeSendsummarytodocumentownerOk() (*bool, bool)`

GetBEzsignfoldertypeSendsummarytodocumentownerOk returns a tuple with the BEzsignfoldertypeSendsummarytodocumentowner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsummarytodocumentowner

`func (o *EzsignfoldertypeRequest) SetBEzsignfoldertypeSendsummarytodocumentowner(v bool)`

SetBEzsignfoldertypeSendsummarytodocumentowner sets BEzsignfoldertypeSendsummarytodocumentowner field to given value.


### GetBEzsignfoldertypeSendsummarytofolderowner

`func (o *EzsignfoldertypeRequest) GetBEzsignfoldertypeSendsummarytofolderowner() bool`

GetBEzsignfoldertypeSendsummarytofolderowner returns the BEzsignfoldertypeSendsummarytofolderowner field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsummarytofolderownerOk

`func (o *EzsignfoldertypeRequest) GetBEzsignfoldertypeSendsummarytofolderownerOk() (*bool, bool)`

GetBEzsignfoldertypeSendsummarytofolderownerOk returns a tuple with the BEzsignfoldertypeSendsummarytofolderowner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsummarytofolderowner

`func (o *EzsignfoldertypeRequest) SetBEzsignfoldertypeSendsummarytofolderowner(v bool)`

SetBEzsignfoldertypeSendsummarytofolderowner sets BEzsignfoldertypeSendsummarytofolderowner field to given value.


### GetBEzsignfoldertypeSendsummarytofullgroup

`func (o *EzsignfoldertypeRequest) GetBEzsignfoldertypeSendsummarytofullgroup() bool`

GetBEzsignfoldertypeSendsummarytofullgroup returns the BEzsignfoldertypeSendsummarytofullgroup field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsummarytofullgroupOk

`func (o *EzsignfoldertypeRequest) GetBEzsignfoldertypeSendsummarytofullgroupOk() (*bool, bool)`

GetBEzsignfoldertypeSendsummarytofullgroupOk returns a tuple with the BEzsignfoldertypeSendsummarytofullgroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsummarytofullgroup

`func (o *EzsignfoldertypeRequest) SetBEzsignfoldertypeSendsummarytofullgroup(v bool)`

SetBEzsignfoldertypeSendsummarytofullgroup sets BEzsignfoldertypeSendsummarytofullgroup field to given value.

### HasBEzsignfoldertypeSendsummarytofullgroup

`func (o *EzsignfoldertypeRequest) HasBEzsignfoldertypeSendsummarytofullgroup() bool`

HasBEzsignfoldertypeSendsummarytofullgroup returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendsummarytolimitedgroup

`func (o *EzsignfoldertypeRequest) GetBEzsignfoldertypeSendsummarytolimitedgroup() bool`

GetBEzsignfoldertypeSendsummarytolimitedgroup returns the BEzsignfoldertypeSendsummarytolimitedgroup field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsummarytolimitedgroupOk

`func (o *EzsignfoldertypeRequest) GetBEzsignfoldertypeSendsummarytolimitedgroupOk() (*bool, bool)`

GetBEzsignfoldertypeSendsummarytolimitedgroupOk returns a tuple with the BEzsignfoldertypeSendsummarytolimitedgroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsummarytolimitedgroup

`func (o *EzsignfoldertypeRequest) SetBEzsignfoldertypeSendsummarytolimitedgroup(v bool)`

SetBEzsignfoldertypeSendsummarytolimitedgroup sets BEzsignfoldertypeSendsummarytolimitedgroup field to given value.

### HasBEzsignfoldertypeSendsummarytolimitedgroup

`func (o *EzsignfoldertypeRequest) HasBEzsignfoldertypeSendsummarytolimitedgroup() bool`

HasBEzsignfoldertypeSendsummarytolimitedgroup returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendsummarytocolleague

`func (o *EzsignfoldertypeRequest) GetBEzsignfoldertypeSendsummarytocolleague() bool`

GetBEzsignfoldertypeSendsummarytocolleague returns the BEzsignfoldertypeSendsummarytocolleague field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsummarytocolleagueOk

`func (o *EzsignfoldertypeRequest) GetBEzsignfoldertypeSendsummarytocolleagueOk() (*bool, bool)`

GetBEzsignfoldertypeSendsummarytocolleagueOk returns a tuple with the BEzsignfoldertypeSendsummarytocolleague field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsummarytocolleague

`func (o *EzsignfoldertypeRequest) SetBEzsignfoldertypeSendsummarytocolleague(v bool)`

SetBEzsignfoldertypeSendsummarytocolleague sets BEzsignfoldertypeSendsummarytocolleague field to given value.


### GetBEzsignfoldertypeIncludeproofsigner

`func (o *EzsignfoldertypeRequest) GetBEzsignfoldertypeIncludeproofsigner() bool`

GetBEzsignfoldertypeIncludeproofsigner returns the BEzsignfoldertypeIncludeproofsigner field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeIncludeproofsignerOk

`func (o *EzsignfoldertypeRequest) GetBEzsignfoldertypeIncludeproofsignerOk() (*bool, bool)`

GetBEzsignfoldertypeIncludeproofsignerOk returns a tuple with the BEzsignfoldertypeIncludeproofsigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeIncludeproofsigner

`func (o *EzsignfoldertypeRequest) SetBEzsignfoldertypeIncludeproofsigner(v bool)`

SetBEzsignfoldertypeIncludeproofsigner sets BEzsignfoldertypeIncludeproofsigner field to given value.


### GetBEzsignfoldertypeIncludeproofuser

`func (o *EzsignfoldertypeRequest) GetBEzsignfoldertypeIncludeproofuser() bool`

GetBEzsignfoldertypeIncludeproofuser returns the BEzsignfoldertypeIncludeproofuser field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeIncludeproofuserOk

`func (o *EzsignfoldertypeRequest) GetBEzsignfoldertypeIncludeproofuserOk() (*bool, bool)`

GetBEzsignfoldertypeIncludeproofuserOk returns a tuple with the BEzsignfoldertypeIncludeproofuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeIncludeproofuser

`func (o *EzsignfoldertypeRequest) SetBEzsignfoldertypeIncludeproofuser(v bool)`

SetBEzsignfoldertypeIncludeproofuser sets BEzsignfoldertypeIncludeproofuser field to given value.


### GetBEzsignfoldertypeIsactive

`func (o *EzsignfoldertypeRequest) GetBEzsignfoldertypeIsactive() bool`

GetBEzsignfoldertypeIsactive returns the BEzsignfoldertypeIsactive field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeIsactiveOk

`func (o *EzsignfoldertypeRequest) GetBEzsignfoldertypeIsactiveOk() (*bool, bool)`

GetBEzsignfoldertypeIsactiveOk returns a tuple with the BEzsignfoldertypeIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeIsactive

`func (o *EzsignfoldertypeRequest) SetBEzsignfoldertypeIsactive(v bool)`

SetBEzsignfoldertypeIsactive sets BEzsignfoldertypeIsactive field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

