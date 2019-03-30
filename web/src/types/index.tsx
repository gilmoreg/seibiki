export interface MeaningData {
    gloss: string;
    partofspeech: string[];
    misc: string[];
}

export interface EntryData {
    sequence: number;
    kanji: string[] | null;
    readings: string[];
    meanings: MeaningData[] | null;
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
    tokens: TokenData[];
}

export interface StoreState {
    selected: number;
    words: WordData[];
}