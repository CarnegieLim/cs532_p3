import React, { useState } from 'react';

import { Card, Form, Row, Col, Input, Button } from 'antd';
import { DownOutlined, UpOutlined } from '@ant-design/icons';

import './TopK.css'


const { Search } = Input;


export default class TopK extends React.Component {

    formRef = React.createRef();

    constructor(props) {
        super(props);
        // this.state = {
        //     result: null
        // }

        this.onFinish = this.onFinish.bind(this)
    }

    // onPredict(sentence) {
    //     // console.log(sentence);
    //     fetch("http://192.168.0.102:8080/predict?sentence=" + sentence)
    //         .then(response => response.json())
    //         .then(data => {
    //             this.setState({ result: data.sentiment })
    //         });
    // }

    onFinish(values){
        console.log('Received values of form: ', values);
    };

    render() {
        // const [form] = Form.useForm();
        // const [expand, setExpand] = useState(false);
        return (
            <Card>
                <Form
                    ref={this.formRef}
                    name="advanced_search"
                    className="ant-advanced-search-form"
                    onFinish={this.onFinish}
                >
                    <Row gutter={24}>
                        <Col span={8}>
                            <Form.Item
                                name="numVote"
                                label="numVote"
                                rules={[
                                    {
                                        required: true,
                                        message: 'Input something!',
                                    },
                                ]}
                            >
                                <Input placeholder="placeholder" />
                            </Form.Item>
                        </Col>
                        <Col span={8}>
                            <Form.Item
                                name="genresA"
                                label="genresA"
                                rules={[
                                    {
                                        required: true,
                                        message: 'Input something!',
                                    },
                                ]}
                            >
                                <Input placeholder="placeholder" />
                            </Form.Item>
                        </Col>
                        <Col span={8}>
                            <Form.Item
                                name="genresB"
                                label="genresB"
                                rules={[
                                    {
                                        required: true,
                                        message: 'Input something!',
                                    },
                                ]}
                            >
                                <Input placeholder="placeholder" />
                            </Form.Item>
                        </Col>
                    </Row>
                    <Row>
                        <Col
                            span={24}
                            style={{
                                textAlign: 'right',
                            }}
                        >
                            <Button type="primary" htmlType="submit">
                                Search
                            </Button>
                            <Button
                                style={{
                                    margin: '0 8px',
                                }}
                                onClick={() => {
                                    form.resetFields();
                                }}
                            >
                                Clear
                            </Button>
                            {/* <a
                                style={{
                                    fontSize: 12,
                                }}
                                onClick={() => {
                                    setExpand(!expand);
                                }}
                            >
                                {expand ? <UpOutlined /> : <DownOutlined />} Collapse
                            </a> */}
                        </Col>
                    </Row>
                </Form>
                <div className="search-result-list">Search Result List</div>
            </Card>
        )
    }
}

