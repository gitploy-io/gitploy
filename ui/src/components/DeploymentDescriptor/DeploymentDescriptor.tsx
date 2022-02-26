import { Button, Descriptions, Modal, Typography } from "antd"
import moment from "moment"

import DeploymentStatusBadge from "../DeploymentStatusBadge"
import UserAvatar from "../UserAvatar"
import CommitChanges from "./CommitChanges"

import { Commit, Deployment } from "../../models"
import { getShortRef } from "../../libs"
import { useState } from "react"

const { Text } = Typography

interface DeploymentDescriptorProps {
    deployment: Deployment
    commits: Commit[]
}

export default function DeploymentDescriptor(props: DeploymentDescriptorProps): JSX.Element {
    const [visible, setVisible] = useState<boolean>(false)

    const showModal = () => {
      setVisible(true)
    }
  
    const hideModal = () => {
      setVisible(false)
    }

    return (
        <Descriptions title="Information" bordered column={1}>
            <Descriptions.Item label="Environment">{props.deployment.env}</Descriptions.Item>
            <Descriptions.Item label="Ref">
                <Text className="gitploy-code" code>{getShortRef(props.deployment)}</Text>
            </Descriptions.Item>
            <Descriptions.Item label="Status">
                <DeploymentStatusBadge deployment={props.deployment}/>
            </Descriptions.Item>
            <Descriptions.Item label="Deployer">
                <UserAvatar user={props.deployment.deployer} boldName={false}/>
            </Descriptions.Item>
            <Descriptions.Item label="Deploy Time">
                {moment(props.deployment.createdAt).format("YYYY-MM-DD HH:mm:ss")}
            </Descriptions.Item>
            <Descriptions.Item label="Changes">
                <Button 
                    type="link"
                    style={{padding: 0, height: 20}}
                    onClick={showModal}
                >
                    View 
                </Button>
                <Modal 
                    title="Changes" 
                    visible={visible}
                    width={800}
                    // Hide OK Button
                    okButtonProps={{style: {display: "none"}}}
                    cancelText="Close"
                    onCancel={hideModal}
                >
                    <CommitChanges changes={props.commits}/>
                </Modal>
            </Descriptions.Item>
        </Descriptions>
    )
}