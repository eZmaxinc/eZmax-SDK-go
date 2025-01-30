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

// checks if the BrandingRequestCompoundV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BrandingRequestCompoundV2{}

// BrandingRequestCompoundV2 A Branding Object and children
type BrandingRequestCompoundV2 struct {
	// The unique ID of the Branding
	PkiBrandingID *int32 `json:"pkiBrandingID,omitempty"`
	ObjBrandingDescription MultilingualBrandingDescription `json:"objBrandingDescription"`
	EBrandingLogo FieldEBrandingLogo `json:"eBrandingLogo"`
	EBrandingAlignlogo *FieldEBrandingAlignlogo `json:"eBrandingAlignlogo,omitempty"`
	// The Base64 encoded binary content of the branding logo. This need to match image type selected in eBrandingLogo if you supply an image. If you select 'Default', the logo will be deleted and the default one will be used.
	SBrandingBase64 *string `json:"sBrandingBase64,omitempty"`
	// The primary color. This is a RGB color converted into integer
	IBrandingColor int32 `json:"iBrandingColor"`
	// The name of the Branding  This value will only be set if you wish to overwrite the default name. If you want to keep the default name, leave this property empty
	SBrandingName *string `json:"sBrandingName,omitempty" validate:"regexp=^.{0,55}$"`
	// The email address.
	SEmailAddress *string "json:\"sEmailAddress,omitempty\" validate:\"regexp=^[\\\\w.%+\\\\-!#$%&'*+\\/=?^`{|}~]+@[a-zA-Z0-9.-]+\\\\.[a-zA-Z]{2,20}$\""
	// Whether the Branding is active or not
	BBrandingIsactive bool `json:"bBrandingIsactive"`
}

type _BrandingRequestCompoundV2 BrandingRequestCompoundV2

// NewBrandingRequestCompoundV2 instantiates a new BrandingRequestCompoundV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrandingRequestCompoundV2(objBrandingDescription MultilingualBrandingDescription, eBrandingLogo FieldEBrandingLogo, iBrandingColor int32, bBrandingIsactive bool) *BrandingRequestCompoundV2 {
	this := BrandingRequestCompoundV2{}
	this.ObjBrandingDescription = objBrandingDescription
	this.EBrandingLogo = eBrandingLogo
	this.IBrandingColor = iBrandingColor
	this.BBrandingIsactive = bBrandingIsactive
	return &this
}

// NewBrandingRequestCompoundV2WithDefaults instantiates a new BrandingRequestCompoundV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrandingRequestCompoundV2WithDefaults() *BrandingRequestCompoundV2 {
	this := BrandingRequestCompoundV2{}
	return &this
}

// GetPkiBrandingID returns the PkiBrandingID field value if set, zero value otherwise.
func (o *BrandingRequestCompoundV2) GetPkiBrandingID() int32 {
	if o == nil || IsNil(o.PkiBrandingID) {
		var ret int32
		return ret
	}
	return *o.PkiBrandingID
}

// GetPkiBrandingIDOk returns a tuple with the PkiBrandingID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandingRequestCompoundV2) GetPkiBrandingIDOk() (*int32, bool) {
	if o == nil || IsNil(o.PkiBrandingID) {
		return nil, false
	}
	return o.PkiBrandingID, true
}

// HasPkiBrandingID returns a boolean if a field has been set.
func (o *BrandingRequestCompoundV2) HasPkiBrandingID() bool {
	if o != nil && !IsNil(o.PkiBrandingID) {
		return true
	}

	return false
}

// SetPkiBrandingID gets a reference to the given int32 and assigns it to the PkiBrandingID field.
func (o *BrandingRequestCompoundV2) SetPkiBrandingID(v int32) {
	o.PkiBrandingID = &v
}

// GetObjBrandingDescription returns the ObjBrandingDescription field value
func (o *BrandingRequestCompoundV2) GetObjBrandingDescription() MultilingualBrandingDescription {
	if o == nil {
		var ret MultilingualBrandingDescription
		return ret
	}

	return o.ObjBrandingDescription
}

// GetObjBrandingDescriptionOk returns a tuple with the ObjBrandingDescription field value
// and a boolean to check if the value has been set.
func (o *BrandingRequestCompoundV2) GetObjBrandingDescriptionOk() (*MultilingualBrandingDescription, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjBrandingDescription, true
}

