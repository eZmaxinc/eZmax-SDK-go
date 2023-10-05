# EzsignsignatureResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsignsignatureID** | **int32** | The unique ID of the Ezsignsignature | 
**FkiEzsigndocumentID** | **int32** | The unique ID of the Ezsigndocument | 
**FkiEzsignfoldersignerassociationID** | **int32** | The unique ID of the Ezsignfoldersignerassociation | 
**IEzsignpagePagenumber** | **int32** | The page number in the Ezsigndocument | 
**IEzsignsignatureX** | **int32** | The X coordinate (Horizontal) where to put the Ezsignsignature on the page.  Coordinate is calculated at 100dpi (dot per inch). So for example, if you want to put the Ezsignsignature 2 inches from the left border of the page, you would use \&quot;200\&quot; for the X coordinate. | 
**IEzsignsignatureY** | **int32** | The Y coordinate (Vertical) where to put the Ezsignsignature on the page.  Coordinate is calculated at 100dpi (dot per inch). So for example, if you want to put the Ezsignsignature 3 inches from the top border of the page, you would use \&quot;300\&quot; for the Y coordinate. | 
**IEzsignsignatureHeight** | Pointer to **int32** | The height of the Ezsignsignature.  Size is calculated at 100dpi (dot per inch). So for example, if you want the Ezsignsignature to have an height of 2 inches, you would use \&quot;200\&quot; for the iEzsignsignatureHeight. | [optional] 
**IEzsignsignatureWidth** | Pointer to **int32** | The width of the Ezsignsignature.  Size is calculated at 100dpi (dot per inch). So for example, if you want the Ezsignsignature to have a width of 2 inches, you would use \&quot;200\&quot; for the iEzsignsignatureWidth. | [optional] 
**IEzsignsignatureStep** | **int32** | The step when the Ezsignsigner will be invited to sign | 
**EEzsignsignatureType** | [**FieldEEzsignsignatureType**](FieldEEzsignsignatureType.md) |  | 
**TEzsignsignatureTooltip** | Pointer to **string** | A tooltip that will be presented to Ezsignsigner about the Ezsignsignature | [optional] 
**EEzsignsignatureTooltipposition** | Pointer to [**FieldEEzsignsignatureTooltipposition**](FieldEEzsignsignatureTooltipposition.md) |  | [optional] 
**EEzsignsignatureFont** | Pointer to [**FieldEEzsignsignatureFont**](FieldEEzsignsignatureFont.md) |  | [optional] 
**IEzsignsignatureValidationstep** | Pointer to **int32** | The step when the Ezsignsigner will be invited to validate the Ezsignsignature of eEzsignsignatureType Attachments | [optional] 
**SEzsignsignatureAttachmentdescription** | Pointer to **string** | The description attached to the attachment name added in Ezsignsignature of eEzsignsignatureType Attachments | [optional] 
**EEzsignsignatureAttachmentnamesource** | Pointer to [**FieldEEzsignsignatureAttachmentnamesource**](FieldEEzsignsignatureAttachmentnamesource.md) |  | [optional] 
**BEzsignsignatureRequired** | Pointer to **bool** | Whether the Ezsignsignature is required or not. This field is relevant only with Ezsignsignature with eEzsignsignatureType &#x3D; Attachments. | [optional] 
**FkiEzsignfoldersignerassociationIDValidation** | Pointer to **int32** | The unique ID of the Ezsignfoldersignerassociation | [optional] 
**DtEzsignsignatureDate** | Pointer to **string** | The date the Ezsignsignature was signed | [optional] 
**IEzsignsignatureattachmentCount** | Pointer to **int32** | The count of Ezsignsignatureattachment | [optional] 
**SEzsignsignatureDescription** | Pointer to **string** | The value entered while signing Ezsignsignature of eEzsignsignatureType **City**, **FieldText** and **FieldTextarea** | [optional] 
**IEzsignsignatureMaxlength** | Pointer to **int32** | The maximum length for the value in the Ezsignsignature  This can only be set if eEzsignsignatureType is **FieldText** or **FieldTextarea** | [optional] 
**EEzsignsignatureTextvalidation** | Pointer to [**EnumTextvalidation**](EnumTextvalidation.md) |  | [optional] 
**EEzsignsignatureDependencyrequirement** | Pointer to [**FieldEEzsignsignatureDependencyrequirement**](FieldEEzsignsignatureDependencyrequirement.md) |  | [optional] 
**SEzsignsignatureRegexp** | Pointer to **string** | A regular expression to indicate what values are acceptable for the Ezsignsignature.  This can only be set if eEzsignsignatureType is **FieldText** or **FieldTextarea** and eEzsignsignatureTextvalidation is **Custom** | [optional] 
**ObjContactName** | [**CustomContactNameResponse**](CustomContactNameResponse.md) |  | 
**ObjContactNameDelegation** | Pointer to [**CustomContactNameResponse**](CustomContactNameResponse.md) |  | [optional] 
**ObjSignature** | Pointer to [**SignatureResponseCompound**](SignatureResponseCompound.md) |  | [optional] 
**BEzsignsignatureCustomdate** | Pointer to **bool** | Whether the Ezsignsignature has a custom date format or not. (Only possible when eEzsignsignatureType is **Name** or **Handwritten**) | [optional] 
**AObjEzsignsignaturecustomdate** | Pointer to [**[]EzsignsignaturecustomdateResponseCompound**](EzsignsignaturecustomdateResponseCompound.md) | An array of custom date blocks that will be filled at the time of signature.  Can only be used if bEzsignsignatureCustomdate is true.  Use an empty array if you don&#39;t want to have a date at all. | [optional] 
**ObjCreditcardtransaction** | Pointer to [**CustomCreditcardtransactionResponse**](CustomCreditcardtransactionResponse.md) |  | [optional] 
**AObjEzsignelementdependency** | Pointer to [**[]EzsignelementdependencyResponseCompound**](EzsignelementdependencyResponseCompound.md) |  | [optional] 

