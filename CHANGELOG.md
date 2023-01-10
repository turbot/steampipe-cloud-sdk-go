# Steampipe Cloud Go SDK

## 0.5.0 [tbd]

_What's new?_

- Identity Pipeline APIs
- Identity Workspace Pipeline APIs

_Breaking changes_

- Modified attribute `Layout` of the WorkspaceSnapshotData struct to have defined type `WorkspaceSnapshotDataLayout` instead of `map[string]interface{}`

## 0.4.0 [2023-01-09]

_Breaking changes_

- Removed attribute `id` in the ErrorModel struct.
- Renamed attribute `errors` to `validation_errors` in the ErrorModel struct.
- Attribute `instance` in the ErrorModel struct will now return the instance of the steampipe cloud error instead of the `id` attribute where it was previously accessed.

## 0.3.0 [2022-12-06]

- Remove SearchPath attribute from WorkspaceSnapshotData model
- Add ExpiresAt attribute to WorkspaceSnapshot model

## 0.2.0 [2022-11-03]

- User Preferences APIs
- User Emails APIs
- Steampipe Login APIs
- Add option to set snapshot titles

## 0.1.3 [2022-08-12]

- Add Snapshot APIs.
- Remove email from all responses where users are returned.
- Various enhancements related to APIs returning redundant data.

## 0.1.2 [2022-07-19]

- Allow users to be added to workspaces directly, bypassing the invite process.
- Add search capability for org members.

## 0.1.1 [2022-07-13]

- Move Orgs for a user under actor orgs.
- Modify response structure for workspaces of an actor.
- Org and Workspace members to be listed via a unified method rather than separate methods for invites and members.

## 0.1.0 [2022-07-01]

- Add Org Workspace Member APIs.
- Modify various existing APIs to include metadata information such as created_by etc. 

## 0.0.4 [2022-05-31]

- Add Mod, Mod Variables APIs

## 0.0.3 [2021-12-15]

- Update the types docs.
- Update the description of input parameters for API.

## 0.0.2 [2021-12-14]

- Updated: Removed the suffix types from the struct.
- Updated: Removed the suffix sperr from the error struct.
- Fixed: Naming of workspace data APIs.

## 0.0.1 [2021-12-13]

- Initial version
