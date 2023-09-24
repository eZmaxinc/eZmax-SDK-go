# EmailRequestCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEmailID** | Pointer to **int32** | The unique ID of the Email | [optional] 
**FkiEmailtypeID** | **int32** | The unique ID of the Emailtype.  Valid values:  |Value|Description| |-|-| |1|Office| |2|Home| | 
**SEmailAddress** | **string** | The email address. | 

## Methods

### NewEmailRequestCompound

`func NewEmailRequestCompound(fkiEmailtypeID int32, sEmailAddress string, ) *EmailRequestCompound`

NewEmailRequestCompound instantiates a new EmailRequestCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailRequestCompoundWithDefaults

`func NewEmailRequestCompoundWithDefaults() *EmailRequestCompound`

NewEmailRequestCompoundWithDefaults instantiates a new EmailRequestCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEmailID

`func (o *EmailRequestCompound) GetPkiEmailID() int32`

GetPkiEmailID returns the PkiEmailID field if non-nil, zero value otherwise.

### GetPkiEmailIDOk

`func (o *EmailRequestCompound) GetPkiEmailIDOk() (*int32, bool)`

GetPkiEmailIDOk returns a tuple with the PkiEmailID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEmailID

`func (o *EmailRequestCompound) SetPkiEmailID(v int32)`

SetPkiEmailID sets PkiEmailID field to given value.

### HasPkiEmailID

`func (o *EmailRequestCompound) HasPkiEmailID() bool`

HasPkiEmailID returns a boolean if a field has been set.

### GetFkiEmailtypeID

`func (o *EmailRequestCompound) GetFkiEmailtypeID() int32`

GetFkiEmailtypeID returns the FkiEmailtypeID field if non-nil, zero value otherwise.

### GetFkiEmailtypeIDOk

`func (o *EmailRequestCompound) GetFkiEmailtypeIDOk() (*int32, bool)`

GetFkiEmailtypeIDOk returns a tuple with the FkiEmailtypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEmailtypeID

`func (o *EmailRequestCompound) SetFkiEmailtypeID(v int32)`

SetFkiEmailtypeID sets FkiEmailtypeID field to given value.


### GetSEmailAddress

`func (o *EmailRequestCompound) GetSEmailAddress() string`

GetSEmailAddress returns the SEmailAddress field if non-nil, zero value otherwise.

### GetSEmailAddressOk

`func (o *EmailRequestCompound) GetSEmailAddressOk() (*string, bool)`

GetSEmailAddressOk returns a tuple with the SEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEmailAddress

`func (o *EmailRequestCompound) SetSEmailAddress(v string)`

SetSEmailAddress sets SEmailAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


