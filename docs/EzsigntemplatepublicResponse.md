# EzsigntemplatepublicResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsigntemplatepublicID** | **int32** | The unique ID of the Ezsigntemplatepublic | 
**FkiEzsignfoldertypeID** | **int32** | The unique ID of the Ezsignfoldertype. | 
**SEzsignfoldertypeNameX** | **string** | The name of the Ezsignfoldertype in the language of the requester | 
**FkiUserlogintypeID** | **int32** | The unique ID of the Userlogintype  Valid values:  |Value|Description|Detail| |-|-|-| |1|**Email Only**|The Ezsignsigner will receive a secure link by email| |2|**Email and phone or SMS**|The Ezsignsigner will receive a secure link by email and will need to authenticate using SMS or Phone call. **Additional fee applies**| |3|**Email and secret question**|The Ezsignsigner will receive a secure link by email and will need to authenticate using a predefined question and answer| |4|**In person only**|The Ezsignsigner will only be able to sign \&quot;In-Person\&quot; and there won&#39;t be any authentication. No email will be sent for invitation to sign. Make sure you evaluate the risk of signature denial and at minimum, we recommend you use a handwritten signature type| |5|**In person with phone or SMS**|The Ezsignsigner will only be able to sign \&quot;In-Person\&quot; and will need to authenticate using SMS or Phone call. No email will be sent for invitation to sign. **Additional fee applies**| |6|**Embedded**|The Ezsignsigner will only be able to sign in the embedded solution. No email will be sent for invitation to sign. **Additional fee applies**|   |7|**Embedded with phone or SMS**|The Ezsignsigner will only be able to sign in the embedded solution and will need to authenticate using SMS or Phone call. No email will be sent for invitation to sign. **Additional fee applies**|   |8|**No validation**|The Ezsignsigner will not receive an email and won&#39;t have to validate his connection using 2 factor. **Additional fee applies**|      |9|**Sms only**|The Ezsignsigner will not receive an email but will will need to authenticate using SMS. **Additional fee applies**|      | 
**SUserlogintypeDescriptionX** | **string** | The description of the Userlogintype in the language of the requester | 
**FkiEzsigntemplateID** | Pointer to **int32** | The unique ID of the Ezsigntemplate | [optional] 
**FkiEzsigntemplatepackageID** | Pointer to **int32** | The unique ID of the Ezsigntemplatepackage | [optional] 
**SEzsigntemplatepublicDescription** | **string** | The description of the Ezsigntemplatepublic | 
**SEzsigntemplatepublicReferenceid** | **string** | The referenceid of the Ezsigntemplatepublic | 
**BEzsigntemplatepublicIsactive** | **bool** | Whether the ezsigntemplatepublic is active or not | 
**TEzsigntemplatepublicNote** | **string** | The note of the Ezsigntemplatepublic | 
**EEzsigntemplatepublicLimittype** | [**FieldEEzsigntemplatepublicLimittype**](FieldEEzsigntemplatepublicLimittype.md) |  | 
**IEzsigntemplatepublicLimit** | **int32** | The limit of the Ezsigntemplatepublic | 
**IEzsigntemplatepublicLimitexceeded** | **int32** | The limitexceeded of the Ezsigntemplatepublic | 
**DtEzsigntemplatepublicLimitexceededsince** | **string** | The limitexceededsince of the Ezsigntemplatepublic | 
**SEzsigntemplatepublicUrl** | **string** | The url of the Ezsigntemplatepublic  You can add these value as query parameters to prefill the corresponding role  |Parameter|Description| |-|-| |sEzsigntemplatesignerDescription|The role to fill| |sContactFirstname|The contact firstname| |sContactLastname|The contact lastname| |sEmailAddress|The contact email| |sPhoneE164|The contact phone number| |sPhoneE164Cell|The contact cell phone number| | 
**SEzsigntemplatepublicEzsigntemplatedescription** | **string** | The Ezsigntemplate/Ezsigntemplatepackage description | 
**ObjAudit** | Pointer to [**CommonAudit**](CommonAudit.md) |  | [optional] 

## Methods

### NewEzsigntemplatepublicResponse

