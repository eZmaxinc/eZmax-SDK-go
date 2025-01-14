# WebsiteResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiWebsiteID** | **int32** | The unique ID of the Website Default | 
**FkiWebsitetypeID** | **int32** | The unique ID of the Websitetype.  Valid values:  |Value|Description| |-|-| |1|Website| |2|Twitter| |3|Facebook| |4|Survey| | 
**SWebsiteAddress** | **string** | The URL of the website. | 

## Methods

### NewWebsiteResponseCompound

`func NewWebsiteResponseCompound(pkiWebsiteID int32, fkiWebsitetypeID int32, sWebsiteAddress string, ) *WebsiteResponseCompound`

NewWebsiteResponseCompound instantiates a new WebsiteResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebsiteResponseCompoundWithDefaults

`func NewWebsiteResponseCompoundWithDefaults() *WebsiteResponseCompound`

NewWebsiteResponseCompoundWithDefaults instantiates a new WebsiteResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiWebsiteID

`func (o *WebsiteResponseCompound) GetPkiWebsiteID() int32`

GetPkiWebsiteID returns the PkiWebsiteID field if non-nil, zero value otherwise.

### GetPkiWebsiteIDOk

`func (o *WebsiteResponseCompound) GetPkiWebsiteIDOk() (*int32, bool)`

GetPkiWebsiteIDOk returns a tuple with the PkiWebsiteID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiWebsiteID

`func (o *WebsiteResponseCompound) SetPkiWebsiteID(v int32)`

SetPkiWebsiteID sets PkiWebsiteID field to given value.


### GetFkiWebsitetypeID

`func (o *WebsiteResponseCompound) GetFkiWebsitetypeID() int32`

GetFkiWebsitetypeID returns the FkiWebsitetypeID field if non-nil, zero value otherwise.

### GetFkiWebsitetypeIDOk

`func (o *WebsiteResponseCompound) GetFkiWebsitetypeIDOk() (*int32, bool)`

GetFkiWebsitetypeIDOk returns a tuple with the FkiWebsitetypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiWebsitetypeID

`func (o *WebsiteResponseCompound) SetFkiWebsitetypeID(v int32)`

SetFkiWebsitetypeID sets FkiWebsitetypeID field to given value.


### GetSWebsiteAddress

`func (o *WebsiteResponseCompound) GetSWebsiteAddress() string`

GetSWebsiteAddress returns the SWebsiteAddress field if non-nil, zero value otherwise.

### GetSWebsiteAddressOk

`func (o *WebsiteResponseCompound) GetSWebsiteAddressOk() (*string, bool)`

GetSWebsiteAddressOk returns a tuple with the SWebsiteAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSWebsiteAddress

`func (o *WebsiteResponseCompound) SetSWebsiteAddress(v string)`

SetSWebsiteAddress sets SWebsiteAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


