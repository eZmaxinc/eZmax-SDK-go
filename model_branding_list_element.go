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

// checks if the BrandingListElement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BrandingListElement{}

// BrandingListElement A Branding List Element
type BrandingListElement struct {
	// The unique ID of the Branding
	PkiBrandingID int32 `json:"pkiBrandingID"`
	// The Description of the Branding in the language of the requester
	SBrandingDescriptionX string `json:"sBrandingDescriptionX"`
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
}

type _BrandingListElement BrandingListElement

// NewBrandingListElement instantiates a new BrandingListElement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrandingListElement(pkiBrandingID int32, sBrandingDescriptionX string, iBrandingColortext int32, iBrandingColortextlinkbox int32, iBrandingColortextbutton int32, iBrandingColorbackground int32, iBrandingColorbackgroundbutton int32, iBrandingColorbackgroundsmallbox int32, bBrandingIsactive bool) *BrandingListElement {
	this := BrandingListElement{}
	this.PkiBrandingID = pkiBrandingID
	this.SBrandingDescriptionX = sBrandingDescriptionX
	this.IBrandingColortext = iBrandingColortext
	this.IBrandingColortextlinkbox = iBrandingColortextlinkbox
	this.IBrandingColortextbutton = iBrandingColortextbutton
	this.IBrandingColorbackground = iBrandingColorbackground
	this.IBrandingColorbackgroundbutton = iBrandingColorbackgroundbutton
	this.IBrandingColorbackgroundsmallbox = iBrandingColorbackgroundsmallbox
	this.BBrandingIsactive = bBrandingIsactive
	return &this
}

// NewBrandingListElementWithDefaults instantiates a new BrandingListElement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrandingListElementWithDefaults() *BrandingListElement {
	this := BrandingListElement{}
	return &this
}

// GetPkiBrandingID returns the PkiBrandingID field value
func (o *BrandingListElement) GetPkiBrandingID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiBrandingID
}

// GetPkiBrandingIDOk returns a tuple with the PkiBrandingID field value
// and a boolean to check if the value has been set.
func (o *BrandingListElement) GetPkiBrandingIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiBrandingID, true
}

// SetPkiBrandingID sets field value
func (o *BrandingListElement) SetPkiBrandingID(v int32) {
	o.PkiBrandingID = v
}

// GetSBrandingDescriptionX returns the SBrandingDescriptionX field value
func (o *BrandingListElement) GetSBrandingDescriptionX() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SBrandingDescriptionX
}

// GetSBrandingDescriptionXOk returns a tuple with the SBrandingDescriptionX field value
// and a boolean to check if the value has been set.
func (o *BrandingListElement) GetSBrandingDescriptionXOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SBrandingDescriptionX, true
}

// SetSBrandingDescriptionX sets field value
func (o *BrandingListElement) SetSBrandingDescriptionX(v string) {
	o.SBrandingDescriptionX = v
}

// GetIBrandingColortext returns the IBrandingColortext field value
func (o *BrandingListElement) GetIBrandingColortext() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IBrandingColortext
}

// GetIBrandingColortextOk returns a tuple with the IBrandingColortext field value
// and a boolean to check if the value has been set.
func (o *BrandingListElement) GetIBrandingColortextOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IBrandingColortext, true
}

// SetIBrandingColortext sets field value
func (o *BrandingListElement) SetIBrandingColortext(v int32) {
	o.IBrandingColortext = v
}

// GetIBrandingColortextlinkbox returns the IBrandingColortextlinkbox field value
func (o *BrandingListElement) GetIBrandingColortextlinkbox() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IBrandingColortextlinkbox
}

// GetIBrandingColortextlinkboxOk returns a tuple with the IBrandingColortextlinkbox field value
// and a boolean to check if the value has been set.
func (o *BrandingListElement) GetIBrandingColortextlinkboxOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IBrandingColortextlinkbox, true
}

// SetIBrandingColortextlinkbox sets field value
func (o *BrandingListElement) SetIBrandingColortextlinkbox(v int32) {
	o.IBrandingColortextlinkbox = v
}

// GetIBrandingColortextbutton returns the IBrandingColortextbutton field value
func (o *BrandingListElement) GetIBrandingColortextbutton() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IBrandingColortextbutton
}

// GetIBrandingColortextbuttonOk returns a tuple with the IBrandingColortextbutton field value
// and a boolean to check if the value has been set.
func (o *BrandingListElement) GetIBrandingColortextbuttonOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IBrandingColortextbutton, true
}

// SetIBrandingColortextbutton sets field value
func (o *BrandingListElement) SetIBrandingColortextbutton(v int32) {
	o.IBrandingColortextbutton = v
}

