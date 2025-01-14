# EzsigntemplatesignatureResponseCompoundV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsigntemplatesignatureID** | **int32** | The unique ID of the Ezsigntemplatesignature | 
**FkiEzsigntemplatedocumentID** | **int32** | The unique ID of the Ezsigntemplatedocument | 
**FkiEzsigntemplatesignerID** | **int32** | The unique ID of the Ezsigntemplatesigner | 
**FkiEzsigntemplatesignerIDValidation** | Pointer to **int32** | The unique ID of the Ezsigntemplatesigner | [optional] 
**BEzsigntemplatesignatureHandwritten** | Pointer to **bool** | Whether the Ezsigntemplatesignature must be handwritten or not when eEzsigntemplatesignatureType &#x3D; Signature. | [optional] 
**BEzsigntemplatesignatureReason** | Pointer to **bool** | Whether the Ezsigntemplatesignature must include a reason or not when eEzsigntemplatesignatureType &#x3D; Signature. | [optional] 
**EEzsigntemplatesignaturePositioning** | Pointer to [**FieldEEzsigntemplatesignaturePositioning**](FieldEEzsigntemplatesignaturePositioning.md) |  | [optional] 
**IEzsigntemplatedocumentpagePagenumber** | **int32** | The page number in the Ezsigntemplatedocument | 
**IEzsigntemplatesignatureX** | Pointer to **int32** | The X coordinate (Horizontal) where to put the Ezsigntemplatesignature on the page.  Coordinate is calculated at 100dpi (dot per inch). So for example, if you want to put the Ezsigntemplatesignature 2 inches from the left border of the page, you would use \&quot;200\&quot; for the X coordinate. | [optional] 
**IEzsigntemplatesignatureY** | Pointer to **int32** | The Y coordinate (Vertical) where to put the Ezsigntemplatesignature on the page.  Coordinate is calculated at 100dpi (dot per inch). So for example, if you want to put the Ezsigntemplatesignature 3 inches from the top border of the page, you would use \&quot;300\&quot; for the Y coordinate. | [optional] 
**IEzsigntemplatesignatureWidth** | Pointer to **int32** | The width of the Ezsigntemplatesignature.  Size is calculated at 100dpi (dot per inch). So for example, if you want the Ezsigntemplatesignature to have a width of 2 inches, you would use \&quot;200\&quot; for the iEzsigntemplatesignatureWidth. | [optional] 
**IEzsigntemplatesignatureHeight** | Pointer to **int32** | The height of the Ezsigntemplatesignature.  Size is calculated at 100dpi (dot per inch). So for example, if you want the Ezsigntemplatesignature to have an height of 2 inches, you would use \&quot;200\&quot; for the iEzsigntemplatesignatureHeight. | [optional] 
**IEzsigntemplatesignatureStep** | **int32** | The step when the Ezsigntemplatesigner will be invited to sign | 
**EEzsigntemplatesignatureType** | [**FieldEEzsigntemplatesignatureType**](FieldEEzsigntemplatesignatureType.md) |  | 
**EEzsigntemplatesignatureConsultationtrigger** | Pointer to [**FieldEEzsigntemplatesignatureConsultationtrigger**](FieldEEzsigntemplatesignatureConsultationtrigger.md) |  | [optional] 
**TEzsigntemplatesignatureTooltip** | Pointer to **string** | A tooltip that will be presented to Ezsigntemplatesigner about the Ezsigntemplatesignature | [optional] 
**EEzsigntemplatesignatureTooltipposition** | Pointer to [**FieldEEzsigntemplatesignatureTooltipposition**](FieldEEzsigntemplatesignatureTooltipposition.md) |  | [optional] 
**EEzsigntemplatesignatureFont** | Pointer to [**FieldEEzsigntemplatesignatureFont**](FieldEEzsigntemplatesignatureFont.md) |  | [optional] 
**IEzsigntemplatesignatureValidationstep** | Pointer to **int32** | The step when the Ezsigntemplatesigner will be invited to validate the Ezsigntemplatesignature of eEzsigntemplatesignatureType Attachments | [optional] 
**SEzsigntemplatesignatureAttachmentdescription** | Pointer to **string** | The description attached to the attachment name added in Ezsigntemplatesignature of eEzsigntemplatesignatureType Attachments | [optional] 
**EEzsigntemplatesignatureAttachmentnamesource** | Pointer to [**FieldEEzsigntemplatesignatureAttachmentnamesource**](FieldEEzsigntemplatesignatureAttachmentnamesource.md) |  | [optional] 
**BEzsigntemplatesignatureRequired** | Pointer to **bool** | Whether the Ezsigntemplatesignature is required or not. This field is relevant only with Ezsigntemplatesignature with eEzsigntemplatesignatureType &#x3D; Attachments. | [optional] 
**IEzsigntemplatesignatureMaxlength** | Pointer to **int32** | The maximum length for the value in the Ezsigntemplatesignature  This can only be set if eEzsigntemplatesignatureType is **FieldText** or **FieldTextarea** | [optional] 
**SEzsigntemplatesignatureDefaultvalue** | Pointer to **string** | The default value for the Ezsigntemplatesignature  You can use the codes below and they will be replaced at signature time.    | Code | Description | Example | | ------------------------- | ------------ | ------------ | | {sUserFirstname} | The first name of the contact | John | | {sUserLastname} | The last name of the contact | Doe | | {sUserJobtitle} | The job title | Sales Representative | | {sCompany} | Company name | eZmax Solutions Inc. | | {sEmailAddress} | The email address | email@example.com | | {sPhoneE164} | A phone number in E.164 Format | +15149901516 | | {sPhoneE164Cell} | A phone number in E.164 Format | +15149901516 | | [optional] 
**SEzsigntemplatesignatureRegexp** | Pointer to **string** | A regular expression to indicate what values are acceptable for the Ezsigntemplatesignature.  This can only be set if eEzsigntemplatesignatureType is **Text** or **Textarea** | [optional] 
**EEzsigntemplatesignatureTextvalidation** | Pointer to [**EnumTextvalidation**](EnumTextvalidation.md) |  | [optional] 
**SEzsigntemplatesignatureTextvalidationcustommessage** | Pointer to **string** | Description of validation rule. Show by signatory. | [optional] 
**EEzsigntemplatesignatureDependencyrequirement** | Pointer to [**FieldEEzsigntemplatesignatureDependencyrequirement**](FieldEEzsigntemplatesignatureDependencyrequirement.md) |  | [optional] 
**SEzsigntemplatesignaturePositioningpattern** | Pointer to **string** | The string pattern to search for the positioning. **This is not a regexp**  This will be required if **eEzsigntemplatesignaturePositioning** is set to **PerCoordinates** | [optional] 
**IEzsigntemplatesignaturePositioningoffsetx** | Pointer to **int32** | The offset X  This will be required if **eEzsigntemplatesignaturePositioning** is set to **PerCoordinates** | [optional] 
**IEzsigntemplatesignaturePositioningoffsety** | Pointer to **int32** | The offset Y  This will be required if **eEzsigntemplatesignaturePositioning** is set to **PerCoordinates** | [optional] 
**EEzsigntemplatesignaturePositioningoccurence** | Pointer to [**FieldEEzsigntemplatesignaturePositioningoccurence**](FieldEEzsigntemplatesignaturePositioningoccurence.md) |  | [optional] 
**BEzsigntemplatesignatureCustomdate** | Pointer to **bool** | Whether the Ezsigntemplatesignature has a custom date format or not. (Only possible when eEzsigntemplatesignatureType is **Name** or **Handwritten**) | [optional] 
**AObjEzsigntemplatesignaturecustomdate** | Pointer to [**[]EzsigntemplatesignaturecustomdateResponseCompoundV2**](EzsigntemplatesignaturecustomdateResponseCompoundV2.md) | An array of custom date blocks that will be filled at the time of signature.  Can only be used if bEzsigntemplatesignatureCustomdate is true.  Use an empty array if you don&#39;t want to have a date at all. | [optional] 
**AObjEzsigntemplateelementdependency** | Pointer to [**[]EzsigntemplateelementdependencyResponseCompound**](EzsigntemplateelementdependencyResponseCompound.md) |  | [optional] 

