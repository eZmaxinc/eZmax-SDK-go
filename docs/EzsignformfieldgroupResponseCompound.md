# EzsignformfieldgroupResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsignformfieldgroupID** | **int32** | The unique ID of the Ezsignformfieldgroup | 
**FkiEzsigndocumentID** | **int32** | The unique ID of the Ezsigndocument | 
**EEzsignformfieldgroupType** | [**FieldEEzsignformfieldgroupType**](FieldEEzsignformfieldgroupType.md) |  | 
**EEzsignformfieldgroupSignerrequirement** | Pointer to [**FieldEEzsignformfieldgroupSignerrequirement**](FieldEEzsignformfieldgroupSignerrequirement.md) |  | [optional] 
**SEzsignformfieldgroupLabel** | **string** | The Label for the Ezsignformfieldgroup | 
**IEzsignformfieldgroupStep** | **int32** | The step when the Ezsignsigner will be invited to fill the form fields | 
**SEzsignformfieldgroupDefaultvalue** | Pointer to **string** | The default value for the Ezsignformfieldgroup  You can use the codes below and they will be replaced at signature time.    | Code | Description | Example | | ------------------------- | ------------ | ------------ | | {sUserFirstname} | The first name of the contact | John | | {sUserLastname} | The last name of the contact | Doe | | {sUserJobtitle} | The job title | Sales Representative | | {sCompany} | Company name | eZmax Solutions Inc. | | {sEmailAddress} | The email address | email@example.com | | {sPhoneE164} | A phone number in E.164 Format | +15149901516 | | {sPhoneE164Cell} | A phone number in E.164 Format | +15149901516 | | [optional] 
**IEzsignformfieldgroupFilledmin** | **int32** | The minimum number of Ezsignformfield that must be filled in the Ezsignformfieldgroup | 
**IEzsignformfieldgroupFilledmax** | **int32** | The maximum number of Ezsignformfield that must be filled in the Ezsignformfieldgroup | 
**BEzsignformfieldgroupReadonly** | **bool** | Whether the Ezsignformfieldgroup is read only or not. | 
**IEzsignformfieldgroupMaxlength** | Pointer to **int32** | The maximum length for the value in the Ezsignformfieldgroup  This can only be set if eEzsignformfieldgroupType is **Text** or **Textarea** | [optional] 
**BEzsignformfieldgroupEncrypted** | Pointer to **bool** | Whether the Ezsignformfieldgroup is encrypted in the database or not. Encrypted values are not displayed on the Ezsigndocument. This can only be set if eEzsignformfieldgroupType is **Text** or **Textarea** | [optional] 
**EEzsignformfieldgroupTextvalidation** | Pointer to [**EnumTextvalidation**](EnumTextvalidation.md) |  | [optional] 
**SEzsignformfieldgroupRegexp** | Pointer to **string** | A regular expression to indicate what values are acceptable for the Ezsignformfieldgroup.  This can only be set if eEzsignformfieldgroupType is **Text** or **Textarea** | [optional] 
**SEzsignformfieldgroupTextvalidationcustommessage** | Pointer to **string** | Description of validation rule. Show by signatory. | [optional] 
**TEzsignformfieldgroupTooltip** | Pointer to **string** | A tooltip that will be presented to Ezsignsigner about the Ezsignformfieldgroup | [optional] 
**EEzsignformfieldgroupTooltipposition** | Pointer to [**FieldEEzsignformfieldgroupTooltipposition**](FieldEEzsignformfieldgroupTooltipposition.md) |  | [optional] 
**AObjEzsignformfield** | [**[]EzsignformfieldResponseCompound**](EzsignformfieldResponseCompound.md) |  | 
**AObjDropdownElement** | Pointer to [**[]CustomDropdownElementResponse**](CustomDropdownElementResponse.md) |  | [optional] 
**AObjEzsignformfieldgroupsigner** | [**[]EzsignformfieldgroupsignerResponse**](EzsignformfieldgroupsignerResponse.md) |  | 

## Methods

### NewEzsignformfieldgroupResponseCompound

