# EzsigntemplateformfieldgroupRequestCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsigntemplateformfieldgroupID** | Pointer to **int32** | The unique ID of the Ezsigntemplateformfieldgroup | [optional] 
**FkiEzsigntemplatedocumentID** | **int32** | The unique ID of the Ezsigntemplatedocument | 
**EEzsigntemplateformfieldgroupType** | [**FieldEEzsigntemplateformfieldgroupType**](FieldEEzsigntemplateformfieldgroupType.md) |  | 
**EEzsigntemplateformfieldgroupSignerrequirement** | Pointer to [**FieldEEzsigntemplateformfieldgroupSignerrequirement**](FieldEEzsigntemplateformfieldgroupSignerrequirement.md) |  | [optional] 
**SEzsigntemplateformfieldgroupLabel** | **string** | The Label for the Ezsigntemplateformfieldgroup | 
**IEzsigntemplateformfieldgroupStep** | **int32** | The step when the Ezsigntemplatesigner will be invited to fill the form fields | 
**SEzsigntemplateformfieldgroupDefaultvalue** | **string** | The default value for the Ezsigntemplateformfieldgroup  You can use the codes below and they will be replaced at signature time.    | Code | Description | Example | | ------------------------- | ------------ | ------------ | | {sUserFirstname} | The first name of the contact | John | | {sUserLastname} | The last name of the contact | Doe | | {sUserJobtitle} | The job title | Sales Representative | | {sEmailAddress} | The email address | email@example.com | | {sPhoneE164} | A phone number in E.164 Format | +15149901516 | | {sPhoneE164Cell} | A phone number in E.164 Format | +15149901516 | | 
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
**AObjEzsigntemplateformfieldgroupsigner** | [**[]EzsigntemplateformfieldgroupsignerRequestCompound**](EzsigntemplateformfieldgroupsignerRequestCompound.md) |  | 
**AObjDropdownElement** | Pointer to [**[]CustomDropdownElementRequestCompound**](CustomDropdownElementRequestCompound.md) |  | [optional] 
**AObjEzsigntemplateformfield** | [**[]EzsigntemplateformfieldRequestCompound**](EzsigntemplateformfieldRequestCompound.md) |  | 

## Methods

### NewEzsigntemplateformfieldgroupRequestCompound

`func NewEzsigntemplateformfieldgroupRequestCompound(fkiEzsigntemplatedocumentID int32, eEzsigntemplateformfieldgroupType FieldEEzsigntemplateformfieldgroupType, sEzsigntemplateformfieldgroupLabel string, iEzsigntemplateformfieldgroupStep int32, sEzsigntemplateformfieldgroupDefaultvalue string, iEzsigntemplateformfieldgroupFilledmin int32, iEzsigntemplateformfieldgroupFilledmax int32, bEzsigntemplateformfieldgroupReadonly bool, aObjEzsigntemplateformfieldgroupsigner []EzsigntemplateformfieldgroupsignerRequestCompound, aObjEzsigntemplateformfield []EzsigntemplateformfieldRequestCompound, ) *EzsigntemplateformfieldgroupRequestCompound`

NewEzsigntemplateformfieldgroupRequestCompound instantiates a new EzsigntemplateformfieldgroupRequestCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigntemplateformfieldgroupRequestCompoundWithDefaults

`func NewEzsigntemplateformfieldgroupRequestCompoundWithDefaults() *EzsigntemplateformfieldgroupRequestCompound`

NewEzsigntemplateformfieldgroupRequestCompoundWithDefaults instantiates a new EzsigntemplateformfieldgroupRequestCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsigntemplateformfieldgroupID

`func (o *EzsigntemplateformfieldgroupRequestCompound) GetPkiEzsigntemplateformfieldgroupID() int32`

GetPkiEzsigntemplateformfieldgroupID returns the PkiEzsigntemplateformfieldgroupID field if non-nil, zero value otherwise.

