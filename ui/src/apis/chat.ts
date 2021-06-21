import { _fetch } from "./_base"
import { instance, headers } from "./settings"
import { StatusCodes } from "http-status-codes"

export const checkSlack = async () => {
    const res = await _fetch(`${instance}/slack/check`, {
        headers,
        credentials: "same-origin",
    })
    if (res.status === StatusCodes.NOT_FOUND) {
        return false
    }

    return true
}