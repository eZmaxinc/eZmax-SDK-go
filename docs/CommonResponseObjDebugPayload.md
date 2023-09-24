# CommonResponseObjDebugPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IVersionMin** | **int32** | The minimum version of the function that can be called | 
**IVersionMax** | **int32** | The maximum version of the function that can be called | 
**ARequiredPermission** | **[]int32** | An array of permissions required to access this function.  If the value \&quot;0\&quot; is present in the array, anyone can call this function.  You must have one of the permission to access the function. You don&#39;t need to have all of them. | 
**BVersionDeprecated** | **bool** | Wheter the current route is deprecated or not | 

## Methods

### NewCommonResponseObjDebugPayload

`func NewCommonResponseObjDebugPayload(iVersionMin int32, iVersionMax int32, aRequiredPermission []int32, bVersionDeprecated bool, ) *CommonResponseObjDebugPayload`

NewCommonResponseObjDebugPayload instantiates a new CommonResponseObjDebugPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonResponseObjDebugPayloadWithDefaults

`func NewCommonResponseObjDebugPayloadWithDefaults() *CommonResponseObjDebugPayload`

NewCommonResponseObjDebugPayloadWithDefaults instantiates a new CommonResponseObjDebugPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIVersionMin

`func (o *CommonResponseObjDebugPayload) GetIVersionMin() int32`

GetIVersionMin returns the IVersionMin field if non-nil, zero value otherwise.

### GetIVersionMinOk

`func (o *CommonResponseObjDebugPayload) GetIVersionMinOk() (*int32, bool)`

GetIVersionMinOk returns a tuple with the IVersionMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIVersionMin

`func (o *CommonResponseObjDebugPayload) SetIVersionMin(v int32)`

SetIVersionMin sets IVersionMin field to given value.


### GetIVersionMax

`func (o *CommonResponseObjDebugPayload) GetIVersionMax() int32`

GetIVersionMax returns the IVersionMax field if non-nil, zero value otherwise.

### GetIVersionMaxOk

`func (o *CommonResponseObjDebugPayload) GetIVersionMaxOk() (*int32, bool)`

GetIVersionMaxOk returns a tuple with the IVersionMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIVersionMax

`func (o *CommonResponseObjDebugPayload) SetIVersionMax(v int32)`

SetIVersionMax sets IVersionMax field to given value.


### GetARequiredPermission

`func (o *CommonResponseObjDebugPayload) GetARequiredPermission() []int32`

GetARequiredPermission returns the ARequiredPermission field if non-nil, zero value otherwise.

### GetARequiredPermissionOk

`func (o *CommonResponseObjDebugPayload) GetARequiredPermissionOk() (*[]int32, bool)`

GetARequiredPermissionOk returns a tuple with the ARequiredPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetARequiredPermission

`func (o *CommonResponseObjDebugPayload) SetARequiredPermission(v []int32)`

SetARequiredPermission sets ARequiredPermission field to given value.


### GetBVersionDeprecated

`func (o *CommonResponseObjDebugPayload) GetBVersionDeprecated() bool`

GetBVersionDeprecated returns the BVersionDeprecated field if non-nil, zero value otherwise.

### GetBVersionDeprecatedOk

`func (o *CommonResponseObjDebugPayload) GetBVersionDeprecatedOk() (*bool, bool)`

GetBVersionDeprecatedOk returns a tuple with the BVersionDeprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBVersionDeprecated

`func (o *CommonResponseObjDebugPayload) SetBVersionDeprecated(v bool)`

SetBVersionDeprecated sets BVersionDeprecated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


