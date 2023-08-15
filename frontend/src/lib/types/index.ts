export interface Room {
    id: number;
    name: string;
}

export interface Location {
    id: number;
    name: string;
    address: string;
    rooms: Room[] | null;
}

export interface ApiResponse {
    responseCode: number;
    data?: any;
    message?: string;
    error?: string;
}
