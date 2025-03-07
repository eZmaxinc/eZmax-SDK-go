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

// checks if the SessionhistoryListElement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SessionhistoryListElement{}

// SessionhistoryListElement A Sessionhistory List Element
type SessionhistoryListElement struct {
	// The unique ID of the Sessionhistory
	PkiSessionhistoryID int32 `json:"pkiSessionhistoryID"`
	// The unique ID of the Computer
	FkiComputerID *int32 `json:"fkiComputerID,omitempty"`
	// The unique ID of the User
	FkiUserID *int32 `json:"fkiUserID,omitempty"`
	// The first hit of the Sessionhistory
	DtSessionhistoryFirsthit string `json:"dtSessionhistoryFirsthit" validate:"regexp=^[0-9]{4}-(0[1-9]|1[0-2])-(0[1-9]|[1-2][0-9]|3[0-1]) ([01]?[0-9]|2[0-3]):([0-5][0-9]):([0-5][0-9])$"`
	// The last hit of the Sessionhistory
	DtSessionhistoryLasthit string `json:"dtSessionhistoryLasthit" validate:"regexp=^[0-9]{4}-(0[1-9]|1[0-2])-(0[1-9]|[1-2][0-9]|3[0-1]) ([01]?[0-9]|2[0-3]):([0-5][0-9]):([0-5][0-9])$"`
	ESessionhistoryEndby FieldESessionhistoryEndby `json:"eSessionhistoryEndby"`
	// The description of the Computer
	SComputerDescription *string `json:"sComputerDescription,omitempty" validate:"regexp=^.{0,50}$"`
	// The duration of the session
	SSessionhistoryDuration string `json:"sSessionhistoryDuration" validate:"regexp=^(0[0-9]{1}|\\\\d{2,}):([0-5][0-9]):([0-5][0-9])$"`
	// Represent an IP address.
	SSessionhistoryIP string `json:"sSessionhistoryIP"`
	// The login name of the User.
	SUserLoginname *string "json:\"sUserLoginname,omitempty\" validate:\"regexp=^(?:([\\\\w.%+\\\\-!#$%&'*+\\/=?^`{|}~]+@[a-zA-Z0-9.-]+\\\\.[a-zA-Z]{2,20})|([a-zA-Z0-9]){1,32})$\""
}

type _SessionhistoryListElement SessionhistoryListElement

// NewSessionhistoryListElement instantiates a new SessionhistoryListElement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionhistoryListElement(pkiSessionhistoryID int32, dtSessionhistoryFirsthit string, dtSessionhistoryLasthit string, eSessionhistoryEndby FieldESessionhistoryEndby, sSessionhistoryDuration string, sSessionhistoryIP string) *SessionhistoryListElement {
	this := SessionhistoryListElement{}
	this.PkiSessionhistoryID = pkiSessionhistoryID
	this.DtSessionhistoryFirsthit = dtSessionhistoryFirsthit
	this.DtSessionhistoryLasthit = dtSessionhistoryLasthit
	this.ESessionhistoryEndby = eSessionhistoryEndby
	this.SSessionhistoryDuration = sSessionhistoryDuration
	this.SSessionhistoryIP = sSessionhistoryIP
	return &this
}

// NewSessionhistoryListElementWithDefaults instantiates a new SessionhistoryListElement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionhistoryListElementWithDefaults() *SessionhistoryListElement {
	this := SessionhistoryListElement{}
	return &this
}

// GetPkiSessionhistoryID returns the PkiSessionhistoryID field value
func (o *SessionhistoryListElement) GetPkiSessionhistoryID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiSessionhistoryID
}

// GetPkiSessionhistoryIDOk returns a tuple with the PkiSessionhistoryID field value
// and a boolean to check if the value has been set.
func (o *SessionhistoryListElement) GetPkiSessionhistoryIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiSessionhistoryID, true
}

// SetPkiSessionhistoryID sets field value
func (o *SessionhistoryListElement) SetPkiSessionhistoryID(v int32) {
	o.PkiSessionhistoryID = v
}

// GetFkiComputerID returns the FkiComputerID field value if set, zero value otherwise.
func (o *SessionhistoryListElement) GetFkiComputerID() int32 {
	if o == nil || IsNil(o.FkiComputerID) {
		var ret int32
		return ret
	}
	return *o.FkiComputerID
}

// GetFkiComputerIDOk returns a tuple with the FkiComputerID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionhistoryListElement) GetFkiComputerIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiComputerID) {
		return nil, false
	}
	return o.FkiComputerID, true
}