`func NewEzsignformfieldgroupResponseCompound(pkiEzsignformfieldgroupID int32, fkiEzsigndocumentID int32, eEzsignformfieldgroupType FieldEEzsignformfieldgroupType, sEzsignformfieldgroupLabel string, iEzsignformfieldgroupStep int32, iEzsignformfieldgroupFilledmin int32, iEzsignformfieldgroupFilledmax int32, bEzsignformfieldgroupReadonly bool, aObjEzsignformfield []EzsignformfieldResponseCompound, aObjEzsignformfieldgroupsigner []EzsignformfieldgroupsignerResponseCompound, ) *EzsignformfieldgroupResponseCompound`

NewEzsignformfieldgroupResponseCompound instantiates a new EzsignformfieldgroupResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignformfieldgroupResponseCompoundWithDefaults

`func NewEzsignformfieldgroupResponseCompoundWithDefaults() *EzsignformfieldgroupResponseCompound`

NewEzsignformfieldgroupResponseCompoundWithDefaults instantiates a new EzsignformfieldgroupResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsignformfieldgroupID

`func (o *EzsignformfieldgroupResponseCompound) GetPkiEzsignformfieldgroupID() int32`

GetPkiEzsignformfieldgroupID returns the PkiEzsignformfieldgroupID field if non-nil, zero value otherwise.

### GetPkiEzsignformfieldgroupIDOk

`func (o *EzsignformfieldgroupResponseCompound) GetPkiEzsignformfieldgroupIDOk() (*int32, bool)`

GetPkiEzsignformfieldgroupIDOk returns a tuple with the PkiEzsignformfieldgroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsignformfieldgroupID

`func (o *EzsignformfieldgroupResponseCompound) SetPkiEzsignformfieldgroupID(v int32)`

SetPkiEzsignformfieldgroupID sets PkiEzsignformfieldgroupID field to given value.


### GetFkiEzsigndocumentID

`func (o *EzsignformfieldgroupResponseCompound) GetFkiEzsigndocumentID() int32`

GetFkiEzsigndocumentID returns the FkiEzsigndocumentID field if non-nil, zero value otherwise.

### GetFkiEzsigndocumentIDOk

`func (o *EzsignformfieldgroupResponseCompound) GetFkiEzsigndocumentIDOk() (*int32, bool)`

GetFkiEzsigndocumentIDOk returns a tuple with the FkiEzsigndocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigndocumentID

`func (o *EzsignformfieldgroupResponseCompound) SetFkiEzsigndocumentID(v int32)`

SetFkiEzsigndocumentID sets FkiEzsigndocumentID field to given value.


### GetEEzsignformfieldgroupType

`func (o *EzsignformfieldgroupResponseCompound) GetEEzsignformfieldgroupType() FieldEEzsignformfieldgroupType`

GetEEzsignformfieldgroupType returns the EEzsignformfieldgroupType field if non-nil, zero value otherwise.

### GetEEzsignformfieldgroupTypeOk

`func (o *EzsignformfieldgroupResponseCompound) GetEEzsignformfieldgroupTypeOk() (*FieldEEzsignformfieldgroupType, bool)`

GetEEzsignformfieldgroupTypeOk returns a tuple with the EEzsignformfieldgroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignformfieldgroupType

`func (o *EzsignformfieldgroupResponseCompound) SetEEzsignformfieldgroupType(v FieldEEzsignformfieldgroupType)`

SetEEzsignformfieldgroupType sets EEzsignformfieldgroupType field to given value.


### GetEEzsignformfieldgroupSignerrequirement

`func (o *EzsignformfieldgroupResponseCompound) GetEEzsignformfieldgroupSignerrequirement() FieldEEzsignformfieldgroupSignerrequirement`

GetEEzsignformfieldgroupSignerrequirement returns the EEzsignformfieldgroupSignerrequirement field if non-nil, zero value otherwise.

### GetEEzsignformfieldgroupSignerrequirementOk

`func (o *EzsignformfieldgroupResponseCompound) GetEEzsignformfieldgroupSignerrequirementOk() (*FieldEEzsignformfieldgroupSignerrequirement, bool)`

GetEEzsignformfieldgroupSignerrequirementOk returns a tuple with the EEzsignformfieldgroupSignerrequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignformfieldgroupSignerrequirement

`func (o *EzsignformfieldgroupResponseCompound) SetEEzsignformfieldgroupSignerrequirement(v FieldEEzsignformfieldgroupSignerrequirement)`

