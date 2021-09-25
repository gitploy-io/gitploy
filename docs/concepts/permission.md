# Permission

Gitploy supports fine-grained access control for the repository. The user has explicit permissions(read, write, and admin) to the repository.

## Synchronization
The permission of the repository is determined while Gitploy synchronizes with Github. *If the access permission has changed, you should re-synchronize it in Gitploy again.*

## Capabilities

Here are capabilities for each permission: 

* **Read** - Users can read all activities that happened in the repository, such as deployments, approvals. And users are also capable of responding to the approval.

* **Write** - Users can lock, deploy, and rollback for the `REF`. 

* **Admin** - Users can configures the repository, such as activating.

Of course, write and admin permission cover the ability of read permission.