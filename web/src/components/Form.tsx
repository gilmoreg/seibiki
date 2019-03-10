import * as React from "react";
import './Form.css'
import { WordData } from 'src/types';

const apiPort = process.env.NODE_ENV === 'production' ? '' : ':3001';

export interface FormProps {
    update: (w: WordData[]) => void;
}

export interface FormState {
    query: string
}

class Form extends React.Component<FormProps, FormState> {
    constructor(props: FormProps) {
        super(props);
        this.state = {
            query: ''
        }

        this.update = this.update.bind(this);
        this.submit = this.submit.bind(this);
    }

    update(e: React.FormEvent<HTMLTextAreaElement>) {
        this.setState({ query: e.currentTarget.value })
    }

    submit(e: React.FormEvent) {
        e.preventDefault();
        console.log()
        fetch(`${window.location.protocol}//${window.location.hostname}${apiPort}/api/lookup`, {
            method: 'POST',
            body: JSON.stringify(this.state),
        })
            .then(res => res.json())
            .then(res => this.props.update(res))
            .catch(err => console.error(err))
    }

    render() {
        return (
            <form className="Form" onSubmit={this.submit}>
                <div><textarea
                    rows={4}
                    cols={200}
                    maxLength={280}
                    placeholder={'Type or paste a Japanese sentence, max 280 characters.'}
                    onInput={this.update}
                >
                </textarea></div>
                <button onClick={this.submit}>Lookup</button>
            </form>
        )
    }
}

export default Form;