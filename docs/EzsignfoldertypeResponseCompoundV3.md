# EzsignfoldertypeResponseCompoundV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsignfoldertypeID** | **int32** | The unique ID of the Ezsignfoldertype. | 
**ObjEzsignfoldertypeName** | [**MultilingualEzsignfoldertypeName**](MultilingualEzsignfoldertypeName.md) |  | 
**FkiBrandingID** | **int32** | The unique ID of the Branding | 
**FkiBillingentityinternalID** | Pointer to **int32** | The unique ID of the Billingentityinternal. | [optional] 
**FkiEzsigntsarequirementID** | Pointer to **int32** | The unique ID of the Ezsigntsarequirement.  Determine if a Time Stamping Authority should add a timestamp on each of the signature. Valid values:  |Value|Description| |-|-| |1|No. TSA Timestamping will requested. This will make all signatures a lot faster since no round-trip to the TSA server will be required. Timestamping will be made using eZsign server&#39;s time.| |2|Best effort. Timestamping from a Time Stamping Authority will be requested but is not mandatory. In the very improbable case it cannot be completed, the timestamping will be made using eZsign server&#39;s time. **Additional fee applies**| |3|Mandatory. Timestamping from a Time Stamping Authority will be requested and is mandatory. In the very improbable case it cannot be completed, the signature will fail and the user will be asked to retry. **Additional fee applies**| | [optional] 
**SBrandingDescriptionX** | **string** | The Description of the Branding in the language of the requester | 
**SBillingentityinternalDescriptionX** | Pointer to **string** | The description of the Billingentityinternal in the language of the requester | [optional] 
**SEzsigntsarequirementDescriptionX** | Pointer to **string** | The description of the Ezsigntsarequirement in the language of the requester | [optional] 
**SEmailAddressSigned** | Pointer to **string** | The email address. | [optional] 
**SEmailAddressSummary** | Pointer to **string** | The email address. | [optional] 
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
**AObjUsergroupAll** | Pointer to [**[]UsergroupResponse**](UsergroupResponse.md) |  | [optional] 
**AObjUsergroupRestricted** | Pointer to [**[]UsergroupResponse**](UsergroupResponse.md) |  | [optional] 
**AObjUsergroupTemplate** | Pointer to [**[]UsergroupResponse**](UsergroupResponse.md) |  | [optional] 
**AFkiUserIDSigned** | Pointer to **[]int32** |  | [optional] 
**AFkiUserIDSummary** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewEzsignfoldertypeResponseCompoundV3

`func NewEzsignfoldertypeResponseCompoundV3(pkiEzsignfoldertypeID int32, objEzsignfoldertypeName MultilingualEzsignfoldertypeName, fkiBrandingID int32, sBrandingDescriptionX string, eEzsignfoldertypePrivacylevel FieldEEzsignfoldertypePrivacylevel, iEzsignfoldertypeArchivaldays int32, eEzsignfoldertypeDisposal FieldEEzsignfoldertypeDisposal, eEzsignfoldertypeCompletion FieldEEzsignfoldertypeCompletion, iEzsignfoldertypeDeadlinedays int32, bEzsignfoldertypeSendsignedtodocumentowner bool, bEzsignfoldertypeSendsignedtofolderowner bool, bEzsignfoldertypeSendsignedtocolleague bool, bEzsignfoldertypeSendsummarytodocumentowner bool, bEzsignfoldertypeSendsummarytofolderowner bool, bEzsignfoldertypeSendsummarytocolleague bool, bEzsignfoldertypeIsactive bool, aObjUserlogintype []UserlogintypeResponse, ) *EzsignfoldertypeResponseCompoundV3`

NewEzsignfoldertypeResponseCompoundV3 instantiates a new EzsignfoldertypeResponseCompoundV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignfoldertypeResponseCompoundV3WithDefaults

`func NewEzsignfoldertypeResponseCompoundV3WithDefaults() *EzsignfoldertypeResponseCompoundV3`

NewEzsignfoldertypeResponseCompoundV3WithDefaults instantiates a new EzsignfoldertypeResponseCompoundV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsignfoldertypeID

`func (o *EzsignfoldertypeResponseCompoundV3) GetPkiEzsignfoldertypeID() int32`

GetPkiEzsignfoldertypeID returns the PkiEzsignfoldertypeID field if non-nil, zero value otherwise.

### GetPkiEzsignfoldertypeIDOk

`func (o *EzsignfoldertypeResponseCompoundV3) GetPkiEzsignfoldertypeIDOk() (*int32, bool)`

GetPkiEzsignfoldertypeIDOk returns a tuple with the PkiEzsignfoldertypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsignfoldertypeID

`func (o *EzsignfoldertypeResponseCompoundV3) SetPkiEzsignfoldertypeID(v int32)`

SetPkiEzsignfoldertypeID sets PkiEzsignfoldertypeID field to given value.


### GetObjEzsignfoldertypeName

`func (o *EzsignfoldertypeResponseCompoundV3) GetObjEzsignfoldertypeName() MultilingualEzsignfoldertypeName`

GetObjEzsignfoldertypeName returns the ObjEzsignfoldertypeName field if non-nil, zero value otherwise.

### GetObjEzsignfoldertypeNameOk

`func (o *EzsignfoldertypeResponseCompoundV3) GetObjEzsignfoldertypeNameOk() (*MultilingualEzsignfoldertypeName, bool)`

GetObjEzsignfoldertypeNameOk returns a tuple with the ObjEzsignfoldertypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjEzsignfoldertypeName

`func (o *EzsignfoldertypeResponseCompoundV3) SetObjEzsignfoldertypeName(v MultilingualEzsignfoldertypeName)`

SetObjEzsignfoldertypeName sets ObjEzsignfoldertypeName field to given value.


### GetFkiBrandingID

`func (o *EzsignfoldertypeResponseCompoundV3) GetFkiBrandingID() int32`

GetFkiBrandingID returns the FkiBrandingID field if non-nil, zero value otherwise.

### GetFkiBrandingIDOk

`func (o *EzsignfoldertypeResponseCompoundV3) GetFkiBrandingIDOk() (*int32, bool)`

GetFkiBrandingIDOk returns a tuple with the FkiBrandingID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiBrandingID

`func (o *EzsignfoldertypeResponseCompoundV3) SetFkiBrandingID(v int32)`

SetFkiBrandingID sets FkiBrandingID field to given value.


### GetFkiBillingentityinternalID

`func (o *EzsignfoldertypeResponseCompoundV3) GetFkiBillingentityinternalID() int32`

GetFkiBillingentityinternalID returns the FkiBillingentityinternalID field if non-nil, zero value otherwise.

### GetFkiBillingentityinternalIDOk

