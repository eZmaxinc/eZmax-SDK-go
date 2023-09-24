# CorsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiCorsID** | **int32** | The unique ID of the Cors | 
**FkiApikeyID** | **int32** | The unique ID of the Apikey | 
**SCorsEntryurl** | **string** | The entryurl of the Cors | 

## Methods

### NewCorsResponse

`func NewCorsResponse(pkiCorsID int32, fkiApikeyID int32, sCorsEntryurl string, ) *CorsResponse`

NewCorsResponse instantiates a new CorsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorsResponseWithDefaults

`func NewCorsResponseWithDefaults() *CorsResponse`

NewCorsResponseWithDefaults instantiates a new CorsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiCorsID

`func (o *CorsResponse) GetPkiCorsID() int32`

GetPkiCorsID returns the PkiCorsID field if non-nil, zero value otherwise.

### GetPkiCorsIDOk

`func (o *CorsResponse) GetPkiCorsIDOk() (*int32, bool)`

GetPkiCorsIDOk returns a tuple with the PkiCorsID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiCorsID

`func (o *CorsResponse) SetPkiCorsID(v int32)`

SetPkiCorsID sets PkiCorsID field to given value.


### GetFkiApikeyID

`func (o *CorsResponse) GetFkiApikeyID() int32`

GetFkiApikeyID returns the FkiApikeyID field if non-nil, zero value otherwise.

### GetFkiApikeyIDOk

`func (o *CorsResponse) GetFkiApikeyIDOk() (*int32, bool)`

GetFkiApikeyIDOk returns a tuple with the FkiApikeyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiApikeyID

`func (o *CorsResponse) SetFkiApikeyID(v int32)`

SetFkiApikeyID sets FkiApikeyID field to given value.


### GetSCorsEntryurl

`func (o *CorsResponse) GetSCorsEntryurl() string`

GetSCorsEntryurl returns the SCorsEntryurl field if non-nil, zero value otherwise.

### GetSCorsEntryurlOk

`func (o *CorsResponse) GetSCorsEntryurlOk() (*string, bool)`

GetSCorsEntryurlOk returns a tuple with the SCorsEntryurl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCorsEntryurl

`func (o *CorsResponse) SetSCorsEntryurl(v string)`

SetSCorsEntryurl sets SCorsEntryurl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


