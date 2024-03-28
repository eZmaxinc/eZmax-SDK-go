# EzsignfoldertypeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsignfoldertypeID** | **int32** | The unique ID of the Ezsignfoldertype. | 
**ObjEzsignfoldertypeName** | [**MultilingualEzsignfoldertypeName**](MultilingualEzsignfoldertypeName.md) |  | 
**FkiBrandingID** | **int32** | The unique ID of the Branding | 
**FkiBillingentityinternalID** | Pointer to **int32** | The unique ID of the Billingentityinternal. | [optional] 
**FkiUsergroupID** | Pointer to **int32** | The unique ID of the Usergroup | [optional] 
**FkiUsergroupIDRestricted** | Pointer to **int32** | The unique ID of the Usergroup | [optional] 
**FkiEzsigntsarequirementID** | Pointer to **int32** | The unique ID of the Ezsigntsarequirement.  Determine if a Time Stamping Authority should add a timestamp on each of the signature. Valid values:  |Value|Description| |-|-| |1|No. TSA Timestamping will requested. This will make all signatures a lot faster since no round-trip to the TSA server will be required. Timestamping will be made using eZsign server&#39;s time.| |2|Best effort. Timestamping from a Time Stamping Authority will be requested but is not mandatory. In the very improbable case it cannot be completed, the timestamping will be made using eZsign server&#39;s time. **Additional fee applies**| |3|Mandatory. Timestamping from a Time Stamping Authority will be requested and is mandatory. In the very improbable case it cannot be completed, the signature will fail and the user will be asked to retry. **Additional fee applies**| | [optional] 
**SBrandingDescriptionX** | **string** | The Description of the Branding in the language of the requester | 
**SBillingentityinternalDescriptionX** | Pointer to **string** | The description of the Billingentityinternal in the language of the requester | [optional] 
**SEzsigntsarequirementDescriptionX** | Pointer to **string** | The description of the Ezsigntsarequirement in the language of the requester | [optional] 
**SEmailAddressSigned** | Pointer to **string** | The email address. | [optional] 
**SEmailAddressSummary** | Pointer to **string** | The email address. | [optional] 
**SUsergroupNameX** | Pointer to **string** | The Name of the Usergroup in the language of the requester | [optional] 
**SUsergroupNameXRestricted** | Pointer to **string** | The Name of the Usergroup in the language of the requester | [optional] 
**EEzsignfoldertypePrivacylevel** | [**FieldEEzsignfoldertypePrivacylevel**](FieldEEzsignfoldertypePrivacylevel.md) |  | 
**EEzsignfoldertypeSendreminderfrequency** | Pointer to [**FieldEEzsignfoldertypeSendreminderfrequency**](FieldEEzsignfoldertypeSendreminderfrequency.md) |  | [optional] 
**IEzsignfoldertypeArchivaldays** | **int32** | The number of days before the archival of Ezsignfolders created using this Ezsignfoldertype | 
**EEzsignfoldertypeDisposal** | [**FieldEEzsignfoldertypeDisposal**](FieldEEzsignfoldertypeDisposal.md) |  | 
**EEzsignfoldertypeCompletion** | [**FieldEEzsignfoldertypeCompletion**](FieldEEzsignfoldertypeCompletion.md) |  | 
**IEzsignfoldertypeDisposaldays** | Pointer to **int32** | The number of days after the archival before the disposal of the Ezsignfolder | [optional] 
**IEzsignfoldertypeDeadlinedays** | **int32** | The number of days to get all Ezsignsignatures | 
**BEzsignfoldertypeDelegate** | Pointer to **bool** | Wheter if delegation of signature is allowed to another user or not | [optional] 
**BEzsignfoldertypeDiscussion** | Pointer to **bool** | Wheter if creating a new Discussion is allowed or not | [optional] 
**BEzsignfoldertypeReassignezsignsigner** | Pointer to **bool** | Wheter if Reassignment of signature is allowed by a signatory to another signatory or not | [optional] 
**BEzsignfoldertypeReassignuser** | Pointer to **bool** | Wheter if Reassignment of signature is allowed by a user to a signatory or another user or not | [optional] 
**BEzsignfoldertypeSendsignedtoezsignsigner** | Pointer to **bool** | Whether we send an email to Ezsignsigner  when document is completed | [optional] 
**BEzsignfoldertypeSendsignedtouser** | Pointer to **bool** | Whether we send an email to User who signed when document is completed | [optional] 
**BEzsignfoldertypeSendattachmentezsignsigner** | Pointer to **bool** | Whether we send the Ezsigndocument in the email to Ezsignsigner | [optional] 
**BEzsignfoldertypeSendproofezsignsigner** | Pointer to **bool** | Whether we send the proof in the email to Ezsignsigner | [optional] 
**BEzsignfoldertypeSendattachmentuser** | Pointer to **bool** | Whether we send the Ezsigndocument in the email to User | [optional] 
**BEzsignfoldertypeSendproofuser** | Pointer to **bool** | Whether we send the proof in the email to User | [optional] 
**BEzsignfoldertypeSendproofemail** | Pointer to **bool** | Whether we send the proof in the email to external recipient | [optional] 
**BEzsignfoldertypeAllowdownloadattachmentezsignsigner** | Pointer to **bool** | Whether we allow the Ezsigndocument to be downloaded by an Ezsignsigner | [optional] 
**BEzsignfoldertypeAllowdownloadproofezsignsigner** | Pointer to **bool** | Whether we allow the proof to be downloaded by an Ezsignsigner | [optional] 
**BEzsignfoldertypeSendproofreceivealldocument** | Pointer to **bool** | Whether we send the proof to user and Ezsignsigner who receive all documents. | [optional] 
**BEzsignfoldertypeSendsignedtodocumentowner** | **bool** | Whether we send the signed Ezsigndocument to the Ezsigndocument&#39;s owner | 
**BEzsignfoldertypeSendsignedtofolderowner** | **bool** | Whether we send the signed Ezsigndocument to the Ezsignfolder&#39;s owner | 
**BEzsignfoldertypeSendsignedtofullgroup** | Pointer to **bool** | Whether we send the signed Ezsigndocument to the Usergroup that has acces to all Ezsignfolders | [optional] 
**BEzsignfoldertypeSendsignedtolimitedgroup** | Pointer to **bool** | THIS FIELD WILL BE DELETED. Whether we send the signed Ezsigndocument to the Usergroup that has acces to only their own Ezsignfolders | [optional] 
**BEzsignfoldertypeSendsignedtocolleague** | **bool** | Whether we send the signed Ezsigndocument to the colleagues | 
**BEzsignfoldertypeSendsummarytodocumentowner** | **bool** | Whether we send the summary to the Ezsigndocument&#39;s owner | 
**BEzsignfoldertypeSendsummarytofolderowner** | **bool** | Whether we send the summary to the Ezsignfolder&#39;s owner | 
**BEzsignfoldertypeSendsummarytofullgroup** | Pointer to **bool** | Whether we send the summary to the Usergroup that has acces to all Ezsignfolders | [optional] 
**BEzsignfoldertypeSendsummarytolimitedgroup** | Pointer to **bool** | Whether we send the summary to the Usergroup that has acces to only their own Ezsignfolders | [optional] 
**BEzsignfoldertypeSendsummarytocolleague** | **bool** | Whether we send the summary to the colleagues | 
**BEzsignfoldertypeIsactive** | **bool** | Whether the Ezsignfoldertype is active or not | 
**AObjUserlogintype** | [**[]UserlogintypeResponse**](UserlogintypeResponse.md) |  | 

