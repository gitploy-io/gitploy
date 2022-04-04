import { useEffect } from "react"
import { shallowEqual } from "react-redux"
import { Input } from "antd"
import { Helmet } from "react-helmet"

import { useAppSelector, useAppDispatch } from "../../redux/hooks"
import { membersSlice as slice, fetchUsers, perPage, updateUser, deleteUser } from "../../redux/members"
import { User } from "../../models"

import Main from '../main'
import MemberList, { MemberListProps } from "./MemberList"
import Pagination from "../../components/Pagination"

const { Search } = Input

export default (): JSX.Element => {
    const { users, page } = useAppSelector(state => state.members, shallowEqual)
    const dispatch = useAppDispatch()

    useEffect(() => {
        dispatch(fetchUsers())
    }, [dispatch])

    const onChangeSwitch = (user: User, checked: boolean) => {
        dispatch(updateUser({user, admin: checked}))
    }

    const onClickDelete = (user: User) => {
        dispatch(deleteUser(user))
    }

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
            <Members 
                users={users}
                page={page}
                onChangeSwitch={onChangeSwitch}
                onClickDelete={onClickDelete}
                onClickPrev={onClickPrev}
                onClickNext={onClickNext}
                onSearch={onSearch}
            />
        </Main>
    )
}

interface MembersProps extends MemberListProps {
    page: number
    onClickPrev(): void
    onClickNext(): void
    onSearch(value: string): void
}

function Members({
    users,
    page,
    onChangeSwitch,
    onClickDelete,
    onClickNext,
    onClickPrev,
    onSearch,
}: MembersProps): JSX.Element {
    return (
        <>
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
                <MemberList 
                    users={users}
                    onChangeSwitch={onChangeSwitch}
                    onClickDelete={onClickDelete}
                />
            </div>
            <div style={{marginTop: "40px", textAlign: "center"}}>
                <Pagination 
                    disabledPrev={page <= 1}
                    disabledNext={users.length < perPage}
                    onClickPrev={onClickPrev}
                    onClickNext={onClickNext}
                />    
            </div>
        </>
    )
}