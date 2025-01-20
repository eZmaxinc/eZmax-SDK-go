# EzsigntemplatesignatureRequestCompoundV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsigntemplatesignatureID** | Pointer to **int32** | The unique ID of the Ezsigntemplatesignature | [optional] 
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
**BEzsigntemplatesignatureRequired** | Pointer to **bool** | Whether the Ezsigntemplatesignature is required or not. This field is relevant only with Ezsigntemplatesignature with eEzsigntemplatesignatureType &#x3D; Attachments. | [optional] 
**EEzsigntemplatesignatureAttachmentnamesource** | Pointer to [**FieldEEzsigntemplatesignatureAttachmentnamesource**](FieldEEzsigntemplatesignatureAttachmentnamesource.md) |  | [optional] 
**SEzsigntemplatesignatureAttachmentdescription** | Pointer to **string** | The description attached to the attachment name added in Ezsigntemplatesignature of eEzsigntemplatesignatureType Attachments | [optional] 
**IEzsigntemplatesignatureValidationstep** | Pointer to **int32** | The step when the Ezsigntemplatesigner will be invited to validate the Ezsigntemplatesignature of eEzsigntemplatesignatureType Attachments | [optional] 
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
**AObjEzsigntemplatesignaturecustomdate** | Pointer to [**[]EzsigntemplatesignaturecustomdateRequestV2**](EzsigntemplatesignaturecustomdateRequestV2.md) | An array of custom date blocks that will be filled at the time of signature.  Can only be used if bEzsigntemplatesignatureCustomdate is true.  Use an empty array if you don&#39;t want to have a date at all. | [optional] 
**AObjEzsigntemplateelementdependency** | Pointer to [**[]EzsigntemplateelementdependencyRequest**](EzsigntemplateelementdependencyRequest.md) |  | [optional] 

## Methods

### NewEzsigntemplatesignatureRequestCompoundV2

`func NewEzsigntemplatesignatureRequestCompoundV2(fkiEzsigntemplatedocumentID int32, fkiEzsigntemplatesignerID int32, iEzsigntemplatedocumentpagePagenumber int32, iEzsigntemplatesignatureStep int32, eEzsigntemplatesignatureType FieldEEzsigntemplatesignatureType, ) *EzsigntemplatesignatureRequestCompoundV2`

NewEzsigntemplatesignatureRequestCompoundV2 instantiates a new EzsigntemplatesignatureRequestCompoundV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigntemplatesignatureRequestCompoundV2WithDefaults

`func NewEzsigntemplatesignatureRequestCompoundV2WithDefaults() *EzsigntemplatesignatureRequestCompoundV2`

NewEzsigntemplatesignatureRequestCompoundV2WithDefaults instantiates a new EzsigntemplatesignatureRequestCompoundV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsigntemplatesignatureID

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetPkiEzsigntemplatesignatureID() int32`

GetPkiEzsigntemplatesignatureID returns the PkiEzsigntemplatesignatureID field if non-nil, zero value otherwise.

### GetPkiEzsigntemplatesignatureIDOk

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetPkiEzsigntemplatesignatureIDOk() (*int32, bool)`

GetPkiEzsigntemplatesignatureIDOk returns a tuple with the PkiEzsigntemplatesignatureID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsigntemplatesignatureID

`func (o *EzsigntemplatesignatureRequestCompoundV2) SetPkiEzsigntemplatesignatureID(v int32)`

SetPkiEzsigntemplatesignatureID sets PkiEzsigntemplatesignatureID field to given value.

### HasPkiEzsigntemplatesignatureID

`func (o *EzsigntemplatesignatureRequestCompoundV2) HasPkiEzsigntemplatesignatureID() bool`

HasPkiEzsigntemplatesignatureID returns a boolean if a field has been set.

### GetFkiEzsigntemplatedocumentID

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetFkiEzsigntemplatedocumentID() int32`

GetFkiEzsigntemplatedocumentID returns the FkiEzsigntemplatedocumentID field if non-nil, zero value otherwise.

### GetFkiEzsigntemplatedocumentIDOk

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetFkiEzsigntemplatedocumentIDOk() (*int32, bool)`

GetFkiEzsigntemplatedocumentIDOk returns a tuple with the FkiEzsigntemplatedocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntemplatedocumentID

`func (o *EzsigntemplatesignatureRequestCompoundV2) SetFkiEzsigntemplatedocumentID(v int32)`

SetFkiEzsigntemplatedocumentID sets FkiEzsigntemplatedocumentID field to given value.


### GetFkiEzsigntemplatesignerID

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetFkiEzsigntemplatesignerID() int32`

GetFkiEzsigntemplatesignerID returns the FkiEzsigntemplatesignerID field if non-nil, zero value otherwise.

### GetFkiEzsigntemplatesignerIDOk

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetFkiEzsigntemplatesignerIDOk() (*int32, bool)`

GetFkiEzsigntemplatesignerIDOk returns a tuple with the FkiEzsigntemplatesignerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntemplatesignerID

`func (o *EzsigntemplatesignatureRequestCompoundV2) SetFkiEzsigntemplatesignerID(v int32)`