## Methods

### NewEzsignfoldertypeResponse

`func NewEzsignfoldertypeResponse(pkiEzsignfoldertypeID int32, objEzsignfoldertypeName MultilingualEzsignfoldertypeName, fkiBrandingID int32, sBrandingDescriptionX string, eEzsignfoldertypePrivacylevel FieldEEzsignfoldertypePrivacylevel, iEzsignfoldertypeArchivaldays int32, eEzsignfoldertypeDisposal FieldEEzsignfoldertypeDisposal, eEzsignfoldertypeCompletion FieldEEzsignfoldertypeCompletion, iEzsignfoldertypeDeadlinedays int32, bEzsignfoldertypeSendsignedtodocumentowner bool, bEzsignfoldertypeSendsignedtofolderowner bool, bEzsignfoldertypeSendsignedtocolleague bool, bEzsignfoldertypeSendsummarytodocumentowner bool, bEzsignfoldertypeSendsummarytofolderowner bool, bEzsignfoldertypeSendsummarytocolleague bool, bEzsignfoldertypeIsactive bool, aObjUserlogintype []UserlogintypeResponse, ) *EzsignfoldertypeResponse`

NewEzsignfoldertypeResponse instantiates a new EzsignfoldertypeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignfoldertypeResponseWithDefaults

`func NewEzsignfoldertypeResponseWithDefaults() *EzsignfoldertypeResponse`

NewEzsignfoldertypeResponseWithDefaults instantiates a new EzsignfoldertypeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsignfoldertypeID

`func (o *EzsignfoldertypeResponse) GetPkiEzsignfoldertypeID() int32`

GetPkiEzsignfoldertypeID returns the PkiEzsignfoldertypeID field if non-nil, zero value otherwise.

### GetPkiEzsignfoldertypeIDOk

`func (o *EzsignfoldertypeResponse) GetPkiEzsignfoldertypeIDOk() (*int32, bool)`

GetPkiEzsignfoldertypeIDOk returns a tuple with the PkiEzsignfoldertypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsignfoldertypeID

`func (o *EzsignfoldertypeResponse) SetPkiEzsignfoldertypeID(v int32)`

SetPkiEzsignfoldertypeID sets PkiEzsignfoldertypeID field to given value.


### GetObjEzsignfoldertypeName

`func (o *EzsignfoldertypeResponse) GetObjEzsignfoldertypeName() MultilingualEzsignfoldertypeName`

GetObjEzsignfoldertypeName returns the ObjEzsignfoldertypeName field if non-nil, zero value otherwise.

### GetObjEzsignfoldertypeNameOk

`func (o *EzsignfoldertypeResponse) GetObjEzsignfoldertypeNameOk() (*MultilingualEzsignfoldertypeName, bool)`

GetObjEzsignfoldertypeNameOk returns a tuple with the ObjEzsignfoldertypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjEzsignfoldertypeName

`func (o *EzsignfoldertypeResponse) SetObjEzsignfoldertypeName(v MultilingualEzsignfoldertypeName)`

SetObjEzsignfoldertypeName sets ObjEzsignfoldertypeName field to given value.


### GetFkiBrandingID

`func (o *EzsignfoldertypeResponse) GetFkiBrandingID() int32`

GetFkiBrandingID returns the FkiBrandingID field if non-nil, zero value otherwise.

### GetFkiBrandingIDOk

`func (o *EzsignfoldertypeResponse) GetFkiBrandingIDOk() (*int32, bool)`

GetFkiBrandingIDOk returns a tuple with the FkiBrandingID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiBrandingID

`func (o *EzsignfoldertypeResponse) SetFkiBrandingID(v int32)`

SetFkiBrandingID sets FkiBrandingID field to given value.


### GetFkiBillingentityinternalID

`func (o *EzsignfoldertypeResponse) GetFkiBillingentityinternalID() int32`

GetFkiBillingentityinternalID returns the FkiBillingentityinternalID field if non-nil, zero value otherwise.

### GetFkiBillingentityinternalIDOk

`func (o *EzsignfoldertypeResponse) GetFkiBillingentityinternalIDOk() (*int32, bool)`

GetFkiBillingentityinternalIDOk returns a tuple with the FkiBillingentityinternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiBillingentityinternalID

`func (o *EzsignfoldertypeResponse) SetFkiBillingentityinternalID(v int32)`

SetFkiBillingentityinternalID sets FkiBillingentityinternalID field to given value.

### HasFkiBillingentityinternalID

`func (o *EzsignfoldertypeResponse) HasFkiBillingentityinternalID() bool`

HasFkiBillingentityinternalID returns a boolean if a field has been set.

### GetFkiUsergroupID

`func (o *EzsignfoldertypeResponse) GetFkiUsergroupID() int32`

GetFkiUsergroupID returns the FkiUsergroupID field if non-nil, zero value otherwise.

### GetFkiUsergroupIDOk

`func (o *EzsignfoldertypeResponse) GetFkiUsergroupIDOk() (*int32, bool)`

GetFkiUsergroupIDOk returns a tuple with the FkiUsergroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUsergroupID

`func (o *EzsignfoldertypeResponse) SetFkiUsergroupID(v int32)`

SetFkiUsergroupID sets FkiUsergroupID field to given value.

### HasFkiUsergroupID

`func (o *EzsignfoldertypeResponse) HasFkiUsergroupID() bool`

HasFkiUsergroupID returns a boolean if a field has been set.

### GetFkiUsergroupIDRestricted

`func (o *EzsignfoldertypeResponse) GetFkiUsergroupIDRestricted() int32`

GetFkiUsergroupIDRestricted returns the FkiUsergroupIDRestricted field if non-nil, zero value otherwise.

### GetFkiUsergroupIDRestrictedOk

`func (o *EzsignfoldertypeResponse) GetFkiUsergroupIDRestrictedOk() (*int32, bool)`

GetFkiUsergroupIDRestrictedOk returns a tuple with the FkiUsergroupIDRestricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUsergroupIDRestricted

`func (o *EzsignfoldertypeResponse) SetFkiUsergroupIDRestricted(v int32)`

SetFkiUsergroupIDRestricted sets FkiUsergroupIDRestricted field to given value.