// SetObjBrandingDescription sets field value
func (o *BrandingRequestCompoundV2) SetObjBrandingDescription(v MultilingualBrandingDescription) {
	o.ObjBrandingDescription = v
}

// GetEBrandingLogo returns the EBrandingLogo field value
func (o *BrandingRequestCompoundV2) GetEBrandingLogo() FieldEBrandingLogo {
	if o == nil {
		var ret FieldEBrandingLogo
		return ret
	}

	return o.EBrandingLogo
}

// GetEBrandingLogoOk returns a tuple with the EBrandingLogo field value
// and a boolean to check if the value has been set.
func (o *BrandingRequestCompoundV2) GetEBrandingLogoOk() (*FieldEBrandingLogo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EBrandingLogo, true
}

// SetEBrandingLogo sets field value
func (o *BrandingRequestCompoundV2) SetEBrandingLogo(v FieldEBrandingLogo) {
	o.EBrandingLogo = v
}

// GetEBrandingAlignlogo returns the EBrandingAlignlogo field value if set, zero value otherwise.
func (o *BrandingRequestCompoundV2) GetEBrandingAlignlogo() FieldEBrandingAlignlogo {
	if o == nil || IsNil(o.EBrandingAlignlogo) {
		var ret FieldEBrandingAlignlogo
		return ret
	}
	return *o.EBrandingAlignlogo
}

// GetEBrandingAlignlogoOk returns a tuple with the EBrandingAlignlogo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandingRequestCompoundV2) GetEBrandingAlignlogoOk() (*FieldEBrandingAlignlogo, bool) {
	if o == nil || IsNil(o.EBrandingAlignlogo) {
		return nil, false
	}
	return o.EBrandingAlignlogo, true
}

// HasEBrandingAlignlogo returns a boolean if a field has been set.
func (o *BrandingRequestCompoundV2) HasEBrandingAlignlogo() bool {
	if o != nil && !IsNil(o.EBrandingAlignlogo) {
		return true
	}

	return false
}

// SetEBrandingAlignlogo gets a reference to the given FieldEBrandingAlignlogo and assigns it to the EBrandingAlignlogo field.
func (o *BrandingRequestCompoundV2) SetEBrandingAlignlogo(v FieldEBrandingAlignlogo) {
	o.EBrandingAlignlogo = &v
}

// GetSBrandingBase64 returns the SBrandingBase64 field value if set, zero value otherwise.
func (o *BrandingRequestCompoundV2) GetSBrandingBase64() string {
	if o == nil || IsNil(o.SBrandingBase64) {
		var ret string
		return ret
	}
	return *o.SBrandingBase64
}

// GetSBrandingBase64Ok returns a tuple with the SBrandingBase64 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandingRequestCompoundV2) GetSBrandingBase64Ok() (*string, bool) {
	if o == nil || IsNil(o.SBrandingBase64) {
		return nil, false
	}
	return o.SBrandingBase64, true
}

// HasSBrandingBase64 returns a boolean if a field has been set.
func (o *BrandingRequestCompoundV2) HasSBrandingBase64() bool {
	if o != nil && !IsNil(o.SBrandingBase64) {
		return true
	}

	return false
}

// SetSBrandingBase64 gets a reference to the given string and assigns it to the SBrandingBase64 field.
func (o *BrandingRequestCompoundV2) SetSBrandingBase64(v string) {
	o.SBrandingBase64 = &v
}

// GetIBrandingColor returns the IBrandingColor field value
func (o *BrandingRequestCompoundV2) GetIBrandingColor() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IBrandingColor
}

// GetIBrandingColorOk returns a tuple with the IBrandingColor field value
// and a boolean to check if the value has been set.
func (o *BrandingRequestCompoundV2) GetIBrandingColorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IBrandingColor, true
}

// SetIBrandingColor sets field value
func (o *BrandingRequestCompoundV2) SetIBrandingColor(v int32) {
	o.IBrandingColor = v
}

// GetSBrandingName returns the SBrandingName field value if set, zero value otherwise.
func (o *BrandingRequestCompoundV2) GetSBrandingName() string {
	if o == nil || IsNil(o.SBrandingName) {
		var ret string
		return ret
	}
	return *o.SBrandingName
}

// GetSBrandingNameOk returns a tuple with the SBrandingName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandingRequestCompoundV2) GetSBrandingNameOk() (*string, bool) {
	if o == nil || IsNil(o.SBrandingName) {
		return nil, false
	}
	return o.SBrandingName, true
}

