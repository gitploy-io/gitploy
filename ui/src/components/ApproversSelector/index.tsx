import { Typography } from "antd"

import ApproversSearch, { ApproversSelectProps } from "./ApproversSearch"
import ApprovalList, { ApprovalListProps } from "./ApprovalList"

const { Text } = Typography

interface ApproversSelectorProps extends ApproversSelectProps, ApprovalListProps {
}

export default function ApproversSelector(props: ApproversSelectorProps): JSX.Element {
    return (
        <div>
            <div style={{paddingLeft: "5px"}}>
                <Text strong>Approvers</Text>
            </div>
            <div style={{marginTop: "5px"}}>
                <ApproversSearch
                    style={{width: "100%"}}
                    value="Select Approvers"
                    approvers={props.approvers}
                    candidates={props.candidates}
                    onSearchCandidates={props.onSearchCandidates}
                    onSelectCandidate={props.onSelectCandidate} />
            </div>
            <div style={{marginTop: "10px", paddingLeft: "5px"}}>
                {(props.approvals.length !== 0) ?
                    <ApprovalList approvals={props.approvals}/>:
                    <Text type="secondary"> No approvers </Text>}
            </div>
        </div>
    )
}