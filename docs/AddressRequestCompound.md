# AddressRequestCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FkiAddresstypeID** | **int32** | The unique ID of the Addresstype.  Valid values:  |Value|Description| |-|-| |1|Office| |2|Home| |3|Real Estate Invoice| |4|Invoicing| |5|Shipping| | 
**SAddressCivic** | **string** | The Civic number. | 
**SAddressStreet** | **string** | The Street Name | 
**SAddressSuite** | **string** | The Suite or appartment number | 
**SAddressCity** | **string** | The City name | 
**FkiProvinceID** | **int32** | The unique ID of the Province.  Here are some common values (Complete list must be retrieved from API):  |Value|Description| |-|-| |1|(Canada) Alberta |2|(Canada) British Columbia| |3|(Canada) Manitoba| |3|(Canada) Manitoba| |4|(Canada) New Brunswick| |5|(Canada) Newfoundland| |6|(Canada) Northwest Territories| |7|(Canada) Nova Scotia| |8|(Canada) Nunavut| |9|(Canada) Ontario| |10|(Canada) Prince Edward Island| |11|(Canada) Quebec| |12|(Canada) Saskatchewan| |13|(Canada) Yukon| |14|(United-States) Alabama| |15|(United-States) Alaska| |16|(United-States) Arizona| |17|(United-States) Arkansas| |18|(United-States) California| |19|(United-States) Colorado| |20|(United-States) Connecticut| |21|(United-States) Delaware| |22|(United-States) District of Columbia| |23|(United-States) Florida| |24|(United-States) Georgia| |25|(United-States) Hawaii| |26|(United-States) Idaho| |27|(United-States) Illinois| |28|(United-States) Indiana| |29|(United-States) Iowa| |30|(United-States) Kansas| |31|(United-States) Kentucky| |32|(United-States) Louisiane| |33|(United-States) Maine| |34|(United-States) Maryland| |35|(United-States) Massachusetts| |36|(United-States) Michigan| |37|(United-States) Minnesota| |38|(United-States) Mississippi| |39|(United-States) Missouri| |40|(United-States) Montana| |41|(United-States) Nebraska| |42|(United-States) Nevada| |43|(United-States) New Hampshire| |44|(United-States) New Jersey| |45|(United-States) New Mexico| |46|(United-States) New York| |47|(United-States) North Carolina| |48|(United-States) North Dakota| |49|(United-States) Ohio| |50|(United-States) Oklahoma| |51|(United-States) Oregon| |52|(United-States) Pennsylvania| |53|(United-States) Rhode Island| |54|(United-States) South Carolina| |55|(United-States) South Dakota| |56|(United-States) Tennessee| |57|(United-States) Texas| |58|(United-States) Utah| |60|(United-States) Vermont| |59|(United-States) Virginia| |61|(United-States) Washington| |62|(United-States) West Virginia| |63|(United-States) Wisconsin| |64|(United-States) Wyoming| | 
**FkiCountryID** | **int32** | The unique ID of the Country.  Here are some common values (Complete list must be retrieved from API):  |Value|Description| |-|-| |1|Canada| |2|United-States| | 
**SAddressZip** | **string** | The Postal/Zip Code  The value must be entered without spaces | 
**FAddressLongitude** | Pointer to **string** | The Longitude of the Address | [optional] 
**FAddressLatitude** | Pointer to **string** | The Latitude of the Address | [optional] 

## Methods

### NewAddressRequestCompound

`func NewAddressRequestCompound(fkiAddresstypeID int32, sAddressCivic string, sAddressStreet string, sAddressSuite string, sAddressCity string, fkiProvinceID int32, fkiCountryID int32, sAddressZip string, ) *AddressRequestCompound`

NewAddressRequestCompound instantiates a new AddressRequestCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressRequestCompoundWithDefaults

`func NewAddressRequestCompoundWithDefaults() *AddressRequestCompound`

NewAddressRequestCompoundWithDefaults instantiates a new AddressRequestCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFkiAddresstypeID

`func (o *AddressRequestCompound) GetFkiAddresstypeID() int32`

GetFkiAddresstypeID returns the FkiAddresstypeID field if non-nil, zero value otherwise.

### GetFkiAddresstypeIDOk

`func (o *AddressRequestCompound) GetFkiAddresstypeIDOk() (*int32, bool)`

GetFkiAddresstypeIDOk returns a tuple with the FkiAddresstypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiAddresstypeID

`func (o *AddressRequestCompound) SetFkiAddresstypeID(v int32)`

SetFkiAddresstypeID sets FkiAddresstypeID field to given value.


### GetSAddressCivic

`func (o *AddressRequestCompound) GetSAddressCivic() string`

GetSAddressCivic returns the SAddressCivic field if non-nil, zero value otherwise.

### GetSAddressCivicOk

`func (o *AddressRequestCompound) GetSAddressCivicOk() (*string, bool)`

