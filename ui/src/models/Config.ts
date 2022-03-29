export default interface Config  {
    envs: Env[]
}

export interface Env {
    name: string
    requiredContexts?: string[]
    dynamicPayload?: DynamicPayload
    review?: {
        enabled: boolean
        reviewers: string[]
    }
}

export interface DynamicPayload {
    enabled: boolean
    inputs: Map<string, DynamicPayloadInput>
}

export interface DynamicPayloadInput {
    type: DynamicPayloadInputTypeEnum
    required: boolean
    default: any
    description: string
    options: string[]
}

export enum DynamicPayloadInputTypeEnum {
    Select = "select",
    String = "string",
    Number = "number",
    Boolean = "boolean"
}
