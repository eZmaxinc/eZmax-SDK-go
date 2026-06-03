# EzsignfoldertypeRequestV4

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
**EEzsignfoldertypeSignaturedatedisplay** | [**FieldEEzsignfoldertypeSignaturedatedisplay**](FieldEEzsignfoldertypeSignaturedatedisplay.md) |  | 
**SEzsignfoldertypeSignaturedatecustomformat** | Pointer to **string** | The custom date format to use  You can use the codes below and they will be replaced at signature time. Text values like month and day names will be rendered in the proper language. Other text will be left as-is.  The codes examples below are based on the following datetime: Thursday, January 6, 2022 at 08:07:09 EST  For example, the format \&quot;Signature date: {MM}/{DD}/{YYYY} {hh}:{mm}\&quot; would become \&quot;Signature date: 01/06/2022 08:07\&quot;  **Year**  | Code | Example | | - | - | | {YYYY} | 2022 | | {YY} | 22 |  **Month**  | Code | Example | | - | - | | {MonthCapitalize} | Janvier | | {Month} | janvier | | {MM} | 01 | | {M} | 1 |  **Day**  | Code | Example | | - | - | | {DayCapitalize} | Jeudi | | {Day} | jeudi | | {DD} | 06 | | {D} | 6 |  **Hour**  | Code | Example | | - | - | | {hh} | 08 |  **Minute**  | Code | Example | | - | - | | {mm} | 07 |  **Second**  | Code | Example | | - | - | | {ss} | 09 |        **Timezone**  | Code | Example | | - | - | | {Z} | EST |       **Time**  | Code | Example | | - | - | | {Time} | 08:07:09 |   | {TimeZ} | 08:07:09 EST |     **Date**  | Code | Example | | - | - | | {Date} | 2022-01-06 |   | {DateText} | 1er Janvier 2022 |  **Full**  | Code | Example | | - | - | | {DateTime} | 2022-01-06 08:07:09 |   | {DateTimeZ} | 2022-01-06 08:07:09 EST |  | [optional] 
**EEzsignfoldertypeDocumentdependency** | Pointer to [**FieldEEzsignfoldertypeDocumentdependency**](FieldEEzsignfoldertypeDocumentdependency.md) |  | [optional] 
**EEzsignfoldertypeDocumentmerge** | Pointer to [**FieldEEzsignfoldertypeDocumentmerge**](FieldEEzsignfoldertypeDocumentmerge.md) |  | [optional] [default to NO]
**SEmailAddressSigned** | Pointer to **string** | The email address. | [optional] 
**SEmailAddressSummary** | Pointer to **string** | The email address. | [optional] 
**EEzsignfoldertypePdfarequirement** | Pointer to [**FieldEEzsignfoldertypePdfarequirement**](FieldEEzsignfoldertypePdfarequirement.md) |  | [optional] 
**EEzsignfoldertypePdfanoncompliantaction** | Pointer to [**FieldEEzsignfoldertypePdfanoncompliantaction**](FieldEEzsignfoldertypePdfanoncompliantaction.md) |  | [optional] 
**EEzsignfoldertypePrivacylevel** | [**FieldEEzsignfoldertypePrivacylevel**](FieldEEzsignfoldertypePrivacylevel.md) |  | 
**IEzsignfoldertypeFontsizeannotation** | Pointer to **int32** | Font size for annotations | [optional] 
**IEzsignfoldertypeFontsizeformfield** | Pointer to **int32** | Font size for form fields | [optional] 
**IEzsignfoldertypeSendreminderfirstdays** | Pointer to **int32** | The number of days before the first reminder sending | [optional] 
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
**BEzsignfoldertypeSenddocumentmergetoemail** | Pointer to **bool** | Whether we send the merged documents in the email to external recipient | [optional] 
**BEzsignfoldertypeSenddocumentmergetoezsignsigner** | Pointer to **bool** | Whether we send the merged documents in the email to Ezsignsigner | [optional] 
**BEzsignfoldertypeSenddocumentmergetoreceivealldocument** | Pointer to **bool** | Whether we send the merged documents in the email to user and Ezsignsigner who receive all documents. | [optional] 
**BEzsignfoldertypeSenddocumentmergetouser** | Pointer to **bool** | Whether we send the merged documents in the email to User | [optional] 
**BEzsignfoldertypeSendsignedtoezsignsigner** | Pointer to **bool** | Whether we send an email to Ezsignsigner  when document is completed | [optional] 
**BEzsignfoldertypeSendsignedtouser** | Pointer to **bool** | Whether we send an email to User who signed when document is completed | [optional] 
**BEzsignfoldertypeSendattachmentezsignsigner** | Pointer to **bool** | Whether we send the Ezsigndocument in the email to Ezsignsigner | [optional] 
**BEzsignfoldertypeSendsignatureattachmentezsignsigner** | Pointer to **bool** | Whether we send the attachments contained in the Ezsignsignatures in the email to Ezsignsigner | [optional] 
**BEzsignfoldertypeSendsignatureattachment** | Pointer to **bool** | Whether we send the attachments contained in the Ezsignsignatures in the email to external recipient | [optional] 
**BEzsignfoldertypeSendproofezsignsigner** | Pointer to **bool** | Whether we send the proof in the email to Ezsignsigner | [optional] 
**BEzsignfoldertypeSendattachmentuser** | Pointer to **bool** | Whether we send the Ezsigndocument in the email to User | [optional] 
**BEzsignfoldertypeSendsignatureattachmentuser** | Pointer to **bool** | Whether we send the attachments contained in the Ezsignsignatures in the email to User | [optional] 
**BEzsignfoldertypeSendproofuser** | Pointer to **bool** | Whether we send the proof in the email to User | [optional] 
**BEzsignfoldertypeSendproofemail** | Pointer to **bool** | Whether we send the proof in the email to external recipient | [optional] 
**BEzsignfoldertypeAllowdownloadattachmentezsignsigner** | Pointer to **bool** | Whether we allow the Ezsigndocument to be downloaded by an Ezsignsigner | [optional] 
**BEzsignfoldertypeAllowdownloadsignatureattachmentezsignsigner** | Pointer to **bool** | Whether we allow the attachments in the Ezsignsignatures to be downloaded by an Ezsignsigner | [optional] 
**BEzsignfoldertypeAllowdownloadproofezsignsigner** | Pointer to **bool** | Whether we allow the proof to be downloaded by an Ezsignsigner | [optional] 
**BEzsignfoldertypeSendproofreceivealldocument** | Pointer to **bool** | Whether we send the proof to user and Ezsignsigner who receive all documents. | [optional] 
**BEzsignfoldertypeSendsignatureattachmentreceivealldocument** | Pointer to **bool** | Whether we send the attachments contained in the Ezsignsignatures to user and Ezsignsigner who receive all documents. | [optional] 
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

### NewEzsignfoldertypeRequestV4

`func NewEzsignfoldertypeRequestV4(objEzsignfoldertypeName MultilingualEzsignfoldertypeName, fkiBrandingID int32, aFkiUserlogintypeID []int32, eEzsignfoldertypeSignaturedatedisplay FieldEEzsignfoldertypeSignaturedatedisplay, eEzsignfoldertypePrivacylevel FieldEEzsignfoldertypePrivacylevel, iEzsignfoldertypeArchivaldays int32, eEzsignfoldertypeDisposal FieldEEzsignfoldertypeDisposal, eEzsignfoldertypeCompletion FieldEEzsignfoldertypeCompletion, iEzsignfoldertypeDeadlinedays int32, bEzsignfoldertypeSendsignedtodocumentowner bool, bEzsignfoldertypeSendsignedtofolderowner bool, bEzsignfoldertypeSendsignedtocolleague bool, bEzsignfoldertypeSendsummarytodocumentowner bool, bEzsignfoldertypeSendsummarytofolderowner bool, bEzsignfoldertypeSendsummarytocolleague bool, bEzsignfoldertypeIsactive bool, ) *EzsignfoldertypeRequestV4`

NewEzsignfoldertypeRequestV4 instantiates a new EzsignfoldertypeRequestV4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignfoldertypeRequestV4WithDefaults

`func NewEzsignfoldertypeRequestV4WithDefaults() *EzsignfoldertypeRequestV4`

NewEzsignfoldertypeRequestV4WithDefaults instantiates a new EzsignfoldertypeRequestV4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsignfoldertypeID

`func (o *EzsignfoldertypeRequestV4) GetPkiEzsignfoldertypeID() int32`

GetPkiEzsignfoldertypeID returns the PkiEzsignfoldertypeID field if non-nil, zero value otherwise.

### GetPkiEzsignfoldertypeIDOk

`func (o *EzsignfoldertypeRequestV4) GetPkiEzsignfoldertypeIDOk() (*int32, bool)`

GetPkiEzsignfoldertypeIDOk returns a tuple with the PkiEzsignfoldertypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsignfoldertypeID

`func (o *EzsignfoldertypeRequestV4) SetPkiEzsignfoldertypeID(v int32)`

SetPkiEzsignfoldertypeID sets PkiEzsignfoldertypeID field to given value.

### HasPkiEzsignfoldertypeID

`func (o *EzsignfoldertypeRequestV4) HasPkiEzsignfoldertypeID() bool`

HasPkiEzsignfoldertypeID returns a boolean if a field has been set.

### GetObjEzsignfoldertypeName

`func (o *EzsignfoldertypeRequestV4) GetObjEzsignfoldertypeName() MultilingualEzsignfoldertypeName`

GetObjEzsignfoldertypeName returns the ObjEzsignfoldertypeName field if non-nil, zero value otherwise.

### GetObjEzsignfoldertypeNameOk

`func (o *EzsignfoldertypeRequestV4) GetObjEzsignfoldertypeNameOk() (*MultilingualEzsignfoldertypeName, bool)`

GetObjEzsignfoldertypeNameOk returns a tuple with the ObjEzsignfoldertypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjEzsignfoldertypeName

`func (o *EzsignfoldertypeRequestV4) SetObjEzsignfoldertypeName(v MultilingualEzsignfoldertypeName)`

