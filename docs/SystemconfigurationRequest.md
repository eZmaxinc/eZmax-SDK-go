# SystemconfigurationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiSystemconfigurationID** | Pointer to **int32** | The unique ID of the Systemconfiguration | [optional] 
**FkiBrandingID** | Pointer to **int32** | The unique ID of the Branding | [optional] 
**ESystemconfigurationNewexternaluseraction** | [**FieldESystemconfigurationNewexternaluseraction**](FieldESystemconfigurationNewexternaluseraction.md) |  | 
**ESystemconfigurationLanguage1** | [**FieldESystemconfigurationLanguage1**](FieldESystemconfigurationLanguage1.md) |  | 
**ESystemconfigurationLanguage2** | [**FieldESystemconfigurationLanguage2**](FieldESystemconfigurationLanguage2.md) |  | 
**ESystemconfigurationEzsign** | Pointer to [**FieldESystemconfigurationEzsign**](FieldESystemconfigurationEzsign.md) |  | [optional] 
**ESystemconfigurationEzsignofficeplan** | Pointer to [**FieldESystemconfigurationEzsignofficeplan**](FieldESystemconfigurationEzsignofficeplan.md) |  | [optional] 
**BSystemconfigurationEzsignpaidbyoffice** | Pointer to **bool** | Whether if Ezsign is paid by the company or not | [optional] 
**BSystemconfigurationEzsignpersonnal** | **bool** | Whether if we allow the creation of personal files in eZsign | 
**BSystemconfigurationSspr** | **bool** | Whether if we allow SSPR | 
**DtSystemconfigurationReadonlyexpirationstart** | Pointer to **string** | The start date where the system will be in read only | [optional] 
**DtSystemconfigurationReadonlyexpirationend** | Pointer to **string** | The end date where the system will be in read only | [optional] 

## Methods

### NewSystemconfigurationRequest

`func NewSystemconfigurationRequest(eSystemconfigurationNewexternaluseraction FieldESystemconfigurationNewexternaluseraction, eSystemconfigurationLanguage1 FieldESystemconfigurationLanguage1, eSystemconfigurationLanguage2 FieldESystemconfigurationLanguage2, bSystemconfigurationEzsignpersonnal bool, bSystemconfigurationSspr bool, ) *SystemconfigurationRequest`

NewSystemconfigurationRequest instantiates a new SystemconfigurationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemconfigurationRequestWithDefaults

`func NewSystemconfigurationRequestWithDefaults() *SystemconfigurationRequest`

NewSystemconfigurationRequestWithDefaults instantiates a new SystemconfigurationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiSystemconfigurationID

`func (o *SystemconfigurationRequest) GetPkiSystemconfigurationID() int32`

GetPkiSystemconfigurationID returns the PkiSystemconfigurationID field if non-nil, zero value otherwise.

### GetPkiSystemconfigurationIDOk

`func (o *SystemconfigurationRequest) GetPkiSystemconfigurationIDOk() (*int32, bool)`

GetPkiSystemconfigurationIDOk returns a tuple with the PkiSystemconfigurationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiSystemconfigurationID

`func (o *SystemconfigurationRequest) SetPkiSystemconfigurationID(v int32)`

SetPkiSystemconfigurationID sets PkiSystemconfigurationID field to given value.

### HasPkiSystemconfigurationID

`func (o *SystemconfigurationRequest) HasPkiSystemconfigurationID() bool`

HasPkiSystemconfigurationID returns a boolean if a field has been set.

### GetFkiBrandingID

`func (o *SystemconfigurationRequest) GetFkiBrandingID() int32`

GetFkiBrandingID returns the FkiBrandingID field if non-nil, zero value otherwise.

### GetFkiBrandingIDOk

`func (o *SystemconfigurationRequest) GetFkiBrandingIDOk() (*int32, bool)`

GetFkiBrandingIDOk returns a tuple with the FkiBrandingID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiBrandingID

`func (o *SystemconfigurationRequest) SetFkiBrandingID(v int32)`

SetFkiBrandingID sets FkiBrandingID field to given value.

### HasFkiBrandingID

`func (o *SystemconfigurationRequest) HasFkiBrandingID() bool`

HasFkiBrandingID returns a boolean if a field has been set.

### GetESystemconfigurationNewexternaluseraction

`func (o *SystemconfigurationRequest) GetESystemconfigurationNewexternaluseraction() FieldESystemconfigurationNewexternaluseraction`

GetESystemconfigurationNewexternaluseraction returns the ESystemconfigurationNewexternaluseraction field if non-nil, zero value otherwise.

### GetESystemconfigurationNewexternaluseractionOk

`func (o *SystemconfigurationRequest) GetESystemconfigurationNewexternaluseractionOk() (*FieldESystemconfigurationNewexternaluseraction, bool)`

