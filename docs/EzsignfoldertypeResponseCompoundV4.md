# EzsignfoldertypeResponseCompoundV4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsignfoldertypeID** | **int32** | The unique ID of the Ezsignfoldertype. | 
**ObjEzsignfoldertypeName** | [**MultilingualEzsignfoldertypeName**](MultilingualEzsignfoldertypeName.md) |  | 
**FkiBrandingID** | **int32** | The unique ID of the Branding | 
**FkiBillingentityinternalID** | Pointer to **int32** | The unique ID of the Billingentityinternal. | [optional] 
**FkiEzsigntsarequirementID** | Pointer to **int32** | The unique ID of the Ezsigntsarequirement.  Determine if a Time Stamping Authority should add a timestamp on each of the signature. Valid values:  |Value|Description| |-|-| |1|No. TSA Timestamping will requested. This will make all signatures a lot faster since no round-trip to the TSA server will be required. Timestamping will be made using eZsign server&#39;s time.| |2|Best effort. Timestamping from a Time Stamping Authority will be requested but is not mandatory. In the very improbable case it cannot be completed, the timestamping will be made using eZsign server&#39;s time. **Additional fee applies**| |3|Mandatory. Timestamping from a Time Stamping Authority will be requested and is mandatory. In the very improbable case it cannot be completed, the signature will fail and the user will be asked to retry. **Additional fee applies**| | [optional] 
**FkiFontIDAnnotation** | Pointer to **int32** | The unique ID of the Font | [optional] 
**FkiFontIDFormfield** | Pointer to **int32** | The unique ID of the Font | [optional] 
**FkiFontIDSignature** | Pointer to **int32** | The unique ID of the Font | [optional] 
**FkiPdfalevelIDConvert** | Pointer to **int32** | The unique ID of the Pdfalevel | [optional] 
**EEzsignfoldertypeDocumentdependency** | Pointer to [**FieldEEzsignfoldertypeDocumentdependency**](FieldEEzsignfoldertypeDocumentdependency.md) |  | [optional] 
**SBrandingDescriptionX** | **string** | The Description of the Branding in the language of the requester | 
**SBillingentityinternalDescriptionX** | Pointer to **string** | The description of the Billingentityinternal in the language of the requester | [optional] 
**SEzsigntsarequirementDescriptionX** | Pointer to **string** | The description of the Ezsigntsarequirement in the language of the requester | [optional] 
**SEmailAddressSigned** | Pointer to **string** | The email address. | [optional] 
**SEmailAddressSummary** | Pointer to **string** | The email address. | [optional] 
**EEzsignfoldertypePdfarequirement** | Pointer to [**FieldEEzsignfoldertypePdfarequirement**](FieldEEzsignfoldertypePdfarequirement.md) |  | [optional] 
**EEzsignfoldertypePdfanoncompliantaction** | Pointer to [**FieldEEzsignfoldertypePdfanoncompliantaction**](FieldEEzsignfoldertypePdfanoncompliantaction.md) |  | [optional] 
**EEzsignfoldertypePrivacylevel** | [**FieldEEzsignfoldertypePrivacylevel**](FieldEEzsignfoldertypePrivacylevel.md) |  | 
**IEzsignfoldertypeFontsizeannotation** | Pointer to **int32** | Font size for annotations | [optional] 
**IEzsignfoldertypeFontsizeformfield** | Pointer to **int32** | Font size for form fields | [optional] 
**IEzsignfoldertypeSendreminderfirstdays** | Pointer to **int32** | The number of days before the the first reminder sending | [optional] 
**IEzsignfoldertypeSendreminderotherdays** | Pointer to **int32** | The number of days after the first reminder sending | [optional] 
**IEzsignfoldertypeArchivaldays** | **int32** | The number of days before the archival of Ezsignfolders created using this Ezsignfoldertype | 
**EEzsignfoldertypeDisposal** | [**FieldEEzsignfoldertypeDisposal**](FieldEEzsignfoldertypeDisposal.md) |  | 
**EEzsignfoldertypeCompletion** | [**FieldEEzsignfoldertypeCompletion**](FieldEEzsignfoldertypeCompletion.md) |  | 
**IEzsignfoldertypeDisposaldays** | Pointer to **int32** | The number of days after the archival before the disposal of the Ezsignfolder | [optional] 
**IEzsignfoldertypeDeadlinedays** | **int32** | The number of days to get all Ezsignsignatures | 
**BEzsignfoldertypePrematurelyendautomatically** | Pointer to **bool** | Wheter if document will be ended prematurely after Ezsignfolder expires. | [optional] 
**IEzsignfoldertypePrematurelyendautomaticallydays** | Pointer to **int32** | Number of days between Ezsignfolder expiration and automatic prematurely end of Ezsigndocuments. | [optional] 
**BEzsignfoldertypeAutomaticsignature** | Pointer to **bool** | Whether we allow the automatic signature by an User | [optional] 
**BEzsignfoldertypeDelegate** | Pointer to **bool** | Wheter if delegation of signature is allowed to another user or not | [optional] 
**BEzsignfoldertypeDiscussion** | Pointer to **bool** | Wheter if creating a new Discussion is allowed or not | [optional] 
**BEzsignfoldertypeLogrecipientinproof** | Pointer to **bool** | Whether we log recipient of signed document in proof | [optional] 
**BEzsignfoldertypeReassignezsignsigner** | Pointer to **bool** | Wheter if Reassignment of signature is allowed by a signatory to another signatory or not | [optional] 
**BEzsignfoldertypeReassignuser** | Pointer to **bool** | Wheter if Reassignment of signature is allowed by a user to a signatory or another user or not | [optional] 
**BEzsignfoldertypeReassigngroup** | Pointer to **bool** | Wheter if Reassignment of signatures of the groups to which the user belongs is authorized by a user to himself | [optional] 
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
**EEzsignfoldertypeSigneraccess** | Pointer to [**FieldEEzsignfoldertypeSigneraccess**](FieldEEzsignfoldertypeSigneraccess.md) |  | [optional] 
**BEzsignfoldertypeIsactive** | **bool** | Whether the Ezsignfoldertype is active or not | 
**AFkiPdfalevelID** | Pointer to **[]int32** |  | [optional] 
**AObjUserlogintype** | [**[]UserlogintypeResponse**](UserlogintypeResponse.md) |  | 
**AObjUsergroupAll** | Pointer to [**[]UsergroupResponse**](UsergroupResponse.md) |  | [optional] 
**AObjUsergroupRestricted** | Pointer to [**[]UsergroupResponse**](UsergroupResponse.md) |  | [optional] 
**AObjUsergroupTemplate** | Pointer to [**[]UsergroupResponse**](UsergroupResponse.md) |  | [optional] 
**ObjAudit** | [**CommonAudit**](CommonAudit.md) |  | 
**AFkiUserIDSigned** | Pointer to **[]int32** |  | [optional] 
**AFkiUserIDSummary** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewEzsignfoldertypeResponseCompoundV4

`func NewEzsignfoldertypeResponseCompoundV4(pkiEzsignfoldertypeID int32, objEzsignfoldertypeName MultilingualEzsignfoldertypeName, fkiBrandingID int32, sBrandingDescriptionX string, eEzsignfoldertypePrivacylevel FieldEEzsignfoldertypePrivacylevel, iEzsignfoldertypeArchivaldays int32, eEzsignfoldertypeDisposal FieldEEzsignfoldertypeDisposal, eEzsignfoldertypeCompletion FieldEEzsignfoldertypeCompletion, iEzsignfoldertypeDeadlinedays int32, bEzsignfoldertypeSendsignedtodocumentowner bool, bEzsignfoldertypeSendsignedtofolderowner bool, bEzsignfoldertypeSendsignedtocolleague bool, bEzsignfoldertypeSendsummarytodocumentowner bool, bEzsignfoldertypeSendsummarytofolderowner bool, bEzsignfoldertypeSendsummarytocolleague bool, bEzsignfoldertypeIsactive bool, aObjUserlogintype []UserlogintypeResponse, objAudit CommonAudit, ) *EzsignfoldertypeResponseCompoundV4`

NewEzsignfoldertypeResponseCompoundV4 instantiates a new EzsignfoldertypeResponseCompoundV4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignfoldertypeResponseCompoundV4WithDefaults

`func NewEzsignfoldertypeResponseCompoundV4WithDefaults() *EzsignfoldertypeResponseCompoundV4`

NewEzsignfoldertypeResponseCompoundV4WithDefaults instantiates a new EzsignfoldertypeResponseCompoundV4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsignfoldertypeID

`func (o *EzsignfoldertypeResponseCompoundV4) GetPkiEzsignfoldertypeID() int32`

GetPkiEzsignfoldertypeID returns the PkiEzsignfoldertypeID field if non-nil, zero value otherwise.

### GetPkiEzsignfoldertypeIDOk

`func (o *EzsignfoldertypeResponseCompoundV4) GetPkiEzsignfoldertypeIDOk() (*int32, bool)`

GetPkiEzsignfoldertypeIDOk returns a tuple with the PkiEzsignfoldertypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsignfoldertypeID

`func (o *EzsignfoldertypeResponseCompoundV4) SetPkiEzsignfoldertypeID(v int32)`

SetPkiEzsignfoldertypeID sets PkiEzsignfoldertypeID field to given value.


### GetObjEzsignfoldertypeName

`func (o *EzsignfoldertypeResponseCompoundV4) GetObjEzsignfoldertypeName() MultilingualEzsignfoldertypeName`

GetObjEzsignfoldertypeName returns the ObjEzsignfoldertypeName field if non-nil, zero value otherwise.

### GetObjEzsignfoldertypeNameOk

`func (o *EzsignfoldertypeResponseCompoundV4) GetObjEzsignfoldertypeNameOk() (*MultilingualEzsignfoldertypeName, bool)`

