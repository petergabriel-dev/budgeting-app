export interface User {
    id: string;
    email: string;
    role?: string;
}

export interface AuthResponse {
    message: string;
    user?: User;
}

export interface LoginInput {
    email: string;
    password: string;
}

export interface RegisterInput {
    email: string;
    password: string;
}
