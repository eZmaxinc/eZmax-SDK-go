# EmployeeListElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEmployeeID** | **int32** | The unique ID of the Employee. | 
**FkiDepartmentID** | **int32** | The unique ID of the Department | 
**SEmployeeCode** | **string** | The code of the Employee | 
**SEmployeeInternalcode** | **string** | The internalcode of the Employee | 
**BEmployeeIsactive** | **bool** | Whether the employee is active or not | 
**DtEmployeeHiredate** | Pointer to **string** | The hiredate of the Employee | [optional] 
**DtEmployeeLeavedate** | Pointer to **string** | The leavedate of the Employee | [optional] 
**SDepartmentNameX** | Pointer to **string** | The Name of the Department in the language of the requester | [optional] 
**SContactFirstname** | Pointer to **string** | The First name of the contact | [optional] 
**SContactLastname** | Pointer to **string** | The Last name of the contact | [optional] 
**SPhoneE164** | Pointer to **string** | A phone number in E.164 Format | [optional] 
**SEmailAddress** | Pointer to **string** | The email address. | [optional] 
**SAddressCivic** | Pointer to **string** | The Civic number. | [optional] 
**SAddressStreet** | Pointer to **string** | The Street Name | [optional] 
**SAddressSuite** | Pointer to **string** | The Suite or appartment number | [optional] 
**SAddressCity** | Pointer to **string** | The City name | [optional] 
**SAddressZip** | Pointer to **string** | The Postal/Zip Code  The value must be entered without spaces | [optional] 
**SProvinceNameX** | Pointer to **string** | The name of the Province in the language of the requester | [optional] 
**SCountryNameX** | Pointer to **string** | The name of the Country in the language of the requester | [optional] 

## Methods

### NewEmployeeListElement

`func NewEmployeeListElement(pkiEmployeeID int32, fkiDepartmentID int32, sEmployeeCode string, sEmployeeInternalcode string, bEmployeeIsactive bool, ) *EmployeeListElement`

NewEmployeeListElement instantiates a new EmployeeListElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmployeeListElementWithDefaults

`func NewEmployeeListElementWithDefaults() *EmployeeListElement`

NewEmployeeListElementWithDefaults instantiates a new EmployeeListElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEmployeeID

`func (o *EmployeeListElement) GetPkiEmployeeID() int32`

GetPkiEmployeeID returns the PkiEmployeeID field if non-nil, zero value otherwise.

### GetPkiEmployeeIDOk

`func (o *EmployeeListElement) GetPkiEmployeeIDOk() (*int32, bool)`

GetPkiEmployeeIDOk returns a tuple with the PkiEmployeeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEmployeeID

`func (o *EmployeeListElement) SetPkiEmployeeID(v int32)`

SetPkiEmployeeID sets PkiEmployeeID field to given value.


### GetFkiDepartmentID

`func (o *EmployeeListElement) GetFkiDepartmentID() int32`

GetFkiDepartmentID returns the FkiDepartmentID field if non-nil, zero value otherwise.

### GetFkiDepartmentIDOk

`func (o *EmployeeListElement) GetFkiDepartmentIDOk() (*int32, bool)`

GetFkiDepartmentIDOk returns a tuple with the FkiDepartmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiDepartmentID

`func (o *EmployeeListElement) SetFkiDepartmentID(v int32)`

SetFkiDepartmentID sets FkiDepartmentID field to given value.


### GetSEmployeeCode

`func (o *EmployeeListElement) GetSEmployeeCode() string`

GetSEmployeeCode returns the SEmployeeCode field if non-nil, zero value otherwise.

### GetSEmployeeCodeOk

`func (o *EmployeeListElement) GetSEmployeeCodeOk() (*string, bool)`

GetSEmployeeCodeOk returns a tuple with the SEmployeeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEmployeeCode

`func (o *EmployeeListElement) SetSEmployeeCode(v string)`

SetSEmployeeCode sets SEmployeeCode field to given value.


### GetSEmployeeInternalcode

`func (o *EmployeeListElement) GetSEmployeeInternalcode() string`

GetSEmployeeInternalcode returns the SEmployeeInternalcode field if non-nil, zero value otherwise.

### GetSEmployeeInternalcodeOk

`func (o *EmployeeListElement) GetSEmployeeInternalcodeOk() (*string, bool)`

GetSEmployeeInternalcodeOk returns a tuple with the SEmployeeInternalcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEmployeeInternalcode

`func (o *EmployeeListElement) SetSEmployeeInternalcode(v string)`

SetSEmployeeInternalcode sets SEmployeeInternalcode field to given value.


### GetBEmployeeIsactive

`func (o *EmployeeListElement) GetBEmployeeIsactive() bool`

GetBEmployeeIsactive returns the BEmployeeIsactive field if non-nil, zero value otherwise.

### GetBEmployeeIsactiveOk

`func (o *EmployeeListElement) GetBEmployeeIsactiveOk() (*bool, bool)`

GetBEmployeeIsactiveOk returns a tuple with the BEmployeeIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEmployeeIsactive

`func (o *EmployeeListElement) SetBEmployeeIsactive(v bool)`

SetBEmployeeIsactive sets BEmployeeIsactive field to given value.


### GetDtEmployeeHiredate

`func (o *EmployeeListElement) GetDtEmployeeHiredate() string`

GetDtEmployeeHiredate returns the DtEmployeeHiredate field if non-nil, zero value otherwise.

### GetDtEmployeeHiredateOk

`func (o *EmployeeListElement) GetDtEmployeeHiredateOk() (*string, bool)`

GetDtEmployeeHiredateOk returns a tuple with the DtEmployeeHiredate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtEmployeeHiredate

`func (o *EmployeeListElement) SetDtEmployeeHiredate(v string)`

SetDtEmployeeHiredate sets DtEmployeeHiredate field to given value.

### HasDtEmployeeHiredate

`func (o *EmployeeListElement) HasDtEmployeeHiredate() bool`

HasDtEmployeeHiredate returns a boolean if a field has been set.

### GetDtEmployeeLeavedate

`func (o *EmployeeListElement) GetDtEmployeeLeavedate() string`

GetDtEmployeeLeavedate returns the DtEmployeeLeavedate field if non-nil, zero value otherwise.

### GetDtEmployeeLeavedateOk

`func (o *EmployeeListElement) GetDtEmployeeLeavedateOk() (*string, bool)`

GetDtEmployeeLeavedateOk returns a tuple with the DtEmployeeLeavedate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtEmployeeLeavedate

`func (o *EmployeeListElement) SetDtEmployeeLeavedate(v string)`

SetDtEmployeeLeavedate sets DtEmployeeLeavedate field to given value.

### HasDtEmployeeLeavedate

`func (o *EmployeeListElement) HasDtEmployeeLeavedate() bool`

HasDtEmployeeLeavedate returns a boolean if a field has been set.

### GetSDepartmentNameX

`func (o *EmployeeListElement) GetSDepartmentNameX() string`

GetSDepartmentNameX returns the SDepartmentNameX field if non-nil, zero value otherwise.

### GetSDepartmentNameXOk

`func (o *EmployeeListElement) GetSDepartmentNameXOk() (*string, bool)`

GetSDepartmentNameXOk returns a tuple with the SDepartmentNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSDepartmentNameX

`func (o *EmployeeListElement) SetSDepartmentNameX(v string)`

SetSDepartmentNameX sets SDepartmentNameX field to given value.

### HasSDepartmentNameX

`func (o *EmployeeListElement) HasSDepartmentNameX() bool`

HasSDepartmentNameX returns a boolean if a field has been set.

### GetSContactFirstname

`func (o *EmployeeListElement) GetSContactFirstname() string`

GetSContactFirstname returns the SContactFirstname field if non-nil, zero value otherwise.

### GetSContactFirstnameOk

`func (o *EmployeeListElement) GetSContactFirstnameOk() (*string, bool)`

GetSContactFirstnameOk returns a tuple with the SContactFirstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSContactFirstname

`func (o *EmployeeListElement) SetSContactFirstname(v string)`

SetSContactFirstname sets SContactFirstname field to given value.

### HasSContactFirstname

`func (o *EmployeeListElement) HasSContactFirstname() bool`

HasSContactFirstname returns a boolean if a field has been set.

### GetSContactLastname

`func (o *EmployeeListElement) GetSContactLastname() string`

GetSContactLastname returns the SContactLastname field if non-nil, zero value otherwise.

### GetSContactLastnameOk

`func (o *EmployeeListElement) GetSContactLastnameOk() (*string, bool)`

GetSContactLastnameOk returns a tuple with the SContactLastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSContactLastname

`func (o *EmployeeListElement) SetSContactLastname(v string)`

SetSContactLastname sets SContactLastname field to given value.

### HasSContactLastname

`func (o *EmployeeListElement) HasSContactLastname() bool`

HasSContactLastname returns a boolean if a field has been set.

### GetSPhoneE164

`func (o *EmployeeListElement) GetSPhoneE164() string`

GetSPhoneE164 returns the SPhoneE164 field if non-nil, zero value otherwise.

### GetSPhoneE164Ok

`func (o *EmployeeListElement) GetSPhoneE164Ok() (*string, bool)`

GetSPhoneE164Ok returns a tuple with the SPhoneE164 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSPhoneE164

`func (o *EmployeeListElement) SetSPhoneE164(v string)`

SetSPhoneE164 sets SPhoneE164 field to given value.

### HasSPhoneE164

`func (o *EmployeeListElement) HasSPhoneE164() bool`

HasSPhoneE164 returns a boolean if a field has been set.

### GetSEmailAddress

`func (o *EmployeeListElement) GetSEmailAddress() string`

GetSEmailAddress returns the SEmailAddress field if non-nil, zero value otherwise.

### GetSEmailAddressOk

`func (o *EmployeeListElement) GetSEmailAddressOk() (*string, bool)`

