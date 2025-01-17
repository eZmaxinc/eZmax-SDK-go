# CommonResponseObjDebugPayloadGetList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IVersionMin** | **int32** | The minimum version of the function that can be called | 
**IVersionMax** | **int32** | The maximum version of the function that can be called | 
**ARequiredPermission** | **[]int32** | An array of permissions required to access this function.  If the value \&quot;0\&quot; is present in the array, anyone can call this function.  You must have one of the permission to access the function. You don&#39;t need to have all of them. | 
**BVersionDeprecated** | **bool** | Wheter the current route is deprecated or not | 
**DtResponseDate** | **string** | Represent a Date Time. The timezone is the one configured in the User&#39;s profile. | 
**AFilter** | [**CommonResponseFilter**](CommonResponseFilter.md) |  | 
**AOrderBy** | **map[string]string** | List of available values for *eOrderBy* | 
**IRowMax** | **int32** | The maximum numbers of results to be returned.  When the content-type is **application/json** there is an implicit default of 10 000.  When it&#39;s **application/vnd.openxmlformats-officedocument.spreadsheetml.sheet** the is no implicit default so if you do not specify iRowMax, all records will be returned. | 
**IRowOffset** | **int32** | The starting element from where to start retrieving the results. For example if you started at iRowOffset&#x3D;0 and asked for iRowMax&#x3D;100, to get the next 100 results, you could specify iRowOffset&#x3D;100&amp;iRowMax&#x3D;100, | [default to 0]

## Methods

### NewCommonResponseObjDebugPayloadGetList

`func NewCommonResponseObjDebugPayloadGetList(iVersionMin int32, iVersionMax int32, aRequiredPermission []int32, bVersionDeprecated bool, dtResponseDate string, aFilter CommonResponseFilter, aOrderBy map[string]string, iRowMax int32, iRowOffset int32, ) *CommonResponseObjDebugPayloadGetList`

NewCommonResponseObjDebugPayloadGetList instantiates a new CommonResponseObjDebugPayloadGetList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonResponseObjDebugPayloadGetListWithDefaults

`func NewCommonResponseObjDebugPayloadGetListWithDefaults() *CommonResponseObjDebugPayloadGetList`

NewCommonResponseObjDebugPayloadGetListWithDefaults instantiates a new CommonResponseObjDebugPayloadGetList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIVersionMin

`func (o *CommonResponseObjDebugPayloadGetList) GetIVersionMin() int32`

GetIVersionMin returns the IVersionMin field if non-nil, zero value otherwise.

### GetIVersionMinOk

`func (o *CommonResponseObjDebugPayloadGetList) GetIVersionMinOk() (*int32, bool)`

GetIVersionMinOk returns a tuple with the IVersionMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIVersionMin

`func (o *CommonResponseObjDebugPayloadGetList) SetIVersionMin(v int32)`

SetIVersionMin sets IVersionMin field to given value.


### GetIVersionMax

`func (o *CommonResponseObjDebugPayloadGetList) GetIVersionMax() int32`

GetIVersionMax returns the IVersionMax field if non-nil, zero value otherwise.

### GetIVersionMaxOk

`func (o *CommonResponseObjDebugPayloadGetList) GetIVersionMaxOk() (*int32, bool)`

GetIVersionMaxOk returns a tuple with the IVersionMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIVersionMax

`func (o *CommonResponseObjDebugPayloadGetList) SetIVersionMax(v int32)`

SetIVersionMax sets IVersionMax field to given value.


### GetARequiredPermission

`func (o *CommonResponseObjDebugPayloadGetList) GetARequiredPermission() []int32`

GetARequiredPermission returns the ARequiredPermission field if non-nil, zero value otherwise.

### GetARequiredPermissionOk

`func (o *CommonResponseObjDebugPayloadGetList) GetARequiredPermissionOk() (*[]int32, bool)`

GetARequiredPermissionOk returns a tuple with the ARequiredPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetARequiredPermission

`func (o *CommonResponseObjDebugPayloadGetList) SetARequiredPermission(v []int32)`

SetARequiredPermission sets ARequiredPermission field to given value.


### GetBVersionDeprecated

`func (o *CommonResponseObjDebugPayloadGetList) GetBVersionDeprecated() bool`

GetBVersionDeprecated returns the BVersionDeprecated field if non-nil, zero value otherwise.

### GetBVersionDeprecatedOk

`func (o *CommonResponseObjDebugPayloadGetList) GetBVersionDeprecatedOk() (*bool, bool)`

GetBVersionDeprecatedOk returns a tuple with the BVersionDeprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBVersionDeprecated

`func (o *CommonResponseObjDebugPayloadGetList) SetBVersionDeprecated(v bool)`

SetBVersionDeprecated sets BVersionDeprecated field to given value.


### GetDtResponseDate

`func (o *CommonResponseObjDebugPayloadGetList) GetDtResponseDate() string`

GetDtResponseDate returns the DtResponseDate field if non-nil, zero value otherwise.

### GetDtResponseDateOk

`func (o *CommonResponseObjDebugPayloadGetList) GetDtResponseDateOk() (*string, bool)`

GetDtResponseDateOk returns a tuple with the DtResponseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtResponseDate

`func (o *CommonResponseObjDebugPayloadGetList) SetDtResponseDate(v string)`

SetDtResponseDate sets DtResponseDate field to given value.


### GetAFilter

`func (o *CommonResponseObjDebugPayloadGetList) GetAFilter() CommonResponseFilter`

GetAFilter returns the AFilter field if non-nil, zero value otherwise.

### GetAFilterOk

`func (o *CommonResponseObjDebugPayloadGetList) GetAFilterOk() (*CommonResponseFilter, bool)`

GetAFilterOk returns a tuple with the AFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAFilter

`func (o *CommonResponseObjDebugPayloadGetList) SetAFilter(v CommonResponseFilter)`

SetAFilter sets AFilter field to given value.


### GetAOrderBy

`func (o *CommonResponseObjDebugPayloadGetList) GetAOrderBy() map[string]string`

GetAOrderBy returns the AOrderBy field if non-nil, zero value otherwise.

### GetAOrderByOk

`func (o *CommonResponseObjDebugPayloadGetList) GetAOrderByOk() (*map[string]string, bool)`

GetAOrderByOk returns a tuple with the AOrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAOrderBy

`func (o *CommonResponseObjDebugPayloadGetList) SetAOrderBy(v map[string]string)`

SetAOrderBy sets AOrderBy field to given value.


### GetIRowMax

`func (o *CommonResponseObjDebugPayloadGetList) GetIRowMax() int32`

GetIRowMax returns the IRowMax field if non-nil, zero value otherwise.

### GetIRowMaxOk

`func (o *CommonResponseObjDebugPayloadGetList) GetIRowMaxOk() (*int32, bool)`

GetIRowMaxOk returns a tuple with the IRowMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowMax

`func (o *CommonResponseObjDebugPayloadGetList) SetIRowMax(v int32)`

SetIRowMax sets IRowMax field to given value.


### GetIRowOffset

`func (o *CommonResponseObjDebugPayloadGetList) GetIRowOffset() int32`

GetIRowOffset returns the IRowOffset field if non-nil, zero value otherwise.

### GetIRowOffsetOk

`func (o *CommonResponseObjDebugPayloadGetList) GetIRowOffsetOk() (*int32, bool)`

GetIRowOffsetOk returns a tuple with the IRowOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowOffset

`func (o *CommonResponseObjDebugPayloadGetList) SetIRowOffset(v int32)`

SetIRowOffset sets IRowOffset field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


