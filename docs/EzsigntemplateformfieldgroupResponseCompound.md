# EzsigntemplateformfieldgroupResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsigntemplateformfieldgroupID** | **int32** | The unique ID of the Ezsigntemplateformfieldgroup | 
**FkiEzsigntemplatedocumentID** | **int32** | The unique ID of the Ezsigntemplatedocument | 
**EEzsigntemplateformfieldgroupType** | [**FieldEEzsigntemplateformfieldgroupType**](FieldEEzsigntemplateformfieldgroupType.md) |  | 
**EEzsigntemplateformfieldgroupSignerrequirement** | Pointer to [**FieldEEzsigntemplateformfieldgroupSignerrequirement**](FieldEEzsigntemplateformfieldgroupSignerrequirement.md) |  | [optional] 
**SEzsigntemplateformfieldgroupLabel** | **string** | The Label for the Ezsigntemplateformfieldgroup | 
**IEzsigntemplateformfieldgroupStep** | **int32** | The step when the Ezsigntemplatesigner will be invited to fill the form fields | 
**SEzsigntemplateformfieldgroupDefaultvalue** | Pointer to **string** | The default value for the Ezsigntemplateformfieldgroup  You can use the codes below and they will be replaced at signature time.    | Code | Description | Example | | ------------------------- | ------------ | ------------ | | {sUserFirstname} | The first name of the contact | John | | {sUserLastname} | The last name of the contact | Doe | | {sUserJobtitle} | The job title | Sales Representative | | {sEmailAddress} | The email address | email@example.com | | {sPhoneE164} | A phone number in E.164 Format | +15149901516 | | {sPhoneE164Cell} | A phone number in E.164 Format | +15149901516 | | [optional] 
**IEzsigntemplateformfieldgroupFilledmin** | **int32** | The minimum number of Ezsigntemplateformfield that must be filled in the Ezsigntemplateformfieldgroup | 
**IEzsigntemplateformfieldgroupFilledmax** | **int32** | The maximum number of Ezsigntemplateformfield that must be filled in the Ezsigntemplateformfieldgroup | 
**BEzsigntemplateformfieldgroupReadonly** | **bool** | Whether the Ezsigntemplateformfieldgroup is read only or not. | 
**IEzsigntemplateformfieldgroupMaxlength** | Pointer to **int32** | The maximum length for the value in the Ezsigntemplateformfieldgroup  This can only be set if eEzsigntemplateformfieldgroupType is **Text** or **Textarea** | [optional] 
**BEzsigntemplateformfieldgroupEncrypted** | Pointer to **bool** | Whether the Ezsigntemplateformfieldgroup is encrypted in the database or not. Encrypted values are not displayed on the Ezsigndocument. This can only be set if eEzsigntemplateformfieldgroupType is **Text** or **Textarea** | [optional] 
**SEzsigntemplateformfieldgroupRegexp** | Pointer to **string** | A regular expression to indicate what values are acceptable for the Ezsigntemplateformfieldgroup.  This can only be set if eEzsigntemplateformfieldgroupType is **Text** or **Textarea** | [optional] 
**SEzsigntemplateformfieldgroupTextvalidationcustommessage** | Pointer to **string** | Description of validation rule. Show by signatory. | [optional] 
**EEzsigntemplateformfieldgroupTextvalidation** | Pointer to [**EnumTextvalidation**](EnumTextvalidation.md) |  | [optional] 
**TEzsigntemplateformfieldgroupTooltip** | Pointer to **string** | A tooltip that will be presented to Ezsigntemplatesigner about the Ezsigntemplateformfieldgroup | [optional] 
**EEzsigntemplateformfieldgroupTooltipposition** | Pointer to [**FieldEEzsigntemplateformfieldgroupTooltipposition**](FieldEEzsigntemplateformfieldgroupTooltipposition.md) |  | [optional] 
**AObjEzsigntemplateformfieldgroupsigner** | [**[]EzsigntemplateformfieldgroupsignerResponseCompound**](EzsigntemplateformfieldgroupsignerResponseCompound.md) |  | 
**AObjDropdownElement** | Pointer to [**[]CustomDropdownElementResponseCompound**](CustomDropdownElementResponseCompound.md) |  | [optional] 
**AObjEzsigntemplateformfield** | [**[]EzsigntemplateformfieldResponseCompound**](EzsigntemplateformfieldResponseCompound.md) |  | 

