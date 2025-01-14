# EzsigntemplatepublicListElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsigntemplatepublicID** | **int32** | The unique ID of the Ezsigntemplatepublic | 
**FkiEzsignfoldertypeID** | **int32** | The unique ID of the Ezsignfoldertype. | 
**SEzsignfoldertypeNameX** | **string** | The name of the Ezsignfoldertype in the language of the requester | 
**FkiUserlogintypeID** | **int32** | The unique ID of the Userlogintype  Valid values:  |Value|Description|Detail| |-|-|-| |1|**Email Only**|The Ezsignsigner will receive a secure link by email| |2|**Email and phone or SMS**|The Ezsignsigner will receive a secure link by email and will need to authenticate using SMS or Phone call. **Additional fee applies**| |3|**Email and secret question**|The Ezsignsigner will receive a secure link by email and will need to authenticate using a predefined question and answer| |4|**In person only**|The Ezsignsigner will only be able to sign \&quot;In-Person\&quot; and there won&#39;t be any authentication. No email will be sent for invitation to sign. Make sure you evaluate the risk of signature denial and at minimum, we recommend you use a handwritten signature type| |5|**In person with phone or SMS**|The Ezsignsigner will only be able to sign \&quot;In-Person\&quot; and will need to authenticate using SMS or Phone call. No email will be sent for invitation to sign. **Additional fee applies**| |6|**Embedded**|The Ezsignsigner will only be able to sign in the embedded solution. No email will be sent for invitation to sign. **Additional fee applies**|   |7|**Embedded with phone or SMS**|The Ezsignsigner will only be able to sign in the embedded solution and will need to authenticate using SMS or Phone call. No email will be sent for invitation to sign. **Additional fee applies**|   |8|**No validation**|The Ezsignsigner will not receive an email and won&#39;t have to validate his connection using 2 factor. **Additional fee applies**|      |9|**Sms only**|The Ezsignsigner will not receive an email but will will need to authenticate using SMS. **Additional fee applies**|      | 
**FkiEzsigntemplateID** | Pointer to **int32** | The unique ID of the Ezsigntemplate | [optional] 
**FkiEzsigntemplatepackageID** | Pointer to **int32** | The unique ID of the Ezsigntemplatepackage | [optional] 
**SEzsigntemplatepublicDescription** | **string** | The description of the Ezsigntemplatepublic | 
**BEzsigntemplatepublicIsactive** | **bool** | Whether the ezsigntemplatepublic is active or not | 
**TEzsigntemplatepublicNote** | **string** | The note of the Ezsigntemplatepublic | 
**EEzsigntemplatepublicLimittype** | [**FieldEEzsigntemplatepublicLimittype**](FieldEEzsigntemplatepublicLimittype.md) |  | 
**IEzsigntemplatepublicLimit** | **int32** | The limit of the Ezsigntemplatepublic | 
**IEzsigntemplatepublicLimitexceeded** | **int32** | The limitexceeded of the Ezsigntemplatepublic | 
**DtEzsigntemplatepublicLimitexceededsince** | **string** | The limitexceededsince of the Ezsigntemplatepublic | 
**IEzsignfolder** | **int32** | The total number of Ezsignfolders using the Ezsigntemplatepublic | 
**IEzsigndocument** | **int32** | The total number of Ezsigndocuments using the Ezsigntemplatepublic | 
**SEzsigntemplatepublicEzsigntemplatedescription** | **string** | The Ezsigntemplate/Ezsigntemplatepackage description | 

## Methods

### NewEzsigntemplatepublicListElement

`func NewEzsigntemplatepublicListElement(pkiEzsigntemplatepublicID int32, fkiEzsignfoldertypeID int32, sEzsignfoldertypeNameX string, fkiUserlogintypeID int32, sEzsigntemplatepublicDescription string, bEzsigntemplatepublicIsactive bool, tEzsigntemplatepublicNote string, eEzsigntemplatepublicLimittype FieldEEzsigntemplatepublicLimittype, iEzsigntemplatepublicLimit int32, iEzsigntemplatepublicLimitexceeded int32, dtEzsigntemplatepublicLimitexceededsince string, iEzsignfolder int32, iEzsigndocument int32, sEzsigntemplatepublicEzsigntemplatedescription string, ) *EzsigntemplatepublicListElement`

NewEzsigntemplatepublicListElement instantiates a new EzsigntemplatepublicListElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigntemplatepublicListElementWithDefaults

`func NewEzsigntemplatepublicListElementWithDefaults() *EzsigntemplatepublicListElement`