// HasSBrandingName returns a boolean if a field has been set.
func (o *BrandingRequestCompoundV2) HasSBrandingName() bool {
	if o != nil && !IsNil(o.SBrandingName) {
		return true
	}

	return false
}

// SetSBrandingName gets a reference to the given string and assigns it to the SBrandingName field.
func (o *BrandingRequestCompoundV2) SetSBrandingName(v string) {
	o.SBrandingName = &v
}

// GetSEmailAddress returns the SEmailAddress field value if set, zero value otherwise.
func (o *BrandingRequestCompoundV2) GetSEmailAddress() string {
	if o == nil || IsNil(o.SEmailAddress) {
		var ret string
		return ret
	}
	return *o.SEmailAddress
}

// GetSEmailAddressOk returns a tuple with the SEmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandingRequestCompoundV2) GetSEmailAddressOk() (*string, bool) {
	if o == nil || IsNil(o.SEmailAddress) {
		return nil, false
	}
	return o.SEmailAddress, true
}

// HasSEmailAddress returns a boolean if a field has been set.
func (o *BrandingRequestCompoundV2) HasSEmailAddress() bool {
	if o != nil && !IsNil(o.SEmailAddress) {
		return true
	}

	return false
}

// SetSEmailAddress gets a reference to the given string and assigns it to the SEmailAddress field.
func (o *BrandingRequestCompoundV2) SetSEmailAddress(v string) {
	o.SEmailAddress = &v
}

// GetBBrandingIsactive returns the BBrandingIsactive field value
func (o *BrandingRequestCompoundV2) GetBBrandingIsactive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BBrandingIsactive
}

// GetBBrandingIsactiveOk returns a tuple with the BBrandingIsactive field value
// and a boolean to check if the value has been set.
func (o *BrandingRequestCompoundV2) GetBBrandingIsactiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BBrandingIsactive, true
}

// SetBBrandingIsactive sets field value
func (o *BrandingRequestCompoundV2) SetBBrandingIsactive(v bool) {
	o.BBrandingIsactive = v
}

func (o BrandingRequestCompoundV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BrandingRequestCompoundV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PkiBrandingID) {
		toSerialize["pkiBrandingID"] = o.PkiBrandingID
	}
	toSerialize["objBrandingDescription"] = o.ObjBrandingDescription
	toSerialize["eBrandingLogo"] = o.EBrandingLogo
	if !IsNil(o.EBrandingAlignlogo) {
		toSerialize["eBrandingAlignlogo"] = o.EBrandingAlignlogo
	}
	if !IsNil(o.SBrandingBase64) {
		toSerialize["sBrandingBase64"] = o.SBrandingBase64
	}
	toSerialize["iBrandingColor"] = o.IBrandingColor
	if !IsNil(o.SBrandingName) {
		toSerialize["sBrandingName"] = o.SBrandingName
	}
	if !IsNil(o.SEmailAddress) {
		toSerialize["sEmailAddress"] = o.SEmailAddress
	}
	toSerialize["bBrandingIsactive"] = o.BBrandingIsactive
	return toSerialize, nil
}

func (o *BrandingRequestCompoundV2) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"objBrandingDescription",
		"eBrandingLogo",
		"iBrandingColor",
		"bBrandingIsactive",
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

	varBrandingRequestCompoundV2 := _BrandingRequestCompoundV2{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBrandingRequestCompoundV2)

	if err != nil {
		return err
	}

	*o = BrandingRequestCompoundV2(varBrandingRequestCompoundV2)

	return err
}

type NullableBrandingRequestCompoundV2 struct {
	value *BrandingRequestCompoundV2
	isSet bool
}

func (v NullableBrandingRequestCompoundV2) Get() *BrandingRequestCompoundV2 {
	return v.value
}

func (v *NullableBrandingRequestCompoundV2) Set(val *BrandingRequestCompoundV2) {
	v.value = val
	v.isSet = true
}

func (v NullableBrandingRequestCompoundV2) IsSet() bool {
	return v.isSet
}

func (v *NullableBrandingRequestCompoundV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrandingRequestCompoundV2(val *BrandingRequestCompoundV2) *NullableBrandingRequestCompoundV2 {
	return &NullableBrandingRequestCompoundV2{value: val, isSet: true}
}

func (v NullableBrandingRequestCompoundV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBrandingRequestCompoundV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


