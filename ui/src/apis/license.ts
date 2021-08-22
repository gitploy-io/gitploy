import { instance, headers } from './setting'
import { _fetch } from "./_base"

import { License } from "../models"

interface LicenseData {
    kind: string
    member_count: number
    member_limit: number
    expired_at: string
}

function mapDataToLicense(data: LicenseData): License {
    return {
        kind: data.kind,
        memberCount: data.member_count,
        memberLimit: data.member_limit,
        expiredAt: new Date(data.expired_at),
    }
}

export const getLicense = async (): Promise<License> => {
    const lic = await _fetch(`${instance}/api/v1/license`, {
        headers,
        credentials: 'same-origin',
    })
        .then(res => res.json())
        .then(data => mapDataToLicense(data))

    return lic
}