import Deployment from "./Deployment"

export default interface Repo { 
    id: string
    namespace: string
    name: string
    description: string
    configPath: string
    active: boolean
    webhookId: number
    syncedAt: Date
    createdAt: Date
    updatedAt: Date 
    deployments?: Deployment[]
}

export interface RepoPayload {
    configPath: string
}
