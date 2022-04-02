import { shallowEqual } from "react-redux"
import { Button, Descriptions } from "antd"

import { useAppSelector } from "../../redux/hooks"

export default function SlackDescriptions(): JSX.Element {
    const { user } = useAppSelector(state => state.settings, shallowEqual)

    const connected = (user?.chatUser)? true : false

    return (
        <Descriptions title="Slack">
            <Descriptions.Item>
                {(connected)? 
                    <Button href="/slack/signout" type="primary" danger>DISCONNECTED</Button>:
                    <Button href="/slack/" type="primary">CONNECT</Button>}
            </Descriptions.Item>
        </Descriptions>
    )
}