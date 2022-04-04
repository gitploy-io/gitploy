import { useEffect  } from "react";
import { useParams } from "react-router-dom";
import { shallowEqual } from "react-redux";
import { PageHeader } from "antd"

import { useAppSelector, useAppDispatch } from "../../redux/hooks"
import { save, deactivate, repoSettingsSlice as slice } from "../../redux/repoSettings"
import { init } from "../../redux/repoSettings"

import SettingsForm, { SettingFormProps } from "./SettingsForm"
import { RequestStatus } from "../../models";

export default (): JSX.Element => {
    const { namespace, name } = useParams<{
        namespace: string
        name: string
    }>()

    const { 
        saving,
        repo 
    } = useAppSelector(state => state.repoSettings, shallowEqual)

    const dispatch = useAppDispatch()

    useEffect(() => {
        const f = async () => {
            await dispatch(init({namespace, name}))
        }
        f()
        // eslint-disable-next-line
    }, [dispatch])

    const onClickFinish = (values: any) => {
        dispatch(slice.actions.setConfigPath(values.config))
        dispatch(save())
    }

    const onClickDeactivate = () => {
        dispatch(deactivate())
    }

    return (
        <RepoSettings 
            saving={saving === RequestStatus.Pending}
            repo={repo}
            onClickFinish={onClickFinish}
            onClickDeactivate={onClickDeactivate}
        />
    )
}

// eslint-disable-next-line
interface RepoSettingsProps extends SettingFormProps {}

function RepoSettings({
    saving,
    repo,
    onClickFinish,
    onClickDeactivate,
}: RepoSettingsProps): JSX.Element {
    return (
        <div>
            <div>
                <PageHeader title="Settings"/>
            </div>
            <div>
                {(repo)?
                    <SettingsForm 
                        saving={saving}
                        repo={repo}
                        onClickFinish={onClickFinish}
                        onClickDeactivate={onClickDeactivate}
                    />
                    :
                    <></>}
            </div>
        </div>
    )
}