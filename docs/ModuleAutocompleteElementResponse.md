# ModuleAutocompleteElementResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiModuleID** | **int32** | The unique ID of the Module | 
**SModuleNameX** | **string** | The Name of the Module in the language of the requester | 
**BModuleIsactive** | **bool** | Whether the Module is active or not | 

## Methods

### NewModuleAutocompleteElementResponse

`func NewModuleAutocompleteElementResponse(pkiModuleID int32, sModuleNameX string, bModuleIsactive bool, ) *ModuleAutocompleteElementResponse`

NewModuleAutocompleteElementResponse instantiates a new ModuleAutocompleteElementResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModuleAutocompleteElementResponseWithDefaults

`func NewModuleAutocompleteElementResponseWithDefaults() *ModuleAutocompleteElementResponse`

NewModuleAutocompleteElementResponseWithDefaults instantiates a new ModuleAutocompleteElementResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiModuleID

`func (o *ModuleAutocompleteElementResponse) GetPkiModuleID() int32`

GetPkiModuleID returns the PkiModuleID field if non-nil, zero value otherwise.

### GetPkiModuleIDOk

`func (o *ModuleAutocompleteElementResponse) GetPkiModuleIDOk() (*int32, bool)`

GetPkiModuleIDOk returns a tuple with the PkiModuleID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiModuleID

`func (o *ModuleAutocompleteElementResponse) SetPkiModuleID(v int32)`

SetPkiModuleID sets PkiModuleID field to given value.


### GetSModuleNameX

`func (o *ModuleAutocompleteElementResponse) GetSModuleNameX() string`

GetSModuleNameX returns the SModuleNameX field if non-nil, zero value otherwise.

### GetSModuleNameXOk

`func (o *ModuleAutocompleteElementResponse) GetSModuleNameXOk() (*string, bool)`

GetSModuleNameXOk returns a tuple with the SModuleNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSModuleNameX

`func (o *ModuleAutocompleteElementResponse) SetSModuleNameX(v string)`

SetSModuleNameX sets SModuleNameX field to given value.


### GetBModuleIsactive

`func (o *ModuleAutocompleteElementResponse) GetBModuleIsactive() bool`

GetBModuleIsactive returns the BModuleIsactive field if non-nil, zero value otherwise.

### GetBModuleIsactiveOk

`func (o *ModuleAutocompleteElementResponse) GetBModuleIsactiveOk() (*bool, bool)`

GetBModuleIsactiveOk returns a tuple with the BModuleIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBModuleIsactive

`func (o *ModuleAutocompleteElementResponse) SetBModuleIsactive(v bool)`

SetBModuleIsactive sets BModuleIsactive field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


