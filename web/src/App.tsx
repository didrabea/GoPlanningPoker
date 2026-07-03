import { useContext, useState } from "react";
import PlanningPoker from "./screens/MainRoom";
import RoomModal from "./screens/RoomModal";
import "./App.scss";
import AppStateContext from "./util/context.ts";

function App() {
  const { state } = useContext(AppStateContext);

  const [showCreate, setShowCreate] = useState(false);
  const [showJoin, setShowJoin] = useState(false);

  return (
    <div className="app-home">
      {state.room.id !== "" ? (
        <PlanningPoker />
      ) : (
        <>
          <div className="hero-card">
            <div className="logo">🃏</div>

            <h1>Planning Poker</h1>

            <p>
              Estimate stories together, reveal votes simultaneously, and keep
              sprint planning fast and collaborative.
            </p>

            <div className="actions">
              <button
                className="btn btn-primary"
                onClick={() => setShowCreate(true)}
              >
                Create Room
              </button>

              <button
                className="btn btn-secondary"
                onClick={() => setShowJoin(true)}
              >
                Join Room
              </button>
            </div>
          </div>

          <RoomModal
            mode="create"
            isOpen={showCreate}
            onClose={() => setShowCreate(false)}
          />

          <RoomModal
            mode="join"
            isOpen={showJoin}
            onClose={() => setShowJoin(false)}
          />
        </>
      )}{" "}
    </div>
  );
}

export default App;
