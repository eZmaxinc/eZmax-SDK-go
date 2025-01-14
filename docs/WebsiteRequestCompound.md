# WebsiteRequestCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiWebsiteID** | Pointer to **int32** | The unique ID of the Website Default | [optional] 
**FkiWebsitetypeID** | **int32** | The unique ID of the Websitetype.  Valid values:  |Value|Description| |-|-| |1|Website| |2|Twitter| |3|Facebook| |4|Survey| | 
**SWebsiteAddress** | **string** | The URL of the website. | 

## Methods

### NewWebsiteRequestCompound

`func NewWebsiteRequestCompound(fkiWebsitetypeID int32, sWebsiteAddress string, ) *WebsiteRequestCompound`

NewWebsiteRequestCompound instantiates a new WebsiteRequestCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebsiteRequestCompoundWithDefaults

`func NewWebsiteRequestCompoundWithDefaults() *WebsiteRequestCompound`

NewWebsiteRequestCompoundWithDefaults instantiates a new WebsiteRequestCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiWebsiteID

`func (o *WebsiteRequestCompound) GetPkiWebsiteID() int32`

GetPkiWebsiteID returns the PkiWebsiteID field if non-nil, zero value otherwise.

### GetPkiWebsiteIDOk

`func (o *WebsiteRequestCompound) GetPkiWebsiteIDOk() (*int32, bool)`

GetPkiWebsiteIDOk returns a tuple with the PkiWebsiteID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiWebsiteID

`func (o *WebsiteRequestCompound) SetPkiWebsiteID(v int32)`

SetPkiWebsiteID sets PkiWebsiteID field to given value.

### HasPkiWebsiteID

`func (o *WebsiteRequestCompound) HasPkiWebsiteID() bool`

HasPkiWebsiteID returns a boolean if a field has been set.

### GetFkiWebsitetypeID

`func (o *WebsiteRequestCompound) GetFkiWebsitetypeID() int32`

GetFkiWebsitetypeID returns the FkiWebsitetypeID field if non-nil, zero value otherwise.

### GetFkiWebsitetypeIDOk

`func (o *WebsiteRequestCompound) GetFkiWebsitetypeIDOk() (*int32, bool)`

GetFkiWebsitetypeIDOk returns a tuple with the FkiWebsitetypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiWebsitetypeID

`func (o *WebsiteRequestCompound) SetFkiWebsitetypeID(v int32)`

SetFkiWebsitetypeID sets FkiWebsitetypeID field to given value.


### GetSWebsiteAddress

`func (o *WebsiteRequestCompound) GetSWebsiteAddress() string`

GetSWebsiteAddress returns the SWebsiteAddress field if non-nil, zero value otherwise.

### GetSWebsiteAddressOk

`func (o *WebsiteRequestCompound) GetSWebsiteAddressOk() (*string, bool)`

GetSWebsiteAddressOk returns a tuple with the SWebsiteAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSWebsiteAddress

`func (o *WebsiteRequestCompound) SetSWebsiteAddress(v string)`

SetSWebsiteAddress sets SWebsiteAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


