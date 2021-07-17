import User from "./User"
import Repo from "./Repo"

export default interface Perm {
    repoPerm: string
    syncedAt: Date
    createdAt: Date
    updatedAt: Date
    user: User
    repo: Repo
}
