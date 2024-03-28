# SystemconfigurationResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiSystemconfigurationID** | **int32** | The unique ID of the Systemconfiguration | 
**FkiSystemconfigurationtypeID** | **int32** | The unique ID of the Systemconfigurationtype | 
**SSystemconfigurationtypeDescriptionX** | **string** | The description of the Systemconfigurationtype in the language of the requester | 
**ESystemconfigurationNewexternaluseraction** | [**FieldESystemconfigurationNewexternaluseraction**](FieldESystemconfigurationNewexternaluseraction.md) |  | 
**ESystemconfigurationLanguage1** | [**FieldESystemconfigurationLanguage1**](FieldESystemconfigurationLanguage1.md) |  | 
**ESystemconfigurationLanguage2** | [**FieldESystemconfigurationLanguage2**](FieldESystemconfigurationLanguage2.md) |  | 
**ESystemconfigurationEzsign** | Pointer to [**FieldESystemconfigurationEzsign**](FieldESystemconfigurationEzsign.md) |  | [optional] 
**ESystemconfigurationEzsignofficeplan** | Pointer to [**FieldESystemconfigurationEzsignofficeplan**](FieldESystemconfigurationEzsignofficeplan.md) |  | [optional] 
**BSystemconfigurationEzsignpaidbyoffice** | Pointer to **bool** | Whether if Ezsign is paid by the company or not | [optional] 
**BSystemconfigurationEzsignpersonnal** | **bool** | Whether if we allow the creation of personal files in eZsign | 
**BSystemconfigurationIsdisposalactive** | Pointer to **bool** | Whether is Disposal processus is active or not | [optional] 
**BSystemconfigurationSspr** | **bool** | Whether if we allow SSPR | 
**DtSystemconfigurationReadonlyexpirationstart** | Pointer to **string** | The start date where the system will be in read only | [optional] 
**DtSystemconfigurationReadonlyexpirationend** | Pointer to **string** | The end date where the system will be in read only | [optional] 

## Methods

### NewSystemconfigurationResponseCompound

`func NewSystemconfigurationResponseCompound(pkiSystemconfigurationID int32, fkiSystemconfigurationtypeID int32, sSystemconfigurationtypeDescriptionX string, eSystemconfigurationNewexternaluseraction FieldESystemconfigurationNewexternaluseraction, eSystemconfigurationLanguage1 FieldESystemconfigurationLanguage1, eSystemconfigurationLanguage2 FieldESystemconfigurationLanguage2, bSystemconfigurationEzsignpersonnal bool, bSystemconfigurationSspr bool, ) *SystemconfigurationResponseCompound`

NewSystemconfigurationResponseCompound instantiates a new SystemconfigurationResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemconfigurationResponseCompoundWithDefaults

`func NewSystemconfigurationResponseCompoundWithDefaults() *SystemconfigurationResponseCompound`

NewSystemconfigurationResponseCompoundWithDefaults instantiates a new SystemconfigurationResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiSystemconfigurationID

`func (o *SystemconfigurationResponseCompound) GetPkiSystemconfigurationID() int32`

GetPkiSystemconfigurationID returns the PkiSystemconfigurationID field if non-nil, zero value otherwise.

### GetPkiSystemconfigurationIDOk

`func (o *SystemconfigurationResponseCompound) GetPkiSystemconfigurationIDOk() (*int32, bool)`

GetPkiSystemconfigurationIDOk returns a tuple with the PkiSystemconfigurationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiSystemconfigurationID

`func (o *SystemconfigurationResponseCompound) SetPkiSystemconfigurationID(v int32)`

SetPkiSystemconfigurationID sets PkiSystemconfigurationID field to given value.


### GetFkiSystemconfigurationtypeID

`func (o *SystemconfigurationResponseCompound) GetFkiSystemconfigurationtypeID() int32`

GetFkiSystemconfigurationtypeID returns the FkiSystemconfigurationtypeID field if non-nil, zero value otherwise.

### GetFkiSystemconfigurationtypeIDOk

`func (o *SystemconfigurationResponseCompound) GetFkiSystemconfigurationtypeIDOk() (*int32, bool)`

GetFkiSystemconfigurationtypeIDOk returns a tuple with the FkiSystemconfigurationtypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiSystemconfigurationtypeID

`func (o *SystemconfigurationResponseCompound) SetFkiSystemconfigurationtypeID(v int32)`

SetFkiSystemconfigurationtypeID sets FkiSystemconfigurationtypeID field to given value.


### GetSSystemconfigurationtypeDescriptionX

`func (o *SystemconfigurationResponseCompound) GetSSystemconfigurationtypeDescriptionX() string`

GetSSystemconfigurationtypeDescriptionX returns the SSystemconfigurationtypeDescriptionX field if non-nil, zero value otherwise.

### GetSSystemconfigurationtypeDescriptionXOk

`func (o *SystemconfigurationResponseCompound) GetSSystemconfigurationtypeDescriptionXOk() (*string, bool)`

GetSSystemconfigurationtypeDescriptionXOk returns a tuple with the SSystemconfigurationtypeDescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSSystemconfigurationtypeDescriptionX

`func (o *SystemconfigurationResponseCompound) SetSSystemconfigurationtypeDescriptionX(v string)`

SetSSystemconfigurationtypeDescriptionX sets SSystemconfigurationtypeDescriptionX field to given value.


### GetESystemconfigurationNewexternaluseraction

`func (o *SystemconfigurationResponseCompound) GetESystemconfigurationNewexternaluseraction() FieldESystemconfigurationNewexternaluseraction`

GetESystemconfigurationNewexternaluseraction returns the ESystemconfigurationNewexternaluseraction field if non-nil, zero value otherwise.

### GetESystemconfigurationNewexternaluseractionOk

`func (o *SystemconfigurationResponseCompound) GetESystemconfigurationNewexternaluseractionOk() (*FieldESystemconfigurationNewexternaluseraction, bool)`

GetESystemconfigurationNewexternaluseractionOk returns a tuple with the ESystemconfigurationNewexternaluseraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetESystemconfigurationNewexternaluseraction

`func (o *SystemconfigurationResponseCompound) SetESystemconfigurationNewexternaluseraction(v FieldESystemconfigurationNewexternaluseraction)`

SetESystemconfigurationNewexternaluseraction sets ESystemconfigurationNewexternaluseraction field to given value.


### GetESystemconfigurationLanguage1

`func (o *SystemconfigurationResponseCompound) GetESystemconfigurationLanguage1() FieldESystemconfigurationLanguage1`

GetESystemconfigurationLanguage1 returns the ESystemconfigurationLanguage1 field if non-nil, zero value otherwise.

### GetESystemconfigurationLanguage1Ok

`func (o *SystemconfigurationResponseCompound) GetESystemconfigurationLanguage1Ok() (*FieldESystemconfigurationLanguage1, bool)`

GetESystemconfigurationLanguage1Ok returns a tuple with the ESystemconfigurationLanguage1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetESystemconfigurationLanguage1

`func (o *SystemconfigurationResponseCompound) SetESystemconfigurationLanguage1(v FieldESystemconfigurationLanguage1)`

SetESystemconfigurationLanguage1 sets ESystemconfigurationLanguage1 field to given value.


### GetESystemconfigurationLanguage2

`func (o *SystemconfigurationResponseCompound) GetESystemconfigurationLanguage2() FieldESystemconfigurationLanguage2`

GetESystemconfigurationLanguage2 returns the ESystemconfigurationLanguage2 field if non-nil, zero value otherwise.

### GetESystemconfigurationLanguage2Ok

`func (o *SystemconfigurationResponseCompound) GetESystemconfigurationLanguage2Ok() (*FieldESystemconfigurationLanguage2, bool)`

GetESystemconfigurationLanguage2Ok returns a tuple with the ESystemconfigurationLanguage2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetESystemconfigurationLanguage2

`func (o *SystemconfigurationResponseCompound) SetESystemconfigurationLanguage2(v FieldESystemconfigurationLanguage2)`

SetESystemconfigurationLanguage2 sets ESystemconfigurationLanguage2 field to given value.


### GetESystemconfigurationEzsign

`func (o *SystemconfigurationResponseCompound) GetESystemconfigurationEzsign() FieldESystemconfigurationEzsign`

GetESystemconfigurationEzsign returns the ESystemconfigurationEzsign field if non-nil, zero value otherwise.

### GetESystemconfigurationEzsignOk

