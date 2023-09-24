# SystemconfigurationRequestCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiSystemconfigurationID** | Pointer to **int32** | The unique ID of the Systemconfiguration | [optional] 
**ESystemconfigurationNewexternaluseraction** | [**FieldESystemconfigurationNewexternaluseraction**](FieldESystemconfigurationNewexternaluseraction.md) |  | 
**ESystemconfigurationLanguage1** | [**FieldESystemconfigurationLanguage1**](FieldESystemconfigurationLanguage1.md) |  | 
**ESystemconfigurationLanguage2** | [**FieldESystemconfigurationLanguage2**](FieldESystemconfigurationLanguage2.md) |  | 
**ESystemconfigurationEzsign** | Pointer to [**FieldESystemconfigurationEzsign**](FieldESystemconfigurationEzsign.md) |  | [optional] 
**BSystemconfigurationEzsignpersonnal** | **bool** | Whether if we allow the creation of personal files in eZsign | 
**BSystemconfigurationSspr** | **bool** | Whether if we allow SSPR | 
**DtSystemconfigurationReadonlyexpirationstart** | Pointer to **string** | The start date where the system will be in read only | [optional] 
**DtSystemconfigurationReadonlyexpirationend** | Pointer to **string** | The end date where the system will be in read only | [optional] 

## Methods

### NewSystemconfigurationRequestCompound

`func NewSystemconfigurationRequestCompound(eSystemconfigurationNewexternaluseraction FieldESystemconfigurationNewexternaluseraction, eSystemconfigurationLanguage1 FieldESystemconfigurationLanguage1, eSystemconfigurationLanguage2 FieldESystemconfigurationLanguage2, bSystemconfigurationEzsignpersonnal bool, bSystemconfigurationSspr bool, ) *SystemconfigurationRequestCompound`

NewSystemconfigurationRequestCompound instantiates a new SystemconfigurationRequestCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemconfigurationRequestCompoundWithDefaults

`func NewSystemconfigurationRequestCompoundWithDefaults() *SystemconfigurationRequestCompound`

NewSystemconfigurationRequestCompoundWithDefaults instantiates a new SystemconfigurationRequestCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiSystemconfigurationID

`func (o *SystemconfigurationRequestCompound) GetPkiSystemconfigurationID() int32`

GetPkiSystemconfigurationID returns the PkiSystemconfigurationID field if non-nil, zero value otherwise.

### GetPkiSystemconfigurationIDOk

`func (o *SystemconfigurationRequestCompound) GetPkiSystemconfigurationIDOk() (*int32, bool)`

GetPkiSystemconfigurationIDOk returns a tuple with the PkiSystemconfigurationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiSystemconfigurationID

`func (o *SystemconfigurationRequestCompound) SetPkiSystemconfigurationID(v int32)`

SetPkiSystemconfigurationID sets PkiSystemconfigurationID field to given value.

### HasPkiSystemconfigurationID

`func (o *SystemconfigurationRequestCompound) HasPkiSystemconfigurationID() bool`

HasPkiSystemconfigurationID returns a boolean if a field has been set.

### GetESystemconfigurationNewexternaluseraction

`func (o *SystemconfigurationRequestCompound) GetESystemconfigurationNewexternaluseraction() FieldESystemconfigurationNewexternaluseraction`

GetESystemconfigurationNewexternaluseraction returns the ESystemconfigurationNewexternaluseraction field if non-nil, zero value otherwise.

### GetESystemconfigurationNewexternaluseractionOk

`func (o *SystemconfigurationRequestCompound) GetESystemconfigurationNewexternaluseractionOk() (*FieldESystemconfigurationNewexternaluseraction, bool)`

GetESystemconfigurationNewexternaluseractionOk returns a tuple with the ESystemconfigurationNewexternaluseraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetESystemconfigurationNewexternaluseraction

`func (o *SystemconfigurationRequestCompound) SetESystemconfigurationNewexternaluseraction(v FieldESystemconfigurationNewexternaluseraction)`

SetESystemconfigurationNewexternaluseraction sets ESystemconfigurationNewexternaluseraction field to given value.


### GetESystemconfigurationLanguage1

`func (o *SystemconfigurationRequestCompound) GetESystemconfigurationLanguage1() FieldESystemconfigurationLanguage1`

GetESystemconfigurationLanguage1 returns the ESystemconfigurationLanguage1 field if non-nil, zero value otherwise.

### GetESystemconfigurationLanguage1Ok

`func (o *SystemconfigurationRequestCompound) GetESystemconfigurationLanguage1Ok() (*FieldESystemconfigurationLanguage1, bool)`

GetESystemconfigurationLanguage1Ok returns a tuple with the ESystemconfigurationLanguage1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetESystemconfigurationLanguage1

`func (o *SystemconfigurationRequestCompound) SetESystemconfigurationLanguage1(v FieldESystemconfigurationLanguage1)`

SetESystemconfigurationLanguage1 sets ESystemconfigurationLanguage1 field to given value.


### GetESystemconfigurationLanguage2

`func (o *SystemconfigurationRequestCompound) GetESystemconfigurationLanguage2() FieldESystemconfigurationLanguage2`

GetESystemconfigurationLanguage2 returns the ESystemconfigurationLanguage2 field if non-nil, zero value otherwise.

### GetESystemconfigurationLanguage2Ok

`func (o *SystemconfigurationRequestCompound) GetESystemconfigurationLanguage2Ok() (*FieldESystemconfigurationLanguage2, bool)`

