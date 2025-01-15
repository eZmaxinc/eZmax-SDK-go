# DomainResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiDomainID** | **int32** | The unique ID of the Domain | 
**SDomainName** | **string** | The name of the Domain | 
**BDomainValiddkim** | **bool** | Whether the DKIM is valid or not | 
**BDomainValidmailfrom** | **bool** | Whether the mail from is valid or not | 
**BDomainValidcustomer** | **bool** | Whether the customer has access to it or not | 
**ObjAudit** | [**CommonAudit**](CommonAudit.md) |  | 
**AObjDnsrecord** | **[]map[string]interface{}** |  | 

## Methods

### NewDomainResponseCompound

`func NewDomainResponseCompound(pkiDomainID int32, sDomainName string, bDomainValiddkim bool, bDomainValidmailfrom bool, bDomainValidcustomer bool, objAudit CommonAudit, aObjDnsrecord []CustomDnsrecordResponse, ) *DomainResponseCompound`

NewDomainResponseCompound instantiates a new DomainResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainResponseCompoundWithDefaults

`func NewDomainResponseCompoundWithDefaults() *DomainResponseCompound`

NewDomainResponseCompoundWithDefaults instantiates a new DomainResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiDomainID

`func (o *DomainResponseCompound) GetPkiDomainID() int32`

GetPkiDomainID returns the PkiDomainID field if non-nil, zero value otherwise.

### GetPkiDomainIDOk

`func (o *DomainResponseCompound) GetPkiDomainIDOk() (*int32, bool)`

GetPkiDomainIDOk returns a tuple with the PkiDomainID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiDomainID

`func (o *DomainResponseCompound) SetPkiDomainID(v int32)`

SetPkiDomainID sets PkiDomainID field to given value.


### GetSDomainName

`func (o *DomainResponseCompound) GetSDomainName() string`

GetSDomainName returns the SDomainName field if non-nil, zero value otherwise.

### GetSDomainNameOk

`func (o *DomainResponseCompound) GetSDomainNameOk() (*string, bool)`

GetSDomainNameOk returns a tuple with the SDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSDomainName

`func (o *DomainResponseCompound) SetSDomainName(v string)`

SetSDomainName sets SDomainName field to given value.


### GetBDomainValiddkim

`func (o *DomainResponseCompound) GetBDomainValiddkim() bool`

GetBDomainValiddkim returns the BDomainValiddkim field if non-nil, zero value otherwise.

### GetBDomainValiddkimOk

`func (o *DomainResponseCompound) GetBDomainValiddkimOk() (*bool, bool)`

GetBDomainValiddkimOk returns a tuple with the BDomainValiddkim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDomainValiddkim

`func (o *DomainResponseCompound) SetBDomainValiddkim(v bool)`

SetBDomainValiddkim sets BDomainValiddkim field to given value.


### GetBDomainValidmailfrom

`func (o *DomainResponseCompound) GetBDomainValidmailfrom() bool`

GetBDomainValidmailfrom returns the BDomainValidmailfrom field if non-nil, zero value otherwise.

### GetBDomainValidmailfromOk

`func (o *DomainResponseCompound) GetBDomainValidmailfromOk() (*bool, bool)`

GetBDomainValidmailfromOk returns a tuple with the BDomainValidmailfrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDomainValidmailfrom

`func (o *DomainResponseCompound) SetBDomainValidmailfrom(v bool)`

SetBDomainValidmailfrom sets BDomainValidmailfrom field to given value.


### GetBDomainValidcustomer

`func (o *DomainResponseCompound) GetBDomainValidcustomer() bool`

GetBDomainValidcustomer returns the BDomainValidcustomer field if non-nil, zero value otherwise.

### GetBDomainValidcustomerOk

`func (o *DomainResponseCompound) GetBDomainValidcustomerOk() (*bool, bool)`

GetBDomainValidcustomerOk returns a tuple with the BDomainValidcustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDomainValidcustomer

`func (o *DomainResponseCompound) SetBDomainValidcustomer(v bool)`

SetBDomainValidcustomer sets BDomainValidcustomer field to given value.


### GetObjAudit

`func (o *DomainResponseCompound) GetObjAudit() CommonAudit`

GetObjAudit returns the ObjAudit field if non-nil, zero value otherwise.

### GetObjAuditOk

`func (o *DomainResponseCompound) GetObjAuditOk() (*CommonAudit, bool)`

GetObjAuditOk returns a tuple with the ObjAudit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjAudit

`func (o *DomainResponseCompound) SetObjAudit(v CommonAudit)`

SetObjAudit sets ObjAudit field to given value.


### GetAObjDnsrecord

`func (o *DomainResponseCompound) GetAObjDnsrecord() []CustomDnsrecordResponse`

GetAObjDnsrecord returns the AObjDnsrecord field if non-nil, zero value otherwise.

### GetAObjDnsrecordOk

`func (o *DomainResponseCompound) GetAObjDnsrecordOk() (*[]CustomDnsrecordResponse, bool)`

GetAObjDnsrecordOk returns a tuple with the AObjDnsrecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjDnsrecord

`func (o *DomainResponseCompound) SetAObjDnsrecord(v []CustomDnsrecordResponse)`

SetAObjDnsrecord sets AObjDnsrecord field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