SetObjEzsignfoldertypeName sets ObjEzsignfoldertypeName field to given value.


### GetFkiBrandingID

`func (o *EzsignfoldertypeRequestV4) GetFkiBrandingID() int32`

GetFkiBrandingID returns the FkiBrandingID field if non-nil, zero value otherwise.

### GetFkiBrandingIDOk

`func (o *EzsignfoldertypeRequestV4) GetFkiBrandingIDOk() (*int32, bool)`

GetFkiBrandingIDOk returns a tuple with the FkiBrandingID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiBrandingID

`func (o *EzsignfoldertypeRequestV4) SetFkiBrandingID(v int32)`

SetFkiBrandingID sets FkiBrandingID field to given value.


### GetFkiBillingentityinternalID

`func (o *EzsignfoldertypeRequestV4) GetFkiBillingentityinternalID() int32`

GetFkiBillingentityinternalID returns the FkiBillingentityinternalID field if non-nil, zero value otherwise.

### GetFkiBillingentityinternalIDOk

`func (o *EzsignfoldertypeRequestV4) GetFkiBillingentityinternalIDOk() (*int32, bool)`

GetFkiBillingentityinternalIDOk returns a tuple with the FkiBillingentityinternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiBillingentityinternalID

`func (o *EzsignfoldertypeRequestV4) SetFkiBillingentityinternalID(v int32)`

SetFkiBillingentityinternalID sets FkiBillingentityinternalID field to given value.

### HasFkiBillingentityinternalID

`func (o *EzsignfoldertypeRequestV4) HasFkiBillingentityinternalID() bool`

HasFkiBillingentityinternalID returns a boolean if a field has been set.

### GetFkiEzsigntsarequirementID

`func (o *EzsignfoldertypeRequestV4) GetFkiEzsigntsarequirementID() int32`

GetFkiEzsigntsarequirementID returns the FkiEzsigntsarequirementID field if non-nil, zero value otherwise.

### GetFkiEzsigntsarequirementIDOk

`func (o *EzsignfoldertypeRequestV4) GetFkiEzsigntsarequirementIDOk() (*int32, bool)`

GetFkiEzsigntsarequirementIDOk returns a tuple with the FkiEzsigntsarequirementID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntsarequirementID

`func (o *EzsignfoldertypeRequestV4) SetFkiEzsigntsarequirementID(v int32)`

SetFkiEzsigntsarequirementID sets FkiEzsigntsarequirementID field to given value.

### HasFkiEzsigntsarequirementID

`func (o *EzsignfoldertypeRequestV4) HasFkiEzsigntsarequirementID() bool`

HasFkiEzsigntsarequirementID returns a boolean if a field has been set.

### GetFkiFontIDAnnotation

`func (o *EzsignfoldertypeRequestV4) GetFkiFontIDAnnotation() int32`

GetFkiFontIDAnnotation returns the FkiFontIDAnnotation field if non-nil, zero value otherwise.

### GetFkiFontIDAnnotationOk

`func (o *EzsignfoldertypeRequestV4) GetFkiFontIDAnnotationOk() (*int32, bool)`

GetFkiFontIDAnnotationOk returns a tuple with the FkiFontIDAnnotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiFontIDAnnotation

`func (o *EzsignfoldertypeRequestV4) SetFkiFontIDAnnotation(v int32)`

SetFkiFontIDAnnotation sets FkiFontIDAnnotation field to given value.

### HasFkiFontIDAnnotation

`func (o *EzsignfoldertypeRequestV4) HasFkiFontIDAnnotation() bool`

HasFkiFontIDAnnotation returns a boolean if a field has been set.

### GetFkiFontIDFormfield

`func (o *EzsignfoldertypeRequestV4) GetFkiFontIDFormfield() int32`

GetFkiFontIDFormfield returns the FkiFontIDFormfield field if non-nil, zero value otherwise.

### GetFkiFontIDFormfieldOk

`func (o *EzsignfoldertypeRequestV4) GetFkiFontIDFormfieldOk() (*int32, bool)`

GetFkiFontIDFormfieldOk returns a tuple with the FkiFontIDFormfield field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiFontIDFormfield

`func (o *EzsignfoldertypeRequestV4) SetFkiFontIDFormfield(v int32)`

SetFkiFontIDFormfield sets FkiFontIDFormfield field to given value.

### HasFkiFontIDFormfield

`func (o *EzsignfoldertypeRequestV4) HasFkiFontIDFormfield() bool`

HasFkiFontIDFormfield returns a boolean if a field has been set.

### GetFkiFontIDSignature

`func (o *EzsignfoldertypeRequestV4) GetFkiFontIDSignature() int32`

GetFkiFontIDSignature returns the FkiFontIDSignature field if non-nil, zero value otherwise.

### GetFkiFontIDSignatureOk

`func (o *EzsignfoldertypeRequestV4) GetFkiFontIDSignatureOk() (*int32, bool)`

GetFkiFontIDSignatureOk returns a tuple with the FkiFontIDSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiFontIDSignature

`func (o *EzsignfoldertypeRequestV4) SetFkiFontIDSignature(v int32)`

SetFkiFontIDSignature sets FkiFontIDSignature field to given value.

### HasFkiFontIDSignature

`func (o *EzsignfoldertypeRequestV4) HasFkiFontIDSignature() bool`

HasFkiFontIDSignature returns a boolean if a field has been set.

### GetFkiPdfalevelIDConvert

`func (o *EzsignfoldertypeRequestV4) GetFkiPdfalevelIDConvert() int32`

GetFkiPdfalevelIDConvert returns the FkiPdfalevelIDConvert field if non-nil, zero value otherwise.

### GetFkiPdfalevelIDConvertOk

`func (o *EzsignfoldertypeRequestV4) GetFkiPdfalevelIDConvertOk() (*int32, bool)`

GetFkiPdfalevelIDConvertOk returns a tuple with the FkiPdfalevelIDConvert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiPdfalevelIDConvert

`func (o *EzsignfoldertypeRequestV4) SetFkiPdfalevelIDConvert(v int32)`

SetFkiPdfalevelIDConvert sets FkiPdfalevelIDConvert field to given value.

### HasFkiPdfalevelIDConvert

`func (o *EzsignfoldertypeRequestV4) HasFkiPdfalevelIDConvert() bool`

HasFkiPdfalevelIDConvert returns a boolean if a field has been set.

### GetAFkiPdfalevelID

`func (o *EzsignfoldertypeRequestV4) GetAFkiPdfalevelID() []int32`

GetAFkiPdfalevelID returns the AFkiPdfalevelID field if non-nil, zero value otherwise.

### GetAFkiPdfalevelIDOk

`func (o *EzsignfoldertypeRequestV4) GetAFkiPdfalevelIDOk() (*[]int32, bool)`

GetAFkiPdfalevelIDOk returns a tuple with the AFkiPdfalevelID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAFkiPdfalevelID

`func (o *EzsignfoldertypeRequestV4) SetAFkiPdfalevelID(v []int32)`

SetAFkiPdfalevelID sets AFkiPdfalevelID field to given value.

### HasAFkiPdfalevelID

`func (o *EzsignfoldertypeRequestV4) HasAFkiPdfalevelID() bool`

HasAFkiPdfalevelID returns a boolean if a field has been set.

### GetAFkiUserlogintypeID

`func (o *EzsignfoldertypeRequestV4) GetAFkiUserlogintypeID() []int32`

GetAFkiUserlogintypeID returns the AFkiUserlogintypeID field if non-nil, zero value otherwise.

### GetAFkiUserlogintypeIDOk

`func (o *EzsignfoldertypeRequestV4) GetAFkiUserlogintypeIDOk() (*[]int32, bool)`

GetAFkiUserlogintypeIDOk returns a tuple with the AFkiUserlogintypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAFkiUserlogintypeID

`func (o *EzsignfoldertypeRequestV4) SetAFkiUserlogintypeID(v []int32)`

SetAFkiUserlogintypeID sets AFkiUserlogintypeID field to given value.


### GetAFkiUsergroupIDAll

`func (o *EzsignfoldertypeRequestV4) GetAFkiUsergroupIDAll() []int32`

GetAFkiUsergroupIDAll returns the AFkiUsergroupIDAll field if non-nil, zero value otherwise.

### GetAFkiUsergroupIDAllOk

`func (o *EzsignfoldertypeRequestV4) GetAFkiUsergroupIDAllOk() (*[]int32, bool)`

GetAFkiUsergroupIDAllOk returns a tuple with the AFkiUsergroupIDAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAFkiUsergroupIDAll

`func (o *EzsignfoldertypeRequestV4) SetAFkiUsergroupIDAll(v []int32)`

SetAFkiUsergroupIDAll sets AFkiUsergroupIDAll field to given value.

### HasAFkiUsergroupIDAll

`func (o *EzsignfoldertypeRequestV4) HasAFkiUsergroupIDAll() bool`

HasAFkiUsergroupIDAll returns a boolean if a field has been set.

### GetAFkiUsergroupIDRestricted

`func (o *EzsignfoldertypeRequestV4) GetAFkiUsergroupIDRestricted() []int32`

GetAFkiUsergroupIDRestricted returns the AFkiUsergroupIDRestricted field if non-nil, zero value otherwise.

### GetAFkiUsergroupIDRestrictedOk

`func (o *EzsignfoldertypeRequestV4) GetAFkiUsergroupIDRestrictedOk() (*[]int32, bool)`

GetAFkiUsergroupIDRestrictedOk returns a tuple with the AFkiUsergroupIDRestricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAFkiUsergroupIDRestricted

`func (o *EzsignfoldertypeRequestV4) SetAFkiUsergroupIDRestricted(v []int32)`

SetAFkiUsergroupIDRestricted sets AFkiUsergroupIDRestricted field to given value.

### HasAFkiUsergroupIDRestricted

`func (o *EzsignfoldertypeRequestV4) HasAFkiUsergroupIDRestricted() bool`

HasAFkiUsergroupIDRestricted returns a boolean if a field has been set.