### GetPkiEzsigntemplateformfieldgroupIDOk

`func (o *EzsigntemplateformfieldgroupRequestCompound) GetPkiEzsigntemplateformfieldgroupIDOk() (*int32, bool)`

GetPkiEzsigntemplateformfieldgroupIDOk returns a tuple with the PkiEzsigntemplateformfieldgroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsigntemplateformfieldgroupID

`func (o *EzsigntemplateformfieldgroupRequestCompound) SetPkiEzsigntemplateformfieldgroupID(v int32)`

SetPkiEzsigntemplateformfieldgroupID sets PkiEzsigntemplateformfieldgroupID field to given value.

### HasPkiEzsigntemplateformfieldgroupID

`func (o *EzsigntemplateformfieldgroupRequestCompound) HasPkiEzsigntemplateformfieldgroupID() bool`

HasPkiEzsigntemplateformfieldgroupID returns a boolean if a field has been set.

### GetFkiEzsigntemplatedocumentID

`func (o *EzsigntemplateformfieldgroupRequestCompound) GetFkiEzsigntemplatedocumentID() int32`

GetFkiEzsigntemplatedocumentID returns the FkiEzsigntemplatedocumentID field if non-nil, zero value otherwise.

### GetFkiEzsigntemplatedocumentIDOk

`func (o *EzsigntemplateformfieldgroupRequestCompound) GetFkiEzsigntemplatedocumentIDOk() (*int32, bool)`

GetFkiEzsigntemplatedocumentIDOk returns a tuple with the FkiEzsigntemplatedocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntemplatedocumentID

`func (o *EzsigntemplateformfieldgroupRequestCompound) SetFkiEzsigntemplatedocumentID(v int32)`

SetFkiEzsigntemplatedocumentID sets FkiEzsigntemplatedocumentID field to given value.


### GetEEzsigntemplateformfieldgroupType

`func (o *EzsigntemplateformfieldgroupRequestCompound) GetEEzsigntemplateformfieldgroupType() FieldEEzsigntemplateformfieldgroupType`

GetEEzsigntemplateformfieldgroupType returns the EEzsigntemplateformfieldgroupType field if non-nil, zero value otherwise.

### GetEEzsigntemplateformfieldgroupTypeOk

`func (o *EzsigntemplateformfieldgroupRequestCompound) GetEEzsigntemplateformfieldgroupTypeOk() (*FieldEEzsigntemplateformfieldgroupType, bool)`

GetEEzsigntemplateformfieldgroupTypeOk returns a tuple with the EEzsigntemplateformfieldgroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplateformfieldgroupType

`func (o *EzsigntemplateformfieldgroupRequestCompound) SetEEzsigntemplateformfieldgroupType(v FieldEEzsigntemplateformfieldgroupType)`

SetEEzsigntemplateformfieldgroupType sets EEzsigntemplateformfieldgroupType field to given value.


### GetEEzsigntemplateformfieldgroupSignerrequirement

`func (o *EzsigntemplateformfieldgroupRequestCompound) GetEEzsigntemplateformfieldgroupSignerrequirement() FieldEEzsigntemplateformfieldgroupSignerrequirement`

GetEEzsigntemplateformfieldgroupSignerrequirement returns the EEzsigntemplateformfieldgroupSignerrequirement field if non-nil, zero value otherwise.

### GetEEzsigntemplateformfieldgroupSignerrequirementOk

`func (o *EzsigntemplateformfieldgroupRequestCompound) GetEEzsigntemplateformfieldgroupSignerrequirementOk() (*FieldEEzsigntemplateformfieldgroupSignerrequirement, bool)`

GetEEzsigntemplateformfieldgroupSignerrequirementOk returns a tuple with the EEzsigntemplateformfieldgroupSignerrequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplateformfieldgroupSignerrequirement

`func (o *EzsigntemplateformfieldgroupRequestCompound) SetEEzsigntemplateformfieldgroupSignerrequirement(v FieldEEzsigntemplateformfieldgroupSignerrequirement)`

