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

// checks if the BrandingResponseCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BrandingResponseCompound{}

// BrandingResponseCompound A Branding Object
type BrandingResponseCompound struct {
	// The unique ID of the Branding
	PkiBrandingID int32 `json:"pkiBrandingID"`
	// The unique ID of the Email
	FkiEmailID *int32 `json:"fkiEmailID,omitempty"`
	ObjBrandingDescription MultilingualBrandingDescription `json:"objBrandingDescription"`
	// The Description of the Branding in the language of the requester
	SBrandingDescriptionX string `json:"sBrandingDescriptionX"`
	// The name of the Branding  This value will only be set if you wish to overwrite the default name. If you want to keep the default name, leave this property empty
	SBrandingName *string `json:"sBrandingName,omitempty"`
	// The email address.
	SEmailAddress *string `json:"sEmailAddress,omitempty"`
	EBrandingLogo FieldEBrandingLogo `json:"eBrandingLogo"`
	// The color of the text. This is a RGB color converted into integer
	IBrandingColortext int32 `json:"iBrandingColortext"`
	// The color of the text in the link box. This is a RGB color converted into integer
	IBrandingColortextlinkbox int32 `json:"iBrandingColortextlinkbox"`
	// The color of the text in the button. This is a RGB color converted into integer
	IBrandingColortextbutton int32 `json:"iBrandingColortextbutton"`
	// The color of the background. This is a RGB color converted into integer
	IBrandingColorbackground int32 `json:"iBrandingColorbackground"`
	// The color of the background of the button. This is a RGB color converted into integer
	IBrandingColorbackgroundbutton int32 `json:"iBrandingColorbackgroundbutton"`
	// The color of the background of the small box. This is a RGB color converted into integer
	IBrandingColorbackgroundsmallbox int32 `json:"iBrandingColorbackgroundsmallbox"`
	// Whether the Branding is active or not
	BBrandingIsactive bool `json:"bBrandingIsactive"`
	// The url of the picture used as logo in the Branding
	SBrandingLogourl *string `json:"sBrandingLogourl,omitempty"`
}

// NewBrandingResponseCompound instantiates a new BrandingResponseCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrandingResponseCompound(pkiBrandingID int32, objBrandingDescription MultilingualBrandingDescription, sBrandingDescriptionX string, eBrandingLogo FieldEBrandingLogo, iBrandingColortext int32, iBrandingColortextlinkbox int32, iBrandingColortextbutton int32, iBrandingColorbackground int32, iBrandingColorbackgroundbutton int32, iBrandingColorbackgroundsmallbox int32, bBrandingIsactive bool) *BrandingResponseCompound {
	this := BrandingResponseCompound{}
	this.PkiBrandingID = pkiBrandingID
	this.ObjBrandingDescription = objBrandingDescription
	this.SBrandingDescriptionX = sBrandingDescriptionX
	this.EBrandingLogo = eBrandingLogo
	this.IBrandingColortext = iBrandingColortext
	this.IBrandingColortextlinkbox = iBrandingColortextlinkbox
	this.IBrandingColortextbutton = iBrandingColortextbutton
	this.IBrandingColorbackground = iBrandingColorbackground
	this.IBrandingColorbackgroundbutton = iBrandingColorbackgroundbutton
	this.IBrandingColorbackgroundsmallbox = iBrandingColorbackgroundsmallbox
	this.BBrandingIsactive = bBrandingIsactive
	return &this
}

// NewBrandingResponseCompoundWithDefaults instantiates a new BrandingResponseCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrandingResponseCompoundWithDefaults() *BrandingResponseCompound {
	this := BrandingResponseCompound{}
	return &this
}

// GetPkiBrandingID returns the PkiBrandingID field value
func (o *BrandingResponseCompound) GetPkiBrandingID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiBrandingID
}

// GetPkiBrandingIDOk returns a tuple with the PkiBrandingID field value
// and a boolean to check if the value has been set.
func (o *BrandingResponseCompound) GetPkiBrandingIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiBrandingID, true
}

// SetPkiBrandingID sets field value
func (o *BrandingResponseCompound) SetPkiBrandingID(v int32) {
	o.PkiBrandingID = v
}