## Methods

### NewEzsigntemplatesignatureResponseCompoundV3

`func NewEzsigntemplatesignatureResponseCompoundV3(pkiEzsigntemplatesignatureID int32, fkiEzsigntemplatedocumentID int32, fkiEzsigntemplatesignerID int32, iEzsigntemplatedocumentpagePagenumber int32, iEzsigntemplatesignatureStep int32, eEzsigntemplatesignatureType FieldEEzsigntemplatesignatureType, ) *EzsigntemplatesignatureResponseCompoundV3`

NewEzsigntemplatesignatureResponseCompoundV3 instantiates a new EzsigntemplatesignatureResponseCompoundV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigntemplatesignatureResponseCompoundV3WithDefaults

`func NewEzsigntemplatesignatureResponseCompoundV3WithDefaults() *EzsigntemplatesignatureResponseCompoundV3`

NewEzsigntemplatesignatureResponseCompoundV3WithDefaults instantiates a new EzsigntemplatesignatureResponseCompoundV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsigntemplatesignatureID

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetPkiEzsigntemplatesignatureID() int32`

GetPkiEzsigntemplatesignatureID returns the PkiEzsigntemplatesignatureID field if non-nil, zero value otherwise.

### GetPkiEzsigntemplatesignatureIDOk

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetPkiEzsigntemplatesignatureIDOk() (*int32, bool)`

GetPkiEzsigntemplatesignatureIDOk returns a tuple with the PkiEzsigntemplatesignatureID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsigntemplatesignatureID

`func (o *EzsigntemplatesignatureResponseCompoundV3) SetPkiEzsigntemplatesignatureID(v int32)`

SetPkiEzsigntemplatesignatureID sets PkiEzsigntemplatesignatureID field to given value.


### GetFkiEzsigntemplatedocumentID

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetFkiEzsigntemplatedocumentID() int32`

GetFkiEzsigntemplatedocumentID returns the FkiEzsigntemplatedocumentID field if non-nil, zero value otherwise.

### GetFkiEzsigntemplatedocumentIDOk

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetFkiEzsigntemplatedocumentIDOk() (*int32, bool)`

GetFkiEzsigntemplatedocumentIDOk returns a tuple with the FkiEzsigntemplatedocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntemplatedocumentID

`func (o *EzsigntemplatesignatureResponseCompoundV3) SetFkiEzsigntemplatedocumentID(v int32)`

SetFkiEzsigntemplatedocumentID sets FkiEzsigntemplatedocumentID field to given value.


### GetFkiEzsigntemplatesignerID

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetFkiEzsigntemplatesignerID() int32`

GetFkiEzsigntemplatesignerID returns the FkiEzsigntemplatesignerID field if non-nil, zero value otherwise.

### GetFkiEzsigntemplatesignerIDOk

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetFkiEzsigntemplatesignerIDOk() (*int32, bool)`

GetFkiEzsigntemplatesignerIDOk returns a tuple with the FkiEzsigntemplatesignerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntemplatesignerID

`func (o *EzsigntemplatesignatureResponseCompoundV3) SetFkiEzsigntemplatesignerID(v int32)`

SetFkiEzsigntemplatesignerID sets FkiEzsigntemplatesignerID field to given value.


### GetFkiEzsigntemplatesignerIDValidation

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetFkiEzsigntemplatesignerIDValidation() int32`

GetFkiEzsigntemplatesignerIDValidation returns the FkiEzsigntemplatesignerIDValidation field if non-nil, zero value otherwise.

### GetFkiEzsigntemplatesignerIDValidationOk

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetFkiEzsigntemplatesignerIDValidationOk() (*int32, bool)`

