# EzsignformfieldgroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsignformfieldgroupID** | Pointer to **int32** | The unique ID of the Ezsignformfieldgroup | [optional] 
**FkiEzsigndocumentID** | **int32** | The unique ID of the Ezsigndocument | 
**EEzsignformfieldgroupType** | [**FieldEEzsignformfieldgroupType**](FieldEEzsignformfieldgroupType.md) |  | 
**EEzsignformfieldgroupSignerrequirement** | Pointer to [**FieldEEzsignformfieldgroupSignerrequirement**](FieldEEzsignformfieldgroupSignerrequirement.md) |  | [optional] 
**SEzsignformfieldgroupLabel** | **string** | The Label for the Ezsignformfieldgroup | 
**IEzsignformfieldgroupStep** | **int32** | The step when the Ezsignsigner will be invited to fill the form fields | 
**SEzsignformfieldgroupDefaultvalue** | Pointer to **string** | The default value for the Ezsignformfieldgroup  You can use the codes below and they will be replaced at signature time.    | Code | Description | Example | | ------------------------- | ------------ | ------------ | | {sUserFirstname} | The first name of the contact | John | | {sUserLastname} | The last name of the contact | Doe | | {sUserJobtitle} | The job title | Sales Representative | | {sEmailAddress} | The email address | email@example.com | | {sPhoneE164} | A phone number in E.164 Format | +15149901516 | | {sPhoneE164Cell} | A phone number in E.164 Format | +15149901516 | | [optional] 
**IEzsignformfieldgroupFilledmin** | **int32** | The minimum number of Ezsignformfield that must be filled in the Ezsignformfieldgroup | 
**IEzsignformfieldgroupFilledmax** | **int32** | The maximum number of Ezsignformfield that must be filled in the Ezsignformfieldgroup | 
**BEzsignformfieldgroupReadonly** | **bool** | Whether the Ezsignformfieldgroup is read only or not. | 
**IEzsignformfieldgroupMaxlength** | Pointer to **int32** | The maximum length for the value in the Ezsignformfieldgroup  This can only be set if eEzsignformfieldgroupType is **Text** or **Textarea** | [optional] 
**BEzsignformfieldgroupEncrypted** | Pointer to **bool** | Whether the Ezsignformfieldgroup is encrypted in the database or not. Encrypted values are not displayed on the Ezsigndocument. This can only be set if eEzsignformfieldgroupType is **Text** or **Textarea** | [optional] 
**SEzsignformfieldgroupRegexp** | Pointer to **string** | A regular expression to indicate what values are acceptable for the Ezsignformfieldgroup.  This can only be set if eEzsignformfieldgroupType is **Text** or **Textarea** | [optional] 
**TEzsignformfieldgroupTooltip** | Pointer to **string** | A tooltip that will be presented to Ezsignsigner about the Ezsignformfieldgroup | [optional] 
**EEzsignformfieldgroupTooltipposition** | Pointer to [**FieldEEzsignformfieldgroupTooltipposition**](FieldEEzsignformfieldgroupTooltipposition.md) |  | [optional] 
**EEzsignformfieldgroupTextvalidation** | Pointer to [**EnumTextvalidation**](EnumTextvalidation.md) |  | [optional] 

## Methods

### NewEzsignformfieldgroupRequest

`func NewEzsignformfieldgroupRequest(fkiEzsigndocumentID int32, eEzsignformfieldgroupType FieldEEzsignformfieldgroupType, sEzsignformfieldgroupLabel string, iEzsignformfieldgroupStep int32, iEzsignformfieldgroupFilledmin int32, iEzsignformfieldgroupFilledmax int32, bEzsignformfieldgroupReadonly bool, ) *EzsignformfieldgroupRequest`

NewEzsignformfieldgroupRequest instantiates a new EzsignformfieldgroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignformfieldgroupRequestWithDefaults

`func NewEzsignformfieldgroupRequestWithDefaults() *EzsignformfieldgroupRequest`

NewEzsignformfieldgroupRequestWithDefaults instantiates a new EzsignformfieldgroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsignformfieldgroupID

