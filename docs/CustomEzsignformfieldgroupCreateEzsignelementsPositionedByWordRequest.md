# CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsignformfieldgroupID** | Pointer to **int32** | The unique ID of the Ezsignformfieldgroup | [optional] 
**FkiEzsigndocumentID** | **int32** | The unique ID of the Ezsigndocument | 
**EEzsignformfieldgroupType** | [**FieldEEzsignformfieldgroupType**](FieldEEzsignformfieldgroupType.md) |  | 
**EEzsignformfieldgroupSignerrequirement** | [**FieldEEzsignformfieldgroupSignerrequirement**](FieldEEzsignformfieldgroupSignerrequirement.md) |  | 
**SEzsignformfieldgroupLabel** | **string** | The Label for the Ezsignformfieldgroup | 
**IEzsignformfieldgroupStep** | **int32** | The step when the Ezsignsigner will be invited to fill the form fields | 
**SEzsignformfieldgroupDefaultvalue** | Pointer to **string** | The default value for the Ezsignformfieldgroup | [optional] 
**IEzsignformfieldgroupFilledmin** | **int32** | The minimum number of Ezsignformfield that must be filled in the Ezsignformfieldgroup | 
**IEzsignformfieldgroupFilledmax** | **int32** | The maximum number of Ezsignformfield that must be filled in the Ezsignformfieldgroup | 
**BEzsignformfieldgroupReadonly** | **bool** | Whether the Ezsignformfieldgroup is read only or not. | 
**IEzsignformfieldgroupMaxlength** | Pointer to **int32** | The maximum length for the value in the Ezsignformfieldgroup  This can only be set if eEzsignformfieldgroupType is **Text** or **Textarea** | [optional] 
**BEzsignformfieldgroupEncrypted** | Pointer to **bool** | Whether the Ezsignformfieldgroup is encrypted in the database or not. Encrypted values are not displayed on the Ezsigndocument. This can only be set if eEzsignformfieldgroupType is **Text** or **Textarea** | [optional] 
**SEzsignformfieldgroupRegexp** | Pointer to **string** | A regular expression to indicate what values are acceptable for the Ezsignformfieldgroup.  This can only be set if eEzsignformfieldgroupType is **Text** or **Textarea** | [optional] 
**TEzsignformfieldgroupTooltip** | Pointer to **string** | A tooltip that will be presented to Ezsignsigner about the Ezsignformfieldgroup | [optional] 
**EEzsignformfieldgroupTooltipposition** | Pointer to [**FieldEEzsignformfieldgroupTooltipposition**](FieldEEzsignformfieldgroupTooltipposition.md) |  | [optional] 
**EEzsignformfieldgroupTextvalidation** | Pointer to [**EnumTextvalidation**](EnumTextvalidation.md) |  | [optional] 
**AObjEzsignformfieldgroupsigner** | [**[]EzsignformfieldgroupsignerRequestCompound**](EzsignformfieldgroupsignerRequestCompound.md) |  | 
**AObjDropdownElement** | Pointer to [**[]CustomDropdownElementRequestCompound**](CustomDropdownElementRequestCompound.md) |  | [optional] 
**AObjEzsignformfield** | [**[]EzsignformfieldRequestCompound**](EzsignformfieldRequestCompound.md) |  | 
**ObjCreateezsignelementspositionedbyword** | [**CustomCreateEzsignelementsPositionedByWordRequest**](CustomCreateEzsignelementsPositionedByWordRequest.md) |  | 

## Methods

### NewCustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest

`func NewCustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest(fkiEzsigndocumentID int32, eEzsignformfieldgroupType FieldEEzsignformfieldgroupType, eEzsignformfieldgroupSignerrequirement FieldEEzsignformfieldgroupSignerrequirement, sEzsignformfieldgroupLabel string, iEzsignformfieldgroupStep int32, iEzsignformfieldgroupFilledmin int32, iEzsignformfieldgroupFilledmax int32, bEzsignformfieldgroupReadonly bool, aObjEzsignformfieldgroupsigner []EzsignformfieldgroupsignerRequestCompound, aObjEzsignformfield []EzsignformfieldRequestCompound, objCreateezsignelementspositionedbyword CustomCreateEzsignelementsPositionedByWordRequest, ) *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest`

NewCustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest instantiates a new CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequestWithDefaults

`func NewCustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequestWithDefaults() *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest`

NewCustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequestWithDefaults instantiates a new CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsignformfieldgroupID

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetPkiEzsignformfieldgroupID() int32`

GetPkiEzsignformfieldgroupID returns the PkiEzsignformfieldgroupID field if non-nil, zero value otherwise.

### GetPkiEzsignformfieldgroupIDOk

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetPkiEzsignformfieldgroupIDOk() (*int32, bool)`

GetPkiEzsignformfieldgroupIDOk returns a tuple with the PkiEzsignformfieldgroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsignformfieldgroupID

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) SetPkiEzsignformfieldgroupID(v int32)`

SetPkiEzsignformfieldgroupID sets PkiEzsignformfieldgroupID field to given value.

### HasPkiEzsignformfieldgroupID

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) HasPkiEzsignformfieldgroupID() bool`

HasPkiEzsignformfieldgroupID returns a boolean if a field has been set.

### GetFkiEzsigndocumentID

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetFkiEzsigndocumentID() int32`

GetFkiEzsigndocumentID returns the FkiEzsigndocumentID field if non-nil, zero value otherwise.

### GetFkiEzsigndocumentIDOk

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetFkiEzsigndocumentIDOk() (*int32, bool)`

GetFkiEzsigndocumentIDOk returns a tuple with the FkiEzsigndocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigndocumentID

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) SetFkiEzsigndocumentID(v int32)`

SetFkiEzsigndocumentID sets FkiEzsigndocumentID field to given value.


### GetEEzsignformfieldgroupType

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetEEzsignformfieldgroupType() FieldEEzsignformfieldgroupType`

GetEEzsignformfieldgroupType returns the EEzsignformfieldgroupType field if non-nil, zero value otherwise.

### GetEEzsignformfieldgroupTypeOk

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetEEzsignformfieldgroupTypeOk() (*FieldEEzsignformfieldgroupType, bool)`

GetEEzsignformfieldgroupTypeOk returns a tuple with the EEzsignformfieldgroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignformfieldgroupType

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) SetEEzsignformfieldgroupType(v FieldEEzsignformfieldgroupType)`

SetEEzsignformfieldgroupType sets EEzsignformfieldgroupType field to given value.


### GetEEzsignformfieldgroupSignerrequirement

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetEEzsignformfieldgroupSignerrequirement() FieldEEzsignformfieldgroupSignerrequirement`

GetEEzsignformfieldgroupSignerrequirement returns the EEzsignformfieldgroupSignerrequirement field if non-nil, zero value otherwise.

### GetEEzsignformfieldgroupSignerrequirementOk

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetEEzsignformfieldgroupSignerrequirementOk() (*FieldEEzsignformfieldgroupSignerrequirement, bool)`

GetEEzsignformfieldgroupSignerrequirementOk returns a tuple with the EEzsignformfieldgroupSignerrequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignformfieldgroupSignerrequirement

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) SetEEzsignformfieldgroupSignerrequirement(v FieldEEzsignformfieldgroupSignerrequirement)`

SetEEzsignformfieldgroupSignerrequirement sets EEzsignformfieldgroupSignerrequirement field to given value.


### GetSEzsignformfieldgroupLabel

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetSEzsignformfieldgroupLabel() string`

GetSEzsignformfieldgroupLabel returns the SEzsignformfieldgroupLabel field if non-nil, zero value otherwise.

### GetSEzsignformfieldgroupLabelOk

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetSEzsignformfieldgroupLabelOk() (*string, bool)`

GetSEzsignformfieldgroupLabelOk returns a tuple with the SEzsignformfieldgroupLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignformfieldgroupLabel

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) SetSEzsignformfieldgroupLabel(v string)`

SetSEzsignformfieldgroupLabel sets SEzsignformfieldgroupLabel field to given value.


### GetIEzsignformfieldgroupStep

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetIEzsignformfieldgroupStep() int32`

GetIEzsignformfieldgroupStep returns the IEzsignformfieldgroupStep field if non-nil, zero value otherwise.

### GetIEzsignformfieldgroupStepOk

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetIEzsignformfieldgroupStepOk() (*int32, bool)`