## Methods

### NewEzsignsignatureResponseCompound

`func NewEzsignsignatureResponseCompound(pkiEzsignsignatureID int32, fkiEzsigndocumentID int32, fkiEzsignfoldersignerassociationID int32, iEzsignpagePagenumber int32, iEzsignsignatureX int32, iEzsignsignatureY int32, iEzsignsignatureStep int32, eEzsignsignatureType FieldEEzsignsignatureType, objContactName CustomContactNameResponse, ) *EzsignsignatureResponseCompound`

NewEzsignsignatureResponseCompound instantiates a new EzsignsignatureResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignsignatureResponseCompoundWithDefaults

`func NewEzsignsignatureResponseCompoundWithDefaults() *EzsignsignatureResponseCompound`

NewEzsignsignatureResponseCompoundWithDefaults instantiates a new EzsignsignatureResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsignsignatureID

`func (o *EzsignsignatureResponseCompound) GetPkiEzsignsignatureID() int32`

GetPkiEzsignsignatureID returns the PkiEzsignsignatureID field if non-nil, zero value otherwise.

### GetPkiEzsignsignatureIDOk

`func (o *EzsignsignatureResponseCompound) GetPkiEzsignsignatureIDOk() (*int32, bool)`

GetPkiEzsignsignatureIDOk returns a tuple with the PkiEzsignsignatureID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsignsignatureID

`func (o *EzsignsignatureResponseCompound) SetPkiEzsignsignatureID(v int32)`

SetPkiEzsignsignatureID sets PkiEzsignsignatureID field to given value.


### GetFkiEzsigndocumentID

`func (o *EzsignsignatureResponseCompound) GetFkiEzsigndocumentID() int32`

GetFkiEzsigndocumentID returns the FkiEzsigndocumentID field if non-nil, zero value otherwise.

### GetFkiEzsigndocumentIDOk

`func (o *EzsignsignatureResponseCompound) GetFkiEzsigndocumentIDOk() (*int32, bool)`

GetFkiEzsigndocumentIDOk returns a tuple with the FkiEzsigndocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigndocumentID

`func (o *EzsignsignatureResponseCompound) SetFkiEzsigndocumentID(v int32)`

SetFkiEzsigndocumentID sets FkiEzsigndocumentID field to given value.


### GetFkiEzsignfoldersignerassociationID

`func (o *EzsignsignatureResponseCompound) GetFkiEzsignfoldersignerassociationID() int32`

GetFkiEzsignfoldersignerassociationID returns the FkiEzsignfoldersignerassociationID field if non-nil, zero value otherwise.

### GetFkiEzsignfoldersignerassociationIDOk

`func (o *EzsignsignatureResponseCompound) GetFkiEzsignfoldersignerassociationIDOk() (*int32, bool)`

GetFkiEzsignfoldersignerassociationIDOk returns a tuple with the FkiEzsignfoldersignerassociationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfoldersignerassociationID

`func (o *EzsignsignatureResponseCompound) SetFkiEzsignfoldersignerassociationID(v int32)`

SetFkiEzsignfoldersignerassociationID sets FkiEzsignfoldersignerassociationID field to given value.


### GetIEzsignpagePagenumber