`func (o *EzsignfoldertypeResponseCompoundV3) GetFkiBillingentityinternalIDOk() (*int32, bool)`

GetFkiBillingentityinternalIDOk returns a tuple with the FkiBillingentityinternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiBillingentityinternalID

`func (o *EzsignfoldertypeResponseCompoundV3) SetFkiBillingentityinternalID(v int32)`

SetFkiBillingentityinternalID sets FkiBillingentityinternalID field to given value.

### HasFkiBillingentityinternalID

`func (o *EzsignfoldertypeResponseCompoundV3) HasFkiBillingentityinternalID() bool`

HasFkiBillingentityinternalID returns a boolean if a field has been set.

### GetFkiEzsigntsarequirementID

`func (o *EzsignfoldertypeResponseCompoundV3) GetFkiEzsigntsarequirementID() int32`

GetFkiEzsigntsarequirementID returns the FkiEzsigntsarequirementID field if non-nil, zero value otherwise.

### GetFkiEzsigntsarequirementIDOk

`func (o *EzsignfoldertypeResponseCompoundV3) GetFkiEzsigntsarequirementIDOk() (*int32, bool)`

GetFkiEzsigntsarequirementIDOk returns a tuple with the FkiEzsigntsarequirementID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntsarequirementID

`func (o *EzsignfoldertypeResponseCompoundV3) SetFkiEzsigntsarequirementID(v int32)`

SetFkiEzsigntsarequirementID sets FkiEzsigntsarequirementID field to given value.

### HasFkiEzsigntsarequirementID

`func (o *EzsignfoldertypeResponseCompoundV3) HasFkiEzsigntsarequirementID() bool`

HasFkiEzsigntsarequirementID returns a boolean if a field has been set.

### GetSBrandingDescriptionX

`func (o *EzsignfoldertypeResponseCompoundV3) GetSBrandingDescriptionX() string`

GetSBrandingDescriptionX returns the SBrandingDescriptionX field if non-nil, zero value otherwise.

### GetSBrandingDescriptionXOk

`func (o *EzsignfoldertypeResponseCompoundV3) GetSBrandingDescriptionXOk() (*string, bool)`

GetSBrandingDescriptionXOk returns a tuple with the SBrandingDescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSBrandingDescriptionX

`func (o *EzsignfoldertypeResponseCompoundV3) SetSBrandingDescriptionX(v string)`

SetSBrandingDescriptionX sets SBrandingDescriptionX field to given value.


### GetSBillingentityinternalDescriptionX

`func (o *EzsignfoldertypeResponseCompoundV3) GetSBillingentityinternalDescriptionX() string`

GetSBillingentityinternalDescriptionX returns the SBillingentityinternalDescriptionX field if non-nil, zero value otherwise.

### GetSBillingentityinternalDescriptionXOk

`func (o *EzsignfoldertypeResponseCompoundV3) GetSBillingentityinternalDescriptionXOk() (*string, bool)`

GetSBillingentityinternalDescriptionXOk returns a tuple with the SBillingentityinternalDescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSBillingentityinternalDescriptionX

`func (o *EzsignfoldertypeResponseCompoundV3) SetSBillingentityinternalDescriptionX(v string)`

SetSBillingentityinternalDescriptionX sets SBillingentityinternalDescriptionX field to given value.

### HasSBillingentityinternalDescriptionX

`func (o *EzsignfoldertypeResponseCompoundV3) HasSBillingentityinternalDescriptionX() bool`

HasSBillingentityinternalDescriptionX returns a boolean if a field has been set.

### GetSEzsigntsarequirementDescriptionX

`func (o *EzsignfoldertypeResponseCompoundV3) GetSEzsigntsarequirementDescriptionX() string`

GetSEzsigntsarequirementDescriptionX returns the SEzsigntsarequirementDescriptionX field if non-nil, zero value otherwise.

### GetSEzsigntsarequirementDescriptionXOk

`func (o *EzsignfoldertypeResponseCompoundV3) GetSEzsigntsarequirementDescriptionXOk() (*string, bool)`

GetSEzsigntsarequirementDescriptionXOk returns a tuple with the SEzsigntsarequirementDescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntsarequirementDescriptionX

`func (o *EzsignfoldertypeResponseCompoundV3) SetSEzsigntsarequirementDescriptionX(v string)`

SetSEzsigntsarequirementDescriptionX sets SEzsigntsarequirementDescriptionX field to given value.

### HasSEzsigntsarequirementDescriptionX

`func (o *EzsignfoldertypeResponseCompoundV3) HasSEzsigntsarequirementDescriptionX() bool`

HasSEzsigntsarequirementDescriptionX returns a boolean if a field has been set.

### GetSEmailAddressSigned

`func (o *EzsignfoldertypeResponseCompoundV3) GetSEmailAddressSigned() string`

GetSEmailAddressSigned returns the SEmailAddressSigned field if non-nil, zero value otherwise.

### GetSEmailAddressSignedOk

`func (o *EzsignfoldertypeResponseCompoundV3) GetSEmailAddressSignedOk() (*string, bool)`

GetSEmailAddressSignedOk returns a tuple with the SEmailAddressSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEmailAddressSigned

`func (o *EzsignfoldertypeResponseCompoundV3) SetSEmailAddressSigned(v string)`

SetSEmailAddressSigned sets SEmailAddressSigned field to given value.

### HasSEmailAddressSigned

`func (o *EzsignfoldertypeResponseCompoundV3) HasSEmailAddressSigned() bool`

HasSEmailAddressSigned returns a boolean if a field has been set.

### GetSEmailAddressSummary

`func (o *EzsignfoldertypeResponseCompoundV3) GetSEmailAddressSummary() string`

GetSEmailAddressSummary returns the SEmailAddressSummary field if non-nil, zero value otherwise.

### GetSEmailAddressSummaryOk

`func (o *EzsignfoldertypeResponseCompoundV3) GetSEmailAddressSummaryOk() (*string, bool)`

GetSEmailAddressSummaryOk returns a tuple with the SEmailAddressSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEmailAddressSummary

`func (o *EzsignfoldertypeResponseCompoundV3) SetSEmailAddressSummary(v string)`

SetSEmailAddressSummary sets SEmailAddressSummary field to given value.

### HasSEmailAddressSummary

`func (o *EzsignfoldertypeResponseCompoundV3) HasSEmailAddressSummary() bool`

HasSEmailAddressSummary returns a boolean if a field has been set.

### GetEEzsignfoldertypePrivacylevel

`func (o *EzsignfoldertypeResponseCompoundV3) GetEEzsignfoldertypePrivacylevel() FieldEEzsignfoldertypePrivacylevel`