`func NewEzsigntemplatepublicResponse(pkiEzsigntemplatepublicID int32, fkiEzsignfoldertypeID int32, sEzsignfoldertypeNameX string, fkiUserlogintypeID int32, sUserlogintypeDescriptionX string, sEzsigntemplatepublicDescription string, sEzsigntemplatepublicReferenceid string, bEzsigntemplatepublicIsactive bool, tEzsigntemplatepublicNote string, eEzsigntemplatepublicLimittype FieldEEzsigntemplatepublicLimittype, iEzsigntemplatepublicLimit int32, iEzsigntemplatepublicLimitexceeded int32, dtEzsigntemplatepublicLimitexceededsince string, sEzsigntemplatepublicUrl string, sEzsigntemplatepublicEzsigntemplatedescription string, ) *EzsigntemplatepublicResponse`

NewEzsigntemplatepublicResponse instantiates a new EzsigntemplatepublicResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigntemplatepublicResponseWithDefaults

`func NewEzsigntemplatepublicResponseWithDefaults() *EzsigntemplatepublicResponse`

NewEzsigntemplatepublicResponseWithDefaults instantiates a new EzsigntemplatepublicResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsigntemplatepublicID

`func (o *EzsigntemplatepublicResponse) GetPkiEzsigntemplatepublicID() int32`

GetPkiEzsigntemplatepublicID returns the PkiEzsigntemplatepublicID field if non-nil, zero value otherwise.

### GetPkiEzsigntemplatepublicIDOk

`func (o *EzsigntemplatepublicResponse) GetPkiEzsigntemplatepublicIDOk() (*int32, bool)`

GetPkiEzsigntemplatepublicIDOk returns a tuple with the PkiEzsigntemplatepublicID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsigntemplatepublicID

`func (o *EzsigntemplatepublicResponse) SetPkiEzsigntemplatepublicID(v int32)`

SetPkiEzsigntemplatepublicID sets PkiEzsigntemplatepublicID field to given value.


### GetFkiEzsignfoldertypeID

`func (o *EzsigntemplatepublicResponse) GetFkiEzsignfoldertypeID() int32`

GetFkiEzsignfoldertypeID returns the FkiEzsignfoldertypeID field if non-nil, zero value otherwise.

### GetFkiEzsignfoldertypeIDOk

`func (o *EzsigntemplatepublicResponse) GetFkiEzsignfoldertypeIDOk() (*int32, bool)`

GetFkiEzsignfoldertypeIDOk returns a tuple with the FkiEzsignfoldertypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfoldertypeID

`func (o *EzsigntemplatepublicResponse) SetFkiEzsignfoldertypeID(v int32)`

SetFkiEzsignfoldertypeID sets FkiEzsignfoldertypeID field to given value.


### GetSEzsignfoldertypeNameX

`func (o *EzsigntemplatepublicResponse) GetSEzsignfoldertypeNameX() string`

GetSEzsignfoldertypeNameX returns the SEzsignfoldertypeNameX field if non-nil, zero value otherwise.

### GetSEzsignfoldertypeNameXOk

`func (o *EzsigntemplatepublicResponse) GetSEzsignfoldertypeNameXOk() (*string, bool)`

GetSEzsignfoldertypeNameXOk returns a tuple with the SEzsignfoldertypeNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignfoldertypeNameX

`func (o *EzsigntemplatepublicResponse) SetSEzsignfoldertypeNameX(v string)`

SetSEzsignfoldertypeNameX sets SEzsignfoldertypeNameX field to given value.


### GetFkiUserlogintypeID

`func (o *EzsigntemplatepublicResponse) GetFkiUserlogintypeID() int32`

GetFkiUserlogintypeID returns the FkiUserlogintypeID field if non-nil, zero value otherwise.

### GetFkiUserlogintypeIDOk

`func (o *EzsigntemplatepublicResponse) GetFkiUserlogintypeIDOk() (*int32, bool)`

GetFkiUserlogintypeIDOk returns a tuple with the FkiUserlogintypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserlogintypeID

`func (o *EzsigntemplatepublicResponse) SetFkiUserlogintypeID(v int32)`

