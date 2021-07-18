import { configureStore, getDefaultMiddleware } from "@reduxjs/toolkit";
import logger from "redux-logger";
import eventsReducer from "./eventSlice";

export const store = configureStore({
  reducer: {
    events: eventsReducer,
  },
  middleware: getDefaultMiddleware().concat(logger),
});

export type AppDispatch = typeof store.dispatch;
export type RootState = ReturnType<typeof store.getState>;