## Methods

### NewEzsigntemplateformfieldgroupResponseCompound

`func NewEzsigntemplateformfieldgroupResponseCompound(pkiEzsigntemplateformfieldgroupID int32, fkiEzsigntemplatedocumentID int32, eEzsigntemplateformfieldgroupType FieldEEzsigntemplateformfieldgroupType, sEzsigntemplateformfieldgroupLabel string, iEzsigntemplateformfieldgroupStep int32, iEzsigntemplateformfieldgroupFilledmin int32, iEzsigntemplateformfieldgroupFilledmax int32, bEzsigntemplateformfieldgroupReadonly bool, aObjEzsigntemplateformfieldgroupsigner []EzsigntemplateformfieldgroupsignerResponseCompound, aObjEzsigntemplateformfield []EzsigntemplateformfieldResponseCompound, ) *EzsigntemplateformfieldgroupResponseCompound`

NewEzsigntemplateformfieldgroupResponseCompound instantiates a new EzsigntemplateformfieldgroupResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigntemplateformfieldgroupResponseCompoundWithDefaults

`func NewEzsigntemplateformfieldgroupResponseCompoundWithDefaults() *EzsigntemplateformfieldgroupResponseCompound`

NewEzsigntemplateformfieldgroupResponseCompoundWithDefaults instantiates a new EzsigntemplateformfieldgroupResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsigntemplateformfieldgroupID

`func (o *EzsigntemplateformfieldgroupResponseCompound) GetPkiEzsigntemplateformfieldgroupID() int32`

GetPkiEzsigntemplateformfieldgroupID returns the PkiEzsigntemplateformfieldgroupID field if non-nil, zero value otherwise.

### GetPkiEzsigntemplateformfieldgroupIDOk

`func (o *EzsigntemplateformfieldgroupResponseCompound) GetPkiEzsigntemplateformfieldgroupIDOk() (*int32, bool)`

GetPkiEzsigntemplateformfieldgroupIDOk returns a tuple with the PkiEzsigntemplateformfieldgroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsigntemplateformfieldgroupID

`func (o *EzsigntemplateformfieldgroupResponseCompound) SetPkiEzsigntemplateformfieldgroupID(v int32)`

SetPkiEzsigntemplateformfieldgroupID sets PkiEzsigntemplateformfieldgroupID field to given value.


### GetFkiEzsigntemplatedocumentID

`func (o *EzsigntemplateformfieldgroupResponseCompound) GetFkiEzsigntemplatedocumentID() int32`

GetFkiEzsigntemplatedocumentID returns the FkiEzsigntemplatedocumentID field if non-nil, zero value otherwise.

### GetFkiEzsigntemplatedocumentIDOk

`func (o *EzsigntemplateformfieldgroupResponseCompound) GetFkiEzsigntemplatedocumentIDOk() (*int32, bool)`

GetFkiEzsigntemplatedocumentIDOk returns a tuple with the FkiEzsigntemplatedocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntemplatedocumentID

`func (o *EzsigntemplateformfieldgroupResponseCompound) SetFkiEzsigntemplatedocumentID(v int32)`

SetFkiEzsigntemplatedocumentID sets FkiEzsigntemplatedocumentID field to given value.


### GetEEzsigntemplateformfieldgroupType

`func (o *EzsigntemplateformfieldgroupResponseCompound) GetEEzsigntemplateformfieldgroupType() FieldEEzsigntemplateformfieldgroupType`

GetEEzsigntemplateformfieldgroupType returns the EEzsigntemplateformfieldgroupType field if non-nil, zero value otherwise.

### GetEEzsigntemplateformfieldgroupTypeOk

`func (o *EzsigntemplateformfieldgroupResponseCompound) GetEEzsigntemplateformfieldgroupTypeOk() (*FieldEEzsigntemplateformfieldgroupType, bool)`

GetEEzsigntemplateformfieldgroupTypeOk returns a tuple with the EEzsigntemplateformfieldgroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplateformfieldgroupType

`func (o *EzsigntemplateformfieldgroupResponseCompound) SetEEzsigntemplateformfieldgroupType(v FieldEEzsigntemplateformfieldgroupType)`

SetEEzsigntemplateformfieldgroupType sets EEzsigntemplateformfieldgroupType field to given value.


### GetEEzsigntemplateformfieldgroupSignerrequirement

`func (o *EzsigntemplateformfieldgroupResponseCompound) GetEEzsigntemplateformfieldgroupSignerrequirement() FieldEEzsigntemplateformfieldgroupSignerrequirement`

GetEEzsigntemplateformfieldgroupSignerrequirement returns the EEzsigntemplateformfieldgroupSignerrequirement field if non-nil, zero value otherwise.

### GetEEzsigntemplateformfieldgroupSignerrequirementOk

`func (o *EzsigntemplateformfieldgroupResponseCompound) GetEEzsigntemplateformfieldgroupSignerrequirementOk() (*FieldEEzsigntemplateformfieldgroupSignerrequirement, bool)`

GetEEzsigntemplateformfieldgroupSignerrequirementOk returns a tuple with the EEzsigntemplateformfieldgroupSignerrequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplateformfieldgroupSignerrequirement

`func (o *EzsigntemplateformfieldgroupResponseCompound) SetEEzsigntemplateformfieldgroupSignerrequirement(v FieldEEzsigntemplateformfieldgroupSignerrequirement)`

SetEEzsigntemplateformfieldgroupSignerrequirement sets EEzsigntemplateformfieldgroupSignerrequirement field to given value.

### HasEEzsigntemplateformfieldgroupSignerrequirement

`func (o *EzsigntemplateformfieldgroupResponseCompound) HasEEzsigntemplateformfieldgroupSignerrequirement() bool`

HasEEzsigntemplateformfieldgroupSignerrequirement returns a boolean if a field has been set.

### GetSEzsigntemplateformfieldgroupLabel

`func (o *EzsigntemplateformfieldgroupResponseCompound) GetSEzsigntemplateformfieldgroupLabel() string`

GetSEzsigntemplateformfieldgroupLabel returns the SEzsigntemplateformfieldgroupLabel field if non-nil, zero value otherwise.

### GetSEzsigntemplateformfieldgroupLabelOk

`func (o *EzsigntemplateformfieldgroupResponseCompound) GetSEzsigntemplateformfieldgroupLabelOk() (*string, bool)`

GetSEzsigntemplateformfieldgroupLabelOk returns a tuple with the SEzsigntemplateformfieldgroupLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplateformfieldgroupLabel

`func (o *EzsigntemplateformfieldgroupResponseCompound) SetSEzsigntemplateformfieldgroupLabel(v string)`

SetSEzsigntemplateformfieldgroupLabel sets SEzsigntemplateformfieldgroupLabel field to given value.


### GetIEzsigntemplateformfieldgroupStep

`func (o *EzsigntemplateformfieldgroupResponseCompound) GetIEzsigntemplateformfieldgroupStep() int32`

GetIEzsigntemplateformfieldgroupStep returns the IEzsigntemplateformfieldgroupStep field if non-nil, zero value otherwise.

### GetIEzsigntemplateformfieldgroupStepOk

`func (o *EzsigntemplateformfieldgroupResponseCompound) GetIEzsigntemplateformfieldgroupStepOk() (*int32, bool)`

GetIEzsigntemplateformfieldgroupStepOk returns a tuple with the IEzsigntemplateformfieldgroupStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplateformfieldgroupStep

`func (o *EzsigntemplateformfieldgroupResponseCompound) SetIEzsigntemplateformfieldgroupStep(v int32)`

SetIEzsigntemplateformfieldgroupStep sets IEzsigntemplateformfieldgroupStep field to given value.


### GetSEzsigntemplateformfieldgroupDefaultvalue

`func (o *EzsigntemplateformfieldgroupResponseCompound) GetSEzsigntemplateformfieldgroupDefaultvalue() string`

GetSEzsigntemplateformfieldgroupDefaultvalue returns the SEzsigntemplateformfieldgroupDefaultvalue field if non-nil, zero value otherwise.

### GetSEzsigntemplateformfieldgroupDefaultvalueOk