SetFkiEzsigntemplatesignerID sets FkiEzsigntemplatesignerID field to given value.


### GetFkiEzsigntemplatesignerIDValidation

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetFkiEzsigntemplatesignerIDValidation() int32`

GetFkiEzsigntemplatesignerIDValidation returns the FkiEzsigntemplatesignerIDValidation field if non-nil, zero value otherwise.

### GetFkiEzsigntemplatesignerIDValidationOk

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetFkiEzsigntemplatesignerIDValidationOk() (*int32, bool)`

GetFkiEzsigntemplatesignerIDValidationOk returns a tuple with the FkiEzsigntemplatesignerIDValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntemplatesignerIDValidation

`func (o *EzsigntemplatesignatureRequestCompoundV2) SetFkiEzsigntemplatesignerIDValidation(v int32)`

SetFkiEzsigntemplatesignerIDValidation sets FkiEzsigntemplatesignerIDValidation field to given value.

### HasFkiEzsigntemplatesignerIDValidation

`func (o *EzsigntemplatesignatureRequestCompoundV2) HasFkiEzsigntemplatesignerIDValidation() bool`

HasFkiEzsigntemplatesignerIDValidation returns a boolean if a field has been set.

### GetBEzsigntemplatesignatureHandwritten

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetBEzsigntemplatesignatureHandwritten() bool`

GetBEzsigntemplatesignatureHandwritten returns the BEzsigntemplatesignatureHandwritten field if non-nil, zero value otherwise.

### GetBEzsigntemplatesignatureHandwrittenOk

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetBEzsigntemplatesignatureHandwrittenOk() (*bool, bool)`

GetBEzsigntemplatesignatureHandwrittenOk returns a tuple with the BEzsigntemplatesignatureHandwritten field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigntemplatesignatureHandwritten

`func (o *EzsigntemplatesignatureRequestCompoundV2) SetBEzsigntemplatesignatureHandwritten(v bool)`

SetBEzsigntemplatesignatureHandwritten sets BEzsigntemplatesignatureHandwritten field to given value.

### HasBEzsigntemplatesignatureHandwritten

`func (o *EzsigntemplatesignatureRequestCompoundV2) HasBEzsigntemplatesignatureHandwritten() bool`

HasBEzsigntemplatesignatureHandwritten returns a boolean if a field has been set.

### GetBEzsigntemplatesignatureReason

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetBEzsigntemplatesignatureReason() bool`

GetBEzsigntemplatesignatureReason returns the BEzsigntemplatesignatureReason field if non-nil, zero value otherwise.

### GetBEzsigntemplatesignatureReasonOk

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetBEzsigntemplatesignatureReasonOk() (*bool, bool)`

GetBEzsigntemplatesignatureReasonOk returns a tuple with the BEzsigntemplatesignatureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigntemplatesignatureReason

`func (o *EzsigntemplatesignatureRequestCompoundV2) SetBEzsigntemplatesignatureReason(v bool)`

SetBEzsigntemplatesignatureReason sets BEzsigntemplatesignatureReason field to given value.

### HasBEzsigntemplatesignatureReason

`func (o *EzsigntemplatesignatureRequestCompoundV2) HasBEzsigntemplatesignatureReason() bool`

HasBEzsigntemplatesignatureReason returns a boolean if a field has been set.

### GetEEzsigntemplatesignaturePositioning

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetEEzsigntemplatesignaturePositioning() FieldEEzsigntemplatesignaturePositioning`

GetEEzsigntemplatesignaturePositioning returns the EEzsigntemplatesignaturePositioning field if non-nil, zero value otherwise.

### GetEEzsigntemplatesignaturePositioningOk

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetEEzsigntemplatesignaturePositioningOk() (*FieldEEzsigntemplatesignaturePositioning, bool)`

GetEEzsigntemplatesignaturePositioningOk returns a tuple with the EEzsigntemplatesignaturePositioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplatesignaturePositioning

`func (o *EzsigntemplatesignatureRequestCompoundV2) SetEEzsigntemplatesignaturePositioning(v FieldEEzsigntemplatesignaturePositioning)`

SetEEzsigntemplatesignaturePositioning sets EEzsigntemplatesignaturePositioning field to given value.

### HasEEzsigntemplatesignaturePositioning

`func (o *EzsigntemplatesignatureRequestCompoundV2) HasEEzsigntemplatesignaturePositioning() bool`

HasEEzsigntemplatesignaturePositioning returns a boolean if a field has been set.

### GetIEzsigntemplatedocumentpagePagenumber

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetIEzsigntemplatedocumentpagePagenumber() int32`

GetIEzsigntemplatedocumentpagePagenumber returns the IEzsigntemplatedocumentpagePagenumber field if non-nil, zero value otherwise.

### GetIEzsigntemplatedocumentpagePagenumberOk

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetIEzsigntemplatedocumentpagePagenumberOk() (*int32, bool)`

GetIEzsigntemplatedocumentpagePagenumberOk returns a tuple with the IEzsigntemplatedocumentpagePagenumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplatedocumentpagePagenumber

`func (o *EzsigntemplatesignatureRequestCompoundV2) SetIEzsigntemplatedocumentpagePagenumber(v int32)`

SetIEzsigntemplatedocumentpagePagenumber sets IEzsigntemplatedocumentpagePagenumber field to given value.


### GetIEzsigntemplatesignatureX

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetIEzsigntemplatesignatureX() int32`

GetIEzsigntemplatesignatureX returns the IEzsigntemplatesignatureX field if non-nil, zero value otherwise.

### GetIEzsigntemplatesignatureXOk

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetIEzsigntemplatesignatureXOk() (*int32, bool)`

GetIEzsigntemplatesignatureXOk returns a tuple with the IEzsigntemplatesignatureX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplatesignatureX

`func (o *EzsigntemplatesignatureRequestCompoundV2) SetIEzsigntemplatesignatureX(v int32)`

SetIEzsigntemplatesignatureX sets IEzsigntemplatesignatureX field to given value.

### HasIEzsigntemplatesignatureX

`func (o *EzsigntemplatesignatureRequestCompoundV2) HasIEzsigntemplatesignatureX() bool`

HasIEzsigntemplatesignatureX returns a boolean if a field has been set.

### GetIEzsigntemplatesignatureY

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetIEzsigntemplatesignatureY() int32`

GetIEzsigntemplatesignatureY returns the IEzsigntemplatesignatureY field if non-nil, zero value otherwise.

### GetIEzsigntemplatesignatureYOk

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetIEzsigntemplatesignatureYOk() (*int32, bool)`

GetIEzsigntemplatesignatureYOk returns a tuple with the IEzsigntemplatesignatureY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplatesignatureY

`func (o *EzsigntemplatesignatureRequestCompoundV2) SetIEzsigntemplatesignatureY(v int32)`

SetIEzsigntemplatesignatureY sets IEzsigntemplatesignatureY field to given value.

### HasIEzsigntemplatesignatureY

`func (o *EzsigntemplatesignatureRequestCompoundV2) HasIEzsigntemplatesignatureY() bool`

HasIEzsigntemplatesignatureY returns a boolean if a field has been set.

### GetIEzsigntemplatesignatureWidth

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetIEzsigntemplatesignatureWidth() int32`

GetIEzsigntemplatesignatureWidth returns the IEzsigntemplatesignatureWidth field if non-nil, zero value otherwise.

### GetIEzsigntemplatesignatureWidthOk

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetIEzsigntemplatesignatureWidthOk() (*int32, bool)`

GetIEzsigntemplatesignatureWidthOk returns a tuple with the IEzsigntemplatesignatureWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplatesignatureWidth

`func (o *EzsigntemplatesignatureRequestCompoundV2) SetIEzsigntemplatesignatureWidth(v int32)`

SetIEzsigntemplatesignatureWidth sets IEzsigntemplatesignatureWidth field to given value.

### HasIEzsigntemplatesignatureWidth

`func (o *EzsigntemplatesignatureRequestCompoundV2) HasIEzsigntemplatesignatureWidth() bool`

HasIEzsigntemplatesignatureWidth returns a boolean if a field has been set.

### GetIEzsigntemplatesignatureHeight

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetIEzsigntemplatesignatureHeight() int32`

GetIEzsigntemplatesignatureHeight returns the IEzsigntemplatesignatureHeight field if non-nil, zero value otherwise.

### GetIEzsigntemplatesignatureHeightOk

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetIEzsigntemplatesignatureHeightOk() (*int32, bool)`

GetIEzsigntemplatesignatureHeightOk returns a tuple with the IEzsigntemplatesignatureHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplatesignatureHeight

`func (o *EzsigntemplatesignatureRequestCompoundV2) SetIEzsigntemplatesignatureHeight(v int32)`

SetIEzsigntemplatesignatureHeight sets IEzsigntemplatesignatureHeight field to given value.

### HasIEzsigntemplatesignatureHeight

`func (o *EzsigntemplatesignatureRequestCompoundV2) HasIEzsigntemplatesignatureHeight() bool`

HasIEzsigntemplatesignatureHeight returns a boolean if a field has been set.

### GetIEzsigntemplatesignatureStep

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetIEzsigntemplatesignatureStep() int32`

GetIEzsigntemplatesignatureStep returns the IEzsigntemplatesignatureStep field if non-nil, zero value otherwise.

### GetIEzsigntemplatesignatureStepOk

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetIEzsigntemplatesignatureStepOk() (*int32, bool)`

GetIEzsigntemplatesignatureStepOk returns a tuple with the IEzsigntemplatesignatureStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplatesignatureStep

`func (o *EzsigntemplatesignatureRequestCompoundV2) SetIEzsigntemplatesignatureStep(v int32)`

SetIEzsigntemplatesignatureStep sets IEzsigntemplatesignatureStep field to given value.


### GetEEzsigntemplatesignatureType

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetEEzsigntemplatesignatureType() FieldEEzsigntemplatesignatureType`

GetEEzsigntemplatesignatureType returns the EEzsigntemplatesignatureType field if non-nil, zero value otherwise.

### GetEEzsigntemplatesignatureTypeOk

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetEEzsigntemplatesignatureTypeOk() (*FieldEEzsigntemplatesignatureType, bool)`

