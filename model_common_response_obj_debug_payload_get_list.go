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

// checks if the CommonResponseObjDebugPayloadGetList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommonResponseObjDebugPayloadGetList{}

// CommonResponseObjDebugPayloadGetList This is a debug object containing debugging information on the actual function
type CommonResponseObjDebugPayloadGetList struct {
	// The minimum version of the function that can be called
	IVersionMin int32 `json:"iVersionMin"`
	// The maximum version of the function that can be called
	IVersionMax int32 `json:"iVersionMax"`
	// An array of permissions required to access this function.  If the value \"0\" is present in the array, anyone can call this function.  You must have one of the permission to access the function. You don't need to have all of them.
	ARequiredPermission []int32 `json:"a_RequiredPermission"`
	// Wheter the current route is deprecated or not
	BVersionDeprecated bool `json:"bVersionDeprecated"`
	// Represent a Date Time. The timezone is the one configured in the User's profile.
	DtResponseDate string `json:"dtResponseDate"`
	AFilter CommonResponseFilter `json:"a_Filter"`
	// List of available values for *eOrderBy*
	AOrderBy map[string]string `json:"a_OrderBy"`
	// The maximum numbers of results to be returned.  When the content-type is **application/json** there is an implicit default of 10 000.  When it's **application/vnd.openxmlformats-officedocument.spreadsheetml.sheet** the is no implicit default so if you do not specify iRowMax, all records will be returned.
	IRowMax int32 `json:"iRowMax"`
	// The starting element from where to start retrieving the results. For example if you started at iRowOffset=0 and asked for iRowMax=100, to get the next 100 results, you could specify iRowOffset=100&iRowMax=100,
	IRowOffset int32 `json:"iRowOffset"`
}

type _CommonResponseObjDebugPayloadGetList CommonResponseObjDebugPayloadGetList

// NewCommonResponseObjDebugPayloadGetList instantiates a new CommonResponseObjDebugPayloadGetList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonResponseObjDebugPayloadGetList(iVersionMin int32, iVersionMax int32, aRequiredPermission []int32, bVersionDeprecated bool, dtResponseDate string, aFilter CommonResponseFilter, aOrderBy map[string]string, iRowMax int32, iRowOffset int32) *CommonResponseObjDebugPayloadGetList {
	this := CommonResponseObjDebugPayloadGetList{}
	this.IVersionMin = iVersionMin
	this.IVersionMax = iVersionMax
	this.ARequiredPermission = aRequiredPermission
	this.BVersionDeprecated = bVersionDeprecated
	this.DtResponseDate = dtResponseDate
	this.AFilter = aFilter
	this.AOrderBy = aOrderBy
	this.IRowMax = iRowMax
	this.IRowOffset = iRowOffset
	return &this
}

// NewCommonResponseObjDebugPayloadGetListWithDefaults instantiates a new CommonResponseObjDebugPayloadGetList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonResponseObjDebugPayloadGetListWithDefaults() *CommonResponseObjDebugPayloadGetList {
	this := CommonResponseObjDebugPayloadGetList{}
	var iRowOffset int32 = 0
	this.IRowOffset = iRowOffset
	return &this
}

// GetIVersionMin returns the IVersionMin field value
func (o *CommonResponseObjDebugPayloadGetList) GetIVersionMin() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IVersionMin
}

// GetIVersionMinOk returns a tuple with the IVersionMin field value
// and a boolean to check if the value has been set.
func (o *CommonResponseObjDebugPayloadGetList) GetIVersionMinOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IVersionMin, true
}

// SetIVersionMin sets field value
func (o *CommonResponseObjDebugPayloadGetList) SetIVersionMin(v int32) {
	o.IVersionMin = v
}

// GetIVersionMax returns the IVersionMax field value
func (o *CommonResponseObjDebugPayloadGetList) GetIVersionMax() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IVersionMax
}

// GetIVersionMaxOk returns a tuple with the IVersionMax field value
// and a boolean to check if the value has been set.
func (o *CommonResponseObjDebugPayloadGetList) GetIVersionMaxOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IVersionMax, true
}

// SetIVersionMax sets field value
func (o *CommonResponseObjDebugPayloadGetList) SetIVersionMax(v int32) {
	o.IVersionMax = v
}

// GetARequiredPermission returns the ARequiredPermission field value
func (o *CommonResponseObjDebugPayloadGetList) GetARequiredPermission() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.ARequiredPermission
}

// GetARequiredPermissionOk returns a tuple with the ARequiredPermission field value
// and a boolean to check if the value has been set.
func (o *CommonResponseObjDebugPayloadGetList) GetARequiredPermissionOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ARequiredPermission, true
}

// SetARequiredPermission sets field value
func (o *CommonResponseObjDebugPayloadGetList) SetARequiredPermission(v []int32) {
	o.ARequiredPermission = v
}

// GetBVersionDeprecated returns the BVersionDeprecated field value
func (o *CommonResponseObjDebugPayloadGetList) GetBVersionDeprecated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BVersionDeprecated
}

// GetBVersionDeprecatedOk returns a tuple with the BVersionDeprecated field value
// and a boolean to check if the value has been set.
func (o *CommonResponseObjDebugPayloadGetList) GetBVersionDeprecatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BVersionDeprecated, true
}