// HasFkiComputerID returns a boolean if a field has been set.
func (o *SessionhistoryListElement) HasFkiComputerID() bool {
	if o != nil && !IsNil(o.FkiComputerID) {
		return true
	}

	return false
}

// SetFkiComputerID gets a reference to the given int32 and assigns it to the FkiComputerID field.
func (o *SessionhistoryListElement) SetFkiComputerID(v int32) {
	o.FkiComputerID = &v
}

// GetFkiUserID returns the FkiUserID field value if set, zero value otherwise.
func (o *SessionhistoryListElement) GetFkiUserID() int32 {
	if o == nil || IsNil(o.FkiUserID) {
		var ret int32
		return ret
	}
	return *o.FkiUserID
}

// GetFkiUserIDOk returns a tuple with the FkiUserID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionhistoryListElement) GetFkiUserIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiUserID) {
		return nil, false
	}
	return o.FkiUserID, true
}

// HasFkiUserID returns a boolean if a field has been set.
func (o *SessionhistoryListElement) HasFkiUserID() bool {
	if o != nil && !IsNil(o.FkiUserID) {
		return true
	}

	return false
}

// SetFkiUserID gets a reference to the given int32 and assigns it to the FkiUserID field.
func (o *SessionhistoryListElement) SetFkiUserID(v int32) {
	o.FkiUserID = &v
}

// GetDtSessionhistoryFirsthit returns the DtSessionhistoryFirsthit field value
func (o *SessionhistoryListElement) GetDtSessionhistoryFirsthit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DtSessionhistoryFirsthit
}

// GetDtSessionhistoryFirsthitOk returns a tuple with the DtSessionhistoryFirsthit field value
// and a boolean to check if the value has been set.
func (o *SessionhistoryListElement) GetDtSessionhistoryFirsthitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DtSessionhistoryFirsthit, true
}

// SetDtSessionhistoryFirsthit sets field value
func (o *SessionhistoryListElement) SetDtSessionhistoryFirsthit(v string) {
	o.DtSessionhistoryFirsthit = v
}

// GetDtSessionhistoryLasthit returns the DtSessionhistoryLasthit field value
func (o *SessionhistoryListElement) GetDtSessionhistoryLasthit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DtSessionhistoryLasthit
}

// GetDtSessionhistoryLasthitOk returns a tuple with the DtSessionhistoryLasthit field value
// and a boolean to check if the value has been set.
func (o *SessionhistoryListElement) GetDtSessionhistoryLasthitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DtSessionhistoryLasthit, true
}

// SetDtSessionhistoryLasthit sets field value
func (o *SessionhistoryListElement) SetDtSessionhistoryLasthit(v string) {
	o.DtSessionhistoryLasthit = v
}

// GetESessionhistoryEndby returns the ESessionhistoryEndby field value
func (o *SessionhistoryListElement) GetESessionhistoryEndby() FieldESessionhistoryEndby {
	if o == nil {
		var ret FieldESessionhistoryEndby
		return ret
	}

	return o.ESessionhistoryEndby
}

// GetESessionhistoryEndbyOk returns a tuple with the ESessionhistoryEndby field value
// and a boolean to check if the value has been set.
func (o *SessionhistoryListElement) GetESessionhistoryEndbyOk() (*FieldESessionhistoryEndby, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ESessionhistoryEndby, true
}

// SetESessionhistoryEndby sets field value
func (o *SessionhistoryListElement) SetESessionhistoryEndby(v FieldESessionhistoryEndby) {
	o.ESessionhistoryEndby = v
}

// GetSComputerDescription returns the SComputerDescription field value if set, zero value otherwise.
func (o *SessionhistoryListElement) GetSComputerDescription() string {
	if o == nil || IsNil(o.SComputerDescription) {
		var ret string
		return ret
	}
	return *o.SComputerDescription
}

// GetSComputerDescriptionOk returns a tuple with the SComputerDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionhistoryListElement) GetSComputerDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.SComputerDescription) {
		return nil, false
	}
	return o.SComputerDescription, true
}

// HasSComputerDescription returns a boolean if a field has been set.
func (o *SessionhistoryListElement) HasSComputerDescription() bool {
	if o != nil && !IsNil(o.SComputerDescription) {
		return true
	}

	return false
}

// SetSComputerDescription gets a reference to the given string and assigns it to the SComputerDescription field.
func (o *SessionhistoryListElement) SetSComputerDescription(v string) {
	o.SComputerDescription = &v
}

