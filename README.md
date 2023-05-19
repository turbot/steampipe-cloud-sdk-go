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
*Auth* | [**LoginTokenCreate**](docs/Auth.md#logintokencreate) | **Post** /login/token | Generate temporary token request
*Auth* | [**LoginTokenDelete**](docs/Auth.md#logintokendelete) | **Delete** /login/token/{temporary_token_request_id} | Delete temporary token request
*Auth* | [**LoginTokenGet**](docs/Auth.md#logintokenget) | **Get** /login/token/{temporary_token_request_id} | Get temporary token request
*Auth* | [**LoginTokenUpdate**](docs/Auth.md#logintokenupdate) | **Patch** /login/token/{temporary_token_request_id} | Update temporary token request
*Auth* | [**Logout**](docs/Auth.md#logout) | **Get** /logout/{provider} | User logout
*Auth* | [**Provider**](docs/Auth.md#provider) | **Get** /auth/{provider} | Auth Provider
*Auth* | [**ProviderCallback**](docs/Auth.md#providercallback) | **Get** /auth/{provider}/callback | Auth provider callback
*Auth* | [**Signup**](docs/Auth.md#signup) | **Post** /signup | User signup
*Identities* | [**Get**](docs/Identities.md#get) | **Get** /identity/{identity_handle} | Get identity
*Identities* | [**GetAvatar**](docs/Identities.md#getavatar) | **Get** /identity/{identity_handle}/avatar | Get identity avatar
*Identities* | [**List**](docs/Identities.md#list) | **Get** /identity | List identities
*OrgConnections* | [**Create**](docs/OrgConnections.md#create) | **Post** /org/{org_handle}/connection | Create org connection
*OrgConnections* | [**CreateDeprecated**](docs/OrgConnections.md#createdeprecated) | **Post** /org/{org_handle}/conn | Create org connection
*OrgConnections* | [**Delete**](docs/OrgConnections.md#delete) | **Delete** /org/{org_handle}/connection/{connection_handle} | Delete org connection
*OrgConnections* | [**DeleteDeprecated**](docs/OrgConnections.md#deletedeprecated) | **Delete** /org/{org_handle}/conn/{conn_handle} | Delete org connection
*OrgConnections* | [**Get**](docs/OrgConnections.md#get) | **Get** /org/{org_handle}/connection/{connection_handle} | Get org connection
*OrgConnections* | [**GetDeprecated**](docs/OrgConnections.md#getdeprecated) | **Get** /org/{org_handle}/conn/{conn_handle} | Get org connection
*OrgConnections* | [**List**](docs/OrgConnections.md#list) | **Get** /org/{org_handle}/connection | List org connections
*OrgConnections* | [**ListDeprecated**](docs/OrgConnections.md#listdeprecated) | **Get** /org/{org_handle}/conn | List org connections
*OrgConnections* | [**ListWorkspaces**](docs/OrgConnections.md#listworkspaces) | **Get** /org/{org_handle}/connection/{connection_handle}/workspace | List org connection workspaces
*OrgConnections* | [**ListWorkspacesDeprecated**](docs/OrgConnections.md#listworkspacesdeprecated) | **Get** /org/{org_handle}/conn/{conn_handle}/workspace | List org connection workspaces
*OrgConnections* | [**Test**](docs/OrgConnections.md#test) | **Post** /org/{org_handle}/connection/{connection_handle}/test | Test org connection
*OrgConnections* | [**TestDeprecated**](docs/OrgConnections.md#testdeprecated) | **Post** /org/{org_handle}/conn/{conn_handle}/test | Test org connection
*OrgConnections* | [**Update**](docs/OrgConnections.md#update) | **Patch** /org/{org_handle}/connection/{connection_handle} | Update org connection
*OrgConnections* | [**UpdateDeprecated**](docs/OrgConnections.md#updatedeprecated) | **Patch** /org/{org_handle}/conn/{conn_handle} | Update org connection
*OrgMembers* | [**ConfirmInvite**](docs/OrgMembers.md#confirminvite) | **Get** /org/{org_handle}/member/invite/confirm | Confirm org member invite
*OrgMembers* | [**Delete**](docs/OrgMembers.md#delete) | **Delete** /org/{org_handle}/member/{user_handle} | Delete org member
*OrgMembers* | [**DeleteInvite**](docs/OrgMembers.md#deleteinvite) | **Delete** /org/{org_handle}/member/invite | Delete org member invite
*OrgMembers* | [**Get**](docs/OrgMembers.md#get) | **Get** /org/{org_handle}/member/{user_handle} | Get org member
*OrgMembers* | [**Invite**](docs/OrgMembers.md#invite) | **Post** /org/{org_handle}/member/invite | Invite org member
*OrgMembers* | [**List**](docs/OrgMembers.md#list) | **Get** /org/{org_handle}/member | List Organization Members
*OrgMembers* | [**Update**](docs/OrgMembers.md#update) | **Patch** /org/{org_handle}/member/{user_handle} | Update org member
*OrgProcesses* | [**Get**](docs/OrgProcesses.md#get) | **Get** /org/{org_handle}/process/{process_id} | Get Org process
*OrgProcesses* | [**List**](docs/OrgProcesses.md#list) | **Get** /org/{org_handle}/process | List Org processes
*OrgProcesses* | [**Log**](docs/OrgProcesses.md#log) | **Get** /org/{org_handle}/process/{process_id}/log/{log_file}.{content_type} | List Org process logs
*OrgWorkspaceAggregators* | [**Create**](docs/OrgWorkspaceAggregators.md#create) | **Post** /org/{org_handle}/workspace/{workspace_handle}/aggregator | Create an aggregator for an org workspace
*OrgWorkspaceAggregators* | [**Delete**](docs/OrgWorkspaceAggregators.md#delete) | **Delete** /org/{org_handle}/workspace/{workspace_handle}/aggregator/{aggregator_handle} | Delete an aggregator for a org workspace
*OrgWorkspaceAggregators* | [**Get**](docs/OrgWorkspaceAggregators.md#get) | **Get** /org/{org_handle}/workspace/{workspace_handle}/aggregator/{aggregator_handle} | Get an aggregator for a org workspace
*OrgWorkspaceAggregators* | [**GetConnection**](docs/OrgWorkspaceAggregators.md#getconnection) | **Get** /org/{org_handle}/workspace/{workspace_handle}/aggregator/{aggregator_handle}/connection/{connection_handle} | Get a connection in the scope of an aggregator for a org workspace
*OrgWorkspaceAggregators* | [**List**](docs/OrgWorkspaceAggregators.md#list) | **Get** /org/{org_handle}/workspace/{workspace_handle}/aggregator | List aggregators for an org workspace
*OrgWorkspaceAggregators* | [**ListConnections**](docs/OrgWorkspaceAggregators.md#listconnections) | **Get** /org/{org_handle}/workspace/{workspace_handle}/aggregator/{aggregator_handle}/connection | List connections in the scope of an aggregator for a org workspace
*OrgWorkspaceAggregators* | [**Update**](docs/OrgWorkspaceAggregators.md#update) | **Patch** /org/{org_handle}/workspace/{workspace_handle}/aggregator/{aggregator_handle} | Update an aggregator for a org workspace
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
*OrgWorkspacePipelines* | [**Command**](docs/OrgWorkspacePipelines.md#command) | **Post** /org/{org_handle}/workspace/{workspace_handle}/pipeline/{pipeline_id}/command | Run org workspace pipeline command
*OrgWorkspacePipelines* | [**Create**](docs/OrgWorkspacePipelines.md#create) | **Post** /org/{org_handle}/workspace/{workspace_handle}/pipeline | Create org workspace pipeline
*OrgWorkspacePipelines* | [**Delete**](docs/OrgWorkspacePipelines.md#delete) | **Delete** /org/{org_handle}/workspace/{workspace_handle}/pipeline/{pipeline_id} | Delete org workspace pipeline
*OrgWorkspacePipelines* | [**Get**](docs/OrgWorkspacePipelines.md#get) | **Get** /org/{org_handle}/workspace/{workspace_handle}/pipeline/{pipeline_id} | Get org workspace pipeline
*OrgWorkspacePipelines* | [**List**](docs/OrgWorkspacePipelines.md#list) | **Get** /org/{org_handle}/workspace/{workspace_handle}/pipeline | List org workspace pipelines
*OrgWorkspacePipelines* | [**Update**](docs/OrgWorkspacePipelines.md#update) | **Patch** /org/{org_handle}/workspace/{workspace_handle}/pipeline/{pipeline_id} | Update org workspace pipeline
*OrgWorkspaceProcesses* | [**Get**](docs/OrgWorkspaceProcesses.md#get) | **Get** /org/{org_handle}/workspace/{workspace_handle}/process/{process_id} | Get org workspace process
*OrgWorkspaceProcesses* | [**List**](docs/OrgWorkspaceProcesses.md#list) | **Get** /org/{org_handle}/workspace/{workspace_handle}/process | List org workspace processes
*OrgWorkspaceProcesses* | [**Log**](docs/OrgWorkspaceProcesses.md#log) | **Get** /org/{org_handle}/workspace/{workspace_handle}/process/{process_id}/log/{log_file}.{content_type} | List org workspace process logs
*OrgWorkspaceSchemas* | [**Attach**](docs/OrgWorkspaceSchemas.md#attach) | **Post** /org/{org_handle}/workspace/{workspace_handle}/schema | Attach a schema to an org workspace
*OrgWorkspaceSchemas* | [**Detach**](docs/OrgWorkspaceSchemas.md#detach) | **Delete** /org/{org_handle}/workspace/{workspace_handle}/schema/{schema_name} | Detach a schema from an org workspace
*OrgWorkspaceSchemas* | [**Get**](docs/OrgWorkspaceSchemas.md#get) | **Get** /org/{org_handle}/workspace/{workspace_handle}/schema/{schema_name} | Get org workspace schema
*OrgWorkspaceSchemas* | [**Get_0**](docs/OrgWorkspaceSchemas.md#get_0) | **Get** /org/{org_handle}/workspace/{workspace_handle}/schema/{schema_name}/table | List org workspace schema tables
*OrgWorkspaceSchemas* | [**List**](docs/OrgWorkspaceSchemas.md#list) | **Get** /org/{org_handle}/workspace/{workspace_handle}/schema | List org workspace schemas
*OrgWorkspaceSnapshots* | [**Create**](docs/OrgWorkspaceSnapshots.md#create) | **Post** /org/{org_handle}/workspace/{workspace_handle}/snapshot | Create org workspace snapshot
*OrgWorkspaceSnapshots* | [**Delete**](docs/OrgWorkspaceSnapshots.md#delete) | **Delete** /org/{org_handle}/workspace/{workspace_handle}/snapshot/{snapshot_id} | Delete org workspace snapshot
*OrgWorkspaceSnapshots* | [**Download**](docs/OrgWorkspaceSnapshots.md#download) | **Get** /org/{org_handle}/workspace/{workspace_handle}/snapshot/{snapshot_id}.{content_type} | Download org workspace snapshot
*OrgWorkspaceSnapshots* | [**Get**](docs/OrgWorkspaceSnapshots.md#get) | **Get** /org/{org_handle}/workspace/{workspace_handle}/snapshot/{snapshot_id} | Get org workspace snapshot
*OrgWorkspaceSnapshots* | [**List**](docs/OrgWorkspaceSnapshots.md#list) | **Get** /org/{org_handle}/workspace/{workspace_handle}/snapshot | List org workspace snapshots
*OrgWorkspaceSnapshots* | [**Update**](docs/OrgWorkspaceSnapshots.md#update) | **Patch** /org/{org_handle}/workspace/{workspace_handle}/snapshot/{snapshot_id} | Update org workspace snapshot
*OrgWorkspaces* | [**Command**](docs/OrgWorkspaces.md#command) | **Post** /org/{org_handle}/workspace/{workspace_handle}/command | Run org workspace command
*OrgWorkspaces* | [**Create**](docs/OrgWorkspaces.md#create) | **Post** /org/{org_handle}/workspace | Create org workspace
*OrgWorkspaces* | [**Delete**](docs/OrgWorkspaces.md#delete) | **Delete** /org/{org_handle}/workspace/{workspace_handle} | Delete org workspace
*OrgWorkspaces* | [**Get**](docs/OrgWorkspaces.md#get) | **Get** /org/{org_handle}/workspace/{workspace_handle} | Get org workspace
*OrgWorkspaces* | [**GetQuery**](docs/OrgWorkspaces.md#getquery) | **Get** /org/{org_handle}/workspace/{workspace_handle}/query | Query org workspace
*OrgWorkspaces* | [**GetQueryWithExtensions**](docs/OrgWorkspaces.md#getquerywithextensions) | **Get** /org/{org_handle}/workspace/{workspace_handle}/query/data.{extensions} | Query org workspace with extensions
*OrgWorkspaces* | [**List**](docs/OrgWorkspaces.md#list) | **Get** /org/{org_handle}/workspace | List org workspaces
*OrgWorkspaces* | [**ListAuditLogs**](docs/OrgWorkspaces.md#listauditlogs) | **Get** /org/{org_handle}/workspace/{workspace_handle}/audit_log | Org workspace audit logs
*OrgWorkspaces* | [**ListDBLogs**](docs/OrgWorkspaces.md#listdblogs) | **Get** /org/{org_handle}/workspace/{workspace_handle}/db_log | Org workspace logs
*OrgWorkspaces* | [**PostQuery**](docs/OrgWorkspaces.md#postquery) | **Post** /org/{org_handle}/workspace/{workspace_handle}/query | Query org workspace
*OrgWorkspaces* | [**PostQueryWithExtensions**](docs/OrgWorkspaces.md#postquerywithextensions) | **Post** /org/{org_handle}/workspace/{workspace_handle}/query/data.{extensions} | Query org workspace with extensions
*OrgWorkspaces* | [**Update**](docs/OrgWorkspaces.md#update) | **Patch** /org/{org_handle}/workspace/{workspace_handle} | Update org workspace
*Orgs* | [**Create**](docs/Orgs.md#create) | **Post** /org | Create org
*Orgs* | [**CreateAvatar**](docs/Orgs.md#createavatar) | **Post** /org/{org_handle}/avatar | Create org avatar
*Orgs* | [**Delete**](docs/Orgs.md#delete) | **Delete** /org/{org_handle} | Delete org
*Orgs* | [**DeleteAvatar**](docs/Orgs.md#deleteavatar) | **Delete** /org/{org_handle}/avatar | Delete org avatar
*Orgs* | [**Get**](docs/Orgs.md#get) | **Get** /org/{org_handle} | Get org
*Orgs* | [**GetQuota**](docs/Orgs.md#getquota) | **Get** /org/{org_handle}/quota | Org quota
*Orgs* | [**List**](docs/Orgs.md#list) | **Get** /org | List orgs
*Orgs* | [**ListAuditLogs**](docs/Orgs.md#listauditlogs) | **Get** /org/{org_handle}/audit_log | Org audit logs
*Orgs* | [**ListConstraints**](docs/Orgs.md#listconstraints) | **Get** /org/{org_handle}/constraint | List org constraints
*Orgs* | [**ListFeatures**](docs/Orgs.md#listfeatures) | **Get** /org/{org_handle}/feature | Org features
*Orgs* | [**Update**](docs/Orgs.md#update) | **Patch** /org/{org_handle} | Update org
*UserConnections* | [**Create**](docs/UserConnections.md#create) | **Post** /user/{user_handle}/connection | Create user connection
*UserConnections* | [**CreateDeprecated**](docs/UserConnections.md#createdeprecated) | **Post** /user/{user_handle}/conn | Create user connection
*UserConnections* | [**Delete**](docs/UserConnections.md#delete) | **Delete** /user/{user_handle}/connection/{connection_handle} | Delete user connection
*UserConnections* | [**DeleteDeprecated**](docs/UserConnections.md#deletedeprecated) | **Delete** /user/{user_handle}/conn/{conn_handle} | Delete user connection
*UserConnections* | [**Get**](docs/UserConnections.md#get) | **Get** /user/{user_handle}/connection/{connection_handle} | Get user connection
*UserConnections* | [**GetDeprecated**](docs/UserConnections.md#getdeprecated) | **Get** /user/{user_handle}/conn/{conn_handle} | Get user connection
*UserConnections* | [**List**](docs/UserConnections.md#list) | **Get** /user/{user_handle}/connection | List user connections
*UserConnections* | [**ListDeprecated**](docs/UserConnections.md#listdeprecated) | **Get** /user/{user_handle}/conn | List user connections
*UserConnections* | [**ListWorkspaces**](docs/UserConnections.md#listworkspaces) | **Get** /user/{user_handle}/connection/{connection_handle}/workspace | List user connection workspaces
*UserConnections* | [**ListWorkspacesDeprecated**](docs/UserConnections.md#listworkspacesdeprecated) | **Get** /user/{user_handle}/conn/{conn_handle}/workspace | List user connection workspaces
*UserConnections* | [**Test**](docs/UserConnections.md#test) | **Post** /user/{user_handle}/connection/{connection_handle}/test | Test user connection
*UserConnections* | [**TestDeprecated**](docs/UserConnections.md#testdeprecated) | **Post** /user/{user_handle}/conn/{conn_handle}/test | Test user connection
*UserConnections* | [**Update**](docs/UserConnections.md#update) | **Patch** /user/{user_handle}/connection/{connection_handle} | Update user connection
*UserConnections* | [**UpdateDeprecated**](docs/UserConnections.md#updatedeprecated) | **Patch** /user/{user_handle}/conn/{conn_handle} | Update user connection
*UserProcesses* | [**Get**](docs/UserProcesses.md#get) | **Get** /user/{user_handle}/process/{process_id} | Get User process
*UserProcesses* | [**List**](docs/UserProcesses.md#list) | **Get** /user/{user_handle}/process | List User processes
*UserProcesses* | [**Log**](docs/UserProcesses.md#log) | **Get** /user/{user_handle}/process/{process_id}/log/{log_file}.{content_type} | List user process logs
*UserTokens* | [**Create**](docs/UserTokens.md#create) | **Post** /user/{user_handle}/token | Create token
*UserTokens* | [**Delete**](docs/UserTokens.md#delete) | **Delete** /user/{user_handle}/token/{token_id} | Delete token
*UserTokens* | [**Get**](docs/UserTokens.md#get) | **Get** /user/{user_handle}/token/{token_id} | Get token
*UserTokens* | [**List**](docs/UserTokens.md#list) | **Get** /user/{user_handle}/token | List tokens
*UserTokens* | [**Update**](docs/UserTokens.md#update) | **Patch** /user/{user_handle}/token/{token_id} | Update token
*UserWorkspaceAggregators* | [**Create**](docs/UserWorkspaceAggregators.md#create) | **Post** /user/{user_handle}/workspace/{workspace_handle}/aggregator | Create an aggregator for a user workspace
*UserWorkspaceAggregators* | [**Delete**](docs/UserWorkspaceAggregators.md#delete) | **Delete** /user/{user_handle}/workspace/{workspace_handle}/aggregator/{aggregator_handle} | Delete an aggregator for a user workspace
*UserWorkspaceAggregators* | [**Get**](docs/UserWorkspaceAggregators.md#get) | **Get** /user/{user_handle}/workspace/{workspace_handle}/aggregator/{aggregator_handle} | Get an aggregator for a user workspace
*UserWorkspaceAggregators* | [**GetConnection**](docs/UserWorkspaceAggregators.md#getconnection) | **Get** /user/{user_handle}/workspace/{workspace_handle}/aggregator/{aggregator_handle}/connection/{connection_handle} | Get a connection in the scope of an aggregator for a user workspace
*UserWorkspaceAggregators* | [**List**](docs/UserWorkspaceAggregators.md#list) | **Get** /user/{user_handle}/workspace/{workspace_handle}/aggregator | List aggregators for a user workspace
*UserWorkspaceAggregators* | [**ListConnections**](docs/UserWorkspaceAggregators.md#listconnections) | **Get** /user/{user_handle}/workspace/{workspace_handle}/aggregator/{aggregator_handle}/connection | List connections in the scope of an aggregator for a user workspace
*UserWorkspaceAggregators* | [**Update**](docs/UserWorkspaceAggregators.md#update) | **Patch** /user/{user_handle}/workspace/{workspace_handle}/aggregator/{aggregator_handle} | Update an aggregator for a user workspace
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
*UserWorkspacePipelines* | [**Command**](docs/UserWorkspacePipelines.md#command) | **Post** /user/{user_handle}/workspace/{workspace_handle}/pipeline/{pipeline_id}/command | Run user workspace pipeline command
*UserWorkspacePipelines* | [**Create**](docs/UserWorkspacePipelines.md#create) | **Post** /user/{user_handle}/workspace/{workspace_handle}/pipeline | Create user workspace pipeline
*UserWorkspacePipelines* | [**Delete**](docs/UserWorkspacePipelines.md#delete) | **Delete** /user/{user_handle}/workspace/{workspace_handle}/pipeline/{pipeline_id} | Delete user workspace pipeline
*UserWorkspacePipelines* | [**Get**](docs/UserWorkspacePipelines.md#get) | **Get** /user/{user_handle}/workspace/{workspace_handle}/pipeline/{pipeline_id} | Get user workspace pipeline
*UserWorkspacePipelines* | [**List**](docs/UserWorkspacePipelines.md#list) | **Get** /user/{user_handle}/workspace/{workspace_handle}/pipeline | List user workspace pipelines
*UserWorkspacePipelines* | [**Update**](docs/UserWorkspacePipelines.md#update) | **Patch** /user/{user_handle}/workspace/{workspace_handle}/pipeline/{pipeline_id} | Update user workspace pipeline
*UserWorkspaceProcesses* | [**Get**](docs/UserWorkspaceProcesses.md#get) | **Get** /user/{user_handle}/workspace/{workspace_handle}/process/{process_id} | Get user workspace process
*UserWorkspaceProcesses* | [**List**](docs/UserWorkspaceProcesses.md#list) | **Get** /user/{user_handle}/workspace/{workspace_handle}/process | List user workspace processes
*UserWorkspaceProcesses* | [**Log**](docs/UserWorkspaceProcesses.md#log) | **Get** /user/{user_handle}/workspace/{workspace_handle}/process/{process_id}/log/{log_file}.{content_type} | List user workspace process logs
*UserWorkspaceSchemas* | [**Attach**](docs/UserWorkspaceSchemas.md#attach) | **Post** /user/{user_handle}/workspace/{workspace_handle}/schema | Attach a schema to a user workspace
*UserWorkspaceSchemas* | [**Detach**](docs/UserWorkspaceSchemas.md#detach) | **Delete** /user/{user_handle}/workspace/{workspace_handle}/schema/{schema_name} | Detach a schema from a user workspace
*UserWorkspaceSchemas* | [**Get**](docs/UserWorkspaceSchemas.md#get) | **Get** /user/{user_handle}/workspace/{workspace_handle}/schema/{schema_name} | Get user workspace schema
*UserWorkspaceSchemas* | [**Get_0**](docs/UserWorkspaceSchemas.md#get_0) | **Get** /user/{user_handle}/workspace/{workspace_handle}/schema/{schema_name}/table | List user workspace schema tables
*UserWorkspaceSchemas* | [**List**](docs/UserWorkspaceSchemas.md#list) | **Get** /user/{user_handle}/workspace/{workspace_handle}/schema | List user workspace schemas
*UserWorkspaceSnapshots* | [**Create**](docs/UserWorkspaceSnapshots.md#create) | **Post** /user/{user_handle}/workspace/{workspace_handle}/snapshot | Create user workspace snapshot
*UserWorkspaceSnapshots* | [**Delete**](docs/UserWorkspaceSnapshots.md#delete) | **Delete** /user/{user_handle}/workspace/{workspace_handle}/snapshot/{snapshot_id} | Delete user workspace snapshot
*UserWorkspaceSnapshots* | [**Download**](docs/UserWorkspaceSnapshots.md#download) | **Get** /user/{user_handle}/workspace/{workspace_handle}/snapshot/{snapshot_id}.{content_type} | Download user workspace snapshot
*UserWorkspaceSnapshots* | [**Get**](docs/UserWorkspaceSnapshots.md#get) | **Get** /user/{user_handle}/workspace/{workspace_handle}/snapshot/{snapshot_id} | Get user workspace snapshot
*UserWorkspaceSnapshots* | [**List**](docs/UserWorkspaceSnapshots.md#list) | **Get** /user/{user_handle}/workspace/{workspace_handle}/snapshot | List user workspace snapshots
*UserWorkspaceSnapshots* | [**Update**](docs/UserWorkspaceSnapshots.md#update) | **Patch** /user/{user_handle}/workspace/{workspace_handle}/snapshot/{snapshot_id} | Update user workspace snapshot
*UserWorkspaces* | [**Command**](docs/UserWorkspaces.md#command) | **Post** /user/{user_handle}/workspace/{workspace_handle}/command | Run user workspace command
*UserWorkspaces* | [**Create**](docs/UserWorkspaces.md#create) | **Post** /user/{user_handle}/workspace | Create user workspace
*UserWorkspaces* | [**Delete**](docs/UserWorkspaces.md#delete) | **Delete** /user/{user_handle}/workspace/{workspace_handle} | Delete user workspace
*UserWorkspaces* | [**Get**](docs/UserWorkspaces.md#get) | **Get** /user/{user_handle}/workspace/{workspace_handle} | Get user workspace
*UserWorkspaces* | [**GetQuery**](docs/UserWorkspaces.md#getquery) | **Get** /user/{user_handle}/workspace/{workspace_handle}/query | Query user workspace
*UserWorkspaces* | [**GetQueryWithExtensions**](docs/UserWorkspaces.md#getquerywithextensions) | **Get** /user/{user_handle}/workspace/{workspace_handle}/query/data.{extensions} | Query user workspace with extensions
*UserWorkspaces* | [**List**](docs/UserWorkspaces.md#list) | **Get** /user/{user_handle}/workspace | List user workspaces
*UserWorkspaces* | [**ListAuditLogs**](docs/UserWorkspaces.md#listauditlogs) | **Get** /user/{user_handle}/workspace/{workspace_handle}/audit_log | User workspace audit logs
*UserWorkspaces* | [**ListDBLogs**](docs/UserWorkspaces.md#listdblogs) | **Get** /user/{user_handle}/workspace/{workspace_handle}/db_log | User workspace logs
*UserWorkspaces* | [**PostQuery**](docs/UserWorkspaces.md#postquery) | **Post** /user/{user_handle}/workspace/{workspace_handle}/query | Query user workspace
*UserWorkspaces* | [**PostQueryWithExtensions**](docs/UserWorkspaces.md#postquerywithextensions) | **Post** /user/{user_handle}/workspace/{workspace_handle}/query/data.{extensions} | Query user workspace with extensions
*UserWorkspaces* | [**Update**](docs/UserWorkspaces.md#update) | **Patch** /user/{user_handle}/workspace/{workspace_handle} | Update user workspace
*Users* | [**Create**](docs/Users.md#create) | **Post** /user | Create user
*Users* | [**CreateAvatar**](docs/Users.md#createavatar) | **Post** /user/{user_handle}/avatar | Create user avatar
*Users* | [**CreateDBPassword**](docs/Users.md#createdbpassword) | **Post** /user/{user_handle}/password | Create user password
*Users* | [**Delete**](docs/Users.md#delete) | **Delete** /user/{user_handle} | Delete user
*Users* | [**DeleteAvatar**](docs/Users.md#deleteavatar) | **Delete** /user/{user_handle}/avatar | Delete user avatar
*Users* | [**Get**](docs/Users.md#get) | **Get** /user/{user_handle} | Get user
*Users* | [**GetDBPassword**](docs/Users.md#getdbpassword) | **Get** /user/{user_handle}/password | Get user password
*Users* | [**GetEmail**](docs/Users.md#getemail) | **Get** /user/{user_handle}/email/{email_id} | Get user email
*Users* | [**GetPreferences**](docs/Users.md#getpreferences) | **Get** /user/{user_handle}/preferences | Get user preferences
*Users* | [**GetQuota**](docs/Users.md#getquota) | **Get** /user/{user_handle}/quota | User quota
*Users* | [**List**](docs/Users.md#list) | **Get** /user | List users
*Users* | [**ListAuditLogs**](docs/Users.md#listauditlogs) | **Get** /user/{user_handle}/audit_log | User audit logs
*Users* | [**ListConstraints**](docs/Users.md#listconstraints) | **Get** /user/{user_handle}/constraint | List user constraints
*Users* | [**ListEmails**](docs/Users.md#listemails) | **Get** /user/{user_handle}/email | List user emails
*Users* | [**ListFeatures**](docs/Users.md#listfeatures) | **Get** /user/{user_handle}/feature | User features
*Users* | [**RequestConstraintOverride**](docs/Users.md#requestconstraintoverride) | **Patch** /user/{user_handle}/constraint | Request user constraint override
*Users* | [**Update**](docs/Users.md#update) | **Patch** /user/{user_handle} | Update user
*Users* | [**UpdatePreferences**](docs/Users.md#updatepreferences) | **Patch** /user/{user_handle}/preferences | Update user preferences


## Documentation For Models

 - [ActorWorkspace](docs/ActorWorkspace.md)
 - [AttachWorkspaceSchemaRequest](docs/AttachWorkspaceSchemaRequest.md)
 - [AuditRecord](docs/AuditRecord.md)
 - [BillingInfo](docs/BillingInfo.md)
 - [Connection](docs/Connection.md)
 - [ConnectionTestResult](docs/ConnectionTestResult.md)
 - [Constraint](docs/Constraint.md)
 - [ConstraintOverrideRequest](docs/ConstraintOverrideRequest.md)
 - [CreateConnectionRequest](docs/CreateConnectionRequest.md)
 - [CreateOrgAvatarResponse](docs/CreateOrgAvatarResponse.md)
 - [CreateOrgRequest](docs/CreateOrgRequest.md)
 - [CreateOrgWorkspaceUserRequest](docs/CreateOrgWorkspaceUserRequest.md)
 - [CreatePipelineRequest](docs/CreatePipelineRequest.md)
 - [CreateUserAvatarResponse](docs/CreateUserAvatarResponse.md)
 - [CreateUserPasswordRequest](docs/CreateUserPasswordRequest.md)
 - [CreateUserRequest](docs/CreateUserRequest.md)
 - [CreateWorkspaceAggregatorRequest](docs/CreateWorkspaceAggregatorRequest.md)
 - [CreateWorkspaceConnRequest](docs/CreateWorkspaceConnRequest.md)
 - [CreateWorkspaceModRequest](docs/CreateWorkspaceModRequest.md)
 - [CreateWorkspaceModVariableSettingRequest](docs/CreateWorkspaceModVariableSettingRequest.md)
 - [CreateWorkspaceRequest](docs/CreateWorkspaceRequest.md)
 - [CreateWorkspaceSnapshotRequest](docs/CreateWorkspaceSnapshotRequest.md)
 - [Datatank](docs/Datatank.md)
 - [DeleteOrgAvatarResponse](docs/DeleteOrgAvatarResponse.md)
 - [DeleteUserAvatarResponse](docs/DeleteUserAvatarResponse.md)
 - [ErrorDetailModel](docs/ErrorDetailModel.md)
 - [ErrorModel](docs/ErrorModel.md)
 - [Feature](docs/Feature.md)
 - [Identity](docs/Identity.md)
 - [InviteOrgUserRequest](docs/InviteOrgUserRequest.md)
 - [JSONTime](docs/JSONTime.md)
 - [ListActorWorkspacesResponse](docs/ListActorWorkspacesResponse.md)
 - [ListAuditLogsResponse](docs/ListAuditLogsResponse.md)
 - [ListConnectionsResponse](docs/ListConnectionsResponse.md)
 - [ListConstraintsResponse](docs/ListConstraintsResponse.md)
 - [ListFeaturesResponse](docs/ListFeaturesResponse.md)
 - [ListIdentitiesResponse](docs/ListIdentitiesResponse.md)
 - [ListLogsResponse](docs/ListLogsResponse.md)
 - [ListOrgUsersResponse](docs/ListOrgUsersResponse.md)
 - [ListOrgWorkspaceUsersResponse](docs/ListOrgWorkspaceUsersResponse.md)
 - [ListOrgsResponse](docs/ListOrgsResponse.md)
 - [ListPipelinesResponse](docs/ListPipelinesResponse.md)
 - [ListProcessesResponse](docs/ListProcessesResponse.md)
 - [ListTokensResponse](docs/ListTokensResponse.md)
 - [ListUserEmailsResponse](docs/ListUserEmailsResponse.md)
 - [ListUserOrgsResponse](docs/ListUserOrgsResponse.md)
 - [ListUsersResponse](docs/ListUsersResponse.md)
 - [ListWorkspaceAggregatorsResponse](docs/ListWorkspaceAggregatorsResponse.md)
 - [ListWorkspaceConnResponse](docs/ListWorkspaceConnResponse.md)
 - [ListWorkspaceConnectionAssociationsResponse](docs/ListWorkspaceConnectionAssociationsResponse.md)
 - [ListWorkspaceConnectionsResponse](docs/ListWorkspaceConnectionsResponse.md)
 - [ListWorkspaceModVariablesResponse](docs/ListWorkspaceModVariablesResponse.md)
 - [ListWorkspaceModsResponse](docs/ListWorkspaceModsResponse.md)
 - [ListWorkspaceSchemaResponse](docs/ListWorkspaceSchemaResponse.md)
 - [ListWorkspaceSchemaTableResponse](docs/ListWorkspaceSchemaTableResponse.md)
 - [ListWorkspaceSnapshotsResponse](docs/ListWorkspaceSnapshotsResponse.md)
 - [ListWorkspacesResponse](docs/ListWorkspacesResponse.md)
 - [LogRecord](docs/LogRecord.md)
 - [Org](docs/Org.md)
 - [OrgQuota](docs/OrgQuota.md)
 - [OrgUser](docs/OrgUser.md)
 - [OrgWorkspaceUser](docs/OrgWorkspaceUser.md)
 - [Pipeline](docs/Pipeline.md)
 - [PipelineCommandRequest](docs/PipelineCommandRequest.md)
 - [PipelineCommandResponse](docs/PipelineCommandResponse.md)
 - [PipelineFrequency](docs/PipelineFrequency.md)
 - [Quota](docs/Quota.md)
 - [SpProcess](docs/SpProcess.md)
 - [TemporaryTokenRequest](docs/TemporaryTokenRequest.md)
 - [TestConnectionRequest](docs/TestConnectionRequest.md)
 - [Token](docs/Token.md)
 - [UpdateConnectionRequest](docs/UpdateConnectionRequest.md)
 - [UpdateOrgRequest](docs/UpdateOrgRequest.md)
 - [UpdateOrgUserRequest](docs/UpdateOrgUserRequest.md)
 - [UpdateOrgWorkspaceUserRequest](docs/UpdateOrgWorkspaceUserRequest.md)
 - [UpdatePipelineRequest](docs/UpdatePipelineRequest.md)
 - [UpdateTemporaryTokenRequest](docs/UpdateTemporaryTokenRequest.md)
 - [UpdateTokenRequest](docs/UpdateTokenRequest.md)
 - [UpdateUserPreferencesRequest](docs/UpdateUserPreferencesRequest.md)
 - [UpdateUserRequest](docs/UpdateUserRequest.md)
 - [UpdateWorkspaceAggregatorRequest](docs/UpdateWorkspaceAggregatorRequest.md)
 - [UpdateWorkspaceModRequest](docs/UpdateWorkspaceModRequest.md)
 - [UpdateWorkspaceModVariableSettingRequest](docs/UpdateWorkspaceModVariableSettingRequest.md)
 - [UpdateWorkspaceRequest](docs/UpdateWorkspaceRequest.md)
 - [UpdateWorkspaceSnapshotRequest](docs/UpdateWorkspaceSnapshotRequest.md)
 - [User](docs/User.md)
 - [UserDatabasePassword](docs/UserDatabasePassword.md)
 - [UserEmail](docs/UserEmail.md)
 - [UserLoginRequest](docs/UserLoginRequest.md)
 - [UserOrg](docs/UserOrg.md)
 - [UserPreferences](docs/UserPreferences.md)
 - [UserQuota](docs/UserQuota.md)
 - [UserSignupRequest](docs/UserSignupRequest.md)
 - [Workspace](docs/Workspace.md)
 - [WorkspaceAggregator](docs/WorkspaceAggregator.md)
 - [WorkspaceCommandRequest](docs/WorkspaceCommandRequest.md)
 - [WorkspaceCommandResponse](docs/WorkspaceCommandResponse.md)
 - [WorkspaceConn](docs/WorkspaceConn.md)
 - [WorkspaceConnection](docs/WorkspaceConnection.md)
 - [WorkspaceConnectionAssociation](docs/WorkspaceConnectionAssociation.md)
 - [WorkspaceMod](docs/WorkspaceMod.md)
 - [WorkspaceModVariable](docs/WorkspaceModVariable.md)
 - [WorkspaceQueryResult](docs/WorkspaceQueryResult.md)
 - [WorkspaceQueryResultColumn](docs/WorkspaceQueryResultColumn.md)
 - [WorkspaceSchema](docs/WorkspaceSchema.md)
 - [WorkspaceSchemaTable](docs/WorkspaceSchemaTable.md)
 - [WorkspaceSchemaTableColumn](docs/WorkspaceSchemaTableColumn.md)
 - [WorkspaceSnapshot](docs/WorkspaceSnapshot.md)
 - [WorkspaceSnapshotData](docs/WorkspaceSnapshotData.md)
 - [WorkspaceSnapshotDataLayout](docs/WorkspaceSnapshotDataLayout.md)


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

