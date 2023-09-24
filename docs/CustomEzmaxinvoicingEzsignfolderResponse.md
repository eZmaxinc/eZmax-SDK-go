# CustomEzmaxinvoicingEzsignfolderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FkiEzsignfolderID** | **int32** | The unique ID of the Ezsignfolder | 
**SEzsignfolderDescription** | **string** | The description of the Ezsignfolder | 
**BEzsigntsarequirementBillable** | **bool** | Whether the TSA requirement is billable or not | 
**BEzsignfolderMfaused** | **bool** | Whether the MFA was used or not for the Ezsignfolder | 
**BEzsignfolderPaymentused** | **bool** | Whether there was a signature is of type payment | 
**BEzsignfolderAllowed** | **bool** | Whether you have access to the Ezsignfolder or not | 

## Methods

### NewCustomEzmaxinvoicingEzsignfolderResponse

`func NewCustomEzmaxinvoicingEzsignfolderResponse(fkiEzsignfolderID int32, sEzsignfolderDescription string, bEzsigntsarequirementBillable bool, bEzsignfolderMfaused bool, bEzsignfolderPaymentused bool, bEzsignfolderAllowed bool, ) *CustomEzmaxinvoicingEzsignfolderResponse`

NewCustomEzmaxinvoicingEzsignfolderResponse instantiates a new CustomEzmaxinvoicingEzsignfolderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomEzmaxinvoicingEzsignfolderResponseWithDefaults

`func NewCustomEzmaxinvoicingEzsignfolderResponseWithDefaults() *CustomEzmaxinvoicingEzsignfolderResponse`

NewCustomEzmaxinvoicingEzsignfolderResponseWithDefaults instantiates a new CustomEzmaxinvoicingEzsignfolderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFkiEzsignfolderID

`func (o *CustomEzmaxinvoicingEzsignfolderResponse) GetFkiEzsignfolderID() int32`

GetFkiEzsignfolderID returns the FkiEzsignfolderID field if non-nil, zero value otherwise.

### GetFkiEzsignfolderIDOk

`func (o *CustomEzmaxinvoicingEzsignfolderResponse) GetFkiEzsignfolderIDOk() (*int32, bool)`

GetFkiEzsignfolderIDOk returns a tuple with the FkiEzsignfolderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfolderID

`func (o *CustomEzmaxinvoicingEzsignfolderResponse) SetFkiEzsignfolderID(v int32)`

SetFkiEzsignfolderID sets FkiEzsignfolderID field to given value.


### GetSEzsignfolderDescription

`func (o *CustomEzmaxinvoicingEzsignfolderResponse) GetSEzsignfolderDescription() string`

GetSEzsignfolderDescription returns the SEzsignfolderDescription field if non-nil, zero value otherwise.

### GetSEzsignfolderDescriptionOk

`func (o *CustomEzmaxinvoicingEzsignfolderResponse) GetSEzsignfolderDescriptionOk() (*string, bool)`

GetSEzsignfolderDescriptionOk returns a tuple with the SEzsignfolderDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignfolderDescription

`func (o *CustomEzmaxinvoicingEzsignfolderResponse) SetSEzsignfolderDescription(v string)`

SetSEzsignfolderDescription sets SEzsignfolderDescription field to given value.


### GetBEzsigntsarequirementBillable

`func (o *CustomEzmaxinvoicingEzsignfolderResponse) GetBEzsigntsarequirementBillable() bool`

GetBEzsigntsarequirementBillable returns the BEzsigntsarequirementBillable field if non-nil, zero value otherwise.

### GetBEzsigntsarequirementBillableOk

`func (o *CustomEzmaxinvoicingEzsignfolderResponse) GetBEzsigntsarequirementBillableOk() (*bool, bool)`

GetBEzsigntsarequirementBillableOk returns a tuple with the BEzsigntsarequirementBillable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigntsarequirementBillable

`func (o *CustomEzmaxinvoicingEzsignfolderResponse) SetBEzsigntsarequirementBillable(v bool)`

SetBEzsigntsarequirementBillable sets BEzsigntsarequirementBillable field to given value.


### GetBEzsignfolderMfaused

`func (o *CustomEzmaxinvoicingEzsignfolderResponse) GetBEzsignfolderMfaused() bool`

GetBEzsignfolderMfaused returns the BEzsignfolderMfaused field if non-nil, zero value otherwise.

### GetBEzsignfolderMfausedOk

`func (o *CustomEzmaxinvoicingEzsignfolderResponse) GetBEzsignfolderMfausedOk() (*bool, bool)`

GetBEzsignfolderMfausedOk returns a tuple with the BEzsignfolderMfaused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfolderMfaused

`func (o *CustomEzmaxinvoicingEzsignfolderResponse) SetBEzsignfolderMfaused(v bool)`

SetBEzsignfolderMfaused sets BEzsignfolderMfaused field to given value.


### GetBEzsignfolderPaymentused

`func (o *CustomEzmaxinvoicingEzsignfolderResponse) GetBEzsignfolderPaymentused() bool`

GetBEzsignfolderPaymentused returns the BEzsignfolderPaymentused field if non-nil, zero value otherwise.

### GetBEzsignfolderPaymentusedOk

`func (o *CustomEzmaxinvoicingEzsignfolderResponse) GetBEzsignfolderPaymentusedOk() (*bool, bool)`

GetBEzsignfolderPaymentusedOk returns a tuple with the BEzsignfolderPaymentused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfolderPaymentused

`func (o *CustomEzmaxinvoicingEzsignfolderResponse) SetBEzsignfolderPaymentused(v bool)`

SetBEzsignfolderPaymentused sets BEzsignfolderPaymentused field to given value.


### GetBEzsignfolderAllowed

`func (o *CustomEzmaxinvoicingEzsignfolderResponse) GetBEzsignfolderAllowed() bool`

GetBEzsignfolderAllowed returns the BEzsignfolderAllowed field if non-nil, zero value otherwise.

### GetBEzsignfolderAllowedOk

`func (o *CustomEzmaxinvoicingEzsignfolderResponse) GetBEzsignfolderAllowedOk() (*bool, bool)`

GetBEzsignfolderAllowedOk returns a tuple with the BEzsignfolderAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfolderAllowed

`func (o *CustomEzmaxinvoicingEzsignfolderResponse) SetBEzsignfolderAllowed(v bool)`

SetBEzsignfolderAllowed sets BEzsignfolderAllowed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