GetEEzsigntemplatesignatureTypeOk returns a tuple with the EEzsigntemplatesignatureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplatesignatureType

`func (o *EzsigntemplatesignatureRequestCompoundV2) SetEEzsigntemplatesignatureType(v FieldEEzsigntemplatesignatureType)`

SetEEzsigntemplatesignatureType sets EEzsigntemplatesignatureType field to given value.


### GetEEzsigntemplatesignatureConsultationtrigger

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetEEzsigntemplatesignatureConsultationtrigger() FieldEEzsigntemplatesignatureConsultationtrigger`

GetEEzsigntemplatesignatureConsultationtrigger returns the EEzsigntemplatesignatureConsultationtrigger field if non-nil, zero value otherwise.

### GetEEzsigntemplatesignatureConsultationtriggerOk

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetEEzsigntemplatesignatureConsultationtriggerOk() (*FieldEEzsigntemplatesignatureConsultationtrigger, bool)`

GetEEzsigntemplatesignatureConsultationtriggerOk returns a tuple with the EEzsigntemplatesignatureConsultationtrigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplatesignatureConsultationtrigger

`func (o *EzsigntemplatesignatureRequestCompoundV2) SetEEzsigntemplatesignatureConsultationtrigger(v FieldEEzsigntemplatesignatureConsultationtrigger)`

SetEEzsigntemplatesignatureConsultationtrigger sets EEzsigntemplatesignatureConsultationtrigger field to given value.

### HasEEzsigntemplatesignatureConsultationtrigger

`func (o *EzsigntemplatesignatureRequestCompoundV2) HasEEzsigntemplatesignatureConsultationtrigger() bool`

HasEEzsigntemplatesignatureConsultationtrigger returns a boolean if a field has been set.

### GetTEzsigntemplatesignatureTooltip

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetTEzsigntemplatesignatureTooltip() string`

GetTEzsigntemplatesignatureTooltip returns the TEzsigntemplatesignatureTooltip field if non-nil, zero value otherwise.

### GetTEzsigntemplatesignatureTooltipOk

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetTEzsigntemplatesignatureTooltipOk() (*string, bool)`

GetTEzsigntemplatesignatureTooltipOk returns a tuple with the TEzsigntemplatesignatureTooltip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEzsigntemplatesignatureTooltip

`func (o *EzsigntemplatesignatureRequestCompoundV2) SetTEzsigntemplatesignatureTooltip(v string)`

SetTEzsigntemplatesignatureTooltip sets TEzsigntemplatesignatureTooltip field to given value.

### HasTEzsigntemplatesignatureTooltip

`func (o *EzsigntemplatesignatureRequestCompoundV2) HasTEzsigntemplatesignatureTooltip() bool`

HasTEzsigntemplatesignatureTooltip returns a boolean if a field has been set.

### GetEEzsigntemplatesignatureTooltipposition

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetEEzsigntemplatesignatureTooltipposition() FieldEEzsigntemplatesignatureTooltipposition`

GetEEzsigntemplatesignatureTooltipposition returns the EEzsigntemplatesignatureTooltipposition field if non-nil, zero value otherwise.

### GetEEzsigntemplatesignatureTooltippositionOk

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetEEzsigntemplatesignatureTooltippositionOk() (*FieldEEzsigntemplatesignatureTooltipposition, bool)`

GetEEzsigntemplatesignatureTooltippositionOk returns a tuple with the EEzsigntemplatesignatureTooltipposition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplatesignatureTooltipposition

`func (o *EzsigntemplatesignatureRequestCompoundV2) SetEEzsigntemplatesignatureTooltipposition(v FieldEEzsigntemplatesignatureTooltipposition)`

SetEEzsigntemplatesignatureTooltipposition sets EEzsigntemplatesignatureTooltipposition field to given value.

### HasEEzsigntemplatesignatureTooltipposition

`func (o *EzsigntemplatesignatureRequestCompoundV2) HasEEzsigntemplatesignatureTooltipposition() bool`

HasEEzsigntemplatesignatureTooltipposition returns a boolean if a field has been set.

### GetEEzsigntemplatesignatureFont

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetEEzsigntemplatesignatureFont() FieldEEzsigntemplatesignatureFont`

GetEEzsigntemplatesignatureFont returns the EEzsigntemplatesignatureFont field if non-nil, zero value otherwise.

### GetEEzsigntemplatesignatureFontOk

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetEEzsigntemplatesignatureFontOk() (*FieldEEzsigntemplatesignatureFont, bool)`

GetEEzsigntemplatesignatureFontOk returns a tuple with the EEzsigntemplatesignatureFont field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplatesignatureFont

`func (o *EzsigntemplatesignatureRequestCompoundV2) SetEEzsigntemplatesignatureFont(v FieldEEzsigntemplatesignatureFont)`

SetEEzsigntemplatesignatureFont sets EEzsigntemplatesignatureFont field to given value.

### HasEEzsigntemplatesignatureFont

