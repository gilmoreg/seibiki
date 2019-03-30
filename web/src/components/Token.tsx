import * as React from 'react';
import { TokenData } from '../types';
import Entry from './Entry';
import './Token.css';

export interface TokenProps {
    token: TokenData;
}

class Token extends React.Component<TokenProps, any> {
    constructor(props: TokenProps) {
        super(props);
    }

    render() {
        const { token } = this.props;
        const pos = token && token.pos && token.pos.length && token.pos.filter(p => p !== '*').join(', ');
        const displayEntries = token.entries && token.entries.map((e, i) =>
            <Entry entry={e} key={token.id + i} />);

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
