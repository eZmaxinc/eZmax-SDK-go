# DomainListElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiDomainID** | **int32** | The unique ID of the Domain | 
**SDomainName** | **string** | The name of the Domain | 

## Methods

### NewDomainListElement

`func NewDomainListElement(pkiDomainID int32, sDomainName string, ) *DomainListElement`

NewDomainListElement instantiates a new DomainListElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainListElementWithDefaults

`func NewDomainListElementWithDefaults() *DomainListElement`

NewDomainListElementWithDefaults instantiates a new DomainListElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiDomainID

`func (o *DomainListElement) GetPkiDomainID() int32`

GetPkiDomainID returns the PkiDomainID field if non-nil, zero value otherwise.

### GetPkiDomainIDOk

`func (o *DomainListElement) GetPkiDomainIDOk() (*int32, bool)`

GetPkiDomainIDOk returns a tuple with the PkiDomainID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiDomainID

`func (o *DomainListElement) SetPkiDomainID(v int32)`

SetPkiDomainID sets PkiDomainID field to given value.


### GetSDomainName

`func (o *DomainListElement) GetSDomainName() string`

GetSDomainName returns the SDomainName field if non-nil, zero value otherwise.

### GetSDomainNameOk

`func (o *DomainListElement) GetSDomainNameOk() (*string, bool)`

GetSDomainNameOk returns a tuple with the SDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSDomainName

`func (o *DomainListElement) SetSDomainName(v string)`

SetSDomainName sets SDomainName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


