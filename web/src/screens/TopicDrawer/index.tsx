import { useState } from "react";
import "./styles.scss";

export interface Topic {
  id: string;
  title: string;
  description?: string;
}

type Props = {
  topics: Topic[];
  activeTopicId?: string;
  isCreator: boolean;
  onSelect: (id: string) => void;
  onCreate: (title: string) => void;
};

export default function TopicDrawer({
  topics,
  activeTopicId,
  isCreator,
  onSelect,
  onCreate,
}: Props) {
  const [open, setOpen] = useState(false);
  const [newTopic, setNewTopic] = useState("");

  const createTopic = () => {
    if (!newTopic.trim()) return;

    onCreate(newTopic);
    setNewTopic("");
  };

  return (
    <>
      <button className="drawer-toggle" onClick={() => setOpen(!open)}>
        {open ? "→" : "Topics"}
      </button>

      <aside className={`topic-drawer ${open ? "open" : ""}`}>
        <div className="drawer-header">
          <h2>Topics</h2>
        </div>

        {isCreator && (
          <div className="new-topic">
            <input
              value={newTopic}
              placeholder="Add story..."
              onChange={(e) => setNewTopic(e.target.value)}
            />

            <button onClick={createTopic}>Add</button>
          </div>
        )}

        <div className="topic-list">
          {topics.map((topic) => (
            <button
              key={topic.id}
              className={`topic-card ${
                activeTopicId === topic.id ? "active" : ""
              }`}
              onClick={() => onSelect(topic.id)}
            >
              <h4>{topic.title}</h4>

              {topic.description && <p>{topic.description}</p>}
            </button>
          ))}
        </div>
      </aside>
    </>
  );
}
