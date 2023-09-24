# EmailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEmailID** | **int32** | The unique ID of the Email | 
**FkiEmailtypeID** | **int32** | The unique ID of the Emailtype.  Valid values:  |Value|Description| |-|-| |1|Office| |2|Home| | 
**SEmailAddress** | **string** | The email address. | 

## Methods

### NewEmailResponse

`func NewEmailResponse(pkiEmailID int32, fkiEmailtypeID int32, sEmailAddress string, ) *EmailResponse`

NewEmailResponse instantiates a new EmailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailResponseWithDefaults

`func NewEmailResponseWithDefaults() *EmailResponse`

NewEmailResponseWithDefaults instantiates a new EmailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEmailID

`func (o *EmailResponse) GetPkiEmailID() int32`

GetPkiEmailID returns the PkiEmailID field if non-nil, zero value otherwise.

### GetPkiEmailIDOk

`func (o *EmailResponse) GetPkiEmailIDOk() (*int32, bool)`

GetPkiEmailIDOk returns a tuple with the PkiEmailID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEmailID

`func (o *EmailResponse) SetPkiEmailID(v int32)`

SetPkiEmailID sets PkiEmailID field to given value.


### GetFkiEmailtypeID

`func (o *EmailResponse) GetFkiEmailtypeID() int32`

GetFkiEmailtypeID returns the FkiEmailtypeID field if non-nil, zero value otherwise.

### GetFkiEmailtypeIDOk

`func (o *EmailResponse) GetFkiEmailtypeIDOk() (*int32, bool)`

GetFkiEmailtypeIDOk returns a tuple with the FkiEmailtypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEmailtypeID

`func (o *EmailResponse) SetFkiEmailtypeID(v int32)`

SetFkiEmailtypeID sets FkiEmailtypeID field to given value.


### GetSEmailAddress

`func (o *EmailResponse) GetSEmailAddress() string`

GetSEmailAddress returns the SEmailAddress field if non-nil, zero value otherwise.

### GetSEmailAddressOk

`func (o *EmailResponse) GetSEmailAddressOk() (*string, bool)`

GetSEmailAddressOk returns a tuple with the SEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEmailAddress

`func (o *EmailResponse) SetSEmailAddress(v string)`

SetSEmailAddress sets SEmailAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


