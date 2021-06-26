import { instance, headers } from "./setting"
import { _fetch } from "./_base"

export const sync = async () => {
    await _fetch(`${instance}/api/v1/sync`, {
        headers,
        credentials: 'same-origin',
        method: "POST",
    })
}