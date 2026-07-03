export interface Room {
  id: string;
  name: string;
  participants: Participant[];
  topics: Record<string, Topic>;
  activeTopic: string;
}

export interface Participant {
  id: string;
  name: string;
  moderator?: boolean;
}

export interface Topic {
  id: string;
  title: string;
  description: string;
  votes: Record<string, string>;
  revealed: boolean;
}

export interface AppState {
  room: Room;
  participant: Participant;
}
