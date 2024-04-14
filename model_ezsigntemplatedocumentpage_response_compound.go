/*
eZmax API Definition (Full)

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.2.1
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the EzsigntemplatedocumentpageResponseCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplatedocumentpageResponseCompound{}

// EzsigntemplatedocumentpageResponseCompound An Ezsigntemplatedocumentpage Object and children to create a complete structure
type EzsigntemplatedocumentpageResponseCompound struct {
	// The unique ID of the Ezsigntemplatedocumentpage
	PkiEzsigntemplatedocumentpageID int32 `json:"pkiEzsigntemplatedocumentpageID"`
	// The Width of the page's image in pixels calculated at 100 DPI
	IEzsigntemplatedocumentpageWidthimage int32 `json:"iEzsigntemplatedocumentpageWidthimage"`
	// The Height of the page's image in pixels calculated at 100 DPI
	IEzsigntemplatedocumentpageHeightimage int32 `json:"iEzsigntemplatedocumentpageHeightimage"`
	// The Width of the page in points calculated at 72 DPI
	IEzsigntemplatedocumentpageWidthpdf int32 `json:"iEzsigntemplatedocumentpageWidthpdf"`
	// The Height of the page in points calculated at 72 DPI
	IEzsigntemplatedocumentpageHeightpdf int32 `json:"iEzsigntemplatedocumentpageHeightpdf"`
	// The page number in the Ezsigntemplatedocument
	IEzsigntemplatedocumentpagePagenumber int32 `json:"iEzsigntemplatedocumentpagePagenumber"`
	// The Url to the Ezsigntemplatedocumentpage's rasterized image.  Url will expire after 5 minutes.
	SComputedImageurl string `json:"sComputedImageurl"`
}

type _EzsigntemplatedocumentpageResponseCompound EzsigntemplatedocumentpageResponseCompound

// NewEzsigntemplatedocumentpageResponseCompound instantiates a new EzsigntemplatedocumentpageResponseCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplatedocumentpageResponseCompound(pkiEzsigntemplatedocumentpageID int32, iEzsigntemplatedocumentpageWidthimage int32, iEzsigntemplatedocumentpageHeightimage int32, iEzsigntemplatedocumentpageWidthpdf int32, iEzsigntemplatedocumentpageHeightpdf int32, iEzsigntemplatedocumentpagePagenumber int32, sComputedImageurl string) *EzsigntemplatedocumentpageResponseCompound {
	this := EzsigntemplatedocumentpageResponseCompound{}
	this.PkiEzsigntemplatedocumentpageID = pkiEzsigntemplatedocumentpageID
	this.IEzsigntemplatedocumentpageWidthimage = iEzsigntemplatedocumentpageWidthimage
	this.IEzsigntemplatedocumentpageHeightimage = iEzsigntemplatedocumentpageHeightimage
	this.IEzsigntemplatedocumentpageWidthpdf = iEzsigntemplatedocumentpageWidthpdf
	this.IEzsigntemplatedocumentpageHeightpdf = iEzsigntemplatedocumentpageHeightpdf
	this.IEzsigntemplatedocumentpagePagenumber = iEzsigntemplatedocumentpagePagenumber
	this.SComputedImageurl = sComputedImageurl
	return &this
}

// NewEzsigntemplatedocumentpageResponseCompoundWithDefaults instantiates a new EzsigntemplatedocumentpageResponseCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplatedocumentpageResponseCompoundWithDefaults() *EzsigntemplatedocumentpageResponseCompound {
	this := EzsigntemplatedocumentpageResponseCompound{}
	return &this
}

// GetPkiEzsigntemplatedocumentpageID returns the PkiEzsigntemplatedocumentpageID field value
func (o *EzsigntemplatedocumentpageResponseCompound) GetPkiEzsigntemplatedocumentpageID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiEzsigntemplatedocumentpageID
}

// GetPkiEzsigntemplatedocumentpageIDOk returns a tuple with the PkiEzsigntemplatedocumentpageID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatedocumentpageResponseCompound) GetPkiEzsigntemplatedocumentpageIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiEzsigntemplatedocumentpageID, true
}

// SetPkiEzsigntemplatedocumentpageID sets field value
func (o *EzsigntemplatedocumentpageResponseCompound) SetPkiEzsigntemplatedocumentpageID(v int32) {
	o.PkiEzsigntemplatedocumentpageID = v
}

// GetIEzsigntemplatedocumentpageWidthimage returns the IEzsigntemplatedocumentpageWidthimage field value
func (o *EzsigntemplatedocumentpageResponseCompound) GetIEzsigntemplatedocumentpageWidthimage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsigntemplatedocumentpageWidthimage
}

// GetIEzsigntemplatedocumentpageWidthimageOk returns a tuple with the IEzsigntemplatedocumentpageWidthimage field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatedocumentpageResponseCompound) GetIEzsigntemplatedocumentpageWidthimageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsigntemplatedocumentpageWidthimage, true
}

// SetIEzsigntemplatedocumentpageWidthimage sets field value
func (o *EzsigntemplatedocumentpageResponseCompound) SetIEzsigntemplatedocumentpageWidthimage(v int32) {
	o.IEzsigntemplatedocumentpageWidthimage = v
}

// GetIEzsigntemplatedocumentpageHeightimage returns the IEzsigntemplatedocumentpageHeightimage field value
func (o *EzsigntemplatedocumentpageResponseCompound) GetIEzsigntemplatedocumentpageHeightimage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsigntemplatedocumentpageHeightimage
}

// GetIEzsigntemplatedocumentpageHeightimageOk returns a tuple with the IEzsigntemplatedocumentpageHeightimage field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatedocumentpageResponseCompound) GetIEzsigntemplatedocumentpageHeightimageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsigntemplatedocumentpageHeightimage, true
}

// SetIEzsigntemplatedocumentpageHeightimage sets field value
func (o *EzsigntemplatedocumentpageResponseCompound) SetIEzsigntemplatedocumentpageHeightimage(v int32) {
	o.IEzsigntemplatedocumentpageHeightimage = v
}

// GetIEzsigntemplatedocumentpageWidthpdf returns the IEzsigntemplatedocumentpageWidthpdf field value
func (o *EzsigntemplatedocumentpageResponseCompound) GetIEzsigntemplatedocumentpageWidthpdf() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsigntemplatedocumentpageWidthpdf
}

// GetIEzsigntemplatedocumentpageWidthpdfOk returns a tuple with the IEzsigntemplatedocumentpageWidthpdf field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatedocumentpageResponseCompound) GetIEzsigntemplatedocumentpageWidthpdfOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsigntemplatedocumentpageWidthpdf, true
}

// SetIEzsigntemplatedocumentpageWidthpdf sets field value
func (o *EzsigntemplatedocumentpageResponseCompound) SetIEzsigntemplatedocumentpageWidthpdf(v int32) {
	o.IEzsigntemplatedocumentpageWidthpdf = v
}

// GetIEzsigntemplatedocumentpageHeightpdf returns the IEzsigntemplatedocumentpageHeightpdf field value
func (o *EzsigntemplatedocumentpageResponseCompound) GetIEzsigntemplatedocumentpageHeightpdf() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsigntemplatedocumentpageHeightpdf
}

// GetIEzsigntemplatedocumentpageHeightpdfOk returns a tuple with the IEzsigntemplatedocumentpageHeightpdf field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatedocumentpageResponseCompound) GetIEzsigntemplatedocumentpageHeightpdfOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsigntemplatedocumentpageHeightpdf, true
}

// SetIEzsigntemplatedocumentpageHeightpdf sets field value
func (o *EzsigntemplatedocumentpageResponseCompound) SetIEzsigntemplatedocumentpageHeightpdf(v int32) {
	o.IEzsigntemplatedocumentpageHeightpdf = v
}

// GetIEzsigntemplatedocumentpagePagenumber returns the IEzsigntemplatedocumentpagePagenumber field value
func (o *EzsigntemplatedocumentpageResponseCompound) GetIEzsigntemplatedocumentpagePagenumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsigntemplatedocumentpagePagenumber
}

// GetIEzsigntemplatedocumentpagePagenumberOk returns a tuple with the IEzsigntemplatedocumentpagePagenumber field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatedocumentpageResponseCompound) GetIEzsigntemplatedocumentpagePagenumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsigntemplatedocumentpagePagenumber, true
}

// SetIEzsigntemplatedocumentpagePagenumber sets field value
func (o *EzsigntemplatedocumentpageResponseCompound) SetIEzsigntemplatedocumentpagePagenumber(v int32) {
	o.IEzsigntemplatedocumentpagePagenumber = v
}

// GetSComputedImageurl returns the SComputedImageurl field value
func (o *EzsigntemplatedocumentpageResponseCompound) GetSComputedImageurl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SComputedImageurl
}

// GetSComputedImageurlOk returns a tuple with the SComputedImageurl field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatedocumentpageResponseCompound) GetSComputedImageurlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SComputedImageurl, true
}

// SetSComputedImageurl sets field value
func (o *EzsigntemplatedocumentpageResponseCompound) SetSComputedImageurl(v string) {
	o.SComputedImageurl = v
}

func (o EzsigntemplatedocumentpageResponseCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplatedocumentpageResponseCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiEzsigntemplatedocumentpageID"] = o.PkiEzsigntemplatedocumentpageID
	toSerialize["iEzsigntemplatedocumentpageWidthimage"] = o.IEzsigntemplatedocumentpageWidthimage
	toSerialize["iEzsigntemplatedocumentpageHeightimage"] = o.IEzsigntemplatedocumentpageHeightimage
	toSerialize["iEzsigntemplatedocumentpageWidthpdf"] = o.IEzsigntemplatedocumentpageWidthpdf
	toSerialize["iEzsigntemplatedocumentpageHeightpdf"] = o.IEzsigntemplatedocumentpageHeightpdf
	toSerialize["iEzsigntemplatedocumentpagePagenumber"] = o.IEzsigntemplatedocumentpagePagenumber
	toSerialize["sComputedImageurl"] = o.SComputedImageurl
	return toSerialize, nil
}

func (o *EzsigntemplatedocumentpageResponseCompound) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiEzsigntemplatedocumentpageID",
		"iEzsigntemplatedocumentpageWidthimage",
		"iEzsigntemplatedocumentpageHeightimage",
		"iEzsigntemplatedocumentpageWidthpdf",
		"iEzsigntemplatedocumentpageHeightpdf",
		"iEzsigntemplatedocumentpagePagenumber",
		"sComputedImageurl",
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

	varEzsigntemplatedocumentpageResponseCompound := _EzsigntemplatedocumentpageResponseCompound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigntemplatedocumentpageResponseCompound)

	if err != nil {
		return err
	}

	*o = EzsigntemplatedocumentpageResponseCompound(varEzsigntemplatedocumentpageResponseCompound)

	return err
}

type NullableEzsigntemplatedocumentpageResponseCompound struct {
	value *EzsigntemplatedocumentpageResponseCompound
	isSet bool
}

func (v NullableEzsigntemplatedocumentpageResponseCompound) Get() *EzsigntemplatedocumentpageResponseCompound {
	return v.value
}

func (v *NullableEzsigntemplatedocumentpageResponseCompound) Set(val *EzsigntemplatedocumentpageResponseCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplatedocumentpageResponseCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplatedocumentpageResponseCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplatedocumentpageResponseCompound(val *EzsigntemplatedocumentpageResponseCompound) *NullableEzsigntemplatedocumentpageResponseCompound {
	return &NullableEzsigntemplatedocumentpageResponseCompound{value: val, isSet: true}
}

func (v NullableEzsigntemplatedocumentpageResponseCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplatedocumentpageResponseCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


