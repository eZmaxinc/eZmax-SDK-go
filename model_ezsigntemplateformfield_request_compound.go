/*
eZmax API Definition (Full)

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.2.2
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the EzsigntemplateformfieldRequestCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplateformfieldRequestCompound{}

// EzsigntemplateformfieldRequestCompound An Ezsigntemplateformfield Object and children to create a complete structure
type EzsigntemplateformfieldRequestCompound struct {
	// The unique ID of the Ezsigntemplateformfield
	PkiEzsigntemplateformfieldID *int32 `json:"pkiEzsigntemplateformfieldID,omitempty"`
	EEzsigntemplateformfieldPositioning *FieldEEzsigntemplateformfieldPositioning `json:"eEzsigntemplateformfieldPositioning,omitempty"`
	// The page number in the Ezsigntemplatedocument
	IEzsigntemplatedocumentpagePagenumber int32 `json:"iEzsigntemplatedocumentpagePagenumber"`
	// The Label for the Ezsigntemplateformfield
	SEzsigntemplateformfieldLabel string `json:"sEzsigntemplateformfieldLabel"`
	// The value for the Ezsigntemplateformfield
	SEzsigntemplateformfieldValue *string `json:"sEzsigntemplateformfieldValue,omitempty"`
	// The X coordinate (Horizontal) where to put the Ezsigntemplateformfield on the Ezsigntemplatepage.  Coordinate is calculated at 100dpi (dot per inch). So for example, if you want to put the Ezsigntemplateformfield 2 inches from the left border of the page, you would use \"200\" for the X coordinate.
	IEzsigntemplateformfieldX *int32 `json:"iEzsigntemplateformfieldX,omitempty"`
	// The Y coordinate (Vertical) where to put the Ezsigntemplateformfield on the Ezsigntemplatepage.  Coordinate is calculated at 100dpi (dot per inch). So for example, if you want to put the Ezsigntemplateformfield 3 inches from the top border of the page, you would use \"300\" for the Y coordinate.
	IEzsigntemplateformfieldY *int32 `json:"iEzsigntemplateformfieldY,omitempty"`
	// The Width of the Ezsigntemplateformfield in pixels calculated at 100 DPI
	IEzsigntemplateformfieldWidth int32 `json:"iEzsigntemplateformfieldWidth"`
	// The Height of the Ezsigntemplateformfield in pixels calculated at 100 DPI 
	IEzsigntemplateformfieldHeight int32 `json:"iEzsigntemplateformfieldHeight"`
	// Whether the Ezsigntemplateformfield allows the use of the autocomplete of the browser.  This can only be set if eEzsigntemplateformfieldgroupType is **Text**
	BEzsigntemplateformfieldAutocomplete *bool `json:"bEzsigntemplateformfieldAutocomplete,omitempty"`
	// Whether the Ezsigntemplateformfield is selected or not by default.  This can only be set if eEzsigntemplateformfieldgroupType is **Checkbox** or **Radio**
	BEzsigntemplateformfieldSelected *bool `json:"bEzsigntemplateformfieldSelected,omitempty"`
	EEzsigntemplateformfieldDependencyrequirement *FieldEEzsigntemplateformfieldDependencyrequirement `json:"eEzsigntemplateformfieldDependencyrequirement,omitempty"`
	// The string pattern to search for the positioning. **This is not a regexp**  This will be required if **eEzsigntemplateformfieldPositioning** is set to **PerCoordinates**
	SEzsigntemplateformfieldPositioningpattern *string `json:"sEzsigntemplateformfieldPositioningpattern,omitempty" validate:"regexp=^.{0,30}$"`
	// The offset X  This will be required if **eEzsigntemplateformfieldPositioning** is set to **PerCoordinates**
	IEzsigntemplateformfieldPositioningoffsetx *int32 `json:"iEzsigntemplateformfieldPositioningoffsetx,omitempty"`
	// The offset Y  This will be required if **eEzsigntemplateformfieldPositioning** is set to **PerCoordinates**
	IEzsigntemplateformfieldPositioningoffsety *int32 `json:"iEzsigntemplateformfieldPositioningoffsety,omitempty"`
	EEzsigntemplateformfieldPositioningoccurence *FieldEEzsigntemplateformfieldPositioningoccurence `json:"eEzsigntemplateformfieldPositioningoccurence,omitempty"`
	EEzsigntemplateformfieldHorizontalalignment *EnumHorizontalalignment `json:"eEzsigntemplateformfieldHorizontalalignment,omitempty"`
	ObjTextstylestatic *TextstylestaticRequestCompound `json:"objTextstylestatic,omitempty"`
	AObjEzsigntemplateelementdependency []EzsigntemplateelementdependencyRequestCompound `json:"a_objEzsigntemplateelementdependency,omitempty"`
}

type _EzsigntemplateformfieldRequestCompound EzsigntemplateformfieldRequestCompound

// NewEzsigntemplateformfieldRequestCompound instantiates a new EzsigntemplateformfieldRequestCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplateformfieldRequestCompound(iEzsigntemplatedocumentpagePagenumber int32, sEzsigntemplateformfieldLabel string, iEzsigntemplateformfieldWidth int32, iEzsigntemplateformfieldHeight int32) *EzsigntemplateformfieldRequestCompound {
	this := EzsigntemplateformfieldRequestCompound{}
	var eEzsigntemplateformfieldPositioning FieldEEzsigntemplateformfieldPositioning = PER_COORDINATES
	this.EEzsigntemplateformfieldPositioning = &eEzsigntemplateformfieldPositioning
	this.IEzsigntemplatedocumentpagePagenumber = iEzsigntemplatedocumentpagePagenumber
	this.SEzsigntemplateformfieldLabel = sEzsigntemplateformfieldLabel
	this.IEzsigntemplateformfieldWidth = iEzsigntemplateformfieldWidth
	this.IEzsigntemplateformfieldHeight = iEzsigntemplateformfieldHeight
	return &this
}

// NewEzsigntemplateformfieldRequestCompoundWithDefaults instantiates a new EzsigntemplateformfieldRequestCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplateformfieldRequestCompoundWithDefaults() *EzsigntemplateformfieldRequestCompound {
	this := EzsigntemplateformfieldRequestCompound{}
	var eEzsigntemplateformfieldPositioning FieldEEzsigntemplateformfieldPositioning = PER_COORDINATES
	this.EEzsigntemplateformfieldPositioning = &eEzsigntemplateformfieldPositioning
	return &this
}

// GetPkiEzsigntemplateformfieldID returns the PkiEzsigntemplateformfieldID field value if set, zero value otherwise.
func (o *EzsigntemplateformfieldRequestCompound) GetPkiEzsigntemplateformfieldID() int32 {
	if o == nil || IsNil(o.PkiEzsigntemplateformfieldID) {
		var ret int32
		return ret
	}
	return *o.PkiEzsigntemplateformfieldID
}

// GetPkiEzsigntemplateformfieldIDOk returns a tuple with the PkiEzsigntemplateformfieldID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldRequestCompound) GetPkiEzsigntemplateformfieldIDOk() (*int32, bool) {
	if o == nil || IsNil(o.PkiEzsigntemplateformfieldID) {
		return nil, false
	}
	return o.PkiEzsigntemplateformfieldID, true
}

// HasPkiEzsigntemplateformfieldID returns a boolean if a field has been set.
func (o *EzsigntemplateformfieldRequestCompound) HasPkiEzsigntemplateformfieldID() bool {
	if o != nil && !IsNil(o.PkiEzsigntemplateformfieldID) {
		return true
	}

	return false
}

// SetPkiEzsigntemplateformfieldID gets a reference to the given int32 and assigns it to the PkiEzsigntemplateformfieldID field.
func (o *EzsigntemplateformfieldRequestCompound) SetPkiEzsigntemplateformfieldID(v int32) {
	o.PkiEzsigntemplateformfieldID = &v
}

// GetEEzsigntemplateformfieldPositioning returns the EEzsigntemplateformfieldPositioning field value if set, zero value otherwise.
func (o *EzsigntemplateformfieldRequestCompound) GetEEzsigntemplateformfieldPositioning() FieldEEzsigntemplateformfieldPositioning {
	if o == nil || IsNil(o.EEzsigntemplateformfieldPositioning) {
		var ret FieldEEzsigntemplateformfieldPositioning
		return ret
	}
	return *o.EEzsigntemplateformfieldPositioning
}

// GetEEzsigntemplateformfieldPositioningOk returns a tuple with the EEzsigntemplateformfieldPositioning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldRequestCompound) GetEEzsigntemplateformfieldPositioningOk() (*FieldEEzsigntemplateformfieldPositioning, bool) {
	if o == nil || IsNil(o.EEzsigntemplateformfieldPositioning) {
		return nil, false
	}
	return o.EEzsigntemplateformfieldPositioning, true
}

// HasEEzsigntemplateformfieldPositioning returns a boolean if a field has been set.
func (o *EzsigntemplateformfieldRequestCompound) HasEEzsigntemplateformfieldPositioning() bool {
	if o != nil && !IsNil(o.EEzsigntemplateformfieldPositioning) {
		return true
	}

	return false
}

// SetEEzsigntemplateformfieldPositioning gets a reference to the given FieldEEzsigntemplateformfieldPositioning and assigns it to the EEzsigntemplateformfieldPositioning field.
func (o *EzsigntemplateformfieldRequestCompound) SetEEzsigntemplateformfieldPositioning(v FieldEEzsigntemplateformfieldPositioning) {
	o.EEzsigntemplateformfieldPositioning = &v
}

// GetIEzsigntemplatedocumentpagePagenumber returns the IEzsigntemplatedocumentpagePagenumber field value
func (o *EzsigntemplateformfieldRequestCompound) GetIEzsigntemplatedocumentpagePagenumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsigntemplatedocumentpagePagenumber
}

// GetIEzsigntemplatedocumentpagePagenumberOk returns a tuple with the IEzsigntemplatedocumentpagePagenumber field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldRequestCompound) GetIEzsigntemplatedocumentpagePagenumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsigntemplatedocumentpagePagenumber, true
}

// SetIEzsigntemplatedocumentpagePagenumber sets field value
func (o *EzsigntemplateformfieldRequestCompound) SetIEzsigntemplatedocumentpagePagenumber(v int32) {
	o.IEzsigntemplatedocumentpagePagenumber = v
}

// GetSEzsigntemplateformfieldLabel returns the SEzsigntemplateformfieldLabel field value
func (o *EzsigntemplateformfieldRequestCompound) GetSEzsigntemplateformfieldLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsigntemplateformfieldLabel
}

// GetSEzsigntemplateformfieldLabelOk returns a tuple with the SEzsigntemplateformfieldLabel field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldRequestCompound) GetSEzsigntemplateformfieldLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsigntemplateformfieldLabel, true
}

// SetSEzsigntemplateformfieldLabel sets field value
func (o *EzsigntemplateformfieldRequestCompound) SetSEzsigntemplateformfieldLabel(v string) {
	o.SEzsigntemplateformfieldLabel = v
}

// GetSEzsigntemplateformfieldValue returns the SEzsigntemplateformfieldValue field value if set, zero value otherwise.
func (o *EzsigntemplateformfieldRequestCompound) GetSEzsigntemplateformfieldValue() string {
	if o == nil || IsNil(o.SEzsigntemplateformfieldValue) {
		var ret string
		return ret
	}
	return *o.SEzsigntemplateformfieldValue
}

// GetSEzsigntemplateformfieldValueOk returns a tuple with the SEzsigntemplateformfieldValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldRequestCompound) GetSEzsigntemplateformfieldValueOk() (*string, bool) {
	if o == nil || IsNil(o.SEzsigntemplateformfieldValue) {
		return nil, false
	}
	return o.SEzsigntemplateformfieldValue, true
}

// HasSEzsigntemplateformfieldValue returns a boolean if a field has been set.
func (o *EzsigntemplateformfieldRequestCompound) HasSEzsigntemplateformfieldValue() bool {
	if o != nil && !IsNil(o.SEzsigntemplateformfieldValue) {
		return true
	}

	return false
}

// SetSEzsigntemplateformfieldValue gets a reference to the given string and assigns it to the SEzsigntemplateformfieldValue field.
func (o *EzsigntemplateformfieldRequestCompound) SetSEzsigntemplateformfieldValue(v string) {
	o.SEzsigntemplateformfieldValue = &v
}

// GetIEzsigntemplateformfieldX returns the IEzsigntemplateformfieldX field value if set, zero value otherwise.
func (o *EzsigntemplateformfieldRequestCompound) GetIEzsigntemplateformfieldX() int32 {
	if o == nil || IsNil(o.IEzsigntemplateformfieldX) {
		var ret int32
		return ret
	}
	return *o.IEzsigntemplateformfieldX
}

// GetIEzsigntemplateformfieldXOk returns a tuple with the IEzsigntemplateformfieldX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldRequestCompound) GetIEzsigntemplateformfieldXOk() (*int32, bool) {
	if o == nil || IsNil(o.IEzsigntemplateformfieldX) {
		return nil, false
	}
	return o.IEzsigntemplateformfieldX, true
}

// HasIEzsigntemplateformfieldX returns a boolean if a field has been set.
func (o *EzsigntemplateformfieldRequestCompound) HasIEzsigntemplateformfieldX() bool {
	if o != nil && !IsNil(o.IEzsigntemplateformfieldX) {
		return true
	}

	return false
}

// SetIEzsigntemplateformfieldX gets a reference to the given int32 and assigns it to the IEzsigntemplateformfieldX field.
func (o *EzsigntemplateformfieldRequestCompound) SetIEzsigntemplateformfieldX(v int32) {
	o.IEzsigntemplateformfieldX = &v
}

// GetIEzsigntemplateformfieldY returns the IEzsigntemplateformfieldY field value if set, zero value otherwise.
func (o *EzsigntemplateformfieldRequestCompound) GetIEzsigntemplateformfieldY() int32 {
	if o == nil || IsNil(o.IEzsigntemplateformfieldY) {
		var ret int32
		return ret
	}
	return *o.IEzsigntemplateformfieldY
}

// GetIEzsigntemplateformfieldYOk returns a tuple with the IEzsigntemplateformfieldY field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldRequestCompound) GetIEzsigntemplateformfieldYOk() (*int32, bool) {
	if o == nil || IsNil(o.IEzsigntemplateformfieldY) {
		return nil, false
	}
	return o.IEzsigntemplateformfieldY, true
}

// HasIEzsigntemplateformfieldY returns a boolean if a field has been set.
func (o *EzsigntemplateformfieldRequestCompound) HasIEzsigntemplateformfieldY() bool {
	if o != nil && !IsNil(o.IEzsigntemplateformfieldY) {
		return true
	}

	return false
}

// SetIEzsigntemplateformfieldY gets a reference to the given int32 and assigns it to the IEzsigntemplateformfieldY field.
func (o *EzsigntemplateformfieldRequestCompound) SetIEzsigntemplateformfieldY(v int32) {
	o.IEzsigntemplateformfieldY = &v
}

// GetIEzsigntemplateformfieldWidth returns the IEzsigntemplateformfieldWidth field value
func (o *EzsigntemplateformfieldRequestCompound) GetIEzsigntemplateformfieldWidth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsigntemplateformfieldWidth
}

// GetIEzsigntemplateformfieldWidthOk returns a tuple with the IEzsigntemplateformfieldWidth field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldRequestCompound) GetIEzsigntemplateformfieldWidthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsigntemplateformfieldWidth, true
}

// SetIEzsigntemplateformfieldWidth sets field value
func (o *EzsigntemplateformfieldRequestCompound) SetIEzsigntemplateformfieldWidth(v int32) {
	o.IEzsigntemplateformfieldWidth = v
}

// GetIEzsigntemplateformfieldHeight returns the IEzsigntemplateformfieldHeight field value
func (o *EzsigntemplateformfieldRequestCompound) GetIEzsigntemplateformfieldHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsigntemplateformfieldHeight
}

// GetIEzsigntemplateformfieldHeightOk returns a tuple with the IEzsigntemplateformfieldHeight field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldRequestCompound) GetIEzsigntemplateformfieldHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsigntemplateformfieldHeight, true
}

// SetIEzsigntemplateformfieldHeight sets field value
func (o *EzsigntemplateformfieldRequestCompound) SetIEzsigntemplateformfieldHeight(v int32) {
	o.IEzsigntemplateformfieldHeight = v
}

// GetBEzsigntemplateformfieldAutocomplete returns the BEzsigntemplateformfieldAutocomplete field value if set, zero value otherwise.
func (o *EzsigntemplateformfieldRequestCompound) GetBEzsigntemplateformfieldAutocomplete() bool {
	if o == nil || IsNil(o.BEzsigntemplateformfieldAutocomplete) {
		var ret bool
		return ret
	}
	return *o.BEzsigntemplateformfieldAutocomplete
}

// GetBEzsigntemplateformfieldAutocompleteOk returns a tuple with the BEzsigntemplateformfieldAutocomplete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldRequestCompound) GetBEzsigntemplateformfieldAutocompleteOk() (*bool, bool) {
	if o == nil || IsNil(o.BEzsigntemplateformfieldAutocomplete) {
		return nil, false
	}
	return o.BEzsigntemplateformfieldAutocomplete, true
}

// HasBEzsigntemplateformfieldAutocomplete returns a boolean if a field has been set.
func (o *EzsigntemplateformfieldRequestCompound) HasBEzsigntemplateformfieldAutocomplete() bool {
	if o != nil && !IsNil(o.BEzsigntemplateformfieldAutocomplete) {
		return true
	}

	return false
}

// SetBEzsigntemplateformfieldAutocomplete gets a reference to the given bool and assigns it to the BEzsigntemplateformfieldAutocomplete field.
func (o *EzsigntemplateformfieldRequestCompound) SetBEzsigntemplateformfieldAutocomplete(v bool) {
	o.BEzsigntemplateformfieldAutocomplete = &v
}

// GetBEzsigntemplateformfieldSelected returns the BEzsigntemplateformfieldSelected field value if set, zero value otherwise.
func (o *EzsigntemplateformfieldRequestCompound) GetBEzsigntemplateformfieldSelected() bool {
	if o == nil || IsNil(o.BEzsigntemplateformfieldSelected) {
		var ret bool
		return ret
	}
	return *o.BEzsigntemplateformfieldSelected
}

// GetBEzsigntemplateformfieldSelectedOk returns a tuple with the BEzsigntemplateformfieldSelected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldRequestCompound) GetBEzsigntemplateformfieldSelectedOk() (*bool, bool) {
	if o == nil || IsNil(o.BEzsigntemplateformfieldSelected) {
		return nil, false
	}
	return o.BEzsigntemplateformfieldSelected, true
}

// HasBEzsigntemplateformfieldSelected returns a boolean if a field has been set.
func (o *EzsigntemplateformfieldRequestCompound) HasBEzsigntemplateformfieldSelected() bool {
	if o != nil && !IsNil(o.BEzsigntemplateformfieldSelected) {
		return true
	}

	return false
}

// SetBEzsigntemplateformfieldSelected gets a reference to the given bool and assigns it to the BEzsigntemplateformfieldSelected field.
func (o *EzsigntemplateformfieldRequestCompound) SetBEzsigntemplateformfieldSelected(v bool) {
	o.BEzsigntemplateformfieldSelected = &v
}

// GetEEzsigntemplateformfieldDependencyrequirement returns the EEzsigntemplateformfieldDependencyrequirement field value if set, zero value otherwise.
func (o *EzsigntemplateformfieldRequestCompound) GetEEzsigntemplateformfieldDependencyrequirement() FieldEEzsigntemplateformfieldDependencyrequirement {
	if o == nil || IsNil(o.EEzsigntemplateformfieldDependencyrequirement) {
		var ret FieldEEzsigntemplateformfieldDependencyrequirement
		return ret
	}
	return *o.EEzsigntemplateformfieldDependencyrequirement
}

// GetEEzsigntemplateformfieldDependencyrequirementOk returns a tuple with the EEzsigntemplateformfieldDependencyrequirement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldRequestCompound) GetEEzsigntemplateformfieldDependencyrequirementOk() (*FieldEEzsigntemplateformfieldDependencyrequirement, bool) {
	if o == nil || IsNil(o.EEzsigntemplateformfieldDependencyrequirement) {
		return nil, false
	}
	return o.EEzsigntemplateformfieldDependencyrequirement, true
}

// HasEEzsigntemplateformfieldDependencyrequirement returns a boolean if a field has been set.
func (o *EzsigntemplateformfieldRequestCompound) HasEEzsigntemplateformfieldDependencyrequirement() bool {
	if o != nil && !IsNil(o.EEzsigntemplateformfieldDependencyrequirement) {
		return true
	}

	return false
}

// SetEEzsigntemplateformfieldDependencyrequirement gets a reference to the given FieldEEzsigntemplateformfieldDependencyrequirement and assigns it to the EEzsigntemplateformfieldDependencyrequirement field.
func (o *EzsigntemplateformfieldRequestCompound) SetEEzsigntemplateformfieldDependencyrequirement(v FieldEEzsigntemplateformfieldDependencyrequirement) {
	o.EEzsigntemplateformfieldDependencyrequirement = &v
}

// GetSEzsigntemplateformfieldPositioningpattern returns the SEzsigntemplateformfieldPositioningpattern field value if set, zero value otherwise.
func (o *EzsigntemplateformfieldRequestCompound) GetSEzsigntemplateformfieldPositioningpattern() string {
	if o == nil || IsNil(o.SEzsigntemplateformfieldPositioningpattern) {
		var ret string
		return ret
	}
	return *o.SEzsigntemplateformfieldPositioningpattern
}

// GetSEzsigntemplateformfieldPositioningpatternOk returns a tuple with the SEzsigntemplateformfieldPositioningpattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldRequestCompound) GetSEzsigntemplateformfieldPositioningpatternOk() (*string, bool) {
	if o == nil || IsNil(o.SEzsigntemplateformfieldPositioningpattern) {
		return nil, false
	}
	return o.SEzsigntemplateformfieldPositioningpattern, true
}

// HasSEzsigntemplateformfieldPositioningpattern returns a boolean if a field has been set.
func (o *EzsigntemplateformfieldRequestCompound) HasSEzsigntemplateformfieldPositioningpattern() bool {
	if o != nil && !IsNil(o.SEzsigntemplateformfieldPositioningpattern) {
		return true
	}

	return false
}

// SetSEzsigntemplateformfieldPositioningpattern gets a reference to the given string and assigns it to the SEzsigntemplateformfieldPositioningpattern field.
func (o *EzsigntemplateformfieldRequestCompound) SetSEzsigntemplateformfieldPositioningpattern(v string) {
	o.SEzsigntemplateformfieldPositioningpattern = &v
}

// GetIEzsigntemplateformfieldPositioningoffsetx returns the IEzsigntemplateformfieldPositioningoffsetx field value if set, zero value otherwise.
func (o *EzsigntemplateformfieldRequestCompound) GetIEzsigntemplateformfieldPositioningoffsetx() int32 {
	if o == nil || IsNil(o.IEzsigntemplateformfieldPositioningoffsetx) {
		var ret int32
		return ret
	}
	return *o.IEzsigntemplateformfieldPositioningoffsetx
}

// GetIEzsigntemplateformfieldPositioningoffsetxOk returns a tuple with the IEzsigntemplateformfieldPositioningoffsetx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldRequestCompound) GetIEzsigntemplateformfieldPositioningoffsetxOk() (*int32, bool) {
	if o == nil || IsNil(o.IEzsigntemplateformfieldPositioningoffsetx) {
		return nil, false
	}
	return o.IEzsigntemplateformfieldPositioningoffsetx, true
}

// HasIEzsigntemplateformfieldPositioningoffsetx returns a boolean if a field has been set.
func (o *EzsigntemplateformfieldRequestCompound) HasIEzsigntemplateformfieldPositioningoffsetx() bool {
	if o != nil && !IsNil(o.IEzsigntemplateformfieldPositioningoffsetx) {
		return true
	}

	return false
}

// SetIEzsigntemplateformfieldPositioningoffsetx gets a reference to the given int32 and assigns it to the IEzsigntemplateformfieldPositioningoffsetx field.
func (o *EzsigntemplateformfieldRequestCompound) SetIEzsigntemplateformfieldPositioningoffsetx(v int32) {
	o.IEzsigntemplateformfieldPositioningoffsetx = &v
}

// GetIEzsigntemplateformfieldPositioningoffsety returns the IEzsigntemplateformfieldPositioningoffsety field value if set, zero value otherwise.
func (o *EzsigntemplateformfieldRequestCompound) GetIEzsigntemplateformfieldPositioningoffsety() int32 {
	if o == nil || IsNil(o.IEzsigntemplateformfieldPositioningoffsety) {
		var ret int32
		return ret
	}
	return *o.IEzsigntemplateformfieldPositioningoffsety
}

// GetIEzsigntemplateformfieldPositioningoffsetyOk returns a tuple with the IEzsigntemplateformfieldPositioningoffsety field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldRequestCompound) GetIEzsigntemplateformfieldPositioningoffsetyOk() (*int32, bool) {
	if o == nil || IsNil(o.IEzsigntemplateformfieldPositioningoffsety) {
		return nil, false
	}
	return o.IEzsigntemplateformfieldPositioningoffsety, true
}

// HasIEzsigntemplateformfieldPositioningoffsety returns a boolean if a field has been set.
func (o *EzsigntemplateformfieldRequestCompound) HasIEzsigntemplateformfieldPositioningoffsety() bool {
	if o != nil && !IsNil(o.IEzsigntemplateformfieldPositioningoffsety) {
		return true
	}

	return false
}

// SetIEzsigntemplateformfieldPositioningoffsety gets a reference to the given int32 and assigns it to the IEzsigntemplateformfieldPositioningoffsety field.
func (o *EzsigntemplateformfieldRequestCompound) SetIEzsigntemplateformfieldPositioningoffsety(v int32) {
	o.IEzsigntemplateformfieldPositioningoffsety = &v
}

// GetEEzsigntemplateformfieldPositioningoccurence returns the EEzsigntemplateformfieldPositioningoccurence field value if set, zero value otherwise.
func (o *EzsigntemplateformfieldRequestCompound) GetEEzsigntemplateformfieldPositioningoccurence() FieldEEzsigntemplateformfieldPositioningoccurence {
	if o == nil || IsNil(o.EEzsigntemplateformfieldPositioningoccurence) {
		var ret FieldEEzsigntemplateformfieldPositioningoccurence
		return ret
	}
	return *o.EEzsigntemplateformfieldPositioningoccurence
}

// GetEEzsigntemplateformfieldPositioningoccurenceOk returns a tuple with the EEzsigntemplateformfieldPositioningoccurence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldRequestCompound) GetEEzsigntemplateformfieldPositioningoccurenceOk() (*FieldEEzsigntemplateformfieldPositioningoccurence, bool) {
	if o == nil || IsNil(o.EEzsigntemplateformfieldPositioningoccurence) {
		return nil, false
	}
	return o.EEzsigntemplateformfieldPositioningoccurence, true
}

// HasEEzsigntemplateformfieldPositioningoccurence returns a boolean if a field has been set.
func (o *EzsigntemplateformfieldRequestCompound) HasEEzsigntemplateformfieldPositioningoccurence() bool {
	if o != nil && !IsNil(o.EEzsigntemplateformfieldPositioningoccurence) {
		return true
	}

	return false
}

// SetEEzsigntemplateformfieldPositioningoccurence gets a reference to the given FieldEEzsigntemplateformfieldPositioningoccurence and assigns it to the EEzsigntemplateformfieldPositioningoccurence field.
func (o *EzsigntemplateformfieldRequestCompound) SetEEzsigntemplateformfieldPositioningoccurence(v FieldEEzsigntemplateformfieldPositioningoccurence) {
	o.EEzsigntemplateformfieldPositioningoccurence = &v
}

// GetEEzsigntemplateformfieldHorizontalalignment returns the EEzsigntemplateformfieldHorizontalalignment field value if set, zero value otherwise.
func (o *EzsigntemplateformfieldRequestCompound) GetEEzsigntemplateformfieldHorizontalalignment() EnumHorizontalalignment {
	if o == nil || IsNil(o.EEzsigntemplateformfieldHorizontalalignment) {
		var ret EnumHorizontalalignment
		return ret
	}
	return *o.EEzsigntemplateformfieldHorizontalalignment
}

// GetEEzsigntemplateformfieldHorizontalalignmentOk returns a tuple with the EEzsigntemplateformfieldHorizontalalignment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldRequestCompound) GetEEzsigntemplateformfieldHorizontalalignmentOk() (*EnumHorizontalalignment, bool) {
	if o == nil || IsNil(o.EEzsigntemplateformfieldHorizontalalignment) {
		return nil, false
	}
	return o.EEzsigntemplateformfieldHorizontalalignment, true
}

// HasEEzsigntemplateformfieldHorizontalalignment returns a boolean if a field has been set.
func (o *EzsigntemplateformfieldRequestCompound) HasEEzsigntemplateformfieldHorizontalalignment() bool {
	if o != nil && !IsNil(o.EEzsigntemplateformfieldHorizontalalignment) {
		return true
	}

	return false
}

// SetEEzsigntemplateformfieldHorizontalalignment gets a reference to the given EnumHorizontalalignment and assigns it to the EEzsigntemplateformfieldHorizontalalignment field.
func (o *EzsigntemplateformfieldRequestCompound) SetEEzsigntemplateformfieldHorizontalalignment(v EnumHorizontalalignment) {
	o.EEzsigntemplateformfieldHorizontalalignment = &v
}

// GetObjTextstylestatic returns the ObjTextstylestatic field value if set, zero value otherwise.
func (o *EzsigntemplateformfieldRequestCompound) GetObjTextstylestatic() TextstylestaticRequestCompound {
	if o == nil || IsNil(o.ObjTextstylestatic) {
		var ret TextstylestaticRequestCompound
		return ret
	}
	return *o.ObjTextstylestatic
}

// GetObjTextstylestaticOk returns a tuple with the ObjTextstylestatic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldRequestCompound) GetObjTextstylestaticOk() (*TextstylestaticRequestCompound, bool) {
	if o == nil || IsNil(o.ObjTextstylestatic) {
		return nil, false
	}
	return o.ObjTextstylestatic, true
}

// HasObjTextstylestatic returns a boolean if a field has been set.
func (o *EzsigntemplateformfieldRequestCompound) HasObjTextstylestatic() bool {
	if o != nil && !IsNil(o.ObjTextstylestatic) {
		return true
	}

	return false
}

// SetObjTextstylestatic gets a reference to the given TextstylestaticRequestCompound and assigns it to the ObjTextstylestatic field.
func (o *EzsigntemplateformfieldRequestCompound) SetObjTextstylestatic(v TextstylestaticRequestCompound) {
	o.ObjTextstylestatic = &v
}

// GetAObjEzsigntemplateelementdependency returns the AObjEzsigntemplateelementdependency field value if set, zero value otherwise.
func (o *EzsigntemplateformfieldRequestCompound) GetAObjEzsigntemplateelementdependency() []EzsigntemplateelementdependencyRequestCompound {
	if o == nil || IsNil(o.AObjEzsigntemplateelementdependency) {
		var ret []EzsigntemplateelementdependencyRequestCompound
		return ret
	}
	return o.AObjEzsigntemplateelementdependency
}

// GetAObjEzsigntemplateelementdependencyOk returns a tuple with the AObjEzsigntemplateelementdependency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldRequestCompound) GetAObjEzsigntemplateelementdependencyOk() ([]EzsigntemplateelementdependencyRequestCompound, bool) {
	if o == nil || IsNil(o.AObjEzsigntemplateelementdependency) {
		return nil, false
	}
	return o.AObjEzsigntemplateelementdependency, true
}

// HasAObjEzsigntemplateelementdependency returns a boolean if a field has been set.
func (o *EzsigntemplateformfieldRequestCompound) HasAObjEzsigntemplateelementdependency() bool {
	if o != nil && !IsNil(o.AObjEzsigntemplateelementdependency) {
		return true
	}

	return false
}

// SetAObjEzsigntemplateelementdependency gets a reference to the given []EzsigntemplateelementdependencyRequestCompound and assigns it to the AObjEzsigntemplateelementdependency field.
func (o *EzsigntemplateformfieldRequestCompound) SetAObjEzsigntemplateelementdependency(v []EzsigntemplateelementdependencyRequestCompound) {
	o.AObjEzsigntemplateelementdependency = v
}

func (o EzsigntemplateformfieldRequestCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplateformfieldRequestCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PkiEzsigntemplateformfieldID) {
		toSerialize["pkiEzsigntemplateformfieldID"] = o.PkiEzsigntemplateformfieldID
	}
	if !IsNil(o.EEzsigntemplateformfieldPositioning) {
		toSerialize["eEzsigntemplateformfieldPositioning"] = o.EEzsigntemplateformfieldPositioning
	}
	toSerialize["iEzsigntemplatedocumentpagePagenumber"] = o.IEzsigntemplatedocumentpagePagenumber
	toSerialize["sEzsigntemplateformfieldLabel"] = o.SEzsigntemplateformfieldLabel
	if !IsNil(o.SEzsigntemplateformfieldValue) {
		toSerialize["sEzsigntemplateformfieldValue"] = o.SEzsigntemplateformfieldValue
	}
	if !IsNil(o.IEzsigntemplateformfieldX) {
		toSerialize["iEzsigntemplateformfieldX"] = o.IEzsigntemplateformfieldX
	}
	if !IsNil(o.IEzsigntemplateformfieldY) {
		toSerialize["iEzsigntemplateformfieldY"] = o.IEzsigntemplateformfieldY
	}
	toSerialize["iEzsigntemplateformfieldWidth"] = o.IEzsigntemplateformfieldWidth
	toSerialize["iEzsigntemplateformfieldHeight"] = o.IEzsigntemplateformfieldHeight
	if !IsNil(o.BEzsigntemplateformfieldAutocomplete) {
		toSerialize["bEzsigntemplateformfieldAutocomplete"] = o.BEzsigntemplateformfieldAutocomplete
	}
	if !IsNil(o.BEzsigntemplateformfieldSelected) {
		toSerialize["bEzsigntemplateformfieldSelected"] = o.BEzsigntemplateformfieldSelected
	}
	if !IsNil(o.EEzsigntemplateformfieldDependencyrequirement) {
		toSerialize["eEzsigntemplateformfieldDependencyrequirement"] = o.EEzsigntemplateformfieldDependencyrequirement
	}
	if !IsNil(o.SEzsigntemplateformfieldPositioningpattern) {
		toSerialize["sEzsigntemplateformfieldPositioningpattern"] = o.SEzsigntemplateformfieldPositioningpattern
	}
	if !IsNil(o.IEzsigntemplateformfieldPositioningoffsetx) {
		toSerialize["iEzsigntemplateformfieldPositioningoffsetx"] = o.IEzsigntemplateformfieldPositioningoffsetx
	}
	if !IsNil(o.IEzsigntemplateformfieldPositioningoffsety) {
		toSerialize["iEzsigntemplateformfieldPositioningoffsety"] = o.IEzsigntemplateformfieldPositioningoffsety
	}
	if !IsNil(o.EEzsigntemplateformfieldPositioningoccurence) {
		toSerialize["eEzsigntemplateformfieldPositioningoccurence"] = o.EEzsigntemplateformfieldPositioningoccurence
	}
	if !IsNil(o.EEzsigntemplateformfieldHorizontalalignment) {
		toSerialize["eEzsigntemplateformfieldHorizontalalignment"] = o.EEzsigntemplateformfieldHorizontalalignment
	}
	if !IsNil(o.ObjTextstylestatic) {
		toSerialize["objTextstylestatic"] = o.ObjTextstylestatic
	}
	if !IsNil(o.AObjEzsigntemplateelementdependency) {
		toSerialize["a_objEzsigntemplateelementdependency"] = o.AObjEzsigntemplateelementdependency
	}
	return toSerialize, nil
}

func (o *EzsigntemplateformfieldRequestCompound) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"iEzsigntemplatedocumentpagePagenumber",
		"sEzsigntemplateformfieldLabel",
		"iEzsigntemplateformfieldWidth",
		"iEzsigntemplateformfieldHeight",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varEzsigntemplateformfieldRequestCompound := _EzsigntemplateformfieldRequestCompound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigntemplateformfieldRequestCompound)

	if err != nil {
		return err
	}

	*o = EzsigntemplateformfieldRequestCompound(varEzsigntemplateformfieldRequestCompound)

	return err
}

type NullableEzsigntemplateformfieldRequestCompound struct {
	value *EzsigntemplateformfieldRequestCompound
	isSet bool
}

func (v NullableEzsigntemplateformfieldRequestCompound) Get() *EzsigntemplateformfieldRequestCompound {
	return v.value
}

func (v *NullableEzsigntemplateformfieldRequestCompound) Set(val *EzsigntemplateformfieldRequestCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplateformfieldRequestCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplateformfieldRequestCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplateformfieldRequestCompound(val *EzsigntemplateformfieldRequestCompound) *NullableEzsigntemplateformfieldRequestCompound {
	return &NullableEzsigntemplateformfieldRequestCompound{value: val, isSet: true}
}

func (v NullableEzsigntemplateformfieldRequestCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplateformfieldRequestCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


