import React, { createContext, type Dispatch } from "react";
import type { Action } from "./reducer.ts";
import type { AppState } from "./interfaces.ts";

const AppStateContext: React.Context<{
  dispatch: Dispatch<Action>;
  state: AppState;
}> = createContext({} as { dispatch: Dispatch<Action>; state: AppState });
export default AppStateContext;