`func (o *EzsignsignatureResponseCompound) GetIEzsignpagePagenumber() int32`

GetIEzsignpagePagenumber returns the IEzsignpagePagenumber field if non-nil, zero value otherwise.

### GetIEzsignpagePagenumberOk

`func (o *EzsignsignatureResponseCompound) GetIEzsignpagePagenumberOk() (*int32, bool)`

GetIEzsignpagePagenumberOk returns a tuple with the IEzsignpagePagenumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignpagePagenumber

`func (o *EzsignsignatureResponseCompound) SetIEzsignpagePagenumber(v int32)`

SetIEzsignpagePagenumber sets IEzsignpagePagenumber field to given value.


### GetIEzsignsignatureX

`func (o *EzsignsignatureResponseCompound) GetIEzsignsignatureX() int32`

GetIEzsignsignatureX returns the IEzsignsignatureX field if non-nil, zero value otherwise.

### GetIEzsignsignatureXOk

`func (o *EzsignsignatureResponseCompound) GetIEzsignsignatureXOk() (*int32, bool)`

GetIEzsignsignatureXOk returns a tuple with the IEzsignsignatureX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignsignatureX

`func (o *EzsignsignatureResponseCompound) SetIEzsignsignatureX(v int32)`

SetIEzsignsignatureX sets IEzsignsignatureX field to given value.


### GetIEzsignsignatureY

`func (o *EzsignsignatureResponseCompound) GetIEzsignsignatureY() int32`

GetIEzsignsignatureY returns the IEzsignsignatureY field if non-nil, zero value otherwise.

### GetIEzsignsignatureYOk

`func (o *EzsignsignatureResponseCompound) GetIEzsignsignatureYOk() (*int32, bool)`

GetIEzsignsignatureYOk returns a tuple with the IEzsignsignatureY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignsignatureY

`func (o *EzsignsignatureResponseCompound) SetIEzsignsignatureY(v int32)`

SetIEzsignsignatureY sets IEzsignsignatureY field to given value.


### GetIEzsignsignatureHeight

`func (o *EzsignsignatureResponseCompound) GetIEzsignsignatureHeight() int32`

GetIEzsignsignatureHeight returns the IEzsignsignatureHeight field if non-nil, zero value otherwise.

### GetIEzsignsignatureHeightOk

`func (o *EzsignsignatureResponseCompound) GetIEzsignsignatureHeightOk() (*int32, bool)`

GetIEzsignsignatureHeightOk returns a tuple with the IEzsignsignatureHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignsignatureHeight

`func (o *EzsignsignatureResponseCompound) SetIEzsignsignatureHeight(v int32)`

SetIEzsignsignatureHeight sets IEzsignsignatureHeight field to given value.

### HasIEzsignsignatureHeight

`func (o *EzsignsignatureResponseCompound) HasIEzsignsignatureHeight() bool`

HasIEzsignsignatureHeight returns a boolean if a field has been set.

### GetIEzsignsignatureWidth

`func (o *EzsignsignatureResponseCompound) GetIEzsignsignatureWidth() int32`

GetIEzsignsignatureWidth returns the IEzsignsignatureWidth field if non-nil, zero value otherwise.

### GetIEzsignsignatureWidthOk

`func (o *EzsignsignatureResponseCompound) GetIEzsignsignatureWidthOk() (*int32, bool)`

GetIEzsignsignatureWidthOk returns a tuple with the IEzsignsignatureWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignsignatureWidth

`func (o *EzsignsignatureResponseCompound) SetIEzsignsignatureWidth(v int32)`

SetIEzsignsignatureWidth sets IEzsignsignatureWidth field to given value.

### HasIEzsignsignatureWidth

`func (o *EzsignsignatureResponseCompound) HasIEzsignsignatureWidth() bool`

HasIEzsignsignatureWidth returns a boolean if a field has been set.

### GetIEzsignsignatureStep

`func (o *EzsignsignatureResponseCompound) GetIEzsignsignatureStep() int32`

GetIEzsignsignatureStep returns the IEzsignsignatureStep field if non-nil, zero value otherwise.

### GetIEzsignsignatureStepOk

`func (o *EzsignsignatureResponseCompound) GetIEzsignsignatureStepOk() (*int32, bool)`

GetIEzsignsignatureStepOk returns a tuple with the IEzsignsignatureStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignsignatureStep

`func (o *EzsignsignatureResponseCompound) SetIEzsignsignatureStep(v int32)`

SetIEzsignsignatureStep sets IEzsignsignatureStep field to given value.


### GetEEzsignsignatureType

`func (o *EzsignsignatureResponseCompound) GetEEzsignsignatureType() FieldEEzsignsignatureType`