SetFkiUserlogintypeID sets FkiUserlogintypeID field to given value.


### GetSUserlogintypeDescriptionX

`func (o *EzsigntemplatepublicResponse) GetSUserlogintypeDescriptionX() string`

GetSUserlogintypeDescriptionX returns the SUserlogintypeDescriptionX field if non-nil, zero value otherwise.

### GetSUserlogintypeDescriptionXOk

`func (o *EzsigntemplatepublicResponse) GetSUserlogintypeDescriptionXOk() (*string, bool)`

GetSUserlogintypeDescriptionXOk returns a tuple with the SUserlogintypeDescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserlogintypeDescriptionX

`func (o *EzsigntemplatepublicResponse) SetSUserlogintypeDescriptionX(v string)`

SetSUserlogintypeDescriptionX sets SUserlogintypeDescriptionX field to given value.


### GetFkiEzsigntemplateID

`func (o *EzsigntemplatepublicResponse) GetFkiEzsigntemplateID() int32`

GetFkiEzsigntemplateID returns the FkiEzsigntemplateID field if non-nil, zero value otherwise.

### GetFkiEzsigntemplateIDOk

`func (o *EzsigntemplatepublicResponse) GetFkiEzsigntemplateIDOk() (*int32, bool)`

GetFkiEzsigntemplateIDOk returns a tuple with the FkiEzsigntemplateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntemplateID

`func (o *EzsigntemplatepublicResponse) SetFkiEzsigntemplateID(v int32)`

SetFkiEzsigntemplateID sets FkiEzsigntemplateID field to given value.

### HasFkiEzsigntemplateID

`func (o *EzsigntemplatepublicResponse) HasFkiEzsigntemplateID() bool`

HasFkiEzsigntemplateID returns a boolean if a field has been set.

### GetFkiEzsigntemplatepackageID

`func (o *EzsigntemplatepublicResponse) GetFkiEzsigntemplatepackageID() int32`

GetFkiEzsigntemplatepackageID returns the FkiEzsigntemplatepackageID field if non-nil, zero value otherwise.

### GetFkiEzsigntemplatepackageIDOk

`func (o *EzsigntemplatepublicResponse) GetFkiEzsigntemplatepackageIDOk() (*int32, bool)`

GetFkiEzsigntemplatepackageIDOk returns a tuple with the FkiEzsigntemplatepackageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntemplatepackageID

`func (o *EzsigntemplatepublicResponse) SetFkiEzsigntemplatepackageID(v int32)`

SetFkiEzsigntemplatepackageID sets FkiEzsigntemplatepackageID field to given value.

### HasFkiEzsigntemplatepackageID

`func (o *EzsigntemplatepublicResponse) HasFkiEzsigntemplatepackageID() bool`

HasFkiEzsigntemplatepackageID returns a boolean if a field has been set.

### GetSEzsigntemplatepublicDescription

`func (o *EzsigntemplatepublicResponse) GetSEzsigntemplatepublicDescription() string`

GetSEzsigntemplatepublicDescription returns the SEzsigntemplatepublicDescription field if non-nil, zero value otherwise.

### GetSEzsigntemplatepublicDescriptionOk

`func (o *EzsigntemplatepublicResponse) GetSEzsigntemplatepublicDescriptionOk() (*string, bool)`

GetSEzsigntemplatepublicDescriptionOk returns a tuple with the SEzsigntemplatepublicDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplatepublicDescription

`func (o *EzsigntemplatepublicResponse) SetSEzsigntemplatepublicDescription(v string)`

SetSEzsigntemplatepublicDescription sets SEzsigntemplatepublicDescription field to given value.


### GetSEzsigntemplatepublicReferenceid

`func (o *EzsigntemplatepublicResponse) GetSEzsigntemplatepublicReferenceid() string`

GetSEzsigntemplatepublicReferenceid returns the SEzsigntemplatepublicReferenceid field if non-nil, zero value otherwise.

### GetSEzsigntemplatepublicReferenceidOk

`func (o *EzsigntemplatepublicResponse) GetSEzsigntemplatepublicReferenceidOk() (*string, bool)`

