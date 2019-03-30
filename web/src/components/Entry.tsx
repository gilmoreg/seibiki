import * as React from 'react';
import { EntryData } from '../types';
import Meaning from './Meaning';
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
        const meanings = entry.meanings && entry.meanings.map((m, i) => <Meaning meaning={m} key={"meaning" + i} />);
        return (
            <span className="Entry">
                <ul>
                    {kanji || ''}
                    <li>Readings: {entry.readings && entry.readings.join(', ')}</li>
                    <ul>{meanings}</ul>
                </ul>
            </span>
        );
    }
}

export default Entry;
