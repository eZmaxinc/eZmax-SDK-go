# DomainRequestCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiDomainID** | Pointer to **int32** | The unique ID of the Domain | [optional] 
**SDomainName** | **string** | The name of the Domain | 

## Methods

### NewDomainRequestCompound

`func NewDomainRequestCompound(sDomainName string, ) *DomainRequestCompound`

NewDomainRequestCompound instantiates a new DomainRequestCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainRequestCompoundWithDefaults

`func NewDomainRequestCompoundWithDefaults() *DomainRequestCompound`

NewDomainRequestCompoundWithDefaults instantiates a new DomainRequestCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiDomainID

`func (o *DomainRequestCompound) GetPkiDomainID() int32`

GetPkiDomainID returns the PkiDomainID field if non-nil, zero value otherwise.

### GetPkiDomainIDOk

`func (o *DomainRequestCompound) GetPkiDomainIDOk() (*int32, bool)`

GetPkiDomainIDOk returns a tuple with the PkiDomainID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiDomainID

`func (o *DomainRequestCompound) SetPkiDomainID(v int32)`

SetPkiDomainID sets PkiDomainID field to given value.

### HasPkiDomainID

`func (o *DomainRequestCompound) HasPkiDomainID() bool`

HasPkiDomainID returns a boolean if a field has been set.

### GetSDomainName

`func (o *DomainRequestCompound) GetSDomainName() string`

GetSDomainName returns the SDomainName field if non-nil, zero value otherwise.

### GetSDomainNameOk

`func (o *DomainRequestCompound) GetSDomainNameOk() (*string, bool)`

GetSDomainNameOk returns a tuple with the SDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSDomainName

`func (o *DomainRequestCompound) SetSDomainName(v string)`

SetSDomainName sets SDomainName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