`func (o *EzsigntemplatesignatureRequestCompoundV2) HasEEzsigntemplatesignatureFont() bool`

HasEEzsigntemplatesignatureFont returns a boolean if a field has been set.

### GetBEzsigntemplatesignatureRequired

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetBEzsigntemplatesignatureRequired() bool`

GetBEzsigntemplatesignatureRequired returns the BEzsigntemplatesignatureRequired field if non-nil, zero value otherwise.

### GetBEzsigntemplatesignatureRequiredOk

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetBEzsigntemplatesignatureRequiredOk() (*bool, bool)`

GetBEzsigntemplatesignatureRequiredOk returns a tuple with the BEzsigntemplatesignatureRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigntemplatesignatureRequired

`func (o *EzsigntemplatesignatureRequestCompoundV2) SetBEzsigntemplatesignatureRequired(v bool)`

SetBEzsigntemplatesignatureRequired sets BEzsigntemplatesignatureRequired field to given value.

### HasBEzsigntemplatesignatureRequired

`func (o *EzsigntemplatesignatureRequestCompoundV2) HasBEzsigntemplatesignatureRequired() bool`

HasBEzsigntemplatesignatureRequired returns a boolean if a field has been set.

### GetEEzsigntemplatesignatureAttachmentnamesource

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetEEzsigntemplatesignatureAttachmentnamesource() FieldEEzsigntemplatesignatureAttachmentnamesource`

GetEEzsigntemplatesignatureAttachmentnamesource returns the EEzsigntemplatesignatureAttachmentnamesource field if non-nil, zero value otherwise.

### GetEEzsigntemplatesignatureAttachmentnamesourceOk

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetEEzsigntemplatesignatureAttachmentnamesourceOk() (*FieldEEzsigntemplatesignatureAttachmentnamesource, bool)`

GetEEzsigntemplatesignatureAttachmentnamesourceOk returns a tuple with the EEzsigntemplatesignatureAttachmentnamesource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplatesignatureAttachmentnamesource

`func (o *EzsigntemplatesignatureRequestCompoundV2) SetEEzsigntemplatesignatureAttachmentnamesource(v FieldEEzsigntemplatesignatureAttachmentnamesource)`

SetEEzsigntemplatesignatureAttachmentnamesource sets EEzsigntemplatesignatureAttachmentnamesource field to given value.

### HasEEzsigntemplatesignatureAttachmentnamesource

`func (o *EzsigntemplatesignatureRequestCompoundV2) HasEEzsigntemplatesignatureAttachmentnamesource() bool`

HasEEzsigntemplatesignatureAttachmentnamesource returns a boolean if a field has been set.

### GetSEzsigntemplatesignatureAttachmentdescription

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetSEzsigntemplatesignatureAttachmentdescription() string`

GetSEzsigntemplatesignatureAttachmentdescription returns the SEzsigntemplatesignatureAttachmentdescription field if non-nil, zero value otherwise.

### GetSEzsigntemplatesignatureAttachmentdescriptionOk

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetSEzsigntemplatesignatureAttachmentdescriptionOk() (*string, bool)`

GetSEzsigntemplatesignatureAttachmentdescriptionOk returns a tuple with the SEzsigntemplatesignatureAttachmentdescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplatesignatureAttachmentdescription

`func (o *EzsigntemplatesignatureRequestCompoundV2) SetSEzsigntemplatesignatureAttachmentdescription(v string)`

SetSEzsigntemplatesignatureAttachmentdescription sets SEzsigntemplatesignatureAttachmentdescription field to given value.

### HasSEzsigntemplatesignatureAttachmentdescription

`func (o *EzsigntemplatesignatureRequestCompoundV2) HasSEzsigntemplatesignatureAttachmentdescription() bool`

HasSEzsigntemplatesignatureAttachmentdescription returns a boolean if a field has been set.

### GetIEzsigntemplatesignatureValidationstep

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetIEzsigntemplatesignatureValidationstep() int32`

GetIEzsigntemplatesignatureValidationstep returns the IEzsigntemplatesignatureValidationstep field if non-nil, zero value otherwise.

### GetIEzsigntemplatesignatureValidationstepOk

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetIEzsigntemplatesignatureValidationstepOk() (*int32, bool)`

GetIEzsigntemplatesignatureValidationstepOk returns a tuple with the IEzsigntemplatesignatureValidationstep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplatesignatureValidationstep

`func (o *EzsigntemplatesignatureRequestCompoundV2) SetIEzsigntemplatesignatureValidationstep(v int32)`

SetIEzsigntemplatesignatureValidationstep sets IEzsigntemplatesignatureValidationstep field to given value.

### HasIEzsigntemplatesignatureValidationstep

`func (o *EzsigntemplatesignatureRequestCompoundV2) HasIEzsigntemplatesignatureValidationstep() bool`

HasIEzsigntemplatesignatureValidationstep returns a boolean if a field has been set.

### GetIEzsigntemplatesignatureMaxlength

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetIEzsigntemplatesignatureMaxlength() int32`

GetIEzsigntemplatesignatureMaxlength returns the IEzsigntemplatesignatureMaxlength field if non-nil, zero value otherwise.

### GetIEzsigntemplatesignatureMaxlengthOk

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetIEzsigntemplatesignatureMaxlengthOk() (*int32, bool)`