### HasFkiUsergroupIDRestricted

`func (o *EzsignfoldertypeResponse) HasFkiUsergroupIDRestricted() bool`

HasFkiUsergroupIDRestricted returns a boolean if a field has been set.

### GetFkiEzsigntsarequirementID

`func (o *EzsignfoldertypeResponse) GetFkiEzsigntsarequirementID() int32`

GetFkiEzsigntsarequirementID returns the FkiEzsigntsarequirementID field if non-nil, zero value otherwise.

### GetFkiEzsigntsarequirementIDOk

`func (o *EzsignfoldertypeResponse) GetFkiEzsigntsarequirementIDOk() (*int32, bool)`

GetFkiEzsigntsarequirementIDOk returns a tuple with the FkiEzsigntsarequirementID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntsarequirementID

`func (o *EzsignfoldertypeResponse) SetFkiEzsigntsarequirementID(v int32)`

SetFkiEzsigntsarequirementID sets FkiEzsigntsarequirementID field to given value.

### HasFkiEzsigntsarequirementID

`func (o *EzsignfoldertypeResponse) HasFkiEzsigntsarequirementID() bool`

HasFkiEzsigntsarequirementID returns a boolean if a field has been set.

### GetSBrandingDescriptionX

`func (o *EzsignfoldertypeResponse) GetSBrandingDescriptionX() string`

GetSBrandingDescriptionX returns the SBrandingDescriptionX field if non-nil, zero value otherwise.

### GetSBrandingDescriptionXOk

`func (o *EzsignfoldertypeResponse) GetSBrandingDescriptionXOk() (*string, bool)`

GetSBrandingDescriptionXOk returns a tuple with the SBrandingDescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSBrandingDescriptionX

`func (o *EzsignfoldertypeResponse) SetSBrandingDescriptionX(v string)`

SetSBrandingDescriptionX sets SBrandingDescriptionX field to given value.


### GetSBillingentityinternalDescriptionX

`func (o *EzsignfoldertypeResponse) GetSBillingentityinternalDescriptionX() string`

GetSBillingentityinternalDescriptionX returns the SBillingentityinternalDescriptionX field if non-nil, zero value otherwise.

### GetSBillingentityinternalDescriptionXOk

`func (o *EzsignfoldertypeResponse) GetSBillingentityinternalDescriptionXOk() (*string, bool)`

GetSBillingentityinternalDescriptionXOk returns a tuple with the SBillingentityinternalDescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSBillingentityinternalDescriptionX

`func (o *EzsignfoldertypeResponse) SetSBillingentityinternalDescriptionX(v string)`

SetSBillingentityinternalDescriptionX sets SBillingentityinternalDescriptionX field to given value.

### HasSBillingentityinternalDescriptionX

`func (o *EzsignfoldertypeResponse) HasSBillingentityinternalDescriptionX() bool`

HasSBillingentityinternalDescriptionX returns a boolean if a field has been set.

### GetSEzsigntsarequirementDescriptionX

`func (o *EzsignfoldertypeResponse) GetSEzsigntsarequirementDescriptionX() string`

GetSEzsigntsarequirementDescriptionX returns the SEzsigntsarequirementDescriptionX field if non-nil, zero value otherwise.

### GetSEzsigntsarequirementDescriptionXOk

`func (o *EzsignfoldertypeResponse) GetSEzsigntsarequirementDescriptionXOk() (*string, bool)`

GetSEzsigntsarequirementDescriptionXOk returns a tuple with the SEzsigntsarequirementDescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntsarequirementDescriptionX

`func (o *EzsignfoldertypeResponse) SetSEzsigntsarequirementDescriptionX(v string)`

SetSEzsigntsarequirementDescriptionX sets SEzsigntsarequirementDescriptionX field to given value.

### HasSEzsigntsarequirementDescriptionX

`func (o *EzsignfoldertypeResponse) HasSEzsigntsarequirementDescriptionX() bool`

HasSEzsigntsarequirementDescriptionX returns a boolean if a field has been set.

### GetSEmailAddressSigned

`func (o *EzsignfoldertypeResponse) GetSEmailAddressSigned() string`

GetSEmailAddressSigned returns the SEmailAddressSigned field if non-nil, zero value otherwise.

### GetSEmailAddressSignedOk

`func (o *EzsignfoldertypeResponse) GetSEmailAddressSignedOk() (*string, bool)`

GetSEmailAddressSignedOk returns a tuple with the SEmailAddressSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEmailAddressSigned

`func (o *EzsignfoldertypeResponse) SetSEmailAddressSigned(v string)`

SetSEmailAddressSigned sets SEmailAddressSigned field to given value.

### HasSEmailAddressSigned

`func (o *EzsignfoldertypeResponse) HasSEmailAddressSigned() bool`

HasSEmailAddressSigned returns a boolean if a field has been set.

### GetSEmailAddressSummary

`func (o *EzsignfoldertypeResponse) GetSEmailAddressSummary() string`

GetSEmailAddressSummary returns the SEmailAddressSummary field if non-nil, zero value otherwise.

### GetSEmailAddressSummaryOk

`func (o *EzsignfoldertypeResponse) GetSEmailAddressSummaryOk() (*string, bool)`

GetSEmailAddressSummaryOk returns a tuple with the SEmailAddressSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEmailAddressSummary

`func (o *EzsignfoldertypeResponse) SetSEmailAddressSummary(v string)`

SetSEmailAddressSummary sets SEmailAddressSummary field to given value.

### HasSEmailAddressSummary

`func (o *EzsignfoldertypeResponse) HasSEmailAddressSummary() bool`

HasSEmailAddressSummary returns a boolean if a field has been set.

### GetSUsergroupNameX

`func (o *EzsignfoldertypeResponse) GetSUsergroupNameX() string`

GetSUsergroupNameX returns the SUsergroupNameX field if non-nil, zero value otherwise.

### GetSUsergroupNameXOk

`func (o *EzsignfoldertypeResponse) GetSUsergroupNameXOk() (*string, bool)`

GetSUsergroupNameXOk returns a tuple with the SUsergroupNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUsergroupNameX

`func (o *EzsignfoldertypeResponse) SetSUsergroupNameX(v string)`

SetSUsergroupNameX sets SUsergroupNameX field to given value.

### HasSUsergroupNameX

`func (o *EzsignfoldertypeResponse) HasSUsergroupNameX() bool`

HasSUsergroupNameX returns a boolean if a field has been set.

### GetSUsergroupNameXRestricted

`func (o *EzsignfoldertypeResponse) GetSUsergroupNameXRestricted() string`

GetSUsergroupNameXRestricted returns the SUsergroupNameXRestricted field if non-nil, zero value otherwise.

### GetSUsergroupNameXRestrictedOk

`func (o *EzsignfoldertypeResponse) GetSUsergroupNameXRestrictedOk() (*string, bool)`

