import { createSlice, createAsyncThunk } from '@reduxjs/toolkit';

import { searchDeployments as _searchDeployments } from '../apis';
import { Deployment } from '../models';

export const perPage = 30;

interface ActivitiesState {
  loading: boolean;
  deployments: Deployment[];
  page: number;
}

const initialState: ActivitiesState = {
  loading: false,
  deployments: [],
  page: 1,
};

export const searchDeployments = createAsyncThunk<
  Deployment[],
  {
    start?: Date;
    end?: Date;
    productionOnly: boolean;
  },
  {
    state: { activities: ActivitiesState };
  }
>(
  'activities/searchDeployments',
  async ({ start, end, productionOnly }, { getState, rejectWithValue }) => {
    const { page } = getState().activities;
    try {
      return await _searchDeployments(
        [],
        false,
        productionOnly,
        start,
        end,
        page,
        perPage
      );
    } catch (e) {
      return rejectWithValue(e);
    }
  }
);

export const activitiesSlice = createSlice({
  name: 'activities',
  initialState,
  reducers: {
    increasePage: (state) => {
      state.page = state.page + 1;
    },
    decreasePage: (state) => {
      state.page = state.page - 1;
    },
  },
  extraReducers: (builder) => {
    builder
      .addCase(searchDeployments.pending, (state) => {
        state.loading = true;
        state.deployments = [];
      })
      .addCase(searchDeployments.fulfilled, (state, action) => {
        state.loading = false;
        state.deployments = action.payload;
      })
      .addCase(searchDeployments.rejected, (state) => {
        state.loading = false;
      });
  },
});
