# BrokerListElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiBrokerID** | **int32** | The unique ID of the Broker. | 
**FkiDepartmentID** | **int32** | The unique ID of the Department | 
**SDepartmentNameX** | Pointer to **string** | The Name of the Department in the language of the requester | [optional] 
**FkiBrokertypeID** | **int32** | The unique ID of the Brokertype | 
**SBrokertypeNameX** | **string** | The name of the Brokertype in the language of the requester | 
**SBrokerCode** | **string** | The code of the Broker | 
**SRealestateboardnumberNumber** | Pointer to **string** | The number of the Realestateboardnumber | [optional] 
**IAgentBannernumber** | Pointer to **int32** | The bannernumber of the Agent | [optional] 
**SLanguageNameX** | Pointer to **string** | The Name of the Language in the language of the requester | [optional] 
**IBrokerPhotocopiercode** | **int32** | The photocopiercode of the Broker | 
**IBrokerLongdistancecode** | **int32** | The longdistancecode of the Broker | 
**SBrokerName** | **string** | The name of the Broker | 
**SBrokerRealestateassociationlicense** | **string** | The realestateassociationlicense of the Broker | 
**DtBrokerHiredate** | **string** | The hiredate of the Broker | 
**DtBrokerLeavedate** | Pointer to **string** | The leavedate of the Broker | [optional] 
**BBrokerTranquillit** | Pointer to **bool** | Whether if it&#39;s an tranquillit | [optional] 
**BBrokerResidentiallicense** | **bool** | Whether if it&#39;s an residentiallicense | 
**BBrokerCommerciallicense** | **bool** | Whether if it&#39;s an commerciallicense | 
**BBrokerMortgagelicense** | **bool** | Whether if it&#39;s an mortgagelicense | 
**BBrokerPaidbyofficetranquillit** | **bool** | Whether if it&#39;s an paidbyofficetranquillit | 
**DtBrokerFintraccertification** | Pointer to **string** | The fintraccertification of the Broker | [optional] 
**BBrokerIsactive** | **bool** | Whether the Broker is active or not | 
**SContactFirstname** | Pointer to **string** | The First name of the contact | [optional] 
**SContactLastname** | Pointer to **string** | The Last name of the contact | [optional] 
**DtContactBirthdate** | Pointer to **string** | The Birth Date of the contact | [optional] 
**SEmailAddress** | Pointer to **string** | The email address. | [optional] 
**SPhoneE164** | Pointer to **string** | A phone number in E.164 Format | [optional] 
**SAddressCivic** | Pointer to **string** | The Civic number. | [optional] 
**SAddressStreet** | Pointer to **string** | The Street Name | [optional] 
**SAddressSuite** | Pointer to **string** | The Suite or appartment number | [optional] 
**SAddressCity** | Pointer to **string** | The City name | [optional] 
**SAddressZip** | Pointer to **string** | The Postal/Zip Code  The value must be entered without spaces | [optional] 
**SProvinceNameX** | Pointer to **string** | The name of the Province in the language of the requester | [optional] 
**SCountryNameX** | Pointer to **string** | The name of the Country in the language of the requester | [optional] 

## Methods

### NewBrokerListElement

`func NewBrokerListElement(pkiBrokerID int32, fkiDepartmentID int32, fkiBrokertypeID int32, sBrokertypeNameX string, sBrokerCode string, iBrokerPhotocopiercode int32, iBrokerLongdistancecode int32, sBrokerName string, sBrokerRealestateassociationlicense string, dtBrokerHiredate string, bBrokerResidentiallicense bool, bBrokerCommerciallicense bool, bBrokerMortgagelicense bool, bBrokerPaidbyofficetranquillit bool, bBrokerIsactive bool, ) *BrokerListElement`

NewBrokerListElement instantiates a new BrokerListElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrokerListElementWithDefaults

`func NewBrokerListElementWithDefaults() *BrokerListElement`

