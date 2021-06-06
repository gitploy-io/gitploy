import { StatusCodes } from 'http-status-codes'

export class HttpRequestError extends Error {
    constructor(public code: number, public m: string) {
        super(m)

        Object.setPrototypeOf(this, HttpRequestError.prototype)
    }
}

export class HttpNotFoundError extends HttpRequestError {
    constructor(public m: string) {
        super(StatusCodes.NOT_FOUND, m)

        Object.setPrototypeOf(this, HttpNotFoundError.prototype)
    }
}