### GetAFkiUsergroupIDTemplate

`func (o *EzsignfoldertypeRequestV4) GetAFkiUsergroupIDTemplate() []int32`

GetAFkiUsergroupIDTemplate returns the AFkiUsergroupIDTemplate field if non-nil, zero value otherwise.

### GetAFkiUsergroupIDTemplateOk

`func (o *EzsignfoldertypeRequestV4) GetAFkiUsergroupIDTemplateOk() (*[]int32, bool)`

GetAFkiUsergroupIDTemplateOk returns a tuple with the AFkiUsergroupIDTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAFkiUsergroupIDTemplate

`func (o *EzsignfoldertypeRequestV4) SetAFkiUsergroupIDTemplate(v []int32)`

SetAFkiUsergroupIDTemplate sets AFkiUsergroupIDTemplate field to given value.

### HasAFkiUsergroupIDTemplate

`func (o *EzsignfoldertypeRequestV4) HasAFkiUsergroupIDTemplate() bool`

HasAFkiUsergroupIDTemplate returns a boolean if a field has been set.

### GetEEzsignfoldertypeSignaturedatedisplay

`func (o *EzsignfoldertypeRequestV4) GetEEzsignfoldertypeSignaturedatedisplay() FieldEEzsignfoldertypeSignaturedatedisplay`

GetEEzsignfoldertypeSignaturedatedisplay returns the EEzsignfoldertypeSignaturedatedisplay field if non-nil, zero value otherwise.

### GetEEzsignfoldertypeSignaturedatedisplayOk

`func (o *EzsignfoldertypeRequestV4) GetEEzsignfoldertypeSignaturedatedisplayOk() (*FieldEEzsignfoldertypeSignaturedatedisplay, bool)`

GetEEzsignfoldertypeSignaturedatedisplayOk returns a tuple with the EEzsignfoldertypeSignaturedatedisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignfoldertypeSignaturedatedisplay

`func (o *EzsignfoldertypeRequestV4) SetEEzsignfoldertypeSignaturedatedisplay(v FieldEEzsignfoldertypeSignaturedatedisplay)`

SetEEzsignfoldertypeSignaturedatedisplay sets EEzsignfoldertypeSignaturedatedisplay field to given value.


### GetSEzsignfoldertypeSignaturedatecustomformat

`func (o *EzsignfoldertypeRequestV4) GetSEzsignfoldertypeSignaturedatecustomformat() string`

GetSEzsignfoldertypeSignaturedatecustomformat returns the SEzsignfoldertypeSignaturedatecustomformat field if non-nil, zero value otherwise.

### GetSEzsignfoldertypeSignaturedatecustomformatOk

`func (o *EzsignfoldertypeRequestV4) GetSEzsignfoldertypeSignaturedatecustomformatOk() (*string, bool)`

GetSEzsignfoldertypeSignaturedatecustomformatOk returns a tuple with the SEzsignfoldertypeSignaturedatecustomformat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignfoldertypeSignaturedatecustomformat

`func (o *EzsignfoldertypeRequestV4) SetSEzsignfoldertypeSignaturedatecustomformat(v string)`

SetSEzsignfoldertypeSignaturedatecustomformat sets SEzsignfoldertypeSignaturedatecustomformat field to given value.

### HasSEzsignfoldertypeSignaturedatecustomformat

`func (o *EzsignfoldertypeRequestV4) HasSEzsignfoldertypeSignaturedatecustomformat() bool`

HasSEzsignfoldertypeSignaturedatecustomformat returns a boolean if a field has been set.

### GetEEzsignfoldertypeDocumentdependency

`func (o *EzsignfoldertypeRequestV4) GetEEzsignfoldertypeDocumentdependency() FieldEEzsignfoldertypeDocumentdependency`

GetEEzsignfoldertypeDocumentdependency returns the EEzsignfoldertypeDocumentdependency field if non-nil, zero value otherwise.

### GetEEzsignfoldertypeDocumentdependencyOk

`func (o *EzsignfoldertypeRequestV4) GetEEzsignfoldertypeDocumentdependencyOk() (*FieldEEzsignfoldertypeDocumentdependency, bool)`

GetEEzsignfoldertypeDocumentdependencyOk returns a tuple with the EEzsignfoldertypeDocumentdependency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignfoldertypeDocumentdependency

`func (o *EzsignfoldertypeRequestV4) SetEEzsignfoldertypeDocumentdependency(v FieldEEzsignfoldertypeDocumentdependency)`

SetEEzsignfoldertypeDocumentdependency sets EEzsignfoldertypeDocumentdependency field to given value.

### HasEEzsignfoldertypeDocumentdependency

`func (o *EzsignfoldertypeRequestV4) HasEEzsignfoldertypeDocumentdependency() bool`

HasEEzsignfoldertypeDocumentdependency returns a boolean if a field has been set.

### GetEEzsignfoldertypeDocumentmerge

`func (o *EzsignfoldertypeRequestV4) GetEEzsignfoldertypeDocumentmerge() FieldEEzsignfoldertypeDocumentmerge`

GetEEzsignfoldertypeDocumentmerge returns the EEzsignfoldertypeDocumentmerge field if non-nil, zero value otherwise.

### GetEEzsignfoldertypeDocumentmergeOk

`func (o *EzsignfoldertypeRequestV4) GetEEzsignfoldertypeDocumentmergeOk() (*FieldEEzsignfoldertypeDocumentmerge, bool)`

GetEEzsignfoldertypeDocumentmergeOk returns a tuple with the EEzsignfoldertypeDocumentmerge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignfoldertypeDocumentmerge

`func (o *EzsignfoldertypeRequestV4) SetEEzsignfoldertypeDocumentmerge(v FieldEEzsignfoldertypeDocumentmerge)`

SetEEzsignfoldertypeDocumentmerge sets EEzsignfoldertypeDocumentmerge field to given value.

### HasEEzsignfoldertypeDocumentmerge

`func (o *EzsignfoldertypeRequestV4) HasEEzsignfoldertypeDocumentmerge() bool`

HasEEzsignfoldertypeDocumentmerge returns a boolean if a field has been set.

### GetSEmailAddressSigned

`func (o *EzsignfoldertypeRequestV4) GetSEmailAddressSigned() string`

GetSEmailAddressSigned returns the SEmailAddressSigned field if non-nil, zero value otherwise.

### GetSEmailAddressSignedOk

`func (o *EzsignfoldertypeRequestV4) GetSEmailAddressSignedOk() (*string, bool)`

GetSEmailAddressSignedOk returns a tuple with the SEmailAddressSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEmailAddressSigned

`func (o *EzsignfoldertypeRequestV4) SetSEmailAddressSigned(v string)`

SetSEmailAddressSigned sets SEmailAddressSigned field to given value.

### HasSEmailAddressSigned

`func (o *EzsignfoldertypeRequestV4) HasSEmailAddressSigned() bool`

HasSEmailAddressSigned returns a boolean if a field has been set.

### GetSEmailAddressSummary

`func (o *EzsignfoldertypeRequestV4) GetSEmailAddressSummary() string`

GetSEmailAddressSummary returns the SEmailAddressSummary field if non-nil, zero value otherwise.

### GetSEmailAddressSummaryOk

`func (o *EzsignfoldertypeRequestV4) GetSEmailAddressSummaryOk() (*string, bool)`

GetSEmailAddressSummaryOk returns a tuple with the SEmailAddressSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEmailAddressSummary

`func (o *EzsignfoldertypeRequestV4) SetSEmailAddressSummary(v string)`

SetSEmailAddressSummary sets SEmailAddressSummary field to given value.

### HasSEmailAddressSummary

`func (o *EzsignfoldertypeRequestV4) HasSEmailAddressSummary() bool`

HasSEmailAddressSummary returns a boolean if a field has been set.

### GetEEzsignfoldertypePdfarequirement

`func (o *EzsignfoldertypeRequestV4) GetEEzsignfoldertypePdfarequirement() FieldEEzsignfoldertypePdfarequirement`

GetEEzsignfoldertypePdfarequirement returns the EEzsignfoldertypePdfarequirement field if non-nil, zero value otherwise.

### GetEEzsignfoldertypePdfarequirementOk

`func (o *EzsignfoldertypeRequestV4) GetEEzsignfoldertypePdfarequirementOk() (*FieldEEzsignfoldertypePdfarequirement, bool)`

GetEEzsignfoldertypePdfarequirementOk returns a tuple with the EEzsignfoldertypePdfarequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignfoldertypePdfarequirement

`func (o *EzsignfoldertypeRequestV4) SetEEzsignfoldertypePdfarequirement(v FieldEEzsignfoldertypePdfarequirement)`

SetEEzsignfoldertypePdfarequirement sets EEzsignfoldertypePdfarequirement field to given value.

### HasEEzsignfoldertypePdfarequirement

`func (o *EzsignfoldertypeRequestV4) HasEEzsignfoldertypePdfarequirement() bool`

HasEEzsignfoldertypePdfarequirement returns a boolean if a field has been set.

### GetEEzsignfoldertypePdfanoncompliantaction

`func (o *EzsignfoldertypeRequestV4) GetEEzsignfoldertypePdfanoncompliantaction() FieldEEzsignfoldertypePdfanoncompliantaction`

GetEEzsignfoldertypePdfanoncompliantaction returns the EEzsignfoldertypePdfanoncompliantaction field if non-nil, zero value otherwise.

### GetEEzsignfoldertypePdfanoncompliantactionOk

`func (o *EzsignfoldertypeRequestV4) GetEEzsignfoldertypePdfanoncompliantactionOk() (*FieldEEzsignfoldertypePdfanoncompliantaction, bool)`

GetEEzsignfoldertypePdfanoncompliantactionOk returns a tuple with the EEzsignfoldertypePdfanoncompliantaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignfoldertypePdfanoncompliantaction

`func (o *EzsignfoldertypeRequestV4) SetEEzsignfoldertypePdfanoncompliantaction(v FieldEEzsignfoldertypePdfanoncompliantaction)`