GetFkiEzsigntemplatesignerIDValidationOk returns a tuple with the FkiEzsigntemplatesignerIDValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntemplatesignerIDValidation

`func (o *EzsigntemplatesignatureResponseCompoundV3) SetFkiEzsigntemplatesignerIDValidation(v int32)`

SetFkiEzsigntemplatesignerIDValidation sets FkiEzsigntemplatesignerIDValidation field to given value.

### HasFkiEzsigntemplatesignerIDValidation

`func (o *EzsigntemplatesignatureResponseCompoundV3) HasFkiEzsigntemplatesignerIDValidation() bool`

HasFkiEzsigntemplatesignerIDValidation returns a boolean if a field has been set.

### GetBEzsigntemplatesignatureHandwritten

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetBEzsigntemplatesignatureHandwritten() bool`

GetBEzsigntemplatesignatureHandwritten returns the BEzsigntemplatesignatureHandwritten field if non-nil, zero value otherwise.

### GetBEzsigntemplatesignatureHandwrittenOk

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetBEzsigntemplatesignatureHandwrittenOk() (*bool, bool)`

GetBEzsigntemplatesignatureHandwrittenOk returns a tuple with the BEzsigntemplatesignatureHandwritten field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigntemplatesignatureHandwritten

`func (o *EzsigntemplatesignatureResponseCompoundV3) SetBEzsigntemplatesignatureHandwritten(v bool)`

SetBEzsigntemplatesignatureHandwritten sets BEzsigntemplatesignatureHandwritten field to given value.

### HasBEzsigntemplatesignatureHandwritten

`func (o *EzsigntemplatesignatureResponseCompoundV3) HasBEzsigntemplatesignatureHandwritten() bool`

HasBEzsigntemplatesignatureHandwritten returns a boolean if a field has been set.

### GetBEzsigntemplatesignatureReason

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetBEzsigntemplatesignatureReason() bool`

GetBEzsigntemplatesignatureReason returns the BEzsigntemplatesignatureReason field if non-nil, zero value otherwise.

### GetBEzsigntemplatesignatureReasonOk

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetBEzsigntemplatesignatureReasonOk() (*bool, bool)`

GetBEzsigntemplatesignatureReasonOk returns a tuple with the BEzsigntemplatesignatureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigntemplatesignatureReason

`func (o *EzsigntemplatesignatureResponseCompoundV3) SetBEzsigntemplatesignatureReason(v bool)`

SetBEzsigntemplatesignatureReason sets BEzsigntemplatesignatureReason field to given value.

### HasBEzsigntemplatesignatureReason

`func (o *EzsigntemplatesignatureResponseCompoundV3) HasBEzsigntemplatesignatureReason() bool`

HasBEzsigntemplatesignatureReason returns a boolean if a field has been set.

### GetEEzsigntemplatesignaturePositioning

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetEEzsigntemplatesignaturePositioning() FieldEEzsigntemplatesignaturePositioning`

GetEEzsigntemplatesignaturePositioning returns the EEzsigntemplatesignaturePositioning field if non-nil, zero value otherwise.

### GetEEzsigntemplatesignaturePositioningOk

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetEEzsigntemplatesignaturePositioningOk() (*FieldEEzsigntemplatesignaturePositioning, bool)`

GetEEzsigntemplatesignaturePositioningOk returns a tuple with the EEzsigntemplatesignaturePositioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplatesignaturePositioning

`func (o *EzsigntemplatesignatureResponseCompoundV3) SetEEzsigntemplatesignaturePositioning(v FieldEEzsigntemplatesignaturePositioning)`

SetEEzsigntemplatesignaturePositioning sets EEzsigntemplatesignaturePositioning field to given value.

### HasEEzsigntemplatesignaturePositioning

`func (o *EzsigntemplatesignatureResponseCompoundV3) HasEEzsigntemplatesignaturePositioning() bool`

HasEEzsigntemplatesignaturePositioning returns a boolean if a field has been set.

### GetIEzsigntemplatedocumentpagePagenumber

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetIEzsigntemplatedocumentpagePagenumber() int32`

GetIEzsigntemplatedocumentpagePagenumber returns the IEzsigntemplatedocumentpagePagenumber field if non-nil, zero value otherwise.

### GetIEzsigntemplatedocumentpagePagenumberOk

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetIEzsigntemplatedocumentpagePagenumberOk() (*int32, bool)`

GetIEzsigntemplatedocumentpagePagenumberOk returns a tuple with the IEzsigntemplatedocumentpagePagenumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplatedocumentpagePagenumber

`func (o *EzsigntemplatesignatureResponseCompoundV3) SetIEzsigntemplatedocumentpagePagenumber(v int32)`

SetIEzsigntemplatedocumentpagePagenumber sets IEzsigntemplatedocumentpagePagenumber field to given value.


### GetIEzsigntemplatesignatureX

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetIEzsigntemplatesignatureX() int32`

GetIEzsigntemplatesignatureX returns the IEzsigntemplatesignatureX field if non-nil, zero value otherwise.

### GetIEzsigntemplatesignatureXOk

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetIEzsigntemplatesignatureXOk() (*int32, bool)`

GetIEzsigntemplatesignatureXOk returns a tuple with the IEzsigntemplatesignatureX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplatesignatureX

`func (o *EzsigntemplatesignatureResponseCompoundV3) SetIEzsigntemplatesignatureX(v int32)`

SetIEzsigntemplatesignatureX sets IEzsigntemplatesignatureX field to given value.

### HasIEzsigntemplatesignatureX

`func (o *EzsigntemplatesignatureResponseCompoundV3) HasIEzsigntemplatesignatureX() bool`

HasIEzsigntemplatesignatureX returns a boolean if a field has been set.

