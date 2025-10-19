# CustomUserNameResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SContacttitleNameX** | Pointer to **string** | The name of the Contacttitle in the language of the requester | [optional] 
**SUserLastname** | **string** | The last name of the user | 
**SUserFirstname** | **string** | The first name of the user | 
**SUserJobtitle** | Pointer to **string** | The job title of the user | [optional] 

## Methods

### NewCustomUserNameResponse

`func NewCustomUserNameResponse(sUserLastname string, sUserFirstname string, ) *CustomUserNameResponse`

NewCustomUserNameResponse instantiates a new CustomUserNameResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomUserNameResponseWithDefaults

`func NewCustomUserNameResponseWithDefaults() *CustomUserNameResponse`

NewCustomUserNameResponseWithDefaults instantiates a new CustomUserNameResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSContacttitleNameX

`func (o *CustomUserNameResponse) GetSContacttitleNameX() string`

GetSContacttitleNameX returns the SContacttitleNameX field if non-nil, zero value otherwise.

### GetSContacttitleNameXOk

`func (o *CustomUserNameResponse) GetSContacttitleNameXOk() (*string, bool)`

GetSContacttitleNameXOk returns a tuple with the SContacttitleNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSContacttitleNameX

`func (o *CustomUserNameResponse) SetSContacttitleNameX(v string)`

SetSContacttitleNameX sets SContacttitleNameX field to given value.

### HasSContacttitleNameX

`func (o *CustomUserNameResponse) HasSContacttitleNameX() bool`

HasSContacttitleNameX returns a boolean if a field has been set.

### GetSUserLastname

`func (o *CustomUserNameResponse) GetSUserLastname() string`

GetSUserLastname returns the SUserLastname field if non-nil, zero value otherwise.

### GetSUserLastnameOk

`func (o *CustomUserNameResponse) GetSUserLastnameOk() (*string, bool)`

GetSUserLastnameOk returns a tuple with the SUserLastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserLastname

`func (o *CustomUserNameResponse) SetSUserLastname(v string)`

SetSUserLastname sets SUserLastname field to given value.


### GetSUserFirstname

`func (o *CustomUserNameResponse) GetSUserFirstname() string`

GetSUserFirstname returns the SUserFirstname field if non-nil, zero value otherwise.

### GetSUserFirstnameOk

`func (o *CustomUserNameResponse) GetSUserFirstnameOk() (*string, bool)`

GetSUserFirstnameOk returns a tuple with the SUserFirstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserFirstname

`func (o *CustomUserNameResponse) SetSUserFirstname(v string)`

SetSUserFirstname sets SUserFirstname field to given value.


### GetSUserJobtitle

`func (o *CustomUserNameResponse) GetSUserJobtitle() string`

GetSUserJobtitle returns the SUserJobtitle field if non-nil, zero value otherwise.

### GetSUserJobtitleOk

`func (o *CustomUserNameResponse) GetSUserJobtitleOk() (*string, bool)`

GetSUserJobtitleOk returns a tuple with the SUserJobtitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserJobtitle

`func (o *CustomUserNameResponse) SetSUserJobtitle(v string)`

SetSUserJobtitle sets SUserJobtitle field to given value.

### HasSUserJobtitle

`func (o *CustomUserNameResponse) HasSUserJobtitle() bool`

HasSUserJobtitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