// GetFkiEmailID returns the FkiEmailID field value if set, zero value otherwise.
func (o *BrandingResponseCompound) GetFkiEmailID() int32 {
	if o == nil || IsNil(o.FkiEmailID) {
		var ret int32
		return ret
	}
	return *o.FkiEmailID
}

// GetFkiEmailIDOk returns a tuple with the FkiEmailID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandingResponseCompound) GetFkiEmailIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiEmailID) {
		return nil, false
	}
	return o.FkiEmailID, true
}

// HasFkiEmailID returns a boolean if a field has been set.
func (o *BrandingResponseCompound) HasFkiEmailID() bool {
	if o != nil && !IsNil(o.FkiEmailID) {
		return true
	}

	return false
}

// SetFkiEmailID gets a reference to the given int32 and assigns it to the FkiEmailID field.
func (o *BrandingResponseCompound) SetFkiEmailID(v int32) {
	o.FkiEmailID = &v
}

// GetObjBrandingDescription returns the ObjBrandingDescription field value
func (o *BrandingResponseCompound) GetObjBrandingDescription() MultilingualBrandingDescription {
	if o == nil {
		var ret MultilingualBrandingDescription
		return ret
	}

	return o.ObjBrandingDescription
}

// GetObjBrandingDescriptionOk returns a tuple with the ObjBrandingDescription field value
// and a boolean to check if the value has been set.
func (o *BrandingResponseCompound) GetObjBrandingDescriptionOk() (*MultilingualBrandingDescription, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjBrandingDescription, true
}

// SetObjBrandingDescription sets field value
func (o *BrandingResponseCompound) SetObjBrandingDescription(v MultilingualBrandingDescription) {
	o.ObjBrandingDescription = v
}

// GetSBrandingDescriptionX returns the SBrandingDescriptionX field value
func (o *BrandingResponseCompound) GetSBrandingDescriptionX() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SBrandingDescriptionX
}

// GetSBrandingDescriptionXOk returns a tuple with the SBrandingDescriptionX field value
// and a boolean to check if the value has been set.
func (o *BrandingResponseCompound) GetSBrandingDescriptionXOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SBrandingDescriptionX, true
}

// SetSBrandingDescriptionX sets field value
func (o *BrandingResponseCompound) SetSBrandingDescriptionX(v string) {
	o.SBrandingDescriptionX = v
}

// GetSBrandingName returns the SBrandingName field value if set, zero value otherwise.
func (o *BrandingResponseCompound) GetSBrandingName() string {
	if o == nil || IsNil(o.SBrandingName) {
		var ret string
		return ret
	}
	return *o.SBrandingName
}

// GetSBrandingNameOk returns a tuple with the SBrandingName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandingResponseCompound) GetSBrandingNameOk() (*string, bool) {
	if o == nil || IsNil(o.SBrandingName) {
		return nil, false
	}
	return o.SBrandingName, true
}

// HasSBrandingName returns a boolean if a field has been set.
func (o *BrandingResponseCompound) HasSBrandingName() bool {
	if o != nil && !IsNil(o.SBrandingName) {
		return true
	}

	return false
}

// SetSBrandingName gets a reference to the given string and assigns it to the SBrandingName field.
func (o *BrandingResponseCompound) SetSBrandingName(v string) {
	o.SBrandingName = &v
}

// GetSEmailAddress returns the SEmailAddress field value if set, zero value otherwise.
func (o *BrandingResponseCompound) GetSEmailAddress() string {
	if o == nil || IsNil(o.SEmailAddress) {
		var ret string
		return ret
	}
	return *o.SEmailAddress
}

// GetSEmailAddressOk returns a tuple with the SEmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandingResponseCompound) GetSEmailAddressOk() (*string, bool) {
	if o == nil || IsNil(o.SEmailAddress) {
		return nil, false
	}
	return o.SEmailAddress, true
}

// HasSEmailAddress returns a boolean if a field has been set.
func (o *BrandingResponseCompound) HasSEmailAddress() bool {
	if o != nil && !IsNil(o.SEmailAddress) {
		return true
	}

	return false
}

// SetSEmailAddress gets a reference to the given string and assigns it to the SEmailAddress field.
func (o *BrandingResponseCompound) SetSEmailAddress(v string) {
	o.SEmailAddress = &v
}