GetIEzsigntemplatesignatureMaxlengthOk returns a tuple with the IEzsigntemplatesignatureMaxlength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplatesignatureMaxlength

`func (o *EzsigntemplatesignatureRequestCompoundV2) SetIEzsigntemplatesignatureMaxlength(v int32)`

SetIEzsigntemplatesignatureMaxlength sets IEzsigntemplatesignatureMaxlength field to given value.

### HasIEzsigntemplatesignatureMaxlength

`func (o *EzsigntemplatesignatureRequestCompoundV2) HasIEzsigntemplatesignatureMaxlength() bool`

HasIEzsigntemplatesignatureMaxlength returns a boolean if a field has been set.

### GetSEzsigntemplatesignatureDefaultvalue

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetSEzsigntemplatesignatureDefaultvalue() string`

GetSEzsigntemplatesignatureDefaultvalue returns the SEzsigntemplatesignatureDefaultvalue field if non-nil, zero value otherwise.

### GetSEzsigntemplatesignatureDefaultvalueOk

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetSEzsigntemplatesignatureDefaultvalueOk() (*string, bool)`

GetSEzsigntemplatesignatureDefaultvalueOk returns a tuple with the SEzsigntemplatesignatureDefaultvalue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplatesignatureDefaultvalue

`func (o *EzsigntemplatesignatureRequestCompoundV2) SetSEzsigntemplatesignatureDefaultvalue(v string)`

SetSEzsigntemplatesignatureDefaultvalue sets SEzsigntemplatesignatureDefaultvalue field to given value.

### HasSEzsigntemplatesignatureDefaultvalue

`func (o *EzsigntemplatesignatureRequestCompoundV2) HasSEzsigntemplatesignatureDefaultvalue() bool`

HasSEzsigntemplatesignatureDefaultvalue returns a boolean if a field has been set.

### GetSEzsigntemplatesignatureRegexp

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetSEzsigntemplatesignatureRegexp() string`

GetSEzsigntemplatesignatureRegexp returns the SEzsigntemplatesignatureRegexp field if non-nil, zero value otherwise.

### GetSEzsigntemplatesignatureRegexpOk

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetSEzsigntemplatesignatureRegexpOk() (*string, bool)`

GetSEzsigntemplatesignatureRegexpOk returns a tuple with the SEzsigntemplatesignatureRegexp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplatesignatureRegexp

`func (o *EzsigntemplatesignatureRequestCompoundV2) SetSEzsigntemplatesignatureRegexp(v string)`

SetSEzsigntemplatesignatureRegexp sets SEzsigntemplatesignatureRegexp field to given value.

### HasSEzsigntemplatesignatureRegexp

`func (o *EzsigntemplatesignatureRequestCompoundV2) HasSEzsigntemplatesignatureRegexp() bool`

HasSEzsigntemplatesignatureRegexp returns a boolean if a field has been set.

### GetEEzsigntemplatesignatureTextvalidation

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetEEzsigntemplatesignatureTextvalidation() EnumTextvalidation`

GetEEzsigntemplatesignatureTextvalidation returns the EEzsigntemplatesignatureTextvalidation field if non-nil, zero value otherwise.

### GetEEzsigntemplatesignatureTextvalidationOk

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetEEzsigntemplatesignatureTextvalidationOk() (*EnumTextvalidation, bool)`

GetEEzsigntemplatesignatureTextvalidationOk returns a tuple with the EEzsigntemplatesignatureTextvalidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplatesignatureTextvalidation

`func (o *EzsigntemplatesignatureRequestCompoundV2) SetEEzsigntemplatesignatureTextvalidation(v EnumTextvalidation)`

SetEEzsigntemplatesignatureTextvalidation sets EEzsigntemplatesignatureTextvalidation field to given value.

### HasEEzsigntemplatesignatureTextvalidation

`func (o *EzsigntemplatesignatureRequestCompoundV2) HasEEzsigntemplatesignatureTextvalidation() bool`

HasEEzsigntemplatesignatureTextvalidation returns a boolean if a field has been set.

### GetSEzsigntemplatesignatureTextvalidationcustommessage

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetSEzsigntemplatesignatureTextvalidationcustommessage() string`

GetSEzsigntemplatesignatureTextvalidationcustommessage returns the SEzsigntemplatesignatureTextvalidationcustommessage field if non-nil, zero value otherwise.

### GetSEzsigntemplatesignatureTextvalidationcustommessageOk

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetSEzsigntemplatesignatureTextvalidationcustommessageOk() (*string, bool)`

GetSEzsigntemplatesignatureTextvalidationcustommessageOk returns a tuple with the SEzsigntemplatesignatureTextvalidationcustommessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplatesignatureTextvalidationcustommessage

`func (o *EzsigntemplatesignatureRequestCompoundV2) SetSEzsigntemplatesignatureTextvalidationcustommessage(v string)`

SetSEzsigntemplatesignatureTextvalidationcustommessage sets SEzsigntemplatesignatureTextvalidationcustommessage field to given value.

