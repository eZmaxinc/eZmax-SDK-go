# CustomEzsignfolderezsigntemplatepublicResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsignfolderID** | **int32** | The unique ID of the Ezsignfolder | 
**SEzsignfolderDescription** | **string** | The description of the Ezsignfolder | 
**EEzsignfolderStep** | [**FieldEEzsignfolderStep**](FieldEEzsignfolderStep.md) |  | 
**IEzsignfolderSignaturetotal** | **int32** | The number of total signatures that were requested in the Ezsignfolder | 
**IEzsignfolderFormfieldtotal** | **int32** | The number of total form fields that were requested in the Ezsignfolder | 
**IEzsignfolderSignaturesigned** | **int32** | The number of signatures that were signed in the Ezsignfolder. | 
**AObjEzsignfolderezsigntemplatepublicSigner** | Pointer to [**[]CustomEzsignfolderezsigntemplatepublicSignerResponse**](CustomEzsignfolderezsigntemplatepublicSignerResponse.md) |  | [optional] 

## Methods

### NewCustomEzsignfolderezsigntemplatepublicResponse

`func NewCustomEzsignfolderezsigntemplatepublicResponse(pkiEzsignfolderID int32, sEzsignfolderDescription string, eEzsignfolderStep FieldEEzsignfolderStep, iEzsignfolderSignaturetotal int32, iEzsignfolderFormfieldtotal int32, iEzsignfolderSignaturesigned int32, ) *CustomEzsignfolderezsigntemplatepublicResponse`

NewCustomEzsignfolderezsigntemplatepublicResponse instantiates a new CustomEzsignfolderezsigntemplatepublicResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomEzsignfolderezsigntemplatepublicResponseWithDefaults

`func NewCustomEzsignfolderezsigntemplatepublicResponseWithDefaults() *CustomEzsignfolderezsigntemplatepublicResponse`

NewCustomEzsignfolderezsigntemplatepublicResponseWithDefaults instantiates a new CustomEzsignfolderezsigntemplatepublicResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsignfolderID

`func (o *CustomEzsignfolderezsigntemplatepublicResponse) GetPkiEzsignfolderID() int32`

GetPkiEzsignfolderID returns the PkiEzsignfolderID field if non-nil, zero value otherwise.

### GetPkiEzsignfolderIDOk

`func (o *CustomEzsignfolderezsigntemplatepublicResponse) GetPkiEzsignfolderIDOk() (*int32, bool)`

GetPkiEzsignfolderIDOk returns a tuple with the PkiEzsignfolderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsignfolderID

`func (o *CustomEzsignfolderezsigntemplatepublicResponse) SetPkiEzsignfolderID(v int32)`

SetPkiEzsignfolderID sets PkiEzsignfolderID field to given value.


### GetSEzsignfolderDescription

`func (o *CustomEzsignfolderezsigntemplatepublicResponse) GetSEzsignfolderDescription() string`

GetSEzsignfolderDescription returns the SEzsignfolderDescription field if non-nil, zero value otherwise.

### GetSEzsignfolderDescriptionOk

`func (o *CustomEzsignfolderezsigntemplatepublicResponse) GetSEzsignfolderDescriptionOk() (*string, bool)`

GetSEzsignfolderDescriptionOk returns a tuple with the SEzsignfolderDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignfolderDescription

`func (o *CustomEzsignfolderezsigntemplatepublicResponse) SetSEzsignfolderDescription(v string)`

SetSEzsignfolderDescription sets SEzsignfolderDescription field to given value.


### GetEEzsignfolderStep

`func (o *CustomEzsignfolderezsigntemplatepublicResponse) GetEEzsignfolderStep() FieldEEzsignfolderStep`

GetEEzsignfolderStep returns the EEzsignfolderStep field if non-nil, zero value otherwise.

### GetEEzsignfolderStepOk

`func (o *CustomEzsignfolderezsigntemplatepublicResponse) GetEEzsignfolderStepOk() (*FieldEEzsignfolderStep, bool)`

GetEEzsignfolderStepOk returns a tuple with the EEzsignfolderStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignfolderStep

`func (o *CustomEzsignfolderezsigntemplatepublicResponse) SetEEzsignfolderStep(v FieldEEzsignfolderStep)`

SetEEzsignfolderStep sets EEzsignfolderStep field to given value.


