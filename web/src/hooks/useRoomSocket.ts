import { useEffect, useRef } from "react";
import { mapRoom } from "../util/mapRoom";

type WSMessage = {
  type: string;
  room: any;
};

export function useRoomSocket(roomId: string, dispatch: any) {
  const wsRef = useRef<WebSocket | null>(null);

  useEffect(() => {
    if (!roomId) return;

    const ws = new WebSocket(`ws://localhost:8080/ws?roomId=${roomId}`);
    wsRef.current = ws;

    ws.onopen = () => {
      console.log("WS connected");
    };

    ws.onmessage = (event) => {
      const msg: WSMessage = JSON.parse(event.data);

      if (msg.type === "room_updated") {
        dispatch({
          type: "SET_ROOM",
          value: mapRoom(msg.room),
        });
      }
    };

    ws.onerror = (err) => {
      console.error("WS error", err);
    };

    ws.onclose = () => {
      console.log("WS closed");
    };

    return () => {
      ws.close();
    };
  }, [roomId, dispatch]);

  return wsRef;
}
