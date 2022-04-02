import { useEffect } from "react"
import { shallowEqual } from "react-redux"
import { Helmet } from "react-helmet"

import { useAppSelector, useAppDispatch } from "../../redux/hooks"
import { fetchMe, checkSlack } from "../../redux/settings"

import Main from "../main"
import UserDescription from "./UserDescriptions"
import SlackDescriptions from "./SlackDescriptions"

export default function Settings(): JSX.Element {
    const { isSlackEnabled } = useAppSelector(state => state.settings, shallowEqual)
    const dispatch = useAppDispatch()

    useEffect(() => {
        dispatch(fetchMe())
        dispatch(checkSlack())
    }, [dispatch])

    return (
        <Main>
            <Helmet>
                <title>Settings</title>
            </Helmet>
            <h1>Settings</h1>
            <div>
                <UserDescription />
            </div>
            {(isSlackEnabled)?
                <div style={{marginTop: "40px", marginBottom: "20px"}}>
                    <SlackDescriptions />
                </div>
                :
                <></>}
        </Main>
    )
}