GetEEzsignsignatureType returns the EEzsignsignatureType field if non-nil, zero value otherwise.

### GetEEzsignsignatureTypeOk

`func (o *EzsignsignatureResponseCompound) GetEEzsignsignatureTypeOk() (*FieldEEzsignsignatureType, bool)`

GetEEzsignsignatureTypeOk returns a tuple with the EEzsignsignatureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignsignatureType

`func (o *EzsignsignatureResponseCompound) SetEEzsignsignatureType(v FieldEEzsignsignatureType)`

SetEEzsignsignatureType sets EEzsignsignatureType field to given value.


### GetTEzsignsignatureTooltip

`func (o *EzsignsignatureResponseCompound) GetTEzsignsignatureTooltip() string`

GetTEzsignsignatureTooltip returns the TEzsignsignatureTooltip field if non-nil, zero value otherwise.

### GetTEzsignsignatureTooltipOk

`func (o *EzsignsignatureResponseCompound) GetTEzsignsignatureTooltipOk() (*string, bool)`

GetTEzsignsignatureTooltipOk returns a tuple with the TEzsignsignatureTooltip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEzsignsignatureTooltip

`func (o *EzsignsignatureResponseCompound) SetTEzsignsignatureTooltip(v string)`

SetTEzsignsignatureTooltip sets TEzsignsignatureTooltip field to given value.

### HasTEzsignsignatureTooltip

`func (o *EzsignsignatureResponseCompound) HasTEzsignsignatureTooltip() bool`

HasTEzsignsignatureTooltip returns a boolean if a field has been set.

### GetEEzsignsignatureTooltipposition

`func (o *EzsignsignatureResponseCompound) GetEEzsignsignatureTooltipposition() FieldEEzsignsignatureTooltipposition`

GetEEzsignsignatureTooltipposition returns the EEzsignsignatureTooltipposition field if non-nil, zero value otherwise.

### GetEEzsignsignatureTooltippositionOk

`func (o *EzsignsignatureResponseCompound) GetEEzsignsignatureTooltippositionOk() (*FieldEEzsignsignatureTooltipposition, bool)`

GetEEzsignsignatureTooltippositionOk returns a tuple with the EEzsignsignatureTooltipposition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignsignatureTooltipposition

`func (o *EzsignsignatureResponseCompound) SetEEzsignsignatureTooltipposition(v FieldEEzsignsignatureTooltipposition)`

SetEEzsignsignatureTooltipposition sets EEzsignsignatureTooltipposition field to given value.

### HasEEzsignsignatureTooltipposition

`func (o *EzsignsignatureResponseCompound) HasEEzsignsignatureTooltipposition() bool`

HasEEzsignsignatureTooltipposition returns a boolean if a field has been set.

### GetEEzsignsignatureFont

`func (o *EzsignsignatureResponseCompound) GetEEzsignsignatureFont() FieldEEzsignsignatureFont`

GetEEzsignsignatureFont returns the EEzsignsignatureFont field if non-nil, zero value otherwise.

### GetEEzsignsignatureFontOk

`func (o *EzsignsignatureResponseCompound) GetEEzsignsignatureFontOk() (*FieldEEzsignsignatureFont, bool)`

GetEEzsignsignatureFontOk returns a tuple with the EEzsignsignatureFont field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignsignatureFont

`func (o *EzsignsignatureResponseCompound) SetEEzsignsignatureFont(v FieldEEzsignsignatureFont)`

SetEEzsignsignatureFont sets EEzsignsignatureFont field to given value.

### HasEEzsignsignatureFont

`func (o *EzsignsignatureResponseCompound) HasEEzsignsignatureFont() bool`

HasEEzsignsignatureFont returns a boolean if a field has been set.

### GetIEzsignsignatureValidationstep

`func (o *EzsignsignatureResponseCompound) GetIEzsignsignatureValidationstep() int32`

GetIEzsignsignatureValidationstep returns the IEzsignsignatureValidationstep field if non-nil, zero value otherwise.

### GetIEzsignsignatureValidationstepOk

`func (o *EzsignsignatureResponseCompound) GetIEzsignsignatureValidationstepOk() (*int32, bool)`

GetIEzsignsignatureValidationstepOk returns a tuple with the IEzsignsignatureValidationstep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignsignatureValidationstep

`func (o *EzsignsignatureResponseCompound) SetIEzsignsignatureValidationstep(v int32)`

SetIEzsignsignatureValidationstep sets IEzsignsignatureValidationstep field to given value.

### HasIEzsignsignatureValidationstep