GetSUsergroupNameXRestrictedOk returns a tuple with the SUsergroupNameXRestricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUsergroupNameXRestricted

`func (o *EzsignfoldertypeResponse) SetSUsergroupNameXRestricted(v string)`

SetSUsergroupNameXRestricted sets SUsergroupNameXRestricted field to given value.

### HasSUsergroupNameXRestricted

`func (o *EzsignfoldertypeResponse) HasSUsergroupNameXRestricted() bool`

HasSUsergroupNameXRestricted returns a boolean if a field has been set.

### GetEEzsignfoldertypePrivacylevel

`func (o *EzsignfoldertypeResponse) GetEEzsignfoldertypePrivacylevel() FieldEEzsignfoldertypePrivacylevel`

GetEEzsignfoldertypePrivacylevel returns the EEzsignfoldertypePrivacylevel field if non-nil, zero value otherwise.

### GetEEzsignfoldertypePrivacylevelOk

`func (o *EzsignfoldertypeResponse) GetEEzsignfoldertypePrivacylevelOk() (*FieldEEzsignfoldertypePrivacylevel, bool)`

GetEEzsignfoldertypePrivacylevelOk returns a tuple with the EEzsignfoldertypePrivacylevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignfoldertypePrivacylevel

`func (o *EzsignfoldertypeResponse) SetEEzsignfoldertypePrivacylevel(v FieldEEzsignfoldertypePrivacylevel)`

SetEEzsignfoldertypePrivacylevel sets EEzsignfoldertypePrivacylevel field to given value.


### GetEEzsignfoldertypeSendreminderfrequency

`func (o *EzsignfoldertypeResponse) GetEEzsignfoldertypeSendreminderfrequency() FieldEEzsignfoldertypeSendreminderfrequency`

GetEEzsignfoldertypeSendreminderfrequency returns the EEzsignfoldertypeSendreminderfrequency field if non-nil, zero value otherwise.

### GetEEzsignfoldertypeSendreminderfrequencyOk

`func (o *EzsignfoldertypeResponse) GetEEzsignfoldertypeSendreminderfrequencyOk() (*FieldEEzsignfoldertypeSendreminderfrequency, bool)`

GetEEzsignfoldertypeSendreminderfrequencyOk returns a tuple with the EEzsignfoldertypeSendreminderfrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignfoldertypeSendreminderfrequency

`func (o *EzsignfoldertypeResponse) SetEEzsignfoldertypeSendreminderfrequency(v FieldEEzsignfoldertypeSendreminderfrequency)`

SetEEzsignfoldertypeSendreminderfrequency sets EEzsignfoldertypeSendreminderfrequency field to given value.

### HasEEzsignfoldertypeSendreminderfrequency

`func (o *EzsignfoldertypeResponse) HasEEzsignfoldertypeSendreminderfrequency() bool`

HasEEzsignfoldertypeSendreminderfrequency returns a boolean if a field has been set.

### GetIEzsignfoldertypeArchivaldays

`func (o *EzsignfoldertypeResponse) GetIEzsignfoldertypeArchivaldays() int32`

GetIEzsignfoldertypeArchivaldays returns the IEzsignfoldertypeArchivaldays field if non-nil, zero value otherwise.

### GetIEzsignfoldertypeArchivaldaysOk

`func (o *EzsignfoldertypeResponse) GetIEzsignfoldertypeArchivaldaysOk() (*int32, bool)`

GetIEzsignfoldertypeArchivaldaysOk returns a tuple with the IEzsignfoldertypeArchivaldays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignfoldertypeArchivaldays

`func (o *EzsignfoldertypeResponse) SetIEzsignfoldertypeArchivaldays(v int32)`

SetIEzsignfoldertypeArchivaldays sets IEzsignfoldertypeArchivaldays field to given value.


### GetEEzsignfoldertypeDisposal

`func (o *EzsignfoldertypeResponse) GetEEzsignfoldertypeDisposal() FieldEEzsignfoldertypeDisposal`

GetEEzsignfoldertypeDisposal returns the EEzsignfoldertypeDisposal field if non-nil, zero value otherwise.

### GetEEzsignfoldertypeDisposalOk

`func (o *EzsignfoldertypeResponse) GetEEzsignfoldertypeDisposalOk() (*FieldEEzsignfoldertypeDisposal, bool)`

GetEEzsignfoldertypeDisposalOk returns a tuple with the EEzsignfoldertypeDisposal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignfoldertypeDisposal

`func (o *EzsignfoldertypeResponse) SetEEzsignfoldertypeDisposal(v FieldEEzsignfoldertypeDisposal)`

SetEEzsignfoldertypeDisposal sets EEzsignfoldertypeDisposal field to given value.


### GetEEzsignfoldertypeCompletion

`func (o *EzsignfoldertypeResponse) GetEEzsignfoldertypeCompletion() FieldEEzsignfoldertypeCompletion`

GetEEzsignfoldertypeCompletion returns the EEzsignfoldertypeCompletion field if non-nil, zero value otherwise.

### GetEEzsignfoldertypeCompletionOk

`func (o *EzsignfoldertypeResponse) GetEEzsignfoldertypeCompletionOk() (*FieldEEzsignfoldertypeCompletion, bool)`

GetEEzsignfoldertypeCompletionOk returns a tuple with the EEzsignfoldertypeCompletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignfoldertypeCompletion

`func (o *EzsignfoldertypeResponse) SetEEzsignfoldertypeCompletion(v FieldEEzsignfoldertypeCompletion)`

SetEEzsignfoldertypeCompletion sets EEzsignfoldertypeCompletion field to given value.


### GetIEzsignfoldertypeDisposaldays

`func (o *EzsignfoldertypeResponse) GetIEzsignfoldertypeDisposaldays() int32`

GetIEzsignfoldertypeDisposaldays returns the IEzsignfoldertypeDisposaldays field if non-nil, zero value otherwise.

### GetIEzsignfoldertypeDisposaldaysOk

`func (o *EzsignfoldertypeResponse) GetIEzsignfoldertypeDisposaldaysOk() (*int32, bool)`

GetIEzsignfoldertypeDisposaldaysOk returns a tuple with the IEzsignfoldertypeDisposaldays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignfoldertypeDisposaldays

`func (o *EzsignfoldertypeResponse) SetIEzsignfoldertypeDisposaldays(v int32)`

SetIEzsignfoldertypeDisposaldays sets IEzsignfoldertypeDisposaldays field to given value.

### HasIEzsignfoldertypeDisposaldays

`func (o *EzsignfoldertypeResponse) HasIEzsignfoldertypeDisposaldays() bool`

HasIEzsignfoldertypeDisposaldays returns a boolean if a field has been set.

### GetIEzsignfoldertypeDeadlinedays

`func (o *EzsignfoldertypeResponse) GetIEzsignfoldertypeDeadlinedays() int32`

