import { CheckCircleTwoTone, CloseCircleTwoTone, ClockCircleTwoTone } from '@ant-design/icons'

import { StatusState } from '../models'

interface StatusStateIconProps {
	state: StatusState
}

export default function StatusStateIcon(props: StatusStateIconProps) {
	const { state } = props
	if (state === StatusState.Pending) {
		return <ClockCircleTwoTone />
	} else if (state === StatusState.Success) {
		return <CheckCircleTwoTone />
	} else if (state === StatusState.Failure) {
		return <CloseCircleTwoTone />
	}
	return <ClockCircleTwoTone />
}