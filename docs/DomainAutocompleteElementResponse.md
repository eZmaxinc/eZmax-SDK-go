# DomainAutocompleteElementResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiDomainID** | **int32** | The unique ID of the Domain | 
**SDomainName** | **string** | The name of the Domain | 
**BDomainIsactive** | **bool** | Whether the Domain is active or not | 

## Methods

### NewDomainAutocompleteElementResponse

`func NewDomainAutocompleteElementResponse(pkiDomainID int32, sDomainName string, bDomainIsactive bool, ) *DomainAutocompleteElementResponse`

NewDomainAutocompleteElementResponse instantiates a new DomainAutocompleteElementResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainAutocompleteElementResponseWithDefaults

`func NewDomainAutocompleteElementResponseWithDefaults() *DomainAutocompleteElementResponse`

NewDomainAutocompleteElementResponseWithDefaults instantiates a new DomainAutocompleteElementResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiDomainID

`func (o *DomainAutocompleteElementResponse) GetPkiDomainID() int32`

GetPkiDomainID returns the PkiDomainID field if non-nil, zero value otherwise.

### GetPkiDomainIDOk

`func (o *DomainAutocompleteElementResponse) GetPkiDomainIDOk() (*int32, bool)`

GetPkiDomainIDOk returns a tuple with the PkiDomainID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiDomainID

`func (o *DomainAutocompleteElementResponse) SetPkiDomainID(v int32)`

SetPkiDomainID sets PkiDomainID field to given value.


### GetSDomainName

`func (o *DomainAutocompleteElementResponse) GetSDomainName() string`

GetSDomainName returns the SDomainName field if non-nil, zero value otherwise.

### GetSDomainNameOk

`func (o *DomainAutocompleteElementResponse) GetSDomainNameOk() (*string, bool)`

GetSDomainNameOk returns a tuple with the SDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSDomainName

`func (o *DomainAutocompleteElementResponse) SetSDomainName(v string)`

SetSDomainName sets SDomainName field to given value.


### GetBDomainIsactive

`func (o *DomainAutocompleteElementResponse) GetBDomainIsactive() bool`

GetBDomainIsactive returns the BDomainIsactive field if non-nil, zero value otherwise.

### GetBDomainIsactiveOk

`func (o *DomainAutocompleteElementResponse) GetBDomainIsactiveOk() (*bool, bool)`

GetBDomainIsactiveOk returns a tuple with the BDomainIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDomainIsactive

`func (o *DomainAutocompleteElementResponse) SetBDomainIsactive(v bool)`

SetBDomainIsactive sets BDomainIsactive field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


