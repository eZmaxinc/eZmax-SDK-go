# CorsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiCorsID** | Pointer to **int32** | The unique ID of the Cors | [optional] 
**FkiApikeyID** | **int32** | The unique ID of the Apikey | 
**SCorsEntryurl** | **string** | The entryurl of the Cors | 

## Methods

### NewCorsRequest

`func NewCorsRequest(fkiApikeyID int32, sCorsEntryurl string, ) *CorsRequest`

NewCorsRequest instantiates a new CorsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorsRequestWithDefaults

`func NewCorsRequestWithDefaults() *CorsRequest`

NewCorsRequestWithDefaults instantiates a new CorsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiCorsID

`func (o *CorsRequest) GetPkiCorsID() int32`

GetPkiCorsID returns the PkiCorsID field if non-nil, zero value otherwise.

### GetPkiCorsIDOk

`func (o *CorsRequest) GetPkiCorsIDOk() (*int32, bool)`

GetPkiCorsIDOk returns a tuple with the PkiCorsID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiCorsID

`func (o *CorsRequest) SetPkiCorsID(v int32)`

SetPkiCorsID sets PkiCorsID field to given value.

### HasPkiCorsID

`func (o *CorsRequest) HasPkiCorsID() bool`

HasPkiCorsID returns a boolean if a field has been set.

### GetFkiApikeyID

`func (o *CorsRequest) GetFkiApikeyID() int32`

GetFkiApikeyID returns the FkiApikeyID field if non-nil, zero value otherwise.

### GetFkiApikeyIDOk

`func (o *CorsRequest) GetFkiApikeyIDOk() (*int32, bool)`

GetFkiApikeyIDOk returns a tuple with the FkiApikeyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiApikeyID

`func (o *CorsRequest) SetFkiApikeyID(v int32)`

SetFkiApikeyID sets FkiApikeyID field to given value.


### GetSCorsEntryurl

`func (o *CorsRequest) GetSCorsEntryurl() string`

GetSCorsEntryurl returns the SCorsEntryurl field if non-nil, zero value otherwise.

### GetSCorsEntryurlOk

`func (o *CorsRequest) GetSCorsEntryurlOk() (*string, bool)`

GetSCorsEntryurlOk returns a tuple with the SCorsEntryurl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCorsEntryurl

`func (o *CorsRequest) SetSCorsEntryurl(v string)`

SetSCorsEntryurl sets SCorsEntryurl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


