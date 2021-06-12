import { useEffect  } from "react";
import { shallowEqual } from "react-redux";
import { useParams } from "react-router-dom";
import { PageHeader } from "antd"

import { useAppSelector, useAppDispatch } from "../redux/hooks"
import { init, save, deactivate } from "../redux/repoSettings"
import { RepoPayload, RequestStatus } from "../models";

import SettingsForm from "../components/SettingsForm"
import Spin from "../components/Spin";

interface Params {
    namespace: string
    name: string
}

export default function RepoSettings() {
    let { namespace, name } = useParams<Params>()
    const { repo, saving } = useAppSelector(state => state.repoSettings, shallowEqual)
    const dispatch = useAppDispatch()

    useEffect(() => {
        const f = async () => {
            await dispatch(init({namespace, name}))
        }
        f()
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [dispatch])

    const onClickSave = (payload: RepoPayload) => {
        dispatch(save(payload))
    }

    const onClickDeactivate = () => {
        dispatch(deactivate())
    }

    if (repo === null) {
        return <div><Spin /></div>
    }

    return (
        <div>
            <div>
                <PageHeader
                    title="Settings"/>
            </div>
            <div>
                <SettingsForm 
                    repo={repo}
                    saving={saving === RequestStatus.Pending}
                    onClickSave={onClickSave}
                    onClickDeactivate={onClickDeactivate}/>
            </div>
        </div>
    )
}