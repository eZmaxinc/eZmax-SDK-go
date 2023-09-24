# EzsigntsarequirementAutocompleteElementResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SEzsigntsarequirementDescriptionX** | **string** | The description of the Ezsigntsarequirement in the language of the requester | 
**PkiEzsigntsarequirementID** | **int32** | The unique ID of the Ezsigntsarequirement.  Determine if a Time Stamping Authority should add a timestamp on each of the signature. Valid values:  |Value|Description| |-|-| |1|No. TSA Timestamping will requested. This will make all signatures a lot faster since no round-trip to the TSA server will be required. Timestamping will be made using eZsign server&#39;s time.| |2|Best effort. Timestamping from a Time Stamping Authority will be requested but is not mandatory. In the very improbable case it cannot be completed, the timestamping will be made using eZsign server&#39;s time. **Additional fee applies**| |3|Mandatory. Timestamping from a Time Stamping Authority will be requested and is mandatory. In the very improbable case it cannot be completed, the signature will fail and the user will be asked to retry. **Additional fee applies**| | 
**BEzsigntsarequirementIsactive** | **bool** | Whether the Ezsigntsarequirement is active or not | 
**BDisabled** | **bool** | Indicates if the element is disabled in the context | 

## Methods

### NewEzsigntsarequirementAutocompleteElementResponse

`func NewEzsigntsarequirementAutocompleteElementResponse(sEzsigntsarequirementDescriptionX string, pkiEzsigntsarequirementID int32, bEzsigntsarequirementIsactive bool, bDisabled bool, ) *EzsigntsarequirementAutocompleteElementResponse`

NewEzsigntsarequirementAutocompleteElementResponse instantiates a new EzsigntsarequirementAutocompleteElementResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigntsarequirementAutocompleteElementResponseWithDefaults

`func NewEzsigntsarequirementAutocompleteElementResponseWithDefaults() *EzsigntsarequirementAutocompleteElementResponse`

NewEzsigntsarequirementAutocompleteElementResponseWithDefaults instantiates a new EzsigntsarequirementAutocompleteElementResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSEzsigntsarequirementDescriptionX

`func (o *EzsigntsarequirementAutocompleteElementResponse) GetSEzsigntsarequirementDescriptionX() string`

GetSEzsigntsarequirementDescriptionX returns the SEzsigntsarequirementDescriptionX field if non-nil, zero value otherwise.

### GetSEzsigntsarequirementDescriptionXOk

`func (o *EzsigntsarequirementAutocompleteElementResponse) GetSEzsigntsarequirementDescriptionXOk() (*string, bool)`

GetSEzsigntsarequirementDescriptionXOk returns a tuple with the SEzsigntsarequirementDescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntsarequirementDescriptionX

`func (o *EzsigntsarequirementAutocompleteElementResponse) SetSEzsigntsarequirementDescriptionX(v string)`

SetSEzsigntsarequirementDescriptionX sets SEzsigntsarequirementDescriptionX field to given value.


### GetPkiEzsigntsarequirementID

`func (o *EzsigntsarequirementAutocompleteElementResponse) GetPkiEzsigntsarequirementID() int32`

GetPkiEzsigntsarequirementID returns the PkiEzsigntsarequirementID field if non-nil, zero value otherwise.

### GetPkiEzsigntsarequirementIDOk

`func (o *EzsigntsarequirementAutocompleteElementResponse) GetPkiEzsigntsarequirementIDOk() (*int32, bool)`

GetPkiEzsigntsarequirementIDOk returns a tuple with the PkiEzsigntsarequirementID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsigntsarequirementID

`func (o *EzsigntsarequirementAutocompleteElementResponse) SetPkiEzsigntsarequirementID(v int32)`

SetPkiEzsigntsarequirementID sets PkiEzsigntsarequirementID field to given value.


### GetBEzsigntsarequirementIsactive

`func (o *EzsigntsarequirementAutocompleteElementResponse) GetBEzsigntsarequirementIsactive() bool`

GetBEzsigntsarequirementIsactive returns the BEzsigntsarequirementIsactive field if non-nil, zero value otherwise.

### GetBEzsigntsarequirementIsactiveOk

`func (o *EzsigntsarequirementAutocompleteElementResponse) GetBEzsigntsarequirementIsactiveOk() (*bool, bool)`

GetBEzsigntsarequirementIsactiveOk returns a tuple with the BEzsigntsarequirementIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigntsarequirementIsactive

`func (o *EzsigntsarequirementAutocompleteElementResponse) SetBEzsigntsarequirementIsactive(v bool)`

SetBEzsigntsarequirementIsactive sets BEzsigntsarequirementIsactive field to given value.


### GetBDisabled

`func (o *EzsigntsarequirementAutocompleteElementResponse) GetBDisabled() bool`

GetBDisabled returns the BDisabled field if non-nil, zero value otherwise.

### GetBDisabledOk

`func (o *EzsigntsarequirementAutocompleteElementResponse) GetBDisabledOk() (*bool, bool)`

GetBDisabledOk returns a tuple with the BDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDisabled

`func (o *EzsigntsarequirementAutocompleteElementResponse) SetBDisabled(v bool)`

SetBDisabled sets BDisabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


