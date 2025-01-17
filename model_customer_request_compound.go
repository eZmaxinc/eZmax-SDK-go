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

// checks if the CustomerRequestCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerRequestCompound{}

// CustomerRequestCompound A Customer Object and children
type CustomerRequestCompound struct {
	CustomerRequest
}

type _CustomerRequestCompound CustomerRequestCompound

// NewCustomerRequestCompound instantiates a new CustomerRequestCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerRequestCompound(fkiCompanyID int32, fkiCustomergroupID int32, sCustomerName string, fkiContactinformationsID int32, fkiContactcontainerID int32, fkiImageID int32, fkiGlaccountcontainerID int32, fkiLanguageID int32, fkiDepartmentID int32, fkiPaymentmethodID int32, fkiElectronicfundstransferbankaccountID int32, fkiElectronicfundstransferbankaccountIDDirectdebit int32, fkiSendingmethodID int32, fkiTaxassignmentID int32, fkiAttendancestatusID int32, fkiAgentIDVariableexpensechargeto int32, fkiBrokerIDVariableexpensechargeto int32, fkiCustomerIDVariableexpensechargeto int32, fkiGlaccountcontainerIDVariableexpensechargeto int32, fkiAgentIDSupplychargechargeto int32, fkiBrokerIDSupplychargechargeto int32, fkiCustomerIDSupplychargechargeto int32, fkiGlaccountcontainerIDSupplychargechargeto int32, fkiInvoicealternatelogoID int32, fkiSynchronizationlinkserverID int32, sCustomerCode string, dCustomerFulltimeequivalent string, iCustomerPhotocopiercode int32, iCustomerLongdistancecode int32, iCustomerTimewindowstart int32, iCustomerTimewindowend int32, dCustomerMinimumchargeableinterests string, dtCustomerBirthdate string, dtCustomerTransfer string, dtCustomerTransferappointment string, dtCustomerTransfersurvey string, bCustomerIsactive bool, bCustomerVariableexpensefinanced bool, bCustomerVariableexpensefinancedtaxes bool, bCustomerSupplychargefinanced bool, bCustomerSupplychargefinancedtaxes bool, bCustomerAttendance bool, eCustomerType FieldECustomerType, eCustomerMarketingcorrespondence FieldECustomerMarketingcorrespondence, bCustomerBlackcopycarbon bool, bCustomerUnsubscribeinfo bool, tCustomerComment string) *CustomerRequestCompound {
	this := CustomerRequestCompound{}
	this.FkiCompanyID = fkiCompanyID
	this.FkiCustomergroupID = fkiCustomergroupID
	this.SCustomerName = sCustomerName
	this.FkiContactinformationsID = fkiContactinformationsID
	this.FkiContactcontainerID = fkiContactcontainerID
	this.FkiImageID = fkiImageID
	this.FkiGlaccountcontainerID = fkiGlaccountcontainerID
	this.FkiLanguageID = fkiLanguageID
	this.FkiDepartmentID = fkiDepartmentID
	this.FkiPaymentmethodID = fkiPaymentmethodID
	this.FkiElectronicfundstransferbankaccountID = fkiElectronicfundstransferbankaccountID
	this.FkiElectronicfundstransferbankaccountIDDirectdebit = fkiElectronicfundstransferbankaccountIDDirectdebit
	this.FkiSendingmethodID = fkiSendingmethodID
	this.FkiTaxassignmentID = fkiTaxassignmentID
	this.FkiAttendancestatusID = fkiAttendancestatusID
	this.FkiAgentIDVariableexpensechargeto = fkiAgentIDVariableexpensechargeto
	this.FkiBrokerIDVariableexpensechargeto = fkiBrokerIDVariableexpensechargeto
	this.FkiCustomerIDVariableexpensechargeto = fkiCustomerIDVariableexpensechargeto
	this.FkiGlaccountcontainerIDVariableexpensechargeto = fkiGlaccountcontainerIDVariableexpensechargeto
	this.FkiAgentIDSupplychargechargeto = fkiAgentIDSupplychargechargeto
	this.FkiBrokerIDSupplychargechargeto = fkiBrokerIDSupplychargechargeto
	this.FkiCustomerIDSupplychargechargeto = fkiCustomerIDSupplychargechargeto
	this.FkiGlaccountcontainerIDSupplychargechargeto = fkiGlaccountcontainerIDSupplychargechargeto
	this.FkiInvoicealternatelogoID = fkiInvoicealternatelogoID
	this.FkiSynchronizationlinkserverID = fkiSynchronizationlinkserverID
	this.SCustomerCode = sCustomerCode
	this.DCustomerFulltimeequivalent = dCustomerFulltimeequivalent
	this.ICustomerPhotocopiercode = iCustomerPhotocopiercode
	this.ICustomerLongdistancecode = iCustomerLongdistancecode
	this.ICustomerTimewindowstart = iCustomerTimewindowstart
	this.ICustomerTimewindowend = iCustomerTimewindowend
	this.DCustomerMinimumchargeableinterests = dCustomerMinimumchargeableinterests
	this.DtCustomerBirthdate = dtCustomerBirthdate
	this.DtCustomerTransfer = dtCustomerTransfer
	this.DtCustomerTransferappointment = dtCustomerTransferappointment
	this.DtCustomerTransfersurvey = dtCustomerTransfersurvey
	this.BCustomerIsactive = bCustomerIsactive
	this.BCustomerVariableexpensefinanced = bCustomerVariableexpensefinanced
	this.BCustomerVariableexpensefinancedtaxes = bCustomerVariableexpensefinancedtaxes
	this.BCustomerSupplychargefinanced = bCustomerSupplychargefinanced
	this.BCustomerSupplychargefinancedtaxes = bCustomerSupplychargefinancedtaxes
	this.BCustomerAttendance = bCustomerAttendance
	this.ECustomerType = eCustomerType
	this.ECustomerMarketingcorrespondence = eCustomerMarketingcorrespondence
	this.BCustomerBlackcopycarbon = bCustomerBlackcopycarbon
	this.BCustomerUnsubscribeinfo = bCustomerUnsubscribeinfo
	this.TCustomerComment = tCustomerComment
	return &this
}

