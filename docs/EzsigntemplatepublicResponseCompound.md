# EzsigntemplatepublicResponseCompound

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
**AObjEzsignfolderezsigntemplatepublic** | [**[]CustomEzsignfolderezsigntemplatepublicResponse**](CustomEzsignfolderezsigntemplatepublicResponse.md) |  | 

## Methods

### NewEzsigntemplatepublicResponseCompound

`func NewEzsigntemplatepublicResponseCompound(pkiEzsigntemplatepublicID int32, fkiEzsignfoldertypeID int32, sEzsignfoldertypeNameX string, fkiUserlogintypeID int32, sUserlogintypeDescriptionX string, sEzsigntemplatepublicDescription string, sEzsigntemplatepublicReferenceid string, bEzsigntemplatepublicIsactive bool, tEzsigntemplatepublicNote string, eEzsigntemplatepublicLimittype FieldEEzsigntemplatepublicLimittype, iEzsigntemplatepublicLimit int32, iEzsigntemplatepublicLimitexceeded int32, dtEzsigntemplatepublicLimitexceededsince string, sEzsigntemplatepublicUrl string, sEzsigntemplatepublicEzsigntemplatedescription string, aObjEzsignfolderezsigntemplatepublic []CustomEzsignfolderezsigntemplatepublicResponse, ) *EzsigntemplatepublicResponseCompound`

NewEzsigntemplatepublicResponseCompound instantiates a new EzsigntemplatepublicResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigntemplatepublicResponseCompoundWithDefaults

`func NewEzsigntemplatepublicResponseCompoundWithDefaults() *EzsigntemplatepublicResponseCompound`

NewEzsigntemplatepublicResponseCompoundWithDefaults instantiates a new EzsigntemplatepublicResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsigntemplatepublicID

`func (o *EzsigntemplatepublicResponseCompound) GetPkiEzsigntemplatepublicID() int32`

GetPkiEzsigntemplatepublicID returns the PkiEzsigntemplatepublicID field if non-nil, zero value otherwise.

### GetPkiEzsigntemplatepublicIDOk

`func (o *EzsigntemplatepublicResponseCompound) GetPkiEzsigntemplatepublicIDOk() (*int32, bool)`

GetPkiEzsigntemplatepublicIDOk returns a tuple with the PkiEzsigntemplatepublicID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsigntemplatepublicID

`func (o *EzsigntemplatepublicResponseCompound) SetPkiEzsigntemplatepublicID(v int32)`

SetPkiEzsigntemplatepublicID sets PkiEzsigntemplatepublicID field to given value.


### GetFkiEzsignfoldertypeID

`func (o *EzsigntemplatepublicResponseCompound) GetFkiEzsignfoldertypeID() int32`

GetFkiEzsignfoldertypeID returns the FkiEzsignfoldertypeID field if non-nil, zero value otherwise.

### GetFkiEzsignfoldertypeIDOk

`func (o *EzsigntemplatepublicResponseCompound) GetFkiEzsignfoldertypeIDOk() (*int32, bool)`

GetFkiEzsignfoldertypeIDOk returns a tuple with the FkiEzsignfoldertypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfoldertypeID

`func (o *EzsigntemplatepublicResponseCompound) SetFkiEzsignfoldertypeID(v int32)`

SetFkiEzsignfoldertypeID sets FkiEzsignfoldertypeID field to given value.


### GetSEzsignfoldertypeNameX

`func (o *EzsigntemplatepublicResponseCompound) GetSEzsignfoldertypeNameX() string`

GetSEzsignfoldertypeNameX returns the SEzsignfoldertypeNameX field if non-nil, zero value otherwise.

### GetSEzsignfoldertypeNameXOk

`func (o *EzsigntemplatepublicResponseCompound) GetSEzsignfoldertypeNameXOk() (*string, bool)`

GetSEzsignfoldertypeNameXOk returns a tuple with the SEzsignfoldertypeNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignfoldertypeNameX

`func (o *EzsigntemplatepublicResponseCompound) SetSEzsignfoldertypeNameX(v string)`

SetSEzsignfoldertypeNameX sets SEzsignfoldertypeNameX field to given value.


### GetFkiUserlogintypeID

`func (o *EzsigntemplatepublicResponseCompound) GetFkiUserlogintypeID() int32`

GetFkiUserlogintypeID returns the FkiUserlogintypeID field if non-nil, zero value otherwise.

### GetFkiUserlogintypeIDOk

`func (o *EzsigntemplatepublicResponseCompound) GetFkiUserlogintypeIDOk() (*int32, bool)`

GetFkiUserlogintypeIDOk returns a tuple with the FkiUserlogintypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserlogintypeID

