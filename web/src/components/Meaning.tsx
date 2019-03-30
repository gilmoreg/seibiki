import * as React from 'react';
import { MeaningData } from '../types';

export interface MeaningProps {
    meaning: MeaningData;
}

class Meaning extends React.Component<MeaningProps, any> {
    constructor(props: MeaningProps) {
        super(props);
    }

    render() {
        const { meaning } = this.props;
        const pos = meaning.partofspeech && meaning.partofspeech.length ?
            ` (${meaning.partofspeech.join(',')})` : '';
        const misc = meaning.misc && meaning.misc.length ?
            ` (${meaning.misc.join(',')})` : '';
        return (
            <li className="Meaning">
                <span>{meaning.gloss}{pos}{misc}</span>
            </li>
        );
    }
}

export default Meaning;
