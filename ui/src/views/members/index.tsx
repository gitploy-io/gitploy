import { useEffect } from "react"
import { shallowEqual } from "react-redux"
import { Input } from "antd"
import { Helmet } from "react-helmet"

import { useAppSelector, useAppDispatch } from "../../redux/hooks"
import { membersSlice as slice, fetchUsers, perPage } from "../../redux/members"

import Main from '../main'
import MemberList from "./MemberList"
import Pagination from "../../components/Pagination"

const { Search } = Input

export default function Members(): JSX.Element {
    const { users, page } = useAppSelector(state => state.members, shallowEqual)
    const dispatch = useAppDispatch()

    useEffect(() => {
        dispatch(fetchUsers())
    }, [dispatch])

    const onSearch = (value: string) => {
        dispatch(slice.actions.setQuery(value))
        dispatch(fetchUsers())
    }

    const onClickPrev = () => {
        dispatch(slice.actions.decreasePage())
        dispatch(fetchUsers())
    }

    const onClickNext = () => {
        dispatch(slice.actions.increasePage())
        dispatch(fetchUsers())
    }

    return (
        <Main>
            <Helmet>
                <title>Members</title>
            </Helmet>
            <div>
                <h1>Members</h1>
            </div>
            <div style={{marginTop: "40px", paddingRight: "20px"}}>
                <Search placeholder="Search user ..." onSearch={onSearch} enterButton />
            </div>
            <div style={{marginTop: "40px"}}>
                <MemberList/>
            </div>
            <div style={{marginTop: "40px", textAlign: "center"}}>
                <Pagination 
                    disabledPrev={page <= 1}
                    disabledNext={users.length < perPage}
                    onClickPrev={onClickPrev}
                    onClickNext={onClickNext}
                />    
            </div>
        </Main>
    )
}