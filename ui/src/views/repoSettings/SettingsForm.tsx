import { Form, Input, Button, Space, Typography } from "antd"

import { Repo } from "../../models"

export interface SettingFormProps {
    saving: boolean
    repo?: Repo
    onClickFinish(values: any): void
    onClickDeactivate(): void
}

export default function SettingForm({
    saving,
    repo,
    onClickFinish,
    onClickDeactivate,
}: SettingFormProps): JSX.Element {
    const layout = {
      labelCol: { span: 5},
      wrapperCol: { span: 12 },
    };

    const submitLayout = {
      wrapperCol: { offset: 5, span: 12 },
    };

    const initialValues = {
        "config": repo?.configPath
    }

    return (
        <Form
            name="setting"
            initialValues={initialValues}
            onFinish={onClickFinish}
        >
            <Form.Item
                label="Config"
                {...layout}
            >
                <Space>
                    <Form.Item
                        name="config"
                        rules={[{required: true}]}
                        noStyle
                    >
                        <Input />
                    </Form.Item>
                    <Typography.Link target="_blank" href={`/link/${repo?.namespace}/${repo?.name}/config`}>
                        Link
                    </Typography.Link>
                </Space>
            </Form.Item>
            <Form.Item {...submitLayout}>
                <Form.Item noStyle>
                    <Button 
                        loading={saving}
                        type="primary" 
                        htmlType="submit"
                    >
                        Submit
                    </Button>
                </Form.Item>
            </Form.Item>
            <Form.Item {...submitLayout}>
                <Form.Item noStyle>
                    <Space>
                        <Button
                            danger
                            type="primary"
                            onClick={onClickDeactivate}
                        >
                            DEACTIVATE
                        </Button>
                    </Space>
                </Form.Item>
            </Form.Item>
        </Form>
    )
}