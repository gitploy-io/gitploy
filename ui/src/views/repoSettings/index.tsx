import { useEffect } from 'react';
import { Params, useParams } from 'react-router-dom';
import { shallowEqual } from 'react-redux';
import { PageHeader } from 'antd';
import { useAppSelector, useAppDispatch } from '../../redux/hooks';
import { save, deactivate } from '../../redux/repoSettings';
import { init } from '../../redux/repoSettings';
import SettingsForm, {
  SettingFormProps,
  SettingFormValues,
} from './SettingsForm';

interface ParamsType extends Params {
  namespace: string;
  name: string;
}

export default (): JSX.Element => {
  const { namespace, name } = useParams() as ParamsType;
  const { repo } = useAppSelector((state) => state.repoSettings, shallowEqual);
  const dispatch = useAppDispatch();

  useEffect(() => {
    const f = async () => {
      await dispatch(init({ namespace, name }));
    };
    f();
    // eslint-disable-next-line
  }, [dispatch]);

  const onClickFinish = (values: SettingFormValues) => {
    const f = async () => {
      await dispatch(save(values));
    };
    f();
  };

  const onClickDeactivate = () => {
    dispatch(deactivate());
  };

  if (!repo) {
    return <></>;
  }

  return (
    <RepoSettings
      configLink={`/link/${repo.namespace}/${repo.name}/config`}
      initialValues={{
        name: repo.name,
        config_path: repo.configPath,
      }}
      onClickFinish={onClickFinish}
      onClickDeactivate={onClickDeactivate}
    />
  );
};

interface RepoSettingsProps extends SettingFormProps {}

function RepoSettings({
  configLink,
  initialValues,
  onClickFinish,
  onClickDeactivate,
}: RepoSettingsProps): JSX.Element {
  return (
    <div>
      <div>
        <PageHeader title="Settings" />
      </div>
      <div>
        <SettingsForm
          configLink={configLink}
          initialValues={initialValues}
          onClickFinish={onClickFinish}
          onClickDeactivate={onClickDeactivate}
        />
      </div>
    </div>
  );
}
