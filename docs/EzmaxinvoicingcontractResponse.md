# EzmaxinvoicingcontractResponse

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

### NewEzmaxinvoicingcontractResponse

`func NewEzmaxinvoicingcontractResponse(pkiEzmaxinvoicingcontractID int32, eEzmaxinvoicingcontractPaymenttype FieldEEzmaxinvoicingcontractPaymenttype, iEzmaxinvoicingcontractLength int32, dtEzmaxinvoicingcontractStart string, dtEzmaxinvoicingcontractEnd string, dEzmaxinvoicingcontractLicense string, dEzmaxinvoicingcontract121qa string, bEzmaxinvoicingcontractEzsignallagents bool, objAudit CommonAudit, ) *EzmaxinvoicingcontractResponse`

NewEzmaxinvoicingcontractResponse instantiates a new EzmaxinvoicingcontractResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzmaxinvoicingcontractResponseWithDefaults

`func NewEzmaxinvoicingcontractResponseWithDefaults() *EzmaxinvoicingcontractResponse`

NewEzmaxinvoicingcontractResponseWithDefaults instantiates a new EzmaxinvoicingcontractResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzmaxinvoicingcontractID

`func (o *EzmaxinvoicingcontractResponse) GetPkiEzmaxinvoicingcontractID() int32`

GetPkiEzmaxinvoicingcontractID returns the PkiEzmaxinvoicingcontractID field if non-nil, zero value otherwise.

### GetPkiEzmaxinvoicingcontractIDOk

`func (o *EzmaxinvoicingcontractResponse) GetPkiEzmaxinvoicingcontractIDOk() (*int32, bool)`

GetPkiEzmaxinvoicingcontractIDOk returns a tuple with the PkiEzmaxinvoicingcontractID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzmaxinvoicingcontractID

`func (o *EzmaxinvoicingcontractResponse) SetPkiEzmaxinvoicingcontractID(v int32)`

SetPkiEzmaxinvoicingcontractID sets PkiEzmaxinvoicingcontractID field to given value.


### GetEEzmaxinvoicingcontractPaymenttype

`func (o *EzmaxinvoicingcontractResponse) GetEEzmaxinvoicingcontractPaymenttype() FieldEEzmaxinvoicingcontractPaymenttype`

GetEEzmaxinvoicingcontractPaymenttype returns the EEzmaxinvoicingcontractPaymenttype field if non-nil, zero value otherwise.

### GetEEzmaxinvoicingcontractPaymenttypeOk

`func (o *EzmaxinvoicingcontractResponse) GetEEzmaxinvoicingcontractPaymenttypeOk() (*FieldEEzmaxinvoicingcontractPaymenttype, bool)`

GetEEzmaxinvoicingcontractPaymenttypeOk returns a tuple with the EEzmaxinvoicingcontractPaymenttype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzmaxinvoicingcontractPaymenttype

`func (o *EzmaxinvoicingcontractResponse) SetEEzmaxinvoicingcontractPaymenttype(v FieldEEzmaxinvoicingcontractPaymenttype)`

SetEEzmaxinvoicingcontractPaymenttype sets EEzmaxinvoicingcontractPaymenttype field to given value.


### GetIEzmaxinvoicingcontractLength

`func (o *EzmaxinvoicingcontractResponse) GetIEzmaxinvoicingcontractLength() int32`

GetIEzmaxinvoicingcontractLength returns the IEzmaxinvoicingcontractLength field if non-nil, zero value otherwise.

### GetIEzmaxinvoicingcontractLengthOk

`func (o *EzmaxinvoicingcontractResponse) GetIEzmaxinvoicingcontractLengthOk() (*int32, bool)`

GetIEzmaxinvoicingcontractLengthOk returns a tuple with the IEzmaxinvoicingcontractLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzmaxinvoicingcontractLength

`func (o *EzmaxinvoicingcontractResponse) SetIEzmaxinvoicingcontractLength(v int32)`

SetIEzmaxinvoicingcontractLength sets IEzmaxinvoicingcontractLength field to given value.


### GetDtEzmaxinvoicingcontractStart

`func (o *EzmaxinvoicingcontractResponse) GetDtEzmaxinvoicingcontractStart() string`

GetDtEzmaxinvoicingcontractStart returns the DtEzmaxinvoicingcontractStart field if non-nil, zero value otherwise.

### GetDtEzmaxinvoicingcontractStartOk

`func (o *EzmaxinvoicingcontractResponse) GetDtEzmaxinvoicingcontractStartOk() (*string, bool)`

