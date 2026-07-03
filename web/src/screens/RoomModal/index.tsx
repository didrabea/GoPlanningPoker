import { useContext, useState } from "react";
import "./styles.scss";
import { api } from "../../util/api.ts";
import AppStateContext from "../../util/context.ts";
import { mapRoom } from "../../util/mapRoom.ts";

interface RoomModalProps {
  mode?: "create" | "join";
  isOpen: boolean;
  onClose: () => void;
}

export default function RoomModal({
  mode = "create",
  isOpen,
  onClose,
}: RoomModalProps) {
  const { dispatch } = useContext(AppStateContext);
  const [userName, setUserName] = useState("");
  const [roomName, setRoomName] = useState("");
  const [roomId, setRoomId] = useState("");

  if (!isOpen) return null;

  const isCreate = mode === "create";

  const createRoom = async () => {
    return await api.post(`/rooms/create`, {
      name: roomName,
      userName: userName,
    });
  };

  const joinRoom = async () => {
    return await api.post(`/rooms/join?id=${roomId}`, { userName: userName });
  };

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    try {
      const response = isCreate ? await createRoom() : await joinRoom();
      console.log(response);
      dispatch({
        type: "SET_ROOM",
        value: mapRoom(response.room),
      });

      dispatch({
        type: "SET_PARTICIPANT",
        value: response.participant,
      });

      onClose();
    } catch (error) {
      console.error(error);
      // todo show error message to participant
    }
  };

  return (
    <div className="modal-backdrop">
      <div className="room-modal">
        <button className="close-btn" onClick={onClose}>
          ×
        </button>

        <h2>{isCreate ? "Create Room" : "Join Room"}</h2>

        <form onSubmit={handleSubmit}>
          <div className="field">
            <label>Your Name</label>

            <input
              value={userName}
              onChange={(e) => setUserName(e.target.value)}
              placeholder="Bob"
            />
          </div>

          {isCreate && (
            <div className="field">
              <label>Room Name</label>

              <input
                value={roomName}
                onChange={(e) => setRoomName(e.target.value)}
                placeholder="Sprint Planning"
              />
            </div>
          )}

          {!isCreate && (
            <div className="field">
              <label>Room ID</label>

              <input
                value={roomId}
                onChange={(e) => setRoomId(e.target.value)}
                placeholder="ABC123"
              />
            </div>
          )}

          <button type="submit" className="submit-btn">
            {isCreate ? "Create Room" : "Join Room"}
          </button>
        </form>
      </div>
    </div>
  );
}