### GetIEzsigntemplatesignatureY

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetIEzsigntemplatesignatureY() int32`

GetIEzsigntemplatesignatureY returns the IEzsigntemplatesignatureY field if non-nil, zero value otherwise.

### GetIEzsigntemplatesignatureYOk

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetIEzsigntemplatesignatureYOk() (*int32, bool)`

GetIEzsigntemplatesignatureYOk returns a tuple with the IEzsigntemplatesignatureY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplatesignatureY

`func (o *EzsigntemplatesignatureResponseCompoundV3) SetIEzsigntemplatesignatureY(v int32)`

SetIEzsigntemplatesignatureY sets IEzsigntemplatesignatureY field to given value.

### HasIEzsigntemplatesignatureY

`func (o *EzsigntemplatesignatureResponseCompoundV3) HasIEzsigntemplatesignatureY() bool`

HasIEzsigntemplatesignatureY returns a boolean if a field has been set.

### GetIEzsigntemplatesignatureWidth

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetIEzsigntemplatesignatureWidth() int32`

GetIEzsigntemplatesignatureWidth returns the IEzsigntemplatesignatureWidth field if non-nil, zero value otherwise.

### GetIEzsigntemplatesignatureWidthOk

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetIEzsigntemplatesignatureWidthOk() (*int32, bool)`

GetIEzsigntemplatesignatureWidthOk returns a tuple with the IEzsigntemplatesignatureWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplatesignatureWidth

`func (o *EzsigntemplatesignatureResponseCompoundV3) SetIEzsigntemplatesignatureWidth(v int32)`

SetIEzsigntemplatesignatureWidth sets IEzsigntemplatesignatureWidth field to given value.

### HasIEzsigntemplatesignatureWidth

`func (o *EzsigntemplatesignatureResponseCompoundV3) HasIEzsigntemplatesignatureWidth() bool`

HasIEzsigntemplatesignatureWidth returns a boolean if a field has been set.

### GetIEzsigntemplatesignatureHeight

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetIEzsigntemplatesignatureHeight() int32`

GetIEzsigntemplatesignatureHeight returns the IEzsigntemplatesignatureHeight field if non-nil, zero value otherwise.

### GetIEzsigntemplatesignatureHeightOk

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetIEzsigntemplatesignatureHeightOk() (*int32, bool)`

GetIEzsigntemplatesignatureHeightOk returns a tuple with the IEzsigntemplatesignatureHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplatesignatureHeight

`func (o *EzsigntemplatesignatureResponseCompoundV3) SetIEzsigntemplatesignatureHeight(v int32)`

SetIEzsigntemplatesignatureHeight sets IEzsigntemplatesignatureHeight field to given value.

### HasIEzsigntemplatesignatureHeight

`func (o *EzsigntemplatesignatureResponseCompoundV3) HasIEzsigntemplatesignatureHeight() bool`

HasIEzsigntemplatesignatureHeight returns a boolean if a field has been set.

### GetIEzsigntemplatesignatureStep

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetIEzsigntemplatesignatureStep() int32`

GetIEzsigntemplatesignatureStep returns the IEzsigntemplatesignatureStep field if non-nil, zero value otherwise.

### GetIEzsigntemplatesignatureStepOk

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetIEzsigntemplatesignatureStepOk() (*int32, bool)`

GetIEzsigntemplatesignatureStepOk returns a tuple with the IEzsigntemplatesignatureStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplatesignatureStep

`func (o *EzsigntemplatesignatureResponseCompoundV3) SetIEzsigntemplatesignatureStep(v int32)`

SetIEzsigntemplatesignatureStep sets IEzsigntemplatesignatureStep field to given value.


### GetEEzsigntemplatesignatureType

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetEEzsigntemplatesignatureType() FieldEEzsigntemplatesignatureType`

GetEEzsigntemplatesignatureType returns the EEzsigntemplatesignatureType field if non-nil, zero value otherwise.

### GetEEzsigntemplatesignatureTypeOk

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetEEzsigntemplatesignatureTypeOk() (*FieldEEzsigntemplatesignatureType, bool)`

GetEEzsigntemplatesignatureTypeOk returns a tuple with the EEzsigntemplatesignatureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplatesignatureType

`func (o *EzsigntemplatesignatureResponseCompoundV3) SetEEzsigntemplatesignatureType(v FieldEEzsigntemplatesignatureType)`

SetEEzsigntemplatesignatureType sets EEzsigntemplatesignatureType field to given value.


### GetEEzsigntemplatesignatureConsultationtrigger

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetEEzsigntemplatesignatureConsultationtrigger() FieldEEzsigntemplatesignatureConsultationtrigger`

GetEEzsigntemplatesignatureConsultationtrigger returns the EEzsigntemplatesignatureConsultationtrigger field if non-nil, zero value otherwise.

### GetEEzsigntemplatesignatureConsultationtriggerOk

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetEEzsigntemplatesignatureConsultationtriggerOk() (*FieldEEzsigntemplatesignatureConsultationtrigger, bool)`

GetEEzsigntemplatesignatureConsultationtriggerOk returns a tuple with the EEzsigntemplatesignatureConsultationtrigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplatesignatureConsultationtrigger

`func (o *EzsigntemplatesignatureResponseCompoundV3) SetEEzsigntemplatesignatureConsultationtrigger(v FieldEEzsigntemplatesignatureConsultationtrigger)`

SetEEzsigntemplatesignatureConsultationtrigger sets EEzsigntemplatesignatureConsultationtrigger field to given value.

### HasEEzsigntemplatesignatureConsultationtrigger

`func (o *EzsigntemplatesignatureResponseCompoundV3) HasEEzsigntemplatesignatureConsultationtrigger() bool`

HasEEzsigntemplatesignatureConsultationtrigger returns a boolean if a field has been set.

