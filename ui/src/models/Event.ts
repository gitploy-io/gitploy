import Deployment from "./Deployment"
import { Review } from "./Review"

export default interface Event {
    id: number
    kind: EventKindEnum
    type: EventTypeEnum
    deployment?: Deployment
    review?: Review
    deletedId: number 
}

export enum EventKindEnum {
    Deployment = "deployment",
    Review = "review"
}

export enum EventTypeEnum {
    Created = "created",
    Updated = "updated",
    Deleted = "deleted"
}
