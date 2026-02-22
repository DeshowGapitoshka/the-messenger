export interface IMessage {
    id: number;
    text: string;
    createdAt: string; // Time
    createdBy: string; // Person
    status: 'sent' | 'got' | 'read' | 'error';
    isChanged: boolean;
}

export const classMembersString = '(text, created_at, created_by, status, is_changed)';