NewBrokerListElementWithDefaults instantiates a new BrokerListElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiBrokerID

`func (o *BrokerListElement) GetPkiBrokerID() int32`

GetPkiBrokerID returns the PkiBrokerID field if non-nil, zero value otherwise.

### GetPkiBrokerIDOk

`func (o *BrokerListElement) GetPkiBrokerIDOk() (*int32, bool)`

GetPkiBrokerIDOk returns a tuple with the PkiBrokerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiBrokerID

`func (o *BrokerListElement) SetPkiBrokerID(v int32)`

SetPkiBrokerID sets PkiBrokerID field to given value.


### GetFkiDepartmentID

`func (o *BrokerListElement) GetFkiDepartmentID() int32`

GetFkiDepartmentID returns the FkiDepartmentID field if non-nil, zero value otherwise.

### GetFkiDepartmentIDOk

`func (o *BrokerListElement) GetFkiDepartmentIDOk() (*int32, bool)`

GetFkiDepartmentIDOk returns a tuple with the FkiDepartmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiDepartmentID

`func (o *BrokerListElement) SetFkiDepartmentID(v int32)`

SetFkiDepartmentID sets FkiDepartmentID field to given value.


### GetSDepartmentNameX

`func (o *BrokerListElement) GetSDepartmentNameX() string`

GetSDepartmentNameX returns the SDepartmentNameX field if non-nil, zero value otherwise.

### GetSDepartmentNameXOk

`func (o *BrokerListElement) GetSDepartmentNameXOk() (*string, bool)`

GetSDepartmentNameXOk returns a tuple with the SDepartmentNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSDepartmentNameX

`func (o *BrokerListElement) SetSDepartmentNameX(v string)`

SetSDepartmentNameX sets SDepartmentNameX field to given value.

### HasSDepartmentNameX

`func (o *BrokerListElement) HasSDepartmentNameX() bool`

HasSDepartmentNameX returns a boolean if a field has been set.

### GetFkiBrokertypeID

`func (o *BrokerListElement) GetFkiBrokertypeID() int32`

GetFkiBrokertypeID returns the FkiBrokertypeID field if non-nil, zero value otherwise.

### GetFkiBrokertypeIDOk

`func (o *BrokerListElement) GetFkiBrokertypeIDOk() (*int32, bool)`

GetFkiBrokertypeIDOk returns a tuple with the FkiBrokertypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiBrokertypeID

`func (o *BrokerListElement) SetFkiBrokertypeID(v int32)`

SetFkiBrokertypeID sets FkiBrokertypeID field to given value.


### GetSBrokertypeNameX

`func (o *BrokerListElement) GetSBrokertypeNameX() string`

GetSBrokertypeNameX returns the SBrokertypeNameX field if non-nil, zero value otherwise.

### GetSBrokertypeNameXOk

`func (o *BrokerListElement) GetSBrokertypeNameXOk() (*string, bool)`

GetSBrokertypeNameXOk returns a tuple with the SBrokertypeNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSBrokertypeNameX

`func (o *BrokerListElement) SetSBrokertypeNameX(v string)`

SetSBrokertypeNameX sets SBrokertypeNameX field to given value.


### GetSBrokerCode

`func (o *BrokerListElement) GetSBrokerCode() string`

GetSBrokerCode returns the SBrokerCode field if non-nil, zero value otherwise.

### GetSBrokerCodeOk

`func (o *BrokerListElement) GetSBrokerCodeOk() (*string, bool)`

GetSBrokerCodeOk returns a tuple with the SBrokerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSBrokerCode

`func (o *BrokerListElement) SetSBrokerCode(v string)`

SetSBrokerCode sets SBrokerCode field to given value.


### GetSRealestateboardnumberNumber

`func (o *BrokerListElement) GetSRealestateboardnumberNumber() string`

