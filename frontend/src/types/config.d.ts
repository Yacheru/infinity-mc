export interface IConfig {
    apiURL: string;
    status: string;
}

enum Status {
    local = "local",
    prod = "prod",
}