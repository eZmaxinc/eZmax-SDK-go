# CommonReportsubsection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AObjReportcolumn** | [**[]CommonReportcolumn**](CommonReportcolumn.md) |  | 
**IReportsubsectionColumncount** | **int32** | The number of Reportcolumns in the Reportsection | 
**IReportsubsectionWidth** | **int32** | The combined width of all the Reportcolumns in the Reportsection | 
**ObjReportsubsectionpartHeader** | [**CommonReportsubsectionpart**](CommonReportsubsectionpart.md) |  | 
**ObjReportsubsectionpartBody** | [**CommonReportsubsectionpart**](CommonReportsubsectionpart.md) |  | 
**ObjReportsubsectionpartFooter** | [**CommonReportsubsectionpart**](CommonReportsubsectionpart.md) |  | 
**SReportsubsectionTitle** | Pointer to **string** | The title of this Reportsubsection | [optional] 

## Methods

### NewCommonReportsubsection

`func NewCommonReportsubsection(aObjReportcolumn []CommonReportcolumn, iReportsubsectionColumncount int32, iReportsubsectionWidth int32, objReportsubsectionpartHeader CommonReportsubsectionpart, objReportsubsectionpartBody CommonReportsubsectionpart, objReportsubsectionpartFooter CommonReportsubsectionpart, ) *CommonReportsubsection`

NewCommonReportsubsection instantiates a new CommonReportsubsection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonReportsubsectionWithDefaults

`func NewCommonReportsubsectionWithDefaults() *CommonReportsubsection`

NewCommonReportsubsectionWithDefaults instantiates a new CommonReportsubsection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAObjReportcolumn

`func (o *CommonReportsubsection) GetAObjReportcolumn() []CommonReportcolumn`

GetAObjReportcolumn returns the AObjReportcolumn field if non-nil, zero value otherwise.

### GetAObjReportcolumnOk

`func (o *CommonReportsubsection) GetAObjReportcolumnOk() (*[]CommonReportcolumn, bool)`

GetAObjReportcolumnOk returns a tuple with the AObjReportcolumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjReportcolumn

`func (o *CommonReportsubsection) SetAObjReportcolumn(v []CommonReportcolumn)`

SetAObjReportcolumn sets AObjReportcolumn field to given value.


### GetIReportsubsectionColumncount

`func (o *CommonReportsubsection) GetIReportsubsectionColumncount() int32`

GetIReportsubsectionColumncount returns the IReportsubsectionColumncount field if non-nil, zero value otherwise.

### GetIReportsubsectionColumncountOk

`func (o *CommonReportsubsection) GetIReportsubsectionColumncountOk() (*int32, bool)`

GetIReportsubsectionColumncountOk returns a tuple with the IReportsubsectionColumncount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIReportsubsectionColumncount

`func (o *CommonReportsubsection) SetIReportsubsectionColumncount(v int32)`

SetIReportsubsectionColumncount sets IReportsubsectionColumncount field to given value.


### GetIReportsubsectionWidth

`func (o *CommonReportsubsection) GetIReportsubsectionWidth() int32`

GetIReportsubsectionWidth returns the IReportsubsectionWidth field if non-nil, zero value otherwise.

### GetIReportsubsectionWidthOk

`func (o *CommonReportsubsection) GetIReportsubsectionWidthOk() (*int32, bool)`

GetIReportsubsectionWidthOk returns a tuple with the IReportsubsectionWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIReportsubsectionWidth

`func (o *CommonReportsubsection) SetIReportsubsectionWidth(v int32)`

SetIReportsubsectionWidth sets IReportsubsectionWidth field to given value.


### GetObjReportsubsectionpartHeader

`func (o *CommonReportsubsection) GetObjReportsubsectionpartHeader() CommonReportsubsectionpart`

GetObjReportsubsectionpartHeader returns the ObjReportsubsectionpartHeader field if non-nil, zero value otherwise.

### GetObjReportsubsectionpartHeaderOk

`func (o *CommonReportsubsection) GetObjReportsubsectionpartHeaderOk() (*CommonReportsubsectionpart, bool)`

GetObjReportsubsectionpartHeaderOk returns a tuple with the ObjReportsubsectionpartHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjReportsubsectionpartHeader

`func (o *CommonReportsubsection) SetObjReportsubsectionpartHeader(v CommonReportsubsectionpart)`

SetObjReportsubsectionpartHeader sets ObjReportsubsectionpartHeader field to given value.


### GetObjReportsubsectionpartBody

`func (o *CommonReportsubsection) GetObjReportsubsectionpartBody() CommonReportsubsectionpart`

GetObjReportsubsectionpartBody returns the ObjReportsubsectionpartBody field if non-nil, zero value otherwise.

### GetObjReportsubsectionpartBodyOk

`func (o *CommonReportsubsection) GetObjReportsubsectionpartBodyOk() (*CommonReportsubsectionpart, bool)`

GetObjReportsubsectionpartBodyOk returns a tuple with the ObjReportsubsectionpartBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjReportsubsectionpartBody

`func (o *CommonReportsubsection) SetObjReportsubsectionpartBody(v CommonReportsubsectionpart)`

SetObjReportsubsectionpartBody sets ObjReportsubsectionpartBody field to given value.


### GetObjReportsubsectionpartFooter

`func (o *CommonReportsubsection) GetObjReportsubsectionpartFooter() CommonReportsubsectionpart`

GetObjReportsubsectionpartFooter returns the ObjReportsubsectionpartFooter field if non-nil, zero value otherwise.

### GetObjReportsubsectionpartFooterOk

`func (o *CommonReportsubsection) GetObjReportsubsectionpartFooterOk() (*CommonReportsubsectionpart, bool)`

GetObjReportsubsectionpartFooterOk returns a tuple with the ObjReportsubsectionpartFooter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjReportsubsectionpartFooter

`func (o *CommonReportsubsection) SetObjReportsubsectionpartFooter(v CommonReportsubsectionpart)`

SetObjReportsubsectionpartFooter sets ObjReportsubsectionpartFooter field to given value.


### GetSReportsubsectionTitle

`func (o *CommonReportsubsection) GetSReportsubsectionTitle() string`

GetSReportsubsectionTitle returns the SReportsubsectionTitle field if non-nil, zero value otherwise.

### GetSReportsubsectionTitleOk

`func (o *CommonReportsubsection) GetSReportsubsectionTitleOk() (*string, bool)`

GetSReportsubsectionTitleOk returns a tuple with the SReportsubsectionTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSReportsubsectionTitle

`func (o *CommonReportsubsection) SetSReportsubsectionTitle(v string)`

SetSReportsubsectionTitle sets SReportsubsectionTitle field to given value.

### HasSReportsubsectionTitle

`func (o *CommonReportsubsection) HasSReportsubsectionTitle() bool`

HasSReportsubsectionTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