GetObjEzsignfoldertypeNameOk returns a tuple with the ObjEzsignfoldertypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjEzsignfoldertypeName

`func (o *EzsignfoldertypeResponseCompoundV4) SetObjEzsignfoldertypeName(v MultilingualEzsignfoldertypeName)`

SetObjEzsignfoldertypeName sets ObjEzsignfoldertypeName field to given value.


### GetFkiBrandingID

`func (o *EzsignfoldertypeResponseCompoundV4) GetFkiBrandingID() int32`

GetFkiBrandingID returns the FkiBrandingID field if non-nil, zero value otherwise.

### GetFkiBrandingIDOk

`func (o *EzsignfoldertypeResponseCompoundV4) GetFkiBrandingIDOk() (*int32, bool)`

GetFkiBrandingIDOk returns a tuple with the FkiBrandingID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiBrandingID

`func (o *EzsignfoldertypeResponseCompoundV4) SetFkiBrandingID(v int32)`

SetFkiBrandingID sets FkiBrandingID field to given value.


### GetFkiBillingentityinternalID

`func (o *EzsignfoldertypeResponseCompoundV4) GetFkiBillingentityinternalID() int32`

GetFkiBillingentityinternalID returns the FkiBillingentityinternalID field if non-nil, zero value otherwise.

### GetFkiBillingentityinternalIDOk

`func (o *EzsignfoldertypeResponseCompoundV4) GetFkiBillingentityinternalIDOk() (*int32, bool)`

GetFkiBillingentityinternalIDOk returns a tuple with the FkiBillingentityinternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiBillingentityinternalID

`func (o *EzsignfoldertypeResponseCompoundV4) SetFkiBillingentityinternalID(v int32)`

SetFkiBillingentityinternalID sets FkiBillingentityinternalID field to given value.

### HasFkiBillingentityinternalID

`func (o *EzsignfoldertypeResponseCompoundV4) HasFkiBillingentityinternalID() bool`

HasFkiBillingentityinternalID returns a boolean if a field has been set.

### GetFkiEzsigntsarequirementID

`func (o *EzsignfoldertypeResponseCompoundV4) GetFkiEzsigntsarequirementID() int32`

GetFkiEzsigntsarequirementID returns the FkiEzsigntsarequirementID field if non-nil, zero value otherwise.

### GetFkiEzsigntsarequirementIDOk

`func (o *EzsignfoldertypeResponseCompoundV4) GetFkiEzsigntsarequirementIDOk() (*int32, bool)`

GetFkiEzsigntsarequirementIDOk returns a tuple with the FkiEzsigntsarequirementID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntsarequirementID

`func (o *EzsignfoldertypeResponseCompoundV4) SetFkiEzsigntsarequirementID(v int32)`

SetFkiEzsigntsarequirementID sets FkiEzsigntsarequirementID field to given value.

### HasFkiEzsigntsarequirementID

`func (o *EzsignfoldertypeResponseCompoundV4) HasFkiEzsigntsarequirementID() bool`

HasFkiEzsigntsarequirementID returns a boolean if a field has been set.

### GetFkiFontIDAnnotation

`func (o *EzsignfoldertypeResponseCompoundV4) GetFkiFontIDAnnotation() int32`

GetFkiFontIDAnnotation returns the FkiFontIDAnnotation field if non-nil, zero value otherwise.

### GetFkiFontIDAnnotationOk

`func (o *EzsignfoldertypeResponseCompoundV4) GetFkiFontIDAnnotationOk() (*int32, bool)`

GetFkiFontIDAnnotationOk returns a tuple with the FkiFontIDAnnotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiFontIDAnnotation

`func (o *EzsignfoldertypeResponseCompoundV4) SetFkiFontIDAnnotation(v int32)`

SetFkiFontIDAnnotation sets FkiFontIDAnnotation field to given value.

### HasFkiFontIDAnnotation

`func (o *EzsignfoldertypeResponseCompoundV4) HasFkiFontIDAnnotation() bool`

HasFkiFontIDAnnotation returns a boolean if a field has been set.

### GetFkiFontIDFormfield

`func (o *EzsignfoldertypeResponseCompoundV4) GetFkiFontIDFormfield() int32`

GetFkiFontIDFormfield returns the FkiFontIDFormfield field if non-nil, zero value otherwise.

### GetFkiFontIDFormfieldOk

`func (o *EzsignfoldertypeResponseCompoundV4) GetFkiFontIDFormfieldOk() (*int32, bool)`

GetFkiFontIDFormfieldOk returns a tuple with the FkiFontIDFormfield field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiFontIDFormfield

`func (o *EzsignfoldertypeResponseCompoundV4) SetFkiFontIDFormfield(v int32)`

SetFkiFontIDFormfield sets FkiFontIDFormfield field to given value.

### HasFkiFontIDFormfield

`func (o *EzsignfoldertypeResponseCompoundV4) HasFkiFontIDFormfield() bool`

HasFkiFontIDFormfield returns a boolean if a field has been set.

### GetFkiFontIDSignature

`func (o *EzsignfoldertypeResponseCompoundV4) GetFkiFontIDSignature() int32`

GetFkiFontIDSignature returns the FkiFontIDSignature field if non-nil, zero value otherwise.

### GetFkiFontIDSignatureOk

`func (o *EzsignfoldertypeResponseCompoundV4) GetFkiFontIDSignatureOk() (*int32, bool)`

GetFkiFontIDSignatureOk returns a tuple with the FkiFontIDSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiFontIDSignature

`func (o *EzsignfoldertypeResponseCompoundV4) SetFkiFontIDSignature(v int32)`

SetFkiFontIDSignature sets FkiFontIDSignature field to given value.

### HasFkiFontIDSignature

`func (o *EzsignfoldertypeResponseCompoundV4) HasFkiFontIDSignature() bool`

HasFkiFontIDSignature returns a boolean if a field has been set.

### GetFkiPdfalevelIDConvert

`func (o *EzsignfoldertypeResponseCompoundV4) GetFkiPdfalevelIDConvert() int32`

GetFkiPdfalevelIDConvert returns the FkiPdfalevelIDConvert field if non-nil, zero value otherwise.

### GetFkiPdfalevelIDConvertOk

`func (o *EzsignfoldertypeResponseCompoundV4) GetFkiPdfalevelIDConvertOk() (*int32, bool)`

GetFkiPdfalevelIDConvertOk returns a tuple with the FkiPdfalevelIDConvert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiPdfalevelIDConvert

`func (o *EzsignfoldertypeResponseCompoundV4) SetFkiPdfalevelIDConvert(v int32)`

SetFkiPdfalevelIDConvert sets FkiPdfalevelIDConvert field to given value.

### HasFkiPdfalevelIDConvert

`func (o *EzsignfoldertypeResponseCompoundV4) HasFkiPdfalevelIDConvert() bool`

HasFkiPdfalevelIDConvert returns a boolean if a field has been set.

### GetEEzsignfoldertypeDocumentdependency

`func (o *EzsignfoldertypeResponseCompoundV4) GetEEzsignfoldertypeDocumentdependency() FieldEEzsignfoldertypeDocumentdependency`

GetEEzsignfoldertypeDocumentdependency returns the EEzsignfoldertypeDocumentdependency field if non-nil, zero value otherwise.

### GetEEzsignfoldertypeDocumentdependencyOk

`func (o *EzsignfoldertypeResponseCompoundV4) GetEEzsignfoldertypeDocumentdependencyOk() (*FieldEEzsignfoldertypeDocumentdependency, bool)`

GetEEzsignfoldertypeDocumentdependencyOk returns a tuple with the EEzsignfoldertypeDocumentdependency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignfoldertypeDocumentdependency

`func (o *EzsignfoldertypeResponseCompoundV4) SetEEzsignfoldertypeDocumentdependency(v FieldEEzsignfoldertypeDocumentdependency)`

SetEEzsignfoldertypeDocumentdependency sets EEzsignfoldertypeDocumentdependency field to given value.

### HasEEzsignfoldertypeDocumentdependency

`func (o *EzsignfoldertypeResponseCompoundV4) HasEEzsignfoldertypeDocumentdependency() bool`

HasEEzsignfoldertypeDocumentdependency returns a boolean if a field has been set.

### GetSBrandingDescriptionX

`func (o *EzsignfoldertypeResponseCompoundV4) GetSBrandingDescriptionX() string`

GetSBrandingDescriptionX returns the SBrandingDescriptionX field if non-nil, zero value otherwise.

### GetSBrandingDescriptionXOk

`func (o *EzsignfoldertypeResponseCompoundV4) GetSBrandingDescriptionXOk() (*string, bool)`

GetSBrandingDescriptionXOk returns a tuple with the SBrandingDescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSBrandingDescriptionX

`func (o *EzsignfoldertypeResponseCompoundV4) SetSBrandingDescriptionX(v string)`

SetSBrandingDescriptionX sets SBrandingDescriptionX field to given value.


### GetSBillingentityinternalDescriptionX

`func (o *EzsignfoldertypeResponseCompoundV4) GetSBillingentityinternalDescriptionX() string`

GetSBillingentityinternalDescriptionX returns the SBillingentityinternalDescriptionX field if non-nil, zero value otherwise.

### GetSBillingentityinternalDescriptionXOk

`func (o *EzsignfoldertypeResponseCompoundV4) GetSBillingentityinternalDescriptionXOk() (*string, bool)`

GetSBillingentityinternalDescriptionXOk returns a tuple with the SBillingentityinternalDescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSBillingentityinternalDescriptionX

`func (o *EzsignfoldertypeResponseCompoundV4) SetSBillingentityinternalDescriptionX(v string)`

SetSBillingentityinternalDescriptionX sets SBillingentityinternalDescriptionX field to given value.

