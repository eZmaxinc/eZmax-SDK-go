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

// checks if the ProvinceAutocompleteElementResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProvinceAutocompleteElementResponse{}

// ProvinceAutocompleteElementResponse A Province AutocompleteElement Response
type ProvinceAutocompleteElementResponse struct {
	// The unique ID of the Province.  Here are some common values (Complete list must be retrieved from API):  |Value|Description| |-|-| |1|(Canada) Alberta |2|(Canada) British Columbia| |3|(Canada) Manitoba| |3|(Canada) Manitoba| |4|(Canada) New Brunswick| |5|(Canada) Newfoundland| |6|(Canada) Northwest Territories| |7|(Canada) Nova Scotia| |8|(Canada) Nunavut| |9|(Canada) Ontario| |10|(Canada) Prince Edward Island| |11|(Canada) Quebec| |12|(Canada) Saskatchewan| |13|(Canada) Yukon| |14|(United-States) Alabama| |15|(United-States) Alaska| |16|(United-States) Arizona| |17|(United-States) Arkansas| |18|(United-States) California| |19|(United-States) Colorado| |20|(United-States) Connecticut| |21|(United-States) Delaware| |22|(United-States) District of Columbia| |23|(United-States) Florida| |24|(United-States) Georgia| |25|(United-States) Hawaii| |26|(United-States) Idaho| |27|(United-States) Illinois| |28|(United-States) Indiana| |29|(United-States) Iowa| |30|(United-States) Kansas| |31|(United-States) Kentucky| |32|(United-States) Louisiane| |33|(United-States) Maine| |34|(United-States) Maryland| |35|(United-States) Massachusetts| |36|(United-States) Michigan| |37|(United-States) Minnesota| |38|(United-States) Mississippi| |39|(United-States) Missouri| |40|(United-States) Montana| |41|(United-States) Nebraska| |42|(United-States) Nevada| |43|(United-States) New Hampshire| |44|(United-States) New Jersey| |45|(United-States) New Mexico| |46|(United-States) New York| |47|(United-States) North Carolina| |48|(United-States) North Dakota| |49|(United-States) Ohio| |50|(United-States) Oklahoma| |51|(United-States) Oregon| |52|(United-States) Pennsylvania| |53|(United-States) Rhode Island| |54|(United-States) South Carolina| |55|(United-States) South Dakota| |56|(United-States) Tennessee| |57|(United-States) Texas| |58|(United-States) Utah| |60|(United-States) Vermont| |59|(United-States) Virginia| |61|(United-States) Washington| |62|(United-States) West Virginia| |63|(United-States) Wisconsin| |64|(United-States) Wyoming|
	PkiProvinceID int32 `json:"pkiProvinceID"`
	// The unique ID of the Country.  Here are some common values (Complete list must be retrieved from API):  |Value|Description| |-|-| |1|Canada| |2|United-States|
	FkiCountryID int32 `json:"fkiCountryID"`
	// The name of the Province in the language of the requester
	SProvinceNameX string `json:"sProvinceNameX" validate:"regexp=^.{0,50}$"`
	// The shortname of the Province
	SProvinceShortname string `json:"sProvinceShortname" validate:"regexp=^.{1,3}$"`
	// Whether the Province is active or not
	BProvinceIsactive bool `json:"bProvinceIsactive"`
}

type _ProvinceAutocompleteElementResponse ProvinceAutocompleteElementResponse

// NewProvinceAutocompleteElementResponse instantiates a new ProvinceAutocompleteElementResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvinceAutocompleteElementResponse(pkiProvinceID int32, fkiCountryID int32, sProvinceNameX string, sProvinceShortname string, bProvinceIsactive bool) *ProvinceAutocompleteElementResponse {
	this := ProvinceAutocompleteElementResponse{}
	this.PkiProvinceID = pkiProvinceID
	this.FkiCountryID = fkiCountryID
	this.SProvinceNameX = sProvinceNameX
	this.SProvinceShortname = sProvinceShortname
	this.BProvinceIsactive = bProvinceIsactive
	return &this
}

// NewProvinceAutocompleteElementResponseWithDefaults instantiates a new ProvinceAutocompleteElementResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvinceAutocompleteElementResponseWithDefaults() *ProvinceAutocompleteElementResponse {
	this := ProvinceAutocompleteElementResponse{}
	return &this
}

// GetPkiProvinceID returns the PkiProvinceID field value
func (o *ProvinceAutocompleteElementResponse) GetPkiProvinceID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiProvinceID
}