SetEEzsignformfieldgroupSignerrequirement sets EEzsignformfieldgroupSignerrequirement field to given value.

### HasEEzsignformfieldgroupSignerrequirement

`func (o *EzsignformfieldgroupResponseCompound) HasEEzsignformfieldgroupSignerrequirement() bool`

HasEEzsignformfieldgroupSignerrequirement returns a boolean if a field has been set.

### GetSEzsignformfieldgroupLabel

`func (o *EzsignformfieldgroupResponseCompound) GetSEzsignformfieldgroupLabel() string`

GetSEzsignformfieldgroupLabel returns the SEzsignformfieldgroupLabel field if non-nil, zero value otherwise.

### GetSEzsignformfieldgroupLabelOk

`func (o *EzsignformfieldgroupResponseCompound) GetSEzsignformfieldgroupLabelOk() (*string, bool)`

GetSEzsignformfieldgroupLabelOk returns a tuple with the SEzsignformfieldgroupLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignformfieldgroupLabel

`func (o *EzsignformfieldgroupResponseCompound) SetSEzsignformfieldgroupLabel(v string)`

SetSEzsignformfieldgroupLabel sets SEzsignformfieldgroupLabel field to given value.


### GetIEzsignformfieldgroupStep

`func (o *EzsignformfieldgroupResponseCompound) GetIEzsignformfieldgroupStep() int32`

GetIEzsignformfieldgroupStep returns the IEzsignformfieldgroupStep field if non-nil, zero value otherwise.

### GetIEzsignformfieldgroupStepOk

`func (o *EzsignformfieldgroupResponseCompound) GetIEzsignformfieldgroupStepOk() (*int32, bool)`

GetIEzsignformfieldgroupStepOk returns a tuple with the IEzsignformfieldgroupStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignformfieldgroupStep

`func (o *EzsignformfieldgroupResponseCompound) SetIEzsignformfieldgroupStep(v int32)`

SetIEzsignformfieldgroupStep sets IEzsignformfieldgroupStep field to given value.


### GetSEzsignformfieldgroupDefaultvalue

`func (o *EzsignformfieldgroupResponseCompound) GetSEzsignformfieldgroupDefaultvalue() string`

GetSEzsignformfieldgroupDefaultvalue returns the SEzsignformfieldgroupDefaultvalue field if non-nil, zero value otherwise.

### GetSEzsignformfieldgroupDefaultvalueOk

`func (o *EzsignformfieldgroupResponseCompound) GetSEzsignformfieldgroupDefaultvalueOk() (*string, bool)`

GetSEzsignformfieldgroupDefaultvalueOk returns a tuple with the SEzsignformfieldgroupDefaultvalue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignformfieldgroupDefaultvalue

`func (o *EzsignformfieldgroupResponseCompound) SetSEzsignformfieldgroupDefaultvalue(v string)`

SetSEzsignformfieldgroupDefaultvalue sets SEzsignformfieldgroupDefaultvalue field to given value.

### HasSEzsignformfieldgroupDefaultvalue

`func (o *EzsignformfieldgroupResponseCompound) HasSEzsignformfieldgroupDefaultvalue() bool`

HasSEzsignformfieldgroupDefaultvalue returns a boolean if a field has been set.

### GetIEzsignformfieldgroupFilledmin

`func (o *EzsignformfieldgroupResponseCompound) GetIEzsignformfieldgroupFilledmin() int32`

GetIEzsignformfieldgroupFilledmin returns the IEzsignformfieldgroupFilledmin field if non-nil, zero value otherwise.

### GetIEzsignformfieldgroupFilledminOk

`func (o *EzsignformfieldgroupResponseCompound) GetIEzsignformfieldgroupFilledminOk() (*int32, bool)`

GetIEzsignformfieldgroupFilledminOk returns a tuple with the IEzsignformfieldgroupFilledmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignformfieldgroupFilledmin

`func (o *EzsignformfieldgroupResponseCompound) SetIEzsignformfieldgroupFilledmin(v int32)`

SetIEzsignformfieldgroupFilledmin sets IEzsignformfieldgroupFilledmin field to given value.


### GetIEzsignformfieldgroupFilledmax

`func (o *EzsignformfieldgroupResponseCompound) GetIEzsignformfieldgroupFilledmax() int32`