GetSAddressCivicOk returns a tuple with the SAddressCivic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAddressCivic

`func (o *AddressRequestCompound) SetSAddressCivic(v string)`

SetSAddressCivic sets SAddressCivic field to given value.


### GetSAddressStreet

`func (o *AddressRequestCompound) GetSAddressStreet() string`

GetSAddressStreet returns the SAddressStreet field if non-nil, zero value otherwise.

### GetSAddressStreetOk

`func (o *AddressRequestCompound) GetSAddressStreetOk() (*string, bool)`

GetSAddressStreetOk returns a tuple with the SAddressStreet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAddressStreet

`func (o *AddressRequestCompound) SetSAddressStreet(v string)`

SetSAddressStreet sets SAddressStreet field to given value.


### GetSAddressSuite

`func (o *AddressRequestCompound) GetSAddressSuite() string`

GetSAddressSuite returns the SAddressSuite field if non-nil, zero value otherwise.

### GetSAddressSuiteOk

`func (o *AddressRequestCompound) GetSAddressSuiteOk() (*string, bool)`

GetSAddressSuiteOk returns a tuple with the SAddressSuite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAddressSuite

`func (o *AddressRequestCompound) SetSAddressSuite(v string)`

SetSAddressSuite sets SAddressSuite field to given value.


### GetSAddressCity

`func (o *AddressRequestCompound) GetSAddressCity() string`

GetSAddressCity returns the SAddressCity field if non-nil, zero value otherwise.

### GetSAddressCityOk

`func (o *AddressRequestCompound) GetSAddressCityOk() (*string, bool)`

GetSAddressCityOk returns a tuple with the SAddressCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAddressCity

`func (o *AddressRequestCompound) SetSAddressCity(v string)`

SetSAddressCity sets SAddressCity field to given value.


### GetFkiProvinceID

`func (o *AddressRequestCompound) GetFkiProvinceID() int32`

GetFkiProvinceID returns the FkiProvinceID field if non-nil, zero value otherwise.

### GetFkiProvinceIDOk

`func (o *AddressRequestCompound) GetFkiProvinceIDOk() (*int32, bool)`

GetFkiProvinceIDOk returns a tuple with the FkiProvinceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiProvinceID

`func (o *AddressRequestCompound) SetFkiProvinceID(v int32)`

SetFkiProvinceID sets FkiProvinceID field to given value.


### GetFkiCountryID

`func (o *AddressRequestCompound) GetFkiCountryID() int32`

GetFkiCountryID returns the FkiCountryID field if non-nil, zero value otherwise.

### GetFkiCountryIDOk

`func (o *AddressRequestCompound) GetFkiCountryIDOk() (*int32, bool)`

GetFkiCountryIDOk returns a tuple with the FkiCountryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiCountryID

`func (o *AddressRequestCompound) SetFkiCountryID(v int32)`

SetFkiCountryID sets FkiCountryID field to given value.


### GetSAddressZip

`func (o *AddressRequestCompound) GetSAddressZip() string`

GetSAddressZip returns the SAddressZip field if non-nil, zero value otherwise.

### GetSAddressZipOk

`func (o *AddressRequestCompound) GetSAddressZipOk() (*string, bool)`

GetSAddressZipOk returns a tuple with the SAddressZip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAddressZip

`func (o *AddressRequestCompound) SetSAddressZip(v string)`

SetSAddressZip sets SAddressZip field to given value.


### GetFAddressLongitude

`func (o *AddressRequestCompound) GetFAddressLongitude() string`

GetFAddressLongitude returns the FAddressLongitude field if non-nil, zero value otherwise.

### GetFAddressLongitudeOk

`func (o *AddressRequestCompound) GetFAddressLongitudeOk() (*string, bool)`

GetFAddressLongitudeOk returns a tuple with the FAddressLongitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFAddressLongitude

`func (o *AddressRequestCompound) SetFAddressLongitude(v string)`

SetFAddressLongitude sets FAddressLongitude field to given value.

### HasFAddressLongitude

`func (o *AddressRequestCompound) HasFAddressLongitude() bool`

HasFAddressLongitude returns a boolean if a field has been set.

### GetFAddressLatitude

`func (o *AddressRequestCompound) GetFAddressLatitude() string`

GetFAddressLatitude returns the FAddressLatitude field if non-nil, zero value otherwise.

### GetFAddressLatitudeOk

`func (o *AddressRequestCompound) GetFAddressLatitudeOk() (*string, bool)`

GetFAddressLatitudeOk returns a tuple with the FAddressLatitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFAddressLatitude

`func (o *AddressRequestCompound) SetFAddressLatitude(v string)`

SetFAddressLatitude sets FAddressLatitude field to given value.

### HasFAddressLatitude

`func (o *AddressRequestCompound) HasFAddressLatitude() bool`

HasFAddressLatitude returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


