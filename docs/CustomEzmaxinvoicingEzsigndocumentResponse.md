# CustomEzmaxinvoicingEzsigndocumentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FkiEzsignfolderID** | **int32** | The unique ID of the Ezsignfolder | 
**FkiBillingentityinternalID** | Pointer to **int32** | The unique ID of the Billingentityinternal. | [optional] 
**SName** | **string** |  | 
**SEzsignfolderDescription** | **string** | The description of the Ezsignfolder | 
**SEzsigndocumentName** | **string** | The name of the document that will be presented to Ezsignfoldersignerassociations | 
**BEzsignfolderAllowed** | **bool** | Whether you have access to the Ezsignfolder or not | 

## Methods

### NewCustomEzmaxinvoicingEzsigndocumentResponse

`func NewCustomEzmaxinvoicingEzsigndocumentResponse(fkiEzsignfolderID int32, sName string, sEzsignfolderDescription string, sEzsigndocumentName string, bEzsignfolderAllowed bool, ) *CustomEzmaxinvoicingEzsigndocumentResponse`

NewCustomEzmaxinvoicingEzsigndocumentResponse instantiates a new CustomEzmaxinvoicingEzsigndocumentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomEzmaxinvoicingEzsigndocumentResponseWithDefaults

`func NewCustomEzmaxinvoicingEzsigndocumentResponseWithDefaults() *CustomEzmaxinvoicingEzsigndocumentResponse`

NewCustomEzmaxinvoicingEzsigndocumentResponseWithDefaults instantiates a new CustomEzmaxinvoicingEzsigndocumentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFkiEzsignfolderID

`func (o *CustomEzmaxinvoicingEzsigndocumentResponse) GetFkiEzsignfolderID() int32`

GetFkiEzsignfolderID returns the FkiEzsignfolderID field if non-nil, zero value otherwise.

### GetFkiEzsignfolderIDOk

`func (o *CustomEzmaxinvoicingEzsigndocumentResponse) GetFkiEzsignfolderIDOk() (*int32, bool)`

GetFkiEzsignfolderIDOk returns a tuple with the FkiEzsignfolderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfolderID

`func (o *CustomEzmaxinvoicingEzsigndocumentResponse) SetFkiEzsignfolderID(v int32)`

SetFkiEzsignfolderID sets FkiEzsignfolderID field to given value.


### GetFkiBillingentityinternalID

`func (o *CustomEzmaxinvoicingEzsigndocumentResponse) GetFkiBillingentityinternalID() int32`

GetFkiBillingentityinternalID returns the FkiBillingentityinternalID field if non-nil, zero value otherwise.

### GetFkiBillingentityinternalIDOk

`func (o *CustomEzmaxinvoicingEzsigndocumentResponse) GetFkiBillingentityinternalIDOk() (*int32, bool)`

GetFkiBillingentityinternalIDOk returns a tuple with the FkiBillingentityinternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiBillingentityinternalID

`func (o *CustomEzmaxinvoicingEzsigndocumentResponse) SetFkiBillingentityinternalID(v int32)`

SetFkiBillingentityinternalID sets FkiBillingentityinternalID field to given value.

### HasFkiBillingentityinternalID

`func (o *CustomEzmaxinvoicingEzsigndocumentResponse) HasFkiBillingentityinternalID() bool`

HasFkiBillingentityinternalID returns a boolean if a field has been set.

### GetSName

`func (o *CustomEzmaxinvoicingEzsigndocumentResponse) GetSName() string`

GetSName returns the SName field if non-nil, zero value otherwise.

### GetSNameOk

`func (o *CustomEzmaxinvoicingEzsigndocumentResponse) GetSNameOk() (*string, bool)`

GetSNameOk returns a tuple with the SName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSName

`func (o *CustomEzmaxinvoicingEzsigndocumentResponse) SetSName(v string)`

SetSName sets SName field to given value.


### GetSEzsignfolderDescription

`func (o *CustomEzmaxinvoicingEzsigndocumentResponse) GetSEzsignfolderDescription() string`

GetSEzsignfolderDescription returns the SEzsignfolderDescription field if non-nil, zero value otherwise.

### GetSEzsignfolderDescriptionOk

`func (o *CustomEzmaxinvoicingEzsigndocumentResponse) GetSEzsignfolderDescriptionOk() (*string, bool)`

GetSEzsignfolderDescriptionOk returns a tuple with the SEzsignfolderDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignfolderDescription

`func (o *CustomEzmaxinvoicingEzsigndocumentResponse) SetSEzsignfolderDescription(v string)`

SetSEzsignfolderDescription sets SEzsignfolderDescription field to given value.


### GetSEzsigndocumentName

`func (o *CustomEzmaxinvoicingEzsigndocumentResponse) GetSEzsigndocumentName() string`

GetSEzsigndocumentName returns the SEzsigndocumentName field if non-nil, zero value otherwise.

### GetSEzsigndocumentNameOk

`func (o *CustomEzmaxinvoicingEzsigndocumentResponse) GetSEzsigndocumentNameOk() (*string, bool)`

GetSEzsigndocumentNameOk returns a tuple with the SEzsigndocumentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigndocumentName

`func (o *CustomEzmaxinvoicingEzsigndocumentResponse) SetSEzsigndocumentName(v string)`

SetSEzsigndocumentName sets SEzsigndocumentName field to given value.


### GetBEzsignfolderAllowed

`func (o *CustomEzmaxinvoicingEzsigndocumentResponse) GetBEzsignfolderAllowed() bool`

GetBEzsignfolderAllowed returns the BEzsignfolderAllowed field if non-nil, zero value otherwise.

### GetBEzsignfolderAllowedOk

`func (o *CustomEzmaxinvoicingEzsigndocumentResponse) GetBEzsignfolderAllowedOk() (*bool, bool)`

GetBEzsignfolderAllowedOk returns a tuple with the BEzsignfolderAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfolderAllowed

`func (o *CustomEzmaxinvoicingEzsigndocumentResponse) SetBEzsignfolderAllowed(v bool)`

SetBEzsignfolderAllowed sets BEzsignfolderAllowed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


