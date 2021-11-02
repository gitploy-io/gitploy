import { Dropdown, Menu, Button } from "antd"

interface ReviewDropdownProps {
    onClickApprove(): void
    onClickReject(): void
}

export default function ReviewDropdown(props: ReviewDropdownProps): JSX.Element {
    return (
        <Dropdown overlay={
            <Menu>
                <Menu.Item key="0">
                    <Button type="text" onClick={props.onClickApprove}>Approve</Button>
                </Menu.Item>
                <Menu.Item key="1">
                    <Button type="text" onClick={props.onClickReject}>Reject</Button>
                </Menu.Item>
            </Menu>
        }>
            <Button >Review</Button>
        </Dropdown>
    )

}