GetSRealestateboardnumberNumber returns the SRealestateboardnumberNumber field if non-nil, zero value otherwise.

### GetSRealestateboardnumberNumberOk

`func (o *BrokerListElement) GetSRealestateboardnumberNumberOk() (*string, bool)`

GetSRealestateboardnumberNumberOk returns a tuple with the SRealestateboardnumberNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSRealestateboardnumberNumber

`func (o *BrokerListElement) SetSRealestateboardnumberNumber(v string)`

SetSRealestateboardnumberNumber sets SRealestateboardnumberNumber field to given value.

### HasSRealestateboardnumberNumber

`func (o *BrokerListElement) HasSRealestateboardnumberNumber() bool`

HasSRealestateboardnumberNumber returns a boolean if a field has been set.

### GetIAgentBannernumber

`func (o *BrokerListElement) GetIAgentBannernumber() int32`

GetIAgentBannernumber returns the IAgentBannernumber field if non-nil, zero value otherwise.

### GetIAgentBannernumberOk

`func (o *BrokerListElement) GetIAgentBannernumberOk() (*int32, bool)`

GetIAgentBannernumberOk returns a tuple with the IAgentBannernumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIAgentBannernumber

`func (o *BrokerListElement) SetIAgentBannernumber(v int32)`

SetIAgentBannernumber sets IAgentBannernumber field to given value.

### HasIAgentBannernumber

`func (o *BrokerListElement) HasIAgentBannernumber() bool`

HasIAgentBannernumber returns a boolean if a field has been set.

### GetSLanguageNameX

`func (o *BrokerListElement) GetSLanguageNameX() string`

GetSLanguageNameX returns the SLanguageNameX field if non-nil, zero value otherwise.

### GetSLanguageNameXOk

`func (o *BrokerListElement) GetSLanguageNameXOk() (*string, bool)`

GetSLanguageNameXOk returns a tuple with the SLanguageNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSLanguageNameX

`func (o *BrokerListElement) SetSLanguageNameX(v string)`

SetSLanguageNameX sets SLanguageNameX field to given value.

### HasSLanguageNameX

`func (o *BrokerListElement) HasSLanguageNameX() bool`

HasSLanguageNameX returns a boolean if a field has been set.

### GetIBrokerPhotocopiercode

`func (o *BrokerListElement) GetIBrokerPhotocopiercode() int32`

GetIBrokerPhotocopiercode returns the IBrokerPhotocopiercode field if non-nil, zero value otherwise.

### GetIBrokerPhotocopiercodeOk

`func (o *BrokerListElement) GetIBrokerPhotocopiercodeOk() (*int32, bool)`

GetIBrokerPhotocopiercodeOk returns a tuple with the IBrokerPhotocopiercode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIBrokerPhotocopiercode

`func (o *BrokerListElement) SetIBrokerPhotocopiercode(v int32)`

SetIBrokerPhotocopiercode sets IBrokerPhotocopiercode field to given value.


### GetIBrokerLongdistancecode

`func (o *BrokerListElement) GetIBrokerLongdistancecode() int32`

GetIBrokerLongdistancecode returns the IBrokerLongdistancecode field if non-nil, zero value otherwise.

### GetIBrokerLongdistancecodeOk

`func (o *BrokerListElement) GetIBrokerLongdistancecodeOk() (*int32, bool)`

GetIBrokerLongdistancecodeOk returns a tuple with the IBrokerLongdistancecode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIBrokerLongdistancecode

`func (o *BrokerListElement) SetIBrokerLongdistancecode(v int32)`

SetIBrokerLongdistancecode sets IBrokerLongdistancecode field to given value.


### GetSBrokerName

`func (o *BrokerListElement) GetSBrokerName() string`

GetSBrokerName returns the SBrokerName field if non-nil, zero value otherwise.

### GetSBrokerNameOk

`func (o *BrokerListElement) GetSBrokerNameOk() (*string, bool)`