// GetEBrandingLogo returns the EBrandingLogo field value
func (o *BrandingResponseCompound) GetEBrandingLogo() FieldEBrandingLogo {
	if o == nil {
		var ret FieldEBrandingLogo
		return ret
	}

	return o.EBrandingLogo
}

// GetEBrandingLogoOk returns a tuple with the EBrandingLogo field value
// and a boolean to check if the value has been set.
func (o *BrandingResponseCompound) GetEBrandingLogoOk() (*FieldEBrandingLogo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EBrandingLogo, true
}

// SetEBrandingLogo sets field value
func (o *BrandingResponseCompound) SetEBrandingLogo(v FieldEBrandingLogo) {
	o.EBrandingLogo = v
}

// GetIBrandingColortext returns the IBrandingColortext field value
func (o *BrandingResponseCompound) GetIBrandingColortext() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IBrandingColortext
}

// GetIBrandingColortextOk returns a tuple with the IBrandingColortext field value
// and a boolean to check if the value has been set.
func (o *BrandingResponseCompound) GetIBrandingColortextOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IBrandingColortext, true
}

// SetIBrandingColortext sets field value
func (o *BrandingResponseCompound) SetIBrandingColortext(v int32) {
	o.IBrandingColortext = v
}

// GetIBrandingColortextlinkbox returns the IBrandingColortextlinkbox field value
func (o *BrandingResponseCompound) GetIBrandingColortextlinkbox() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IBrandingColortextlinkbox
}

// GetIBrandingColortextlinkboxOk returns a tuple with the IBrandingColortextlinkbox field value
// and a boolean to check if the value has been set.
func (o *BrandingResponseCompound) GetIBrandingColortextlinkboxOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IBrandingColortextlinkbox, true
}

// SetIBrandingColortextlinkbox sets field value
func (o *BrandingResponseCompound) SetIBrandingColortextlinkbox(v int32) {
	o.IBrandingColortextlinkbox = v
}

// GetIBrandingColortextbutton returns the IBrandingColortextbutton field value
func (o *BrandingResponseCompound) GetIBrandingColortextbutton() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IBrandingColortextbutton
}

// GetIBrandingColortextbuttonOk returns a tuple with the IBrandingColortextbutton field value
// and a boolean to check if the value has been set.
func (o *BrandingResponseCompound) GetIBrandingColortextbuttonOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IBrandingColortextbutton, true
}

// SetIBrandingColortextbutton sets field value
func (o *BrandingResponseCompound) SetIBrandingColortextbutton(v int32) {
	o.IBrandingColortextbutton = v
}

// GetIBrandingColorbackground returns the IBrandingColorbackground field value
func (o *BrandingResponseCompound) GetIBrandingColorbackground() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IBrandingColorbackground
}

// GetIBrandingColorbackgroundOk returns a tuple with the IBrandingColorbackground field value
// and a boolean to check if the value has been set.
func (o *BrandingResponseCompound) GetIBrandingColorbackgroundOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IBrandingColorbackground, true
}

// SetIBrandingColorbackground sets field value
func (o *BrandingResponseCompound) SetIBrandingColorbackground(v int32) {
	o.IBrandingColorbackground = v
}

// GetIBrandingColorbackgroundbutton returns the IBrandingColorbackgroundbutton field value
func (o *BrandingResponseCompound) GetIBrandingColorbackgroundbutton() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IBrandingColorbackgroundbutton
}

// GetIBrandingColorbackgroundbuttonOk returns a tuple with the IBrandingColorbackgroundbutton field value
// and a boolean to check if the value has been set.
func (o *BrandingResponseCompound) GetIBrandingColorbackgroundbuttonOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IBrandingColorbackgroundbutton, true
}

// SetIBrandingColorbackgroundbutton sets field value
func (o *BrandingResponseCompound) SetIBrandingColorbackgroundbutton(v int32) {
	o.IBrandingColorbackgroundbutton = v
}

// GetIBrandingColorbackgroundsmallbox returns the IBrandingColorbackgroundsmallbox field value
func (o *BrandingResponseCompound) GetIBrandingColorbackgroundsmallbox() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IBrandingColorbackgroundsmallbox
}