GetIEzsignfoldertypeDeadlinedays returns the IEzsignfoldertypeDeadlinedays field if non-nil, zero value otherwise.

### GetIEzsignfoldertypeDeadlinedaysOk

`func (o *EzsignfoldertypeResponse) GetIEzsignfoldertypeDeadlinedaysOk() (*int32, bool)`

GetIEzsignfoldertypeDeadlinedaysOk returns a tuple with the IEzsignfoldertypeDeadlinedays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignfoldertypeDeadlinedays

`func (o *EzsignfoldertypeResponse) SetIEzsignfoldertypeDeadlinedays(v int32)`

SetIEzsignfoldertypeDeadlinedays sets IEzsignfoldertypeDeadlinedays field to given value.


### GetBEzsignfoldertypeDelegate

`func (o *EzsignfoldertypeResponse) GetBEzsignfoldertypeDelegate() bool`

GetBEzsignfoldertypeDelegate returns the BEzsignfoldertypeDelegate field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeDelegateOk

`func (o *EzsignfoldertypeResponse) GetBEzsignfoldertypeDelegateOk() (*bool, bool)`

GetBEzsignfoldertypeDelegateOk returns a tuple with the BEzsignfoldertypeDelegate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeDelegate

`func (o *EzsignfoldertypeResponse) SetBEzsignfoldertypeDelegate(v bool)`

SetBEzsignfoldertypeDelegate sets BEzsignfoldertypeDelegate field to given value.

### HasBEzsignfoldertypeDelegate

`func (o *EzsignfoldertypeResponse) HasBEzsignfoldertypeDelegate() bool`

HasBEzsignfoldertypeDelegate returns a boolean if a field has been set.

### GetBEzsignfoldertypeDiscussion

`func (o *EzsignfoldertypeResponse) GetBEzsignfoldertypeDiscussion() bool`

GetBEzsignfoldertypeDiscussion returns the BEzsignfoldertypeDiscussion field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeDiscussionOk

`func (o *EzsignfoldertypeResponse) GetBEzsignfoldertypeDiscussionOk() (*bool, bool)`

GetBEzsignfoldertypeDiscussionOk returns a tuple with the BEzsignfoldertypeDiscussion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeDiscussion

`func (o *EzsignfoldertypeResponse) SetBEzsignfoldertypeDiscussion(v bool)`

SetBEzsignfoldertypeDiscussion sets BEzsignfoldertypeDiscussion field to given value.

### HasBEzsignfoldertypeDiscussion

`func (o *EzsignfoldertypeResponse) HasBEzsignfoldertypeDiscussion() bool`

HasBEzsignfoldertypeDiscussion returns a boolean if a field has been set.

### GetBEzsignfoldertypeReassignezsignsigner

`func (o *EzsignfoldertypeResponse) GetBEzsignfoldertypeReassignezsignsigner() bool`

GetBEzsignfoldertypeReassignezsignsigner returns the BEzsignfoldertypeReassignezsignsigner field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeReassignezsignsignerOk

`func (o *EzsignfoldertypeResponse) GetBEzsignfoldertypeReassignezsignsignerOk() (*bool, bool)`

GetBEzsignfoldertypeReassignezsignsignerOk returns a tuple with the BEzsignfoldertypeReassignezsignsigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeReassignezsignsigner

`func (o *EzsignfoldertypeResponse) SetBEzsignfoldertypeReassignezsignsigner(v bool)`

SetBEzsignfoldertypeReassignezsignsigner sets BEzsignfoldertypeReassignezsignsigner field to given value.

### HasBEzsignfoldertypeReassignezsignsigner

`func (o *EzsignfoldertypeResponse) HasBEzsignfoldertypeReassignezsignsigner() bool`

HasBEzsignfoldertypeReassignezsignsigner returns a boolean if a field has been set.

### GetBEzsignfoldertypeReassignuser

`func (o *EzsignfoldertypeResponse) GetBEzsignfoldertypeReassignuser() bool`

GetBEzsignfoldertypeReassignuser returns the BEzsignfoldertypeReassignuser field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeReassignuserOk

`func (o *EzsignfoldertypeResponse) GetBEzsignfoldertypeReassignuserOk() (*bool, bool)`

GetBEzsignfoldertypeReassignuserOk returns a tuple with the BEzsignfoldertypeReassignuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeReassignuser

`func (o *EzsignfoldertypeResponse) SetBEzsignfoldertypeReassignuser(v bool)`

SetBEzsignfoldertypeReassignuser sets BEzsignfoldertypeReassignuser field to given value.

### HasBEzsignfoldertypeReassignuser

`func (o *EzsignfoldertypeResponse) HasBEzsignfoldertypeReassignuser() bool`

HasBEzsignfoldertypeReassignuser returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendsignedtoezsignsigner

`func (o *EzsignfoldertypeResponse) GetBEzsignfoldertypeSendsignedtoezsignsigner() bool`

GetBEzsignfoldertypeSendsignedtoezsignsigner returns the BEzsignfoldertypeSendsignedtoezsignsigner field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsignedtoezsignsignerOk

`func (o *EzsignfoldertypeResponse) GetBEzsignfoldertypeSendsignedtoezsignsignerOk() (*bool, bool)`

GetBEzsignfoldertypeSendsignedtoezsignsignerOk returns a tuple with the BEzsignfoldertypeSendsignedtoezsignsigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsignedtoezsignsigner

`func (o *EzsignfoldertypeResponse) SetBEzsignfoldertypeSendsignedtoezsignsigner(v bool)`

SetBEzsignfoldertypeSendsignedtoezsignsigner sets BEzsignfoldertypeSendsignedtoezsignsigner field to given value.

### HasBEzsignfoldertypeSendsignedtoezsignsigner

`func (o *EzsignfoldertypeResponse) HasBEzsignfoldertypeSendsignedtoezsignsigner() bool`

HasBEzsignfoldertypeSendsignedtoezsignsigner returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendsignedtouser

`func (o *EzsignfoldertypeResponse) GetBEzsignfoldertypeSendsignedtouser() bool`

GetBEzsignfoldertypeSendsignedtouser returns the BEzsignfoldertypeSendsignedtouser field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsignedtouserOk

`func (o *EzsignfoldertypeResponse) GetBEzsignfoldertypeSendsignedtouserOk() (*bool, bool)`

GetBEzsignfoldertypeSendsignedtouserOk returns a tuple with the BEzsignfoldertypeSendsignedtouser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsignedtouser

`func (o *EzsignfoldertypeResponse) SetBEzsignfoldertypeSendsignedtouser(v bool)`

SetBEzsignfoldertypeSendsignedtouser sets BEzsignfoldertypeSendsignedtouser field to given value.

### HasBEzsignfoldertypeSendsignedtouser

`func (o *EzsignfoldertypeResponse) HasBEzsignfoldertypeSendsignedtouser() bool`