GetSBrokerNameOk returns a tuple with the SBrokerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSBrokerName

`func (o *BrokerListElement) SetSBrokerName(v string)`

SetSBrokerName sets SBrokerName field to given value.


### GetSBrokerRealestateassociationlicense

`func (o *BrokerListElement) GetSBrokerRealestateassociationlicense() string`

GetSBrokerRealestateassociationlicense returns the SBrokerRealestateassociationlicense field if non-nil, zero value otherwise.

### GetSBrokerRealestateassociationlicenseOk

`func (o *BrokerListElement) GetSBrokerRealestateassociationlicenseOk() (*string, bool)`

GetSBrokerRealestateassociationlicenseOk returns a tuple with the SBrokerRealestateassociationlicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSBrokerRealestateassociationlicense

`func (o *BrokerListElement) SetSBrokerRealestateassociationlicense(v string)`

SetSBrokerRealestateassociationlicense sets SBrokerRealestateassociationlicense field to given value.


### GetDtBrokerHiredate

`func (o *BrokerListElement) GetDtBrokerHiredate() string`

GetDtBrokerHiredate returns the DtBrokerHiredate field if non-nil, zero value otherwise.

### GetDtBrokerHiredateOk

`func (o *BrokerListElement) GetDtBrokerHiredateOk() (*string, bool)`

GetDtBrokerHiredateOk returns a tuple with the DtBrokerHiredate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtBrokerHiredate

`func (o *BrokerListElement) SetDtBrokerHiredate(v string)`

SetDtBrokerHiredate sets DtBrokerHiredate field to given value.


### GetDtBrokerLeavedate

`func (o *BrokerListElement) GetDtBrokerLeavedate() string`

GetDtBrokerLeavedate returns the DtBrokerLeavedate field if non-nil, zero value otherwise.

### GetDtBrokerLeavedateOk

`func (o *BrokerListElement) GetDtBrokerLeavedateOk() (*string, bool)`

GetDtBrokerLeavedateOk returns a tuple with the DtBrokerLeavedate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtBrokerLeavedate

`func (o *BrokerListElement) SetDtBrokerLeavedate(v string)`

SetDtBrokerLeavedate sets DtBrokerLeavedate field to given value.

### HasDtBrokerLeavedate

`func (o *BrokerListElement) HasDtBrokerLeavedate() bool`

HasDtBrokerLeavedate returns a boolean if a field has been set.

### GetBBrokerTranquillit

`func (o *BrokerListElement) GetBBrokerTranquillit() bool`

GetBBrokerTranquillit returns the BBrokerTranquillit field if non-nil, zero value otherwise.

### GetBBrokerTranquillitOk

`func (o *BrokerListElement) GetBBrokerTranquillitOk() (*bool, bool)`

GetBBrokerTranquillitOk returns a tuple with the BBrokerTranquillit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBBrokerTranquillit

`func (o *BrokerListElement) SetBBrokerTranquillit(v bool)`

SetBBrokerTranquillit sets BBrokerTranquillit field to given value.

### HasBBrokerTranquillit

`func (o *BrokerListElement) HasBBrokerTranquillit() bool`

HasBBrokerTranquillit returns a boolean if a field has been set.

### GetBBrokerResidentiallicense

`func (o *BrokerListElement) GetBBrokerResidentiallicense() bool`

GetBBrokerResidentiallicense returns the BBrokerResidentiallicense field if non-nil, zero value otherwise.

### GetBBrokerResidentiallicenseOk

`func (o *BrokerListElement) GetBBrokerResidentiallicenseOk() (*bool, bool)`

GetBBrokerResidentiallicenseOk returns a tuple with the BBrokerResidentiallicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBBrokerResidentiallicense

`func (o *BrokerListElement) SetBBrokerResidentiallicense(v bool)`

SetBBrokerResidentiallicense sets BBrokerResidentiallicense field to given value.


### GetBBrokerCommerciallicense

