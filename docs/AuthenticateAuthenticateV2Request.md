# AuthenticateAuthenticateV2Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PksCustomerCode** | **string** | The customer code assigned to your account | 
**SEmailAddress** | Pointer to **string** | The email address. | [optional] 
**SUserLoginname** | Pointer to **string** | The Login name of the User. | [optional] 
**SPassword** | Pointer to **string** | A Password.  Must meet complexity requirements | [optional] 
**SPasswordEncrypted** | Pointer to **string** | A Password encrypted and encoded in Base64  Must meet complexity requirements | [optional] 

## Methods

### NewAuthenticateAuthenticateV2Request

`func NewAuthenticateAuthenticateV2Request(pksCustomerCode string, ) *AuthenticateAuthenticateV2Request`

NewAuthenticateAuthenticateV2Request instantiates a new AuthenticateAuthenticateV2Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticateAuthenticateV2RequestWithDefaults

`func NewAuthenticateAuthenticateV2RequestWithDefaults() *AuthenticateAuthenticateV2Request`

NewAuthenticateAuthenticateV2RequestWithDefaults instantiates a new AuthenticateAuthenticateV2Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPksCustomerCode

`func (o *AuthenticateAuthenticateV2Request) GetPksCustomerCode() string`

GetPksCustomerCode returns the PksCustomerCode field if non-nil, zero value otherwise.

### GetPksCustomerCodeOk

`func (o *AuthenticateAuthenticateV2Request) GetPksCustomerCodeOk() (*string, bool)`

GetPksCustomerCodeOk returns a tuple with the PksCustomerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPksCustomerCode

`func (o *AuthenticateAuthenticateV2Request) SetPksCustomerCode(v string)`

SetPksCustomerCode sets PksCustomerCode field to given value.


### GetSEmailAddress

`func (o *AuthenticateAuthenticateV2Request) GetSEmailAddress() string`

GetSEmailAddress returns the SEmailAddress field if non-nil, zero value otherwise.

### GetSEmailAddressOk

`func (o *AuthenticateAuthenticateV2Request) GetSEmailAddressOk() (*string, bool)`

GetSEmailAddressOk returns a tuple with the SEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEmailAddress

`func (o *AuthenticateAuthenticateV2Request) SetSEmailAddress(v string)`

SetSEmailAddress sets SEmailAddress field to given value.

### HasSEmailAddress

`func (o *AuthenticateAuthenticateV2Request) HasSEmailAddress() bool`

HasSEmailAddress returns a boolean if a field has been set.

### GetSUserLoginname

`func (o *AuthenticateAuthenticateV2Request) GetSUserLoginname() string`

GetSUserLoginname returns the SUserLoginname field if non-nil, zero value otherwise.

### GetSUserLoginnameOk

`func (o *AuthenticateAuthenticateV2Request) GetSUserLoginnameOk() (*string, bool)`

GetSUserLoginnameOk returns a tuple with the SUserLoginname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserLoginname

`func (o *AuthenticateAuthenticateV2Request) SetSUserLoginname(v string)`

SetSUserLoginname sets SUserLoginname field to given value.

### HasSUserLoginname

`func (o *AuthenticateAuthenticateV2Request) HasSUserLoginname() bool`

HasSUserLoginname returns a boolean if a field has been set.

### GetSPassword

`func (o *AuthenticateAuthenticateV2Request) GetSPassword() string`

GetSPassword returns the SPassword field if non-nil, zero value otherwise.

### GetSPasswordOk

`func (o *AuthenticateAuthenticateV2Request) GetSPasswordOk() (*string, bool)`

GetSPasswordOk returns a tuple with the SPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSPassword

`func (o *AuthenticateAuthenticateV2Request) SetSPassword(v string)`

SetSPassword sets SPassword field to given value.

### HasSPassword

`func (o *AuthenticateAuthenticateV2Request) HasSPassword() bool`

HasSPassword returns a boolean if a field has been set.

### GetSPasswordEncrypted

`func (o *AuthenticateAuthenticateV2Request) GetSPasswordEncrypted() string`

GetSPasswordEncrypted returns the SPasswordEncrypted field if non-nil, zero value otherwise.

### GetSPasswordEncryptedOk

`func (o *AuthenticateAuthenticateV2Request) GetSPasswordEncryptedOk() (*string, bool)`

GetSPasswordEncryptedOk returns a tuple with the SPasswordEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSPasswordEncrypted

`func (o *AuthenticateAuthenticateV2Request) SetSPasswordEncrypted(v string)`

SetSPasswordEncrypted sets SPasswordEncrypted field to given value.

### HasSPasswordEncrypted

`func (o *AuthenticateAuthenticateV2Request) HasSPasswordEncrypted() bool`

HasSPasswordEncrypted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