### HasSBillingentityinternalDescriptionX

`func (o *EzsignfoldertypeResponseCompoundV4) HasSBillingentityinternalDescriptionX() bool`

HasSBillingentityinternalDescriptionX returns a boolean if a field has been set.

### GetSEzsigntsarequirementDescriptionX

`func (o *EzsignfoldertypeResponseCompoundV4) GetSEzsigntsarequirementDescriptionX() string`

GetSEzsigntsarequirementDescriptionX returns the SEzsigntsarequirementDescriptionX field if non-nil, zero value otherwise.

### GetSEzsigntsarequirementDescriptionXOk

`func (o *EzsignfoldertypeResponseCompoundV4) GetSEzsigntsarequirementDescriptionXOk() (*string, bool)`

GetSEzsigntsarequirementDescriptionXOk returns a tuple with the SEzsigntsarequirementDescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntsarequirementDescriptionX

`func (o *EzsignfoldertypeResponseCompoundV4) SetSEzsigntsarequirementDescriptionX(v string)`

SetSEzsigntsarequirementDescriptionX sets SEzsigntsarequirementDescriptionX field to given value.

### HasSEzsigntsarequirementDescriptionX

`func (o *EzsignfoldertypeResponseCompoundV4) HasSEzsigntsarequirementDescriptionX() bool`

HasSEzsigntsarequirementDescriptionX returns a boolean if a field has been set.

### GetSEmailAddressSigned

`func (o *EzsignfoldertypeResponseCompoundV4) GetSEmailAddressSigned() string`

GetSEmailAddressSigned returns the SEmailAddressSigned field if non-nil, zero value otherwise.

### GetSEmailAddressSignedOk

`func (o *EzsignfoldertypeResponseCompoundV4) GetSEmailAddressSignedOk() (*string, bool)`

GetSEmailAddressSignedOk returns a tuple with the SEmailAddressSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEmailAddressSigned

`func (o *EzsignfoldertypeResponseCompoundV4) SetSEmailAddressSigned(v string)`

SetSEmailAddressSigned sets SEmailAddressSigned field to given value.

### HasSEmailAddressSigned

`func (o *EzsignfoldertypeResponseCompoundV4) HasSEmailAddressSigned() bool`

HasSEmailAddressSigned returns a boolean if a field has been set.

### GetSEmailAddressSummary

`func (o *EzsignfoldertypeResponseCompoundV4) GetSEmailAddressSummary() string`

GetSEmailAddressSummary returns the SEmailAddressSummary field if non-nil, zero value otherwise.

### GetSEmailAddressSummaryOk

`func (o *EzsignfoldertypeResponseCompoundV4) GetSEmailAddressSummaryOk() (*string, bool)`

GetSEmailAddressSummaryOk returns a tuple with the SEmailAddressSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEmailAddressSummary

`func (o *EzsignfoldertypeResponseCompoundV4) SetSEmailAddressSummary(v string)`

SetSEmailAddressSummary sets SEmailAddressSummary field to given value.

### HasSEmailAddressSummary

`func (o *EzsignfoldertypeResponseCompoundV4) HasSEmailAddressSummary() bool`

HasSEmailAddressSummary returns a boolean if a field has been set.

### GetEEzsignfoldertypePdfarequirement

`func (o *EzsignfoldertypeResponseCompoundV4) GetEEzsignfoldertypePdfarequirement() FieldEEzsignfoldertypePdfarequirement`

GetEEzsignfoldertypePdfarequirement returns the EEzsignfoldertypePdfarequirement field if non-nil, zero value otherwise.

### GetEEzsignfoldertypePdfarequirementOk

`func (o *EzsignfoldertypeResponseCompoundV4) GetEEzsignfoldertypePdfarequirementOk() (*FieldEEzsignfoldertypePdfarequirement, bool)`

GetEEzsignfoldertypePdfarequirementOk returns a tuple with the EEzsignfoldertypePdfarequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignfoldertypePdfarequirement

`func (o *EzsignfoldertypeResponseCompoundV4) SetEEzsignfoldertypePdfarequirement(v FieldEEzsignfoldertypePdfarequirement)`

SetEEzsignfoldertypePdfarequirement sets EEzsignfoldertypePdfarequirement field to given value.

### HasEEzsignfoldertypePdfarequirement

`func (o *EzsignfoldertypeResponseCompoundV4) HasEEzsignfoldertypePdfarequirement() bool`

HasEEzsignfoldertypePdfarequirement returns a boolean if a field has been set.

### GetEEzsignfoldertypePdfanoncompliantaction

`func (o *EzsignfoldertypeResponseCompoundV4) GetEEzsignfoldertypePdfanoncompliantaction() FieldEEzsignfoldertypePdfanoncompliantaction`

GetEEzsignfoldertypePdfanoncompliantaction returns the EEzsignfoldertypePdfanoncompliantaction field if non-nil, zero value otherwise.

### GetEEzsignfoldertypePdfanoncompliantactionOk

`func (o *EzsignfoldertypeResponseCompoundV4) GetEEzsignfoldertypePdfanoncompliantactionOk() (*FieldEEzsignfoldertypePdfanoncompliantaction, bool)`

GetEEzsignfoldertypePdfanoncompliantactionOk returns a tuple with the EEzsignfoldertypePdfanoncompliantaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignfoldertypePdfanoncompliantaction

`func (o *EzsignfoldertypeResponseCompoundV4) SetEEzsignfoldertypePdfanoncompliantaction(v FieldEEzsignfoldertypePdfanoncompliantaction)`

SetEEzsignfoldertypePdfanoncompliantaction sets EEzsignfoldertypePdfanoncompliantaction field to given value.

### HasEEzsignfoldertypePdfanoncompliantaction

`func (o *EzsignfoldertypeResponseCompoundV4) HasEEzsignfoldertypePdfanoncompliantaction() bool`

HasEEzsignfoldertypePdfanoncompliantaction returns a boolean if a field has been set.

### GetEEzsignfoldertypePrivacylevel

`func (o *EzsignfoldertypeResponseCompoundV4) GetEEzsignfoldertypePrivacylevel() FieldEEzsignfoldertypePrivacylevel`

GetEEzsignfoldertypePrivacylevel returns the EEzsignfoldertypePrivacylevel field if non-nil, zero value otherwise.

### GetEEzsignfoldertypePrivacylevelOk

`func (o *EzsignfoldertypeResponseCompoundV4) GetEEzsignfoldertypePrivacylevelOk() (*FieldEEzsignfoldertypePrivacylevel, bool)`

GetEEzsignfoldertypePrivacylevelOk returns a tuple with the EEzsignfoldertypePrivacylevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignfoldertypePrivacylevel

`func (o *EzsignfoldertypeResponseCompoundV4) SetEEzsignfoldertypePrivacylevel(v FieldEEzsignfoldertypePrivacylevel)`

SetEEzsignfoldertypePrivacylevel sets EEzsignfoldertypePrivacylevel field to given value.


### GetIEzsignfoldertypeFontsizeannotation

`func (o *EzsignfoldertypeResponseCompoundV4) GetIEzsignfoldertypeFontsizeannotation() int32`

GetIEzsignfoldertypeFontsizeannotation returns the IEzsignfoldertypeFontsizeannotation field if non-nil, zero value otherwise.

### GetIEzsignfoldertypeFontsizeannotationOk

`func (o *EzsignfoldertypeResponseCompoundV4) GetIEzsignfoldertypeFontsizeannotationOk() (*int32, bool)`

GetIEzsignfoldertypeFontsizeannotationOk returns a tuple with the IEzsignfoldertypeFontsizeannotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignfoldertypeFontsizeannotation

`func (o *EzsignfoldertypeResponseCompoundV4) SetIEzsignfoldertypeFontsizeannotation(v int32)`

SetIEzsignfoldertypeFontsizeannotation sets IEzsignfoldertypeFontsizeannotation field to given value.

### HasIEzsignfoldertypeFontsizeannotation

`func (o *EzsignfoldertypeResponseCompoundV4) HasIEzsignfoldertypeFontsizeannotation() bool`

HasIEzsignfoldertypeFontsizeannotation returns a boolean if a field has been set.

### GetIEzsignfoldertypeFontsizeformfield

`func (o *EzsignfoldertypeResponseCompoundV4) GetIEzsignfoldertypeFontsizeformfield() int32`

GetIEzsignfoldertypeFontsizeformfield returns the IEzsignfoldertypeFontsizeformfield field if non-nil, zero value otherwise.

### GetIEzsignfoldertypeFontsizeformfieldOk

`func (o *EzsignfoldertypeResponseCompoundV4) GetIEzsignfoldertypeFontsizeformfieldOk() (*int32, bool)`

GetIEzsignfoldertypeFontsizeformfieldOk returns a tuple with the IEzsignfoldertypeFontsizeformfield field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignfoldertypeFontsizeformfield

`func (o *EzsignfoldertypeResponseCompoundV4) SetIEzsignfoldertypeFontsizeformfield(v int32)`

SetIEzsignfoldertypeFontsizeformfield sets IEzsignfoldertypeFontsizeformfield field to given value.

### HasIEzsignfoldertypeFontsizeformfield

`func (o *EzsignfoldertypeResponseCompoundV4) HasIEzsignfoldertypeFontsizeformfield() bool`

HasIEzsignfoldertypeFontsizeformfield returns a boolean if a field has been set.

### GetIEzsignfoldertypeSendreminderfirstdays

`func (o *EzsignfoldertypeResponseCompoundV4) GetIEzsignfoldertypeSendreminderfirstdays() int32`

GetIEzsignfoldertypeSendreminderfirstdays returns the IEzsignfoldertypeSendreminderfirstdays field if non-nil, zero value otherwise.