`func (o *EzsigntemplatepublicResponseCompound) SetFkiUserlogintypeID(v int32)`

SetFkiUserlogintypeID sets FkiUserlogintypeID field to given value.


### GetSUserlogintypeDescriptionX

`func (o *EzsigntemplatepublicResponseCompound) GetSUserlogintypeDescriptionX() string`

GetSUserlogintypeDescriptionX returns the SUserlogintypeDescriptionX field if non-nil, zero value otherwise.

### GetSUserlogintypeDescriptionXOk

`func (o *EzsigntemplatepublicResponseCompound) GetSUserlogintypeDescriptionXOk() (*string, bool)`

GetSUserlogintypeDescriptionXOk returns a tuple with the SUserlogintypeDescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserlogintypeDescriptionX

`func (o *EzsigntemplatepublicResponseCompound) SetSUserlogintypeDescriptionX(v string)`

SetSUserlogintypeDescriptionX sets SUserlogintypeDescriptionX field to given value.


### GetFkiEzsigntemplateID

`func (o *EzsigntemplatepublicResponseCompound) GetFkiEzsigntemplateID() int32`

GetFkiEzsigntemplateID returns the FkiEzsigntemplateID field if non-nil, zero value otherwise.

### GetFkiEzsigntemplateIDOk

`func (o *EzsigntemplatepublicResponseCompound) GetFkiEzsigntemplateIDOk() (*int32, bool)`

GetFkiEzsigntemplateIDOk returns a tuple with the FkiEzsigntemplateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntemplateID

`func (o *EzsigntemplatepublicResponseCompound) SetFkiEzsigntemplateID(v int32)`

SetFkiEzsigntemplateID sets FkiEzsigntemplateID field to given value.

### HasFkiEzsigntemplateID

`func (o *EzsigntemplatepublicResponseCompound) HasFkiEzsigntemplateID() bool`

HasFkiEzsigntemplateID returns a boolean if a field has been set.

### GetFkiEzsigntemplatepackageID

`func (o *EzsigntemplatepublicResponseCompound) GetFkiEzsigntemplatepackageID() int32`

GetFkiEzsigntemplatepackageID returns the FkiEzsigntemplatepackageID field if non-nil, zero value otherwise.

### GetFkiEzsigntemplatepackageIDOk

`func (o *EzsigntemplatepublicResponseCompound) GetFkiEzsigntemplatepackageIDOk() (*int32, bool)`

GetFkiEzsigntemplatepackageIDOk returns a tuple with the FkiEzsigntemplatepackageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntemplatepackageID

`func (o *EzsigntemplatepublicResponseCompound) SetFkiEzsigntemplatepackageID(v int32)`

SetFkiEzsigntemplatepackageID sets FkiEzsigntemplatepackageID field to given value.

### HasFkiEzsigntemplatepackageID

`func (o *EzsigntemplatepublicResponseCompound) HasFkiEzsigntemplatepackageID() bool`

HasFkiEzsigntemplatepackageID returns a boolean if a field has been set.

### GetSEzsigntemplatepublicDescription

`func (o *EzsigntemplatepublicResponseCompound) GetSEzsigntemplatepublicDescription() string`

GetSEzsigntemplatepublicDescription returns the SEzsigntemplatepublicDescription field if non-nil, zero value otherwise.

### GetSEzsigntemplatepublicDescriptionOk

`func (o *EzsigntemplatepublicResponseCompound) GetSEzsigntemplatepublicDescriptionOk() (*string, bool)`

GetSEzsigntemplatepublicDescriptionOk returns a tuple with the SEzsigntemplatepublicDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplatepublicDescription

`func (o *EzsigntemplatepublicResponseCompound) SetSEzsigntemplatepublicDescription(v string)`

SetSEzsigntemplatepublicDescription sets SEzsigntemplatepublicDescription field to given value.


### GetSEzsigntemplatepublicReferenceid

`func (o *EzsigntemplatepublicResponseCompound) GetSEzsigntemplatepublicReferenceid() string`

GetSEzsigntemplatepublicReferenceid returns the SEzsigntemplatepublicReferenceid field if non-nil, zero value otherwise.

### GetSEzsigntemplatepublicReferenceidOk

`func (o *EzsigntemplatepublicResponseCompound) GetSEzsigntemplatepublicReferenceidOk() (*string, bool)`

GetSEzsigntemplatepublicReferenceidOk returns a tuple with the SEzsigntemplatepublicReferenceid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplatepublicReferenceid

`func (o *EzsigntemplatepublicResponseCompound) SetSEzsigntemplatepublicReferenceid(v string)`

