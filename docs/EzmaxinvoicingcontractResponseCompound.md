# EzmaxinvoicingcontractResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzmaxinvoicingcontractID** | **int32** | The unique ID of the Ezmaxinvoicingcontract | 
**EEzmaxinvoicingcontractPaymenttype** | [**FieldEEzmaxinvoicingcontractPaymenttype**](FieldEEzmaxinvoicingcontractPaymenttype.md) |  | 
**IEzmaxinvoicingcontractLength** | **int32** | The length in years of the Ezmaxinvoicingcontract | 
**DtEzmaxinvoicingcontractStart** | **string** | The start date of the Ezmaxinvoicingcontract | 
**DtEzmaxinvoicingcontractEnd** | **string** | The end date of the Ezmaxinvoicingcontract | 
**DEzmaxinvoicingcontractLicense** | **string** | The price of the license | 
**DEzmaxinvoicingcontract121qa** | **string** | The price for 121QA | 
**BEzmaxinvoicingcontractEzsignallagents** | **bool** | Whether eZsign is for all agents | 
**ObjAudit** | [**CommonAudit**](CommonAudit.md) |  | 

## Methods

### NewEzmaxinvoicingcontractResponseCompound

`func NewEzmaxinvoicingcontractResponseCompound(pkiEzmaxinvoicingcontractID int32, eEzmaxinvoicingcontractPaymenttype FieldEEzmaxinvoicingcontractPaymenttype, iEzmaxinvoicingcontractLength int32, dtEzmaxinvoicingcontractStart string, dtEzmaxinvoicingcontractEnd string, dEzmaxinvoicingcontractLicense string, dEzmaxinvoicingcontract121qa string, bEzmaxinvoicingcontractEzsignallagents bool, objAudit CommonAudit, ) *EzmaxinvoicingcontractResponseCompound`

NewEzmaxinvoicingcontractResponseCompound instantiates a new EzmaxinvoicingcontractResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzmaxinvoicingcontractResponseCompoundWithDefaults

`func NewEzmaxinvoicingcontractResponseCompoundWithDefaults() *EzmaxinvoicingcontractResponseCompound`

NewEzmaxinvoicingcontractResponseCompoundWithDefaults instantiates a new EzmaxinvoicingcontractResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzmaxinvoicingcontractID

`func (o *EzmaxinvoicingcontractResponseCompound) GetPkiEzmaxinvoicingcontractID() int32`

GetPkiEzmaxinvoicingcontractID returns the PkiEzmaxinvoicingcontractID field if non-nil, zero value otherwise.

### GetPkiEzmaxinvoicingcontractIDOk

`func (o *EzmaxinvoicingcontractResponseCompound) GetPkiEzmaxinvoicingcontractIDOk() (*int32, bool)`

GetPkiEzmaxinvoicingcontractIDOk returns a tuple with the PkiEzmaxinvoicingcontractID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzmaxinvoicingcontractID

`func (o *EzmaxinvoicingcontractResponseCompound) SetPkiEzmaxinvoicingcontractID(v int32)`

SetPkiEzmaxinvoicingcontractID sets PkiEzmaxinvoicingcontractID field to given value.


### GetEEzmaxinvoicingcontractPaymenttype

`func (o *EzmaxinvoicingcontractResponseCompound) GetEEzmaxinvoicingcontractPaymenttype() FieldEEzmaxinvoicingcontractPaymenttype`

GetEEzmaxinvoicingcontractPaymenttype returns the EEzmaxinvoicingcontractPaymenttype field if non-nil, zero value otherwise.

### GetEEzmaxinvoicingcontractPaymenttypeOk

`func (o *EzmaxinvoicingcontractResponseCompound) GetEEzmaxinvoicingcontractPaymenttypeOk() (*FieldEEzmaxinvoicingcontractPaymenttype, bool)`

GetEEzmaxinvoicingcontractPaymenttypeOk returns a tuple with the EEzmaxinvoicingcontractPaymenttype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzmaxinvoicingcontractPaymenttype

`func (o *EzmaxinvoicingcontractResponseCompound) SetEEzmaxinvoicingcontractPaymenttype(v FieldEEzmaxinvoicingcontractPaymenttype)`

SetEEzmaxinvoicingcontractPaymenttype sets EEzmaxinvoicingcontractPaymenttype field to given value.


### GetIEzmaxinvoicingcontractLength

`func (o *EzmaxinvoicingcontractResponseCompound) GetIEzmaxinvoicingcontractLength() int32`

GetIEzmaxinvoicingcontractLength returns the IEzmaxinvoicingcontractLength field if non-nil, zero value otherwise.

### GetIEzmaxinvoicingcontractLengthOk

`func (o *EzmaxinvoicingcontractResponseCompound) GetIEzmaxinvoicingcontractLengthOk() (*int32, bool)`

GetIEzmaxinvoicingcontractLengthOk returns a tuple with the IEzmaxinvoicingcontractLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzmaxinvoicingcontractLength

`func (o *EzmaxinvoicingcontractResponseCompound) SetIEzmaxinvoicingcontractLength(v int32)`

SetIEzmaxinvoicingcontractLength sets IEzmaxinvoicingcontractLength field to given value.


### GetDtEzmaxinvoicingcontractStart

