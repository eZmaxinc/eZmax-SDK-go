# CountryAutocompleteElementResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiCountryID** | **int32** | The unique ID of the Country.  Here are some common values (Complete list must be retrieved from API):  |Value|Description| |-|-| |1|Canada| |2|United-States| | 
**SCountryNameX** | **string** | The name of the Country in the language of the requester | 
**SCountryShortname** | **string** | The shortname of the Country | 
**BCountryIsactive** | **bool** | Whether the Country is active or not | 

## Methods

### NewCountryAutocompleteElementResponse

`func NewCountryAutocompleteElementResponse(pkiCountryID int32, sCountryNameX string, sCountryShortname string, bCountryIsactive bool, ) *CountryAutocompleteElementResponse`

NewCountryAutocompleteElementResponse instantiates a new CountryAutocompleteElementResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCountryAutocompleteElementResponseWithDefaults

`func NewCountryAutocompleteElementResponseWithDefaults() *CountryAutocompleteElementResponse`

NewCountryAutocompleteElementResponseWithDefaults instantiates a new CountryAutocompleteElementResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiCountryID

`func (o *CountryAutocompleteElementResponse) GetPkiCountryID() int32`

GetPkiCountryID returns the PkiCountryID field if non-nil, zero value otherwise.

### GetPkiCountryIDOk

`func (o *CountryAutocompleteElementResponse) GetPkiCountryIDOk() (*int32, bool)`

GetPkiCountryIDOk returns a tuple with the PkiCountryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiCountryID

`func (o *CountryAutocompleteElementResponse) SetPkiCountryID(v int32)`

SetPkiCountryID sets PkiCountryID field to given value.


### GetSCountryNameX

`func (o *CountryAutocompleteElementResponse) GetSCountryNameX() string`

GetSCountryNameX returns the SCountryNameX field if non-nil, zero value otherwise.

### GetSCountryNameXOk

`func (o *CountryAutocompleteElementResponse) GetSCountryNameXOk() (*string, bool)`

GetSCountryNameXOk returns a tuple with the SCountryNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCountryNameX

`func (o *CountryAutocompleteElementResponse) SetSCountryNameX(v string)`

SetSCountryNameX sets SCountryNameX field to given value.


### GetSCountryShortname

`func (o *CountryAutocompleteElementResponse) GetSCountryShortname() string`

GetSCountryShortname returns the SCountryShortname field if non-nil, zero value otherwise.

### GetSCountryShortnameOk

`func (o *CountryAutocompleteElementResponse) GetSCountryShortnameOk() (*string, bool)`

GetSCountryShortnameOk returns a tuple with the SCountryShortname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCountryShortname

`func (o *CountryAutocompleteElementResponse) SetSCountryShortname(v string)`

SetSCountryShortname sets SCountryShortname field to given value.


### GetBCountryIsactive

`func (o *CountryAutocompleteElementResponse) GetBCountryIsactive() bool`

GetBCountryIsactive returns the BCountryIsactive field if non-nil, zero value otherwise.

### GetBCountryIsactiveOk

`func (o *CountryAutocompleteElementResponse) GetBCountryIsactiveOk() (*bool, bool)`

GetBCountryIsactiveOk returns a tuple with the BCountryIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBCountryIsactive

`func (o *CountryAutocompleteElementResponse) SetBCountryIsactive(v bool)`

SetBCountryIsactive sets BCountryIsactive field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


