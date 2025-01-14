# DomainResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiDomainID** | **int32** | The unique ID of the Domain | 
**SDomainName** | **string** | The name of the Domain | 
**BDomainValiddkim** | **bool** | Whether the DKIM is valid or not | 
**BDomainValidmailfrom** | **bool** | Whether the mail from is valid or not | 
**BDomainValidcustomer** | **bool** | Whether the customer has access to it or not | 
**ObjAudit** | [**CommonAudit**](CommonAudit.md) |  | 

## Methods

### NewDomainResponse

`func NewDomainResponse(pkiDomainID int32, sDomainName string, bDomainValiddkim bool, bDomainValidmailfrom bool, bDomainValidcustomer bool, objAudit CommonAudit, ) *DomainResponse`

NewDomainResponse instantiates a new DomainResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainResponseWithDefaults

`func NewDomainResponseWithDefaults() *DomainResponse`

NewDomainResponseWithDefaults instantiates a new DomainResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiDomainID

`func (o *DomainResponse) GetPkiDomainID() int32`

GetPkiDomainID returns the PkiDomainID field if non-nil, zero value otherwise.

### GetPkiDomainIDOk

`func (o *DomainResponse) GetPkiDomainIDOk() (*int32, bool)`

GetPkiDomainIDOk returns a tuple with the PkiDomainID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiDomainID

`func (o *DomainResponse) SetPkiDomainID(v int32)`

SetPkiDomainID sets PkiDomainID field to given value.


### GetSDomainName

`func (o *DomainResponse) GetSDomainName() string`

GetSDomainName returns the SDomainName field if non-nil, zero value otherwise.

### GetSDomainNameOk

`func (o *DomainResponse) GetSDomainNameOk() (*string, bool)`

GetSDomainNameOk returns a tuple with the SDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSDomainName

`func (o *DomainResponse) SetSDomainName(v string)`

SetSDomainName sets SDomainName field to given value.


### GetBDomainValiddkim

`func (o *DomainResponse) GetBDomainValiddkim() bool`

GetBDomainValiddkim returns the BDomainValiddkim field if non-nil, zero value otherwise.

### GetBDomainValiddkimOk

`func (o *DomainResponse) GetBDomainValiddkimOk() (*bool, bool)`

GetBDomainValiddkimOk returns a tuple with the BDomainValiddkim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDomainValiddkim

`func (o *DomainResponse) SetBDomainValiddkim(v bool)`

SetBDomainValiddkim sets BDomainValiddkim field to given value.


### GetBDomainValidmailfrom

`func (o *DomainResponse) GetBDomainValidmailfrom() bool`

GetBDomainValidmailfrom returns the BDomainValidmailfrom field if non-nil, zero value otherwise.

### GetBDomainValidmailfromOk

`func (o *DomainResponse) GetBDomainValidmailfromOk() (*bool, bool)`

GetBDomainValidmailfromOk returns a tuple with the BDomainValidmailfrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDomainValidmailfrom

`func (o *DomainResponse) SetBDomainValidmailfrom(v bool)`

SetBDomainValidmailfrom sets BDomainValidmailfrom field to given value.


### GetBDomainValidcustomer

`func (o *DomainResponse) GetBDomainValidcustomer() bool`

GetBDomainValidcustomer returns the BDomainValidcustomer field if non-nil, zero value otherwise.

### GetBDomainValidcustomerOk

`func (o *DomainResponse) GetBDomainValidcustomerOk() (*bool, bool)`

GetBDomainValidcustomerOk returns a tuple with the BDomainValidcustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDomainValidcustomer

`func (o *DomainResponse) SetBDomainValidcustomer(v bool)`

SetBDomainValidcustomer sets BDomainValidcustomer field to given value.


### GetObjAudit

`func (o *DomainResponse) GetObjAudit() CommonAudit`

GetObjAudit returns the ObjAudit field if non-nil, zero value otherwise.

### GetObjAuditOk

`func (o *DomainResponse) GetObjAuditOk() (*CommonAudit, bool)`

GetObjAuditOk returns a tuple with the ObjAudit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjAudit

`func (o *DomainResponse) SetObjAudit(v CommonAudit)`

SetObjAudit sets ObjAudit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


