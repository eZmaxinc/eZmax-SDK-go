# EzsignfoldertypeRequestV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsignfoldertypeID** | Pointer to **int32** | The unique ID of the Ezsignfoldertype. | [optional] 
**ObjEzsignfoldertypeName** | [**MultilingualEzsignfoldertypeName**](MultilingualEzsignfoldertypeName.md) |  | 
**FkiBrandingID** | **int32** | The unique ID of the Branding | 
**FkiBillingentityinternalID** | Pointer to **int32** | The unique ID of the Billingentityinternal. | [optional] 
**FkiEzsigntsarequirementID** | Pointer to **int32** | The unique ID of the Ezsigntsarequirement.  Determine if a Time Stamping Authority should add a timestamp on each of the signature. Valid values:  |Value|Description| |-|-| |1|No. TSA Timestamping will requested. This will make all signatures a lot faster since no round-trip to the TSA server will be required. Timestamping will be made using eZsign server&#39;s time.| |2|Best effort. Timestamping from a Time Stamping Authority will be requested but is not mandatory. In the very improbable case it cannot be completed, the timestamping will be made using eZsign server&#39;s time. **Additional fee applies**| |3|Mandatory. Timestamping from a Time Stamping Authority will be requested and is mandatory. In the very improbable case it cannot be completed, the signature will fail and the user will be asked to retry. **Additional fee applies**| | [optional] 
**FkiFontIDAnnotation** | Pointer to **int32** | The unique ID of the Font | [optional] 
**FkiFontIDFormfield** | Pointer to **int32** | The unique ID of the Font | [optional] 
**FkiFontIDSignature** | Pointer to **int32** | The unique ID of the Font | [optional] 
**FkiPdfalevelIDConvert** | Pointer to **int32** | The unique ID of the Pdfalevel | [optional] 
**AFkiPdfalevelID** | Pointer to **[]int32** |  | [optional] 
**AFkiUserlogintypeID** | **[]int32** |  | 
**AFkiUsergroupIDAll** | Pointer to **[]int32** |  | [optional] 
**AFkiUsergroupIDRestricted** | Pointer to **[]int32** |  | [optional] 
**AFkiUsergroupIDTemplate** | Pointer to **[]int32** |  | [optional] 
**EEzsignfoldertypeDocumentdependency** | Pointer to [**FieldEEzsignfoldertypeDocumentdependency**](FieldEEzsignfoldertypeDocumentdependency.md) |  | [optional] 
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
**BEzsignfoldertypeSendattachmentreceivecopy** | Pointer to **bool** | Whether we send the Ezsigndocument in the email to Ezsignsigner or User when bEzsignfoldersignerassociationReceivecopy &#x3D; 1 | [optional] 
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

## Methods

### NewEzsignfoldertypeRequestV3

`func NewEzsignfoldertypeRequestV3(objEzsignfoldertypeName MultilingualEzsignfoldertypeName, fkiBrandingID int32, aFkiUserlogintypeID []int32, eEzsignfoldertypePrivacylevel FieldEEzsignfoldertypePrivacylevel, iEzsignfoldertypeArchivaldays int32, eEzsignfoldertypeDisposal FieldEEzsignfoldertypeDisposal, eEzsignfoldertypeCompletion FieldEEzsignfoldertypeCompletion, iEzsignfoldertypeDeadlinedays int32, bEzsignfoldertypeSendsignedtodocumentowner bool, bEzsignfoldertypeSendsignedtofolderowner bool, bEzsignfoldertypeSendsignedtocolleague bool, bEzsignfoldertypeSendsummarytodocumentowner bool, bEzsignfoldertypeSendsummarytofolderowner bool, bEzsignfoldertypeSendsummarytocolleague bool, bEzsignfoldertypeIsactive bool, ) *EzsignfoldertypeRequestV3`

NewEzsignfoldertypeRequestV3 instantiates a new EzsignfoldertypeRequestV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignfoldertypeRequestV3WithDefaults

`func NewEzsignfoldertypeRequestV3WithDefaults() *EzsignfoldertypeRequestV3`

NewEzsignfoldertypeRequestV3WithDefaults instantiates a new EzsignfoldertypeRequestV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsignfoldertypeID

`func (o *EzsignfoldertypeRequestV3) GetPkiEzsignfoldertypeID() int32`

GetPkiEzsignfoldertypeID returns the PkiEzsignfoldertypeID field if non-nil, zero value otherwise.

### GetPkiEzsignfoldertypeIDOk

`func (o *EzsignfoldertypeRequestV3) GetPkiEzsignfoldertypeIDOk() (*int32, bool)`

GetPkiEzsignfoldertypeIDOk returns a tuple with the PkiEzsignfoldertypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsignfoldertypeID

`func (o *EzsignfoldertypeRequestV3) SetPkiEzsignfoldertypeID(v int32)`

SetPkiEzsignfoldertypeID sets PkiEzsignfoldertypeID field to given value.

### HasPkiEzsignfoldertypeID

`func (o *EzsignfoldertypeRequestV3) HasPkiEzsignfoldertypeID() bool`

HasPkiEzsignfoldertypeID returns a boolean if a field has been set.

### GetObjEzsignfoldertypeName

`func (o *EzsignfoldertypeRequestV3) GetObjEzsignfoldertypeName() MultilingualEzsignfoldertypeName`

GetObjEzsignfoldertypeName returns the ObjEzsignfoldertypeName field if non-nil, zero value otherwise.

### GetObjEzsignfoldertypeNameOk

`func (o *EzsignfoldertypeRequestV3) GetObjEzsignfoldertypeNameOk() (*MultilingualEzsignfoldertypeName, bool)`

GetObjEzsignfoldertypeNameOk returns a tuple with the ObjEzsignfoldertypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjEzsignfoldertypeName

`func (o *EzsignfoldertypeRequestV3) SetObjEzsignfoldertypeName(v MultilingualEzsignfoldertypeName)`

SetObjEzsignfoldertypeName sets ObjEzsignfoldertypeName field to given value.


### GetFkiBrandingID

`func (o *EzsignfoldertypeRequestV3) GetFkiBrandingID() int32`

GetFkiBrandingID returns the FkiBrandingID field if non-nil, zero value otherwise.

### GetFkiBrandingIDOk

`func (o *EzsignfoldertypeRequestV3) GetFkiBrandingIDOk() (*int32, bool)`

GetFkiBrandingIDOk returns a tuple with the FkiBrandingID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiBrandingID

`func (o *EzsignfoldertypeRequestV3) SetFkiBrandingID(v int32)`

SetFkiBrandingID sets FkiBrandingID field to given value.


### GetFkiBillingentityinternalID

`func (o *EzsignfoldertypeRequestV3) GetFkiBillingentityinternalID() int32`

GetFkiBillingentityinternalID returns the FkiBillingentityinternalID field if non-nil, zero value otherwise.

### GetFkiBillingentityinternalIDOk

`func (o *EzsignfoldertypeRequestV3) GetFkiBillingentityinternalIDOk() (*int32, bool)`

GetFkiBillingentityinternalIDOk returns a tuple with the FkiBillingentityinternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiBillingentityinternalID

`func (o *EzsignfoldertypeRequestV3) SetFkiBillingentityinternalID(v int32)`

SetFkiBillingentityinternalID sets FkiBillingentityinternalID field to given value.

### HasFkiBillingentityinternalID

`func (o *EzsignfoldertypeRequestV3) HasFkiBillingentityinternalID() bool`

HasFkiBillingentityinternalID returns a boolean if a field has been set.

### GetFkiEzsigntsarequirementID

`func (o *EzsignfoldertypeRequestV3) GetFkiEzsigntsarequirementID() int32`

GetFkiEzsigntsarequirementID returns the FkiEzsigntsarequirementID field if non-nil, zero value otherwise.

### GetFkiEzsigntsarequirementIDOk

`func (o *EzsignfoldertypeRequestV3) GetFkiEzsigntsarequirementIDOk() (*int32, bool)`

GetFkiEzsigntsarequirementIDOk returns a tuple with the FkiEzsigntsarequirementID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntsarequirementID

`func (o *EzsignfoldertypeRequestV3) SetFkiEzsigntsarequirementID(v int32)`

SetFkiEzsigntsarequirementID sets FkiEzsigntsarequirementID field to given value.

### HasFkiEzsigntsarequirementID

`func (o *EzsignfoldertypeRequestV3) HasFkiEzsigntsarequirementID() bool`

HasFkiEzsigntsarequirementID returns a boolean if a field has been set.

