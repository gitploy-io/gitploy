import { useEffect } from "react"
import { shallowEqual } from "react-redux"
import { Input } from "antd"

import { useAppSelector, useAppDispatch } from "../redux/hooks"
import { membersSlice as slice, fetchUsers, updateUser, deleteUser, perPage } from "../redux/members"

import { User } from "../models"

import Main from './Main'
import MemberList from "../components/MemberList"
import Pagination from "../components/Pagination"

const { Search } = Input

export default function Member(): JSX.Element {
    const { users, page } = useAppSelector(state => state.members, shallowEqual)
    const dispatch = useAppDispatch()

    useEffect(() => {
        dispatch(fetchUsers())
    }, [dispatch])

    const onSearch = (value: string) => {
        dispatch(slice.actions.setQuery(value))
        dispatch(fetchUsers())
    }

    const onChangeSwitch = (user: User, checked: boolean) => {
        dispatch(updateUser({user, admin: checked}))
    }

    const onClickDelete = (user: User) => {
        dispatch(deleteUser(user))
    }

    const onClickPrev = () => {
        dispatch(slice.actions.decreasePage())
        dispatch(fetchUsers())
    }

    const onClickNext = () => {
        dispatch(slice.actions.increasePage())
        dispatch(fetchUsers())
    }

    return <Main>
        <div>
            <h1>Members</h1>
        </div>
        <div style={{marginTop: "40px", paddingRight: "20px"}}>
            <Search placeholder="Search user ..." onSearch={onSearch} enterButton />
        </div>
        <div style={{marginTop: "40px"}}>
            <MemberList
                users={users}
                onChangeSwitch={onChangeSwitch}
                onClickDelete={onClickDelete}
            />
        </div>
        <div style={{marginTop: "40px", textAlign: "center"}}>
            <Pagination 
                page={page}
                isLast={users.length !== perPage}
                onClickPrev={onClickPrev}
                onClickNext={onClickNext}
            />    
        </div>
    </Main>
}