export default interface Config  {
    envs: Env[]
}

export interface Env {
    name: string
    requiredContexts: string[]
    approvalEnabled: boolean
}