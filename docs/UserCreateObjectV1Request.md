# UserCreateObjectV1Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AObjUser** | [**[]UserRequestCompound**](UserRequestCompound.md) |  | 

## Methods

### NewUserCreateObjectV1Request

`func NewUserCreateObjectV1Request(aObjUser []UserRequestCompound, ) *UserCreateObjectV1Request`

NewUserCreateObjectV1Request instantiates a new UserCreateObjectV1Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserCreateObjectV1RequestWithDefaults

`func NewUserCreateObjectV1RequestWithDefaults() *UserCreateObjectV1Request`

NewUserCreateObjectV1RequestWithDefaults instantiates a new UserCreateObjectV1Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAObjUser

`func (o *UserCreateObjectV1Request) GetAObjUser() []UserRequestCompound`

GetAObjUser returns the AObjUser field if non-nil, zero value otherwise.

### GetAObjUserOk

`func (o *UserCreateObjectV1Request) GetAObjUserOk() (*[]UserRequestCompound, bool)`

GetAObjUserOk returns a tuple with the AObjUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjUser

`func (o *UserCreateObjectV1Request) SetAObjUser(v []UserRequestCompound)`

SetAObjUser sets AObjUser field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


