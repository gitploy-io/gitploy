import { createSlice, createAsyncThunk, PayloadAction } from '@reduxjs/toolkit';

import { searchDeployments as _searchDeployments } from '../apis';
import { Deployment } from '../models';

export const perPage = 30;

interface ActivitiesState {
  loading: boolean;
  startDate?: Date;
  endDate?: Date;
  productionOnly: boolean;
  deployments: Deployment[];
  page: number;
}

const initialState: ActivitiesState = {
  loading: false,
  productionOnly: false,
  deployments: [],
  page: 1,
};

export const searchDeployments = createAsyncThunk<
  Deployment[],
  void,
  {
    state: { activities: ActivitiesState };
  }
>('activities/searchDeployments', async (_, { getState, rejectWithValue }) => {
  const { startDate, endDate, productionOnly, page } = getState().activities;
  try {
    return await _searchDeployments(
      [],
      false,
      productionOnly,
      startDate,
      endDate,
      page,
      perPage
    );
  } catch (e) {
    return rejectWithValue(e);
  }
});

export const activitiesSlice = createSlice({
  name: 'activities',
  initialState,
  reducers: {
    setSearchOptions: (
      state: ActivitiesState,
      action: PayloadAction<{
        startDate: Date;
        endDate: Date;
        productionOnly: boolean;
      }>
    ) => {
      state.startDate = action.payload.startDate;
      state.endDate = action.payload.endDate;
      state.productionOnly = action.payload.productionOnly;
    },
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