### GetTEzsigntemplatesignatureTooltip

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetTEzsigntemplatesignatureTooltip() string`

GetTEzsigntemplatesignatureTooltip returns the TEzsigntemplatesignatureTooltip field if non-nil, zero value otherwise.

### GetTEzsigntemplatesignatureTooltipOk

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetTEzsigntemplatesignatureTooltipOk() (*string, bool)`

GetTEzsigntemplatesignatureTooltipOk returns a tuple with the TEzsigntemplatesignatureTooltip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEzsigntemplatesignatureTooltip

`func (o *EzsigntemplatesignatureResponseCompoundV3) SetTEzsigntemplatesignatureTooltip(v string)`

SetTEzsigntemplatesignatureTooltip sets TEzsigntemplatesignatureTooltip field to given value.

### HasTEzsigntemplatesignatureTooltip

`func (o *EzsigntemplatesignatureResponseCompoundV3) HasTEzsigntemplatesignatureTooltip() bool`

HasTEzsigntemplatesignatureTooltip returns a boolean if a field has been set.

### GetEEzsigntemplatesignatureTooltipposition

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetEEzsigntemplatesignatureTooltipposition() FieldEEzsigntemplatesignatureTooltipposition`

GetEEzsigntemplatesignatureTooltipposition returns the EEzsigntemplatesignatureTooltipposition field if non-nil, zero value otherwise.

### GetEEzsigntemplatesignatureTooltippositionOk

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetEEzsigntemplatesignatureTooltippositionOk() (*FieldEEzsigntemplatesignatureTooltipposition, bool)`

GetEEzsigntemplatesignatureTooltippositionOk returns a tuple with the EEzsigntemplatesignatureTooltipposition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplatesignatureTooltipposition

`func (o *EzsigntemplatesignatureResponseCompoundV3) SetEEzsigntemplatesignatureTooltipposition(v FieldEEzsigntemplatesignatureTooltipposition)`

SetEEzsigntemplatesignatureTooltipposition sets EEzsigntemplatesignatureTooltipposition field to given value.

### HasEEzsigntemplatesignatureTooltipposition

`func (o *EzsigntemplatesignatureResponseCompoundV3) HasEEzsigntemplatesignatureTooltipposition() bool`

HasEEzsigntemplatesignatureTooltipposition returns a boolean if a field has been set.

### GetEEzsigntemplatesignatureFont

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetEEzsigntemplatesignatureFont() FieldEEzsigntemplatesignatureFont`

GetEEzsigntemplatesignatureFont returns the EEzsigntemplatesignatureFont field if non-nil, zero value otherwise.

### GetEEzsigntemplatesignatureFontOk

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetEEzsigntemplatesignatureFontOk() (*FieldEEzsigntemplatesignatureFont, bool)`

GetEEzsigntemplatesignatureFontOk returns a tuple with the EEzsigntemplatesignatureFont field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplatesignatureFont

`func (o *EzsigntemplatesignatureResponseCompoundV3) SetEEzsigntemplatesignatureFont(v FieldEEzsigntemplatesignatureFont)`

SetEEzsigntemplatesignatureFont sets EEzsigntemplatesignatureFont field to given value.

### HasEEzsigntemplatesignatureFont

`func (o *EzsigntemplatesignatureResponseCompoundV3) HasEEzsigntemplatesignatureFont() bool`

HasEEzsigntemplatesignatureFont returns a boolean if a field has been set.

### GetIEzsigntemplatesignatureValidationstep

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetIEzsigntemplatesignatureValidationstep() int32`

GetIEzsigntemplatesignatureValidationstep returns the IEzsigntemplatesignatureValidationstep field if non-nil, zero value otherwise.

### GetIEzsigntemplatesignatureValidationstepOk

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetIEzsigntemplatesignatureValidationstepOk() (*int32, bool)`

GetIEzsigntemplatesignatureValidationstepOk returns a tuple with the IEzsigntemplatesignatureValidationstep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplatesignatureValidationstep

`func (o *EzsigntemplatesignatureResponseCompoundV3) SetIEzsigntemplatesignatureValidationstep(v int32)`

SetIEzsigntemplatesignatureValidationstep sets IEzsigntemplatesignatureValidationstep field to given value.

### HasIEzsigntemplatesignatureValidationstep

`func (o *EzsigntemplatesignatureResponseCompoundV3) HasIEzsigntemplatesignatureValidationstep() bool`

HasIEzsigntemplatesignatureValidationstep returns a boolean if a field has been set.

### GetSEzsigntemplatesignatureAttachmentdescription

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetSEzsigntemplatesignatureAttachmentdescription() string`

GetSEzsigntemplatesignatureAttachmentdescription returns the SEzsigntemplatesignatureAttachmentdescription field if non-nil, zero value otherwise.

### GetSEzsigntemplatesignatureAttachmentdescriptionOk

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetSEzsigntemplatesignatureAttachmentdescriptionOk() (*string, bool)`

GetSEzsigntemplatesignatureAttachmentdescriptionOk returns a tuple with the SEzsigntemplatesignatureAttachmentdescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplatesignatureAttachmentdescription

`func (o *EzsigntemplatesignatureResponseCompoundV3) SetSEzsigntemplatesignatureAttachmentdescription(v string)`

SetSEzsigntemplatesignatureAttachmentdescription sets SEzsigntemplatesignatureAttachmentdescription field to given value.

### HasSEzsigntemplatesignatureAttachmentdescription

`func (o *EzsigntemplatesignatureResponseCompoundV3) HasSEzsigntemplatesignatureAttachmentdescription() bool`

HasSEzsigntemplatesignatureAttachmentdescription returns a boolean if a field has been set.

### GetEEzsigntemplatesignatureAttachmentnamesource

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetEEzsigntemplatesignatureAttachmentnamesource() FieldEEzsigntemplatesignatureAttachmentnamesource`

GetEEzsigntemplatesignatureAttachmentnamesource returns the EEzsigntemplatesignatureAttachmentnamesource field if non-nil, zero value otherwise.

### GetEEzsigntemplatesignatureAttachmentnamesourceOk

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetEEzsigntemplatesignatureAttachmentnamesourceOk() (*FieldEEzsigntemplatesignatureAttachmentnamesource, bool)`