`func (o *EzsignformfieldgroupRequest) GetPkiEzsignformfieldgroupID() int32`

GetPkiEzsignformfieldgroupID returns the PkiEzsignformfieldgroupID field if non-nil, zero value otherwise.

### GetPkiEzsignformfieldgroupIDOk

`func (o *EzsignformfieldgroupRequest) GetPkiEzsignformfieldgroupIDOk() (*int32, bool)`

GetPkiEzsignformfieldgroupIDOk returns a tuple with the PkiEzsignformfieldgroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsignformfieldgroupID

`func (o *EzsignformfieldgroupRequest) SetPkiEzsignformfieldgroupID(v int32)`

SetPkiEzsignformfieldgroupID sets PkiEzsignformfieldgroupID field to given value.

### HasPkiEzsignformfieldgroupID

`func (o *EzsignformfieldgroupRequest) HasPkiEzsignformfieldgroupID() bool`

HasPkiEzsignformfieldgroupID returns a boolean if a field has been set.

### GetFkiEzsigndocumentID

`func (o *EzsignformfieldgroupRequest) GetFkiEzsigndocumentID() int32`

GetFkiEzsigndocumentID returns the FkiEzsigndocumentID field if non-nil, zero value otherwise.

### GetFkiEzsigndocumentIDOk

`func (o *EzsignformfieldgroupRequest) GetFkiEzsigndocumentIDOk() (*int32, bool)`

GetFkiEzsigndocumentIDOk returns a tuple with the FkiEzsigndocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigndocumentID

`func (o *EzsignformfieldgroupRequest) SetFkiEzsigndocumentID(v int32)`

SetFkiEzsigndocumentID sets FkiEzsigndocumentID field to given value.


### GetEEzsignformfieldgroupType

`func (o *EzsignformfieldgroupRequest) GetEEzsignformfieldgroupType() FieldEEzsignformfieldgroupType`

GetEEzsignformfieldgroupType returns the EEzsignformfieldgroupType field if non-nil, zero value otherwise.

### GetEEzsignformfieldgroupTypeOk

`func (o *EzsignformfieldgroupRequest) GetEEzsignformfieldgroupTypeOk() (*FieldEEzsignformfieldgroupType, bool)`

GetEEzsignformfieldgroupTypeOk returns a tuple with the EEzsignformfieldgroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignformfieldgroupType

`func (o *EzsignformfieldgroupRequest) SetEEzsignformfieldgroupType(v FieldEEzsignformfieldgroupType)`

SetEEzsignformfieldgroupType sets EEzsignformfieldgroupType field to given value.


### GetEEzsignformfieldgroupSignerrequirement

`func (o *EzsignformfieldgroupRequest) GetEEzsignformfieldgroupSignerrequirement() FieldEEzsignformfieldgroupSignerrequirement`

GetEEzsignformfieldgroupSignerrequirement returns the EEzsignformfieldgroupSignerrequirement field if non-nil, zero value otherwise.

### GetEEzsignformfieldgroupSignerrequirementOk

`func (o *EzsignformfieldgroupRequest) GetEEzsignformfieldgroupSignerrequirementOk() (*FieldEEzsignformfieldgroupSignerrequirement, bool)`

GetEEzsignformfieldgroupSignerrequirementOk returns a tuple with the EEzsignformfieldgroupSignerrequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignformfieldgroupSignerrequirement

`func (o *EzsignformfieldgroupRequest) SetEEzsignformfieldgroupSignerrequirement(v FieldEEzsignformfieldgroupSignerrequirement)`

SetEEzsignformfieldgroupSignerrequirement sets EEzsignformfieldgroupSignerrequirement field to given value.

### HasEEzsignformfieldgroupSignerrequirement

`func (o *EzsignformfieldgroupRequest) HasEEzsignformfieldgroupSignerrequirement() bool`

HasEEzsignformfieldgroupSignerrequirement returns a boolean if a field has been set.

### GetSEzsignformfieldgroupLabel