GetSEzsigntemplatepublicReferenceidOk returns a tuple with the SEzsigntemplatepublicReferenceid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplatepublicReferenceid

`func (o *EzsigntemplatepublicResponse) SetSEzsigntemplatepublicReferenceid(v string)`

SetSEzsigntemplatepublicReferenceid sets SEzsigntemplatepublicReferenceid field to given value.


### GetBEzsigntemplatepublicIsactive

`func (o *EzsigntemplatepublicResponse) GetBEzsigntemplatepublicIsactive() bool`

GetBEzsigntemplatepublicIsactive returns the BEzsigntemplatepublicIsactive field if non-nil, zero value otherwise.

### GetBEzsigntemplatepublicIsactiveOk

`func (o *EzsigntemplatepublicResponse) GetBEzsigntemplatepublicIsactiveOk() (*bool, bool)`

GetBEzsigntemplatepublicIsactiveOk returns a tuple with the BEzsigntemplatepublicIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigntemplatepublicIsactive

`func (o *EzsigntemplatepublicResponse) SetBEzsigntemplatepublicIsactive(v bool)`

SetBEzsigntemplatepublicIsactive sets BEzsigntemplatepublicIsactive field to given value.


### GetTEzsigntemplatepublicNote

`func (o *EzsigntemplatepublicResponse) GetTEzsigntemplatepublicNote() string`

GetTEzsigntemplatepublicNote returns the TEzsigntemplatepublicNote field if non-nil, zero value otherwise.

### GetTEzsigntemplatepublicNoteOk

`func (o *EzsigntemplatepublicResponse) GetTEzsigntemplatepublicNoteOk() (*string, bool)`

GetTEzsigntemplatepublicNoteOk returns a tuple with the TEzsigntemplatepublicNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEzsigntemplatepublicNote

`func (o *EzsigntemplatepublicResponse) SetTEzsigntemplatepublicNote(v string)`

SetTEzsigntemplatepublicNote sets TEzsigntemplatepublicNote field to given value.


### GetEEzsigntemplatepublicLimittype

`func (o *EzsigntemplatepublicResponse) GetEEzsigntemplatepublicLimittype() FieldEEzsigntemplatepublicLimittype`

GetEEzsigntemplatepublicLimittype returns the EEzsigntemplatepublicLimittype field if non-nil, zero value otherwise.

### GetEEzsigntemplatepublicLimittypeOk

`func (o *EzsigntemplatepublicResponse) GetEEzsigntemplatepublicLimittypeOk() (*FieldEEzsigntemplatepublicLimittype, bool)`

GetEEzsigntemplatepublicLimittypeOk returns a tuple with the EEzsigntemplatepublicLimittype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplatepublicLimittype

`func (o *EzsigntemplatepublicResponse) SetEEzsigntemplatepublicLimittype(v FieldEEzsigntemplatepublicLimittype)`

SetEEzsigntemplatepublicLimittype sets EEzsigntemplatepublicLimittype field to given value.


### GetIEzsigntemplatepublicLimit

`func (o *EzsigntemplatepublicResponse) GetIEzsigntemplatepublicLimit() int32`

GetIEzsigntemplatepublicLimit returns the IEzsigntemplatepublicLimit field if non-nil, zero value otherwise.

### GetIEzsigntemplatepublicLimitOk

`func (o *EzsigntemplatepublicResponse) GetIEzsigntemplatepublicLimitOk() (*int32, bool)`

GetIEzsigntemplatepublicLimitOk returns a tuple with the IEzsigntemplatepublicLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplatepublicLimit

`func (o *EzsigntemplatepublicResponse) SetIEzsigntemplatepublicLimit(v int32)`

SetIEzsigntemplatepublicLimit sets IEzsigntemplatepublicLimit field to given value.


### GetIEzsigntemplatepublicLimitexceeded

`func (o *EzsigntemplatepublicResponse) GetIEzsigntemplatepublicLimitexceeded() int32`

GetIEzsigntemplatepublicLimitexceeded returns the IEzsigntemplatepublicLimitexceeded field if non-nil, zero value otherwise.