GetEEzsigntemplatesignatureAttachmentnamesourceOk returns a tuple with the EEzsigntemplatesignatureAttachmentnamesource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplatesignatureAttachmentnamesource

`func (o *EzsigntemplatesignatureResponseCompoundV3) SetEEzsigntemplatesignatureAttachmentnamesource(v FieldEEzsigntemplatesignatureAttachmentnamesource)`

SetEEzsigntemplatesignatureAttachmentnamesource sets EEzsigntemplatesignatureAttachmentnamesource field to given value.

### HasEEzsigntemplatesignatureAttachmentnamesource

`func (o *EzsigntemplatesignatureResponseCompoundV3) HasEEzsigntemplatesignatureAttachmentnamesource() bool`

HasEEzsigntemplatesignatureAttachmentnamesource returns a boolean if a field has been set.

### GetBEzsigntemplatesignatureRequired

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetBEzsigntemplatesignatureRequired() bool`

GetBEzsigntemplatesignatureRequired returns the BEzsigntemplatesignatureRequired field if non-nil, zero value otherwise.

### GetBEzsigntemplatesignatureRequiredOk

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetBEzsigntemplatesignatureRequiredOk() (*bool, bool)`

GetBEzsigntemplatesignatureRequiredOk returns a tuple with the BEzsigntemplatesignatureRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigntemplatesignatureRequired

`func (o *EzsigntemplatesignatureResponseCompoundV3) SetBEzsigntemplatesignatureRequired(v bool)`

SetBEzsigntemplatesignatureRequired sets BEzsigntemplatesignatureRequired field to given value.

### HasBEzsigntemplatesignatureRequired

`func (o *EzsigntemplatesignatureResponseCompoundV3) HasBEzsigntemplatesignatureRequired() bool`

HasBEzsigntemplatesignatureRequired returns a boolean if a field has been set.

### GetIEzsigntemplatesignatureMaxlength

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetIEzsigntemplatesignatureMaxlength() int32`

GetIEzsigntemplatesignatureMaxlength returns the IEzsigntemplatesignatureMaxlength field if non-nil, zero value otherwise.

### GetIEzsigntemplatesignatureMaxlengthOk

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetIEzsigntemplatesignatureMaxlengthOk() (*int32, bool)`

GetIEzsigntemplatesignatureMaxlengthOk returns a tuple with the IEzsigntemplatesignatureMaxlength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplatesignatureMaxlength

`func (o *EzsigntemplatesignatureResponseCompoundV3) SetIEzsigntemplatesignatureMaxlength(v int32)`

SetIEzsigntemplatesignatureMaxlength sets IEzsigntemplatesignatureMaxlength field to given value.

### HasIEzsigntemplatesignatureMaxlength

`func (o *EzsigntemplatesignatureResponseCompoundV3) HasIEzsigntemplatesignatureMaxlength() bool`

HasIEzsigntemplatesignatureMaxlength returns a boolean if a field has been set.

### GetSEzsigntemplatesignatureDefaultvalue

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetSEzsigntemplatesignatureDefaultvalue() string`

GetSEzsigntemplatesignatureDefaultvalue returns the SEzsigntemplatesignatureDefaultvalue field if non-nil, zero value otherwise.

### GetSEzsigntemplatesignatureDefaultvalueOk

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetSEzsigntemplatesignatureDefaultvalueOk() (*string, bool)`

GetSEzsigntemplatesignatureDefaultvalueOk returns a tuple with the SEzsigntemplatesignatureDefaultvalue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplatesignatureDefaultvalue

`func (o *EzsigntemplatesignatureResponseCompoundV3) SetSEzsigntemplatesignatureDefaultvalue(v string)`

SetSEzsigntemplatesignatureDefaultvalue sets SEzsigntemplatesignatureDefaultvalue field to given value.

### HasSEzsigntemplatesignatureDefaultvalue

`func (o *EzsigntemplatesignatureResponseCompoundV3) HasSEzsigntemplatesignatureDefaultvalue() bool`

HasSEzsigntemplatesignatureDefaultvalue returns a boolean if a field has been set.

### GetSEzsigntemplatesignatureRegexp

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetSEzsigntemplatesignatureRegexp() string`

GetSEzsigntemplatesignatureRegexp returns the SEzsigntemplatesignatureRegexp field if non-nil, zero value otherwise.

### GetSEzsigntemplatesignatureRegexpOk

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetSEzsigntemplatesignatureRegexpOk() (*string, bool)`

GetSEzsigntemplatesignatureRegexpOk returns a tuple with the SEzsigntemplatesignatureRegexp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplatesignatureRegexp

`func (o *EzsigntemplatesignatureResponseCompoundV3) SetSEzsigntemplatesignatureRegexp(v string)`

SetSEzsigntemplatesignatureRegexp sets SEzsigntemplatesignatureRegexp field to given value.

### HasSEzsigntemplatesignatureRegexp

`func (o *EzsigntemplatesignatureResponseCompoundV3) HasSEzsigntemplatesignatureRegexp() bool`

HasSEzsigntemplatesignatureRegexp returns a boolean if a field has been set.

### GetEEzsigntemplatesignatureTextvalidation

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetEEzsigntemplatesignatureTextvalidation() EnumTextvalidation`

GetEEzsigntemplatesignatureTextvalidation returns the EEzsigntemplatesignatureTextvalidation field if non-nil, zero value otherwise.

### GetEEzsigntemplatesignatureTextvalidationOk

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetEEzsigntemplatesignatureTextvalidationOk() (*EnumTextvalidation, bool)`

GetEEzsigntemplatesignatureTextvalidationOk returns a tuple with the EEzsigntemplatesignatureTextvalidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplatesignatureTextvalidation

`func (o *EzsigntemplatesignatureResponseCompoundV3) SetEEzsigntemplatesignatureTextvalidation(v EnumTextvalidation)`

SetEEzsigntemplatesignatureTextvalidation sets EEzsigntemplatesignatureTextvalidation field to given value.

### HasEEzsigntemplatesignatureTextvalidation

`func (o *EzsigntemplatesignatureResponseCompoundV3) HasEEzsigntemplatesignatureTextvalidation() bool`

HasEEzsigntemplatesignatureTextvalidation returns a boolean if a field has been set.

### GetSEzsigntemplatesignatureTextvalidationcustommessage

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetSEzsigntemplatesignatureTextvalidationcustommessage() string`

GetSEzsigntemplatesignatureTextvalidationcustommessage returns the SEzsigntemplatesignatureTextvalidationcustommessage field if non-nil, zero value otherwise.

### GetSEzsigntemplatesignatureTextvalidationcustommessageOk

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetSEzsigntemplatesignatureTextvalidationcustommessageOk() (*string, bool)`

GetSEzsigntemplatesignatureTextvalidationcustommessageOk returns a tuple with the SEzsigntemplatesignatureTextvalidationcustommessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplatesignatureTextvalidationcustommessage

`func (o *EzsigntemplatesignatureResponseCompoundV3) SetSEzsigntemplatesignatureTextvalidationcustommessage(v string)`

SetSEzsigntemplatesignatureTextvalidationcustommessage sets SEzsigntemplatesignatureTextvalidationcustommessage field to given value.

### HasSEzsigntemplatesignatureTextvalidationcustommessage

`func (o *EzsigntemplatesignatureResponseCompoundV3) HasSEzsigntemplatesignatureTextvalidationcustommessage() bool`

HasSEzsigntemplatesignatureTextvalidationcustommessage returns a boolean if a field has been set.

### GetEEzsigntemplatesignatureDependencyrequirement

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetEEzsigntemplatesignatureDependencyrequirement() FieldEEzsigntemplatesignatureDependencyrequirement`

GetEEzsigntemplatesignatureDependencyrequirement returns the EEzsigntemplatesignatureDependencyrequirement field if non-nil, zero value otherwise.

### GetEEzsigntemplatesignatureDependencyrequirementOk

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetEEzsigntemplatesignatureDependencyrequirementOk() (*FieldEEzsigntemplatesignatureDependencyrequirement, bool)`

GetEEzsigntemplatesignatureDependencyrequirementOk returns a tuple with the EEzsigntemplatesignatureDependencyrequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplatesignatureDependencyrequirement

`func (o *EzsigntemplatesignatureResponseCompoundV3) SetEEzsigntemplatesignatureDependencyrequirement(v FieldEEzsigntemplatesignatureDependencyrequirement)`

SetEEzsigntemplatesignatureDependencyrequirement sets EEzsigntemplatesignatureDependencyrequirement field to given value.

### HasEEzsigntemplatesignatureDependencyrequirement

`func (o *EzsigntemplatesignatureResponseCompoundV3) HasEEzsigntemplatesignatureDependencyrequirement() bool`

HasEEzsigntemplatesignatureDependencyrequirement returns a boolean if a field has been set.

### GetSEzsigntemplatesignaturePositioningpattern

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetSEzsigntemplatesignaturePositioningpattern() string`

GetSEzsigntemplatesignaturePositioningpattern returns the SEzsigntemplatesignaturePositioningpattern field if non-nil, zero value otherwise.

### GetSEzsigntemplatesignaturePositioningpatternOk

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetSEzsigntemplatesignaturePositioningpatternOk() (*string, bool)`

GetSEzsigntemplatesignaturePositioningpatternOk returns a tuple with the SEzsigntemplatesignaturePositioningpattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplatesignaturePositioningpattern

`func (o *EzsigntemplatesignatureResponseCompoundV3) SetSEzsigntemplatesignaturePositioningpattern(v string)`

SetSEzsigntemplatesignaturePositioningpattern sets SEzsigntemplatesignaturePositioningpattern field to given value.

### HasSEzsigntemplatesignaturePositioningpattern

`func (o *EzsigntemplatesignatureResponseCompoundV3) HasSEzsigntemplatesignaturePositioningpattern() bool`

HasSEzsigntemplatesignaturePositioningpattern returns a boolean if a field has been set.

### GetIEzsigntemplatesignaturePositioningoffsetx

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetIEzsigntemplatesignaturePositioningoffsetx() int32`

GetIEzsigntemplatesignaturePositioningoffsetx returns the IEzsigntemplatesignaturePositioningoffsetx field if non-nil, zero value otherwise.

### GetIEzsigntemplatesignaturePositioningoffsetxOk

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetIEzsigntemplatesignaturePositioningoffsetxOk() (*int32, bool)`

GetIEzsigntemplatesignaturePositioningoffsetxOk returns a tuple with the IEzsigntemplatesignaturePositioningoffsetx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplatesignaturePositioningoffsetx

`func (o *EzsigntemplatesignatureResponseCompoundV3) SetIEzsigntemplatesignaturePositioningoffsetx(v int32)`

SetIEzsigntemplatesignaturePositioningoffsetx sets IEzsigntemplatesignaturePositioningoffsetx field to given value.

### HasIEzsigntemplatesignaturePositioningoffsetx

`func (o *EzsigntemplatesignatureResponseCompoundV3) HasIEzsigntemplatesignaturePositioningoffsetx() bool`

HasIEzsigntemplatesignaturePositioningoffsetx returns a boolean if a field has been set.

### GetIEzsigntemplatesignaturePositioningoffsety

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetIEzsigntemplatesignaturePositioningoffsety() int32`

GetIEzsigntemplatesignaturePositioningoffsety returns the IEzsigntemplatesignaturePositioningoffsety field if non-nil, zero value otherwise.

### GetIEzsigntemplatesignaturePositioningoffsetyOk

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetIEzsigntemplatesignaturePositioningoffsetyOk() (*int32, bool)`