### GetIEzsignfoldertypeSendreminderfirstdaysOk

`func (o *EzsignfoldertypeResponseCompoundV4) GetIEzsignfoldertypeSendreminderfirstdaysOk() (*int32, bool)`

GetIEzsignfoldertypeSendreminderfirstdaysOk returns a tuple with the IEzsignfoldertypeSendreminderfirstdays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignfoldertypeSendreminderfirstdays

`func (o *EzsignfoldertypeResponseCompoundV4) SetIEzsignfoldertypeSendreminderfirstdays(v int32)`

SetIEzsignfoldertypeSendreminderfirstdays sets IEzsignfoldertypeSendreminderfirstdays field to given value.

### HasIEzsignfoldertypeSendreminderfirstdays

`func (o *EzsignfoldertypeResponseCompoundV4) HasIEzsignfoldertypeSendreminderfirstdays() bool`

HasIEzsignfoldertypeSendreminderfirstdays returns a boolean if a field has been set.

### GetIEzsignfoldertypeSendreminderotherdays

`func (o *EzsignfoldertypeResponseCompoundV4) GetIEzsignfoldertypeSendreminderotherdays() int32`

GetIEzsignfoldertypeSendreminderotherdays returns the IEzsignfoldertypeSendreminderotherdays field if non-nil, zero value otherwise.

### GetIEzsignfoldertypeSendreminderotherdaysOk

`func (o *EzsignfoldertypeResponseCompoundV4) GetIEzsignfoldertypeSendreminderotherdaysOk() (*int32, bool)`

GetIEzsignfoldertypeSendreminderotherdaysOk returns a tuple with the IEzsignfoldertypeSendreminderotherdays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignfoldertypeSendreminderotherdays

`func (o *EzsignfoldertypeResponseCompoundV4) SetIEzsignfoldertypeSendreminderotherdays(v int32)`

SetIEzsignfoldertypeSendreminderotherdays sets IEzsignfoldertypeSendreminderotherdays field to given value.

### HasIEzsignfoldertypeSendreminderotherdays

`func (o *EzsignfoldertypeResponseCompoundV4) HasIEzsignfoldertypeSendreminderotherdays() bool`

HasIEzsignfoldertypeSendreminderotherdays returns a boolean if a field has been set.

### GetIEzsignfoldertypeArchivaldays

`func (o *EzsignfoldertypeResponseCompoundV4) GetIEzsignfoldertypeArchivaldays() int32`

GetIEzsignfoldertypeArchivaldays returns the IEzsignfoldertypeArchivaldays field if non-nil, zero value otherwise.

### GetIEzsignfoldertypeArchivaldaysOk

`func (o *EzsignfoldertypeResponseCompoundV4) GetIEzsignfoldertypeArchivaldaysOk() (*int32, bool)`

GetIEzsignfoldertypeArchivaldaysOk returns a tuple with the IEzsignfoldertypeArchivaldays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignfoldertypeArchivaldays

`func (o *EzsignfoldertypeResponseCompoundV4) SetIEzsignfoldertypeArchivaldays(v int32)`

SetIEzsignfoldertypeArchivaldays sets IEzsignfoldertypeArchivaldays field to given value.


### GetEEzsignfoldertypeDisposal

`func (o *EzsignfoldertypeResponseCompoundV4) GetEEzsignfoldertypeDisposal() FieldEEzsignfoldertypeDisposal`

GetEEzsignfoldertypeDisposal returns the EEzsignfoldertypeDisposal field if non-nil, zero value otherwise.

### GetEEzsignfoldertypeDisposalOk

`func (o *EzsignfoldertypeResponseCompoundV4) GetEEzsignfoldertypeDisposalOk() (*FieldEEzsignfoldertypeDisposal, bool)`

GetEEzsignfoldertypeDisposalOk returns a tuple with the EEzsignfoldertypeDisposal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignfoldertypeDisposal

`func (o *EzsignfoldertypeResponseCompoundV4) SetEEzsignfoldertypeDisposal(v FieldEEzsignfoldertypeDisposal)`

SetEEzsignfoldertypeDisposal sets EEzsignfoldertypeDisposal field to given value.


### GetEEzsignfoldertypeCompletion

`func (o *EzsignfoldertypeResponseCompoundV4) GetEEzsignfoldertypeCompletion() FieldEEzsignfoldertypeCompletion`

GetEEzsignfoldertypeCompletion returns the EEzsignfoldertypeCompletion field if non-nil, zero value otherwise.

### GetEEzsignfoldertypeCompletionOk

`func (o *EzsignfoldertypeResponseCompoundV4) GetEEzsignfoldertypeCompletionOk() (*FieldEEzsignfoldertypeCompletion, bool)`

GetEEzsignfoldertypeCompletionOk returns a tuple with the EEzsignfoldertypeCompletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignfoldertypeCompletion

`func (o *EzsignfoldertypeResponseCompoundV4) SetEEzsignfoldertypeCompletion(v FieldEEzsignfoldertypeCompletion)`

SetEEzsignfoldertypeCompletion sets EEzsignfoldertypeCompletion field to given value.


### GetIEzsignfoldertypeDisposaldays

`func (o *EzsignfoldertypeResponseCompoundV4) GetIEzsignfoldertypeDisposaldays() int32`

GetIEzsignfoldertypeDisposaldays returns the IEzsignfoldertypeDisposaldays field if non-nil, zero value otherwise.

### GetIEzsignfoldertypeDisposaldaysOk

`func (o *EzsignfoldertypeResponseCompoundV4) GetIEzsignfoldertypeDisposaldaysOk() (*int32, bool)`

GetIEzsignfoldertypeDisposaldaysOk returns a tuple with the IEzsignfoldertypeDisposaldays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignfoldertypeDisposaldays

`func (o *EzsignfoldertypeResponseCompoundV4) SetIEzsignfoldertypeDisposaldays(v int32)`

SetIEzsignfoldertypeDisposaldays sets IEzsignfoldertypeDisposaldays field to given value.

### HasIEzsignfoldertypeDisposaldays

`func (o *EzsignfoldertypeResponseCompoundV4) HasIEzsignfoldertypeDisposaldays() bool`

HasIEzsignfoldertypeDisposaldays returns a boolean if a field has been set.

### GetIEzsignfoldertypeDeadlinedays

`func (o *EzsignfoldertypeResponseCompoundV4) GetIEzsignfoldertypeDeadlinedays() int32`

GetIEzsignfoldertypeDeadlinedays returns the IEzsignfoldertypeDeadlinedays field if non-nil, zero value otherwise.

### GetIEzsignfoldertypeDeadlinedaysOk

`func (o *EzsignfoldertypeResponseCompoundV4) GetIEzsignfoldertypeDeadlinedaysOk() (*int32, bool)`

GetIEzsignfoldertypeDeadlinedaysOk returns a tuple with the IEzsignfoldertypeDeadlinedays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignfoldertypeDeadlinedays

`func (o *EzsignfoldertypeResponseCompoundV4) SetIEzsignfoldertypeDeadlinedays(v int32)`

SetIEzsignfoldertypeDeadlinedays sets IEzsignfoldertypeDeadlinedays field to given value.


### GetBEzsignfoldertypePrematurelyendautomatically

`func (o *EzsignfoldertypeResponseCompoundV4) GetBEzsignfoldertypePrematurelyendautomatically() bool`

GetBEzsignfoldertypePrematurelyendautomatically returns the BEzsignfoldertypePrematurelyendautomatically field if non-nil, zero value otherwise.

### GetBEzsignfoldertypePrematurelyendautomaticallyOk

`func (o *EzsignfoldertypeResponseCompoundV4) GetBEzsignfoldertypePrematurelyendautomaticallyOk() (*bool, bool)`

GetBEzsignfoldertypePrematurelyendautomaticallyOk returns a tuple with the BEzsignfoldertypePrematurelyendautomatically field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypePrematurelyendautomatically

`func (o *EzsignfoldertypeResponseCompoundV4) SetBEzsignfoldertypePrematurelyendautomatically(v bool)`

SetBEzsignfoldertypePrematurelyendautomatically sets BEzsignfoldertypePrematurelyendautomatically field to given value.

### HasBEzsignfoldertypePrematurelyendautomatically

`func (o *EzsignfoldertypeResponseCompoundV4) HasBEzsignfoldertypePrematurelyendautomatically() bool`

HasBEzsignfoldertypePrematurelyendautomatically returns a boolean if a field has been set.

### GetIEzsignfoldertypePrematurelyendautomaticallydays

`func (o *EzsignfoldertypeResponseCompoundV4) GetIEzsignfoldertypePrematurelyendautomaticallydays() int32`

GetIEzsignfoldertypePrematurelyendautomaticallydays returns the IEzsignfoldertypePrematurelyendautomaticallydays field if non-nil, zero value otherwise.

### GetIEzsignfoldertypePrematurelyendautomaticallydaysOk

`func (o *EzsignfoldertypeResponseCompoundV4) GetIEzsignfoldertypePrematurelyendautomaticallydaysOk() (*int32, bool)`

GetIEzsignfoldertypePrematurelyendautomaticallydaysOk returns a tuple with the IEzsignfoldertypePrematurelyendautomaticallydays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignfoldertypePrematurelyendautomaticallydays

`func (o *EzsignfoldertypeResponseCompoundV4) SetIEzsignfoldertypePrematurelyendautomaticallydays(v int32)`

SetIEzsignfoldertypePrematurelyendautomaticallydays sets IEzsignfoldertypePrematurelyendautomaticallydays field to given value.

### HasIEzsignfoldertypePrematurelyendautomaticallydays

`func (o *EzsignfoldertypeResponseCompoundV4) HasIEzsignfoldertypePrematurelyendautomaticallydays() bool`

