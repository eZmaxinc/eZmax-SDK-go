# SystemconfigurationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiSystemconfigurationID** | **int32** | The unique ID of the Systemconfiguration | 
**FkiSystemconfigurationtypeID** | **int32** | The unique ID of the Systemconfigurationtype | 
**SSystemconfigurationtypeDescriptionX** | **string** | The description of the Systemconfigurationtype in the language of the requester | 
**ESystemconfigurationNewexternaluseraction** | [**FieldESystemconfigurationNewexternaluseraction**](FieldESystemconfigurationNewexternaluseraction.md) |  | 
**ESystemconfigurationLanguage1** | [**FieldESystemconfigurationLanguage1**](FieldESystemconfigurationLanguage1.md) |  | 
**ESystemconfigurationLanguage2** | [**FieldESystemconfigurationLanguage2**](FieldESystemconfigurationLanguage2.md) |  | 
**ESystemconfigurationEzsign** | [**FieldESystemconfigurationEzsign**](FieldESystemconfigurationEzsign.md) |  | 
**BSystemconfigurationEzsignpersonnal** | **bool** | Whether if we allow the creation of personal files in eZsign | 
**BSystemconfigurationSspr** | **bool** | Whether if we allow SSPR | 
**DtSystemconfigurationReadonlyexpirationstart** | Pointer to **string** | The start date where the system will be in read only | [optional] 
**DtSystemconfigurationReadonlyexpirationend** | Pointer to **string** | The end date where the system will be in read only | [optional] 

## Methods

### NewSystemconfigurationResponse

`func NewSystemconfigurationResponse(pkiSystemconfigurationID int32, fkiSystemconfigurationtypeID int32, sSystemconfigurationtypeDescriptionX string, eSystemconfigurationNewexternaluseraction FieldESystemconfigurationNewexternaluseraction, eSystemconfigurationLanguage1 FieldESystemconfigurationLanguage1, eSystemconfigurationLanguage2 FieldESystemconfigurationLanguage2, eSystemconfigurationEzsign FieldESystemconfigurationEzsign, bSystemconfigurationEzsignpersonnal bool, bSystemconfigurationSspr bool, ) *SystemconfigurationResponse`

NewSystemconfigurationResponse instantiates a new SystemconfigurationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemconfigurationResponseWithDefaults

`func NewSystemconfigurationResponseWithDefaults() *SystemconfigurationResponse`

NewSystemconfigurationResponseWithDefaults instantiates a new SystemconfigurationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiSystemconfigurationID

`func (o *SystemconfigurationResponse) GetPkiSystemconfigurationID() int32`

GetPkiSystemconfigurationID returns the PkiSystemconfigurationID field if non-nil, zero value otherwise.

### GetPkiSystemconfigurationIDOk

`func (o *SystemconfigurationResponse) GetPkiSystemconfigurationIDOk() (*int32, bool)`

GetPkiSystemconfigurationIDOk returns a tuple with the PkiSystemconfigurationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiSystemconfigurationID

`func (o *SystemconfigurationResponse) SetPkiSystemconfigurationID(v int32)`

SetPkiSystemconfigurationID sets PkiSystemconfigurationID field to given value.


### GetFkiSystemconfigurationtypeID

`func (o *SystemconfigurationResponse) GetFkiSystemconfigurationtypeID() int32`

GetFkiSystemconfigurationtypeID returns the FkiSystemconfigurationtypeID field if non-nil, zero value otherwise.

### GetFkiSystemconfigurationtypeIDOk

`func (o *SystemconfigurationResponse) GetFkiSystemconfigurationtypeIDOk() (*int32, bool)`

GetFkiSystemconfigurationtypeIDOk returns a tuple with the FkiSystemconfigurationtypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiSystemconfigurationtypeID

`func (o *SystemconfigurationResponse) SetFkiSystemconfigurationtypeID(v int32)`

SetFkiSystemconfigurationtypeID sets FkiSystemconfigurationtypeID field to given value.


### GetSSystemconfigurationtypeDescriptionX

`func (o *SystemconfigurationResponse) GetSSystemconfigurationtypeDescriptionX() string`

