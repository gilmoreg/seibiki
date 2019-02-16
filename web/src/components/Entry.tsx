import * as React from 'react';
import { EntryData } from '../types';
import './Entry.css';

export interface EntryProps {
    entry: EntryData;
}

class Entry extends React.Component<EntryProps, any> {
    constructor(props: EntryProps) {
        super(props);
    }

    render() {
        const { entry } = this.props;
        const kanji = entry.kanji && entry.kanji.length && <li>Kanji: {entry.kanji.join(', ')}</li>;
        return (
            <span className="Entry">
                <ul>
                    {kanji || ''}
                    <li>Readings: {entry.readings && entry.readings.join(', ')}</li>
                    <li>Meanings: {entry.meanings && entry.meanings.join(', ')}</li>
                    <li>Part of speech: {entry.partofspeech}</li>
                </ul>
            </span>
        );
    }
}

export default Entry;
