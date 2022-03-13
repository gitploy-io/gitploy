import { useEffect } from "react"
import { shallowEqual } from 'react-redux'
import { Helmet } from "react-helmet"

import { useAppSelector, useAppDispatch } from "../redux/hooks"
import { perPage, activitiesSlice, searchDeployments } from "../redux/activities"

import Main from "./Main"
import SearchActivities from "../components/SearchActivities"
import ActivityLogs from "../components/ActivityLogs"
import Pagination from "../components/Pagination"
import Spin from '../components/Spin'

const { actions } = activitiesSlice

export default function Activities(): JSX.Element {
    const { 
        loading,
        deployments,
        page,
    } = useAppSelector(state => state.activities, shallowEqual)

    const dispatch = useAppDispatch()

    useEffect(() => {
        dispatch(searchDeployments())
        // eslint-disable-next-line 
    }, [dispatch])

    const onChangePeriod = (start: Date, end: Date) => {
        dispatch(actions.setStart(start))
        dispatch(actions.setEnd(end))
    }

    const onClickSearch = () => dispatch(searchDeployments())

    const onClickPrev = () => dispatch(actions.decreasePage())

    const onClickNext = () => dispatch(actions.increasePage())

    return (
        <Main>
            <Helmet>
                <title>Activities</title>
            </Helmet>
            <h1>Activities</h1>
            <div style={{marginTop: 30}}>
                <h2>Search</h2>
                <SearchActivities 
                    onChangePeriod={onChangePeriod}
                    onClickSearch={onClickSearch}
                />
            </div>
            <div style={{marginTop: 50}}>
                <h2>History</h2>
                <div style={{marginTop: 30}}>
                    {(loading)? 
                        <div style={{textAlign: "center"}}>
                            <Spin />
                        </div> 
                        :
                        <ActivityLogs deployments={deployments}/>}
                </div>
            </div>
            <div style={{marginTop: 30, textAlign: "center"}}>
                <Pagination 
                    page={page} 
                    isLast={deployments.length !== perPage} 
                    onClickPrev={onClickPrev} 
                    onClickNext={onClickNext} />
            </div>
        </Main>
    )
}