// GetSSessionhistoryDuration returns the SSessionhistoryDuration field value
func (o *SessionhistoryListElement) GetSSessionhistoryDuration() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SSessionhistoryDuration
}

// GetSSessionhistoryDurationOk returns a tuple with the SSessionhistoryDuration field value
// and a boolean to check if the value has been set.
func (o *SessionhistoryListElement) GetSSessionhistoryDurationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SSessionhistoryDuration, true
}

// SetSSessionhistoryDuration sets field value
func (o *SessionhistoryListElement) SetSSessionhistoryDuration(v string) {
	o.SSessionhistoryDuration = v
}

// GetSSessionhistoryIP returns the SSessionhistoryIP field value
func (o *SessionhistoryListElement) GetSSessionhistoryIP() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SSessionhistoryIP
}

// GetSSessionhistoryIPOk returns a tuple with the SSessionhistoryIP field value
// and a boolean to check if the value has been set.
func (o *SessionhistoryListElement) GetSSessionhistoryIPOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SSessionhistoryIP, true
}

// SetSSessionhistoryIP sets field value
func (o *SessionhistoryListElement) SetSSessionhistoryIP(v string) {
	o.SSessionhistoryIP = v
}

// GetSUserLoginname returns the SUserLoginname field value if set, zero value otherwise.
func (o *SessionhistoryListElement) GetSUserLoginname() string {
	if o == nil || IsNil(o.SUserLoginname) {
		var ret string
		return ret
	}
	return *o.SUserLoginname
}

// GetSUserLoginnameOk returns a tuple with the SUserLoginname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionhistoryListElement) GetSUserLoginnameOk() (*string, bool) {
	if o == nil || IsNil(o.SUserLoginname) {
		return nil, false
	}
	return o.SUserLoginname, true
}

// HasSUserLoginname returns a boolean if a field has been set.
func (o *SessionhistoryListElement) HasSUserLoginname() bool {
	if o != nil && !IsNil(o.SUserLoginname) {
		return true
	}

	return false
}

// SetSUserLoginname gets a reference to the given string and assigns it to the SUserLoginname field.
func (o *SessionhistoryListElement) SetSUserLoginname(v string) {
	o.SUserLoginname = &v
}

func (o SessionhistoryListElement) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SessionhistoryListElement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiSessionhistoryID"] = o.PkiSessionhistoryID
	if !IsNil(o.FkiComputerID) {
		toSerialize["fkiComputerID"] = o.FkiComputerID
	}
	if !IsNil(o.FkiUserID) {
		toSerialize["fkiUserID"] = o.FkiUserID
	}
	toSerialize["dtSessionhistoryFirsthit"] = o.DtSessionhistoryFirsthit
	toSerialize["dtSessionhistoryLasthit"] = o.DtSessionhistoryLasthit
	toSerialize["eSessionhistoryEndby"] = o.ESessionhistoryEndby
	if !IsNil(o.SComputerDescription) {
		toSerialize["sComputerDescription"] = o.SComputerDescription
	}
	toSerialize["sSessionhistoryDuration"] = o.SSessionhistoryDuration
	toSerialize["sSessionhistoryIP"] = o.SSessionhistoryIP
	if !IsNil(o.SUserLoginname) {
		toSerialize["sUserLoginname"] = o.SUserLoginname
	}
	return toSerialize, nil
}

func (o *SessionhistoryListElement) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiSessionhistoryID",
		"dtSessionhistoryFirsthit",
		"dtSessionhistoryLasthit",
		"eSessionhistoryEndby",
		"sSessionhistoryDuration",
		"sSessionhistoryIP",
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

	varSessionhistoryListElement := _SessionhistoryListElement{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSessionhistoryListElement)

	if err != nil {
		return err
	}

	*o = SessionhistoryListElement(varSessionhistoryListElement)

	return err
}

type NullableSessionhistoryListElement struct {
	value *SessionhistoryListElement
	isSet bool
}

func (v NullableSessionhistoryListElement) Get() *SessionhistoryListElement {
	return v.value
}

func (v *NullableSessionhistoryListElement) Set(val *SessionhistoryListElement) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionhistoryListElement) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionhistoryListElement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionhistoryListElement(val *SessionhistoryListElement) *NullableSessionhistoryListElement {
	return &NullableSessionhistoryListElement{value: val, isSet: true}
}

func (v NullableSessionhistoryListElement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionhistoryListElement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


