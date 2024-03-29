/*
eZmax API Definition

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.1.3
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
)

// ListpresentationRequest A Listpresentation element
type ListpresentationRequest struct {
	// A descriptive for the list presentation
	SListpresentationDescription string `json:"sListpresentationDescription"`
	// The filter to apply to the request to limit results.
	SListpresentationFilter string `json:"sListpresentationFilter"`
	// The order by the user chose
	SListpresentationOrderby string `json:"sListpresentationOrderby"`
	// An array of column names that the user chose to bee visible
	ASColumnName []string `json:"a_sColumnName"`
	// The maximum numbers of results to be returned
	IListpresentationRowMax int32 `json:"iListpresentationRowMax"`
	// The starting element from where to start retrieving the results. For example if you started at iRowOffset=0 and asked for iRowMax=100, to get the next 100 results, you could specify iRowOffset=100&iRowMax=100,
	IListpresentationRowOffset int32 `json:"iListpresentationRowOffset"`
	// Set to true if the user chose this Listpresentation as the default one. A single element should be set to true
	BListpresentationDefault bool `json:"bListpresentationDefault"`
}

// NewListpresentationRequest instantiates a new ListpresentationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListpresentationRequest(sListpresentationDescription string, sListpresentationFilter string, sListpresentationOrderby string, aSColumnName []string, iListpresentationRowMax int32, iListpresentationRowOffset int32, bListpresentationDefault bool) *ListpresentationRequest {
	this := ListpresentationRequest{}
	this.SListpresentationDescription = sListpresentationDescription
	this.SListpresentationFilter = sListpresentationFilter
	this.SListpresentationOrderby = sListpresentationOrderby
	this.ASColumnName = aSColumnName
	this.IListpresentationRowMax = iListpresentationRowMax
	this.IListpresentationRowOffset = iListpresentationRowOffset
	this.BListpresentationDefault = bListpresentationDefault
	return &this
}

// NewListpresentationRequestWithDefaults instantiates a new ListpresentationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListpresentationRequestWithDefaults() *ListpresentationRequest {
	this := ListpresentationRequest{}
	return &this
}

// GetSListpresentationDescription returns the SListpresentationDescription field value
func (o *ListpresentationRequest) GetSListpresentationDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SListpresentationDescription
}

// GetSListpresentationDescriptionOk returns a tuple with the SListpresentationDescription field value
// and a boolean to check if the value has been set.
func (o *ListpresentationRequest) GetSListpresentationDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SListpresentationDescription, true
}

// SetSListpresentationDescription sets field value
func (o *ListpresentationRequest) SetSListpresentationDescription(v string) {
	o.SListpresentationDescription = v
}

// GetSListpresentationFilter returns the SListpresentationFilter field value
func (o *ListpresentationRequest) GetSListpresentationFilter() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SListpresentationFilter
}

// GetSListpresentationFilterOk returns a tuple with the SListpresentationFilter field value
// and a boolean to check if the value has been set.
func (o *ListpresentationRequest) GetSListpresentationFilterOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SListpresentationFilter, true
}

// SetSListpresentationFilter sets field value
func (o *ListpresentationRequest) SetSListpresentationFilter(v string) {
	o.SListpresentationFilter = v
}

// GetSListpresentationOrderby returns the SListpresentationOrderby field value
func (o *ListpresentationRequest) GetSListpresentationOrderby() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SListpresentationOrderby
}

// GetSListpresentationOrderbyOk returns a tuple with the SListpresentationOrderby field value
// and a boolean to check if the value has been set.
func (o *ListpresentationRequest) GetSListpresentationOrderbyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SListpresentationOrderby, true
}

// SetSListpresentationOrderby sets field value
func (o *ListpresentationRequest) SetSListpresentationOrderby(v string) {
	o.SListpresentationOrderby = v
}

// GetASColumnName returns the ASColumnName field value
func (o *ListpresentationRequest) GetASColumnName() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ASColumnName
}

// GetASColumnNameOk returns a tuple with the ASColumnName field value
// and a boolean to check if the value has been set.
func (o *ListpresentationRequest) GetASColumnNameOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ASColumnName, true
}

// SetASColumnName sets field value
func (o *ListpresentationRequest) SetASColumnName(v []string) {
	o.ASColumnName = v
}

// GetIListpresentationRowMax returns the IListpresentationRowMax field value
func (o *ListpresentationRequest) GetIListpresentationRowMax() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IListpresentationRowMax
}

// GetIListpresentationRowMaxOk returns a tuple with the IListpresentationRowMax field value
// and a boolean to check if the value has been set.
func (o *ListpresentationRequest) GetIListpresentationRowMaxOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IListpresentationRowMax, true
}

// SetIListpresentationRowMax sets field value
func (o *ListpresentationRequest) SetIListpresentationRowMax(v int32) {
	o.IListpresentationRowMax = v
}

// GetIListpresentationRowOffset returns the IListpresentationRowOffset field value
func (o *ListpresentationRequest) GetIListpresentationRowOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IListpresentationRowOffset
}

// GetIListpresentationRowOffsetOk returns a tuple with the IListpresentationRowOffset field value
// and a boolean to check if the value has been set.
func (o *ListpresentationRequest) GetIListpresentationRowOffsetOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IListpresentationRowOffset, true
}

// SetIListpresentationRowOffset sets field value
func (o *ListpresentationRequest) SetIListpresentationRowOffset(v int32) {
	o.IListpresentationRowOffset = v
}

// GetBListpresentationDefault returns the BListpresentationDefault field value
func (o *ListpresentationRequest) GetBListpresentationDefault() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BListpresentationDefault
}

// GetBListpresentationDefaultOk returns a tuple with the BListpresentationDefault field value
// and a boolean to check if the value has been set.
func (o *ListpresentationRequest) GetBListpresentationDefaultOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BListpresentationDefault, true
}

// SetBListpresentationDefault sets field value
func (o *ListpresentationRequest) SetBListpresentationDefault(v bool) {
	o.BListpresentationDefault = v
}

func (o ListpresentationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sListpresentationDescription"] = o.SListpresentationDescription
	}
	if true {
		toSerialize["sListpresentationFilter"] = o.SListpresentationFilter
	}
	if true {
		toSerialize["sListpresentationOrderby"] = o.SListpresentationOrderby
	}
	if true {
		toSerialize["a_sColumnName"] = o.ASColumnName
	}
	if true {
		toSerialize["iListpresentationRowMax"] = o.IListpresentationRowMax
	}
	if true {
		toSerialize["iListpresentationRowOffset"] = o.IListpresentationRowOffset
	}
	if true {
		toSerialize["bListpresentationDefault"] = o.BListpresentationDefault
	}
	return json.Marshal(toSerialize)
}

type NullableListpresentationRequest struct {
	value *ListpresentationRequest
	isSet bool
}

func (v NullableListpresentationRequest) Get() *ListpresentationRequest {
	return v.value
}

func (v *NullableListpresentationRequest) Set(val *ListpresentationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableListpresentationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableListpresentationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListpresentationRequest(val *ListpresentationRequest) *NullableListpresentationRequest {
	return &NullableListpresentationRequest{value: val, isSet: true}
}

func (v NullableListpresentationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListpresentationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


