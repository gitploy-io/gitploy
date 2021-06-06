export default interface Commit {
    sha: string
    message: string
    isPullRequest: boolean
}