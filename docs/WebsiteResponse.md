# WebsiteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiWebsiteID** | **int32** | The unique ID of the Website Default | 
**FkiWebsitetypeID** | **int32** | The unique ID of the Websitetype.  Valid values:  |Value|Description| |-|-| |1|Website| |2|Twitter| |3|Facebook| |4|Survey| | 
**SWebsiteAddress** | **string** | The URL of the website. | 

## Methods

### NewWebsiteResponse

`func NewWebsiteResponse(pkiWebsiteID int32, fkiWebsitetypeID int32, sWebsiteAddress string, ) *WebsiteResponse`

NewWebsiteResponse instantiates a new WebsiteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebsiteResponseWithDefaults

`func NewWebsiteResponseWithDefaults() *WebsiteResponse`

NewWebsiteResponseWithDefaults instantiates a new WebsiteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiWebsiteID

`func (o *WebsiteResponse) GetPkiWebsiteID() int32`

GetPkiWebsiteID returns the PkiWebsiteID field if non-nil, zero value otherwise.

### GetPkiWebsiteIDOk

`func (o *WebsiteResponse) GetPkiWebsiteIDOk() (*int32, bool)`

GetPkiWebsiteIDOk returns a tuple with the PkiWebsiteID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiWebsiteID

`func (o *WebsiteResponse) SetPkiWebsiteID(v int32)`

SetPkiWebsiteID sets PkiWebsiteID field to given value.


### GetFkiWebsitetypeID

`func (o *WebsiteResponse) GetFkiWebsitetypeID() int32`

GetFkiWebsitetypeID returns the FkiWebsitetypeID field if non-nil, zero value otherwise.

### GetFkiWebsitetypeIDOk

`func (o *WebsiteResponse) GetFkiWebsitetypeIDOk() (*int32, bool)`

GetFkiWebsitetypeIDOk returns a tuple with the FkiWebsitetypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiWebsitetypeID

`func (o *WebsiteResponse) SetFkiWebsitetypeID(v int32)`

SetFkiWebsitetypeID sets FkiWebsitetypeID field to given value.


### GetSWebsiteAddress

`func (o *WebsiteResponse) GetSWebsiteAddress() string`

GetSWebsiteAddress returns the SWebsiteAddress field if non-nil, zero value otherwise.

### GetSWebsiteAddressOk

`func (o *WebsiteResponse) GetSWebsiteAddressOk() (*string, bool)`

GetSWebsiteAddressOk returns a tuple with the SWebsiteAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSWebsiteAddress

`func (o *WebsiteResponse) SetSWebsiteAddress(v string)`

SetSWebsiteAddress sets SWebsiteAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