`func (o *EzsignformfieldgroupRequest) GetSEzsignformfieldgroupLabel() string`

GetSEzsignformfieldgroupLabel returns the SEzsignformfieldgroupLabel field if non-nil, zero value otherwise.

### GetSEzsignformfieldgroupLabelOk

`func (o *EzsignformfieldgroupRequest) GetSEzsignformfieldgroupLabelOk() (*string, bool)`

GetSEzsignformfieldgroupLabelOk returns a tuple with the SEzsignformfieldgroupLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignformfieldgroupLabel

`func (o *EzsignformfieldgroupRequest) SetSEzsignformfieldgroupLabel(v string)`

SetSEzsignformfieldgroupLabel sets SEzsignformfieldgroupLabel field to given value.


### GetIEzsignformfieldgroupStep

`func (o *EzsignformfieldgroupRequest) GetIEzsignformfieldgroupStep() int32`

GetIEzsignformfieldgroupStep returns the IEzsignformfieldgroupStep field if non-nil, zero value otherwise.

### GetIEzsignformfieldgroupStepOk

`func (o *EzsignformfieldgroupRequest) GetIEzsignformfieldgroupStepOk() (*int32, bool)`

GetIEzsignformfieldgroupStepOk returns a tuple with the IEzsignformfieldgroupStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignformfieldgroupStep

`func (o *EzsignformfieldgroupRequest) SetIEzsignformfieldgroupStep(v int32)`

SetIEzsignformfieldgroupStep sets IEzsignformfieldgroupStep field to given value.


### GetSEzsignformfieldgroupDefaultvalue

`func (o *EzsignformfieldgroupRequest) GetSEzsignformfieldgroupDefaultvalue() string`

GetSEzsignformfieldgroupDefaultvalue returns the SEzsignformfieldgroupDefaultvalue field if non-nil, zero value otherwise.

### GetSEzsignformfieldgroupDefaultvalueOk

`func (o *EzsignformfieldgroupRequest) GetSEzsignformfieldgroupDefaultvalueOk() (*string, bool)`

GetSEzsignformfieldgroupDefaultvalueOk returns a tuple with the SEzsignformfieldgroupDefaultvalue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignformfieldgroupDefaultvalue

`func (o *EzsignformfieldgroupRequest) SetSEzsignformfieldgroupDefaultvalue(v string)`

SetSEzsignformfieldgroupDefaultvalue sets SEzsignformfieldgroupDefaultvalue field to given value.

### HasSEzsignformfieldgroupDefaultvalue

`func (o *EzsignformfieldgroupRequest) HasSEzsignformfieldgroupDefaultvalue() bool`

HasSEzsignformfieldgroupDefaultvalue returns a boolean if a field has been set.

### GetIEzsignformfieldgroupFilledmin

`func (o *EzsignformfieldgroupRequest) GetIEzsignformfieldgroupFilledmin() int32`

GetIEzsignformfieldgroupFilledmin returns the IEzsignformfieldgroupFilledmin field if non-nil, zero value otherwise.

### GetIEzsignformfieldgroupFilledminOk

`func (o *EzsignformfieldgroupRequest) GetIEzsignformfieldgroupFilledminOk() (*int32, bool)`

GetIEzsignformfieldgroupFilledminOk returns a tuple with the IEzsignformfieldgroupFilledmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignformfieldgroupFilledmin

`func (o *EzsignformfieldgroupRequest) SetIEzsignformfieldgroupFilledmin(v int32)`

SetIEzsignformfieldgroupFilledmin sets IEzsignformfieldgroupFilledmin field to given value.


### GetIEzsignformfieldgroupFilledmax

`func (o *EzsignformfieldgroupRequest) GetIEzsignformfieldgroupFilledmax() int32`

GetIEzsignformfieldgroupFilledmax returns the IEzsignformfieldgroupFilledmax field if non-nil, zero value otherwise.

### GetIEzsignformfieldgroupFilledmaxOk

`func (o *EzsignformfieldgroupRequest) GetIEzsignformfieldgroupFilledmaxOk() (*int32, bool)`

