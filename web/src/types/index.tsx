export interface EntryData {
    sequence: number;
    kanji: string[] | null;
    readings: string[];
    meanings: string[];
    partofspeech: string;
}

export interface TokenData {
    id: number;
    class: string; // DUMMY, KNOWN, UNKNOWN, USER
    surface: string;
    pos: string[];
    base: string;
    reading: string;
    pron: string;
    entries: EntryData[] | null;
}

export interface WordData {
    surface: string;
    entries: EntryData[] | null;
    tokens: TokenData[];
}

export interface StoreState {
    selected: number;
    words: WordData[];
}