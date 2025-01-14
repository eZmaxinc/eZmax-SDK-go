# EzsigntemplatedocumentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsigntemplatedocumentID** | **int32** | The unique ID of the Ezsigntemplatedocument | 
**FkiEzsigntemplateID** | **int32** | The unique ID of the Ezsigntemplate | 
**SEzsigntemplatedocumentName** | **string** | The name of the Ezsigntemplatedocument. | 
**IEzsigntemplatedocumentPagetotal** | **int32** | The number of pages in the Ezsigntemplatedocument. | 
**IEzsigntemplatedocumentSignaturetotal** | **int32** | The number of total signatures in the Ezsigntemplate. | 
**IEzsigntemplatedocumentFormfieldtotal** | **int32** | The number of total form fields in the Ezsigntemplate. | 
**BEzsigntemplatedocumentHassignedsignatures** | **bool** | If the Ezsigntemplatedocument contains signed signatures (From internal or external sources) | 

## Methods

### NewEzsigntemplatedocumentResponse

`func NewEzsigntemplatedocumentResponse(pkiEzsigntemplatedocumentID int32, fkiEzsigntemplateID int32, sEzsigntemplatedocumentName string, iEzsigntemplatedocumentPagetotal int32, iEzsigntemplatedocumentSignaturetotal int32, iEzsigntemplatedocumentFormfieldtotal int32, bEzsigntemplatedocumentHassignedsignatures bool, ) *EzsigntemplatedocumentResponse`

NewEzsigntemplatedocumentResponse instantiates a new EzsigntemplatedocumentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigntemplatedocumentResponseWithDefaults

`func NewEzsigntemplatedocumentResponseWithDefaults() *EzsigntemplatedocumentResponse`

NewEzsigntemplatedocumentResponseWithDefaults instantiates a new EzsigntemplatedocumentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsigntemplatedocumentID

`func (o *EzsigntemplatedocumentResponse) GetPkiEzsigntemplatedocumentID() int32`

GetPkiEzsigntemplatedocumentID returns the PkiEzsigntemplatedocumentID field if non-nil, zero value otherwise.

### GetPkiEzsigntemplatedocumentIDOk

`func (o *EzsigntemplatedocumentResponse) GetPkiEzsigntemplatedocumentIDOk() (*int32, bool)`

GetPkiEzsigntemplatedocumentIDOk returns a tuple with the PkiEzsigntemplatedocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsigntemplatedocumentID

`func (o *EzsigntemplatedocumentResponse) SetPkiEzsigntemplatedocumentID(v int32)`

SetPkiEzsigntemplatedocumentID sets PkiEzsigntemplatedocumentID field to given value.


### GetFkiEzsigntemplateID

`func (o *EzsigntemplatedocumentResponse) GetFkiEzsigntemplateID() int32`

GetFkiEzsigntemplateID returns the FkiEzsigntemplateID field if non-nil, zero value otherwise.

### GetFkiEzsigntemplateIDOk

`func (o *EzsigntemplatedocumentResponse) GetFkiEzsigntemplateIDOk() (*int32, bool)`

GetFkiEzsigntemplateIDOk returns a tuple with the FkiEzsigntemplateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntemplateID

`func (o *EzsigntemplatedocumentResponse) SetFkiEzsigntemplateID(v int32)`

SetFkiEzsigntemplateID sets FkiEzsigntemplateID field to given value.


### GetSEzsigntemplatedocumentName

`func (o *EzsigntemplatedocumentResponse) GetSEzsigntemplatedocumentName() string`

GetSEzsigntemplatedocumentName returns the SEzsigntemplatedocumentName field if non-nil, zero value otherwise.

### GetSEzsigntemplatedocumentNameOk

`func (o *EzsigntemplatedocumentResponse) GetSEzsigntemplatedocumentNameOk() (*string, bool)`

GetSEzsigntemplatedocumentNameOk returns a tuple with the SEzsigntemplatedocumentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplatedocumentName

