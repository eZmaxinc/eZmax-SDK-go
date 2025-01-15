# CustomEzsignformfieldgroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsignformfieldgroupID** | Pointer to **int32** | The unique ID of the Ezsignformfieldgroup | [optional] 
**SEzsignformfieldgroupLabel** | Pointer to **string** | The Label for the Ezsignformfieldgroup | [optional] 
**AObjEzsignformfield** | **[]map[string]interface{}** | An array containing all the values to fill the Ezsignform. | 

## Methods

### NewCustomEzsignformfieldgroupRequest

`func NewCustomEzsignformfieldgroupRequest(aObjEzsignformfield []CustomEzsignformfieldRequest, ) *CustomEzsignformfieldgroupRequest`

NewCustomEzsignformfieldgroupRequest instantiates a new CustomEzsignformfieldgroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomEzsignformfieldgroupRequestWithDefaults

`func NewCustomEzsignformfieldgroupRequestWithDefaults() *CustomEzsignformfieldgroupRequest`

NewCustomEzsignformfieldgroupRequestWithDefaults instantiates a new CustomEzsignformfieldgroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsignformfieldgroupID

`func (o *CustomEzsignformfieldgroupRequest) GetPkiEzsignformfieldgroupID() int32`

GetPkiEzsignformfieldgroupID returns the PkiEzsignformfieldgroupID field if non-nil, zero value otherwise.

### GetPkiEzsignformfieldgroupIDOk

`func (o *CustomEzsignformfieldgroupRequest) GetPkiEzsignformfieldgroupIDOk() (*int32, bool)`

GetPkiEzsignformfieldgroupIDOk returns a tuple with the PkiEzsignformfieldgroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsignformfieldgroupID

`func (o *CustomEzsignformfieldgroupRequest) SetPkiEzsignformfieldgroupID(v int32)`

SetPkiEzsignformfieldgroupID sets PkiEzsignformfieldgroupID field to given value.

### HasPkiEzsignformfieldgroupID

`func (o *CustomEzsignformfieldgroupRequest) HasPkiEzsignformfieldgroupID() bool`

HasPkiEzsignformfieldgroupID returns a boolean if a field has been set.

### GetSEzsignformfieldgroupLabel

`func (o *CustomEzsignformfieldgroupRequest) GetSEzsignformfieldgroupLabel() string`

GetSEzsignformfieldgroupLabel returns the SEzsignformfieldgroupLabel field if non-nil, zero value otherwise.

### GetSEzsignformfieldgroupLabelOk

`func (o *CustomEzsignformfieldgroupRequest) GetSEzsignformfieldgroupLabelOk() (*string, bool)`

GetSEzsignformfieldgroupLabelOk returns a tuple with the SEzsignformfieldgroupLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignformfieldgroupLabel

`func (o *CustomEzsignformfieldgroupRequest) SetSEzsignformfieldgroupLabel(v string)`

SetSEzsignformfieldgroupLabel sets SEzsignformfieldgroupLabel field to given value.

### HasSEzsignformfieldgroupLabel

`func (o *CustomEzsignformfieldgroupRequest) HasSEzsignformfieldgroupLabel() bool`

HasSEzsignformfieldgroupLabel returns a boolean if a field has been set.

### GetAObjEzsignformfield

`func (o *CustomEzsignformfieldgroupRequest) GetAObjEzsignformfield() []CustomEzsignformfieldRequest`

GetAObjEzsignformfield returns the AObjEzsignformfield field if non-nil, zero value otherwise.

### GetAObjEzsignformfieldOk

`func (o *CustomEzsignformfieldgroupRequest) GetAObjEzsignformfieldOk() (*[]CustomEzsignformfieldRequest, bool)`

GetAObjEzsignformfieldOk returns a tuple with the AObjEzsignformfield field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzsignformfield

`func (o *CustomEzsignformfieldgroupRequest) SetAObjEzsignformfield(v []CustomEzsignformfieldRequest)`

SetAObjEzsignformfield sets AObjEzsignformfield field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