`func (o *EzmaxinvoicingcontractResponseCompound) GetDtEzmaxinvoicingcontractStart() string`

GetDtEzmaxinvoicingcontractStart returns the DtEzmaxinvoicingcontractStart field if non-nil, zero value otherwise.

### GetDtEzmaxinvoicingcontractStartOk

`func (o *EzmaxinvoicingcontractResponseCompound) GetDtEzmaxinvoicingcontractStartOk() (*string, bool)`

GetDtEzmaxinvoicingcontractStartOk returns a tuple with the DtEzmaxinvoicingcontractStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtEzmaxinvoicingcontractStart

`func (o *EzmaxinvoicingcontractResponseCompound) SetDtEzmaxinvoicingcontractStart(v string)`

SetDtEzmaxinvoicingcontractStart sets DtEzmaxinvoicingcontractStart field to given value.


### GetDtEzmaxinvoicingcontractEnd

`func (o *EzmaxinvoicingcontractResponseCompound) GetDtEzmaxinvoicingcontractEnd() string`

GetDtEzmaxinvoicingcontractEnd returns the DtEzmaxinvoicingcontractEnd field if non-nil, zero value otherwise.

### GetDtEzmaxinvoicingcontractEndOk

`func (o *EzmaxinvoicingcontractResponseCompound) GetDtEzmaxinvoicingcontractEndOk() (*string, bool)`

GetDtEzmaxinvoicingcontractEndOk returns a tuple with the DtEzmaxinvoicingcontractEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtEzmaxinvoicingcontractEnd

`func (o *EzmaxinvoicingcontractResponseCompound) SetDtEzmaxinvoicingcontractEnd(v string)`

SetDtEzmaxinvoicingcontractEnd sets DtEzmaxinvoicingcontractEnd field to given value.


### GetDEzmaxinvoicingcontractLicense

`func (o *EzmaxinvoicingcontractResponseCompound) GetDEzmaxinvoicingcontractLicense() string`

GetDEzmaxinvoicingcontractLicense returns the DEzmaxinvoicingcontractLicense field if non-nil, zero value otherwise.

### GetDEzmaxinvoicingcontractLicenseOk

`func (o *EzmaxinvoicingcontractResponseCompound) GetDEzmaxinvoicingcontractLicenseOk() (*string, bool)`

GetDEzmaxinvoicingcontractLicenseOk returns a tuple with the DEzmaxinvoicingcontractLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEzmaxinvoicingcontractLicense

`func (o *EzmaxinvoicingcontractResponseCompound) SetDEzmaxinvoicingcontractLicense(v string)`

SetDEzmaxinvoicingcontractLicense sets DEzmaxinvoicingcontractLicense field to given value.


### GetDEzmaxinvoicingcontract121qa

`func (o *EzmaxinvoicingcontractResponseCompound) GetDEzmaxinvoicingcontract121qa() string`

GetDEzmaxinvoicingcontract121qa returns the DEzmaxinvoicingcontract121qa field if non-nil, zero value otherwise.

### GetDEzmaxinvoicingcontract121qaOk

`func (o *EzmaxinvoicingcontractResponseCompound) GetDEzmaxinvoicingcontract121qaOk() (*string, bool)`

GetDEzmaxinvoicingcontract121qaOk returns a tuple with the DEzmaxinvoicingcontract121qa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEzmaxinvoicingcontract121qa

`func (o *EzmaxinvoicingcontractResponseCompound) SetDEzmaxinvoicingcontract121qa(v string)`

SetDEzmaxinvoicingcontract121qa sets DEzmaxinvoicingcontract121qa field to given value.


### GetBEzmaxinvoicingcontractEzsignallagents

`func (o *EzmaxinvoicingcontractResponseCompound) GetBEzmaxinvoicingcontractEzsignallagents() bool`

GetBEzmaxinvoicingcontractEzsignallagents returns the BEzmaxinvoicingcontractEzsignallagents field if non-nil, zero value otherwise.

### GetBEzmaxinvoicingcontractEzsignallagentsOk

`func (o *EzmaxinvoicingcontractResponseCompound) GetBEzmaxinvoicingcontractEzsignallagentsOk() (*bool, bool)`

GetBEzmaxinvoicingcontractEzsignallagentsOk returns a tuple with the BEzmaxinvoicingcontractEzsignallagents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzmaxinvoicingcontractEzsignallagents

`func (o *EzmaxinvoicingcontractResponseCompound) SetBEzmaxinvoicingcontractEzsignallagents(v bool)`

SetBEzmaxinvoicingcontractEzsignallagents sets BEzmaxinvoicingcontractEzsignallagents field to given value.


### GetObjAudit

`func (o *EzmaxinvoicingcontractResponseCompound) GetObjAudit() CommonAudit`

GetObjAudit returns the ObjAudit field if non-nil, zero value otherwise.

### GetObjAuditOk

`func (o *EzmaxinvoicingcontractResponseCompound) GetObjAuditOk() (*CommonAudit, bool)`

GetObjAuditOk returns a tuple with the ObjAudit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjAudit

`func (o *EzmaxinvoicingcontractResponseCompound) SetObjAudit(v CommonAudit)`

SetObjAudit sets ObjAudit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