SetSEzsigntemplatepublicReferenceid sets SEzsigntemplatepublicReferenceid field to given value.


### GetBEzsigntemplatepublicIsactive

`func (o *EzsigntemplatepublicResponseCompound) GetBEzsigntemplatepublicIsactive() bool`

GetBEzsigntemplatepublicIsactive returns the BEzsigntemplatepublicIsactive field if non-nil, zero value otherwise.

### GetBEzsigntemplatepublicIsactiveOk

`func (o *EzsigntemplatepublicResponseCompound) GetBEzsigntemplatepublicIsactiveOk() (*bool, bool)`

GetBEzsigntemplatepublicIsactiveOk returns a tuple with the BEzsigntemplatepublicIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigntemplatepublicIsactive

`func (o *EzsigntemplatepublicResponseCompound) SetBEzsigntemplatepublicIsactive(v bool)`

SetBEzsigntemplatepublicIsactive sets BEzsigntemplatepublicIsactive field to given value.


### GetTEzsigntemplatepublicNote

`func (o *EzsigntemplatepublicResponseCompound) GetTEzsigntemplatepublicNote() string`

GetTEzsigntemplatepublicNote returns the TEzsigntemplatepublicNote field if non-nil, zero value otherwise.

### GetTEzsigntemplatepublicNoteOk

`func (o *EzsigntemplatepublicResponseCompound) GetTEzsigntemplatepublicNoteOk() (*string, bool)`

GetTEzsigntemplatepublicNoteOk returns a tuple with the TEzsigntemplatepublicNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEzsigntemplatepublicNote

`func (o *EzsigntemplatepublicResponseCompound) SetTEzsigntemplatepublicNote(v string)`

SetTEzsigntemplatepublicNote sets TEzsigntemplatepublicNote field to given value.


### GetEEzsigntemplatepublicLimittype

`func (o *EzsigntemplatepublicResponseCompound) GetEEzsigntemplatepublicLimittype() FieldEEzsigntemplatepublicLimittype`

GetEEzsigntemplatepublicLimittype returns the EEzsigntemplatepublicLimittype field if non-nil, zero value otherwise.

### GetEEzsigntemplatepublicLimittypeOk

`func (o *EzsigntemplatepublicResponseCompound) GetEEzsigntemplatepublicLimittypeOk() (*FieldEEzsigntemplatepublicLimittype, bool)`

GetEEzsigntemplatepublicLimittypeOk returns a tuple with the EEzsigntemplatepublicLimittype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplatepublicLimittype

`func (o *EzsigntemplatepublicResponseCompound) SetEEzsigntemplatepublicLimittype(v FieldEEzsigntemplatepublicLimittype)`

SetEEzsigntemplatepublicLimittype sets EEzsigntemplatepublicLimittype field to given value.


### GetIEzsigntemplatepublicLimit

`func (o *EzsigntemplatepublicResponseCompound) GetIEzsigntemplatepublicLimit() int32`

GetIEzsigntemplatepublicLimit returns the IEzsigntemplatepublicLimit field if non-nil, zero value otherwise.

### GetIEzsigntemplatepublicLimitOk

`func (o *EzsigntemplatepublicResponseCompound) GetIEzsigntemplatepublicLimitOk() (*int32, bool)`

GetIEzsigntemplatepublicLimitOk returns a tuple with the IEzsigntemplatepublicLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplatepublicLimit

`func (o *EzsigntemplatepublicResponseCompound) SetIEzsigntemplatepublicLimit(v int32)`

SetIEzsigntemplatepublicLimit sets IEzsigntemplatepublicLimit field to given value.


### GetIEzsigntemplatepublicLimitexceeded

`func (o *EzsigntemplatepublicResponseCompound) GetIEzsigntemplatepublicLimitexceeded() int32`

GetIEzsigntemplatepublicLimitexceeded returns the IEzsigntemplatepublicLimitexceeded field if non-nil, zero value otherwise.

### GetIEzsigntemplatepublicLimitexceededOk

`func (o *EzsigntemplatepublicResponseCompound) GetIEzsigntemplatepublicLimitexceededOk() (*int32, bool)`

GetIEzsigntemplatepublicLimitexceededOk returns a tuple with the IEzsigntemplatepublicLimitexceeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplatepublicLimitexceeded

`func (o *EzsigntemplatepublicResponseCompound) SetIEzsigntemplatepublicLimitexceeded(v int32)`

SetIEzsigntemplatepublicLimitexceeded sets IEzsigntemplatepublicLimitexceeded field to given value.


### GetDtEzsigntemplatepublicLimitexceededsince

`func (o *EzsigntemplatepublicResponseCompound) GetDtEzsigntemplatepublicLimitexceededsince() string`

