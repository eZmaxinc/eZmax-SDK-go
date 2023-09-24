# TaxassignmentAutocompleteElementResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**STaxassignmentDescriptionX** | **string** | The description of the Taxassignment  in the language of the requester | 
**PkiTaxassignmentID** | **int32** | The unique ID of the Taxassignment.  Valid values:  |Value|Description| |-|-| |1|No tax| |2|GST| |3|HST (ON)| |4|HST (NB)| |5|HST (NS)| |6|HST (NL)| |7|HST (PE)| |8|GST + QST (QC)| |9|GST + QST (QC) Non-Recoverable| |10|GST + PST (BC)| |11|GST + PST (SK)| |12|GST + RST (MB)| |13|GST + PST (BC) Non-Recoverable| |14|GST + PST (SK) Non-Recoverable| |15|GST + RST (MB) Non-Recoverable| | 
**BTaxassignmentIsactive** | **bool** | Whether the Taxassignment is active or not | 

## Methods

### NewTaxassignmentAutocompleteElementResponse

`func NewTaxassignmentAutocompleteElementResponse(sTaxassignmentDescriptionX string, pkiTaxassignmentID int32, bTaxassignmentIsactive bool, ) *TaxassignmentAutocompleteElementResponse`

NewTaxassignmentAutocompleteElementResponse instantiates a new TaxassignmentAutocompleteElementResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxassignmentAutocompleteElementResponseWithDefaults

`func NewTaxassignmentAutocompleteElementResponseWithDefaults() *TaxassignmentAutocompleteElementResponse`

NewTaxassignmentAutocompleteElementResponseWithDefaults instantiates a new TaxassignmentAutocompleteElementResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSTaxassignmentDescriptionX

`func (o *TaxassignmentAutocompleteElementResponse) GetSTaxassignmentDescriptionX() string`

GetSTaxassignmentDescriptionX returns the STaxassignmentDescriptionX field if non-nil, zero value otherwise.

### GetSTaxassignmentDescriptionXOk

`func (o *TaxassignmentAutocompleteElementResponse) GetSTaxassignmentDescriptionXOk() (*string, bool)`

GetSTaxassignmentDescriptionXOk returns a tuple with the STaxassignmentDescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSTaxassignmentDescriptionX

`func (o *TaxassignmentAutocompleteElementResponse) SetSTaxassignmentDescriptionX(v string)`

SetSTaxassignmentDescriptionX sets STaxassignmentDescriptionX field to given value.


### GetPkiTaxassignmentID

`func (o *TaxassignmentAutocompleteElementResponse) GetPkiTaxassignmentID() int32`

GetPkiTaxassignmentID returns the PkiTaxassignmentID field if non-nil, zero value otherwise.

### GetPkiTaxassignmentIDOk

`func (o *TaxassignmentAutocompleteElementResponse) GetPkiTaxassignmentIDOk() (*int32, bool)`

GetPkiTaxassignmentIDOk returns a tuple with the PkiTaxassignmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiTaxassignmentID

`func (o *TaxassignmentAutocompleteElementResponse) SetPkiTaxassignmentID(v int32)`

SetPkiTaxassignmentID sets PkiTaxassignmentID field to given value.


### GetBTaxassignmentIsactive

`func (o *TaxassignmentAutocompleteElementResponse) GetBTaxassignmentIsactive() bool`

GetBTaxassignmentIsactive returns the BTaxassignmentIsactive field if non-nil, zero value otherwise.

### GetBTaxassignmentIsactiveOk

`func (o *TaxassignmentAutocompleteElementResponse) GetBTaxassignmentIsactiveOk() (*bool, bool)`

GetBTaxassignmentIsactiveOk returns a tuple with the BTaxassignmentIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBTaxassignmentIsactive

`func (o *TaxassignmentAutocompleteElementResponse) SetBTaxassignmentIsactive(v bool)`

SetBTaxassignmentIsactive sets BTaxassignmentIsactive field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


