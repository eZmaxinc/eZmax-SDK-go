# CommonResponseObjDebug

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SMemoryUsage** | **string** | The peak memory allocated during the API request execution. Formatted as a human readable string | 
**SRunTime** | **string** | The total server execution time of the API request execution. Formatted as a human readable string | 
**ISQLSelects** | **int32** | The number of SQL SELECT queries that were sent to the database server during the API request execution | 
**ISQLQueries** | **int32** | The number of SQL INSERT/UPDATE/DELETE queries that were sent to the database server during the API request execution | 
**AObjSQLQuery** | [**[]CommonResponseObjSQLQuery**](CommonResponseObjSQLQuery.md) | An array of the SQL Queries that were executed during the API request execution | 

## Methods

### NewCommonResponseObjDebug

`func NewCommonResponseObjDebug(sMemoryUsage string, sRunTime string, iSQLSelects int32, iSQLQueries int32, aObjSQLQuery []CommonResponseObjSQLQuery, ) *CommonResponseObjDebug`

NewCommonResponseObjDebug instantiates a new CommonResponseObjDebug object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonResponseObjDebugWithDefaults

`func NewCommonResponseObjDebugWithDefaults() *CommonResponseObjDebug`

NewCommonResponseObjDebugWithDefaults instantiates a new CommonResponseObjDebug object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSMemoryUsage

`func (o *CommonResponseObjDebug) GetSMemoryUsage() string`

GetSMemoryUsage returns the SMemoryUsage field if non-nil, zero value otherwise.

### GetSMemoryUsageOk

`func (o *CommonResponseObjDebug) GetSMemoryUsageOk() (*string, bool)`

GetSMemoryUsageOk returns a tuple with the SMemoryUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSMemoryUsage

`func (o *CommonResponseObjDebug) SetSMemoryUsage(v string)`

SetSMemoryUsage sets SMemoryUsage field to given value.


### GetSRunTime

`func (o *CommonResponseObjDebug) GetSRunTime() string`

GetSRunTime returns the SRunTime field if non-nil, zero value otherwise.

### GetSRunTimeOk

`func (o *CommonResponseObjDebug) GetSRunTimeOk() (*string, bool)`

GetSRunTimeOk returns a tuple with the SRunTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSRunTime

`func (o *CommonResponseObjDebug) SetSRunTime(v string)`

SetSRunTime sets SRunTime field to given value.


### GetISQLSelects

`func (o *CommonResponseObjDebug) GetISQLSelects() int32`

GetISQLSelects returns the ISQLSelects field if non-nil, zero value otherwise.

### GetISQLSelectsOk

`func (o *CommonResponseObjDebug) GetISQLSelectsOk() (*int32, bool)`

GetISQLSelectsOk returns a tuple with the ISQLSelects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetISQLSelects

`func (o *CommonResponseObjDebug) SetISQLSelects(v int32)`

SetISQLSelects sets ISQLSelects field to given value.


### GetISQLQueries

`func (o *CommonResponseObjDebug) GetISQLQueries() int32`

GetISQLQueries returns the ISQLQueries field if non-nil, zero value otherwise.

### GetISQLQueriesOk

`func (o *CommonResponseObjDebug) GetISQLQueriesOk() (*int32, bool)`

GetISQLQueriesOk returns a tuple with the ISQLQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetISQLQueries

`func (o *CommonResponseObjDebug) SetISQLQueries(v int32)`

SetISQLQueries sets ISQLQueries field to given value.


### GetAObjSQLQuery

`func (o *CommonResponseObjDebug) GetAObjSQLQuery() []CommonResponseObjSQLQuery`

GetAObjSQLQuery returns the AObjSQLQuery field if non-nil, zero value otherwise.

### GetAObjSQLQueryOk

`func (o *CommonResponseObjDebug) GetAObjSQLQueryOk() (*[]CommonResponseObjSQLQuery, bool)`

GetAObjSQLQueryOk returns a tuple with the AObjSQLQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjSQLQuery

`func (o *CommonResponseObjDebug) SetAObjSQLQuery(v []CommonResponseObjSQLQuery)`

SetAObjSQLQuery sets AObjSQLQuery field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