`func (o *SystemconfigurationResponseCompound) GetESystemconfigurationEzsignOk() (*FieldESystemconfigurationEzsign, bool)`

GetESystemconfigurationEzsignOk returns a tuple with the ESystemconfigurationEzsign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetESystemconfigurationEzsign

`func (o *SystemconfigurationResponseCompound) SetESystemconfigurationEzsign(v FieldESystemconfigurationEzsign)`

SetESystemconfigurationEzsign sets ESystemconfigurationEzsign field to given value.

### HasESystemconfigurationEzsign

`func (o *SystemconfigurationResponseCompound) HasESystemconfigurationEzsign() bool`

HasESystemconfigurationEzsign returns a boolean if a field has been set.

### GetESystemconfigurationEzsignofficeplan

`func (o *SystemconfigurationResponseCompound) GetESystemconfigurationEzsignofficeplan() FieldESystemconfigurationEzsignofficeplan`

GetESystemconfigurationEzsignofficeplan returns the ESystemconfigurationEzsignofficeplan field if non-nil, zero value otherwise.

### GetESystemconfigurationEzsignofficeplanOk

`func (o *SystemconfigurationResponseCompound) GetESystemconfigurationEzsignofficeplanOk() (*FieldESystemconfigurationEzsignofficeplan, bool)`

GetESystemconfigurationEzsignofficeplanOk returns a tuple with the ESystemconfigurationEzsignofficeplan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetESystemconfigurationEzsignofficeplan

`func (o *SystemconfigurationResponseCompound) SetESystemconfigurationEzsignofficeplan(v FieldESystemconfigurationEzsignofficeplan)`

SetESystemconfigurationEzsignofficeplan sets ESystemconfigurationEzsignofficeplan field to given value.

### HasESystemconfigurationEzsignofficeplan

`func (o *SystemconfigurationResponseCompound) HasESystemconfigurationEzsignofficeplan() bool`

HasESystemconfigurationEzsignofficeplan returns a boolean if a field has been set.

### GetBSystemconfigurationEzsignpaidbyoffice

`func (o *SystemconfigurationResponseCompound) GetBSystemconfigurationEzsignpaidbyoffice() bool`

GetBSystemconfigurationEzsignpaidbyoffice returns the BSystemconfigurationEzsignpaidbyoffice field if non-nil, zero value otherwise.

### GetBSystemconfigurationEzsignpaidbyofficeOk

`func (o *SystemconfigurationResponseCompound) GetBSystemconfigurationEzsignpaidbyofficeOk() (*bool, bool)`

GetBSystemconfigurationEzsignpaidbyofficeOk returns a tuple with the BSystemconfigurationEzsignpaidbyoffice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBSystemconfigurationEzsignpaidbyoffice

`func (o *SystemconfigurationResponseCompound) SetBSystemconfigurationEzsignpaidbyoffice(v bool)`

SetBSystemconfigurationEzsignpaidbyoffice sets BSystemconfigurationEzsignpaidbyoffice field to given value.

### HasBSystemconfigurationEzsignpaidbyoffice

`func (o *SystemconfigurationResponseCompound) HasBSystemconfigurationEzsignpaidbyoffice() bool`

HasBSystemconfigurationEzsignpaidbyoffice returns a boolean if a field has been set.

### GetBSystemconfigurationEzsignpersonnal

`func (o *SystemconfigurationResponseCompound) GetBSystemconfigurationEzsignpersonnal() bool`

GetBSystemconfigurationEzsignpersonnal returns the BSystemconfigurationEzsignpersonnal field if non-nil, zero value otherwise.

### GetBSystemconfigurationEzsignpersonnalOk

`func (o *SystemconfigurationResponseCompound) GetBSystemconfigurationEzsignpersonnalOk() (*bool, bool)`

GetBSystemconfigurationEzsignpersonnalOk returns a tuple with the BSystemconfigurationEzsignpersonnal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBSystemconfigurationEzsignpersonnal

`func (o *SystemconfigurationResponseCompound) SetBSystemconfigurationEzsignpersonnal(v bool)`

SetBSystemconfigurationEzsignpersonnal sets BSystemconfigurationEzsignpersonnal field to given value.


### GetBSystemconfigurationIsdisposalactive

