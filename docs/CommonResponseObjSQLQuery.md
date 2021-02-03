# CommonResponseObjSQLQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SQuery** | **string** | The SQL Query | 
**FDuration** | **float32** | Execution time of the SQL Query in seconds | 

## Methods

### NewCommonResponseObjSQLQuery

`func NewCommonResponseObjSQLQuery(sQuery string, fDuration float32, ) *CommonResponseObjSQLQuery`

NewCommonResponseObjSQLQuery instantiates a new CommonResponseObjSQLQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonResponseObjSQLQueryWithDefaults

`func NewCommonResponseObjSQLQueryWithDefaults() *CommonResponseObjSQLQuery`

NewCommonResponseObjSQLQueryWithDefaults instantiates a new CommonResponseObjSQLQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSQuery

`func (o *CommonResponseObjSQLQuery) GetSQuery() string`

GetSQuery returns the SQuery field if non-nil, zero value otherwise.

### GetSQueryOk

`func (o *CommonResponseObjSQLQuery) GetSQueryOk() (*string, bool)`

GetSQueryOk returns a tuple with the SQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSQuery

`func (o *CommonResponseObjSQLQuery) SetSQuery(v string)`

SetSQuery sets SQuery field to given value.


### GetFDuration

`func (o *CommonResponseObjSQLQuery) GetFDuration() float32`

GetFDuration returns the FDuration field if non-nil, zero value otherwise.

### GetFDurationOk

`func (o *CommonResponseObjSQLQuery) GetFDurationOk() (*float32, bool)`

GetFDurationOk returns a tuple with the FDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFDuration

`func (o *CommonResponseObjSQLQuery) SetFDuration(v float32)`

SetFDuration sets FDuration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