GetIEzsignformfieldgroupStepOk returns a tuple with the IEzsignformfieldgroupStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignformfieldgroupStep

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) SetIEzsignformfieldgroupStep(v int32)`

SetIEzsignformfieldgroupStep sets IEzsignformfieldgroupStep field to given value.


### GetSEzsignformfieldgroupDefaultvalue

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetSEzsignformfieldgroupDefaultvalue() string`

GetSEzsignformfieldgroupDefaultvalue returns the SEzsignformfieldgroupDefaultvalue field if non-nil, zero value otherwise.

### GetSEzsignformfieldgroupDefaultvalueOk

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetSEzsignformfieldgroupDefaultvalueOk() (*string, bool)`

GetSEzsignformfieldgroupDefaultvalueOk returns a tuple with the SEzsignformfieldgroupDefaultvalue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignformfieldgroupDefaultvalue

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) SetSEzsignformfieldgroupDefaultvalue(v string)`

SetSEzsignformfieldgroupDefaultvalue sets SEzsignformfieldgroupDefaultvalue field to given value.

### HasSEzsignformfieldgroupDefaultvalue

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) HasSEzsignformfieldgroupDefaultvalue() bool`

HasSEzsignformfieldgroupDefaultvalue returns a boolean if a field has been set.

### GetIEzsignformfieldgroupFilledmin

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetIEzsignformfieldgroupFilledmin() int32`

GetIEzsignformfieldgroupFilledmin returns the IEzsignformfieldgroupFilledmin field if non-nil, zero value otherwise.

### GetIEzsignformfieldgroupFilledminOk

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetIEzsignformfieldgroupFilledminOk() (*int32, bool)`

GetIEzsignformfieldgroupFilledminOk returns a tuple with the IEzsignformfieldgroupFilledmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignformfieldgroupFilledmin

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) SetIEzsignformfieldgroupFilledmin(v int32)`

SetIEzsignformfieldgroupFilledmin sets IEzsignformfieldgroupFilledmin field to given value.


### GetIEzsignformfieldgroupFilledmax

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetIEzsignformfieldgroupFilledmax() int32`

GetIEzsignformfieldgroupFilledmax returns the IEzsignformfieldgroupFilledmax field if non-nil, zero value otherwise.

### GetIEzsignformfieldgroupFilledmaxOk

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetIEzsignformfieldgroupFilledmaxOk() (*int32, bool)`

GetIEzsignformfieldgroupFilledmaxOk returns a tuple with the IEzsignformfieldgroupFilledmax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignformfieldgroupFilledmax

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) SetIEzsignformfieldgroupFilledmax(v int32)`

SetIEzsignformfieldgroupFilledmax sets IEzsignformfieldgroupFilledmax field to given value.


### GetBEzsignformfieldgroupReadonly

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetBEzsignformfieldgroupReadonly() bool`

GetBEzsignformfieldgroupReadonly returns the BEzsignformfieldgroupReadonly field if non-nil, zero value otherwise.

### GetBEzsignformfieldgroupReadonlyOk

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetBEzsignformfieldgroupReadonlyOk() (*bool, bool)`

GetBEzsignformfieldgroupReadonlyOk returns a tuple with the BEzsignformfieldgroupReadonly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignformfieldgroupReadonly

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) SetBEzsignformfieldgroupReadonly(v bool)`

SetBEzsignformfieldgroupReadonly sets BEzsignformfieldgroupReadonly field to given value.


### GetIEzsignformfieldgroupMaxlength

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetIEzsignformfieldgroupMaxlength() int32`

GetIEzsignformfieldgroupMaxlength returns the IEzsignformfieldgroupMaxlength field if non-nil, zero value otherwise.

### GetIEzsignformfieldgroupMaxlengthOk

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetIEzsignformfieldgroupMaxlengthOk() (*int32, bool)`