GetIEzsignformfieldgroupFilledmaxOk returns a tuple with the IEzsignformfieldgroupFilledmax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignformfieldgroupFilledmax

`func (o *EzsignformfieldgroupRequest) SetIEzsignformfieldgroupFilledmax(v int32)`

SetIEzsignformfieldgroupFilledmax sets IEzsignformfieldgroupFilledmax field to given value.


### GetBEzsignformfieldgroupReadonly

`func (o *EzsignformfieldgroupRequest) GetBEzsignformfieldgroupReadonly() bool`

GetBEzsignformfieldgroupReadonly returns the BEzsignformfieldgroupReadonly field if non-nil, zero value otherwise.

### GetBEzsignformfieldgroupReadonlyOk

`func (o *EzsignformfieldgroupRequest) GetBEzsignformfieldgroupReadonlyOk() (*bool, bool)`

GetBEzsignformfieldgroupReadonlyOk returns a tuple with the BEzsignformfieldgroupReadonly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignformfieldgroupReadonly

`func (o *EzsignformfieldgroupRequest) SetBEzsignformfieldgroupReadonly(v bool)`

SetBEzsignformfieldgroupReadonly sets BEzsignformfieldgroupReadonly field to given value.


### GetIEzsignformfieldgroupMaxlength

`func (o *EzsignformfieldgroupRequest) GetIEzsignformfieldgroupMaxlength() int32`

GetIEzsignformfieldgroupMaxlength returns the IEzsignformfieldgroupMaxlength field if non-nil, zero value otherwise.

### GetIEzsignformfieldgroupMaxlengthOk

`func (o *EzsignformfieldgroupRequest) GetIEzsignformfieldgroupMaxlengthOk() (*int32, bool)`

GetIEzsignformfieldgroupMaxlengthOk returns a tuple with the IEzsignformfieldgroupMaxlength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignformfieldgroupMaxlength

`func (o *EzsignformfieldgroupRequest) SetIEzsignformfieldgroupMaxlength(v int32)`

SetIEzsignformfieldgroupMaxlength sets IEzsignformfieldgroupMaxlength field to given value.

### HasIEzsignformfieldgroupMaxlength

`func (o *EzsignformfieldgroupRequest) HasIEzsignformfieldgroupMaxlength() bool`

HasIEzsignformfieldgroupMaxlength returns a boolean if a field has been set.

### GetBEzsignformfieldgroupEncrypted

`func (o *EzsignformfieldgroupRequest) GetBEzsignformfieldgroupEncrypted() bool`

GetBEzsignformfieldgroupEncrypted returns the BEzsignformfieldgroupEncrypted field if non-nil, zero value otherwise.

### GetBEzsignformfieldgroupEncryptedOk

`func (o *EzsignformfieldgroupRequest) GetBEzsignformfieldgroupEncryptedOk() (*bool, bool)`

GetBEzsignformfieldgroupEncryptedOk returns a tuple with the BEzsignformfieldgroupEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignformfieldgroupEncrypted

`func (o *EzsignformfieldgroupRequest) SetBEzsignformfieldgroupEncrypted(v bool)`

SetBEzsignformfieldgroupEncrypted sets BEzsignformfieldgroupEncrypted field to given value.

### HasBEzsignformfieldgroupEncrypted

`func (o *EzsignformfieldgroupRequest) HasBEzsignformfieldgroupEncrypted() bool`

HasBEzsignformfieldgroupEncrypted returns a boolean if a field has been set.

### GetSEzsignformfieldgroupRegexp

`func (o *EzsignformfieldgroupRequest) GetSEzsignformfieldgroupRegexp() string`

GetSEzsignformfieldgroupRegexp returns the SEzsignformfieldgroupRegexp field if non-nil, zero value otherwise.

### GetSEzsignformfieldgroupRegexpOk

`func (o *EzsignformfieldgroupRequest) GetSEzsignformfieldgroupRegexpOk() (*string, bool)`

GetSEzsignformfieldgroupRegexpOk returns a tuple with the SEzsignformfieldgroupRegexp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignformfieldgroupRegexp

`func (o *EzsignformfieldgroupRequest) SetSEzsignformfieldgroupRegexp(v string)`

SetSEzsignformfieldgroupRegexp sets SEzsignformfieldgroupRegexp field to given value.

### HasSEzsignformfieldgroupRegexp

`func (o *EzsignformfieldgroupRequest) HasSEzsignformfieldgroupRegexp() bool`

HasSEzsignformfieldgroupRegexp returns a boolean if a field has been set.

### GetTEzsignformfieldgroupTooltip

`func (o *EzsignformfieldgroupRequest) GetTEzsignformfieldgroupTooltip() string`

GetTEzsignformfieldgroupTooltip returns the TEzsignformfieldgroupTooltip field if non-nil, zero value otherwise.

### GetTEzsignformfieldgroupTooltipOk

`func (o *EzsignformfieldgroupRequest) GetTEzsignformfieldgroupTooltipOk() (*string, bool)`

GetTEzsignformfieldgroupTooltipOk returns a tuple with the TEzsignformfieldgroupTooltip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEzsignformfieldgroupTooltip

`func (o *EzsignformfieldgroupRequest) SetTEzsignformfieldgroupTooltip(v string)`

SetTEzsignformfieldgroupTooltip sets TEzsignformfieldgroupTooltip field to given value.

### HasTEzsignformfieldgroupTooltip

`func (o *EzsignformfieldgroupRequest) HasTEzsignformfieldgroupTooltip() bool`

HasTEzsignformfieldgroupTooltip returns a boolean if a field has been set.

### GetEEzsignformfieldgroupTooltipposition

`func (o *EzsignformfieldgroupRequest) GetEEzsignformfieldgroupTooltipposition() FieldEEzsignformfieldgroupTooltipposition`

GetEEzsignformfieldgroupTooltipposition returns the EEzsignformfieldgroupTooltipposition field if non-nil, zero value otherwise.

### GetEEzsignformfieldgroupTooltippositionOk

`func (o *EzsignformfieldgroupRequest) GetEEzsignformfieldgroupTooltippositionOk() (*FieldEEzsignformfieldgroupTooltipposition, bool)`

GetEEzsignformfieldgroupTooltippositionOk returns a tuple with the EEzsignformfieldgroupTooltipposition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignformfieldgroupTooltipposition

`func (o *EzsignformfieldgroupRequest) SetEEzsignformfieldgroupTooltipposition(v FieldEEzsignformfieldgroupTooltipposition)`

SetEEzsignformfieldgroupTooltipposition sets EEzsignformfieldgroupTooltipposition field to given value.

### HasEEzsignformfieldgroupTooltipposition

`func (o *EzsignformfieldgroupRequest) HasEEzsignformfieldgroupTooltipposition() bool`

HasEEzsignformfieldgroupTooltipposition returns a boolean if a field has been set.

### GetEEzsignformfieldgroupTextvalidation

`func (o *EzsignformfieldgroupRequest) GetEEzsignformfieldgroupTextvalidation() EnumTextvalidation`

GetEEzsignformfieldgroupTextvalidation returns the EEzsignformfieldgroupTextvalidation field if non-nil, zero value otherwise.

### GetEEzsignformfieldgroupTextvalidationOk

`func (o *EzsignformfieldgroupRequest) GetEEzsignformfieldgroupTextvalidationOk() (*EnumTextvalidation, bool)`

GetEEzsignformfieldgroupTextvalidationOk returns a tuple with the EEzsignformfieldgroupTextvalidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignformfieldgroupTextvalidation

`func (o *EzsignformfieldgroupRequest) SetEEzsignformfieldgroupTextvalidation(v EnumTextvalidation)`

SetEEzsignformfieldgroupTextvalidation sets EEzsignformfieldgroupTextvalidation field to given value.

### HasEEzsignformfieldgroupTextvalidation

`func (o *EzsignformfieldgroupRequest) HasEEzsignformfieldgroupTextvalidation() bool`

HasEEzsignformfieldgroupTextvalidation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


