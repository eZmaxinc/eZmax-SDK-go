# UsergroupListElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiUsergroupID** | **int32** | The unique ID of the Usergroup | 
**SUsergroupNameX** | **string** | The Name of the Usergroup in the language of the requester | 
**ICountUser** | **int32** | Number of users in group | 

## Methods

### NewUsergroupListElement

`func NewUsergroupListElement(pkiUsergroupID int32, sUsergroupNameX string, iCountUser int32, ) *UsergroupListElement`

NewUsergroupListElement instantiates a new UsergroupListElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsergroupListElementWithDefaults

`func NewUsergroupListElementWithDefaults() *UsergroupListElement`

NewUsergroupListElementWithDefaults instantiates a new UsergroupListElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiUsergroupID

`func (o *UsergroupListElement) GetPkiUsergroupID() int32`

GetPkiUsergroupID returns the PkiUsergroupID field if non-nil, zero value otherwise.

### GetPkiUsergroupIDOk

`func (o *UsergroupListElement) GetPkiUsergroupIDOk() (*int32, bool)`

GetPkiUsergroupIDOk returns a tuple with the PkiUsergroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiUsergroupID

`func (o *UsergroupListElement) SetPkiUsergroupID(v int32)`

SetPkiUsergroupID sets PkiUsergroupID field to given value.


### GetSUsergroupNameX

`func (o *UsergroupListElement) GetSUsergroupNameX() string`

GetSUsergroupNameX returns the SUsergroupNameX field if non-nil, zero value otherwise.

### GetSUsergroupNameXOk

`func (o *UsergroupListElement) GetSUsergroupNameXOk() (*string, bool)`

GetSUsergroupNameXOk returns a tuple with the SUsergroupNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUsergroupNameX

`func (o *UsergroupListElement) SetSUsergroupNameX(v string)`

SetSUsergroupNameX sets SUsergroupNameX field to given value.


### GetICountUser

`func (o *UsergroupListElement) GetICountUser() int32`

GetICountUser returns the ICountUser field if non-nil, zero value otherwise.

### GetICountUserOk

`func (o *UsergroupListElement) GetICountUserOk() (*int32, bool)`

GetICountUserOk returns a tuple with the ICountUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetICountUser

`func (o *UsergroupListElement) SetICountUser(v int32)`

SetICountUser sets ICountUser field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


