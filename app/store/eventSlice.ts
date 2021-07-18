import {
  createAsyncThunk,
  createEntityAdapter,
  createSlice,
} from "@reduxjs/toolkit";
import { RootState } from "./store";
import { firestore } from "../services/firebase";

export interface EventData {
  id: string;
  created: string;
  labels: Array<{ description: string; score: number }>;
}

export const getEvents = createAsyncThunk("events/getEvents", async () => {
  const response = await firestore.collection("Uploads").get();
  return response.docs.map((d) => ({
    id: d.id,
    created: d.get("created").toDate(),
    labels: d.get("labels"),
  })) as EventData[];
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
