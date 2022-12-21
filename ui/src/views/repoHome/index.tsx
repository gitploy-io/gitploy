import { useEffect } from 'react';
import { Params, useParams, useSearchParams } from 'react-router-dom';
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

interface ParamsType extends Params {
  namespace: string;
  name: string;
}

const parseSearchParams = (
  searchParams: URLSearchParams
): { env: string; page: number } => {
  return {
    env: searchParams.get('env') as string,
    page: parseInt(searchParams.get('page') as string, 10),
  };
};

export default (): JSX.Element => {
  const { namespace, name } = useParams() as ParamsType;
  const [searchParams, setSearchParams] = useSearchParams({
    env: '',
    page: '1',
  });
  const { env, page } = parseSearchParams(searchParams);
  const { loading, deployments, envs } = useAppSelector(
    (state) => state.repoHome,
    shallowEqual
  );
  const dispatch = useAppDispatch();

  useEffect(() => {
    const f = async () => {
      await dispatch(slice.actions.init({ namespace, name }));
      await dispatch(fetchEnvs());
      await dispatch(fetchDeployments({ env, page }));
    };
    f();
  }, []);

  // Fetch deployments when search parameters change.
  useEffect(() => {
    dispatch(fetchDeployments({ env, page }));
  }, [env, page]);

  const onChangeEnv = (env: string) => {
    // When the environment is changed,
    // unconditionally return to the first page.
    setSearchParams({ env, page: '1' });
  };

  const onClickPrev = () => {
    // Subtract one for the page param.
    const prevPage = page != 1 ? page - 1 : 1;
    searchParams.set('page', prevPage.toString());
    setSearchParams(searchParams);
  };

  const onClickNext = () => {
    // Plus one for the page param.
    const nextPage = page + 1;
    searchParams.set('page', nextPage.toString());
    setSearchParams(searchParams);
  };

  return (
    <RepoHome
      loading={loading}
      deployments={deployments}
      env={env}
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
  env: string;
  envs: string[];
  onChangeEnv(env: string): void;
}

export function RepoHome({
  // Deployments
  loading,
  deployments,
  // Environment Selector
  env,
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
              defaultValue={env}
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