`func (o *BrokerListElement) GetBBrokerCommerciallicense() bool`

GetBBrokerCommerciallicense returns the BBrokerCommerciallicense field if non-nil, zero value otherwise.

### GetBBrokerCommerciallicenseOk

`func (o *BrokerListElement) GetBBrokerCommerciallicenseOk() (*bool, bool)`

GetBBrokerCommerciallicenseOk returns a tuple with the BBrokerCommerciallicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBBrokerCommerciallicense

`func (o *BrokerListElement) SetBBrokerCommerciallicense(v bool)`

SetBBrokerCommerciallicense sets BBrokerCommerciallicense field to given value.


### GetBBrokerMortgagelicense

`func (o *BrokerListElement) GetBBrokerMortgagelicense() bool`

GetBBrokerMortgagelicense returns the BBrokerMortgagelicense field if non-nil, zero value otherwise.

### GetBBrokerMortgagelicenseOk

`func (o *BrokerListElement) GetBBrokerMortgagelicenseOk() (*bool, bool)`

GetBBrokerMortgagelicenseOk returns a tuple with the BBrokerMortgagelicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBBrokerMortgagelicense

`func (o *BrokerListElement) SetBBrokerMortgagelicense(v bool)`

SetBBrokerMortgagelicense sets BBrokerMortgagelicense field to given value.


### GetBBrokerPaidbyofficetranquillit

`func (o *BrokerListElement) GetBBrokerPaidbyofficetranquillit() bool`

GetBBrokerPaidbyofficetranquillit returns the BBrokerPaidbyofficetranquillit field if non-nil, zero value otherwise.

### GetBBrokerPaidbyofficetranquillitOk

`func (o *BrokerListElement) GetBBrokerPaidbyofficetranquillitOk() (*bool, bool)`

GetBBrokerPaidbyofficetranquillitOk returns a tuple with the BBrokerPaidbyofficetranquillit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBBrokerPaidbyofficetranquillit

`func (o *BrokerListElement) SetBBrokerPaidbyofficetranquillit(v bool)`

SetBBrokerPaidbyofficetranquillit sets BBrokerPaidbyofficetranquillit field to given value.


### GetDtBrokerFintraccertification

`func (o *BrokerListElement) GetDtBrokerFintraccertification() string`

GetDtBrokerFintraccertification returns the DtBrokerFintraccertification field if non-nil, zero value otherwise.

### GetDtBrokerFintraccertificationOk

`func (o *BrokerListElement) GetDtBrokerFintraccertificationOk() (*string, bool)`

GetDtBrokerFintraccertificationOk returns a tuple with the DtBrokerFintraccertification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtBrokerFintraccertification

`func (o *BrokerListElement) SetDtBrokerFintraccertification(v string)`

SetDtBrokerFintraccertification sets DtBrokerFintraccertification field to given value.

### HasDtBrokerFintraccertification

`func (o *BrokerListElement) HasDtBrokerFintraccertification() bool`

HasDtBrokerFintraccertification returns a boolean if a field has been set.

### GetBBrokerIsactive

`func (o *BrokerListElement) GetBBrokerIsactive() bool`

GetBBrokerIsactive returns the BBrokerIsactive field if non-nil, zero value otherwise.

### GetBBrokerIsactiveOk

`func (o *BrokerListElement) GetBBrokerIsactiveOk() (*bool, bool)`

GetBBrokerIsactiveOk returns a tuple with the BBrokerIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBBrokerIsactive

`func (o *BrokerListElement) SetBBrokerIsactive(v bool)`

SetBBrokerIsactive sets BBrokerIsactive field to given value.


### GetSContactFirstname

`func (o *BrokerListElement) GetSContactFirstname() string`

GetSContactFirstname returns the SContactFirstname field if non-nil, zero value otherwise.

### GetSContactFirstnameOk

`func (o *BrokerListElement) GetSContactFirstnameOk() (*string, bool)`

