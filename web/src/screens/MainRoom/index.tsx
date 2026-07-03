import { useContext, useEffect, useState } from "react";
import "./styles.scss";
import { UserAvatar } from "../../components/UserAvatar.tsx";
import AppStateContext from "../../util/context.ts";
import { api } from "../../util/api.ts";
import { mapRoom } from "../../util/mapRoom.ts";
import { useRoomSocket } from "../../hooks/useRoomSocket";
import TopicDrawer from "../TopicDrawer";

const voteOptions = ["0", "1", "2", "3", "5", "8", "13", "21", "?"];

export default function PlanningPoker() {
  const { state, dispatch } = useContext(AppStateContext);
  const [selectedVote, setSelectedVote] = useState("");
  const activeTopicId = state.room.activeTopic;
  const topic = state.room.topics?.[activeTopicId];
  const votes = topic ? Object.values(topic.votes) : [];
  const average =
    votes.length > 0
      ? votes.reduce((sum, vote) => {
          const v = Number(vote);
          return sum + (isNaN(v) ? 0 : v);
        }, 0) / votes.length
      : 0;

  useEffect(() => {
    console.log(
      state.room.topics?.[activeTopicId]?.votes[state.participant.id],
    );
    if (
      state.room.topics?.[activeTopicId]?.votes[state.participant.id] ===
      undefined
    )
      setSelectedVote("");
  }, [state.room.topics?.[activeTopicId]?.votes[state.participant.id]]);
  const topics = Object.values(state.room.topics || {});
  useRoomSocket(state.room.id, dispatch);

  const sendVote = async (value: string) => {
    setSelectedVote(value);

    console.log(state.room);

    await api.post(`/rooms/vote?id=${state.room.id}`, {
      userId: state.participant.id,
      vote: value,
    });
  };

  const resetVotes = async () => {
    const res = await api.post(`/rooms/reset?id=${state.room.id}`, {});
    dispatch({
      type: "SET_ROOM",
      value: mapRoom(res.room),
    });
  };

  const revealVotes = async () => {
    const res = await api.post(`/rooms/reveal?id=${state.room.id}`, {});
    dispatch({
      type: "SET_ROOM",
      value: mapRoom(res.room),
    });
  };

  return (
    <div className="app">
      <header className="header">
        <h1>{state.room.name}</h1>
        <span className="room-id">{state.room.id}</span>
      </header>

      <TopicDrawer
        topics={topics}
        activeTopicId={activeTopicId}
        isCreator={true}
        onSelect={(id) => {
          api.post(`/rooms/topics/set-active?id=${state.room.id}`, {
            topicId: id,
          });
        }}
        onCreate={(title) => {
          console.log(title);
        }}
      />

      <div className="board">
        <div className="table">
          <h2>{state.room.topics[activeTopicId].title}</h2>
          <p>
            {state.room.topics[activeTopicId].description &&
              state.room.topics[activeTopicId].description}
          </p>

          {topic?.revealed && (
            <>
              <p>Average: {average}</p>{" "}
            </>
          )}

          <div className="controls">
            <button onClick={() => revealVotes()}>Reveal</button>
            <button onClick={() => resetVotes()}>Reset</button>
          </div>
        </div>

        {state.room.participants.map((participant, index) => (
          <UserAvatar
            key={participant.id}
            participant={participant}
            index={index}
            total={state.room.participants.length}
            revealed={topic?.revealed || false}
            vote={topic?.votes[participant.id]}
          />
        ))}
      </div>

      <div className="vote-panel">
        <h3>Your Vote</h3>

        <div className="vote-grid">
          {voteOptions.map((vote) => (
            <button
              key={vote}
              className={`vote-card ${selectedVote === vote ? "selected" : ""}`}
              onClick={() => sendVote(vote)}
            >
              {vote}
            </button>
          ))}
        </div>
      </div>
    </div>
  );
}
