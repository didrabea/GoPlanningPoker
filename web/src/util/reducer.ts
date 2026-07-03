import type { AppState, Room, Participant } from "./interfaces.ts";

export type Action =
  | { type: "SET_ROOM"; value: Room }
  | { type: "SET_PARTICIPANT"; value: Participant };

export const reducer = (state: AppState, action: Action) => {
  switch (action.type) {
    case "SET_ROOM":
      return { ...state, room: action.value };
    case "SET_PARTICIPANT":
      return { ...state, participant: action.value };

    default:
      return state;
  }
};
