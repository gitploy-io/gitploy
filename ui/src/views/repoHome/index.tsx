import { useEffect } from 'react';
import { useParams } from 'react-router-dom';
import { shallowEqual } from 'react-redux';
import { PageHeader, Select } from 'antd';

import { useAppSelector, useAppDispatch } from '../../redux/hooks';
import {
  repoHomeSlice as slice,
  fetchEnvs,
  fetchDeployments,
  perPage,
} from '../../redux/repoHome';

import ActivityLogs, { ActivityLogsProps } from './ActivityLogs';
import Spin from '../../components/Spin';
import Pagination, { PaginationProps } from '../../components/Pagination';

const { Option } = Select;

export default (): JSX.Element => {
  const { namespace, name } = useParams<{
    namespace: string;
    name: string;
  }>();

  const { loading, deployments, envs, page } = useAppSelector(
    (state) => state.repoHome,
    shallowEqual
  );

  const dispatch = useAppDispatch();

  useEffect(() => {
    const f = async () => {
      await dispatch(slice.actions.init({ namespace, name }));
      await dispatch(fetchEnvs());
      await dispatch(fetchDeployments());
    };
    f();
  }, [dispatch]);

  const onChangeEnv = (env: string) => {
    dispatch(slice.actions.setEnv(env));
    dispatch(fetchDeployments());
  };

  const onClickPrev = () => {
    dispatch(slice.actions.decreasePage());
    dispatch(fetchDeployments());
  };

  const onClickNext = () => {
    dispatch(slice.actions.increasePage());
    dispatch(fetchDeployments());
  };
  return (
    <RepoHome
      loading={loading}
      deployments={deployments}
      envs={envs}
      onChangeEnv={onChangeEnv}
      disabledPrev={page <= 1}
      disabledNext={deployments.length < perPage}
      onClickPrev={onClickPrev}
      onClickNext={onClickNext}
    />
  );
};

interface RepoHomeProps extends ActivityLogsProps, PaginationProps {
  loading: boolean;
  envs: string[];
  onChangeEnv(env: string): void;
}

export function RepoHome({
  // Deployments
  loading,
  deployments,
  // Environment Selector
  envs,
  onChangeEnv,
  // Pagination
  disabledPrev,
  disabledNext,
  onClickNext,
  onClickPrev,
}: RepoHomeProps): JSX.Element {
  return (
    <div>
      <div>
        <PageHeader
          title="Activity Log"
          extra={[
            <Select
              key="1"
              style={{ width: 150 }}
              defaultValue=""
              onChange={onChangeEnv}
            >
              <Option value="">All Environments</Option>
              {envs.map((env, idx) => {
                return (
                  <Option key={idx} value={env}>
                    {env}
                  </Option>
                );
              })}
            </Select>,
          ]}
        />
      </div>
      <div style={{ marginTop: '30px', padding: '16px 24px' }}>
        {loading ? (
          <div style={{ textAlign: 'center' }}>
            <Spin />
          </div>
        ) : (
          <ActivityLogs deployments={deployments} />
        )}
      </div>
      <div style={{ marginTop: '20px', textAlign: 'center' }}>
        <Pagination
          disabledPrev={disabledPrev}
          disabledNext={disabledNext}
          onClickPrev={onClickPrev}
          onClickNext={onClickNext}
        />
      </div>
    </div>
  );
}
