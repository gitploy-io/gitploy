import React from "react"
import { Select, SelectProps, Avatar, Spin } from "antd"
import debounce from "lodash.debounce"

import { User } from "../models"

export interface ApproversSelectProps extends SelectProps<string>{
    candidates: User[]
    onSelectCandidate(id: string): void
    onDeselectCandidate(id: string): void
    onSearchCandidates(login: string): void
}

export default function ApproversSelect(props: ApproversSelectProps) {
    const [ searching, setSearching ] = React.useState<boolean>(false)

    // Clone Select props only
    // https://stackoverflow.com/questions/34698905/how-can-i-clone-a-javascript-object-except-for-one-key
    const {candidates, onSelectCandidate, onDeselectCandidate, onSearchCandidates, ...selectProps} = props

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
            mode="multiple"
            filterOption={false}
            placeholder="Select approvers"
            notFoundContent={searching ? <Spin size="small" /> : null}
            onSelect={props.onSelectCandidate}
            onDeselect={props.onDeselectCandidate}
            onSearch={onSearch} >
                {props.candidates.map((candidate, idx) => {
                    return (
                        <Select.Option 
                            key={idx}
                            value={candidate.id}>
                            <span><Avatar size="small" src={candidate.avatar}/> {candidate.login}</span>
                        </Select.Option>
                    )
                })}
        </Select>
    )
}