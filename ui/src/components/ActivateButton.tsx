import { Result, Button } from "antd"

export interface ActivateButtonProps {
    onClickActivate(): void
}

export default function ActivateButton(props: ActivateButtonProps) {
    return (
        <Result
            title="Activate your repository"
            extra={[
                <Button
                    key={0}
                    onClick={props.onClickActivate}
                    type="primary"
                    size="large">
                    ACTIVATE
                </Button>,
            ]}
        />
    )
}