### GetFkiFontIDAnnotation

`func (o *EzsignfoldertypeRequestV3) GetFkiFontIDAnnotation() int32`

GetFkiFontIDAnnotation returns the FkiFontIDAnnotation field if non-nil, zero value otherwise.

### GetFkiFontIDAnnotationOk

`func (o *EzsignfoldertypeRequestV3) GetFkiFontIDAnnotationOk() (*int32, bool)`

GetFkiFontIDAnnotationOk returns a tuple with the FkiFontIDAnnotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiFontIDAnnotation

`func (o *EzsignfoldertypeRequestV3) SetFkiFontIDAnnotation(v int32)`

SetFkiFontIDAnnotation sets FkiFontIDAnnotation field to given value.

### HasFkiFontIDAnnotation

`func (o *EzsignfoldertypeRequestV3) HasFkiFontIDAnnotation() bool`

HasFkiFontIDAnnotation returns a boolean if a field has been set.

### GetFkiFontIDFormfield

`func (o *EzsignfoldertypeRequestV3) GetFkiFontIDFormfield() int32`

GetFkiFontIDFormfield returns the FkiFontIDFormfield field if non-nil, zero value otherwise.

### GetFkiFontIDFormfieldOk

`func (o *EzsignfoldertypeRequestV3) GetFkiFontIDFormfieldOk() (*int32, bool)`

GetFkiFontIDFormfieldOk returns a tuple with the FkiFontIDFormfield field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiFontIDFormfield

`func (o *EzsignfoldertypeRequestV3) SetFkiFontIDFormfield(v int32)`

SetFkiFontIDFormfield sets FkiFontIDFormfield field to given value.

### HasFkiFontIDFormfield

`func (o *EzsignfoldertypeRequestV3) HasFkiFontIDFormfield() bool`

HasFkiFontIDFormfield returns a boolean if a field has been set.

### GetFkiFontIDSignature

`func (o *EzsignfoldertypeRequestV3) GetFkiFontIDSignature() int32`

GetFkiFontIDSignature returns the FkiFontIDSignature field if non-nil, zero value otherwise.

### GetFkiFontIDSignatureOk

`func (o *EzsignfoldertypeRequestV3) GetFkiFontIDSignatureOk() (*int32, bool)`

GetFkiFontIDSignatureOk returns a tuple with the FkiFontIDSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiFontIDSignature

`func (o *EzsignfoldertypeRequestV3) SetFkiFontIDSignature(v int32)`

SetFkiFontIDSignature sets FkiFontIDSignature field to given value.

### HasFkiFontIDSignature

`func (o *EzsignfoldertypeRequestV3) HasFkiFontIDSignature() bool`

HasFkiFontIDSignature returns a boolean if a field has been set.

### GetFkiPdfalevelIDConvert

`func (o *EzsignfoldertypeRequestV3) GetFkiPdfalevelIDConvert() int32`

GetFkiPdfalevelIDConvert returns the FkiPdfalevelIDConvert field if non-nil, zero value otherwise.

### GetFkiPdfalevelIDConvertOk

`func (o *EzsignfoldertypeRequestV3) GetFkiPdfalevelIDConvertOk() (*int32, bool)`

GetFkiPdfalevelIDConvertOk returns a tuple with the FkiPdfalevelIDConvert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiPdfalevelIDConvert

`func (o *EzsignfoldertypeRequestV3) SetFkiPdfalevelIDConvert(v int32)`

SetFkiPdfalevelIDConvert sets FkiPdfalevelIDConvert field to given value.

### HasFkiPdfalevelIDConvert

`func (o *EzsignfoldertypeRequestV3) HasFkiPdfalevelIDConvert() bool`

HasFkiPdfalevelIDConvert returns a boolean if a field has been set.

### GetAFkiPdfalevelID

`func (o *EzsignfoldertypeRequestV3) GetAFkiPdfalevelID() []int32`

GetAFkiPdfalevelID returns the AFkiPdfalevelID field if non-nil, zero value otherwise.

### GetAFkiPdfalevelIDOk

`func (o *EzsignfoldertypeRequestV3) GetAFkiPdfalevelIDOk() (*[]int32, bool)`

GetAFkiPdfalevelIDOk returns a tuple with the AFkiPdfalevelID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAFkiPdfalevelID

`func (o *EzsignfoldertypeRequestV3) SetAFkiPdfalevelID(v []int32)`

SetAFkiPdfalevelID sets AFkiPdfalevelID field to given value.

### HasAFkiPdfalevelID

`func (o *EzsignfoldertypeRequestV3) HasAFkiPdfalevelID() bool`

HasAFkiPdfalevelID returns a boolean if a field has been set.

### GetAFkiUserlogintypeID

`func (o *EzsignfoldertypeRequestV3) GetAFkiUserlogintypeID() []int32`

GetAFkiUserlogintypeID returns the AFkiUserlogintypeID field if non-nil, zero value otherwise.

### GetAFkiUserlogintypeIDOk

`func (o *EzsignfoldertypeRequestV3) GetAFkiUserlogintypeIDOk() (*[]int32, bool)`

GetAFkiUserlogintypeIDOk returns a tuple with the AFkiUserlogintypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAFkiUserlogintypeID

`func (o *EzsignfoldertypeRequestV3) SetAFkiUserlogintypeID(v []int32)`

SetAFkiUserlogintypeID sets AFkiUserlogintypeID field to given value.


### GetAFkiUsergroupIDAll

`func (o *EzsignfoldertypeRequestV3) GetAFkiUsergroupIDAll() []int32`

GetAFkiUsergroupIDAll returns the AFkiUsergroupIDAll field if non-nil, zero value otherwise.

### GetAFkiUsergroupIDAllOk

`func (o *EzsignfoldertypeRequestV3) GetAFkiUsergroupIDAllOk() (*[]int32, bool)`

GetAFkiUsergroupIDAllOk returns a tuple with the AFkiUsergroupIDAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAFkiUsergroupIDAll

`func (o *EzsignfoldertypeRequestV3) SetAFkiUsergroupIDAll(v []int32)`

SetAFkiUsergroupIDAll sets AFkiUsergroupIDAll field to given value.

### HasAFkiUsergroupIDAll

`func (o *EzsignfoldertypeRequestV3) HasAFkiUsergroupIDAll() bool`

HasAFkiUsergroupIDAll returns a boolean if a field has been set.

### GetAFkiUsergroupIDRestricted

`func (o *EzsignfoldertypeRequestV3) GetAFkiUsergroupIDRestricted() []int32`

GetAFkiUsergroupIDRestricted returns the AFkiUsergroupIDRestricted field if non-nil, zero value otherwise.

### GetAFkiUsergroupIDRestrictedOk

`func (o *EzsignfoldertypeRequestV3) GetAFkiUsergroupIDRestrictedOk() (*[]int32, bool)`

GetAFkiUsergroupIDRestrictedOk returns a tuple with the AFkiUsergroupIDRestricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAFkiUsergroupIDRestricted

`func (o *EzsignfoldertypeRequestV3) SetAFkiUsergroupIDRestricted(v []int32)`

SetAFkiUsergroupIDRestricted sets AFkiUsergroupIDRestricted field to given value.

### HasAFkiUsergroupIDRestricted

`func (o *EzsignfoldertypeRequestV3) HasAFkiUsergroupIDRestricted() bool`

HasAFkiUsergroupIDRestricted returns a boolean if a field has been set.

### GetAFkiUsergroupIDTemplate

`func (o *EzsignfoldertypeRequestV3) GetAFkiUsergroupIDTemplate() []int32`

GetAFkiUsergroupIDTemplate returns the AFkiUsergroupIDTemplate field if non-nil, zero value otherwise.

### GetAFkiUsergroupIDTemplateOk

`func (o *EzsignfoldertypeRequestV3) GetAFkiUsergroupIDTemplateOk() (*[]int32, bool)`

GetAFkiUsergroupIDTemplateOk returns a tuple with the AFkiUsergroupIDTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAFkiUsergroupIDTemplate

`func (o *EzsignfoldertypeRequestV3) SetAFkiUsergroupIDTemplate(v []int32)`

SetAFkiUsergroupIDTemplate sets AFkiUsergroupIDTemplate field to given value.

### HasAFkiUsergroupIDTemplate

`func (o *EzsignfoldertypeRequestV3) HasAFkiUsergroupIDTemplate() bool`

HasAFkiUsergroupIDTemplate returns a boolean if a field has been set.

