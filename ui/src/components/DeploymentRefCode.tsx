import { Typography } from "antd"

import { Deployment, DeploymentType } from "../models"

const { Text } = Typography

interface DeploymentRefCodeProps {
    deployment: Deployment
}

export default function DeploymentRefCode(props: DeploymentRefCodeProps): JSX.Element {
    let ref: string
    if (props.deployment.type === DeploymentType.Commit) {
        ref = props.deployment.ref.substr(0, 7)
    } else if (props.deployment.type === DeploymentType.Branch && props.deployment.sha !== "") {
        ref = `${props.deployment.ref}(${props.deployment.sha.substr(0, 7)})`
    } else {
        ref = props.deployment.ref
    }

    return <Text className="gitploy-code" code>
        {ref}
    </Text>
}