`func (o *EzsigntemplateformfieldgroupResponseCompound) GetSEzsigntemplateformfieldgroupDefaultvalueOk() (*string, bool)`

GetSEzsigntemplateformfieldgroupDefaultvalueOk returns a tuple with the SEzsigntemplateformfieldgroupDefaultvalue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplateformfieldgroupDefaultvalue

`func (o *EzsigntemplateformfieldgroupResponseCompound) SetSEzsigntemplateformfieldgroupDefaultvalue(v string)`

SetSEzsigntemplateformfieldgroupDefaultvalue sets SEzsigntemplateformfieldgroupDefaultvalue field to given value.

### HasSEzsigntemplateformfieldgroupDefaultvalue

`func (o *EzsigntemplateformfieldgroupResponseCompound) HasSEzsigntemplateformfieldgroupDefaultvalue() bool`

HasSEzsigntemplateformfieldgroupDefaultvalue returns a boolean if a field has been set.

### GetIEzsigntemplateformfieldgroupFilledmin

`func (o *EzsigntemplateformfieldgroupResponseCompound) GetIEzsigntemplateformfieldgroupFilledmin() int32`

GetIEzsigntemplateformfieldgroupFilledmin returns the IEzsigntemplateformfieldgroupFilledmin field if non-nil, zero value otherwise.

### GetIEzsigntemplateformfieldgroupFilledminOk

`func (o *EzsigntemplateformfieldgroupResponseCompound) GetIEzsigntemplateformfieldgroupFilledminOk() (*int32, bool)`

GetIEzsigntemplateformfieldgroupFilledminOk returns a tuple with the IEzsigntemplateformfieldgroupFilledmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplateformfieldgroupFilledmin

`func (o *EzsigntemplateformfieldgroupResponseCompound) SetIEzsigntemplateformfieldgroupFilledmin(v int32)`

SetIEzsigntemplateformfieldgroupFilledmin sets IEzsigntemplateformfieldgroupFilledmin field to given value.


### GetIEzsigntemplateformfieldgroupFilledmax

`func (o *EzsigntemplateformfieldgroupResponseCompound) GetIEzsigntemplateformfieldgroupFilledmax() int32`

GetIEzsigntemplateformfieldgroupFilledmax returns the IEzsigntemplateformfieldgroupFilledmax field if non-nil, zero value otherwise.

### GetIEzsigntemplateformfieldgroupFilledmaxOk

`func (o *EzsigntemplateformfieldgroupResponseCompound) GetIEzsigntemplateformfieldgroupFilledmaxOk() (*int32, bool)`

GetIEzsigntemplateformfieldgroupFilledmaxOk returns a tuple with the IEzsigntemplateformfieldgroupFilledmax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplateformfieldgroupFilledmax

`func (o *EzsigntemplateformfieldgroupResponseCompound) SetIEzsigntemplateformfieldgroupFilledmax(v int32)`

SetIEzsigntemplateformfieldgroupFilledmax sets IEzsigntemplateformfieldgroupFilledmax field to given value.


### GetBEzsigntemplateformfieldgroupReadonly

`func (o *EzsigntemplateformfieldgroupResponseCompound) GetBEzsigntemplateformfieldgroupReadonly() bool`

GetBEzsigntemplateformfieldgroupReadonly returns the BEzsigntemplateformfieldgroupReadonly field if non-nil, zero value otherwise.

### GetBEzsigntemplateformfieldgroupReadonlyOk

`func (o *EzsigntemplateformfieldgroupResponseCompound) GetBEzsigntemplateformfieldgroupReadonlyOk() (*bool, bool)`

GetBEzsigntemplateformfieldgroupReadonlyOk returns a tuple with the BEzsigntemplateformfieldgroupReadonly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigntemplateformfieldgroupReadonly

`func (o *EzsigntemplateformfieldgroupResponseCompound) SetBEzsigntemplateformfieldgroupReadonly(v bool)`

SetBEzsigntemplateformfieldgroupReadonly sets BEzsigntemplateformfieldgroupReadonly field to given value.


### GetIEzsigntemplateformfieldgroupMaxlength

`func (o *EzsigntemplateformfieldgroupResponseCompound) GetIEzsigntemplateformfieldgroupMaxlength() int32`

