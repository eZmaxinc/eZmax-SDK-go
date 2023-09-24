# UserAutocompleteElementResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EUserType** | [**FieldEUserType**](FieldEUserType.md) |  | 
**SUserName** | **string** | The description of the User in the language of the requester | 
**PkiUserID** | **int32** | The unique ID of the User | 
**BUserIsactive** | **bool** | Whether the User is active or not | 

## Methods

### NewUserAutocompleteElementResponse

`func NewUserAutocompleteElementResponse(eUserType FieldEUserType, sUserName string, pkiUserID int32, bUserIsactive bool, ) *UserAutocompleteElementResponse`

NewUserAutocompleteElementResponse instantiates a new UserAutocompleteElementResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAutocompleteElementResponseWithDefaults

`func NewUserAutocompleteElementResponseWithDefaults() *UserAutocompleteElementResponse`

NewUserAutocompleteElementResponseWithDefaults instantiates a new UserAutocompleteElementResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEUserType

`func (o *UserAutocompleteElementResponse) GetEUserType() FieldEUserType`

GetEUserType returns the EUserType field if non-nil, zero value otherwise.

### GetEUserTypeOk

`func (o *UserAutocompleteElementResponse) GetEUserTypeOk() (*FieldEUserType, bool)`

GetEUserTypeOk returns a tuple with the EUserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEUserType

`func (o *UserAutocompleteElementResponse) SetEUserType(v FieldEUserType)`

SetEUserType sets EUserType field to given value.


### GetSUserName

`func (o *UserAutocompleteElementResponse) GetSUserName() string`

GetSUserName returns the SUserName field if non-nil, zero value otherwise.

### GetSUserNameOk

`func (o *UserAutocompleteElementResponse) GetSUserNameOk() (*string, bool)`

GetSUserNameOk returns a tuple with the SUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserName

`func (o *UserAutocompleteElementResponse) SetSUserName(v string)`

SetSUserName sets SUserName field to given value.


### GetPkiUserID

`func (o *UserAutocompleteElementResponse) GetPkiUserID() int32`

GetPkiUserID returns the PkiUserID field if non-nil, zero value otherwise.

### GetPkiUserIDOk

`func (o *UserAutocompleteElementResponse) GetPkiUserIDOk() (*int32, bool)`

GetPkiUserIDOk returns a tuple with the PkiUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiUserID

`func (o *UserAutocompleteElementResponse) SetPkiUserID(v int32)`

SetPkiUserID sets PkiUserID field to given value.


### GetBUserIsactive

`func (o *UserAutocompleteElementResponse) GetBUserIsactive() bool`

GetBUserIsactive returns the BUserIsactive field if non-nil, zero value otherwise.

### GetBUserIsactiveOk

`func (o *UserAutocompleteElementResponse) GetBUserIsactiveOk() (*bool, bool)`

GetBUserIsactiveOk returns a tuple with the BUserIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBUserIsactive

`func (o *UserAutocompleteElementResponse) SetBUserIsactive(v bool)`

SetBUserIsactive sets BUserIsactive field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


