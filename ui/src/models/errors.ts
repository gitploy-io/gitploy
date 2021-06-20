import { StatusCodes } from 'http-status-codes'

export class HttpRequestError extends Error {
    constructor(public code: number, public m: string) {
        super(m)

        Object.setPrototypeOf(this, HttpRequestError.prototype)
    }
}

export class HttpInternalServerError extends HttpRequestError {
    constructor(public m: string) {
        super(StatusCodes.INTERNAL_SERVER_ERROR, m)

        Object.setPrototypeOf(this, HttpInternalServerError.prototype)
    }
}

export class HttpUnauthorizedError extends HttpRequestError {
    constructor(public m: string) {
        super(StatusCodes.UNAUTHORIZED, m)

        Object.setPrototypeOf(this, HttpUnauthorizedError.prototype)
    }
}

export class HttpForbiddenError extends HttpRequestError {
    constructor(public m: string) {
        super(StatusCodes.FORBIDDEN, m)

        Object.setPrototypeOf(this, HttpForbiddenError.prototype)
    }
}

export class HttpNotFoundError extends HttpRequestError {
    constructor(public m: string) {
        super(StatusCodes.NOT_FOUND, m)

        Object.setPrototypeOf(this, HttpNotFoundError.prototype)
    }
}