// GetIBrandingColorbackgroundsmallboxOk returns a tuple with the IBrandingColorbackgroundsmallbox field value
// and a boolean to check if the value has been set.
func (o *BrandingResponseCompound) GetIBrandingColorbackgroundsmallboxOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IBrandingColorbackgroundsmallbox, true
}

// SetIBrandingColorbackgroundsmallbox sets field value
func (o *BrandingResponseCompound) SetIBrandingColorbackgroundsmallbox(v int32) {
	o.IBrandingColorbackgroundsmallbox = v
}

// GetBBrandingIsactive returns the BBrandingIsactive field value
func (o *BrandingResponseCompound) GetBBrandingIsactive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BBrandingIsactive
}

// GetBBrandingIsactiveOk returns a tuple with the BBrandingIsactive field value
// and a boolean to check if the value has been set.
func (o *BrandingResponseCompound) GetBBrandingIsactiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BBrandingIsactive, true
}

// SetBBrandingIsactive sets field value
func (o *BrandingResponseCompound) SetBBrandingIsactive(v bool) {
	o.BBrandingIsactive = v
}

// GetSBrandingLogourl returns the SBrandingLogourl field value if set, zero value otherwise.
func (o *BrandingResponseCompound) GetSBrandingLogourl() string {
	if o == nil || IsNil(o.SBrandingLogourl) {
		var ret string
		return ret
	}
	return *o.SBrandingLogourl
}

// GetSBrandingLogourlOk returns a tuple with the SBrandingLogourl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandingResponseCompound) GetSBrandingLogourlOk() (*string, bool) {
	if o == nil || IsNil(o.SBrandingLogourl) {
		return nil, false
	}
	return o.SBrandingLogourl, true
}

// HasSBrandingLogourl returns a boolean if a field has been set.
func (o *BrandingResponseCompound) HasSBrandingLogourl() bool {
	if o != nil && !IsNil(o.SBrandingLogourl) {
		return true
	}

	return false
}

// SetSBrandingLogourl gets a reference to the given string and assigns it to the SBrandingLogourl field.
func (o *BrandingResponseCompound) SetSBrandingLogourl(v string) {
	o.SBrandingLogourl = &v
}

func (o BrandingResponseCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BrandingResponseCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiBrandingID"] = o.PkiBrandingID
	if !IsNil(o.FkiEmailID) {
		toSerialize["fkiEmailID"] = o.FkiEmailID
	}
	toSerialize["objBrandingDescription"] = o.ObjBrandingDescription
	toSerialize["sBrandingDescriptionX"] = o.SBrandingDescriptionX
	if !IsNil(o.SBrandingName) {
		toSerialize["sBrandingName"] = o.SBrandingName
	}
	if !IsNil(o.SEmailAddress) {
		toSerialize["sEmailAddress"] = o.SEmailAddress
	}
	toSerialize["eBrandingLogo"] = o.EBrandingLogo
	toSerialize["iBrandingColortext"] = o.IBrandingColortext
	toSerialize["iBrandingColortextlinkbox"] = o.IBrandingColortextlinkbox
	toSerialize["iBrandingColortextbutton"] = o.IBrandingColortextbutton
	toSerialize["iBrandingColorbackground"] = o.IBrandingColorbackground
	toSerialize["iBrandingColorbackgroundbutton"] = o.IBrandingColorbackgroundbutton
	toSerialize["iBrandingColorbackgroundsmallbox"] = o.IBrandingColorbackgroundsmallbox
	toSerialize["bBrandingIsactive"] = o.BBrandingIsactive
	if !IsNil(o.SBrandingLogourl) {
		toSerialize["sBrandingLogourl"] = o.SBrandingLogourl
	}
	return toSerialize, nil
}

type NullableBrandingResponseCompound struct {
	value *BrandingResponseCompound
	isSet bool
}

func (v NullableBrandingResponseCompound) Get() *BrandingResponseCompound {
	return v.value
}

func (v *NullableBrandingResponseCompound) Set(val *BrandingResponseCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableBrandingResponseCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableBrandingResponseCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrandingResponseCompound(val *BrandingResponseCompound) *NullableBrandingResponseCompound {
	return &NullableBrandingResponseCompound{value: val, isSet: true}
}

func (v NullableBrandingResponseCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBrandingResponseCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

