import React from 'react';
import { useState } from 'react';
import api from './api';
import { Select, Button, Form, DatePicker, Row, Col } from 'antd';
import { HistoryInfoContext } from "./App";
import "./Calculation.css"

const { Option } = Select;

function Calculation() {
    const [error, setError] = useState(null);
    const [data, setData] = useState({
        fiatAmounts: 0,
        cryptoAmount: 0,
        cryptoAVG: 0
    });
    const [loading, setLoading] = useState(false);

    const historyInfo = React.useContext(HistoryInfoContext);

    const [form] = Form.useForm();

    const onSubmit = (fieldsValue) => {
        const values = {
            ...fieldsValue,
            from: fieldsValue.from.format('YYYY-MM-DD HH:mm:ss'),
            to: fieldsValue.to.format('YYYY-MM-DD HH:mm:ss')
        };

        setLoading(true);
        api.get(`/history/calculate?from=${values.from}&to=${values.to}&payMethod=${values.payMethod}&cryptoCurrency=${values.cryptoCurrency}`)
            .then(res => {
                setData(res.data)
            })
            .catch(err => {
                setError(err)
            })
        setLoading(false);
    }

    if (error) {
        return (<div>{error}</div>)
    } else if (loading) {
        return (<div>Loading...</div>)
    } else {
        return (
            <Row>
                <Col span={8}>
                    <Form
                        className="history-form"
                        form={form}
                        onFinish={onSubmit}
                        initialValues={{ cryptoCurrency: "BTC" }}
                        layout="vertical"
                    >
                        <Form.Item
                            label="Crypto Currency"
                            name="cryptoCurrency"
                            rules={[
                                {
                                    required: true,
                                },
                            ]}
                        >
                            <Select style={{ width: 205 }}>
                                {(historyInfo.cryptoCurrencies || []).map(cryptoCurrency => (
                                    <Option key={cryptoCurrency}>{cryptoCurrency}</Option>
                                ))}
                            </Select>
                        </Form.Item>
                        <Form.Item
                            label="Pay Method"
                            name="payMethod"
                            rules={[
                                {
                                    required: true,
                                },
                            ]}
                        >
                            <Select style={{ width: 205 }}>
                                {(historyInfo.payMethods || []).map(payMethod => (
                                    <Option key={payMethod}>{payMethod}</Option>
                                ))}
                            </Select>
                        </Form.Item>
                        <Form.Item
                            label="From"
                            name="from"
                            rules={[
                                {
                                    required: true,
                                },
                            ]}
                        >
                            <DatePicker format="YYYY-MM-DD HH:mm:ss" />
                        </Form.Item>
                        <Form.Item
                            label="To"
                            name="to"
                            rules={[
                                {
                                    required: true,
                                },
                            ]}
                        >
                            <DatePicker format="YYYY-MM-DD HH:mm:ss" />
                        </Form.Item>

                        <Form.Item shouldUpdate>
                            {() => (
                                <Button
                                    type="primary"
                                    htmlType="submit"
                                    style={{ width: 205 }}
                                    disabled={
                                        !!form.getFieldsError().filter(({ errors }) => errors.length).length
                                    }
                                >
                                    Calculate
                                </Button>
                            )}
                        </Form.Item>
                    </Form>
                </Col>
                <Col span={8} offset={8}>
                    <div>
                        <div><b>Fiat amounts:</b> {data.fiatAmounts}</div>
                        <div><b>Crypto amount:</b> {data.cryptoAmount}</div>
                        <div><b>Crypto AVG:</b> {data.cryptoAVG}</div>
                    </div>
                </Col>
            </Row>
        )
    }
}

export default Calculation;