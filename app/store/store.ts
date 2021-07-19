import { configureStore, getDefaultMiddleware } from "@reduxjs/toolkit";
import eventsReducer from "./eventSlice";

export const store = configureStore({
  reducer: {
    events: eventsReducer,
  },
  middleware: getDefaultMiddleware({
    serializableCheck: false,
  }),
});

export type AppDispatch = typeof store.dispatch;
export type RootState = ReturnType<typeof store.getState>;