GetDtEzsigntemplatepublicLimitexceededsince returns the DtEzsigntemplatepublicLimitexceededsince field if non-nil, zero value otherwise.

### GetDtEzsigntemplatepublicLimitexceededsinceOk

`func (o *EzsigntemplatepublicResponseCompound) GetDtEzsigntemplatepublicLimitexceededsinceOk() (*string, bool)`

GetDtEzsigntemplatepublicLimitexceededsinceOk returns a tuple with the DtEzsigntemplatepublicLimitexceededsince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtEzsigntemplatepublicLimitexceededsince

`func (o *EzsigntemplatepublicResponseCompound) SetDtEzsigntemplatepublicLimitexceededsince(v string)`

SetDtEzsigntemplatepublicLimitexceededsince sets DtEzsigntemplatepublicLimitexceededsince field to given value.


### GetSEzsigntemplatepublicUrl

`func (o *EzsigntemplatepublicResponseCompound) GetSEzsigntemplatepublicUrl() string`

GetSEzsigntemplatepublicUrl returns the SEzsigntemplatepublicUrl field if non-nil, zero value otherwise.

### GetSEzsigntemplatepublicUrlOk

`func (o *EzsigntemplatepublicResponseCompound) GetSEzsigntemplatepublicUrlOk() (*string, bool)`

GetSEzsigntemplatepublicUrlOk returns a tuple with the SEzsigntemplatepublicUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplatepublicUrl

`func (o *EzsigntemplatepublicResponseCompound) SetSEzsigntemplatepublicUrl(v string)`

SetSEzsigntemplatepublicUrl sets SEzsigntemplatepublicUrl field to given value.


### GetSEzsigntemplatepublicEzsigntemplatedescription

`func (o *EzsigntemplatepublicResponseCompound) GetSEzsigntemplatepublicEzsigntemplatedescription() string`

GetSEzsigntemplatepublicEzsigntemplatedescription returns the SEzsigntemplatepublicEzsigntemplatedescription field if non-nil, zero value otherwise.

### GetSEzsigntemplatepublicEzsigntemplatedescriptionOk

`func (o *EzsigntemplatepublicResponseCompound) GetSEzsigntemplatepublicEzsigntemplatedescriptionOk() (*string, bool)`

GetSEzsigntemplatepublicEzsigntemplatedescriptionOk returns a tuple with the SEzsigntemplatepublicEzsigntemplatedescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplatepublicEzsigntemplatedescription

`func (o *EzsigntemplatepublicResponseCompound) SetSEzsigntemplatepublicEzsigntemplatedescription(v string)`

SetSEzsigntemplatepublicEzsigntemplatedescription sets SEzsigntemplatepublicEzsigntemplatedescription field to given value.


### GetObjAudit

`func (o *EzsigntemplatepublicResponseCompound) GetObjAudit() CommonAudit`

GetObjAudit returns the ObjAudit field if non-nil, zero value otherwise.

### GetObjAuditOk

`func (o *EzsigntemplatepublicResponseCompound) GetObjAuditOk() (*CommonAudit, bool)`

GetObjAuditOk returns a tuple with the ObjAudit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjAudit

`func (o *EzsigntemplatepublicResponseCompound) SetObjAudit(v CommonAudit)`

SetObjAudit sets ObjAudit field to given value.

### HasObjAudit

`func (o *EzsigntemplatepublicResponseCompound) HasObjAudit() bool`

HasObjAudit returns a boolean if a field has been set.

### GetAObjEzsignfolderezsigntemplatepublic

`func (o *EzsigntemplatepublicResponseCompound) GetAObjEzsignfolderezsigntemplatepublic() []CustomEzsignfolderezsigntemplatepublicResponse`

GetAObjEzsignfolderezsigntemplatepublic returns the AObjEzsignfolderezsigntemplatepublic field if non-nil, zero value otherwise.

### GetAObjEzsignfolderezsigntemplatepublicOk

`func (o *EzsigntemplatepublicResponseCompound) GetAObjEzsignfolderezsigntemplatepublicOk() (*[]CustomEzsignfolderezsigntemplatepublicResponse, bool)`

GetAObjEzsignfolderezsigntemplatepublicOk returns a tuple with the AObjEzsignfolderezsigntemplatepublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzsignfolderezsigntemplatepublic

`func (o *EzsigntemplatepublicResponseCompound) SetAObjEzsignfolderezsigntemplatepublic(v []CustomEzsignfolderezsigntemplatepublicResponse)`

SetAObjEzsignfolderezsigntemplatepublic sets AObjEzsignfolderezsigntemplatepublic field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


