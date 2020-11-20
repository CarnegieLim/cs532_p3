import React from 'react';

import {Card, Input, Tag} from 'antd';


const { Search } = Input;

export default class Predict extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            result: null
        }

        this.onPredict = this.onPredict.bind(this)
    }

    onPredict(sentence) {
        // console.log(sentence);
        fetch("http://192.168.0.102:8080/predict?sentence="+sentence)
            .then(response => response.json())
            .then(data => {
                this.setState({result: data.sentiment})
            });
    }

    render() {
        return (
            <Card>

                <Search
                    placeholder="input search text"
                    allowClear
                    enterButton="Search"
                    size="large"
                    onSearch={this.onPredict}
                />

                <Tag color="#f50">{this.state.result}</Tag>
            </Card>
        )
    }
}