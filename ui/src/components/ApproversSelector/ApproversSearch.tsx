import React from "react"
import { Select, SelectProps, Avatar, Spin } from "antd"
import { CheckOutlined } from "@ant-design/icons"
import debounce from "lodash.debounce"

import { User } from "../../models"

export interface ApproversSelectProps extends SelectProps<string>{
    approvers: User[]
    candidates: User[]
    // The type of parameter have to be string
    // because candidates could be dynamically changed with search.
    onSelectCandidate(id: string): void
    onSearchCandidates(login: string): void
}

export default function ApproversSelect(props: ApproversSelectProps): JSX.Element {
    const [ searching, setSearching ] = React.useState<boolean>(false)

    // Clone Select props only
    // https://stackoverflow.com/questions/34698905/how-can-i-clone-a-javascript-object-except-for-one-key
    const {candidates, onSelectCandidate, onSearchCandidates, ...selectProps} = props // eslint-disable-line

    // debounce search action.
    const onSearch = debounce((login: string) => {
        setSearching(true)
        setTimeout(() => {
           setSearching(false) 
        }, 500);

        props.onSearchCandidates(login)
    }, 800)


    return (
        <Select
            {...selectProps}
            showSearch
            filterOption={false}
            placeholder="Select approvers"
            notFoundContent={searching ? <Spin size="small" /> : null}
            onSelect={props.onSelectCandidate}
            onSearch={onSearch} >
                {props.candidates.map((candidate, idx) => {
                    const approver = props.approvers.find(approver => approver.id === candidate.id)
                    const checked = approver !== undefined

                    return (
                        <Select.Option 
                            key={idx}
                            value={candidate.id}>
                            <span>
                                <Avatar size="small" src={candidate.avatar}/>&nbsp;
                                {candidate.login}&nbsp;
                                {(checked) ? 
                                    <CheckOutlined style={{color: "purple"}}/> :
                                    null}
                            </span>
                        </Select.Option>
                    )
                })}
        </Select>
    )
}