GetSContactFirstnameOk returns a tuple with the SContactFirstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSContactFirstname

`func (o *BrokerListElement) SetSContactFirstname(v string)`

SetSContactFirstname sets SContactFirstname field to given value.

### HasSContactFirstname

`func (o *BrokerListElement) HasSContactFirstname() bool`

HasSContactFirstname returns a boolean if a field has been set.

### GetSContactLastname

`func (o *BrokerListElement) GetSContactLastname() string`

GetSContactLastname returns the SContactLastname field if non-nil, zero value otherwise.

### GetSContactLastnameOk

`func (o *BrokerListElement) GetSContactLastnameOk() (*string, bool)`

GetSContactLastnameOk returns a tuple with the SContactLastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSContactLastname

`func (o *BrokerListElement) SetSContactLastname(v string)`

SetSContactLastname sets SContactLastname field to given value.

### HasSContactLastname

`func (o *BrokerListElement) HasSContactLastname() bool`

HasSContactLastname returns a boolean if a field has been set.

### GetDtContactBirthdate

`func (o *BrokerListElement) GetDtContactBirthdate() string`

GetDtContactBirthdate returns the DtContactBirthdate field if non-nil, zero value otherwise.

### GetDtContactBirthdateOk

`func (o *BrokerListElement) GetDtContactBirthdateOk() (*string, bool)`

GetDtContactBirthdateOk returns a tuple with the DtContactBirthdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtContactBirthdate

`func (o *BrokerListElement) SetDtContactBirthdate(v string)`

SetDtContactBirthdate sets DtContactBirthdate field to given value.

### HasDtContactBirthdate

`func (o *BrokerListElement) HasDtContactBirthdate() bool`

HasDtContactBirthdate returns a boolean if a field has been set.

### GetSEmailAddress

`func (o *BrokerListElement) GetSEmailAddress() string`

GetSEmailAddress returns the SEmailAddress field if non-nil, zero value otherwise.

### GetSEmailAddressOk

`func (o *BrokerListElement) GetSEmailAddressOk() (*string, bool)`

GetSEmailAddressOk returns a tuple with the SEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEmailAddress

`func (o *BrokerListElement) SetSEmailAddress(v string)`

SetSEmailAddress sets SEmailAddress field to given value.

### HasSEmailAddress

`func (o *BrokerListElement) HasSEmailAddress() bool`

HasSEmailAddress returns a boolean if a field has been set.

### GetSPhoneE164

`func (o *BrokerListElement) GetSPhoneE164() string`

GetSPhoneE164 returns the SPhoneE164 field if non-nil, zero value otherwise.

### GetSPhoneE164Ok

`func (o *BrokerListElement) GetSPhoneE164Ok() (*string, bool)`

GetSPhoneE164Ok returns a tuple with the SPhoneE164 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSPhoneE164

`func (o *BrokerListElement) SetSPhoneE164(v string)`

SetSPhoneE164 sets SPhoneE164 field to given value.

### HasSPhoneE164

`func (o *BrokerListElement) HasSPhoneE164() bool`

HasSPhoneE164 returns a boolean if a field has been set.

### GetSAddressCivic

`func (o *BrokerListElement) GetSAddressCivic() string`

GetSAddressCivic returns the SAddressCivic field if non-nil, zero value otherwise.

### GetSAddressCivicOk

`func (o *BrokerListElement) GetSAddressCivicOk() (*string, bool)`

GetSAddressCivicOk returns a tuple with the SAddressCivic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAddressCivic

`func (o *BrokerListElement) SetSAddressCivic(v string)`

SetSAddressCivic sets SAddressCivic field to given value.

### HasSAddressCivic

`func (o *BrokerListElement) HasSAddressCivic() bool`

HasSAddressCivic returns a boolean if a field has been set.

### GetSAddressStreet

`func (o *BrokerListElement) GetSAddressStreet() string`