// GetPkiProvinceIDOk returns a tuple with the PkiProvinceID field value
// and a boolean to check if the value has been set.
func (o *ProvinceAutocompleteElementResponse) GetPkiProvinceIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiProvinceID, true
}

// SetPkiProvinceID sets field value
func (o *ProvinceAutocompleteElementResponse) SetPkiProvinceID(v int32) {
	o.PkiProvinceID = v
}

// GetFkiCountryID returns the FkiCountryID field value
func (o *ProvinceAutocompleteElementResponse) GetFkiCountryID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiCountryID
}

// GetFkiCountryIDOk returns a tuple with the FkiCountryID field value
// and a boolean to check if the value has been set.
func (o *ProvinceAutocompleteElementResponse) GetFkiCountryIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiCountryID, true
}

// SetFkiCountryID sets field value
func (o *ProvinceAutocompleteElementResponse) SetFkiCountryID(v int32) {
	o.FkiCountryID = v
}

// GetSProvinceNameX returns the SProvinceNameX field value
func (o *ProvinceAutocompleteElementResponse) GetSProvinceNameX() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SProvinceNameX
}

// GetSProvinceNameXOk returns a tuple with the SProvinceNameX field value
// and a boolean to check if the value has been set.
func (o *ProvinceAutocompleteElementResponse) GetSProvinceNameXOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SProvinceNameX, true
}

// SetSProvinceNameX sets field value
func (o *ProvinceAutocompleteElementResponse) SetSProvinceNameX(v string) {
	o.SProvinceNameX = v
}

// GetSProvinceShortname returns the SProvinceShortname field value
func (o *ProvinceAutocompleteElementResponse) GetSProvinceShortname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SProvinceShortname
}

// GetSProvinceShortnameOk returns a tuple with the SProvinceShortname field value
// and a boolean to check if the value has been set.
func (o *ProvinceAutocompleteElementResponse) GetSProvinceShortnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SProvinceShortname, true
}

// SetSProvinceShortname sets field value
func (o *ProvinceAutocompleteElementResponse) SetSProvinceShortname(v string) {
	o.SProvinceShortname = v
}

// GetBProvinceIsactive returns the BProvinceIsactive field value
func (o *ProvinceAutocompleteElementResponse) GetBProvinceIsactive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BProvinceIsactive
}

// GetBProvinceIsactiveOk returns a tuple with the BProvinceIsactive field value
// and a boolean to check if the value has been set.
func (o *ProvinceAutocompleteElementResponse) GetBProvinceIsactiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BProvinceIsactive, true
}

// SetBProvinceIsactive sets field value
func (o *ProvinceAutocompleteElementResponse) SetBProvinceIsactive(v bool) {
	o.BProvinceIsactive = v
}

func (o ProvinceAutocompleteElementResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProvinceAutocompleteElementResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiProvinceID"] = o.PkiProvinceID
	toSerialize["fkiCountryID"] = o.FkiCountryID
	toSerialize["sProvinceNameX"] = o.SProvinceNameX
	toSerialize["sProvinceShortname"] = o.SProvinceShortname
	toSerialize["bProvinceIsactive"] = o.BProvinceIsactive
	return toSerialize, nil
}

func (o *ProvinceAutocompleteElementResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiProvinceID",
		"fkiCountryID",
		"sProvinceNameX",
		"sProvinceShortname",
		"bProvinceIsactive",
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

	varProvinceAutocompleteElementResponse := _ProvinceAutocompleteElementResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProvinceAutocompleteElementResponse)

	if err != nil {
		return err
	}

	*o = ProvinceAutocompleteElementResponse(varProvinceAutocompleteElementResponse)

	return err
}

type NullableProvinceAutocompleteElementResponse struct {
	value *ProvinceAutocompleteElementResponse
	isSet bool
}

func (v NullableProvinceAutocompleteElementResponse) Get() *ProvinceAutocompleteElementResponse {
	return v.value
}

func (v *NullableProvinceAutocompleteElementResponse) Set(val *ProvinceAutocompleteElementResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProvinceAutocompleteElementResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProvinceAutocompleteElementResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvinceAutocompleteElementResponse(val *ProvinceAutocompleteElementResponse) *NullableProvinceAutocompleteElementResponse {
	return &NullableProvinceAutocompleteElementResponse{value: val, isSet: true}
}

func (v NullableProvinceAutocompleteElementResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvinceAutocompleteElementResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