`func (o *EzsignsignatureResponseCompound) HasIEzsignsignatureValidationstep() bool`

HasIEzsignsignatureValidationstep returns a boolean if a field has been set.

### GetSEzsignsignatureAttachmentdescription

`func (o *EzsignsignatureResponseCompound) GetSEzsignsignatureAttachmentdescription() string`

GetSEzsignsignatureAttachmentdescription returns the SEzsignsignatureAttachmentdescription field if non-nil, zero value otherwise.

### GetSEzsignsignatureAttachmentdescriptionOk

`func (o *EzsignsignatureResponseCompound) GetSEzsignsignatureAttachmentdescriptionOk() (*string, bool)`

GetSEzsignsignatureAttachmentdescriptionOk returns a tuple with the SEzsignsignatureAttachmentdescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignsignatureAttachmentdescription

`func (o *EzsignsignatureResponseCompound) SetSEzsignsignatureAttachmentdescription(v string)`

SetSEzsignsignatureAttachmentdescription sets SEzsignsignatureAttachmentdescription field to given value.

### HasSEzsignsignatureAttachmentdescription

`func (o *EzsignsignatureResponseCompound) HasSEzsignsignatureAttachmentdescription() bool`

HasSEzsignsignatureAttachmentdescription returns a boolean if a field has been set.

### GetEEzsignsignatureAttachmentnamesource

`func (o *EzsignsignatureResponseCompound) GetEEzsignsignatureAttachmentnamesource() FieldEEzsignsignatureAttachmentnamesource`

GetEEzsignsignatureAttachmentnamesource returns the EEzsignsignatureAttachmentnamesource field if non-nil, zero value otherwise.

### GetEEzsignsignatureAttachmentnamesourceOk

`func (o *EzsignsignatureResponseCompound) GetEEzsignsignatureAttachmentnamesourceOk() (*FieldEEzsignsignatureAttachmentnamesource, bool)`

GetEEzsignsignatureAttachmentnamesourceOk returns a tuple with the EEzsignsignatureAttachmentnamesource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignsignatureAttachmentnamesource

`func (o *EzsignsignatureResponseCompound) SetEEzsignsignatureAttachmentnamesource(v FieldEEzsignsignatureAttachmentnamesource)`

SetEEzsignsignatureAttachmentnamesource sets EEzsignsignatureAttachmentnamesource field to given value.

### HasEEzsignsignatureAttachmentnamesource

`func (o *EzsignsignatureResponseCompound) HasEEzsignsignatureAttachmentnamesource() bool`

HasEEzsignsignatureAttachmentnamesource returns a boolean if a field has been set.

### GetBEzsignsignatureRequired

`func (o *EzsignsignatureResponseCompound) GetBEzsignsignatureRequired() bool`

GetBEzsignsignatureRequired returns the BEzsignsignatureRequired field if non-nil, zero value otherwise.

### GetBEzsignsignatureRequiredOk

`func (o *EzsignsignatureResponseCompound) GetBEzsignsignatureRequiredOk() (*bool, bool)`

GetBEzsignsignatureRequiredOk returns a tuple with the BEzsignsignatureRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignsignatureRequired

`func (o *EzsignsignatureResponseCompound) SetBEzsignsignatureRequired(v bool)`

SetBEzsignsignatureRequired sets BEzsignsignatureRequired field to given value.

### HasBEzsignsignatureRequired

`func (o *EzsignsignatureResponseCompound) HasBEzsignsignatureRequired() bool`

HasBEzsignsignatureRequired returns a boolean if a field has been set.

### GetFkiEzsignfoldersignerassociationIDValidation

`func (o *EzsignsignatureResponseCompound) GetFkiEzsignfoldersignerassociationIDValidation() int32`

GetFkiEzsignfoldersignerassociationIDValidation returns the FkiEzsignfoldersignerassociationIDValidation field if non-nil, zero value otherwise.

### GetFkiEzsignfoldersignerassociationIDValidationOk

`func (o *EzsignsignatureResponseCompound) GetFkiEzsignfoldersignerassociationIDValidationOk() (*int32, bool)`

GetFkiEzsignfoldersignerassociationIDValidationOk returns a tuple with the FkiEzsignfoldersignerassociationIDValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfoldersignerassociationIDValidation

`func (o *EzsignsignatureResponseCompound) SetFkiEzsignfoldersignerassociationIDValidation(v int32)`

SetFkiEzsignfoldersignerassociationIDValidation sets FkiEzsignfoldersignerassociationIDValidation field to given value.

### HasFkiEzsignfoldersignerassociationIDValidation

`func (o *EzsignsignatureResponseCompound) HasFkiEzsignfoldersignerassociationIDValidation() bool`