HasBEzsignfoldertypeSendsignedtouser returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendattachmentezsignsigner

`func (o *EzsignfoldertypeResponse) GetBEzsignfoldertypeSendattachmentezsignsigner() bool`

GetBEzsignfoldertypeSendattachmentezsignsigner returns the BEzsignfoldertypeSendattachmentezsignsigner field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendattachmentezsignsignerOk

`func (o *EzsignfoldertypeResponse) GetBEzsignfoldertypeSendattachmentezsignsignerOk() (*bool, bool)`

GetBEzsignfoldertypeSendattachmentezsignsignerOk returns a tuple with the BEzsignfoldertypeSendattachmentezsignsigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendattachmentezsignsigner

`func (o *EzsignfoldertypeResponse) SetBEzsignfoldertypeSendattachmentezsignsigner(v bool)`

SetBEzsignfoldertypeSendattachmentezsignsigner sets BEzsignfoldertypeSendattachmentezsignsigner field to given value.

### HasBEzsignfoldertypeSendattachmentezsignsigner

`func (o *EzsignfoldertypeResponse) HasBEzsignfoldertypeSendattachmentezsignsigner() bool`

HasBEzsignfoldertypeSendattachmentezsignsigner returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendproofezsignsigner

`func (o *EzsignfoldertypeResponse) GetBEzsignfoldertypeSendproofezsignsigner() bool`

GetBEzsignfoldertypeSendproofezsignsigner returns the BEzsignfoldertypeSendproofezsignsigner field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendproofezsignsignerOk

`func (o *EzsignfoldertypeResponse) GetBEzsignfoldertypeSendproofezsignsignerOk() (*bool, bool)`

GetBEzsignfoldertypeSendproofezsignsignerOk returns a tuple with the BEzsignfoldertypeSendproofezsignsigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendproofezsignsigner

`func (o *EzsignfoldertypeResponse) SetBEzsignfoldertypeSendproofezsignsigner(v bool)`

SetBEzsignfoldertypeSendproofezsignsigner sets BEzsignfoldertypeSendproofezsignsigner field to given value.

### HasBEzsignfoldertypeSendproofezsignsigner

`func (o *EzsignfoldertypeResponse) HasBEzsignfoldertypeSendproofezsignsigner() bool`

HasBEzsignfoldertypeSendproofezsignsigner returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendattachmentuser

`func (o *EzsignfoldertypeResponse) GetBEzsignfoldertypeSendattachmentuser() bool`

GetBEzsignfoldertypeSendattachmentuser returns the BEzsignfoldertypeSendattachmentuser field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendattachmentuserOk

`func (o *EzsignfoldertypeResponse) GetBEzsignfoldertypeSendattachmentuserOk() (*bool, bool)`

GetBEzsignfoldertypeSendattachmentuserOk returns a tuple with the BEzsignfoldertypeSendattachmentuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendattachmentuser

`func (o *EzsignfoldertypeResponse) SetBEzsignfoldertypeSendattachmentuser(v bool)`

SetBEzsignfoldertypeSendattachmentuser sets BEzsignfoldertypeSendattachmentuser field to given value.

### HasBEzsignfoldertypeSendattachmentuser

`func (o *EzsignfoldertypeResponse) HasBEzsignfoldertypeSendattachmentuser() bool`

HasBEzsignfoldertypeSendattachmentuser returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendproofuser

`func (o *EzsignfoldertypeResponse) GetBEzsignfoldertypeSendproofuser() bool`

GetBEzsignfoldertypeSendproofuser returns the BEzsignfoldertypeSendproofuser field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendproofuserOk

`func (o *EzsignfoldertypeResponse) GetBEzsignfoldertypeSendproofuserOk() (*bool, bool)`

GetBEzsignfoldertypeSendproofuserOk returns a tuple with the BEzsignfoldertypeSendproofuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendproofuser

`func (o *EzsignfoldertypeResponse) SetBEzsignfoldertypeSendproofuser(v bool)`

SetBEzsignfoldertypeSendproofuser sets BEzsignfoldertypeSendproofuser field to given value.

### HasBEzsignfoldertypeSendproofuser

`func (o *EzsignfoldertypeResponse) HasBEzsignfoldertypeSendproofuser() bool`

HasBEzsignfoldertypeSendproofuser returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendproofemail

`func (o *EzsignfoldertypeResponse) GetBEzsignfoldertypeSendproofemail() bool`

GetBEzsignfoldertypeSendproofemail returns the BEzsignfoldertypeSendproofemail field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendproofemailOk

`func (o *EzsignfoldertypeResponse) GetBEzsignfoldertypeSendproofemailOk() (*bool, bool)`

GetBEzsignfoldertypeSendproofemailOk returns a tuple with the BEzsignfoldertypeSendproofemail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendproofemail

`func (o *EzsignfoldertypeResponse) SetBEzsignfoldertypeSendproofemail(v bool)`

SetBEzsignfoldertypeSendproofemail sets BEzsignfoldertypeSendproofemail field to given value.

### HasBEzsignfoldertypeSendproofemail

`func (o *EzsignfoldertypeResponse) HasBEzsignfoldertypeSendproofemail() bool`

HasBEzsignfoldertypeSendproofemail returns a boolean if a field has been set.

### GetBEzsignfoldertypeAllowdownloadattachmentezsignsigner

`func (o *EzsignfoldertypeResponse) GetBEzsignfoldertypeAllowdownloadattachmentezsignsigner() bool`

GetBEzsignfoldertypeAllowdownloadattachmentezsignsigner returns the BEzsignfoldertypeAllowdownloadattachmentezsignsigner field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeAllowdownloadattachmentezsignsignerOk

`func (o *EzsignfoldertypeResponse) GetBEzsignfoldertypeAllowdownloadattachmentezsignsignerOk() (*bool, bool)`

GetBEzsignfoldertypeAllowdownloadattachmentezsignsignerOk returns a tuple with the BEzsignfoldertypeAllowdownloadattachmentezsignsigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeAllowdownloadattachmentezsignsigner

`func (o *EzsignfoldertypeResponse) SetBEzsignfoldertypeAllowdownloadattachmentezsignsigner(v bool)`

SetBEzsignfoldertypeAllowdownloadattachmentezsignsigner sets BEzsignfoldertypeAllowdownloadattachmentezsignsigner field to given value.

### HasBEzsignfoldertypeAllowdownloadattachmentezsignsigner

`func (o *EzsignfoldertypeResponse) HasBEzsignfoldertypeAllowdownloadattachmentezsignsigner() bool`

HasBEzsignfoldertypeAllowdownloadattachmentezsignsigner returns a boolean if a field has been set.

### GetBEzsignfoldertypeAllowdownloadproofezsignsigner

`func (o *EzsignfoldertypeResponse) GetBEzsignfoldertypeAllowdownloadproofezsignsigner() bool`