// NewCustomerRequestCompoundWithDefaults instantiates a new CustomerRequestCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerRequestCompoundWithDefaults() *CustomerRequestCompound {
	this := CustomerRequestCompound{}
	return &this
}

func (o CustomerRequestCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerRequestCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

func (o *CustomerRequestCompound) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fkiCompanyID",
		"fkiCustomergroupID",
		"sCustomerName",
		"fkiContactinformationsID",
		"fkiContactcontainerID",
		"fkiImageID",
		"fkiGlaccountcontainerID",
		"fkiLanguageID",
		"fkiDepartmentID",
		"fkiPaymentmethodID",
		"fkiElectronicfundstransferbankaccountID",
		"fkiElectronicfundstransferbankaccountIDDirectdebit",
		"fkiSendingmethodID",
		"fkiTaxassignmentID",
		"fkiAttendancestatusID",
		"fkiAgentIDVariableexpensechargeto",
		"fkiBrokerIDVariableexpensechargeto",
		"fkiCustomerIDVariableexpensechargeto",
		"fkiGlaccountcontainerIDVariableexpensechargeto",
		"fkiAgentIDSupplychargechargeto",
		"fkiBrokerIDSupplychargechargeto",
		"fkiCustomerIDSupplychargechargeto",
		"fkiGlaccountcontainerIDSupplychargechargeto",
		"fkiInvoicealternatelogoID",
		"fkiSynchronizationlinkserverID",
		"sCustomerCode",
		"dCustomerFulltimeequivalent",
		"iCustomerPhotocopiercode",
		"iCustomerLongdistancecode",
		"iCustomerTimewindowstart",
		"iCustomerTimewindowend",
		"dCustomerMinimumchargeableinterests",
		"dtCustomerBirthdate",
		"dtCustomerTransfer",
		"dtCustomerTransferappointment",
		"dtCustomerTransfersurvey",
		"bCustomerIsactive",
		"bCustomerVariableexpensefinanced",
		"bCustomerVariableexpensefinancedtaxes",
		"bCustomerSupplychargefinanced",
		"bCustomerSupplychargefinancedtaxes",
		"bCustomerAttendance",
		"eCustomerType",
		"eCustomerMarketingcorrespondence",
		"bCustomerBlackcopycarbon",
		"bCustomerUnsubscribeinfo",
		"tCustomerComment",
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

	varCustomerRequestCompound := _CustomerRequestCompound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCustomerRequestCompound)

	if err != nil {
		return err
	}

	*o = CustomerRequestCompound(varCustomerRequestCompound)

	return err
}

type NullableCustomerRequestCompound struct {
	value *CustomerRequestCompound
	isSet bool
}

func (v NullableCustomerRequestCompound) Get() *CustomerRequestCompound {
	return v.value
}

func (v *NullableCustomerRequestCompound) Set(val *CustomerRequestCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerRequestCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerRequestCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerRequestCompound(val *CustomerRequestCompound) *NullableCustomerRequestCompound {
	return &NullableCustomerRequestCompound{value: val, isSet: true}
}

func (v NullableCustomerRequestCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerRequestCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


