# Go API client for eZmaxinc/eZmax-SDK-go

This API expose all the functionnalities for the eZmax and eZsign application.

We provide SDKs for customers. They are generated using OpenAPI codegen, we encourage customers to use them as we also provide samples for them.

You can choose to build your own implementation manually or can use any compatible OpenAPI 3.0 generator like Swagger Codegen, OpenAPI codegen or any commercial generators.

If you need helping understanding how to use this API, don't waste too much time looking for it. Contact support-api@ezmax.ca, we're here to help. We are developpers so we know programmers don't like bad documentation. If you don't find what you need in the documentation, let us know, we'll improve it and put you rapidly up on track.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.28
- Package version: 1.0.28
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://www.ezmax.ca/en/contact](https://www.ezmax.ca/en/contact)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sw "./eZmaxinc/eZmax-SDK-go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identifield by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ModuleUserApi* | [**UserCreateEzsignuserV1**](docs/ModuleUserApi.md#usercreateezsignuserv1) | **Post** /1/module/user/createezsignuser | Create a new User of type Ezsignuser
*ObjectActivesessionApi* | [**ActivesessionGetCurrentV1**](docs/ObjectActivesessionApi.md#activesessiongetcurrentv1) | **Get** /1/object/activesession/getCurrent | Get Current Activesession
*ObjectApikeyApi* | [**ApikeyCreateObjectV1**](docs/ObjectApikeyApi.md#apikeycreateobjectv1) | **Post** /1/object/apikey | Create a new Apikey
*ObjectEzsigndocumentApi* | [**EzsigndocumentApplyEzsigntemplateV1**](docs/ObjectEzsigndocumentApi.md#ezsigndocumentapplyezsigntemplatev1) | **Post** /1/object/ezsigndocument/{pkiEzsigndocumentID}/applyezsigntemplate | Apply an Ezsign Template to the Ezsigndocument.
*ObjectEzsigndocumentApi* | [**EzsigndocumentCreateObjectV1**](docs/ObjectEzsigndocumentApi.md#ezsigndocumentcreateobjectv1) | **Post** /1/object/ezsigndocument | Create a new Ezsigndocument
*ObjectEzsigndocumentApi* | [**EzsigndocumentDeleteObjectV1**](docs/ObjectEzsigndocumentApi.md#ezsigndocumentdeleteobjectv1) | **Delete** /1/object/ezsigndocument/{pkiEzsigndocumentID} | Delete an existing Ezsigndocument
*ObjectEzsigndocumentApi* | [**EzsigndocumentEditObjectV1**](docs/ObjectEzsigndocumentApi.md#ezsigndocumenteditobjectv1) | **Put** /1/object/ezsigndocument/{pkiEzsigndocumentID} | Modify an existing Ezsigndocument
*ObjectEzsigndocumentApi* | [**EzsigndocumentGetDownloadUrlV1**](docs/ObjectEzsigndocumentApi.md#ezsigndocumentgetdownloadurlv1) | **Get** /1/object/ezsigndocument/{pkiEzsigndocumentID}/getDownloadUrl/{eDocumentType} | Retrieve a URL to download documents.
*ObjectEzsigndocumentApi* | [**EzsigndocumentGetObjectGetChildrenV1**](docs/ObjectEzsigndocumentApi.md#ezsigndocumentgetobjectgetchildrenv1) | **Get** /1/object/ezsigndocument/{pkiEzsigndocumentID}/getChildren | Retrieve an existing Ezsigndocument&#39;s children IDs
*ObjectEzsigndocumentApi* | [**EzsigndocumentGetObjectV1**](docs/ObjectEzsigndocumentApi.md#ezsigndocumentgetobjectv1) | **Get** /1/object/ezsigndocument/{pkiEzsigndocumentID} | Retrieve an existing Ezsigndocument
*ObjectEzsignfolderApi* | [**EzsignfolderCreateObjectV1**](docs/ObjectEzsignfolderApi.md#ezsignfoldercreateobjectv1) | **Post** /1/object/ezsignfolder | Create a new Ezsignfolder
*ObjectEzsignfolderApi* | [**EzsignfolderDeleteObjectV1**](docs/ObjectEzsignfolderApi.md#ezsignfolderdeleteobjectv1) | **Delete** /1/object/ezsignfolder/{pkiEzsignfolderID} | Delete an existing Ezsignfolder
*ObjectEzsignfolderApi* | [**EzsignfolderEditObjectV1**](docs/ObjectEzsignfolderApi.md#ezsignfoldereditobjectv1) | **Put** /1/object/ezsignfolder/{pkiEzsignfolderID} | Modify an existing Ezsignfolder
*ObjectEzsignfolderApi* | [**EzsignfolderGetObjectGetChildrenV1**](docs/ObjectEzsignfolderApi.md#ezsignfoldergetobjectgetchildrenv1) | **Get** /1/object/ezsignfolder/{pkiEzsignfolderID}/getChildren | Retrieve an existing Ezsignfolder&#39;s children IDs
*ObjectEzsignfolderApi* | [**EzsignfolderGetObjectV1**](docs/ObjectEzsignfolderApi.md#ezsignfoldergetobjectv1) | **Get** /1/object/ezsignfolder/{pkiEzsignfolderID} | Retrieve an existing Ezsignfolder
*ObjectEzsignfolderApi* | [**EzsignfolderSendV1**](docs/ObjectEzsignfolderApi.md#ezsignfoldersendv1) | **Post** /1/object/ezsignfolder/{pkiEzsignfolderID}/send | Send the Ezsignfolder to the signatories for signature
*ObjectEzsignfoldersignerassociationApi* | [**EzsignfoldersignerassociationCreateObjectV1**](docs/ObjectEzsignfoldersignerassociationApi.md#ezsignfoldersignerassociationcreateobjectv1) | **Post** /1/object/ezsignfoldersignerassociation | Create a new Ezsignfoldersignerassociation
*ObjectEzsignfoldersignerassociationApi* | [**EzsignfoldersignerassociationDeleteObjectV1**](docs/ObjectEzsignfoldersignerassociationApi.md#ezsignfoldersignerassociationdeleteobjectv1) | **Delete** /1/object/ezsignfoldersignerassociation/{pkiEzsignfoldersignerassociationID} | Delete an existing Ezsignfoldersignerassociation
*ObjectEzsignfoldersignerassociationApi* | [**EzsignfoldersignerassociationEditObjectV1**](docs/ObjectEzsignfoldersignerassociationApi.md#ezsignfoldersignerassociationeditobjectv1) | **Put** /1/object/ezsignfoldersignerassociation/{pkiEzsignfoldersignerassociationID} | Modify an existing Ezsignfoldersignerassociation
*ObjectEzsignfoldersignerassociationApi* | [**EzsignfoldersignerassociationGetInPersonLoginUrlV1**](docs/ObjectEzsignfoldersignerassociationApi.md#ezsignfoldersignerassociationgetinpersonloginurlv1) | **Get** /1/object/ezsignfoldersignerassociation/{pkiEzsignfoldersignerassociationID}/getInPersonLoginUrl | Retrieve a Login Url to allow In-Person signing
*ObjectEzsignfoldersignerassociationApi* | [**EzsignfoldersignerassociationGetObjectGetChildrenV1**](docs/ObjectEzsignfoldersignerassociationApi.md#ezsignfoldersignerassociationgetobjectgetchildrenv1) | **Get** /1/object/ezsignfoldersignerassociation/{pkiEzsignfoldersignerassociationID}/getChildren | Retrieve an existing Ezsignfoldersignerassociation&#39;s children IDs
*ObjectEzsignfoldersignerassociationApi* | [**EzsignfoldersignerassociationGetObjectV1**](docs/ObjectEzsignfoldersignerassociationApi.md#ezsignfoldersignerassociationgetobjectv1) | **Get** /1/object/ezsignfoldersignerassociation/{pkiEzsignfoldersignerassociationID} | Retrieve an existing Ezsignfoldersignerassociation
*ObjectEzsignsignatureApi* | [**EzsignsignatureCreateObjectV1**](docs/ObjectEzsignsignatureApi.md#ezsignsignaturecreateobjectv1) | **Post** /1/object/ezsignsignature | Create a new Ezsignsignature
*ObjectEzsignsignatureApi* | [**EzsignsignatureDeleteObjectV1**](docs/ObjectEzsignsignatureApi.md#ezsignsignaturedeleteobjectv1) | **Delete** /1/object/ezsignsignature/{pkiEzsignsignatureID} | Delete an existing Ezsignsignature
*ObjectEzsignsignatureApi* | [**EzsignsignatureEditObjectV1**](docs/ObjectEzsignsignatureApi.md#ezsignsignatureeditobjectv1) | **Put** /1/object/ezsignsignature/{pkiEzsignsignatureID} | Modify an existing Ezsignsignature
*ObjectEzsignsignatureApi* | [**EzsignsignatureGetObjectGetChildrenV1**](docs/ObjectEzsignsignatureApi.md#ezsignsignaturegetobjectgetchildrenv1) | **Get** /1/object/ezsignsignature/{pkiEzsignsignatureID}/getChildren | Retrieve an existing Ezsignsignature&#39;s children IDs
*ObjectEzsignsignatureApi* | [**EzsignsignatureGetObjectV1**](docs/ObjectEzsignsignatureApi.md#ezsignsignaturegetobjectv1) | **Get** /1/object/ezsignsignature/{pkiEzsignsignatureID} | Retrieve an existing Ezsignsignature
*ObjectFranchisebrokerApi* | [**FranchisebrokerGetAutocompleteV1**](docs/ObjectFranchisebrokerApi.md#franchisebrokergetautocompletev1) | **Get** /1/object/franchisebroker/getAutocomplete/{sSelector} | Retrieve Franchisebrokers and IDs
*ObjectFranchiseofficeApi* | [**FranchiseofficeGetAutocompleteV1**](docs/ObjectFranchiseofficeApi.md#franchiseofficegetautocompletev1) | **Get** /1/object/franchiseoffice/getAutocomplete/{sSelector} | Retrieve Franchiseoffices and IDs
*ObjectFranchisereferalincomeApi* | [**FranchisereferalincomeCreateObjectV1**](docs/ObjectFranchisereferalincomeApi.md#franchisereferalincomecreateobjectv1) | **Post** /1/object/franchisereferalincome | Create a new Franchisereferalincome
*ObjectPeriodApi* | [**PeriodGetAutocompleteV1**](docs/ObjectPeriodApi.md#periodgetautocompletev1) | **Get** /1/object/period/getAutocomplete/{sSelector} | Retrieve Periods and IDs


## Documentation For Models

 - [ActivesessionGetCurrentV1Response](docs/ActivesessionGetCurrentV1Response.md)
 - [ActivesessionGetCurrentV1ResponseMPayload](docs/ActivesessionGetCurrentV1ResponseMPayload.md)
 - [AddressRequest](docs/AddressRequest.md)
 - [ApikeyCreateObjectV1Request](docs/ApikeyCreateObjectV1Request.md)
 - [ApikeyCreateObjectV1Response](docs/ApikeyCreateObjectV1Response.md)
 - [ApikeyCreateObjectV1ResponseMPayload](docs/ApikeyCreateObjectV1ResponseMPayload.md)
 - [ApikeyRequest](docs/ApikeyRequest.md)
 - [ApikeyRequestCompound](docs/ApikeyRequestCompound.md)
 - [ApikeyResponse](docs/ApikeyResponse.md)
 - [AttemptResponse](docs/AttemptResponse.md)
 - [CommonAudit](docs/CommonAudit.md)
 - [CommonGetAutocompleteV1Response](docs/CommonGetAutocompleteV1Response.md)
 - [CommonGetAutocompleteV1ResponseMPayload](docs/CommonGetAutocompleteV1ResponseMPayload.md)
 - [CommonResponse](docs/CommonResponse.md)
 - [CommonResponseError](docs/CommonResponseError.md)
 - [CommonResponseObjDebug](docs/CommonResponseObjDebug.md)
 - [CommonResponseObjDebugPayload](docs/CommonResponseObjDebugPayload.md)
 - [CommonResponseObjSQLQuery](docs/CommonResponseObjSQLQuery.md)
 - [CommonWebhook](docs/CommonWebhook.md)
 - [ContactRequest](docs/ContactRequest.md)
 - [ContactRequestCompound](docs/ContactRequestCompound.md)
 - [ContactRequestCompoundAllOf](docs/ContactRequestCompoundAllOf.md)
 - [ContactinformationsRequest](docs/ContactinformationsRequest.md)
 - [ContactinformationsRequestCompound](docs/ContactinformationsRequestCompound.md)
 - [EmailRequest](docs/EmailRequest.md)
 - [EzsigndocumentApplyEzsigntemplateV1Request](docs/EzsigndocumentApplyEzsigntemplateV1Request.md)
 - [EzsigndocumentApplyEzsigntemplateV1Response](docs/EzsigndocumentApplyEzsigntemplateV1Response.md)
 - [EzsigndocumentCreateObjectV1Request](docs/EzsigndocumentCreateObjectV1Request.md)
 - [EzsigndocumentCreateObjectV1Response](docs/EzsigndocumentCreateObjectV1Response.md)
 - [EzsigndocumentCreateObjectV1ResponseMPayload](docs/EzsigndocumentCreateObjectV1ResponseMPayload.md)
 - [EzsigndocumentDeleteObjectV1Response](docs/EzsigndocumentDeleteObjectV1Response.md)
 - [EzsigndocumentEditObjectV1Request](docs/EzsigndocumentEditObjectV1Request.md)
 - [EzsigndocumentEditObjectV1Response](docs/EzsigndocumentEditObjectV1Response.md)
 - [EzsigndocumentGetDownloadUrlV1Response](docs/EzsigndocumentGetDownloadUrlV1Response.md)
 - [EzsigndocumentGetDownloadUrlV1ResponseMPayload](docs/EzsigndocumentGetDownloadUrlV1ResponseMPayload.md)
 - [EzsigndocumentGetObjectV1Response](docs/EzsigndocumentGetObjectV1Response.md)
 - [EzsigndocumentGetObjectV1ResponseMPayload](docs/EzsigndocumentGetObjectV1ResponseMPayload.md)
 - [EzsigndocumentRequest](docs/EzsigndocumentRequest.md)
 - [EzsigndocumentRequestCompound](docs/EzsigndocumentRequestCompound.md)
 - [EzsigndocumentResponse](docs/EzsigndocumentResponse.md)
 - [EzsigndocumentResponseAllOf](docs/EzsigndocumentResponseAllOf.md)
 - [EzsigndocumentResponseCompound](docs/EzsigndocumentResponseCompound.md)
 - [EzsignfolderCreateObjectV1Request](docs/EzsignfolderCreateObjectV1Request.md)
 - [EzsignfolderCreateObjectV1Response](docs/EzsignfolderCreateObjectV1Response.md)
 - [EzsignfolderCreateObjectV1ResponseMPayload](docs/EzsignfolderCreateObjectV1ResponseMPayload.md)
 - [EzsignfolderDeleteObjectV1Response](docs/EzsignfolderDeleteObjectV1Response.md)
 - [EzsignfolderEditObjectV1Request](docs/EzsignfolderEditObjectV1Request.md)
 - [EzsignfolderEditObjectV1Response](docs/EzsignfolderEditObjectV1Response.md)
 - [EzsignfolderGetObjectV1Response](docs/EzsignfolderGetObjectV1Response.md)
 - [EzsignfolderGetObjectV1ResponseMPayload](docs/EzsignfolderGetObjectV1ResponseMPayload.md)
 - [EzsignfolderRequest](docs/EzsignfolderRequest.md)
 - [EzsignfolderRequestCompound](docs/EzsignfolderRequestCompound.md)
 - [EzsignfolderResponse](docs/EzsignfolderResponse.md)
 - [EzsignfolderResponseAllOf](docs/EzsignfolderResponseAllOf.md)
 - [EzsignfolderResponseCompound](docs/EzsignfolderResponseCompound.md)
 - [EzsignfolderSendV1Request](docs/EzsignfolderSendV1Request.md)
 - [EzsignfolderSendV1Response](docs/EzsignfolderSendV1Response.md)
 - [EzsignfoldersignerassociationCreateObjectV1Request](docs/EzsignfoldersignerassociationCreateObjectV1Request.md)
 - [EzsignfoldersignerassociationCreateObjectV1Response](docs/EzsignfoldersignerassociationCreateObjectV1Response.md)
 - [EzsignfoldersignerassociationCreateObjectV1ResponseMPayload](docs/EzsignfoldersignerassociationCreateObjectV1ResponseMPayload.md)
 - [EzsignfoldersignerassociationDeleteObjectV1Response](docs/EzsignfoldersignerassociationDeleteObjectV1Response.md)
 - [EzsignfoldersignerassociationEditObjectV1Request](docs/EzsignfoldersignerassociationEditObjectV1Request.md)
 - [EzsignfoldersignerassociationEditObjectV1Response](docs/EzsignfoldersignerassociationEditObjectV1Response.md)
 - [EzsignfoldersignerassociationGetInPersonLoginUrlV1Response](docs/EzsignfoldersignerassociationGetInPersonLoginUrlV1Response.md)
 - [EzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseMPayload](docs/EzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseMPayload.md)
 - [EzsignfoldersignerassociationGetObjectV1Response](docs/EzsignfoldersignerassociationGetObjectV1Response.md)
 - [EzsignfoldersignerassociationRequest](docs/EzsignfoldersignerassociationRequest.md)
 - [EzsignfoldersignerassociationRequestCompound](docs/EzsignfoldersignerassociationRequestCompound.md)
 - [EzsignsignatureCreateObjectV1Request](docs/EzsignsignatureCreateObjectV1Request.md)
 - [EzsignsignatureCreateObjectV1Response](docs/EzsignsignatureCreateObjectV1Response.md)
 - [EzsignsignatureCreateObjectV1ResponseMPayload](docs/EzsignsignatureCreateObjectV1ResponseMPayload.md)
 - [EzsignsignatureDeleteObjectV1Response](docs/EzsignsignatureDeleteObjectV1Response.md)
 - [EzsignsignatureEditObjectV1Request](docs/EzsignsignatureEditObjectV1Request.md)
 - [EzsignsignatureEditObjectV1Response](docs/EzsignsignatureEditObjectV1Response.md)
 - [EzsignsignatureGetObjectV1Response](docs/EzsignsignatureGetObjectV1Response.md)
 - [EzsignsignatureRequest](docs/EzsignsignatureRequest.md)
 - [EzsignsignatureRequestCompound](docs/EzsignsignatureRequestCompound.md)
 - [EzsignsignerRequest](docs/EzsignsignerRequest.md)
 - [EzsignsignerRequestCompound](docs/EzsignsignerRequestCompound.md)
 - [EzsignsignerRequestCompoundContact](docs/EzsignsignerRequestCompoundContact.md)
 - [FieldEEzsigndocumentStep](docs/FieldEEzsigndocumentStep.md)
 - [FieldEEzsignfolderSendreminderfrequency](docs/FieldEEzsignfolderSendreminderfrequency.md)
 - [FieldEEzsignfolderStep](docs/FieldEEzsignfolderStep.md)
 - [FieldEPhoneType](docs/FieldEPhoneType.md)
 - [FieldEUserType](docs/FieldEUserType.md)
 - [FranchisereferalincomeCreateObjectV1Request](docs/FranchisereferalincomeCreateObjectV1Request.md)
 - [FranchisereferalincomeCreateObjectV1Response](docs/FranchisereferalincomeCreateObjectV1Response.md)
 - [FranchisereferalincomeCreateObjectV1ResponseMPayload](docs/FranchisereferalincomeCreateObjectV1ResponseMPayload.md)
 - [FranchisereferalincomeRequest](docs/FranchisereferalincomeRequest.md)
 - [FranchisereferalincomeRequestCompound](docs/FranchisereferalincomeRequestCompound.md)
 - [MultilingualApikeyDescription](docs/MultilingualApikeyDescription.md)
 - [PhoneRequest](docs/PhoneRequest.md)
 - [UserCreateEzsignuserV1Request](docs/UserCreateEzsignuserV1Request.md)
 - [UserCreateEzsignuserV1Response](docs/UserCreateEzsignuserV1Response.md)
 - [UserCreateEzsignuserV1ResponseMPayload](docs/UserCreateEzsignuserV1ResponseMPayload.md)
 - [UserResponse](docs/UserResponse.md)
 - [UserResponseAllOf](docs/UserResponseAllOf.md)
 - [WebhookEzsignDocumentCompleted](docs/WebhookEzsignDocumentCompleted.md)
 - [WebhookEzsignDocumentCompletedAllOf](docs/WebhookEzsignDocumentCompletedAllOf.md)
 - [WebhookEzsignFolderCompleted](docs/WebhookEzsignFolderCompleted.md)
 - [WebhookEzsignFolderCompletedAllOf](docs/WebhookEzsignFolderCompletedAllOf.md)
 - [WebhookResponse](docs/WebhookResponse.md)
 - [WebhookUserUserCreated](docs/WebhookUserUserCreated.md)
 - [WebhookUserUserCreatedAllOf](docs/WebhookUserUserCreatedAllOf.md)
 - [WebsiteRequest](docs/WebsiteRequest.md)


## Documentation For Authorization



### Authorization

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Authorization and passed in as the auth context for each request.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

support-api@ezmax.ca