### GetEEzsignfoldertypeDocumentdependency

`func (o *EzsignfoldertypeRequestV3) GetEEzsignfoldertypeDocumentdependency() FieldEEzsignfoldertypeDocumentdependency`

GetEEzsignfoldertypeDocumentdependency returns the EEzsignfoldertypeDocumentdependency field if non-nil, zero value otherwise.

### GetEEzsignfoldertypeDocumentdependencyOk

`func (o *EzsignfoldertypeRequestV3) GetEEzsignfoldertypeDocumentdependencyOk() (*FieldEEzsignfoldertypeDocumentdependency, bool)`

GetEEzsignfoldertypeDocumentdependencyOk returns a tuple with the EEzsignfoldertypeDocumentdependency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignfoldertypeDocumentdependency

`func (o *EzsignfoldertypeRequestV3) SetEEzsignfoldertypeDocumentdependency(v FieldEEzsignfoldertypeDocumentdependency)`

SetEEzsignfoldertypeDocumentdependency sets EEzsignfoldertypeDocumentdependency field to given value.

### HasEEzsignfoldertypeDocumentdependency

`func (o *EzsignfoldertypeRequestV3) HasEEzsignfoldertypeDocumentdependency() bool`

HasEEzsignfoldertypeDocumentdependency returns a boolean if a field has been set.

### GetSEmailAddressSigned

`func (o *EzsignfoldertypeRequestV3) GetSEmailAddressSigned() string`

GetSEmailAddressSigned returns the SEmailAddressSigned field if non-nil, zero value otherwise.

### GetSEmailAddressSignedOk

`func (o *EzsignfoldertypeRequestV3) GetSEmailAddressSignedOk() (*string, bool)`

GetSEmailAddressSignedOk returns a tuple with the SEmailAddressSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEmailAddressSigned

`func (o *EzsignfoldertypeRequestV3) SetSEmailAddressSigned(v string)`

SetSEmailAddressSigned sets SEmailAddressSigned field to given value.

### HasSEmailAddressSigned

`func (o *EzsignfoldertypeRequestV3) HasSEmailAddressSigned() bool`

HasSEmailAddressSigned returns a boolean if a field has been set.

### GetSEmailAddressSummary

`func (o *EzsignfoldertypeRequestV3) GetSEmailAddressSummary() string`

GetSEmailAddressSummary returns the SEmailAddressSummary field if non-nil, zero value otherwise.

### GetSEmailAddressSummaryOk

`func (o *EzsignfoldertypeRequestV3) GetSEmailAddressSummaryOk() (*string, bool)`

GetSEmailAddressSummaryOk returns a tuple with the SEmailAddressSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEmailAddressSummary

`func (o *EzsignfoldertypeRequestV3) SetSEmailAddressSummary(v string)`

SetSEmailAddressSummary sets SEmailAddressSummary field to given value.

### HasSEmailAddressSummary

`func (o *EzsignfoldertypeRequestV3) HasSEmailAddressSummary() bool`

HasSEmailAddressSummary returns a boolean if a field has been set.

### GetEEzsignfoldertypePdfarequirement

`func (o *EzsignfoldertypeRequestV3) GetEEzsignfoldertypePdfarequirement() FieldEEzsignfoldertypePdfarequirement`

GetEEzsignfoldertypePdfarequirement returns the EEzsignfoldertypePdfarequirement field if non-nil, zero value otherwise.

### GetEEzsignfoldertypePdfarequirementOk

`func (o *EzsignfoldertypeRequestV3) GetEEzsignfoldertypePdfarequirementOk() (*FieldEEzsignfoldertypePdfarequirement, bool)`

GetEEzsignfoldertypePdfarequirementOk returns a tuple with the EEzsignfoldertypePdfarequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignfoldertypePdfarequirement

`func (o *EzsignfoldertypeRequestV3) SetEEzsignfoldertypePdfarequirement(v FieldEEzsignfoldertypePdfarequirement)`

SetEEzsignfoldertypePdfarequirement sets EEzsignfoldertypePdfarequirement field to given value.

### HasEEzsignfoldertypePdfarequirement

`func (o *EzsignfoldertypeRequestV3) HasEEzsignfoldertypePdfarequirement() bool`

HasEEzsignfoldertypePdfarequirement returns a boolean if a field has been set.

### GetEEzsignfoldertypePdfanoncompliantaction

`func (o *EzsignfoldertypeRequestV3) GetEEzsignfoldertypePdfanoncompliantaction() FieldEEzsignfoldertypePdfanoncompliantaction`

GetEEzsignfoldertypePdfanoncompliantaction returns the EEzsignfoldertypePdfanoncompliantaction field if non-nil, zero value otherwise.

### GetEEzsignfoldertypePdfanoncompliantactionOk

`func (o *EzsignfoldertypeRequestV3) GetEEzsignfoldertypePdfanoncompliantactionOk() (*FieldEEzsignfoldertypePdfanoncompliantaction, bool)`

GetEEzsignfoldertypePdfanoncompliantactionOk returns a tuple with the EEzsignfoldertypePdfanoncompliantaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignfoldertypePdfanoncompliantaction

`func (o *EzsignfoldertypeRequestV3) SetEEzsignfoldertypePdfanoncompliantaction(v FieldEEzsignfoldertypePdfanoncompliantaction)`

SetEEzsignfoldertypePdfanoncompliantaction sets EEzsignfoldertypePdfanoncompliantaction field to given value.

### HasEEzsignfoldertypePdfanoncompliantaction

`func (o *EzsignfoldertypeRequestV3) HasEEzsignfoldertypePdfanoncompliantaction() bool`

HasEEzsignfoldertypePdfanoncompliantaction returns a boolean if a field has been set.

### GetEEzsignfoldertypePrivacylevel

`func (o *EzsignfoldertypeRequestV3) GetEEzsignfoldertypePrivacylevel() FieldEEzsignfoldertypePrivacylevel`

GetEEzsignfoldertypePrivacylevel returns the EEzsignfoldertypePrivacylevel field if non-nil, zero value otherwise.

### GetEEzsignfoldertypePrivacylevelOk

`func (o *EzsignfoldertypeRequestV3) GetEEzsignfoldertypePrivacylevelOk() (*FieldEEzsignfoldertypePrivacylevel, bool)`

GetEEzsignfoldertypePrivacylevelOk returns a tuple with the EEzsignfoldertypePrivacylevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignfoldertypePrivacylevel

`func (o *EzsignfoldertypeRequestV3) SetEEzsignfoldertypePrivacylevel(v FieldEEzsignfoldertypePrivacylevel)`

SetEEzsignfoldertypePrivacylevel sets EEzsignfoldertypePrivacylevel field to given value.


### GetIEzsignfoldertypeFontsizeannotation

`func (o *EzsignfoldertypeRequestV3) GetIEzsignfoldertypeFontsizeannotation() int32`

GetIEzsignfoldertypeFontsizeannotation returns the IEzsignfoldertypeFontsizeannotation field if non-nil, zero value otherwise.

### GetIEzsignfoldertypeFontsizeannotationOk

`func (o *EzsignfoldertypeRequestV3) GetIEzsignfoldertypeFontsizeannotationOk() (*int32, bool)`

GetIEzsignfoldertypeFontsizeannotationOk returns a tuple with the IEzsignfoldertypeFontsizeannotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignfoldertypeFontsizeannotation

`func (o *EzsignfoldertypeRequestV3) SetIEzsignfoldertypeFontsizeannotation(v int32)`

SetIEzsignfoldertypeFontsizeannotation sets IEzsignfoldertypeFontsizeannotation field to given value.

### HasIEzsignfoldertypeFontsizeannotation

`func (o *EzsignfoldertypeRequestV3) HasIEzsignfoldertypeFontsizeannotation() bool`

HasIEzsignfoldertypeFontsizeannotation returns a boolean if a field has been set.

### GetIEzsignfoldertypeFontsizeformfield

`func (o *EzsignfoldertypeRequestV3) GetIEzsignfoldertypeFontsizeformfield() int32`

GetIEzsignfoldertypeFontsizeformfield returns the IEzsignfoldertypeFontsizeformfield field if non-nil, zero value otherwise.

### GetIEzsignfoldertypeFontsizeformfieldOk

`func (o *EzsignfoldertypeRequestV3) GetIEzsignfoldertypeFontsizeformfieldOk() (*int32, bool)`

GetIEzsignfoldertypeFontsizeformfieldOk returns a tuple with the IEzsignfoldertypeFontsizeformfield field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignfoldertypeFontsizeformfield

`func (o *EzsignfoldertypeRequestV3) SetIEzsignfoldertypeFontsizeformfield(v int32)`

SetIEzsignfoldertypeFontsizeformfield sets IEzsignfoldertypeFontsizeformfield field to given value.

### HasIEzsignfoldertypeFontsizeformfield

`func (o *EzsignfoldertypeRequestV3) HasIEzsignfoldertypeFontsizeformfield() bool`

HasIEzsignfoldertypeFontsizeformfield returns a boolean if a field has been set.

### GetIEzsignfoldertypeSendreminderfirstdays

`func (o *EzsignfoldertypeRequestV3) GetIEzsignfoldertypeSendreminderfirstdays() int32`

GetIEzsignfoldertypeSendreminderfirstdays returns the IEzsignfoldertypeSendreminderfirstdays field if non-nil, zero value otherwise.

### GetIEzsignfoldertypeSendreminderfirstdaysOk

`func (o *EzsignfoldertypeRequestV3) GetIEzsignfoldertypeSendreminderfirstdaysOk() (*int32, bool)`

GetIEzsignfoldertypeSendreminderfirstdaysOk returns a tuple with the IEzsignfoldertypeSendreminderfirstdays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignfoldertypeSendreminderfirstdays

`func (o *EzsignfoldertypeRequestV3) SetIEzsignfoldertypeSendreminderfirstdays(v int32)`

SetIEzsignfoldertypeSendreminderfirstdays sets IEzsignfoldertypeSendreminderfirstdays field to given value.

### HasIEzsignfoldertypeSendreminderfirstdays

`func (o *EzsignfoldertypeRequestV3) HasIEzsignfoldertypeSendreminderfirstdays() bool`

HasIEzsignfoldertypeSendreminderfirstdays returns a boolean if a field has been set.

### GetIEzsignfoldertypeSendreminderotherdays

`func (o *EzsignfoldertypeRequestV3) GetIEzsignfoldertypeSendreminderotherdays() int32`

GetIEzsignfoldertypeSendreminderotherdays returns the IEzsignfoldertypeSendreminderotherdays field if non-nil, zero value otherwise.

### GetIEzsignfoldertypeSendreminderotherdaysOk

`func (o *EzsignfoldertypeRequestV3) GetIEzsignfoldertypeSendreminderotherdaysOk() (*int32, bool)`

GetIEzsignfoldertypeSendreminderotherdaysOk returns a tuple with the IEzsignfoldertypeSendreminderotherdays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignfoldertypeSendreminderotherdays

`func (o *EzsignfoldertypeRequestV3) SetIEzsignfoldertypeSendreminderotherdays(v int32)`

SetIEzsignfoldertypeSendreminderotherdays sets IEzsignfoldertypeSendreminderotherdays field to given value.

### HasIEzsignfoldertypeSendreminderotherdays

`func (o *EzsignfoldertypeRequestV3) HasIEzsignfoldertypeSendreminderotherdays() bool`

HasIEzsignfoldertypeSendreminderotherdays returns a boolean if a field has been set.

### GetIEzsignfoldertypeArchivaldays

`func (o *EzsignfoldertypeRequestV3) GetIEzsignfoldertypeArchivaldays() int32`

GetIEzsignfoldertypeArchivaldays returns the IEzsignfoldertypeArchivaldays field if non-nil, zero value otherwise.

### GetIEzsignfoldertypeArchivaldaysOk

`func (o *EzsignfoldertypeRequestV3) GetIEzsignfoldertypeArchivaldaysOk() (*int32, bool)`

GetIEzsignfoldertypeArchivaldaysOk returns a tuple with the IEzsignfoldertypeArchivaldays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignfoldertypeArchivaldays

`func (o *EzsignfoldertypeRequestV3) SetIEzsignfoldertypeArchivaldays(v int32)`

SetIEzsignfoldertypeArchivaldays sets IEzsignfoldertypeArchivaldays field to given value.


### GetEEzsignfoldertypeDisposal

`func (o *EzsignfoldertypeRequestV3) GetEEzsignfoldertypeDisposal() FieldEEzsignfoldertypeDisposal`

GetEEzsignfoldertypeDisposal returns the EEzsignfoldertypeDisposal field if non-nil, zero value otherwise.

### GetEEzsignfoldertypeDisposalOk

`func (o *EzsignfoldertypeRequestV3) GetEEzsignfoldertypeDisposalOk() (*FieldEEzsignfoldertypeDisposal, bool)`

GetEEzsignfoldertypeDisposalOk returns a tuple with the EEzsignfoldertypeDisposal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignfoldertypeDisposal

`func (o *EzsignfoldertypeRequestV3) SetEEzsignfoldertypeDisposal(v FieldEEzsignfoldertypeDisposal)`

SetEEzsignfoldertypeDisposal sets EEzsignfoldertypeDisposal field to given value.


### GetEEzsignfoldertypeCompletion

`func (o *EzsignfoldertypeRequestV3) GetEEzsignfoldertypeCompletion() FieldEEzsignfoldertypeCompletion`

GetEEzsignfoldertypeCompletion returns the EEzsignfoldertypeCompletion field if non-nil, zero value otherwise.

### GetEEzsignfoldertypeCompletionOk

`func (o *EzsignfoldertypeRequestV3) GetEEzsignfoldertypeCompletionOk() (*FieldEEzsignfoldertypeCompletion, bool)`

GetEEzsignfoldertypeCompletionOk returns a tuple with the EEzsignfoldertypeCompletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignfoldertypeCompletion

`func (o *EzsignfoldertypeRequestV3) SetEEzsignfoldertypeCompletion(v FieldEEzsignfoldertypeCompletion)`

SetEEzsignfoldertypeCompletion sets EEzsignfoldertypeCompletion field to given value.


### GetIEzsignfoldertypeDisposaldays

`func (o *EzsignfoldertypeRequestV3) GetIEzsignfoldertypeDisposaldays() int32`

GetIEzsignfoldertypeDisposaldays returns the IEzsignfoldertypeDisposaldays field if non-nil, zero value otherwise.

### GetIEzsignfoldertypeDisposaldaysOk

`func (o *EzsignfoldertypeRequestV3) GetIEzsignfoldertypeDisposaldaysOk() (*int32, bool)`

GetIEzsignfoldertypeDisposaldaysOk returns a tuple with the IEzsignfoldertypeDisposaldays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignfoldertypeDisposaldays

`func (o *EzsignfoldertypeRequestV3) SetIEzsignfoldertypeDisposaldays(v int32)`

SetIEzsignfoldertypeDisposaldays sets IEzsignfoldertypeDisposaldays field to given value.

### HasIEzsignfoldertypeDisposaldays

`func (o *EzsignfoldertypeRequestV3) HasIEzsignfoldertypeDisposaldays() bool`

HasIEzsignfoldertypeDisposaldays returns a boolean if a field has been set.

### GetIEzsignfoldertypeDeadlinedays

`func (o *EzsignfoldertypeRequestV3) GetIEzsignfoldertypeDeadlinedays() int32`

GetIEzsignfoldertypeDeadlinedays returns the IEzsignfoldertypeDeadlinedays field if non-nil, zero value otherwise.

### GetIEzsignfoldertypeDeadlinedaysOk

`func (o *EzsignfoldertypeRequestV3) GetIEzsignfoldertypeDeadlinedaysOk() (*int32, bool)`

GetIEzsignfoldertypeDeadlinedaysOk returns a tuple with the IEzsignfoldertypeDeadlinedays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignfoldertypeDeadlinedays

`func (o *EzsignfoldertypeRequestV3) SetIEzsignfoldertypeDeadlinedays(v int32)`

SetIEzsignfoldertypeDeadlinedays sets IEzsignfoldertypeDeadlinedays field to given value.


### GetBEzsignfoldertypePrematurelyendautomatically

`func (o *EzsignfoldertypeRequestV3) GetBEzsignfoldertypePrematurelyendautomatically() bool`

GetBEzsignfoldertypePrematurelyendautomatically returns the BEzsignfoldertypePrematurelyendautomatically field if non-nil, zero value otherwise.

### GetBEzsignfoldertypePrematurelyendautomaticallyOk

`func (o *EzsignfoldertypeRequestV3) GetBEzsignfoldertypePrematurelyendautomaticallyOk() (*bool, bool)`

GetBEzsignfoldertypePrematurelyendautomaticallyOk returns a tuple with the BEzsignfoldertypePrematurelyendautomatically field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypePrematurelyendautomatically

`func (o *EzsignfoldertypeRequestV3) SetBEzsignfoldertypePrematurelyendautomatically(v bool)`

SetBEzsignfoldertypePrematurelyendautomatically sets BEzsignfoldertypePrematurelyendautomatically field to given value.

### HasBEzsignfoldertypePrematurelyendautomatically

`func (o *EzsignfoldertypeRequestV3) HasBEzsignfoldertypePrematurelyendautomatically() bool`

HasBEzsignfoldertypePrematurelyendautomatically returns a boolean if a field has been set.

### GetIEzsignfoldertypePrematurelyendautomaticallydays

`func (o *EzsignfoldertypeRequestV3) GetIEzsignfoldertypePrematurelyendautomaticallydays() int32`

GetIEzsignfoldertypePrematurelyendautomaticallydays returns the IEzsignfoldertypePrematurelyendautomaticallydays field if non-nil, zero value otherwise.

### GetIEzsignfoldertypePrematurelyendautomaticallydaysOk

`func (o *EzsignfoldertypeRequestV3) GetIEzsignfoldertypePrematurelyendautomaticallydaysOk() (*int32, bool)`

GetIEzsignfoldertypePrematurelyendautomaticallydaysOk returns a tuple with the IEzsignfoldertypePrematurelyendautomaticallydays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignfoldertypePrematurelyendautomaticallydays

`func (o *EzsignfoldertypeRequestV3) SetIEzsignfoldertypePrematurelyendautomaticallydays(v int32)`

SetIEzsignfoldertypePrematurelyendautomaticallydays sets IEzsignfoldertypePrematurelyendautomaticallydays field to given value.

### HasIEzsignfoldertypePrematurelyendautomaticallydays

`func (o *EzsignfoldertypeRequestV3) HasIEzsignfoldertypePrematurelyendautomaticallydays() bool`

HasIEzsignfoldertypePrematurelyendautomaticallydays returns a boolean if a field has been set.

### GetBEzsignfoldertypeAutomaticsignature

`func (o *EzsignfoldertypeRequestV3) GetBEzsignfoldertypeAutomaticsignature() bool`

GetBEzsignfoldertypeAutomaticsignature returns the BEzsignfoldertypeAutomaticsignature field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeAutomaticsignatureOk

`func (o *EzsignfoldertypeRequestV3) GetBEzsignfoldertypeAutomaticsignatureOk() (*bool, bool)`

GetBEzsignfoldertypeAutomaticsignatureOk returns a tuple with the BEzsignfoldertypeAutomaticsignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeAutomaticsignature

`func (o *EzsignfoldertypeRequestV3) SetBEzsignfoldertypeAutomaticsignature(v bool)`

SetBEzsignfoldertypeAutomaticsignature sets BEzsignfoldertypeAutomaticsignature field to given value.

### HasBEzsignfoldertypeAutomaticsignature

`func (o *EzsignfoldertypeRequestV3) HasBEzsignfoldertypeAutomaticsignature() bool`

HasBEzsignfoldertypeAutomaticsignature returns a boolean if a field has been set.

### GetBEzsignfoldertypeDelegate

`func (o *EzsignfoldertypeRequestV3) GetBEzsignfoldertypeDelegate() bool`

GetBEzsignfoldertypeDelegate returns the BEzsignfoldertypeDelegate field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeDelegateOk

`func (o *EzsignfoldertypeRequestV3) GetBEzsignfoldertypeDelegateOk() (*bool, bool)`

GetBEzsignfoldertypeDelegateOk returns a tuple with the BEzsignfoldertypeDelegate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeDelegate

`func (o *EzsignfoldertypeRequestV3) SetBEzsignfoldertypeDelegate(v bool)`

SetBEzsignfoldertypeDelegate sets BEzsignfoldertypeDelegate field to given value.

### HasBEzsignfoldertypeDelegate

`func (o *EzsignfoldertypeRequestV3) HasBEzsignfoldertypeDelegate() bool`

HasBEzsignfoldertypeDelegate returns a boolean if a field has been set.

### GetBEzsignfoldertypeDiscussion

`func (o *EzsignfoldertypeRequestV3) GetBEzsignfoldertypeDiscussion() bool`

GetBEzsignfoldertypeDiscussion returns the BEzsignfoldertypeDiscussion field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeDiscussionOk

`func (o *EzsignfoldertypeRequestV3) GetBEzsignfoldertypeDiscussionOk() (*bool, bool)`

GetBEzsignfoldertypeDiscussionOk returns a tuple with the BEzsignfoldertypeDiscussion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeDiscussion

`func (o *EzsignfoldertypeRequestV3) SetBEzsignfoldertypeDiscussion(v bool)`

SetBEzsignfoldertypeDiscussion sets BEzsignfoldertypeDiscussion field to given value.

### HasBEzsignfoldertypeDiscussion

`func (o *EzsignfoldertypeRequestV3) HasBEzsignfoldertypeDiscussion() bool`

HasBEzsignfoldertypeDiscussion returns a boolean if a field has been set.

### GetBEzsignfoldertypeLogrecipientinproof

`func (o *EzsignfoldertypeRequestV3) GetBEzsignfoldertypeLogrecipientinproof() bool`

GetBEzsignfoldertypeLogrecipientinproof returns the BEzsignfoldertypeLogrecipientinproof field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeLogrecipientinproofOk

`func (o *EzsignfoldertypeRequestV3) GetBEzsignfoldertypeLogrecipientinproofOk() (*bool, bool)`

GetBEzsignfoldertypeLogrecipientinproofOk returns a tuple with the BEzsignfoldertypeLogrecipientinproof field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeLogrecipientinproof

`func (o *EzsignfoldertypeRequestV3) SetBEzsignfoldertypeLogrecipientinproof(v bool)`

SetBEzsignfoldertypeLogrecipientinproof sets BEzsignfoldertypeLogrecipientinproof field to given value.

### HasBEzsignfoldertypeLogrecipientinproof

`func (o *EzsignfoldertypeRequestV3) HasBEzsignfoldertypeLogrecipientinproof() bool`

HasBEzsignfoldertypeLogrecipientinproof returns a boolean if a field has been set.

### GetBEzsignfoldertypeReassignezsignsigner

`func (o *EzsignfoldertypeRequestV3) GetBEzsignfoldertypeReassignezsignsigner() bool`

GetBEzsignfoldertypeReassignezsignsigner returns the BEzsignfoldertypeReassignezsignsigner field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeReassignezsignsignerOk

`func (o *EzsignfoldertypeRequestV3) GetBEzsignfoldertypeReassignezsignsignerOk() (*bool, bool)`

GetBEzsignfoldertypeReassignezsignsignerOk returns a tuple with the BEzsignfoldertypeReassignezsignsigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeReassignezsignsigner

`func (o *EzsignfoldertypeRequestV3) SetBEzsignfoldertypeReassignezsignsigner(v bool)`

SetBEzsignfoldertypeReassignezsignsigner sets BEzsignfoldertypeReassignezsignsigner field to given value.

### HasBEzsignfoldertypeReassignezsignsigner

`func (o *EzsignfoldertypeRequestV3) HasBEzsignfoldertypeReassignezsignsigner() bool`

HasBEzsignfoldertypeReassignezsignsigner returns a boolean if a field has been set.

### GetBEzsignfoldertypeReassignuser

`func (o *EzsignfoldertypeRequestV3) GetBEzsignfoldertypeReassignuser() bool`

GetBEzsignfoldertypeReassignuser returns the BEzsignfoldertypeReassignuser field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeReassignuserOk

`func (o *EzsignfoldertypeRequestV3) GetBEzsignfoldertypeReassignuserOk() (*bool, bool)`

GetBEzsignfoldertypeReassignuserOk returns a tuple with the BEzsignfoldertypeReassignuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeReassignuser

`func (o *EzsignfoldertypeRequestV3) SetBEzsignfoldertypeReassignuser(v bool)`

SetBEzsignfoldertypeReassignuser sets BEzsignfoldertypeReassignuser field to given value.

### HasBEzsignfoldertypeReassignuser

`func (o *EzsignfoldertypeRequestV3) HasBEzsignfoldertypeReassignuser() bool`

HasBEzsignfoldertypeReassignuser returns a boolean if a field has been set.

### GetBEzsignfoldertypeReassigngroup

`func (o *EzsignfoldertypeRequestV3) GetBEzsignfoldertypeReassigngroup() bool`