HasFkiEzsignfoldersignerassociationIDValidation returns a boolean if a field has been set.

### GetDtEzsignsignatureDate

`func (o *EzsignsignatureResponseCompound) GetDtEzsignsignatureDate() string`

GetDtEzsignsignatureDate returns the DtEzsignsignatureDate field if non-nil, zero value otherwise.

### GetDtEzsignsignatureDateOk

`func (o *EzsignsignatureResponseCompound) GetDtEzsignsignatureDateOk() (*string, bool)`

GetDtEzsignsignatureDateOk returns a tuple with the DtEzsignsignatureDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtEzsignsignatureDate

`func (o *EzsignsignatureResponseCompound) SetDtEzsignsignatureDate(v string)`

SetDtEzsignsignatureDate sets DtEzsignsignatureDate field to given value.

### HasDtEzsignsignatureDate

`func (o *EzsignsignatureResponseCompound) HasDtEzsignsignatureDate() bool`

HasDtEzsignsignatureDate returns a boolean if a field has been set.

### GetIEzsignsignatureattachmentCount

`func (o *EzsignsignatureResponseCompound) GetIEzsignsignatureattachmentCount() int32`

GetIEzsignsignatureattachmentCount returns the IEzsignsignatureattachmentCount field if non-nil, zero value otherwise.

### GetIEzsignsignatureattachmentCountOk

`func (o *EzsignsignatureResponseCompound) GetIEzsignsignatureattachmentCountOk() (*int32, bool)`

GetIEzsignsignatureattachmentCountOk returns a tuple with the IEzsignsignatureattachmentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignsignatureattachmentCount

`func (o *EzsignsignatureResponseCompound) SetIEzsignsignatureattachmentCount(v int32)`

SetIEzsignsignatureattachmentCount sets IEzsignsignatureattachmentCount field to given value.

### HasIEzsignsignatureattachmentCount

`func (o *EzsignsignatureResponseCompound) HasIEzsignsignatureattachmentCount() bool`

HasIEzsignsignatureattachmentCount returns a boolean if a field has been set.

### GetSEzsignsignatureDescription

`func (o *EzsignsignatureResponseCompound) GetSEzsignsignatureDescription() string`

GetSEzsignsignatureDescription returns the SEzsignsignatureDescription field if non-nil, zero value otherwise.

### GetSEzsignsignatureDescriptionOk

`func (o *EzsignsignatureResponseCompound) GetSEzsignsignatureDescriptionOk() (*string, bool)`

GetSEzsignsignatureDescriptionOk returns a tuple with the SEzsignsignatureDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignsignatureDescription

`func (o *EzsignsignatureResponseCompound) SetSEzsignsignatureDescription(v string)`

SetSEzsignsignatureDescription sets SEzsignsignatureDescription field to given value.

### HasSEzsignsignatureDescription

`func (o *EzsignsignatureResponseCompound) HasSEzsignsignatureDescription() bool`

HasSEzsignsignatureDescription returns a boolean if a field has been set.

### GetIEzsignsignatureMaxlength

`func (o *EzsignsignatureResponseCompound) GetIEzsignsignatureMaxlength() int32`

GetIEzsignsignatureMaxlength returns the IEzsignsignatureMaxlength field if non-nil, zero value otherwise.

### GetIEzsignsignatureMaxlengthOk

`func (o *EzsignsignatureResponseCompound) GetIEzsignsignatureMaxlengthOk() (*int32, bool)`

GetIEzsignsignatureMaxlengthOk returns a tuple with the IEzsignsignatureMaxlength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignsignatureMaxlength

`func (o *EzsignsignatureResponseCompound) SetIEzsignsignatureMaxlength(v int32)`

SetIEzsignsignatureMaxlength sets IEzsignsignatureMaxlength field to given value.

### HasIEzsignsignatureMaxlength

`func (o *EzsignsignatureResponseCompound) HasIEzsignsignatureMaxlength() bool`

HasIEzsignsignatureMaxlength returns a boolean if a field has been set.

### GetEEzsignsignatureTextvalidation

`func (o *EzsignsignatureResponseCompound) GetEEzsignsignatureTextvalidation() EnumTextvalidation`

GetEEzsignsignatureTextvalidation returns the EEzsignsignatureTextvalidation field if non-nil, zero value otherwise.

### GetEEzsignsignatureTextvalidationOk

`func (o *EzsignsignatureResponseCompound) GetEEzsignsignatureTextvalidationOk() (*EnumTextvalidation, bool)`

GetEEzsignsignatureTextvalidationOk returns a tuple with the EEzsignsignatureTextvalidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignsignatureTextvalidation

`func (o *EzsignsignatureResponseCompound) SetEEzsignsignatureTextvalidation(v EnumTextvalidation)`

SetEEzsignsignatureTextvalidation sets EEzsignsignatureTextvalidation field to given value.

### HasEEzsignsignatureTextvalidation

`func (o *EzsignsignatureResponseCompound) HasEEzsignsignatureTextvalidation() bool`

HasEEzsignsignatureTextvalidation returns a boolean if a field has been set.

### GetEEzsignsignatureDependencyrequirement

`func (o *EzsignsignatureResponseCompound) GetEEzsignsignatureDependencyrequirement() FieldEEzsignsignatureDependencyrequirement`

GetEEzsignsignatureDependencyrequirement returns the EEzsignsignatureDependencyrequirement field if non-nil, zero value otherwise.

### GetEEzsignsignatureDependencyrequirementOk

`func (o *EzsignsignatureResponseCompound) GetEEzsignsignatureDependencyrequirementOk() (*FieldEEzsignsignatureDependencyrequirement, bool)`

GetEEzsignsignatureDependencyrequirementOk returns a tuple with the EEzsignsignatureDependencyrequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignsignatureDependencyrequirement

`func (o *EzsignsignatureResponseCompound) SetEEzsignsignatureDependencyrequirement(v FieldEEzsignsignatureDependencyrequirement)`

SetEEzsignsignatureDependencyrequirement sets EEzsignsignatureDependencyrequirement field to given value.

### HasEEzsignsignatureDependencyrequirement

`func (o *EzsignsignatureResponseCompound) HasEEzsignsignatureDependencyrequirement() bool`

HasEEzsignsignatureDependencyrequirement returns a boolean if a field has been set.

### GetSEzsignsignatureRegexp

`func (o *EzsignsignatureResponseCompound) GetSEzsignsignatureRegexp() string`

GetSEzsignsignatureRegexp returns the SEzsignsignatureRegexp field if non-nil, zero value otherwise.

### GetSEzsignsignatureRegexpOk

`func (o *EzsignsignatureResponseCompound) GetSEzsignsignatureRegexpOk() (*string, bool)`

GetSEzsignsignatureRegexpOk returns a tuple with the SEzsignsignatureRegexp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignsignatureRegexp

`func (o *EzsignsignatureResponseCompound) SetSEzsignsignatureRegexp(v string)`

SetSEzsignsignatureRegexp sets SEzsignsignatureRegexp field to given value.

### HasSEzsignsignatureRegexp

`func (o *EzsignsignatureResponseCompound) HasSEzsignsignatureRegexp() bool`

HasSEzsignsignatureRegexp returns a boolean if a field has been set.

### GetObjContactName

`func (o *EzsignsignatureResponseCompound) GetObjContactName() CustomContactNameResponse`

GetObjContactName returns the ObjContactName field if non-nil, zero value otherwise.

### GetObjContactNameOk

`func (o *EzsignsignatureResponseCompound) GetObjContactNameOk() (*CustomContactNameResponse, bool)`

GetObjContactNameOk returns a tuple with the ObjContactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjContactName

`func (o *EzsignsignatureResponseCompound) SetObjContactName(v CustomContactNameResponse)`

SetObjContactName sets ObjContactName field to given value.


### GetObjContactNameDelegation

`func (o *EzsignsignatureResponseCompound) GetObjContactNameDelegation() CustomContactNameResponse`

GetObjContactNameDelegation returns the ObjContactNameDelegation field if non-nil, zero value otherwise.

### GetObjContactNameDelegationOk

`func (o *EzsignsignatureResponseCompound) GetObjContactNameDelegationOk() (*CustomContactNameResponse, bool)`

GetObjContactNameDelegationOk returns a tuple with the ObjContactNameDelegation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjContactNameDelegation

`func (o *EzsignsignatureResponseCompound) SetObjContactNameDelegation(v CustomContactNameResponse)`

SetObjContactNameDelegation sets ObjContactNameDelegation field to given value.

### HasObjContactNameDelegation

`func (o *EzsignsignatureResponseCompound) HasObjContactNameDelegation() bool`

HasObjContactNameDelegation returns a boolean if a field has been set.

### GetObjSignature

`func (o *EzsignsignatureResponseCompound) GetObjSignature() SignatureResponseCompound`

GetObjSignature returns the ObjSignature field if non-nil, zero value otherwise.

### GetObjSignatureOk

`func (o *EzsignsignatureResponseCompound) GetObjSignatureOk() (*SignatureResponseCompound, bool)`