SetEEzsignfoldertypePdfanoncompliantaction sets EEzsignfoldertypePdfanoncompliantaction field to given value.

### HasEEzsignfoldertypePdfanoncompliantaction

`func (o *EzsignfoldertypeRequestV4) HasEEzsignfoldertypePdfanoncompliantaction() bool`

HasEEzsignfoldertypePdfanoncompliantaction returns a boolean if a field has been set.

### GetEEzsignfoldertypePrivacylevel

`func (o *EzsignfoldertypeRequestV4) GetEEzsignfoldertypePrivacylevel() FieldEEzsignfoldertypePrivacylevel`

GetEEzsignfoldertypePrivacylevel returns the EEzsignfoldertypePrivacylevel field if non-nil, zero value otherwise.

### GetEEzsignfoldertypePrivacylevelOk

`func (o *EzsignfoldertypeRequestV4) GetEEzsignfoldertypePrivacylevelOk() (*FieldEEzsignfoldertypePrivacylevel, bool)`

GetEEzsignfoldertypePrivacylevelOk returns a tuple with the EEzsignfoldertypePrivacylevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignfoldertypePrivacylevel

`func (o *EzsignfoldertypeRequestV4) SetEEzsignfoldertypePrivacylevel(v FieldEEzsignfoldertypePrivacylevel)`

SetEEzsignfoldertypePrivacylevel sets EEzsignfoldertypePrivacylevel field to given value.


### GetIEzsignfoldertypeFontsizeannotation

`func (o *EzsignfoldertypeRequestV4) GetIEzsignfoldertypeFontsizeannotation() int32`

GetIEzsignfoldertypeFontsizeannotation returns the IEzsignfoldertypeFontsizeannotation field if non-nil, zero value otherwise.

### GetIEzsignfoldertypeFontsizeannotationOk

`func (o *EzsignfoldertypeRequestV4) GetIEzsignfoldertypeFontsizeannotationOk() (*int32, bool)`

GetIEzsignfoldertypeFontsizeannotationOk returns a tuple with the IEzsignfoldertypeFontsizeannotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignfoldertypeFontsizeannotation

`func (o *EzsignfoldertypeRequestV4) SetIEzsignfoldertypeFontsizeannotation(v int32)`

SetIEzsignfoldertypeFontsizeannotation sets IEzsignfoldertypeFontsizeannotation field to given value.

### HasIEzsignfoldertypeFontsizeannotation

`func (o *EzsignfoldertypeRequestV4) HasIEzsignfoldertypeFontsizeannotation() bool`

HasIEzsignfoldertypeFontsizeannotation returns a boolean if a field has been set.

### GetIEzsignfoldertypeFontsizeformfield

`func (o *EzsignfoldertypeRequestV4) GetIEzsignfoldertypeFontsizeformfield() int32`

GetIEzsignfoldertypeFontsizeformfield returns the IEzsignfoldertypeFontsizeformfield field if non-nil, zero value otherwise.

### GetIEzsignfoldertypeFontsizeformfieldOk

`func (o *EzsignfoldertypeRequestV4) GetIEzsignfoldertypeFontsizeformfieldOk() (*int32, bool)`

GetIEzsignfoldertypeFontsizeformfieldOk returns a tuple with the IEzsignfoldertypeFontsizeformfield field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignfoldertypeFontsizeformfield

`func (o *EzsignfoldertypeRequestV4) SetIEzsignfoldertypeFontsizeformfield(v int32)`

SetIEzsignfoldertypeFontsizeformfield sets IEzsignfoldertypeFontsizeformfield field to given value.

### HasIEzsignfoldertypeFontsizeformfield

`func (o *EzsignfoldertypeRequestV4) HasIEzsignfoldertypeFontsizeformfield() bool`

HasIEzsignfoldertypeFontsizeformfield returns a boolean if a field has been set.

### GetIEzsignfoldertypeSendreminderfirstdays

`func (o *EzsignfoldertypeRequestV4) GetIEzsignfoldertypeSendreminderfirstdays() int32`

GetIEzsignfoldertypeSendreminderfirstdays returns the IEzsignfoldertypeSendreminderfirstdays field if non-nil, zero value otherwise.

### GetIEzsignfoldertypeSendreminderfirstdaysOk

`func (o *EzsignfoldertypeRequestV4) GetIEzsignfoldertypeSendreminderfirstdaysOk() (*int32, bool)`

GetIEzsignfoldertypeSendreminderfirstdaysOk returns a tuple with the IEzsignfoldertypeSendreminderfirstdays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignfoldertypeSendreminderfirstdays

`func (o *EzsignfoldertypeRequestV4) SetIEzsignfoldertypeSendreminderfirstdays(v int32)`

SetIEzsignfoldertypeSendreminderfirstdays sets IEzsignfoldertypeSendreminderfirstdays field to given value.

### HasIEzsignfoldertypeSendreminderfirstdays

`func (o *EzsignfoldertypeRequestV4) HasIEzsignfoldertypeSendreminderfirstdays() bool`

HasIEzsignfoldertypeSendreminderfirstdays returns a boolean if a field has been set.

### GetIEzsignfoldertypeSendreminderotherdays

`func (o *EzsignfoldertypeRequestV4) GetIEzsignfoldertypeSendreminderotherdays() int32`

GetIEzsignfoldertypeSendreminderotherdays returns the IEzsignfoldertypeSendreminderotherdays field if non-nil, zero value otherwise.

### GetIEzsignfoldertypeSendreminderotherdaysOk

`func (o *EzsignfoldertypeRequestV4) GetIEzsignfoldertypeSendreminderotherdaysOk() (*int32, bool)`

GetIEzsignfoldertypeSendreminderotherdaysOk returns a tuple with the IEzsignfoldertypeSendreminderotherdays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignfoldertypeSendreminderotherdays

`func (o *EzsignfoldertypeRequestV4) SetIEzsignfoldertypeSendreminderotherdays(v int32)`

SetIEzsignfoldertypeSendreminderotherdays sets IEzsignfoldertypeSendreminderotherdays field to given value.

### HasIEzsignfoldertypeSendreminderotherdays

`func (o *EzsignfoldertypeRequestV4) HasIEzsignfoldertypeSendreminderotherdays() bool`

HasIEzsignfoldertypeSendreminderotherdays returns a boolean if a field has been set.

### GetIEzsignfoldertypeArchivaldays

`func (o *EzsignfoldertypeRequestV4) GetIEzsignfoldertypeArchivaldays() int32`

GetIEzsignfoldertypeArchivaldays returns the IEzsignfoldertypeArchivaldays field if non-nil, zero value otherwise.

### GetIEzsignfoldertypeArchivaldaysOk

`func (o *EzsignfoldertypeRequestV4) GetIEzsignfoldertypeArchivaldaysOk() (*int32, bool)`

GetIEzsignfoldertypeArchivaldaysOk returns a tuple with the IEzsignfoldertypeArchivaldays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignfoldertypeArchivaldays

`func (o *EzsignfoldertypeRequestV4) SetIEzsignfoldertypeArchivaldays(v int32)`

SetIEzsignfoldertypeArchivaldays sets IEzsignfoldertypeArchivaldays field to given value.


### GetEEzsignfoldertypeDisposal

`func (o *EzsignfoldertypeRequestV4) GetEEzsignfoldertypeDisposal() FieldEEzsignfoldertypeDisposal`

GetEEzsignfoldertypeDisposal returns the EEzsignfoldertypeDisposal field if non-nil, zero value otherwise.

### GetEEzsignfoldertypeDisposalOk

`func (o *EzsignfoldertypeRequestV4) GetEEzsignfoldertypeDisposalOk() (*FieldEEzsignfoldertypeDisposal, bool)`

GetEEzsignfoldertypeDisposalOk returns a tuple with the EEzsignfoldertypeDisposal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignfoldertypeDisposal

`func (o *EzsignfoldertypeRequestV4) SetEEzsignfoldertypeDisposal(v FieldEEzsignfoldertypeDisposal)`

SetEEzsignfoldertypeDisposal sets EEzsignfoldertypeDisposal field to given value.


### GetEEzsignfoldertypeCompletion

`func (o *EzsignfoldertypeRequestV4) GetEEzsignfoldertypeCompletion() FieldEEzsignfoldertypeCompletion`

GetEEzsignfoldertypeCompletion returns the EEzsignfoldertypeCompletion field if non-nil, zero value otherwise.

### GetEEzsignfoldertypeCompletionOk

`func (o *EzsignfoldertypeRequestV4) GetEEzsignfoldertypeCompletionOk() (*FieldEEzsignfoldertypeCompletion, bool)`

GetEEzsignfoldertypeCompletionOk returns a tuple with the EEzsignfoldertypeCompletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignfoldertypeCompletion

`func (o *EzsignfoldertypeRequestV4) SetEEzsignfoldertypeCompletion(v FieldEEzsignfoldertypeCompletion)`

SetEEzsignfoldertypeCompletion sets EEzsignfoldertypeCompletion field to given value.


### GetIEzsignfoldertypeDisposaldays

`func (o *EzsignfoldertypeRequestV4) GetIEzsignfoldertypeDisposaldays() int32`

GetIEzsignfoldertypeDisposaldays returns the IEzsignfoldertypeDisposaldays field if non-nil, zero value otherwise.

### GetIEzsignfoldertypeDisposaldaysOk

`func (o *EzsignfoldertypeRequestV4) GetIEzsignfoldertypeDisposaldaysOk() (*int32, bool)`

GetIEzsignfoldertypeDisposaldaysOk returns a tuple with the IEzsignfoldertypeDisposaldays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignfoldertypeDisposaldays

`func (o *EzsignfoldertypeRequestV4) SetIEzsignfoldertypeDisposaldays(v int32)`

SetIEzsignfoldertypeDisposaldays sets IEzsignfoldertypeDisposaldays field to given value.

### HasIEzsignfoldertypeDisposaldays

`func (o *EzsignfoldertypeRequestV4) HasIEzsignfoldertypeDisposaldays() bool`

