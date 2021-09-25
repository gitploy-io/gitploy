import { List, Button } from "antd"
import { LockOutlined, UnlockOutlined } from "@ant-design/icons"
import moment from 'moment'

import { Env, Lock } from "../models"
import UserAvatar from './UserAvatar'

interface LockListProps {
    envs: Env[]
    locks: Lock[]
    onClickLock(env: string): void
    onClickUnlock(env: string): void
}

export default function LockList(props: LockListProps): JSX.Element {
    return (
        <List
            dataSource={props.envs}
            renderItem={(env) => {
                const lock = props.locks.find((lock) => lock.env === env.name)

                return (lock)? 
                    <List.Item>
                        <List.Item.Meta 
                            title={<span>{env.name.toUpperCase()} <LockOutlined /></span>}
                            description={<LockDescription lock={lock}/>}
                        />
                        <Button danger
                            onClick={() => {props.onClickUnlock(env.name)}}
                        >
                            UNLOCK
                        </Button>
                    </List.Item>:
                    <List.Item>
                        <List.Item.Meta 
                            title={<span>{env.name.toUpperCase()} <UnlockOutlined /></span>}
                            description={<LockDescription lock={lock}/>}
                        />
                        <Button danger
                            type="primary"
                            onClick={() => {props.onClickLock(env.name)}}
                        >
                            LOCK
                        </Button>
                    </List.Item>
            }}
        />
    )
}

interface LockDescriptionProps {
    lock?: Lock
}

function LockDescription(props: LockDescriptionProps) {

    return (
        (props.lock)?
            <span style={{color: "black"}}>Locked by <UserAvatar user={props.lock.user} />  {moment(props.lock.createdAt).fromNow()}</span>:
            <span></span>
    )
}