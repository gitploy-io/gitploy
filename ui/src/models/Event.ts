import Deployment from "./Deployment"
import Approval from "./Approval"
import { Review } from "./Review"

export default interface Event {
    id: number
    kind: EventKindEnum
    type: EventTypeEnum
    deployment?: Deployment
    approval?: Approval
    review?: Review
    deletedId: number 
}

export enum EventKindEnum {
    Deployment = "deployment",
    Approval = "approval",
    Review = "review"
}

export enum EventTypeEnum {
    Created = "created",
    Updated = "updated",
    Deleted = "deleted"
}