GetBEzsignfoldertypeReassigngroup returns the BEzsignfoldertypeReassigngroup field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeReassigngroupOk

`func (o *EzsignfoldertypeRequestV3) GetBEzsignfoldertypeReassigngroupOk() (*bool, bool)`

GetBEzsignfoldertypeReassigngroupOk returns a tuple with the BEzsignfoldertypeReassigngroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeReassigngroup

`func (o *EzsignfoldertypeRequestV3) SetBEzsignfoldertypeReassigngroup(v bool)`

SetBEzsignfoldertypeReassigngroup sets BEzsignfoldertypeReassigngroup field to given value.

### HasBEzsignfoldertypeReassigngroup

`func (o *EzsignfoldertypeRequestV3) HasBEzsignfoldertypeReassigngroup() bool`

HasBEzsignfoldertypeReassigngroup returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendsignedtoezsignsigner

`func (o *EzsignfoldertypeRequestV3) GetBEzsignfoldertypeSendsignedtoezsignsigner() bool`

GetBEzsignfoldertypeSendsignedtoezsignsigner returns the BEzsignfoldertypeSendsignedtoezsignsigner field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsignedtoezsignsignerOk

`func (o *EzsignfoldertypeRequestV3) GetBEzsignfoldertypeSendsignedtoezsignsignerOk() (*bool, bool)`

GetBEzsignfoldertypeSendsignedtoezsignsignerOk returns a tuple with the BEzsignfoldertypeSendsignedtoezsignsigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsignedtoezsignsigner

`func (o *EzsignfoldertypeRequestV3) SetBEzsignfoldertypeSendsignedtoezsignsigner(v bool)`

SetBEzsignfoldertypeSendsignedtoezsignsigner sets BEzsignfoldertypeSendsignedtoezsignsigner field to given value.

### HasBEzsignfoldertypeSendsignedtoezsignsigner

`func (o *EzsignfoldertypeRequestV3) HasBEzsignfoldertypeSendsignedtoezsignsigner() bool`

HasBEzsignfoldertypeSendsignedtoezsignsigner returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendsignedtouser

`func (o *EzsignfoldertypeRequestV3) GetBEzsignfoldertypeSendsignedtouser() bool`

GetBEzsignfoldertypeSendsignedtouser returns the BEzsignfoldertypeSendsignedtouser field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsignedtouserOk

`func (o *EzsignfoldertypeRequestV3) GetBEzsignfoldertypeSendsignedtouserOk() (*bool, bool)`

GetBEzsignfoldertypeSendsignedtouserOk returns a tuple with the BEzsignfoldertypeSendsignedtouser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsignedtouser

`func (o *EzsignfoldertypeRequestV3) SetBEzsignfoldertypeSendsignedtouser(v bool)`

SetBEzsignfoldertypeSendsignedtouser sets BEzsignfoldertypeSendsignedtouser field to given value.

### HasBEzsignfoldertypeSendsignedtouser

`func (o *EzsignfoldertypeRequestV3) HasBEzsignfoldertypeSendsignedtouser() bool`

HasBEzsignfoldertypeSendsignedtouser returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendattachmentezsignsigner

`func (o *EzsignfoldertypeRequestV3) GetBEzsignfoldertypeSendattachmentezsignsigner() bool`

GetBEzsignfoldertypeSendattachmentezsignsigner returns the BEzsignfoldertypeSendattachmentezsignsigner field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendattachmentezsignsignerOk

`func (o *EzsignfoldertypeRequestV3) GetBEzsignfoldertypeSendattachmentezsignsignerOk() (*bool, bool)`

GetBEzsignfoldertypeSendattachmentezsignsignerOk returns a tuple with the BEzsignfoldertypeSendattachmentezsignsigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendattachmentezsignsigner

`func (o *EzsignfoldertypeRequestV3) SetBEzsignfoldertypeSendattachmentezsignsigner(v bool)`

SetBEzsignfoldertypeSendattachmentezsignsigner sets BEzsignfoldertypeSendattachmentezsignsigner field to given value.

### HasBEzsignfoldertypeSendattachmentezsignsigner

`func (o *EzsignfoldertypeRequestV3) HasBEzsignfoldertypeSendattachmentezsignsigner() bool`

HasBEzsignfoldertypeSendattachmentezsignsigner returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendproofezsignsigner

`func (o *EzsignfoldertypeRequestV3) GetBEzsignfoldertypeSendproofezsignsigner() bool`

GetBEzsignfoldertypeSendproofezsignsigner returns the BEzsignfoldertypeSendproofezsignsigner field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendproofezsignsignerOk

`func (o *EzsignfoldertypeRequestV3) GetBEzsignfoldertypeSendproofezsignsignerOk() (*bool, bool)`

GetBEzsignfoldertypeSendproofezsignsignerOk returns a tuple with the BEzsignfoldertypeSendproofezsignsigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendproofezsignsigner

`func (o *EzsignfoldertypeRequestV3) SetBEzsignfoldertypeSendproofezsignsigner(v bool)`

SetBEzsignfoldertypeSendproofezsignsigner sets BEzsignfoldertypeSendproofezsignsigner field to given value.

### HasBEzsignfoldertypeSendproofezsignsigner

`func (o *EzsignfoldertypeRequestV3) HasBEzsignfoldertypeSendproofezsignsigner() bool`

HasBEzsignfoldertypeSendproofezsignsigner returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendattachmentreceivecopy

`func (o *EzsignfoldertypeRequestV3) GetBEzsignfoldertypeSendattachmentreceivecopy() bool`

GetBEzsignfoldertypeSendattachmentreceivecopy returns the BEzsignfoldertypeSendattachmentreceivecopy field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendattachmentreceivecopyOk

`func (o *EzsignfoldertypeRequestV3) GetBEzsignfoldertypeSendattachmentreceivecopyOk() (*bool, bool)`

GetBEzsignfoldertypeSendattachmentreceivecopyOk returns a tuple with the BEzsignfoldertypeSendattachmentreceivecopy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendattachmentreceivecopy

`func (o *EzsignfoldertypeRequestV3) SetBEzsignfoldertypeSendattachmentreceivecopy(v bool)`

SetBEzsignfoldertypeSendattachmentreceivecopy sets BEzsignfoldertypeSendattachmentreceivecopy field to given value.

### HasBEzsignfoldertypeSendattachmentreceivecopy

`func (o *EzsignfoldertypeRequestV3) HasBEzsignfoldertypeSendattachmentreceivecopy() bool`

HasBEzsignfoldertypeSendattachmentreceivecopy returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendattachmentuser

`func (o *EzsignfoldertypeRequestV3) GetBEzsignfoldertypeSendattachmentuser() bool`

GetBEzsignfoldertypeSendattachmentuser returns the BEzsignfoldertypeSendattachmentuser field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendattachmentuserOk

`func (o *EzsignfoldertypeRequestV3) GetBEzsignfoldertypeSendattachmentuserOk() (*bool, bool)`

GetBEzsignfoldertypeSendattachmentuserOk returns a tuple with the BEzsignfoldertypeSendattachmentuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendattachmentuser

`func (o *EzsignfoldertypeRequestV3) SetBEzsignfoldertypeSendattachmentuser(v bool)`

SetBEzsignfoldertypeSendattachmentuser sets BEzsignfoldertypeSendattachmentuser field to given value.

### HasBEzsignfoldertypeSendattachmentuser

`func (o *EzsignfoldertypeRequestV3) HasBEzsignfoldertypeSendattachmentuser() bool`

HasBEzsignfoldertypeSendattachmentuser returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendproofuser

`func (o *EzsignfoldertypeRequestV3) GetBEzsignfoldertypeSendproofuser() bool`

GetBEzsignfoldertypeSendproofuser returns the BEzsignfoldertypeSendproofuser field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendproofuserOk

`func (o *EzsignfoldertypeRequestV3) GetBEzsignfoldertypeSendproofuserOk() (*bool, bool)`

GetBEzsignfoldertypeSendproofuserOk returns a tuple with the BEzsignfoldertypeSendproofuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendproofuser

`func (o *EzsignfoldertypeRequestV3) SetBEzsignfoldertypeSendproofuser(v bool)`

SetBEzsignfoldertypeSendproofuser sets BEzsignfoldertypeSendproofuser field to given value.

### HasBEzsignfoldertypeSendproofuser

`func (o *EzsignfoldertypeRequestV3) HasBEzsignfoldertypeSendproofuser() bool`

HasBEzsignfoldertypeSendproofuser returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendproofemail

