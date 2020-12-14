import { writable } from 'svelte/store';

export const username = writable('Guest');
export const code = writable('');
export const participantId = writable('');
export const action = writable('');
export const currentWordCount = writable(0)

export const wsStore = writable('');
export const WS_ACTIVE = writable(false);
