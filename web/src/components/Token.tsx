import * as React from 'react';
import { TokenData } from '../types';
import Entry from './Entry';
import './Token.css';

export interface TokenProps {
    token: TokenData;
}

// http://www.edrdg.org/jmdictdb/cgi-bin/edhelp.py?svc=jmdict&sid=#kw_fld
const posMapping = {
    '助動詞': '&aux-v;',
    '形容詞': '&adj-i;',
    '副詞': '&adv;'
}

class Token extends React.Component<TokenProps, any> {
    constructor(props: TokenProps) {
        super(props);
    }

    render() {
        const { token } = this.props;
        const pos = token && token.pos && token.pos.length && token.pos.filter(p => p !== '*').join(', ');
        const targetPos = pos && pos.length && posMapping[token.pos[0]];
        const filteredEntries = targetPos && token.entries && token.entries.filter(e => e.partofspeech === targetPos);
        const entries = filteredEntries && filteredEntries.length ? filteredEntries : token.entries;
        // const entries = token.entries && token.entries
        //    
        // const entries = (guessedEntries.length ? guessedEntries : entries) 
        const displayEntries = entries && entries.map((e, i) => <Entry entry={e} key={token.id + i} />);

        return (
            <div className="Token">
                <p>{token.surface}</p>
                <ul>
                    <li>Part of Speech: {pos}</li>
                    <li>Base: {token.base}</li>
                    <li>Reading: {token.reading}</li>
                    <li>Proununciation: {token.pron}</li>
                </ul>
                {displayEntries}
            </div>
        );
    }
}

export default Token;