GetIEzsigntemplateformfieldgroupMaxlength returns the IEzsigntemplateformfieldgroupMaxlength field if non-nil, zero value otherwise.

### GetIEzsigntemplateformfieldgroupMaxlengthOk

`func (o *EzsigntemplateformfieldgroupResponseCompound) GetIEzsigntemplateformfieldgroupMaxlengthOk() (*int32, bool)`

GetIEzsigntemplateformfieldgroupMaxlengthOk returns a tuple with the IEzsigntemplateformfieldgroupMaxlength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplateformfieldgroupMaxlength

`func (o *EzsigntemplateformfieldgroupResponseCompound) SetIEzsigntemplateformfieldgroupMaxlength(v int32)`

SetIEzsigntemplateformfieldgroupMaxlength sets IEzsigntemplateformfieldgroupMaxlength field to given value.

### HasIEzsigntemplateformfieldgroupMaxlength

`func (o *EzsigntemplateformfieldgroupResponseCompound) HasIEzsigntemplateformfieldgroupMaxlength() bool`

HasIEzsigntemplateformfieldgroupMaxlength returns a boolean if a field has been set.

### GetBEzsigntemplateformfieldgroupEncrypted

`func (o *EzsigntemplateformfieldgroupResponseCompound) GetBEzsigntemplateformfieldgroupEncrypted() bool`

GetBEzsigntemplateformfieldgroupEncrypted returns the BEzsigntemplateformfieldgroupEncrypted field if non-nil, zero value otherwise.

### GetBEzsigntemplateformfieldgroupEncryptedOk

`func (o *EzsigntemplateformfieldgroupResponseCompound) GetBEzsigntemplateformfieldgroupEncryptedOk() (*bool, bool)`

GetBEzsigntemplateformfieldgroupEncryptedOk returns a tuple with the BEzsigntemplateformfieldgroupEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigntemplateformfieldgroupEncrypted

`func (o *EzsigntemplateformfieldgroupResponseCompound) SetBEzsigntemplateformfieldgroupEncrypted(v bool)`

SetBEzsigntemplateformfieldgroupEncrypted sets BEzsigntemplateformfieldgroupEncrypted field to given value.

### HasBEzsigntemplateformfieldgroupEncrypted

`func (o *EzsigntemplateformfieldgroupResponseCompound) HasBEzsigntemplateformfieldgroupEncrypted() bool`

HasBEzsigntemplateformfieldgroupEncrypted returns a boolean if a field has been set.

### GetSEzsigntemplateformfieldgroupRegexp

`func (o *EzsigntemplateformfieldgroupResponseCompound) GetSEzsigntemplateformfieldgroupRegexp() string`

GetSEzsigntemplateformfieldgroupRegexp returns the SEzsigntemplateformfieldgroupRegexp field if non-nil, zero value otherwise.

### GetSEzsigntemplateformfieldgroupRegexpOk

`func (o *EzsigntemplateformfieldgroupResponseCompound) GetSEzsigntemplateformfieldgroupRegexpOk() (*string, bool)`

GetSEzsigntemplateformfieldgroupRegexpOk returns a tuple with the SEzsigntemplateformfieldgroupRegexp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplateformfieldgroupRegexp

`func (o *EzsigntemplateformfieldgroupResponseCompound) SetSEzsigntemplateformfieldgroupRegexp(v string)`

SetSEzsigntemplateformfieldgroupRegexp sets SEzsigntemplateformfieldgroupRegexp field to given value.

### HasSEzsigntemplateformfieldgroupRegexp

`func (o *EzsigntemplateformfieldgroupResponseCompound) HasSEzsigntemplateformfieldgroupRegexp() bool`

HasSEzsigntemplateformfieldgroupRegexp returns a boolean if a field has been set.

### GetSEzsigntemplateformfieldgroupTextvalidationcustommessage

`func (o *EzsigntemplateformfieldgroupResponseCompound) GetSEzsigntemplateformfieldgroupTextvalidationcustommessage() string`

GetSEzsigntemplateformfieldgroupTextvalidationcustommessage returns the SEzsigntemplateformfieldgroupTextvalidationcustommessage field if non-nil, zero value otherwise.

### GetSEzsigntemplateformfieldgroupTextvalidationcustommessageOk

