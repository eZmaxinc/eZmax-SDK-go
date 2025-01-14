# CustomEzsignfoldertransmissionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsignfolderID** | **int32** | The unique ID of the Ezsignfolder | 
**SEzsignfolderDescription** | **string** | The description of the Ezsignfolder | 
**EEzsignfolderStep** | [**FieldEEzsignfolderStep**](FieldEEzsignfolderStep.md) |  | 
**IEzsignfolderSignaturetotal** | **int32** | The number of total signatures that were requested in the Ezsignfolder | 
**IEzsignfolderFormfieldtotal** | **int32** | The number of total form fields that were requested in the Ezsignfolder | 
**IEzsignfolderSignaturesigned** | **int32** | The number of signatures that were signed in the Ezsignfolder. | 
**AObjEzsignfoldertransmissionSigner** | [**[]CustomEzsignfoldertransmissionSignerResponse**](CustomEzsignfoldertransmissionSignerResponse.md) |  | 

## Methods

### NewCustomEzsignfoldertransmissionResponse

`func NewCustomEzsignfoldertransmissionResponse(pkiEzsignfolderID int32, sEzsignfolderDescription string, eEzsignfolderStep FieldEEzsignfolderStep, iEzsignfolderSignaturetotal int32, iEzsignfolderFormfieldtotal int32, iEzsignfolderSignaturesigned int32, aObjEzsignfoldertransmissionSigner []CustomEzsignfoldertransmissionSignerResponse, ) *CustomEzsignfoldertransmissionResponse`

NewCustomEzsignfoldertransmissionResponse instantiates a new CustomEzsignfoldertransmissionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomEzsignfoldertransmissionResponseWithDefaults

`func NewCustomEzsignfoldertransmissionResponseWithDefaults() *CustomEzsignfoldertransmissionResponse`

NewCustomEzsignfoldertransmissionResponseWithDefaults instantiates a new CustomEzsignfoldertransmissionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsignfolderID

`func (o *CustomEzsignfoldertransmissionResponse) GetPkiEzsignfolderID() int32`

GetPkiEzsignfolderID returns the PkiEzsignfolderID field if non-nil, zero value otherwise.

### GetPkiEzsignfolderIDOk

`func (o *CustomEzsignfoldertransmissionResponse) GetPkiEzsignfolderIDOk() (*int32, bool)`

GetPkiEzsignfolderIDOk returns a tuple with the PkiEzsignfolderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsignfolderID

`func (o *CustomEzsignfoldertransmissionResponse) SetPkiEzsignfolderID(v int32)`

SetPkiEzsignfolderID sets PkiEzsignfolderID field to given value.


### GetSEzsignfolderDescription

`func (o *CustomEzsignfoldertransmissionResponse) GetSEzsignfolderDescription() string`

GetSEzsignfolderDescription returns the SEzsignfolderDescription field if non-nil, zero value otherwise.

### GetSEzsignfolderDescriptionOk

`func (o *CustomEzsignfoldertransmissionResponse) GetSEzsignfolderDescriptionOk() (*string, bool)`

GetSEzsignfolderDescriptionOk returns a tuple with the SEzsignfolderDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignfolderDescription

`func (o *CustomEzsignfoldertransmissionResponse) SetSEzsignfolderDescription(v string)`

SetSEzsignfolderDescription sets SEzsignfolderDescription field to given value.


### GetEEzsignfolderStep

`func (o *CustomEzsignfoldertransmissionResponse) GetEEzsignfolderStep() FieldEEzsignfolderStep`

GetEEzsignfolderStep returns the EEzsignfolderStep field if non-nil, zero value otherwise.

### GetEEzsignfolderStepOk

`func (o *CustomEzsignfoldertransmissionResponse) GetEEzsignfolderStepOk() (*FieldEEzsignfolderStep, bool)`

GetEEzsignfolderStepOk returns a tuple with the EEzsignfolderStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignfolderStep

`func (o *CustomEzsignfoldertransmissionResponse) SetEEzsignfolderStep(v FieldEEzsignfolderStep)`

SetEEzsignfolderStep sets EEzsignfolderStep field to given value.


### GetIEzsignfolderSignaturetotal

`func (o *CustomEzsignfoldertransmissionResponse) GetIEzsignfolderSignaturetotal() int32`

GetIEzsignfolderSignaturetotal returns the IEzsignfolderSignaturetotal field if non-nil, zero value otherwise.

### GetIEzsignfolderSignaturetotalOk

`func (o *CustomEzsignfoldertransmissionResponse) GetIEzsignfolderSignaturetotalOk() (*int32, bool)`

GetIEzsignfolderSignaturetotalOk returns a tuple with the IEzsignfolderSignaturetotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignfolderSignaturetotal

`func (o *CustomEzsignfoldertransmissionResponse) SetIEzsignfolderSignaturetotal(v int32)`

SetIEzsignfolderSignaturetotal sets IEzsignfolderSignaturetotal field to given value.


### GetIEzsignfolderFormfieldtotal

`func (o *CustomEzsignfoldertransmissionResponse) GetIEzsignfolderFormfieldtotal() int32`

GetIEzsignfolderFormfieldtotal returns the IEzsignfolderFormfieldtotal field if non-nil, zero value otherwise.

### GetIEzsignfolderFormfieldtotalOk

`func (o *CustomEzsignfoldertransmissionResponse) GetIEzsignfolderFormfieldtotalOk() (*int32, bool)`

GetIEzsignfolderFormfieldtotalOk returns a tuple with the IEzsignfolderFormfieldtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignfolderFormfieldtotal

`func (o *CustomEzsignfoldertransmissionResponse) SetIEzsignfolderFormfieldtotal(v int32)`

SetIEzsignfolderFormfieldtotal sets IEzsignfolderFormfieldtotal field to given value.


### GetIEzsignfolderSignaturesigned

`func (o *CustomEzsignfoldertransmissionResponse) GetIEzsignfolderSignaturesigned() int32`

GetIEzsignfolderSignaturesigned returns the IEzsignfolderSignaturesigned field if non-nil, zero value otherwise.

### GetIEzsignfolderSignaturesignedOk

`func (o *CustomEzsignfoldertransmissionResponse) GetIEzsignfolderSignaturesignedOk() (*int32, bool)`

GetIEzsignfolderSignaturesignedOk returns a tuple with the IEzsignfolderSignaturesigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignfolderSignaturesigned

`func (o *CustomEzsignfoldertransmissionResponse) SetIEzsignfolderSignaturesigned(v int32)`

SetIEzsignfolderSignaturesigned sets IEzsignfolderSignaturesigned field to given value.


### GetAObjEzsignfoldertransmissionSigner

`func (o *CustomEzsignfoldertransmissionResponse) GetAObjEzsignfoldertransmissionSigner() []CustomEzsignfoldertransmissionSignerResponse`

GetAObjEzsignfoldertransmissionSigner returns the AObjEzsignfoldertransmissionSigner field if non-nil, zero value otherwise.

### GetAObjEzsignfoldertransmissionSignerOk

`func (o *CustomEzsignfoldertransmissionResponse) GetAObjEzsignfoldertransmissionSignerOk() (*[]CustomEzsignfoldertransmissionSignerResponse, bool)`

GetAObjEzsignfoldertransmissionSignerOk returns a tuple with the AObjEzsignfoldertransmissionSigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzsignfoldertransmissionSigner

`func (o *CustomEzsignfoldertransmissionResponse) SetAObjEzsignfoldertransmissionSigner(v []CustomEzsignfoldertransmissionSignerResponse)`

SetAObjEzsignfoldertransmissionSigner sets AObjEzsignfoldertransmissionSigner field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


