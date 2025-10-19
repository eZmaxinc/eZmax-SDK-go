# BrokerAutocompleteElementResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiBrokerID** | **int32** | The unique ID of the Broker. | 
**FkiDepartmentID** | **int32** | The unique ID of the Department | 
**SBrokerName** | **string** | The name of the Broker | 
**BBrokerIsactive** | **bool** | Whether the Broker is active or not | 

## Methods

### NewBrokerAutocompleteElementResponse

`func NewBrokerAutocompleteElementResponse(pkiBrokerID int32, fkiDepartmentID int32, sBrokerName string, bBrokerIsactive bool, ) *BrokerAutocompleteElementResponse`

NewBrokerAutocompleteElementResponse instantiates a new BrokerAutocompleteElementResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrokerAutocompleteElementResponseWithDefaults

`func NewBrokerAutocompleteElementResponseWithDefaults() *BrokerAutocompleteElementResponse`

NewBrokerAutocompleteElementResponseWithDefaults instantiates a new BrokerAutocompleteElementResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiBrokerID

`func (o *BrokerAutocompleteElementResponse) GetPkiBrokerID() int32`

GetPkiBrokerID returns the PkiBrokerID field if non-nil, zero value otherwise.

### GetPkiBrokerIDOk

`func (o *BrokerAutocompleteElementResponse) GetPkiBrokerIDOk() (*int32, bool)`

GetPkiBrokerIDOk returns a tuple with the PkiBrokerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiBrokerID

`func (o *BrokerAutocompleteElementResponse) SetPkiBrokerID(v int32)`

SetPkiBrokerID sets PkiBrokerID field to given value.


### GetFkiDepartmentID

`func (o *BrokerAutocompleteElementResponse) GetFkiDepartmentID() int32`

GetFkiDepartmentID returns the FkiDepartmentID field if non-nil, zero value otherwise.

### GetFkiDepartmentIDOk

`func (o *BrokerAutocompleteElementResponse) GetFkiDepartmentIDOk() (*int32, bool)`

GetFkiDepartmentIDOk returns a tuple with the FkiDepartmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiDepartmentID

`func (o *BrokerAutocompleteElementResponse) SetFkiDepartmentID(v int32)`

SetFkiDepartmentID sets FkiDepartmentID field to given value.


### GetSBrokerName

`func (o *BrokerAutocompleteElementResponse) GetSBrokerName() string`

GetSBrokerName returns the SBrokerName field if non-nil, zero value otherwise.

### GetSBrokerNameOk

`func (o *BrokerAutocompleteElementResponse) GetSBrokerNameOk() (*string, bool)`

GetSBrokerNameOk returns a tuple with the SBrokerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSBrokerName

`func (o *BrokerAutocompleteElementResponse) SetSBrokerName(v string)`

SetSBrokerName sets SBrokerName field to given value.


### GetBBrokerIsactive

`func (o *BrokerAutocompleteElementResponse) GetBBrokerIsactive() bool`

GetBBrokerIsactive returns the BBrokerIsactive field if non-nil, zero value otherwise.

### GetBBrokerIsactiveOk

`func (o *BrokerAutocompleteElementResponse) GetBBrokerIsactiveOk() (*bool, bool)`

GetBBrokerIsactiveOk returns a tuple with the BBrokerIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBBrokerIsactive

`func (o *BrokerAutocompleteElementResponse) SetBBrokerIsactive(v bool)`

SetBBrokerIsactive sets BBrokerIsactive field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