SetEEzsigntemplateformfieldgroupSignerrequirement sets EEzsigntemplateformfieldgroupSignerrequirement field to given value.

### HasEEzsigntemplateformfieldgroupSignerrequirement

`func (o *EzsigntemplateformfieldgroupRequestCompound) HasEEzsigntemplateformfieldgroupSignerrequirement() bool`

HasEEzsigntemplateformfieldgroupSignerrequirement returns a boolean if a field has been set.

### GetSEzsigntemplateformfieldgroupLabel

`func (o *EzsigntemplateformfieldgroupRequestCompound) GetSEzsigntemplateformfieldgroupLabel() string`

GetSEzsigntemplateformfieldgroupLabel returns the SEzsigntemplateformfieldgroupLabel field if non-nil, zero value otherwise.

### GetSEzsigntemplateformfieldgroupLabelOk

`func (o *EzsigntemplateformfieldgroupRequestCompound) GetSEzsigntemplateformfieldgroupLabelOk() (*string, bool)`

GetSEzsigntemplateformfieldgroupLabelOk returns a tuple with the SEzsigntemplateformfieldgroupLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplateformfieldgroupLabel

`func (o *EzsigntemplateformfieldgroupRequestCompound) SetSEzsigntemplateformfieldgroupLabel(v string)`

SetSEzsigntemplateformfieldgroupLabel sets SEzsigntemplateformfieldgroupLabel field to given value.


### GetIEzsigntemplateformfieldgroupStep

`func (o *EzsigntemplateformfieldgroupRequestCompound) GetIEzsigntemplateformfieldgroupStep() int32`

GetIEzsigntemplateformfieldgroupStep returns the IEzsigntemplateformfieldgroupStep field if non-nil, zero value otherwise.

### GetIEzsigntemplateformfieldgroupStepOk

`func (o *EzsigntemplateformfieldgroupRequestCompound) GetIEzsigntemplateformfieldgroupStepOk() (*int32, bool)`

GetIEzsigntemplateformfieldgroupStepOk returns a tuple with the IEzsigntemplateformfieldgroupStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplateformfieldgroupStep

`func (o *EzsigntemplateformfieldgroupRequestCompound) SetIEzsigntemplateformfieldgroupStep(v int32)`

SetIEzsigntemplateformfieldgroupStep sets IEzsigntemplateformfieldgroupStep field to given value.


### GetSEzsigntemplateformfieldgroupDefaultvalue

`func (o *EzsigntemplateformfieldgroupRequestCompound) GetSEzsigntemplateformfieldgroupDefaultvalue() string`

GetSEzsigntemplateformfieldgroupDefaultvalue returns the SEzsigntemplateformfieldgroupDefaultvalue field if non-nil, zero value otherwise.

### GetSEzsigntemplateformfieldgroupDefaultvalueOk

`func (o *EzsigntemplateformfieldgroupRequestCompound) GetSEzsigntemplateformfieldgroupDefaultvalueOk() (*string, bool)`

GetSEzsigntemplateformfieldgroupDefaultvalueOk returns a tuple with the SEzsigntemplateformfieldgroupDefaultvalue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplateformfieldgroupDefaultvalue

`func (o *EzsigntemplateformfieldgroupRequestCompound) SetSEzsigntemplateformfieldgroupDefaultvalue(v string)`

SetSEzsigntemplateformfieldgroupDefaultvalue sets SEzsigntemplateformfieldgroupDefaultvalue field to given value.


### GetIEzsigntemplateformfieldgroupFilledmin

`func (o *EzsigntemplateformfieldgroupRequestCompound) GetIEzsigntemplateformfieldgroupFilledmin() int32`

GetIEzsigntemplateformfieldgroupFilledmin returns the IEzsigntemplateformfieldgroupFilledmin field if non-nil, zero value otherwise.

