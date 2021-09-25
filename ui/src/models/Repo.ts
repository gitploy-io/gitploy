import Deployment from "./Deployment"

export default interface Repo { 
    id: string
    namespace: string
    name: string
    description: string
    configPath: string
    active: boolean
    webhookId: number
    createdAt: Date
    updatedAt: Date 
    deployments?: Deployment[]
}
