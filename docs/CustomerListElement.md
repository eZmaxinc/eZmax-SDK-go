# CustomerListElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiCustomerID** | **int32** | The unique ID of the Customer. | 
**SCustomerName** | **string** | The name of the Customer | 
**SCustomerNote** | Pointer to **string** | A note for the Customer | [optional] 
**SCustomerCode** | **string** | The code of the Customer | 
**BCustomerIsactive** | **bool** | Whether the customer is active or not | 
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

### NewCustomerListElement

`func NewCustomerListElement(pkiCustomerID int32, sCustomerName string, sCustomerCode string, bCustomerIsactive bool, ) *CustomerListElement`

NewCustomerListElement instantiates a new CustomerListElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerListElementWithDefaults

`func NewCustomerListElementWithDefaults() *CustomerListElement`

NewCustomerListElementWithDefaults instantiates a new CustomerListElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiCustomerID

`func (o *CustomerListElement) GetPkiCustomerID() int32`

GetPkiCustomerID returns the PkiCustomerID field if non-nil, zero value otherwise.

### GetPkiCustomerIDOk

`func (o *CustomerListElement) GetPkiCustomerIDOk() (*int32, bool)`

GetPkiCustomerIDOk returns a tuple with the PkiCustomerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiCustomerID

`func (o *CustomerListElement) SetPkiCustomerID(v int32)`

SetPkiCustomerID sets PkiCustomerID field to given value.


### GetSCustomerName

`func (o *CustomerListElement) GetSCustomerName() string`

GetSCustomerName returns the SCustomerName field if non-nil, zero value otherwise.

### GetSCustomerNameOk

`func (o *CustomerListElement) GetSCustomerNameOk() (*string, bool)`

GetSCustomerNameOk returns a tuple with the SCustomerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCustomerName

`func (o *CustomerListElement) SetSCustomerName(v string)`

SetSCustomerName sets SCustomerName field to given value.


### GetSCustomerNote

`func (o *CustomerListElement) GetSCustomerNote() string`

GetSCustomerNote returns the SCustomerNote field if non-nil, zero value otherwise.

### GetSCustomerNoteOk

`func (o *CustomerListElement) GetSCustomerNoteOk() (*string, bool)`

GetSCustomerNoteOk returns a tuple with the SCustomerNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCustomerNote

`func (o *CustomerListElement) SetSCustomerNote(v string)`

SetSCustomerNote sets SCustomerNote field to given value.

### HasSCustomerNote

`func (o *CustomerListElement) HasSCustomerNote() bool`

HasSCustomerNote returns a boolean if a field has been set.

### GetSCustomerCode

`func (o *CustomerListElement) GetSCustomerCode() string`

GetSCustomerCode returns the SCustomerCode field if non-nil, zero value otherwise.

### GetSCustomerCodeOk

`func (o *CustomerListElement) GetSCustomerCodeOk() (*string, bool)`

GetSCustomerCodeOk returns a tuple with the SCustomerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCustomerCode

`func (o *CustomerListElement) SetSCustomerCode(v string)`

SetSCustomerCode sets SCustomerCode field to given value.


### GetBCustomerIsactive

`func (o *CustomerListElement) GetBCustomerIsactive() bool`

GetBCustomerIsactive returns the BCustomerIsactive field if non-nil, zero value otherwise.

### GetBCustomerIsactiveOk

`func (o *CustomerListElement) GetBCustomerIsactiveOk() (*bool, bool)`

GetBCustomerIsactiveOk returns a tuple with the BCustomerIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBCustomerIsactive

`func (o *CustomerListElement) SetBCustomerIsactive(v bool)`

SetBCustomerIsactive sets BCustomerIsactive field to given value.


### GetSPhoneE164

`func (o *CustomerListElement) GetSPhoneE164() string`

GetSPhoneE164 returns the SPhoneE164 field if non-nil, zero value otherwise.

### GetSPhoneE164Ok

`func (o *CustomerListElement) GetSPhoneE164Ok() (*string, bool)`

GetSPhoneE164Ok returns a tuple with the SPhoneE164 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSPhoneE164

`func (o *CustomerListElement) SetSPhoneE164(v string)`

SetSPhoneE164 sets SPhoneE164 field to given value.

### HasSPhoneE164

`func (o *CustomerListElement) HasSPhoneE164() bool`

HasSPhoneE164 returns a boolean if a field has been set.

### GetSEmailAddress

`func (o *CustomerListElement) GetSEmailAddress() string`

GetSEmailAddress returns the SEmailAddress field if non-nil, zero value otherwise.

### GetSEmailAddressOk

`func (o *CustomerListElement) GetSEmailAddressOk() (*string, bool)`

GetSEmailAddressOk returns a tuple with the SEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEmailAddress

`func (o *CustomerListElement) SetSEmailAddress(v string)`