### GetIEzsigntemplateformfieldgroupFilledminOk

`func (o *EzsigntemplateformfieldgroupRequestCompound) GetIEzsigntemplateformfieldgroupFilledminOk() (*int32, bool)`

GetIEzsigntemplateformfieldgroupFilledminOk returns a tuple with the IEzsigntemplateformfieldgroupFilledmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplateformfieldgroupFilledmin

`func (o *EzsigntemplateformfieldgroupRequestCompound) SetIEzsigntemplateformfieldgroupFilledmin(v int32)`

SetIEzsigntemplateformfieldgroupFilledmin sets IEzsigntemplateformfieldgroupFilledmin field to given value.


### GetIEzsigntemplateformfieldgroupFilledmax

`func (o *EzsigntemplateformfieldgroupRequestCompound) GetIEzsigntemplateformfieldgroupFilledmax() int32`

GetIEzsigntemplateformfieldgroupFilledmax returns the IEzsigntemplateformfieldgroupFilledmax field if non-nil, zero value otherwise.

### GetIEzsigntemplateformfieldgroupFilledmaxOk

`func (o *EzsigntemplateformfieldgroupRequestCompound) GetIEzsigntemplateformfieldgroupFilledmaxOk() (*int32, bool)`

GetIEzsigntemplateformfieldgroupFilledmaxOk returns a tuple with the IEzsigntemplateformfieldgroupFilledmax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplateformfieldgroupFilledmax

`func (o *EzsigntemplateformfieldgroupRequestCompound) SetIEzsigntemplateformfieldgroupFilledmax(v int32)`

SetIEzsigntemplateformfieldgroupFilledmax sets IEzsigntemplateformfieldgroupFilledmax field to given value.


### GetBEzsigntemplateformfieldgroupReadonly

`func (o *EzsigntemplateformfieldgroupRequestCompound) GetBEzsigntemplateformfieldgroupReadonly() bool`

GetBEzsigntemplateformfieldgroupReadonly returns the BEzsigntemplateformfieldgroupReadonly field if non-nil, zero value otherwise.

### GetBEzsigntemplateformfieldgroupReadonlyOk

`func (o *EzsigntemplateformfieldgroupRequestCompound) GetBEzsigntemplateformfieldgroupReadonlyOk() (*bool, bool)`

GetBEzsigntemplateformfieldgroupReadonlyOk returns a tuple with the BEzsigntemplateformfieldgroupReadonly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigntemplateformfieldgroupReadonly

`func (o *EzsigntemplateformfieldgroupRequestCompound) SetBEzsigntemplateformfieldgroupReadonly(v bool)`

SetBEzsigntemplateformfieldgroupReadonly sets BEzsigntemplateformfieldgroupReadonly field to given value.


### GetIEzsigntemplateformfieldgroupMaxlength

`func (o *EzsigntemplateformfieldgroupRequestCompound) GetIEzsigntemplateformfieldgroupMaxlength() int32`

GetIEzsigntemplateformfieldgroupMaxlength returns the IEzsigntemplateformfieldgroupMaxlength field if non-nil, zero value otherwise.

### GetIEzsigntemplateformfieldgroupMaxlengthOk

`func (o *EzsigntemplateformfieldgroupRequestCompound) GetIEzsigntemplateformfieldgroupMaxlengthOk() (*int32, bool)`

GetIEzsigntemplateformfieldgroupMaxlengthOk returns a tuple with the IEzsigntemplateformfieldgroupMaxlength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplateformfieldgroupMaxlength

`func (o *EzsigntemplateformfieldgroupRequestCompound) SetIEzsigntemplateformfieldgroupMaxlength(v int32)`

SetIEzsigntemplateformfieldgroupMaxlength sets IEzsigntemplateformfieldgroupMaxlength field to given value.

### HasIEzsigntemplateformfieldgroupMaxlength

`func (o *EzsigntemplateformfieldgroupRequestCompound) HasIEzsigntemplateformfieldgroupMaxlength() bool`