`func (o *EzsigntemplateformfieldgroupResponseCompound) GetSEzsigntemplateformfieldgroupTextvalidationcustommessageOk() (*string, bool)`

GetSEzsigntemplateformfieldgroupTextvalidationcustommessageOk returns a tuple with the SEzsigntemplateformfieldgroupTextvalidationcustommessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplateformfieldgroupTextvalidationcustommessage

`func (o *EzsigntemplateformfieldgroupResponseCompound) SetSEzsigntemplateformfieldgroupTextvalidationcustommessage(v string)`

SetSEzsigntemplateformfieldgroupTextvalidationcustommessage sets SEzsigntemplateformfieldgroupTextvalidationcustommessage field to given value.

### HasSEzsigntemplateformfieldgroupTextvalidationcustommessage

`func (o *EzsigntemplateformfieldgroupResponseCompound) HasSEzsigntemplateformfieldgroupTextvalidationcustommessage() bool`

HasSEzsigntemplateformfieldgroupTextvalidationcustommessage returns a boolean if a field has been set.

### GetEEzsigntemplateformfieldgroupTextvalidation

`func (o *EzsigntemplateformfieldgroupResponseCompound) GetEEzsigntemplateformfieldgroupTextvalidation() EnumTextvalidation`

GetEEzsigntemplateformfieldgroupTextvalidation returns the EEzsigntemplateformfieldgroupTextvalidation field if non-nil, zero value otherwise.

### GetEEzsigntemplateformfieldgroupTextvalidationOk

`func (o *EzsigntemplateformfieldgroupResponseCompound) GetEEzsigntemplateformfieldgroupTextvalidationOk() (*EnumTextvalidation, bool)`

GetEEzsigntemplateformfieldgroupTextvalidationOk returns a tuple with the EEzsigntemplateformfieldgroupTextvalidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplateformfieldgroupTextvalidation

`func (o *EzsigntemplateformfieldgroupResponseCompound) SetEEzsigntemplateformfieldgroupTextvalidation(v EnumTextvalidation)`

SetEEzsigntemplateformfieldgroupTextvalidation sets EEzsigntemplateformfieldgroupTextvalidation field to given value.

### HasEEzsigntemplateformfieldgroupTextvalidation

`func (o *EzsigntemplateformfieldgroupResponseCompound) HasEEzsigntemplateformfieldgroupTextvalidation() bool`

HasEEzsigntemplateformfieldgroupTextvalidation returns a boolean if a field has been set.

### GetTEzsigntemplateformfieldgroupTooltip

`func (o *EzsigntemplateformfieldgroupResponseCompound) GetTEzsigntemplateformfieldgroupTooltip() string`

GetTEzsigntemplateformfieldgroupTooltip returns the TEzsigntemplateformfieldgroupTooltip field if non-nil, zero value otherwise.

### GetTEzsigntemplateformfieldgroupTooltipOk

`func (o *EzsigntemplateformfieldgroupResponseCompound) GetTEzsigntemplateformfieldgroupTooltipOk() (*string, bool)`

GetTEzsigntemplateformfieldgroupTooltipOk returns a tuple with the TEzsigntemplateformfieldgroupTooltip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEzsigntemplateformfieldgroupTooltip

`func (o *EzsigntemplateformfieldgroupResponseCompound) SetTEzsigntemplateformfieldgroupTooltip(v string)`

SetTEzsigntemplateformfieldgroupTooltip sets TEzsigntemplateformfieldgroupTooltip field to given value.

### HasTEzsigntemplateformfieldgroupTooltip

`func (o *EzsigntemplateformfieldgroupResponseCompound) HasTEzsigntemplateformfieldgroupTooltip() bool`

HasTEzsigntemplateformfieldgroupTooltip returns a boolean if a field has been set.

### GetEEzsigntemplateformfieldgroupTooltipposition

`func (o *EzsigntemplateformfieldgroupResponseCompound) GetEEzsigntemplateformfieldgroupTooltipposition() FieldEEzsigntemplateformfieldgroupTooltipposition`

GetEEzsigntemplateformfieldgroupTooltipposition returns the EEzsigntemplateformfieldgroupTooltipposition field if non-nil, zero value otherwise.

### GetEEzsigntemplateformfieldgroupTooltippositionOk