GetBEzsignfoldertypeAllowdownloadproofezsignsigner returns the BEzsignfoldertypeAllowdownloadproofezsignsigner field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeAllowdownloadproofezsignsignerOk

`func (o *EzsignfoldertypeResponse) GetBEzsignfoldertypeAllowdownloadproofezsignsignerOk() (*bool, bool)`

GetBEzsignfoldertypeAllowdownloadproofezsignsignerOk returns a tuple with the BEzsignfoldertypeAllowdownloadproofezsignsigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeAllowdownloadproofezsignsigner

`func (o *EzsignfoldertypeResponse) SetBEzsignfoldertypeAllowdownloadproofezsignsigner(v bool)`

SetBEzsignfoldertypeAllowdownloadproofezsignsigner sets BEzsignfoldertypeAllowdownloadproofezsignsigner field to given value.

### HasBEzsignfoldertypeAllowdownloadproofezsignsigner

`func (o *EzsignfoldertypeResponse) HasBEzsignfoldertypeAllowdownloadproofezsignsigner() bool`

HasBEzsignfoldertypeAllowdownloadproofezsignsigner returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendproofreceivealldocument

`func (o *EzsignfoldertypeResponse) GetBEzsignfoldertypeSendproofreceivealldocument() bool`

GetBEzsignfoldertypeSendproofreceivealldocument returns the BEzsignfoldertypeSendproofreceivealldocument field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendproofreceivealldocumentOk

`func (o *EzsignfoldertypeResponse) GetBEzsignfoldertypeSendproofreceivealldocumentOk() (*bool, bool)`

GetBEzsignfoldertypeSendproofreceivealldocumentOk returns a tuple with the BEzsignfoldertypeSendproofreceivealldocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendproofreceivealldocument

`func (o *EzsignfoldertypeResponse) SetBEzsignfoldertypeSendproofreceivealldocument(v bool)`

SetBEzsignfoldertypeSendproofreceivealldocument sets BEzsignfoldertypeSendproofreceivealldocument field to given value.

### HasBEzsignfoldertypeSendproofreceivealldocument

`func (o *EzsignfoldertypeResponse) HasBEzsignfoldertypeSendproofreceivealldocument() bool`

HasBEzsignfoldertypeSendproofreceivealldocument returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendsignedtodocumentowner

`func (o *EzsignfoldertypeResponse) GetBEzsignfoldertypeSendsignedtodocumentowner() bool`

GetBEzsignfoldertypeSendsignedtodocumentowner returns the BEzsignfoldertypeSendsignedtodocumentowner field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsignedtodocumentownerOk

`func (o *EzsignfoldertypeResponse) GetBEzsignfoldertypeSendsignedtodocumentownerOk() (*bool, bool)`

GetBEzsignfoldertypeSendsignedtodocumentownerOk returns a tuple with the BEzsignfoldertypeSendsignedtodocumentowner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsignedtodocumentowner

`func (o *EzsignfoldertypeResponse) SetBEzsignfoldertypeSendsignedtodocumentowner(v bool)`

SetBEzsignfoldertypeSendsignedtodocumentowner sets BEzsignfoldertypeSendsignedtodocumentowner field to given value.


### GetBEzsignfoldertypeSendsignedtofolderowner

`func (o *EzsignfoldertypeResponse) GetBEzsignfoldertypeSendsignedtofolderowner() bool`

GetBEzsignfoldertypeSendsignedtofolderowner returns the BEzsignfoldertypeSendsignedtofolderowner field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsignedtofolderownerOk

`func (o *EzsignfoldertypeResponse) GetBEzsignfoldertypeSendsignedtofolderownerOk() (*bool, bool)`

GetBEzsignfoldertypeSendsignedtofolderownerOk returns a tuple with the BEzsignfoldertypeSendsignedtofolderowner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsignedtofolderowner

`func (o *EzsignfoldertypeResponse) SetBEzsignfoldertypeSendsignedtofolderowner(v bool)`

SetBEzsignfoldertypeSendsignedtofolderowner sets BEzsignfoldertypeSendsignedtofolderowner field to given value.


### GetBEzsignfoldertypeSendsignedtofullgroup

`func (o *EzsignfoldertypeResponse) GetBEzsignfoldertypeSendsignedtofullgroup() bool`

GetBEzsignfoldertypeSendsignedtofullgroup returns the BEzsignfoldertypeSendsignedtofullgroup field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsignedtofullgroupOk

`func (o *EzsignfoldertypeResponse) GetBEzsignfoldertypeSendsignedtofullgroupOk() (*bool, bool)`

GetBEzsignfoldertypeSendsignedtofullgroupOk returns a tuple with the BEzsignfoldertypeSendsignedtofullgroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsignedtofullgroup

`func (o *EzsignfoldertypeResponse) SetBEzsignfoldertypeSendsignedtofullgroup(v bool)`

SetBEzsignfoldertypeSendsignedtofullgroup sets BEzsignfoldertypeSendsignedtofullgroup field to given value.

### HasBEzsignfoldertypeSendsignedtofullgroup

`func (o *EzsignfoldertypeResponse) HasBEzsignfoldertypeSendsignedtofullgroup() bool`

HasBEzsignfoldertypeSendsignedtofullgroup returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendsignedtolimitedgroup

`func (o *EzsignfoldertypeResponse) GetBEzsignfoldertypeSendsignedtolimitedgroup() bool`

GetBEzsignfoldertypeSendsignedtolimitedgroup returns the BEzsignfoldertypeSendsignedtolimitedgroup field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsignedtolimitedgroupOk

`func (o *EzsignfoldertypeResponse) GetBEzsignfoldertypeSendsignedtolimitedgroupOk() (*bool, bool)`

GetBEzsignfoldertypeSendsignedtolimitedgroupOk returns a tuple with the BEzsignfoldertypeSendsignedtolimitedgroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsignedtolimitedgroup

`func (o *EzsignfoldertypeResponse) SetBEzsignfoldertypeSendsignedtolimitedgroup(v bool)`

SetBEzsignfoldertypeSendsignedtolimitedgroup sets BEzsignfoldertypeSendsignedtolimitedgroup field to given value.

### HasBEzsignfoldertypeSendsignedtolimitedgroup

`func (o *EzsignfoldertypeResponse) HasBEzsignfoldertypeSendsignedtolimitedgroup() bool`

HasBEzsignfoldertypeSendsignedtolimitedgroup returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendsignedtocolleague

`func (o *EzsignfoldertypeResponse) GetBEzsignfoldertypeSendsignedtocolleague() bool`

GetBEzsignfoldertypeSendsignedtocolleague returns the BEzsignfoldertypeSendsignedtocolleague field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsignedtocolleagueOk

`func (o *EzsignfoldertypeResponse) GetBEzsignfoldertypeSendsignedtocolleagueOk() (*bool, bool)`

