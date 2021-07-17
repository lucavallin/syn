import { configureStore, ThunkAction, Action } from "@reduxjs/toolkit";
import eventsReducer from "../components/events/eventSlice";

export const store = configureStore({
  reducer: {
    counter: eventsReducer,
  },
});

export type AppDispatch = typeof store.dispatch;
export type RootState = ReturnType<typeof store.getState>;
export type AppThunk<ReturnType = void> = ThunkAction<
  ReturnType,
  RootState,
  unknown,
  Action<string>
>;