`func (o *EzsigntemplateformfieldgroupResponseCompound) GetEEzsigntemplateformfieldgroupTooltippositionOk() (*FieldEEzsigntemplateformfieldgroupTooltipposition, bool)`

GetEEzsigntemplateformfieldgroupTooltippositionOk returns a tuple with the EEzsigntemplateformfieldgroupTooltipposition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplateformfieldgroupTooltipposition

`func (o *EzsigntemplateformfieldgroupResponseCompound) SetEEzsigntemplateformfieldgroupTooltipposition(v FieldEEzsigntemplateformfieldgroupTooltipposition)`

SetEEzsigntemplateformfieldgroupTooltipposition sets EEzsigntemplateformfieldgroupTooltipposition field to given value.

### HasEEzsigntemplateformfieldgroupTooltipposition

`func (o *EzsigntemplateformfieldgroupResponseCompound) HasEEzsigntemplateformfieldgroupTooltipposition() bool`

HasEEzsigntemplateformfieldgroupTooltipposition returns a boolean if a field has been set.

### GetAObjEzsigntemplateformfieldgroupsigner

`func (o *EzsigntemplateformfieldgroupResponseCompound) GetAObjEzsigntemplateformfieldgroupsigner() []EzsigntemplateformfieldgroupsignerResponseCompound`

GetAObjEzsigntemplateformfieldgroupsigner returns the AObjEzsigntemplateformfieldgroupsigner field if non-nil, zero value otherwise.

### GetAObjEzsigntemplateformfieldgroupsignerOk

`func (o *EzsigntemplateformfieldgroupResponseCompound) GetAObjEzsigntemplateformfieldgroupsignerOk() (*[]EzsigntemplateformfieldgroupsignerResponseCompound, bool)`

GetAObjEzsigntemplateformfieldgroupsignerOk returns a tuple with the AObjEzsigntemplateformfieldgroupsigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzsigntemplateformfieldgroupsigner

`func (o *EzsigntemplateformfieldgroupResponseCompound) SetAObjEzsigntemplateformfieldgroupsigner(v []EzsigntemplateformfieldgroupsignerResponseCompound)`

SetAObjEzsigntemplateformfieldgroupsigner sets AObjEzsigntemplateformfieldgroupsigner field to given value.


### GetAObjDropdownElement

`func (o *EzsigntemplateformfieldgroupResponseCompound) GetAObjDropdownElement() []CustomDropdownElementResponseCompound`

GetAObjDropdownElement returns the AObjDropdownElement field if non-nil, zero value otherwise.

### GetAObjDropdownElementOk

`func (o *EzsigntemplateformfieldgroupResponseCompound) GetAObjDropdownElementOk() (*[]CustomDropdownElementResponseCompound, bool)`

GetAObjDropdownElementOk returns a tuple with the AObjDropdownElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjDropdownElement

`func (o *EzsigntemplateformfieldgroupResponseCompound) SetAObjDropdownElement(v []CustomDropdownElementResponseCompound)`

SetAObjDropdownElement sets AObjDropdownElement field to given value.

### HasAObjDropdownElement

`func (o *EzsigntemplateformfieldgroupResponseCompound) HasAObjDropdownElement() bool`

HasAObjDropdownElement returns a boolean if a field has been set.

### GetAObjEzsigntemplateformfield

`func (o *EzsigntemplateformfieldgroupResponseCompound) GetAObjEzsigntemplateformfield() []EzsigntemplateformfieldResponseCompound`

GetAObjEzsigntemplateformfield returns the AObjEzsigntemplateformfield field if non-nil, zero value otherwise.

### GetAObjEzsigntemplateformfieldOk

`func (o *EzsigntemplateformfieldgroupResponseCompound) GetAObjEzsigntemplateformfieldOk() (*[]EzsigntemplateformfieldResponseCompound, bool)`

GetAObjEzsigntemplateformfieldOk returns a tuple with the AObjEzsigntemplateformfield field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzsigntemplateformfield

`func (o *EzsigntemplateformfieldgroupResponseCompound) SetAObjEzsigntemplateformfield(v []EzsigntemplateformfieldResponseCompound)`

SetAObjEzsigntemplateformfield sets AObjEzsigntemplateformfield field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


