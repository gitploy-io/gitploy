import { StatusCodes } from "http-status-codes"
import { HttpInternalServerError, HttpUnauthorizedError } from "../models/errors"

export const _fetch = async (input: RequestInfo, init?: RequestInit): Promise<Response>  => {
    const response = await fetch(input, init)
    
    // Throw exception when the general status code is received.
    if (response.status === StatusCodes.INTERNAL_SERVER_ERROR) {
        throw new HttpInternalServerError("The internal server error occurs.")
    } else if (response.status === StatusCodes.UNAUTHORIZED) {
        throw new HttpUnauthorizedError("The session is expired.")
    } 

    return response
}