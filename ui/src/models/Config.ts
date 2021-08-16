export default interface Config  {
    envs: Env[]
}

export interface Env {
    name: string
    requiredContexts?: string[]
    approval?: EnvApproval
}

export interface EnvApproval {
    enabled: boolean
    required_count: number
}