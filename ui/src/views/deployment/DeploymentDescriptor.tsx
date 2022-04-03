import { useState } from "react"
import { Button, Descriptions, Modal, Typography } from "antd"
import moment from "moment"

import CommitChanges from "./CommitChanges"
import DeploymentStatusBadge from "../../components/DeploymentStatusBadge"
import UserAvatar from "../../components/UserAvatar"

import { Commit, Deployment } from "../../models"
import { getShortRef } from "../../libs"

const { Text } = Typography

export interface DeploymentDescriptorProps {
    deployment: Deployment
    changes: Commit[]
}

export default function DeploymentDescriptor({
    deployment, 
    changes
}: DeploymentDescriptorProps): JSX.Element {
    const [visible, setVisible] = useState<boolean>(false)

    const showModal = () => {
      setVisible(true)
    }
  
    const hideModal = () => {
      setVisible(false)
    }

    return (
        <Descriptions title="Information" bordered column={1}>
            <Descriptions.Item label="Environment">
                {deployment.env}
            </Descriptions.Item>
            <Descriptions.Item label="Ref">
                <Text className="gitploy-code" code>{getShortRef(deployment)}</Text>
            </Descriptions.Item>
            <Descriptions.Item label="Status">
                <DeploymentStatusBadge deployment={deployment}/>
            </Descriptions.Item>
            <Descriptions.Item label="Deployer">
                <UserAvatar user={deployment.deployer} boldName={false}/>
            </Descriptions.Item>
            <Descriptions.Item label="Deploy Time">
                {moment(deployment.createdAt).format("YYYY-MM-DD HH:mm:ss")}
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
                    <CommitChanges changes={changes}/>
                </Modal>
            </Descriptions.Item>
        </Descriptions>
    )
}