import { Helmet } from "react-helmet"

import Main from "./Main"
import SearchActivities from "../components/SearchActivities"
import ActivityLogs from "../components/ActivityLogs"
import Pagination from "../components/Pagination"

export default function Activities(): JSX.Element {
    return (
        <Main>
            <Helmet>
                <title>Activities</title>
            </Helmet>
            <h1>Activities</h1>
            <div style={{marginTop: 30}}>
                <h2>Search</h2>
                <SearchActivities 
                    onChangePeriod={() => console.log("period")}
                    onClickSearch={() => console.log("search")}
                />
            </div>
            <div style={{marginTop: 30}}>
                <ActivityLogs deployments={[]}/>
            </div>
            <div style={{marginTop: 30, textAlign: "center"}}>
                <Pagination page={0} isLast={true} onClickPrev={() => {console.log('')}} onClickNext={() => {console.log('')}} />
            </div>
        </Main>
    )
}