GetEEzsignfoldertypePrivacylevel returns the EEzsignfoldertypePrivacylevel field if non-nil, zero value otherwise.

### GetEEzsignfoldertypePrivacylevelOk

`func (o *EzsignfoldertypeResponseCompoundV3) GetEEzsignfoldertypePrivacylevelOk() (*FieldEEzsignfoldertypePrivacylevel, bool)`

GetEEzsignfoldertypePrivacylevelOk returns a tuple with the EEzsignfoldertypePrivacylevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignfoldertypePrivacylevel

`func (o *EzsignfoldertypeResponseCompoundV3) SetEEzsignfoldertypePrivacylevel(v FieldEEzsignfoldertypePrivacylevel)`

SetEEzsignfoldertypePrivacylevel sets EEzsignfoldertypePrivacylevel field to given value.


### GetEEzsignfoldertypeSendreminderfrequency

`func (o *EzsignfoldertypeResponseCompoundV3) GetEEzsignfoldertypeSendreminderfrequency() FieldEEzsignfoldertypeSendreminderfrequency`

GetEEzsignfoldertypeSendreminderfrequency returns the EEzsignfoldertypeSendreminderfrequency field if non-nil, zero value otherwise.

### GetEEzsignfoldertypeSendreminderfrequencyOk

`func (o *EzsignfoldertypeResponseCompoundV3) GetEEzsignfoldertypeSendreminderfrequencyOk() (*FieldEEzsignfoldertypeSendreminderfrequency, bool)`

GetEEzsignfoldertypeSendreminderfrequencyOk returns a tuple with the EEzsignfoldertypeSendreminderfrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignfoldertypeSendreminderfrequency

`func (o *EzsignfoldertypeResponseCompoundV3) SetEEzsignfoldertypeSendreminderfrequency(v FieldEEzsignfoldertypeSendreminderfrequency)`

SetEEzsignfoldertypeSendreminderfrequency sets EEzsignfoldertypeSendreminderfrequency field to given value.

### HasEEzsignfoldertypeSendreminderfrequency

`func (o *EzsignfoldertypeResponseCompoundV3) HasEEzsignfoldertypeSendreminderfrequency() bool`

HasEEzsignfoldertypeSendreminderfrequency returns a boolean if a field has been set.

### GetIEzsignfoldertypeArchivaldays

`func (o *EzsignfoldertypeResponseCompoundV3) GetIEzsignfoldertypeArchivaldays() int32`

GetIEzsignfoldertypeArchivaldays returns the IEzsignfoldertypeArchivaldays field if non-nil, zero value otherwise.

### GetIEzsignfoldertypeArchivaldaysOk

`func (o *EzsignfoldertypeResponseCompoundV3) GetIEzsignfoldertypeArchivaldaysOk() (*int32, bool)`

GetIEzsignfoldertypeArchivaldaysOk returns a tuple with the IEzsignfoldertypeArchivaldays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignfoldertypeArchivaldays

`func (o *EzsignfoldertypeResponseCompoundV3) SetIEzsignfoldertypeArchivaldays(v int32)`

SetIEzsignfoldertypeArchivaldays sets IEzsignfoldertypeArchivaldays field to given value.


### GetEEzsignfoldertypeDisposal

`func (o *EzsignfoldertypeResponseCompoundV3) GetEEzsignfoldertypeDisposal() FieldEEzsignfoldertypeDisposal`

GetEEzsignfoldertypeDisposal returns the EEzsignfoldertypeDisposal field if non-nil, zero value otherwise.

### GetEEzsignfoldertypeDisposalOk

`func (o *EzsignfoldertypeResponseCompoundV3) GetEEzsignfoldertypeDisposalOk() (*FieldEEzsignfoldertypeDisposal, bool)`

GetEEzsignfoldertypeDisposalOk returns a tuple with the EEzsignfoldertypeDisposal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignfoldertypeDisposal

`func (o *EzsignfoldertypeResponseCompoundV3) SetEEzsignfoldertypeDisposal(v FieldEEzsignfoldertypeDisposal)`

SetEEzsignfoldertypeDisposal sets EEzsignfoldertypeDisposal field to given value.


### GetEEzsignfoldertypeCompletion

`func (o *EzsignfoldertypeResponseCompoundV3) GetEEzsignfoldertypeCompletion() FieldEEzsignfoldertypeCompletion`

GetEEzsignfoldertypeCompletion returns the EEzsignfoldertypeCompletion field if non-nil, zero value otherwise.

### GetEEzsignfoldertypeCompletionOk

`func (o *EzsignfoldertypeResponseCompoundV3) GetEEzsignfoldertypeCompletionOk() (*FieldEEzsignfoldertypeCompletion, bool)`

GetEEzsignfoldertypeCompletionOk returns a tuple with the EEzsignfoldertypeCompletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignfoldertypeCompletion

`func (o *EzsignfoldertypeResponseCompoundV3) SetEEzsignfoldertypeCompletion(v FieldEEzsignfoldertypeCompletion)`

SetEEzsignfoldertypeCompletion sets EEzsignfoldertypeCompletion field to given value.


### GetIEzsignfoldertypeDisposaldays

`func (o *EzsignfoldertypeResponseCompoundV3) GetIEzsignfoldertypeDisposaldays() int32`

GetIEzsignfoldertypeDisposaldays returns the IEzsignfoldertypeDisposaldays field if non-nil, zero value otherwise.

### GetIEzsignfoldertypeDisposaldaysOk

`func (o *EzsignfoldertypeResponseCompoundV3) GetIEzsignfoldertypeDisposaldaysOk() (*int32, bool)`

GetIEzsignfoldertypeDisposaldaysOk returns a tuple with the IEzsignfoldertypeDisposaldays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignfoldertypeDisposaldays

`func (o *EzsignfoldertypeResponseCompoundV3) SetIEzsignfoldertypeDisposaldays(v int32)`

SetIEzsignfoldertypeDisposaldays sets IEzsignfoldertypeDisposaldays field to given value.

### HasIEzsignfoldertypeDisposaldays

`func (o *EzsignfoldertypeResponseCompoundV3) HasIEzsignfoldertypeDisposaldays() bool`

HasIEzsignfoldertypeDisposaldays returns a boolean if a field has been set.

### GetIEzsignfoldertypeDeadlinedays

`func (o *EzsignfoldertypeResponseCompoundV3) GetIEzsignfoldertypeDeadlinedays() int32`

GetIEzsignfoldertypeDeadlinedays returns the IEzsignfoldertypeDeadlinedays field if non-nil, zero value otherwise.

### GetIEzsignfoldertypeDeadlinedaysOk

`func (o *EzsignfoldertypeResponseCompoundV3) GetIEzsignfoldertypeDeadlinedaysOk() (*int32, bool)`

