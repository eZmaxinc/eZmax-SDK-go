# UserCreateEzsignuserV1Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FkiLanguageID** | **int32** | The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English| | 
**SUserFirstname** | **string** | The first name of the user | 
**SUserLastname** | **string** | The last name of the user | 
**SEmailAddress** | **string** | The email address. | 
**SPhoneRegion** | **string** | The region of the phone number. (For a North America Number only)  The region is the \&quot;514\&quot; section in this sample phone number: (514) 990-1516 x123 | 
**SPhoneExchange** | **string** | The exchange of the phone number. (For a North America Number only)  The exchange is the \&quot;990\&quot; section in this sample phone number: (514) 990-1516 x123 | 
**SPhoneNumber** | **string** | The number of the phone number. (For a North America Number only)  The number is the \&quot;1516\&quot; section in this sample phone number: (514) 990-1516 x123 | 
**SPhoneExtension** | Pointer to **string** | The extension of the phone number.  The extension is the \&quot;123\&quot; section in this sample phone number: (514) 990-1516 x123.  It can also be used with international phone numbers | [optional] 

## Methods

### NewUserCreateEzsignuserV1Request

`func NewUserCreateEzsignuserV1Request(fkiLanguageID int32, sUserFirstname string, sUserLastname string, sEmailAddress string, sPhoneRegion string, sPhoneExchange string, sPhoneNumber string, ) *UserCreateEzsignuserV1Request`

NewUserCreateEzsignuserV1Request instantiates a new UserCreateEzsignuserV1Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserCreateEzsignuserV1RequestWithDefaults

`func NewUserCreateEzsignuserV1RequestWithDefaults() *UserCreateEzsignuserV1Request`

NewUserCreateEzsignuserV1RequestWithDefaults instantiates a new UserCreateEzsignuserV1Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFkiLanguageID

`func (o *UserCreateEzsignuserV1Request) GetFkiLanguageID() int32`

GetFkiLanguageID returns the FkiLanguageID field if non-nil, zero value otherwise.

### GetFkiLanguageIDOk

`func (o *UserCreateEzsignuserV1Request) GetFkiLanguageIDOk() (*int32, bool)`

GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiLanguageID

`func (o *UserCreateEzsignuserV1Request) SetFkiLanguageID(v int32)`

SetFkiLanguageID sets FkiLanguageID field to given value.


### GetSUserFirstname

`func (o *UserCreateEzsignuserV1Request) GetSUserFirstname() string`

GetSUserFirstname returns the SUserFirstname field if non-nil, zero value otherwise.

### GetSUserFirstnameOk

`func (o *UserCreateEzsignuserV1Request) GetSUserFirstnameOk() (*string, bool)`

GetSUserFirstnameOk returns a tuple with the SUserFirstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserFirstname

`func (o *UserCreateEzsignuserV1Request) SetSUserFirstname(v string)`

SetSUserFirstname sets SUserFirstname field to given value.


### GetSUserLastname

`func (o *UserCreateEzsignuserV1Request) GetSUserLastname() string`

GetSUserLastname returns the SUserLastname field if non-nil, zero value otherwise.

### GetSUserLastnameOk

`func (o *UserCreateEzsignuserV1Request) GetSUserLastnameOk() (*string, bool)`

GetSUserLastnameOk returns a tuple with the SUserLastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserLastname

`func (o *UserCreateEzsignuserV1Request) SetSUserLastname(v string)`

SetSUserLastname sets SUserLastname field to given value.


### GetSEmailAddress

`func (o *UserCreateEzsignuserV1Request) GetSEmailAddress() string`

GetSEmailAddress returns the SEmailAddress field if non-nil, zero value otherwise.

### GetSEmailAddressOk

`func (o *UserCreateEzsignuserV1Request) GetSEmailAddressOk() (*string, bool)`

GetSEmailAddressOk returns a tuple with the SEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEmailAddress

`func (o *UserCreateEzsignuserV1Request) SetSEmailAddress(v string)`

SetSEmailAddress sets SEmailAddress field to given value.


### GetSPhoneRegion

`func (o *UserCreateEzsignuserV1Request) GetSPhoneRegion() string`

GetSPhoneRegion returns the SPhoneRegion field if non-nil, zero value otherwise.

### GetSPhoneRegionOk

`func (o *UserCreateEzsignuserV1Request) GetSPhoneRegionOk() (*string, bool)`

GetSPhoneRegionOk returns a tuple with the SPhoneRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSPhoneRegion

`func (o *UserCreateEzsignuserV1Request) SetSPhoneRegion(v string)`

SetSPhoneRegion sets SPhoneRegion field to given value.


### GetSPhoneExchange

`func (o *UserCreateEzsignuserV1Request) GetSPhoneExchange() string`

GetSPhoneExchange returns the SPhoneExchange field if non-nil, zero value otherwise.

### GetSPhoneExchangeOk

`func (o *UserCreateEzsignuserV1Request) GetSPhoneExchangeOk() (*string, bool)`

GetSPhoneExchangeOk returns a tuple with the SPhoneExchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSPhoneExchange

`func (o *UserCreateEzsignuserV1Request) SetSPhoneExchange(v string)`

SetSPhoneExchange sets SPhoneExchange field to given value.


### GetSPhoneNumber

`func (o *UserCreateEzsignuserV1Request) GetSPhoneNumber() string`

GetSPhoneNumber returns the SPhoneNumber field if non-nil, zero value otherwise.

### GetSPhoneNumberOk

`func (o *UserCreateEzsignuserV1Request) GetSPhoneNumberOk() (*string, bool)`

GetSPhoneNumberOk returns a tuple with the SPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSPhoneNumber

`func (o *UserCreateEzsignuserV1Request) SetSPhoneNumber(v string)`

SetSPhoneNumber sets SPhoneNumber field to given value.


### GetSPhoneExtension

`func (o *UserCreateEzsignuserV1Request) GetSPhoneExtension() string`

GetSPhoneExtension returns the SPhoneExtension field if non-nil, zero value otherwise.

### GetSPhoneExtensionOk

`func (o *UserCreateEzsignuserV1Request) GetSPhoneExtensionOk() (*string, bool)`

GetSPhoneExtensionOk returns a tuple with the SPhoneExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSPhoneExtension

`func (o *UserCreateEzsignuserV1Request) SetSPhoneExtension(v string)`

SetSPhoneExtension sets SPhoneExtension field to given value.

### HasSPhoneExtension

`func (o *UserCreateEzsignuserV1Request) HasSPhoneExtension() bool`

HasSPhoneExtension returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


