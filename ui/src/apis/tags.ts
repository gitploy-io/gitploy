import { StatusCodes } from 'http-status-codes'

import { instance, headers } from './settings'
import { _fetch } from "./_base"
import { Tag, HttpNotFoundError } from '../models'

export const listTags = async (repoId: string, page: number = 1, perPage: number = 30) => {
    const tags: Tag[] = await _fetch(`${instance}/api/v1/repos/${repoId}/tags?page=${page}&per_page=${perPage}`, {
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

export const getTag = async (repoId: string, name: string) => {
    const response = await _fetch(`${instance}/api/v1/repos/${repoId}/tags/${name}`, {
        headers,
        credentials: "same-origin",
    })
    if (response.status === StatusCodes.NOT_FOUND) {
        const message = await response.json().then(data => data.message)
        throw new HttpNotFoundError(message)
    }

    const tag:Tag = await response
            .json()
            .then(t => ({
                name: t.name,
                commitSha: t.commit_sha,
        }))
    return tag
}