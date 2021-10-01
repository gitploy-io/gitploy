import React from "react"
import { Select, SelectProps, Avatar, Spin } from "antd"
import debounce from "lodash.debounce"

import { User } from "../models"

export interface ApproversSelectProps extends SelectProps<string>{
    candidates: User[]
    // The type of parameter have to be string
    // because candidates could be dynamically changed with search.
    onSelectCandidate(candidate: User): void
    onDeselectCandidate(candidate: User): void
    onSearchCandidates(login: string): void
}

export default function ApproversSelect(props: ApproversSelectProps): JSX.Element {
    const [ searching, setSearching ] = React.useState<boolean>(false)
    const [ value, setValue ] = React.useState<User[]>([])

    // Clone Select props only
    // https://stackoverflow.com/questions/34698905/how-can-i-clone-a-javascript-object-except-for-one-key
    const {candidates, onSelectCandidate, onDeselectCandidate, onSearchCandidates, ...selectProps} = props // eslint-disable-line

    // debounce search action.
    const onSearch = debounce((login: string) => {
        setSearching(true)
        setTimeout(() => {
           setSearching(false) 
        }, 500);

        props.onSearchCandidates(login)
    }, 800)

    const _onSelectCandidate = (id: string) => {
        const candidate = props.candidates.find(c => c.id.toString() ===id)
        if (candidate === undefined) {
            return
        }

        onSelectCandidate(candidate)

        // Save value for the deselect event.
        value.push(candidate)
        setValue(value)
    }

    const _onDeselectCandidate = (id: string) => {
        const candidate = value.find(c => c.id.toString() === id)
        if (candidate === undefined) {
            return
        }

        onDeselectCandidate(candidate)

        const candidates = value.filter(c => c.id.toString() !== id)
        setValue(candidates)
    }

    return (
        <Select
            {...selectProps}
            mode="multiple"
            filterOption={false}
            placeholder="Select approvers"
            notFoundContent={searching ? <Spin size="small" /> : null}
            onSelect={_onSelectCandidate}
            onDeselect={_onDeselectCandidate}
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