GetObjSignatureOk returns a tuple with the ObjSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjSignature

`func (o *EzsignsignatureResponseCompound) SetObjSignature(v SignatureResponseCompound)`

SetObjSignature sets ObjSignature field to given value.

### HasObjSignature

`func (o *EzsignsignatureResponseCompound) HasObjSignature() bool`

HasObjSignature returns a boolean if a field has been set.

### GetBEzsignsignatureCustomdate

`func (o *EzsignsignatureResponseCompound) GetBEzsignsignatureCustomdate() bool`

GetBEzsignsignatureCustomdate returns the BEzsignsignatureCustomdate field if non-nil, zero value otherwise.

### GetBEzsignsignatureCustomdateOk

`func (o *EzsignsignatureResponseCompound) GetBEzsignsignatureCustomdateOk() (*bool, bool)`

GetBEzsignsignatureCustomdateOk returns a tuple with the BEzsignsignatureCustomdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignsignatureCustomdate

`func (o *EzsignsignatureResponseCompound) SetBEzsignsignatureCustomdate(v bool)`

SetBEzsignsignatureCustomdate sets BEzsignsignatureCustomdate field to given value.

### HasBEzsignsignatureCustomdate

`func (o *EzsignsignatureResponseCompound) HasBEzsignsignatureCustomdate() bool`

HasBEzsignsignatureCustomdate returns a boolean if a field has been set.

### GetAObjEzsignsignaturecustomdate

`func (o *EzsignsignatureResponseCompound) GetAObjEzsignsignaturecustomdate() []EzsignsignaturecustomdateResponseCompound`

GetAObjEzsignsignaturecustomdate returns the AObjEzsignsignaturecustomdate field if non-nil, zero value otherwise.

### GetAObjEzsignsignaturecustomdateOk

`func (o *EzsignsignatureResponseCompound) GetAObjEzsignsignaturecustomdateOk() (*[]EzsignsignaturecustomdateResponseCompound, bool)`

GetAObjEzsignsignaturecustomdateOk returns a tuple with the AObjEzsignsignaturecustomdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzsignsignaturecustomdate

`func (o *EzsignsignatureResponseCompound) SetAObjEzsignsignaturecustomdate(v []EzsignsignaturecustomdateResponseCompound)`

SetAObjEzsignsignaturecustomdate sets AObjEzsignsignaturecustomdate field to given value.

### HasAObjEzsignsignaturecustomdate

`func (o *EzsignsignatureResponseCompound) HasAObjEzsignsignaturecustomdate() bool`

HasAObjEzsignsignaturecustomdate returns a boolean if a field has been set.

### GetObjCreditcardtransaction

`func (o *EzsignsignatureResponseCompound) GetObjCreditcardtransaction() CustomCreditcardtransactionResponse`

GetObjCreditcardtransaction returns the ObjCreditcardtransaction field if non-nil, zero value otherwise.

### GetObjCreditcardtransactionOk

`func (o *EzsignsignatureResponseCompound) GetObjCreditcardtransactionOk() (*CustomCreditcardtransactionResponse, bool)`

GetObjCreditcardtransactionOk returns a tuple with the ObjCreditcardtransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjCreditcardtransaction

`func (o *EzsignsignatureResponseCompound) SetObjCreditcardtransaction(v CustomCreditcardtransactionResponse)`

SetObjCreditcardtransaction sets ObjCreditcardtransaction field to given value.

### HasObjCreditcardtransaction

`func (o *EzsignsignatureResponseCompound) HasObjCreditcardtransaction() bool`

HasObjCreditcardtransaction returns a boolean if a field has been set.

### GetAObjEzsignelementdependency

`func (o *EzsignsignatureResponseCompound) GetAObjEzsignelementdependency() []EzsignelementdependencyResponseCompound`

GetAObjEzsignelementdependency returns the AObjEzsignelementdependency field if non-nil, zero value otherwise.

### GetAObjEzsignelementdependencyOk

`func (o *EzsignsignatureResponseCompound) GetAObjEzsignelementdependencyOk() (*[]EzsignelementdependencyResponseCompound, bool)`

GetAObjEzsignelementdependencyOk returns a tuple with the AObjEzsignelementdependency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzsignelementdependency

`func (o *EzsignsignatureResponseCompound) SetAObjEzsignelementdependency(v []EzsignelementdependencyResponseCompound)`

SetAObjEzsignelementdependency sets AObjEzsignelementdependency field to given value.

### HasAObjEzsignelementdependency

`func (o *EzsignsignatureResponseCompound) HasAObjEzsignelementdependency() bool`

HasAObjEzsignelementdependency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


