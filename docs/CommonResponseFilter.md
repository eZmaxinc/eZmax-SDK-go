# CommonResponseFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AAutoType** | Pointer to **map[string]string** | List of filters that can be used in *sFilter* (Automatic types) | [optional] 
**AAutoTypeHaving** | Pointer to **map[string]string** | List of computed filters that can be used in *sFilter* (Automatic types) | [optional] 
**AEnum** | Pointer to **map[string]map[string]string** | List of filters that can be used in *sFilter* (Enum types) | [optional] 

## Methods

### NewCommonResponseFilter

`func NewCommonResponseFilter() *CommonResponseFilter`

NewCommonResponseFilter instantiates a new CommonResponseFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonResponseFilterWithDefaults

`func NewCommonResponseFilterWithDefaults() *CommonResponseFilter`

NewCommonResponseFilterWithDefaults instantiates a new CommonResponseFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAAutoType

`func (o *CommonResponseFilter) GetAAutoType() map[string]string`

GetAAutoType returns the AAutoType field if non-nil, zero value otherwise.

### GetAAutoTypeOk

`func (o *CommonResponseFilter) GetAAutoTypeOk() (*map[string]string, bool)`

GetAAutoTypeOk returns a tuple with the AAutoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAAutoType

`func (o *CommonResponseFilter) SetAAutoType(v map[string]string)`

SetAAutoType sets AAutoType field to given value.

### HasAAutoType

`func (o *CommonResponseFilter) HasAAutoType() bool`

HasAAutoType returns a boolean if a field has been set.

### GetAAutoTypeHaving

`func (o *CommonResponseFilter) GetAAutoTypeHaving() map[string]string`

GetAAutoTypeHaving returns the AAutoTypeHaving field if non-nil, zero value otherwise.

### GetAAutoTypeHavingOk

`func (o *CommonResponseFilter) GetAAutoTypeHavingOk() (*map[string]string, bool)`

GetAAutoTypeHavingOk returns a tuple with the AAutoTypeHaving field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAAutoTypeHaving

`func (o *CommonResponseFilter) SetAAutoTypeHaving(v map[string]string)`

SetAAutoTypeHaving sets AAutoTypeHaving field to given value.

### HasAAutoTypeHaving

`func (o *CommonResponseFilter) HasAAutoTypeHaving() bool`

HasAAutoTypeHaving returns a boolean if a field has been set.

### GetAEnum

`func (o *CommonResponseFilter) GetAEnum() map[string]map[string]string`

GetAEnum returns the AEnum field if non-nil, zero value otherwise.

### GetAEnumOk

`func (o *CommonResponseFilter) GetAEnumOk() (*map[string]map[string]string, bool)`

GetAEnumOk returns a tuple with the AEnum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAEnum

`func (o *CommonResponseFilter) SetAEnum(v map[string]map[string]string)`

SetAEnum sets AEnum field to given value.

### HasAEnum

`func (o *CommonResponseFilter) HasAEnum() bool`

HasAEnum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