GetIEzsigntemplatesignaturePositioningoffsetyOk returns a tuple with the IEzsigntemplatesignaturePositioningoffsety field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplatesignaturePositioningoffsety

`func (o *EzsigntemplatesignatureResponseCompoundV3) SetIEzsigntemplatesignaturePositioningoffsety(v int32)`

SetIEzsigntemplatesignaturePositioningoffsety sets IEzsigntemplatesignaturePositioningoffsety field to given value.

### HasIEzsigntemplatesignaturePositioningoffsety

`func (o *EzsigntemplatesignatureResponseCompoundV3) HasIEzsigntemplatesignaturePositioningoffsety() bool`

HasIEzsigntemplatesignaturePositioningoffsety returns a boolean if a field has been set.

### GetEEzsigntemplatesignaturePositioningoccurence

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetEEzsigntemplatesignaturePositioningoccurence() FieldEEzsigntemplatesignaturePositioningoccurence`

GetEEzsigntemplatesignaturePositioningoccurence returns the EEzsigntemplatesignaturePositioningoccurence field if non-nil, zero value otherwise.

### GetEEzsigntemplatesignaturePositioningoccurenceOk

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetEEzsigntemplatesignaturePositioningoccurenceOk() (*FieldEEzsigntemplatesignaturePositioningoccurence, bool)`

GetEEzsigntemplatesignaturePositioningoccurenceOk returns a tuple with the EEzsigntemplatesignaturePositioningoccurence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplatesignaturePositioningoccurence

`func (o *EzsigntemplatesignatureResponseCompoundV3) SetEEzsigntemplatesignaturePositioningoccurence(v FieldEEzsigntemplatesignaturePositioningoccurence)`

SetEEzsigntemplatesignaturePositioningoccurence sets EEzsigntemplatesignaturePositioningoccurence field to given value.

### HasEEzsigntemplatesignaturePositioningoccurence

`func (o *EzsigntemplatesignatureResponseCompoundV3) HasEEzsigntemplatesignaturePositioningoccurence() bool`

HasEEzsigntemplatesignaturePositioningoccurence returns a boolean if a field has been set.

### GetBEzsigntemplatesignatureCustomdate

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetBEzsigntemplatesignatureCustomdate() bool`

GetBEzsigntemplatesignatureCustomdate returns the BEzsigntemplatesignatureCustomdate field if non-nil, zero value otherwise.

### GetBEzsigntemplatesignatureCustomdateOk

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetBEzsigntemplatesignatureCustomdateOk() (*bool, bool)`

GetBEzsigntemplatesignatureCustomdateOk returns a tuple with the BEzsigntemplatesignatureCustomdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigntemplatesignatureCustomdate

`func (o *EzsigntemplatesignatureResponseCompoundV3) SetBEzsigntemplatesignatureCustomdate(v bool)`

SetBEzsigntemplatesignatureCustomdate sets BEzsigntemplatesignatureCustomdate field to given value.

### HasBEzsigntemplatesignatureCustomdate

`func (o *EzsigntemplatesignatureResponseCompoundV3) HasBEzsigntemplatesignatureCustomdate() bool`

HasBEzsigntemplatesignatureCustomdate returns a boolean if a field has been set.

### GetAObjEzsigntemplatesignaturecustomdate

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetAObjEzsigntemplatesignaturecustomdate() []EzsigntemplatesignaturecustomdateResponseCompoundV2`

GetAObjEzsigntemplatesignaturecustomdate returns the AObjEzsigntemplatesignaturecustomdate field if non-nil, zero value otherwise.

### GetAObjEzsigntemplatesignaturecustomdateOk

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetAObjEzsigntemplatesignaturecustomdateOk() (*[]EzsigntemplatesignaturecustomdateResponseCompoundV2, bool)`

GetAObjEzsigntemplatesignaturecustomdateOk returns a tuple with the AObjEzsigntemplatesignaturecustomdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzsigntemplatesignaturecustomdate

`func (o *EzsigntemplatesignatureResponseCompoundV3) SetAObjEzsigntemplatesignaturecustomdate(v []EzsigntemplatesignaturecustomdateResponseCompoundV2)`

SetAObjEzsigntemplatesignaturecustomdate sets AObjEzsigntemplatesignaturecustomdate field to given value.

### HasAObjEzsigntemplatesignaturecustomdate

`func (o *EzsigntemplatesignatureResponseCompoundV3) HasAObjEzsigntemplatesignaturecustomdate() bool`

HasAObjEzsigntemplatesignaturecustomdate returns a boolean if a field has been set.

### GetAObjEzsigntemplateelementdependency

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetAObjEzsigntemplateelementdependency() []EzsigntemplateelementdependencyResponseCompound`

GetAObjEzsigntemplateelementdependency returns the AObjEzsigntemplateelementdependency field if non-nil, zero value otherwise.

### GetAObjEzsigntemplateelementdependencyOk

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetAObjEzsigntemplateelementdependencyOk() (*[]EzsigntemplateelementdependencyResponseCompound, bool)`

GetAObjEzsigntemplateelementdependencyOk returns a tuple with the AObjEzsigntemplateelementdependency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzsigntemplateelementdependency

`func (o *EzsigntemplatesignatureResponseCompoundV3) SetAObjEzsigntemplateelementdependency(v []EzsigntemplateelementdependencyResponseCompound)`

SetAObjEzsigntemplateelementdependency sets AObjEzsigntemplateelementdependency field to given value.

### HasAObjEzsigntemplateelementdependency

`func (o *EzsigntemplatesignatureResponseCompoundV3) HasAObjEzsigntemplateelementdependency() bool`

HasAObjEzsigntemplateelementdependency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