SetSEmailAddress sets SEmailAddress field to given value.

### HasSEmailAddress

`func (o *CustomerListElement) HasSEmailAddress() bool`

HasSEmailAddress returns a boolean if a field has been set.

### GetSAddressCivic

`func (o *CustomerListElement) GetSAddressCivic() string`

GetSAddressCivic returns the SAddressCivic field if non-nil, zero value otherwise.

### GetSAddressCivicOk

`func (o *CustomerListElement) GetSAddressCivicOk() (*string, bool)`

GetSAddressCivicOk returns a tuple with the SAddressCivic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAddressCivic

`func (o *CustomerListElement) SetSAddressCivic(v string)`

SetSAddressCivic sets SAddressCivic field to given value.

### HasSAddressCivic

`func (o *CustomerListElement) HasSAddressCivic() bool`

HasSAddressCivic returns a boolean if a field has been set.

### GetSAddressStreet

`func (o *CustomerListElement) GetSAddressStreet() string`

GetSAddressStreet returns the SAddressStreet field if non-nil, zero value otherwise.

### GetSAddressStreetOk

`func (o *CustomerListElement) GetSAddressStreetOk() (*string, bool)`

GetSAddressStreetOk returns a tuple with the SAddressStreet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAddressStreet

`func (o *CustomerListElement) SetSAddressStreet(v string)`

SetSAddressStreet sets SAddressStreet field to given value.

### HasSAddressStreet

`func (o *CustomerListElement) HasSAddressStreet() bool`

HasSAddressStreet returns a boolean if a field has been set.

### GetSAddressSuite

`func (o *CustomerListElement) GetSAddressSuite() string`

GetSAddressSuite returns the SAddressSuite field if non-nil, zero value otherwise.

### GetSAddressSuiteOk

`func (o *CustomerListElement) GetSAddressSuiteOk() (*string, bool)`

GetSAddressSuiteOk returns a tuple with the SAddressSuite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAddressSuite

`func (o *CustomerListElement) SetSAddressSuite(v string)`

SetSAddressSuite sets SAddressSuite field to given value.

### HasSAddressSuite

`func (o *CustomerListElement) HasSAddressSuite() bool`

HasSAddressSuite returns a boolean if a field has been set.

### GetSAddressCity

`func (o *CustomerListElement) GetSAddressCity() string`

GetSAddressCity returns the SAddressCity field if non-nil, zero value otherwise.

### GetSAddressCityOk

`func (o *CustomerListElement) GetSAddressCityOk() (*string, bool)`

GetSAddressCityOk returns a tuple with the SAddressCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAddressCity

`func (o *CustomerListElement) SetSAddressCity(v string)`

SetSAddressCity sets SAddressCity field to given value.

### HasSAddressCity

`func (o *CustomerListElement) HasSAddressCity() bool`

HasSAddressCity returns a boolean if a field has been set.

### GetSAddressZip

`func (o *CustomerListElement) GetSAddressZip() string`

GetSAddressZip returns the SAddressZip field if non-nil, zero value otherwise.

### GetSAddressZipOk

`func (o *CustomerListElement) GetSAddressZipOk() (*string, bool)`

GetSAddressZipOk returns a tuple with the SAddressZip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAddressZip

`func (o *CustomerListElement) SetSAddressZip(v string)`

SetSAddressZip sets SAddressZip field to given value.

### HasSAddressZip

`func (o *CustomerListElement) HasSAddressZip() bool`

HasSAddressZip returns a boolean if a field has been set.

### GetSProvinceNameX

`func (o *CustomerListElement) GetSProvinceNameX() string`

GetSProvinceNameX returns the SProvinceNameX field if non-nil, zero value otherwise.

### GetSProvinceNameXOk

`func (o *CustomerListElement) GetSProvinceNameXOk() (*string, bool)`

GetSProvinceNameXOk returns a tuple with the SProvinceNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSProvinceNameX

`func (o *CustomerListElement) SetSProvinceNameX(v string)`

SetSProvinceNameX sets SProvinceNameX field to given value.

### HasSProvinceNameX

`func (o *CustomerListElement) HasSProvinceNameX() bool`

HasSProvinceNameX returns a boolean if a field has been set.

### GetSCountryNameX

`func (o *CustomerListElement) GetSCountryNameX() string`

GetSCountryNameX returns the SCountryNameX field if non-nil, zero value otherwise.

### GetSCountryNameXOk

`func (o *CustomerListElement) GetSCountryNameXOk() (*string, bool)`

GetSCountryNameXOk returns a tuple with the SCountryNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCountryNameX

`func (o *CustomerListElement) SetSCountryNameX(v string)`

SetSCountryNameX sets SCountryNameX field to given value.

### HasSCountryNameX

`func (o *CustomerListElement) HasSCountryNameX() bool`

HasSCountryNameX returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


