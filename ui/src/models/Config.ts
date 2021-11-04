export default interface Config  {
    envs: Env[]
}

export interface Env {
    name: string
    requiredContexts?: string[]
    review?: {
        enabled: boolean
        reviewers: string[]
    }
}
