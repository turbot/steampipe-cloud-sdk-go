# Go SDK for Steampipe Cloud

Steampipe Cloud is a hosted version of Steampipe (https://steampipe.io), an open source tool to instantly query your cloud services (e.g. AWS, Azure, GCP and more) with SQL. No DB required.

For help on getting started with Steampipe Cloud, please visit https://steampipe.io/docs/cloud/getting-started.

## Getting Started

Here's an example of listing the workspaces for your user:

```go
package main

import (
    "context"
    "fmt"
    "os"

    steampipecloud "github.com/turbot/steampipe-cloud-sdk-go"
)

func main() {
    // Create a default configuration
    configuration := steampipecloud.NewConfiguration()

    // Add your Steampipe Cloud user token as an auth header
    configuration.AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("STEAMPIPE_CLOUD_TOKEN")))

    // Create a client
    client := steampipecloud.NewAPIClient(configuration)

    // Find your authenticated user info
    actor, _, err := client.Actors.Get(context.Background()).Execute()

    if err != nil {
      // Do something with the error
      return
    }

    // List your workspaces
    workspaces, _, err := client.UserWorkspaces.List(context.Background(), actor.Handle).Execute()

    if err != nil {
      // Do something with the error
      return
    }
}
```

## Usages

For more detailed examples of using the SDK, please check out the following open source projects:

- https://github.com/turbot/steampipe-plugin-steampipecloud
  - The official Steampipe plugin for accessing your Steampipe Cloud resources via SQL.
- https://github.com/turbot/terraform-provider-steampipecloud
  - The Terraform provider for managing your Steampipe Cloud resources using Terraform.

## Documentation for API Endpoints

All URIs are relative to *https://cloud.steampipe.io/api/v0*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*Actors* | [**Get**](docs/Actors.md#get) | **Get** /actor | Actor information
*Actors* | [**ListActivity**](docs/Actors.md#listactivity) | **Get** /actor/activity | List actor activity
*Actors* | [**ListConnections**](docs/Actors.md#listconnections) | **Get** /actor/conn | List actor connections
*Actors* | [**ListOrgInvites**](docs/Actors.md#listorginvites) | **Get** /actor/org/invite | List org invites for actor
*Actors* | [**ListOrgs**](docs/Actors.md#listorgs) | **Get** /actor/org | List actor orgs
*Actors* | [**ListWorkspaces**](docs/Actors.md#listworkspaces) | **Get** /actor/workspace | List actor workspaces
*Auth* | [**ConfirmLogin**](docs/Auth.md#confirmlogin) | **Get** /login/confirm | Confirm user login
*Auth* | [**ConfirmSignup**](docs/Auth.md#confirmsignup) | **Get** /signup/confirm | Confirm user signup
*Auth* | [**Login**](docs/Auth.md#login) | **Post** /login | User login
*Auth* | [**Logout**](docs/Auth.md#logout) | **Get** /logout/{provider} | User logout
*Auth* | [**Provider**](docs/Auth.md#provider) | **Get** /auth/{provider} | Auth Provider
*Auth* | [**ProviderCallback**](docs/Auth.md#providercallback) | **Get** /auth/{provider}/callback | Auth provider callback
*Auth* | [**Signup**](docs/Auth.md#signup) | **Post** /signup | User signup
*Identities* | [**Get**](docs/Identities.md#get) | **Get** /identity/{identity_handle} | Get identity
*Identities* | [**List**](docs/Identities.md#list) | **Get** /identity | List identities
*OrgConnections* | [**Create**](docs/OrgConnections.md#create) | **Post** /org/{org_handle}/conn | Create org connection
*OrgConnections* | [**Delete**](docs/OrgConnections.md#delete) | **Delete** /org/{org_handle}/conn/{conn_handle} | Delete org connection
*OrgConnections* | [**Get**](docs/OrgConnections.md#get) | **Get** /org/{org_handle}/conn/{conn_handle} | Get org connection
*OrgConnections* | [**List**](docs/OrgConnections.md#list) | **Get** /org/{org_handle}/conn | List org connections
*OrgConnections* | [**Test**](docs/OrgConnections.md#test) | **Post** /org/{org_handle}/conn/{conn_handle}/test | Test org connection
*OrgConnections* | [**Update**](docs/OrgConnections.md#update) | **Patch** /org/{org_handle}/conn/{conn_handle} | Update org connection
*OrgMembers* | [**ConfirmInvite**](docs/OrgMembers.md#confirminvite) | **Get** /org/{org_handle}/member/invite/confirm | Confirm org member invite
*OrgMembers* | [**Delete**](docs/OrgMembers.md#delete) | **Delete** /org/{org_handle}/member/{user_handle} | Delete org member
*OrgMembers* | [**DeleteInvite**](docs/OrgMembers.md#deleteinvite) | **Delete** /org/{org_handle}/member/invite | Delete org member invite
*OrgMembers* | [**Get**](docs/OrgMembers.md#get) | **Get** /org/{org_handle}/member/{user_handle} | Get org member
*OrgMembers* | [**Invite**](docs/OrgMembers.md#invite) | **Post** /org/{org_handle}/member/invite | Invite org member
*OrgMembers* | [**List**](docs/OrgMembers.md#list) | **Get** /org/{org_handle}/member | List Organization Members
*OrgMembers* | [**Update**](docs/OrgMembers.md#update) | **Patch** /org/{org_handle}/member/{user_handle} | Update org member
*OrgWorkspaceConnectionAssociations* | [**Create**](docs/OrgWorkspaceConnectionAssociations.md#create) | **Post** /org/{org_handle}/workspace/{workspace_handle}/conn | Create org workspace connection association
*OrgWorkspaceConnectionAssociations* | [**Delete**](docs/OrgWorkspaceConnectionAssociations.md#delete) | **Delete** /org/{org_handle}/workspace/{workspace_handle}/conn/{conn_handle} | Delete org workspace connection association
*OrgWorkspaceConnectionAssociations* | [**Get**](docs/OrgWorkspaceConnectionAssociations.md#get) | **Get** /org/{org_handle}/workspace/{workspace_handle}/conn/{conn_handle} | Get org workspace connection association
*OrgWorkspaceConnectionAssociations* | [**List**](docs/OrgWorkspaceConnectionAssociations.md#list) | **Get** /org/{org_handle}/workspace/{workspace_handle}/conn | List org workspace connection associations
*OrgWorkspaceMembers* | [**Create**](docs/OrgWorkspaceMembers.md#create) | **Post** /org/{org_handle}/workspace/{workspace_handle}/member | Create Org Workspace Member
*OrgWorkspaceMembers* | [**Delete**](docs/OrgWorkspaceMembers.md#delete) | **Delete** /org/{org_handle}/workspace/{workspace_handle}/member/{user_handle} | Delete Org Workspace Member
*OrgWorkspaceMembers* | [**Get**](docs/OrgWorkspaceMembers.md#get) | **Get** /org/{org_handle}/workspace/{workspace_handle}/member/{user_handle} | Get Org Workspace Member
*OrgWorkspaceMembers* | [**List**](docs/OrgWorkspaceMembers.md#list) | **Get** /org/{org_handle}/workspace/{workspace_handle}/member | List Organization Workspace Members
*OrgWorkspaceMembers* | [**Update**](docs/OrgWorkspaceMembers.md#update) | **Patch** /org/{org_handle}/workspace/{workspace_handle}/member/{user_handle} | Update Org Workspace Member
*OrgWorkspaceModVariables* | [**CreateSetting**](docs/OrgWorkspaceModVariables.md#createsetting) | **Post** /org/{org_handle}/workspace/{workspace_handle}/mod/{mod_alias}/variable | Create a setting for a mod variable in an organization workspace
*OrgWorkspaceModVariables* | [**DeleteSetting**](docs/OrgWorkspaceModVariables.md#deletesetting) | **Delete** /org/{org_handle}/workspace/{workspace_handle}/mod/{mod_alias}/variable/{variable_name} | Delete setting for a mod variable in an organization workspace
*OrgWorkspaceModVariables* | [**GetSetting**](docs/OrgWorkspaceModVariables.md#getsetting) | **Get** /org/{org_handle}/workspace/{workspace_handle}/mod/{mod_alias}/variable/{variable_name} | Get setting for a mod variable in an organization workspace
*OrgWorkspaceModVariables* | [**List**](docs/OrgWorkspaceModVariables.md#list) | **Get** /org/{org_handle}/workspace/{workspace_handle}/mod/{mod_alias}/variable | List variables in an organization workspace mod
*OrgWorkspaceModVariables* | [**UpdateSetting**](docs/OrgWorkspaceModVariables.md#updatesetting) | **Patch** /org/{org_handle}/workspace/{workspace_handle}/mod/{mod_alias}/variable/{variable_name} | Update setting for a mod variable in an organization workspace
*OrgWorkspaceMods* | [**Get**](docs/OrgWorkspaceMods.md#get) | **Get** /org/{org_handle}/workspace/{workspace_handle}/mod/{mod_alias} | Get organization workspace installed mod
*OrgWorkspaceMods* | [**Install**](docs/OrgWorkspaceMods.md#install) | **Post** /org/{org_handle}/workspace/{workspace_handle}/mod | Install a mod to an organization workspace
*OrgWorkspaceMods* | [**List**](docs/OrgWorkspaceMods.md#list) | **Get** /org/{org_handle}/workspace/{workspace_handle}/mod | List organization workspace installed mods
*OrgWorkspaceMods* | [**Uninstall**](docs/OrgWorkspaceMods.md#uninstall) | **Delete** /org/{org_handle}/workspace/{workspace_handle}/mod/{mod_alias} | Uninstall mod from organization workspace.
*OrgWorkspaceMods* | [**Update**](docs/OrgWorkspaceMods.md#update) | **Patch** /org/{org_handle}/workspace/{workspace_handle}/mod/{mod_alias} | Update a mod in an organization workspace
*OrgWorkspaces* | [**Create**](docs/OrgWorkspaces.md#create) | **Post** /org/{org_handle}/workspace | Create org workspace
*OrgWorkspaces* | [**Delete**](docs/OrgWorkspaces.md#delete) | **Delete** /org/{org_handle}/workspace/{workspace_handle} | Delete org workspace
*OrgWorkspaces* | [**Get**](docs/OrgWorkspaces.md#get) | **Get** /org/{org_handle}/workspace/{workspace_handle} | Get org workspace
*OrgWorkspaces* | [**GetQuery**](docs/OrgWorkspaces.md#getquery) | **Get** /org/{org_handle}/workspace/{workspace_handle}/query | Query org workspace
*OrgWorkspaces* | [**GetQueryWithExtensions**](docs/OrgWorkspaces.md#getquerywithextensions) | **Get** /org/{org_handle}/workspace/{workspace_handle}/query/data.{extensions} | Query org workspace with extensions
*OrgWorkspaces* | [**GetSchema**](docs/OrgWorkspaces.md#getschema) | **Get** /org/{org_handle}/workspace/{workspace_handle}/schema | Get org workspace schemas
*OrgWorkspaces* | [**List**](docs/OrgWorkspaces.md#list) | **Get** /org/{org_handle}/workspace | List org workspaces
*OrgWorkspaces* | [**ListDBLogs**](docs/OrgWorkspaces.md#listdblogs) | **Get** /org/{org_handle}/workspace/{workspace_handle}/logs | Org workspace logs
*OrgWorkspaces* | [**PostQuery**](docs/OrgWorkspaces.md#postquery) | **Post** /org/{org_handle}/workspace/{workspace_handle}/query | Query org workspace
*OrgWorkspaces* | [**PostQueryWithExtensions**](docs/OrgWorkspaces.md#postquerywithextensions) | **Post** /org/{org_handle}/workspace/{workspace_handle}/query/data.{extensions} | Query org workspace with extensions
*OrgWorkspaces* | [**Update**](docs/OrgWorkspaces.md#update) | **Patch** /org/{org_handle}/workspace/{workspace_handle} | Update org workspace
*Orgs* | [**Create**](docs/Orgs.md#create) | **Post** /org | Create org
*Orgs* | [**Delete**](docs/Orgs.md#delete) | **Delete** /org/{org_handle} | Delete org
*Orgs* | [**Get**](docs/Orgs.md#get) | **Get** /org/{org_handle} | Get org
*Orgs* | [**GetQuota**](docs/Orgs.md#getquota) | **Get** /org/{org_handle}/quota | Org quota
*Orgs* | [**List**](docs/Orgs.md#list) | **Get** /org | List orgs
*Orgs* | [**ListAuditLogs**](docs/Orgs.md#listauditlogs) | **Get** /org/{org_handle}/audit | Org audit logs
*Orgs* | [**ListFeatures**](docs/Orgs.md#listfeatures) | **Get** /org/{org_handle}/feature | Org features
*Orgs* | [**ListWorkspaceAuditLogs**](docs/Orgs.md#listworkspaceauditlogs) | **Get** /org/{org_handle}/workspace/{workspace_handle}/audit | Org workspace audit logs
*Orgs* | [**Update**](docs/Orgs.md#update) | **Patch** /org/{org_handle} | Update org
*UserConnections* | [**Create**](docs/UserConnections.md#create) | **Post** /user/{user_handle}/conn | Create user connection
*UserConnections* | [**Delete**](docs/UserConnections.md#delete) | **Delete** /user/{user_handle}/conn/{conn_handle} | Delete user connection
*UserConnections* | [**Get**](docs/UserConnections.md#get) | **Get** /user/{user_handle}/conn/{conn_handle} | Get user connection
*UserConnections* | [**List**](docs/UserConnections.md#list) | **Get** /user/{user_handle}/conn | List user connections
*UserConnections* | [**Test**](docs/UserConnections.md#test) | **Post** /user/{user_handle}/conn/{conn_handle}/test | Test user connection
*UserConnections* | [**Update**](docs/UserConnections.md#update) | **Patch** /user/{user_handle}/conn/{conn_handle} | Update user connection
*UserTokens* | [**Create**](docs/UserTokens.md#create) | **Post** /user/{user_handle}/token | Create token
*UserTokens* | [**Delete**](docs/UserTokens.md#delete) | **Delete** /user/{user_handle}/token/{token_id} | Delete token
*UserTokens* | [**Get**](docs/UserTokens.md#get) | **Get** /user/{user_handle}/token/{token_id} | Get token
*UserTokens* | [**List**](docs/UserTokens.md#list) | **Get** /user/{user_handle}/token | List tokens
*UserTokens* | [**Update**](docs/UserTokens.md#update) | **Patch** /user/{user_handle}/token/{token_id} | Update token
*UserWorkspaceConnectionAssociations* | [**Create**](docs/UserWorkspaceConnectionAssociations.md#create) | **Post** /user/{user_handle}/workspace/{workspace_handle}/conn | Create user workspace connection association
*UserWorkspaceConnectionAssociations* | [**Delete**](docs/UserWorkspaceConnectionAssociations.md#delete) | **Delete** /user/{user_handle}/workspace/{workspace_handle}/conn/{conn_handle} | Delete user workspace connection association
*UserWorkspaceConnectionAssociations* | [**Get**](docs/UserWorkspaceConnectionAssociations.md#get) | **Get** /user/{user_handle}/workspace/{workspace_handle}/conn/{conn_handle} | Get user workspace connection association
*UserWorkspaceConnectionAssociations* | [**List**](docs/UserWorkspaceConnectionAssociations.md#list) | **Get** /user/{user_handle}/workspace/{workspace_handle}/conn | List user workspace connection associations
*UserWorkspaceModVariables* | [**CreateSetting**](docs/UserWorkspaceModVariables.md#createsetting) | **Post** /user/{user_handle}/workspace/{workspace_handle}/mod/{mod_alias}/variable | Create a setting for a mod variable in a user workspace
*UserWorkspaceModVariables* | [**DeleteSetting**](docs/UserWorkspaceModVariables.md#deletesetting) | **Delete** /user/{user_handle}/workspace/{workspace_handle}/mod/{mod_alias}/variable/{variable_name} | Delete setting for a mod variable in a user workspace
*UserWorkspaceModVariables* | [**GetSetting**](docs/UserWorkspaceModVariables.md#getsetting) | **Get** /user/{user_handle}/workspace/{workspace_handle}/mod/{mod_alias}/variable/{variable_name} | Get setting for a mod variable in a user workspace
*UserWorkspaceModVariables* | [**List**](docs/UserWorkspaceModVariables.md#list) | **Get** /user/{user_handle}/workspace/{workspace_handle}/mod/{mod_alias}/variable | List variables for a user workspace mod
*UserWorkspaceModVariables* | [**UpdateSetting**](docs/UserWorkspaceModVariables.md#updatesetting) | **Patch** /user/{user_handle}/workspace/{workspace_handle}/mod/{mod_alias}/variable/{variable_name} | Update setting for a mod variable in a user workspace
*UserWorkspaceMods* | [**Get**](docs/UserWorkspaceMods.md#get) | **Get** /user/{user_handle}/workspace/{workspace_handle}/mod/{mod_alias} | Get user workspace installed mod
*UserWorkspaceMods* | [**Install**](docs/UserWorkspaceMods.md#install) | **Post** /user/{user_handle}/workspace/{workspace_handle}/mod | Install a mod to a user&#39;s workspace
*UserWorkspaceMods* | [**List**](docs/UserWorkspaceMods.md#list) | **Get** /user/{user_handle}/workspace/{workspace_handle}/mod | List user workspace installed mods
*UserWorkspaceMods* | [**Uninstall**](docs/UserWorkspaceMods.md#uninstall) | **Delete** /user/{user_handle}/workspace/{workspace_handle}/mod/{mod_alias} | Uninstall mod from a user&#39;s workspace.
*UserWorkspaceMods* | [**Update**](docs/UserWorkspaceMods.md#update) | **Patch** /user/{user_handle}/workspace/{workspace_handle}/mod/{mod_alias} | Update a mod in a user&#39;s workspace
*UserWorkspaces* | [**Create**](docs/UserWorkspaces.md#create) | **Post** /user/{user_handle}/workspace | Create user workspace
*UserWorkspaces* | [**Delete**](docs/UserWorkspaces.md#delete) | **Delete** /user/{user_handle}/workspace/{workspace_handle} | Delete user workspace
*UserWorkspaces* | [**Get**](docs/UserWorkspaces.md#get) | **Get** /user/{user_handle}/workspace/{workspace_handle} | Get user workspace
*UserWorkspaces* | [**GetQuery**](docs/UserWorkspaces.md#getquery) | **Get** /user/{user_handle}/workspace/{workspace_handle}/query | Query user workspace
*UserWorkspaces* | [**GetQueryWithExtensions**](docs/UserWorkspaces.md#getquerywithextensions) | **Get** /user/{user_handle}/workspace/{workspace_handle}/query/data.{extensions} | Query user workspace with extensions
*UserWorkspaces* | [**GetSchema**](docs/UserWorkspaces.md#getschema) | **Get** /user/{user_handle}/workspace/{workspace_handle}/schema | Get user workspace schemas
*UserWorkspaces* | [**List**](docs/UserWorkspaces.md#list) | **Get** /user/{user_handle}/workspace | List user workspaces
*UserWorkspaces* | [**ListDBLogs**](docs/UserWorkspaces.md#listdblogs) | **Get** /user/{user_handle}/workspace/{workspace_handle}/logs | User workspace logs
*UserWorkspaces* | [**PostQuery**](docs/UserWorkspaces.md#postquery) | **Post** /user/{user_handle}/workspace/{workspace_handle}/query | Query user workspace
*UserWorkspaces* | [**PostQueryWithExtensions**](docs/UserWorkspaces.md#postquerywithextensions) | **Post** /user/{user_handle}/workspace/{workspace_handle}/query/data.{extensions} | Query user workspace with extensions
*UserWorkspaces* | [**Update**](docs/UserWorkspaces.md#update) | **Patch** /user/{user_handle}/workspace/{workspace_handle} | Update user workspace
*Users* | [**Create**](docs/Users.md#create) | **Post** /user | Create user
*Users* | [**CreateDBPassword**](docs/Users.md#createdbpassword) | **Post** /user/{user_handle}/password | Create user password
*Users* | [**Delete**](docs/Users.md#delete) | **Delete** /user/{user_handle} | Delete user
*Users* | [**Get**](docs/Users.md#get) | **Get** /user/{user_handle} | Get user
*Users* | [**GetDBPassword**](docs/Users.md#getdbpassword) | **Get** /user/{user_handle}/password | Get user password
*Users* | [**GetQuota**](docs/Users.md#getquota) | **Get** /user/{user_handle}/quota | User quota
*Users* | [**List**](docs/Users.md#list) | **Get** /user | List users
*Users* | [**ListAuditLogs**](docs/Users.md#listauditlogs) | **Get** /user/{user_handle}/audit | User audit logs
*Users* | [**ListFeatures**](docs/Users.md#listfeatures) | **Get** /user/{user_handle}/feature | User features
*Users* | [**ListWorkspaceAuditLogs**](docs/Users.md#listworkspaceauditlogs) | **Get** /user/{user_handle}/workspace/{workspace_handle}/audit | User workspace audit logs
*Users* | [**Update**](docs/Users.md#update) | **Patch** /user/{user_handle} | Update user


## Documentation For Models

 - [ActorOrg](docs/ActorOrg.md)
 - [ActorOrgWorkspace](docs/ActorOrgWorkspace.md)
 - [ActorWorkspace](docs/ActorWorkspace.md)
 - [AuditRecord](docs/AuditRecord.md)
 - [Connection](docs/Connection.md)
 - [ConnectionTestResult](docs/ConnectionTestResult.md)
 - [CreateConnectionRequest](docs/CreateConnectionRequest.md)
 - [CreateOrgRequest](docs/CreateOrgRequest.md)
 - [CreateOrgWorkspaceUserRequest](docs/CreateOrgWorkspaceUserRequest.md)
 - [CreateUserPasswordRequest](docs/CreateUserPasswordRequest.md)
 - [CreateUserRequest](docs/CreateUserRequest.md)
 - [CreateWorkspaceConnRequest](docs/CreateWorkspaceConnRequest.md)
 - [CreateWorkspaceModRequest](docs/CreateWorkspaceModRequest.md)
 - [CreateWorkspaceModVariableSettingRequest](docs/CreateWorkspaceModVariableSettingRequest.md)
 - [CreateWorkspaceRequest](docs/CreateWorkspaceRequest.md)
 - [ErrorDetailModel](docs/ErrorDetailModel.md)
 - [ErrorModel](docs/ErrorModel.md)
 - [Feature](docs/Feature.md)
 - [Identity](docs/Identity.md)
 - [InviteOrgUserRequest](docs/InviteOrgUserRequest.md)
 - [ListActorOrgsResponse](docs/ListActorOrgsResponse.md)
 - [ListActorWorkspacesResponse](docs/ListActorWorkspacesResponse.md)
 - [ListAuditLogsResponse](docs/ListAuditLogsResponse.md)
 - [ListConnectionsResponse](docs/ListConnectionsResponse.md)
 - [ListFeaturesResponse](docs/ListFeaturesResponse.md)
 - [ListIdentitiesResponse](docs/ListIdentitiesResponse.md)
 - [ListLogsResponse](docs/ListLogsResponse.md)
 - [ListOrgUsersResponse](docs/ListOrgUsersResponse.md)
 - [ListOrgWorkspaceUsersResponse](docs/ListOrgWorkspaceUsersResponse.md)
 - [ListOrgsResponse](docs/ListOrgsResponse.md)
 - [ListTokensResponse](docs/ListTokensResponse.md)
 - [ListUserOrgsResponse](docs/ListUserOrgsResponse.md)
 - [ListUsersResponse](docs/ListUsersResponse.md)
 - [ListWorkspaceConnResponse](docs/ListWorkspaceConnResponse.md)
 - [ListWorkspaceModVariablesResponse](docs/ListWorkspaceModVariablesResponse.md)
 - [ListWorkspaceModsResponse](docs/ListWorkspaceModsResponse.md)
 - [ListWorkspacesResponse](docs/ListWorkspacesResponse.md)
 - [LogRecord](docs/LogRecord.md)
 - [Org](docs/Org.md)
 - [OrgQuota](docs/OrgQuota.md)
 - [OrgUser](docs/OrgUser.md)
 - [OrgWorkspaceUser](docs/OrgWorkspaceUser.md)
 - [Quota](docs/Quota.md)
 - [SchemaInfo](docs/SchemaInfo.md)
 - [SchemaTable](docs/SchemaTable.md)
 - [SchemaTableColumn](docs/SchemaTableColumn.md)
 - [Token](docs/Token.md)
 - [UpdateConnectionRequest](docs/UpdateConnectionRequest.md)
 - [UpdateOrgRequest](docs/UpdateOrgRequest.md)
 - [UpdateOrgUserRequest](docs/UpdateOrgUserRequest.md)
 - [UpdateOrgWorkspaceUserRequest](docs/UpdateOrgWorkspaceUserRequest.md)
 - [UpdateTokenRequest](docs/UpdateTokenRequest.md)
 - [UpdateUserRequest](docs/UpdateUserRequest.md)
 - [UpdateWorkspaceModRequest](docs/UpdateWorkspaceModRequest.md)
 - [UpdateWorkspaceModVariableSettingRequest](docs/UpdateWorkspaceModVariableSettingRequest.md)
 - [UpdateWorkspaceRequest](docs/UpdateWorkspaceRequest.md)
 - [User](docs/User.md)
 - [UserDatabasePassword](docs/UserDatabasePassword.md)
 - [UserLoginRequest](docs/UserLoginRequest.md)
 - [UserOrg](docs/UserOrg.md)
 - [UserQuota](docs/UserQuota.md)
 - [UserSignupRequest](docs/UserSignupRequest.md)
 - [Workspace](docs/Workspace.md)
 - [WorkspaceConn](docs/WorkspaceConn.md)
 - [WorkspaceMod](docs/WorkspaceMod.md)
 - [WorkspaceModVariable](docs/WorkspaceModVariable.md)
 - [WorkspaceQueryResult](docs/WorkspaceQueryResult.md)
 - [WorkspaceQueryResultColumn](docs/WorkspaceQueryResultColumn.md)
 - [WorkspaceSchema](docs/WorkspaceSchema.md)


## Documentation For Authorization



### ApiKeyAuth

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

help@steampipe.io

