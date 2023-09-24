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

// checks if the AddressRequestCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddressRequestCompound{}

// AddressRequestCompound An Address Object and children to create a complete structure
type AddressRequestCompound struct {
	// The unique ID of the Addresstype.  Valid values:  |Value|Description| |-|-| |1|Office| |2|Home| |3|Real Estate Invoice| |4|Invoicing| |5|Shipping|
	FkiAddresstypeID int32 `json:"fkiAddresstypeID"`
	// The Civic number.
	SAddressCivic string `json:"sAddressCivic"`
	// The Street Name
	SAddressStreet string `json:"sAddressStreet"`
	// The Suite or appartment number
	SAddressSuite string `json:"sAddressSuite"`
	// The City name
	SAddressCity string `json:"sAddressCity"`
	// The unique ID of the Province.  Here are some common values (Complete list must be retrieved from API):  |Value|Description| |-|-| |1|(Canada) Alberta |2|(Canada) British Columbia| |3|(Canada) Manitoba| |3|(Canada) Manitoba| |4|(Canada) New Brunswick| |5|(Canada) Newfoundland| |6|(Canada) Northwest Territories| |7|(Canada) Nova Scotia| |8|(Canada) Nunavut| |9|(Canada) Ontario| |10|(Canada) Prince Edward Island| |11|(Canada) Quebec| |12|(Canada) Saskatchewan| |13|(Canada) Yukon| |14|(United-States) Alabama| |15|(United-States) Alaska| |16|(United-States) Arizona| |17|(United-States) Arkansas| |18|(United-States) California| |19|(United-States) Colorado| |20|(United-States) Connecticut| |21|(United-States) Delaware| |22|(United-States) District of Columbia| |23|(United-States) Florida| |24|(United-States) Georgia| |25|(United-States) Hawaii| |26|(United-States) Idaho| |27|(United-States) Illinois| |28|(United-States) Indiana| |29|(United-States) Iowa| |30|(United-States) Kansas| |31|(United-States) Kentucky| |32|(United-States) Louisiane| |33|(United-States) Maine| |34|(United-States) Maryland| |35|(United-States) Massachusetts| |36|(United-States) Michigan| |37|(United-States) Minnesota| |38|(United-States) Mississippi| |39|(United-States) Missouri| |40|(United-States) Montana| |41|(United-States) Nebraska| |42|(United-States) Nevada| |43|(United-States) New Hampshire| |44|(United-States) New Jersey| |45|(United-States) New Mexico| |46|(United-States) New York| |47|(United-States) North Carolina| |48|(United-States) North Dakota| |49|(United-States) Ohio| |50|(United-States) Oklahoma| |51|(United-States) Oregon| |52|(United-States) Pennsylvania| |53|(United-States) Rhode Island| |54|(United-States) South Carolina| |55|(United-States) South Dakota| |56|(United-States) Tennessee| |57|(United-States) Texas| |58|(United-States) Utah| |60|(United-States) Vermont| |59|(United-States) Virginia| |61|(United-States) Washington| |62|(United-States) West Virginia| |63|(United-States) Wisconsin| |64|(United-States) Wyoming|
	FkiProvinceID int32 `json:"fkiProvinceID"`
	// The unique ID of the Country.  Here are some common values (Complete list must be retrieved from API):  |Value|Description| |-|-| |1|Canada| |2|United-States|
	FkiCountryID int32 `json:"fkiCountryID"`
	// The Postal/Zip Code  The value must be entered without spaces
	SAddressZip string `json:"sAddressZip"`
}

// NewAddressRequestCompound instantiates a new AddressRequestCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressRequestCompound(fkiAddresstypeID int32, sAddressCivic string, sAddressStreet string, sAddressSuite string, sAddressCity string, fkiProvinceID int32, fkiCountryID int32, sAddressZip string) *AddressRequestCompound {
	this := AddressRequestCompound{}
	this.FkiAddresstypeID = fkiAddresstypeID
	this.SAddressCivic = sAddressCivic
	this.SAddressStreet = sAddressStreet
	this.SAddressSuite = sAddressSuite
	this.SAddressCity = sAddressCity
	this.FkiProvinceID = fkiProvinceID
	this.FkiCountryID = fkiCountryID
	this.SAddressZip = sAddressZip
	return &this
}

// NewAddressRequestCompoundWithDefaults instantiates a new AddressRequestCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressRequestCompoundWithDefaults() *AddressRequestCompound {
	this := AddressRequestCompound{}
	return &this
}

// GetFkiAddresstypeID returns the FkiAddresstypeID field value
func (o *AddressRequestCompound) GetFkiAddresstypeID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiAddresstypeID
}

// GetFkiAddresstypeIDOk returns a tuple with the FkiAddresstypeID field value
// and a boolean to check if the value has been set.
func (o *AddressRequestCompound) GetFkiAddresstypeIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiAddresstypeID, true
}