HasIEzsignfoldertypeDisposaldays returns a boolean if a field has been set.

### GetIEzsignfoldertypeDeadlinedays

`func (o *EzsignfoldertypeRequestV4) GetIEzsignfoldertypeDeadlinedays() int32`

GetIEzsignfoldertypeDeadlinedays returns the IEzsignfoldertypeDeadlinedays field if non-nil, zero value otherwise.

### GetIEzsignfoldertypeDeadlinedaysOk

`func (o *EzsignfoldertypeRequestV4) GetIEzsignfoldertypeDeadlinedaysOk() (*int32, bool)`

GetIEzsignfoldertypeDeadlinedaysOk returns a tuple with the IEzsignfoldertypeDeadlinedays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignfoldertypeDeadlinedays

`func (o *EzsignfoldertypeRequestV4) SetIEzsignfoldertypeDeadlinedays(v int32)`

SetIEzsignfoldertypeDeadlinedays sets IEzsignfoldertypeDeadlinedays field to given value.


### GetBEzsignfoldertypePrematurelyendautomatically

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypePrematurelyendautomatically() bool`

GetBEzsignfoldertypePrematurelyendautomatically returns the BEzsignfoldertypePrematurelyendautomatically field if non-nil, zero value otherwise.

### GetBEzsignfoldertypePrematurelyendautomaticallyOk

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypePrematurelyendautomaticallyOk() (*bool, bool)`

GetBEzsignfoldertypePrematurelyendautomaticallyOk returns a tuple with the BEzsignfoldertypePrematurelyendautomatically field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypePrematurelyendautomatically

`func (o *EzsignfoldertypeRequestV4) SetBEzsignfoldertypePrematurelyendautomatically(v bool)`

SetBEzsignfoldertypePrematurelyendautomatically sets BEzsignfoldertypePrematurelyendautomatically field to given value.

### HasBEzsignfoldertypePrematurelyendautomatically

`func (o *EzsignfoldertypeRequestV4) HasBEzsignfoldertypePrematurelyendautomatically() bool`

HasBEzsignfoldertypePrematurelyendautomatically returns a boolean if a field has been set.

### GetIEzsignfoldertypePrematurelyendautomaticallydays

`func (o *EzsignfoldertypeRequestV4) GetIEzsignfoldertypePrematurelyendautomaticallydays() int32`

GetIEzsignfoldertypePrematurelyendautomaticallydays returns the IEzsignfoldertypePrematurelyendautomaticallydays field if non-nil, zero value otherwise.

### GetIEzsignfoldertypePrematurelyendautomaticallydaysOk

`func (o *EzsignfoldertypeRequestV4) GetIEzsignfoldertypePrematurelyendautomaticallydaysOk() (*int32, bool)`

GetIEzsignfoldertypePrematurelyendautomaticallydaysOk returns a tuple with the IEzsignfoldertypePrematurelyendautomaticallydays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignfoldertypePrematurelyendautomaticallydays

`func (o *EzsignfoldertypeRequestV4) SetIEzsignfoldertypePrematurelyendautomaticallydays(v int32)`

SetIEzsignfoldertypePrematurelyendautomaticallydays sets IEzsignfoldertypePrematurelyendautomaticallydays field to given value.

### HasIEzsignfoldertypePrematurelyendautomaticallydays

`func (o *EzsignfoldertypeRequestV4) HasIEzsignfoldertypePrematurelyendautomaticallydays() bool`

HasIEzsignfoldertypePrematurelyendautomaticallydays returns a boolean if a field has been set.

### GetBEzsignfoldertypeAutomaticsignature

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeAutomaticsignature() bool`

GetBEzsignfoldertypeAutomaticsignature returns the BEzsignfoldertypeAutomaticsignature field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeAutomaticsignatureOk

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeAutomaticsignatureOk() (*bool, bool)`

GetBEzsignfoldertypeAutomaticsignatureOk returns a tuple with the BEzsignfoldertypeAutomaticsignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeAutomaticsignature

`func (o *EzsignfoldertypeRequestV4) SetBEzsignfoldertypeAutomaticsignature(v bool)`

SetBEzsignfoldertypeAutomaticsignature sets BEzsignfoldertypeAutomaticsignature field to given value.

### HasBEzsignfoldertypeAutomaticsignature

`func (o *EzsignfoldertypeRequestV4) HasBEzsignfoldertypeAutomaticsignature() bool`

HasBEzsignfoldertypeAutomaticsignature returns a boolean if a field has been set.

### GetBEzsignfoldertypeDelegate

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeDelegate() bool`

GetBEzsignfoldertypeDelegate returns the BEzsignfoldertypeDelegate field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeDelegateOk

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeDelegateOk() (*bool, bool)`

GetBEzsignfoldertypeDelegateOk returns a tuple with the BEzsignfoldertypeDelegate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeDelegate

`func (o *EzsignfoldertypeRequestV4) SetBEzsignfoldertypeDelegate(v bool)`

SetBEzsignfoldertypeDelegate sets BEzsignfoldertypeDelegate field to given value.

### HasBEzsignfoldertypeDelegate

`func (o *EzsignfoldertypeRequestV4) HasBEzsignfoldertypeDelegate() bool`

HasBEzsignfoldertypeDelegate returns a boolean if a field has been set.

### GetBEzsignfoldertypeDiscussion

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeDiscussion() bool`

GetBEzsignfoldertypeDiscussion returns the BEzsignfoldertypeDiscussion field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeDiscussionOk

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeDiscussionOk() (*bool, bool)`

GetBEzsignfoldertypeDiscussionOk returns a tuple with the BEzsignfoldertypeDiscussion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeDiscussion

`func (o *EzsignfoldertypeRequestV4) SetBEzsignfoldertypeDiscussion(v bool)`

SetBEzsignfoldertypeDiscussion sets BEzsignfoldertypeDiscussion field to given value.

### HasBEzsignfoldertypeDiscussion

`func (o *EzsignfoldertypeRequestV4) HasBEzsignfoldertypeDiscussion() bool`

HasBEzsignfoldertypeDiscussion returns a boolean if a field has been set.

### GetBEzsignfoldertypeLogrecipientinproof

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeLogrecipientinproof() bool`

GetBEzsignfoldertypeLogrecipientinproof returns the BEzsignfoldertypeLogrecipientinproof field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeLogrecipientinproofOk

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeLogrecipientinproofOk() (*bool, bool)`

GetBEzsignfoldertypeLogrecipientinproofOk returns a tuple with the BEzsignfoldertypeLogrecipientinproof field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeLogrecipientinproof

`func (o *EzsignfoldertypeRequestV4) SetBEzsignfoldertypeLogrecipientinproof(v bool)`

SetBEzsignfoldertypeLogrecipientinproof sets BEzsignfoldertypeLogrecipientinproof field to given value.

### HasBEzsignfoldertypeLogrecipientinproof

`func (o *EzsignfoldertypeRequestV4) HasBEzsignfoldertypeLogrecipientinproof() bool`

HasBEzsignfoldertypeLogrecipientinproof returns a boolean if a field has been set.

### GetBEzsignfoldertypeReassignezsignsigner

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeReassignezsignsigner() bool`

GetBEzsignfoldertypeReassignezsignsigner returns the BEzsignfoldertypeReassignezsignsigner field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeReassignezsignsignerOk

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeReassignezsignsignerOk() (*bool, bool)`

GetBEzsignfoldertypeReassignezsignsignerOk returns a tuple with the BEzsignfoldertypeReassignezsignsigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeReassignezsignsigner

`func (o *EzsignfoldertypeRequestV4) SetBEzsignfoldertypeReassignezsignsigner(v bool)`

SetBEzsignfoldertypeReassignezsignsigner sets BEzsignfoldertypeReassignezsignsigner field to given value.

### HasBEzsignfoldertypeReassignezsignsigner

`func (o *EzsignfoldertypeRequestV4) HasBEzsignfoldertypeReassignezsignsigner() bool`

HasBEzsignfoldertypeReassignezsignsigner returns a boolean if a field has been set.

### GetBEzsignfoldertypeReassignuser

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeReassignuser() bool`

GetBEzsignfoldertypeReassignuser returns the BEzsignfoldertypeReassignuser field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeReassignuserOk

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeReassignuserOk() (*bool, bool)`

GetBEzsignfoldertypeReassignuserOk returns a tuple with the BEzsignfoldertypeReassignuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeReassignuser

`func (o *EzsignfoldertypeRequestV4) SetBEzsignfoldertypeReassignuser(v bool)`

SetBEzsignfoldertypeReassignuser sets BEzsignfoldertypeReassignuser field to given value.

### HasBEzsignfoldertypeReassignuser

`func (o *EzsignfoldertypeRequestV4) HasBEzsignfoldertypeReassignuser() bool`

HasBEzsignfoldertypeReassignuser returns a boolean if a field has been set.

### GetBEzsignfoldertypeReassigngroup

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeReassigngroup() bool`

GetBEzsignfoldertypeReassigngroup returns the BEzsignfoldertypeReassigngroup field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeReassigngroupOk

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeReassigngroupOk() (*bool, bool)`

GetBEzsignfoldertypeReassigngroupOk returns a tuple with the BEzsignfoldertypeReassigngroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeReassigngroup

`func (o *EzsignfoldertypeRequestV4) SetBEzsignfoldertypeReassigngroup(v bool)`

SetBEzsignfoldertypeReassigngroup sets BEzsignfoldertypeReassigngroup field to given value.

### HasBEzsignfoldertypeReassigngroup

`func (o *EzsignfoldertypeRequestV4) HasBEzsignfoldertypeReassigngroup() bool`

HasBEzsignfoldertypeReassigngroup returns a boolean if a field has been set.

### GetBEzsignfoldertypeSenddocumentmergetoemail

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeSenddocumentmergetoemail() bool`

GetBEzsignfoldertypeSenddocumentmergetoemail returns the BEzsignfoldertypeSenddocumentmergetoemail field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSenddocumentmergetoemailOk

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeSenddocumentmergetoemailOk() (*bool, bool)`

