import { instance, headers } from './settings'
import { Tag } from '../models'

export const listTags = async (repoId: string, page: number = 1, perPage: number = 30) => {
    const tags: Tag[] = await fetch(`${instance}/v1/repos/${repoId}/tags?page=${page}&per_page=${perPage}`, {
        headers,
        credentials: "same-origin",
    })
        .then(response => response.json())
        .then(tags => tags.map((t: any): Tag => {
            return {
                name: t.name,
                commitSha: t.commit_sha,
            } 
        }))
    
    return tags
}