import { Form, DatePicker, Button, Switch } from "antd"
import moment from "moment"

interface SearchActivitiesProps {
    onChangePeriod(start: moment.Moment, end: moment.Moment): void
    onClickSearch(): void
}

export default function SearchActivities(props: SearchActivitiesProps): JSX.Element {
    return (
        <Form
            layout="inline"
        >
            <Form.Item label="Period" required>
                <DatePicker.RangePicker 
                    format="YYYY-MM-DD HH:mm"
                    showTime={{
                        showSecond: false, 
                        defaultValue: [moment().hour(0).minute(0), moment().hour(23).minute(59)],
                    }}
                    onChange={(_, dateStrings) => props.onChangePeriod(moment(dateStrings[0]), moment(dateStrings[1])) }
                />
            </Form.Item>
            <Form.Item label="Production">
                <Switch size="small" />
            </Form.Item>
            <Form.Item >
                <Button type="primary" onClick={props.onClickSearch}>Search</Button>
            </Form.Item>
        </Form>
    )
}