`func (o *EzsignfoldertypeRequestV3) GetBEzsignfoldertypeSendproofemail() bool`

GetBEzsignfoldertypeSendproofemail returns the BEzsignfoldertypeSendproofemail field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendproofemailOk

`func (o *EzsignfoldertypeRequestV3) GetBEzsignfoldertypeSendproofemailOk() (*bool, bool)`

GetBEzsignfoldertypeSendproofemailOk returns a tuple with the BEzsignfoldertypeSendproofemail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendproofemail

`func (o *EzsignfoldertypeRequestV3) SetBEzsignfoldertypeSendproofemail(v bool)`

SetBEzsignfoldertypeSendproofemail sets BEzsignfoldertypeSendproofemail field to given value.

### HasBEzsignfoldertypeSendproofemail

`func (o *EzsignfoldertypeRequestV3) HasBEzsignfoldertypeSendproofemail() bool`

HasBEzsignfoldertypeSendproofemail returns a boolean if a field has been set.

### GetBEzsignfoldertypeAllowdownloadattachmentezsignsigner

`func (o *EzsignfoldertypeRequestV3) GetBEzsignfoldertypeAllowdownloadattachmentezsignsigner() bool`

GetBEzsignfoldertypeAllowdownloadattachmentezsignsigner returns the BEzsignfoldertypeAllowdownloadattachmentezsignsigner field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeAllowdownloadattachmentezsignsignerOk

`func (o *EzsignfoldertypeRequestV3) GetBEzsignfoldertypeAllowdownloadattachmentezsignsignerOk() (*bool, bool)`

GetBEzsignfoldertypeAllowdownloadattachmentezsignsignerOk returns a tuple with the BEzsignfoldertypeAllowdownloadattachmentezsignsigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeAllowdownloadattachmentezsignsigner

`func (o *EzsignfoldertypeRequestV3) SetBEzsignfoldertypeAllowdownloadattachmentezsignsigner(v bool)`

SetBEzsignfoldertypeAllowdownloadattachmentezsignsigner sets BEzsignfoldertypeAllowdownloadattachmentezsignsigner field to given value.

### HasBEzsignfoldertypeAllowdownloadattachmentezsignsigner

`func (o *EzsignfoldertypeRequestV3) HasBEzsignfoldertypeAllowdownloadattachmentezsignsigner() bool`

HasBEzsignfoldertypeAllowdownloadattachmentezsignsigner returns a boolean if a field has been set.

### GetBEzsignfoldertypeAllowdownloadproofezsignsigner

`func (o *EzsignfoldertypeRequestV3) GetBEzsignfoldertypeAllowdownloadproofezsignsigner() bool`

GetBEzsignfoldertypeAllowdownloadproofezsignsigner returns the BEzsignfoldertypeAllowdownloadproofezsignsigner field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeAllowdownloadproofezsignsignerOk

`func (o *EzsignfoldertypeRequestV3) GetBEzsignfoldertypeAllowdownloadproofezsignsignerOk() (*bool, bool)`

GetBEzsignfoldertypeAllowdownloadproofezsignsignerOk returns a tuple with the BEzsignfoldertypeAllowdownloadproofezsignsigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeAllowdownloadproofezsignsigner

`func (o *EzsignfoldertypeRequestV3) SetBEzsignfoldertypeAllowdownloadproofezsignsigner(v bool)`

SetBEzsignfoldertypeAllowdownloadproofezsignsigner sets BEzsignfoldertypeAllowdownloadproofezsignsigner field to given value.

### HasBEzsignfoldertypeAllowdownloadproofezsignsigner

`func (o *EzsignfoldertypeRequestV3) HasBEzsignfoldertypeAllowdownloadproofezsignsigner() bool`

HasBEzsignfoldertypeAllowdownloadproofezsignsigner returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendproofreceivealldocument

`func (o *EzsignfoldertypeRequestV3) GetBEzsignfoldertypeSendproofreceivealldocument() bool`

GetBEzsignfoldertypeSendproofreceivealldocument returns the BEzsignfoldertypeSendproofreceivealldocument field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendproofreceivealldocumentOk

`func (o *EzsignfoldertypeRequestV3) GetBEzsignfoldertypeSendproofreceivealldocumentOk() (*bool, bool)`

GetBEzsignfoldertypeSendproofreceivealldocumentOk returns a tuple with the BEzsignfoldertypeSendproofreceivealldocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendproofreceivealldocument

`func (o *EzsignfoldertypeRequestV3) SetBEzsignfoldertypeSendproofreceivealldocument(v bool)`

SetBEzsignfoldertypeSendproofreceivealldocument sets BEzsignfoldertypeSendproofreceivealldocument field to given value.

### HasBEzsignfoldertypeSendproofreceivealldocument

`func (o *EzsignfoldertypeRequestV3) HasBEzsignfoldertypeSendproofreceivealldocument() bool`

HasBEzsignfoldertypeSendproofreceivealldocument returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendsignedtodocumentowner

`func (o *EzsignfoldertypeRequestV3) GetBEzsignfoldertypeSendsignedtodocumentowner() bool`

GetBEzsignfoldertypeSendsignedtodocumentowner returns the BEzsignfoldertypeSendsignedtodocumentowner field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsignedtodocumentownerOk

`func (o *EzsignfoldertypeRequestV3) GetBEzsignfoldertypeSendsignedtodocumentownerOk() (*bool, bool)`

GetBEzsignfoldertypeSendsignedtodocumentownerOk returns a tuple with the BEzsignfoldertypeSendsignedtodocumentowner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsignedtodocumentowner

`func (o *EzsignfoldertypeRequestV3) SetBEzsignfoldertypeSendsignedtodocumentowner(v bool)`

SetBEzsignfoldertypeSendsignedtodocumentowner sets BEzsignfoldertypeSendsignedtodocumentowner field to given value.


### GetBEzsignfoldertypeSendsignedtofolderowner

`func (o *EzsignfoldertypeRequestV3) GetBEzsignfoldertypeSendsignedtofolderowner() bool`

GetBEzsignfoldertypeSendsignedtofolderowner returns the BEzsignfoldertypeSendsignedtofolderowner field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsignedtofolderownerOk

`func (o *EzsignfoldertypeRequestV3) GetBEzsignfoldertypeSendsignedtofolderownerOk() (*bool, bool)`

GetBEzsignfoldertypeSendsignedtofolderownerOk returns a tuple with the BEzsignfoldertypeSendsignedtofolderowner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsignedtofolderowner

`func (o *EzsignfoldertypeRequestV3) SetBEzsignfoldertypeSendsignedtofolderowner(v bool)`

SetBEzsignfoldertypeSendsignedtofolderowner sets BEzsignfoldertypeSendsignedtofolderowner field to given value.


### GetBEzsignfoldertypeSendsignedtofullgroup

`func (o *EzsignfoldertypeRequestV3) GetBEzsignfoldertypeSendsignedtofullgroup() bool`

GetBEzsignfoldertypeSendsignedtofullgroup returns the BEzsignfoldertypeSendsignedtofullgroup field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsignedtofullgroupOk

`func (o *EzsignfoldertypeRequestV3) GetBEzsignfoldertypeSendsignedtofullgroupOk() (*bool, bool)`

GetBEzsignfoldertypeSendsignedtofullgroupOk returns a tuple with the BEzsignfoldertypeSendsignedtofullgroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsignedtofullgroup

`func (o *EzsignfoldertypeRequestV3) SetBEzsignfoldertypeSendsignedtofullgroup(v bool)`

SetBEzsignfoldertypeSendsignedtofullgroup sets BEzsignfoldertypeSendsignedtofullgroup field to given value.

### HasBEzsignfoldertypeSendsignedtofullgroup

`func (o *EzsignfoldertypeRequestV3) HasBEzsignfoldertypeSendsignedtofullgroup() bool`

HasBEzsignfoldertypeSendsignedtofullgroup returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendsignedtolimitedgroup

`func (o *EzsignfoldertypeRequestV3) GetBEzsignfoldertypeSendsignedtolimitedgroup() bool`

GetBEzsignfoldertypeSendsignedtolimitedgroup returns the BEzsignfoldertypeSendsignedtolimitedgroup field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsignedtolimitedgroupOk

`func (o *EzsignfoldertypeRequestV3) GetBEzsignfoldertypeSendsignedtolimitedgroupOk() (*bool, bool)`

GetBEzsignfoldertypeSendsignedtolimitedgroupOk returns a tuple with the BEzsignfoldertypeSendsignedtolimitedgroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsignedtolimitedgroup

