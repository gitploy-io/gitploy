import { useEffect } from 'react';
import { shallowEqual } from 'react-redux';
import { Helmet } from 'react-helmet';

import { useAppSelector, useAppDispatch } from '../../redux/hooks';
import { fetchMe, checkSlack } from '../../redux/settings';

import Main from '../main';
import UserDescription, { UserDescriptionsProps } from './UserDescriptions';
import SlackDescriptions from './SlackDescriptions';

export default (): JSX.Element => {
  const { user, isSlackEnabled } = useAppSelector(
    (state) => state.settings,
    shallowEqual
  );

  const dispatch = useAppDispatch();

  useEffect(() => {
    dispatch(fetchMe());
    dispatch(checkSlack());
  }, [dispatch]);

  return (
    <Main>
      <Settings user={user} isSlackEnabled={isSlackEnabled} />
    </Main>
  );
};

interface SettingsProps extends UserDescriptionsProps {
  isSlackEnabled: boolean;
}

function Settings({ user, isSlackEnabled }: SettingsProps): JSX.Element {
  return (
    <>
      <Helmet>
        <title>Settings</title>
      </Helmet>
      <h1>Settings</h1>
      <div>
        <UserDescription user={user} />
      </div>
      {isSlackEnabled ? (
        <div style={{ marginTop: '40px', marginBottom: '20px' }}>
          <SlackDescriptions user={user} />
        </div>
      ) : (
        <></>
      )}
    </>
  );
}
