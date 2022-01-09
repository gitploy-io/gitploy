import { Deployment, DeploymentType } from "../models"

/**
 * The function returns the short-formatted ref string.
 * @param deployment 
 * @returns 
 */
export const getShortRef = (deployment: Deployment): string => {
    return deployment.type === DeploymentType.Commit? deployment.ref.substring(0, 7) : deployment.ref
}