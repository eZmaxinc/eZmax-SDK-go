# WebsiteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiWebsiteID** | Pointer to **int32** | The unique ID of the Website Default | [optional] 
**FkiWebsitetypeID** | **int32** | The unique ID of the Websitetype.  Valid values:  |Value|Description| |-|-| |1|Website| |2|Twitter| |3|Facebook| |4|Survey| | 
**SWebsiteAddress** | **string** | The URL of the website. | 

## Methods

### NewWebsiteRequest

`func NewWebsiteRequest(fkiWebsitetypeID int32, sWebsiteAddress string, ) *WebsiteRequest`

NewWebsiteRequest instantiates a new WebsiteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebsiteRequestWithDefaults

`func NewWebsiteRequestWithDefaults() *WebsiteRequest`

NewWebsiteRequestWithDefaults instantiates a new WebsiteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiWebsiteID

`func (o *WebsiteRequest) GetPkiWebsiteID() int32`

GetPkiWebsiteID returns the PkiWebsiteID field if non-nil, zero value otherwise.

### GetPkiWebsiteIDOk

`func (o *WebsiteRequest) GetPkiWebsiteIDOk() (*int32, bool)`

GetPkiWebsiteIDOk returns a tuple with the PkiWebsiteID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiWebsiteID

`func (o *WebsiteRequest) SetPkiWebsiteID(v int32)`

SetPkiWebsiteID sets PkiWebsiteID field to given value.

### HasPkiWebsiteID

`func (o *WebsiteRequest) HasPkiWebsiteID() bool`

HasPkiWebsiteID returns a boolean if a field has been set.

### GetFkiWebsitetypeID

`func (o *WebsiteRequest) GetFkiWebsitetypeID() int32`

GetFkiWebsitetypeID returns the FkiWebsitetypeID field if non-nil, zero value otherwise.

### GetFkiWebsitetypeIDOk

`func (o *WebsiteRequest) GetFkiWebsitetypeIDOk() (*int32, bool)`

GetFkiWebsitetypeIDOk returns a tuple with the FkiWebsitetypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiWebsitetypeID

`func (o *WebsiteRequest) SetFkiWebsitetypeID(v int32)`

SetFkiWebsitetypeID sets FkiWebsitetypeID field to given value.


### GetSWebsiteAddress

`func (o *WebsiteRequest) GetSWebsiteAddress() string`

GetSWebsiteAddress returns the SWebsiteAddress field if non-nil, zero value otherwise.

### GetSWebsiteAddressOk

`func (o *WebsiteRequest) GetSWebsiteAddressOk() (*string, bool)`

GetSWebsiteAddressOk returns a tuple with the SWebsiteAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSWebsiteAddress

`func (o *WebsiteRequest) SetSWebsiteAddress(v string)`

SetSWebsiteAddress sets SWebsiteAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


