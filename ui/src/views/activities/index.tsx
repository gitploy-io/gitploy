import { useEffect } from 'react';
import { shallowEqual } from 'react-redux';
import { Helmet } from 'react-helmet';

import { useAppSelector, useAppDispatch } from '../../redux/hooks';
import {
  perPage,
  activitiesSlice,
  searchDeployments,
} from '../../redux/activities';

import Main from '../main';
import SearchActivities, {
  SearchActivitiesProps,
  SearchActivitiesValues,
} from './SearchActivities';
import ActivityHistory, { ActivityHistoryProps } from './ActivityHistory';
import Pagination, { PaginationProps } from '../../components/Pagination';
import Spin from '../../components/Spin';

const { actions } = activitiesSlice;

export default (): JSX.Element => {
  const { loading, deployments, page } = useAppSelector(
    (state) => state.activities,
    shallowEqual
  );

  const dispatch = useAppDispatch();

  useEffect(() => {
    dispatch(
      searchDeployments({
        productionOnly: false,
      })
    );

    // eslint-disable-next-line
  }, [dispatch]);

  const onClickSearch = (values: SearchActivitiesValues) => {
    dispatch(
      searchDeployments({
        start: values.period ? values.period[0].toDate() : undefined,
        end: values.period ? values.period[1].toDate() : undefined,
        productionOnly: values.productionOnly ? values.productionOnly : false,
      })
    );
  };

  const onClickPrev = () => dispatch(actions.decreasePage());

  const onClickNext = () => dispatch(actions.increasePage());

  return (
    <Main>
      <Activities
        onClickSearch={onClickSearch}
        loading={loading}
        deployments={deployments}
        disabledPrev={page <= 1}
        disabledNext={deployments.length != perPage}
        onClickPrev={onClickPrev}
        onClickNext={onClickNext}
      />
    </Main>
  );
};

interface ActivitiesProps
  extends SearchActivitiesProps,
    ActivityHistoryProps,
    PaginationProps {
  loading: boolean;
}

function Activities({
  // Properties to search.
  initialValues,
  onClickSearch,
  // Properties for the deployment history.
  loading,
  deployments,
  // Pagination for the pagination.
  disabledPrev,
  disabledNext,
  onClickPrev,
  onClickNext,
}: ActivitiesProps): JSX.Element {
  return (
    <>
      <Helmet>
        <title>Activities</title>
      </Helmet>
      <h1>Activities</h1>
      <div style={{ marginTop: 30 }}>
        <h2>Search</h2>
        <SearchActivities
          initialValues={initialValues}
          onClickSearch={onClickSearch}
        />
      </div>
      <div style={{ marginTop: 50 }}>
        <h2>History</h2>
        <div style={{ marginTop: 30 }}>
          {loading ? (
            <div style={{ textAlign: 'center' }}>
              <Spin />
            </div>
          ) : (
            <ActivityHistory deployments={deployments} />
          )}
        </div>
      </div>
      <div style={{ marginTop: 30, textAlign: 'center' }}>
        <Pagination
          disabledPrev={disabledPrev}
          disabledNext={disabledNext}
          onClickPrev={onClickPrev}
          onClickNext={onClickNext}
        />
      </div>
    </>
  );
}