GetBEzsignfoldertypeSenddocumentmergetoemailOk returns a tuple with the BEzsignfoldertypeSenddocumentmergetoemail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSenddocumentmergetoemail

`func (o *EzsignfoldertypeRequestV4) SetBEzsignfoldertypeSenddocumentmergetoemail(v bool)`

SetBEzsignfoldertypeSenddocumentmergetoemail sets BEzsignfoldertypeSenddocumentmergetoemail field to given value.

### HasBEzsignfoldertypeSenddocumentmergetoemail

`func (o *EzsignfoldertypeRequestV4) HasBEzsignfoldertypeSenddocumentmergetoemail() bool`

HasBEzsignfoldertypeSenddocumentmergetoemail returns a boolean if a field has been set.

### GetBEzsignfoldertypeSenddocumentmergetoezsignsigner

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeSenddocumentmergetoezsignsigner() bool`

GetBEzsignfoldertypeSenddocumentmergetoezsignsigner returns the BEzsignfoldertypeSenddocumentmergetoezsignsigner field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSenddocumentmergetoezsignsignerOk

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeSenddocumentmergetoezsignsignerOk() (*bool, bool)`

GetBEzsignfoldertypeSenddocumentmergetoezsignsignerOk returns a tuple with the BEzsignfoldertypeSenddocumentmergetoezsignsigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSenddocumentmergetoezsignsigner

`func (o *EzsignfoldertypeRequestV4) SetBEzsignfoldertypeSenddocumentmergetoezsignsigner(v bool)`

SetBEzsignfoldertypeSenddocumentmergetoezsignsigner sets BEzsignfoldertypeSenddocumentmergetoezsignsigner field to given value.

### HasBEzsignfoldertypeSenddocumentmergetoezsignsigner

`func (o *EzsignfoldertypeRequestV4) HasBEzsignfoldertypeSenddocumentmergetoezsignsigner() bool`

HasBEzsignfoldertypeSenddocumentmergetoezsignsigner returns a boolean if a field has been set.

### GetBEzsignfoldertypeSenddocumentmergetoreceivealldocument

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeSenddocumentmergetoreceivealldocument() bool`

GetBEzsignfoldertypeSenddocumentmergetoreceivealldocument returns the BEzsignfoldertypeSenddocumentmergetoreceivealldocument field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSenddocumentmergetoreceivealldocumentOk

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeSenddocumentmergetoreceivealldocumentOk() (*bool, bool)`

GetBEzsignfoldertypeSenddocumentmergetoreceivealldocumentOk returns a tuple with the BEzsignfoldertypeSenddocumentmergetoreceivealldocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSenddocumentmergetoreceivealldocument

`func (o *EzsignfoldertypeRequestV4) SetBEzsignfoldertypeSenddocumentmergetoreceivealldocument(v bool)`

SetBEzsignfoldertypeSenddocumentmergetoreceivealldocument sets BEzsignfoldertypeSenddocumentmergetoreceivealldocument field to given value.

### HasBEzsignfoldertypeSenddocumentmergetoreceivealldocument

`func (o *EzsignfoldertypeRequestV4) HasBEzsignfoldertypeSenddocumentmergetoreceivealldocument() bool`

HasBEzsignfoldertypeSenddocumentmergetoreceivealldocument returns a boolean if a field has been set.

### GetBEzsignfoldertypeSenddocumentmergetouser

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeSenddocumentmergetouser() bool`

GetBEzsignfoldertypeSenddocumentmergetouser returns the BEzsignfoldertypeSenddocumentmergetouser field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSenddocumentmergetouserOk

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeSenddocumentmergetouserOk() (*bool, bool)`

GetBEzsignfoldertypeSenddocumentmergetouserOk returns a tuple with the BEzsignfoldertypeSenddocumentmergetouser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSenddocumentmergetouser

`func (o *EzsignfoldertypeRequestV4) SetBEzsignfoldertypeSenddocumentmergetouser(v bool)`

SetBEzsignfoldertypeSenddocumentmergetouser sets BEzsignfoldertypeSenddocumentmergetouser field to given value.

### HasBEzsignfoldertypeSenddocumentmergetouser

`func (o *EzsignfoldertypeRequestV4) HasBEzsignfoldertypeSenddocumentmergetouser() bool`

HasBEzsignfoldertypeSenddocumentmergetouser returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendsignedtoezsignsigner

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeSendsignedtoezsignsigner() bool`

GetBEzsignfoldertypeSendsignedtoezsignsigner returns the BEzsignfoldertypeSendsignedtoezsignsigner field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsignedtoezsignsignerOk

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeSendsignedtoezsignsignerOk() (*bool, bool)`

GetBEzsignfoldertypeSendsignedtoezsignsignerOk returns a tuple with the BEzsignfoldertypeSendsignedtoezsignsigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsignedtoezsignsigner

`func (o *EzsignfoldertypeRequestV4) SetBEzsignfoldertypeSendsignedtoezsignsigner(v bool)`

SetBEzsignfoldertypeSendsignedtoezsignsigner sets BEzsignfoldertypeSendsignedtoezsignsigner field to given value.

### HasBEzsignfoldertypeSendsignedtoezsignsigner

`func (o *EzsignfoldertypeRequestV4) HasBEzsignfoldertypeSendsignedtoezsignsigner() bool`

HasBEzsignfoldertypeSendsignedtoezsignsigner returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendsignedtouser

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeSendsignedtouser() bool`

GetBEzsignfoldertypeSendsignedtouser returns the BEzsignfoldertypeSendsignedtouser field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsignedtouserOk

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeSendsignedtouserOk() (*bool, bool)`

GetBEzsignfoldertypeSendsignedtouserOk returns a tuple with the BEzsignfoldertypeSendsignedtouser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsignedtouser

`func (o *EzsignfoldertypeRequestV4) SetBEzsignfoldertypeSendsignedtouser(v bool)`

SetBEzsignfoldertypeSendsignedtouser sets BEzsignfoldertypeSendsignedtouser field to given value.

### HasBEzsignfoldertypeSendsignedtouser

`func (o *EzsignfoldertypeRequestV4) HasBEzsignfoldertypeSendsignedtouser() bool`

HasBEzsignfoldertypeSendsignedtouser returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendattachmentezsignsigner

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeSendattachmentezsignsigner() bool`

GetBEzsignfoldertypeSendattachmentezsignsigner returns the BEzsignfoldertypeSendattachmentezsignsigner field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendattachmentezsignsignerOk

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeSendattachmentezsignsignerOk() (*bool, bool)`

GetBEzsignfoldertypeSendattachmentezsignsignerOk returns a tuple with the BEzsignfoldertypeSendattachmentezsignsigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendattachmentezsignsigner

`func (o *EzsignfoldertypeRequestV4) SetBEzsignfoldertypeSendattachmentezsignsigner(v bool)`

SetBEzsignfoldertypeSendattachmentezsignsigner sets BEzsignfoldertypeSendattachmentezsignsigner field to given value.

### HasBEzsignfoldertypeSendattachmentezsignsigner

`func (o *EzsignfoldertypeRequestV4) HasBEzsignfoldertypeSendattachmentezsignsigner() bool`

HasBEzsignfoldertypeSendattachmentezsignsigner returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendsignatureattachmentezsignsigner

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeSendsignatureattachmentezsignsigner() bool`

GetBEzsignfoldertypeSendsignatureattachmentezsignsigner returns the BEzsignfoldertypeSendsignatureattachmentezsignsigner field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsignatureattachmentezsignsignerOk

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeSendsignatureattachmentezsignsignerOk() (*bool, bool)`

GetBEzsignfoldertypeSendsignatureattachmentezsignsignerOk returns a tuple with the BEzsignfoldertypeSendsignatureattachmentezsignsigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsignatureattachmentezsignsigner

`func (o *EzsignfoldertypeRequestV4) SetBEzsignfoldertypeSendsignatureattachmentezsignsigner(v bool)`

SetBEzsignfoldertypeSendsignatureattachmentezsignsigner sets BEzsignfoldertypeSendsignatureattachmentezsignsigner field to given value.

### HasBEzsignfoldertypeSendsignatureattachmentezsignsigner

`func (o *EzsignfoldertypeRequestV4) HasBEzsignfoldertypeSendsignatureattachmentezsignsigner() bool`

HasBEzsignfoldertypeSendsignatureattachmentezsignsigner returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendsignatureattachment

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeSendsignatureattachment() bool`

GetBEzsignfoldertypeSendsignatureattachment returns the BEzsignfoldertypeSendsignatureattachment field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsignatureattachmentOk

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeSendsignatureattachmentOk() (*bool, bool)`

GetBEzsignfoldertypeSendsignatureattachmentOk returns a tuple with the BEzsignfoldertypeSendsignatureattachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsignatureattachment

`func (o *EzsignfoldertypeRequestV4) SetBEzsignfoldertypeSendsignatureattachment(v bool)`

SetBEzsignfoldertypeSendsignatureattachment sets BEzsignfoldertypeSendsignatureattachment field to given value.

### HasBEzsignfoldertypeSendsignatureattachment

`func (o *EzsignfoldertypeRequestV4) HasBEzsignfoldertypeSendsignatureattachment() bool`

HasBEzsignfoldertypeSendsignatureattachment returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendproofezsignsigner

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeSendproofezsignsigner() bool`

GetBEzsignfoldertypeSendproofezsignsigner returns the BEzsignfoldertypeSendproofezsignsigner field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendproofezsignsignerOk

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeSendproofezsignsignerOk() (*bool, bool)`

GetBEzsignfoldertypeSendproofezsignsignerOk returns a tuple with the BEzsignfoldertypeSendproofezsignsigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendproofezsignsigner

`func (o *EzsignfoldertypeRequestV4) SetBEzsignfoldertypeSendproofezsignsigner(v bool)`

SetBEzsignfoldertypeSendproofezsignsigner sets BEzsignfoldertypeSendproofezsignsigner field to given value.

### HasBEzsignfoldertypeSendproofezsignsigner

