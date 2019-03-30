import * as React from "react";
import { WordData } from '../types';
import ShortWord from './ShortWord';
import Word from './Word';

import './Answer.css';

export interface AnswerProps {
    words?: WordData[];
}

export interface AnswerState {
    selected: number;
}

class Answer extends React.Component<AnswerProps, AnswerState> {
    constructor(props: AnswerProps) {
        super(props);
        this.state = { selected: -1 };

        this.select = this.select.bind(this);
    }

    select(id: number) {
        this.setState({ selected: id });
    }

    render() {
        const wordComponents = this.props.words && this.props.words.map((w, i) =>
            <ShortWord
                word={w.surface}
                index={i}
                key={'wordComponents' + w.surface + i}
                select={this.select}
            />);

        const selectedWord = (this.state.selected >= 0) ?
            // @ts-ignore
            <Word word={this.props.words[this.state.selected]} /> :
            <span />

        return (
            <div>
                <div className="Answer">{wordComponents}</div>
                <div>{selectedWord}</div>
            </div>
        );
    }
}

export default Answer;