GetSEmailAddressOk returns a tuple with the SEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEmailAddress

`func (o *EmployeeListElement) SetSEmailAddress(v string)`

SetSEmailAddress sets SEmailAddress field to given value.

### HasSEmailAddress

`func (o *EmployeeListElement) HasSEmailAddress() bool`

HasSEmailAddress returns a boolean if a field has been set.

### GetSAddressCivic

`func (o *EmployeeListElement) GetSAddressCivic() string`

GetSAddressCivic returns the SAddressCivic field if non-nil, zero value otherwise.

### GetSAddressCivicOk

`func (o *EmployeeListElement) GetSAddressCivicOk() (*string, bool)`

GetSAddressCivicOk returns a tuple with the SAddressCivic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAddressCivic

`func (o *EmployeeListElement) SetSAddressCivic(v string)`

SetSAddressCivic sets SAddressCivic field to given value.

### HasSAddressCivic

`func (o *EmployeeListElement) HasSAddressCivic() bool`

HasSAddressCivic returns a boolean if a field has been set.

### GetSAddressStreet

`func (o *EmployeeListElement) GetSAddressStreet() string`

GetSAddressStreet returns the SAddressStreet field if non-nil, zero value otherwise.

### GetSAddressStreetOk

`func (o *EmployeeListElement) GetSAddressStreetOk() (*string, bool)`

GetSAddressStreetOk returns a tuple with the SAddressStreet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAddressStreet

`func (o *EmployeeListElement) SetSAddressStreet(v string)`

SetSAddressStreet sets SAddressStreet field to given value.

### HasSAddressStreet

`func (o *EmployeeListElement) HasSAddressStreet() bool`

HasSAddressStreet returns a boolean if a field has been set.

### GetSAddressSuite

`func (o *EmployeeListElement) GetSAddressSuite() string`

GetSAddressSuite returns the SAddressSuite field if non-nil, zero value otherwise.

### GetSAddressSuiteOk

`func (o *EmployeeListElement) GetSAddressSuiteOk() (*string, bool)`

GetSAddressSuiteOk returns a tuple with the SAddressSuite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAddressSuite

`func (o *EmployeeListElement) SetSAddressSuite(v string)`

SetSAddressSuite sets SAddressSuite field to given value.

### HasSAddressSuite

`func (o *EmployeeListElement) HasSAddressSuite() bool`

HasSAddressSuite returns a boolean if a field has been set.

### GetSAddressCity

`func (o *EmployeeListElement) GetSAddressCity() string`

GetSAddressCity returns the SAddressCity field if non-nil, zero value otherwise.

### GetSAddressCityOk

`func (o *EmployeeListElement) GetSAddressCityOk() (*string, bool)`

GetSAddressCityOk returns a tuple with the SAddressCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAddressCity

`func (o *EmployeeListElement) SetSAddressCity(v string)`

SetSAddressCity sets SAddressCity field to given value.

### HasSAddressCity

`func (o *EmployeeListElement) HasSAddressCity() bool`

HasSAddressCity returns a boolean if a field has been set.

### GetSAddressZip

`func (o *EmployeeListElement) GetSAddressZip() string`

GetSAddressZip returns the SAddressZip field if non-nil, zero value otherwise.

### GetSAddressZipOk

`func (o *EmployeeListElement) GetSAddressZipOk() (*string, bool)`

GetSAddressZipOk returns a tuple with the SAddressZip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAddressZip

`func (o *EmployeeListElement) SetSAddressZip(v string)`

SetSAddressZip sets SAddressZip field to given value.

### HasSAddressZip

`func (o *EmployeeListElement) HasSAddressZip() bool`

HasSAddressZip returns a boolean if a field has been set.

### GetSProvinceNameX

`func (o *EmployeeListElement) GetSProvinceNameX() string`

GetSProvinceNameX returns the SProvinceNameX field if non-nil, zero value otherwise.

### GetSProvinceNameXOk

`func (o *EmployeeListElement) GetSProvinceNameXOk() (*string, bool)`

GetSProvinceNameXOk returns a tuple with the SProvinceNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSProvinceNameX

`func (o *EmployeeListElement) SetSProvinceNameX(v string)`

SetSProvinceNameX sets SProvinceNameX field to given value.

### HasSProvinceNameX

`func (o *EmployeeListElement) HasSProvinceNameX() bool`

HasSProvinceNameX returns a boolean if a field has been set.

### GetSCountryNameX

`func (o *EmployeeListElement) GetSCountryNameX() string`

GetSCountryNameX returns the SCountryNameX field if non-nil, zero value otherwise.

### GetSCountryNameXOk

`func (o *EmployeeListElement) GetSCountryNameXOk() (*string, bool)`

GetSCountryNameXOk returns a tuple with the SCountryNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCountryNameX

`func (o *EmployeeListElement) SetSCountryNameX(v string)`

SetSCountryNameX sets SCountryNameX field to given value.

### HasSCountryNameX

`func (o *EmployeeListElement) HasSCountryNameX() bool`

HasSCountryNameX returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