HasIEzsigntemplateformfieldgroupMaxlength returns a boolean if a field has been set.

### GetBEzsigntemplateformfieldgroupEncrypted

`func (o *EzsigntemplateformfieldgroupRequestCompound) GetBEzsigntemplateformfieldgroupEncrypted() bool`

GetBEzsigntemplateformfieldgroupEncrypted returns the BEzsigntemplateformfieldgroupEncrypted field if non-nil, zero value otherwise.

### GetBEzsigntemplateformfieldgroupEncryptedOk

`func (o *EzsigntemplateformfieldgroupRequestCompound) GetBEzsigntemplateformfieldgroupEncryptedOk() (*bool, bool)`

GetBEzsigntemplateformfieldgroupEncryptedOk returns a tuple with the BEzsigntemplateformfieldgroupEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigntemplateformfieldgroupEncrypted

`func (o *EzsigntemplateformfieldgroupRequestCompound) SetBEzsigntemplateformfieldgroupEncrypted(v bool)`

SetBEzsigntemplateformfieldgroupEncrypted sets BEzsigntemplateformfieldgroupEncrypted field to given value.

### HasBEzsigntemplateformfieldgroupEncrypted

`func (o *EzsigntemplateformfieldgroupRequestCompound) HasBEzsigntemplateformfieldgroupEncrypted() bool`

HasBEzsigntemplateformfieldgroupEncrypted returns a boolean if a field has been set.

### GetSEzsigntemplateformfieldgroupRegexp

`func (o *EzsigntemplateformfieldgroupRequestCompound) GetSEzsigntemplateformfieldgroupRegexp() string`

GetSEzsigntemplateformfieldgroupRegexp returns the SEzsigntemplateformfieldgroupRegexp field if non-nil, zero value otherwise.

### GetSEzsigntemplateformfieldgroupRegexpOk

`func (o *EzsigntemplateformfieldgroupRequestCompound) GetSEzsigntemplateformfieldgroupRegexpOk() (*string, bool)`

GetSEzsigntemplateformfieldgroupRegexpOk returns a tuple with the SEzsigntemplateformfieldgroupRegexp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplateformfieldgroupRegexp

`func (o *EzsigntemplateformfieldgroupRequestCompound) SetSEzsigntemplateformfieldgroupRegexp(v string)`

SetSEzsigntemplateformfieldgroupRegexp sets SEzsigntemplateformfieldgroupRegexp field to given value.

### HasSEzsigntemplateformfieldgroupRegexp

`func (o *EzsigntemplateformfieldgroupRequestCompound) HasSEzsigntemplateformfieldgroupRegexp() bool`

HasSEzsigntemplateformfieldgroupRegexp returns a boolean if a field has been set.

### GetSEzsigntemplateformfieldgroupTextvalidationcustommessage

`func (o *EzsigntemplateformfieldgroupRequestCompound) GetSEzsigntemplateformfieldgroupTextvalidationcustommessage() string`

GetSEzsigntemplateformfieldgroupTextvalidationcustommessage returns the SEzsigntemplateformfieldgroupTextvalidationcustommessage field if non-nil, zero value otherwise.

### GetSEzsigntemplateformfieldgroupTextvalidationcustommessageOk

`func (o *EzsigntemplateformfieldgroupRequestCompound) GetSEzsigntemplateformfieldgroupTextvalidationcustommessageOk() (*string, bool)`

GetSEzsigntemplateformfieldgroupTextvalidationcustommessageOk returns a tuple with the SEzsigntemplateformfieldgroupTextvalidationcustommessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplateformfieldgroupTextvalidationcustommessage

`func (o *EzsigntemplateformfieldgroupRequestCompound) SetSEzsigntemplateformfieldgroupTextvalidationcustommessage(v string)`

SetSEzsigntemplateformfieldgroupTextvalidationcustommessage sets SEzsigntemplateformfieldgroupTextvalidationcustommessage field to given value.