### GetIEzsignfolderSignaturetotal

`func (o *CustomEzsignfolderezsigntemplatepublicResponse) GetIEzsignfolderSignaturetotal() int32`

GetIEzsignfolderSignaturetotal returns the IEzsignfolderSignaturetotal field if non-nil, zero value otherwise.

### GetIEzsignfolderSignaturetotalOk

`func (o *CustomEzsignfolderezsigntemplatepublicResponse) GetIEzsignfolderSignaturetotalOk() (*int32, bool)`

GetIEzsignfolderSignaturetotalOk returns a tuple with the IEzsignfolderSignaturetotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignfolderSignaturetotal

`func (o *CustomEzsignfolderezsigntemplatepublicResponse) SetIEzsignfolderSignaturetotal(v int32)`

SetIEzsignfolderSignaturetotal sets IEzsignfolderSignaturetotal field to given value.


### GetIEzsignfolderFormfieldtotal

`func (o *CustomEzsignfolderezsigntemplatepublicResponse) GetIEzsignfolderFormfieldtotal() int32`

GetIEzsignfolderFormfieldtotal returns the IEzsignfolderFormfieldtotal field if non-nil, zero value otherwise.

### GetIEzsignfolderFormfieldtotalOk

`func (o *CustomEzsignfolderezsigntemplatepublicResponse) GetIEzsignfolderFormfieldtotalOk() (*int32, bool)`

GetIEzsignfolderFormfieldtotalOk returns a tuple with the IEzsignfolderFormfieldtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignfolderFormfieldtotal

`func (o *CustomEzsignfolderezsigntemplatepublicResponse) SetIEzsignfolderFormfieldtotal(v int32)`

SetIEzsignfolderFormfieldtotal sets IEzsignfolderFormfieldtotal field to given value.


### GetIEzsignfolderSignaturesigned

`func (o *CustomEzsignfolderezsigntemplatepublicResponse) GetIEzsignfolderSignaturesigned() int32`

GetIEzsignfolderSignaturesigned returns the IEzsignfolderSignaturesigned field if non-nil, zero value otherwise.

### GetIEzsignfolderSignaturesignedOk

`func (o *CustomEzsignfolderezsigntemplatepublicResponse) GetIEzsignfolderSignaturesignedOk() (*int32, bool)`

GetIEzsignfolderSignaturesignedOk returns a tuple with the IEzsignfolderSignaturesigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignfolderSignaturesigned

`func (o *CustomEzsignfolderezsigntemplatepublicResponse) SetIEzsignfolderSignaturesigned(v int32)`

SetIEzsignfolderSignaturesigned sets IEzsignfolderSignaturesigned field to given value.


### GetAObjEzsignfolderezsigntemplatepublicSigner

`func (o *CustomEzsignfolderezsigntemplatepublicResponse) GetAObjEzsignfolderezsigntemplatepublicSigner() []CustomEzsignfolderezsigntemplatepublicSignerResponse`

GetAObjEzsignfolderezsigntemplatepublicSigner returns the AObjEzsignfolderezsigntemplatepublicSigner field if non-nil, zero value otherwise.

### GetAObjEzsignfolderezsigntemplatepublicSignerOk

`func (o *CustomEzsignfolderezsigntemplatepublicResponse) GetAObjEzsignfolderezsigntemplatepublicSignerOk() (*[]CustomEzsignfolderezsigntemplatepublicSignerResponse, bool)`

GetAObjEzsignfolderezsigntemplatepublicSignerOk returns a tuple with the AObjEzsignfolderezsigntemplatepublicSigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzsignfolderezsigntemplatepublicSigner

`func (o *CustomEzsignfolderezsigntemplatepublicResponse) SetAObjEzsignfolderezsigntemplatepublicSigner(v []CustomEzsignfolderezsigntemplatepublicSignerResponse)`

SetAObjEzsignfolderezsigntemplatepublicSigner sets AObjEzsignfolderezsigntemplatepublicSigner field to given value.

### HasAObjEzsignfolderezsigntemplatepublicSigner

`func (o *CustomEzsignfolderezsigntemplatepublicResponse) HasAObjEzsignfolderezsigntemplatepublicSigner() bool`

HasAObjEzsignfolderezsigntemplatepublicSigner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


