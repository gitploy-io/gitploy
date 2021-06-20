import { instance, headers } from "./settings"

export const sync = async () => {
    await fetch(`${instance}/api/v1/sync`, {
        headers,
        credentials: 'same-origin',
        method: "POST",
    })
}