import { defineStore } from "pinia";
import { useBackend, User } from "../backend";

type AuthStoreState = {
    user: User | null;
    users: User[];
}

const TOKEN_KEY_NAME = "token";

export const useAuthStore = defineStore("auth", {
    state: (): AuthStoreState => ({
        user: null,
        users: [],
    }),
    getters: {
        isAuthenticated(): boolean {
            return !! this.user;
        },
    },
    actions: {
        async initialize() {
            await this.refreshAuth();

        },
        async fetchUser() {
            this.user = await useBackend().fetchUser();
        },

        async fetchUsers() {
            this.users = await useBackend().fetchUsers();
        },

        /**
         * Refresh the authentication state.
         */
        async refreshAuth(): Promise<boolean> {
            try {
                await this.fetchUser();
            } catch (error) {
            }

            return this.isAuthenticated;
        },
        async login({ username, password }) {
            const backend = useBackend();
            const { success } = await backend.login({ username, password });
            if ( success ) {
                await Promise.all([
                    this.fetchUser(),
                    this.fetchUsers(),
                ]);
            }
        },
    },
});
