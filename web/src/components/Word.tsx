import * as React from 'react';
import { WordData } from '../types';
import Entry from './Entry';
import Token from './Token';
import './Word.css';

export interface WordProps {
    word: WordData;
}

class Word extends React.Component<WordProps, any> {
    constructor(props: WordProps) {
        super(props);
    }

    render() {
        const { word } = this.props;
        const entries = word.entries && word.entries.map((e, i) => <Entry entry={e} key={e.sequence + i} />);
        const tokens = word.tokens.map((t, i) => <Token token={t} key={t.id + i} />);

        return (
            <div className="Word">
                <h1>{word.surface}</h1>
                {entries ? <h2>Entries</h2> : ''}
                {entries ? <div className="entries">{entries}</div> : ''}
                <h2>Parts</h2>
                <div className="tokens">
                    {tokens}
                </div>
            </div>
        );
    }
}

export default Word;