HasIEzsignfoldertypePrematurelyendautomaticallydays returns a boolean if a field has been set.

### GetBEzsignfoldertypeAutomaticsignature

`func (o *EzsignfoldertypeResponseCompoundV4) GetBEzsignfoldertypeAutomaticsignature() bool`

GetBEzsignfoldertypeAutomaticsignature returns the BEzsignfoldertypeAutomaticsignature field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeAutomaticsignatureOk

`func (o *EzsignfoldertypeResponseCompoundV4) GetBEzsignfoldertypeAutomaticsignatureOk() (*bool, bool)`

GetBEzsignfoldertypeAutomaticsignatureOk returns a tuple with the BEzsignfoldertypeAutomaticsignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeAutomaticsignature

`func (o *EzsignfoldertypeResponseCompoundV4) SetBEzsignfoldertypeAutomaticsignature(v bool)`

SetBEzsignfoldertypeAutomaticsignature sets BEzsignfoldertypeAutomaticsignature field to given value.

### HasBEzsignfoldertypeAutomaticsignature

`func (o *EzsignfoldertypeResponseCompoundV4) HasBEzsignfoldertypeAutomaticsignature() bool`

HasBEzsignfoldertypeAutomaticsignature returns a boolean if a field has been set.

### GetBEzsignfoldertypeDelegate

`func (o *EzsignfoldertypeResponseCompoundV4) GetBEzsignfoldertypeDelegate() bool`

GetBEzsignfoldertypeDelegate returns the BEzsignfoldertypeDelegate field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeDelegateOk

`func (o *EzsignfoldertypeResponseCompoundV4) GetBEzsignfoldertypeDelegateOk() (*bool, bool)`

GetBEzsignfoldertypeDelegateOk returns a tuple with the BEzsignfoldertypeDelegate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeDelegate

`func (o *EzsignfoldertypeResponseCompoundV4) SetBEzsignfoldertypeDelegate(v bool)`

SetBEzsignfoldertypeDelegate sets BEzsignfoldertypeDelegate field to given value.

### HasBEzsignfoldertypeDelegate

`func (o *EzsignfoldertypeResponseCompoundV4) HasBEzsignfoldertypeDelegate() bool`

HasBEzsignfoldertypeDelegate returns a boolean if a field has been set.

### GetBEzsignfoldertypeDiscussion

`func (o *EzsignfoldertypeResponseCompoundV4) GetBEzsignfoldertypeDiscussion() bool`

GetBEzsignfoldertypeDiscussion returns the BEzsignfoldertypeDiscussion field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeDiscussionOk

`func (o *EzsignfoldertypeResponseCompoundV4) GetBEzsignfoldertypeDiscussionOk() (*bool, bool)`

GetBEzsignfoldertypeDiscussionOk returns a tuple with the BEzsignfoldertypeDiscussion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeDiscussion

`func (o *EzsignfoldertypeResponseCompoundV4) SetBEzsignfoldertypeDiscussion(v bool)`

SetBEzsignfoldertypeDiscussion sets BEzsignfoldertypeDiscussion field to given value.

### HasBEzsignfoldertypeDiscussion

`func (o *EzsignfoldertypeResponseCompoundV4) HasBEzsignfoldertypeDiscussion() bool`

HasBEzsignfoldertypeDiscussion returns a boolean if a field has been set.

### GetBEzsignfoldertypeLogrecipientinproof

`func (o *EzsignfoldertypeResponseCompoundV4) GetBEzsignfoldertypeLogrecipientinproof() bool`

GetBEzsignfoldertypeLogrecipientinproof returns the BEzsignfoldertypeLogrecipientinproof field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeLogrecipientinproofOk

`func (o *EzsignfoldertypeResponseCompoundV4) GetBEzsignfoldertypeLogrecipientinproofOk() (*bool, bool)`

GetBEzsignfoldertypeLogrecipientinproofOk returns a tuple with the BEzsignfoldertypeLogrecipientinproof field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeLogrecipientinproof

`func (o *EzsignfoldertypeResponseCompoundV4) SetBEzsignfoldertypeLogrecipientinproof(v bool)`

SetBEzsignfoldertypeLogrecipientinproof sets BEzsignfoldertypeLogrecipientinproof field to given value.

### HasBEzsignfoldertypeLogrecipientinproof

`func (o *EzsignfoldertypeResponseCompoundV4) HasBEzsignfoldertypeLogrecipientinproof() bool`

HasBEzsignfoldertypeLogrecipientinproof returns a boolean if a field has been set.

### GetBEzsignfoldertypeReassignezsignsigner

`func (o *EzsignfoldertypeResponseCompoundV4) GetBEzsignfoldertypeReassignezsignsigner() bool`

GetBEzsignfoldertypeReassignezsignsigner returns the BEzsignfoldertypeReassignezsignsigner field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeReassignezsignsignerOk

`func (o *EzsignfoldertypeResponseCompoundV4) GetBEzsignfoldertypeReassignezsignsignerOk() (*bool, bool)`

GetBEzsignfoldertypeReassignezsignsignerOk returns a tuple with the BEzsignfoldertypeReassignezsignsigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeReassignezsignsigner

`func (o *EzsignfoldertypeResponseCompoundV4) SetBEzsignfoldertypeReassignezsignsigner(v bool)`

SetBEzsignfoldertypeReassignezsignsigner sets BEzsignfoldertypeReassignezsignsigner field to given value.

### HasBEzsignfoldertypeReassignezsignsigner

`func (o *EzsignfoldertypeResponseCompoundV4) HasBEzsignfoldertypeReassignezsignsigner() bool`

HasBEzsignfoldertypeReassignezsignsigner returns a boolean if a field has been set.

### GetBEzsignfoldertypeReassignuser

`func (o *EzsignfoldertypeResponseCompoundV4) GetBEzsignfoldertypeReassignuser() bool`

GetBEzsignfoldertypeReassignuser returns the BEzsignfoldertypeReassignuser field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeReassignuserOk

`func (o *EzsignfoldertypeResponseCompoundV4) GetBEzsignfoldertypeReassignuserOk() (*bool, bool)`

GetBEzsignfoldertypeReassignuserOk returns a tuple with the BEzsignfoldertypeReassignuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeReassignuser

`func (o *EzsignfoldertypeResponseCompoundV4) SetBEzsignfoldertypeReassignuser(v bool)`

SetBEzsignfoldertypeReassignuser sets BEzsignfoldertypeReassignuser field to given value.

### HasBEzsignfoldertypeReassignuser

`func (o *EzsignfoldertypeResponseCompoundV4) HasBEzsignfoldertypeReassignuser() bool`

HasBEzsignfoldertypeReassignuser returns a boolean if a field has been set.

### GetBEzsignfoldertypeReassigngroup

`func (o *EzsignfoldertypeResponseCompoundV4) GetBEzsignfoldertypeReassigngroup() bool`

GetBEzsignfoldertypeReassigngroup returns the BEzsignfoldertypeReassigngroup field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeReassigngroupOk

`func (o *EzsignfoldertypeResponseCompoundV4) GetBEzsignfoldertypeReassigngroupOk() (*bool, bool)`

GetBEzsignfoldertypeReassigngroupOk returns a tuple with the BEzsignfoldertypeReassigngroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeReassigngroup

`func (o *EzsignfoldertypeResponseCompoundV4) SetBEzsignfoldertypeReassigngroup(v bool)`

SetBEzsignfoldertypeReassigngroup sets BEzsignfoldertypeReassigngroup field to given value.

### HasBEzsignfoldertypeReassigngroup

`func (o *EzsignfoldertypeResponseCompoundV4) HasBEzsignfoldertypeReassigngroup() bool`

HasBEzsignfoldertypeReassigngroup returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendsignedtoezsignsigner

`func (o *EzsignfoldertypeResponseCompoundV4) GetBEzsignfoldertypeSendsignedtoezsignsigner() bool`

GetBEzsignfoldertypeSendsignedtoezsignsigner returns the BEzsignfoldertypeSendsignedtoezsignsigner field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsignedtoezsignsignerOk

`func (o *EzsignfoldertypeResponseCompoundV4) GetBEzsignfoldertypeSendsignedtoezsignsignerOk() (*bool, bool)`

GetBEzsignfoldertypeSendsignedtoezsignsignerOk returns a tuple with the BEzsignfoldertypeSendsignedtoezsignsigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsignedtoezsignsigner

`func (o *EzsignfoldertypeResponseCompoundV4) SetBEzsignfoldertypeSendsignedtoezsignsigner(v bool)`

SetBEzsignfoldertypeSendsignedtoezsignsigner sets BEzsignfoldertypeSendsignedtoezsignsigner field to given value.

### HasBEzsignfoldertypeSendsignedtoezsignsigner

`func (o *EzsignfoldertypeResponseCompoundV4) HasBEzsignfoldertypeSendsignedtoezsignsigner() bool`

HasBEzsignfoldertypeSendsignedtoezsignsigner returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendsignedtouser

`func (o *EzsignfoldertypeResponseCompoundV4) GetBEzsignfoldertypeSendsignedtouser() bool`

GetBEzsignfoldertypeSendsignedtouser returns the BEzsignfoldertypeSendsignedtouser field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsignedtouserOk

`func (o *EzsignfoldertypeResponseCompoundV4) GetBEzsignfoldertypeSendsignedtouserOk() (*bool, bool)`

GetBEzsignfoldertypeSendsignedtouserOk returns a tuple with the BEzsignfoldertypeSendsignedtouser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsignedtouser

`func (o *EzsignfoldertypeResponseCompoundV4) SetBEzsignfoldertypeSendsignedtouser(v bool)`

SetBEzsignfoldertypeSendsignedtouser sets BEzsignfoldertypeSendsignedtouser field to given value.