// SetBVersionDeprecated sets field value
func (o *CommonResponseObjDebugPayloadGetList) SetBVersionDeprecated(v bool) {
	o.BVersionDeprecated = v
}

// GetDtResponseDate returns the DtResponseDate field value
func (o *CommonResponseObjDebugPayloadGetList) GetDtResponseDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DtResponseDate
}

// GetDtResponseDateOk returns a tuple with the DtResponseDate field value
// and a boolean to check if the value has been set.
func (o *CommonResponseObjDebugPayloadGetList) GetDtResponseDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DtResponseDate, true
}

// SetDtResponseDate sets field value
func (o *CommonResponseObjDebugPayloadGetList) SetDtResponseDate(v string) {
	o.DtResponseDate = v
}

// GetAFilter returns the AFilter field value
func (o *CommonResponseObjDebugPayloadGetList) GetAFilter() CommonResponseFilter {
	if o == nil {
		var ret CommonResponseFilter
		return ret
	}

	return o.AFilter
}

// GetAFilterOk returns a tuple with the AFilter field value
// and a boolean to check if the value has been set.
func (o *CommonResponseObjDebugPayloadGetList) GetAFilterOk() (*CommonResponseFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AFilter, true
}

// SetAFilter sets field value
func (o *CommonResponseObjDebugPayloadGetList) SetAFilter(v CommonResponseFilter) {
	o.AFilter = v
}

// GetAOrderBy returns the AOrderBy field value
func (o *CommonResponseObjDebugPayloadGetList) GetAOrderBy() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.AOrderBy
}

// GetAOrderByOk returns a tuple with the AOrderBy field value
// and a boolean to check if the value has been set.
func (o *CommonResponseObjDebugPayloadGetList) GetAOrderByOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AOrderBy, true
}

// SetAOrderBy sets field value
func (o *CommonResponseObjDebugPayloadGetList) SetAOrderBy(v map[string]string) {
	o.AOrderBy = v
}

// GetIRowMax returns the IRowMax field value
func (o *CommonResponseObjDebugPayloadGetList) GetIRowMax() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IRowMax
}

// GetIRowMaxOk returns a tuple with the IRowMax field value
// and a boolean to check if the value has been set.
func (o *CommonResponseObjDebugPayloadGetList) GetIRowMaxOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IRowMax, true
}

// SetIRowMax sets field value
func (o *CommonResponseObjDebugPayloadGetList) SetIRowMax(v int32) {
	o.IRowMax = v
}

// GetIRowOffset returns the IRowOffset field value
func (o *CommonResponseObjDebugPayloadGetList) GetIRowOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IRowOffset
}

// GetIRowOffsetOk returns a tuple with the IRowOffset field value
// and a boolean to check if the value has been set.
func (o *CommonResponseObjDebugPayloadGetList) GetIRowOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IRowOffset, true
}

// SetIRowOffset sets field value
func (o *CommonResponseObjDebugPayloadGetList) SetIRowOffset(v int32) {
	o.IRowOffset = v
}

func (o CommonResponseObjDebugPayloadGetList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommonResponseObjDebugPayloadGetList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["iVersionMin"] = o.IVersionMin
	toSerialize["iVersionMax"] = o.IVersionMax
	toSerialize["a_RequiredPermission"] = o.ARequiredPermission
	toSerialize["bVersionDeprecated"] = o.BVersionDeprecated
	toSerialize["dtResponseDate"] = o.DtResponseDate
	toSerialize["a_Filter"] = o.AFilter
	toSerialize["a_OrderBy"] = o.AOrderBy
	toSerialize["iRowMax"] = o.IRowMax
	toSerialize["iRowOffset"] = o.IRowOffset
	return toSerialize, nil
}

func (o *CommonResponseObjDebugPayloadGetList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"iVersionMin",
		"iVersionMax",
		"a_RequiredPermission",
		"bVersionDeprecated",
		"dtResponseDate",
		"a_Filter",
		"a_OrderBy",
		"iRowMax",
		"iRowOffset",
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

	varCommonResponseObjDebugPayloadGetList := _CommonResponseObjDebugPayloadGetList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCommonResponseObjDebugPayloadGetList)

	if err != nil {
		return err
	}

	*o = CommonResponseObjDebugPayloadGetList(varCommonResponseObjDebugPayloadGetList)

	return err
}

type NullableCommonResponseObjDebugPayloadGetList struct {
	value *CommonResponseObjDebugPayloadGetList
	isSet bool
}

func (v NullableCommonResponseObjDebugPayloadGetList) Get() *CommonResponseObjDebugPayloadGetList {
	return v.value
}

func (v *NullableCommonResponseObjDebugPayloadGetList) Set(val *CommonResponseObjDebugPayloadGetList) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonResponseObjDebugPayloadGetList) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonResponseObjDebugPayloadGetList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonResponseObjDebugPayloadGetList(val *CommonResponseObjDebugPayloadGetList) *NullableCommonResponseObjDebugPayloadGetList {
	return &NullableCommonResponseObjDebugPayloadGetList{value: val, isSet: true}
}

func (v NullableCommonResponseObjDebugPayloadGetList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonResponseObjDebugPayloadGetList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


