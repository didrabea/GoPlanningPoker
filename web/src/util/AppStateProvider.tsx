import React, { useCallback, useReducer } from "react";
import type { AppState } from "./interfaces";
import AppStateContext from "./context";
import { reducer } from "./reducer";

export default function AppStateProvider(props: { children: React.ReactNode }) {
  const INITIAL_STATE: AppState = {
    room: { id: "", name: "", participants: [], topics: {}, activeTopic: "" },
    participant: { id: "", name: "" },
  };

  const [state, dispatch] = useReducer(reducer, INITIAL_STATE);

  return (
    <AppStateContext.Provider
      value={useCallback(
        () => ({
          state,
          dispatch,
        }),
        [state, dispatch],
      )()}
    >
      {props.children}
    </AppStateContext.Provider>
  );
}
