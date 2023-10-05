/*
eZmax API Definition (Full)

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.2.0
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
)

// checks if the EzsigntemplatesignatureResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplatesignatureResponse{}

// EzsigntemplatesignatureResponse A Ezsigntemplatesignature Object
type EzsigntemplatesignatureResponse struct {
	// The unique ID of the Ezsigntemplatesignature
	PkiEzsigntemplatesignatureID int32 `json:"pkiEzsigntemplatesignatureID"`
	// The unique ID of the Ezsigntemplatedocument
	FkiEzsigntemplatedocumentID int32 `json:"fkiEzsigntemplatedocumentID"`
	// The unique ID of the Ezsigntemplatesigner
	FkiEzsigntemplatesignerID int32 `json:"fkiEzsigntemplatesignerID"`
	// The unique ID of the Ezsigntemplatesigner
	FkiEzsigntemplatesignerIDValidation *int32 `json:"fkiEzsigntemplatesignerIDValidation,omitempty"`
	// The page number in the Ezsigntemplatedocument
	IEzsigntemplatedocumentpagePagenumber int32 `json:"iEzsigntemplatedocumentpagePagenumber"`
	// The X coordinate (Horizontal) where to put the Ezsigntemplatesignature on the page.  Coordinate is calculated at 100dpi (dot per inch). So for example, if you want to put the Ezsigntemplatesignature 2 inches from the left border of the page, you would use \"200\" for the X coordinate.
	IEzsigntemplatesignatureX int32 `json:"iEzsigntemplatesignatureX"`
	// The Y coordinate (Vertical) where to put the Ezsigntemplatesignature on the page.  Coordinate is calculated at 100dpi (dot per inch). So for example, if you want to put the Ezsigntemplatesignature 3 inches from the top border of the page, you would use \"300\" for the Y coordinate.
	IEzsigntemplatesignatureY int32 `json:"iEzsigntemplatesignatureY"`
	// The width of the Ezsigntemplatesignature.  Size is calculated at 100dpi (dot per inch). So for example, if you want the Ezsigntemplatesignature to have a width of 2 inches, you would use \"200\" for the iEzsigntemplatesignatureWidth.
	IEzsigntemplatesignatureWidth *int32 `json:"iEzsigntemplatesignatureWidth,omitempty"`
	// The height of the Ezsigntemplatesignature.  Size is calculated at 100dpi (dot per inch). So for example, if you want the Ezsigntemplatesignature to have an height of 2 inches, you would use \"200\" for the iEzsigntemplatesignatureHeight.
	IEzsigntemplatesignatureHeight *int32 `json:"iEzsigntemplatesignatureHeight,omitempty"`
	// The step when the Ezsigntemplatesigner will be invited to sign
	IEzsigntemplatesignatureStep int32 `json:"iEzsigntemplatesignatureStep"`
	EEzsigntemplatesignatureType FieldEEzsigntemplatesignatureType `json:"eEzsigntemplatesignatureType"`
	// A tooltip that will be presented to Ezsigntemplatesigner about the Ezsigntemplatesignature
	TEzsigntemplatesignatureTooltip *string `json:"tEzsigntemplatesignatureTooltip,omitempty"`
	EEzsigntemplatesignatureTooltipposition *FieldEEzsigntemplatesignatureTooltipposition `json:"eEzsigntemplatesignatureTooltipposition,omitempty"`
	EEzsigntemplatesignatureFont *FieldEEzsigntemplatesignatureFont `json:"eEzsigntemplatesignatureFont,omitempty"`
	// The step when the Ezsigntemplatesigner will be invited to validate the Ezsigntemplatesignature of eEzsigntemplatesignatureType Attachments
	IEzsigntemplatesignatureValidationstep *int32 `json:"iEzsigntemplatesignatureValidationstep,omitempty"`
	// The description attached to the attachment name added in Ezsigntemplatesignature of eEzsigntemplatesignatureType Attachments
	SEzsigntemplatesignatureAttachmentdescription *string `json:"sEzsigntemplatesignatureAttachmentdescription,omitempty"`
	EEzsigntemplatesignatureAttachmentnamesource *FieldEEzsigntemplatesignatureAttachmentnamesource `json:"eEzsigntemplatesignatureAttachmentnamesource,omitempty"`
	// Whether the Ezsigntemplatesignature is required or not. This field is relevant only with Ezsigntemplatesignature with eEzsigntemplatesignatureType = Attachments.
	BEzsigntemplatesignatureRequired *bool `json:"bEzsigntemplatesignatureRequired,omitempty"`
	// The maximum length for the value in the Ezsigntemplatesignature  This can only be set if eEzsigntemplatesignatureType is **FieldText** or **FieldTextarea**
	IEzsigntemplatesignatureMaxlength *int32 `json:"iEzsigntemplatesignatureMaxlength,omitempty"`
	// A regular expression to indicate what values are acceptable for the Ezsigntemplatesignature.  This can only be set if eEzsigntemplatesignatureType is **Text** or **Textarea**
	SEzsigntemplatesignatureRegexp *string `json:"sEzsigntemplatesignatureRegexp,omitempty"`
	EEzsigntemplatesignatureTextvalidation *EnumTextvalidation `json:"eEzsigntemplatesignatureTextvalidation,omitempty"`
	EEzsigntemplatesignatureDependencyrequirement *FieldEEzsigntemplatesignatureDependencyrequirement `json:"eEzsigntemplatesignatureDependencyrequirement,omitempty"`
}

// NewEzsigntemplatesignatureResponse instantiates a new EzsigntemplatesignatureResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplatesignatureResponse(pkiEzsigntemplatesignatureID int32, fkiEzsigntemplatedocumentID int32, fkiEzsigntemplatesignerID int32, iEzsigntemplatedocumentpagePagenumber int32, iEzsigntemplatesignatureX int32, iEzsigntemplatesignatureY int32, iEzsigntemplatesignatureStep int32, eEzsigntemplatesignatureType FieldEEzsigntemplatesignatureType) *EzsigntemplatesignatureResponse {
	this := EzsigntemplatesignatureResponse{}
	this.PkiEzsigntemplatesignatureID = pkiEzsigntemplatesignatureID
	this.FkiEzsigntemplatedocumentID = fkiEzsigntemplatedocumentID
	this.FkiEzsigntemplatesignerID = fkiEzsigntemplatesignerID
	this.IEzsigntemplatedocumentpagePagenumber = iEzsigntemplatedocumentpagePagenumber
	this.IEzsigntemplatesignatureX = iEzsigntemplatesignatureX
	this.IEzsigntemplatesignatureY = iEzsigntemplatesignatureY
	this.IEzsigntemplatesignatureStep = iEzsigntemplatesignatureStep
	this.EEzsigntemplatesignatureType = eEzsigntemplatesignatureType
	return &this
}

// NewEzsigntemplatesignatureResponseWithDefaults instantiates a new EzsigntemplatesignatureResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplatesignatureResponseWithDefaults() *EzsigntemplatesignatureResponse {
	this := EzsigntemplatesignatureResponse{}
	return &this
}

// GetPkiEzsigntemplatesignatureID returns the PkiEzsigntemplatesignatureID field value
func (o *EzsigntemplatesignatureResponse) GetPkiEzsigntemplatesignatureID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiEzsigntemplatesignatureID
}

// GetPkiEzsigntemplatesignatureIDOk returns a tuple with the PkiEzsigntemplatesignatureID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignatureResponse) GetPkiEzsigntemplatesignatureIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiEzsigntemplatesignatureID, true
}

// SetPkiEzsigntemplatesignatureID sets field value
func (o *EzsigntemplatesignatureResponse) SetPkiEzsigntemplatesignatureID(v int32) {
	o.PkiEzsigntemplatesignatureID = v
}

// GetFkiEzsigntemplatedocumentID returns the FkiEzsigntemplatedocumentID field value
func (o *EzsigntemplatesignatureResponse) GetFkiEzsigntemplatedocumentID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsigntemplatedocumentID
}

// GetFkiEzsigntemplatedocumentIDOk returns a tuple with the FkiEzsigntemplatedocumentID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignatureResponse) GetFkiEzsigntemplatedocumentIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzsigntemplatedocumentID, true
}

// SetFkiEzsigntemplatedocumentID sets field value
func (o *EzsigntemplatesignatureResponse) SetFkiEzsigntemplatedocumentID(v int32) {
	o.FkiEzsigntemplatedocumentID = v
}

// GetFkiEzsigntemplatesignerID returns the FkiEzsigntemplatesignerID field value
func (o *EzsigntemplatesignatureResponse) GetFkiEzsigntemplatesignerID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsigntemplatesignerID
}

// GetFkiEzsigntemplatesignerIDOk returns a tuple with the FkiEzsigntemplatesignerID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignatureResponse) GetFkiEzsigntemplatesignerIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzsigntemplatesignerID, true
}

// SetFkiEzsigntemplatesignerID sets field value
func (o *EzsigntemplatesignatureResponse) SetFkiEzsigntemplatesignerID(v int32) {
	o.FkiEzsigntemplatesignerID = v
}

// GetFkiEzsigntemplatesignerIDValidation returns the FkiEzsigntemplatesignerIDValidation field value if set, zero value otherwise.
func (o *EzsigntemplatesignatureResponse) GetFkiEzsigntemplatesignerIDValidation() int32 {
	if o == nil || IsNil(o.FkiEzsigntemplatesignerIDValidation) {
		var ret int32
		return ret
	}
	return *o.FkiEzsigntemplatesignerIDValidation
}

// GetFkiEzsigntemplatesignerIDValidationOk returns a tuple with the FkiEzsigntemplatesignerIDValidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignatureResponse) GetFkiEzsigntemplatesignerIDValidationOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiEzsigntemplatesignerIDValidation) {
		return nil, false
	}
	return o.FkiEzsigntemplatesignerIDValidation, true
}

// HasFkiEzsigntemplatesignerIDValidation returns a boolean if a field has been set.
func (o *EzsigntemplatesignatureResponse) HasFkiEzsigntemplatesignerIDValidation() bool {
	if o != nil && !IsNil(o.FkiEzsigntemplatesignerIDValidation) {
		return true
	}

	return false
}

// SetFkiEzsigntemplatesignerIDValidation gets a reference to the given int32 and assigns it to the FkiEzsigntemplatesignerIDValidation field.
func (o *EzsigntemplatesignatureResponse) SetFkiEzsigntemplatesignerIDValidation(v int32) {
	o.FkiEzsigntemplatesignerIDValidation = &v
}

// GetIEzsigntemplatedocumentpagePagenumber returns the IEzsigntemplatedocumentpagePagenumber field value
func (o *EzsigntemplatesignatureResponse) GetIEzsigntemplatedocumentpagePagenumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsigntemplatedocumentpagePagenumber
}

// GetIEzsigntemplatedocumentpagePagenumberOk returns a tuple with the IEzsigntemplatedocumentpagePagenumber field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignatureResponse) GetIEzsigntemplatedocumentpagePagenumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsigntemplatedocumentpagePagenumber, true
}

// SetIEzsigntemplatedocumentpagePagenumber sets field value
func (o *EzsigntemplatesignatureResponse) SetIEzsigntemplatedocumentpagePagenumber(v int32) {
	o.IEzsigntemplatedocumentpagePagenumber = v
}

// GetIEzsigntemplatesignatureX returns the IEzsigntemplatesignatureX field value
func (o *EzsigntemplatesignatureResponse) GetIEzsigntemplatesignatureX() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsigntemplatesignatureX
}

// GetIEzsigntemplatesignatureXOk returns a tuple with the IEzsigntemplatesignatureX field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignatureResponse) GetIEzsigntemplatesignatureXOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsigntemplatesignatureX, true
}

// SetIEzsigntemplatesignatureX sets field value
func (o *EzsigntemplatesignatureResponse) SetIEzsigntemplatesignatureX(v int32) {
	o.IEzsigntemplatesignatureX = v
}

// GetIEzsigntemplatesignatureY returns the IEzsigntemplatesignatureY field value
func (o *EzsigntemplatesignatureResponse) GetIEzsigntemplatesignatureY() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsigntemplatesignatureY
}

// GetIEzsigntemplatesignatureYOk returns a tuple with the IEzsigntemplatesignatureY field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignatureResponse) GetIEzsigntemplatesignatureYOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsigntemplatesignatureY, true
}

// SetIEzsigntemplatesignatureY sets field value
func (o *EzsigntemplatesignatureResponse) SetIEzsigntemplatesignatureY(v int32) {
	o.IEzsigntemplatesignatureY = v
}

// GetIEzsigntemplatesignatureWidth returns the IEzsigntemplatesignatureWidth field value if set, zero value otherwise.
func (o *EzsigntemplatesignatureResponse) GetIEzsigntemplatesignatureWidth() int32 {
	if o == nil || IsNil(o.IEzsigntemplatesignatureWidth) {
		var ret int32
		return ret
	}
	return *o.IEzsigntemplatesignatureWidth
}

// GetIEzsigntemplatesignatureWidthOk returns a tuple with the IEzsigntemplatesignatureWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignatureResponse) GetIEzsigntemplatesignatureWidthOk() (*int32, bool) {
	if o == nil || IsNil(o.IEzsigntemplatesignatureWidth) {
		return nil, false
	}
	return o.IEzsigntemplatesignatureWidth, true
}

// HasIEzsigntemplatesignatureWidth returns a boolean if a field has been set.
func (o *EzsigntemplatesignatureResponse) HasIEzsigntemplatesignatureWidth() bool {
	if o != nil && !IsNil(o.IEzsigntemplatesignatureWidth) {
		return true
	}

	return false
}

// SetIEzsigntemplatesignatureWidth gets a reference to the given int32 and assigns it to the IEzsigntemplatesignatureWidth field.
func (o *EzsigntemplatesignatureResponse) SetIEzsigntemplatesignatureWidth(v int32) {
	o.IEzsigntemplatesignatureWidth = &v
}

// GetIEzsigntemplatesignatureHeight returns the IEzsigntemplatesignatureHeight field value if set, zero value otherwise.
func (o *EzsigntemplatesignatureResponse) GetIEzsigntemplatesignatureHeight() int32 {
	if o == nil || IsNil(o.IEzsigntemplatesignatureHeight) {
		var ret int32
		return ret
	}
	return *o.IEzsigntemplatesignatureHeight
}

// GetIEzsigntemplatesignatureHeightOk returns a tuple with the IEzsigntemplatesignatureHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignatureResponse) GetIEzsigntemplatesignatureHeightOk() (*int32, bool) {
	if o == nil || IsNil(o.IEzsigntemplatesignatureHeight) {
		return nil, false
	}
	return o.IEzsigntemplatesignatureHeight, true
}

// HasIEzsigntemplatesignatureHeight returns a boolean if a field has been set.
func (o *EzsigntemplatesignatureResponse) HasIEzsigntemplatesignatureHeight() bool {
	if o != nil && !IsNil(o.IEzsigntemplatesignatureHeight) {
		return true
	}

	return false
}

// SetIEzsigntemplatesignatureHeight gets a reference to the given int32 and assigns it to the IEzsigntemplatesignatureHeight field.
func (o *EzsigntemplatesignatureResponse) SetIEzsigntemplatesignatureHeight(v int32) {
	o.IEzsigntemplatesignatureHeight = &v
}

// GetIEzsigntemplatesignatureStep returns the IEzsigntemplatesignatureStep field value
func (o *EzsigntemplatesignatureResponse) GetIEzsigntemplatesignatureStep() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsigntemplatesignatureStep
}

// GetIEzsigntemplatesignatureStepOk returns a tuple with the IEzsigntemplatesignatureStep field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignatureResponse) GetIEzsigntemplatesignatureStepOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsigntemplatesignatureStep, true
}

// SetIEzsigntemplatesignatureStep sets field value
func (o *EzsigntemplatesignatureResponse) SetIEzsigntemplatesignatureStep(v int32) {
	o.IEzsigntemplatesignatureStep = v
}

// GetEEzsigntemplatesignatureType returns the EEzsigntemplatesignatureType field value
func (o *EzsigntemplatesignatureResponse) GetEEzsigntemplatesignatureType() FieldEEzsigntemplatesignatureType {
	if o == nil {
		var ret FieldEEzsigntemplatesignatureType
		return ret
	}

	return o.EEzsigntemplatesignatureType
}

// GetEEzsigntemplatesignatureTypeOk returns a tuple with the EEzsigntemplatesignatureType field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignatureResponse) GetEEzsigntemplatesignatureTypeOk() (*FieldEEzsigntemplatesignatureType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EEzsigntemplatesignatureType, true
}

// SetEEzsigntemplatesignatureType sets field value
func (o *EzsigntemplatesignatureResponse) SetEEzsigntemplatesignatureType(v FieldEEzsigntemplatesignatureType) {
	o.EEzsigntemplatesignatureType = v
}

// GetTEzsigntemplatesignatureTooltip returns the TEzsigntemplatesignatureTooltip field value if set, zero value otherwise.
func (o *EzsigntemplatesignatureResponse) GetTEzsigntemplatesignatureTooltip() string {
	if o == nil || IsNil(o.TEzsigntemplatesignatureTooltip) {
		var ret string
		return ret
	}
	return *o.TEzsigntemplatesignatureTooltip
}

// GetTEzsigntemplatesignatureTooltipOk returns a tuple with the TEzsigntemplatesignatureTooltip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignatureResponse) GetTEzsigntemplatesignatureTooltipOk() (*string, bool) {
	if o == nil || IsNil(o.TEzsigntemplatesignatureTooltip) {
		return nil, false
	}
	return o.TEzsigntemplatesignatureTooltip, true
}

// HasTEzsigntemplatesignatureTooltip returns a boolean if a field has been set.
func (o *EzsigntemplatesignatureResponse) HasTEzsigntemplatesignatureTooltip() bool {
	if o != nil && !IsNil(o.TEzsigntemplatesignatureTooltip) {
		return true
	}

	return false
}

// SetTEzsigntemplatesignatureTooltip gets a reference to the given string and assigns it to the TEzsigntemplatesignatureTooltip field.
func (o *EzsigntemplatesignatureResponse) SetTEzsigntemplatesignatureTooltip(v string) {
	o.TEzsigntemplatesignatureTooltip = &v
}

// GetEEzsigntemplatesignatureTooltipposition returns the EEzsigntemplatesignatureTooltipposition field value if set, zero value otherwise.
func (o *EzsigntemplatesignatureResponse) GetEEzsigntemplatesignatureTooltipposition() FieldEEzsigntemplatesignatureTooltipposition {
	if o == nil || IsNil(o.EEzsigntemplatesignatureTooltipposition) {
		var ret FieldEEzsigntemplatesignatureTooltipposition
		return ret
	}
	return *o.EEzsigntemplatesignatureTooltipposition
}

// GetEEzsigntemplatesignatureTooltippositionOk returns a tuple with the EEzsigntemplatesignatureTooltipposition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignatureResponse) GetEEzsigntemplatesignatureTooltippositionOk() (*FieldEEzsigntemplatesignatureTooltipposition, bool) {
	if o == nil || IsNil(o.EEzsigntemplatesignatureTooltipposition) {
		return nil, false
	}
	return o.EEzsigntemplatesignatureTooltipposition, true
}

// HasEEzsigntemplatesignatureTooltipposition returns a boolean if a field has been set.
func (o *EzsigntemplatesignatureResponse) HasEEzsigntemplatesignatureTooltipposition() bool {
	if o != nil && !IsNil(o.EEzsigntemplatesignatureTooltipposition) {
		return true
	}

	return false
}

// SetEEzsigntemplatesignatureTooltipposition gets a reference to the given FieldEEzsigntemplatesignatureTooltipposition and assigns it to the EEzsigntemplatesignatureTooltipposition field.
func (o *EzsigntemplatesignatureResponse) SetEEzsigntemplatesignatureTooltipposition(v FieldEEzsigntemplatesignatureTooltipposition) {
	o.EEzsigntemplatesignatureTooltipposition = &v
}

// GetEEzsigntemplatesignatureFont returns the EEzsigntemplatesignatureFont field value if set, zero value otherwise.
func (o *EzsigntemplatesignatureResponse) GetEEzsigntemplatesignatureFont() FieldEEzsigntemplatesignatureFont {
	if o == nil || IsNil(o.EEzsigntemplatesignatureFont) {
		var ret FieldEEzsigntemplatesignatureFont
		return ret
	}
	return *o.EEzsigntemplatesignatureFont
}

// GetEEzsigntemplatesignatureFontOk returns a tuple with the EEzsigntemplatesignatureFont field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignatureResponse) GetEEzsigntemplatesignatureFontOk() (*FieldEEzsigntemplatesignatureFont, bool) {
	if o == nil || IsNil(o.EEzsigntemplatesignatureFont) {
		return nil, false
	}
	return o.EEzsigntemplatesignatureFont, true
}

// HasEEzsigntemplatesignatureFont returns a boolean if a field has been set.
func (o *EzsigntemplatesignatureResponse) HasEEzsigntemplatesignatureFont() bool {
	if o != nil && !IsNil(o.EEzsigntemplatesignatureFont) {
		return true
	}

	return false
}

// SetEEzsigntemplatesignatureFont gets a reference to the given FieldEEzsigntemplatesignatureFont and assigns it to the EEzsigntemplatesignatureFont field.
func (o *EzsigntemplatesignatureResponse) SetEEzsigntemplatesignatureFont(v FieldEEzsigntemplatesignatureFont) {
	o.EEzsigntemplatesignatureFont = &v
}

// GetIEzsigntemplatesignatureValidationstep returns the IEzsigntemplatesignatureValidationstep field value if set, zero value otherwise.
func (o *EzsigntemplatesignatureResponse) GetIEzsigntemplatesignatureValidationstep() int32 {
	if o == nil || IsNil(o.IEzsigntemplatesignatureValidationstep) {
		var ret int32
		return ret
	}
	return *o.IEzsigntemplatesignatureValidationstep
}

// GetIEzsigntemplatesignatureValidationstepOk returns a tuple with the IEzsigntemplatesignatureValidationstep field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignatureResponse) GetIEzsigntemplatesignatureValidationstepOk() (*int32, bool) {
	if o == nil || IsNil(o.IEzsigntemplatesignatureValidationstep) {
		return nil, false
	}
	return o.IEzsigntemplatesignatureValidationstep, true
}

// HasIEzsigntemplatesignatureValidationstep returns a boolean if a field has been set.
func (o *EzsigntemplatesignatureResponse) HasIEzsigntemplatesignatureValidationstep() bool {
	if o != nil && !IsNil(o.IEzsigntemplatesignatureValidationstep) {
		return true
	}

	return false
}

// SetIEzsigntemplatesignatureValidationstep gets a reference to the given int32 and assigns it to the IEzsigntemplatesignatureValidationstep field.
func (o *EzsigntemplatesignatureResponse) SetIEzsigntemplatesignatureValidationstep(v int32) {
	o.IEzsigntemplatesignatureValidationstep = &v
}

// GetSEzsigntemplatesignatureAttachmentdescription returns the SEzsigntemplatesignatureAttachmentdescription field value if set, zero value otherwise.
func (o *EzsigntemplatesignatureResponse) GetSEzsigntemplatesignatureAttachmentdescription() string {
	if o == nil || IsNil(o.SEzsigntemplatesignatureAttachmentdescription) {
		var ret string
		return ret
	}
	return *o.SEzsigntemplatesignatureAttachmentdescription
}

// GetSEzsigntemplatesignatureAttachmentdescriptionOk returns a tuple with the SEzsigntemplatesignatureAttachmentdescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignatureResponse) GetSEzsigntemplatesignatureAttachmentdescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.SEzsigntemplatesignatureAttachmentdescription) {
		return nil, false
	}
	return o.SEzsigntemplatesignatureAttachmentdescription, true
}

// HasSEzsigntemplatesignatureAttachmentdescription returns a boolean if a field has been set.
func (o *EzsigntemplatesignatureResponse) HasSEzsigntemplatesignatureAttachmentdescription() bool {
	if o != nil && !IsNil(o.SEzsigntemplatesignatureAttachmentdescription) {
		return true
	}

	return false
}

// SetSEzsigntemplatesignatureAttachmentdescription gets a reference to the given string and assigns it to the SEzsigntemplatesignatureAttachmentdescription field.
func (o *EzsigntemplatesignatureResponse) SetSEzsigntemplatesignatureAttachmentdescription(v string) {
	o.SEzsigntemplatesignatureAttachmentdescription = &v
}

// GetEEzsigntemplatesignatureAttachmentnamesource returns the EEzsigntemplatesignatureAttachmentnamesource field value if set, zero value otherwise.
func (o *EzsigntemplatesignatureResponse) GetEEzsigntemplatesignatureAttachmentnamesource() FieldEEzsigntemplatesignatureAttachmentnamesource {
	if o == nil || IsNil(o.EEzsigntemplatesignatureAttachmentnamesource) {
		var ret FieldEEzsigntemplatesignatureAttachmentnamesource
		return ret
	}
	return *o.EEzsigntemplatesignatureAttachmentnamesource
}

// GetEEzsigntemplatesignatureAttachmentnamesourceOk returns a tuple with the EEzsigntemplatesignatureAttachmentnamesource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignatureResponse) GetEEzsigntemplatesignatureAttachmentnamesourceOk() (*FieldEEzsigntemplatesignatureAttachmentnamesource, bool) {
	if o == nil || IsNil(o.EEzsigntemplatesignatureAttachmentnamesource) {
		return nil, false
	}
	return o.EEzsigntemplatesignatureAttachmentnamesource, true
}

// HasEEzsigntemplatesignatureAttachmentnamesource returns a boolean if a field has been set.
func (o *EzsigntemplatesignatureResponse) HasEEzsigntemplatesignatureAttachmentnamesource() bool {
	if o != nil && !IsNil(o.EEzsigntemplatesignatureAttachmentnamesource) {
		return true
	}

	return false
}

// SetEEzsigntemplatesignatureAttachmentnamesource gets a reference to the given FieldEEzsigntemplatesignatureAttachmentnamesource and assigns it to the EEzsigntemplatesignatureAttachmentnamesource field.
func (o *EzsigntemplatesignatureResponse) SetEEzsigntemplatesignatureAttachmentnamesource(v FieldEEzsigntemplatesignatureAttachmentnamesource) {
	o.EEzsigntemplatesignatureAttachmentnamesource = &v
}

// GetBEzsigntemplatesignatureRequired returns the BEzsigntemplatesignatureRequired field value if set, zero value otherwise.
func (o *EzsigntemplatesignatureResponse) GetBEzsigntemplatesignatureRequired() bool {
	if o == nil || IsNil(o.BEzsigntemplatesignatureRequired) {
		var ret bool
		return ret
	}
	return *o.BEzsigntemplatesignatureRequired
}

// GetBEzsigntemplatesignatureRequiredOk returns a tuple with the BEzsigntemplatesignatureRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignatureResponse) GetBEzsigntemplatesignatureRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.BEzsigntemplatesignatureRequired) {
		return nil, false
	}
	return o.BEzsigntemplatesignatureRequired, true
}

// HasBEzsigntemplatesignatureRequired returns a boolean if a field has been set.
func (o *EzsigntemplatesignatureResponse) HasBEzsigntemplatesignatureRequired() bool {
	if o != nil && !IsNil(o.BEzsigntemplatesignatureRequired) {
		return true
	}

	return false
}

// SetBEzsigntemplatesignatureRequired gets a reference to the given bool and assigns it to the BEzsigntemplatesignatureRequired field.
func (o *EzsigntemplatesignatureResponse) SetBEzsigntemplatesignatureRequired(v bool) {
	o.BEzsigntemplatesignatureRequired = &v
}

// GetIEzsigntemplatesignatureMaxlength returns the IEzsigntemplatesignatureMaxlength field value if set, zero value otherwise.
func (o *EzsigntemplatesignatureResponse) GetIEzsigntemplatesignatureMaxlength() int32 {
	if o == nil || IsNil(o.IEzsigntemplatesignatureMaxlength) {
		var ret int32
		return ret
	}
	return *o.IEzsigntemplatesignatureMaxlength
}

// GetIEzsigntemplatesignatureMaxlengthOk returns a tuple with the IEzsigntemplatesignatureMaxlength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignatureResponse) GetIEzsigntemplatesignatureMaxlengthOk() (*int32, bool) {
	if o == nil || IsNil(o.IEzsigntemplatesignatureMaxlength) {
		return nil, false
	}
	return o.IEzsigntemplatesignatureMaxlength, true
}

// HasIEzsigntemplatesignatureMaxlength returns a boolean if a field has been set.
func (o *EzsigntemplatesignatureResponse) HasIEzsigntemplatesignatureMaxlength() bool {
	if o != nil && !IsNil(o.IEzsigntemplatesignatureMaxlength) {
		return true
	}

	return false
}

// SetIEzsigntemplatesignatureMaxlength gets a reference to the given int32 and assigns it to the IEzsigntemplatesignatureMaxlength field.
func (o *EzsigntemplatesignatureResponse) SetIEzsigntemplatesignatureMaxlength(v int32) {
	o.IEzsigntemplatesignatureMaxlength = &v
}

// GetSEzsigntemplatesignatureRegexp returns the SEzsigntemplatesignatureRegexp field value if set, zero value otherwise.
func (o *EzsigntemplatesignatureResponse) GetSEzsigntemplatesignatureRegexp() string {
	if o == nil || IsNil(o.SEzsigntemplatesignatureRegexp) {
		var ret string
		return ret
	}
	return *o.SEzsigntemplatesignatureRegexp
}

// GetSEzsigntemplatesignatureRegexpOk returns a tuple with the SEzsigntemplatesignatureRegexp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignatureResponse) GetSEzsigntemplatesignatureRegexpOk() (*string, bool) {
	if o == nil || IsNil(o.SEzsigntemplatesignatureRegexp) {
		return nil, false
	}
	return o.SEzsigntemplatesignatureRegexp, true
}

// HasSEzsigntemplatesignatureRegexp returns a boolean if a field has been set.
func (o *EzsigntemplatesignatureResponse) HasSEzsigntemplatesignatureRegexp() bool {
	if o != nil && !IsNil(o.SEzsigntemplatesignatureRegexp) {
		return true
	}

	return false
}

// SetSEzsigntemplatesignatureRegexp gets a reference to the given string and assigns it to the SEzsigntemplatesignatureRegexp field.
func (o *EzsigntemplatesignatureResponse) SetSEzsigntemplatesignatureRegexp(v string) {
	o.SEzsigntemplatesignatureRegexp = &v
}

// GetEEzsigntemplatesignatureTextvalidation returns the EEzsigntemplatesignatureTextvalidation field value if set, zero value otherwise.
func (o *EzsigntemplatesignatureResponse) GetEEzsigntemplatesignatureTextvalidation() EnumTextvalidation {
	if o == nil || IsNil(o.EEzsigntemplatesignatureTextvalidation) {
		var ret EnumTextvalidation
		return ret
	}
	return *o.EEzsigntemplatesignatureTextvalidation
}

// GetEEzsigntemplatesignatureTextvalidationOk returns a tuple with the EEzsigntemplatesignatureTextvalidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignatureResponse) GetEEzsigntemplatesignatureTextvalidationOk() (*EnumTextvalidation, bool) {
	if o == nil || IsNil(o.EEzsigntemplatesignatureTextvalidation) {
		return nil, false
	}
	return o.EEzsigntemplatesignatureTextvalidation, true
}

// HasEEzsigntemplatesignatureTextvalidation returns a boolean if a field has been set.
func (o *EzsigntemplatesignatureResponse) HasEEzsigntemplatesignatureTextvalidation() bool {
	if o != nil && !IsNil(o.EEzsigntemplatesignatureTextvalidation) {
		return true
	}

	return false
}

// SetEEzsigntemplatesignatureTextvalidation gets a reference to the given EnumTextvalidation and assigns it to the EEzsigntemplatesignatureTextvalidation field.
func (o *EzsigntemplatesignatureResponse) SetEEzsigntemplatesignatureTextvalidation(v EnumTextvalidation) {
	o.EEzsigntemplatesignatureTextvalidation = &v
}

// GetEEzsigntemplatesignatureDependencyrequirement returns the EEzsigntemplatesignatureDependencyrequirement field value if set, zero value otherwise.
func (o *EzsigntemplatesignatureResponse) GetEEzsigntemplatesignatureDependencyrequirement() FieldEEzsigntemplatesignatureDependencyrequirement {
	if o == nil || IsNil(o.EEzsigntemplatesignatureDependencyrequirement) {
		var ret FieldEEzsigntemplatesignatureDependencyrequirement
		return ret
	}
	return *o.EEzsigntemplatesignatureDependencyrequirement
}

// GetEEzsigntemplatesignatureDependencyrequirementOk returns a tuple with the EEzsigntemplatesignatureDependencyrequirement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignatureResponse) GetEEzsigntemplatesignatureDependencyrequirementOk() (*FieldEEzsigntemplatesignatureDependencyrequirement, bool) {
	if o == nil || IsNil(o.EEzsigntemplatesignatureDependencyrequirement) {
		return nil, false
	}
	return o.EEzsigntemplatesignatureDependencyrequirement, true
}

// HasEEzsigntemplatesignatureDependencyrequirement returns a boolean if a field has been set.
func (o *EzsigntemplatesignatureResponse) HasEEzsigntemplatesignatureDependencyrequirement() bool {
	if o != nil && !IsNil(o.EEzsigntemplatesignatureDependencyrequirement) {
		return true
	}

	return false
}

// SetEEzsigntemplatesignatureDependencyrequirement gets a reference to the given FieldEEzsigntemplatesignatureDependencyrequirement and assigns it to the EEzsigntemplatesignatureDependencyrequirement field.
func (o *EzsigntemplatesignatureResponse) SetEEzsigntemplatesignatureDependencyrequirement(v FieldEEzsigntemplatesignatureDependencyrequirement) {
	o.EEzsigntemplatesignatureDependencyrequirement = &v
}

func (o EzsigntemplatesignatureResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplatesignatureResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiEzsigntemplatesignatureID"] = o.PkiEzsigntemplatesignatureID
	toSerialize["fkiEzsigntemplatedocumentID"] = o.FkiEzsigntemplatedocumentID
	toSerialize["fkiEzsigntemplatesignerID"] = o.FkiEzsigntemplatesignerID
	if !IsNil(o.FkiEzsigntemplatesignerIDValidation) {
		toSerialize["fkiEzsigntemplatesignerIDValidation"] = o.FkiEzsigntemplatesignerIDValidation
	}
	toSerialize["iEzsigntemplatedocumentpagePagenumber"] = o.IEzsigntemplatedocumentpagePagenumber
	toSerialize["iEzsigntemplatesignatureX"] = o.IEzsigntemplatesignatureX
	toSerialize["iEzsigntemplatesignatureY"] = o.IEzsigntemplatesignatureY
	if !IsNil(o.IEzsigntemplatesignatureWidth) {
		toSerialize["iEzsigntemplatesignatureWidth"] = o.IEzsigntemplatesignatureWidth
	}
	if !IsNil(o.IEzsigntemplatesignatureHeight) {
		toSerialize["iEzsigntemplatesignatureHeight"] = o.IEzsigntemplatesignatureHeight
	}
	toSerialize["iEzsigntemplatesignatureStep"] = o.IEzsigntemplatesignatureStep
	toSerialize["eEzsigntemplatesignatureType"] = o.EEzsigntemplatesignatureType
	if !IsNil(o.TEzsigntemplatesignatureTooltip) {
		toSerialize["tEzsigntemplatesignatureTooltip"] = o.TEzsigntemplatesignatureTooltip
	}
	if !IsNil(o.EEzsigntemplatesignatureTooltipposition) {
		toSerialize["eEzsigntemplatesignatureTooltipposition"] = o.EEzsigntemplatesignatureTooltipposition
	}
	if !IsNil(o.EEzsigntemplatesignatureFont) {
		toSerialize["eEzsigntemplatesignatureFont"] = o.EEzsigntemplatesignatureFont
	}
	if !IsNil(o.IEzsigntemplatesignatureValidationstep) {
		toSerialize["iEzsigntemplatesignatureValidationstep"] = o.IEzsigntemplatesignatureValidationstep
	}
	if !IsNil(o.SEzsigntemplatesignatureAttachmentdescription) {
		toSerialize["sEzsigntemplatesignatureAttachmentdescription"] = o.SEzsigntemplatesignatureAttachmentdescription
	}
	if !IsNil(o.EEzsigntemplatesignatureAttachmentnamesource) {
		toSerialize["eEzsigntemplatesignatureAttachmentnamesource"] = o.EEzsigntemplatesignatureAttachmentnamesource
	}
	if !IsNil(o.BEzsigntemplatesignatureRequired) {
		toSerialize["bEzsigntemplatesignatureRequired"] = o.BEzsigntemplatesignatureRequired
	}
	if !IsNil(o.IEzsigntemplatesignatureMaxlength) {
		toSerialize["iEzsigntemplatesignatureMaxlength"] = o.IEzsigntemplatesignatureMaxlength
	}
	if !IsNil(o.SEzsigntemplatesignatureRegexp) {
		toSerialize["sEzsigntemplatesignatureRegexp"] = o.SEzsigntemplatesignatureRegexp
	}
	if !IsNil(o.EEzsigntemplatesignatureTextvalidation) {
		toSerialize["eEzsigntemplatesignatureTextvalidation"] = o.EEzsigntemplatesignatureTextvalidation
	}
	if !IsNil(o.EEzsigntemplatesignatureDependencyrequirement) {
		toSerialize["eEzsigntemplatesignatureDependencyrequirement"] = o.EEzsigntemplatesignatureDependencyrequirement
	}
	return toSerialize, nil
}

type NullableEzsigntemplatesignatureResponse struct {
	value *EzsigntemplatesignatureResponse
	isSet bool
}

func (v NullableEzsigntemplatesignatureResponse) Get() *EzsigntemplatesignatureResponse {
	return v.value
}

func (v *NullableEzsigntemplatesignatureResponse) Set(val *EzsigntemplatesignatureResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplatesignatureResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplatesignatureResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplatesignatureResponse(val *EzsigntemplatesignatureResponse) *NullableEzsigntemplatesignatureResponse {
	return &NullableEzsigntemplatesignatureResponse{value: val, isSet: true}
}

func (v NullableEzsigntemplatesignatureResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplatesignatureResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


