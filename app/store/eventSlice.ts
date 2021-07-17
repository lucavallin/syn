import {
  createAsyncThunk,
  createEntityAdapter,
  createSlice,
} from "@reduxjs/toolkit";
import { RootState } from "./store";

interface EventData {
  id: string;
}

export const getEvents = createAsyncThunk("events/getEvents", async () => {
  const response = await fetch("https://reqres.in/api/users?delay=1");
  return (await response.json()).data as EventData[];
});

export const eventsAdapter = createEntityAdapter<EventData>();

const eventsSlice = createSlice({
  name: "events",
  initialState: eventsAdapter.getInitialState({
    loading: false,
  }),
  reducers: {},
  extraReducers: (builder) => {
    builder.addCase(getEvents.pending, (state) => {
      state.loading = true;
    });
    builder.addCase(getEvents.fulfilled, (state, action) => {
      eventsAdapter.setAll(state, action.payload);
      state.loading = false;
    });
    builder.addCase(getEvents.rejected, (state) => {
      state.loading = false;
    });
  },
});

export const { selectAll: selectAllEvents } = eventsAdapter.getSelectors(
  (state: RootState) => state.events
);

export default eventsSlice.reducer;
