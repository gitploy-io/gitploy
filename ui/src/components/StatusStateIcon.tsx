import { Popover, Avatar, Typography, Row, Col, Divider } from "antd"
import { CheckOutlined, CloseOutlined } from "@ant-design/icons"

import { Status, StatusState } from "../models"

const colorSuccess = "#1a7f37"
const colorFailure = "#cf222e"

const { Text, Link } = Typography

interface StatusStateIconProps {
	statuses: Status[]
}

export default function StatusStateIcon(props: StatusStateIconProps): JSX.Element {
	const states = props.statuses.map((status) => status.state)
	const content: JSX.Element = <div style={{width: "350px"}}>
		{props.statuses.map((status, idx) => {
			return <Row key={idx}>
				<Col span="12">
					{mapStateToIcon(status.state)}&nbsp;&nbsp;
					<Avatar size="small" src={status.avatarUrl} shape="square"/>&nbsp;&nbsp;
					<Text strong>{status.context}</Text>
				</Col>
				<Col span="11" style={{textAlign: "right"}}>
					<Link href={status.targetUrl} target="_blank">Details</Link>
				</Col>
				{(idx !== props.statuses.length - 1)? 
					<Divider style={{margin: "5px 0px"}}/> : null}
			</Row>
		})}
	</div>

	return (
		<Popover title="Statuses" content={content}>
			{mapStateToIcon(mergeStatusStates(states))}
		</Popover>
	)
}

function mapStateToIcon(state: StatusState): JSX.Element {
	switch (state) {
		case StatusState.Null:
			return <span></span>
		case StatusState.Pending:
			return <span>
				<span className="gitploy-pending-icon"></span>&nbsp;&nbsp;
			</span>
		case StatusState.Success:
			return <CheckOutlined style={{color: colorSuccess}}/>
		case StatusState.Failure:
			return <CloseOutlined style={{color: colorFailure}}/>
		default:
			return <span>
				<span className="gitploy-pending-icon"></span>&nbsp;&nbsp;
			</span>
	}
}

function mergeStatusStates(states: StatusState[]): StatusState {
	if (states.length === 0) {
		return StatusState.Null
	}

	// The state is failure if one of them is failure.
	for (let idx = 0; idx < states.length; idx++) {
		if (states[idx] === StatusState.Failure) {
			return StatusState.Failure
		}
	}

	for (let idx = 0; idx < states.length; idx++) {
		if (states[idx] === StatusState.Pending) {
			return StatusState.Pending
		}
	}

	return StatusState.Success
}