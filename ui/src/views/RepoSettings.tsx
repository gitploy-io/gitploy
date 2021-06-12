import { useEffect  } from "react";
import { shallowEqual } from "react-redux";
import { useParams } from "react-router-dom";
import { message, PageHeader } from "antd"

import { useAppSelector, useAppDispatch } from "../redux/hooks"
import { repoSettingsSlice, init, save, deactivate } from "../redux/repoSettings"
import { RepoPayload, RequestStatus } from "../models";

import SettingsForm from "../components/SettingsForm"
import Spin from "../components/Spin";

const { actions } = repoSettingsSlice

interface Params {
    namespace: string
    name: string
}

export default function RepoSettings() {
    let { namespace, name } = useParams<Params>()
    const { repo, saving, deactivating } = useAppSelector(state => state.repoSettings, shallowEqual)
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

    const handleSaving = () => {
        if (saving === RequestStatus.Success) {
            dispatch(actions.unsetSaving())
            message.success("It has succeed to save.", 3)
        } else if (saving === RequestStatus.Failure) {
            dispatch(actions.unsetSaving())
            message.error("It has failed to save.", 3)
        }
    }

    const handleDeactivating = () => {
        if (deactivating === RequestStatus.Success) {
            dispatch(actions.unsetDeactivating())
            window.location.reload()
        } else if (deactivating === RequestStatus.Failure) {
            message.error("Only admin permission can deactivate.", 3)
            dispatch(actions.unsetDeactivating())
        }
    }

    handleSaving()
    handleDeactivating()

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