GetIEzsignformfieldgroupFilledmax returns the IEzsignformfieldgroupFilledmax field if non-nil, zero value otherwise.

### GetIEzsignformfieldgroupFilledmaxOk

`func (o *EzsignformfieldgroupResponseCompound) GetIEzsignformfieldgroupFilledmaxOk() (*int32, bool)`

GetIEzsignformfieldgroupFilledmaxOk returns a tuple with the IEzsignformfieldgroupFilledmax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignformfieldgroupFilledmax

`func (o *EzsignformfieldgroupResponseCompound) SetIEzsignformfieldgroupFilledmax(v int32)`

SetIEzsignformfieldgroupFilledmax sets IEzsignformfieldgroupFilledmax field to given value.


### GetBEzsignformfieldgroupReadonly

`func (o *EzsignformfieldgroupResponseCompound) GetBEzsignformfieldgroupReadonly() bool`

GetBEzsignformfieldgroupReadonly returns the BEzsignformfieldgroupReadonly field if non-nil, zero value otherwise.

### GetBEzsignformfieldgroupReadonlyOk

`func (o *EzsignformfieldgroupResponseCompound) GetBEzsignformfieldgroupReadonlyOk() (*bool, bool)`

GetBEzsignformfieldgroupReadonlyOk returns a tuple with the BEzsignformfieldgroupReadonly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignformfieldgroupReadonly

`func (o *EzsignformfieldgroupResponseCompound) SetBEzsignformfieldgroupReadonly(v bool)`

SetBEzsignformfieldgroupReadonly sets BEzsignformfieldgroupReadonly field to given value.


### GetIEzsignformfieldgroupMaxlength

`func (o *EzsignformfieldgroupResponseCompound) GetIEzsignformfieldgroupMaxlength() int32`

GetIEzsignformfieldgroupMaxlength returns the IEzsignformfieldgroupMaxlength field if non-nil, zero value otherwise.

### GetIEzsignformfieldgroupMaxlengthOk

`func (o *EzsignformfieldgroupResponseCompound) GetIEzsignformfieldgroupMaxlengthOk() (*int32, bool)`

GetIEzsignformfieldgroupMaxlengthOk returns a tuple with the IEzsignformfieldgroupMaxlength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignformfieldgroupMaxlength

`func (o *EzsignformfieldgroupResponseCompound) SetIEzsignformfieldgroupMaxlength(v int32)`

SetIEzsignformfieldgroupMaxlength sets IEzsignformfieldgroupMaxlength field to given value.

### HasIEzsignformfieldgroupMaxlength

`func (o *EzsignformfieldgroupResponseCompound) HasIEzsignformfieldgroupMaxlength() bool`

HasIEzsignformfieldgroupMaxlength returns a boolean if a field has been set.

### GetBEzsignformfieldgroupEncrypted

`func (o *EzsignformfieldgroupResponseCompound) GetBEzsignformfieldgroupEncrypted() bool`

GetBEzsignformfieldgroupEncrypted returns the BEzsignformfieldgroupEncrypted field if non-nil, zero value otherwise.

### GetBEzsignformfieldgroupEncryptedOk

`func (o *EzsignformfieldgroupResponseCompound) GetBEzsignformfieldgroupEncryptedOk() (*bool, bool)`

GetBEzsignformfieldgroupEncryptedOk returns a tuple with the BEzsignformfieldgroupEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignformfieldgroupEncrypted

`func (o *EzsignformfieldgroupResponseCompound) SetBEzsignformfieldgroupEncrypted(v bool)`

SetBEzsignformfieldgroupEncrypted sets BEzsignformfieldgroupEncrypted field to given value.

### HasBEzsignformfieldgroupEncrypted

`func (o *EzsignformfieldgroupResponseCompound) HasBEzsignformfieldgroupEncrypted() bool`

HasBEzsignformfieldgroupEncrypted returns a boolean if a field has been set.

### GetEEzsignformfieldgroupTextvalidation

`func (o *EzsignformfieldgroupResponseCompound) GetEEzsignformfieldgroupTextvalidation() EnumTextvalidation`

GetEEzsignformfieldgroupTextvalidation returns the EEzsignformfieldgroupTextvalidation field if non-nil, zero value otherwise.

### GetEEzsignformfieldgroupTextvalidationOk