GetSAddressStreet returns the SAddressStreet field if non-nil, zero value otherwise.

### GetSAddressStreetOk

`func (o *BrokerListElement) GetSAddressStreetOk() (*string, bool)`

GetSAddressStreetOk returns a tuple with the SAddressStreet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAddressStreet

`func (o *BrokerListElement) SetSAddressStreet(v string)`

SetSAddressStreet sets SAddressStreet field to given value.

### HasSAddressStreet

`func (o *BrokerListElement) HasSAddressStreet() bool`

HasSAddressStreet returns a boolean if a field has been set.

### GetSAddressSuite

`func (o *BrokerListElement) GetSAddressSuite() string`

GetSAddressSuite returns the SAddressSuite field if non-nil, zero value otherwise.

### GetSAddressSuiteOk

`func (o *BrokerListElement) GetSAddressSuiteOk() (*string, bool)`

GetSAddressSuiteOk returns a tuple with the SAddressSuite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAddressSuite

`func (o *BrokerListElement) SetSAddressSuite(v string)`

SetSAddressSuite sets SAddressSuite field to given value.

### HasSAddressSuite

`func (o *BrokerListElement) HasSAddressSuite() bool`

HasSAddressSuite returns a boolean if a field has been set.

### GetSAddressCity

`func (o *BrokerListElement) GetSAddressCity() string`

GetSAddressCity returns the SAddressCity field if non-nil, zero value otherwise.

### GetSAddressCityOk

`func (o *BrokerListElement) GetSAddressCityOk() (*string, bool)`

GetSAddressCityOk returns a tuple with the SAddressCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAddressCity

`func (o *BrokerListElement) SetSAddressCity(v string)`

SetSAddressCity sets SAddressCity field to given value.

### HasSAddressCity

`func (o *BrokerListElement) HasSAddressCity() bool`

HasSAddressCity returns a boolean if a field has been set.

### GetSAddressZip

`func (o *BrokerListElement) GetSAddressZip() string`

GetSAddressZip returns the SAddressZip field if non-nil, zero value otherwise.

### GetSAddressZipOk

`func (o *BrokerListElement) GetSAddressZipOk() (*string, bool)`

GetSAddressZipOk returns a tuple with the SAddressZip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAddressZip

`func (o *BrokerListElement) SetSAddressZip(v string)`

SetSAddressZip sets SAddressZip field to given value.

### HasSAddressZip

`func (o *BrokerListElement) HasSAddressZip() bool`

HasSAddressZip returns a boolean if a field has been set.

### GetSProvinceNameX

`func (o *BrokerListElement) GetSProvinceNameX() string`

GetSProvinceNameX returns the SProvinceNameX field if non-nil, zero value otherwise.

### GetSProvinceNameXOk

`func (o *BrokerListElement) GetSProvinceNameXOk() (*string, bool)`

GetSProvinceNameXOk returns a tuple with the SProvinceNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSProvinceNameX

`func (o *BrokerListElement) SetSProvinceNameX(v string)`

SetSProvinceNameX sets SProvinceNameX field to given value.

### HasSProvinceNameX

`func (o *BrokerListElement) HasSProvinceNameX() bool`

HasSProvinceNameX returns a boolean if a field has been set.

### GetSCountryNameX

`func (o *BrokerListElement) GetSCountryNameX() string`

GetSCountryNameX returns the SCountryNameX field if non-nil, zero value otherwise.

### GetSCountryNameXOk

`func (o *BrokerListElement) GetSCountryNameXOk() (*string, bool)`

GetSCountryNameXOk returns a tuple with the SCountryNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCountryNameX

`func (o *BrokerListElement) SetSCountryNameX(v string)`

SetSCountryNameX sets SCountryNameX field to given value.

### HasSCountryNameX

`func (o *BrokerListElement) HasSCountryNameX() bool`

HasSCountryNameX returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


