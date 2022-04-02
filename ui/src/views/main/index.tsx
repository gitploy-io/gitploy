import { useEffect } from "react"
import { shallowEqual } from "react-redux"
import { Layout } from "antd"
import { Helmet } from "react-helmet"
import moment from "moment"

import { useAppSelector, useAppDispatch } from "../../redux/hooks"
import { subscribeEvents } from "../../apis"
import { 
    init, 
    searchDeployments, 
    searchReviews, 
    fetchLicense, 
    notifyDeploymentEvent, 
    notifyReviewmentEvent, 
    mainSlice as slice 
} from "../../redux/main"

import MainHeader from "./Header"
import MainContent from "./Content"
import LicenseWarning from "./LicenseWarningFooter"

const { Header, Content, Footer } = Layout

// eslint-disable-next-line
export default function Main(props: React.PropsWithChildren<{}>) { 
    const { 
        deployments,
        reviews,
    } = useAppSelector(state => state.main, shallowEqual)
    const dispatch = useAppDispatch()

    useEffect(() => {
        dispatch(init())
        dispatch(searchDeployments())
        dispatch(searchReviews())
        dispatch(fetchLicense())

        const sub = subscribeEvents((event) => {
            dispatch(slice.actions.handleDeploymentEvent(event))
            dispatch(slice.actions.handleReviewEvent(event))
            dispatch(notifyDeploymentEvent(event))
            dispatch(notifyReviewmentEvent(event))
        })

        return () => {
            sub.close()
        }
    }, [dispatch])

    return (
        <Layout className="layout">
            <Helmet>
                <title>Gitploy</title>
                {(deployments.length + reviews.length > 0)?
                    <link rel="icon" href="/spinner.ico" />
                    :
                    <link rel="icon" href="/favicon.ico" />}
            </Helmet>
            <Header>
                <MainHeader />
            </Header>
            <Content style={{ padding: "50px 50px" }}>
                <MainContent>
                    {props.children}
                </MainContent>
            </Content>
            <Footer style={{ textAlign: "center" }}>
                <div>
                    <LicenseWarning />
                </div>
                <div>
                    Gitploy ©{moment().format("YYYY")} Created by Gitploy.IO 
                </div>
            </Footer>
        </Layout>
    )
}
