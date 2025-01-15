# PermissionEditObjectV1Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjPermission** | [**PermissionRequest**](PermissionRequest.md) | A Permission Object and children to create a complete structure | 

## Methods

### NewPermissionEditObjectV1Request

`func NewPermissionEditObjectV1Request(objPermission PermissionRequest, ) *PermissionEditObjectV1Request`

NewPermissionEditObjectV1Request instantiates a new PermissionEditObjectV1Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionEditObjectV1RequestWithDefaults

`func NewPermissionEditObjectV1RequestWithDefaults() *PermissionEditObjectV1Request`

NewPermissionEditObjectV1RequestWithDefaults instantiates a new PermissionEditObjectV1Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjPermission

`func (o *PermissionEditObjectV1Request) GetObjPermission() PermissionRequest`

GetObjPermission returns the ObjPermission field if non-nil, zero value otherwise.

### GetObjPermissionOk

`func (o *PermissionEditObjectV1Request) GetObjPermissionOk() (*PermissionRequest, bool)`

GetObjPermissionOk returns a tuple with the ObjPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjPermission

`func (o *PermissionEditObjectV1Request) SetObjPermission(v PermissionRequest)`

SetObjPermission sets ObjPermission field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