`func (o *EzsignformfieldgroupResponseCompound) GetEEzsignformfieldgroupTextvalidationOk() (*EnumTextvalidation, bool)`

GetEEzsignformfieldgroupTextvalidationOk returns a tuple with the EEzsignformfieldgroupTextvalidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignformfieldgroupTextvalidation

`func (o *EzsignformfieldgroupResponseCompound) SetEEzsignformfieldgroupTextvalidation(v EnumTextvalidation)`

SetEEzsignformfieldgroupTextvalidation sets EEzsignformfieldgroupTextvalidation field to given value.

### HasEEzsignformfieldgroupTextvalidation

`func (o *EzsignformfieldgroupResponseCompound) HasEEzsignformfieldgroupTextvalidation() bool`

HasEEzsignformfieldgroupTextvalidation returns a boolean if a field has been set.

### GetSEzsignformfieldgroupRegexp

`func (o *EzsignformfieldgroupResponseCompound) GetSEzsignformfieldgroupRegexp() string`

GetSEzsignformfieldgroupRegexp returns the SEzsignformfieldgroupRegexp field if non-nil, zero value otherwise.

### GetSEzsignformfieldgroupRegexpOk

`func (o *EzsignformfieldgroupResponseCompound) GetSEzsignformfieldgroupRegexpOk() (*string, bool)`

GetSEzsignformfieldgroupRegexpOk returns a tuple with the SEzsignformfieldgroupRegexp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignformfieldgroupRegexp

`func (o *EzsignformfieldgroupResponseCompound) SetSEzsignformfieldgroupRegexp(v string)`

SetSEzsignformfieldgroupRegexp sets SEzsignformfieldgroupRegexp field to given value.

### HasSEzsignformfieldgroupRegexp

`func (o *EzsignformfieldgroupResponseCompound) HasSEzsignformfieldgroupRegexp() bool`

HasSEzsignformfieldgroupRegexp returns a boolean if a field has been set.

### GetSEzsignformfieldgroupTextvalidationcustommessage

`func (o *EzsignformfieldgroupResponseCompound) GetSEzsignformfieldgroupTextvalidationcustommessage() string`

GetSEzsignformfieldgroupTextvalidationcustommessage returns the SEzsignformfieldgroupTextvalidationcustommessage field if non-nil, zero value otherwise.

### GetSEzsignformfieldgroupTextvalidationcustommessageOk

`func (o *EzsignformfieldgroupResponseCompound) GetSEzsignformfieldgroupTextvalidationcustommessageOk() (*string, bool)`

GetSEzsignformfieldgroupTextvalidationcustommessageOk returns a tuple with the SEzsignformfieldgroupTextvalidationcustommessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignformfieldgroupTextvalidationcustommessage

`func (o *EzsignformfieldgroupResponseCompound) SetSEzsignformfieldgroupTextvalidationcustommessage(v string)`

SetSEzsignformfieldgroupTextvalidationcustommessage sets SEzsignformfieldgroupTextvalidationcustommessage field to given value.

### HasSEzsignformfieldgroupTextvalidationcustommessage

`func (o *EzsignformfieldgroupResponseCompound) HasSEzsignformfieldgroupTextvalidationcustommessage() bool`

HasSEzsignformfieldgroupTextvalidationcustommessage returns a boolean if a field has been set.

### GetTEzsignformfieldgroupTooltip

`func (o *EzsignformfieldgroupResponseCompound) GetTEzsignformfieldgroupTooltip() string`

GetTEzsignformfieldgroupTooltip returns the TEzsignformfieldgroupTooltip field if non-nil, zero value otherwise.

### GetTEzsignformfieldgroupTooltipOk

`func (o *EzsignformfieldgroupResponseCompound) GetTEzsignformfieldgroupTooltipOk() (*string, bool)`

GetTEzsignformfieldgroupTooltipOk returns a tuple with the TEzsignformfieldgroupTooltip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEzsignformfieldgroupTooltip

`func (o *EzsignformfieldgroupResponseCompound) SetTEzsignformfieldgroupTooltip(v string)`

SetTEzsignformfieldgroupTooltip sets TEzsignformfieldgroupTooltip field to given value.

### HasTEzsignformfieldgroupTooltip

