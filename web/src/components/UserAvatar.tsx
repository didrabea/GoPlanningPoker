import type { Participant } from "../util/interfaces.ts";

interface UserAvatarProps {
  participant: Participant;
  index: number;
  total: number;
  revealed: boolean;
  vote?: string;
}

export function UserAvatar({
  participant,
  index,
  total,
  revealed,
  vote,
}: UserAvatarProps) {
  const radius = 260;
  const center = 350;

  const angle = (index / total) * 2 * Math.PI - Math.PI / 2;

  const x = center + radius * Math.cos(angle);
  const y = center + radius * Math.sin(angle);

  return (
    <div
      className="participant"
      style={{
        left: x,
        top: y,
      }}
    >
      <div className="avatar" />

      <div className="participant-name">{participant.name}</div>

      <div className="vote-state">{revealed ? vote : !!vote ? "✓" : "..."}</div>
    </div>
  );
}
