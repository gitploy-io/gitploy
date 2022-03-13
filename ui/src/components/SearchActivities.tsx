import { Row, Col, Form, DatePicker, Button, Switch } from "antd"
import moment from "moment"

interface SearchActivitiesProps {
    onChangePeriod(start: Date, end: Date): void
    onClickSearch(): void
}

export default function SearchActivities(props: SearchActivitiesProps): JSX.Element {
    const content = (
        <>
            <Form.Item label="Period">
                <DatePicker.RangePicker 
                    format="YYYY-MM-DD HH:mm"
                    showTime={{
                        showSecond: false, 
                        defaultValue: [moment().hour(0).minute(0), moment().hour(23).minute(59)],
                    }}
                    onChange={(_, dateStrings) => props.onChangePeriod(new Date(dateStrings[0]), new Date(dateStrings[1])) }
                />
            </Form.Item>
            <Form.Item label="Production">
                <Switch size="small" />
            </Form.Item>
            <Form.Item >
                <Button 
                    htmlType="submit" 
                    type="primary" 
                    onClick={props.onClickSearch}
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
                <Form layout="horizontal" >
                    {content}
                </Form>
            </Col>
            {/* Laptop */}
            <Col span={0} lg={24}>
                <Form layout="inline" >
                    {content}
                </Form>
            </Col>
        </Row>
    )
}