GetDtEzmaxinvoicingcontractStartOk returns a tuple with the DtEzmaxinvoicingcontractStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtEzmaxinvoicingcontractStart

`func (o *EzmaxinvoicingcontractResponse) SetDtEzmaxinvoicingcontractStart(v string)`

SetDtEzmaxinvoicingcontractStart sets DtEzmaxinvoicingcontractStart field to given value.


### GetDtEzmaxinvoicingcontractEnd

`func (o *EzmaxinvoicingcontractResponse) GetDtEzmaxinvoicingcontractEnd() string`

GetDtEzmaxinvoicingcontractEnd returns the DtEzmaxinvoicingcontractEnd field if non-nil, zero value otherwise.

### GetDtEzmaxinvoicingcontractEndOk

`func (o *EzmaxinvoicingcontractResponse) GetDtEzmaxinvoicingcontractEndOk() (*string, bool)`

GetDtEzmaxinvoicingcontractEndOk returns a tuple with the DtEzmaxinvoicingcontractEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtEzmaxinvoicingcontractEnd

`func (o *EzmaxinvoicingcontractResponse) SetDtEzmaxinvoicingcontractEnd(v string)`

SetDtEzmaxinvoicingcontractEnd sets DtEzmaxinvoicingcontractEnd field to given value.


### GetDEzmaxinvoicingcontractLicense

`func (o *EzmaxinvoicingcontractResponse) GetDEzmaxinvoicingcontractLicense() string`

GetDEzmaxinvoicingcontractLicense returns the DEzmaxinvoicingcontractLicense field if non-nil, zero value otherwise.

### GetDEzmaxinvoicingcontractLicenseOk

`func (o *EzmaxinvoicingcontractResponse) GetDEzmaxinvoicingcontractLicenseOk() (*string, bool)`

GetDEzmaxinvoicingcontractLicenseOk returns a tuple with the DEzmaxinvoicingcontractLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEzmaxinvoicingcontractLicense

`func (o *EzmaxinvoicingcontractResponse) SetDEzmaxinvoicingcontractLicense(v string)`

SetDEzmaxinvoicingcontractLicense sets DEzmaxinvoicingcontractLicense field to given value.


### GetDEzmaxinvoicingcontract121qa

`func (o *EzmaxinvoicingcontractResponse) GetDEzmaxinvoicingcontract121qa() string`

GetDEzmaxinvoicingcontract121qa returns the DEzmaxinvoicingcontract121qa field if non-nil, zero value otherwise.

### GetDEzmaxinvoicingcontract121qaOk

`func (o *EzmaxinvoicingcontractResponse) GetDEzmaxinvoicingcontract121qaOk() (*string, bool)`

GetDEzmaxinvoicingcontract121qaOk returns a tuple with the DEzmaxinvoicingcontract121qa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEzmaxinvoicingcontract121qa

`func (o *EzmaxinvoicingcontractResponse) SetDEzmaxinvoicingcontract121qa(v string)`

SetDEzmaxinvoicingcontract121qa sets DEzmaxinvoicingcontract121qa field to given value.


### GetBEzmaxinvoicingcontractEzsignallagents

`func (o *EzmaxinvoicingcontractResponse) GetBEzmaxinvoicingcontractEzsignallagents() bool`

GetBEzmaxinvoicingcontractEzsignallagents returns the BEzmaxinvoicingcontractEzsignallagents field if non-nil, zero value otherwise.

### GetBEzmaxinvoicingcontractEzsignallagentsOk

`func (o *EzmaxinvoicingcontractResponse) GetBEzmaxinvoicingcontractEzsignallagentsOk() (*bool, bool)`

GetBEzmaxinvoicingcontractEzsignallagentsOk returns a tuple with the BEzmaxinvoicingcontractEzsignallagents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzmaxinvoicingcontractEzsignallagents

`func (o *EzmaxinvoicingcontractResponse) SetBEzmaxinvoicingcontractEzsignallagents(v bool)`

SetBEzmaxinvoicingcontractEzsignallagents sets BEzmaxinvoicingcontractEzsignallagents field to given value.


### GetObjAudit

`func (o *EzmaxinvoicingcontractResponse) GetObjAudit() CommonAudit`

GetObjAudit returns the ObjAudit field if non-nil, zero value otherwise.

### GetObjAuditOk

`func (o *EzmaxinvoicingcontractResponse) GetObjAuditOk() (*CommonAudit, bool)`

GetObjAuditOk returns a tuple with the ObjAudit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjAudit

`func (o *EzmaxinvoicingcontractResponse) SetObjAudit(v CommonAudit)`

SetObjAudit sets ObjAudit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