GetIEzsignfoldertypeDeadlinedaysOk returns a tuple with the IEzsignfoldertypeDeadlinedays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignfoldertypeDeadlinedays

`func (o *EzsignfoldertypeResponseCompoundV3) SetIEzsignfoldertypeDeadlinedays(v int32)`

SetIEzsignfoldertypeDeadlinedays sets IEzsignfoldertypeDeadlinedays field to given value.


### GetBEzsignfoldertypeDelegate

`func (o *EzsignfoldertypeResponseCompoundV3) GetBEzsignfoldertypeDelegate() bool`

GetBEzsignfoldertypeDelegate returns the BEzsignfoldertypeDelegate field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeDelegateOk

`func (o *EzsignfoldertypeResponseCompoundV3) GetBEzsignfoldertypeDelegateOk() (*bool, bool)`

GetBEzsignfoldertypeDelegateOk returns a tuple with the BEzsignfoldertypeDelegate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeDelegate

`func (o *EzsignfoldertypeResponseCompoundV3) SetBEzsignfoldertypeDelegate(v bool)`

SetBEzsignfoldertypeDelegate sets BEzsignfoldertypeDelegate field to given value.

### HasBEzsignfoldertypeDelegate

`func (o *EzsignfoldertypeResponseCompoundV3) HasBEzsignfoldertypeDelegate() bool`

HasBEzsignfoldertypeDelegate returns a boolean if a field has been set.

### GetBEzsignfoldertypeDiscussion

`func (o *EzsignfoldertypeResponseCompoundV3) GetBEzsignfoldertypeDiscussion() bool`

GetBEzsignfoldertypeDiscussion returns the BEzsignfoldertypeDiscussion field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeDiscussionOk

`func (o *EzsignfoldertypeResponseCompoundV3) GetBEzsignfoldertypeDiscussionOk() (*bool, bool)`

GetBEzsignfoldertypeDiscussionOk returns a tuple with the BEzsignfoldertypeDiscussion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeDiscussion

`func (o *EzsignfoldertypeResponseCompoundV3) SetBEzsignfoldertypeDiscussion(v bool)`

SetBEzsignfoldertypeDiscussion sets BEzsignfoldertypeDiscussion field to given value.

### HasBEzsignfoldertypeDiscussion

`func (o *EzsignfoldertypeResponseCompoundV3) HasBEzsignfoldertypeDiscussion() bool`

HasBEzsignfoldertypeDiscussion returns a boolean if a field has been set.

### GetBEzsignfoldertypeReassignezsignsigner

`func (o *EzsignfoldertypeResponseCompoundV3) GetBEzsignfoldertypeReassignezsignsigner() bool`

GetBEzsignfoldertypeReassignezsignsigner returns the BEzsignfoldertypeReassignezsignsigner field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeReassignezsignsignerOk

`func (o *EzsignfoldertypeResponseCompoundV3) GetBEzsignfoldertypeReassignezsignsignerOk() (*bool, bool)`

GetBEzsignfoldertypeReassignezsignsignerOk returns a tuple with the BEzsignfoldertypeReassignezsignsigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeReassignezsignsigner

`func (o *EzsignfoldertypeResponseCompoundV3) SetBEzsignfoldertypeReassignezsignsigner(v bool)`

SetBEzsignfoldertypeReassignezsignsigner sets BEzsignfoldertypeReassignezsignsigner field to given value.

### HasBEzsignfoldertypeReassignezsignsigner

`func (o *EzsignfoldertypeResponseCompoundV3) HasBEzsignfoldertypeReassignezsignsigner() bool`

HasBEzsignfoldertypeReassignezsignsigner returns a boolean if a field has been set.

### GetBEzsignfoldertypeReassignuser

`func (o *EzsignfoldertypeResponseCompoundV3) GetBEzsignfoldertypeReassignuser() bool`

GetBEzsignfoldertypeReassignuser returns the BEzsignfoldertypeReassignuser field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeReassignuserOk

`func (o *EzsignfoldertypeResponseCompoundV3) GetBEzsignfoldertypeReassignuserOk() (*bool, bool)`

GetBEzsignfoldertypeReassignuserOk returns a tuple with the BEzsignfoldertypeReassignuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeReassignuser

`func (o *EzsignfoldertypeResponseCompoundV3) SetBEzsignfoldertypeReassignuser(v bool)`

SetBEzsignfoldertypeReassignuser sets BEzsignfoldertypeReassignuser field to given value.

### HasBEzsignfoldertypeReassignuser

`func (o *EzsignfoldertypeResponseCompoundV3) HasBEzsignfoldertypeReassignuser() bool`

HasBEzsignfoldertypeReassignuser returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendsignedtoezsignsigner

`func (o *EzsignfoldertypeResponseCompoundV3) GetBEzsignfoldertypeSendsignedtoezsignsigner() bool`

GetBEzsignfoldertypeSendsignedtoezsignsigner returns the BEzsignfoldertypeSendsignedtoezsignsigner field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsignedtoezsignsignerOk

`func (o *EzsignfoldertypeResponseCompoundV3) GetBEzsignfoldertypeSendsignedtoezsignsignerOk() (*bool, bool)`

GetBEzsignfoldertypeSendsignedtoezsignsignerOk returns a tuple with the BEzsignfoldertypeSendsignedtoezsignsigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsignedtoezsignsigner

`func (o *EzsignfoldertypeResponseCompoundV3) SetBEzsignfoldertypeSendsignedtoezsignsigner(v bool)`

SetBEzsignfoldertypeSendsignedtoezsignsigner sets BEzsignfoldertypeSendsignedtoezsignsigner field to given value.

### HasBEzsignfoldertypeSendsignedtoezsignsigner

`func (o *EzsignfoldertypeResponseCompoundV3) HasBEzsignfoldertypeSendsignedtoezsignsigner() bool`

HasBEzsignfoldertypeSendsignedtoezsignsigner returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendsignedtouser

`func (o *EzsignfoldertypeResponseCompoundV3) GetBEzsignfoldertypeSendsignedtouser() bool`

GetBEzsignfoldertypeSendsignedtouser returns the BEzsignfoldertypeSendsignedtouser field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsignedtouserOk

`func (o *EzsignfoldertypeResponseCompoundV3) GetBEzsignfoldertypeSendsignedtouserOk() (*bool, bool)`

GetBEzsignfoldertypeSendsignedtouserOk returns a tuple with the BEzsignfoldertypeSendsignedtouser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsignedtouser

`func (o *EzsignfoldertypeResponseCompoundV3) SetBEzsignfoldertypeSendsignedtouser(v bool)`