`func (o *EzsignfoldertypeRequestV4) HasBEzsignfoldertypeSendproofezsignsigner() bool`

HasBEzsignfoldertypeSendproofezsignsigner returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendattachmentuser

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeSendattachmentuser() bool`

GetBEzsignfoldertypeSendattachmentuser returns the BEzsignfoldertypeSendattachmentuser field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendattachmentuserOk

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeSendattachmentuserOk() (*bool, bool)`

GetBEzsignfoldertypeSendattachmentuserOk returns a tuple with the BEzsignfoldertypeSendattachmentuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendattachmentuser

`func (o *EzsignfoldertypeRequestV4) SetBEzsignfoldertypeSendattachmentuser(v bool)`

SetBEzsignfoldertypeSendattachmentuser sets BEzsignfoldertypeSendattachmentuser field to given value.

### HasBEzsignfoldertypeSendattachmentuser

`func (o *EzsignfoldertypeRequestV4) HasBEzsignfoldertypeSendattachmentuser() bool`

HasBEzsignfoldertypeSendattachmentuser returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendsignatureattachmentuser

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeSendsignatureattachmentuser() bool`

GetBEzsignfoldertypeSendsignatureattachmentuser returns the BEzsignfoldertypeSendsignatureattachmentuser field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsignatureattachmentuserOk

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeSendsignatureattachmentuserOk() (*bool, bool)`

GetBEzsignfoldertypeSendsignatureattachmentuserOk returns a tuple with the BEzsignfoldertypeSendsignatureattachmentuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsignatureattachmentuser

`func (o *EzsignfoldertypeRequestV4) SetBEzsignfoldertypeSendsignatureattachmentuser(v bool)`

SetBEzsignfoldertypeSendsignatureattachmentuser sets BEzsignfoldertypeSendsignatureattachmentuser field to given value.

### HasBEzsignfoldertypeSendsignatureattachmentuser

`func (o *EzsignfoldertypeRequestV4) HasBEzsignfoldertypeSendsignatureattachmentuser() bool`

HasBEzsignfoldertypeSendsignatureattachmentuser returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendproofuser

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeSendproofuser() bool`

GetBEzsignfoldertypeSendproofuser returns the BEzsignfoldertypeSendproofuser field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendproofuserOk

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeSendproofuserOk() (*bool, bool)`

GetBEzsignfoldertypeSendproofuserOk returns a tuple with the BEzsignfoldertypeSendproofuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendproofuser

`func (o *EzsignfoldertypeRequestV4) SetBEzsignfoldertypeSendproofuser(v bool)`

SetBEzsignfoldertypeSendproofuser sets BEzsignfoldertypeSendproofuser field to given value.

### HasBEzsignfoldertypeSendproofuser

`func (o *EzsignfoldertypeRequestV4) HasBEzsignfoldertypeSendproofuser() bool`

HasBEzsignfoldertypeSendproofuser returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendproofemail

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeSendproofemail() bool`

GetBEzsignfoldertypeSendproofemail returns the BEzsignfoldertypeSendproofemail field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendproofemailOk

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeSendproofemailOk() (*bool, bool)`

GetBEzsignfoldertypeSendproofemailOk returns a tuple with the BEzsignfoldertypeSendproofemail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendproofemail

`func (o *EzsignfoldertypeRequestV4) SetBEzsignfoldertypeSendproofemail(v bool)`

SetBEzsignfoldertypeSendproofemail sets BEzsignfoldertypeSendproofemail field to given value.

### HasBEzsignfoldertypeSendproofemail

`func (o *EzsignfoldertypeRequestV4) HasBEzsignfoldertypeSendproofemail() bool`

HasBEzsignfoldertypeSendproofemail returns a boolean if a field has been set.

### GetBEzsignfoldertypeAllowdownloadattachmentezsignsigner

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeAllowdownloadattachmentezsignsigner() bool`

GetBEzsignfoldertypeAllowdownloadattachmentezsignsigner returns the BEzsignfoldertypeAllowdownloadattachmentezsignsigner field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeAllowdownloadattachmentezsignsignerOk

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeAllowdownloadattachmentezsignsignerOk() (*bool, bool)`

GetBEzsignfoldertypeAllowdownloadattachmentezsignsignerOk returns a tuple with the BEzsignfoldertypeAllowdownloadattachmentezsignsigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeAllowdownloadattachmentezsignsigner

`func (o *EzsignfoldertypeRequestV4) SetBEzsignfoldertypeAllowdownloadattachmentezsignsigner(v bool)`

SetBEzsignfoldertypeAllowdownloadattachmentezsignsigner sets BEzsignfoldertypeAllowdownloadattachmentezsignsigner field to given value.

### HasBEzsignfoldertypeAllowdownloadattachmentezsignsigner

`func (o *EzsignfoldertypeRequestV4) HasBEzsignfoldertypeAllowdownloadattachmentezsignsigner() bool`

HasBEzsignfoldertypeAllowdownloadattachmentezsignsigner returns a boolean if a field has been set.

### GetBEzsignfoldertypeAllowdownloadsignatureattachmentezsignsigner

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeAllowdownloadsignatureattachmentezsignsigner() bool`

GetBEzsignfoldertypeAllowdownloadsignatureattachmentezsignsigner returns the BEzsignfoldertypeAllowdownloadsignatureattachmentezsignsigner field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeAllowdownloadsignatureattachmentezsignsignerOk

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeAllowdownloadsignatureattachmentezsignsignerOk() (*bool, bool)`

GetBEzsignfoldertypeAllowdownloadsignatureattachmentezsignsignerOk returns a tuple with the BEzsignfoldertypeAllowdownloadsignatureattachmentezsignsigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeAllowdownloadsignatureattachmentezsignsigner

`func (o *EzsignfoldertypeRequestV4) SetBEzsignfoldertypeAllowdownloadsignatureattachmentezsignsigner(v bool)`

SetBEzsignfoldertypeAllowdownloadsignatureattachmentezsignsigner sets BEzsignfoldertypeAllowdownloadsignatureattachmentezsignsigner field to given value.

### HasBEzsignfoldertypeAllowdownloadsignatureattachmentezsignsigner

`func (o *EzsignfoldertypeRequestV4) HasBEzsignfoldertypeAllowdownloadsignatureattachmentezsignsigner() bool`

HasBEzsignfoldertypeAllowdownloadsignatureattachmentezsignsigner returns a boolean if a field has been set.

### GetBEzsignfoldertypeAllowdownloadproofezsignsigner

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeAllowdownloadproofezsignsigner() bool`

GetBEzsignfoldertypeAllowdownloadproofezsignsigner returns the BEzsignfoldertypeAllowdownloadproofezsignsigner field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeAllowdownloadproofezsignsignerOk

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeAllowdownloadproofezsignsignerOk() (*bool, bool)`

GetBEzsignfoldertypeAllowdownloadproofezsignsignerOk returns a tuple with the BEzsignfoldertypeAllowdownloadproofezsignsigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeAllowdownloadproofezsignsigner

`func (o *EzsignfoldertypeRequestV4) SetBEzsignfoldertypeAllowdownloadproofezsignsigner(v bool)`

SetBEzsignfoldertypeAllowdownloadproofezsignsigner sets BEzsignfoldertypeAllowdownloadproofezsignsigner field to given value.

### HasBEzsignfoldertypeAllowdownloadproofezsignsigner

`func (o *EzsignfoldertypeRequestV4) HasBEzsignfoldertypeAllowdownloadproofezsignsigner() bool`

HasBEzsignfoldertypeAllowdownloadproofezsignsigner returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendproofreceivealldocument

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeSendproofreceivealldocument() bool`

GetBEzsignfoldertypeSendproofreceivealldocument returns the BEzsignfoldertypeSendproofreceivealldocument field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendproofreceivealldocumentOk

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeSendproofreceivealldocumentOk() (*bool, bool)`

GetBEzsignfoldertypeSendproofreceivealldocumentOk returns a tuple with the BEzsignfoldertypeSendproofreceivealldocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendproofreceivealldocument

`func (o *EzsignfoldertypeRequestV4) SetBEzsignfoldertypeSendproofreceivealldocument(v bool)`

SetBEzsignfoldertypeSendproofreceivealldocument sets BEzsignfoldertypeSendproofreceivealldocument field to given value.

### HasBEzsignfoldertypeSendproofreceivealldocument

`func (o *EzsignfoldertypeRequestV4) HasBEzsignfoldertypeSendproofreceivealldocument() bool`

HasBEzsignfoldertypeSendproofreceivealldocument returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendsignatureattachmentreceivealldocument

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeSendsignatureattachmentreceivealldocument() bool`

GetBEzsignfoldertypeSendsignatureattachmentreceivealldocument returns the BEzsignfoldertypeSendsignatureattachmentreceivealldocument field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsignatureattachmentreceivealldocumentOk

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeSendsignatureattachmentreceivealldocumentOk() (*bool, bool)`

GetBEzsignfoldertypeSendsignatureattachmentreceivealldocumentOk returns a tuple with the BEzsignfoldertypeSendsignatureattachmentreceivealldocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsignatureattachmentreceivealldocument

`func (o *EzsignfoldertypeRequestV4) SetBEzsignfoldertypeSendsignatureattachmentreceivealldocument(v bool)`

SetBEzsignfoldertypeSendsignatureattachmentreceivealldocument sets BEzsignfoldertypeSendsignatureattachmentreceivealldocument field to given value.

### HasBEzsignfoldertypeSendsignatureattachmentreceivealldocument

`func (o *EzsignfoldertypeRequestV4) HasBEzsignfoldertypeSendsignatureattachmentreceivealldocument() bool`

HasBEzsignfoldertypeSendsignatureattachmentreceivealldocument returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendsignedtodocumentowner

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeSendsignedtodocumentowner() bool`

GetBEzsignfoldertypeSendsignedtodocumentowner returns the BEzsignfoldertypeSendsignedtodocumentowner field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsignedtodocumentownerOk

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeSendsignedtodocumentownerOk() (*bool, bool)`

