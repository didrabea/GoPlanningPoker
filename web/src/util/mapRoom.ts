import type { Room, Participant, Topic } from "./interfaces";

type ApiRoom = {
  id: string;
  name: string;
  participants: Record<string, Participant>;
  topics: Record<string, Topic>;
  activeTopic: string;
};

export function mapRoom(apiRoom: ApiRoom): Room {
  if (!apiRoom) {
    throw new Error("mapRoom received empty room");
  }

  return {
    id: apiRoom.id,
    name: apiRoom.name,
    participants: Object.values(apiRoom.participants || {}),
    topics: apiRoom.topics,
    activeTopic: apiRoom.activeTopic,
  };
}