`func (o *EzsignformfieldgroupResponseCompound) HasTEzsignformfieldgroupTooltip() bool`

HasTEzsignformfieldgroupTooltip returns a boolean if a field has been set.

### GetEEzsignformfieldgroupTooltipposition

`func (o *EzsignformfieldgroupResponseCompound) GetEEzsignformfieldgroupTooltipposition() FieldEEzsignformfieldgroupTooltipposition`

GetEEzsignformfieldgroupTooltipposition returns the EEzsignformfieldgroupTooltipposition field if non-nil, zero value otherwise.

### GetEEzsignformfieldgroupTooltippositionOk

`func (o *EzsignformfieldgroupResponseCompound) GetEEzsignformfieldgroupTooltippositionOk() (*FieldEEzsignformfieldgroupTooltipposition, bool)`

GetEEzsignformfieldgroupTooltippositionOk returns a tuple with the EEzsignformfieldgroupTooltipposition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignformfieldgroupTooltipposition

`func (o *EzsignformfieldgroupResponseCompound) SetEEzsignformfieldgroupTooltipposition(v FieldEEzsignformfieldgroupTooltipposition)`

SetEEzsignformfieldgroupTooltipposition sets EEzsignformfieldgroupTooltipposition field to given value.

### HasEEzsignformfieldgroupTooltipposition

`func (o *EzsignformfieldgroupResponseCompound) HasEEzsignformfieldgroupTooltipposition() bool`

HasEEzsignformfieldgroupTooltipposition returns a boolean if a field has been set.

### GetAObjEzsignformfield

`func (o *EzsignformfieldgroupResponseCompound) GetAObjEzsignformfield() []EzsignformfieldResponseCompound`

GetAObjEzsignformfield returns the AObjEzsignformfield field if non-nil, zero value otherwise.

### GetAObjEzsignformfieldOk

`func (o *EzsignformfieldgroupResponseCompound) GetAObjEzsignformfieldOk() (*[]EzsignformfieldResponseCompound, bool)`

GetAObjEzsignformfieldOk returns a tuple with the AObjEzsignformfield field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzsignformfield

`func (o *EzsignformfieldgroupResponseCompound) SetAObjEzsignformfield(v []EzsignformfieldResponseCompound)`

SetAObjEzsignformfield sets AObjEzsignformfield field to given value.


### GetAObjDropdownElement

`func (o *EzsignformfieldgroupResponseCompound) GetAObjDropdownElement() []CustomDropdownElementResponseCompound`

GetAObjDropdownElement returns the AObjDropdownElement field if non-nil, zero value otherwise.

### GetAObjDropdownElementOk

`func (o *EzsignformfieldgroupResponseCompound) GetAObjDropdownElementOk() (*[]CustomDropdownElementResponseCompound, bool)`

GetAObjDropdownElementOk returns a tuple with the AObjDropdownElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjDropdownElement

`func (o *EzsignformfieldgroupResponseCompound) SetAObjDropdownElement(v []CustomDropdownElementResponseCompound)`

SetAObjDropdownElement sets AObjDropdownElement field to given value.

### HasAObjDropdownElement

`func (o *EzsignformfieldgroupResponseCompound) HasAObjDropdownElement() bool`

HasAObjDropdownElement returns a boolean if a field has been set.

### GetAObjEzsignformfieldgroupsigner

`func (o *EzsignformfieldgroupResponseCompound) GetAObjEzsignformfieldgroupsigner() []EzsignformfieldgroupsignerResponseCompound`

GetAObjEzsignformfieldgroupsigner returns the AObjEzsignformfieldgroupsigner field if non-nil, zero value otherwise.

### GetAObjEzsignformfieldgroupsignerOk

`func (o *EzsignformfieldgroupResponseCompound) GetAObjEzsignformfieldgroupsignerOk() (*[]EzsignformfieldgroupsignerResponseCompound, bool)`

GetAObjEzsignformfieldgroupsignerOk returns a tuple with the AObjEzsignformfieldgroupsigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzsignformfieldgroupsigner

`func (o *EzsignformfieldgroupResponseCompound) SetAObjEzsignformfieldgroupsigner(v []EzsignformfieldgroupsignerResponseCompound)`

SetAObjEzsignformfieldgroupsigner sets AObjEzsignformfieldgroupsigner field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


