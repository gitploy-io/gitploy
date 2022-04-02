import { shallowEqual } from "react-redux"
import { Row, Col, Result, Button} from "antd"

import { useAppSelector, useAppDispatch } from "../../redux/hooks"
import { mainSlice as slice } from "../../redux/main"
import React from "react"

export default function Content(props: React.PropsWithChildren<any>): JSX.Element {
    const { 
        available, 
        authorized, 
        expired,
    } = useAppSelector(state => state.main, shallowEqual)
    const dispatch = useAppDispatch()

    const onClickRetry = () => {
        dispatch(slice.actions.setAvailable(true))
        dispatch(slice.actions.setExpired(false))
    }

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
        content = props.children
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