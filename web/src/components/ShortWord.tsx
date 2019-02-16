import * as React from 'react';
import './ShortWord.css';

export interface ShortWordProps {
    word: string;
    index: number;
    select: (i: number) => void;
}

class ShortWord extends React.Component<ShortWordProps, any> {
    constructor(props: ShortWordProps) {
        super(props);

        this.hover = this.hover.bind(this);
    }

    hover(e: React.MouseEvent<HTMLSpanElement, MouseEvent>) {
        this.props.select && this.props.select(this.props.index);
    }

    render() {
        return (
            <span className="ShortWord" onMouseEnter={this.hover}>{this.props.word}</span>
        )
    }
}

export default ShortWord;