`func (o *EzsigntemplatedocumentResponse) SetSEzsigntemplatedocumentName(v string)`

SetSEzsigntemplatedocumentName sets SEzsigntemplatedocumentName field to given value.


### GetIEzsigntemplatedocumentPagetotal

`func (o *EzsigntemplatedocumentResponse) GetIEzsigntemplatedocumentPagetotal() int32`

GetIEzsigntemplatedocumentPagetotal returns the IEzsigntemplatedocumentPagetotal field if non-nil, zero value otherwise.

### GetIEzsigntemplatedocumentPagetotalOk

`func (o *EzsigntemplatedocumentResponse) GetIEzsigntemplatedocumentPagetotalOk() (*int32, bool)`

GetIEzsigntemplatedocumentPagetotalOk returns a tuple with the IEzsigntemplatedocumentPagetotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplatedocumentPagetotal

`func (o *EzsigntemplatedocumentResponse) SetIEzsigntemplatedocumentPagetotal(v int32)`

SetIEzsigntemplatedocumentPagetotal sets IEzsigntemplatedocumentPagetotal field to given value.


### GetIEzsigntemplatedocumentSignaturetotal

`func (o *EzsigntemplatedocumentResponse) GetIEzsigntemplatedocumentSignaturetotal() int32`

GetIEzsigntemplatedocumentSignaturetotal returns the IEzsigntemplatedocumentSignaturetotal field if non-nil, zero value otherwise.

### GetIEzsigntemplatedocumentSignaturetotalOk

`func (o *EzsigntemplatedocumentResponse) GetIEzsigntemplatedocumentSignaturetotalOk() (*int32, bool)`

GetIEzsigntemplatedocumentSignaturetotalOk returns a tuple with the IEzsigntemplatedocumentSignaturetotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplatedocumentSignaturetotal

`func (o *EzsigntemplatedocumentResponse) SetIEzsigntemplatedocumentSignaturetotal(v int32)`

SetIEzsigntemplatedocumentSignaturetotal sets IEzsigntemplatedocumentSignaturetotal field to given value.


### GetIEzsigntemplatedocumentFormfieldtotal

`func (o *EzsigntemplatedocumentResponse) GetIEzsigntemplatedocumentFormfieldtotal() int32`

GetIEzsigntemplatedocumentFormfieldtotal returns the IEzsigntemplatedocumentFormfieldtotal field if non-nil, zero value otherwise.

### GetIEzsigntemplatedocumentFormfieldtotalOk

`func (o *EzsigntemplatedocumentResponse) GetIEzsigntemplatedocumentFormfieldtotalOk() (*int32, bool)`

GetIEzsigntemplatedocumentFormfieldtotalOk returns a tuple with the IEzsigntemplatedocumentFormfieldtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplatedocumentFormfieldtotal

`func (o *EzsigntemplatedocumentResponse) SetIEzsigntemplatedocumentFormfieldtotal(v int32)`

SetIEzsigntemplatedocumentFormfieldtotal sets IEzsigntemplatedocumentFormfieldtotal field to given value.


### GetBEzsigntemplatedocumentHassignedsignatures

`func (o *EzsigntemplatedocumentResponse) GetBEzsigntemplatedocumentHassignedsignatures() bool`

GetBEzsigntemplatedocumentHassignedsignatures returns the BEzsigntemplatedocumentHassignedsignatures field if non-nil, zero value otherwise.

### GetBEzsigntemplatedocumentHassignedsignaturesOk

`func (o *EzsigntemplatedocumentResponse) GetBEzsigntemplatedocumentHassignedsignaturesOk() (*bool, bool)`

GetBEzsigntemplatedocumentHassignedsignaturesOk returns a tuple with the BEzsigntemplatedocumentHassignedsignatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigntemplatedocumentHassignedsignatures

`func (o *EzsigntemplatedocumentResponse) SetBEzsigntemplatedocumentHassignedsignatures(v bool)`

SetBEzsigntemplatedocumentHassignedsignatures sets BEzsigntemplatedocumentHassignedsignatures field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