GetSSystemconfigurationtypeDescriptionX returns the SSystemconfigurationtypeDescriptionX field if non-nil, zero value otherwise.

### GetSSystemconfigurationtypeDescriptionXOk

`func (o *SystemconfigurationResponse) GetSSystemconfigurationtypeDescriptionXOk() (*string, bool)`

GetSSystemconfigurationtypeDescriptionXOk returns a tuple with the SSystemconfigurationtypeDescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSSystemconfigurationtypeDescriptionX

`func (o *SystemconfigurationResponse) SetSSystemconfigurationtypeDescriptionX(v string)`

SetSSystemconfigurationtypeDescriptionX sets SSystemconfigurationtypeDescriptionX field to given value.


### GetESystemconfigurationNewexternaluseraction

`func (o *SystemconfigurationResponse) GetESystemconfigurationNewexternaluseraction() FieldESystemconfigurationNewexternaluseraction`

GetESystemconfigurationNewexternaluseraction returns the ESystemconfigurationNewexternaluseraction field if non-nil, zero value otherwise.

### GetESystemconfigurationNewexternaluseractionOk

`func (o *SystemconfigurationResponse) GetESystemconfigurationNewexternaluseractionOk() (*FieldESystemconfigurationNewexternaluseraction, bool)`

GetESystemconfigurationNewexternaluseractionOk returns a tuple with the ESystemconfigurationNewexternaluseraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetESystemconfigurationNewexternaluseraction

`func (o *SystemconfigurationResponse) SetESystemconfigurationNewexternaluseraction(v FieldESystemconfigurationNewexternaluseraction)`

SetESystemconfigurationNewexternaluseraction sets ESystemconfigurationNewexternaluseraction field to given value.


### GetESystemconfigurationLanguage1

`func (o *SystemconfigurationResponse) GetESystemconfigurationLanguage1() FieldESystemconfigurationLanguage1`

GetESystemconfigurationLanguage1 returns the ESystemconfigurationLanguage1 field if non-nil, zero value otherwise.

### GetESystemconfigurationLanguage1Ok

`func (o *SystemconfigurationResponse) GetESystemconfigurationLanguage1Ok() (*FieldESystemconfigurationLanguage1, bool)`

GetESystemconfigurationLanguage1Ok returns a tuple with the ESystemconfigurationLanguage1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetESystemconfigurationLanguage1

`func (o *SystemconfigurationResponse) SetESystemconfigurationLanguage1(v FieldESystemconfigurationLanguage1)`

SetESystemconfigurationLanguage1 sets ESystemconfigurationLanguage1 field to given value.


### GetESystemconfigurationLanguage2

`func (o *SystemconfigurationResponse) GetESystemconfigurationLanguage2() FieldESystemconfigurationLanguage2`

GetESystemconfigurationLanguage2 returns the ESystemconfigurationLanguage2 field if non-nil, zero value otherwise.

### GetESystemconfigurationLanguage2Ok

`func (o *SystemconfigurationResponse) GetESystemconfigurationLanguage2Ok() (*FieldESystemconfigurationLanguage2, bool)`

GetESystemconfigurationLanguage2Ok returns a tuple with the ESystemconfigurationLanguage2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetESystemconfigurationLanguage2

`func (o *SystemconfigurationResponse) SetESystemconfigurationLanguage2(v FieldESystemconfigurationLanguage2)`

SetESystemconfigurationLanguage2 sets ESystemconfigurationLanguage2 field to given value.


### GetESystemconfigurationEzsign

`func (o *SystemconfigurationResponse) GetESystemconfigurationEzsign() FieldESystemconfigurationEzsign`

GetESystemconfigurationEzsign returns the ESystemconfigurationEzsign field if non-nil, zero value otherwise.

### GetESystemconfigurationEzsignOk

`func (o *SystemconfigurationResponse) GetESystemconfigurationEzsignOk() (*FieldESystemconfigurationEzsign, bool)`

GetESystemconfigurationEzsignOk returns a tuple with the ESystemconfigurationEzsign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetESystemconfigurationEzsign

`func (o *SystemconfigurationResponse) SetESystemconfigurationEzsign(v FieldESystemconfigurationEzsign)`