SetBEzsignfoldertypeSendsignedtouser sets BEzsignfoldertypeSendsignedtouser field to given value.

### HasBEzsignfoldertypeSendsignedtouser

`func (o *EzsignfoldertypeResponseCompoundV3) HasBEzsignfoldertypeSendsignedtouser() bool`

HasBEzsignfoldertypeSendsignedtouser returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendattachmentezsignsigner

`func (o *EzsignfoldertypeResponseCompoundV3) GetBEzsignfoldertypeSendattachmentezsignsigner() bool`

GetBEzsignfoldertypeSendattachmentezsignsigner returns the BEzsignfoldertypeSendattachmentezsignsigner field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendattachmentezsignsignerOk

`func (o *EzsignfoldertypeResponseCompoundV3) GetBEzsignfoldertypeSendattachmentezsignsignerOk() (*bool, bool)`

GetBEzsignfoldertypeSendattachmentezsignsignerOk returns a tuple with the BEzsignfoldertypeSendattachmentezsignsigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendattachmentezsignsigner

`func (o *EzsignfoldertypeResponseCompoundV3) SetBEzsignfoldertypeSendattachmentezsignsigner(v bool)`

SetBEzsignfoldertypeSendattachmentezsignsigner sets BEzsignfoldertypeSendattachmentezsignsigner field to given value.

### HasBEzsignfoldertypeSendattachmentezsignsigner

`func (o *EzsignfoldertypeResponseCompoundV3) HasBEzsignfoldertypeSendattachmentezsignsigner() bool`

HasBEzsignfoldertypeSendattachmentezsignsigner returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendproofezsignsigner

`func (o *EzsignfoldertypeResponseCompoundV3) GetBEzsignfoldertypeSendproofezsignsigner() bool`

GetBEzsignfoldertypeSendproofezsignsigner returns the BEzsignfoldertypeSendproofezsignsigner field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendproofezsignsignerOk

`func (o *EzsignfoldertypeResponseCompoundV3) GetBEzsignfoldertypeSendproofezsignsignerOk() (*bool, bool)`

GetBEzsignfoldertypeSendproofezsignsignerOk returns a tuple with the BEzsignfoldertypeSendproofezsignsigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendproofezsignsigner

`func (o *EzsignfoldertypeResponseCompoundV3) SetBEzsignfoldertypeSendproofezsignsigner(v bool)`

SetBEzsignfoldertypeSendproofezsignsigner sets BEzsignfoldertypeSendproofezsignsigner field to given value.

### HasBEzsignfoldertypeSendproofezsignsigner

`func (o *EzsignfoldertypeResponseCompoundV3) HasBEzsignfoldertypeSendproofezsignsigner() bool`

HasBEzsignfoldertypeSendproofezsignsigner returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendattachmentuser

`func (o *EzsignfoldertypeResponseCompoundV3) GetBEzsignfoldertypeSendattachmentuser() bool`

GetBEzsignfoldertypeSendattachmentuser returns the BEzsignfoldertypeSendattachmentuser field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendattachmentuserOk

`func (o *EzsignfoldertypeResponseCompoundV3) GetBEzsignfoldertypeSendattachmentuserOk() (*bool, bool)`

GetBEzsignfoldertypeSendattachmentuserOk returns a tuple with the BEzsignfoldertypeSendattachmentuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendattachmentuser

`func (o *EzsignfoldertypeResponseCompoundV3) SetBEzsignfoldertypeSendattachmentuser(v bool)`

SetBEzsignfoldertypeSendattachmentuser sets BEzsignfoldertypeSendattachmentuser field to given value.

### HasBEzsignfoldertypeSendattachmentuser

`func (o *EzsignfoldertypeResponseCompoundV3) HasBEzsignfoldertypeSendattachmentuser() bool`

HasBEzsignfoldertypeSendattachmentuser returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendproofuser

`func (o *EzsignfoldertypeResponseCompoundV3) GetBEzsignfoldertypeSendproofuser() bool`

GetBEzsignfoldertypeSendproofuser returns the BEzsignfoldertypeSendproofuser field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendproofuserOk

`func (o *EzsignfoldertypeResponseCompoundV3) GetBEzsignfoldertypeSendproofuserOk() (*bool, bool)`

GetBEzsignfoldertypeSendproofuserOk returns a tuple with the BEzsignfoldertypeSendproofuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendproofuser

`func (o *EzsignfoldertypeResponseCompoundV3) SetBEzsignfoldertypeSendproofuser(v bool)`

SetBEzsignfoldertypeSendproofuser sets BEzsignfoldertypeSendproofuser field to given value.

### HasBEzsignfoldertypeSendproofuser

`func (o *EzsignfoldertypeResponseCompoundV3) HasBEzsignfoldertypeSendproofuser() bool`

HasBEzsignfoldertypeSendproofuser returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendproofemail

`func (o *EzsignfoldertypeResponseCompoundV3) GetBEzsignfoldertypeSendproofemail() bool`

GetBEzsignfoldertypeSendproofemail returns the BEzsignfoldertypeSendproofemail field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendproofemailOk

`func (o *EzsignfoldertypeResponseCompoundV3) GetBEzsignfoldertypeSendproofemailOk() (*bool, bool)`

GetBEzsignfoldertypeSendproofemailOk returns a tuple with the BEzsignfoldertypeSendproofemail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendproofemail

`func (o *EzsignfoldertypeResponseCompoundV3) SetBEzsignfoldertypeSendproofemail(v bool)`

SetBEzsignfoldertypeSendproofemail sets BEzsignfoldertypeSendproofemail field to given value.

### HasBEzsignfoldertypeSendproofemail

`func (o *EzsignfoldertypeResponseCompoundV3) HasBEzsignfoldertypeSendproofemail() bool`

HasBEzsignfoldertypeSendproofemail returns a boolean if a field has been set.

### GetBEzsignfoldertypeAllowdownloadattachmentezsignsigner

`func (o *EzsignfoldertypeResponseCompoundV3) GetBEzsignfoldertypeAllowdownloadattachmentezsignsigner() bool`

GetBEzsignfoldertypeAllowdownloadattachmentezsignsigner returns the BEzsignfoldertypeAllowdownloadattachmentezsignsigner field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeAllowdownloadattachmentezsignsignerOk

`func (o *EzsignfoldertypeResponseCompoundV3) GetBEzsignfoldertypeAllowdownloadattachmentezsignsignerOk() (*bool, bool)`

GetBEzsignfoldertypeAllowdownloadattachmentezsignsignerOk returns a tuple with the BEzsignfoldertypeAllowdownloadattachmentezsignsigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeAllowdownloadattachmentezsignsigner

`func (o *EzsignfoldertypeResponseCompoundV3) SetBEzsignfoldertypeAllowdownloadattachmentezsignsigner(v bool)`

SetBEzsignfoldertypeAllowdownloadattachmentezsignsigner sets BEzsignfoldertypeAllowdownloadattachmentezsignsigner field to given value.

### HasBEzsignfoldertypeAllowdownloadattachmentezsignsigner

`func (o *EzsignfoldertypeResponseCompoundV3) HasBEzsignfoldertypeAllowdownloadattachmentezsignsigner() bool`

HasBEzsignfoldertypeAllowdownloadattachmentezsignsigner returns a boolean if a field has been set.

### GetBEzsignfoldertypeAllowdownloadproofezsignsigner

`func (o *EzsignfoldertypeResponseCompoundV3) GetBEzsignfoldertypeAllowdownloadproofezsignsigner() bool`

GetBEzsignfoldertypeAllowdownloadproofezsignsigner returns the BEzsignfoldertypeAllowdownloadproofezsignsigner field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeAllowdownloadproofezsignsignerOk

`func (o *EzsignfoldertypeResponseCompoundV3) GetBEzsignfoldertypeAllowdownloadproofezsignsignerOk() (*bool, bool)`

GetBEzsignfoldertypeAllowdownloadproofezsignsignerOk returns a tuple with the BEzsignfoldertypeAllowdownloadproofezsignsigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeAllowdownloadproofezsignsigner

`func (o *EzsignfoldertypeResponseCompoundV3) SetBEzsignfoldertypeAllowdownloadproofezsignsigner(v bool)`

SetBEzsignfoldertypeAllowdownloadproofezsignsigner sets BEzsignfoldertypeAllowdownloadproofezsignsigner field to given value.

### HasBEzsignfoldertypeAllowdownloadproofezsignsigner

`func (o *EzsignfoldertypeResponseCompoundV3) HasBEzsignfoldertypeAllowdownloadproofezsignsigner() bool`

HasBEzsignfoldertypeAllowdownloadproofezsignsigner returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendproofreceivealldocument

`func (o *EzsignfoldertypeResponseCompoundV3) GetBEzsignfoldertypeSendproofreceivealldocument() bool`

GetBEzsignfoldertypeSendproofreceivealldocument returns the BEzsignfoldertypeSendproofreceivealldocument field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendproofreceivealldocumentOk

`func (o *EzsignfoldertypeResponseCompoundV3) GetBEzsignfoldertypeSendproofreceivealldocumentOk() (*bool, bool)`

GetBEzsignfoldertypeSendproofreceivealldocumentOk returns a tuple with the BEzsignfoldertypeSendproofreceivealldocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendproofreceivealldocument

`func (o *EzsignfoldertypeResponseCompoundV3) SetBEzsignfoldertypeSendproofreceivealldocument(v bool)`

SetBEzsignfoldertypeSendproofreceivealldocument sets BEzsignfoldertypeSendproofreceivealldocument field to given value.

### HasBEzsignfoldertypeSendproofreceivealldocument

`func (o *EzsignfoldertypeResponseCompoundV3) HasBEzsignfoldertypeSendproofreceivealldocument() bool`

HasBEzsignfoldertypeSendproofreceivealldocument returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendsignedtodocumentowner

`func (o *EzsignfoldertypeResponseCompoundV3) GetBEzsignfoldertypeSendsignedtodocumentowner() bool`

GetBEzsignfoldertypeSendsignedtodocumentowner returns the BEzsignfoldertypeSendsignedtodocumentowner field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsignedtodocumentownerOk

`func (o *EzsignfoldertypeResponseCompoundV3) GetBEzsignfoldertypeSendsignedtodocumentownerOk() (*bool, bool)`

GetBEzsignfoldertypeSendsignedtodocumentownerOk returns a tuple with the BEzsignfoldertypeSendsignedtodocumentowner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsignedtodocumentowner

`func (o *EzsignfoldertypeResponseCompoundV3) SetBEzsignfoldertypeSendsignedtodocumentowner(v bool)`

SetBEzsignfoldertypeSendsignedtodocumentowner sets BEzsignfoldertypeSendsignedtodocumentowner field to given value.


### GetBEzsignfoldertypeSendsignedtofolderowner

`func (o *EzsignfoldertypeResponseCompoundV3) GetBEzsignfoldertypeSendsignedtofolderowner() bool`

GetBEzsignfoldertypeSendsignedtofolderowner returns the BEzsignfoldertypeSendsignedtofolderowner field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsignedtofolderownerOk

`func (o *EzsignfoldertypeResponseCompoundV3) GetBEzsignfoldertypeSendsignedtofolderownerOk() (*bool, bool)`

GetBEzsignfoldertypeSendsignedtofolderownerOk returns a tuple with the BEzsignfoldertypeSendsignedtofolderowner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsignedtofolderowner

`func (o *EzsignfoldertypeResponseCompoundV3) SetBEzsignfoldertypeSendsignedtofolderowner(v bool)`

SetBEzsignfoldertypeSendsignedtofolderowner sets BEzsignfoldertypeSendsignedtofolderowner field to given value.


### GetBEzsignfoldertypeSendsignedtofullgroup

`func (o *EzsignfoldertypeResponseCompoundV3) GetBEzsignfoldertypeSendsignedtofullgroup() bool`

GetBEzsignfoldertypeSendsignedtofullgroup returns the BEzsignfoldertypeSendsignedtofullgroup field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsignedtofullgroupOk

`func (o *EzsignfoldertypeResponseCompoundV3) GetBEzsignfoldertypeSendsignedtofullgroupOk() (*bool, bool)`

GetBEzsignfoldertypeSendsignedtofullgroupOk returns a tuple with the BEzsignfoldertypeSendsignedtofullgroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsignedtofullgroup

`func (o *EzsignfoldertypeResponseCompoundV3) SetBEzsignfoldertypeSendsignedtofullgroup(v bool)`

SetBEzsignfoldertypeSendsignedtofullgroup sets BEzsignfoldertypeSendsignedtofullgroup field to given value.

### HasBEzsignfoldertypeSendsignedtofullgroup

`func (o *EzsignfoldertypeResponseCompoundV3) HasBEzsignfoldertypeSendsignedtofullgroup() bool`

HasBEzsignfoldertypeSendsignedtofullgroup returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendsignedtolimitedgroup

`func (o *EzsignfoldertypeResponseCompoundV3) GetBEzsignfoldertypeSendsignedtolimitedgroup() bool`

GetBEzsignfoldertypeSendsignedtolimitedgroup returns the BEzsignfoldertypeSendsignedtolimitedgroup field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsignedtolimitedgroupOk

`func (o *EzsignfoldertypeResponseCompoundV3) GetBEzsignfoldertypeSendsignedtolimitedgroupOk() (*bool, bool)`

GetBEzsignfoldertypeSendsignedtolimitedgroupOk returns a tuple with the BEzsignfoldertypeSendsignedtolimitedgroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsignedtolimitedgroup

`func (o *EzsignfoldertypeResponseCompoundV3) SetBEzsignfoldertypeSendsignedtolimitedgroup(v bool)`

SetBEzsignfoldertypeSendsignedtolimitedgroup sets BEzsignfoldertypeSendsignedtolimitedgroup field to given value.

### HasBEzsignfoldertypeSendsignedtolimitedgroup

`func (o *EzsignfoldertypeResponseCompoundV3) HasBEzsignfoldertypeSendsignedtolimitedgroup() bool`

HasBEzsignfoldertypeSendsignedtolimitedgroup returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendsignedtocolleague

`func (o *EzsignfoldertypeResponseCompoundV3) GetBEzsignfoldertypeSendsignedtocolleague() bool`

GetBEzsignfoldertypeSendsignedtocolleague returns the BEzsignfoldertypeSendsignedtocolleague field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsignedtocolleagueOk

`func (o *EzsignfoldertypeResponseCompoundV3) GetBEzsignfoldertypeSendsignedtocolleagueOk() (*bool, bool)`

GetBEzsignfoldertypeSendsignedtocolleagueOk returns a tuple with the BEzsignfoldertypeSendsignedtocolleague field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsignedtocolleague

`func (o *EzsignfoldertypeResponseCompoundV3) SetBEzsignfoldertypeSendsignedtocolleague(v bool)`

SetBEzsignfoldertypeSendsignedtocolleague sets BEzsignfoldertypeSendsignedtocolleague field to given value.


### GetBEzsignfoldertypeSendsummarytodocumentowner

`func (o *EzsignfoldertypeResponseCompoundV3) GetBEzsignfoldertypeSendsummarytodocumentowner() bool`

GetBEzsignfoldertypeSendsummarytodocumentowner returns the BEzsignfoldertypeSendsummarytodocumentowner field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsummarytodocumentownerOk

`func (o *EzsignfoldertypeResponseCompoundV3) GetBEzsignfoldertypeSendsummarytodocumentownerOk() (*bool, bool)`

GetBEzsignfoldertypeSendsummarytodocumentownerOk returns a tuple with the BEzsignfoldertypeSendsummarytodocumentowner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsummarytodocumentowner

`func (o *EzsignfoldertypeResponseCompoundV3) SetBEzsignfoldertypeSendsummarytodocumentowner(v bool)`

SetBEzsignfoldertypeSendsummarytodocumentowner sets BEzsignfoldertypeSendsummarytodocumentowner field to given value.


### GetBEzsignfoldertypeSendsummarytofolderowner

`func (o *EzsignfoldertypeResponseCompoundV3) GetBEzsignfoldertypeSendsummarytofolderowner() bool`

GetBEzsignfoldertypeSendsummarytofolderowner returns the BEzsignfoldertypeSendsummarytofolderowner field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsummarytofolderownerOk

`func (o *EzsignfoldertypeResponseCompoundV3) GetBEzsignfoldertypeSendsummarytofolderownerOk() (*bool, bool)`

GetBEzsignfoldertypeSendsummarytofolderownerOk returns a tuple with the BEzsignfoldertypeSendsummarytofolderowner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsummarytofolderowner

`func (o *EzsignfoldertypeResponseCompoundV3) SetBEzsignfoldertypeSendsummarytofolderowner(v bool)`

SetBEzsignfoldertypeSendsummarytofolderowner sets BEzsignfoldertypeSendsummarytofolderowner field to given value.


### GetBEzsignfoldertypeSendsummarytofullgroup

`func (o *EzsignfoldertypeResponseCompoundV3) GetBEzsignfoldertypeSendsummarytofullgroup() bool`

GetBEzsignfoldertypeSendsummarytofullgroup returns the BEzsignfoldertypeSendsummarytofullgroup field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsummarytofullgroupOk

`func (o *EzsignfoldertypeResponseCompoundV3) GetBEzsignfoldertypeSendsummarytofullgroupOk() (*bool, bool)`

GetBEzsignfoldertypeSendsummarytofullgroupOk returns a tuple with the BEzsignfoldertypeSendsummarytofullgroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsummarytofullgroup

`func (o *EzsignfoldertypeResponseCompoundV3) SetBEzsignfoldertypeSendsummarytofullgroup(v bool)`

SetBEzsignfoldertypeSendsummarytofullgroup sets BEzsignfoldertypeSendsummarytofullgroup field to given value.

### HasBEzsignfoldertypeSendsummarytofullgroup

`func (o *EzsignfoldertypeResponseCompoundV3) HasBEzsignfoldertypeSendsummarytofullgroup() bool`

HasBEzsignfoldertypeSendsummarytofullgroup returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendsummarytolimitedgroup

`func (o *EzsignfoldertypeResponseCompoundV3) GetBEzsignfoldertypeSendsummarytolimitedgroup() bool`

GetBEzsignfoldertypeSendsummarytolimitedgroup returns the BEzsignfoldertypeSendsummarytolimitedgroup field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsummarytolimitedgroupOk

`func (o *EzsignfoldertypeResponseCompoundV3) GetBEzsignfoldertypeSendsummarytolimitedgroupOk() (*bool, bool)`

GetBEzsignfoldertypeSendsummarytolimitedgroupOk returns a tuple with the BEzsignfoldertypeSendsummarytolimitedgroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsummarytolimitedgroup

`func (o *EzsignfoldertypeResponseCompoundV3) SetBEzsignfoldertypeSendsummarytolimitedgroup(v bool)`

SetBEzsignfoldertypeSendsummarytolimitedgroup sets BEzsignfoldertypeSendsummarytolimitedgroup field to given value.

### HasBEzsignfoldertypeSendsummarytolimitedgroup

`func (o *EzsignfoldertypeResponseCompoundV3) HasBEzsignfoldertypeSendsummarytolimitedgroup() bool`

HasBEzsignfoldertypeSendsummarytolimitedgroup returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendsummarytocolleague

`func (o *EzsignfoldertypeResponseCompoundV3) GetBEzsignfoldertypeSendsummarytocolleague() bool`

GetBEzsignfoldertypeSendsummarytocolleague returns the BEzsignfoldertypeSendsummarytocolleague field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsummarytocolleagueOk

`func (o *EzsignfoldertypeResponseCompoundV3) GetBEzsignfoldertypeSendsummarytocolleagueOk() (*bool, bool)`

GetBEzsignfoldertypeSendsummarytocolleagueOk returns a tuple with the BEzsignfoldertypeSendsummarytocolleague field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsummarytocolleague

`func (o *EzsignfoldertypeResponseCompoundV3) SetBEzsignfoldertypeSendsummarytocolleague(v bool)`

SetBEzsignfoldertypeSendsummarytocolleague sets BEzsignfoldertypeSendsummarytocolleague field to given value.


### GetBEzsignfoldertypeIsactive

`func (o *EzsignfoldertypeResponseCompoundV3) GetBEzsignfoldertypeIsactive() bool`

GetBEzsignfoldertypeIsactive returns the BEzsignfoldertypeIsactive field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeIsactiveOk

`func (o *EzsignfoldertypeResponseCompoundV3) GetBEzsignfoldertypeIsactiveOk() (*bool, bool)`

GetBEzsignfoldertypeIsactiveOk returns a tuple with the BEzsignfoldertypeIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeIsactive

`func (o *EzsignfoldertypeResponseCompoundV3) SetBEzsignfoldertypeIsactive(v bool)`

SetBEzsignfoldertypeIsactive sets BEzsignfoldertypeIsactive field to given value.


### GetAObjUserlogintype

`func (o *EzsignfoldertypeResponseCompoundV3) GetAObjUserlogintype() []UserlogintypeResponse`

GetAObjUserlogintype returns the AObjUserlogintype field if non-nil, zero value otherwise.

### GetAObjUserlogintypeOk

`func (o *EzsignfoldertypeResponseCompoundV3) GetAObjUserlogintypeOk() (*[]UserlogintypeResponse, bool)`

GetAObjUserlogintypeOk returns a tuple with the AObjUserlogintype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjUserlogintype

`func (o *EzsignfoldertypeResponseCompoundV3) SetAObjUserlogintype(v []UserlogintypeResponse)`

SetAObjUserlogintype sets AObjUserlogintype field to given value.


### GetAObjUsergroupAll

`func (o *EzsignfoldertypeResponseCompoundV3) GetAObjUsergroupAll() []UsergroupResponse`

GetAObjUsergroupAll returns the AObjUsergroupAll field if non-nil, zero value otherwise.

### GetAObjUsergroupAllOk

`func (o *EzsignfoldertypeResponseCompoundV3) GetAObjUsergroupAllOk() (*[]UsergroupResponse, bool)`

GetAObjUsergroupAllOk returns a tuple with the AObjUsergroupAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjUsergroupAll

`func (o *EzsignfoldertypeResponseCompoundV3) SetAObjUsergroupAll(v []UsergroupResponse)`

SetAObjUsergroupAll sets AObjUsergroupAll field to given value.

### HasAObjUsergroupAll

`func (o *EzsignfoldertypeResponseCompoundV3) HasAObjUsergroupAll() bool`

HasAObjUsergroupAll returns a boolean if a field has been set.

### GetAObjUsergroupRestricted

`func (o *EzsignfoldertypeResponseCompoundV3) GetAObjUsergroupRestricted() []UsergroupResponse`

GetAObjUsergroupRestricted returns the AObjUsergroupRestricted field if non-nil, zero value otherwise.

### GetAObjUsergroupRestrictedOk

`func (o *EzsignfoldertypeResponseCompoundV3) GetAObjUsergroupRestrictedOk() (*[]UsergroupResponse, bool)`

GetAObjUsergroupRestrictedOk returns a tuple with the AObjUsergroupRestricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjUsergroupRestricted

`func (o *EzsignfoldertypeResponseCompoundV3) SetAObjUsergroupRestricted(v []UsergroupResponse)`

SetAObjUsergroupRestricted sets AObjUsergroupRestricted field to given value.

### HasAObjUsergroupRestricted

`func (o *EzsignfoldertypeResponseCompoundV3) HasAObjUsergroupRestricted() bool`

HasAObjUsergroupRestricted returns a boolean if a field has been set.

### GetAObjUsergroupTemplate

`func (o *EzsignfoldertypeResponseCompoundV3) GetAObjUsergroupTemplate() []UsergroupResponse`

GetAObjUsergroupTemplate returns the AObjUsergroupTemplate field if non-nil, zero value otherwise.

### GetAObjUsergroupTemplateOk

`func (o *EzsignfoldertypeResponseCompoundV3) GetAObjUsergroupTemplateOk() (*[]UsergroupResponse, bool)`

GetAObjUsergroupTemplateOk returns a tuple with the AObjUsergroupTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjUsergroupTemplate

`func (o *EzsignfoldertypeResponseCompoundV3) SetAObjUsergroupTemplate(v []UsergroupResponse)`

SetAObjUsergroupTemplate sets AObjUsergroupTemplate field to given value.

### HasAObjUsergroupTemplate

`func (o *EzsignfoldertypeResponseCompoundV3) HasAObjUsergroupTemplate() bool`

HasAObjUsergroupTemplate returns a boolean if a field has been set.

### GetAFkiUserIDSigned

`func (o *EzsignfoldertypeResponseCompoundV3) GetAFkiUserIDSigned() []int32`

GetAFkiUserIDSigned returns the AFkiUserIDSigned field if non-nil, zero value otherwise.

### GetAFkiUserIDSignedOk

`func (o *EzsignfoldertypeResponseCompoundV3) GetAFkiUserIDSignedOk() (*[]int32, bool)`

GetAFkiUserIDSignedOk returns a tuple with the AFkiUserIDSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAFkiUserIDSigned

`func (o *EzsignfoldertypeResponseCompoundV3) SetAFkiUserIDSigned(v []int32)`

SetAFkiUserIDSigned sets AFkiUserIDSigned field to given value.

### HasAFkiUserIDSigned

`func (o *EzsignfoldertypeResponseCompoundV3) HasAFkiUserIDSigned() bool`

HasAFkiUserIDSigned returns a boolean if a field has been set.

### GetAFkiUserIDSummary

`func (o *EzsignfoldertypeResponseCompoundV3) GetAFkiUserIDSummary() []int32`

GetAFkiUserIDSummary returns the AFkiUserIDSummary field if non-nil, zero value otherwise.

### GetAFkiUserIDSummaryOk

`func (o *EzsignfoldertypeResponseCompoundV3) GetAFkiUserIDSummaryOk() (*[]int32, bool)`

GetAFkiUserIDSummaryOk returns a tuple with the AFkiUserIDSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAFkiUserIDSummary

`func (o *EzsignfoldertypeResponseCompoundV3) SetAFkiUserIDSummary(v []int32)`

SetAFkiUserIDSummary sets AFkiUserIDSummary field to given value.

### HasAFkiUserIDSummary

`func (o *EzsignfoldertypeResponseCompoundV3) HasAFkiUserIDSummary() bool`

HasAFkiUserIDSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


