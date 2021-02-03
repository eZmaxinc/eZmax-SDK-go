# CommonResponseObjDebugPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IVersionMin** | **int32** | The minimum version of the function that can be called | 
**IVersionMax** | **int32** | The maximum version of the function that can be called | 
**ARequiredPermissions** | **[]int32** | An array of permissions required to access this function.  If the value \&quot;0\&quot; is present in the array, anyone can call this function.  You must have one of the permission to access the function. You don&#39;t need to have all of them. | 

## Methods

### NewCommonResponseObjDebugPayload

`func NewCommonResponseObjDebugPayload(iVersionMin int32, iVersionMax int32, aRequiredPermissions []int32, ) *CommonResponseObjDebugPayload`

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


### GetARequiredPermissions

`func (o *CommonResponseObjDebugPayload) GetARequiredPermissions() []int32`

GetARequiredPermissions returns the ARequiredPermissions field if non-nil, zero value otherwise.

### GetARequiredPermissionsOk

`func (o *CommonResponseObjDebugPayload) GetARequiredPermissionsOk() (*[]int32, bool)`

GetARequiredPermissionsOk returns a tuple with the ARequiredPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetARequiredPermissions

`func (o *CommonResponseObjDebugPayload) SetARequiredPermissions(v []int32)`

SetARequiredPermissions sets ARequiredPermissions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