GetBEzsignfoldertypeSendsignedtocolleagueOk returns a tuple with the BEzsignfoldertypeSendsignedtocolleague field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsignedtocolleague

`func (o *EzsignfoldertypeResponse) SetBEzsignfoldertypeSendsignedtocolleague(v bool)`

SetBEzsignfoldertypeSendsignedtocolleague sets BEzsignfoldertypeSendsignedtocolleague field to given value.


### GetBEzsignfoldertypeSendsummarytodocumentowner

`func (o *EzsignfoldertypeResponse) GetBEzsignfoldertypeSendsummarytodocumentowner() bool`

GetBEzsignfoldertypeSendsummarytodocumentowner returns the BEzsignfoldertypeSendsummarytodocumentowner field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsummarytodocumentownerOk

`func (o *EzsignfoldertypeResponse) GetBEzsignfoldertypeSendsummarytodocumentownerOk() (*bool, bool)`

GetBEzsignfoldertypeSendsummarytodocumentownerOk returns a tuple with the BEzsignfoldertypeSendsummarytodocumentowner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsummarytodocumentowner

`func (o *EzsignfoldertypeResponse) SetBEzsignfoldertypeSendsummarytodocumentowner(v bool)`

SetBEzsignfoldertypeSendsummarytodocumentowner sets BEzsignfoldertypeSendsummarytodocumentowner field to given value.


### GetBEzsignfoldertypeSendsummarytofolderowner

`func (o *EzsignfoldertypeResponse) GetBEzsignfoldertypeSendsummarytofolderowner() bool`

GetBEzsignfoldertypeSendsummarytofolderowner returns the BEzsignfoldertypeSendsummarytofolderowner field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsummarytofolderownerOk

`func (o *EzsignfoldertypeResponse) GetBEzsignfoldertypeSendsummarytofolderownerOk() (*bool, bool)`

GetBEzsignfoldertypeSendsummarytofolderownerOk returns a tuple with the BEzsignfoldertypeSendsummarytofolderowner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsummarytofolderowner

`func (o *EzsignfoldertypeResponse) SetBEzsignfoldertypeSendsummarytofolderowner(v bool)`

SetBEzsignfoldertypeSendsummarytofolderowner sets BEzsignfoldertypeSendsummarytofolderowner field to given value.


### GetBEzsignfoldertypeSendsummarytofullgroup

`func (o *EzsignfoldertypeResponse) GetBEzsignfoldertypeSendsummarytofullgroup() bool`

GetBEzsignfoldertypeSendsummarytofullgroup returns the BEzsignfoldertypeSendsummarytofullgroup field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsummarytofullgroupOk

`func (o *EzsignfoldertypeResponse) GetBEzsignfoldertypeSendsummarytofullgroupOk() (*bool, bool)`

GetBEzsignfoldertypeSendsummarytofullgroupOk returns a tuple with the BEzsignfoldertypeSendsummarytofullgroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsummarytofullgroup

`func (o *EzsignfoldertypeResponse) SetBEzsignfoldertypeSendsummarytofullgroup(v bool)`

SetBEzsignfoldertypeSendsummarytofullgroup sets BEzsignfoldertypeSendsummarytofullgroup field to given value.

### HasBEzsignfoldertypeSendsummarytofullgroup

`func (o *EzsignfoldertypeResponse) HasBEzsignfoldertypeSendsummarytofullgroup() bool`

HasBEzsignfoldertypeSendsummarytofullgroup returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendsummarytolimitedgroup

`func (o *EzsignfoldertypeResponse) GetBEzsignfoldertypeSendsummarytolimitedgroup() bool`

GetBEzsignfoldertypeSendsummarytolimitedgroup returns the BEzsignfoldertypeSendsummarytolimitedgroup field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsummarytolimitedgroupOk

`func (o *EzsignfoldertypeResponse) GetBEzsignfoldertypeSendsummarytolimitedgroupOk() (*bool, bool)`

GetBEzsignfoldertypeSendsummarytolimitedgroupOk returns a tuple with the BEzsignfoldertypeSendsummarytolimitedgroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsummarytolimitedgroup

`func (o *EzsignfoldertypeResponse) SetBEzsignfoldertypeSendsummarytolimitedgroup(v bool)`

SetBEzsignfoldertypeSendsummarytolimitedgroup sets BEzsignfoldertypeSendsummarytolimitedgroup field to given value.

### HasBEzsignfoldertypeSendsummarytolimitedgroup

`func (o *EzsignfoldertypeResponse) HasBEzsignfoldertypeSendsummarytolimitedgroup() bool`

HasBEzsignfoldertypeSendsummarytolimitedgroup returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendsummarytocolleague

`func (o *EzsignfoldertypeResponse) GetBEzsignfoldertypeSendsummarytocolleague() bool`

GetBEzsignfoldertypeSendsummarytocolleague returns the BEzsignfoldertypeSendsummarytocolleague field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsummarytocolleagueOk

`func (o *EzsignfoldertypeResponse) GetBEzsignfoldertypeSendsummarytocolleagueOk() (*bool, bool)`

GetBEzsignfoldertypeSendsummarytocolleagueOk returns a tuple with the BEzsignfoldertypeSendsummarytocolleague field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsummarytocolleague

`func (o *EzsignfoldertypeResponse) SetBEzsignfoldertypeSendsummarytocolleague(v bool)`

SetBEzsignfoldertypeSendsummarytocolleague sets BEzsignfoldertypeSendsummarytocolleague field to given value.


### GetBEzsignfoldertypeIsactive

`func (o *EzsignfoldertypeResponse) GetBEzsignfoldertypeIsactive() bool`

GetBEzsignfoldertypeIsactive returns the BEzsignfoldertypeIsactive field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeIsactiveOk

`func (o *EzsignfoldertypeResponse) GetBEzsignfoldertypeIsactiveOk() (*bool, bool)`

GetBEzsignfoldertypeIsactiveOk returns a tuple with the BEzsignfoldertypeIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeIsactive

`func (o *EzsignfoldertypeResponse) SetBEzsignfoldertypeIsactive(v bool)`

SetBEzsignfoldertypeIsactive sets BEzsignfoldertypeIsactive field to given value.


### GetAObjUserlogintype

`func (o *EzsignfoldertypeResponse) GetAObjUserlogintype() []UserlogintypeResponse`

GetAObjUserlogintype returns the AObjUserlogintype field if non-nil, zero value otherwise.

### GetAObjUserlogintypeOk

`func (o *EzsignfoldertypeResponse) GetAObjUserlogintypeOk() (*[]UserlogintypeResponse, bool)`

GetAObjUserlogintypeOk returns a tuple with the AObjUserlogintype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjUserlogintype

`func (o *EzsignfoldertypeResponse) SetAObjUserlogintype(v []UserlogintypeResponse)`

SetAObjUserlogintype sets AObjUserlogintype field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