NewEzsigntemplatepublicListElementWithDefaults instantiates a new EzsigntemplatepublicListElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsigntemplatepublicID

`func (o *EzsigntemplatepublicListElement) GetPkiEzsigntemplatepublicID() int32`

GetPkiEzsigntemplatepublicID returns the PkiEzsigntemplatepublicID field if non-nil, zero value otherwise.

### GetPkiEzsigntemplatepublicIDOk

`func (o *EzsigntemplatepublicListElement) GetPkiEzsigntemplatepublicIDOk() (*int32, bool)`

GetPkiEzsigntemplatepublicIDOk returns a tuple with the PkiEzsigntemplatepublicID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsigntemplatepublicID

`func (o *EzsigntemplatepublicListElement) SetPkiEzsigntemplatepublicID(v int32)`

SetPkiEzsigntemplatepublicID sets PkiEzsigntemplatepublicID field to given value.


### GetFkiEzsignfoldertypeID

`func (o *EzsigntemplatepublicListElement) GetFkiEzsignfoldertypeID() int32`

GetFkiEzsignfoldertypeID returns the FkiEzsignfoldertypeID field if non-nil, zero value otherwise.

### GetFkiEzsignfoldertypeIDOk

`func (o *EzsigntemplatepublicListElement) GetFkiEzsignfoldertypeIDOk() (*int32, bool)`

GetFkiEzsignfoldertypeIDOk returns a tuple with the FkiEzsignfoldertypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfoldertypeID

`func (o *EzsigntemplatepublicListElement) SetFkiEzsignfoldertypeID(v int32)`

SetFkiEzsignfoldertypeID sets FkiEzsignfoldertypeID field to given value.


### GetSEzsignfoldertypeNameX

`func (o *EzsigntemplatepublicListElement) GetSEzsignfoldertypeNameX() string`

GetSEzsignfoldertypeNameX returns the SEzsignfoldertypeNameX field if non-nil, zero value otherwise.

### GetSEzsignfoldertypeNameXOk

`func (o *EzsigntemplatepublicListElement) GetSEzsignfoldertypeNameXOk() (*string, bool)`

GetSEzsignfoldertypeNameXOk returns a tuple with the SEzsignfoldertypeNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignfoldertypeNameX

`func (o *EzsigntemplatepublicListElement) SetSEzsignfoldertypeNameX(v string)`

SetSEzsignfoldertypeNameX sets SEzsignfoldertypeNameX field to given value.


### GetFkiUserlogintypeID

`func (o *EzsigntemplatepublicListElement) GetFkiUserlogintypeID() int32`

GetFkiUserlogintypeID returns the FkiUserlogintypeID field if non-nil, zero value otherwise.

### GetFkiUserlogintypeIDOk

`func (o *EzsigntemplatepublicListElement) GetFkiUserlogintypeIDOk() (*int32, bool)`

GetFkiUserlogintypeIDOk returns a tuple with the FkiUserlogintypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserlogintypeID

`func (o *EzsigntemplatepublicListElement) SetFkiUserlogintypeID(v int32)`

SetFkiUserlogintypeID sets FkiUserlogintypeID field to given value.


### GetFkiEzsigntemplateID

`func (o *EzsigntemplatepublicListElement) GetFkiEzsigntemplateID() int32`

GetFkiEzsigntemplateID returns the FkiEzsigntemplateID field if non-nil, zero value otherwise.

### GetFkiEzsigntemplateIDOk

`func (o *EzsigntemplatepublicListElement) GetFkiEzsigntemplateIDOk() (*int32, bool)`

GetFkiEzsigntemplateIDOk returns a tuple with the FkiEzsigntemplateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntemplateID

`func (o *EzsigntemplatepublicListElement) SetFkiEzsigntemplateID(v int32)`

SetFkiEzsigntemplateID sets FkiEzsigntemplateID field to given value.

### HasFkiEzsigntemplateID

`func (o *EzsigntemplatepublicListElement) HasFkiEzsigntemplateID() bool`

HasFkiEzsigntemplateID returns a boolean if a field has been set.

### GetFkiEzsigntemplatepackageID

`func (o *EzsigntemplatepublicListElement) GetFkiEzsigntemplatepackageID() int32`

GetFkiEzsigntemplatepackageID returns the FkiEzsigntemplatepackageID field if non-nil, zero value otherwise.

### GetFkiEzsigntemplatepackageIDOk