### GetIEzsigntemplatepublicLimitexceededOk

`func (o *EzsigntemplatepublicResponse) GetIEzsigntemplatepublicLimitexceededOk() (*int32, bool)`

GetIEzsigntemplatepublicLimitexceededOk returns a tuple with the IEzsigntemplatepublicLimitexceeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplatepublicLimitexceeded

`func (o *EzsigntemplatepublicResponse) SetIEzsigntemplatepublicLimitexceeded(v int32)`

SetIEzsigntemplatepublicLimitexceeded sets IEzsigntemplatepublicLimitexceeded field to given value.


### GetDtEzsigntemplatepublicLimitexceededsince

`func (o *EzsigntemplatepublicResponse) GetDtEzsigntemplatepublicLimitexceededsince() string`

GetDtEzsigntemplatepublicLimitexceededsince returns the DtEzsigntemplatepublicLimitexceededsince field if non-nil, zero value otherwise.

### GetDtEzsigntemplatepublicLimitexceededsinceOk

`func (o *EzsigntemplatepublicResponse) GetDtEzsigntemplatepublicLimitexceededsinceOk() (*string, bool)`

GetDtEzsigntemplatepublicLimitexceededsinceOk returns a tuple with the DtEzsigntemplatepublicLimitexceededsince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtEzsigntemplatepublicLimitexceededsince

`func (o *EzsigntemplatepublicResponse) SetDtEzsigntemplatepublicLimitexceededsince(v string)`

SetDtEzsigntemplatepublicLimitexceededsince sets DtEzsigntemplatepublicLimitexceededsince field to given value.


### GetSEzsigntemplatepublicUrl

`func (o *EzsigntemplatepublicResponse) GetSEzsigntemplatepublicUrl() string`

GetSEzsigntemplatepublicUrl returns the SEzsigntemplatepublicUrl field if non-nil, zero value otherwise.

### GetSEzsigntemplatepublicUrlOk

`func (o *EzsigntemplatepublicResponse) GetSEzsigntemplatepublicUrlOk() (*string, bool)`

GetSEzsigntemplatepublicUrlOk returns a tuple with the SEzsigntemplatepublicUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplatepublicUrl

`func (o *EzsigntemplatepublicResponse) SetSEzsigntemplatepublicUrl(v string)`

SetSEzsigntemplatepublicUrl sets SEzsigntemplatepublicUrl field to given value.


### GetSEzsigntemplatepublicEzsigntemplatedescription

`func (o *EzsigntemplatepublicResponse) GetSEzsigntemplatepublicEzsigntemplatedescription() string`

GetSEzsigntemplatepublicEzsigntemplatedescription returns the SEzsigntemplatepublicEzsigntemplatedescription field if non-nil, zero value otherwise.

### GetSEzsigntemplatepublicEzsigntemplatedescriptionOk

`func (o *EzsigntemplatepublicResponse) GetSEzsigntemplatepublicEzsigntemplatedescriptionOk() (*string, bool)`

GetSEzsigntemplatepublicEzsigntemplatedescriptionOk returns a tuple with the SEzsigntemplatepublicEzsigntemplatedescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplatepublicEzsigntemplatedescription

`func (o *EzsigntemplatepublicResponse) SetSEzsigntemplatepublicEzsigntemplatedescription(v string)`

SetSEzsigntemplatepublicEzsigntemplatedescription sets SEzsigntemplatepublicEzsigntemplatedescription field to given value.


### GetObjAudit

`func (o *EzsigntemplatepublicResponse) GetObjAudit() CommonAudit`

GetObjAudit returns the ObjAudit field if non-nil, zero value otherwise.

### GetObjAuditOk

`func (o *EzsigntemplatepublicResponse) GetObjAuditOk() (*CommonAudit, bool)`

GetObjAuditOk returns a tuple with the ObjAudit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjAudit

`func (o *EzsigntemplatepublicResponse) SetObjAudit(v CommonAudit)`

SetObjAudit sets ObjAudit field to given value.

### HasObjAudit

`func (o *EzsigntemplatepublicResponse) HasObjAudit() bool`

HasObjAudit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