`func (o *SystemconfigurationResponseCompound) GetBSystemconfigurationIsdisposalactive() bool`

GetBSystemconfigurationIsdisposalactive returns the BSystemconfigurationIsdisposalactive field if non-nil, zero value otherwise.

### GetBSystemconfigurationIsdisposalactiveOk

`func (o *SystemconfigurationResponseCompound) GetBSystemconfigurationIsdisposalactiveOk() (*bool, bool)`

GetBSystemconfigurationIsdisposalactiveOk returns a tuple with the BSystemconfigurationIsdisposalactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBSystemconfigurationIsdisposalactive

`func (o *SystemconfigurationResponseCompound) SetBSystemconfigurationIsdisposalactive(v bool)`

SetBSystemconfigurationIsdisposalactive sets BSystemconfigurationIsdisposalactive field to given value.

### HasBSystemconfigurationIsdisposalactive

`func (o *SystemconfigurationResponseCompound) HasBSystemconfigurationIsdisposalactive() bool`

HasBSystemconfigurationIsdisposalactive returns a boolean if a field has been set.

### GetBSystemconfigurationSspr

`func (o *SystemconfigurationResponseCompound) GetBSystemconfigurationSspr() bool`

GetBSystemconfigurationSspr returns the BSystemconfigurationSspr field if non-nil, zero value otherwise.

### GetBSystemconfigurationSsprOk

`func (o *SystemconfigurationResponseCompound) GetBSystemconfigurationSsprOk() (*bool, bool)`

GetBSystemconfigurationSsprOk returns a tuple with the BSystemconfigurationSspr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBSystemconfigurationSspr

`func (o *SystemconfigurationResponseCompound) SetBSystemconfigurationSspr(v bool)`

SetBSystemconfigurationSspr sets BSystemconfigurationSspr field to given value.


### GetDtSystemconfigurationReadonlyexpirationstart

`func (o *SystemconfigurationResponseCompound) GetDtSystemconfigurationReadonlyexpirationstart() string`

GetDtSystemconfigurationReadonlyexpirationstart returns the DtSystemconfigurationReadonlyexpirationstart field if non-nil, zero value otherwise.

### GetDtSystemconfigurationReadonlyexpirationstartOk

`func (o *SystemconfigurationResponseCompound) GetDtSystemconfigurationReadonlyexpirationstartOk() (*string, bool)`

GetDtSystemconfigurationReadonlyexpirationstartOk returns a tuple with the DtSystemconfigurationReadonlyexpirationstart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtSystemconfigurationReadonlyexpirationstart

`func (o *SystemconfigurationResponseCompound) SetDtSystemconfigurationReadonlyexpirationstart(v string)`

SetDtSystemconfigurationReadonlyexpirationstart sets DtSystemconfigurationReadonlyexpirationstart field to given value.

### HasDtSystemconfigurationReadonlyexpirationstart

`func (o *SystemconfigurationResponseCompound) HasDtSystemconfigurationReadonlyexpirationstart() bool`

HasDtSystemconfigurationReadonlyexpirationstart returns a boolean if a field has been set.

### GetDtSystemconfigurationReadonlyexpirationend

`func (o *SystemconfigurationResponseCompound) GetDtSystemconfigurationReadonlyexpirationend() string`

GetDtSystemconfigurationReadonlyexpirationend returns the DtSystemconfigurationReadonlyexpirationend field if non-nil, zero value otherwise.

### GetDtSystemconfigurationReadonlyexpirationendOk

`func (o *SystemconfigurationResponseCompound) GetDtSystemconfigurationReadonlyexpirationendOk() (*string, bool)`

GetDtSystemconfigurationReadonlyexpirationendOk returns a tuple with the DtSystemconfigurationReadonlyexpirationend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtSystemconfigurationReadonlyexpirationend

`func (o *SystemconfigurationResponseCompound) SetDtSystemconfigurationReadonlyexpirationend(v string)`

SetDtSystemconfigurationReadonlyexpirationend sets DtSystemconfigurationReadonlyexpirationend field to given value.

### HasDtSystemconfigurationReadonlyexpirationend

`func (o *SystemconfigurationResponseCompound) HasDtSystemconfigurationReadonlyexpirationend() bool`

HasDtSystemconfigurationReadonlyexpirationend returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