### HasSEzsigntemplatesignatureTextvalidationcustommessage

`func (o *EzsigntemplatesignatureRequestCompoundV2) HasSEzsigntemplatesignatureTextvalidationcustommessage() bool`

HasSEzsigntemplatesignatureTextvalidationcustommessage returns a boolean if a field has been set.

### GetEEzsigntemplatesignatureDependencyrequirement

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetEEzsigntemplatesignatureDependencyrequirement() FieldEEzsigntemplatesignatureDependencyrequirement`

GetEEzsigntemplatesignatureDependencyrequirement returns the EEzsigntemplatesignatureDependencyrequirement field if non-nil, zero value otherwise.

### GetEEzsigntemplatesignatureDependencyrequirementOk

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetEEzsigntemplatesignatureDependencyrequirementOk() (*FieldEEzsigntemplatesignatureDependencyrequirement, bool)`

GetEEzsigntemplatesignatureDependencyrequirementOk returns a tuple with the EEzsigntemplatesignatureDependencyrequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplatesignatureDependencyrequirement

`func (o *EzsigntemplatesignatureRequestCompoundV2) SetEEzsigntemplatesignatureDependencyrequirement(v FieldEEzsigntemplatesignatureDependencyrequirement)`

SetEEzsigntemplatesignatureDependencyrequirement sets EEzsigntemplatesignatureDependencyrequirement field to given value.

### HasEEzsigntemplatesignatureDependencyrequirement

`func (o *EzsigntemplatesignatureRequestCompoundV2) HasEEzsigntemplatesignatureDependencyrequirement() bool`

HasEEzsigntemplatesignatureDependencyrequirement returns a boolean if a field has been set.

### GetSEzsigntemplatesignaturePositioningpattern

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetSEzsigntemplatesignaturePositioningpattern() string`

GetSEzsigntemplatesignaturePositioningpattern returns the SEzsigntemplatesignaturePositioningpattern field if non-nil, zero value otherwise.

### GetSEzsigntemplatesignaturePositioningpatternOk

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetSEzsigntemplatesignaturePositioningpatternOk() (*string, bool)`

GetSEzsigntemplatesignaturePositioningpatternOk returns a tuple with the SEzsigntemplatesignaturePositioningpattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplatesignaturePositioningpattern

`func (o *EzsigntemplatesignatureRequestCompoundV2) SetSEzsigntemplatesignaturePositioningpattern(v string)`

SetSEzsigntemplatesignaturePositioningpattern sets SEzsigntemplatesignaturePositioningpattern field to given value.

### HasSEzsigntemplatesignaturePositioningpattern

`func (o *EzsigntemplatesignatureRequestCompoundV2) HasSEzsigntemplatesignaturePositioningpattern() bool`

HasSEzsigntemplatesignaturePositioningpattern returns a boolean if a field has been set.

### GetIEzsigntemplatesignaturePositioningoffsetx

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetIEzsigntemplatesignaturePositioningoffsetx() int32`

GetIEzsigntemplatesignaturePositioningoffsetx returns the IEzsigntemplatesignaturePositioningoffsetx field if non-nil, zero value otherwise.

### GetIEzsigntemplatesignaturePositioningoffsetxOk

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetIEzsigntemplatesignaturePositioningoffsetxOk() (*int32, bool)`

GetIEzsigntemplatesignaturePositioningoffsetxOk returns a tuple with the IEzsigntemplatesignaturePositioningoffsetx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplatesignaturePositioningoffsetx

`func (o *EzsigntemplatesignatureRequestCompoundV2) SetIEzsigntemplatesignaturePositioningoffsetx(v int32)`

SetIEzsigntemplatesignaturePositioningoffsetx sets IEzsigntemplatesignaturePositioningoffsetx field to given value.

### HasIEzsigntemplatesignaturePositioningoffsetx

`func (o *EzsigntemplatesignatureRequestCompoundV2) HasIEzsigntemplatesignaturePositioningoffsetx() bool`

HasIEzsigntemplatesignaturePositioningoffsetx returns a boolean if a field has been set.

### GetIEzsigntemplatesignaturePositioningoffsety

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetIEzsigntemplatesignaturePositioningoffsety() int32`

GetIEzsigntemplatesignaturePositioningoffsety returns the IEzsigntemplatesignaturePositioningoffsety field if non-nil, zero value otherwise.

### GetIEzsigntemplatesignaturePositioningoffsetyOk

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetIEzsigntemplatesignaturePositioningoffsetyOk() (*int32, bool)`

GetIEzsigntemplatesignaturePositioningoffsetyOk returns a tuple with the IEzsigntemplatesignaturePositioningoffsety field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplatesignaturePositioningoffsety

`func (o *EzsigntemplatesignatureRequestCompoundV2) SetIEzsigntemplatesignaturePositioningoffsety(v int32)`

SetIEzsigntemplatesignaturePositioningoffsety sets IEzsigntemplatesignaturePositioningoffsety field to given value.

### HasIEzsigntemplatesignaturePositioningoffsety

