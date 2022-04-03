import { Row, Col, Result, Button} from "antd"

import React from "react"

export interface ContentProps {
    available: boolean 
    authorized: boolean
    expired: boolean
    onClickRetry(): void
}

export default function Content({
    available,
    authorized,
    expired,
    children,
    onClickRetry,
}: React.PropsWithChildren<ContentProps>): JSX.Element {


    let content: React.ReactNode
    if (!available) {
        content = <Result
            style={{paddingTop: '120px'}}
            status="error"
            title="Server Internal Error"
            subTitle="Sorry, something went wrong. Contact administrator."
            extra={[<Button key="console" type="primary" onClick={onClickRetry}>Retry</Button>]}
        />
    } else if (!authorized) {
        content = <Result
            style={{paddingTop: '120px'}}
            status="warning"
            title="Unauthorized Error"
            subTitle="Sorry, you are not authorized."
            extra={[<Button key="console" type="primary" href="/">Sign in</Button>]}
        />
    } else if (expired) {
        content = <Result
            style={{paddingTop: '120px'}}
            status="warning"
            title="License Expired"
            subTitle="Sorry, the license is expired."
            extra={[<Button key="console" type="primary" onClick={onClickRetry}>Retry</Button>]}
        />
    } else {
        content = children
    }

    return (
        <Row>
            <Col 
                span={22}
                offset={1}
                md={{span: 14, offset: 5}} 
                lg={{span: 12, offset: 6}} 
                xxl={{span: 10, offset: 7}}
            >
                {content}
            </Col>
    </Row>
    )
}