GetIEzsignformfieldgroupMaxlengthOk returns a tuple with the IEzsignformfieldgroupMaxlength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignformfieldgroupMaxlength

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) SetIEzsignformfieldgroupMaxlength(v int32)`

SetIEzsignformfieldgroupMaxlength sets IEzsignformfieldgroupMaxlength field to given value.

### HasIEzsignformfieldgroupMaxlength

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) HasIEzsignformfieldgroupMaxlength() bool`

HasIEzsignformfieldgroupMaxlength returns a boolean if a field has been set.

### GetBEzsignformfieldgroupEncrypted

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetBEzsignformfieldgroupEncrypted() bool`

GetBEzsignformfieldgroupEncrypted returns the BEzsignformfieldgroupEncrypted field if non-nil, zero value otherwise.

### GetBEzsignformfieldgroupEncryptedOk

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetBEzsignformfieldgroupEncryptedOk() (*bool, bool)`

GetBEzsignformfieldgroupEncryptedOk returns a tuple with the BEzsignformfieldgroupEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignformfieldgroupEncrypted

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) SetBEzsignformfieldgroupEncrypted(v bool)`

SetBEzsignformfieldgroupEncrypted sets BEzsignformfieldgroupEncrypted field to given value.

### HasBEzsignformfieldgroupEncrypted

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) HasBEzsignformfieldgroupEncrypted() bool`

HasBEzsignformfieldgroupEncrypted returns a boolean if a field has been set.

### GetSEzsignformfieldgroupRegexp

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetSEzsignformfieldgroupRegexp() string`

GetSEzsignformfieldgroupRegexp returns the SEzsignformfieldgroupRegexp field if non-nil, zero value otherwise.

### GetSEzsignformfieldgroupRegexpOk

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetSEzsignformfieldgroupRegexpOk() (*string, bool)`

GetSEzsignformfieldgroupRegexpOk returns a tuple with the SEzsignformfieldgroupRegexp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignformfieldgroupRegexp

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) SetSEzsignformfieldgroupRegexp(v string)`

SetSEzsignformfieldgroupRegexp sets SEzsignformfieldgroupRegexp field to given value.

### HasSEzsignformfieldgroupRegexp

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) HasSEzsignformfieldgroupRegexp() bool`

HasSEzsignformfieldgroupRegexp returns a boolean if a field has been set.

### GetTEzsignformfieldgroupTooltip

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetTEzsignformfieldgroupTooltip() string`

GetTEzsignformfieldgroupTooltip returns the TEzsignformfieldgroupTooltip field if non-nil, zero value otherwise.

### GetTEzsignformfieldgroupTooltipOk

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetTEzsignformfieldgroupTooltipOk() (*string, bool)`

GetTEzsignformfieldgroupTooltipOk returns a tuple with the TEzsignformfieldgroupTooltip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEzsignformfieldgroupTooltip

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) SetTEzsignformfieldgroupTooltip(v string)`

SetTEzsignformfieldgroupTooltip sets TEzsignformfieldgroupTooltip field to given value.

### HasTEzsignformfieldgroupTooltip

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) HasTEzsignformfieldgroupTooltip() bool`

HasTEzsignformfieldgroupTooltip returns a boolean if a field has been set.

### GetEEzsignformfieldgroupTooltipposition

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetEEzsignformfieldgroupTooltipposition() FieldEEzsignformfieldgroupTooltipposition`

GetEEzsignformfieldgroupTooltipposition returns the EEzsignformfieldgroupTooltipposition field if non-nil, zero value otherwise.

### GetEEzsignformfieldgroupTooltippositionOk

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetEEzsignformfieldgroupTooltippositionOk() (*FieldEEzsignformfieldgroupTooltipposition, bool)`

GetEEzsignformfieldgroupTooltippositionOk returns a tuple with the EEzsignformfieldgroupTooltipposition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignformfieldgroupTooltipposition

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) SetEEzsignformfieldgroupTooltipposition(v FieldEEzsignformfieldgroupTooltipposition)`

SetEEzsignformfieldgroupTooltipposition sets EEzsignformfieldgroupTooltipposition field to given value.

### HasEEzsignformfieldgroupTooltipposition

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) HasEEzsignformfieldgroupTooltipposition() bool`

HasEEzsignformfieldgroupTooltipposition returns a boolean if a field has been set.