GetESystemconfigurationNewexternaluseractionOk returns a tuple with the ESystemconfigurationNewexternaluseraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetESystemconfigurationNewexternaluseraction

`func (o *SystemconfigurationRequest) SetESystemconfigurationNewexternaluseraction(v FieldESystemconfigurationNewexternaluseraction)`

SetESystemconfigurationNewexternaluseraction sets ESystemconfigurationNewexternaluseraction field to given value.


### GetESystemconfigurationLanguage1

`func (o *SystemconfigurationRequest) GetESystemconfigurationLanguage1() FieldESystemconfigurationLanguage1`

GetESystemconfigurationLanguage1 returns the ESystemconfigurationLanguage1 field if non-nil, zero value otherwise.

### GetESystemconfigurationLanguage1Ok

`func (o *SystemconfigurationRequest) GetESystemconfigurationLanguage1Ok() (*FieldESystemconfigurationLanguage1, bool)`

GetESystemconfigurationLanguage1Ok returns a tuple with the ESystemconfigurationLanguage1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetESystemconfigurationLanguage1

`func (o *SystemconfigurationRequest) SetESystemconfigurationLanguage1(v FieldESystemconfigurationLanguage1)`

SetESystemconfigurationLanguage1 sets ESystemconfigurationLanguage1 field to given value.


### GetESystemconfigurationLanguage2

`func (o *SystemconfigurationRequest) GetESystemconfigurationLanguage2() FieldESystemconfigurationLanguage2`

GetESystemconfigurationLanguage2 returns the ESystemconfigurationLanguage2 field if non-nil, zero value otherwise.

### GetESystemconfigurationLanguage2Ok

`func (o *SystemconfigurationRequest) GetESystemconfigurationLanguage2Ok() (*FieldESystemconfigurationLanguage2, bool)`

GetESystemconfigurationLanguage2Ok returns a tuple with the ESystemconfigurationLanguage2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetESystemconfigurationLanguage2

`func (o *SystemconfigurationRequest) SetESystemconfigurationLanguage2(v FieldESystemconfigurationLanguage2)`

SetESystemconfigurationLanguage2 sets ESystemconfigurationLanguage2 field to given value.


### GetESystemconfigurationEzsign

`func (o *SystemconfigurationRequest) GetESystemconfigurationEzsign() FieldESystemconfigurationEzsign`

GetESystemconfigurationEzsign returns the ESystemconfigurationEzsign field if non-nil, zero value otherwise.

### GetESystemconfigurationEzsignOk

`func (o *SystemconfigurationRequest) GetESystemconfigurationEzsignOk() (*FieldESystemconfigurationEzsign, bool)`

GetESystemconfigurationEzsignOk returns a tuple with the ESystemconfigurationEzsign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetESystemconfigurationEzsign

`func (o *SystemconfigurationRequest) SetESystemconfigurationEzsign(v FieldESystemconfigurationEzsign)`

SetESystemconfigurationEzsign sets ESystemconfigurationEzsign field to given value.

### HasESystemconfigurationEzsign

`func (o *SystemconfigurationRequest) HasESystemconfigurationEzsign() bool`

HasESystemconfigurationEzsign returns a boolean if a field has been set.

### GetESystemconfigurationEzsignofficeplan

`func (o *SystemconfigurationRequest) GetESystemconfigurationEzsignofficeplan() FieldESystemconfigurationEzsignofficeplan`

GetESystemconfigurationEzsignofficeplan returns the ESystemconfigurationEzsignofficeplan field if non-nil, zero value otherwise.

### GetESystemconfigurationEzsignofficeplanOk

`func (o *SystemconfigurationRequest) GetESystemconfigurationEzsignofficeplanOk() (*FieldESystemconfigurationEzsignofficeplan, bool)`

GetESystemconfigurationEzsignofficeplanOk returns a tuple with the ESystemconfigurationEzsignofficeplan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetESystemconfigurationEzsignofficeplan

`func (o *SystemconfigurationRequest) SetESystemconfigurationEzsignofficeplan(v FieldESystemconfigurationEzsignofficeplan)`

SetESystemconfigurationEzsignofficeplan sets ESystemconfigurationEzsignofficeplan field to given value.

### HasESystemconfigurationEzsignofficeplan

`func (o *SystemconfigurationRequest) HasESystemconfigurationEzsignofficeplan() bool`

HasESystemconfigurationEzsignofficeplan returns a boolean if a field has been set.

### GetBSystemconfigurationEzsignpaidbyoffice

`func (o *SystemconfigurationRequest) GetBSystemconfigurationEzsignpaidbyoffice() bool`

GetBSystemconfigurationEzsignpaidbyoffice returns the BSystemconfigurationEzsignpaidbyoffice field if non-nil, zero value otherwise.

### GetBSystemconfigurationEzsignpaidbyofficeOk

