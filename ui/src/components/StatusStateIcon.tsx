import { CheckCircleTwoTone, CloseCircleTwoTone, ClockCircleTwoTone } from '@ant-design/icons'

import { StatusState } from '../models'

interface StatusStateIconProps {
	state: StatusState
}

export default function StatusStateIcon(props: StatusStateIconProps): JSX.Element {
	const { state } = props
	if (state === StatusState.Pending) {
		return <ClockCircleTwoTone twoToneColor="#d9d9d9"/>
	} else if (state === StatusState.Success) {
		return <CheckCircleTwoTone twoToneColor="#52c41a"/>
	} else if (state === StatusState.Failure) {
		return <CloseCircleTwoTone twoToneColor="red"/>
	}
	return <span></span>
}