`func (o *EzsigntemplatepublicListElement) GetFkiEzsigntemplatepackageIDOk() (*int32, bool)`

GetFkiEzsigntemplatepackageIDOk returns a tuple with the FkiEzsigntemplatepackageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntemplatepackageID

`func (o *EzsigntemplatepublicListElement) SetFkiEzsigntemplatepackageID(v int32)`

SetFkiEzsigntemplatepackageID sets FkiEzsigntemplatepackageID field to given value.

### HasFkiEzsigntemplatepackageID

`func (o *EzsigntemplatepublicListElement) HasFkiEzsigntemplatepackageID() bool`

HasFkiEzsigntemplatepackageID returns a boolean if a field has been set.

### GetSEzsigntemplatepublicDescription

`func (o *EzsigntemplatepublicListElement) GetSEzsigntemplatepublicDescription() string`

GetSEzsigntemplatepublicDescription returns the SEzsigntemplatepublicDescription field if non-nil, zero value otherwise.

### GetSEzsigntemplatepublicDescriptionOk

`func (o *EzsigntemplatepublicListElement) GetSEzsigntemplatepublicDescriptionOk() (*string, bool)`

GetSEzsigntemplatepublicDescriptionOk returns a tuple with the SEzsigntemplatepublicDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplatepublicDescription

`func (o *EzsigntemplatepublicListElement) SetSEzsigntemplatepublicDescription(v string)`

SetSEzsigntemplatepublicDescription sets SEzsigntemplatepublicDescription field to given value.


### GetBEzsigntemplatepublicIsactive

`func (o *EzsigntemplatepublicListElement) GetBEzsigntemplatepublicIsactive() bool`

GetBEzsigntemplatepublicIsactive returns the BEzsigntemplatepublicIsactive field if non-nil, zero value otherwise.

### GetBEzsigntemplatepublicIsactiveOk

`func (o *EzsigntemplatepublicListElement) GetBEzsigntemplatepublicIsactiveOk() (*bool, bool)`

GetBEzsigntemplatepublicIsactiveOk returns a tuple with the BEzsigntemplatepublicIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigntemplatepublicIsactive

`func (o *EzsigntemplatepublicListElement) SetBEzsigntemplatepublicIsactive(v bool)`

SetBEzsigntemplatepublicIsactive sets BEzsigntemplatepublicIsactive field to given value.


### GetTEzsigntemplatepublicNote

`func (o *EzsigntemplatepublicListElement) GetTEzsigntemplatepublicNote() string`

GetTEzsigntemplatepublicNote returns the TEzsigntemplatepublicNote field if non-nil, zero value otherwise.

### GetTEzsigntemplatepublicNoteOk

`func (o *EzsigntemplatepublicListElement) GetTEzsigntemplatepublicNoteOk() (*string, bool)`

GetTEzsigntemplatepublicNoteOk returns a tuple with the TEzsigntemplatepublicNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEzsigntemplatepublicNote

`func (o *EzsigntemplatepublicListElement) SetTEzsigntemplatepublicNote(v string)`

SetTEzsigntemplatepublicNote sets TEzsigntemplatepublicNote field to given value.


### GetEEzsigntemplatepublicLimittype

`func (o *EzsigntemplatepublicListElement) GetEEzsigntemplatepublicLimittype() FieldEEzsigntemplatepublicLimittype`

GetEEzsigntemplatepublicLimittype returns the EEzsigntemplatepublicLimittype field if non-nil, zero value otherwise.

### GetEEzsigntemplatepublicLimittypeOk

`func (o *EzsigntemplatepublicListElement) GetEEzsigntemplatepublicLimittypeOk() (*FieldEEzsigntemplatepublicLimittype, bool)`

GetEEzsigntemplatepublicLimittypeOk returns a tuple with the EEzsigntemplatepublicLimittype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplatepublicLimittype

`func (o *EzsigntemplatepublicListElement) SetEEzsigntemplatepublicLimittype(v FieldEEzsigntemplatepublicLimittype)`

SetEEzsigntemplatepublicLimittype sets EEzsigntemplatepublicLimittype field to given value.


### GetIEzsigntemplatepublicLimit

`func (o *EzsigntemplatepublicListElement) GetIEzsigntemplatepublicLimit() int32`

GetIEzsigntemplatepublicLimit returns the IEzsigntemplatepublicLimit field if non-nil, zero value otherwise.

### GetIEzsigntemplatepublicLimitOk

`func (o *EzsigntemplatepublicListElement) GetIEzsigntemplatepublicLimitOk() (*int32, bool)`

