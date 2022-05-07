import { useEffect } from 'react';
import { shallowEqual } from 'react-redux';
import { useParams } from 'react-router-dom';
import { PageHeader, Button } from 'antd';
import { Result } from 'antd';

import { useAppSelector, useAppDispatch } from '../../redux/hooks';
import {
  fetchConfig,
  listLocks,
  lock,
  unlock,
  repoLockSlice as slice,
  setAutoUnlock,
} from '../../redux/repoLock';
import LockList, { LockListProps } from './LockList';

export default (): JSX.Element => {
  const { namespace, name } = useParams<{
    namespace: string;
    name: string;
  }>();

  const { display, config, locks } = useAppSelector(
    (state) => state.repoLock,
    shallowEqual
  );

  const dispatch = useAppDispatch();

  useEffect(() => {
    const f = async () => {
      await dispatch(slice.actions.init({ namespace, name }));
      await dispatch(fetchConfig());
      await dispatch(listLocks());
      await dispatch(slice.actions.setDisplay(true));
    };
    f();
    // eslint-disable-next-line
  }, [dispatch]);

  const onClickLock = (env: string) => {
    dispatch(lock(env));
  };

  const onClickUnlock = (env: string) => {
    dispatch(unlock(env));
  };

  const onChangeExpiredAt = (env: string, expiredAt: Date) => {
    dispatch(setAutoUnlock({ env, expiredAt }));
  };

  if (!display) {
    return <></>;
  }

  if (!config) {
    return (
      <Result
        status="warning"
        title="There is no configuration file."
        extra={[
          <Button
            type="primary"
            key="console"
            target="_blank"
            href="https://www.gitploy.io/docs/concepts/deploy.yml"
          >
            Read Document
          </Button>,
          <Button
            type="link"
            key="link"
            target="_blank"
            href={`/link/${namespace}/${name}/config/new`}
          >
            New Configuration
          </Button>,
        ]}
      />
    );
  }

  return (
    <RepoLock
      envs={config ? config.envs : []}
      locks={locks}
      onClickLock={onClickLock}
      onClickUnlock={onClickUnlock}
      onChangeExpiredAt={onChangeExpiredAt}
    />
  );
};

interface RepoLockProps extends LockListProps {}

function RepoLock({
  envs,
  locks,
  onChangeExpiredAt,
  onClickLock,
  onClickUnlock,
}: RepoLockProps): JSX.Element {
  return (
    <div>
      <div>
        <PageHeader title="Lock" subTitle="Lock the environment." />
      </div>
      <div style={{ padding: '16px 24px' }}>
        <LockList
          envs={envs}
          locks={locks}
          onClickLock={onClickLock}
          onClickUnlock={onClickUnlock}
          onChangeExpiredAt={onChangeExpiredAt}
        />
      </div>
    </div>
  );
}