GetESystemconfigurationLanguage2Ok returns a tuple with the ESystemconfigurationLanguage2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetESystemconfigurationLanguage2

`func (o *SystemconfigurationRequestCompound) SetESystemconfigurationLanguage2(v FieldESystemconfigurationLanguage2)`

SetESystemconfigurationLanguage2 sets ESystemconfigurationLanguage2 field to given value.


### GetESystemconfigurationEzsign

`func (o *SystemconfigurationRequestCompound) GetESystemconfigurationEzsign() FieldESystemconfigurationEzsign`

GetESystemconfigurationEzsign returns the ESystemconfigurationEzsign field if non-nil, zero value otherwise.

### GetESystemconfigurationEzsignOk

`func (o *SystemconfigurationRequestCompound) GetESystemconfigurationEzsignOk() (*FieldESystemconfigurationEzsign, bool)`

GetESystemconfigurationEzsignOk returns a tuple with the ESystemconfigurationEzsign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetESystemconfigurationEzsign

`func (o *SystemconfigurationRequestCompound) SetESystemconfigurationEzsign(v FieldESystemconfigurationEzsign)`

SetESystemconfigurationEzsign sets ESystemconfigurationEzsign field to given value.

### HasESystemconfigurationEzsign

`func (o *SystemconfigurationRequestCompound) HasESystemconfigurationEzsign() bool`

HasESystemconfigurationEzsign returns a boolean if a field has been set.

### GetBSystemconfigurationEzsignpersonnal

`func (o *SystemconfigurationRequestCompound) GetBSystemconfigurationEzsignpersonnal() bool`

GetBSystemconfigurationEzsignpersonnal returns the BSystemconfigurationEzsignpersonnal field if non-nil, zero value otherwise.

### GetBSystemconfigurationEzsignpersonnalOk

`func (o *SystemconfigurationRequestCompound) GetBSystemconfigurationEzsignpersonnalOk() (*bool, bool)`

GetBSystemconfigurationEzsignpersonnalOk returns a tuple with the BSystemconfigurationEzsignpersonnal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBSystemconfigurationEzsignpersonnal

`func (o *SystemconfigurationRequestCompound) SetBSystemconfigurationEzsignpersonnal(v bool)`

SetBSystemconfigurationEzsignpersonnal sets BSystemconfigurationEzsignpersonnal field to given value.


### GetBSystemconfigurationSspr

`func (o *SystemconfigurationRequestCompound) GetBSystemconfigurationSspr() bool`

GetBSystemconfigurationSspr returns the BSystemconfigurationSspr field if non-nil, zero value otherwise.

### GetBSystemconfigurationSsprOk

`func (o *SystemconfigurationRequestCompound) GetBSystemconfigurationSsprOk() (*bool, bool)`

GetBSystemconfigurationSsprOk returns a tuple with the BSystemconfigurationSspr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBSystemconfigurationSspr

`func (o *SystemconfigurationRequestCompound) SetBSystemconfigurationSspr(v bool)`

SetBSystemconfigurationSspr sets BSystemconfigurationSspr field to given value.


### GetDtSystemconfigurationReadonlyexpirationstart

`func (o *SystemconfigurationRequestCompound) GetDtSystemconfigurationReadonlyexpirationstart() string`

GetDtSystemconfigurationReadonlyexpirationstart returns the DtSystemconfigurationReadonlyexpirationstart field if non-nil, zero value otherwise.

### GetDtSystemconfigurationReadonlyexpirationstartOk

`func (o *SystemconfigurationRequestCompound) GetDtSystemconfigurationReadonlyexpirationstartOk() (*string, bool)`

GetDtSystemconfigurationReadonlyexpirationstartOk returns a tuple with the DtSystemconfigurationReadonlyexpirationstart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtSystemconfigurationReadonlyexpirationstart

`func (o *SystemconfigurationRequestCompound) SetDtSystemconfigurationReadonlyexpirationstart(v string)`

SetDtSystemconfigurationReadonlyexpirationstart sets DtSystemconfigurationReadonlyexpirationstart field to given value.

### HasDtSystemconfigurationReadonlyexpirationstart

`func (o *SystemconfigurationRequestCompound) HasDtSystemconfigurationReadonlyexpirationstart() bool`

HasDtSystemconfigurationReadonlyexpirationstart returns a boolean if a field has been set.

### GetDtSystemconfigurationReadonlyexpirationend

`func (o *SystemconfigurationRequestCompound) GetDtSystemconfigurationReadonlyexpirationend() string`

GetDtSystemconfigurationReadonlyexpirationend returns the DtSystemconfigurationReadonlyexpirationend field if non-nil, zero value otherwise.

### GetDtSystemconfigurationReadonlyexpirationendOk

`func (o *SystemconfigurationRequestCompound) GetDtSystemconfigurationReadonlyexpirationendOk() (*string, bool)`

GetDtSystemconfigurationReadonlyexpirationendOk returns a tuple with the DtSystemconfigurationReadonlyexpirationend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtSystemconfigurationReadonlyexpirationend

`func (o *SystemconfigurationRequestCompound) SetDtSystemconfigurationReadonlyexpirationend(v string)`

SetDtSystemconfigurationReadonlyexpirationend sets DtSystemconfigurationReadonlyexpirationend field to given value.

### HasDtSystemconfigurationReadonlyexpirationend

`func (o *SystemconfigurationRequestCompound) HasDtSystemconfigurationReadonlyexpirationend() bool`

HasDtSystemconfigurationReadonlyexpirationend returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


