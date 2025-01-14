# DomainRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiDomainID** | Pointer to **int32** | The unique ID of the Domain | [optional] 
**SDomainName** | **string** | The name of the Domain | 

## Methods

### NewDomainRequest

`func NewDomainRequest(sDomainName string, ) *DomainRequest`

NewDomainRequest instantiates a new DomainRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainRequestWithDefaults

`func NewDomainRequestWithDefaults() *DomainRequest`

NewDomainRequestWithDefaults instantiates a new DomainRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiDomainID

`func (o *DomainRequest) GetPkiDomainID() int32`

GetPkiDomainID returns the PkiDomainID field if non-nil, zero value otherwise.

### GetPkiDomainIDOk

`func (o *DomainRequest) GetPkiDomainIDOk() (*int32, bool)`

GetPkiDomainIDOk returns a tuple with the PkiDomainID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiDomainID

`func (o *DomainRequest) SetPkiDomainID(v int32)`

SetPkiDomainID sets PkiDomainID field to given value.

### HasPkiDomainID

`func (o *DomainRequest) HasPkiDomainID() bool`

HasPkiDomainID returns a boolean if a field has been set.

### GetSDomainName

`func (o *DomainRequest) GetSDomainName() string`

GetSDomainName returns the SDomainName field if non-nil, zero value otherwise.

### GetSDomainNameOk

`func (o *DomainRequest) GetSDomainNameOk() (*string, bool)`

GetSDomainNameOk returns a tuple with the SDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSDomainName

`func (o *DomainRequest) SetSDomainName(v string)`

SetSDomainName sets SDomainName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