GetBEzsignfoldertypeSendsignedtodocumentownerOk returns a tuple with the BEzsignfoldertypeSendsignedtodocumentowner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsignedtodocumentowner

`func (o *EzsignfoldertypeRequestV4) SetBEzsignfoldertypeSendsignedtodocumentowner(v bool)`

SetBEzsignfoldertypeSendsignedtodocumentowner sets BEzsignfoldertypeSendsignedtodocumentowner field to given value.


### GetBEzsignfoldertypeSendsignedtofolderowner

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeSendsignedtofolderowner() bool`

GetBEzsignfoldertypeSendsignedtofolderowner returns the BEzsignfoldertypeSendsignedtofolderowner field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsignedtofolderownerOk

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeSendsignedtofolderownerOk() (*bool, bool)`

GetBEzsignfoldertypeSendsignedtofolderownerOk returns a tuple with the BEzsignfoldertypeSendsignedtofolderowner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsignedtofolderowner

`func (o *EzsignfoldertypeRequestV4) SetBEzsignfoldertypeSendsignedtofolderowner(v bool)`

SetBEzsignfoldertypeSendsignedtofolderowner sets BEzsignfoldertypeSendsignedtofolderowner field to given value.


### GetBEzsignfoldertypeSendsignedtofullgroup

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeSendsignedtofullgroup() bool`

GetBEzsignfoldertypeSendsignedtofullgroup returns the BEzsignfoldertypeSendsignedtofullgroup field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsignedtofullgroupOk

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeSendsignedtofullgroupOk() (*bool, bool)`

GetBEzsignfoldertypeSendsignedtofullgroupOk returns a tuple with the BEzsignfoldertypeSendsignedtofullgroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsignedtofullgroup

`func (o *EzsignfoldertypeRequestV4) SetBEzsignfoldertypeSendsignedtofullgroup(v bool)`

SetBEzsignfoldertypeSendsignedtofullgroup sets BEzsignfoldertypeSendsignedtofullgroup field to given value.

### HasBEzsignfoldertypeSendsignedtofullgroup

`func (o *EzsignfoldertypeRequestV4) HasBEzsignfoldertypeSendsignedtofullgroup() bool`

HasBEzsignfoldertypeSendsignedtofullgroup returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendsignedtolimitedgroup

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeSendsignedtolimitedgroup() bool`

GetBEzsignfoldertypeSendsignedtolimitedgroup returns the BEzsignfoldertypeSendsignedtolimitedgroup field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsignedtolimitedgroupOk

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeSendsignedtolimitedgroupOk() (*bool, bool)`

GetBEzsignfoldertypeSendsignedtolimitedgroupOk returns a tuple with the BEzsignfoldertypeSendsignedtolimitedgroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsignedtolimitedgroup

`func (o *EzsignfoldertypeRequestV4) SetBEzsignfoldertypeSendsignedtolimitedgroup(v bool)`

SetBEzsignfoldertypeSendsignedtolimitedgroup sets BEzsignfoldertypeSendsignedtolimitedgroup field to given value.

### HasBEzsignfoldertypeSendsignedtolimitedgroup

`func (o *EzsignfoldertypeRequestV4) HasBEzsignfoldertypeSendsignedtolimitedgroup() bool`

HasBEzsignfoldertypeSendsignedtolimitedgroup returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendsignedtocolleague

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeSendsignedtocolleague() bool`

GetBEzsignfoldertypeSendsignedtocolleague returns the BEzsignfoldertypeSendsignedtocolleague field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsignedtocolleagueOk

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeSendsignedtocolleagueOk() (*bool, bool)`

GetBEzsignfoldertypeSendsignedtocolleagueOk returns a tuple with the BEzsignfoldertypeSendsignedtocolleague field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsignedtocolleague

`func (o *EzsignfoldertypeRequestV4) SetBEzsignfoldertypeSendsignedtocolleague(v bool)`

SetBEzsignfoldertypeSendsignedtocolleague sets BEzsignfoldertypeSendsignedtocolleague field to given value.


### GetBEzsignfoldertypeSendsummarytodocumentowner

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeSendsummarytodocumentowner() bool`

GetBEzsignfoldertypeSendsummarytodocumentowner returns the BEzsignfoldertypeSendsummarytodocumentowner field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsummarytodocumentownerOk

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeSendsummarytodocumentownerOk() (*bool, bool)`

GetBEzsignfoldertypeSendsummarytodocumentownerOk returns a tuple with the BEzsignfoldertypeSendsummarytodocumentowner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsummarytodocumentowner

`func (o *EzsignfoldertypeRequestV4) SetBEzsignfoldertypeSendsummarytodocumentowner(v bool)`

SetBEzsignfoldertypeSendsummarytodocumentowner sets BEzsignfoldertypeSendsummarytodocumentowner field to given value.


### GetBEzsignfoldertypeSendsummarytofolderowner

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeSendsummarytofolderowner() bool`

GetBEzsignfoldertypeSendsummarytofolderowner returns the BEzsignfoldertypeSendsummarytofolderowner field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsummarytofolderownerOk

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeSendsummarytofolderownerOk() (*bool, bool)`

GetBEzsignfoldertypeSendsummarytofolderownerOk returns a tuple with the BEzsignfoldertypeSendsummarytofolderowner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsummarytofolderowner

`func (o *EzsignfoldertypeRequestV4) SetBEzsignfoldertypeSendsummarytofolderowner(v bool)`

SetBEzsignfoldertypeSendsummarytofolderowner sets BEzsignfoldertypeSendsummarytofolderowner field to given value.


### GetBEzsignfoldertypeSendsummarytofullgroup

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeSendsummarytofullgroup() bool`

GetBEzsignfoldertypeSendsummarytofullgroup returns the BEzsignfoldertypeSendsummarytofullgroup field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsummarytofullgroupOk

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeSendsummarytofullgroupOk() (*bool, bool)`

GetBEzsignfoldertypeSendsummarytofullgroupOk returns a tuple with the BEzsignfoldertypeSendsummarytofullgroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsummarytofullgroup

`func (o *EzsignfoldertypeRequestV4) SetBEzsignfoldertypeSendsummarytofullgroup(v bool)`

SetBEzsignfoldertypeSendsummarytofullgroup sets BEzsignfoldertypeSendsummarytofullgroup field to given value.

### HasBEzsignfoldertypeSendsummarytofullgroup

`func (o *EzsignfoldertypeRequestV4) HasBEzsignfoldertypeSendsummarytofullgroup() bool`

HasBEzsignfoldertypeSendsummarytofullgroup returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendsummarytolimitedgroup

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeSendsummarytolimitedgroup() bool`

GetBEzsignfoldertypeSendsummarytolimitedgroup returns the BEzsignfoldertypeSendsummarytolimitedgroup field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsummarytolimitedgroupOk

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeSendsummarytolimitedgroupOk() (*bool, bool)`

GetBEzsignfoldertypeSendsummarytolimitedgroupOk returns a tuple with the BEzsignfoldertypeSendsummarytolimitedgroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsummarytolimitedgroup

`func (o *EzsignfoldertypeRequestV4) SetBEzsignfoldertypeSendsummarytolimitedgroup(v bool)`

SetBEzsignfoldertypeSendsummarytolimitedgroup sets BEzsignfoldertypeSendsummarytolimitedgroup field to given value.

### HasBEzsignfoldertypeSendsummarytolimitedgroup

`func (o *EzsignfoldertypeRequestV4) HasBEzsignfoldertypeSendsummarytolimitedgroup() bool`

HasBEzsignfoldertypeSendsummarytolimitedgroup returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendsummarytocolleague

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeSendsummarytocolleague() bool`

GetBEzsignfoldertypeSendsummarytocolleague returns the BEzsignfoldertypeSendsummarytocolleague field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendsummarytocolleagueOk

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeSendsummarytocolleagueOk() (*bool, bool)`

GetBEzsignfoldertypeSendsummarytocolleagueOk returns a tuple with the BEzsignfoldertypeSendsummarytocolleague field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendsummarytocolleague

`func (o *EzsignfoldertypeRequestV4) SetBEzsignfoldertypeSendsummarytocolleague(v bool)`

SetBEzsignfoldertypeSendsummarytocolleague sets BEzsignfoldertypeSendsummarytocolleague field to given value.


### GetEEzsignfoldertypeSigneraccess

`func (o *EzsignfoldertypeRequestV4) GetEEzsignfoldertypeSigneraccess() FieldEEzsignfoldertypeSigneraccess`

GetEEzsignfoldertypeSigneraccess returns the EEzsignfoldertypeSigneraccess field if non-nil, zero value otherwise.

### GetEEzsignfoldertypeSigneraccessOk

`func (o *EzsignfoldertypeRequestV4) GetEEzsignfoldertypeSigneraccessOk() (*FieldEEzsignfoldertypeSigneraccess, bool)`

GetEEzsignfoldertypeSigneraccessOk returns a tuple with the EEzsignfoldertypeSigneraccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignfoldertypeSigneraccess

`func (o *EzsignfoldertypeRequestV4) SetEEzsignfoldertypeSigneraccess(v FieldEEzsignfoldertypeSigneraccess)`

SetEEzsignfoldertypeSigneraccess sets EEzsignfoldertypeSigneraccess field to given value.

### HasEEzsignfoldertypeSigneraccess

`func (o *EzsignfoldertypeRequestV4) HasEEzsignfoldertypeSigneraccess() bool`

HasEEzsignfoldertypeSigneraccess returns a boolean if a field has been set.

### GetBEzsignfoldertypeIsactive

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeIsactive() bool`

GetBEzsignfoldertypeIsactive returns the BEzsignfoldertypeIsactive field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeIsactiveOk

`func (o *EzsignfoldertypeRequestV4) GetBEzsignfoldertypeIsactiveOk() (*bool, bool)`

GetBEzsignfoldertypeIsactiveOk returns a tuple with the BEzsignfoldertypeIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeIsactive

`func (o *EzsignfoldertypeRequestV4) SetBEzsignfoldertypeIsactive(v bool)`

SetBEzsignfoldertypeIsactive sets BEzsignfoldertypeIsactive field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


