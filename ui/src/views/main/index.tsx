import { useEffect } from "react"
import { shallowEqual } from "react-redux"
import { Layout } from "antd"
import { Helmet } from "react-helmet"
import moment from "moment"

import { useAppSelector, useAppDispatch } from "../../redux/hooks"
import { subscribeDeploymentEvents, subscribeReviewEvents } from "../../apis"
import { 
    init, 
    searchDeployments, 
    searchReviews, 
    fetchLicense, 
    notifyDeploymentEvent, 
    notifyReviewmentEvent, 
    mainSlice as slice 
} from "../../redux/main"

import MainHeader, { HeaderProps } from "./Header"
import MainContent, { ContentProps } from "./Content"
import LicenseWarning, { LicenseWarningFooterProps } from "./LicenseWarningFooter"

const { Header, Content, Footer } = Layout

export default (props: React.PropsWithChildren<any>): JSX.Element => { 
    const { 
        authorized,
        available,
        expired,
        user,
        deployments,
        reviews,
    } = useAppSelector(state => state.main, shallowEqual)

    const dispatch = useAppDispatch()

    useEffect(() => {
        dispatch(init())
        dispatch(searchDeployments())
        dispatch(searchReviews())
        dispatch(fetchLicense())

        const deploymentEvents = subscribeDeploymentEvents((deployment) => {
            dispatch(slice.actions.handleDeploymentEvent(deployment))
            dispatch(notifyDeploymentEvent(deployment))
        })

        const reviewEvents = subscribeReviewEvents((review) => {
            dispatch(slice.actions.handleReviewEvent(review))
            dispatch(notifyReviewmentEvent(review))
        })

        return () => {
            deploymentEvents.close()
            reviewEvents.close()
        }
    }, [dispatch])

    const onClickRetry = () => {
        dispatch(slice.actions.setAvailable(true))
        dispatch(slice.actions.setExpired(false))
    }

    return (
        <Main 
            authorized={authorized}
            available={available}
            expired={expired}
            user={user}
            deployments={deployments}
            reviews={reviews}
            onClickRetry={onClickRetry}
            children={props.children}
        />
    )
}

interface MainProps extends HeaderProps, ContentProps, LicenseWarningFooterProps {}

function Main({
    authorized,
    available,
    expired,
    user,
    deployments,
    reviews,
    license,
    children,
    onClickRetry
}: React.PropsWithChildren<MainProps>) { 
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
                <MainHeader 
                    user={user}
                    deployments={deployments}
                    reviews={reviews}
                />
            </Header>
            <Content style={{ padding: "50px 50px" }}>
                <MainContent
                    authorized={authorized}
                    available={available}
                    expired={expired}
                    onClickRetry={onClickRetry}
                >
                    {children}
                </MainContent>
            </Content>
            <Footer style={{ textAlign: "center" }}>
                <div>
                    <LicenseWarning 
                        license={license}
                    />
                </div>
                <div>
                    Gitploy ©{moment().format("YYYY")} Created by Gitploy.IO 
                </div>
            </Footer>
        </Layout>
    )
}
