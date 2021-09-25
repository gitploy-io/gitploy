import { useEffect  } from "react";
import { shallowEqual } from "react-redux";
import { useParams } from "react-router-dom";
import { PageHeader } from "antd"

import { useAppSelector, useAppDispatch } from "../redux/hooks"
import { init, save, deactivate, repoSettingsSlice as slice } from "../redux/repoSettings"
import { RequestStatus } from "../models";

import RepoSettingsForm from "../components/RepoSettingsForm"
import Spin from "../components/Spin";

interface Params {
    namespace: string
    name: string
}

export default function RepoSettings(): JSX.Element {
    const { namespace, name } = useParams<Params>()
    const { repo, saving } = useAppSelector(state => state.repoSettings, shallowEqual)
    const dispatch = useAppDispatch()

    useEffect(() => {
        const f = async () => {
            await dispatch(init({namespace, name}))
        }
        f()
        // eslint-disable-next-line
    }, [dispatch])

    const onClickSave = (payload: {configPath: string}) => {
        dispatch(slice.actions.setConfigPath(payload.configPath))
        dispatch(save())
    }

    const onClickDeactivate = () => {
        dispatch(deactivate())
    }

    if (!repo) {
        return <div><Spin /></div>
    }

    return (
        <div>
            <div>
                <PageHeader
                    title="Settings"/>
            </div>
            <div>
                <RepoSettingsForm 
                    repo={repo}
                    saving={saving === RequestStatus.Pending}
                    onClickSave={onClickSave}
                    onClickDeactivate={onClickDeactivate}
                />
            </div>
        </div>
    )
}