// SetFkiAddresstypeID sets field value
func (o *AddressRequestCompound) SetFkiAddresstypeID(v int32) {
	o.FkiAddresstypeID = v
}

// GetSAddressCivic returns the SAddressCivic field value
func (o *AddressRequestCompound) GetSAddressCivic() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SAddressCivic
}

// GetSAddressCivicOk returns a tuple with the SAddressCivic field value
// and a boolean to check if the value has been set.
func (o *AddressRequestCompound) GetSAddressCivicOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SAddressCivic, true
}

// SetSAddressCivic sets field value
func (o *AddressRequestCompound) SetSAddressCivic(v string) {
	o.SAddressCivic = v
}

// GetSAddressStreet returns the SAddressStreet field value
func (o *AddressRequestCompound) GetSAddressStreet() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SAddressStreet
}

// GetSAddressStreetOk returns a tuple with the SAddressStreet field value
// and a boolean to check if the value has been set.
func (o *AddressRequestCompound) GetSAddressStreetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SAddressStreet, true
}

// SetSAddressStreet sets field value
func (o *AddressRequestCompound) SetSAddressStreet(v string) {
	o.SAddressStreet = v
}

// GetSAddressSuite returns the SAddressSuite field value
func (o *AddressRequestCompound) GetSAddressSuite() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SAddressSuite
}

// GetSAddressSuiteOk returns a tuple with the SAddressSuite field value
// and a boolean to check if the value has been set.
func (o *AddressRequestCompound) GetSAddressSuiteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SAddressSuite, true
}

// SetSAddressSuite sets field value
func (o *AddressRequestCompound) SetSAddressSuite(v string) {
	o.SAddressSuite = v
}

// GetSAddressCity returns the SAddressCity field value
func (o *AddressRequestCompound) GetSAddressCity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SAddressCity
}

// GetSAddressCityOk returns a tuple with the SAddressCity field value
// and a boolean to check if the value has been set.
func (o *AddressRequestCompound) GetSAddressCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SAddressCity, true
}

// SetSAddressCity sets field value
func (o *AddressRequestCompound) SetSAddressCity(v string) {
	o.SAddressCity = v
}

// GetFkiProvinceID returns the FkiProvinceID field value
func (o *AddressRequestCompound) GetFkiProvinceID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiProvinceID
}

// GetFkiProvinceIDOk returns a tuple with the FkiProvinceID field value
// and a boolean to check if the value has been set.
func (o *AddressRequestCompound) GetFkiProvinceIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiProvinceID, true
}

// SetFkiProvinceID sets field value
func (o *AddressRequestCompound) SetFkiProvinceID(v int32) {
	o.FkiProvinceID = v
}

// GetFkiCountryID returns the FkiCountryID field value
func (o *AddressRequestCompound) GetFkiCountryID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiCountryID
}

// GetFkiCountryIDOk returns a tuple with the FkiCountryID field value
// and a boolean to check if the value has been set.
func (o *AddressRequestCompound) GetFkiCountryIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiCountryID, true
}

// SetFkiCountryID sets field value
func (o *AddressRequestCompound) SetFkiCountryID(v int32) {
	o.FkiCountryID = v
}

// GetSAddressZip returns the SAddressZip field value
func (o *AddressRequestCompound) GetSAddressZip() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SAddressZip
}

// GetSAddressZipOk returns a tuple with the SAddressZip field value
// and a boolean to check if the value has been set.
func (o *AddressRequestCompound) GetSAddressZipOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SAddressZip, true
}

// SetSAddressZip sets field value
func (o *AddressRequestCompound) SetSAddressZip(v string) {
	o.SAddressZip = v
}

func (o AddressRequestCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressRequestCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fkiAddresstypeID"] = o.FkiAddresstypeID
	toSerialize["sAddressCivic"] = o.SAddressCivic
	toSerialize["sAddressStreet"] = o.SAddressStreet
	toSerialize["sAddressSuite"] = o.SAddressSuite
	toSerialize["sAddressCity"] = o.SAddressCity
	toSerialize["fkiProvinceID"] = o.FkiProvinceID
	toSerialize["fkiCountryID"] = o.FkiCountryID
	toSerialize["sAddressZip"] = o.SAddressZip
	return toSerialize, nil
}

type NullableAddressRequestCompound struct {
	value *AddressRequestCompound
	isSet bool
}

func (v NullableAddressRequestCompound) Get() *AddressRequestCompound {
	return v.value
}

func (v *NullableAddressRequestCompound) Set(val *AddressRequestCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressRequestCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressRequestCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressRequestCompound(val *AddressRequestCompound) *NullableAddressRequestCompound {
	return &NullableAddressRequestCompound{value: val, isSet: true}
}

func (v NullableAddressRequestCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressRequestCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

