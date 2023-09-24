# EzsigntemplatedocumentResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsigntemplatedocumentID** | **int32** | The unique ID of the Ezsigntemplatedocument | 
**FkiEzsigntemplateID** | **int32** | The unique ID of the Ezsigntemplate | 
**SEzsigntemplatedocumentName** | **string** | The name of the Ezsigntemplatedocument. | 
**IEzsigntemplatedocumentPagetotal** | **int32** | The number of pages in the Ezsigntemplatedocument. | 
**IEzsigntemplatedocumentSignaturetotal** | **int32** | The number of total signatures in the Ezsigntemplate. | 
**BEzsigntemplatedocumentHassignedsignatures** | **bool** | If the Ezsigntemplatedocument contains signed signatures (From internal or external sources) | 

## Methods

### NewEzsigntemplatedocumentResponseCompound

`func NewEzsigntemplatedocumentResponseCompound(pkiEzsigntemplatedocumentID int32, fkiEzsigntemplateID int32, sEzsigntemplatedocumentName string, iEzsigntemplatedocumentPagetotal int32, iEzsigntemplatedocumentSignaturetotal int32, bEzsigntemplatedocumentHassignedsignatures bool, ) *EzsigntemplatedocumentResponseCompound`

NewEzsigntemplatedocumentResponseCompound instantiates a new EzsigntemplatedocumentResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigntemplatedocumentResponseCompoundWithDefaults

`func NewEzsigntemplatedocumentResponseCompoundWithDefaults() *EzsigntemplatedocumentResponseCompound`

NewEzsigntemplatedocumentResponseCompoundWithDefaults instantiates a new EzsigntemplatedocumentResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsigntemplatedocumentID

`func (o *EzsigntemplatedocumentResponseCompound) GetPkiEzsigntemplatedocumentID() int32`

GetPkiEzsigntemplatedocumentID returns the PkiEzsigntemplatedocumentID field if non-nil, zero value otherwise.

### GetPkiEzsigntemplatedocumentIDOk

`func (o *EzsigntemplatedocumentResponseCompound) GetPkiEzsigntemplatedocumentIDOk() (*int32, bool)`

GetPkiEzsigntemplatedocumentIDOk returns a tuple with the PkiEzsigntemplatedocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsigntemplatedocumentID

`func (o *EzsigntemplatedocumentResponseCompound) SetPkiEzsigntemplatedocumentID(v int32)`

SetPkiEzsigntemplatedocumentID sets PkiEzsigntemplatedocumentID field to given value.


### GetFkiEzsigntemplateID

`func (o *EzsigntemplatedocumentResponseCompound) GetFkiEzsigntemplateID() int32`

GetFkiEzsigntemplateID returns the FkiEzsigntemplateID field if non-nil, zero value otherwise.

### GetFkiEzsigntemplateIDOk

`func (o *EzsigntemplatedocumentResponseCompound) GetFkiEzsigntemplateIDOk() (*int32, bool)`

GetFkiEzsigntemplateIDOk returns a tuple with the FkiEzsigntemplateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntemplateID

`func (o *EzsigntemplatedocumentResponseCompound) SetFkiEzsigntemplateID(v int32)`

SetFkiEzsigntemplateID sets FkiEzsigntemplateID field to given value.


### GetSEzsigntemplatedocumentName

`func (o *EzsigntemplatedocumentResponseCompound) GetSEzsigntemplatedocumentName() string`

GetSEzsigntemplatedocumentName returns the SEzsigntemplatedocumentName field if non-nil, zero value otherwise.

### GetSEzsigntemplatedocumentNameOk

`func (o *EzsigntemplatedocumentResponseCompound) GetSEzsigntemplatedocumentNameOk() (*string, bool)`

GetSEzsigntemplatedocumentNameOk returns a tuple with the SEzsigntemplatedocumentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplatedocumentName

`func (o *EzsigntemplatedocumentResponseCompound) SetSEzsigntemplatedocumentName(v string)`

SetSEzsigntemplatedocumentName sets SEzsigntemplatedocumentName field to given value.


### GetIEzsigntemplatedocumentPagetotal

`func (o *EzsigntemplatedocumentResponseCompound) GetIEzsigntemplatedocumentPagetotal() int32`

GetIEzsigntemplatedocumentPagetotal returns the IEzsigntemplatedocumentPagetotal field if non-nil, zero value otherwise.

### GetIEzsigntemplatedocumentPagetotalOk

`func (o *EzsigntemplatedocumentResponseCompound) GetIEzsigntemplatedocumentPagetotalOk() (*int32, bool)`

GetIEzsigntemplatedocumentPagetotalOk returns a tuple with the IEzsigntemplatedocumentPagetotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplatedocumentPagetotal

`func (o *EzsigntemplatedocumentResponseCompound) SetIEzsigntemplatedocumentPagetotal(v int32)`

SetIEzsigntemplatedocumentPagetotal sets IEzsigntemplatedocumentPagetotal field to given value.


### GetIEzsigntemplatedocumentSignaturetotal

`func (o *EzsigntemplatedocumentResponseCompound) GetIEzsigntemplatedocumentSignaturetotal() int32`

GetIEzsigntemplatedocumentSignaturetotal returns the IEzsigntemplatedocumentSignaturetotal field if non-nil, zero value otherwise.

### GetIEzsigntemplatedocumentSignaturetotalOk

`func (o *EzsigntemplatedocumentResponseCompound) GetIEzsigntemplatedocumentSignaturetotalOk() (*int32, bool)`

GetIEzsigntemplatedocumentSignaturetotalOk returns a tuple with the IEzsigntemplatedocumentSignaturetotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplatedocumentSignaturetotal

`func (o *EzsigntemplatedocumentResponseCompound) SetIEzsigntemplatedocumentSignaturetotal(v int32)`

SetIEzsigntemplatedocumentSignaturetotal sets IEzsigntemplatedocumentSignaturetotal field to given value.


### GetBEzsigntemplatedocumentHassignedsignatures

`func (o *EzsigntemplatedocumentResponseCompound) GetBEzsigntemplatedocumentHassignedsignatures() bool`

GetBEzsigntemplatedocumentHassignedsignatures returns the BEzsigntemplatedocumentHassignedsignatures field if non-nil, zero value otherwise.

### GetBEzsigntemplatedocumentHassignedsignaturesOk

`func (o *EzsigntemplatedocumentResponseCompound) GetBEzsigntemplatedocumentHassignedsignaturesOk() (*bool, bool)`

GetBEzsigntemplatedocumentHassignedsignaturesOk returns a tuple with the BEzsigntemplatedocumentHassignedsignatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigntemplatedocumentHassignedsignatures

`func (o *EzsigntemplatedocumentResponseCompound) SetBEzsigntemplatedocumentHassignedsignatures(v bool)`

SetBEzsigntemplatedocumentHassignedsignatures sets BEzsigntemplatedocumentHassignedsignatures field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