SetESystemconfigurationEzsign sets ESystemconfigurationEzsign field to given value.


### GetBSystemconfigurationEzsignpersonnal

`func (o *SystemconfigurationResponse) GetBSystemconfigurationEzsignpersonnal() bool`

GetBSystemconfigurationEzsignpersonnal returns the BSystemconfigurationEzsignpersonnal field if non-nil, zero value otherwise.

### GetBSystemconfigurationEzsignpersonnalOk

`func (o *SystemconfigurationResponse) GetBSystemconfigurationEzsignpersonnalOk() (*bool, bool)`

GetBSystemconfigurationEzsignpersonnalOk returns a tuple with the BSystemconfigurationEzsignpersonnal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBSystemconfigurationEzsignpersonnal

`func (o *SystemconfigurationResponse) SetBSystemconfigurationEzsignpersonnal(v bool)`

SetBSystemconfigurationEzsignpersonnal sets BSystemconfigurationEzsignpersonnal field to given value.


### GetBSystemconfigurationSspr

`func (o *SystemconfigurationResponse) GetBSystemconfigurationSspr() bool`

GetBSystemconfigurationSspr returns the BSystemconfigurationSspr field if non-nil, zero value otherwise.

### GetBSystemconfigurationSsprOk

`func (o *SystemconfigurationResponse) GetBSystemconfigurationSsprOk() (*bool, bool)`

GetBSystemconfigurationSsprOk returns a tuple with the BSystemconfigurationSspr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBSystemconfigurationSspr

`func (o *SystemconfigurationResponse) SetBSystemconfigurationSspr(v bool)`

SetBSystemconfigurationSspr sets BSystemconfigurationSspr field to given value.


### GetDtSystemconfigurationReadonlyexpirationstart

`func (o *SystemconfigurationResponse) GetDtSystemconfigurationReadonlyexpirationstart() string`

GetDtSystemconfigurationReadonlyexpirationstart returns the DtSystemconfigurationReadonlyexpirationstart field if non-nil, zero value otherwise.

### GetDtSystemconfigurationReadonlyexpirationstartOk

`func (o *SystemconfigurationResponse) GetDtSystemconfigurationReadonlyexpirationstartOk() (*string, bool)`

GetDtSystemconfigurationReadonlyexpirationstartOk returns a tuple with the DtSystemconfigurationReadonlyexpirationstart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtSystemconfigurationReadonlyexpirationstart

`func (o *SystemconfigurationResponse) SetDtSystemconfigurationReadonlyexpirationstart(v string)`

SetDtSystemconfigurationReadonlyexpirationstart sets DtSystemconfigurationReadonlyexpirationstart field to given value.

### HasDtSystemconfigurationReadonlyexpirationstart

`func (o *SystemconfigurationResponse) HasDtSystemconfigurationReadonlyexpirationstart() bool`

HasDtSystemconfigurationReadonlyexpirationstart returns a boolean if a field has been set.

### GetDtSystemconfigurationReadonlyexpirationend

`func (o *SystemconfigurationResponse) GetDtSystemconfigurationReadonlyexpirationend() string`

GetDtSystemconfigurationReadonlyexpirationend returns the DtSystemconfigurationReadonlyexpirationend field if non-nil, zero value otherwise.

### GetDtSystemconfigurationReadonlyexpirationendOk

`func (o *SystemconfigurationResponse) GetDtSystemconfigurationReadonlyexpirationendOk() (*string, bool)`

GetDtSystemconfigurationReadonlyexpirationendOk returns a tuple with the DtSystemconfigurationReadonlyexpirationend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtSystemconfigurationReadonlyexpirationend

`func (o *SystemconfigurationResponse) SetDtSystemconfigurationReadonlyexpirationend(v string)`

SetDtSystemconfigurationReadonlyexpirationend sets DtSystemconfigurationReadonlyexpirationend field to given value.

### HasDtSystemconfigurationReadonlyexpirationend

`func (o *SystemconfigurationResponse) HasDtSystemconfigurationReadonlyexpirationend() bool`

HasDtSystemconfigurationReadonlyexpirationend returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


