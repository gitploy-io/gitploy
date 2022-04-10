import { Row, Col, Form, DatePicker, Button, Switch } from "antd"
import moment, { Moment } from "moment"

export interface SearchActivitiesValues {
    period?: [Moment, Moment]
    productionOnly?: boolean
}

export interface SearchActivitiesProps {
    initialValues?: SearchActivitiesValues 
    onClickSearch(values: SearchActivitiesValues): void
}

export default function SearchActivities(props: SearchActivitiesProps): JSX.Element {
    const content = (
        <>
            <Form.Item label="Period" name="period">
                <DatePicker.RangePicker 
                    format="YYYY-MM-DD HH:mm"
                    showTime={{
                        showSecond: false, 
                        defaultValue: [moment().hour(0).minute(0), moment().hour(23).minute(59)],
                    }}
                />
            </Form.Item>
            <Form.Item label="Production" name="productionOnly" valuePropName="checked">
                <Switch size="small" />
            </Form.Item>
            <Form.Item >
                <Button 
                    htmlType="submit" 
                    type="primary" 
                >
                    Search
                </Button>
            </Form.Item>
        </>
    )
    return (
        <Row>
            {/* Mobile view */}
            <Col span={24} lg={0}>
                <Form 
                    layout="horizontal" 
                    initialValues={props.initialValues}
                    onFinish={props.onClickSearch}
                >
                    {content}
                </Form>
            </Col>
            {/* Laptop */}
            <Col span={0} lg={24}>
                <Form 
                    layout="inline" 
                    initialValues={props.initialValues}
                    onFinish={props.onClickSearch}
                >
                    {content}
                </Form>
            </Col>
        </Row>
    )
}