### GetEEzsignformfieldgroupTextvalidation

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetEEzsignformfieldgroupTextvalidation() EnumTextvalidation`

GetEEzsignformfieldgroupTextvalidation returns the EEzsignformfieldgroupTextvalidation field if non-nil, zero value otherwise.

### GetEEzsignformfieldgroupTextvalidationOk

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetEEzsignformfieldgroupTextvalidationOk() (*EnumTextvalidation, bool)`

GetEEzsignformfieldgroupTextvalidationOk returns a tuple with the EEzsignformfieldgroupTextvalidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignformfieldgroupTextvalidation

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) SetEEzsignformfieldgroupTextvalidation(v EnumTextvalidation)`

SetEEzsignformfieldgroupTextvalidation sets EEzsignformfieldgroupTextvalidation field to given value.

### HasEEzsignformfieldgroupTextvalidation

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) HasEEzsignformfieldgroupTextvalidation() bool`

HasEEzsignformfieldgroupTextvalidation returns a boolean if a field has been set.

### GetAObjEzsignformfieldgroupsigner

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetAObjEzsignformfieldgroupsigner() []EzsignformfieldgroupsignerRequestCompound`

GetAObjEzsignformfieldgroupsigner returns the AObjEzsignformfieldgroupsigner field if non-nil, zero value otherwise.

### GetAObjEzsignformfieldgroupsignerOk

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetAObjEzsignformfieldgroupsignerOk() (*[]EzsignformfieldgroupsignerRequestCompound, bool)`

GetAObjEzsignformfieldgroupsignerOk returns a tuple with the AObjEzsignformfieldgroupsigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzsignformfieldgroupsigner

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) SetAObjEzsignformfieldgroupsigner(v []EzsignformfieldgroupsignerRequestCompound)`

SetAObjEzsignformfieldgroupsigner sets AObjEzsignformfieldgroupsigner field to given value.


### GetAObjDropdownElement

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetAObjDropdownElement() []CustomDropdownElementRequestCompound`

GetAObjDropdownElement returns the AObjDropdownElement field if non-nil, zero value otherwise.

### GetAObjDropdownElementOk

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetAObjDropdownElementOk() (*[]CustomDropdownElementRequestCompound, bool)`

GetAObjDropdownElementOk returns a tuple with the AObjDropdownElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjDropdownElement

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) SetAObjDropdownElement(v []CustomDropdownElementRequestCompound)`

SetAObjDropdownElement sets AObjDropdownElement field to given value.

### HasAObjDropdownElement

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) HasAObjDropdownElement() bool`

HasAObjDropdownElement returns a boolean if a field has been set.

### GetAObjEzsignformfield

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetAObjEzsignformfield() []EzsignformfieldRequestCompound`

GetAObjEzsignformfield returns the AObjEzsignformfield field if non-nil, zero value otherwise.

### GetAObjEzsignformfieldOk

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetAObjEzsignformfieldOk() (*[]EzsignformfieldRequestCompound, bool)`

GetAObjEzsignformfieldOk returns a tuple with the AObjEzsignformfield field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzsignformfield

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) SetAObjEzsignformfield(v []EzsignformfieldRequestCompound)`

SetAObjEzsignformfield sets AObjEzsignformfield field to given value.


### GetObjCreateezsignelementspositionedbyword

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetObjCreateezsignelementspositionedbyword() CustomCreateEzsignelementsPositionedByWordRequest`

GetObjCreateezsignelementspositionedbyword returns the ObjCreateezsignelementspositionedbyword field if non-nil, zero value otherwise.

### GetObjCreateezsignelementspositionedbywordOk

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetObjCreateezsignelementspositionedbywordOk() (*CustomCreateEzsignelementsPositionedByWordRequest, bool)`

GetObjCreateezsignelementspositionedbywordOk returns a tuple with the ObjCreateezsignelementspositionedbyword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjCreateezsignelementspositionedbyword

`func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) SetObjCreateezsignelementspositionedbyword(v CustomCreateEzsignelementsPositionedByWordRequest)`

SetObjCreateezsignelementspositionedbyword sets ObjCreateezsignelementspositionedbyword field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


