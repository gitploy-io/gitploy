import { Dropdown, Menu, Button } from "antd"

interface ApprovalDropdownProps {
    onClickApprove(): void
    onClickDecline(): void
}

export default function ApprovalDropdown(props: ApprovalDropdownProps) {
    return (
        <Dropdown overlay={<Menu>
                <Menu.Item key="0">
                    <Button type="text" onClick={props.onClickApprove}>Approve</Button>
                </Menu.Item>
                <Menu.Item key="1">
                    <Button type="text" onClick={props.onClickDecline}>Decline</Button>
                </Menu.Item>
            </Menu>
        }>
            <Button >Approval</Button>
        </Dropdown>
    )

}