### HasBEzsignfoldertypeSendsignedtouser

`func (o *EzsignfoldertypeResponseCompoundV4) HasBEzsignfoldertypeSendsignedtouser() bool`

HasBEzsignfoldertypeSendsignedtouser returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendattachmentezsignsigner

`func (o *EzsignfoldertypeResponseCompoundV4) GetBEzsignfoldertypeSendattachmentezsignsigner() bool`

GetBEzsignfoldertypeSendattachmentezsignsigner returns the BEzsignfoldertypeSendattachmentezsignsigner field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendattachmentezsignsignerOk

`func (o *EzsignfoldertypeResponseCompoundV4) GetBEzsignfoldertypeSendattachmentezsignsignerOk() (*bool, bool)`

GetBEzsignfoldertypeSendattachmentezsignsignerOk returns a tuple with the BEzsignfoldertypeSendattachmentezsignsigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendattachmentezsignsigner

`func (o *EzsignfoldertypeResponseCompoundV4) SetBEzsignfoldertypeSendattachmentezsignsigner(v bool)`

SetBEzsignfoldertypeSendattachmentezsignsigner sets BEzsignfoldertypeSendattachmentezsignsigner field to given value.

### HasBEzsignfoldertypeSendattachmentezsignsigner

`func (o *EzsignfoldertypeResponseCompoundV4) HasBEzsignfoldertypeSendattachmentezsignsigner() bool`

HasBEzsignfoldertypeSendattachmentezsignsigner returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendproofezsignsigner

`func (o *EzsignfoldertypeResponseCompoundV4) GetBEzsignfoldertypeSendproofezsignsigner() bool`

GetBEzsignfoldertypeSendproofezsignsigner returns the BEzsignfoldertypeSendproofezsignsigner field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendproofezsignsignerOk

`func (o *EzsignfoldertypeResponseCompoundV4) GetBEzsignfoldertypeSendproofezsignsignerOk() (*bool, bool)`

GetBEzsignfoldertypeSendproofezsignsignerOk returns a tuple with the BEzsignfoldertypeSendproofezsignsigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendproofezsignsigner

`func (o *EzsignfoldertypeResponseCompoundV4) SetBEzsignfoldertypeSendproofezsignsigner(v bool)`

SetBEzsignfoldertypeSendproofezsignsigner sets BEzsignfoldertypeSendproofezsignsigner field to given value.

### HasBEzsignfoldertypeSendproofezsignsigner

`func (o *EzsignfoldertypeResponseCompoundV4) HasBEzsignfoldertypeSendproofezsignsigner() bool`

HasBEzsignfoldertypeSendproofezsignsigner returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendattachmentuser

`func (o *EzsignfoldertypeResponseCompoundV4) GetBEzsignfoldertypeSendattachmentuser() bool`

GetBEzsignfoldertypeSendattachmentuser returns the BEzsignfoldertypeSendattachmentuser field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendattachmentuserOk

`func (o *EzsignfoldertypeResponseCompoundV4) GetBEzsignfoldertypeSendattachmentuserOk() (*bool, bool)`

GetBEzsignfoldertypeSendattachmentuserOk returns a tuple with the BEzsignfoldertypeSendattachmentuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendattachmentuser

`func (o *EzsignfoldertypeResponseCompoundV4) SetBEzsignfoldertypeSendattachmentuser(v bool)`

SetBEzsignfoldertypeSendattachmentuser sets BEzsignfoldertypeSendattachmentuser field to given value.

### HasBEzsignfoldertypeSendattachmentuser

`func (o *EzsignfoldertypeResponseCompoundV4) HasBEzsignfoldertypeSendattachmentuser() bool`

HasBEzsignfoldertypeSendattachmentuser returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendproofuser

`func (o *EzsignfoldertypeResponseCompoundV4) GetBEzsignfoldertypeSendproofuser() bool`

GetBEzsignfoldertypeSendproofuser returns the BEzsignfoldertypeSendproofuser field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendproofuserOk

`func (o *EzsignfoldertypeResponseCompoundV4) GetBEzsignfoldertypeSendproofuserOk() (*bool, bool)`

GetBEzsignfoldertypeSendproofuserOk returns a tuple with the BEzsignfoldertypeSendproofuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendproofuser

`func (o *EzsignfoldertypeResponseCompoundV4) SetBEzsignfoldertypeSendproofuser(v bool)`

SetBEzsignfoldertypeSendproofuser sets BEzsignfoldertypeSendproofuser field to given value.

### HasBEzsignfoldertypeSendproofuser

`func (o *EzsignfoldertypeResponseCompoundV4) HasBEzsignfoldertypeSendproofuser() bool`

HasBEzsignfoldertypeSendproofuser returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendproofemail

`func (o *EzsignfoldertypeResponseCompoundV4) GetBEzsignfoldertypeSendproofemail() bool`

GetBEzsignfoldertypeSendproofemail returns the BEzsignfoldertypeSendproofemail field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendproofemailOk

`func (o *EzsignfoldertypeResponseCompoundV4) GetBEzsignfoldertypeSendproofemailOk() (*bool, bool)`

GetBEzsignfoldertypeSendproofemailOk returns a tuple with the BEzsignfoldertypeSendproofemail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendproofemail

`func (o *EzsignfoldertypeResponseCompoundV4) SetBEzsignfoldertypeSendproofemail(v bool)`

SetBEzsignfoldertypeSendproofemail sets BEzsignfoldertypeSendproofemail field to given value.

### HasBEzsignfoldertypeSendproofemail

`func (o *EzsignfoldertypeResponseCompoundV4) HasBEzsignfoldertypeSendproofemail() bool`

HasBEzsignfoldertypeSendproofemail returns a boolean if a field has been set.

### GetBEzsignfoldertypeAllowdownloadattachmentezsignsigner

`func (o *EzsignfoldertypeResponseCompoundV4) GetBEzsignfoldertypeAllowdownloadattachmentezsignsigner() bool`

GetBEzsignfoldertypeAllowdownloadattachmentezsignsigner returns the BEzsignfoldertypeAllowdownloadattachmentezsignsigner field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeAllowdownloadattachmentezsignsignerOk

`func (o *EzsignfoldertypeResponseCompoundV4) GetBEzsignfoldertypeAllowdownloadattachmentezsignsignerOk() (*bool, bool)`

GetBEzsignfoldertypeAllowdownloadattachmentezsignsignerOk returns a tuple with the BEzsignfoldertypeAllowdownloadattachmentezsignsigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeAllowdownloadattachmentezsignsigner

`func (o *EzsignfoldertypeResponseCompoundV4) SetBEzsignfoldertypeAllowdownloadattachmentezsignsigner(v bool)`

SetBEzsignfoldertypeAllowdownloadattachmentezsignsigner sets BEzsignfoldertypeAllowdownloadattachmentezsignsigner field to given value.

### HasBEzsignfoldertypeAllowdownloadattachmentezsignsigner

`func (o *EzsignfoldertypeResponseCompoundV4) HasBEzsignfoldertypeAllowdownloadattachmentezsignsigner() bool`

HasBEzsignfoldertypeAllowdownloadattachmentezsignsigner returns a boolean if a field has been set.

### GetBEzsignfoldertypeAllowdownloadproofezsignsigner

`func (o *EzsignfoldertypeResponseCompoundV4) GetBEzsignfoldertypeAllowdownloadproofezsignsigner() bool`

GetBEzsignfoldertypeAllowdownloadproofezsignsigner returns the BEzsignfoldertypeAllowdownloadproofezsignsigner field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeAllowdownloadproofezsignsignerOk

`func (o *EzsignfoldertypeResponseCompoundV4) GetBEzsignfoldertypeAllowdownloadproofezsignsignerOk() (*bool, bool)`

GetBEzsignfoldertypeAllowdownloadproofezsignsignerOk returns a tuple with the BEzsignfoldertypeAllowdownloadproofezsignsigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeAllowdownloadproofezsignsigner

`func (o *EzsignfoldertypeResponseCompoundV4) SetBEzsignfoldertypeAllowdownloadproofezsignsigner(v bool)`

SetBEzsignfoldertypeAllowdownloadproofezsignsigner sets BEzsignfoldertypeAllowdownloadproofezsignsigner field to given value.

### HasBEzsignfoldertypeAllowdownloadproofezsignsigner

`func (o *EzsignfoldertypeResponseCompoundV4) HasBEzsignfoldertypeAllowdownloadproofezsignsigner() bool`

HasBEzsignfoldertypeAllowdownloadproofezsignsigner returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendproofreceivealldocument

`func (o *EzsignfoldertypeResponseCompoundV4) GetBEzsignfoldertypeSendproofreceivealldocument() bool`

GetBEzsignfoldertypeSendproofreceivealldocument returns the BEzsignfoldertypeSendproofreceivealldocument field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendproofreceivealldocumentOk

`func (o *EzsignfoldertypeResponseCompoundV4) GetBEzsignfoldertypeSendproofreceivealldocumentOk() (*bool, bool)`

GetBEzsignfoldertypeSendproofreceivealldocumentOk returns a tuple with the BEzsignfoldertypeSendproofreceivealldocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendproofreceivealldocument

`func (o *EzsignfoldertypeResponseCompoundV4) SetBEzsignfoldertypeSendproofreceivealldocument(v bool)`

SetBEzsignfoldertypeSendproofreceivealldocument sets BEzsignfoldertypeSendproofreceivealldocument field to given value.

### HasBEzsignfoldertypeSendproofreceivealldocument

`func (o *EzsignfoldertypeResponseCompoundV4) HasBEzsignfoldertypeSendproofreceivealldocument() bool`

HasBEzsignfoldertypeSendproofreceivealldocument returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendsignedtodocumentowner

`func (o *EzsignfoldertypeResponseCompoundV4) GetBEzsignfoldertypeSendsignedtodocumentowner() bool`

GetBEzsignfoldertypeSendsignedtodocumentowner returns the BEzsignfoldertypeSendsignedtodocumentowner field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsignedtodocumentownerOk

`func (o *EzsignfoldertypeResponseCompoundV4) GetBEzsignfoldertypeSendsignedtodocumentownerOk() (*bool, bool)`

GetBEzsignfoldertypeSendsignedtodocumentownerOk returns a tuple with the BEzsignfoldertypeSendsignedtodocumentowner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsignedtodocumentowner

`func (o *EzsignfoldertypeResponseCompoundV4) SetBEzsignfoldertypeSendsignedtodocumentowner(v bool)`

SetBEzsignfoldertypeSendsignedtodocumentowner sets BEzsignfoldertypeSendsignedtodocumentowner field to given value.


### GetBEzsignfoldertypeSendsignedtofolderowner

`func (o *EzsignfoldertypeResponseCompoundV4) GetBEzsignfoldertypeSendsignedtofolderowner() bool`

GetBEzsignfoldertypeSendsignedtofolderowner returns the BEzsignfoldertypeSendsignedtofolderowner field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsignedtofolderownerOk

`func (o *EzsignfoldertypeResponseCompoundV4) GetBEzsignfoldertypeSendsignedtofolderownerOk() (*bool, bool)`

GetBEzsignfoldertypeSendsignedtofolderownerOk returns a tuple with the BEzsignfoldertypeSendsignedtofolderowner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsignedtofolderowner

`func (o *EzsignfoldertypeResponseCompoundV4) SetBEzsignfoldertypeSendsignedtofolderowner(v bool)`

SetBEzsignfoldertypeSendsignedtofolderowner sets BEzsignfoldertypeSendsignedtofolderowner field to given value.


### GetBEzsignfoldertypeSendsignedtofullgroup

`func (o *EzsignfoldertypeResponseCompoundV4) GetBEzsignfoldertypeSendsignedtofullgroup() bool`

GetBEzsignfoldertypeSendsignedtofullgroup returns the BEzsignfoldertypeSendsignedtofullgroup field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsignedtofullgroupOk

`func (o *EzsignfoldertypeResponseCompoundV4) GetBEzsignfoldertypeSendsignedtofullgroupOk() (*bool, bool)`

GetBEzsignfoldertypeSendsignedtofullgroupOk returns a tuple with the BEzsignfoldertypeSendsignedtofullgroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsignedtofullgroup

`func (o *EzsignfoldertypeResponseCompoundV4) SetBEzsignfoldertypeSendsignedtofullgroup(v bool)`

SetBEzsignfoldertypeSendsignedtofullgroup sets BEzsignfoldertypeSendsignedtofullgroup field to given value.

### HasBEzsignfoldertypeSendsignedtofullgroup

`func (o *EzsignfoldertypeResponseCompoundV4) HasBEzsignfoldertypeSendsignedtofullgroup() bool`

HasBEzsignfoldertypeSendsignedtofullgroup returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendsignedtolimitedgroup

`func (o *EzsignfoldertypeResponseCompoundV4) GetBEzsignfoldertypeSendsignedtolimitedgroup() bool`

GetBEzsignfoldertypeSendsignedtolimitedgroup returns the BEzsignfoldertypeSendsignedtolimitedgroup field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsignedtolimitedgroupOk

`func (o *EzsignfoldertypeResponseCompoundV4) GetBEzsignfoldertypeSendsignedtolimitedgroupOk() (*bool, bool)`

GetBEzsignfoldertypeSendsignedtolimitedgroupOk returns a tuple with the BEzsignfoldertypeSendsignedtolimitedgroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsignedtolimitedgroup

`func (o *EzsignfoldertypeResponseCompoundV4) SetBEzsignfoldertypeSendsignedtolimitedgroup(v bool)`

SetBEzsignfoldertypeSendsignedtolimitedgroup sets BEzsignfoldertypeSendsignedtolimitedgroup field to given value.

### HasBEzsignfoldertypeSendsignedtolimitedgroup

`func (o *EzsignfoldertypeResponseCompoundV4) HasBEzsignfoldertypeSendsignedtolimitedgroup() bool`

HasBEzsignfoldertypeSendsignedtolimitedgroup returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendsignedtocolleague

`func (o *EzsignfoldertypeResponseCompoundV4) GetBEzsignfoldertypeSendsignedtocolleague() bool`

GetBEzsignfoldertypeSendsignedtocolleague returns the BEzsignfoldertypeSendsignedtocolleague field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsignedtocolleagueOk

`func (o *EzsignfoldertypeResponseCompoundV4) GetBEzsignfoldertypeSendsignedtocolleagueOk() (*bool, bool)`

GetBEzsignfoldertypeSendsignedtocolleagueOk returns a tuple with the BEzsignfoldertypeSendsignedtocolleague field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsignedtocolleague

`func (o *EzsignfoldertypeResponseCompoundV4) SetBEzsignfoldertypeSendsignedtocolleague(v bool)`

SetBEzsignfoldertypeSendsignedtocolleague sets BEzsignfoldertypeSendsignedtocolleague field to given value.


### GetBEzsignfoldertypeSendsummarytodocumentowner

`func (o *EzsignfoldertypeResponseCompoundV4) GetBEzsignfoldertypeSendsummarytodocumentowner() bool`

GetBEzsignfoldertypeSendsummarytodocumentowner returns the BEzsignfoldertypeSendsummarytodocumentowner field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsummarytodocumentownerOk

`func (o *EzsignfoldertypeResponseCompoundV4) GetBEzsignfoldertypeSendsummarytodocumentownerOk() (*bool, bool)`

GetBEzsignfoldertypeSendsummarytodocumentownerOk returns a tuple with the BEzsignfoldertypeSendsummarytodocumentowner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsummarytodocumentowner

`func (o *EzsignfoldertypeResponseCompoundV4) SetBEzsignfoldertypeSendsummarytodocumentowner(v bool)`

SetBEzsignfoldertypeSendsummarytodocumentowner sets BEzsignfoldertypeSendsummarytodocumentowner field to given value.


### GetBEzsignfoldertypeSendsummarytofolderowner

`func (o *EzsignfoldertypeResponseCompoundV4) GetBEzsignfoldertypeSendsummarytofolderowner() bool`

GetBEzsignfoldertypeSendsummarytofolderowner returns the BEzsignfoldertypeSendsummarytofolderowner field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsummarytofolderownerOk

`func (o *EzsignfoldertypeResponseCompoundV4) GetBEzsignfoldertypeSendsummarytofolderownerOk() (*bool, bool)`

GetBEzsignfoldertypeSendsummarytofolderownerOk returns a tuple with the BEzsignfoldertypeSendsummarytofolderowner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsummarytofolderowner

`func (o *EzsignfoldertypeResponseCompoundV4) SetBEzsignfoldertypeSendsummarytofolderowner(v bool)`

SetBEzsignfoldertypeSendsummarytofolderowner sets BEzsignfoldertypeSendsummarytofolderowner field to given value.


### GetBEzsignfoldertypeSendsummarytofullgroup

`func (o *EzsignfoldertypeResponseCompoundV4) GetBEzsignfoldertypeSendsummarytofullgroup() bool`

GetBEzsignfoldertypeSendsummarytofullgroup returns the BEzsignfoldertypeSendsummarytofullgroup field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsummarytofullgroupOk

`func (o *EzsignfoldertypeResponseCompoundV4) GetBEzsignfoldertypeSendsummarytofullgroupOk() (*bool, bool)`

GetBEzsignfoldertypeSendsummarytofullgroupOk returns a tuple with the BEzsignfoldertypeSendsummarytofullgroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsummarytofullgroup

`func (o *EzsignfoldertypeResponseCompoundV4) SetBEzsignfoldertypeSendsummarytofullgroup(v bool)`

SetBEzsignfoldertypeSendsummarytofullgroup sets BEzsignfoldertypeSendsummarytofullgroup field to given value.

### HasBEzsignfoldertypeSendsummarytofullgroup

`func (o *EzsignfoldertypeResponseCompoundV4) HasBEzsignfoldertypeSendsummarytofullgroup() bool`

HasBEzsignfoldertypeSendsummarytofullgroup returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendsummarytolimitedgroup

`func (o *EzsignfoldertypeResponseCompoundV4) GetBEzsignfoldertypeSendsummarytolimitedgroup() bool`

GetBEzsignfoldertypeSendsummarytolimitedgroup returns the BEzsignfoldertypeSendsummarytolimitedgroup field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsummarytolimitedgroupOk

`func (o *EzsignfoldertypeResponseCompoundV4) GetBEzsignfoldertypeSendsummarytolimitedgroupOk() (*bool, bool)`

GetBEzsignfoldertypeSendsummarytolimitedgroupOk returns a tuple with the BEzsignfoldertypeSendsummarytolimitedgroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsummarytolimitedgroup

`func (o *EzsignfoldertypeResponseCompoundV4) SetBEzsignfoldertypeSendsummarytolimitedgroup(v bool)`

SetBEzsignfoldertypeSendsummarytolimitedgroup sets BEzsignfoldertypeSendsummarytolimitedgroup field to given value.

### HasBEzsignfoldertypeSendsummarytolimitedgroup

`func (o *EzsignfoldertypeResponseCompoundV4) HasBEzsignfoldertypeSendsummarytolimitedgroup() bool`

HasBEzsignfoldertypeSendsummarytolimitedgroup returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendsummarytocolleague

`func (o *EzsignfoldertypeResponseCompoundV4) GetBEzsignfoldertypeSendsummarytocolleague() bool`

GetBEzsignfoldertypeSendsummarytocolleague returns the BEzsignfoldertypeSendsummarytocolleague field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsummarytocolleagueOk

`func (o *EzsignfoldertypeResponseCompoundV4) GetBEzsignfoldertypeSendsummarytocolleagueOk() (*bool, bool)`

GetBEzsignfoldertypeSendsummarytocolleagueOk returns a tuple with the BEzsignfoldertypeSendsummarytocolleague field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsummarytocolleague

`func (o *EzsignfoldertypeResponseCompoundV4) SetBEzsignfoldertypeSendsummarytocolleague(v bool)`

SetBEzsignfoldertypeSendsummarytocolleague sets BEzsignfoldertypeSendsummarytocolleague field to given value.


### GetEEzsignfoldertypeSigneraccess

`func (o *EzsignfoldertypeResponseCompoundV4) GetEEzsignfoldertypeSigneraccess() FieldEEzsignfoldertypeSigneraccess`

GetEEzsignfoldertypeSigneraccess returns the EEzsignfoldertypeSigneraccess field if non-nil, zero value otherwise.

### GetEEzsignfoldertypeSigneraccessOk

`func (o *EzsignfoldertypeResponseCompoundV4) GetEEzsignfoldertypeSigneraccessOk() (*FieldEEzsignfoldertypeSigneraccess, bool)`

GetEEzsignfoldertypeSigneraccessOk returns a tuple with the EEzsignfoldertypeSigneraccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignfoldertypeSigneraccess

`func (o *EzsignfoldertypeResponseCompoundV4) SetEEzsignfoldertypeSigneraccess(v FieldEEzsignfoldertypeSigneraccess)`

SetEEzsignfoldertypeSigneraccess sets EEzsignfoldertypeSigneraccess field to given value.

### HasEEzsignfoldertypeSigneraccess

`func (o *EzsignfoldertypeResponseCompoundV4) HasEEzsignfoldertypeSigneraccess() bool`

HasEEzsignfoldertypeSigneraccess returns a boolean if a field has been set.

### GetBEzsignfoldertypeIsactive

`func (o *EzsignfoldertypeResponseCompoundV4) GetBEzsignfoldertypeIsactive() bool`

GetBEzsignfoldertypeIsactive returns the BEzsignfoldertypeIsactive field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeIsactiveOk

`func (o *EzsignfoldertypeResponseCompoundV4) GetBEzsignfoldertypeIsactiveOk() (*bool, bool)`

GetBEzsignfoldertypeIsactiveOk returns a tuple with the BEzsignfoldertypeIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeIsactive

`func (o *EzsignfoldertypeResponseCompoundV4) SetBEzsignfoldertypeIsactive(v bool)`

SetBEzsignfoldertypeIsactive sets BEzsignfoldertypeIsactive field to given value.


### GetAFkiPdfalevelID

`func (o *EzsignfoldertypeResponseCompoundV4) GetAFkiPdfalevelID() []int32`

GetAFkiPdfalevelID returns the AFkiPdfalevelID field if non-nil, zero value otherwise.

### GetAFkiPdfalevelIDOk

`func (o *EzsignfoldertypeResponseCompoundV4) GetAFkiPdfalevelIDOk() (*[]int32, bool)`

GetAFkiPdfalevelIDOk returns a tuple with the AFkiPdfalevelID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAFkiPdfalevelID

`func (o *EzsignfoldertypeResponseCompoundV4) SetAFkiPdfalevelID(v []int32)`

SetAFkiPdfalevelID sets AFkiPdfalevelID field to given value.

### HasAFkiPdfalevelID

`func (o *EzsignfoldertypeResponseCompoundV4) HasAFkiPdfalevelID() bool`

HasAFkiPdfalevelID returns a boolean if a field has been set.

### GetAObjUserlogintype

`func (o *EzsignfoldertypeResponseCompoundV4) GetAObjUserlogintype() []UserlogintypeResponse`

GetAObjUserlogintype returns the AObjUserlogintype field if non-nil, zero value otherwise.

### GetAObjUserlogintypeOk

`func (o *EzsignfoldertypeResponseCompoundV4) GetAObjUserlogintypeOk() (*[]UserlogintypeResponse, bool)`

GetAObjUserlogintypeOk returns a tuple with the AObjUserlogintype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjUserlogintype

`func (o *EzsignfoldertypeResponseCompoundV4) SetAObjUserlogintype(v []UserlogintypeResponse)`

SetAObjUserlogintype sets AObjUserlogintype field to given value.


### GetAObjUsergroupAll

`func (o *EzsignfoldertypeResponseCompoundV4) GetAObjUsergroupAll() []UsergroupResponse`

GetAObjUsergroupAll returns the AObjUsergroupAll field if non-nil, zero value otherwise.

### GetAObjUsergroupAllOk

`func (o *EzsignfoldertypeResponseCompoundV4) GetAObjUsergroupAllOk() (*[]UsergroupResponse, bool)`

GetAObjUsergroupAllOk returns a tuple with the AObjUsergroupAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjUsergroupAll

`func (o *EzsignfoldertypeResponseCompoundV4) SetAObjUsergroupAll(v []UsergroupResponse)`

SetAObjUsergroupAll sets AObjUsergroupAll field to given value.

### HasAObjUsergroupAll

`func (o *EzsignfoldertypeResponseCompoundV4) HasAObjUsergroupAll() bool`

HasAObjUsergroupAll returns a boolean if a field has been set.

### GetAObjUsergroupRestricted

`func (o *EzsignfoldertypeResponseCompoundV4) GetAObjUsergroupRestricted() []UsergroupResponse`

GetAObjUsergroupRestricted returns the AObjUsergroupRestricted field if non-nil, zero value otherwise.

### GetAObjUsergroupRestrictedOk

`func (o *EzsignfoldertypeResponseCompoundV4) GetAObjUsergroupRestrictedOk() (*[]UsergroupResponse, bool)`

GetAObjUsergroupRestrictedOk returns a tuple with the AObjUsergroupRestricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjUsergroupRestricted

`func (o *EzsignfoldertypeResponseCompoundV4) SetAObjUsergroupRestricted(v []UsergroupResponse)`

SetAObjUsergroupRestricted sets AObjUsergroupRestricted field to given value.

### HasAObjUsergroupRestricted

`func (o *EzsignfoldertypeResponseCompoundV4) HasAObjUsergroupRestricted() bool`

HasAObjUsergroupRestricted returns a boolean if a field has been set.

### GetAObjUsergroupTemplate

`func (o *EzsignfoldertypeResponseCompoundV4) GetAObjUsergroupTemplate() []UsergroupResponse`

GetAObjUsergroupTemplate returns the AObjUsergroupTemplate field if non-nil, zero value otherwise.

### GetAObjUsergroupTemplateOk

`func (o *EzsignfoldertypeResponseCompoundV4) GetAObjUsergroupTemplateOk() (*[]UsergroupResponse, bool)`

GetAObjUsergroupTemplateOk returns a tuple with the AObjUsergroupTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjUsergroupTemplate

`func (o *EzsignfoldertypeResponseCompoundV4) SetAObjUsergroupTemplate(v []UsergroupResponse)`

SetAObjUsergroupTemplate sets AObjUsergroupTemplate field to given value.

### HasAObjUsergroupTemplate

`func (o *EzsignfoldertypeResponseCompoundV4) HasAObjUsergroupTemplate() bool`

HasAObjUsergroupTemplate returns a boolean if a field has been set.

### GetObjAudit

`func (o *EzsignfoldertypeResponseCompoundV4) GetObjAudit() CommonAudit`

GetObjAudit returns the ObjAudit field if non-nil, zero value otherwise.

### GetObjAuditOk

`func (o *EzsignfoldertypeResponseCompoundV4) GetObjAuditOk() (*CommonAudit, bool)`

GetObjAuditOk returns a tuple with the ObjAudit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjAudit

`func (o *EzsignfoldertypeResponseCompoundV4) SetObjAudit(v CommonAudit)`

SetObjAudit sets ObjAudit field to given value.


### GetAFkiUserIDSigned

`func (o *EzsignfoldertypeResponseCompoundV4) GetAFkiUserIDSigned() []int32`

GetAFkiUserIDSigned returns the AFkiUserIDSigned field if non-nil, zero value otherwise.

### GetAFkiUserIDSignedOk

`func (o *EzsignfoldertypeResponseCompoundV4) GetAFkiUserIDSignedOk() (*[]int32, bool)`

GetAFkiUserIDSignedOk returns a tuple with the AFkiUserIDSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAFkiUserIDSigned

`func (o *EzsignfoldertypeResponseCompoundV4) SetAFkiUserIDSigned(v []int32)`

SetAFkiUserIDSigned sets AFkiUserIDSigned field to given value.

### HasAFkiUserIDSigned

`func (o *EzsignfoldertypeResponseCompoundV4) HasAFkiUserIDSigned() bool`

HasAFkiUserIDSigned returns a boolean if a field has been set.

### GetAFkiUserIDSummary

`func (o *EzsignfoldertypeResponseCompoundV4) GetAFkiUserIDSummary() []int32`

GetAFkiUserIDSummary returns the AFkiUserIDSummary field if non-nil, zero value otherwise.

### GetAFkiUserIDSummaryOk

`func (o *EzsignfoldertypeResponseCompoundV4) GetAFkiUserIDSummaryOk() (*[]int32, bool)`

GetAFkiUserIDSummaryOk returns a tuple with the AFkiUserIDSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAFkiUserIDSummary

`func (o *EzsignfoldertypeResponseCompoundV4) SetAFkiUserIDSummary(v []int32)`

SetAFkiUserIDSummary sets AFkiUserIDSummary field to given value.

### HasAFkiUserIDSummary

`func (o *EzsignfoldertypeResponseCompoundV4) HasAFkiUserIDSummary() bool`

HasAFkiUserIDSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