// GetIBrandingColorbackground returns the IBrandingColorbackground field value
func (o *BrandingListElement) GetIBrandingColorbackground() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IBrandingColorbackground
}

// GetIBrandingColorbackgroundOk returns a tuple with the IBrandingColorbackground field value
// and a boolean to check if the value has been set.
func (o *BrandingListElement) GetIBrandingColorbackgroundOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IBrandingColorbackground, true
}

// SetIBrandingColorbackground sets field value
func (o *BrandingListElement) SetIBrandingColorbackground(v int32) {
	o.IBrandingColorbackground = v
}

// GetIBrandingColorbackgroundbutton returns the IBrandingColorbackgroundbutton field value
func (o *BrandingListElement) GetIBrandingColorbackgroundbutton() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IBrandingColorbackgroundbutton
}

// GetIBrandingColorbackgroundbuttonOk returns a tuple with the IBrandingColorbackgroundbutton field value
// and a boolean to check if the value has been set.
func (o *BrandingListElement) GetIBrandingColorbackgroundbuttonOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IBrandingColorbackgroundbutton, true
}

// SetIBrandingColorbackgroundbutton sets field value
func (o *BrandingListElement) SetIBrandingColorbackgroundbutton(v int32) {
	o.IBrandingColorbackgroundbutton = v
}

// GetIBrandingColorbackgroundsmallbox returns the IBrandingColorbackgroundsmallbox field value
func (o *BrandingListElement) GetIBrandingColorbackgroundsmallbox() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IBrandingColorbackgroundsmallbox
}

// GetIBrandingColorbackgroundsmallboxOk returns a tuple with the IBrandingColorbackgroundsmallbox field value
// and a boolean to check if the value has been set.
func (o *BrandingListElement) GetIBrandingColorbackgroundsmallboxOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IBrandingColorbackgroundsmallbox, true
}

// SetIBrandingColorbackgroundsmallbox sets field value
func (o *BrandingListElement) SetIBrandingColorbackgroundsmallbox(v int32) {
	o.IBrandingColorbackgroundsmallbox = v
}

// GetBBrandingIsactive returns the BBrandingIsactive field value
func (o *BrandingListElement) GetBBrandingIsactive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BBrandingIsactive
}

// GetBBrandingIsactiveOk returns a tuple with the BBrandingIsactive field value
// and a boolean to check if the value has been set.
func (o *BrandingListElement) GetBBrandingIsactiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BBrandingIsactive, true
}

// SetBBrandingIsactive sets field value
func (o *BrandingListElement) SetBBrandingIsactive(v bool) {
	o.BBrandingIsactive = v
}

func (o BrandingListElement) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BrandingListElement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiBrandingID"] = o.PkiBrandingID
	toSerialize["sBrandingDescriptionX"] = o.SBrandingDescriptionX
	toSerialize["iBrandingColortext"] = o.IBrandingColortext
	toSerialize["iBrandingColortextlinkbox"] = o.IBrandingColortextlinkbox
	toSerialize["iBrandingColortextbutton"] = o.IBrandingColortextbutton
	toSerialize["iBrandingColorbackground"] = o.IBrandingColorbackground
	toSerialize["iBrandingColorbackgroundbutton"] = o.IBrandingColorbackgroundbutton
	toSerialize["iBrandingColorbackgroundsmallbox"] = o.IBrandingColorbackgroundsmallbox
	toSerialize["bBrandingIsactive"] = o.BBrandingIsactive
	return toSerialize, nil
}

func (o *BrandingListElement) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiBrandingID",
		"sBrandingDescriptionX",
		"iBrandingColortext",
		"iBrandingColortextlinkbox",
		"iBrandingColortextbutton",
		"iBrandingColorbackground",
		"iBrandingColorbackgroundbutton",
		"iBrandingColorbackgroundsmallbox",
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

	varBrandingListElement := _BrandingListElement{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBrandingListElement)

	if err != nil {
		return err
	}

	*o = BrandingListElement(varBrandingListElement)

	return err
}

type NullableBrandingListElement struct {
	value *BrandingListElement
	isSet bool
}

func (v NullableBrandingListElement) Get() *BrandingListElement {
	return v.value
}

func (v *NullableBrandingListElement) Set(val *BrandingListElement) {
	v.value = val
	v.isSet = true
}

func (v NullableBrandingListElement) IsSet() bool {
	return v.isSet
}

func (v *NullableBrandingListElement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrandingListElement(val *BrandingListElement) *NullableBrandingListElement {
	return &NullableBrandingListElement{value: val, isSet: true}
}

func (v NullableBrandingListElement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBrandingListElement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


