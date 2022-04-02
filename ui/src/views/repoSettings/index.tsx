import { useEffect  } from "react";
import { useParams } from "react-router-dom";
import { shallowEqual } from "react-redux";
import { PageHeader } from "antd"

import { useAppSelector, useAppDispatch } from "../../redux/hooks"
import { init } from "../../redux/repoSettings"

import SettingsForm from "./SettingsForm"

interface Params {
    namespace: string
    name: string
}

export default function RepoSettings(): JSX.Element {
    const { namespace, name } = useParams<Params>()
    const { repo } = useAppSelector(state => state.repoSettings, shallowEqual)
    const dispatch = useAppDispatch()

    useEffect(() => {
        const f = async () => {
            await dispatch(init({namespace, name}))
        }
        f()
        // eslint-disable-next-line
    }, [dispatch])

    return (
        <div>
            <div>
                <PageHeader title="Settings"/>
            </div>
            <div>
                {(repo)?
                    <SettingsForm />
                    :<></>}
            </div>
        </div>
    )
}