import { StatusCodes } from 'http-status-codes'

export class HttpRequestError extends Error {
    constructor(public code: number, m: string) {
        super(m)

        Object.setPrototypeOf(this, HttpRequestError.prototype)
    }
}

export class NotFoundError extends HttpRequestError {
    constructor(m: string) {
        super(StatusCodes.NOT_FOUND, m)

        Object.setPrototypeOf(this, HttpRequestError.prototype)
    }
}