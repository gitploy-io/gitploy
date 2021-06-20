import { Button } from "antd"
import { RedoOutlined } from "@ant-design/icons"

export interface SyncButtonProps {
    loading: boolean
    onClickSync(): void
}

export default function SyncButton(props: SyncButtonProps) {
    return (
        <Button
            loading={props.loading}
            icon={<RedoOutlined />}
            onClick={props.onClickSync}>
            Sync
        </Button>
    )
}