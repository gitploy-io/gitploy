import Deployment from "./Deployment"
import Approval from "./Approval"

export default interface Event {
    id: number
    kind: EventKindEnum
    type: EventTypeEnum
    deployment?: Deployment
    approval?: Approval

    // This field identifies which approval was deleted.
    approvalId?: number 
}

export enum EventKindEnum {
    Deployment = "deployment",
    Approval = "approval"
}

export enum EventTypeEnum {
    Created = "created",
    Updated = "updated",
    Deleted = "deleted"
}