`func (o *SystemconfigurationRequest) GetBSystemconfigurationEzsignpaidbyofficeOk() (*bool, bool)`

GetBSystemconfigurationEzsignpaidbyofficeOk returns a tuple with the BSystemconfigurationEzsignpaidbyoffice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBSystemconfigurationEzsignpaidbyoffice

`func (o *SystemconfigurationRequest) SetBSystemconfigurationEzsignpaidbyoffice(v bool)`

SetBSystemconfigurationEzsignpaidbyoffice sets BSystemconfigurationEzsignpaidbyoffice field to given value.

### HasBSystemconfigurationEzsignpaidbyoffice

`func (o *SystemconfigurationRequest) HasBSystemconfigurationEzsignpaidbyoffice() bool`

HasBSystemconfigurationEzsignpaidbyoffice returns a boolean if a field has been set.

### GetBSystemconfigurationEzsignpersonnal

`func (o *SystemconfigurationRequest) GetBSystemconfigurationEzsignpersonnal() bool`

GetBSystemconfigurationEzsignpersonnal returns the BSystemconfigurationEzsignpersonnal field if non-nil, zero value otherwise.

### GetBSystemconfigurationEzsignpersonnalOk

`func (o *SystemconfigurationRequest) GetBSystemconfigurationEzsignpersonnalOk() (*bool, bool)`

GetBSystemconfigurationEzsignpersonnalOk returns a tuple with the BSystemconfigurationEzsignpersonnal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBSystemconfigurationEzsignpersonnal

`func (o *SystemconfigurationRequest) SetBSystemconfigurationEzsignpersonnal(v bool)`

SetBSystemconfigurationEzsignpersonnal sets BSystemconfigurationEzsignpersonnal field to given value.


### GetBSystemconfigurationSspr

`func (o *SystemconfigurationRequest) GetBSystemconfigurationSspr() bool`

GetBSystemconfigurationSspr returns the BSystemconfigurationSspr field if non-nil, zero value otherwise.

### GetBSystemconfigurationSsprOk

`func (o *SystemconfigurationRequest) GetBSystemconfigurationSsprOk() (*bool, bool)`

GetBSystemconfigurationSsprOk returns a tuple with the BSystemconfigurationSspr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBSystemconfigurationSspr

`func (o *SystemconfigurationRequest) SetBSystemconfigurationSspr(v bool)`

SetBSystemconfigurationSspr sets BSystemconfigurationSspr field to given value.


### GetDtSystemconfigurationReadonlyexpirationstart

`func (o *SystemconfigurationRequest) GetDtSystemconfigurationReadonlyexpirationstart() string`

GetDtSystemconfigurationReadonlyexpirationstart returns the DtSystemconfigurationReadonlyexpirationstart field if non-nil, zero value otherwise.

### GetDtSystemconfigurationReadonlyexpirationstartOk

`func (o *SystemconfigurationRequest) GetDtSystemconfigurationReadonlyexpirationstartOk() (*string, bool)`

GetDtSystemconfigurationReadonlyexpirationstartOk returns a tuple with the DtSystemconfigurationReadonlyexpirationstart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtSystemconfigurationReadonlyexpirationstart

`func (o *SystemconfigurationRequest) SetDtSystemconfigurationReadonlyexpirationstart(v string)`

SetDtSystemconfigurationReadonlyexpirationstart sets DtSystemconfigurationReadonlyexpirationstart field to given value.

### HasDtSystemconfigurationReadonlyexpirationstart

`func (o *SystemconfigurationRequest) HasDtSystemconfigurationReadonlyexpirationstart() bool`

HasDtSystemconfigurationReadonlyexpirationstart returns a boolean if a field has been set.

### GetDtSystemconfigurationReadonlyexpirationend

`func (o *SystemconfigurationRequest) GetDtSystemconfigurationReadonlyexpirationend() string`

GetDtSystemconfigurationReadonlyexpirationend returns the DtSystemconfigurationReadonlyexpirationend field if non-nil, zero value otherwise.

### GetDtSystemconfigurationReadonlyexpirationendOk

`func (o *SystemconfigurationRequest) GetDtSystemconfigurationReadonlyexpirationendOk() (*string, bool)`

GetDtSystemconfigurationReadonlyexpirationendOk returns a tuple with the DtSystemconfigurationReadonlyexpirationend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtSystemconfigurationReadonlyexpirationend

`func (o *SystemconfigurationRequest) SetDtSystemconfigurationReadonlyexpirationend(v string)`

SetDtSystemconfigurationReadonlyexpirationend sets DtSystemconfigurationReadonlyexpirationend field to given value.

### HasDtSystemconfigurationReadonlyexpirationend

`func (o *SystemconfigurationRequest) HasDtSystemconfigurationReadonlyexpirationend() bool`

HasDtSystemconfigurationReadonlyexpirationend returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