### HasSEzsigntemplateformfieldgroupTextvalidationcustommessage

`func (o *EzsigntemplateformfieldgroupRequestCompound) HasSEzsigntemplateformfieldgroupTextvalidationcustommessage() bool`

HasSEzsigntemplateformfieldgroupTextvalidationcustommessage returns a boolean if a field has been set.

### GetEEzsigntemplateformfieldgroupTextvalidation

`func (o *EzsigntemplateformfieldgroupRequestCompound) GetEEzsigntemplateformfieldgroupTextvalidation() EnumTextvalidation`

GetEEzsigntemplateformfieldgroupTextvalidation returns the EEzsigntemplateformfieldgroupTextvalidation field if non-nil, zero value otherwise.

### GetEEzsigntemplateformfieldgroupTextvalidationOk

`func (o *EzsigntemplateformfieldgroupRequestCompound) GetEEzsigntemplateformfieldgroupTextvalidationOk() (*EnumTextvalidation, bool)`

GetEEzsigntemplateformfieldgroupTextvalidationOk returns a tuple with the EEzsigntemplateformfieldgroupTextvalidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplateformfieldgroupTextvalidation

`func (o *EzsigntemplateformfieldgroupRequestCompound) SetEEzsigntemplateformfieldgroupTextvalidation(v EnumTextvalidation)`

SetEEzsigntemplateformfieldgroupTextvalidation sets EEzsigntemplateformfieldgroupTextvalidation field to given value.

### HasEEzsigntemplateformfieldgroupTextvalidation

`func (o *EzsigntemplateformfieldgroupRequestCompound) HasEEzsigntemplateformfieldgroupTextvalidation() bool`

HasEEzsigntemplateformfieldgroupTextvalidation returns a boolean if a field has been set.

### GetTEzsigntemplateformfieldgroupTooltip

`func (o *EzsigntemplateformfieldgroupRequestCompound) GetTEzsigntemplateformfieldgroupTooltip() string`

GetTEzsigntemplateformfieldgroupTooltip returns the TEzsigntemplateformfieldgroupTooltip field if non-nil, zero value otherwise.

### GetTEzsigntemplateformfieldgroupTooltipOk

`func (o *EzsigntemplateformfieldgroupRequestCompound) GetTEzsigntemplateformfieldgroupTooltipOk() (*string, bool)`

GetTEzsigntemplateformfieldgroupTooltipOk returns a tuple with the TEzsigntemplateformfieldgroupTooltip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEzsigntemplateformfieldgroupTooltip

`func (o *EzsigntemplateformfieldgroupRequestCompound) SetTEzsigntemplateformfieldgroupTooltip(v string)`

SetTEzsigntemplateformfieldgroupTooltip sets TEzsigntemplateformfieldgroupTooltip field to given value.

### HasTEzsigntemplateformfieldgroupTooltip

`func (o *EzsigntemplateformfieldgroupRequestCompound) HasTEzsigntemplateformfieldgroupTooltip() bool`

HasTEzsigntemplateformfieldgroupTooltip returns a boolean if a field has been set.

### GetEEzsigntemplateformfieldgroupTooltipposition

`func (o *EzsigntemplateformfieldgroupRequestCompound) GetEEzsigntemplateformfieldgroupTooltipposition() FieldEEzsigntemplateformfieldgroupTooltipposition`

GetEEzsigntemplateformfieldgroupTooltipposition returns the EEzsigntemplateformfieldgroupTooltipposition field if non-nil, zero value otherwise.

### GetEEzsigntemplateformfieldgroupTooltippositionOk

`func (o *EzsigntemplateformfieldgroupRequestCompound) GetEEzsigntemplateformfieldgroupTooltippositionOk() (*FieldEEzsigntemplateformfieldgroupTooltipposition, bool)`

GetEEzsigntemplateformfieldgroupTooltippositionOk returns a tuple with the EEzsigntemplateformfieldgroupTooltipposition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplateformfieldgroupTooltipposition

`func (o *EzsigntemplateformfieldgroupRequestCompound) SetEEzsigntemplateformfieldgroupTooltipposition(v FieldEEzsigntemplateformfieldgroupTooltipposition)`

SetEEzsigntemplateformfieldgroupTooltipposition sets EEzsigntemplateformfieldgroupTooltipposition field to given value.

### HasEEzsigntemplateformfieldgroupTooltipposition

`func (o *EzsigntemplateformfieldgroupRequestCompound) HasEEzsigntemplateformfieldgroupTooltipposition() bool`

HasEEzsigntemplateformfieldgroupTooltipposition returns a boolean if a field has been set.

### GetAObjEzsigntemplateformfieldgroupsigner

`func (o *EzsigntemplateformfieldgroupRequestCompound) GetAObjEzsigntemplateformfieldgroupsigner() []EzsigntemplateformfieldgroupsignerRequestCompound`

GetAObjEzsigntemplateformfieldgroupsigner returns the AObjEzsigntemplateformfieldgroupsigner field if non-nil, zero value otherwise.

### GetAObjEzsigntemplateformfieldgroupsignerOk

`func (o *EzsigntemplateformfieldgroupRequestCompound) GetAObjEzsigntemplateformfieldgroupsignerOk() (*[]EzsigntemplateformfieldgroupsignerRequestCompound, bool)`

GetAObjEzsigntemplateformfieldgroupsignerOk returns a tuple with the AObjEzsigntemplateformfieldgroupsigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzsigntemplateformfieldgroupsigner

`func (o *EzsigntemplateformfieldgroupRequestCompound) SetAObjEzsigntemplateformfieldgroupsigner(v []EzsigntemplateformfieldgroupsignerRequestCompound)`

SetAObjEzsigntemplateformfieldgroupsigner sets AObjEzsigntemplateformfieldgroupsigner field to given value.


### GetAObjDropdownElement

`func (o *EzsigntemplateformfieldgroupRequestCompound) GetAObjDropdownElement() []CustomDropdownElementRequestCompound`

GetAObjDropdownElement returns the AObjDropdownElement field if non-nil, zero value otherwise.

### GetAObjDropdownElementOk

`func (o *EzsigntemplateformfieldgroupRequestCompound) GetAObjDropdownElementOk() (*[]CustomDropdownElementRequestCompound, bool)`

GetAObjDropdownElementOk returns a tuple with the AObjDropdownElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjDropdownElement

`func (o *EzsigntemplateformfieldgroupRequestCompound) SetAObjDropdownElement(v []CustomDropdownElementRequestCompound)`

SetAObjDropdownElement sets AObjDropdownElement field to given value.

### HasAObjDropdownElement

`func (o *EzsigntemplateformfieldgroupRequestCompound) HasAObjDropdownElement() bool`

HasAObjDropdownElement returns a boolean if a field has been set.

### GetAObjEzsigntemplateformfield

`func (o *EzsigntemplateformfieldgroupRequestCompound) GetAObjEzsigntemplateformfield() []EzsigntemplateformfieldRequestCompound`

GetAObjEzsigntemplateformfield returns the AObjEzsigntemplateformfield field if non-nil, zero value otherwise.

### GetAObjEzsigntemplateformfieldOk

`func (o *EzsigntemplateformfieldgroupRequestCompound) GetAObjEzsigntemplateformfieldOk() (*[]EzsigntemplateformfieldRequestCompound, bool)`

GetAObjEzsigntemplateformfieldOk returns a tuple with the AObjEzsigntemplateformfield field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzsigntemplateformfield

`func (o *EzsigntemplateformfieldgroupRequestCompound) SetAObjEzsigntemplateformfield(v []EzsigntemplateformfieldRequestCompound)`

SetAObjEzsigntemplateformfield sets AObjEzsigntemplateformfield field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


