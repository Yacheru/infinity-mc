export interface IData {
    hronon: {
        title: string;
        description: string;
        costPlaceholder: string
        costs: {
            "2592000": string[];
            "7776000": string[];
            "15552000": string[];
        };
        durations: Durations[];
    };
    nickname: {
        title: string;
        description: string;
        costPlaceholder: string
        costs: {
            "15552000": string[];
        };
        durations: Durations[];
    };
    badge: {
        title: string;
        description: string;
        costPlaceholder: string
        costs: {
            "15552000": string[];
        };
        durations: Durations[];
    }
}

export type Durations = "2592000" | "7776000" | "15552000"