`func (o *EzsigntemplatesignatureRequestCompoundV2) HasIEzsigntemplatesignaturePositioningoffsety() bool`

HasIEzsigntemplatesignaturePositioningoffsety returns a boolean if a field has been set.

### GetEEzsigntemplatesignaturePositioningoccurence

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetEEzsigntemplatesignaturePositioningoccurence() FieldEEzsigntemplatesignaturePositioningoccurence`

GetEEzsigntemplatesignaturePositioningoccurence returns the EEzsigntemplatesignaturePositioningoccurence field if non-nil, zero value otherwise.

### GetEEzsigntemplatesignaturePositioningoccurenceOk

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetEEzsigntemplatesignaturePositioningoccurenceOk() (*FieldEEzsigntemplatesignaturePositioningoccurence, bool)`

GetEEzsigntemplatesignaturePositioningoccurenceOk returns a tuple with the EEzsigntemplatesignaturePositioningoccurence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplatesignaturePositioningoccurence

`func (o *EzsigntemplatesignatureRequestCompoundV2) SetEEzsigntemplatesignaturePositioningoccurence(v FieldEEzsigntemplatesignaturePositioningoccurence)`

SetEEzsigntemplatesignaturePositioningoccurence sets EEzsigntemplatesignaturePositioningoccurence field to given value.

### HasEEzsigntemplatesignaturePositioningoccurence

`func (o *EzsigntemplatesignatureRequestCompoundV2) HasEEzsigntemplatesignaturePositioningoccurence() bool`

HasEEzsigntemplatesignaturePositioningoccurence returns a boolean if a field has been set.

### GetBEzsigntemplatesignatureCustomdate

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetBEzsigntemplatesignatureCustomdate() bool`

GetBEzsigntemplatesignatureCustomdate returns the BEzsigntemplatesignatureCustomdate field if non-nil, zero value otherwise.

### GetBEzsigntemplatesignatureCustomdateOk

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetBEzsigntemplatesignatureCustomdateOk() (*bool, bool)`

GetBEzsigntemplatesignatureCustomdateOk returns a tuple with the BEzsigntemplatesignatureCustomdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigntemplatesignatureCustomdate

`func (o *EzsigntemplatesignatureRequestCompoundV2) SetBEzsigntemplatesignatureCustomdate(v bool)`

SetBEzsigntemplatesignatureCustomdate sets BEzsigntemplatesignatureCustomdate field to given value.

### HasBEzsigntemplatesignatureCustomdate

`func (o *EzsigntemplatesignatureRequestCompoundV2) HasBEzsigntemplatesignatureCustomdate() bool`

HasBEzsigntemplatesignatureCustomdate returns a boolean if a field has been set.

### GetAObjEzsigntemplatesignaturecustomdate

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetAObjEzsigntemplatesignaturecustomdate() []EzsigntemplatesignaturecustomdateRequestCompoundV2`

GetAObjEzsigntemplatesignaturecustomdate returns the AObjEzsigntemplatesignaturecustomdate field if non-nil, zero value otherwise.

### GetAObjEzsigntemplatesignaturecustomdateOk

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetAObjEzsigntemplatesignaturecustomdateOk() (*[]EzsigntemplatesignaturecustomdateRequestCompoundV2, bool)`

GetAObjEzsigntemplatesignaturecustomdateOk returns a tuple with the AObjEzsigntemplatesignaturecustomdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzsigntemplatesignaturecustomdate

`func (o *EzsigntemplatesignatureRequestCompoundV2) SetAObjEzsigntemplatesignaturecustomdate(v []EzsigntemplatesignaturecustomdateRequestCompoundV2)`

SetAObjEzsigntemplatesignaturecustomdate sets AObjEzsigntemplatesignaturecustomdate field to given value.

### HasAObjEzsigntemplatesignaturecustomdate

`func (o *EzsigntemplatesignatureRequestCompoundV2) HasAObjEzsigntemplatesignaturecustomdate() bool`

HasAObjEzsigntemplatesignaturecustomdate returns a boolean if a field has been set.

### GetAObjEzsigntemplateelementdependency

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetAObjEzsigntemplateelementdependency() []EzsigntemplateelementdependencyRequestCompound`

GetAObjEzsigntemplateelementdependency returns the AObjEzsigntemplateelementdependency field if non-nil, zero value otherwise.

### GetAObjEzsigntemplateelementdependencyOk

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetAObjEzsigntemplateelementdependencyOk() (*[]EzsigntemplateelementdependencyRequestCompound, bool)`

GetAObjEzsigntemplateelementdependencyOk returns a tuple with the AObjEzsigntemplateelementdependency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzsigntemplateelementdependency

`func (o *EzsigntemplatesignatureRequestCompoundV2) SetAObjEzsigntemplateelementdependency(v []EzsigntemplateelementdependencyRequestCompound)`

SetAObjEzsigntemplateelementdependency sets AObjEzsigntemplateelementdependency field to given value.

### HasAObjEzsigntemplateelementdependency

`func (o *EzsigntemplatesignatureRequestCompoundV2) HasAObjEzsigntemplateelementdependency() bool`

HasAObjEzsigntemplateelementdependency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