GetIEzsigntemplatepublicLimitOk returns a tuple with the IEzsigntemplatepublicLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplatepublicLimit

`func (o *EzsigntemplatepublicListElement) SetIEzsigntemplatepublicLimit(v int32)`

SetIEzsigntemplatepublicLimit sets IEzsigntemplatepublicLimit field to given value.


### GetIEzsigntemplatepublicLimitexceeded

`func (o *EzsigntemplatepublicListElement) GetIEzsigntemplatepublicLimitexceeded() int32`

GetIEzsigntemplatepublicLimitexceeded returns the IEzsigntemplatepublicLimitexceeded field if non-nil, zero value otherwise.

### GetIEzsigntemplatepublicLimitexceededOk

`func (o *EzsigntemplatepublicListElement) GetIEzsigntemplatepublicLimitexceededOk() (*int32, bool)`

GetIEzsigntemplatepublicLimitexceededOk returns a tuple with the IEzsigntemplatepublicLimitexceeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplatepublicLimitexceeded

`func (o *EzsigntemplatepublicListElement) SetIEzsigntemplatepublicLimitexceeded(v int32)`

SetIEzsigntemplatepublicLimitexceeded sets IEzsigntemplatepublicLimitexceeded field to given value.


### GetDtEzsigntemplatepublicLimitexceededsince

`func (o *EzsigntemplatepublicListElement) GetDtEzsigntemplatepublicLimitexceededsince() string`

GetDtEzsigntemplatepublicLimitexceededsince returns the DtEzsigntemplatepublicLimitexceededsince field if non-nil, zero value otherwise.

### GetDtEzsigntemplatepublicLimitexceededsinceOk

`func (o *EzsigntemplatepublicListElement) GetDtEzsigntemplatepublicLimitexceededsinceOk() (*string, bool)`

GetDtEzsigntemplatepublicLimitexceededsinceOk returns a tuple with the DtEzsigntemplatepublicLimitexceededsince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtEzsigntemplatepublicLimitexceededsince

`func (o *EzsigntemplatepublicListElement) SetDtEzsigntemplatepublicLimitexceededsince(v string)`

SetDtEzsigntemplatepublicLimitexceededsince sets DtEzsigntemplatepublicLimitexceededsince field to given value.


### GetIEzsignfolder

`func (o *EzsigntemplatepublicListElement) GetIEzsignfolder() int32`

GetIEzsignfolder returns the IEzsignfolder field if non-nil, zero value otherwise.

### GetIEzsignfolderOk

`func (o *EzsigntemplatepublicListElement) GetIEzsignfolderOk() (*int32, bool)`

GetIEzsignfolderOk returns a tuple with the IEzsignfolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignfolder

`func (o *EzsigntemplatepublicListElement) SetIEzsignfolder(v int32)`

SetIEzsignfolder sets IEzsignfolder field to given value.


### GetIEzsigndocument

`func (o *EzsigntemplatepublicListElement) GetIEzsigndocument() int32`

GetIEzsigndocument returns the IEzsigndocument field if non-nil, zero value otherwise.

### GetIEzsigndocumentOk

`func (o *EzsigntemplatepublicListElement) GetIEzsigndocumentOk() (*int32, bool)`

GetIEzsigndocumentOk returns a tuple with the IEzsigndocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigndocument

`func (o *EzsigntemplatepublicListElement) SetIEzsigndocument(v int32)`

SetIEzsigndocument sets IEzsigndocument field to given value.


### GetSEzsigntemplatepublicEzsigntemplatedescription

`func (o *EzsigntemplatepublicListElement) GetSEzsigntemplatepublicEzsigntemplatedescription() string`

GetSEzsigntemplatepublicEzsigntemplatedescription returns the SEzsigntemplatepublicEzsigntemplatedescription field if non-nil, zero value otherwise.

### GetSEzsigntemplatepublicEzsigntemplatedescriptionOk

`func (o *EzsigntemplatepublicListElement) GetSEzsigntemplatepublicEzsigntemplatedescriptionOk() (*string, bool)`

GetSEzsigntemplatepublicEzsigntemplatedescriptionOk returns a tuple with the SEzsigntemplatepublicEzsigntemplatedescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplatepublicEzsigntemplatedescription

`func (o *EzsigntemplatepublicListElement) SetSEzsigntemplatepublicEzsigntemplatedescription(v string)`

SetSEzsigntemplatepublicEzsigntemplatedescription sets SEzsigntemplatepublicEzsigntemplatedescription field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