`func (o *EzsignfoldertypeRequestV3) SetBEzsignfoldertypeSendsignedtolimitedgroup(v bool)`

SetBEzsignfoldertypeSendsignedtolimitedgroup sets BEzsignfoldertypeSendsignedtolimitedgroup field to given value.

### HasBEzsignfoldertypeSendsignedtolimitedgroup

`func (o *EzsignfoldertypeRequestV3) HasBEzsignfoldertypeSendsignedtolimitedgroup() bool`

HasBEzsignfoldertypeSendsignedtolimitedgroup returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendsignedtocolleague

`func (o *EzsignfoldertypeRequestV3) GetBEzsignfoldertypeSendsignedtocolleague() bool`

GetBEzsignfoldertypeSendsignedtocolleague returns the BEzsignfoldertypeSendsignedtocolleague field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsignedtocolleagueOk

`func (o *EzsignfoldertypeRequestV3) GetBEzsignfoldertypeSendsignedtocolleagueOk() (*bool, bool)`

GetBEzsignfoldertypeSendsignedtocolleagueOk returns a tuple with the BEzsignfoldertypeSendsignedtocolleague field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsignedtocolleague

`func (o *EzsignfoldertypeRequestV3) SetBEzsignfoldertypeSendsignedtocolleague(v bool)`

SetBEzsignfoldertypeSendsignedtocolleague sets BEzsignfoldertypeSendsignedtocolleague field to given value.


### GetBEzsignfoldertypeSendsummarytodocumentowner

`func (o *EzsignfoldertypeRequestV3) GetBEzsignfoldertypeSendsummarytodocumentowner() bool`

GetBEzsignfoldertypeSendsummarytodocumentowner returns the BEzsignfoldertypeSendsummarytodocumentowner field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsummarytodocumentownerOk

`func (o *EzsignfoldertypeRequestV3) GetBEzsignfoldertypeSendsummarytodocumentownerOk() (*bool, bool)`

GetBEzsignfoldertypeSendsummarytodocumentownerOk returns a tuple with the BEzsignfoldertypeSendsummarytodocumentowner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsummarytodocumentowner

`func (o *EzsignfoldertypeRequestV3) SetBEzsignfoldertypeSendsummarytodocumentowner(v bool)`

SetBEzsignfoldertypeSendsummarytodocumentowner sets BEzsignfoldertypeSendsummarytodocumentowner field to given value.


### GetBEzsignfoldertypeSendsummarytofolderowner

`func (o *EzsignfoldertypeRequestV3) GetBEzsignfoldertypeSendsummarytofolderowner() bool`

GetBEzsignfoldertypeSendsummarytofolderowner returns the BEzsignfoldertypeSendsummarytofolderowner field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsummarytofolderownerOk

`func (o *EzsignfoldertypeRequestV3) GetBEzsignfoldertypeSendsummarytofolderownerOk() (*bool, bool)`

GetBEzsignfoldertypeSendsummarytofolderownerOk returns a tuple with the BEzsignfoldertypeSendsummarytofolderowner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsummarytofolderowner

`func (o *EzsignfoldertypeRequestV3) SetBEzsignfoldertypeSendsummarytofolderowner(v bool)`

SetBEzsignfoldertypeSendsummarytofolderowner sets BEzsignfoldertypeSendsummarytofolderowner field to given value.


### GetBEzsignfoldertypeSendsummarytofullgroup

`func (o *EzsignfoldertypeRequestV3) GetBEzsignfoldertypeSendsummarytofullgroup() bool`

GetBEzsignfoldertypeSendsummarytofullgroup returns the BEzsignfoldertypeSendsummarytofullgroup field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsummarytofullgroupOk

`func (o *EzsignfoldertypeRequestV3) GetBEzsignfoldertypeSendsummarytofullgroupOk() (*bool, bool)`

GetBEzsignfoldertypeSendsummarytofullgroupOk returns a tuple with the BEzsignfoldertypeSendsummarytofullgroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsummarytofullgroup

`func (o *EzsignfoldertypeRequestV3) SetBEzsignfoldertypeSendsummarytofullgroup(v bool)`

SetBEzsignfoldertypeSendsummarytofullgroup sets BEzsignfoldertypeSendsummarytofullgroup field to given value.

### HasBEzsignfoldertypeSendsummarytofullgroup

`func (o *EzsignfoldertypeRequestV3) HasBEzsignfoldertypeSendsummarytofullgroup() bool`

HasBEzsignfoldertypeSendsummarytofullgroup returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendsummarytolimitedgroup

`func (o *EzsignfoldertypeRequestV3) GetBEzsignfoldertypeSendsummarytolimitedgroup() bool`

GetBEzsignfoldertypeSendsummarytolimitedgroup returns the BEzsignfoldertypeSendsummarytolimitedgroup field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsummarytolimitedgroupOk

`func (o *EzsignfoldertypeRequestV3) GetBEzsignfoldertypeSendsummarytolimitedgroupOk() (*bool, bool)`

GetBEzsignfoldertypeSendsummarytolimitedgroupOk returns a tuple with the BEzsignfoldertypeSendsummarytolimitedgroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsummarytolimitedgroup

`func (o *EzsignfoldertypeRequestV3) SetBEzsignfoldertypeSendsummarytolimitedgroup(v bool)`

SetBEzsignfoldertypeSendsummarytolimitedgroup sets BEzsignfoldertypeSendsummarytolimitedgroup field to given value.

### HasBEzsignfoldertypeSendsummarytolimitedgroup

`func (o *EzsignfoldertypeRequestV3) HasBEzsignfoldertypeSendsummarytolimitedgroup() bool`

HasBEzsignfoldertypeSendsummarytolimitedgroup returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendsummarytocolleague

`func (o *EzsignfoldertypeRequestV3) GetBEzsignfoldertypeSendsummarytocolleague() bool`

GetBEzsignfoldertypeSendsummarytocolleague returns the BEzsignfoldertypeSendsummarytocolleague field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsummarytocolleagueOk

`func (o *EzsignfoldertypeRequestV3) GetBEzsignfoldertypeSendsummarytocolleagueOk() (*bool, bool)`

GetBEzsignfoldertypeSendsummarytocolleagueOk returns a tuple with the BEzsignfoldertypeSendsummarytocolleague field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsummarytocolleague

`func (o *EzsignfoldertypeRequestV3) SetBEzsignfoldertypeSendsummarytocolleague(v bool)`

SetBEzsignfoldertypeSendsummarytocolleague sets BEzsignfoldertypeSendsummarytocolleague field to given value.


### GetEEzsignfoldertypeSigneraccess

`func (o *EzsignfoldertypeRequestV3) GetEEzsignfoldertypeSigneraccess() FieldEEzsignfoldertypeSigneraccess`

GetEEzsignfoldertypeSigneraccess returns the EEzsignfoldertypeSigneraccess field if non-nil, zero value otherwise.

### GetEEzsignfoldertypeSigneraccessOk

`func (o *EzsignfoldertypeRequestV3) GetEEzsignfoldertypeSigneraccessOk() (*FieldEEzsignfoldertypeSigneraccess, bool)`

GetEEzsignfoldertypeSigneraccessOk returns a tuple with the EEzsignfoldertypeSigneraccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignfoldertypeSigneraccess

`func (o *EzsignfoldertypeRequestV3) SetEEzsignfoldertypeSigneraccess(v FieldEEzsignfoldertypeSigneraccess)`

SetEEzsignfoldertypeSigneraccess sets EEzsignfoldertypeSigneraccess field to given value.

### HasEEzsignfoldertypeSigneraccess

`func (o *EzsignfoldertypeRequestV3) HasEEzsignfoldertypeSigneraccess() bool`

HasEEzsignfoldertypeSigneraccess returns a boolean if a field has been set.

### GetBEzsignfoldertypeIsactive

`func (o *EzsignfoldertypeRequestV3) GetBEzsignfoldertypeIsactive() bool`

GetBEzsignfoldertypeIsactive returns the BEzsignfoldertypeIsactive field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeIsactiveOk

`func (o *EzsignfoldertypeRequestV3) GetBEzsignfoldertypeIsactiveOk() (*bool, bool)`

GetBEzsignfoldertypeIsactiveOk returns a tuple with the BEzsignfoldertypeIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeIsactive

`func (o *EzsignfoldertypeRequestV3) SetBEzsignfoldertypeIsactive(v bool)`

SetBEzsignfoldertypeIsactive sets BEzsignfoldertypeIsactive field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


